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

// TODO: add locking Duckies alongside Mythic
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

	// TODO: merge these mappings into a struct
	mapping(address => uint256[]) public lockedMythicsOf;
	mapping(address => uint256) public cumulativeLockedRankOf;
	mapping(address => uint256) public rewardPaidTo;
	mapping(address => mapping(uint256 => uint64)) public unlockableAt;

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

	function getLockedMythicsOf(address account) external view override returns (uint256[] memory) {
		return lockedMythicsOf[account];
	}

	function isMythicLocked(uint256 tokenId) external view override returns (bool) {
		return unlockableAt[msg.sender][tokenId] != 0;
	}

	function getVotingPowerOf(address account) external view override returns (uint256) {
		// TODO: discuss should the power depend on the rank
		return lockedMythicsOf[account].length * powerPerMythic;
	}

	function pendingYield(address account) external view override returns (uint256) {
		if (block.number < rewardStartBlock) {
			return 0;
		}

		uint256 blocks = _getBlocksBetween(lastClaimedAtBlock, block.number);
		uint256 reward = blocks * rewardPerBlock;
		uint256 adjustedTokenPerShare = accRewardPerShare +
			(reward * PRECISION_FACTOR) /
			cumulativeLockedRank;

		return
			(cumulativeLockedRankOf[account] * adjustedTokenPerShare) /
			PRECISION_FACTOR -
			rewardPaidTo[account];
	}

	// -------- Lock --------

	function lockMythic(uint256 tokenId) external override nonReentrant {
		address account = msg.sender;

		if (ducklings.ownerOf(tokenId) != account) {
			revert CallerNotOwner(tokenId);
		}

		if (unlockableAt[account][tokenId] != 0) {
			revert AlreadyLocked(tokenId);
		}

		uint256 genome = ducklings.getGenome(tokenId);
		if (genome.getGene(Genome.COLLECTION_GENE_IDX) != DuckyGenome.mythicCollectionId) {
			revert InvalidCollection(
				DuckyGenome.mythicCollectionId,
				genome.getGene(Genome.COLLECTION_GENE_IDX)
			);
		}

		_claimYield(account);

		ducklings.transferFrom(account, address(this), tokenId);

		uint64 _unlockableAt = uint64(block.timestamp) + lockupPeriodSeconds;
		unlockableAt[account][tokenId] = _unlockableAt;
		lockedMythicsOf[account].push(tokenId);

		uint256 mythicRank = genome.getGene(uint8(IDuckyFamily.MythicGenes.UniqId));
		cumulativeLockedRankOf[account] += mythicRank;
		cumulativeLockedRank += mythicRank;

		_updateAccRewardPerShare();

		emit MythicLocked(account, tokenId);
	}

	function unlockMythic(uint256 tokenId) external override nonReentrant {
		address account = msg.sender;

		uint64 _unlockableAt = unlockableAt[account][tokenId];

		if (_unlockableAt == 0) {
			revert NotLocked(tokenId);
		}

		if (_unlockableAt > block.timestamp) {
			revert NotUnlockable(tokenId);
		}

		_claimYield(account);

		ducklings.transferFrom(address(this), account, tokenId);

		unlockableAt[account][tokenId] = 0;
		_removeLockedMythicOf(account, tokenId);

		uint256 genome = ducklings.getGenome(tokenId);
		uint256 mythicRank = genome.getGene(uint8(IDuckyFamily.MythicGenes.UniqId));
		cumulativeLockedRankOf[account] -= mythicRank;
		cumulativeLockedRank -= mythicRank;

		_updateAccRewardPerShare();

		emit MythicUnlocked(account, tokenId);
	}

	// -------- Yield --------

	/// @dev Does not affect cumulative locked rank, thus `accRewardPerShare` update is unnecessary.
	function claimYield() external override nonReentrant {
		_claimYield(msg.sender);
	}

	// -------- Internal --------

	function _removeLockedMythicOf(address account, uint256 tokenId) internal {
		uint256[] storage tokenIds = lockedMythicsOf[account];
		uint256 length = tokenIds.length;
		for (uint256 i = 0; i < length; i++) {
			if (tokenIds[i] == tokenId) {
				tokenIds[i] = tokenIds[length - 1];
				tokenIds.pop();
				break;
			}
		}
	}

	function _claimYield(address account) internal {
		if (
			block.number < rewardStartBlock ||
			rewardEndBlock < block.number ||
			cumulativeLockedRank == 0
		) {
			return;
		}

		uint256 yield = cumulativeLockedRankOf[account] * accRewardPerShare - rewardPaidTo[account];

		if (yield > 0) {
			rewardToken.transfer(account, yield);
			rewardPaidTo[account] += yield;
			emit YieldClaimed(account, yield);
		}

		lastClaimedAtBlock = block.number;
	}

	function _updateAccRewardPerShare() internal {
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
