// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '../Random.sol';

contract RandomTestConsumer is Random {
	event NumberGenerated(uint256 number);

	function random(uint256 max, bytes3 seed) external useRandom {
		uint256 num = _random(max, seed);
		emit NumberGenerated(num);
	}

	function randomWeightedNumber(uint32[] memory weights, bytes3 seed) external useRandom {
		uint256 num = _randomWeightedNumber(weights, seed);
		emit NumberGenerated(num);
	}
}
