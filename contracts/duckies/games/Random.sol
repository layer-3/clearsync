// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

// TODO: dev docs
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
	bytes32 private pepper;
	uint256 private nonce;

	function _setPepper(bytes32 newPepper) internal {
		pepper = newPepper;
	}

	function _randomSeed() internal returns (bytes32) {
		// use old salt to generate a new one, so that user's predictions are invalid after function that uses random is called
		salt = keccak256(abi.encode(salt, msg.sender, block.timestamp));
		unchecked {
			nonce++;
		}

		return keccak256(abi.encode(salt, pepper, nonce, msg.sender, block.timestamp));
	}

	// circular shift of 3 bytes to the left
	function _shiftSeedSlice(bytes32 seed) internal pure returns (bytes3, bytes32) {
		bytes3 slice = bytes3(seed);
		return (slice, (seed << 24) | (bytes32(slice) >> 232));
	}

	/**
	 * @notice Generates a random number in range [0, max).
	 * @dev Calculates hash of encoded salt, nonce, msg sender block timestamp to the number, and returns modulo `max`.
	 * @param max Upper bound of the range.
	 * @param bitSlice Upper bound of the range.
	 * @return Random number in range [0, max).
	 */
	function _max(bytes3 bitSlice, uint24 max) internal pure returns (uint24) {
		return uint24(bitSlice) % max;
	}

	/**
	 * @notice Generates a weighted random number given the `weights` array in range [0, weights.length).
	 * @dev Number `x` is generated with probability `weights[x] / sum(weights)`.
	 * @param weights Array of weights.
	 * @return Random number in range [0, weights.length).
	 */
	function _randomWeightedNumber(
		uint32[] memory weights,
		bytes3 bitSlice
	) internal pure returns (uint8) {
		// no sense in empty weights array
		if (weights.length == 0) revert InvalidWeights(weights);

		uint256 randomNumber = _max(bitSlice, uint24(_sum(weights)));

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
