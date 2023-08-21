// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

interface IPond {
	// -------- Errors --------
	error CallerNotOwner(uint256 tokenId);
	error AlreadyLocked(uint256 tokenId);
	error NotLocked(uint256 tokenId);
	error NotUnlockable(uint256 tokenId);
	error InvalidCollection(uint8 expected, uint8 actual);

	// -------- Events --------

	event LockupSecondsSet(uint64 lockupSeconds);
	event PowerPerMythicSet(uint256 powerPerMythic);
	event MythicLocked(address indexed account, uint256 indexed tokenId);
	event MythicUnlocked(address indexed account, uint256 indexed tokenId);

	// -------- View --------
	function getLockedMythicsOf(address account) external view returns (uint256[] memory);

	function isMythicLocked(uint256 tokenId) external view returns (bool);

	function getVotingPowerOf(address account) external view returns (uint256);

	// -------- Lock --------

	function lockMythic(uint256 tokenId) external;

	function unlockMythic(uint256 tokenId) external;

	// -------- Yield --------

	function claimYield() external;
}
