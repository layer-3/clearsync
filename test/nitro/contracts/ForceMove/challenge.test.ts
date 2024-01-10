import { BigNumber, Contract, Signature, Wallet, ethers } from 'ethers';
import hre from 'hardhat';
import { before, describe, it } from 'mocha';
import { expect } from 'chai';

import { expectRevert } from '../../../helpers/expect-revert';
import { getChannelId } from '../../../../src/nitro/contract/channel';
import { ChannelData, channelDataToStatus } from '../../../../src/nitro/contract/channel-storage';
import {
  State,
  getFixedPart,
  getVariablePart,
  separateProofAndCandidate,
} from '../../../../src/nitro/contract/state';
import {
  CHALLENGER_NON_PARTICIPANT,
  CHANNEL_FINALIZED,
  COUNTING_APP_INVALID_TRANSITION,
  INVALID_NUMBER_OF_PROOF_STATES,
  INVALID_SIGNATURE,
  TURN_NUM_RECORD_DECREASED,
  TURN_NUM_RECORD_NOT_INCREASED,
} from '../../../../src/nitro/contract/transaction-creators/revert-reasons';
import { Outcome, SignedState, getRandomNonce } from '../../../../src/nitro/index';
import {
  bindSignatures,
  signChallengeMessage,
  signData,
  signState,
  signStates,
} from '../../../../src/nitro/signatures';
import {
  clearedChallengeFingerprint,
  finalizedFingerprint,
  largeOutcome,
  nonParticipant,
  ongoingChallengeFingerprint,
  parseOutcomeEventResult,
  setupContract,
} from '../../test-helpers';
import { NITRO_MAX_GAS, createChallengeTransaction } from '../../../../src/nitro/transactions';
import { hashChallengeMessage } from '../../../../src/nitro/contract/challenge';
import { MAX_OUTCOME_ITEMS } from '../../../../src/nitro/contract/outcome';

import type { transitionType } from './types';
import type { CountingApp, TESTForceMove } from '../../../../typechain-types';

const { HashZero } = ethers.constants;
const { defaultAbiCoder } = ethers.utils;

let forceMove: Contract & TESTForceMove;
let countingApp: Contract & CountingApp;

const participants = ['', '', ''];
const wallets = Array.from({ length: 3 });
const challengeDuration = 86_400; // 1 day
const outcome: Outcome = [
  {
    allocations: [],
    asset: Wallet.createRandom().address,
    assetMetadata: { assetType: 0, metadata: '0x' },
  },
];
const keys = [
  '0x8624ebe7364bb776f891ca339f0aaa820cc64cc9fca6a28eec71e6d8fc950f29',
  '0x275a2e2cd9314f53b42246694034a80119963097e3adf495fbf6d821dc8b6c8e',
  '0x1b7598002c59e7d9131d7e7c9d0ec48ed065a3ed04af56674497d6b0048f2d84',
];

// Populate wallets and participants array
for (let i = 0; i < 3; i++) {
  wallets[i] = new Wallet(keys[i]);
  participants[i] = wallets[i].address;
}

before(async () => {
  forceMove = await setupContract<TESTForceMove>('TESTForceMove');
  countingApp = await setupContract<CountingApp>('CountingApp');
});

