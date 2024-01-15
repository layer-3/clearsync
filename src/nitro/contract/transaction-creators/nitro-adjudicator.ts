import { Signature, constants, providers, utils } from 'ethers';

import NitroAdjudicatorArtifact from '../../../../artifacts/contracts/nitro/NitroAdjudicator.sol/NitroAdjudicator.json';
import { bindSignatures, getChannelId, hashState } from '../..';
import { encodeOutcome } from '../outcome';
import { State, getFixedPart, getVariablePart, separateProofAndCandidate } from '../state';

// https://github.com/ethers-io/ethers.js/issues/602#issuecomment-574671078
const NitroAdjudicatorContractInterface = new utils.Interface(NitroAdjudicatorArtifact.abi);

export function concludeAndTransferAllAssetsArgs(
  states: State[],
  signatures: Signature[],
  whoSignedWhat: number[],
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
): any[] {
  // Sanity checks on expected lengths
  if (states.length === 0) {
    throw new Error('No states provided');
  }
  const { participants } = states[0];
  if (participants.length !== signatures.length) {
    throw new Error(
      `Participants (length:${participants.length}) and signatures (length:${signatures.length}) need to be the same length`,
    );
  }

  const fixedPart = getFixedPart(states[0]);
  const variableParts = states.map((s) => getVariablePart(s));
  const { proof, candidate } = separateProofAndCandidate(
    bindSignatures(variableParts, signatures, whoSignedWhat),
  );

  return [fixedPart, proof, candidate];
}

export function createConcludeAndTransferAllAssetsTransaction(
  states: State[],
  signatures: Signature[],
  whoSignedWhat: number[],
): providers.TransactionRequest {
  return {
    data: NitroAdjudicatorContractInterface.encodeFunctionData(
      'concludeAndTransferAllAssets',
      concludeAndTransferAllAssetsArgs(states, signatures, whoSignedWhat),
    ),
  };
}

function transferAllAssetsArgs(
  state: State,
  overrideStateHash = false, // set to true if channel concluded happily
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
): any[] {
  const channelId = getChannelId(getFixedPart(state));
  const outcomeBytes = encodeOutcome(state.outcome);
  const stateHash = overrideStateHash ? constants.HashZero : hashState(state);
  return [channelId, outcomeBytes, stateHash];
}

export function createTransferAllAssetsTransaction(
  state: State,
  overrideStateHash = false, // set to true if channel concluded happily
): providers.TransactionRequest {
  return {
    data: NitroAdjudicatorContractInterface.encodeFunctionData(
      'transferAllAssets',
      transferAllAssetsArgs(state, overrideStateHash),
    ),
  };
}
