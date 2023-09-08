// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import '@statechannels/nitro-protocol/contracts/interfaces/IForceMoveApp.sol';
import '@statechannels/nitro-protocol/contracts/libraries/NitroUtils.sol';

/**
 * @dev The EscrowApp contracts complies with the ForceMoveApp interface and uses consensus signatures logic.
 */
contract EscrowApp is IForceMoveApp {
	/**
	 * @notice Encodes application-specific rules for a particular ForceMove-compliant state channel.
	 * @dev Encodes application-specific rules for a particular ForceMove-compliant state channel.
	 * @param fixedPart Fixed part of the state channel.
	 * @param proof Array of recovered variable parts which constitutes a support proof for the candidate.
	 * @param candidate Recovered variable part the proof was supplied for.
	 */
	function stateIsSupported(
		FixedPart calldata fixedPart,
		RecoveredVariablePart[] calldata proof,
		RecoveredVariablePart calldata candidate
	) external pure override returns (bool, string memory) {
		require(proof.length == 0, '|proof|!=0');
		require(
			NitroUtils.getClaimedSignersNum(candidate.signedBy) == fixedPart.participants.length,
			'!unanimous'
		);

        // TODO: state 2+ requires postfund state to be supplied
		// if (proof.length == 1) {
		// 	// postfund checks
		// 	_requireProofOfUnanimousConsensusOnPostFund(proof[0], nParticipants);
        // }

        if candidate.variablePart.turnNum == 2 {
            require(candidate.variablePart.isFinal == true, '!final; turnNum==2')
        }

		return (true, '');
	}
}
