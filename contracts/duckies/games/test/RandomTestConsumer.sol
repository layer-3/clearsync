// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '../Random.sol';

contract RandomTestConsumer is Random {
	event SeedGenerated(bytes32 seed);

	function randomSeed() external {
		emit SeedGenerated(_randomSeed());
	}

	function shiftSeedSlice(bytes32 seed) external pure returns (bytes3, bytes32) {
		return _shiftSeedSlice(seed);
	}

	function max(bytes3 seed, uint24 max_) external pure returns (uint256) {
		return _max(seed, max_);
	}

	function randomWeightedNumber(
		uint32[] memory weights,
		bytes3 seed
	) external pure returns (uint256) {
		return _randomWeightedNumber(weights, seed);
	}
}
