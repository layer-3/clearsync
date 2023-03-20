// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

interface IDucklings {
	// events
	event Minted(address to, uint256 tokenId, uint256 genome, uint256 chainId);

	function isOwnerOf(address account, uint256 tokenIds) external view returns (bool);

	function isOwnerOf(address account, uint256[] calldata tokenIds) external view returns (bool);

	function getGenome(uint256 tokenId) external view returns (uint256);

	function getGenomes(uint256[] calldata tokenIds) external view returns (uint256[] memory);

	function mintTo(address to, uint256 genome) external returns (uint256);

	function burn(uint256 tokenId) external;

	function burn(uint256[] calldata tokenIds) external;
}