describe('challenge', () => {
  const threeStates = { appDatas: [0, 1, 2], whoSignedWhat: [0, 1, 2] };
  const fourStates = { appDatas: [0, 1, 2, 3], whoSignedWhat: [0, 1, 2, 0] };
  const oneState = { appDatas: [2], whoSignedWhat: [0, 0, 0] };
  const invalid = { appDatas: [0, 2, 1], whoSignedWhat: [0, 1, 2] };
  const largestTurnNum = 8;
  const isFinalCount = 0;
  const challenger = wallets[2];
  const empty = HashZero; // Equivalent to openAtZero
  const openAtFive = clearedChallengeFingerprint(5);
  const openAtLargestTurnNum = clearedChallengeFingerprint(largestTurnNum);
  const openAtTwenty = clearedChallengeFingerprint(20);
  const challengeAtFive = ongoingChallengeFingerprint(5);
  const challengeAtLargestTurnNum = ongoingChallengeFingerprint(largestTurnNum);
  const challengeAtTwenty = ongoingChallengeFingerprint(20);
  const finalizedAtFive = finalizedFingerprint(5);

  let channelNonce = getRandomNonce('challenge');
  beforeEach(() => (channelNonce = BigNumber.from(channelNonce).add(1).toHexString()));

  // Scenarios are synonymous with channelNonce.
  // For the purposes of this test, participants are fixed, making channelId 1-1 with channelNonce.
  const testCases = [
    {
      description:
        'It accepts for an open channel, and updates storage correctly, ' +
        'when the slot is empty, 3 states submitted',
      initialFingerprint: empty,
      stateData: threeStates,
      challengeSignatureType: 'correct',
      reasonString: undefined,
    },
    {
      description:
        'It accepts for an open channel, and updates storage correctly, ' +
        'when the slot is not empty, 3 states submitted',
      initialFingerprint: openAtFive,
      stateData: threeStates,
      challengeSignatureType: 'correct',
      reasonString: undefined,
    },
    {
      description:
        'It accepts for an open channel, and updates storage correctly, ' +
        'when the slot is not empty, 3 states submitted, open at largestTurnNum',
      initialFingerprint: openAtLargestTurnNum,
      stateData: threeStates,
      challengeSignatureType: 'correct',
      reasonString: undefined,
    },
    {
      description:
        'It accepts when a challenge is present, and updates storage correctly, ' +
        'when the turnNumRecord increases, 3 states',
      initialFingerprint: challengeAtFive,
      stateData: threeStates,
      challengeSignatureType: 'correct',
      reasonString: undefined,
    },
    {
      description: 'It reverts for an open channel if ' + 'the turnNumRecord does not increase',
      initialFingerprint: openAtTwenty,
      stateData: threeStates,
      challengeSignatureType: 'correct',
      reasonString: TURN_NUM_RECORD_DECREASED,
    },
    {
      description: 'It reverts for an open channel if ' + 'the challengerSig is incorrect',
      initialFingerprint: empty,
      stateData: threeStates,
      challengeSignatureType: 'incorrect',
      reasonString: CHALLENGER_NON_PARTICIPANT,
    },
    {
      description: 'It reverts for an open channel if ' + 'the challengerSig is invalid',
      initialFingerprint: empty,
      stateData: threeStates,
      challengeSignatureType: 'invalid',
      reasonString: INVALID_SIGNATURE,
    },
    {
      description:
        'It reverts for an open channel if ' + 'the states do not form a validTransition chain',
      initialFingerprint: empty,
      stateData: invalid,
      challengeSignatureType: 'correct',
      reasonString: COUNTING_APP_INVALID_TRANSITION,
    },
    {
      description: 'It reverts when a challenge is present if the turnNumRecord does not increase',
      initialFingerprint: challengeAtTwenty,
      stateData: threeStates,
      challengeSignatureType: 'correct',
      reasonString: TURN_NUM_RECORD_NOT_INCREASED,
    },
    {
      description: 'It reverts when a challenge is present if the turnNumRecord does not increase',
      initialFingerprint: challengeAtLargestTurnNum,
      stateData: threeStates,
      challengeSignatureType: 'correct',
      reasonString: TURN_NUM_RECORD_NOT_INCREASED,
    },
    {
      description: 'It reverts when the channel is finalized',
      initialFingerprint: finalizedAtFive,
      stateData: threeStates,
      challengeSignatureType: 'correct',
      reasonString: CHANNEL_FINALIZED,
    },
    {
      description: 'It reverts when too few states are submitted',
      initialFingerprint: empty,
      stateData: oneState,
      challengeSignatureType: 'correct',
      reasonString: INVALID_NUMBER_OF_PROOF_STATES,
    },
    {
      description: 'It reverts when too many states are submitted',
      initialFingerprint: empty,
      stateData: fourStates,
      challengeSignatureType: 'correct',
      reasonString: INVALID_NUMBER_OF_PROOF_STATES,
    },
  ];

  for (const tc of testCases)
    it(tc.description, async () => {
      const { reasonString, challengeSignatureType, stateData, initialFingerprint } =
        tc as unknown as {
          initialFingerprint: string;
          stateData: transitionType;
          challengeSignatureType: string;
          reasonString: undefined | string;
        };
      const { appDatas, whoSignedWhat } = stateData;

      const states: State[] = appDatas.map((data, idx) => ({
        turnNum: largestTurnNum - appDatas.length + 1 + idx,
        isFinal: idx > appDatas.length - isFinalCount,
        participants,
        channelNonce,
        challengeDuration,
        outcome,
        appDefinition: countingApp.address,
        appData: defaultAbiCoder.encode(['uint256'], [data]),
      }));
      const variableParts = states.map((state) => getVariablePart(state));
      const fixedPart = getFixedPart(states[0]);
      const channelId = getChannelId(fixedPart);

      // Sign the states
      const signatures = await signStates(states, wallets, whoSignedWhat);
      const { proof, candidate } = separateProofAndCandidate(
        bindSignatures(variableParts, signatures, whoSignedWhat),
      );

      const challengeState: SignedState = {
        state: states.at(-1),
        signature: { v: 0, r: '', s: '', _vs: '', recoveryParam: 0 } as Signature,
      };

      const correctChallengeSignature = signChallengeMessage(
        [challengeState],
        challenger.privateKey,
      );
      let challengeSignature: ethers.Signature;

      switch (challengeSignatureType) {
        case 'incorrect': {
          challengeSignature = signChallengeMessageByNonParticipant([challengeState]);
          break;
        }
        case 'invalid': {
          challengeSignature = { v: 1, s: HashZero, r: HashZero } as ethers.Signature;
          break;
        }
        case 'correct':
        default: {
          challengeSignature = correctChallengeSignature;
        }
      }

      // Set current channelStorageHashes value
      await (await forceMove.setStatus(channelId, initialFingerprint)).wait();

      const tx = forceMove.challenge(fixedPart, proof, candidate, challengeSignature);
      if (reasonString) {
        await expectRevert(() => tx, reasonString);
      } else {
        const receipt = await (await tx).wait();
        const event = receipt.events.pop();

        // Catch ChallengeRegistered event
        const {
          channelId: eventChannelId,
          finalizesAt: eventFinalizesAt,

          proof: eventProof,
          candidate: eventCandidate,
        } = event.args;

        // Check this information is enough to respond
        expect(eventChannelId).to.equal(channelId);

        if (proof.length > 0) {
          const res = parseOutcomeEventResult(eventProof.at(-1).variablePart.outcome);
          expect(res).to.deep.equal(proof.at(-1).variablePart.outcome);
          expect(eventProof.at(-1).variablePart.appData).to.equal(
            proof.at(-1).variablePart.appData,
          );
        }

        const res = parseOutcomeEventResult(eventCandidate.variablePart.outcome);
        expect(res).to.deep.equal(candidate.variablePart.outcome);
        expect(eventCandidate.variablePart.appData).to.equal(candidate.variablePart.appData);

        const expectedChannelStorage: ChannelData = {
          turnNumRecord: largestTurnNum,
          finalizesAt: eventFinalizesAt,
          state: states.at(-1),
          outcome,
        };
        const expectedFingerprint = channelDataToStatus(expectedChannelStorage);

        // Check channelStorageHash against the expected value
        expect(await forceMove.statusOf(channelId)).to.equal(expectedFingerprint);
      }
    });
});

