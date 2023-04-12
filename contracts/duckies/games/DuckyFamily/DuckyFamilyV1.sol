// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts/access/AccessControl.sol';

import '@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol';
import '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';

import '../../../interfaces/IVoucher.sol';
import '../../../interfaces/IDucklings.sol';
import '../Random.sol';
import '../Genome.sol';

contract DuckyFamilyV1 is IVoucher, AccessControl, Random {
	using Genome for uint256;
	using ECDSA for bytes32;

	// errors
	error InvalidMintParams(MintParams mintParams);
	error InvalidMeldParams(MeldParams meldParams);

	error MintingRulesViolated(uint8 collectionId, uint8 amount);
	error MeldingRulesViolated(uint256[] tokenIds);
	error IncorrectGenomesForMelding(uint256[] genomes);

	// events
	event Melded(address owner, uint256[] meldingTokenIds, uint256 meldedTokenId, uint256 chainId);

	// roles
	bytes32 public constant MAINTAINER_ROLE = keccak256('MAINTAINER_ROLE');

	// ------- IVoucher -------

	enum VoucherActions {
		MintPack,
		MeldFlock
	}

	struct MintParams {
		address to;
		uint8 size;
		bool isTransferable;
	}

	struct MeldParams {
		address owner;
		uint256[] tokenIds;
		bool isTransferable;
	}

	address public issuer;

	// Store the vouchers to avoid replay attacks
	mapping(bytes32 => bool) internal _usedVouchers;

	// ------- Ducklings Game -------

	// for now, Solidity does not support starting value for enum
	// enum Collections {
	// 	Duckling = 0,
	// 	Zombeak,
	// 	Mythic
	// }

	uint8 internal constant ducklingCollectionId = 0;
	uint8 internal constant zombeakCollectionId = 1;
	uint8 internal constant mythicCollectionId = 2;

	uint8 internal constant RARITIES_NUM = 4;

	enum Rarities {
		Common,
		Rare,
		Epic,
		Legendary
	}

	enum GeneDistributionTypes {
		Even,
		Uneven
	}

	enum GenerativeGenes {
		Collection,
		Rarity,
		Color,
		Family,
		Body,
		Head
	}

	enum MythicGenes {
		Collection,
		UniqId
	}

	// ------- Internal values -------

	uint8 public constant MAX_PACK_SIZE = 50;
	uint8 public constant FLOCK_SIZE = 5;

	uint8 internal constant collectionGeneIdx = Genome.COLLECTION_GENE_IDX;
	uint8 internal constant rarityGeneIdx = 1;
	uint8 internal constant flagsGeneIdx = Genome.FLAGS_GENE_IDX;
	// general genes start after Collection and Rarity
	uint8 internal constant generativeGenesOffset = 2;

	// number of values for each gene for Duckling and Zombeak collections
	uint8[][2] internal collectionsGeneValuesNum = [
		// Duckling genes: (Collection, Rarity), Color, Family, Body, Head, Eyes, Beak, Wings, FirstName, Temper, Skill, Habitat, Breed
		[4, 5, 10, 25, 30, 14, 10, 36, 16, 12, 5, 28],
		// Zombeak genes: (Collection, Rarity), Color, Family, Body, Head, Eyes, Beak, Wings, FirstName, Temper, Skill, Habitat, Breed
		[2, 3, 7, 6, 9, 7, 10, 36, 16, 12, 5, 28]
	];
	// distribution type of each gene for Duckling and Zombeak collections
	uint32[2] internal collectionsGeneDistributionTypes = [
		2940, // reverse(001111101101) = 101101111100
		2940 // reverse(001111101101) = 101101111100
	];

	// peculiarity is a sum of uneven gene values for Ducklings
	uint16 internal maxPeculiarity;
	uint8 internal constant MYTHIC_DISPERSION = 5;
	uint8 internal mythicAmount = 59;

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

	// ------- Initializer -------

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

		maxPeculiarity = _calcMaxPeculiarity();
	}

	// ------- Vouchers -------

	function setIssuer(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
		issuer = account;
	}

	function useVouchers(Voucher[] calldata vouchers, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(vouchers), signature, issuer);
		for (uint8 i = 0; i < vouchers.length; i++) {
			_useVoucher(vouchers[i]);
		}
	}

	function useVoucher(Voucher calldata voucher, bytes calldata signature) external {
		_requireCorrectSigner(abi.encode(voucher), signature, issuer);
		_useVoucher(voucher);
	}

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

	function _requireCorrectSigner(
		bytes memory encodedData,
		bytes memory signature,
		address signer
	) internal pure {
		address actualSigner = keccak256(encodedData).toEthSignedMessageHash().recover(signature);
		if (actualSigner != signer) revert IncorrectSigner(signer, actualSigner);
	}

	// -------- Config --------

	function setMintPrice(uint256 price) external onlyRole(MAINTAINER_ROLE) {
		mintPrice = price * 10 ** duckiesContract.decimals();
	}

	function getMeldPrices() external view returns (uint256[RARITIES_NUM] memory) {
		return meldPrices;
	}

	function setMeldPrices(
		uint256[RARITIES_NUM] calldata prices
	) external onlyRole(MAINTAINER_ROLE) {
		for (uint8 i = 0; i < RARITIES_NUM; i++) {
			meldPrices[i] = prices[i] * 10 ** duckiesContract.decimals();
		}
	}

	function getCollectionsGeneValues() external view returns (uint8[][2] memory, uint8) {
		return (collectionsGeneValuesNum, mythicAmount);
	}

	function getCollectionsGeneDistributionTypes() external view returns (uint32[2] memory) {
		return collectionsGeneDistributionTypes;
	}

	function setDucklingGeneValues(
		uint8[] memory duckingGeneValuesNum
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneValuesNum[0] = duckingGeneValuesNum;
		maxPeculiarity = _calcMaxPeculiarity();
	}

	function setDucklingGeneDistributionTypes(
		uint32 ducklingGeneDistrTypes
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneDistributionTypes[0] = ducklingGeneDistrTypes;
		maxPeculiarity = _calcMaxPeculiarity();
	}

	function setZombeakGeneValues(
		uint8[] memory zombeaksGeneValuesNum
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneValuesNum[1] = zombeaksGeneValuesNum;
	}

	function setZombeakGeneDistributionTypes(
		uint32 zombeakGeneDistrTypes
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneDistributionTypes[1] = zombeakGeneDistrTypes;
	}

	function setMythicAmount(uint8 amount) external onlyRole(DEFAULT_ADMIN_ROLE) {
		mythicAmount = amount;
	}

	// ------- Mint -------

	function mintPack(uint8 size) external UseRandom {
		duckiesContract.transferFrom(msg.sender, treasureVaultAddress, mintPrice * size);
		_mintPackTo(msg.sender, size, true);
	}

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

	function _generateGenome(uint8 collectionId) internal returns (uint256) {
		if (collectionId != ducklingCollectionId && collectionId != zombeakCollectionId) {
			revert MintingRulesViolated(collectionId, 1);
		}

		uint256 genome;

		genome = genome.setGene(collectionGeneIdx, collectionId);
		genome = genome.setGene(rarityGeneIdx, uint8(_generateRarity()));
		genome = _generateAndSetGenes(genome, collectionId);
		genome = genome.setGene(Genome.MAGIC_NUMBER_GENE_IDX, Genome.BASE_MAGIC_NUMBER);

		return genome;
	}

	function _generateRarity() internal returns (Rarities) {
		return Rarities(_randomWeightedNumber(rarityChances));
	}

	function _generateAndSetGenes(uint256 genome, uint8 collectionId) internal returns (uint256) {
		uint8[] memory geneValuesNum = collectionsGeneValuesNum[collectionId];
		uint32 geneDistributionTypes = collectionsGeneDistributionTypes[collectionId];

		// generate and set each gene
		for (uint8 i = 0; i < geneValuesNum.length; i++) {
			GeneDistributionTypes distrType = _getDistributionType(geneDistributionTypes, i);
			uint8 geneValue;

			if (distrType == GeneDistributionTypes.Even) {
				geneValue = uint8(_randomMaxNumber(geneValuesNum[i]));
			} else {
				geneValue = uint8(_generateUnevenGeneValue(geneValuesNum[i]));
			}

			// gene with value 0 means it is a default value, thus this   \/
			genome = genome.setGene(generativeGenesOffset + i, geneValue + 1);
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

	function _generateMythicGenome(uint256[] memory genomes) internal returns (uint256) {
		uint16 sumPeculiarity = 0;
		uint16 maxSumPeculiarity = maxPeculiarity * uint16(genomes.length);

		for (uint8 i = 0; i < genomes.length; i++) {
			sumPeculiarity += _calcPeculiarity(genomes[i]);
		}

		uint16 maxUniqId = mythicAmount - 1;
		uint16 pivotalUniqId = uint16((uint64(sumPeculiarity) * maxUniqId) / maxSumPeculiarity); // multiply and then divide to avoid float numbers
		uint16 leftEndUniqId;
		uint16 uniqIdSegmentLength;

		if (pivotalUniqId < MYTHIC_DISPERSION) {
			// mythic id range overlaps with left dispersion border
			leftEndUniqId = 0;
			uniqIdSegmentLength = pivotalUniqId + MYTHIC_DISPERSION;
		} else if (maxUniqId < pivotalUniqId + MYTHIC_DISPERSION) {
			// mythic id range overlaps with right dispersion border
			leftEndUniqId = pivotalUniqId - MYTHIC_DISPERSION;
			uniqIdSegmentLength = maxUniqId - pivotalUniqId + MYTHIC_DISPERSION;
		} else {
			// mythic id range does not overlap with dispersion borders
			leftEndUniqId = pivotalUniqId - MYTHIC_DISPERSION;
			uniqIdSegmentLength = 2 * MYTHIC_DISPERSION;
		}

		uint16 uniqId = leftEndUniqId + uint16(Random._randomMaxNumber(uniqIdSegmentLength));

		uint256 genome;
		genome = genome.setGene(collectionGeneIdx, mythicCollectionId);
		genome = genome.setGene(uint8(MythicGenes.UniqId), uint8(uniqId));
		genome = genome.setGene(Genome.MAGIC_NUMBER_GENE_IDX, Genome.MYTHIC_MAGIC_NUMBER);

		return genome;
	}

	// ------- Meld -------

	function meldFlock(uint256[] calldata meldingTokenIds) external UseRandom {
		// assume all tokens have the same rarity. This is checked later.
		uint256 meldPrice = meldPrices[
			ducklingsContract.getGenome(meldingTokenIds[0]).getGene(rarityGeneIdx)
		];
		duckiesContract.transferFrom(msg.sender, treasureVaultAddress, meldPrice);

		_meldOf(msg.sender, meldingTokenIds, true);
	}

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

	function _meldGenomes(uint256[] memory genomes) internal returns (uint256) {
		uint8 collectionId = genomes[0].getGene(collectionGeneIdx);
		Rarities rarity = Rarities(genomes[0].getGene(rarityGeneIdx));

		// if melding Duckling, they can mutate or evolve into Mythic
		if (collectionId == ducklingCollectionId) {
			if (_isCollectionMutating(rarity)) {
				uint256 zombeakGenome = _generateGenome(zombeakCollectionId);
				return zombeakGenome.setGene(rarityGeneIdx, uint8(rarity));
			}

			if (rarity == Rarities.Legendary) {
				return _generateMythicGenome(genomes);
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
			uint8 geneValue = _meldGenes(
				genomes,
				generativeGenesOffset + i,
				geneValuesNum[i],
				_getDistributionType(geneDistTypes, i)
			);
			meldedGenome = meldedGenome.setGene(generativeGenesOffset + i, geneValue);
		}

		meldedGenome = meldedGenome.setGene(Genome.MAGIC_NUMBER_GENE_IDX, Genome.BASE_MAGIC_NUMBER);

		return meldedGenome;
	}

	function _isCollectionMutating(Rarities rarity) internal returns (bool) {
		// check if mutating chance for this rarity is present
		if (collectionMutationChances.length <= uint8(rarity)) {
			return false;
		}

		uint32 mutationPercentage = collectionMutationChances[uint8(rarity)];
		// dynamic array is needed for `_randomWeightedNumber()`
		uint32[] memory chances = new uint32[](2);
		chances[0] = mutationPercentage;
		chances[1] = 1000 - mutationPercentage; // 1000 as changes are represented in per mil
		return _randomWeightedNumber(chances) == 0;
	}

	function _meldGenes(
		uint256[] memory genomes,
		uint8 gene,
		uint8 maxGeneValue,
		GeneDistributionTypes geneDistrType
	) internal returns (uint8) {
		// gene mutation
		if (
			geneDistrType == GeneDistributionTypes.Uneven &&
			_randomWeightedNumber(geneMutationChance) == 1
		) {
			uint8 maxPresentGeneValue = Genome._maxGene(genomes, gene);
			return maxPresentGeneValue == maxGeneValue ? maxGeneValue : maxPresentGeneValue + 1;
		}

		// gene inheritance
		uint8 inheritanceIdx = _randomWeightedNumber(geneInheritanceChances);
		return genomes[inheritanceIdx].getGene(gene);
	}

	// ------- Gene distribution -------

	function _getDistributionType(
		uint32 distributionTypes,
		uint8 idx
	) internal pure returns (GeneDistributionTypes) {
		return
			distributionTypes & (1 << idx) == 0
				? GeneDistributionTypes.Even
				: GeneDistributionTypes.Uneven;
	}

	function _generateUnevenGeneValue(uint8 valuesNum) internal returns (uint8) {
		// using quadratic algorithm
		// chance of each gene value to be generated is (N - v)^2 / S
		// N - number of gene values, v - gene value, S - sum of Squared gene values

		uint256 N = uint256(valuesNum);
		uint256 S = (N * (N + 1) * (2 * N + 1)) / 6;
		uint256 num = _randomMaxNumber(S);
		uint256 accumNum = 0;

		for (uint8 i = 0; i < N; i++) {
			accumNum += (N - i) ** 2;

			if (num < accumNum) {
				return i;
			}
		}

		// code execution should never reach this
		return 0;
	}

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

	function _calcPeculiarity(uint256 genome) internal view returns (uint16) {
		uint16 sum = 0;
		uint32 ducklingDistrTypes = collectionsGeneDistributionTypes[ducklingCollectionId];

		for (uint8 i = 0; i < collectionsGeneValuesNum[ducklingCollectionId].length; i++) {
			if (_getDistributionType(ducklingDistrTypes, i) == GeneDistributionTypes.Uneven) {
				// add number of values and not actual values as actual values start with 1, which means number of values and actual values are equal
				sum += genome.getGene(i + generativeGenesOffset);
			}
		}

		return sum;
	}
}
