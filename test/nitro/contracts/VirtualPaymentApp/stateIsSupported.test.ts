import { BigNumber, Contract, Wallet } from 'ethers';
import { ethers } from 'hardhat';
import { beforeEach, describe, it } from 'mocha';
import { expect } from 'chai';

import {
  MAGIC_ADDRESS_INDICATING_ETH,
  Voucher,
  computeOutcome,
  convertAddressToBytes32,
  encodeVoucherAmountAndSignature,
  getChannelId,
  signVoucher,
} from '../../../../src/nitro';
import {
  FixedPart,
  RecoveredVariablePart,
  State,
  getFixedPart,
  getVariablePart,
} from '../../../../src/nitro/contract/state';
import { generateParticipants, setupContract } from '../../test-helpers';

import type { VirtualPaymentApp } from '../../../../typechain-types';

const { HashZero } = ethers.constants;

const PANIC_ARITHMETIC = '0x11';

let virtualPaymentApp: Contract;
let participantsWallets: Wallet[];
let baseState: State;
let fixedPart: FixedPart;
let channelId: string;
let alice: string;
let bob: string;

beforeEach(async () => {
  virtualPaymentApp = await setupContract<VirtualPaymentApp>('VirtualPaymentApp');

  const nParticipants = 4;
  const { wallets, participants } = generateParticipants(nParticipants);
  participantsWallets = wallets;

  baseState = {
    turnNum: 0,
    isFinal: false,
    channelNonce: '0x8',
    participants,
    challengeDuration: 0x1_00,
    outcome: [],
    appData: HashZero,
    appDefinition: virtualPaymentApp.address,
  };
  fixedPart = getFixedPart(baseState);
  channelId = getChannelId(fixedPart);

  // NOTE these destinations do not necessarily need to be related to participant addresses
  alice = convertAddressToBytes32(participants[0]);
  bob = convertAddressToBytes32(participants[2]);
});

describe('stateIsSupported (lone candidate route)', () => {
  interface TestCase {
    turnNum: number;
    isFinal: boolean;
    reason?: string;
  }

  const testcases: TestCase[] = [
    { turnNum: 0, isFinal: false, reason: undefined },
    { turnNum: 1, isFinal: false, reason: undefined },
    { turnNum: 2, isFinal: false, reason: 'bad candidate turnNum; |proof|=0' },
    { turnNum: 4, isFinal: false, reason: 'bad candidate turnNum; |proof|=0' },
  ];

  for (const tc of testcases) {
    it(`${tc.reason ? 'reverts        ' : 'does not revert'} for unanimous consensus on ${
      tc.isFinal ? 'final' : 'nonfinal'
    } state with turnNum ${tc.turnNum}`, async () => {
      const state: State = {
        ...baseState,
        turnNum: tc.turnNum,
        isFinal: tc.isFinal,
      };

      const variablePart = getVariablePart(state);

      const candidate: RecoveredVariablePart = {
        variablePart,
        signedBy: BigNumber.from(0b1111).toHexString(),
      };

      if (tc.reason) {
        await expect(
          virtualPaymentApp.stateIsSupported(fixedPart, [], candidate),
        ).to.be.revertedWith(tc.reason);
      } else {
        await virtualPaymentApp.stateIsSupported(fixedPart, [], candidate);
      }
    });
  }
});

