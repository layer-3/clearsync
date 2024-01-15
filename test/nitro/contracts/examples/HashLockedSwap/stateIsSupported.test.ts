import { Allocation, AllocationType } from '@statechannels/exit-format';
import { BigNumber, utils } from 'ethers';
import { ethers } from 'hardhat';
import { before, beforeEach, describe, it } from 'mocha';
import { expect } from 'chai';

import {
  AssetOutcomeShortHand,
  Bytes32,
  bindSignaturesWithSignedByBitfield,
  getRandomNonce,
  signStates,
} from '../../../../../src/nitro';
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
import { expectSupportedState } from '../../../tx-expect-wrappers';
import { replaceAddressesAndBigNumberify } from '../../../../../src/nitro/helpers';

import type { Bytes } from '../../../../../src/nitro/contract/types';
import type { Outcome } from '../../../../../src/nitro/contract/outcome';
import type { HashLockedSwap } from '../../../../../typechain-types';

const { HashZero, AddressZero } = ethers.constants;

// Utilities
// TODO: move to a src file
interface HashLockedSwapData {
  h: Bytes32;
  preImage: Bytes;
}

function encodeHashLockedSwapData(data: HashLockedSwapData): string {
  return utils.defaultAbiCoder.encode(['tuple(bytes32 h, bytes preImage)'], [data]);
}

let hashLockedSwap: HashLockedSwap;

const addresses = {
  // Participants
  Sender: randomExternalDestination(),
  Receiver: randomExternalDestination(),
};

const nParticipants = 2;
const { wallets, participants } = generateParticipants(nParticipants);

const challengeDuration = 0x1_00;
const whoSignedWhat = [1, 0];

before(async () => {
  hashLockedSwap = await setupContract<HashLockedSwap>('HashLockedSwap');
});

const preImage = '0xdeadbeef';
const conditionalPayment: HashLockedSwapData = {
  h: utils.sha256(preImage),
  // ^^^^ important field (SENDER)
  preImage: HashZero,
};

const correctPreImage: HashLockedSwapData = {
  preImage: preImage,
  // ^^^^ important field (RECEIVER)
  h: HashZero,
};

const incorrectPreImage: HashLockedSwapData = {
  preImage: '0xdeadc0de',
  // ^^^^ important field (RECEIVER)
  h: HashZero,
};

interface testParams {
  description: string;
  isValid: boolean;
  dataA: HashLockedSwapData;
  balancesA: AssetOutcomeShortHand;
  turnNumB: number;
  dataB: HashLockedSwapData;
  balancesB: AssetOutcomeShortHand;
}

describe('stateIsSupported', () => {
  let channelNonce = getRandomNonce('HashLockedSwap');
  beforeEach(() => (channelNonce = BigNumber.from(channelNonce).add(1).toHexString()));

  const testCases: testParams[] = [
    {
      description: 'Receiver unlocks the conditional payment',
      isValid: true,
      dataA: conditionalPayment,
      balancesA: { Sender: 1, Receiver: 0 },
      turnNumB: 4,
      dataB: correctPreImage,
      balancesB: { Sender: 0, Receiver: 1 },
    },
    {
      description: 'Receiver cannot unlock with incorrect preimage',
      isValid: false,
      dataA: conditionalPayment,
      balancesA: { Sender: 1, Receiver: 0 },
      turnNumB: 4,
      dataB: incorrectPreImage,
      balancesB: { Sender: 0, Receiver: 1 },
    },
  ];

  for (const tc of testCases)
    it(tc.description, async () => {
      const { isValid, dataA, turnNumB, dataB } = tc as unknown as testParams;
      let { balancesA, balancesB } = tc as unknown as testParams;
      const turnNumA = turnNumB - 1;
      balancesA = replaceAddressesAndBigNumberify(balancesA, addresses) as AssetOutcomeShortHand;
      const allocationsA: Allocation[] = [];
      for (const key of Object.keys(balancesA))
        allocationsA.push({
          destination: key,
          // balancesA[key] is a BigNumberish and `.toString()` will not return [object Object]
          // eslint-disable-next-line @typescript-eslint/no-base-to-string
          amount: balancesA[key].toString(),
          allocationType: AllocationType.simple,
          metadata: '0x',
        });
      const outcomeA: Outcome = [
        {
          asset: AddressZero,
          allocations: allocationsA,
          assetMetadata: { assetType: 0, metadata: '0x' },
        },
      ];
      balancesB = replaceAddressesAndBigNumberify(balancesB, addresses) as AssetOutcomeShortHand;
      const allocationsB: Allocation[] = [];
      for (const key of Object.keys(balancesB))
        allocationsB.push({
          destination: key,
          // balancesB[key] is a BigNumberish and `.toString()` will not return [object Object]
          // eslint-disable-next-line @typescript-eslint/no-base-to-string
          amount: balancesB[key].toString(),
          allocationType: AllocationType.simple,
          metadata: '0x',
        });
      const outcomeB: Outcome = [
        {
          asset: AddressZero,
          allocations: allocationsB,
          assetMetadata: { assetType: 0, metadata: '0x' },
        },
      ];
      const states: State[] = [
        {
          turnNum: turnNumA,
          isFinal: false,
          channelNonce,
          participants,
          challengeDuration,
          outcome: outcomeA,
          appData: encodeHashLockedSwapData(dataA),
          appDefinition: hashLockedSwap.address,
        },
        {
          turnNum: turnNumB,
          isFinal: false,
          channelNonce,
          participants,
          challengeDuration,
          outcome: outcomeB,
          appData: encodeHashLockedSwapData(dataB),
          appDefinition: hashLockedSwap.address,
        },
      ];
      const fixedPart = getFixedPart(states[0]);
      const variableParts = states.map((s) => getVariablePart(s));

      // Sign the states
      const signatures = await signStates(states, wallets, whoSignedWhat);
      const { proof, candidate } = separateProofAndCandidate(
        bindSignaturesWithSignedByBitfield(variableParts, signatures, whoSignedWhat),
      );

      if (isValid) {
        await expectSupportedState(() =>
          hashLockedSwap.stateIsSupported(fixedPart, proof, candidate),
        );
      } else {
        await expect(
          hashLockedSwap.stateIsSupported(fixedPart, proof, candidate),
        ).to.be.revertedWith('incorrect preimage');
      }
    });
});
