import { BigNumber, constants } from 'ethers';
import { ethers } from 'hardhat';
import { before, beforeEach, describe, it } from 'mocha';
import { expect } from 'chai';

import { getChannelId } from '../../../../src/nitro/contract/channel';
import {
  FixedPart,
  State,
  getFixedPart,
  getVariablePart,
  separateProofAndCandidate,
} from '../../../../src/nitro/contract/state';
import {
  generateParticipants,
  randomChannelId,
  randomExternalDestination,
  setupContract,
} from '../../test-helpers';
import {
  OutcomeShortHand,
  bindSignatures,
  channelDataToStatus,
  signStates,
} from '../../../../src/nitro';
import { MAGIC_ADDRESS_INDICATING_ETH, NITRO_MAX_GAS } from '../../../../src/nitro/transactions';
import {
  computeOutcome,
  getRandomNonce,
  replaceAddressesAndBigNumberify,
} from '../../../../src/nitro/helpers';

import type { CountingApp, TESTNitroAdjudicator, Token } from '../../../../typechain-types';
import type { Outcome } from '../../../../src/nitro/contract/outcome';

interface addressesT {
  [index: string]: string | undefined;
  At: string;
  Bt: string;
}

type payoutsT = Record<string, number>;

interface TestCase {
  description: string;
  outcomeShortHand: OutcomeShortHand;
  heldBefore: OutcomeShortHand;
  heldAfter: OutcomeShortHand;
  newOutcome: OutcomeShortHand;
  payouts: OutcomeShortHand;
  reasonString: string | undefined;
}