describe('challenge with transaction generator', () => {
  let twoPartyFixedPart: {
    channelNonce: string;
    participants: string[];
    appDefinition: string;
    challengeDuration: number;
  } = {
    participants: [],
  };

  beforeEach(async () => {
    twoPartyFixedPart = {
      channelNonce: '0x1',
      participants: [wallets[0].address, wallets[1].address],
      appDefinition: countingApp.address,
      challengeDuration,
    };
    await (await forceMove.setStatus(getChannelId(twoPartyFixedPart), HashZero)).wait();
  });

  const testCases = [
    {
      description: 'challenge(0,1) accepted',
      appData: [0, 1],
      outcome: [],
      turnNums: [1, 2],
      challenger: 1,
    },
    {
      description: 'challenge(1,2) accepted',
      appData: [0, 1],
      outcome: [],
      turnNums: [2, 3],
      challenger: 0,
    },
    {
      description: 'challenge(2,3) accepted, MAX_OUTCOME_ITEMS',
      appData: [0, 1],
      outcome: largeOutcome(MAX_OUTCOME_ITEMS),
      turnNums: [3, 4],
      challenger: 0,
    },
  ];

  // FIX: even if dropping channel status before each test, turn nums from prev tests are saved and can cause reverts
  for (const tc of testCases)
    it(tc.description, async () => {
      const { appData, turnNums, challenger } = tc as unknown as {
        appData: number[];
        turnNums: number[];
        challenger: number;
      };
      const transactionRequest: ethers.providers.TransactionRequest = createChallengeTransaction(
        [
          await createTwoPartySignedCountingAppState(appData[0], turnNums[0]),
          await createTwoPartySignedCountingAppState(appData[1], turnNums[1]),
        ],
        wallets[challenger].privateKey,
      );

      const signer = hre.ethers.provider.getSigner();
      const response = await signer.sendTransaction({
        to: forceMove.address,
        ...transactionRequest,
      });
      expect(BigNumber.from((await response.wait()).gasUsed).lt(BigNumber.from(NITRO_MAX_GAS))).to
        .be.true;
    });
});

async function createTwoPartySignedCountingAppState(
  appData: number,
  turnNum: number,
  outcome: Outcome = [],
) {
  return signState(
    {
      turnNum,
      isFinal: false,
      appDefinition: countingApp.address,
      appData: defaultAbiCoder.encode(['uint256'], [appData]),
      outcome,
      channelNonce: '0x1',
      participants: [wallets[0].address, wallets[1].address],
      challengeDuration: 0xf_ff,
    },
    wallets[turnNum % 2].privateKey,
  );
}

function signChallengeMessageByNonParticipant(signedStates: SignedState[]): Signature {
  if (signedStates.length === 0) {
    throw new Error('At least one signed state must be provided');
  }
  const challengeState = signedStates.at(-1).state;
  const challengeHash = hashChallengeMessage(challengeState);
  return signData(challengeHash, nonParticipant.privateKey);
}
