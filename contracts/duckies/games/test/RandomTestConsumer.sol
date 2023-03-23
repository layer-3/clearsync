// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '../RandomUpgradeable.sol';

contract RandomTestConsumer is RandomUpgradeable {
	event NumberGenerated(uint256 number);

	function randomMaxNumber(uint256 max) external UseRandom {
		uint256 num = _randomMaxNumber(max);
		emit NumberGenerated(num);
	}

	function randomWeightedNumber(uint8[] memory weights) external UseRandom {
		uint256 num = _randomWeightedNumber(weights);
		emit NumberGenerated(num);
	}
}
