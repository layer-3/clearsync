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

// TODO:  1. Add trait struct for each collection
//        2. Dynamic price per mint
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

	struct Duckling {
		uint256 gene;
		uint64 birthdate;
	}

	bytes32 public constant UPGRADER_ROLE = keccak256('UPGRADER_ROLE');
	bytes32 public constant API_SETTER_ROLE = keccak256('API_SETTER_ROLE');
	bytes32 public constant ROYALTIES_COLLECTOR_ROLE = keccak256('ROYALTIES_COLLECTOR_ROLE');

	uint32 private constant ROYALTY_FEE = 1000; // 10%

	uint8 public constant MAX_MINT_PACK_SIZE = 10;
	uint256 public BASE_DUCKIES_PER_MINT;

	address private _royaltiesCollector;

	string public apiBaseURL;
	// TODO: remove
	bytes32 private salt;

	CountersUpgradeable.Counter nextNewTokenId;

	mapping(uint256 => Duckling) public tokenIdToDuckling;

	/*
	 * Traits can have even and uneven probabilities, which are represented in a different ways.
	 *
	 * Even probabilities are stored just as a number of trait values.
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
	uint8[][] public traitWeights;
	uint8[] public meldWeights;
	uint8[][3] meldingZombeakWeights;

	ERC20BurnableUpgradeable public duckiesContract;

	// events
	event Minted(uint256 mintedTokenId, uint256 mintedGene, address owner, uint256 chainId);
	event Melded(
		uint256[5] meldingTokenIds,
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
		_grantRole(API_SETTER_ROLE, msg.sender);
		_grantRole(ROYALTIES_COLLECTOR_ROLE, msg.sender);

		_setDefaultRoyalty(msg.sender, ROYALTY_FEE);
		setRoyaltyCollector(msg.sender);

		// Arrays are pushed due to difference in size. Solidity is bad in converting fixed-size memory array into dynamic one.
		// Class
		traitWeights.push([74, 20, 5, 1]);
		// Body
		traitWeights.push([0, 17, 16, 14, 14, 12, 10, 9, 8]);
		// Head
		traitWeights.push([0, 9, 9, 7, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 3, 3]);
		// Background
		traitWeights.push([4]);
		// Element
		traitWeights.push([5]);
		// Eyes
		traitWeights.push([8, 7, 6, 6, 6, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 3, 3, 3, 2, 2]);
		// Beak
		traitWeights.push([15, 14, 14, 11, 10, 9, 8, 7, 7, 6]);
		// Wings
		traitWeights.push([20, 19, 15, 13, 12, 11, 10]);
		// First name
		traitWeights.push([32]);
		// Last name
		traitWeights.push([17]);
		// Temper
		// TODO: define probabilities
		traitWeights.push([16]);
		// Peculiarity - for SuperLegendary
		// TODO: how many?
		traitWeights.push([100]);

		meldWeights = [6, 5, 4, 3, 2, 1];
		meldingZombeakWeights = [[10, 90], [5, 90], [2, 90]];

		duckiesContract = ERC20BurnableUpgradeable(ducklingsAddress);
		BASE_DUCKIES_PER_MINT = 10_000 * 10 ** duckiesContract.decimals();
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
		return
			bytes(apiBaseURL).length > 0
				? string(
					abi.encodePacked(
						apiBaseURL,
						tokenIdToDuckling[tokenId].gene.toString(),
						'-',
						uint256(tokenIdToDuckling[tokenId].birthdate).toString()
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
		_grantRole(ROYALTIES_COLLECTOR_ROLE, account);
	}

	function getRoyaltyCollector() public view returns (address) {
		return _royaltiesCollector;
	}

	// public

	function setAPIBaseURL(string calldata apiBaseURL_) external onlyRole(API_SETTER_ROLE) {
		apiBaseURL = apiBaseURL_;
	}

	function ducklingsPerMint() external view returns (uint256) {
		return _ducklingsPerMint();
	}

	function mintPack(uint8 amount) external UseRandom {
		require(amount <= MAX_MINT_PACK_SIZE, 'pack size exceeded');

		duckiesContract.burnFrom(msg.sender, _ducklingsPerMint() * amount);

		for (uint256 i = 0; i < amount; i++) {
			(uint256 tokenId, uint256 gene) = _mint();
			emit Minted(tokenId, gene, msg.sender, _getChainId());
		}
	}

	function meld(uint256[5] calldata meldingTokenIds) external UseRandom {
		_requireCallerIsOwner(meldingTokenIds);

		uint256[5] memory genesToMeld = _burnTokensAndGetGenes(meldingTokenIds);

		uint256 meldedGene = _meldGenes(genesToMeld);

		uint256 meldedTokenId = nextNewTokenId.current();
		tokenIdToDuckling[meldedTokenId] = Duckling(meldedGene, uint64(block.timestamp));
		_safeMint(msg.sender, meldedTokenId);
		nextNewTokenId.increment();

		emit Melded(meldingTokenIds, meldedTokenId, meldedGene, msg.sender, _getChainId());
	}

	// internal minting

	function _mint() internal returns (uint256 tokenId, uint256 gene) {
		gene = _generateGene();

		tokenId = nextNewTokenId.current();
		tokenIdToDuckling[tokenId] = Duckling(gene, uint64(block.timestamp));
		_safeMint(msg.sender, tokenId);
		nextNewTokenId.increment();
	}

	function _generateGene() internal returns (uint256) {
		uint256 gene = 0;

		// generate class
		uint8 class = _generateTrait(Gene.Traits.Class);
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
		for (; traitIdx < traitWeights.length; traitIdx++) {
			uint8 trait = _generateTrait(Gene.Traits(traitIdx));
			// FIX: inconsistency where Traits.UniquenessIdx != traitWeights.length
			gene = gene.setTrait(Gene.Traits(traitIdx), trait);
		}

		return gene;
	}

	// seed must be different for each invocation
	function _generateTrait(Gene.Traits trait) internal returns (uint8) {
		// check whether trait values has the same probabilities
		if (traitWeights[uint8(trait)].length == 1) {
			// the same probabilities, just get random of the values
			return uint8(_randomMaxNumber(traitWeights[uint8(trait)][0]));
		} else {
			// uneven probabilities, generate weighted number
			return _randomWeightedNumber(traitWeights[uint8(trait)]);
		}
	}

	// internal melding

	function _meldGenes(uint256[5] memory genes) internal returns (uint256) {
		// Classes should be the same
		// Classes should not be Legendary or SuperLegendary
		_requireCorrectMeldingClasses(genes);
		Gene.Classes meldingClass = Gene.Classes(genes[0].getTrait(Gene.Traits.Class));

		uint256 gene = 0;

		if (meldingClass == Gene.Classes.Legendary) {
			// cards mush have the same Background
			// cards must be of each Element
			_requireCorrectMeldingTraits(genes);
		} else {
			// Common, Rare, Epic
			if (_checkZombeak(meldingClass)) {
				gene = gene.setTrait(Gene.Traits.Class, uint8(Gene.Classes.Zombie));
				return gene;
			}
		}

		// increment class
		gene = gene.setTrait(Gene.Traits.Class, uint8(meldingClass) + 1);

		// if melded SuperLegendary, set only Peculiarity trait
		if (meldingClass == Gene.Classes.Legendary) {
			return gene.setTrait(Gene.Traits.Peculiarity, _generateTrait(Gene.Traits.Peculiarity));
		}

		for (uint256 traitIdx = 1; traitIdx < traitWeights.length; traitIdx++) {
			uint8 meldedTrait = _meldTraits(genes, Gene.Traits(traitIdx));
			gene = gene.setTrait(Gene.Traits(traitIdx), meldedTrait);
		}

		return gene;
	}

	function _meldTraits(uint256[5] memory genes, Gene.Traits trait) internal returns (uint8) {
		uint8 meldedTraitIdx = _randomWeightedNumber(meldWeights);

		// mutation, return upgraded best trait value
		if (meldedTraitIdx == 5) {
			uint8 bestTraitValue = _maxTrait(genes, trait);
			return
				bestTraitValue == traitWeights[uint8(trait)].length - 1
					? bestTraitValue
					: bestTraitValue + 1;
		}

		// no mutation, return selected trait value
		return genes[meldedTraitIdx].getTrait(trait);
	}

	function _requireCorrectMeldingClasses(uint256[5] memory genes) internal pure {
		Gene.Classes class = Gene.Classes(genes[0].getTrait(Gene.Traits.Class));

		for (uint256 i = 1; i < genes.length; i++) {
			require(
				genes[i].getTrait(Gene.Traits.Class) == uint8(class),
				'melding Classes must be equal'
			);
		}

		require(
			class != Gene.Classes.SuperLegendary && class != Gene.Classes.Zombie,
			'wrong Class to meld'
		);
	}

	// applies only to melding Legendary into SuperLegendary
	// cards mush have the same Background
	// cards must be of each Element
	function _requireCorrectMeldingTraits(uint256[5] memory genes) internal pure {
		uint8 color = genes[0].getTrait(Gene.Traits.Background);
		uint256 elementsPresentBitfield = 0;

		for (uint256 i = 1; i < genes.length; i++) {
			require(
				genes[i].getTrait(Gene.Traits.Background) == color,
				'melding legendary Backgrounds must be equal'
			);
			elementsPresentBitfield += 2 ** genes[i].getTrait(Gene.Traits.Element);
		}

		uint256 allElementsBitfield = 2 ** 5 - 1; // in bin representation there is a '1' set at indexes for each Element value
		require(
			elementsPresentBitfield == allElementsBitfield,
			'mending legendary requires all Elements'
		);
	}

	function _checkZombeak(Gene.Classes class) internal returns (bool) {
		if (uint8(class) <= uint8(Gene.Classes.Epic)) {
			return _randomWeightedNumber(meldingZombeakWeights[uint8(class)]) == 0;
		} else {
			return false;
		}
	}

	// other internal

	function _ducklingsPerMint() internal view returns (uint256) {
		return BASE_DUCKIES_PER_MINT;
	}

	function _requireCallerIsOwner(uint256[5] memory tokenIds) internal view {
		address account = msg.sender;

		for (uint256 i = 0; i < tokenIds.length; i++) {
			require(account == ownerOf(tokenIds[i]), 'caller is not token owner');
		}
	}

	function _burnTokensAndGetGenes(
		uint256[5] memory tokenIds
	) internal returns (uint256[5] memory) {
		uint256[5] memory genes;

		for (uint256 i = 0; i < tokenIds.length; i++) {
			genes[i] = tokenIdToDuckling[tokenIds[i]].gene;
			_burn(tokenIds[i]);
		}

		return genes;
	}

	function _maxTrait(uint256[5] memory genes, Gene.Traits trait) internal pure returns (uint8) {
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

	function _getChainId() internal view returns (uint256 id) {
		assembly {
			id := chainid()
		}
	}
}
