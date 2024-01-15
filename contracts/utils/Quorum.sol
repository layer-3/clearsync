// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';
import '@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol';

// TODO: try using OZ's `SignatureChecker` to support ERC1271
contract Quorum {
	using ECDSA for bytes32;
	using MessageHashUtils for bytes32;

	enum QuorumConfigurationChange {
		addValidator,
		removeValidator,
		setQuorum
	}
	uint256 public quorumConfigurationNonce;
	// TODO: make public?
	address[] private _validators;
	uint8 private _quorum;

	// modifiers
	modifier onlyValidator() {
		int256 vIdx = _getValidatorIndex(msg.sender);
		require(vIdx != -1, 'Caller not validator');
		_;
	}

	// constructor
	constructor(address[] memory validators, uint8 quorum) {
		_requireNonZeroQuorum(quorum);
		_requireCorrectValidatorsQuorumRatio(validators.length, quorum);
		_requireValidValidators(validators);

		_validators = validators;
		_quorum = quorum;
	}

	// read
	function getValidators() external view onlyValidator returns (address[] memory) {
		return _validators;
	}

	function getQuorum() external view onlyValidator returns (uint8) {
		return _quorum;
	}

	// TODO: add `requireQuorum` modifier

	// TODO: decide which approach to take: setQuorumConfiguration or multiple functions

	// modify
	function addValidator(address newValidator, bytes[] calldata signatures) external {
		int256 vIdx = _getValidatorIndex(newValidator);
		require(vIdx == -1, 'Validator already present');

		_requireValidAddress(newValidator);
		requireQuorum(
			abi.encode(
				QuorumConfigurationChange.addValidator,
				quorumConfigurationNonce,
				newValidator
			),
			signatures
		);

		_validators.push(newValidator);
		quorumConfigurationNonce++;

		emit ValidatorAdded(newValidator);
	}

	function removeValidator(address validator, bytes[] calldata signatures) external {
		int256 vIdx = _getValidatorIndex(validator);
		require(vIdx != -1, 'Validator not present');

		_requireCorrectValidatorsQuorumRatio(_validators.length - 1, _quorum);

		_requireValidAddress(validator);
		requireQuorum(
			abi.encode(
				QuorumConfigurationChange.removeValidator,
				quorumConfigurationNonce,
				validator
			),
			signatures
		);

		_removeValidatorAtIndex(uint256(vIdx));
		quorumConfigurationNonce++;

		emit ValidatorRemoved(validator);
	}

	function setQuorum(uint8 newQuorum, bytes[] calldata signatures) external {
		_requireNonZeroQuorum(newQuorum);
		_requireCorrectValidatorsQuorumRatio(_validators.length, newQuorum);

		requireQuorum(
			abi.encode(QuorumConfigurationChange.setQuorum, quorumConfigurationNonce, newQuorum),
			signatures
		);

		_quorum = newQuorum;
		quorumConfigurationNonce++;

		emit QuorumChanged(newQuorum);
	}

	function setQuorumConfiguration(
		address[] calldata newValidators,
		uint8 newQuorum,
		bytes[] calldata signatures
	) external {
		requireQuorum(abi.encode(newValidators, newQuorum, quorumConfigurationNonce), signatures);

		_requireNonZeroQuorum(newQuorum);
		_requireCorrectValidatorsQuorumRatio(newValidators.length, newQuorum);
		_requireValidValidators(newValidators);

		quorumConfigurationNonce++;

		_validators = newValidators;
		_quorum = newQuorum;
	}

	// require quorum
	function requireQuorum(bytes memory encodedData, bytes[] memory signatures) public view {
		require(signatures.length >= _quorum, 'Not enough signatures for quorum');

		uint8 signers = 0;
		// bit field
		uint256 signedSoFar = 0;

		for (uint256 sigIdx = 0; sigIdx < signatures.length; sigIdx++) {
			address signer = _recoverSigner(encodedData, signatures[sigIdx]);
			_requireValidAddress(signer);

			int256 vIdx = _getValidatorIndex(signer);
			require(vIdx != -1, 'Signer not validator');

			// if signer has not been confirmed
			if ((signedSoFar >> uint256(vIdx)) % 2 != 1) {
				// confirm signer
				signedSoFar += 2 ** uint256(vIdx);

				signers++;
			}
		}

		require(signers >= _quorum, 'Quorum not achieved');
	}

	// internal
	function _requireNonZeroQuorum(uint8 quorum) internal pure {
		require(quorum > 0, 'Zero quorum');
	}

	function _requireCorrectValidatorsQuorumRatio(
		uint256 numValidators,
		uint8 quorum
	) internal pure {
		// quorum != |_validators| to protect from one validator blocking all quorum
		require(numValidators > quorum, 'Quorum too large');
	}

	function _recoverSigner(
		bytes memory encodedData,
		bytes memory signature
	) internal pure returns (address) {
		// FIXME: try OZ's `SignatureChecker` here
		return keccak256(encodedData).toEthSignedMessageHash().recover(signature);
	}

	function _requireValidAddress(address address_) internal pure {
		require(address_ != address(0), 'Invalid address');
	}

	function _requireValidValidators(address[] memory validators) internal pure {
		for (uint256 i = 0; i < validators.length; i++) {
			_requireValidAddress(validators[i]);

			for (uint256 j = i + 1; j < validators.length; j++) {
				require(validators[i] != validators[j], 'Duplicate validators');
			}
		}
	}

	function _getValidatorIndex(address validator) internal view returns (int256) {
		address[] memory validators = _validators;
		for (uint256 vIdx = 0; vIdx < validators.length; vIdx++) {
			if (validator == validators[vIdx]) {
				return int256(vIdx);
			}
		}

		return -1;
	}

	function _removeValidatorAtIndex(uint256 index) internal {
		uint256 length = _validators.length;
		_validators[index] = _validators[length - 1];
		delete _validators[length - 1];
	}

	// events
	event ValidatorAdded(address newValidator);

	event ValidatorRemoved(address validator);

	event QuorumChanged(uint8 newQuorum);
}
