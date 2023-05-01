// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts/access/AccessControl.sol';

import '@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol';
import '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';

import '../../../interfaces/IDuckyFamily.sol';
import '../../../interfaces/IDucklings.sol';
import '../Random.sol';
import '../Genome.sol';

/**
 * @title DuckyFamilyV1
 *
 * @notice DuckyFamily contract defines rules of Ducky Family game, which is a card game similar to Happy Families and Uno games.
 * This game also incorporates vouchers as defined in IVoucher interface.
 *
 * DuckyFamily operates on Ducklings NFT, which is defined in a corresponding contract. DuckyFamily can mint, burn and query information about NFTs
 * by calling Ducklings contract.
 *
 * Users can buy NFT (card) packs of different size. When a pack is bought, a number of cards is generated and assigned to the user.
 * The packs can be bought with Duckies token, so user should approve DuckyFamily contract to spend Duckies on his behalf.
 *
 * Each card has a unique genome, which is a 256-bit number. The genome is a combination of different genes, which describe the card and its properties.
 * There are 3 types of cards introduced in this game, which are differentiated by the 'collection' gene: Duckling, Zombeak and Mythic.
 * Duckling and Zombeak NFTs have a class system, which is defined by 'rarity' gene: Common, Rare, Epic and Legendary.
 * Mythic NFTs are not part of the class system and are considered to be the most rare and powerful cards in the game.
 *
 * All cards have a set of generative genes, which are used to describe the card, its rarity and image.
 * There are 2 types of generative genes: with even and uneven chance for each value of that gene.
 * All values of even genes are generated with equal probability, while uneven genes have a higher chance for the first values and lower for the last values.
 * Thus, while even genes can describe the card, uneven can set the rarity of the card.
 *
 * Note: don't confuse 'rarity' gene with rarity of the card. 'Rarity' gene is a part of the game logic, while rarity of the card is a value this card represents.
 * Henceforth, if a 'Common' rarity gene card has uneven generative genes with high values (which means this card has a tiny chance to being generated),
 * then this card can be more rare than some 'Rare' rarity gene cards.
 * So, when we mean 'rarity' gene, we will use quotes, while when we mean rarity of the card, we will use it without quotes.
 *
 * Duckling are the main cards in the game, as they are the only way users can get Mythic cards.
 * However, users are not obliged to use every Duckling cards to help them get Mythic, they can improve them and collect the rarest ones.
 * Users can get Duckling cards from minting packs.
 *
 * Users can improve the 'rarity' of the card by melding them. Melding is a process of combining a flock of 5 cards to create a new one.
 * The new card will have the same 'collection' gene as the first card in the flock, but the 'rarity' gene will be incremented.
 * However, users must oblige to specific rules when melding cards:
 * 1. All cards in the flock must have the same 'collection' gene.
 * 2. All cards in the flock must have the same 'rarity' gene.
 * 3a. When melding Common cards, all cards in the flock must have either the same Color or Family gene values.
 * 3b. When melding Rare and Epic cards, all cards in the flock must have both the same Color and Family gene values.
 * 3c. When melding Legendary cards, all cards in the flock must have the same Color and different Family gene values.
 * 4. Mythic cards cannot be melded.
 * 5. Legendary Zombeak cards cannot be melded.
 *
 * Other generative genes of the melded card are not random, but are calculated from the genes of the source cards.
 * This process is called 'inheritance' and is the following:
 * 1. Each generative gene is inherited separately
 * 2. A gene has a high chance of being inherited from the first card in the flock, and this chance is lower for each next card in the flock.
 * 3. A gene has a mere chance of 'positive mutation', which sets inherited gene value to be bigger than the biggest value of this gene in the flock.
 *
 * Melding is not free and has a different cost for each 'rarity' of the cards being melded.
 *
 * Zombeak are secondary cards, that you can only get when melding mutates. There is a different chance (defined in Config section below) for each 'rarity' of the Duckling cards that are being melded,
 * that the melding result card will mutate to Zombeak. If the melding mutates, then the new card will have the same 'rarity' gene as the source cards.
 * This logic makes Zombeak cards more rare than some Duckling cards, as they can only be obtained by melding mutating.
 * However, Zombeak cards cannot be melded into Mythic, which means their main value is rarity.
 *
 * Mythic are the most rare and powerful cards in the game. They can only be obtained by melding Legendary Duckling cards with special rules described above.
 * The rarity of the Mythic card is defined by the 'UniqId' gene, which corresponds to the picture of the card. The higher the 'UniqId' gene value, the rarer the card.
 * The 'UniqId' value is correlated with the 'peculiarity' of the flock that creates the Mythic: the higher the peculiarity, the higher the 'UniqId' value.
 * Peculiarity of the card is a sum of all uneven gene values of this card, and peculiarity of the flock is a sum of peculiarities of all cards in the flock.
 *
 * Mythic cards give bonuses to their owned depending on their rarity. These bonuses will be revealed in the future, but they may include
 * free Yellow tokens (with vesting claim mechanism), an ability to change existing cards, stealing / fighting other cards, etc.
 */
