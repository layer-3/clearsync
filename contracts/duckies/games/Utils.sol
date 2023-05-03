// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';

import '../../interfaces/IVoucher.sol';

/**
 * @title Utils
 * @notice Utility functions for games, that include pseudo random number generation, signature verification, etc.
 *
 * Pseudo random number generation is based on the bitSlices, which are part of a seed created in `Seeding.sol`.
 */
library Utils {
	using ECDSA for bytes32;

	/**
	 * @notice Invalid weights error while trying to generate a weighted random number.
	 * @param weights Empty weights array.
	 */
	error InvalidWeights(uint32[] weights);

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
		uint32[] memory weights, // chances are represented in per mil
		bytes3 bitSlice
	) internal pure returns (uint8) {
		uint24 sum = uint24(_sum(weights));

		// no sense in empty weights array
		if (weights.length == 0 || sum == 0) revert InvalidWeights(weights);

		uint256 randomNumber = _max(bitSlice, sum);

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

	/**
	 * @notice Check that `signatures is `encodedData` signed by `signer`. Reverts if not.
	 * @dev Check that `signatures is `encodedData` signed by `signer`. Reverts if not.
	 * @param encodedData Data to check.
	 * @param signature Signature to check.
	 * @param signer Address of the signer.
	 */
	function _requireCorrectSigner(
		bytes memory encodedData,
		bytes memory signature,
		address signer
	) internal pure {
		address actualSigner = keccak256(encodedData).toEthSignedMessageHash().recover(signature);
		if (actualSigner != signer) revert IVoucher.IncorrectSigner(signer, actualSigner);
	}
}
