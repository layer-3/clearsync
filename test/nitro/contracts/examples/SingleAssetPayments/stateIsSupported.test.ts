import { Allocation, AllocationType } from '@statechannels/exit-format';
import { BigNumber, Contract } from 'ethers';
import { ethers } from 'hardhat';
import { before, beforeEach, describe, it } from 'mocha';

import { expectRevert } from '../../../../helpers/expect-revert';
import { Outcome, encodeGuaranteeData } from '../../../../../src/nitro/contract/outcome';
import {
  State,
  getFixedPart,
  getVariablePart,
  separateProofAndCandidate,
} from '../../../../../src/nitro/contract/state';
import {
  generateParticipants,
  randomExternalDestination,
  setupContract,
} from '../../../test-helpers';
import {
  AssetOutcomeShortHand,
  bindSignaturesWithSignedByBitfield,
  getRandomNonce,
  signStates,
} from '../../../../../src/nitro';
import { INVALID_SIGNED_BY } from '../../../../../src/nitro/contract/transaction-creators/revert-reasons';
import { expectSupportedState } from '../../../tx-expect-wrappers';
import { replaceAddressesAndBigNumberify } from '../../../../../src/nitro/helpers';

import type { SingleAssetPayments } from '../../../../../typechain-types';

const { HashZero } = ethers.constants;

let singleAssetPayments: Contract;

const addresses = {
  // Participants
  A: randomExternalDestination(),
  B: randomExternalDestination(),
};

const nParticipants = 2;
const { wallets, participants } = generateParticipants(nParticipants);

const challengeDuration = 0x1_00;
const guaranteeData = { left: addresses.A, right: addresses.B };
let channelNonce = getRandomNonce('SingleAssetPayments');

const whoSignedWhatA = [1, 0];
const whoSignedWhatB = [0, 1];

before(async () => {
  singleAssetPayments = await setupContract<SingleAssetPayments>('SingleAssetPayments');
});

describe('stateIsSupported', () => {
  beforeEach(() => (channelNonce = BigNumber.from(channelNonce).add(1).toHexString()));

  const testCases = [
    {
      numAssets: [1, 1],
      isAllocation: [true, true],
      balancesA: { A: 1, B: 1 },
      turnNums: [3, 4],
      balancesB: { A: 0, B: 2 },
      whoSignedWhat: whoSignedWhatA,
      reason: undefined,
      description: 'A pays B 1 wei',
    },
    {
      numAssets: [1, 1],
      isAllocation: [true, true],
      balancesA: { A: 1, B: 1 },
      turnNums: [2, 3],
      balancesB: { A: 2, B: 0 },
      whoSignedWhat: whoSignedWhatB,
      reason: undefined,
      description: 'B pays A 1 wei',
    },
    {
      numAssets: [1, 1],
      isAllocation: [true, true],
      balancesA: { A: 1, B: 1 },
      turnNums: [2, 3],
      balancesB: { A: 0, B: 2 },
      whoSignedWhat: whoSignedWhatA,
      reason: INVALID_SIGNED_BY,
      description: 'A pays B 1 wei (not their move)',
    },
    {
      numAssets: [1, 1],
      isAllocation: [false, false],
      balancesA: { A: 1, B: 1 },
      turnNums: [3, 4],
      balancesB: { A: 0, B: 2 },
      whoSignedWhat: whoSignedWhatA,
      reason: 'not a simple allocation',
      description: 'Guarantee',
    },
    {
      numAssets: [1, 1],
      isAllocation: [true, true],
      balancesA: { A: 1, B: 1 },
      turnNums: [3, 4],
      balancesB: { A: 1, B: 2 },
      whoSignedWhat: whoSignedWhatA,
      reason: 'Total allocated cannot change',
      description: 'Total amounts increase',
    },
    {
      numAssets: [2, 2],
      isAllocation: [true, true],
      balancesA: { A: 1, B: 1 },
      turnNums: [3, 4],
      balancesB: { A: 2, B: 0 },
      whoSignedWhat: whoSignedWhatA,
      reason: 'outcome: Only one asset allowed',
      description: 'More than one asset',
    },
  ];

  for (const tc of testCases) it(tc.description, async () => {
      const { isAllocation, numAssets, turnNums, whoSignedWhat, reason } = tc as unknown as {
        isAllocation: boolean[];
        numAssets: number[];
        balancesA: AssetOutcomeShortHand;
        turnNums: number[];
        balancesB: AssetOutcomeShortHand;
        whoSignedWhat: number[];
        reason?: string;
      };
      let balancesA = tc.balancesA as AssetOutcomeShortHand;
      balancesA = replaceAddressesAndBigNumberify(balancesA, addresses) as AssetOutcomeShortHand;
      const allocationsA: Allocation[] = [];
      for (const key of Object.keys(balancesA)) allocationsA.push({
          destination: key,
          amount: balancesA[key].toString(),
          allocationType: isAllocation[0] ? AllocationType.simple : AllocationType.guarantee,
          metadata: isAllocation[0] ? '0x' : encodeGuaranteeData(guaranteeData),
        })
      ;
      const outcomeA: Outcome = [
        {
          asset: ethers.constants.AddressZero,
          assetMetadata: { assetType: 0, metadata: '0x' },
          allocations: allocationsA,
        },
      ];

      if (numAssets[0] === 2) {
        outcomeA.push(outcomeA[0]);
      }
      let balancesB = tc.balancesB as AssetOutcomeShortHand;
      balancesB = replaceAddressesAndBigNumberify(balancesB, addresses) as AssetOutcomeShortHand;
      const allocationsB: Allocation[] = [];

      for (const key of Object.keys(balancesB)) allocationsB.push({
          destination: key,
          amount: balancesB[key].toString(),
          allocationType: isAllocation[1] ? AllocationType.simple : AllocationType.guarantee,
          metadata: isAllocation[1] ? '0x' : encodeGuaranteeData(guaranteeData),
        })
      ;

      const outcomeB: Outcome = [
        {
          asset: ethers.constants.AddressZero,
          assetMetadata: { assetType: 0, metadata: '0x' },
          allocations: allocationsB,
        },
      ];

      if (numAssets[1] === 2) {
        outcomeB.push(outcomeB[0]);
      }

      const states: State[] = [
        {
          turnNum: turnNums[0],
          isFinal: false,
          channelNonce,
          participants,
          challengeDuration,
          outcome: outcomeA,
          appData: HashZero,
          appDefinition: singleAssetPayments.address,
        },
        {
          turnNum: turnNums[1],
          isFinal: false,
          channelNonce,
          participants,
          challengeDuration,
          outcome: outcomeB,
          appData: HashZero,
          appDefinition: singleAssetPayments.address,
        },
      ];
      const fixedPart = getFixedPart(states[0]);
      const variableParts = states.map((s) => getVariablePart(s));

      // Sign the states
      const signatures = await signStates(states, wallets, whoSignedWhat);
      const recoveredVariableParts = bindSignaturesWithSignedByBitfield(
        variableParts,
        signatures,
        whoSignedWhat,
      );

      const { proof, candidate } = separateProofAndCandidate(recoveredVariableParts);

      if (reason) {
        await expectRevert(
          () => singleAssetPayments.stateIsSupported(fixedPart, proof, candidate),
          reason,
        );
      } else {
        await expectSupportedState(() =>
          singleAssetPayments.stateIsSupported(fixedPart, proof, candidate),
        );
      }
    })
  ;
});
