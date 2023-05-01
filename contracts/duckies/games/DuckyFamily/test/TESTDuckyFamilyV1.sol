// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '../DuckyFamilyV1.sol';

/**
 * @title TESTDuckyFamilyV1
 * @notice Contract for testing DuckyFamilyV1 contract. NOT FOR USE IN PRODUCTION.
 */
contract TESTDuckyFamilyV1 is DuckyFamilyV1 {
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

	constructor(
		address duckiesAddress,
		address ducklingsAddress,
		address treasureVaultAddress
	) DuckyFamilyV1(duckiesAddress, ducklingsAddress, treasureVaultAddress) {}

	// allow setting config for better testing

	/**
	 * @notice Sets the rarity chances.
	 * @dev Used for testing.
	 * @param chances Array of rarity chances.
	 */
	function setRarityChances(uint32[] calldata chances) external {
		rarityChances = chances;
	}

	/**
	 * @notice Sets the collection mutation chances.
	 * @dev Used for testing.
	 * @param chances Array of collection mutation chances.
	 */
	function setCollectionMutationChances(uint32[] calldata chances) external {
		collectionMutationChances = chances;
	}

	/**
	 * @notice Sets the gene mutation chances.
	 * @dev Used for testing.
	 * @param chance Chance of gene mutation.
	 */
	function setGeneMutationChance(uint32 chance) external {
		geneMutationChance = [1000 - chance, chance];
	}

	/**
	 * @notice Sets the gene inheritance chances.
	 * @dev Used for testing.
	 * @param chances Array of gene inheritance chances.
	 */
	function setGeneInheritanceChances(uint32[] calldata chances) external {
		geneInheritanceChances = chances;
	}

	// ===============

	/**
	 * @notice Generates a genome. Emits GenomeReturned event.
	 * @dev Exposed for testing.
	 * @param collectionId Collection Id to generate genome for.
	 */
	function generateGenome(uint8 collectionId) external {
		emit GenomeReturned(_generateGenome(collectionId));
	}

	/**
	 * @notice Generates and sets genes to genome. Emits GenomeReturned event.
	 * @dev Exposed for testing.
	 * @param genome Genome to set genes to.
	 * @param collectionId Collection Id to generate genes for.
	 * @param bitSlice Bit slice for randomization.
	 */
	function generateAndSetGenes(uint256 genome, uint8 collectionId, bytes3 bitSlice) external {
		emit GenomeReturned(_generateAndSetGenes(genome, collectionId, bitSlice));
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
		GeneDistributionTypes distrType,
		bytes3 bitSlice
	) external {
		emit GenomeReturned(
			_generateAndSetGene(genome, geneIdx, geneValuesNum, distrType, bitSlice)
		);
	}

	/**
	 * @notice Generates a mythic genome. Emits GenomeReturned event.
	 * @dev Exposed for testing.
	 * @param genomes Genomes to generate mythic genome from.
	 * @param seed Seed for randomization.
	 */
	function generateMythicGenome(uint256[] calldata genomes, bytes32 seed) external {
		emit GenomeReturned(_generateMythicGenome(genomes, seed));
	}

	/**
	 * @notice Checks if genomes satisfy melding, reverting if not.
	 * @dev Exposed for testing.
	 * @param genomes Genomes to check.
	 */
	function requireGenomesSatisfyMelding(uint256[] calldata genomes) external pure {
		_requireGenomesSatisfyMelding(genomes);
	}

	/**
	 * @notice Melds genomes. Emits GenomeReturned event.
	 * @dev Exposed for testing.
	 * @param genomes Genomes to meld.
	 */
	function meldGenomes(uint256[] calldata genomes) external {
		emit GenomeReturned(_meldGenomes(genomes));
	}

	/**
	 * @notice Checks if a collection is mutating. Emits BoolReturned event.
	 * @dev Exposed for testing.
	 * @param rarity Rarity of collection.
	 * @param bitSlice Bit slice for randomization.
	 */
	function isCollectionMutating(Rarities rarity, bytes3 bitSlice) external {
		emit BoolReturned(_isCollectionMutating(rarity, bitSlice));
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
		uint256[] calldata genomes,
		uint8 gene,
		uint8 maxGeneValue,
		GeneDistributionTypes geneDistrType,
		bytes3 bitSlice
	) external {
		emit GeneReturned(_meldGenes(genomes, gene, maxGeneValue, geneDistrType, bitSlice));
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
	) external pure returns (GeneDistributionTypes) {
		return _getDistributionType(distributionTypes, idx);
	}

	/**
	 * @notice Generate uneven gene value.
	 * @dev Exposed for testing. Not pure to measure gas consumption.
	 * @param valuesNum Number of gene values.
	 * @param bitSlice Bit slice for randomization.
	 * @return Uneven gene value.
	 */
	function generateUnevenGeneValue(uint8 valuesNum, bytes3 bitSlice) external returns (uint8) {
		return _generateUnevenGeneValue(valuesNum, bitSlice);
	}

	/**
	 * @notice Calculates square root of `x` rounded down.
	 * @dev Exposed for testing.
	 * @param x Number to calculate square root of.
	 * @return result Square root of `x` rounded down.
	 */
	function floorSqrt(uint256 x) external pure returns (uint256) {
		return _floorSqrt(x);
	}

	/**
	 * @notice Calculate maximum (config) peculiarity.
	 * @dev Exposed for testing.
	 * @return Maximum peculiarity.
	 */
	function calcMaxPeculiarity() external view returns (uint16) {
		return _calcMaxPeculiarity();
	}

	/**
	 * @notice Calculate peculiarity of the genome supplied.
	 * @dev Exposed for testing.
	 * @param genome Genome to calculate peculiarity for.
	 * @return peculiarity Peculiarity.
	 */
	function calcPeculiarity(uint256 genome) external view returns (uint16) {
		return _calcPeculiarity(genome);
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
		uint16 maxUniqId
	) external pure returns (uint16 leftEndUniqId, uint16 uniqIdSegmentLength) {
		return _calcUniqIdGenerationParams(pivotalUniqId, maxUniqId);
	}
}
