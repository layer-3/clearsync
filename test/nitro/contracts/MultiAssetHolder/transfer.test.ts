import { BigNumber, constants, Contract } from 'ethers';
import { ethers } from 'hardhat';
import { describe, before, it } from 'mocha';
import { Allocation, AllocationType } from '@statechannels/exit-format';

import { expectRevert } from '../../../helpers/expect-revert';
import { randomChannelId, randomExternalDestination, setupContract } from '../../test-helpers';
import { encodeOutcome, hashOutcome, Outcome } from '../../../../src/nitro/contract/outcome';
import type { TESTNitroAdjudicator } from '../../../../typechain-types';
import TESTNitroAdjudicatorArtifact from '../../../../artifacts/contracts/nitro/test/TESTNitroAdjudicator.sol/TESTNitroAdjudicator.json';
import { channelDataToStatus, isExternalDestination } from '../../../../src/nitro';
import { MAGIC_ADDRESS_INDICATING_ETH } from '../../../../src/nitro/transactions';
import {
  AssetOutcomeShortHand,
  replaceAddressesAndBigNumberify,
} from '../../../../src/nitro/helpers';
import { expect } from 'chai';

interface addressT {
  [index: string]: string | undefined;
}

let testNitroAdjudicator: TESTNitroAdjudicator;
let addresses: addressT;

before(async () => {
  testNitroAdjudicator = await setupContract<TESTNitroAdjudicator>('TESTNitroAdjudicator');
  addresses = {
    // Channels
    c: undefined as string | undefined,
    C: randomChannelId(),
    X: randomChannelId(),
    // Externals
    A: randomExternalDestination(),
    B: randomExternalDestination(),
  };
});

