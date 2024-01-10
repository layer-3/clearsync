import { BigNumber, Contract, Wallet, ethers } from 'ethers';
import { before, describe, it } from 'mocha';
import { expect } from 'chai';

import { expectRevert } from '../../../helpers/expect-revert';
import { getChannelId } from '../../../../src/nitro/contract/channel';
import { channelDataToStatus } from '../../../../src/nitro/contract/channel-storage';
import {
  State,
  getFixedPart,
  getVariablePart,
  separateProofAndCandidate,
} from '../../../../src/nitro/contract/state';
import { generateParticipants, setupContract } from '../../test-helpers';
import { bindSignatures, getRandomNonce, signStates } from '../../../../src/nitro';
import {
  CHANNEL_FINALIZED,
  COUNTING_APP_INVALID_TRANSITION,
  INVALID_SIGNED_BY,
  TURN_NUM_RECORD_NOT_INCREASED,
} from '../../../../src/nitro/contract/transaction-creators/revert-reasons';

import type { CountingApp, TESTForceMove } from '../../../../typechain-types';
import type { Outcome } from '../../../../src/nitro/contract/outcome';

const { HashZero } = ethers.constants;
const { defaultAbiCoder } = ethers.utils;

interface transitionType {
  whoSignedWhat: number[];
  appDatas: number[];
}

interface testParams {
  largestTurnNum: number;
  support: transitionType;
  finalizesAt: number | undefined;
  reason: string | undefined;
}

let forceMove: Contract & TESTForceMove;
let countingApp: Contract & CountingApp;

const participantsNum = 3;
const { wallets, participants } = generateParticipants(participantsNum);

const challengeDuration = 0x10_00;
const asset = Wallet.createRandom().address;
const defaultOutcome: Outcome = [
  { asset, allocations: [], assetMetadata: { assetType: 0, metadata: '0x' } },
];

before(async () => {
  forceMove = await setupContract<TESTForceMove>('TESTForceMove');
  countingApp = await setupContract<CountingApp>('CountingApp');
});

const valid = {
  whoSignedWhat: [0, 1, 2],
  appDatas: [0, 1, 2],
};
const invalidTransition = {
  whoSignedWhat: [0, 1, 2],
  appDatas: [0, 2, 1],
};
const unsupported = {
  whoSignedWhat: [0, 0, 0],
  appDatas: [0, 1, 2],
};

const future = 1e12;
const past = 1;
const never = '0x00';
const turnNumRecord = 7;

