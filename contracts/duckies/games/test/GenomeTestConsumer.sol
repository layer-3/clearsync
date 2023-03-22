// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '../Genome.sol';

contract GenomeTestConsumer {
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
		return Genome._geneValuesAreEqual(genomes, gene);
	}
}
