// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts/access/Ownable.sol';
import '@openzeppelin/contracts/security/ReentrancyGuard.sol';
import '@openzeppelin/contracts/token/ERC20/IERC20.sol';

import '../interfaces/IPond.sol';
import '../interfaces/IERC20MintableBurnable.sol';
import '../interfaces/IDucklings.sol';
import './games/DuckyFamily/DuckyGenome.sol';
import './games/Genome.sol';
import '../interfaces/IDuckyFamily.sol';

/**
 * Terminology:
 *
 * Pool - locked Mythic tokens and a strategy of paying yield for them.
 * Yield - reward for locking Mythic Ducklings in the Pool.
 * (User) Share - the ratio of total rank of User-locked Mythic Ducklings to the total rank of all locked Mythic Ducklings.
 */

// TODO: add locking Duckies / LP tokens alongside Mythic
contract Pond is IPond, Ownable, ReentrancyGuard {
	using Genome for uint256;

	IDucklings public immutable ducklings;
	IERC20 public immutable yieldToken;

	uint64 public lockupPeriodSeconds;
	uint256 public powerPerMythic;

	// Accrued reward token per locked mythic rank
	uint256 public accrYieldPerShare;
	uint256 public yieldStartBlock;
	uint256 public yieldEndBlock;
	uint256 public latestYieldBlock;
	uint256 public yieldPerBlock;
	uint256 public constant PRECISION_FACTOR = 10 ** 16;

	struct UserInfo {
		uint256[] lockedMythics; // Locked Mythic token IDs
		uint256 totalLockedRank; // Total rank of locked Mythic
		mapping(uint256 => uint64) unlockableAt; // Un-lockable at timestamp
		uint256 notApplicableYield; // Yield that was already paid or is not to be paid to the User (yield calculation specificity)
	}

	mapping(address => UserInfo) public userInfoOf;

	uint256 public totalLockedRank;

	constructor(
		IDucklings _ducklings,
		IERC20 _yieldToken,
		uint256 _powerPerMythic,
		uint256 _yieldStartBlock,
		uint256 _yieldEndBlock,
		uint256 _yieldPerBlock
	) {
		ducklings = _ducklings;
		yieldToken = _yieldToken;
		powerPerMythic = _powerPerMythic;

		yieldStartBlock = _yieldStartBlock;
		yieldEndBlock = _yieldEndBlock;
		latestYieldBlock = _yieldStartBlock;
		yieldPerBlock = _yieldPerBlock;
	}

	// -------- Setters --------

	function setLockupPeriodSeconds(uint64 _lockupPeriodSeconds) external onlyOwner {
		lockupPeriodSeconds = _lockupPeriodSeconds;
	}

	function setPowerPerMythic(uint256 _powerPerMythic) external onlyOwner {
		powerPerMythic = _powerPerMythic;
	}

	// -------- View --------

	function getLockedMythicsOf(address user) external view override returns (uint256[] memory) {
		return userInfoOf[user].lockedMythics;
	}

	function isMythicLocked(uint256 tokenId) external view override returns (bool) {
		return userInfoOf[msg.sender].unlockableAt[tokenId] != 0;
	}

	function getVotingPowerOf(address user) external view override returns (uint256) {
		// TODO: discuss should the power depend on the rank
		return userInfoOf[user].lockedMythics.length * powerPerMythic;
	}

	function getPendingYield(address user) external view override returns (uint256) {
		if (block.number < yieldStartBlock) {
			return 0;
		}

		uint256 blocks = _getBlocksBetween(latestYieldBlock, block.number);
		uint256 reward = blocks * yieldPerBlock;
		uint256 adjustedTokenPerShare = accrYieldPerShare +
			(reward * PRECISION_FACTOR) /
			totalLockedRank;

		return
			(userInfoOf[user].totalLockedRank * adjustedTokenPerShare) /
			PRECISION_FACTOR -
			userInfoOf[user].notApplicableYield;
	}

	// -------- Lock --------

	function lockMythic(uint256 tokenId) external override nonReentrant {
		address user = msg.sender;
		UserInfo storage userInfo = userInfoOf[user];

		if (ducklings.ownerOf(tokenId) != user) {
			revert CallerNotOwner(tokenId);
		}

		if (userInfoOf[user].unlockableAt[tokenId] != 0) {
			revert AlreadyLocked(tokenId);
		}

		uint256 genome = ducklings.getGenome(tokenId);
		if (genome.getGene(Genome.COLLECTION_GENE_IDX) != DuckyGenome.mythicCollectionId) {
			revert InvalidCollection(
				DuckyGenome.mythicCollectionId,
				genome.getGene(Genome.COLLECTION_GENE_IDX)
			);
		}

		_updatePool();
		_claimYield(user);

		ducklings.transferFrom(user, address(this), tokenId);

		userInfo.unlockableAt[tokenId] = uint64(block.timestamp) + lockupPeriodSeconds;
		userInfo.lockedMythics.push(tokenId);

		uint256 mythicRank = genome.getGene(uint8(IDuckyFamily.MythicGenes.UniqId));
		userInfo.totalLockedRank += mythicRank;
		totalLockedRank += mythicRank;

		_applyYieldAndShareChange(user);

		emit MythicLocked(user, tokenId);
	}

	function unlockMythic(uint256 tokenId) external override nonReentrant {
		address user = msg.sender;
		UserInfo storage userInfo = userInfoOf[user];

		uint64 _unlockableAt = userInfo.unlockableAt[tokenId];

		if (_unlockableAt == 0) {
			revert NotLocked(tokenId);
		}

		if (_unlockableAt > block.timestamp) {
			revert NotUnlockable(tokenId);
		}

		_updatePool();
		_claimYield(user);

		ducklings.transferFrom(address(this), user, tokenId);

		userInfo.unlockableAt[tokenId] = 0;
		_removeLockedMythicOf(user, tokenId);

		uint256 genome = ducklings.getGenome(tokenId);
		uint256 mythicRank = genome.getGene(uint8(IDuckyFamily.MythicGenes.UniqId));
		userInfo.totalLockedRank -= mythicRank;
		totalLockedRank -= mythicRank;

		_applyYieldAndShareChange(user);

		emit MythicUnlocked(user, tokenId);
	}

	// -------- Yield --------

	function claimYield() external override nonReentrant {
		_updatePool();
		_claimYield(msg.sender);
		_applyYieldAndShareChange(msg.sender);
	}

	// -------- Internal --------

	function _removeLockedMythicOf(address user, uint256 tokenId) internal {
		uint256[] storage tokenIds = userInfoOf[user].lockedMythics;
		uint256 length = tokenIds.length;
		for (uint256 i = 0; i < length; i++) {
			if (tokenIds[i] == tokenId) {
				tokenIds[i] = tokenIds[length - 1];
				tokenIds.pop();
				break;
			}
		}
	}

	function _updatePool() internal {
		if (block.number <= latestYieldBlock) {
			return;
		}

		if (totalLockedRank == 0) {
			latestYieldBlock = block.number;
			return;
		}

		uint256 blocks = _getBlocksBetween(latestYieldBlock, block.number);
		uint256 reward = blocks * yieldPerBlock;
		accrYieldPerShare += (reward * PRECISION_FACTOR) / totalLockedRank;
	}

	function _claimYield(address user) internal {
		if (
			block.number < yieldStartBlock || yieldEndBlock < block.number || totalLockedRank == 0
		) {
			return;
		}

		UserInfo storage userInfo = userInfoOf[user];
		uint256 yield = (userInfo.totalLockedRank * accrYieldPerShare) /
			PRECISION_FACTOR -
			userInfo.notApplicableYield;

		if (yield > 0) {
			yieldToken.transfer(user, yield);
			emit YieldClaimed(user, yield);
		}

		latestYieldBlock = block.number;
	}

	function _applyYieldAndShareChange(address user) internal {
		UserInfo storage userInfo = userInfoOf[user];
		userInfo.notApplicableYield =
			(userInfo.totalLockedRank * accrYieldPerShare) /
			PRECISION_FACTOR;
	}

	function _getBlocksBetween(uint256 _from, uint256 _to) internal view returns (uint256) {
		uint256 _yieldEndBlock = yieldEndBlock;

		if (_to <= _yieldEndBlock) {
			return _to - _from;
		} else if (_from >= _yieldEndBlock) {
			return 0;
		} else {
			return _yieldEndBlock - _from;
		}
	}
}
