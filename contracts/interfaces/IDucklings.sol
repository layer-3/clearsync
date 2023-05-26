// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol';

/**
 * @title IDucklings
 * @notice This interface defines the Ducklings ERC721-compatible contract,
 * which provides basic functionality for minting, burning and querying information about the tokens.
 */
interface IDucklings is IERC721Upgradeable {
	/**
	 * @notice Incorrect transferability error. Is thrown when expected transferability of `tokenId` to be `expected` but is the opposite.
	 * @param tokenId Token Id that has incorrect transferability.
	 * @param expected Expected transferability of `tokenId`.
	 */
	error IncorrectTransferability(uint256 tokenId, bool expected);

	/**
	 * @notice Invalid magic number error. Is used when trying to mint a token with an invalid magic number.
	 * @param magicNumber Magic number that is invalid.
	 */
	error InvalidMagicNumber(uint8 magicNumber);

	struct Duckling {
		uint256 genome;
		uint64 birthdate;
	}

	// events
	/**
	 * @notice Minted event. Is emitted when a token is minted.
	 * @param to Address of the token owner.
	 * @param tokenId Id of the minted token.
	 * @param genome Genome of the minted token.
	 * @param birthdate Birthdate of the minted token.
	 * @param chainId Id of the chain where the token was minted.
	 */
	event Minted(address to, uint256 tokenId, uint256 genome, uint64 birthdate, uint256 chainId);

	/**
	 * @notice Check whether `account` is owner of `tokenId`.
	 * @dev Revert if `account` is address(0) or `tokenId` does not exist.
	 * @param account Address to check.
	 * @param tokenId Token Id to check.
	 * @return isOwnerOf True if `account` is owner of `tokenId`, false otherwise.
	 */
	function isOwnerOf(address account, uint256 tokenId) external view returns (bool);

	/**
	 * @notice Check whether `account` is owner of `tokenIds`.
	 * @dev Revert if `account` is address(0) or any of `tokenIds` do not exist.
	 * @param account Address to check.
	 * @param tokenIds Token Ids to check.
	 * @return isOwnerOfBatch True if `account` is owner of `tokenIds`, false otherwise.
	 */
	function isOwnerOfBatch(
		address account,
		uint256[] calldata tokenIds
	) external view returns (bool);

	/**
	 * @notice Get genome of `tokenId`.
	 * @dev Revert if `tokenId` does not exist.
	 * @param tokenId Token Id to get the genome of.
	 * @return genome Genome of `tokenId`.
	 */
	function getGenome(uint256 tokenId) external view returns (uint256);

	/**
	 * @notice Get genomes of `tokenIds`.
	 * @dev Revert if any of `tokenIds` do not exist.
	 * @param tokenIds Token Ids to get the genomes of.
	 * @return genomes Genomes of `tokenIds`.
	 */
	function getGenomes(uint256[] calldata tokenIds) external view returns (uint256[] memory);

	/**
	 * @notice Mint token with `genome` to `to`. Emits Minted event.
	 * @dev Revert if `to` is address(0) or `genome` has wrong magic number.
	 * @param to Address to mint token to.
	 * @param genome Genome of the token to mint.
	 * @return tokenId Id of the minted token.
	 */
	function mintTo(address to, uint256 genome) external returns (uint256);

	/**
	 * @notice Mint tokens with `genomes` to `to`. Emits Minted event for each token.
	 * @dev Revert if `to` is address(0) or any of `genomes` has wrong magic number.
	 * @param to Address to mint tokens to.
	 * @param genomes Genomes of the tokens to mint.
	 * @return tokenIds Ids of the minted tokens.
	 */
	function mintBatchTo(
		address to,
		uint256[] calldata genomes
	) external returns (uint256[] memory);

	/**
	 * @notice Burn token with `tokenId`.
	 * @dev Revert if `tokenId` does not exist.
	 * @param tokenId Id of the token to burn.
	 */
	function burn(uint256 tokenId) external;

	/**
	 * @notice Burn tokens with `tokenIds`.
	 * @dev Revert if any of `tokenIds` do not exist.
	 * @param tokenIds Ids of the tokens to burn.
	 */
	function burnBatch(uint256[] calldata tokenIds) external;
}
