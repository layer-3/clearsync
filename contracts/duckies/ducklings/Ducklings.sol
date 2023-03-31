// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721RoyaltyUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol';

import '@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/StringsUpgradeable.sol';

import '../../interfaces/IDucklings.sol';

contract Ducklings is
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

	// roles
	bytes32 public constant UPGRADER_ROLE = keccak256('UPGRADER_ROLE');
	bytes32 public constant MAINTAINER_ROLE = keccak256('MAINTAINER_ROLE');
	bytes32 public constant GAME_ROLE = keccak256('GAME_ROLE');

	// royalty
	address private _royaltiesCollector;
	uint32 private constant ROYALTY_FEE = 1000; // 10%

	string public apiBaseURL;

	CountersUpgradeable.Counter public nextNewTokenId;
	mapping(uint256 => Duckling) public idToDuckling;

	// ------- Initializer -------

	function initialize() external initializer {
		__ERC721_init('Yellow Ducklings NFT Collection', 'YDNC');
		__ERC721Royalty_init();
		__AccessControl_init();
		__UUPSUpgradeable_init();

		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(UPGRADER_ROLE, msg.sender);
		_grantRole(MAINTAINER_ROLE, msg.sender);

		setRoyaltyCollector(msg.sender);
		_setDefaultRoyalty(msg.sender, ROYALTY_FEE);
	}

	// -------- Upgrades --------

	function _authorizeUpgrade(
		address newImplementation
	) internal override onlyRole(UPGRADER_ROLE) {}

	// -------- ERC721 --------

	function _burn(uint256 tokenId) internal override(ERC721RoyaltyUpgradeable, ERC721Upgradeable) {
		super._burn(tokenId);
	}

	function tokenURI(
		uint256 tokenId
	) public view override(ERC721Upgradeable) returns (string memory) {
		Duckling memory duckling = idToDuckling[tokenId];

		return
			bytes(apiBaseURL).length > 0
				? string.concat(
					apiBaseURL,
					duckling.genome.toString(),
					'-',
					uint256(duckling.birthdate).toString()
				)
				: '';
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

	function setAPIBaseURL(string calldata apiBaseURL_) external onlyRole(MAINTAINER_ROLE) {
		apiBaseURL = apiBaseURL_;
	}

	// -------- IDucklings --------

	function isOwnerOf(address account, uint256 tokenId) external view returns (bool) {
		return account == ownerOf(tokenId);
	}

	function isOwnerOf(address account, uint256[] calldata tokenIds) external view returns (bool) {
		for (uint256 i = 0; i < tokenIds.length; i++) {
			if (account != ownerOf(tokenIds[i])) {
				return false;
			}
		}

		return true;
	}

	function getGenome(uint256 tokenId) external view returns (uint256) {
		if (!_exists(tokenId)) revert InvalidTokenId(tokenId);
		return idToDuckling[tokenId].genome;
	}

	function getGenomes(uint256[] calldata tokenIds) external view returns (uint256[] memory) {
		// explicitly specify array length
		uint256[] memory genomes = new uint256[](tokenIds.length);

		for (uint256 i = 0; i < tokenIds.length; i++) {
			if (!_exists(tokenIds[i])) revert InvalidTokenId(tokenIds[i]);
			genomes[i] = idToDuckling[tokenIds[i]].genome;
		}

		return genomes;
	}

	function mintTo(
		address to,
		uint256 genome
	) external onlyRole(GAME_ROLE) returns (uint256 tokenId) {
		tokenId = nextNewTokenId.current();
		uint64 birthdate = uint64(block.timestamp);
		idToDuckling[tokenId] = Duckling(genome, birthdate);

		_safeMint(to, tokenId);

		nextNewTokenId.increment();
		emit Minted(to, tokenId, genome, birthdate, block.chainid);
	}

	function burn(uint256 tokenId) external onlyRole(GAME_ROLE) {
		_burn(tokenId);
	}

	function burn(uint256[] calldata tokenIds) external onlyRole(GAME_ROLE) {
		for (uint256 i = 0; i < tokenIds.length; i++) _burn(tokenIds[i]);
	}
}
