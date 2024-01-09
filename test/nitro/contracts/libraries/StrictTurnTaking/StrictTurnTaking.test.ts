import { BigNumber, Contract, ContractInterface, Wallet } from 'ethers';
import { describe, before, beforeEach, it } from 'mocha';

import { expectRevert } from '../../../../helpers/expect-revert';
import {
  shortenedToRecoveredVariableParts,
  TurnNumToShortenedVariablePart,
} from '../../../../../src/nitro/signatures';
import { generateParticipants, setupContract } from '../../../test-helpers';
import type { CountingApp, TESTStrictTurnTaking } from '../../../../../typechain-types';
import {
  getFixedPart,
  getRandomNonce,
  getVariablePart,
  Outcome,
  State,
} from '../../../../../src/nitro';
import {
  INVALID_NUMBER_OF_PROOF_STATES,
  INVALID_SIGNED_BY,
  TOO_MANY_PARTICIPANTS,
  WRONG_TURN_NUM,
} from '../../../../../src/nitro/contract/transaction-creators/revert-reasons';
import {
  RecoveredVariablePart,
  separateProofAndCandidate,
} from '../../../../../src/nitro/contract/state';
import { getSignedBy } from '../../../../../src/nitro/bitfield-utils';
import { expectSucceedWithNoReturnValues } from '../../../tx-expect-wrappers';
import { expect } from 'chai';

let strictTurnTaking: Contract & TESTStrictTurnTaking;
let countingApp: Contract & CountingApp;

const challengeDuration = 0x1000;
const asset = Wallet.createRandom().address;
const defaultOutcome: Outcome = [
  { asset, allocations: [], assetMetadata: { assetType: 0, metadata: '0x' } },
];
const nParticipants = 3;
const { wallets, participants } = generateParticipants(nParticipants);

before(async () => {
  strictTurnTaking = await setupContract<TESTStrictTurnTaking>('TESTStrictTurnTaking');
  countingApp = await setupContract<CountingApp>('CountingApp');
});

let channelNonce = getRandomNonce('StrictTurnTaking');
beforeEach(() => (channelNonce = BigNumber.from(channelNonce).add(1).toHexString()));

describe('isSignedByMover', () => {
  const testCases = [
    {
      description: 'should not revert when signed only by mover',
      turnNum: 3,
      signedBy: [0],
      reason: undefined,
    },
    {
      description: 'should revert when not signed by mover',
      turnNum: 3,
      signedBy: [2],
      reason: INVALID_SIGNED_BY,
    },
    {
      description: 'should revert when signed not only by mover',
      turnNum: 3,
      signedBy: [0, 1],
      reason: INVALID_SIGNED_BY,
    },
  ];

  testCases.forEach((tc) =>
    it(tc.description, async () => {
      const { turnNum, signedBy, reason } = tc as unknown as {
        turnNum: number;
        signedBy: number[];
        reason: string | undefined;
      };

      const state: State = {
        turnNum,
        isFinal: false,
        participants,
        channelNonce,
        challengeDuration,
        outcome: defaultOutcome,
        appDefinition: countingApp.address,
        appData: '0x',
      };

      const variablePart = getVariablePart(state);
      const fixedPart = getFixedPart(state);

      const rvp: RecoveredVariablePart = {
        variablePart,
        signedBy: BigNumber.from(getSignedBy(signedBy)).toHexString(),
      };

      if (reason) {
        await expectRevert(() => strictTurnTaking.isSignedByMover(fixedPart, rvp), reason);
      } else {
        await expectSucceedWithNoReturnValues(() =>
          strictTurnTaking.isSignedByMover(fixedPart, rvp),
        );
      }
    }),
  );
});

describe('moverAddress', () => {
  const testCases = [
    {
      description: 'return correct mover',
      turnNum: 0,
      expectedParticipantIdx: 0,
    },
    {
      description: 'return correct mover',
      turnNum: 1,
      expectedParticipantIdx: 1,
    },
    {
      description: 'return correct mover',
      turnNum: 2,
      expectedParticipantIdx: 2,
    },
    {
      description: 'return correct mover for turnNum >= numParticipants',
      turnNum: 3,
      expectedParticipantIdx: 0,
    },
    {
      description: 'return correct mover for turnNum >= numParticipants',
      turnNum: 7,
      expectedParticipantIdx: 1,
    },
  ];

  testCases.forEach((tc) =>
    it(tc.description, async () => {
      const { turnNum, expectedParticipantIdx } = tc as unknown as {
        turnNum: number;
        expectedParticipantIdx: number;
      };

      expect(await strictTurnTaking.moverAddress(participants, turnNum)).to.equal(
        wallets[expectedParticipantIdx].address,
      );
    }),
  );
});

