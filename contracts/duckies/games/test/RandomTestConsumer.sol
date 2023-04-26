// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '../Random.sol';

contract RandomTestConsumer is Random {
	event SeedGenerated(bytes32 seed);

	function randomSeed() external {
		emit SeedGenerated(_randomSeed());
	}

	function rotateSeedChunk(bytes32 seed) external pure returns (bytes3, bytes32) {
		return _rotateSeedChunk(seed);
	}

	function random(uint256 max, bytes3 seed) external view returns (uint256) {
		return _random(max, seed);
	}

	function randomWeightedNumber(
		uint32[] memory weights,
		bytes3 seed
	) external view returns (uint256) {
		return _randomWeightedNumber(weights, seed);
	}
}
