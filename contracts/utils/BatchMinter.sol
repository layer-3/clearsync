// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '../duckies/ducklings/DucklingsV1.sol';

contract BatchMinter {
	DucklingsV1 public ducklings;

	constructor(address _ducklings) {
		ducklings = DucklingsV1(_ducklings);
	}

	function mintToAddresses(uint256[] calldata genomes, address[] calldata addresses) external {
		require(genomes.length > 0 && genomes.length == addresses.length, 'Invalid input lengths');

		for (uint256 i = 0; i < genomes.length; i++) {
			ducklings.mintTo(addresses[i], genomes[i]);
		}
	}
}
