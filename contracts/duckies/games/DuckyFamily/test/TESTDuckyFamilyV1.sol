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
	 * @param seed Seed for randomization.
	 */
	function generateGenome(uint8 collectionId, bytes32 seed) external {
		emit GenomeReturned(_generateGenome(collectionId, seed));
	}

	/**
	 * @notice Generates a mythic genome. Emits GenomeReturned event.
	 * @dev Exposed for testing.
	 * @param genomes Genomes to generate mythic genome from.
	 * @param maxPeculiarity Max peculiarity for mythic genome.
	 * @param mythicAmount Amount of mythic genes to generate.
	 * @param seed Seed for randomization.
	 */
	function generateMythicGenome(
		uint256[] calldata genomes,
		uint16 maxPeculiarity,
		uint16 mythicAmount,
		bytes32 seed
	) external {
		emit GenomeReturned(_generateMythicGenome(genomes, maxPeculiarity, mythicAmount, seed));
	}

	/**
	 * @notice Melds genomes. Emits GenomeReturned event.
	 * @dev Exposed for testing.
	 * @param genomes Genomes to meld.
	 * @param seed Seed for randomization.
	 */
	function meldGenomes(uint256[] calldata genomes, bytes32 seed) external {
		emit GenomeReturned(_meldGenomes(genomes, seed));
	}
}