contract DuckyFamilyV1 is IDuckyFamily, AccessControl, Random {
	using Genome for uint256;
	using ECDSA for bytes32;

	// Roles
	bytes32 public constant MAINTAINER_ROLE = keccak256('MAINTAINER_ROLE'); // can change minting and melding price

	address public issuer; // issuer of Vouchers

	// Store the vouchers to avoid replay attacks
	mapping(bytes32 => bool) internal _usedVouchers;

	// ------- Config -------

	uint8 internal constant ducklingCollectionId = 0;
	uint8 internal constant zombeakCollectionId = 1;
	uint8 internal constant mythicCollectionId = 2;
	uint8 internal constant RARITIES_NUM = 4;

	uint8 public constant MAX_PACK_SIZE = 50;
	uint8 public constant FLOCK_SIZE = 5;

	uint8 internal constant collectionGeneIdx = Genome.COLLECTION_GENE_IDX;
	uint8 internal constant rarityGeneIdx = 1;
	uint8 internal constant flagsGeneIdx = Genome.FLAGS_GENE_IDX;
	// general genes start after Collection and Rarity
	uint8 internal constant generativeGenesOffset = 2;

	// number of values for each gene for Duckling and Zombeak collections
	uint8[][3] internal collectionsGeneValuesNum; // set in constructor

	// distribution type of each gene for Duckling and Zombeak collections (0 - even, 1 - uneven)
	uint32[3] internal collectionsGeneDistributionTypes = [
		2940, // reverse(001111101101) = 101101111100
		2940, // reverse(001111101101) = 101101111100
		107 // reverse(11010110) = 01101011
	];

	// peculiarity is a sum of uneven gene values for Ducklings
	uint16 internal maxPeculiarity;
	// mythic dispersion define the interval size in which UniqId value is generated
	uint8 internal constant MYTHIC_DISPERSION = 5;
	uint8 internal mythicAmount = 60;

	// chance of a Duckling of a certain rarity to be generated
	uint32[] internal rarityChances = [850, 120, 25, 5]; // per mil

	// chance of a Duckling of certain rarity to mutate to Zombeak while melding
	uint32[] internal collectionMutationChances = [150, 100, 50, 10]; // per mil

	uint32[] internal geneMutationChance = [955, 45]; // 4.5% to mutate gene value
	uint32[] internal geneInheritanceChances = [400, 300, 150, 100, 50]; // per mil

	// ------- Public values -------

	ERC20Burnable public duckiesContract;
	IDucklings public ducklingsContract;
	address public treasureVaultAddress;

	uint256 public mintPrice;
	uint256[RARITIES_NUM] public meldPrices; // [0] - melding Commons, [1] - melding Rares...

	// ------- Constructor -------

	/**
	 * @notice Sets Duckies, Ducklings and Treasure Vault addresses, minting and melding prices and other game config.
	 * @dev Grants DEFAULT_ADMIN_ROLE and MAINTAINER_ROLE to the deployer.
	 * @param duckiesAddress Address of Duckies ERC20 contract.
	 * @param ducklingsAddress Address of Ducklings ERC721 contract.
	 * @param treasureVaultAddress_ Address of Treasure Vault contract.
	 */
	constructor(address duckiesAddress, address ducklingsAddress, address treasureVaultAddress_) {
		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(MAINTAINER_ROLE, msg.sender);

		duckiesContract = ERC20Burnable(duckiesAddress);
		ducklingsContract = IDucklings(ducklingsAddress);
		treasureVaultAddress = treasureVaultAddress_;

		uint256 decimalsMultiplier = 10 ** duckiesContract.decimals();

		mintPrice = 50 * decimalsMultiplier;
		meldPrices = [
			100 * decimalsMultiplier,
			200 * decimalsMultiplier,
			500 * decimalsMultiplier,
			1000 * decimalsMultiplier
		];

		// Duckling genes: (Collection, Rarity), Color, Family, Body, Head, Eyes, Beak, Wings, FirstName, Temper, Skill, Habitat, Breed
		collectionsGeneValuesNum[0] = [4, 5, 10, 25, 30, 14, 10, 36, 16, 12, 5, 28];
		// Zombeak genes: (Collection, Rarity), Color, Family, Body, Head, Eyes, Beak, Wings, FirstName, Temper, Skill, Habitat, Breed
		collectionsGeneValuesNum[1] = [2, 3, 7, 6, 9, 7, 10, 36, 16, 12, 5, 28];
		// Mythic genes: (Collection, UniqId), Temper, Skill, Habitat, Breed, Birthplace, Quirk, Favorite Food, Favorite Color
		collectionsGeneValuesNum[2] = [16, 12, 5, 28, 5, 10, 8, 4];

		maxPeculiarity = _calcMaxPeculiarity();
	}

	// ------- Random -------

	/**
	 * @notice Sets the pepper for random generator.
	 * @dev Require MAINTAINER_ROLE to call. Pepper is a random data changed periodically by external entity.
	 * @param pepper New pepper.
	 */
	function setPepper(bytes32 pepper) external onlyRole(MAINTAINER_ROLE) {
		_setPepper(pepper);
	}

	// ------- Vouchers -------

	/**
	 * @notice Sets the issuer of Vouchers.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param account Address of a new issuer.
	 */
	function setIssuer(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
		issuer = account;
	}

	/**
	 * @notice Use multiple Vouchers. Check the signature and invoke internal function for each voucher.
	 * @dev Vouchers are issued by the Back-End and signed by the issuer.
	 * @param vouchers Array of Vouchers to use.
	 * @param signature Vouchers signed by the issuer.
	 */
	function useVouchers(Voucher[] calldata vouchers, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(vouchers), signature, issuer);
		for (uint8 i = 0; i < vouchers.length; i++) {
			_useVoucher(vouchers[i]);
		}
	}

	/**
	 * @notice Use a single Voucher. Check the signature and invoke internal function.
	 * @dev Vouchers are issued by the Back-End and signed by the issuer.
	 * @param voucher Voucher to use.
	 * @param signature Voucher signed by the issuer.
	 */
	function useVoucher(Voucher calldata voucher, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(voucher), signature, issuer);
		_useVoucher(voucher);
	}

	/**
	 * @notice Check the validity of a voucher, decode voucher params and mint or meld tokens depending on voucher's type. Emits VoucherUsed event. Internal function.
	 * @dev Vouchers are issued by the Back-End and signed by the issuer.
	 * @param voucher Voucher to use.
	 */
	function _useVoucher(Voucher memory voucher) internal {
		_requireValidVoucher(voucher);

		_usedVouchers[voucher.voucherCodeHash] = true;

		// parse & process Voucher
		if (voucher.action == uint8(VoucherActions.MintPack)) {
			MintParams memory mintParams = abi.decode(voucher.encodedParams, (MintParams));

			// mintParams checks
			if (
				mintParams.to == address(0) ||
				mintParams.size == 0 ||
				mintParams.size > MAX_PACK_SIZE
			) revert InvalidMintParams(mintParams);

			_mintPackTo(mintParams.to, mintParams.size, mintParams.isTransferable);
		} else if (voucher.action == uint8(VoucherActions.MeldFlock)) {
			MeldParams memory meldParams = abi.decode(voucher.encodedParams, (MeldParams));

			// meldParams checks
			if (meldParams.owner == address(0) || meldParams.tokenIds.length != FLOCK_SIZE)
				revert InvalidMeldParams(meldParams);

			_meldOf(meldParams.owner, meldParams.tokenIds, meldParams.isTransferable);
		} else {
			revert InvalidVoucher(voucher);
		}

		emit VoucherUsed(
			voucher.beneficiary,
			voucher.action,
			voucher.voucherCodeHash,
			voucher.chainId
		);
	}

	/**
	 * @notice Check the validity of a voucher, reverts if invalid.
	 * @dev Voucher address must be this contract, beneficiary must be msg.sender, voucher must not be used before, voucher must not be expired.
	 * @param voucher Voucher to check.
	 */
	function _requireValidVoucher(Voucher memory voucher) internal view {
		if (_usedVouchers[voucher.voucherCodeHash])
			revert VoucherAlreadyUsed(voucher.voucherCodeHash);

		if (
			voucher.target != address(this) ||
			voucher.beneficiary != msg.sender ||
			block.timestamp > voucher.expire ||
			voucher.chainId != block.chainid
		) revert InvalidVoucher(voucher);
	}

	/**
	 * @notice Check that `signatures is `encodedData` signed by `signer`. Reverts if not.
	 * @dev Check that `signatures is `encodedData` signed by `signer`. Reverts if not.
	 * @param encodedData Data to check.
	 * @param signature Signature to check.
	 * @param signer Address of the signer.
	 */
	function _requireCorrectSigner(
		bytes memory encodedData,
		bytes memory signature,
		address signer
	) internal pure {
		address actualSigner = keccak256(encodedData).toEthSignedMessageHash().recover(signature);
		if (actualSigner != signer) revert IncorrectSigner(signer, actualSigner);
	}

	// -------- Config --------

	/**
	 * @notice Get the mint price in Duckies with decimals.
	 * @dev Get the mint price in Duckies with decimals.
	 * @return mintPrice Mint price in Duckies with decimals.
	 */
	function getMintPrice() external view returns (uint256) {
		return mintPrice;
	}

	/**
	 * @notice Set the mint price in Duckies without decimals.
	 * @dev Require MAINTAINER_ROLE to call.
	 * @param price Mint price in Duckies without decimals.
	 */
	function setMintPrice(uint256 price) external onlyRole(MAINTAINER_ROLE) {
		mintPrice = price * 10 ** duckiesContract.decimals();
	}

	/**
	 * @notice Get the meld price for each 'rarity' in Duckies with decimals.
	 * @dev Get the meld price for each 'rarity' in Duckies with decimals.
	 * @return meldPrices Array of meld prices in Duckies with decimals.
	 */
	function getMeldPrices() external view returns (uint256[RARITIES_NUM] memory) {
		return meldPrices;
	}

	/**
	 * @notice Set the meld price for each 'rarity' in Duckies without decimals.
	 * @dev Require MAINTAINER_ROLE to call.
	 * @param prices Array of meld prices in Duckies without decimals.
	 */
	function setMeldPrices(
		uint256[RARITIES_NUM] calldata prices
	) external onlyRole(MAINTAINER_ROLE) {
		for (uint8 i = 0; i < RARITIES_NUM; i++) {
			meldPrices[i] = prices[i] * 10 ** duckiesContract.decimals();
		}
	}

	/**
	 * @notice Get number of gene values for all collections and a number of different Mythic tokens.
	 * @dev Get number of gene values for all collections and a number of different Mythic tokens.
	 * @return collectionsGeneValuesNum Arrays of number of gene values for all collections and a mythic amount.
	 */
	function getCollectionsGeneValues() external view returns (uint8[][3] memory, uint8) {
		return (collectionsGeneValuesNum, mythicAmount);
	}

	/**
	 * @notice Get gene distribution types for all collections.
	 * @dev Get gene distribution types for all collections.
	 * @return collectionsGeneDistributionTypes Arrays of gene distribution types for all collections.
	 */
	function getCollectionsGeneDistributionTypes() external view returns (uint32[3] memory) {
		return collectionsGeneDistributionTypes;
	}

	/**
	 * @notice Set gene values number for each gene for Duckling collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param duckingGeneValuesNum Array of gene values number for each gene for Duckling collection.
	 */
	function setDucklingGeneValues(
		uint8[] memory duckingGeneValuesNum
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneValuesNum[0] = duckingGeneValuesNum;
		maxPeculiarity = _calcMaxPeculiarity();
	}

	/**
	 * @notice Set gene distribution types for Duckling collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param ducklingGeneDistrTypes Gene distribution types for Duckling collection.
	 */
	function setDucklingGeneDistributionTypes(
		uint32 ducklingGeneDistrTypes
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneDistributionTypes[0] = ducklingGeneDistrTypes;
		maxPeculiarity = _calcMaxPeculiarity();
	}

	/**
	 * @notice Set gene values number for each gene for Zombeak collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param zombeakGeneValuesNum Array of gene values number for each gene for Duckling collection.
	 */
	function setZombeakGeneValues(
		uint8[] memory zombeakGeneValuesNum
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneValuesNum[1] = zombeakGeneValuesNum;
	}

	/**
	 * @notice Set gene distribution types for Zombeak collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param zombeakGeneDistrTypes Gene distribution types for Zombeak collection.
	 */
	function setZombeakGeneDistributionTypes(
		uint32 zombeakGeneDistrTypes
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneDistributionTypes[1] = zombeakGeneDistrTypes;
	}

	/**
	 * @notice Set number of different Mythic tokens.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param amount Number of different Mythic tokens.
	 */
	function setMythicAmount(uint8 amount) external onlyRole(DEFAULT_ADMIN_ROLE) {
		mythicAmount = amount;
	}

	/**
	 * @notice Set gene values number for each gene for Mythic collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param mythicGeneValuesNum Array of gene values number for each gene for Mythic collection.
	 */
	function setMythicGeneValues(
		uint8[] memory mythicGeneValuesNum
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneValuesNum[2] = mythicGeneValuesNum;
	}

	/**
	 * @notice Set gene distribution types for Mythic collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param mythicGeneDistrTypes Gene distribution types for Mythic collection.
	 */
	function setMythicGeneDistributionTypes(
		uint32 mythicGeneDistrTypes
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneDistributionTypes[2] = mythicGeneDistrTypes;
	}

	// ------- Mint -------

	/**
	 * @notice Mint a pack with `size` of Ducklings. Transfer Duckies from the sender to the TreasureVault.
	 * @dev `Size` must be less than or equal to `MAX_PACK_SIZE`.
	 * @param size Number of Ducklings in the pack.
	 */
	function mintPack(uint8 size) external {
		duckiesContract.transferFrom(msg.sender, treasureVaultAddress, mintPrice * size);
		_mintPackTo(msg.sender, size, true);
	}

	/**
	 * @notice Mint a pack with `amount` of Ducklings to `to` and set transferable flag for each token. Internal function.
	 * @dev `amount` must be less than or equal to `MAX_PACK_SIZE`.
	 * @param to Address to mint the pack to.
	 * @param amount Number of Ducklings in the pack.
	 * @param isTransferable Transferable flag for each token.
	 * @return tokenIds Array of minted token IDs.
	 */
	function _mintPackTo(
		address to,
		uint8 amount,
		bool isTransferable
	) internal returns (uint256[] memory tokenIds) {
		if (amount == 0 || amount > MAX_PACK_SIZE)
			revert MintingRulesViolated(ducklingCollectionId, amount);

		tokenIds = new uint256[](amount);
		uint256[] memory tokenGenomes = new uint256[](amount);

		for (uint256 i = 0; i < amount; i++) {
			tokenGenomes[i] = _generateGenome(ducklingCollectionId).setFlag(
				Genome.FLAG_TRANSFERABLE,
				isTransferable
			);
		}

		tokenIds = ducklingsContract.mintBatchTo(to, tokenGenomes);
	}

	/**
	 * @notice Generate genome for Duckling or Zombeak.
	 * @dev Generate and set all genes from a corresponding collection.
	 * @param collectionId Collection ID.
	 * @return genome Generated genome.
	 */
	function _generateGenome(uint8 collectionId) internal returns (uint256) {
		if (collectionId != ducklingCollectionId && collectionId != zombeakCollectionId) {
			revert MintingRulesViolated(collectionId, 1);
		}

		(bytes3 bitSlice, bytes32 seed) = _shiftSeedSlice(_randomSeed());

		uint256 genome;

		genome = genome.setGene(collectionGeneIdx, collectionId);
		genome = genome.setGene(rarityGeneIdx, _randomWeightedNumber(rarityChances, bitSlice));
		genome = _generateAndSetGenes(genome, collectionId, seed);
		genome = genome.setGene(Genome.MAGIC_NUMBER_GENE_IDX, Genome.BASE_MAGIC_NUMBER);

		return genome;
	}

	/**
	 * @notice Generate and set all genes from a corresponding collection to `genome`.
	 * @dev Generate and set all genes from a corresponding collection to `genome`.
	 * @param genome Genome to set genes to.
	 * @param collectionId Collection ID.
	 * @param seed Random seed to generate genes from.
	 * @return genome Genome with set genes.
	 */
	function _generateAndSetGenes(
		uint256 genome,
		uint8 collectionId,
		bytes32 seed
	) internal view returns (uint256) {
		uint8[] memory geneValuesNum = collectionsGeneValuesNum[collectionId];
		uint32 geneDistributionTypes = collectionsGeneDistributionTypes[collectionId];
		bytes32 newSeed;

		// generate and set each gene
		for (uint8 i = 0; i < geneValuesNum.length; i++) {
			GeneDistributionTypes distrType = _getDistributionType(geneDistributionTypes, i);
			bytes3 bitSlice;
			(bitSlice, newSeed) = _shiftSeedSlice(seed);
			genome = _generateAndSetGene(
				genome,
				generativeGenesOffset + i,
				geneValuesNum[i],
				distrType,
				bitSlice
			);
		}

		// set default values for Ducklings
		if (collectionId == ducklingCollectionId) {
			Rarities rarity = Rarities(genome.getGene(rarityGeneIdx));

			if (rarity == Rarities.Common) {
				genome = genome.setGene(uint8(GenerativeGenes.Body), 0);
				genome = genome.setGene(uint8(GenerativeGenes.Head), 0);
			} else if (rarity == Rarities.Rare) {
				genome = genome.setGene(uint8(GenerativeGenes.Head), 0);
			}
		}

		return genome;
	}

	/**
	 * @notice Generate and set a gene with `geneIdx` to `genome`.
	 * @dev Generate and set a gene with `geneIdx` to `genome`.
	 * @param genome Genome to set a gene to.
	 * @param geneIdx Gene index.
	 * @param geneValuesNum Number of gene values.
	 * @param distrType Gene distribution type.
	 * @param bitSlice Random bit slice to generate a gene from.
	 * @return genome Genome with set gene.
	 */
	function _generateAndSetGene(
		uint256 genome,
		uint8 geneIdx,
		uint8 geneValuesNum,
		GeneDistributionTypes distrType,
		bytes3 bitSlice
	) internal pure returns (uint256) {
		uint8 geneValue;

		if (distrType == GeneDistributionTypes.Even) {
			geneValue = uint8(_max(bitSlice, geneValuesNum));
		} else {
			geneValue = uint8(_generateUnevenGeneValue(geneValuesNum, bitSlice));
		}

		// gene with value 0 means it is a default value, thus this   \/
		genome = genome.setGene(geneIdx, geneValue + 1);

		return genome;
	}

	/**
	 * @notice Generate mythic genome based on melding `genomes`.
	 * @dev Calculates flock peculiarity, and randomizes UniqId corresponding to the peculiarity.
	 * @param genomes Array of genomes to meld into Mythic.
	 * @param seed Random seed to generate mythic genome from.
	 * @return genome Generated Mythic genome.
	 */
	function _generateMythicGenome(
		uint256[] memory genomes,
		bytes32 seed
	) internal view returns (uint256) {
		uint16 sumPeculiarity = 0;
		uint16 maxSumPeculiarity = maxPeculiarity * uint16(genomes.length);

		for (uint8 i = 0; i < genomes.length; i++) {
			sumPeculiarity += _calcPeculiarity(genomes[i]);
		}

		uint16 maxUniqId = mythicAmount - 1;
		uint16 pivotalUniqId = uint16((uint64(sumPeculiarity) * maxUniqId) / maxSumPeculiarity); // multiply and then divide to avoid float numbers
		(uint16 leftEndUniqId, uint16 uniqIdSegmentLength) = _calcUniqIdGenerationParams(
			pivotalUniqId,
			maxUniqId
		);

		(bytes3 bitSlice, bytes32 newSeed) = _shiftSeedSlice(seed);
		uint16 uniqId = leftEndUniqId + uint16(_max(bitSlice, uniqIdSegmentLength));

		uint256 genome;
		genome = genome.setGene(collectionGeneIdx, mythicCollectionId);
		genome = genome.setGene(uint8(MythicGenes.UniqId), uint8(uniqId));
		genome = _generateAndSetGenes(genome, mythicCollectionId, newSeed);
		genome = genome.setGene(Genome.MAGIC_NUMBER_GENE_IDX, Genome.MYTHIC_MAGIC_NUMBER);

		return genome;
	}

	// ------- Meld -------

	/**
	 * @notice Meld tokens with `meldingTokenIds` into a new token. Calls internal function.
	 * @dev Meld tokens with `meldingTokenIds` into a new token.
	 * @param meldingTokenIds Array of token IDs to meld.
	 */
	function meldFlock(uint256[] calldata meldingTokenIds) external {
		// assume all tokens have the same rarity. This is checked later.
		uint256 meldPrice = meldPrices[
			ducklingsContract.getGenome(meldingTokenIds[0]).getGene(rarityGeneIdx)
		];
		duckiesContract.transferFrom(msg.sender, treasureVaultAddress, meldPrice);

		_meldOf(msg.sender, meldingTokenIds, true);
	}

	/**
	 * @notice Meld tokens with `meldingTokenIds` into a new token. Internal function.
	 * @dev Check `owner` is indeed the owner of `meldingTokenIds`. Burn NFTs with `meldingTokenIds`. Transfers Duckies to the TreasureVault.
	 * @param meldingTokenIds Array of token IDs to meld.
	 * @param isTransferable Whether the new token is transferable.
	 * @return meldedTokenId ID of the new token.
	 */
	function _meldOf(
		address owner,
		uint256[] memory meldingTokenIds,
		bool isTransferable
	) internal returns (uint256) {
		if (meldingTokenIds.length != FLOCK_SIZE) revert MeldingRulesViolated(meldingTokenIds);
		if (!ducklingsContract.isOwnerOfBatch(owner, meldingTokenIds))
			revert MeldingRulesViolated(meldingTokenIds);

		uint256[] memory meldingGenomes = ducklingsContract.getGenomes(meldingTokenIds);
		_requireGenomesSatisfyMelding(meldingGenomes);

		ducklingsContract.burnBatch(meldingTokenIds);

		uint256 meldedGenome = _meldGenomes(meldingGenomes).setFlag(
			Genome.FLAG_TRANSFERABLE,
			isTransferable
		);
		uint256 meldedTokenId = ducklingsContract.mintTo(owner, meldedGenome);

		emit Melded(owner, meldingTokenIds, meldedTokenId, block.chainid);

		return meldedTokenId;
	}

	/**
	 * @notice Check that `genomes` satisfy melding rules. Reverts if not.
	 * @dev Check that `genomes` satisfy melding rules. Reverts if not.
	 * @param genomes Array of genomes to check.
	 */
	function _requireGenomesSatisfyMelding(uint256[] memory genomes) internal pure {
		if (
			// equal collections
			!Genome._geneValuesAreEqual(genomes, collectionGeneIdx) ||
			// Rarities must be the same
			!Genome._geneValuesAreEqual(genomes, rarityGeneIdx) ||
			// not Mythic
			genomes[0].getGene(collectionGeneIdx) == mythicCollectionId
		) revert IncorrectGenomesForMelding(genomes);

		Rarities rarity = Rarities(genomes[0].getGene(rarityGeneIdx));
		bool sameColors = Genome._geneValuesAreEqual(genomes, uint8(GenerativeGenes.Color));
		bool sameFamilies = Genome._geneValuesAreEqual(genomes, uint8(GenerativeGenes.Family));
		bool uniqueFamilies = Genome._geneValuesAreUnique(genomes, uint8(GenerativeGenes.Family));

		// specific melding rules
		if (rarity == Rarities.Common) {
			// Common
			if (
				// cards must have the same Color OR the same Family
				!sameColors && !sameFamilies
			) revert IncorrectGenomesForMelding(genomes);
		} else {
			// Rare, Epic
			if (rarity == Rarities.Rare || rarity == Rarities.Epic) {
				if (
					// cards must have the same Color AND the same Family
					!sameColors || !sameFamilies
				) revert IncorrectGenomesForMelding(genomes);
			} else {
				// Legendary
				if (
					// not Legendary Zombeak
					genomes[0].getGene(collectionGeneIdx) == zombeakCollectionId ||
					// cards must have the same Color AND be of each Family
					!sameColors ||
					!uniqueFamilies
				) revert IncorrectGenomesForMelding(genomes);
			}
		}
	}

	/**
	 * @notice Meld `genomes` into a new genome.
	 * @dev Meld `genomes` into a new genome gene by gene. Set the corresponding collection
	 * @param genomes Array of genomes to meld.
	 * @return meldedGenome Melded genome.
	 */
	function _meldGenomes(uint256[] memory genomes) internal returns (uint256) {
		uint8 collectionId = genomes[0].getGene(collectionGeneIdx);
		Rarities rarity = Rarities(genomes[0].getGene(rarityGeneIdx));

		(bytes3 bitSlice, bytes32 seed) = _shiftSeedSlice(_randomSeed());

		// if melding Duckling, they can mutate or evolve into Mythic
		if (collectionId == ducklingCollectionId) {
			if (_isCollectionMutating(rarity, bitSlice)) {
				uint256 zombeakGenome = _generateGenome(zombeakCollectionId);
				return zombeakGenome.setGene(rarityGeneIdx, uint8(rarity));
			}

			if (rarity == Rarities.Legendary) {
				return _generateMythicGenome(genomes, seed);
			}
		}

		uint256 meldedGenome;

		// set the same collection
		meldedGenome = meldedGenome.setGene(collectionGeneIdx, collectionId);
		// increase rarity
		meldedGenome = meldedGenome.setGene(rarityGeneIdx, genomes[0].getGene(rarityGeneIdx) + 1);

		uint8[] memory geneValuesNum = collectionsGeneValuesNum[collectionId];
		uint32 geneDistTypes = collectionsGeneDistributionTypes[collectionId];

		for (uint8 i = 0; i < geneValuesNum.length; i++) {
			(bitSlice, seed) = _shiftSeedSlice(seed);
			uint8 geneValue = _meldGenes(
				genomes,
				generativeGenesOffset + i,
				geneValuesNum[i],
				_getDistributionType(geneDistTypes, i),
				bitSlice
			);
			meldedGenome = meldedGenome.setGene(generativeGenesOffset + i, geneValue);
		}

		// randomize Body for Common and Head for Rare for Ducklings
		if (collectionId == ducklingCollectionId) {
			(bitSlice, seed) = _shiftSeedSlice(seed);
			if (rarity == Rarities.Common) {
				meldedGenome = _generateAndSetGene(
					meldedGenome,
					uint8(GenerativeGenes.Body),
					geneValuesNum[uint8(GenerativeGenes.Body) - generativeGenesOffset],
					GeneDistributionTypes.Uneven,
					bitSlice
				);
			} else if (rarity == Rarities.Rare) {
				meldedGenome = _generateAndSetGene(
					meldedGenome,
					uint8(GenerativeGenes.Head),
					geneValuesNum[uint8(GenerativeGenes.Head) - generativeGenesOffset],
					GeneDistributionTypes.Uneven,
					bitSlice
				);
			}
		}

		meldedGenome = meldedGenome.setGene(Genome.MAGIC_NUMBER_GENE_IDX, Genome.BASE_MAGIC_NUMBER);

		return meldedGenome;
	}

	/**
	 * @notice Randomize if collection is mutating.
	 * @dev Randomize if collection is mutating.
	 * @param rarity Rarity of the collection.
	 * @param bitSlice Bit slice to use for randomization.
	 * @return isMutating True if mutating, false otherwise.
	 */
	function _isCollectionMutating(Rarities rarity, bytes3 bitSlice) internal view returns (bool) {
		// check if mutating chance for this rarity is present
		if (collectionMutationChances.length <= uint8(rarity)) {
			return false;
		}

		uint32 mutationPercentage = collectionMutationChances[uint8(rarity)];
		// dynamic array is needed for `_randomWeightedNumber()`
		uint32[] memory chances = new uint32[](2);
		chances[0] = mutationPercentage;
		chances[1] = 1000 - mutationPercentage; // 1000 as changes are represented in per mil
		return _randomWeightedNumber(chances, bitSlice) == 0;
	}

	/**
	 * @notice Meld `gene` from `genomes` into a new gene value.
	 * @dev Meld `gene` from `genomes` into a new gene value. Gene mutation and inheritance are applied.
	 * @param genomes Array of genomes to meld.
	 * @param gene Gene to be meld.
	 * @param maxGeneValue Max gene value.
	 * @param geneDistrType Gene distribution type.
	 * @param bitSlice Bit slice to use for randomization.
	 * @return geneValue Melded gene value.
	 */
	function _meldGenes(
		uint256[] memory genomes,
		uint8 gene,
		uint8 maxGeneValue,
		GeneDistributionTypes geneDistrType,
		bytes3 bitSlice
	) internal view returns (uint8) {
		// gene mutation
		if (
			geneDistrType == GeneDistributionTypes.Uneven &&
			_randomWeightedNumber(geneMutationChance, bitSlice) == 1
		) {
			uint8 maxPresentGeneValue = Genome._maxGene(genomes, gene);
			return maxPresentGeneValue == maxGeneValue ? maxGeneValue : maxPresentGeneValue + 1;
		}

		// gene inheritance
		uint8 inheritanceIdx = _randomWeightedNumber(geneInheritanceChances, bitSlice);
		return genomes[inheritanceIdx].getGene(gene);
	}

	// ------- Helpers -------

	/**
	 * @notice Get gene distribution type.
	 * @dev Get gene distribution type.
	 * @param distributionTypes Distribution types.
	 * @param idx Index of the gene.
	 * @return Gene distribution type.
	 */
	function _getDistributionType(
		uint32 distributionTypes,
		uint8 idx
	) internal pure returns (GeneDistributionTypes) {
		return
			distributionTypes & (1 << idx) == 0
				? GeneDistributionTypes.Even
				: GeneDistributionTypes.Uneven;
	}

	/**
	 * @notice Generate uneven gene value given the maximum number of values.
	 * @dev Generate uneven gene value using weighed distribution created by squares difference.
	 * @param valuesNum Maximum number of gene values.
	 * @param bitSlice Bit slice to use for randomization.
	 * @return geneValue Gene value.
	 */
	function _generateUnevenGeneValue(
		uint8 valuesNum,
		bytes3 bitSlice
	) internal pure returns (uint8) {
		// number axis is divided into segments by squares of natural numbers (e.g., 1, 4, 9, 16 up to valuesNum ** 2)
		// then a number x is generated in range [0, valuesNum ** 2)
		// the number of segment S number x lies in represents generated gene value
		// however, as larger segments are located further away from 0 at number axis, we need to subtract S from the total number of segments

		uint256 x = _max(bitSlice, uint24(valuesNum) ** 2);
		return valuesNum - uint8(_floorSqrt(x)) - 1;
	}

	function _floorSqrt(uint256 n) internal pure returns (uint256) {
		unchecked {
			if (n > 0) {
				uint256 x = n / 2 + 1;
				uint256 y = (x + n / x) / 2;
				while (x > y) {
					x = y;
					y = (x + n / x) / 2;
				}
				return x;
			}
			return 0;
		}
	}

	/**
	 * @notice Calculate max peculiarity for a current Duckling config.
	 * @dev Sum up number of uneven gene values.
	 * @return maxPeculiarity Max peculiarity.
	 */
	function _calcMaxPeculiarity() internal view returns (uint16) {
		uint16 sum = 0;
		uint32 ducklingDistrTypes = collectionsGeneDistributionTypes[ducklingCollectionId];
		uint8[] memory ducklingGeneValuesNum = collectionsGeneValuesNum[ducklingCollectionId];

		for (uint8 i = 0; i < ducklingGeneValuesNum.length; i++) {
			if (_getDistributionType(ducklingDistrTypes, i) == GeneDistributionTypes.Uneven) {
				// add number of values and not actual values as actual values start with 1, which means number of values and actual values are equal
				sum += ducklingGeneValuesNum[i];
			}
		}

		return sum;
	}

	/**
	 * @notice Calculate peculiarity for a given genome.
	 * @dev Sum up number of uneven gene values.
	 * @param genome Genome.
	 * @return peculiarity Peculiarity.
	 */
	function _calcPeculiarity(uint256 genome) internal view returns (uint16) {
		uint16 sum = 0;
		uint32 ducklingDistrTypes = collectionsGeneDistributionTypes[ducklingCollectionId];

		uint256 genesNum = collectionsGeneValuesNum[ducklingCollectionId].length;
		for (uint8 i = 0; i < genesNum; i++) {
			if (_getDistributionType(ducklingDistrTypes, i) == GeneDistributionTypes.Uneven) {
				// add number of values and not actual values as actual values start with 1, which means number of values and actual values are equal
				sum += genome.getGene(i + generativeGenesOffset);
			}
		}

		return sum;
	}

	/**
	 * @notice Calculate `leftEndUniqId` and `uniqIdSegmentLength` for UniqId generation.
	 * @dev Then UniqId is generated by adding a random number [0, `uniqIdSegmentLength`) to `leftEndUniqId`.
	 * @param pivotalUniqId Pivotal UniqId.
	 * @param maxUniqId Max UniqId.
	 * @return leftEndUniqId Left end of the UniqId segment.
	 * @return uniqIdSegmentLength Length of the UniqId segment.
	 */
	function _calcUniqIdGenerationParams(
		uint16 pivotalUniqId,
		uint16 maxUniqId
	) internal pure returns (uint16 leftEndUniqId, uint16 uniqIdSegmentLength) {
		if (pivotalUniqId < MYTHIC_DISPERSION) {
			// mythic id range overlaps with left dispersion border
			leftEndUniqId = 0;
			uniqIdSegmentLength = pivotalUniqId + MYTHIC_DISPERSION;
		} else if (maxUniqId < pivotalUniqId + MYTHIC_DISPERSION) {
			// mythic id range overlaps with right dispersion border
			leftEndUniqId = pivotalUniqId - MYTHIC_DISPERSION;
			uniqIdSegmentLength = maxUniqId - leftEndUniqId + 1; // +1 to include right border, where the last UniqId is located
		} else {
			// mythic id range does not overlap with dispersion borders
			leftEndUniqId = pivotalUniqId - MYTHIC_DISPERSION;
			uniqIdSegmentLength = 2 * MYTHIC_DISPERSION;
		}
	}
}
