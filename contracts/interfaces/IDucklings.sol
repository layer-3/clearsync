// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol';

interface IDucklings is IERC721Upgradeable {
	error TokenNotTransferable(uint256 tokenId);

	struct Duckling {
		uint256 genome;
		uint64 birthdate;
	}

	// events
	event Minted(
		address to,
		uint256 tokenId,
		bool isTransferable,
		uint256 genome,
		uint64 birthdate,
		uint256 chainId
	);

	event TransferableSet(uint256 tokenId, bool isTransferable);

	function isOwnerOf(address account, uint256 tokenIds) external view returns (bool);

	function isOwnerOf(address account, uint256[] calldata tokenIds) external view returns (bool);

	function getGenome(uint256 tokenId) external view returns (uint256);

	function getGenomes(uint256[] calldata tokenIds) external view returns (uint256[] memory);

	function setTransferable(uint256 tokenId, bool isTransferable) external;

	function mintTo(address to, uint256 genome, bool isTransferable) external returns (uint256);

	function mintBatchTo(
		address to,
		uint256[] calldata genomes,
		bool isTransferable
	) external returns (uint256[] memory);

	function burn(uint256 tokenId) external;

	function burnBatch(uint256[] calldata tokenIds) external;
}
