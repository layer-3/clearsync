// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '../Genome.sol';

contract GenomeTestConsumer {
	function getFlags(uint256 genome) external pure returns (uint8) {
		return Genome.getFlags(genome);
	}

	function getFlag(uint256 genome, uint8 flag) external pure returns (bool) {
		return Genome.getFlag(genome, flag);
	}

	function setFlag(uint256 genome, uint8 flag, bool value) external pure returns (uint256) {
		return Genome.setFlag(genome, flag, value);
	}

	function setGene(uint256 genome, uint8 gene, uint8 value) external pure returns (uint256) {
		return Genome.setGene(genome, gene, value);
	}

	function getGene(uint256 genome, uint8 gene) external pure returns (uint8) {
		return Genome.getGene(genome, gene);
	}

	function maxGene(uint256[] memory genomes, uint8 gene) external pure returns (uint8) {
		return Genome._maxGene(genomes, gene);
	}

	function geneValuesAreEqual(uint256[] memory genomes, uint8 gene) external pure returns (bool) {
		return Genome._geneValuesAreEqual(genomes, gene);
	}

	function geneValuesAreUnique(
		uint256[] memory genomes,
		uint8 gene
	) external pure returns (bool) {
		return Genome._geneValuesAreUnique(genomes, gene);
	}
}
