// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721RoyaltyUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol';

import '@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/StringsUpgradeable.sol';

import '../../interfaces/IDucklings.sol';
import '../games/Genome.sol';

/**
 * @title Ducklings
 * @notice This contract implements ERC721 and ERC2981 standards, stores and provides basic functionality for Ducklings NFT.
 * Ducklings expects other Game contracts to define more specific logic for Ducklings NFT.
 *
 * Game contracts should be granted GAME_ROLE to be able to mint and burn tokens.
 * Ducklings defines specific query methods for Game contracts to retrieve specific NFT data.
 *
 * Ducklings can be upgraded by an account with UPGRADER_ROLE to add certain functionality if needed.
 */
contract OldDucklingsV1 is
	Initializable,
	IDucklings,
	ERC721EnumerableUpgradeable,
	ERC721RoyaltyUpgradeable,
	UUPSUpgradeable,
	AccessControlUpgradeable
{
	using CountersUpgradeable for CountersUpgradeable.Counter;
	using {StringsUpgradeable.toString} for uint256;
	using Genome for uint256;

	/**
	 * @notice Is thrown when token with given Id does not exist.
	 * @param tokenId Id of the token.
	 */
	error InvalidTokenId(uint256 tokenId);

	/**
	 * @notice Is thrown when given address is address(0).
	 * @param addr Invalid address.
	 */
	error InvalidAddress(address addr);

	// Roles
	bytes32 public constant UPGRADER_ROLE = keccak256('UPGRADER_ROLE');
	bytes32 public constant GAME_ROLE = keccak256('GAME_ROLE');

	// Royalty
	address internal _royaltiesCollector;
	uint32 internal _royaltyFee;

	// Server address that is prepended to tokenURI
	string public apiBaseURL;

	CountersUpgradeable.Counter public nextNewTokenId;
	mapping(uint256 => Duckling) public tokenToDuckling;

	// ------- Initializer -------

	/**
	 * @notice Initializes the contract.
	 * Grants DEFAULT_ADMIN_ROLE and UPGRADER_ROLE to the deployer.
	 * Sets deployer to be Royalty collector, set royalty fee to 10%.
	 * @dev This function is called only once during contract deployment.
	 */
	function initialize() external initializer {
		__ERC721_init('Yellow Ducklings', 'DUCKLING');
		__ERC721Royalty_init();
		__AccessControl_init();
		__UUPSUpgradeable_init();

		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(UPGRADER_ROLE, msg.sender);

		_royaltiesCollector = msg.sender;
		_royaltyFee = 1000; // 10%
		_setDefaultRoyalty(_royaltiesCollector, _royaltyFee);
	}

	// ------- Upgradable -------

	/**
	 * @notice Upgrades the contract.
	 * @dev Requires UPGRADER_ROLE to invoke.
	 * @param newImplementation Address of the new implementation.
	 */
	function _authorizeUpgrade(
		address newImplementation
	) internal override onlyRole(UPGRADER_ROLE) {}

	// -------- ERC721 --------

	/**
	 * @notice Necessary override to specify what implementation of _burn to use.
	 * @dev Necessary override to specify what implementation of _burn to use.
	 */
	function _burn(uint256 tokenId) internal override(ERC721RoyaltyUpgradeable, ERC721Upgradeable) {
		// check on token existence is performed in ERC721Upgradeable._burn
		super._burn(tokenId);
	}

	/**
	 * @notice Composes an API URL for a given token that returns metadata.json.
	 * @dev Concatenates `apiBaseURL` with token genome, dash (-) and token birthdate.
	 * @param tokenId Id of the token.
	 * @return uri URL for the token metadata.
	 */
	function tokenURI(
		uint256 tokenId
	) public view override(ERC721Upgradeable) returns (string memory) {
		if (!_exists(tokenId)) revert InvalidTokenId(tokenId);

		Duckling memory duckling = tokenToDuckling[tokenId];
		string memory genome = duckling.genome.toString();
		string memory birthdate = uint256(duckling.birthdate).toString();
		string memory uri = string.concat(apiBaseURL, genome, '-', birthdate);

		return uri;
	}

	/**
	 * @notice Checks whether supplied `interface` is supported by the contract.
	 * @dev Checks whether supplied `interface` is supported by the contract.
	 * @param interfaceId Id of the interface.
	 * @return interfaceSupported true if interface is supported, false otherwise.
	 */
	function supportsInterface(
		bytes4 interfaceId
	)
		public
		view
		virtual
		override(
			IERC165Upgradeable,
			ERC721EnumerableUpgradeable,
			ERC721RoyaltyUpgradeable,
			AccessControlUpgradeable
		)
		returns (bool)
	{
		return super.supportsInterface(interfaceId);
	}

	// -------- ERC2981 Royalties --------

	/**
	 * @notice Sets royalties collector.
	 * @dev Requires DEFAULT_ADMIN_ROLE to invoke.
	 * @param account Address of the royalties collector.
	 */
	function setRoyaltyCollector(address account) public onlyRole(DEFAULT_ADMIN_ROLE) {
		_royaltiesCollector = account;
		_setDefaultRoyalty(account, _royaltyFee);
	}

	/**
	 * @notice Returns royalties collector.
	 * @dev Returns royalties collector.
	 * @return address Address of the royalties collector.
	 */
	function getRoyaltyCollector() public view returns (address) {
		return _royaltiesCollector;
	}

	/**
	 * @notice Sets royalties fee.
	 * @dev Requires DEFAULT_ADMIN_ROLE to invoke.
	 * @param fee Royalties fee in permyriad.
	 */
	function setRoyaltyFee(uint32 fee) public onlyRole(DEFAULT_ADMIN_ROLE) {
		_royaltyFee = fee;
		_setDefaultRoyalty(_royaltiesCollector, fee);
	}

	/**
	 * @notice Returns royalties fee.
	 * @dev Returns royalties fee.
	 * @return uint32 Royalties fee in permyriad.
	 */
	function getRoyaltyFee() public view returns (uint32) {
		return _royaltyFee;
	}

	// -------- API URL --------

	/**
	 * @notice Sets api server endpoint that is prepended to the tokenURI.
	 * @dev Requires DEFAULT_ADMIN_ROLE to invoke.
	 * @param apiBaseURL_ URL of the api server.
	 */
	function setAPIBaseURL(string calldata apiBaseURL_) external onlyRole(DEFAULT_ADMIN_ROLE) {
		apiBaseURL = apiBaseURL_;
	}

	// -------- IDucklings --------

	/**
	 * @notice Check whether `account` is owner of `tokenId`.
	 * @dev Revert if `account` is address(0) or `tokenId` does not exist.
	 * @param account Address to check.
	 * @param tokenId Token Id to check.
	 * @return isOwnerOf True if `account` is owner of `tokenId`, false otherwise.
	 */
	function isOwnerOf(address account, uint256 tokenId) external view returns (bool) {
		if (account == address(0)) revert InvalidAddress(account);
		if (!_exists(tokenId)) revert InvalidTokenId(tokenId);

		return account == ownerOf(tokenId);
	}

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
	) external view returns (bool) {
		if (account == address(0)) revert InvalidAddress(account);

		// first check if all tokens exist
		for (uint256 i = 0; i < tokenIds.length; i++) {
			if (!_exists(tokenIds[i])) revert InvalidTokenId(tokenIds[i]);
		}

		// then check if all tokens belong to the account
		for (uint256 i = 0; i < tokenIds.length; i++) {
			if (account != ownerOf(tokenIds[i])) {
				return false;
			}
		}

		return true;
	}

	/**
	 * @notice Get genome of `tokenId`.
	 * @dev Revert if `tokenId` does not exist.
	 * @param tokenId Token Id to get the genome of.
	 * @return genome Genome of `tokenId`.
	 */
	function getGenome(uint256 tokenId) external view returns (uint256) {
		if (!_exists(tokenId)) revert InvalidTokenId(tokenId);
		return tokenToDuckling[tokenId].genome;
	}

	/**
	 * @notice Get genomes of `tokenIds`.
	 * @dev Revert if any of `tokenIds` do not exist.
	 * @param tokenIds Token Ids to get the genomes of.
	 * @return genomes Genomes of `tokenIds`.
	 */
	function getGenomes(uint256[] calldata tokenIds) external view returns (uint256[] memory) {
		// explicitly specify array length
		uint256[] memory genomes = new uint256[](tokenIds.length);

		for (uint256 i = 0; i < tokenIds.length; i++) {
			if (!_exists(tokenIds[i])) revert InvalidTokenId(tokenIds[i]);
			genomes[i] = tokenToDuckling[tokenIds[i]].genome;
		}

		return genomes;
	}

	/**
	 * @notice Check whether token with `tokenId` is transferable.
	 * @dev Revert if `tokenId` does not exist.
	 * @param tokenId Token Id to check.
	 * @return isTransferable True if token with `tokenId` is transferable, false otherwise.
	 */
	function isTransferable(uint256 tokenId) external view returns (bool) {
		return _isTransferable(tokenId);
	}

	/**
	 * @notice Check whether token with `tokenId` is transferable. Internal function.
	 * @dev Revert if `tokenId` does not exist.
	 * @param tokenId Token Id to check.
	 * @return isTransferable True if token with `tokenId` is transferable, false otherwise.
	 */
	function _isTransferable(uint256 tokenId) internal view returns (bool) {
		if (!_exists(tokenId)) revert InvalidTokenId(tokenId);

		return tokenToDuckling[tokenId].genome.getFlag(Genome.FLAG_TRANSFERABLE);
	}

	/**
	 * @notice Check whether token that is being transferred is transferable. Revert if not.
	 * @dev Revert if token is not transferable.
	 * @param from Address of the sender.
	 * @param to Address of the recipient.
	 * @param firstTokenId Id of the token being transferred.
	 */
	function _beforeTokenTransfer(
		address from,
		address to,
		uint256 firstTokenId,
		uint256 batchSize
	) internal override(ERC721Upgradeable, ERC721EnumerableUpgradeable) {
		super._beforeTokenTransfer(from, to, firstTokenId, batchSize);

		// mint and burn for not transferable is allowed
		if (from == address(0) || to == address(0)) return;

		if (!_isTransferable(firstTokenId)) revert TokenNotTransferable(firstTokenId);
	}

	/**
	 * @notice Mint token with `genome` to `to`. Emits Minted event.
	 * @dev Revert if `to` is address(0) or `genome` has wrong magic number.
	 * @param to Address to mint token to.
	 * @param genome Genome of the token to mint.
	 * @return tokenId Id of the minted token.
	 */
	function mintTo(
		address to,
		uint256 genome
	) external onlyRole(GAME_ROLE) returns (uint256 tokenId) {
		return _mintTo(to, genome);
	}

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
	) external onlyRole(GAME_ROLE) returns (uint256[] memory tokenIds) {
		tokenIds = new uint256[](genomes.length);

		for (uint8 i = 0; i < genomes.length; i++) {
			tokenIds[i] = _mintTo(to, genomes[i]);
		}
	}

	/**
	 * @notice Mint token with `genome` to `to`. Emits Minted event. Internal function.
	 * @dev Revert if `to` is address(0) or `genome` has wrong magic number.
	 * @param to Address to mint token to.
	 * @param genome Genome of the token to mint.
	 * @return tokenId Id of the minted token.
	 */
	function _mintTo(address to, uint256 genome) internal returns (uint256 tokenId) {
		if (to == address(0)) revert InvalidAddress(to);

		uint8 magicNum = genome.getGene(Genome.MAGIC_NUMBER_GENE_IDX);

		if (magicNum != Genome.BASE_MAGIC_NUMBER && magicNum != Genome.MYTHIC_MAGIC_NUMBER)
			revert InvalidMagicNumber(magicNum);

		tokenId = nextNewTokenId.current();
		uint64 birthdate = uint64(block.timestamp);
		tokenToDuckling[tokenId] = Duckling(genome, birthdate);
		_safeMint(to, tokenId);
		nextNewTokenId.increment();

		emit Minted(to, tokenId, genome, birthdate, block.chainid);
	}

	/**
	 * @notice Burn token with `tokenId`.
	 * @dev Revert if `tokenId` does not exist.
	 * @param tokenId Id of the token to burn.
	 */
	function burn(uint256 tokenId) external onlyRole(GAME_ROLE) {
		_burn(tokenId);
	}

	/**
	 * @notice Burn tokens with `tokenIds`.
	 * @dev Revert if any of `tokenIds` do not exist.
	 * @param tokenIds Ids of the tokens to burn.
	 */
	function burnBatch(uint256[] calldata tokenIds) external onlyRole(GAME_ROLE) {
		for (uint256 i = 0; i < tokenIds.length; i++) _burn(tokenIds[i]);
	}
}
