// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '../Genome.sol';

/**
 * @title GenomeTestConsumer
 * @notice Contract for testing Genome contract. NOT FOR USE IN PRODUCTION.
 */
contract GenomeTestConsumer {
	/**
	 * @notice Read flags gene from genome.
	 * @dev Expose for testing.
	 * @param genome Genome to get flags gene from.
	 * @return flags Flags gene.
	 */
	function getFlags(uint256 genome) external pure returns (uint8) {
		return Genome.getFlags(genome);
	}

	/**
	 * @notice Read specific bit mask flag from genome.
	 * @dev Expose for testing.
	 * @param genome Genome to read flag from.
	 * @param flag Bit mask flag to read.
	 * @return value Value of the flag.
	 */
	function getFlag(uint256 genome, uint8 flag) external pure returns (bool) {
		return Genome.getFlag(genome, flag);
	}

	/**
	 * @notice Set specific bit mask flag in genome.
	 * @dev Expose for testing.
	 * @param genome Genome to set flag in.
	 * @param flag Bit mask flag to set.
	 * @param value Value of the flag.
	 * @return genome Genome with the flag set.
	 */
	function setFlag(uint256 genome, uint8 flag, bool value) external pure returns (uint256) {
		return Genome.setFlag(genome, flag, value);
	}

	/**
	 * @notice Set `value` to `gene` in genome.
	 * @dev Expose for testing.
	 * @param genome Genome to set gene in.
	 * @param gene Gene to set.
	 * @param value Value to set.
	 * @return genome Genome with the gene set.
	 */
	function setGene(uint256 genome, uint8 gene, uint8 value) external pure returns (uint256) {
		return Genome.setGene(genome, gene, value);
	}

	/**
	 * @notice Get `gene` value from genome.
	 * @dev Expose for testing.
	 * @param genome Genome to get gene from.
	 * @param gene Gene to get.
	 * @return geneValue Gene value.
	 */
	function getGene(uint256 genome, uint8 gene) external pure returns (uint8) {
		return Genome.getGene(genome, gene);
	}

	/**
	 * @notice Get largest value of a `gene` in `genomes`.
	 * @dev Expose for testing.
	 * @param genomes Genomes to get gene from.
	 * @param gene Gene to get.
	 * @return maxValue Largest value of a `gene` in `genomes`.
	 */
	function maxGene(uint256[] memory genomes, uint8 gene) external pure returns (uint8) {
		return Genome.maxGene(genomes, gene);
	}

	/**
	 * @notice Check if values of `gene` in `genomes` are equal.
	 * @dev Expose for testing.
	 * @param genomes Genomes to check.
	 * @param gene Gene to check.
	 * @return isEqual True if values of `gene` in `genomes` are equal, false otherwise.
	 */
	function geneValuesAreEqual(uint256[] memory genomes, uint8 gene) external pure returns (bool) {
		return Genome.geneValuesAreEqual(genomes, gene);
	}

	/**
	 * @notice Check if values of `gene` in `genomes` are unique.
	 * @dev Expose for testing.
	 * @param genomes Genomes to check.
	 * @param gene Gene to check.
	 * @return isUnique True if values of `gene` in `genomes` are unique, false otherwise.
	 */
	function geneValuesAreUnique(
		uint256[] memory genomes,
		uint8 gene
	) external pure returns (bool) {
		return Genome.geneValuesAreUnique(genomes, gene);
	}
}
