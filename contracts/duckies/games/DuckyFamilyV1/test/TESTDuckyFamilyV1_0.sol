// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '../DuckyFamilyV1_0.sol';

contract TESTDuckyFamilyV1_0 is DuckyFamilyV1_0 {
	event GenomeReturned(uint256 genome);
	event GeneReturned(uint8 gene);
	event BoolReturned(bool returnedBool);
	event Uint8Returned(uint8 returnedUint8);

	function generateGenome(uint8 collectionId) external {
		emit GenomeReturned(_generateGenome(collectionId));
	}

	function generateAndSetGenes(uint256 genome, uint8 collectionId) external {
		emit GenomeReturned(_generateAndSetGenes(genome, collectionId));
	}

	function requireGenomesSatisfyMelding(uint256[] calldata genomes) external pure {
		_requireGenomesSatisfyMelding(genomes);
	}

	function meldGenomes(uint256[] calldata genomes) external {
		emit GenomeReturned(_meldGenomes(genomes));
	}

	function isCollectionMutating(Rarities rarity) external {
		emit BoolReturned(_isCollectionMutating(rarity));
	}

	function meldGenes(
		uint256[] calldata genomes,
		uint8 gene,
		uint8 maxGeneValue,
		GeneDistributionTypes geneDistrType
	) external {
		emit GeneReturned(_meldGenes(genomes, gene, maxGeneValue, geneDistrType));
	}

	function getDistibutionType(
		uint32 distributionTypes,
		uint8 idx
	) external pure returns (GeneDistributionTypes) {
		return _getDistibutionType(distributionTypes, idx);
	}

	function generateUnevenGeneValue(uint8 valuesNum) external {
		emit Uint8Returned(_generateUnevenGeneValue(valuesNum));
	}
}