// eslint-disable-next-line sonarjs/cognitive-complexity
describe('concludeAndTransferAllAssets', () => {
  const nParticipants = 3;
  const { wallets, participants } = generateParticipants(nParticipants);

  const challengeDuration = 0x10_00;
  let countingApp: CountingApp;
  let testNitroAdjudicator: TESTNitroAdjudicator;
  let token: Token;

  const addresses: addressesT = {
    // Channels
    c: undefined,
    C: randomChannelId(),
    X: randomChannelId(),
    // Externals
    A: randomExternalDestination(),
    B: randomExternalDestination(),
    // // Externals preloaded with TOK (cheaper to pay to)
    At: randomExternalDestination(),
    Bt: randomExternalDestination(),
    // Asset Holders
    ETH: undefined,
    ETH2: undefined,
    ERC20: undefined,
  };

  const tenPayouts = { ERC20: {} as payoutsT };
  const fiftyPayouts = { ERC20: {} as payoutsT };
  const oneHundredPayouts = { ERC20: {} as payoutsT };
  for (let i = 0; i < 100; i++) {
    addresses[i.toString()] =
      '0x000000000000000000000000e0c3b40fdff77c786dd3737837887c85' +
      (0x23_92_fa_22 + i).toString(16); // they need to be distinct because JS objects
    if (i < 10) tenPayouts.ERC20[i.toString()] = 1;
    if (i < 50) fiftyPayouts.ERC20[i.toString()] = 1;
    if (i < 100) oneHundredPayouts.ERC20[i.toString()] = 1;
  }

  const oneState = {
    whoSignedWhat: [0, 0, 0],
    appData: [ethers.constants.HashZero],
  };
  const turnNumRecord = 5;
  let channelNonce = getRandomNonce('concludeAndTransferAllAssets');

  before(async () => {
    countingApp = await setupContract<CountingApp>('CountingApp');
    testNitroAdjudicator = await setupContract<TESTNitroAdjudicator>('TESTNitroAdjudicator');
    token = await setupContract<Token>('Token', '0x6B8B2958795a5E9c00A2E8D4B0b90b870cbAB4b6');

    addresses.ETH = MAGIC_ADDRESS_INDICATING_ETH;
    addresses.ERC20 = token.address;

    // Preload At and Bt with TOK
    const aliceTx = await token.transfer('0x' + addresses.At.slice(26), BigNumber.from(1));
    await aliceTx.wait();
    const bobTx = await token.transfer('0x' + addresses.Bt.slice(26), BigNumber.from(1));
    await bobTx.wait();
  });

  beforeEach(() => (channelNonce = BigNumber.from(channelNonce).add(1).toHexString()));

  // For the purposes of this test, participants are fixed, making channelId 1-1 with channelNonce
  const testCases: TestCase[] = [
    {
      description: '{ETH: {A: 1}}',
      outcomeShortHand: { ETH: { A: 1 } },
      heldBefore: { ETH: { c: 1 } },
      heldAfter: { ETH: { c: 0 } },
      newOutcome: {},
      payouts: { ETH: { A: 1 } },
      reasonString: undefined,
    },
    {
      description: '{ETH: {A: 1}}',
      outcomeShortHand: { ETH: { A: 1 } },
      heldBefore: { ETH: { c: 1 } },
      heldAfter: { ETH: { c: 0 } },
      newOutcome: {},
      payouts: { ETH: { A: 1 } },
      reasonString: undefined,
    },
    {
      description: '{ETH: {A: 1, B: 1}}',
      outcomeShortHand: { ETH: { A: 1, B: 1 } },
      heldBefore: { ETH: { c: 2 } },
      heldAfter: { ETH: { c: 0 } },
      newOutcome: {},
      payouts: { ETH: { A: 1, B: 1 } },
      reasonString: undefined,
    },
    {
      description: '{ERC20: {A: 1, B: 1}}',
      outcomeShortHand: { ERC20: { A: 1, B: 1 } },
      heldBefore: { ERC20: { c: 2 } },
      heldAfter: { ERC20: { c: 0 } },
      newOutcome: {},
      payouts: { ETH: { A: 1, B: 1 } },
      reasonString: undefined,
    },
    {
      description: '{ERC20: {A: 1}}',
      outcomeShortHand: { ERC20: { A: 1 } },
      heldBefore: { ERC20: { c: 1 } },
      heldAfter: { ERC20: { c: 0 } },
      newOutcome: {},
      payouts: { ERC20: { A: 1 } },
      reasonString: undefined,
    },
    {
      description: '{ERC20: {At: 1, Bt: 1}} (At and Bt already have some TOK)',
      outcomeShortHand: { ERC20: { At: 1, Bt: 1 } },
      heldBefore: { ERC20: { c: 2 } },
      heldAfter: { ERC20: { c: 0 } },
      newOutcome: {},
      payouts: { ERC20: { At: 1, Bt: 1 } },
      reasonString: undefined,
    },
    {
      description: '10 TOK payouts',
      outcomeShortHand: tenPayouts,
      heldBefore: { ERC20: { c: 10 } },
      heldAfter: { ERC20: { c: 0 } },
      newOutcome: {},
      payouts: tenPayouts,
      reasonString: undefined,
    },
    {
      description: '50 TOK payouts',
      outcomeShortHand: fiftyPayouts,
      heldBefore: { ERC20: { c: 50 } },
      heldAfter: { ERC20: { c: 0 } },
      newOutcome: {},
      payouts: fiftyPayouts,
      reasonString: undefined,
    },
    {
      description: '100 TOK payouts',
      outcomeShortHand: oneHundredPayouts,
      heldBefore: { ERC20: { c: 100 } },
      heldAfter: { ERC20: { c: 0 } },
      newOutcome: {},
      payouts: oneHundredPayouts,
      reasonString: undefined,
    },
  ];

  for (const tc of testCases) {
    it(tc.description, async () => {
      let outcomeShortHand = tc.outcomeShortHand;
      let heldBefore = tc.heldBefore;
      let heldAfter = tc.heldAfter;
      let newOutcome = tc.newOutcome;
      let payouts = tc.payouts;
      const reasonString = tc.reasonString;

      const fixedPart: FixedPart = {
        participants,
        channelNonce,
        appDefinition: countingApp.address,
        challengeDuration,
      };
      const channelId = getChannelId(fixedPart);
      addresses.c = channelId;
      const support = oneState;
      const { appData, whoSignedWhat } = support;
      const numStates = appData.length;
      const largestTurnNum = turnNumRecord + 1;

      // Transfer some tokens into the relevant AssetHolder
      // Do this step before transforming input data (easier)
      if ('ERC20' in heldBefore) {
        const increaseTx = await token.increaseAllowance(
          testNitroAdjudicator.address,
          heldBefore.ERC20.c,
        );
        await increaseTx.wait();

        const depositTx = await testNitroAdjudicator.deposit(
          token.address,
          channelId,
          '0x00',
          heldBefore.ERC20.c,
        );
        await depositTx.wait();
      }
      if ('ETH' in heldBefore) {
        const tx = await testNitroAdjudicator.deposit(
          MAGIC_ADDRESS_INDICATING_ETH,
          channelId,
          '0x00',
          heldBefore.ETH.c,
          {
            value: heldBefore.ETH.c,
          },
        );
        await tx.wait();
      }

      // Transform input data (unpack addresses and BigNumberify amounts)
      [heldBefore, outcomeShortHand, newOutcome, heldAfter, payouts] = [
        heldBefore,
        outcomeShortHand,
        newOutcome,
        heldAfter,
        payouts,
      ].map((object) => replaceAddressesAndBigNumberify(object, addresses) as OutcomeShortHand);

      // Compute the outcome.
      const outcome: Outcome = computeOutcome(outcomeShortHand);

      // Construct states
      const states: State[] = [];
      for (let i = 1; i <= numStates; i++) {
        states.push({
          isFinal: true,
          participants,
          channelNonce,
          outcome,
          appDefinition: countingApp.address,
          appData: appData[i - 1],
          challengeDuration,
          turnNum: largestTurnNum + i - numStates,
        });
      }

      const variableParts = states.map((state) => getVariablePart(state));

      // Sign the states
      const signatures = await signStates(states, wallets, whoSignedWhat);
      const { candidate } = separateProofAndCandidate(
        bindSignatures(variableParts, signatures, whoSignedWhat),
      );

      // Form transaction
      const pendingTx = testNitroAdjudicator.concludeAndTransferAllAssets(
        getFixedPart(states[0]),
        candidate,
        { gasLimit: NITRO_MAX_GAS },
      );

      // Switch on overall test expectation
      if (reasonString) {
        await expect(pendingTx).to.be.revertedWith(reasonString);
      } else {
        const tx = await pendingTx;
        const receipt = await tx.wait();

        expect(BigNumber.from(receipt.gasUsed).lt(BigNumber.from(NITRO_MAX_GAS))).to.equal(true);

        // Compute expected ChannelDataHash
        const block = await ethers.provider.getBlock(receipt.blockNumber);
        const expectedFingerprint =
          // eslint-disable-next-line @typescript-eslint/no-unnecessary-condition
          newOutcome.length === undefined
            ? constants.HashZero
            : channelDataToStatus({
                turnNumRecord: 0,
                finalizesAt: block.timestamp,
                outcome: computeOutcome(newOutcome),
              });

        // Check fingerprint against the expected value
        expect(await testNitroAdjudicator.statusOf(channelId)).to.equal(expectedFingerprint);

        // Check new holdings
        await Promise.all(
          // For each asset
          Object.keys(heldAfter).map(async (asset) => {
            await Promise.all(
              Object.keys(heldAfter[asset]).map(async (destination) => {
                // for each channel
                const amount = heldAfter[asset][destination];
                const res = await testNitroAdjudicator.holdings(asset, destination);
                expect(res.eq(amount)).to.eq(true);
              }),
            );
          }),
        );
      }
    });
  }
});