// eslint-disable-next-line sonarjs/cognitive-complexity
describe('stateIsSupported (candidate plus single proof state route)', () => {
  interface TestCase {
    proofTurnNum: number;
    candidateTurnNum: number;
    unanimityOnProof: boolean;
    bobSignedCandidate: boolean;
    voucherForThisChannel: boolean;
    voucherSignedByAlice: boolean;
    aliceAdjustedCorrectly: boolean;
    bobAdjustedCorrectly: boolean;
    nativeAsset: boolean;
    multipleAssets: boolean;
    aliceUnderflow: boolean;
    reason?: string;
  }

  const vVR: TestCase = {
    // valid voucher redemption
    proofTurnNum: 1,
    candidateTurnNum: 2,
    unanimityOnProof: true,
    bobSignedCandidate: true,
    voucherForThisChannel: true,
    voucherSignedByAlice: true,
    aliceAdjustedCorrectly: true,
    bobAdjustedCorrectly: true,
    nativeAsset: true,
    multipleAssets: false,
    aliceUnderflow: false,
    reason: undefined,
  };
  const testcases: TestCase[] = [
    vVR,
    { ...vVR, proofTurnNum: 0, reason: 'bad proof[0].turnNum; |proof|=1' },
    { ...vVR, unanimityOnProof: false, reason: 'postfund !unanimous; |proof|=1' },
    { ...vVR, bobSignedCandidate: false, reason: 'redemption not signed by Bob' },
    { ...vVR, voucherSignedByAlice: false, reason: 'invalid signature for voucher' },
    { ...vVR, voucherForThisChannel: false, reason: 'invalid signature for voucher' },
    { ...vVR, aliceAdjustedCorrectly: false, reason: 'Alice not adjusted correctly' },
    { ...vVR, bobAdjustedCorrectly: false, reason: 'Bob not adjusted correctly' },
    { ...vVR, nativeAsset: false, reason: 'only native asset allowed' },
    { ...vVR, multipleAssets: true, reason: 'only native asset allowed' },
    // we expect transaction to revert with a panic code 0x11 (Arithmetic operation underflowed or overflowed outside of an unchecked block)
    { ...vVR, aliceUnderflow: true, reason: PANIC_ARITHMETIC },
  ];

  for (const tc of testcases) {
    it(`${
      tc.reason ? 'reverts        ' : 'does not revert'
    } for a redemption transition with ${JSON.stringify(tc)}`, async () => {
      const proofState: State = {
        ...baseState,
        turnNum: tc.proofTurnNum,
        isFinal: false,
        outcome: computeOutcome({
          [MAGIC_ADDRESS_INDICATING_ETH]: { [alice]: 10, [bob]: 10 },
        }),
      };

      // construct voucher with the (in)appropriate channelId
      const amount = tc.aliceUnderflow
        ? BigNumber.from(999_999_999_999).toHexString() // much larger than Alice's original balance
        : BigNumber.from(7).toHexString();

      const voucher: Voucher = {
        channelId: tc.voucherForThisChannel
          ? channelId
          : convertAddressToBytes32(MAGIC_ADDRESS_INDICATING_ETH),
        amount,
      };

      // make an (in)valid signature
      const signature = await signVoucher(voucher, participantsWallets[0]);
      if (!tc.voucherSignedByAlice) {
        // (conditionally) corrupt the signature
        signature.s = signature.r;
      }

      // embed voucher into candidate state
      const encodedVoucherAmountAndSignature = encodeVoucherAmountAndSignature(amount, signature);
      const candidateState: State = {
        ...proofState,
        outcome: computeOutcome({
          [MAGIC_ADDRESS_INDICATING_ETH]: {
            [alice]: tc.aliceAdjustedCorrectly ? 3 : 2,
            [bob]: tc.bobAdjustedCorrectly ? 7 : 99,
          },
        }),
        turnNum: tc.candidateTurnNum,
        appData: encodedVoucherAmountAndSignature,
      };

      if (!tc.nativeAsset) {
        candidateState.outcome[0].asset = virtualPaymentApp.address;
      }
      if (tc.multipleAssets) {
        candidateState.outcome.push(candidateState.outcome[0]);
      }

      // Sign the proof state (should be everyone)
      const proof: RecoveredVariablePart[] = [
        {
          variablePart: getVariablePart(proofState),
          signedBy: BigNumber.from(tc.unanimityOnProof ? 0b1111 : 0b1011).toHexString(),
        },
      ];

      // Sign the candidate state (should be just Bob)
      const candidate: RecoveredVariablePart = {
        variablePart: getVariablePart(candidateState),
        signedBy: BigNumber.from(tc.bobSignedCandidate ? 0b1000 : 0b0000).toHexString(), // 0b1000 signed by Bob only
      };

      if (tc.reason) {
        if (tc.reason === PANIC_ARITHMETIC) {
          await expect(
            virtualPaymentApp.stateIsSupported(fixedPart, proof, candidate),
          ).to.be.revertedWithPanic(PANIC_ARITHMETIC);
        } else {
          await expect(
            virtualPaymentApp.stateIsSupported(fixedPart, proof, candidate),
          ).to.be.revertedWith(tc.reason);
        }
      } else {
        await virtualPaymentApp.stateIsSupported(fixedPart, proof, candidate);
      }
    });
  }
});

describe('stateIsSupported (longer proof state route)', () => {
  it(`reverts for |support|>1`, async () => {
    const variablePart = getVariablePart(baseState);

    const candidate: RecoveredVariablePart = {
      variablePart,
      signedBy: BigNumber.from(0b1111).toHexString(),
    };

    await expect(
      virtualPaymentApp.stateIsSupported(fixedPart, [candidate, candidate], candidate),
    ).to.be.revertedWith('bad proof length');
  });
});
