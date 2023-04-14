// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

// chances are represented in per mil, thus uint32
/**
 * @title Random
 * @notice A contract that provides pseudo random number generation.
 * Pseudo random number generation is based on the block timestamp, sender address, salt and nonce.
 * Salt is based on block timestamp and msg sender, and is calculated every time a user-function that uses Random logic is called.
 * Nonce is incremented every time a random number is generated.
 */
contract Random {
	/**
	 * @notice Invalid weights error while trying to generate a weighted random number.
	 * @param weights Empty weights array.
	 */
	error InvalidWeights(uint32[] weights);

	bytes32 private salt;
	uint256 private nonce;

	/**
	 * @notice Specifies that calling function uses random number generation.
	 * @dev Modifier that updates salt after calling function is invoked.
	 */
	modifier UseRandom() {
		_;
		_updateSalt();
	}

	/**
	 * @notice Updates nonce
	 */
	function _updateNonce() private {
		unchecked {
			nonce++;
		}
	}

	/**
	 * @notice Updates salt.
	 * @dev Salt is a hash of encoded block timestamp and msg sender.
	 */
	function _updateSalt() private {
		salt = keccak256(abi.encode(msg.sender, block.timestamp));
	}

	/**
	 * @notice Generates a random number in range [0, max).
	 * @dev Cast hash of encoded salt, nonce, msg sender block timestamp to the number, and returns modulo `max`.
	 * @param max Upper bound of the range.
	 * @return Random number in range [0, max).
	 */
	function _randomMaxNumber(uint256 max) internal returns (uint256) {
		_updateNonce();
		return uint256(keccak256(abi.encode(salt, nonce, msg.sender, block.timestamp))) % max;
	}

	/**
	 * @notice Generates a weighted random number given the `weights` array in range [0, weights.length).
	 * @dev Number `x` is generated with probability `weights[x] / sum(weights)`.
	 * @param weights Array of weights.
	 * @return Random number in range [0, weights.length).
	 */
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

	/**
	 * @notice Calculates sum of all elements in array.
	 * @dev Calculates sum of all elements in array.
	 * @param numbers Array of numbers.
	 * @return sum Sum of all elements in array.
	 */
	function _sum(uint32[] memory numbers) internal pure returns (uint256 sum) {
		for (uint256 i = 0; i < numbers.length; i++) sum += numbers[i];
	}
}
