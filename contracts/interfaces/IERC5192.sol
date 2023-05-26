// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

interface IERC5192 {
	/// @notice Emitted when the locking status is changed to locked.
	/// @dev If a token is minted and the status is locked, this event should be emitted.
	/// @param tokenId The identifier for a token.
	event Locked(uint256 tokenId);

	/// @notice Emitted when the locking status is changed to unlocked.
	/// @dev If a token is minted and the status is unlocked, this event should be emitted.
	/// @param tokenId The identifier for a token.
	event Unlocked(uint256 tokenId);

	/// @notice Returns the locking status of an Soulbound Token.
	/// @dev SBTs assigned to zero address are considered invalid, and queries  about them do throw.
	/// @param tokenId The identifier for an SBT.
	function locked(uint256 tokenId) external view returns (bool);
}