describe('checkpoint', () => {
  let channelNonce = getRandomNonce('checkpoint');
  beforeEach(() => (channelNonce = BigNumber.from(channelNonce).add(1).toHexString()));

  const testCases = [
    {
      description:
        'It accepts valid input, and clears any existing challenge, if ' + 'the slot is empty',
      largestTurnNum: turnNumRecord + 1,
      support: valid,
      finalizesAt: undefined,
      reason: undefined,
    },
    {
      description:
        'It accepts valid input, and clears any existing challenge, if ' +
        'there is a challenge and the existing turnNumRecord is increased',
      largestTurnNum: turnNumRecord + 1,
      support: valid,
      finalizesAt: future,
      reason: undefined,
    },
    {
      description:
        'It accepts valid input, and clears any existing challenge, if ' +
        'there is no challenge and the existing turnNumRecord is increased',
      largestTurnNum: turnNumRecord + 1 + participantsNum,
      support: valid,
      finalizesAt: never,
      reason: undefined,
    },
    {
      description:
        'It reverts when the channel is open, but ' + 'the turnNumRecord is not increased.',
      largestTurnNum: turnNumRecord,
      support: valid,
      finalizesAt: never,
      reason: TURN_NUM_RECORD_NOT_INCREASED,
    },
    {
      description: 'It reverts when the channel is open, but ' + 'there is an invalid transition',
      largestTurnNum: turnNumRecord + 1,
      support: invalidTransition,
      finalizesAt: never,
      reason: COUNTING_APP_INVALID_TRANSITION,
    },
    {
      description: 'It reverts when the channel is open, but ' + 'the final state is not supported',
      largestTurnNum: turnNumRecord + 1,
      support: unsupported,
      finalizesAt: never,
      reason: INVALID_SIGNED_BY,
    },
    {
      description:
        'It reverts when there is an ongoing challenge, but ' +
        'the turnNumRecord is not increased.',
      largestTurnNum: turnNumRecord,
      support: valid,
      finalizesAt: future,
      reason: TURN_NUM_RECORD_NOT_INCREASED,
    },
    {
      description:
        'It reverts when there is an ongoing challenge, but ' + 'there is an invalid transition',
      largestTurnNum: turnNumRecord + 1,
      support: invalidTransition,
      finalizesAt: future,
      reason: COUNTING_APP_INVALID_TRANSITION,
    },
    {
      description:
        'It reverts when there is an ongoing challenge, but ' + 'the final state is not supported',
      largestTurnNum: turnNumRecord + 1,
      support: unsupported,
      finalizesAt: future,
      reason: INVALID_SIGNED_BY,
    },
    {
      description: 'It reverts when a challenge has expired',
      largestTurnNum: turnNumRecord + 1,
      support: valid,
      finalizesAt: past,
      reason: CHANNEL_FINALIZED,
    },
  ];

  for (const tc of testCases) it(tc.description, async () => {
      const { largestTurnNum, support, finalizesAt, reason } = tc as unknown as testParams;
      const { appDatas, whoSignedWhat } = support;

      const states: State[] = appDatas.map((data, idx) => ({
        turnNum: largestTurnNum - appDatas.length + 1 + idx,
        isFinal: false,
        channelNonce,
        participants,
        challengeDuration,
        outcome: defaultOutcome,
        appData: defaultAbiCoder.encode(['uint256'], [data]),
        appDefinition: countingApp.address,
      }));

      const variableParts = states.map((state) => getVariablePart(state));
      const fixedPart = getFixedPart(states[0]);
      const channelId = getChannelId(fixedPart);

      // Sign the states
      const signatures = await signStates(states, wallets, whoSignedWhat);
      const { proof, candidate } = separateProofAndCandidate(
        bindSignatures(variableParts, signatures, whoSignedWhat),
      );

      const isChallenged = finalizesAt && finalizesAt > Math.floor(Date.now() / 1000);
      const outcome = isChallenged ? defaultOutcome : [];

      const challengeState: State | undefined = isChallenged
        ? {
            turnNum: turnNumRecord,
            isFinal: false,
            channelNonce,
            participants,
            outcome,
            appData: defaultAbiCoder.encode(['uint256'], [appDatas[0]]),
            appDefinition: countingApp.address,
            challengeDuration,
          }
        : undefined;

      const fingerprint = finalizesAt
        ? channelDataToStatus({
            turnNumRecord,
            finalizesAt,
            state: challengeState,
            outcome,
          })
        : HashZero;

      // Call public wrapper to set state (only works on test contract)
      await (await forceMove.setStatus(channelId, fingerprint)).wait();
      expect(await forceMove.statusOf(channelId)).to.equal(fingerprint);

      const tx = forceMove.checkpoint(fixedPart, proof, candidate);
      if (reason) {
        await expectRevert(() => tx, reason);
      } else {
        const receipt = await (await tx).wait();
        const event = receipt.events.pop();

        expect(event.event).to.equal(isChallenged ? 'ChallengeCleared' : 'Checkpointed');
        const expectedEvent = {
          0: channelId,
          1: largestTurnNum,
          channelId: channelId,
          newTurnNumRecord: largestTurnNum,
        };
        for (const [key, value] of Object.entries(expectedEvent)) {
          expect(event.args[key]).to.equal(value);
        }

        const expectedChannelStorageHash = channelDataToStatus({
          turnNumRecord: largestTurnNum,
          finalizesAt: 0x0,
        });

        // Check channelStorageHash against the expected value
        expect(await forceMove.statusOf(channelId)).to.equal(expectedChannelStorageHash);
      }
    })
  ;
});
