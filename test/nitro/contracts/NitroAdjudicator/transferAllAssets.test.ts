import { constants } from 'ethers';
import { ethers } from 'hardhat';
import { assert, expect } from 'chai';
import { before, describe, it } from 'mocha';

import { getChannelId } from '../../../../src/nitro/contract/channel';
import { Outcome, hashOutcome } from '../../../../src/nitro/contract/outcome';
import {
  generateParticipants,
  randomChannelId,
  randomExternalDestination,
  setupContract,
} from '../../test-helpers';
import {
  OutcomeShortHand,
  channelDataToStatus,
  computeOutcome,
  convertBytes32ToAddress,
  getRandomNonce,
} from '../../../../src/nitro';
import { MAGIC_ADDRESS_INDICATING_ETH } from '../../../../src/nitro/transactions';
import { replaceAddressesAndBigNumberify } from '../../../../src/nitro/helpers';

import type { CountingApp, TESTNitroAdjudicator, Token } from '../../../../typechain-types';

// Constants for this test suite
const nParticipants = 3;
const { participants } = generateParticipants(nParticipants);

const challengeDuration = 0x10_00;
let countingApp: CountingApp;
let testNitroAdjudicator: TESTNitroAdjudicator;
let token: Token;
let addresses: Record<string, string | undefined>;

interface testParams {
  setOutcome: OutcomeShortHand;
  heldBefore: OutcomeShortHand;
  newOutcome: OutcomeShortHand;
  heldAfter: OutcomeShortHand;
  payouts: OutcomeShortHand;
  reasonString: string | undefined;
}

before(async () => {
  countingApp = await setupContract<CountingApp>('CountingApp');
  testNitroAdjudicator = await setupContract<TESTNitroAdjudicator>('TESTNitroAdjudicator');
  token = await setupContract<Token>('Token', '0x6B8B2958795a5E9c00A2E8D4B0b90b870cbAB4b6');
  addresses = {
    // Channels
    c: undefined as string | undefined,
    C: randomChannelId(),
    X: randomChannelId(),
    // Externals
    A: randomExternalDestination(),
    B: randomExternalDestination(),
    ETH: MAGIC_ADDRESS_INDICATING_ETH,
    ERC20: token.address,
  };
});

// eslint-disable-next-line @typescript-eslint/no-misused-promises, sonarjs/cognitive-complexity
describe('transferAllAssets', () => {
  const testCases = [
    {
      description:
        'testNitroAdjudicator accepts a transferAllAssets tx for a finalized channel, and 2x Asset types transferred',
      setOutcome: { ETH: { A: 1 }, ERC20: { A: 2 } },
      heldBefore: { ETH: { c: 1 }, ERC20: { c: 2 } },
      newOutcome: {},
      heldAfter: { ETH: { c: 0 }, ERC20: { c: 0 } },
      payouts: { ETH: { A: 1 }, ERC20: { A: 2 } },
      reasonString: undefined,
    },
  ];

  for (const tc of testCases)
    it(tc.description, async () => {
      const channelNonce = getRandomNonce('transferAllAssets');
      const channelId = getChannelId({
        channelNonce,
        participants,
        appDefinition: countingApp.address,
        challengeDuration,
      });

      let { setOutcome, heldBefore, newOutcome, heldAfter, payouts } = tc as unknown as testParams;
      const { reasonString } = tc as unknown as testParams;

      addresses.c = channelId;

      // Transform input data (unpack addresses and BigNumberify amounts)
      [heldBefore, setOutcome, newOutcome, heldAfter, payouts] = [
        heldBefore,
        setOutcome,
        newOutcome,
        heldAfter,
        payouts,
      ].map((object) => replaceAddressesAndBigNumberify(object, addresses) as OutcomeShortHand);

      // Deposit into channels
      await Promise.all(
        // For each asset
        Object.keys(heldBefore).map(async (asset) => {
          await Promise.all(
            Object.keys(heldBefore[asset]).map(async (destination) => {
              // for each channel
              const amount = heldBefore[asset][destination];
              if (asset != MAGIC_ADDRESS_INDICATING_ETH) {
                // Increase allowance
                const increaseTx = await token.approve(testNitroAdjudicator.address, amount);
                await increaseTx.wait(); // Approve enough for setup and main test
              }
              const depositTx = await testNitroAdjudicator.deposit(asset, destination, 0, amount, {
                value: asset == MAGIC_ADDRESS_INDICATING_ETH ? amount : 0,
              });
              await depositTx.wait();

              const holdings = await testNitroAdjudicator.holdings(asset, destination);
              expect(holdings).to.equal(amount);
            }),
          );
        }),
      );

      // Compute the outcome.
      const outcome: Outcome = computeOutcome(setOutcome);
      const outcomeHash = hashOutcome(outcome);
      // Call public wrapper to set state (only works on test contract)
      const stateHash = constants.HashZero;
      const finalizesAt = 42;
      const turnNumRecord = 7;
      const setStatusTx = await testNitroAdjudicator.setStatusFromChannelData(channelId, {
        turnNumRecord,
        finalizesAt,
        stateHash,
        outcomeHash,
      });
      await setStatusTx.wait();

      const pendingTx = testNitroAdjudicator.transferAllAssets(channelId, outcome, stateHash);

      // Call method in a slightly different way if expecting a revert
      if (reasonString) {
        const regex = new RegExp(
          '^' + 'VM Exception while processing transaction: revert ' + reasonString + '$',
        );
        await expect(pendingTx).to.be.revertedWith(regex);
      } else {
        const tx = await pendingTx;
        const { events: eventsFromTx } = await tx.wait();

        expect(eventsFromTx).not.to.equal(undefined);
        if (eventsFromTx === undefined) {
          assert.fail('eventsFromTx is undefined');
        }

        // expect an event per asset
        expect(eventsFromTx[0].event).to.equal('AllocationUpdated');
        expect(eventsFromTx[1].event).to.equal('AllocationUpdated');

        // Check new status
        const outcomeAfter: Outcome = computeOutcome(newOutcome);

        const expectedStatusAfter =
          // eslint-disable-next-line @typescript-eslint/no-unnecessary-condition
          newOutcome.length === undefined
            ? constants.HashZero
            : channelDataToStatus({
                turnNumRecord,
                finalizesAt,
                // stateHash will be set to HashZero by this helper fn
                // if state property of this object is undefined
                outcome: outcomeAfter,
              });
        expect(await testNitroAdjudicator.statusOf(channelId)).to.equal(expectedStatusAfter);

        // Check payouts
        await Promise.all(
          // For each asset
          Object.keys(payouts).map(async (asset) => {
            await Promise.all(
              Object.keys(payouts[asset]).map(async (destination) => {
                const address = convertBytes32ToAddress(destination);
                // for each channel
                const amount = payouts[asset][destination];
                if (asset == MAGIC_ADDRESS_INDICATING_ETH) {
                  const balance = await ethers.provider.getBalance(address);
                  expect(balance).to.equal(amount);
                } else {
                  const balance = await token.balanceOf(address);
                  expect(balance).equal(amount);
                }
              }),
            );
          }),
        );

        // Check new holdings
        await Promise.all(
          // For each asset
          Object.keys(heldAfter).map(async (asset) => {
            await Promise.all(
              Object.keys(heldAfter[asset]).map(async (destination) => {
                // for each channel
                const amount = heldAfter[asset][destination];
                const holdings = await testNitroAdjudicator.holdings(asset, destination);
                expect(holdings).to.equal(amount);
              }),
            );
          }),
        );
      }
    });
});
