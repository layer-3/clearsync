import { Contract, constants } from 'ethers';
import { ethers } from 'hardhat';
import { expect } from 'chai';
import { before, describe, it } from 'mocha';

import { expectRevert } from '../../../helpers/expect-revert';
import { getChannelId } from '../../../../src/nitro/contract/channel';
import { hashOutcome, Outcome } from '../../../../src/nitro/contract/outcome';
import {
  generateParticipants,
  randomChannelId,
  randomExternalDestination,
  setupContract,
} from '../../test-helpers';
import type { CountingApp, TESTNitroAdjudicator, Token } from '../../../../typechain-types';
import {
  channelDataToStatus,
  computeOutcome,
  convertBytes32ToAddress,
  getRandomNonce,
  OutcomeShortHand,
} from '../../../../src/nitro';
import { MAGIC_ADDRESS_INDICATING_ETH } from '../../../../src/nitro/transactions';
import { replaceAddressesAndBigNumberify } from '../../../../src/nitro/helpers';

// Constants for this test suite
const nParticipants = 3;
const { participants } = generateParticipants(nParticipants);

const challengeDuration = 0x1000;
let countingApp: Contract;
let testNitroAdjudicator: Contract;
let token: Contract;
let addresses: any;

type testParams = {
  setOutcome: OutcomeShortHand;
  heldBefore: OutcomeShortHand;
  newOutcome: OutcomeShortHand;
  heldAfter: OutcomeShortHand;
  payouts: OutcomeShortHand;
  reasonString: string | undefined;
};

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

describe('transferAllAssets', async () => {
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

  testCases.forEach((tc) =>
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
                await (await token.increaseAllowance(testNitroAdjudicator.address, amount)).wait(); // Approve enough for setup and main test
              }
              await (
                await testNitroAdjudicator.deposit(asset, destination, 0, amount, {
                  value: asset == MAGIC_ADDRESS_INDICATING_ETH ? amount : 0,
                })
              ).wait();
              expect((await testNitroAdjudicator.holdings(asset, destination)).eq(amount)).to.equal(
                true,
              );
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
      await (
        await testNitroAdjudicator.setStatusFromChannelData(channelId, {
          turnNumRecord,
          finalizesAt,
          stateHash,
          outcomeHash,
        })
      ).wait();

      const tx1 = testNitroAdjudicator.transferAllAssets(channelId, outcome, stateHash);

      // Call method in a slightly different way if expecting a revert
      if (reasonString) {
        const regex = new RegExp(
          '^' + 'VM Exception while processing transaction: revert ' + reasonString + '$',
        );
        await expectRevert(() => tx1, regex);
      } else {
        const { events: eventsFromTx } = await (await tx1).wait();

        expect(eventsFromTx).not.to.equal(undefined);
        if (eventsFromTx === undefined) {
          return;
        }

        // expect an event per asset
        expect(eventsFromTx[0].event).to.equal('AllocationUpdated');
        expect(eventsFromTx[1].event).to.equal('AllocationUpdated');

        // Check new status
        const outcomeAfter: Outcome = computeOutcome(newOutcome);

        const expectedStatusAfter = newOutcome.length
          ? channelDataToStatus({
              turnNumRecord,
              finalizesAt,
              // stateHash will be set to HashZero by this helper fn
              // if state property of this object is undefined
              outcome: outcomeAfter,
            })
          : constants.HashZero;
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
                if (asset != MAGIC_ADDRESS_INDICATING_ETH) {
                  expect((await token.balanceOf(address)).eq(amount)).to.equal(true);
                } else {
                  expect((await ethers.provider.getBalance(address)).eq(amount)).to.equal(true);
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
                expect(
                  (await testNitroAdjudicator.holdings(asset, destination)).eq(amount),
                ).to.equal(true);
              }),
            );
          }),
        );
      }
    }),
  );
});
