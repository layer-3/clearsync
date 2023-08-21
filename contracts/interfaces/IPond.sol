// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

interface IPond {
	// errors

	// events
	event LockupSecondsSet(uint64 lockupSeconds);
	event PowerPerMythicSet(uint256 powerPerMythic);
	event MythicLocked(address indexed account, uint256 indexed tokenId);
	event MythicUnlocked(address indexed account, uint256 indexed tokenId);

	// view
	function lockedMythicsOf(address account) external view returns (uint256[] memory);

	function isMythicLocked(uint256 tokenId) external view returns (bool);

	function votingPowerOf(address account) external view returns (uint256);

	// lock
	function lockMythic(uint256 tokenId) external;

	function unlockMythic(uint256 tokenId) external;

	// yield
	function claimYield() external;
}
