import { BigNumber, Contract, Wallet } from 'ethers';

import { describe, before, beforeEach, it } from 'mocha';
import { expectRevert } from '../../../../helpers/expect-revert';
import { generateParticipants, setupContract } from '../../../test-helpers';
import type { TESTConsensus } from '../../../../../typechain-types';
import {
  getFixedPart,
  getRandomNonce,
  Outcome,
  shortenedToRecoveredVariableParts,
  State,
} from '../../../../../src/nitro';
import {
  NOT_UNANIMOUS,
  PROOF_SUPPLIED,
} from '../../../../../src/nitro/contract/transaction-creators/revert-reasons';
import { separateProofAndCandidate } from '../../../../../src/nitro/contract/state';
import { expectSucceedWithNoReturnValues } from '../../../tx-expect-wrappers';

let consensusApp: Contract & TESTConsensus;

const challengeDuration = 0x1000;
const asset = Wallet.createRandom().address;
const defaultOutcome: Outcome = [
  { asset, allocations: [], assetMetadata: { assetType: 0, metadata: '0x' } },
];

const nParticipants = 3;
const { participants } = generateParticipants(nParticipants);
let channelNonce = getRandomNonce('Consensus');

before(async () => {
  consensusApp = await setupContract<TESTConsensus>('TESTConsensus');
});

describe('requireConsensus', () => {
  beforeEach(() => (channelNonce = BigNumber.from(channelNonce).add(1).toHexString()));

  const testCases = [
    {
      description: 'accept when signed by all (one turnNum)',
      turnNumToShortenedVariablePart: new Map([[0, [0, 1, 2]]]),
      reason: undefined,
    },
    {
      description: 'accept when signed by all (other turnNum)',
      turnNumToShortenedVariablePart: new Map([[2, [0, 1, 2]]]),
      reason: undefined,
    },
    {
      description: 'revert when not signed by all',
      turnNumToShortenedVariablePart: new Map([[0, [0, 1]]]),
      reason: NOT_UNANIMOUS,
    },
    {
      description: 'revert when not signed at all',
      turnNumToShortenedVariablePart: new Map([[0, []]]),
      reason: NOT_UNANIMOUS,
    },
    {
      description: 'revert when supplied proof state',
      turnNumToShortenedVariablePart: new Map([
        [0, [0]],
        [1, [0, 1, 2]],
      ]),
      reason: PROOF_SUPPLIED,
    },
  ];

  testCases.forEach((tc) =>
    it(tc.description, async () => {
      const { turnNumToShortenedVariablePart, reason } = tc;

      const state: State = {
        turnNum: 0,
        isFinal: false,
        participants,
        channelNonce,
        challengeDuration,
        outcome: defaultOutcome,
        appDefinition: consensusApp.address,
        appData: '0x',
      };

      const fixedPart = getFixedPart(state);

      const recoveredVP = shortenedToRecoveredVariableParts(turnNumToShortenedVariablePart);
      const { proof, candidate } = separateProofAndCandidate(recoveredVP);

      if (reason) {
        await expectRevert(() => consensusApp.requireConsensus(fixedPart, proof, candidate));
      } else {
        await expectSucceedWithNoReturnValues(() =>
          consensusApp.requireConsensus(fixedPart, proof, candidate),
        );
      }
    }),
  );
});
