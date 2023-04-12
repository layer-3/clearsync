// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

/*
 * Genome is a number with a special structure that defines Duckling genes.
 * All genes are packed consequently in the reversed order in the Genome, meaning the first gene being stored in the last Genome bits.
 * Each gene takes up the block of 8 bits in genome, thus having 256 possible values.
 *
 * Example of genome, following genes Rarity, Head and Body are defined:
 *
 * 00000001|00000010|00000011
 *   Body    Head     Rarity
 *
 * This genome can be represented in uint24 as 66051.
 * Genes have the following values: Body = 1, Head = 2, Rarity = 3.
 */
library Genome {
	uint8 public constant BITS_PER_GENE = 8;

	uint8 public constant COLLECTION_GENE_IDX = 0;
	uint8 public constant FLAGS_GENE_IDX = 30;
	uint8 public constant MAGIC_NUMBER_GENE_IDX = 31;

	uint8 public constant FLAG_TRANSFERABLE = 1; // 0b0000_0001

	function getFlags(uint256 genome) internal pure returns (uint8) {
		return getGene(genome, FLAGS_GENE_IDX);
	}

	function setGene(
		uint256 genome,
		uint8 gene,
		// by specifying uint8 we set maxCap for gene values, which is 256
		uint8 value
	) internal pure returns (uint256) {
		// number of bytes from genome's rightmost and geneBlock's rightmost
		// NOTE: maximum index of a gene is actually uint5
		uint8 shiftingBy = gene * BITS_PER_GENE;

		// remember genes we will shift off
		uint256 shiftedPart = genome & ((1 << shiftingBy) - 1);

		// shift right so that genome's rightmost bit is the geneBlock's rightmost
		genome >>= shiftingBy;

		// clear previous gene value by shifting it off
		genome >>= BITS_PER_GENE;
		genome <<= BITS_PER_GENE;

		// update gene's value
		genome += value;

		// reserve space for restoring previously shifted off values
		genome <<= shiftingBy;

		// restore previously shifted off values
		genome += shiftedPart;

		return genome;
	}

	function getGene(uint256 genome, uint8 gene) internal pure returns (uint8) {
		// number of bytes from genome's rightmost and geneBlock's rightmost
		// NOTE: maximum index of a gene is actually uint5
		uint8 shiftingBy = gene * BITS_PER_GENE;

		uint256 temp = genome >> shiftingBy;
		return uint8(temp & ((1 << BITS_PER_GENE) - 1));
	}

	function _maxGene(uint256[] memory genomes, uint8 gene) internal pure returns (uint8) {
		uint8 maxValue = 0;

		for (uint256 i = 0; i < genomes.length; i++) {
			uint8 geneValue = getGene(genomes[i], gene);
			if (maxValue < geneValue) {
				maxValue = geneValue;
			}
		}

		return maxValue;
	}

	function _geneValuesAreEqual(
		uint256[] memory genomes,
		uint8 gene
	) internal pure returns (bool) {
		uint8 geneValue = getGene(genomes[0], gene);

		for (uint256 i = 1; i < genomes.length; i++) {
			if (getGene(genomes[i], gene) != geneValue) {
				return false;
			}
		}

		return true;
	}

	function _geneValuesAreUnique(
		uint256[] memory genomes,
		uint8 gene
	) internal pure returns (bool) {
		uint256 valuesPresentBitfield = 1 << getGene(genomes[0], gene);

		for (uint256 i = 1; i < genomes.length; i++) {
			if (valuesPresentBitfield & (1 << getGene(genomes[i], gene)) != 0) {
				return false;
			}
			valuesPresentBitfield |= 1 << getGene(genomes[i], gene);
		}

		return true;
	}
}
