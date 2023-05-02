// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

/**
 * @title Seeding
 * @notice A contract that provides seeds for pseudo random number generation.
 * Seed is created from the salt, pepper, nonce, sender address, and block timestamp.
 * Seed is divided into 32 bit slices, and each slice is later used to generate a random number.
 * Seed user must keep track of used bit slices to avoid reusing them.
 * Salt is a data based on block timestamp and msg sender, and is calculated every time a seed is generated.
 * Pepper is a random data changed periodically by external entity.
 * Nonce is incremented every time a random number is generated.
 */
contract Seeding {
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
}
