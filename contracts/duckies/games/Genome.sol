// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

/**
 * @title Genome
 *
 * @notice The library to work with NFT genomes.
 *
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
	/// @notice Number of bits each gene constitutes. Thus, each gene can have 2^8 = 256 possible values.
	uint8 public constant BITS_PER_GENE = 8;

	uint8 public constant COLLECTION_GENE_IDX = 0;

	// Flags
	/// @notice Reserve 30th gene for bool flags, which are stored as a bit field.
	uint8 public constant FLAGS_GENE_IDX = 30;
	uint8 public constant FLAG_TRANSFERABLE = 1; // 0b0000_0001

	// Magic number
	/// @notice Reserve 31th gene for magic number, which is used as an extension for genomes.
	/// Genomes with wrong extension are considered invalid.
	uint8 public constant MAGIC_NUMBER_GENE_IDX = 31;
	uint8 public constant BASE_MAGIC_NUMBER = 209; // Ð
	uint8 public constant MYTHIC_MAGIC_NUMBER = 210; // Ð + 1

	/**
	 * @notice Read flags gene from genome.
	 * @dev Read flags gene from genome.
	 * @param self Genome to get flags gene from.
	 * @return flags Flags gene.
	 */
	function getFlags(uint256 self) internal pure returns (uint8) {
		return getGene(self, FLAGS_GENE_IDX);
	}

	/**
	 * @notice Read specific bit mask flag from genome.
	 * @dev Read specific bit mask flag from genome.
	 * @param self Genome to read flag from.
	 * @param flag Bit mask flag to read.
	 * @return value Value of the flag.
	 */
	function getFlag(uint256 self, uint8 flag) internal pure returns (bool) {
		return getGene(self, FLAGS_GENE_IDX) & flag > 0;
	}

	/**
	 * @notice Set specific bit mask flag in genome.
	 * @dev Set specific bit mask flag in genome.
	 * @param self Genome to set flag in.
	 * @param flag Bit mask flag to set.
	 * @param value Value of the flag.
	 * @return genome Genome with the flag set.
	 */
	function setFlag(uint256 self, uint8 flag, bool value) internal pure returns (uint256) {
		uint8 flags = getGene(self, FLAGS_GENE_IDX);
		if (value) {
			flags |= flag;
		} else {
			flags &= ~flag;
		}
		return setGene(self, FLAGS_GENE_IDX, flags);
	}

	/**
	 * @notice Set `value` to `gene` in genome.
	 * @dev Set `value` to `gene` in genome.
	 * @param self Genome to set gene in.
	 * @param gene Gene to set.
	 * @param value Value to set.
	 * @return genome Genome with the gene set.
	 */
	function setGene(
		uint256 self,
		uint8 gene,
		// by specifying uint8 we set maxCap for gene values, which is 256
		uint8 value
	) internal pure returns (uint256) {
		// number of bytes from genome's rightmost and geneBlock's rightmost
		// NOTE: maximum index of a gene is actually uint5
		uint8 shiftingBy = gene * BITS_PER_GENE;

		// remember genes we will shift off
		uint256 shiftedPart = self & ((1 << shiftingBy) - 1);

		// shift right so that genome's rightmost bit is the geneBlock's rightmost
		self >>= shiftingBy;

		// clear previous gene value by shifting it off
		self >>= BITS_PER_GENE;
		self <<= BITS_PER_GENE;

		// update gene's value
		self += value;

		// reserve space for restoring previously shifted off values
		self <<= shiftingBy;

		// restore previously shifted off values
		self += shiftedPart;

		return self;
	}

	/**
	 * @notice Get `gene` value from genome.
	 * @dev Get `gene` value from genome.
	 * @param self Genome to get gene from.
	 * @param gene Gene to get.
	 * @return geneValue Gene value.
	 */
	function getGene(uint256 self, uint8 gene) internal pure returns (uint8) {
		// number of bytes from genome's rightmost and geneBlock's rightmost
		// NOTE: maximum index of a gene is actually uint5
		uint8 shiftingBy = gene * BITS_PER_GENE;

		uint256 temp = self >> shiftingBy;
		return uint8(temp & ((1 << BITS_PER_GENE) - 1));
	}

	/**
	 * @notice Get largest value of a `gene` in `genomes`.
	 * @dev Get largest value of a `gene` in `genomes`.
	 * @param genomes Genomes to get gene from.
	 * @param gene Gene to get.
	 * @return maxValue Largest value of a `gene` in `genomes`.
	 */
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

	/**
	 * @notice Check if values of `gene` in `genomes` are equal.
	 * @dev Check if values of `gene` in `genomes` are equal.
	 * @param genomes Genomes to check.
	 * @param gene Gene to check.
	 * @return isEqual True if values of `gene` in `genomes` are equal, false otherwise.
	 */
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

	/**
	 * @notice Check if values of `gene` in `genomes` are unique.
	 * @dev Check if values of `gene` in `genomes` are unique.
	 * @param genomes Genomes to check.
	 * @param gene Gene to check.
	 * @return isUnique True if values of `gene` in `genomes` are unique, false otherwise.
	 */
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
