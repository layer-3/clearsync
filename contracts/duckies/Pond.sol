// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts/access/Ownable.sol';

import '../interfaces/IPond.sol';
import '../interfaces/IERC20MintableBurnable.sol';
import '../interfaces/IDucklings.sol';
import './games/DuckyFamily/DuckyGenome.sol';
import './games/Genome.sol';

contract Pond is IPond, Ownable {
	using Genome for uint256;

	IDucklings public immutable ducklings;

	uint64 public lockupSeconds;
	uint256 public powerPerMythic;

	mapping(address => uint256[]) public lockedMythicsOf;
	mapping(address => mapping(uint256 => uint64)) public unlockableAt;

	constructor(IDucklings _ducklings, uint256 _powerPerMythic) {
		ducklings = _ducklings;
		powerPerMythic = _powerPerMythic;
	}

	// -------- Setters --------

	function setLockupSeconds(uint64 _lockupSeconds) external onlyOwner {
		lockupSeconds = _lockupSeconds;
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
		return lockedMythicsOf[account].length * powerPerMythic;
	}

	// -------- Lock --------

	function lockMythic(uint256 tokenId) external override {
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

		ducklings.transferFrom(account, address(this), tokenId);

		uint64 _unlockableAt = uint64(block.timestamp) + lockupSeconds;
		unlockableAt[account][tokenId] = _unlockableAt;
		lockedMythicsOf[account].push(tokenId);

		emit MythicLocked(account, tokenId);
	}

	function unlockMythic(uint256 tokenId) external override {
		address account = msg.sender;

		uint64 _unlockableAt = unlockableAt[account][tokenId];

		if (_unlockableAt == 0) {
			revert NotLocked(tokenId);
		}

		if (_unlockableAt > block.timestamp) {
			revert NotUnlockable(tokenId);
		}

		ducklings.transferFrom(address(this), account, tokenId);

		unlockableAt[account][tokenId] = 0;
		_removeLockedMythicOf(account, tokenId);

		emit MythicUnlocked(account, tokenId);
	}

	// -------- Yield --------

	function claimYield() external override {}

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
}
