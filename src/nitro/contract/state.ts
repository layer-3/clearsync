import {utils} from 'ethers';
import {Signature} from '@ethersproject/bytes';
import {ParamType} from 'ethers/lib/utils';

import {getChannelId} from './channel';
import {Outcome} from './outcome';
import {Address, Bytes, Bytes32, Uint256, Uint48, Uint64} from './types';

export type State = FixedPart & VariablePart;

/**
 * The part of a State which does not ordinarily change during state channel updates
 */
export interface FixedPart {
  participants: Address[];
  channelNonce: Uint64;
  appDefinition: Address;
  challengeDuration: Uint48;
}

/**
 * The part of a State which usually changes during state channel updates
 */
export interface VariablePart {
  outcome: Outcome;
  appData: Bytes; // any encoded app-related type encoded once more as bytes
  turnNum: Uint48;
  isFinal: boolean;
}

/**
 * Variable part with its signatures created by participants
 */
export interface SignedVariablePart {
  variablePart: VariablePart;
  sigs: Signature[];
}

/**
 * Variable part with a signedBy bitfield.
 */
export interface RecoveredVariablePart {
  variablePart: VariablePart;
  signedBy: Uint256;
}

/**
 * Separates supplied array into proof and candidate.
 * @param svps Array of SignedVariableParts or RecoveredVariableParts.
 * @returns proof and candidate, of the same type as input.
 */
export function separateProofAndCandidate<T = SignedVariablePart | RecoveredVariablePart>(
  svps: T[]
): {
  proof: T[];
  candidate: T;
} {
  const proof = svps.slice(0, -1);
  const candidate = svps[svps.length - 1];
  if (candidate == undefined) {
    throw Error('insufficient length array');
  }

  return {proof, candidate};
}

/**
 * Extracts the VariablePart of a state
 * @param state a State
 * @returns the VariablePart, which usually changes during state channel updates
 */
export function getVariablePart(state: State): VariablePart {
  return {
    outcome: state.outcome,
    appData: state.appData,
    turnNum: state.turnNum,
    isFinal: state.isFinal,
  };
}

/**
 * Extracts the FixedPart of a state
 * @param state a State
 * @returns the FixedPart, which does not ordinarily change during state channel updates
 */
export function getFixedPart(state: State): FixedPart {
  const {appDefinition, challengeDuration, participants, channelNonce} = state;
  return {participants, channelNonce, appDefinition, challengeDuration};
}

/**
 * Encodes and hashes the AppPart of a state
 * @param state a State
 * @returns a 32 byte keccak256 hash
 */
export function hashAppPart(state: State): Bytes32 {
  const {challengeDuration, appDefinition, appData} = state;
  return utils.keccak256(
    utils.defaultAbiCoder.encode(
      ['uint256', 'address', 'bytes'],
      [challengeDuration, appDefinition, appData]
    )
  );
}

/**
 * Encodes a state
 * @param state a State
 * @returns bytes array encoding
 */
export function encodeState(state: State): Bytes {
  const {turnNum, isFinal, appData, outcome} = state;
  const channelId = getChannelId(getFixedPart(state));

  return utils.defaultAbiCoder.encode(
    [
      'bytes32',
      'bytes',
      {
        type: 'tuple[]',
        components: [
          {name: 'asset', type: 'address'},
          {
            name: 'assetMetadata',
            type: 'tuple',
            components: [
              {name: 'assetType', type: 'uint8'},
              {name: 'metadata', type: 'bytes'},
            ],
          },
          {
            type: 'tuple[]',
            name: 'allocations',
            components: [
              {name: 'destination', type: 'bytes32'},
              {name: 'amount', type: 'uint256'},
              {name: 'allocationType', type: 'uint8'},
              {name: 'metadata', type: 'bytes'},
            ],
          } as ParamType,
        ],
      } as ParamType,
      'uint256',
      'bool',
    ],
    [channelId, appData, outcome, turnNum, isFinal]
  );
}

/**
 * Hashes a state
 * @param state a State
 * @returns a 32 byte keccak256 hash
 */
export function hashState(state: State): Bytes32 {
  return utils.keccak256(encodeState(state));
}
