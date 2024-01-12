import { Wallet, ethers } from 'ethers';
import { before, describe, it } from 'mocha';
import { expect } from 'chai';

import { SignedState, State, getRandomNonce } from '../../../src/nitro';
import { MAX_OUTCOME_ITEMS } from '../../../src/nitro/contract/outcome';
import { signState } from '../../../src/nitro/signatures';
import {
  MAX_TX_DATA_SIZE,
  createChallengeTransaction,
  createCheckpointTransaction,
  createConcludeTransaction,
  createSignatureArguments,
} from '../../../src/nitro/transactions';
import { largeOutcome } from '../test-helpers';

const walletA = Wallet.createRandom();
const walletB = Wallet.createRandom();

// TODO: use 3x participants to match other tests

const state: State = {
  turnNum: 0,
  isFinal: false,
  appDefinition: ethers.constants.AddressZero,
  appData: '0x00',
  outcome: [],
  channelNonce: getRandomNonce('transactions'),
  participants: [walletA.address, walletB.address], // 2 participants is the most common use case

  challengeDuration: 0x1,
};
let signedStateA: SignedState;
let signedStateB: SignedState;
const stateWithLargeOutcome = { ...state, outcome: largeOutcome(MAX_OUTCOME_ITEMS) };

before(() => {
  signedStateA = signState(state, walletA.privateKey);
  signedStateB = signState(state, walletB.privateKey);
});

describe('transaction-generators', () => {
  it('creates a challenge transaction with MAX_OUTCOME_ITEMS outcome items that is smaller than MAX_TX_DATA_SIZE', () => {
    const transactionRequest: ethers.providers.TransactionRequest = createChallengeTransaction(
      [
        signState(stateWithLargeOutcome, walletA.privateKey),
        signState(stateWithLargeOutcome, walletB.privateKey),
      ],
      walletA.privateKey,
    );
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion, @typescript-eslint/no-base-to-string
    expect(transactionRequest.data!.toString().slice(2).length / 2).to.be.lt(MAX_TX_DATA_SIZE); // it's a hex string, so divide by 2 for bytes
  });

  it('creates a challenge transaction', () => {
    const transactionRequest: ethers.providers.TransactionRequest = createChallengeTransaction(
      [signedStateA, signedStateB],
      walletA.privateKey,
    );

    expect(transactionRequest.data).to.exist;
  });

  it('creates a conclude from open transaction', () => {
    const transactionRequest: ethers.providers.TransactionRequest = createConcludeTransaction([
      signedStateA,
      signedStateB,
    ]);

    expect(transactionRequest.data).to.exist;
  });

  it('creates a conclude from challenged transaction', () => {
    const transactionRequest: ethers.providers.TransactionRequest = createConcludeTransaction([
      signedStateA,
      signedStateB,
    ]);

    expect(transactionRequest.data).to.exist;
  });

  const testCases = [
    {
      turnNum: [0, 1],
      expectedWhoSignedWhat: [0, 1],
    },
    {
      turnNum: [1, 2],
      expectedWhoSignedWhat: [1, 0],
    },
  ];

  for (const tc of testCases)
    it(`creates a correct signature arguments when handling multiple states (turnNum=${String(
      tc.turnNum,
    )}, expectedWhoSignedWhat=${String(tc.expectedWhoSignedWhat)})`, () => {
      const { turnNum, expectedWhoSignedWhat } = tc;
      const wallet2 = Wallet.createRandom();

      const signedStates = [
        signState(
          {
            turnNum: turnNum[0],
            isFinal: false,
            appDefinition: ethers.constants.AddressZero,
            appData: '0x00',
            outcome: [],
            channelNonce: getRandomNonce('transactions'),
            participants: [walletA.address, wallet2.address], // 2 participants is the most common use case
            challengeDuration: 0x0,
          },
          turnNum[0] % 2 === 0 ? walletA.privateKey : wallet2.privateKey,
        ),
        signState(
          {
            turnNum: turnNum[1],
            isFinal: false,
            appDefinition: ethers.constants.AddressZero,
            appData: '0x00',
            outcome: [],
            channelNonce: getRandomNonce('transactions'),
            participants: [walletA.address, wallet2.address], // 2 participants is the most common use case
            challengeDuration: 0x0,
          },
          turnNum[1] % 2 === 0 ? walletA.privateKey : wallet2.privateKey,
        ),
      ];
      const { states, signatures, whoSignedWhat } = createSignatureArguments(signedStates);

      expect(states).to.have.lengthOf(2);
      expect(signatures).to.have.lengthOf(2);
      expect(whoSignedWhat).to.deep.equal(expectedWhoSignedWhat);
    });

  describe('checkpoint transactions', () => {
    it('creates a transaction when there is a challenge state', () => {
      const transactionRequest: ethers.providers.TransactionRequest = createCheckpointTransaction([
        signedStateA,
        signedStateB,
      ]);

      expect(transactionRequest.data).to.exist;
    });

    it('creates a transaction when the channel is open', () => {
      const transactionRequest: ethers.providers.TransactionRequest = createCheckpointTransaction([
        signedStateA,
        signedStateB,
      ]);

      expect(transactionRequest.data).to.exist;
    });
  });
});