// c is the channel we are transferring from.
describe('transfer', () => {
  const testCases = [
    {
      name: ' 0. channel not finalized        ',
      heldBefore: { c: 1 },
      isSimple: true,
      setOutcome: {},
      indices: [0],
      newOutcome: {},
      heldAfter: {},
      payouts: { A: 1 },
      reason: 'Channel not finalized',
    },
    {
      name: ' 1. funded          -> 1 EOA',
      heldBefore: { c: 1 },
      isSimple: true,
      setOutcome: { A: 1 },
      indices: [0],
      newOutcome: { A: 0 },
      heldAfter: {},
      payouts: { A: 1 },
      reason: undefined,
    },
    {
      name: ' 2. overfunded      -> 1 EOA',
      heldBefore: { c: 2 },
      isSimple: true,
      setOutcome: { A: 1 },
      indices: [0],
      newOutcome: { A: 0 },
      heldAfter: { c: 1 },
      payouts: { A: 1 },
      reason: undefined,
    },
    {
      name: ' 3. underfunded     -> 1 EOA',
      heldBefore: { c: 1 },
      isSimple: true,
      setOutcome: { A: 2 },
      indices: [0],
      newOutcome: { A: 1 },
      heldAfter: {},
      payouts: { A: 1 },
      reason: undefined,
    },
    {
      name: ' 4. funded      -> 1 channel',
      heldBefore: { c: 1 },
      isSimple: true,
      setOutcome: { C: 1 },
      indices: [0],
      newOutcome: { C: 0 },
      heldAfter: { c: 0, C: 1 },
      payouts: {},
      reason: undefined,
    },
    {
      name: ' 5. overfunded  -> 1 channel',
      heldBefore: { c: 2 },
      isSimple: true,
      setOutcome: { C: 1 },
      indices: [0],
      newOutcome: { C: 0 },
      heldAfter: { c: 1, C: 1 },
      payouts: {},
      reason: undefined,
    },
    {
      name: ' 6. underfunded -> 1 channel',
      heldBefore: { c: 1 },
      isSimple: true,
      setOutcome: { C: 2 },
      indices: [0],
      newOutcome: { C: 1 },
      heldAfter: { c: 0, C: 1 },
      payouts: {},
      reason: undefined,
    },
    {
      name: ' 7. -> 2 EOA         1 index',
      heldBefore: { c: 2 },
      isSimple: true,
      setOutcome: { A: 1, B: 1 },
      indices: [0],
      newOutcome: { A: 0, B: 1 },
      heldAfter: { c: 1 },
      payouts: { A: 1 },
      reason: undefined,
    },
    {
      name: ' 8. -> 2 EOA         1 index',
      heldBefore: { c: 1 },
      isSimple: true,
      setOutcome: { A: 1, B: 1 },
      indices: [0],
      newOutcome: { A: 0, B: 1 },
      heldAfter: { c: 0 },
      payouts: { A: 1 },
      reason: undefined,
    },
    {
      name: ' 9. -> 2 EOA         partial',
      heldBefore: { c: 3 },
      isSimple: true,
      setOutcome: { A: 2, B: 2 },
      indices: [1],
      newOutcome: { A: 2, B: 1 },
      heldAfter: { c: 2 },
      payouts: { B: 1 },
      reason: undefined,
    },
    {
      name: '10. -> 2 chan             no',
      heldBefore: { c: 1 },
      isSimple: true,
      setOutcome: { C: 1, X: 1 },
      indices: [1],
      newOutcome: { C: 1, X: 1 },
      heldAfter: { c: 1 },
      payouts: {},
      reason: undefined,
    },
    {
      name: '11. -> 2 chan           full',
      heldBefore: { c: 1 },
      isSimple: true,
      setOutcome: { C: 1, X: 1 },
      indices: [0],
      newOutcome: { C: 0, X: 1 },
      heldAfter: { c: 0, C: 1 },
      payouts: {},
      reason: undefined,
    },
    {
      name: '12. -> 2 chan        partial',
      heldBefore: { c: 3 },
      isSimple: true,
      setOutcome: { C: 2, X: 2 },
      indices: [1],
      newOutcome: { C: 2, X: 1 },
      heldAfter: { c: 2, X: 1 },
      payouts: {},
      reason: undefined,
    },
    {
      name: '13. -> 2 indices',
      heldBefore: { c: 3 },
      isSimple: true,
      setOutcome: { C: 2, X: 2 },
      indices: [0, 1],
      newOutcome: { C: 0, X: 1 },
      heldAfter: { c: 0, X: 1 },
      payouts: { C: 2 },
      reason: undefined,
    },
    {
      name: '14. -> 3 indices',
      heldBefore: { c: 5 },
      isSimple: true,
      setOutcome: { A: 1, C: 2, X: 2 },
      indices: [0, 1, 2],
      newOutcome: { A: 0, C: 0, X: 0 },
      heldAfter: { c: 0, X: 2 },
      payouts: { A: 1, C: 2 },
      reason: undefined,
    },
    {
      name: '15. -> reverse order (see 13)',
      heldBefore: { c: 3 },
      isSimple: true,
      setOutcome: { C: 2, X: 2 },
      indices: [1, 0],
      newOutcome: { C: 2, X: 1 },
      heldAfter: { c: 2, X: 1 },
      payouts: {},
      reason: 'Indices must be sorted',
    },
    {
      name: '16. incorrect fingerprint        ',
      heldBefore: { c: 1 },
      isSimple: true,
      setOutcome: {},
      indices: [0],
      newOutcome: {},
      heldAfter: {},
      payouts: { A: 1 },
      reason: 'incorrect fingerprint',
    },
    {
      name: '17. guarantee allocationType',
      heldBefore: { c: 1 },
      isSimple: false,
      setOutcome: { A: 1 },
      indices: [0],
      newOutcome: { A: 0 },
      heldAfter: {},
      payouts: { A: 1 },
      reason: 'cannot transfer a guarantee',
    },
  ];

  testCases.forEach((tc) =>
    it(tc.name, async () => {
      let heldBefore = tc.heldBefore as AssetOutcomeShortHand;
      let setOutcome = tc.setOutcome as AssetOutcomeShortHand;
      let newOutcome = tc.newOutcome as AssetOutcomeShortHand;
      let heldAfter = tc.heldAfter as AssetOutcomeShortHand;
      let payouts = tc.payouts as AssetOutcomeShortHand;
      const reason = tc.reason as string;

      // Compute channelId
      addresses.c = randomChannelId();
      const channelId = addresses.c;
      addresses.C = randomChannelId();
      addresses.X = randomChannelId();
      addresses.A = randomExternalDestination();
      addresses.B = randomExternalDestination();

      // Transform input data (unpack addresses and BigNumberify amounts)
      heldBefore = replaceAddressesAndBigNumberify(heldBefore, addresses) as AssetOutcomeShortHand;
      setOutcome = replaceAddressesAndBigNumberify(setOutcome, addresses) as AssetOutcomeShortHand;
      newOutcome = replaceAddressesAndBigNumberify(newOutcome, addresses) as AssetOutcomeShortHand;
      heldAfter = replaceAddressesAndBigNumberify(heldAfter, addresses) as AssetOutcomeShortHand;
      payouts = replaceAddressesAndBigNumberify(payouts, addresses) as AssetOutcomeShortHand;

      // Deposit into channels

      await Promise.all(
        Object.keys(heldBefore).map(async (key) => {
          // Key must be either in heldBefore or heldAfter or both
          const amount = heldBefore[key];
          await (
            await testNitroAdjudicator.deposit(MAGIC_ADDRESS_INDICATING_ETH, key, 0, amount, {
              value: amount,
            })
          ).wait();
          expect(
            (await testNitroAdjudicator.holdings(MAGIC_ADDRESS_INDICATING_ETH, key)).eq(amount),
          ).to.equal(true);
        }),
      );

      // Compute an appropriate allocation.
      const allocations: Allocation[] = [];
      Object.keys(setOutcome).forEach((key) =>
        allocations.push({
          destination: key,
          amount: BigNumber.from(setOutcome[key]).toHexString(),
          metadata: '0x',
          allocationType: (tc.isSimple as boolean)
            ? AllocationType.simple
            : AllocationType.guarantee,
        }),
      );
      const outcomeHash = hashOutcome([
        {
          asset: MAGIC_ADDRESS_INDICATING_ETH,
          assetMetadata: { assetType: 0, metadata: '0x' },
          allocations,
        },
      ]);
      const outcomeBytes = encodeOutcome([
        {
          asset: MAGIC_ADDRESS_INDICATING_ETH,
          assetMetadata: { assetType: 0, metadata: '0x' },
          allocations,
        },
      ]);

      // Set adjudicator status
      const stateHash = constants.HashZero; // not realistic, but OK for purpose of this test
      const finalizesAt = 42;
      const turnNumRecord = 7;

      if (reason != 'Channel not finalized') {
        await (
          await testNitroAdjudicator.setStatusFromChannelData(channelId, {
            turnNumRecord,
            finalizesAt,
            stateHash,
            outcomeHash,
          })
        ).wait();
      }

      const tx = testNitroAdjudicator.transfer(
        MAGIC_ADDRESS_INDICATING_ETH,
        channelId,
        reason == 'incorrect fingerprint' ? '0xdeadbeef' : outcomeBytes,
        stateHash,
        tc.indices as number[],
      );

      // Call method in a slightly different way if expecting a revert
      if (reason) {
        await expectRevert(() => tx, reason);
      } else {
        const { events: eventsFromTx } = await (await tx).wait();
        // Check new holdings
        await Promise.all(
          Object.keys(heldAfter).map(async (key) =>
            expect(await testNitroAdjudicator.holdings(MAGIC_ADDRESS_INDICATING_ETH, key)).to.equal(
              heldAfter[key],
            ),
          ),
        );

        // Check new status
        const allocationsAfter: Allocation[] = [];
        Object.keys(newOutcome).forEach((key) => {
          allocationsAfter.push({
            destination: key,
            amount: BigNumber.from(newOutcome[key]).toHexString(),
            metadata: '0x',
            allocationType: AllocationType.simple,
          });
        });
        const outcomeAfter: Outcome = [
          {
            asset: MAGIC_ADDRESS_INDICATING_ETH,
            assetMetadata: { assetType: 0, metadata: '0x' },
            allocations: allocationsAfter,
          },
        ];
        const expectedStatusAfter = channelDataToStatus({
          turnNumRecord,
          finalizesAt,
          // stateHash will be set to HashZero by this helper fn
          // if state property of this object is undefined
          outcome: outcomeAfter,
        });
        expect(await testNitroAdjudicator.statusOf(channelId)).to.equal(expectedStatusAfter);

        const expectedEvents = [
          {
            event: 'AllocationUpdated',
            args: {
              channelId,
              assetIndex: BigNumber.from(0),
              initialHoldings: heldBefore[addresses.c],
            },
          },
        ];

        expectedEvents.forEach((expectedEvent, index) => {
          const actualEvent = eventsFromTx[index];

          // Assert the 'event' field
          expect(actualEvent.event).to.equal(expectedEvent.event);

          // Assert each field in 'args'
          Object.entries(expectedEvent.args).forEach(([key, value]) => {
            expect(actualEvent.args[key]).to.deep.equal(value);
          });
        });

        // Check payouts
        for (const destination of Object.keys(payouts)) {
          if (isExternalDestination(destination)) {
            const asAddress = '0x' + destination.substring(26);
            const balance = await ethers.provider.getBalance(asAddress);
            // console.log(`checking balance of ${destination}: ${balance.toString()}`);
            expect(balance).to.equal(payouts[destination]);
          } else {
            const holdings = await testNitroAdjudicator.holdings(
              MAGIC_ADDRESS_INDICATING_ETH,
              destination,
            );
            // console.log(`checking holdings of ${destination}: ${holdings.toString()}`);
            expect(holdings).to.equal(payouts[destination]);
          }
        }
      }
    }),
  );
});
