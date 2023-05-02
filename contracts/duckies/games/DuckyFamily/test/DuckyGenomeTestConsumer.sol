// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '../../../../interfaces/IDuckyFamily.sol';
import '../DuckyGenome.sol';

contract DuckyGenomeTestConsumer {
	/**
	 * @notice Event emitted when a genome is generated.
	 * @dev Used for testing.
	 * @param genome Generated genome.
	 */
	event GenomeReturned(uint256 genome);

	/**
	 * @notice Event emitted when a gene is generated.
	 * @dev Used for testing.
	 * @param gene Generated gene.
	 */
	event GeneReturned(uint8 gene);

	/**
	 * @notice Event emitted when a bool is returned.
	 * @dev Used for testing.
	 * @param returnedBool Returned bool.
	 */
	event BoolReturned(bool returnedBool);

	/**
	 * @notice Event emitted when a uint8 is returned.
	 * @dev Used for testing.
	 * @param returnedUint8 Returned uint8.
	 */
	event Uint8Returned(uint8 returnedUint8);

	/**
	 * @notice Generates and sets genes to genome. Emits GenomeReturned event.
	 * @dev Exposed for testing.
	 * @param genome Genome to set genes to.
	 * @param collectionId Collection Id to generate genes for.
	 * @param seed Seed for randomization.
	 */
	function generateAndSetGenes(
		uint256 genome,
		uint8 collectionId,
		uint8[] memory geneValues,
		uint32 geneDistributionTypes,
		bytes32 seed
	) external {
		emit GenomeReturned(
			DuckyGenome._generateAndSetGenes(
				genome,
				collectionId,
				geneValues,
				geneDistributionTypes,
				seed
			)
		);
	}

	/**
	 * @notice Generates and sets gene to genome. Emits GenomeReturned event.
	 * @dev Exposed for testing.
	 * @param genome Genome to set gene to.
	 * @param geneIdx Index of gene to set.
	 * @param geneValuesNum Number of gene values.
	 * @param distrType Gene distribution type.
	 * @param bitSlice Bit slice for randomization.
	 */
	function generateAndSetGene(
		uint256 genome,
		uint8 geneIdx,
		uint8 geneValuesNum,
		IDuckyFamily.GeneDistributionTypes distrType,
		bytes3 bitSlice
	) external {
		emit GenomeReturned(
			DuckyGenome._generateAndSetGene(genome, geneIdx, geneValuesNum, distrType, bitSlice)
		);
	}

	/**
	 * @notice Checks if genomes satisfy melding, reverting if not.
	 * @dev Exposed for testing.
	 * @param genomes Genomes to check.
	 */
	function requireGenomesSatisfyMelding(uint256[] calldata genomes) external pure {
		DuckyGenome._requireGenomesSatisfyMelding(genomes);
	}

	/**
	 * @notice Checks if a collection is mutating. Emits BoolReturned event.
	 * @dev Exposed for testing.
	 * @param rarity Rarity of collection.
	 * @param bitSlice Bit slice for randomization.
	 */
	function isCollectionMutating(
		IDuckyFamily.Rarities rarity,
		uint32[] memory mutationChances,
		bytes3 bitSlice
	) external {
		emit BoolReturned(DuckyGenome._isCollectionMutating(rarity, mutationChances, bitSlice));
	}

	/**
	 * @notice Melds genes. Emits GeneReturned event.
	 * @dev Exposed for testing.
	 * @param genomes Genomes to meld genes from.
	 * @param gene Gene to meld.
	 * @param maxGeneValue Max gene value.
	 * @param geneDistrType Gene distribution type.
	 * @param bitSlice Bit slice for randomization.
	 */
	function meldGenes(
		uint256[] memory genomes,
		uint8 gene,
		uint8 maxGeneValue,
		IDuckyFamily.GeneDistributionTypes geneDistrType,
		uint32[] memory mutationChance,
		uint32[] memory inheritanceChances,
		bytes3 bitSlice
	) external {
		emit GeneReturned(
			DuckyGenome._meldGenes(
				genomes,
				gene,
				maxGeneValue,
				geneDistrType,
				mutationChance,
				inheritanceChances,
				bitSlice
			)
		);
	}

	/**
	 * @notice Get gene distribution type.
	 * @dev Exposed for testing.
	 * @param distributionTypes Distribution types.
	 * @param idx Index of the gene.
	 * @return Gene distribution type.
	 */
	function getDistributionType(
		uint32 distributionTypes,
		uint8 idx
	) external pure returns (IDuckyFamily.GeneDistributionTypes) {
		return DuckyGenome._getDistributionType(distributionTypes, idx);
	}

	/**
	 * @notice Generate uneven gene value. Emits Uint8Returned event.
	 * @dev Exposed for testing. Not pure to measure gas consumption.
	 * @param valuesNum Number of gene values.
	 * @param bitSlice Bit slice for randomization.
	 */
	function generateUnevenGeneValue(uint8 valuesNum, bytes3 bitSlice) external {
		emit Uint8Returned(DuckyGenome._generateUnevenGeneValue(valuesNum, bitSlice));
	}

	/**
	 * @notice Calculate maximum (config) peculiarity.
	 * @dev Exposed for testing.
	 * @return Maximum peculiarity.
	 */
	function _calcConfigPeculiarity(
		uint8[] memory geneValuesNum,
		uint32 geneDistrTypes
	) external pure returns (uint16) {
		return DuckyGenome._calcConfigPeculiarity(geneValuesNum, geneDistrTypes);
	}

	/**
	 * @notice Calculate peculiarity of the genome supplied.
	 * @dev Exposed for testing.
	 * @param genome Genome to calculate peculiarity for.
	 * @return peculiarity Peculiarity.
	 */
	function calcPeculiarity(
		uint256 genome,
		uint8 genesNum,
		uint32 geneDistrTypes
	) external pure returns (uint16) {
		return DuckyGenome._calcPeculiarity(genome, genesNum, geneDistrTypes);
	}

	/**
	 * @notice Calculate `leftEndUniqId` and `uniqIdSegmentLength` for UniqId generation.
	 * @dev Exposed for testing. Then UniqId is generated by adding a random number [0, `uniqIdSegmentLength`) to `leftEndUniqId`.
	 * @param pivotalUniqId Pivotal UniqId.
	 * @param maxUniqId Max UniqId.
	 * @return leftEndUniqId Left end of the UniqId segment.
	 * @return uniqIdSegmentLength Length of the UniqId segment.
	 */
	function calcUniqIdGenerationParams(
		uint16 pivotalUniqId,
		uint16 maxUniqId,
		uint16 mythicDispersion
	) external pure returns (uint16 leftEndUniqId, uint16 uniqIdSegmentLength) {
		return DuckyGenome._calcUniqIdGenerationParams(pivotalUniqId, maxUniqId, mythicDispersion);
	}
}
