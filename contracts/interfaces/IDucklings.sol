// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol';

interface IDucklings is IERC721Upgradeable {
	error TokenNotTransferable(uint256 tokenId);
	error InvalidMagicNumber(uint8 magicNumber);

	struct Duckling {
		uint256 genome;
		uint64 birthdate;
	}

	// events
	event Minted(address to, uint256 tokenId, uint256 genome, uint64 birthdate, uint256 chainId);

	function isOwnerOf(address account, uint256 tokenIds) external view returns (bool);

	function isOwnerOfBatch(
		address account,
		uint256[] calldata tokenIds
	) external view returns (bool);

	function getGenome(uint256 tokenId) external view returns (uint256);

	function getGenomes(uint256[] calldata tokenIds) external view returns (uint256[] memory);

	function mintTo(address to, uint256 genome) external returns (uint256);

	function mintBatchTo(
		address to,
		uint256[] calldata genomes
	) external returns (uint256[] memory);

	function burn(uint256 tokenId) external;

	function burnBatch(uint256[] calldata tokenIds) external;
}
