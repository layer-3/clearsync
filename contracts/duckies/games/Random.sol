// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

// chances are represented in per mil, thus uint32
/**
 * @title Random
 * @notice A contract that provides pseudo random number generation.
 * Pseudo random number generation is based on the seed created from the salt, pepper, nonce, sender address, and block timestamp.
 * Seed is divided into 32 bit slices, and each slice is used to generate a random number.
 * User of this contract must keep track of used bit slices to avoid reusing them.
 * Salt is a data based on block timestamp and msg sender, and is calculated every time a seed is generated.
 * Pepper is a random data changed periodically by external entity.
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

	/**
	 * @notice Sets the pepper.
	 * @dev Pepper is a random data changed periodically by external entity.
	 * @param newPepper New pepper.
	 */
	function _setPepper(bytes32 newPepper) internal {
		pepper = newPepper;
	}

	/**
	 * @notice Creates a new seed based on the salt, pepper, nonce, sender address, and block timestamp.
	 * @dev Creates a new seed based on the salt, pepper, nonce, sender address, and block timestamp.
	 * @return New seed.
	 */
	function _randomSeed() internal returns (bytes32) {
		// use old salt to generate a new one, so that user's predictions are invalid after function that uses random is called
		salt = keccak256(abi.encode(salt, msg.sender, block.timestamp));
		unchecked {
			nonce++;
		}

		return keccak256(abi.encode(salt, pepper, nonce, msg.sender, block.timestamp));
	}

	/**
	 * @notice Perform circular shift on the seed by 3 bytes to the left, and returns the shifted slice and the updated seed.
	 * @dev User of this contract must keep track of used bit slices to avoid reusing them.
	 * @param seed Seed to shift and extract the shifted slice from.
	 * @return bitSlice Shifted bit slice.
	 * @return updatedSeed Shifted seed.
	 */
	function _shiftSeedSlice(bytes32 seed) internal pure returns (bytes3, bytes32) {
		bytes3 slice = bytes3(seed);
		return (slice, (seed << 24) | (bytes32(slice) >> 232));
	}

	/**
	 * @notice Extracts a number from the bit slice in range [0, max).
	 * @dev Extracts a number from the bit slice in range [0, max).
	 * @param bitSlice Bit slice to extract the number from.
	 * @param max Max number to extract.
	 * @return Extracted number in range [0, max).
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
