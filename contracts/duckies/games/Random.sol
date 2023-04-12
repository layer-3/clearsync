// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

// chances are represented in per mil, thus uint32
contract Random {
	error InvalidWeights(uint32[] weights);

	bytes32 private salt;
	uint256 private nonce;

	// specifies an external function which uses Random logic
	modifier UseRandom() {
		_;
		_updateSalt();
	}

	function _updateNonce() private {
		unchecked {
			nonce++;
		}
	}

	function _updateSalt() private {
		salt = keccak256(abi.encode(msg.sender, block.timestamp));
	}

	function _randomMaxNumber(uint256 max) internal returns (uint256) {
		_updateNonce();
		return uint256(keccak256(abi.encode(salt, nonce, msg.sender, block.timestamp))) % max;
	}

	function _randomWeightedNumber(uint32[] memory weights) internal returns (uint8) {
		// no sense in empty weights array
		if (weights.length == 0) revert InvalidWeights(weights);

		// generated number should be strictly less than right \/ segment boundary
		uint256 randomNumber = _randomMaxNumber(_sum(weights));

		uint256 segmentRightBoundary = 0;

		for (uint8 i = 0; i < weights.length; i++) {
			segmentRightBoundary += weights[i];
			if (randomNumber < segmentRightBoundary) {
				return i;
			}
		}

		// execution should never reach this
		return uint8(weights.length - 1);
	}

	function _sum(uint32[] memory numbers) internal pure returns (uint256 sum) {
		for (uint256 i = 0; i < numbers.length; i++) sum += numbers[i];
	}
}
