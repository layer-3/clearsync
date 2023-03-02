// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

/*
 * Gene is a number with a special structure that defines Duckling traits.
 * All traits are packed consequently in the reversed order in the Gene, meaning the first trait being stored in the last Gene bits.
 * Each trait takes up the block of 8 bits in gene, thus having 256 possible values.
 *
 * Example of gene, following traits Class, Head and Body are defined:
 *
 * 00000001|00000010|00000011
 *   Body    Head     Class
 *
 * This gene can be represented in uint24 as 66051.
 * Traits have the following values: Body = 1, Head = 2, Class = 3.
 */
library Gene {
	uint8 public constant BYTES_PER_TRAIT = 8;

	enum Classes {
		Common,
		Rare,
		Epic,
		Legendary,
		SuperLegendary
	}

	// order is important, see _generateGene()
	enum Traits {
		Class,
		Body,
		Head,
		Background,
		Element,
		Eyes,
		Beak,
		Wings,
		Firstname,
		Lastname,
		Temper,
		Peculiarity
	}

	function setTrait(
		uint256 gene,
		Traits trait,
		// by specifying uint8 we set maxCap for trait values, which is 256
		uint8 value
	) internal pure returns (uint256) {
		// number of bytes from gene's rightmost and traitBlock's rightmost
		// NOTE: maximum index of a trait is actually uint5
		uint8 shiftingBy = uint8(trait) * BYTES_PER_TRAIT;

		// remember traits we will shift off
		uint256 shiftedPart = gene % 10 ** shiftingBy;

		// shift right so that gene's rightmost bit is the traitBlock's rightmost
		gene >>= shiftingBy;

		// clear previous trait value by shifting it off
		gene >>= BYTES_PER_TRAIT;
		gene <<= BYTES_PER_TRAIT;

		// update trait's value
		gene += value;

		// reserve space for restoring previously shifted off values
		gene <<= shiftingBy;

		// restore previously shifted off values
		gene += shiftedPart;

		return gene;
	}

	function getTrait(uint256 gene, Traits trait) internal pure returns (uint8) {
		// number of bytes from gene's rightmost and traitBlock's rightmost
		// NOTE: maximum index of a trait is actually uint5
		uint8 shiftingBy = uint8(trait) * BYTES_PER_TRAIT;

		uint256 temp = gene >> shiftingBy;
		return uint8(temp % 10 ** shiftingBy);
	}

	function _maxTrait(uint256[] memory genes, Gene.Traits trait) internal pure returns (uint8) {
		uint8 maxValue = 0;

		for (uint256 i = 0; i < genes.length; i++) {
			uint8 traitValue = getTrait(genes[i], trait);
			if (maxValue < traitValue) {
				maxValue = traitValue;
			}
		}

		return maxValue;
	}

	function _traitValuesAreEqual(
		uint256[] memory genes,
		Gene.Traits trait
	) internal pure returns (bool) {
		uint8 traitValue = getTrait(genes[0], trait);

		for (uint256 i = 1; i < genes.length; i++) {
			if (getTrait(genes[i], trait) != traitValue) {
				return false;
			}
		}

		return true;
	}

	function _traitValuesAreUnique(
		uint256[] memory genes,
		Gene.Traits trait
	) internal pure returns (bool) {
		uint256 valuesPresentBitfield = 0;

		for (uint256 i = 1; i < genes.length; i++) {
			if (valuesPresentBitfield % 2 ** getTrait(genes[i], trait) == 1) {
				return false;
			}
			valuesPresentBitfield += 2 ** getTrait(genes[i], trait);
		}

		return true;
	}
}
