import { ethers, Wallet } from 'ethers';
import { describe, before, it } from 'mocha';
import { expect } from 'chai';

import { getRandomNonce, SignedState, State } from '../../../src/nitro';
import { MAX_OUTCOME_ITEMS } from '../../../src/nitro/contract/outcome';
import { signState } from '../../../src/nitro/signatures';
import {
  createCheckpointTransaction,
  createConcludeTransaction,
  createChallengeTransaction,
  createSignatureArguments,
  MAX_TX_DATA_SIZE,
} from '../../../src/nitro/transactions';
import { largeOutcome } from '../test-helpers';

const walletA = Wallet.createRandom();
const walletB = Wallet.createRandom();

// TODO use 3x participants to match other tests

const state: State = {
  turnNum: 0,
  isFinal: false,
  appDefinition: ethers.constants.AddressZero,
  appData: '0x00',
  outcome: [],
  channelNonce: getRandomNonce('transactions'),
  participants: [walletA.address, walletB.address], // 2 participants is the most common usecase

  challengeDuration: 0x1,
};
let signedStateA: SignedState;
let signedStateB: SignedState;
const stateWithLargeOutcome = { ...state, outcome: largeOutcome(MAX_OUTCOME_ITEMS) };

before(async () => {
  signedStateA = signState(state, walletA.privateKey);
  signedStateB = signState(state, walletB.privateKey);
});

describe('transaction-generators', () => {
  it('creates a challenge transaction with MAX_OUTCOME_ITEMS outcome items that is smaller than MAX_TX_DATA_SIZE', async () => {
    const transactionRequest: ethers.providers.TransactionRequest = createChallengeTransaction(
      [
        signState(stateWithLargeOutcome, walletA.privateKey),
        signState(stateWithLargeOutcome, walletB.privateKey),
      ],
      walletA.privateKey,
    );
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    expect(transactionRequest.data!.toString().slice(2).length / 2).to.be.lt(MAX_TX_DATA_SIZE); // it's a hex string, so divide by 2 for bytes
  });

  it('creates a challenge transaction', async () => {
    const transactionRequest: ethers.providers.TransactionRequest = createChallengeTransaction(
      [signedStateA, signedStateB],
      walletA.privateKey,
    );

    expect(transactionRequest.data).to.exist;
  });

  it('creates a conclude from open transaction', async () => {
    const transactionRequest: ethers.providers.TransactionRequest = createConcludeTransaction([
      signedStateA,
      signedStateB,
    ]);

    expect(transactionRequest.data).to.exist;
  });

  it('creates a conclude from challenged transaction', async () => {
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

  testCases.forEach((tc) =>
    it(`creates a correct signature arguments when handling multiple states (turnNum=${tc.turnNum}, expectedWhoSignedWhat=${tc.expectedWhoSignedWhat})`, async () => {
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
            participants: [walletA.address, wallet2.address], // 2 participants is the most common usecase
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
            participants: [walletA.address, wallet2.address], // 2 participants is the most common usecase
            challengeDuration: 0x0,
          },
          turnNum[1] % 2 === 0 ? walletA.privateKey : wallet2.privateKey,
        ),
      ];
      const { states, signatures, whoSignedWhat } = createSignatureArguments(signedStates);

      expect(states).to.have.lengthOf(2);
      expect(signatures).to.have.lengthOf(2);
      expect(whoSignedWhat).to.deep.equal(expectedWhoSignedWhat);
    }),
  );

  describe('checkpoint transactions', () => {
    it('creates a transaction when there is a challenge state', async () => {
      const transactionRequest: ethers.providers.TransactionRequest = createCheckpointTransaction([
        signedStateA,
        signedStateB,
      ]);

      expect(transactionRequest.data).to.exist;
    });

    it('creates a transaction when the channel is open', async () => {
      const transactionRequest: ethers.providers.TransactionRequest = createCheckpointTransaction([
        signedStateA,
        signedStateB,
      ]);

      expect(transactionRequest.data).to.exist;
    });
  });
});
