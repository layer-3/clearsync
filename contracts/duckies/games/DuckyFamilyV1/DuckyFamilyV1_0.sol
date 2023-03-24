// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol';

import '@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20BurnableUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol';
import '@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol';

import '../../../interfaces/IVoucher.sol';
import '../../../interfaces/IDucklings.sol';
import '../RandomUpgradeable.sol';
import '../Genome.sol';

contract DuckyFamilyV1_0 is
	IVoucher,
	Initializable,
	UUPSUpgradeable,
	AccessControlUpgradeable,
	RandomUpgradeable
{
	using CountersUpgradeable for CountersUpgradeable.Counter;
	using Genome for uint256;
	using ECDSAUpgradeable for bytes32;

	// errors
	error InvalidMintParams(MintParams mintParams);
	error InvalidMeldParams(MeldParams meldParams);

	error MintingRulesViolated(uint8 collectionId, uint8 amount);
	error MeldingRulesViolated(uint256[] tokenIds);
	error IncorrectGenomesForMelding(uint256[] genomes);

	// events
	event Melded(address owner, uint256[] meldingTokenIds, uint256 meldedTokenId, uint256 chainId);

	// roles
	bytes32 public constant UPGRADER_ROLE = keccak256('UPGRADER_ROLE');
	bytes32 public constant MAINTAINER_ROLE = keccak256('MAINTAINER_ROLE');

	// ------- IVoucher -------

	enum VoucherActions {
		MintPack,
		MeldFlock
	}

	struct MintParams {
		address to;
		uint8 size;
	}

	struct MeldParams {
		address owner;
		uint256[] tokenIds;
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

	// Duckling genes: Collection, Rarity, Color, Family, Body, Head, Eyes, Beak, Wings, FirstName, Temper
	// Zombeak genes: Collection, Rarity, Color, Family, Body, Head, Eyes, Beak, Wings

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

	uint8 internal constant collectionGeneIdx = 0;
	uint8 internal constant rarityGeneIdx = 1;
	// general genes start after Collection and Rarity
	uint8 internal constant generativeGenesOffset = 2;

	// number of values for each gene for Duckling and Zombeak collections
	uint8[][2] internal collectionsGeneValuesNum;
	// distribution type of each gene for Duckling and Zombeak collections
	uint32[2] internal collectionsGeneDistributionTypes;

	uint8 internal maxMythicId;

	// chance of a Duckling of a certain rarity to be generated
	uint8[] internal rarityChances; // 70, 20, 5, 1

	// chance of a Duckling of certain rarity to mutate to Zombeak while melding
	uint8[] internal collectionMutationChances; // 10, 5, 2, 0

	uint8[] internal geneMutationChance; // 955, 45 (4.5% to mutate gene value)
	uint8[] internal geneInheritanceChanges; // 5, 4, 3, 2, 1

	// ------- Public values -------

	uint8 public constant MAX_PACK_SIZE = 100;
	uint8 public constant FLOCK_SIZE = 5;

	ERC20BurnableUpgradeable public duckiesContract;
	IDucklings public ducklingsContract;
	address public treasureVaultAddress;

	uint256 public mintPrice;
	uint256 public meldPrice;

	CountersUpgradeable.Counter public nextMythicId;
	CountersUpgradeable.Counter public nextMythicZombeakId;

	// ------- Initializer -------

	function initialize(
		address duckiesAddress,
		address ducklingsAddress,
		address treasureVaultAddress_
	) external initializer {
		__AccessControl_init();
		__UUPSUpgradeable_init();
		__Random_init();

		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(UPGRADER_ROLE, msg.sender);
		_grantRole(MAINTAINER_ROLE, msg.sender);

		duckiesContract = ERC20BurnableUpgradeable(duckiesAddress);
		ducklingsContract = IDucklings(ducklingsAddress);
		treasureVaultAddress = treasureVaultAddress_;

		mintPrice = 5 * 10 ** duckiesContract.decimals();
		meldPrice = 5 * 10 ** duckiesContract.decimals();

		// config
		// duckling
		// TODO: confirm gene values
		collectionsGeneValuesNum[0] = [4, 5, 16, 32, 32, 16, 8, 32, 16, 12, 5, 26];
		collectionsGeneDistributionTypes[0] = 2940; // reverse(001111101101) = 101101111100
		// zombeak
		// TODO: confirm gene values
		collectionsGeneValuesNum[1] = [2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4];
		collectionsGeneDistributionTypes[1] = 2940; // reverse(001111101101) = 101101111100

		// TODO: confirm
		maxMythicId = 64;

		rarityChances = [70, 20, 5, 1];

		collectionMutationChances = [10, 5, 2, 0];
	}

	// -------- Upgrades --------

	function _authorizeUpgrade(
		address newImplementation
	) internal override onlyRole(UPGRADER_ROLE) {}

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

			_mintPackTo(mintParams.to, mintParams.size);
		} else if (voucher.action == uint8(VoucherActions.MeldFlock)) {
			MeldParams memory meldParams = abi.decode(voucher.encodedParams, (MeldParams));

			// meldParams checks
			if (meldParams.owner == address(0) || meldParams.tokenIds.length != FLOCK_SIZE)
				revert InvalidMeldParams(meldParams);

			_meldOf(meldParams.owner, meldParams.tokenIds);
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

	// -------- Price --------

	function setMintPrice(uint256 price) external onlyRole(MAINTAINER_ROLE) {
		mintPrice = price;
	}

	function setMeldPrice(uint256 price) external onlyRole(MAINTAINER_ROLE) {
		meldPrice = price;
	}

	// ------- Mint -------

	function mintPack(uint8 size) external UseRandom {
		duckiesContract.transferFrom(msg.sender, treasureVaultAddress, mintPrice * size);
		_mintPackTo(msg.sender, size);
	}

	function _mintPackTo(address to, uint8 amount) internal {
		if (amount > MAX_PACK_SIZE) revert MintingRulesViolated(ducklingCollectionId, amount);
		for (uint256 i = 0; i < amount; i++) {
			uint256 genome = _generateGenome(ducklingCollectionId);
			ducklingsContract.mintTo(to, genome);
		}
	}

	function _generateGenome(uint8 collectionId) internal returns (uint256) {
		uint256 genome;

		genome = genome.setGene(collectionGeneIdx, collectionId);

		if (collectionId == mythicCollectionId) {
			if (nextMythicId.current() > maxMythicId)
				revert MintingRulesViolated(mythicCollectionId, 1);

			genome = genome.setGene(uint8(MythicGenes.UniqId), uint8(nextMythicId.current()));
			return genome;
		}

		genome = genome.setGene(rarityGeneIdx, uint8(_generateRarity()));
		genome = _generateAndSetGenes(genome, collectionId);

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
			GeneDistributionTypes distrType = _getDistibutionType(geneDistributionTypes, i);
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

	// ------- Meld -------

	function meldFlock(uint256[] calldata meldingTokenIds) external UseRandom {
		duckiesContract.transferFrom(msg.sender, treasureVaultAddress, meldPrice);
		_meldOf(msg.sender, meldingTokenIds);
	}

	function _meldOf(address owner, uint256[] memory meldingTokenIds) internal {
		if (meldingTokenIds.length != FLOCK_SIZE) revert MeldingRulesViolated(meldingTokenIds);
		if (!ducklingsContract.isOwnerOf(owner, meldingTokenIds))
			revert MeldingRulesViolated(meldingTokenIds);

		uint256[] memory meldingGenomes = ducklingsContract.getGenomes(meldingTokenIds);
		_requireGenomesSatisfyMelding(meldingGenomes);

		ducklingsContract.burn(meldingTokenIds);

		uint256 meldedGenome = _meldGenomes(meldingGenomes);
		uint256 meldedTokenId = ducklingsContract.mintTo(owner, meldedGenome);

		emit Melded(owner, meldingTokenIds, meldedTokenId, block.chainid);
	}

	function _requireGenomesSatisfyMelding(uint256[] memory genomes) internal pure {
		if (
			// equal collections
			!Genome._geneValuesAreEqual(genomes, collectionGeneIdx) ||
			// not Mythic
			genomes[0].getGene(collectionGeneIdx) == mythicCollectionId ||
			// Rarities must be the same
			!Genome._geneValuesAreEqual(genomes, rarityGeneIdx)
		) revert IncorrectGenomesForMelding(genomes);

		// specific melding rules
		if (genomes[0].getGene(rarityGeneIdx) == uint8(Rarities.Legendary)) {
			if (
				// not Legendary Zombeak
				genomes[0].getGene(collectionGeneIdx) == zombeakCollectionId ||
				// cards must have the same Color
				!Genome._geneValuesAreEqual(genomes, uint8(GenerativeGenes.Color)) ||
				// cards must be of each Family
				!Genome._geneValuesAreUnique(genomes, uint8(GenerativeGenes.Family))
			) revert IncorrectGenomesForMelding(genomes);
		} else {
			//   Common, Rare, Epic
			if (
				// cards must have the same Color or the same Family
				!Genome._geneValuesAreEqual(genomes, uint8(GenerativeGenes.Color)) &&
				!Genome._geneValuesAreEqual(genomes, uint8(GenerativeGenes.Family))
			) revert IncorrectGenomesForMelding(genomes);
		}
	}

	function _meldGenomes(uint256[] memory genomes) internal returns (uint256) {
		uint8 collectionId = genomes[0].getGene(collectionGeneIdx);
		Rarities rarity = Rarities(genomes[0].getGene(rarityGeneIdx));

		// if melding Duckling, they can mutate or evolve into Mythic
		if (collectionId == ducklingCollectionId) {
			if (_isCollectionMutating(rarity)) {
				return _generateGenome(zombeakCollectionId);
			}

			if (rarity == Rarities.Legendary) {
				nextMythicId.increment();
				return _generateGenome(mythicCollectionId);
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
				_getDistibutionType(geneDistTypes, i)
			);
			meldedGenome = meldedGenome.setGene(generativeGenesOffset + i, geneValue);
		}

		return meldedGenome;
	}

	function _isCollectionMutating(Rarities rarity) internal returns (bool) {
		if (rarity <= Rarities.Epic) {
			uint8 mutationPercentage = collectionMutationChances[uint8(rarity)];
			// dynamic array is needed for `_randomWeighterNumber()`
			uint8[] memory chances;
			chances[0] = mutationPercentage;
			chances[1] = 100 - mutationPercentage;
			return _randomWeightedNumber(chances) == 0;
		} else {
			return false;
		}
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
		uint8 inheritanceIdx = _randomWeightedNumber(geneInheritanceChanges);
		return genomes[inheritanceIdx].getGene(gene);
	}

	// ------- Gene distribution -------

	function _getDistibutionType(
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
}