describe('requireValidInput', () => {
  const testCases = [
    {
      description: 'accept when all rules are satisfied',
      nParticipants: 2,
      numProof: 1,
      reason: undefined,
    },
    {
      description: 'accept when all rules are satisfied',
      nParticipants: 4,
      numProof: 3,
      reason: undefined,
    },
    {
      description: 'revert when supplied zero proof states',
      nParticipants: 2,
      numProof: 0,
      reason: INVALID_NUMBER_OF_PROOF_STATES,
    },
    {
      description: 'revert when supplied not enough proof states',
      nParticipants: 4,
      numProof: 1,
      reason: INVALID_NUMBER_OF_PROOF_STATES,
    },
    {
      description: 'revert when supplied excessive proof states',
      nParticipants: 2,
      numProof: 2,
      reason: INVALID_NUMBER_OF_PROOF_STATES,
    },
    {
      description: 'revert when too many participants',
      nParticipants: 256,
      numProof: 255,
      reason: TOO_MANY_PARTICIPANTS,
    },
  ];

  testCases.forEach((tc) =>
    it(tc.description, async () => {
      if (tc.reason) {
        await expectRevert(
          () => strictTurnTaking.requireValidInput(tc.nParticipants, tc.numProof),
          tc.reason,
        );
      } else {
        await expectSucceedWithNoReturnValues(() =>
          strictTurnTaking.requireValidInput(tc.nParticipants, tc.numProof),
        );
      }
    }),
  );
});

describe('requireValidTurnTaking', () => {
  const testCases = [
    {
      description: 'accept when strict turn taking from 0',
      turnNumToShortenedVariablePart: new Map([
        [0, [0]],
        [1, [1]],
        [2, [2]],
      ]),
      reason: undefined,
    },
    {
      description: 'accept when strict turn taking not from 0',
      turnNumToShortenedVariablePart: new Map([
        [3, [0]],
        [4, [1]],
        [5, [2]],
      ]),
      reason: undefined,
    },
    {
      description: 'revert when insufficient states',
      turnNumToShortenedVariablePart: new Map([
        [0, [0]],
        [1, [1]],
      ]),
      reason: INVALID_NUMBER_OF_PROOF_STATES,
    },
    {
      description: 'revert when excess states',
      turnNumToShortenedVariablePart: new Map([
        [0, [0]],
        [1, [1]],
        [2, [2]],
        [3, [0]],
      ]),
      reason: INVALID_NUMBER_OF_PROOF_STATES,
    },
    {
      description: 'revert when a state is signed by multiple participants',
      turnNumToShortenedVariablePart: new Map([
        [0, [0]],
        [1, [1, 2]],
        [2, [2]],
      ]),
      reason: INVALID_SIGNED_BY,
    },
    {
      description: 'revert when a state is not signed',
      turnNumToShortenedVariablePart: new Map([
        [0, [0]],
        [1, []],
        [2, [2]],
      ]),
      reason: INVALID_SIGNED_BY,
    },
    {
      description: 'revert when a state signed by non mover',
      turnNumToShortenedVariablePart: new Map([
        [0, [0]],
        [1, [2]],
        [2, [1]],
      ]),
      reason: INVALID_SIGNED_BY,
    },
    {
      description: 'revert when a turn number is skipped',
      turnNumToShortenedVariablePart: new Map([
        [0, [0]],
        [2, [1]],
        [3, [2]],
      ]),
      reason: WRONG_TURN_NUM,
    },
  ];

  testCases.forEach((tc) =>
    it(tc.description, async () => {
      const { reason, turnNumToShortenedVariablePart } = tc as unknown as {
        reason: string | undefined;
        turnNumToShortenedVariablePart: TurnNumToShortenedVariablePart;
      };

      const state: State = {
        turnNum: 0,
        isFinal: false,
        participants,
        channelNonce,
        challengeDuration,
        outcome: defaultOutcome,
        appDefinition: countingApp.address,
        appData: '0x',
      };

      const fixedPart = getFixedPart(state);

      const recoveredVP = shortenedToRecoveredVariableParts(turnNumToShortenedVariablePart);
      const { proof, candidate } = separateProofAndCandidate(recoveredVP);

      if (reason) {
        await expectRevert(() =>
          strictTurnTaking.requireValidTurnTaking(fixedPart, proof, candidate),
        );
      } else {
        await expectSucceedWithNoReturnValues(() =>
          strictTurnTaking.requireValidTurnTaking(fixedPart, proof, candidate),
        );
      }
    }),
  );
});
