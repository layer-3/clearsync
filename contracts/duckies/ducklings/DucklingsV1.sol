// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721RoyaltyUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol';

import '@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/StringsUpgradeable.sol';

import '../../interfaces/IDucklings.sol';

contract DucklingsV1 is
	Initializable,
	IDucklings,
	ERC721Upgradeable,
	ERC721RoyaltyUpgradeable,
	UUPSUpgradeable,
	AccessControlUpgradeable
{
	using CountersUpgradeable for CountersUpgradeable.Counter;
	using {StringsUpgradeable.toString} for uint256;

	error InvalidTokenId(uint256 tokenId);

	// Constants
	uint8 public constant FLAG_TRANSFERABLE = 0;

	// Roles
	bytes32 public constant UPGRADER_ROLE = keccak256('UPGRADER_ROLE');
	bytes32 public constant GAME_ROLE = keccak256('GAME_ROLE');

	// Royalty
	address private _royaltiesCollector;
	uint32 private constant ROYALTY_FEE = 1000; // 10%

	string public apiBaseURL;

	CountersUpgradeable.Counter public nextNewTokenId;
	mapping(uint256 => Duckling) public tokenToDuckling;

	// ------- Initializer -------

	function initialize() external initializer {
		__ERC721_init('Yellow Ducklings', 'DUCKLING');
		__ERC721Royalty_init();
		__AccessControl_init();
		__UUPSUpgradeable_init();

		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(UPGRADER_ROLE, msg.sender);

		setRoyaltyCollector(msg.sender);
		_setDefaultRoyalty(msg.sender, ROYALTY_FEE);
	}

	// ------- Upgradable -------
	function _authorizeUpgrade(
		address newImplementation
	) internal override onlyRole(UPGRADER_ROLE) {}

	// -------- ERC721 --------

	function _burn(uint256 tokenId) internal override(ERC721RoyaltyUpgradeable, ERC721Upgradeable) {
		require(_exists(tokenId), 'Token does not exist');
		require(_isApprovedOrOwner(_msgSender(), tokenId), 'Caller is not owner nor approved');
		super._burn(tokenId);
	}

	function tokenURI(
		uint256 tokenId
	) public view override(ERC721Upgradeable) returns (string memory) {
		require(_exists(tokenId), 'Ducklings: token does not exist');
		Duckling memory duckling = tokenToDuckling[tokenId];
		string memory genome = duckling.genome.toString();
		string memory birthdate = uint256(duckling.birthdate).toString();
		string memory uri = string.concat(apiBaseURL, genome, '-', birthdate);

		return uri;
	}

	function supportsInterface(
		bytes4 interfaceId
	)
		public
		view
		virtual
		override(
			IERC165Upgradeable,
			ERC721RoyaltyUpgradeable,
			ERC721Upgradeable,
			AccessControlUpgradeable
		)
		returns (bool)
	{
		return super.supportsInterface(interfaceId);
	}

	// -------- ERC2981 Royalties --------

	// TODO: add full customize functions
	function setRoyaltyCollector(address account) public onlyRole(DEFAULT_ADMIN_ROLE) {
		_royaltiesCollector = account;
	}

	function getRoyaltyCollector() public view returns (address) {
		return _royaltiesCollector;
	}

	// -------- API URL --------

	function setAPIBaseURL(string calldata apiBaseURL_) external onlyRole(DEFAULT_ADMIN_ROLE) {
		apiBaseURL = apiBaseURL_;
	}

	// -------- IDucklings --------

	function isOwnerOf(address account, uint256 tokenId) external view returns (bool) {
		require(account != address(0), 'ERC721: owner query for the zero address');
		return account == ownerOf(tokenId);
	}

	function isOwnerOf(address account, uint256[] calldata tokenIds) external view returns (bool) {
		require(account != address(0), 'ERC721: owner query for the zero address');
		for (uint256 i = 0; i < tokenIds.length; i++) {
			if (account != ownerOf(tokenIds[i])) {
				return false;
			}
		}

		return true;
	}

	function getGenome(uint256 tokenId) external view returns (uint256) {
		if (!_exists(tokenId)) revert InvalidTokenId(tokenId);
		return tokenToDuckling[tokenId].genome;
	}

	function getGenomes(uint256[] calldata tokenIds) external view returns (uint256[] memory) {
		// explicitly specify array length
		uint256[] memory genomes = new uint256[](tokenIds.length);

		for (uint256 i = 0; i < tokenIds.length; i++) {
			if (!_exists(tokenIds[i])) revert InvalidTokenId(tokenIds[i]);
			genomes[i] = tokenToDuckling[tokenIds[i]].genome;
		}

		return genomes;
	}

	function isTransferable(uint256 tokenId) external view returns (bool) {
		return !_isTransferable(tokenId);
	}

	function _isTransferable(uint256 tokenId) internal view returns (bool) {
		if (!_exists(tokenId)) revert InvalidTokenId(tokenId);

		uint8 headGene = uint8(tokenToDuckling[tokenId].genome >> 248);
		return headGene & (1 << FLAG_TRANSFERABLE) == 0;
	}

	function _beforeTokenTransfer(
		address, // from,
		address to,
		uint256 firstTokenId,
		uint256 // batchSize - always 1 in ERC721
	) internal view override {
		// burn for not transferable is allowed
		if (to == address(0)) return;

		if (!_isTransferable(firstTokenId)) revert TokenNotTransferable(firstTokenId);
	}

	function mintTo(
		address to,
		uint256 genome
	) external onlyRole(GAME_ROLE) returns (uint256 tokenId) {
		return _mintTo(to, genome);
	}

	function mintBatchTo(
		address to,
		uint256[] calldata genomes
	) external onlyRole(GAME_ROLE) returns (uint256[] memory tokenIds) {
		tokenIds = new uint256[](genomes.length);
		for (uint8 i = 0; i < genomes.length; i++) {
			tokenIds[i] = _mintTo(to, genomes[i]);
		}
	}

	function _mintTo(address to, uint256 genome) internal returns (uint256 tokenId) {
		require(to != address(0), 'ERC721: mint to the zero address');
		require(genome != 0, 'ERC721: mint with zero genome');
		tokenId = nextNewTokenId.current();
		uint64 birthdate = uint64(block.timestamp);
		tokenToDuckling[tokenId] = Duckling(genome, birthdate);
		_safeMint(to, tokenId);
		nextNewTokenId.increment();
		emit Minted(to, tokenId, genome, birthdate, block.chainid);
	}

	function burn(uint256 tokenId) external onlyRole(GAME_ROLE) {
		_burn(tokenId);
	}

	function burnBatch(uint256[] calldata tokenIds) external onlyRole(GAME_ROLE) {
		for (uint256 i = 0; i < tokenIds.length; i++) _burn(tokenIds[i]);
	}
}
