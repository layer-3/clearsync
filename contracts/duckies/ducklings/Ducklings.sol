// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import '@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721BurnableUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721RoyaltyUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol';

import '@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/StringsUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol';

import './RandomUpgradeable.sol';
import './Gene.sol';

// TODO:  1. Benchmark and optimize if necessary
contract Ducklings is
	Initializable,
	ERC721Upgradeable,
	ERC721BurnableUpgradeable,
	ERC721RoyaltyUpgradeable,
	UUPSUpgradeable,
	AccessControlUpgradeable,
	RandomUpgradeable
{
	using CountersUpgradeable for CountersUpgradeable.Counter;
	using Gene for uint256;
	using {StringsUpgradeable.toString} for uint256;
	using ECDSAUpgradeable for bytes32;

	error InvalidCollection(Collection collection);
	error CollectionNotAvailable(uint8 collectionId);
	error MintingRulesViolated(uint8 collectionId, uint8 amount);
	error MeldingRulesViolated(uint256[5] tokenIdsToMeld);
	error IncorrectGenesForMelding(uint256[5] genes);

	struct Collection {
		uint64 availableBefore;
		bool isMeldable;
		uint8[][] traitWeights;
	}

	struct Duckling {
		uint256 gene;
		uint64 birthdate;
		uint8 collectionId;
	}

	// TODO: review roles
	bytes32 public constant UPGRADER_ROLE = keccak256('UPGRADER_ROLE');
	bytes32 public constant MAINTAINER_ROLE = keccak256('MAINTAINER_ROLE');
	bytes32 public constant VOUCHER_ISSUER_ROLE = keccak256('VOUCHER_ISSUER_ROLE');

	uint32 private constant ROYALTY_FEE = 1000; // 10%

	uint8 public constant ZOMBEAK_COLLECTION = 0;

	uint8 public constant MAX_MINT_PACK_SIZE = 10;
	uint8 public constant MELD_TOKENS_AMOUNT = 5;

	uint256 public mintPrice;
	uint256 public meldPrice;

	address private _royaltiesCollector;

	string public apiBaseURL;

	CountersUpgradeable.Counter nextNewTokenId;

	CountersUpgradeable.Counter nextCollectionId;

	mapping(uint256 => Duckling) public tokenIdToDuckling;

	mapping(uint8 => Collection) public collectionOfId;

	/*
	 * Traits can have even and uneven probabilities, which are represented in a different ways.
	 *
	 * Even probabilities are stored just as a number of trait values.
	 *
	 * weights = [15] means there are 15 possible trait values that have equal chance of being generated.
	 *
	 * A random number V is generated in the bounds [0, N), where N is number of trait values. V is set to be the value of the trait.
	 *
	 * Uneven traits generation uses weighted random, which means specifying the probability for each trait value.
	 * This probability is represented as an array of N numbers, which divide the segment (bounded by 0 and the sum of weights) into N smaller segments.
	 * Each n-th trait value has a chance of `n-th segment / full segment` for being generated.
	 *
	 * Below is an array and a visual representation for trait values A, B, and C with weights 3, 4, and 3.
	 *
	 *   A   B    C
	 * |---|----|---|       weights = [3,4,3]
	 * 0   3    7   10
	 *
	 * Here is the algorithm for selecting a trait value given the weights array above:
	 * 1. a random number R in limits [0, 10) is generated.
	 * 2. find S, which is an index of the segment R lies in.
	 * 3. set the value of the trait as S
	 *
	 * Also, for default trait values, that should not be generated randomly, a segment [0, 0) is used
	 *
	 * Moreover, trait weights must be stored in ascending order, so that common trait values are preceding the rare ones.
	 */

	uint8[] public meldWeights;
	uint8[][3] meldingZombeakWeights;

	ERC20BurnableUpgradeable public duckiesContract;

	// events
	event Minted(uint256 mintedTokenId, uint256 mintedGene, address owner, uint256 chainId);
	event Melded(
		uint256[MELD_TOKENS_AMOUNT] meldingTokenIds,
		uint256 meldedTokenId,
		uint256 meldedGene,
		address owner,
		uint256 chainId
	);

	// constructor

	function initialize(address ducklingsAddress) public initializer {
		__ERC721_init('Yellow Ducklings NFT Collection', 'YDNC');
		__ERC721Burnable_init();
		__ERC721Royalty_init();
		__AccessControl_init();
		__UUPSUpgradeable_init();
		__Random_init();

		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(UPGRADER_ROLE, msg.sender);
		_grantRole(MAINTAINER_ROLE, msg.sender);

		_setDefaultRoyalty(msg.sender, ROYALTY_FEE);
		setRoyaltyCollector(msg.sender);

		meldWeights = [6, 5, 4, 3, 2, 1];
		meldingZombeakWeights = [[10, 90], [5, 90], [2, 90]];

		// TODO: define price
		mintPrice = 10_000 * 10 ** duckiesContract.decimals();
		meldPrice = 10_000 * 10 ** duckiesContract.decimals();

		duckiesContract = ERC20BurnableUpgradeable(ducklingsAddress);
	}

	// upgrades

	function _authorizeUpgrade(
		address newImplementation
	) internal override onlyRole(UPGRADER_ROLE) {}

	// ERC721

	function _burn(uint256 tokenId) internal override(ERC721RoyaltyUpgradeable, ERC721Upgradeable) {
		super._burn(tokenId);
	}

	function tokenURI(
		uint256 tokenId
	) public view override(ERC721Upgradeable) returns (string memory) {
		Duckling memory duckling = tokenIdToDuckling[tokenId];

		return
			bytes(apiBaseURL).length > 0
				? string(
					abi.encodePacked(
						apiBaseURL,
						duckling.gene.toString(),
						'-',
						uint256(duckling.birthdate).toString(),
						'-',
						uint256(duckling.collectionId).toString()
					)
				)
				: '';
	}

	function supportsInterface(
		bytes4 interfaceId
	)
		public
		view
		virtual
		override(ERC721RoyaltyUpgradeable, ERC721Upgradeable, AccessControlUpgradeable)
		returns (bool)
	{
		return super.supportsInterface(interfaceId);
	}

	// ERC2981 Royalties

	function setRoyaltyCollector(address account) public onlyRole(DEFAULT_ADMIN_ROLE) {
		_royaltiesCollector = account;
	}

	function getRoyaltyCollector() public view returns (address) {
		return _royaltiesCollector;
	}

	// public

	function setAPIBaseURL(string calldata apiBaseURL_) external onlyRole(MAINTAINER_ROLE) {
		apiBaseURL = apiBaseURL_;
	}

	function setMintPrice(uint256 price) external onlyRole(MAINTAINER_ROLE) {
		mintPrice = price;
	}

	function setMeldPrice(uint256 price) external onlyRole(MAINTAINER_ROLE) {
		meldPrice = price;
	}

	// collections

	function addCollection(Collection calldata collection) external onlyRole(MAINTAINER_ROLE) {
		if (
			collection.availableBefore <= block.timestamp ||
			collection.traitWeights.length != uint8(type(Gene.Traits).max) + 1
		) revert InvalidCollection(collection);

		collectionOfId[uint8(nextCollectionId.current())] = collection;

		nextCollectionId.increment();
	}

	function obsoleteCollection(uint8 collectionId) external onlyRole(MAINTAINER_ROLE) {
		_requireValidCollection(collectionId);

		collectionOfId[collectionId].availableBefore = uint64(block.timestamp);
	}

	// mint, meld

	function mintByPayment(uint8 collectionId, uint8 amount) external UseRandom {
		duckiesContract.burnFrom(msg.sender, mintPrice * amount);
		_mintPackTo(msg.sender, collectionId, amount);
	}

	function mintByVoucher(
		address to,
		uint8 collectionId,
		uint8 amount
	) external UseRandom onlyRole(VOUCHER_ISSUER_ROLE) {
		_mintPackTo(to, collectionId, amount);
	}

	function meldByPayment(
		uint256[MELD_TOKENS_AMOUNT] calldata meldingTokenIds
	) external UseRandom {
		duckiesContract.burnFrom(msg.sender, meldPrice);
		_meldOf(msg.sender, meldingTokenIds);
	}

	function meldByVoucher(
		address owner,
		uint256[MELD_TOKENS_AMOUNT] calldata meldingTokenIds
	) external UseRandom onlyRole(VOUCHER_ISSUER_ROLE) {
		_meldOf(owner, meldingTokenIds);
	}

	// internal minting

	function _mintPackTo(address to, uint8 collectionId, uint8 amount) internal {
		_requireValidCollection(collectionId);

		if (amount > MAX_MINT_PACK_SIZE) revert MintingRulesViolated(collectionId, amount);

		for (uint256 i = 0; i < amount; i++) {
			(uint256 tokenId, uint256 gene) = _mintTo(to, collectionId);
			emit Minted(tokenId, gene, msg.sender, block.chainid);
		}
	}

	function _mintTo(
		address to,
		uint8 collectionId
	) internal returns (uint256 tokenId, uint256 gene) {
		gene = _generateGeneFromCollection(collectionId);

		tokenId = nextNewTokenId.current();
		tokenIdToDuckling[tokenId] = Duckling(gene, uint64(block.timestamp), collectionId);
		_safeMint(to, tokenId);
		nextNewTokenId.increment();
	}

	function _generateGeneFromCollection(uint8 collectionId) internal returns (uint256) {
		uint256 gene = 0;

		// generate class
		uint8 class = _generateTraitFromCollection(Gene.Traits.Class, collectionId);
		gene = gene.setTrait(Gene.Traits.Class, class);

		uint8 traitIdx = uint8(Gene.Traits.Class) + 1;

		// if Common, skip Head and Body traits (they are defined as 0)
		if (Gene.Classes(class) == Gene.Classes.Common) {
			traitIdx = uint8(Gene.Traits.Head) + 1;
		}

		// if Rare, skip Body trait (it is defined as 0)
		if (Gene.Classes(class) == Gene.Classes.Rare) {
			traitIdx = uint8(Gene.Traits.Body) + 1;
		}

		// generate and write to gene other traits starting at 'traitIdx'
		for (; traitIdx < uint8(type(Gene.Traits).max); traitIdx++) {
			uint8 trait = _generateTraitFromCollection(Gene.Traits(traitIdx), collectionId);
			gene = gene.setTrait(Gene.Traits(traitIdx), trait);
		}

		return gene;
	}

	// seed must be different for each invocation
	function _generateTraitFromCollection(
		Gene.Traits trait,
		uint8 collectionId
	) internal returns (uint8) {
		uint8[] memory traitWeights = collectionOfId[collectionId].traitWeights[uint8(trait)];

		// check whether trait values has the same probabilities
		if (traitWeights.length == 1) {
			// the same probabilities, just get random of the values
			return uint8(_randomMaxNumber(traitWeights[0]));
		} else {
			// uneven probabilities, generate weighted number
			return _randomWeightedNumber(traitWeights);
		}
	}

	// internal melding

	function _meldOf(address owner, uint256[MELD_TOKENS_AMOUNT] memory meldingTokenIds) internal {
		uint8 collectionId = tokenIdToDuckling[meldingTokenIds[0]].collectionId;

		_requireIsOwnerOf(owner, meldingTokenIds);
		if (!collectionOfId[collectionId].isMeldable) revert MeldingRulesViolated(meldingTokenIds);
		_requireEqualCollection(meldingTokenIds);

		uint256[MELD_TOKENS_AMOUNT] memory genesToMeld = _burnDucklingsAndGetGenes(meldingTokenIds);

		_requireGenesSatisfyMelding(genesToMeld);

		uint256 meldedGene;

		if (_checkZombeak(genesToMeld[0])) {
			meldedGene = _generateGeneFromCollection(ZOMBEAK_COLLECTION);
			collectionId = ZOMBEAK_COLLECTION;
		} else {
			meldedGene = _meldGenes(genesToMeld, collectionId);
		}

		uint256 meldedTokenId = nextNewTokenId.current();
		tokenIdToDuckling[meldedTokenId] = Duckling(
			meldedGene,
			uint64(block.timestamp),
			collectionId
		);
		_safeMint(msg.sender, meldedTokenId);
		nextNewTokenId.increment();

		emit Melded(meldingTokenIds, meldedTokenId, meldedGene, msg.sender, block.chainid);
	}

	function _requireGenesSatisfyMelding(uint256[MELD_TOKENS_AMOUNT] memory genes) internal pure {
		// Classes should be the same
		// Classes should not be SuperLegendary
		_requireCorrectMeldingClasses(genes);
		Gene.Classes meldingClass = Gene.Classes(genes[0].getTrait(Gene.Traits.Class));

		if (meldingClass == Gene.Classes.Legendary) {
			// Legendary

			// cards must have the same Background
			// cards must be of each Element
			if (
				!_traitValuesAreEqual(genes, Gene.Traits.Background) ||
				!_traitValuesAreUnique(genes, Gene.Traits.Element)
			) revert IncorrectGenesForMelding(genes);
		} else {
			// Common, Rare, Epic

			// cards must have the same Background or the same Element
			if (
				!_traitValuesAreEqual(genes, Gene.Traits.Background) &&
				!_traitValuesAreEqual(genes, Gene.Traits.Element)
			) revert IncorrectGenesForMelding(genes);
		}
	}

	function _meldGenes(
		uint256[MELD_TOKENS_AMOUNT] memory genes,
		uint8 collectionId
	) internal returns (uint256) {
		Gene.Classes meldingClass = Gene.Classes(genes[0].getTrait(Gene.Traits.Class));

		uint256 gene = 0;

		// increment class
		gene = gene.setTrait(Gene.Traits.Class, uint8(meldingClass) + 1);

		// if melded SuperLegendary, set only Peculiarity trait
		if (meldingClass == Gene.Classes.Legendary) {
			return
				gene.setTrait(
					Gene.Traits.Peculiarity,
					_generateTraitFromCollection(Gene.Traits.Peculiarity, collectionId)
				);
		}

		// optimization: read only once from the storage is much more efficient than reading in every iteration
		uint8[][] memory traitWeights = collectionOfId[collectionId].traitWeights;

		for (uint8 traitIdx = 1; traitIdx < uint8(type(Gene.Traits).max); traitIdx++) {
			uint8 maxTraitValue = uint8(traitWeights[traitIdx].length) - 1;
			uint8 meldedTrait = _meldTraits(genes, Gene.Traits(traitIdx), maxTraitValue);
			gene = gene.setTrait(Gene.Traits(traitIdx), meldedTrait);
		}

		return gene;
	}

	function _meldTraits(
		uint256[MELD_TOKENS_AMOUNT] memory genes,
		Gene.Traits trait,
		uint8 maxTraitValue
	) internal returns (uint8) {
		uint8 meldedTraitIdx = _randomWeightedNumber(meldWeights);

		// mutation, return upgraded best trait value
		if (meldedTraitIdx == MELD_TOKENS_AMOUNT) {
			uint8 bestTraitValue = _maxTrait(genes, trait);
			return bestTraitValue == maxTraitValue ? bestTraitValue : bestTraitValue + 1;
		}

		// no mutation, return selected trait value
		return genes[meldedTraitIdx].getTrait(trait);
	}

	function _requireCorrectMeldingClasses(uint256[MELD_TOKENS_AMOUNT] memory genes) internal pure {
		Gene.Classes class = Gene.Classes(genes[0].getTrait(Gene.Traits.Class));

		for (uint256 i = 1; i < genes.length; i++) {
			if (genes[i].getTrait(Gene.Traits.Class) != uint8(class))
				revert IncorrectGenesForMelding(genes);
		}

		if (class == Gene.Classes.SuperLegendary) revert IncorrectGenesForMelding(genes);
	}

	function _checkZombeak(uint256 gene) internal returns (bool) {
		Gene.Classes class = Gene.Classes(gene.getTrait(Gene.Traits.Class));

		if (uint8(class) <= uint8(Gene.Classes.Epic)) {
			return _randomWeightedNumber(meldingZombeakWeights[uint8(class)]) == 0;
		} else {
			return false;
		}
	}

	// other internal

	function _requireValidCollection(uint8 collectionId) internal view {
		if (
			collectionId >= nextCollectionId.current() ||
			collectionOfId[collectionId].availableBefore < block.timestamp
		) revert CollectionNotAvailable(collectionId);
	}

	function _requireIsOwnerOf(
		address owner,
		uint256[MELD_TOKENS_AMOUNT] memory tokenIds
	) internal view {
		for (uint256 i = 0; i < tokenIds.length; i++) {
			if (owner != ownerOf(tokenIds[i])) revert MeldingRulesViolated(tokenIds);
		}
	}

	function _requireEqualCollection(uint256[MELD_TOKENS_AMOUNT] memory tokenIds) internal view {
		uint8 collection = tokenIdToDuckling[tokenIds[0]].collectionId;

		for (uint256 i = 1; i < tokenIds.length; i++) {
			if (collection != tokenIdToDuckling[tokenIds[i]].collectionId)
				revert MeldingRulesViolated(tokenIds);
		}
	}

	function _burnDucklingsAndGetGenes(
		uint256[MELD_TOKENS_AMOUNT] memory tokenIds
	) internal returns (uint256[MELD_TOKENS_AMOUNT] memory) {
		uint256[MELD_TOKENS_AMOUNT] memory genes;

		for (uint256 i = 0; i < tokenIds.length; i++) {
			genes[i] = tokenIdToDuckling[tokenIds[i]].gene;
			_burn(tokenIds[i]);
		}

		return genes;
	}

	function _maxTrait(
		uint256[MELD_TOKENS_AMOUNT] memory genes,
		Gene.Traits trait
	) internal pure returns (uint8) {
		uint8 maxValue = 0;

		for (uint256 i = 0; i < genes.length; i++) {
			uint8 traitValue = genes[i].getTrait(trait);
			if (maxValue < traitValue) {
				maxValue = traitValue;
			}
		}

		return maxValue;
	}

	function _randomWeightedNumber(uint8[] memory weights) internal returns (uint8) {
		// generated number should be strictly less than right \/ segment boundary
		uint256 randomNumber = _randomMaxNumber(_sum(weights) - 1);

		uint256 segmentRightBoundary = 0;

		for (uint8 i = 0; i < weights.length; i++) {
			segmentRightBoundary += weights[i];
			if (randomNumber < segmentRightBoundary) {
				return i;
			}
		}

		// execution should never reach this
		return uint8(weights.length - 1);
	}

	function _sum(uint8[] memory numbers) internal pure returns (uint256) {
		uint8 sum = 0;

		for (uint256 i = 0; i < numbers.length; i++) {
			sum += numbers[i];
		}

		return sum;
	}

	function _traitValuesAreEqual(
		uint256[MELD_TOKENS_AMOUNT] memory genes,
		Gene.Traits trait
	) internal pure returns (bool) {
		uint8 traitValue = genes[0].getTrait(trait);

		for (uint256 i = 1; i < genes.length; i++) {
			if (genes[i].getTrait(trait) != traitValue) {
				return false;
			}
		}

		return true;
	}

	function _traitValuesAreUnique(
		uint256[MELD_TOKENS_AMOUNT] memory genes,
		Gene.Traits trait
	) internal pure returns (bool) {
		uint256 valuesPresentBitfield = 0;

		for (uint256 i = 1; i < genes.length; i++) {
			if (valuesPresentBitfield % 2 ** genes[i].getTrait(trait) == 1) {
				return false;
			}
			valuesPresentBitfield += 2 ** genes[i].getTrait(trait);
		}

		return true;
	}
}
