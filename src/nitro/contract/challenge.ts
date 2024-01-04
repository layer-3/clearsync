import {ethers, BigNumber, utils, Signature} from 'ethers';
const {Interface, keccak256, defaultAbiCoder} = utils;

import NitroAdjudicatorArtifact from '../../artifacts/contracts/NitroAdjudicator.sol/NitroAdjudicator.json';
import {SignedState} from '../signatures';

import {decodeOutcome} from './outcome';
import {FixedPart, hashState, State, VariablePart} from './state';
import {Address, Bytes32, Uint8, Uint48} from './types';

export function hashChallengeMessage(challengeState: State): Bytes32 {
  return keccak256(
    defaultAbiCoder.encode(['bytes32', 'string'], [hashState(challengeState), 'forceMove'])
  );
}

/**
 * Holds information from a ChallengeRegistered event in a convenient form
 */
export interface ChallengeRegisteredEvent {
  channelId: Bytes32; // The id of the channel that was challenged
  finalizesAt: number; // The timestamp when the channel will finalize if the challenge is not cleared
  challengeStates: SignedState[]; // An array of states used to generate the challenge
}
export interface ChallengeRegisteredStruct {
  channelId: Bytes32;
  turnNumRecord: Uint48;
  finalizesAt: Uint48;
  challenger: Address;
  isFinal: boolean;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  fixedPart: Array<any>;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  variableParts: Array<any>;
  sigs: Signature[];
  whoSignedWhat: Uint8[];
}

/**
 * Extracts a ChallengeRegisteredEvent (containing challengeStates) from the supplied eventResult.
 * @param eventResult the event itself
 * @returns a ChallengeRegisteredEvent
 */
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function getChallengeRegisteredEvent(eventResult: any[]): ChallengeRegisteredEvent {
  const {
    channelId,
    turnNumRecord,
    finalizesAt,
    isFinal,
    fixedPart,
    variableParts: variablePartsUnstructured,
    sigs,
  }: ChallengeRegisteredStruct = eventResult.slice(-1)[0].args;

  // Fixed part
  const participants = fixedPart[1].map((p: string) => BigNumber.from(p).toHexString());
  const channelNonce = fixedPart[2];
  const appDefinition = fixedPart[3];
  const challengeDuration = BigNumber.from(fixedPart[4]).toNumber();

  // Variable part
  const variableParts: VariablePart[] = variablePartsUnstructured.map(v => {
    const [outcome, appData, turnNum, isFinal] = v;
    return {outcome, appData, turnNum, isFinal};
  });

  const challengeStates: SignedState[] = variableParts.map((v, i) => {
    const turnNum = turnNumRecord - (variableParts.length - i - 1);
    const signature = sigs[i];
    const state: State = {
      turnNum,
      channelNonce,
      participants,
      outcome: decodeOutcome(v.outcome),
      appData: v.appData,
      challengeDuration,
      appDefinition,
      isFinal,
    };
    return {state, signature};
  });
  return {channelId, challengeStates, finalizesAt};
}

export interface ChallengeClearedEvent {
  kind: 'respond' | 'checkpoint';
  newStates: SignedState[];
}
export interface ChallengeClearedStruct {
  channelId: string;
  newTurnNumRecord: string;
}
export interface RespondTransactionArguments {
  challenger: string;
  isFinalAb: [boolean, boolean];
  fixedPart: FixedPart;
  variablePartAB: [VariablePart, VariablePart];
  sig: Signature;
}

/**
 * Extracts a ChallengeClearedEvent (containing a new signedState) from the logs of a respond or checkpoint transaction
 * @param tx a suitable transaction causing a ChallengeCleared event to be emitted
 * @param eventResult the event itself
 * @returns a ChallengeClearedEvent
 */
export function getChallengeClearedEvent(
  tx: ethers.Transaction,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  eventResult: any[]
): ChallengeClearedEvent {
  const {newTurnNumRecord}: ChallengeClearedStruct = eventResult.slice(-1)[0].args;

  // https://github.com/ethers-io/ethers.js/issues/602#issuecomment-574671078
  const decodedTransaction = new Interface(NitroAdjudicatorArtifact.abi).parseTransaction(tx);

  if (decodedTransaction.name === 'respond') {
    // NOTE: args value is an array of the inputted arguments, not an object with labelled keys
    // ethers.js should change this, and when it does, we can use the commented out type
    const args /* RespondTransactionArguments */ = decodedTransaction.args;
    const [participants, channelNonce, appDefinition, challengeDuration] = args[2];
    const isFinal = args[1][1];
    const outcome = decodeOutcome(args[3][1][0]);
    const appData = args[3][1][1];
    const signature: Signature = {
      v: args[4][0],
      r: args[4][1],
      s: args[4][2],
      _vs: args[4][3],
      recoveryParam: args[4][4],
    } as Signature;

    const signedState: SignedState = {
      signature,
      state: {
        challengeDuration,
        appDefinition,
        isFinal,
        outcome,
        appData,
        channelNonce,
        participants,
        turnNum: BigNumber.from(newTurnNumRecord).toNumber(),
      },
    };

    return {
      kind: 'respond',
      newStates: [signedState],
    };
  } else if (decodedTransaction.name === 'checkpoint') {
    throw new Error('UnimplementedError');
  } else {
    throw new Error(
      'Unexpected call to getChallengeClearedEvent with invalid or unrelated transaction data'
    );
  }
}
