// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '../Utils.sol';

/**
 * @title UtilsTestConsumer
 * @notice Contract for testing Utils contract. NOT FOR USE IN PRODUCTION.
 */
contract UtilsTestConsumer {
	/**
	 * @notice Perform circular shift on the seed by 3 bytes to the left, and returns the shifted slice and the updated seed.
	 * @dev Expose internal _shiftSeedSlice function. User of this contract must keep track of used bit slices to avoid reusing them.
	 * @param seed Seed to shift and extract the shifted slice from.
	 * @return bitSlice Shifted bit slice.
	 * @return updatedSeed Shifted seed.
	 */
	function shiftSeedSlice(bytes32 seed) external pure returns (bytes3, bytes32) {
		return Utils._shiftSeedSlice(seed);
	}

	/**
	 * @notice Extracts a number from the bit slice in range [0, max).
	 * @dev Expose internal _max function. Extracts a number from the bit slice in range [0, max).
	 * @param bitSlice Bit slice to extract the number from.
	 * @param max_ Max number to extract.
	 * @return Extracted number in range [0, max).
	 */
	function max(bytes3 bitSlice, uint24 max_) external pure returns (uint256) {
		return Utils._max(bitSlice, max_);
	}

	/**
	 * @notice Generates a weighted random number given the `weights` array in range [0, weights.length).
	 * @dev Expose internal _randomWeightedNumber function. Number `x` is generated with probability `weights[x] / sum(weights)`.
	 * @param weights Array of weights.
	 * @return Random number in range [0, weights.length).
	 */
	function randomWeightedNumber(
		uint32[] memory weights,
		bytes3 bitSlice
	) external pure returns (uint256) {
		return Utils._randomWeightedNumber(weights, bitSlice);
	}

	/**
	 * @notice Check that `signatures is `encodedData` signed by `signer`. Reverts if not.
	 * @dev Expose internal _requireCorrectSigner function. Check that `signatures is `encodedData` signed by `signer`. Reverts if not.
	 * @param encodedData Data to check.
	 * @param signature Signature to check.
	 * @param signer Address of the signer.
	 */
	function requireCorrectSigner(
		bytes memory encodedData,
		bytes memory signature,
		address signer
	) external pure {
		Utils._requireCorrectSigner(encodedData, signature, signer);
	}
}
