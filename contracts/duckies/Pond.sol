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

// TODO: add locking Duckies / LP tokens alongside Mythic
contract Pond is IPond, Ownable, ReentrancyGuard {
	using Genome for uint256;

	IDucklings public immutable ducklings;
	IERC20 public immutable rewardToken;

	uint64 public lockupPeriodSeconds;
	uint256 public powerPerMythic;

	// Accrued reward token per locked mythic rank
	uint256 public accRewardPerShare;
	uint256 public rewardStartBlock;
	uint256 public rewardEndBlock;
	uint256 public lastClaimedAtBlock;
	uint256 public rewardPerBlock;
	uint256 public constant PRECISION_FACTOR = 10 ** 16;

	struct UserInfo {
		uint256[] lockedMythics; // Locked Mythic token IDs
		uint256 cumulativeLockedRank; // Cumulative rank of locked Mythic
		uint256 rewardDebt; // Reward debt
		mapping(uint256 => uint64) unlockableAt; // Un-lockable at timestamp
	}

	mapping(address => UserInfo) public userInfoOf;

	uint256 public cumulativeLockedRank;

	constructor(
		IDucklings _ducklings,
		IERC20 _rewardToken,
		uint256 _powerPerMythic,
		uint256 _rewardStartBlock,
		uint256 _rewardEndBlock,
		uint256 _rewardPerBlock
	) {
		ducklings = _ducklings;
		rewardToken = _rewardToken;
		powerPerMythic = _powerPerMythic;

		rewardStartBlock = _rewardStartBlock;
		rewardEndBlock = _rewardEndBlock;
		lastClaimedAtBlock = _rewardStartBlock;
		rewardPerBlock = _rewardPerBlock;
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

	function pendingYield(address user) external view override returns (uint256) {
		if (block.number < rewardStartBlock) {
			return 0;
		}

		uint256 blocks = _getBlocksBetween(lastClaimedAtBlock, block.number);
		uint256 reward = blocks * rewardPerBlock;
		uint256 adjustedTokenPerShare = accRewardPerShare +
			(reward * PRECISION_FACTOR) /
			cumulativeLockedRank;

		return
			(userInfoOf[user].cumulativeLockedRank * adjustedTokenPerShare) /
			PRECISION_FACTOR -
			userInfoOf[user].rewardDebt;
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

		_updateYieldParams();
		_claimYield(user);

		ducklings.transferFrom(user, address(this), tokenId);

		userInfo.unlockableAt[tokenId] = uint64(block.timestamp) + lockupPeriodSeconds;
		userInfo.lockedMythics.push(tokenId);

		uint256 mythicRank = genome.getGene(uint8(IDuckyFamily.MythicGenes.UniqId));
		userInfo.cumulativeLockedRank += mythicRank;
		cumulativeLockedRank += mythicRank;

		_postAccountChangeUpdate(user);

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

		_updateYieldParams();
		_claimYield(user);

		ducklings.transferFrom(address(this), user, tokenId);

		userInfo.unlockableAt[tokenId] = 0;
		_removeLockedMythicOf(user, tokenId);

		uint256 genome = ducklings.getGenome(tokenId);
		uint256 mythicRank = genome.getGene(uint8(IDuckyFamily.MythicGenes.UniqId));
		userInfo.cumulativeLockedRank -= mythicRank;
		cumulativeLockedRank -= mythicRank;

		_postAccountChangeUpdate(user);

		emit MythicUnlocked(user, tokenId);
	}

	// -------- Yield --------

	function claimYield() external override nonReentrant {
		_updateYieldParams();
		_claimYield(msg.sender);
		_postAccountChangeUpdate(msg.sender);
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

	function _claimYield(address user) internal {
		if (
			block.number < rewardStartBlock ||
			rewardEndBlock < block.number ||
			cumulativeLockedRank == 0
		) {
			return;
		}

		UserInfo storage userInfo = userInfoOf[user];
		uint256 yield = (userInfo.cumulativeLockedRank * accRewardPerShare) /
			PRECISION_FACTOR -
			userInfo.rewardDebt;

		if (yield > 0) {
			rewardToken.transfer(user, yield);
			emit YieldClaimed(user, yield);
		}

		lastClaimedAtBlock = block.number;
	}

	function _updateYieldParams() internal {
		if (block.number <= lastClaimedAtBlock) {
			return;
		}

		if (cumulativeLockedRank == 0) {
			lastClaimedAtBlock = block.number;
			return;
		}

		uint256 blocks = _getBlocksBetween(lastClaimedAtBlock, block.number);
		uint256 reward = blocks * rewardPerBlock;
		accRewardPerShare += (reward * PRECISION_FACTOR) / cumulativeLockedRank;
	}

	function _postAccountChangeUpdate(address user) internal {
		UserInfo storage userInfo = userInfoOf[user];
		userInfo.rewardDebt =
			(userInfo.cumulativeLockedRank * accRewardPerShare) /
			PRECISION_FACTOR;
	}

	function _getBlocksBetween(uint256 _from, uint256 _to) internal view returns (uint256) {
		uint256 _rewardEndBlock = rewardEndBlock;

		if (_to <= _rewardEndBlock) {
			return _to - _from;
		} else if (_from >= _rewardEndBlock) {
			return 0;
		} else {
			return _rewardEndBlock - _from;
		}
	}
}
