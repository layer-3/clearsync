// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '../Random.sol';

contract RandomTestConsumer is Random {
	event NumberGenerated(uint256 number);

	function randomMaxNumber(uint256 max) external UseRandom {
		uint256 num = _randomMaxNumber(max);
		emit NumberGenerated(num);
	}

	function randomWeightedNumber(uint32[] memory weights) external UseRandom {
		uint256 num = _randomWeightedNumber(weights);
		emit NumberGenerated(num);
	}
}
