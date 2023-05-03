// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts/utils/math/Math.sol';

import '../../../interfaces/IDuckyFamily.sol';
import '../Genome.sol';
import '../Utils.sol';

/**
 * @title DuckyGenome
 * @notice Library for generating Duckies genomes.
 * @dev Contains functions for generating Duckies genomes.
 */
library DuckyGenome {
	using Genome for uint256;

	/// @dev constants must be duplicated both here and in DuckyFamilyV1 as Solidity does not see Library constants as constants, see https://github.com/ethereum/solidity/issues/12248
	uint8 internal constant ducklingCollectionId = 0;
	uint8 internal constant zombeakCollectionId = 1;
	uint8 internal constant mythicCollectionId = 2;
	uint8 internal constant RARITIES_NUM = 4;

	uint8 internal constant collectionGeneIdx = Genome.COLLECTION_GENE_IDX;
	uint8 internal constant rarityGeneIdx = 1;
	uint8 internal constant flagsGeneIdx = Genome.FLAGS_GENE_IDX;
	uint8 internal constant generativeGenesOffset = 2;

	/**
	 * @notice Generates and sets genes to genome. Emits GenomeReturned event.
	 * @dev Generates and sets genes to genome. Emits GenomeReturned event.
	 * @param genome Genome to set genes to.
	 * @param collectionId Collection Id to generate genes for.
	 * @param geneValuesNum Number of gene values for each gene.
	 * @param geneDistributionTypes Gene distribution types.
	 * @param seed Seed for randomization.
	 */
	function generateAndSetGenes(
		uint256 genome,
		uint8 collectionId,
		uint8[] memory geneValuesNum,
		uint32 geneDistributionTypes,
		bytes32 seed
	) internal pure returns (uint256) {
		uint8 genesNum = uint8(geneValuesNum.length);
		bytes32 newSeed;

		// generate and set each gene
		for (uint8 i = 0; i < genesNum; i++) {
			IDuckyFamily.GeneDistributionTypes distrType = getDistributionType(
				geneDistributionTypes,
				i
			);
			bytes3 bitSlice;
			(bitSlice, newSeed) = Utils.shiftSeedSlice(seed);
			genome = generateAndSetGene(
				genome,
				generativeGenesOffset + i,
				geneValuesNum[i],
				distrType,
				bitSlice
			);
		}

		// set default values for Ducklings
		if (collectionId == ducklingCollectionId) {
			IDuckyFamily.Rarities rarity = IDuckyFamily.Rarities(genome.getGene(rarityGeneIdx));

			if (rarity == IDuckyFamily.Rarities.Common) {
				genome = genome.setGene(uint8(IDuckyFamily.GenerativeGenes.Body), 0);
				genome = genome.setGene(uint8(IDuckyFamily.GenerativeGenes.Head), 0);
			} else if (rarity == IDuckyFamily.Rarities.Rare) {
				genome = genome.setGene(uint8(IDuckyFamily.GenerativeGenes.Head), 0);
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
	function generateAndSetGene(
		uint256 genome,
		uint8 geneIdx,
		uint8 geneValuesNum,
		IDuckyFamily.GeneDistributionTypes distrType,
		bytes3 bitSlice
	) internal pure returns (uint256) {
		uint8 geneValue;

		if (distrType == IDuckyFamily.GeneDistributionTypes.Even) {
			geneValue = uint8(Utils.randomNumber(bitSlice, geneValuesNum));
		} else {
			geneValue = uint8(generateUnevenGeneValue(geneValuesNum, bitSlice));
		}

		// gene with value 0 means it is a default value, thus this   \/
		genome = genome.setGene(geneIdx, geneValue + 1);

		return genome;
	}

	/**
	 * @notice Check that `genomes` satisfy melding rules. Reverts if not.
	 * @dev Check that `genomes` satisfy melding rules. Reverts if not.
	 * @param genomes Array of genomes to check.
	 */
	function requireGenomesSatisfyMelding(uint256[] memory genomes) internal pure {
		if (
			// equal collections
			!Genome.geneValuesAreEqual(genomes, collectionGeneIdx) ||
			// Rarities must be the same
			!Genome.geneValuesAreEqual(genomes, rarityGeneIdx) ||
			// not Mythic
			genomes[0].getGene(collectionGeneIdx) == mythicCollectionId
		) revert IDuckyFamily.IncorrectGenomesForMelding(genomes);

		IDuckyFamily.Rarities rarity = IDuckyFamily.Rarities(genomes[0].getGene(rarityGeneIdx));
		bool sameColors = Genome.geneValuesAreEqual(
			genomes,
			uint8(IDuckyFamily.GenerativeGenes.Color)
		);
		bool sameFamilies = Genome.geneValuesAreEqual(
			genomes,
			uint8(IDuckyFamily.GenerativeGenes.Family)
		);
		bool uniqueFamilies = Genome.geneValuesAreUnique(
			genomes,
			uint8(IDuckyFamily.GenerativeGenes.Family)
		);

		// specific melding rules
		if (rarity == IDuckyFamily.Rarities.Common) {
			// Common
			if (
				// cards must have the same Color OR the same Family
				!sameColors && !sameFamilies
			) revert IDuckyFamily.IncorrectGenomesForMelding(genomes);
		} else {
			// Rare, Epic
			if (rarity == IDuckyFamily.Rarities.Rare || rarity == IDuckyFamily.Rarities.Epic) {
				if (
					// cards must have the same Color AND the same Family
					!sameColors || !sameFamilies
				) revert IDuckyFamily.IncorrectGenomesForMelding(genomes);
			} else {
				// Legendary
				if (
					// not Legendary Zombeak
					genomes[0].getGene(collectionGeneIdx) == zombeakCollectionId ||
					// cards must have the same Color AND be of each Family
					!sameColors ||
					!uniqueFamilies
				) revert IDuckyFamily.IncorrectGenomesForMelding(genomes);
			}
		}
	}

	/**
	 * @notice Randomize if collection is mutating.
	 * @dev Randomize if collection is mutating.
	 * @param rarity Rarity of the collection.
	 * @param mutationChances Array of mutation chances for each rarity.
	 * @param bitSlice Bit slice to use for randomization.
	 * @return isMutating True if mutating, false otherwise.
	 */
	function isCollectionMutating(
		IDuckyFamily.Rarities rarity,
		uint32[] memory mutationChances,
		bytes3 bitSlice
	) internal pure returns (bool) {
		// check if mutating chance for this rarity is present
		if (mutationChances.length <= uint8(rarity)) {
			return false;
		}

		uint32 mutationPercentage = mutationChances[uint8(rarity)];
		// dynamic array is needed for `randomWeightedNumber()`
		uint32[] memory chances = new uint32[](2);
		chances[0] = mutationPercentage;
		chances[1] = 1000 - mutationPercentage; // 1000 as changes are represented in per mil
		return Utils.randomWeightedNumber(chances, bitSlice) == 0;
	}

	/**
	 * @notice Meld `gene` from `genomes` into a new gene value.
	 * @dev Meld `gene` from `genomes` into a new gene value. Gene mutation and inheritance are applied.
	 * @param genomes Array of genomes to meld.
	 * @param gene Gene to be meld.
	 * @param maxGeneValue Max gene value.
	 * @param geneDistrType Gene distribution type.
	 * @param mutationChance Mutation chance. Represented as [chance of no mutation, chance of mutation] in per mil.
	 * @param inheritanceChances Array of inheritance chances for each rarity.
	 * @param bitSlice Bit slice to use for randomization.
	 * @return geneValue Melded gene value.
	 */
	function meldGenes(
		uint256[] memory genomes,
		uint8 gene,
		uint8 maxGeneValue,
		IDuckyFamily.GeneDistributionTypes geneDistrType,
		uint32[] memory mutationChance,
		uint32[] memory inheritanceChances,
		bytes3 bitSlice
	) internal pure returns (uint8) {
		// gene mutation
		if (
			geneDistrType == IDuckyFamily.GeneDistributionTypes.Uneven &&
			Utils.randomWeightedNumber(mutationChance, bitSlice) == 1
		) {
			uint8 maxPresentGeneValue = Genome.maxGene(genomes, gene);
			return maxPresentGeneValue == maxGeneValue ? maxGeneValue : maxPresentGeneValue + 1;
		}

		// gene inheritance
		uint8 inheritanceIdx = Utils.randomWeightedNumber(inheritanceChances, bitSlice);
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
	function getDistributionType(
		uint32 distributionTypes,
		uint8 idx
	) internal pure returns (IDuckyFamily.GeneDistributionTypes) {
		return
			distributionTypes & (1 << idx) == 0
				? IDuckyFamily.GeneDistributionTypes.Even
				: IDuckyFamily.GeneDistributionTypes.Uneven;
	}

	/**
	 * @notice Generate uneven gene value given the maximum number of values.
	 * @dev Generate uneven gene value using reciprocal distribution described below.
	 * @param valuesNum Maximum number of gene values.
	 * @param bitSlice Bit slice to use for randomization.
	 * @return geneValue Gene value.
	 */
	function generateUnevenGeneValue(
		uint8 valuesNum,
		bytes3 bitSlice
	) internal pure returns (uint8) {
		// using reciprocal distribution
		// gene value is selected as ceil[(2N/(x+1))-N],
		// where x is random number between 0 and 1
		// Because of shape of reciprocal graph,
		// evenly distributed x values will result in unevenly distributed y values.

		// N - number of gene values
		uint256 N = uint256(valuesNum);
		// Generates number from 1 to 10^6
		uint256 x = 1 + Utils.randomNumber(bitSlice, 1_000_000);
		// Calculates uneven distributed y, value of y is between 0 and N
		uint256 y = (2 * N * 1_000) / (Math.sqrt(x) + 1_000) - N;
		return uint8(y);
	}

	/**
	 * @notice Calculate max peculiarity for a supplied config.
	 * @dev Sum up number of uneven gene values.
	 * @param geneValuesNum Array of number of gene values for each gene.
	 * @param geneDistrTypes Gene distribution types.
	 * @return maxPeculiarity Max peculiarity.
	 */
	function calcConfigPeculiarity(
		uint8[] memory geneValuesNum,
		uint32 geneDistrTypes
	) internal pure returns (uint16) {
		uint16 sum = 0;

		uint8 genesNum = uint8(geneValuesNum.length);
		for (uint8 i = 0; i < genesNum; i++) {
			if (
				getDistributionType(geneDistrTypes, i) == IDuckyFamily.GeneDistributionTypes.Uneven
			) {
				// add number of values and not actual values as actual values start with 1, which means number of values and actual values are equal
				sum += geneValuesNum[i];
			}
		}

		return sum;
	}

	/**
	 * @notice Calculate peculiarity for a given genome.
	 * @dev Sum up number of uneven gene values.
	 * @param genome Genome.
	 * @param genesNum Number of genes.
	 * @param geneDistrTypes Gene distribution types.
	 * @return peculiarity Peculiarity.
	 */
	function calcPeculiarity(
		uint256 genome,
		uint8 genesNum,
		uint32 geneDistrTypes
	) internal pure returns (uint16) {
		uint16 sum = 0;

		for (uint8 i = 0; i < genesNum; i++) {
			if (
				getDistributionType(geneDistrTypes, i) == IDuckyFamily.GeneDistributionTypes.Uneven
			) {
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
	 * @param mythicDispersion Half of the segment length in which mythic UniqIds are generated.
	 * @return leftEndUniqId Left end of the UniqId segment.
	 * @return uniqIdSegmentLength Length of the UniqId segment.
	 */
	function calcUniqIdGenerationParams(
		uint16 pivotalUniqId,
		uint16 maxUniqId,
		uint16 mythicDispersion
	) internal pure returns (uint16 leftEndUniqId, uint16 uniqIdSegmentLength) {
		if (pivotalUniqId < mythicDispersion) {
			// mythic id range overlaps with left dispersion border
			leftEndUniqId = 0;
			uniqIdSegmentLength = pivotalUniqId + mythicDispersion;
		} else if (maxUniqId < pivotalUniqId + mythicDispersion) {
			// mythic id range overlaps with right dispersion border
			leftEndUniqId = pivotalUniqId - mythicDispersion;
			uniqIdSegmentLength = maxUniqId - leftEndUniqId + 1; // +1 to include right border, where the last UniqId is located
		} else {
			// mythic id range does not overlap with dispersion borders
			leftEndUniqId = pivotalUniqId - mythicDispersion;
			uniqIdSegmentLength = 2 * mythicDispersion;
		}
	}
}
