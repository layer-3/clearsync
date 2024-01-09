import { Allocation, AllocationType } from '@statechannels/exit-format';
import { BigNumber, Contract, utils } from 'ethers';
import { ethers } from 'hardhat';
import { describe, before, beforeEach, it } from 'mocha';

const { HashZero } = ethers.constants;
import { expectRevert } from '../../../../helpers/expect-revert';
import {
  AssetOutcomeShortHand,
  bindSignaturesWithSignedByBitfield,
  Bytes32,
  getRandomNonce,
  signStates,
} from '../../../../../src/nitro';
import type { Outcome } from '../../../../../src/nitro/contract/outcome';
import {
  getFixedPart,
  getVariablePart,
  separateProofAndCandidate,
  State,
} from '../../../../../src/nitro/contract/state';
import type { Bytes } from '../../../../../src/nitro/contract/types';
import {
  randomExternalDestination,
  setupContract,
  generateParticipants,
} from '../../../test-helpers';
import { expectSupportedState } from '../../../tx-expect-wrappers';
import { replaceAddressesAndBigNumberify } from '../../../../../src/nitro/helpers';
import type { HashLockedSwap } from '../../../../../typechain-types';

// Utilities
// TODO: move to a src file
interface HashLockedSwapData {
  h: Bytes32;
  preImage: Bytes;
}

function encodeHashLockedSwapData(data: HashLockedSwapData): string {
  return utils.defaultAbiCoder.encode(['tuple(bytes32 h, bytes preImage)'], [data]);
}
// *****

let hashTimeLock: Contract;

const addresses = {
  // Participants
  Sender: randomExternalDestination(),
  Receiver: randomExternalDestination(),
};

const nParticipants = 2;
const { wallets, participants } = generateParticipants(nParticipants);

const challengeDuration = 0x100;
const whoSignedWhat = [1, 0];

before(async () => {
  hashTimeLock = await setupContract<HashLockedSwap>('HashLockedSwap');
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

type testParams = {
  description: string;
  isValid: boolean;
  dataA: HashLockedSwapData;
  balancesA: AssetOutcomeShortHand;
  turnNumB: number;
  dataB: HashLockedSwapData;
  balancesB: AssetOutcomeShortHand;
};

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

  testCases.forEach((tc) =>
    it(tc.description, async () => {
      const { isValid, dataA, turnNumB, dataB } = tc as unknown as testParams;
      let { balancesA, balancesB } = tc as unknown as testParams;
      const turnNumA = turnNumB - 1;
      balancesA = replaceAddressesAndBigNumberify(balancesA, addresses) as AssetOutcomeShortHand;
      const allocationsA: Allocation[] = [];
      Object.keys(balancesA).forEach((key) =>
        allocationsA.push({
          destination: key,
          amount: balancesA[key].toString(),
          allocationType: AllocationType.simple,
          metadata: '0x',
        }),
      );
      const outcomeA: Outcome = [
        {
          asset: ethers.constants.AddressZero,
          allocations: allocationsA,
          assetMetadata: { assetType: 0, metadata: '0x' },
        },
      ];
      balancesB = replaceAddressesAndBigNumberify(balancesB, addresses) as AssetOutcomeShortHand;
      const allocationsB: Allocation[] = [];
      Object.keys(balancesB).forEach((key) =>
        allocationsB.push({
          destination: key,
          amount: balancesB[key].toString(),
          allocationType: AllocationType.simple,
          metadata: '0x',
        }),
      );
      const outcomeB: Outcome = [
        {
          asset: ethers.constants.AddressZero,
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
          appDefinition: hashTimeLock.address,
        },
        {
          turnNum: turnNumB,
          isFinal: false,
          channelNonce,
          participants,
          challengeDuration,
          outcome: outcomeB,
          appData: encodeHashLockedSwapData(dataB),
          appDefinition: hashTimeLock.address,
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
          hashTimeLock.stateIsSupported(fixedPart, proof, candidate),
        );
      } else {
        await expectRevert(
          () => hashTimeLock.stateIsSupported(fixedPart, proof, candidate),
          'incorrect preimage',
        );
      }
    }),
  );
});
