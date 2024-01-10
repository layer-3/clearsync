import { Allocation, AllocationType } from '@statechannels/exit-format';
import { BigNumber, constants } from 'ethers';
import { before, describe, it } from 'mocha';
import { expect } from 'chai';

import { randomChannelId, randomExternalDestination, setupContract } from '../../test-helpers';
import { Outcome, channelDataToStatus, encodeOutcome, hashOutcome } from '../../../../src/nitro';
import { MAGIC_ADDRESS_INDICATING_ETH } from '../../../../src/nitro/transactions';
import { encodeGuaranteeData } from '../../../../src/nitro/contract/outcome';

import type { TESTNitroAdjudicator } from '../../../../typechain-types';

let testNitroAdjudicator: TESTNitroAdjudicator;

before(async () => {
  testNitroAdjudicator = await setupContract<TESTNitroAdjudicator>('TESTNitroAdjudicator');
});

// Amounts are valueString representations of wei
describe('reclaim', () => {
  // TODO: add a test case to show off a multihop reclaim, where we have Alice, Irene, Ivan and Bob.
  it('handles a simple case as expected', async () => {
    const targetId = randomChannelId();
    const sourceId = randomChannelId();
    const Alice = randomExternalDestination();
    const Bob = randomExternalDestination();
    const Irene = randomExternalDestination();

    // prepare an appropriate virtual channel outcome and finalize

    const vAllocations: Allocation[] = [
      {
        destination: Alice,
        amount: BigNumber.from(7).toHexString(),
        allocationType: AllocationType.simple,
        metadata: '0x',
      },
      {
        destination: Bob,
        amount: BigNumber.from(3).toHexString(),
        allocationType: AllocationType.simple,
        metadata: '0x',
      },
    ];

    const vOutcome: Outcome = [
      {
        asset: MAGIC_ADDRESS_INDICATING_ETH,
        allocations: vAllocations,
        assetMetadata: { assetType: 0, metadata: '0x' },
      },
    ];
    const vOutcomeHash = hashOutcome(vOutcome);
    await (
      await testNitroAdjudicator.setStatusFromChannelData(targetId, {
        turnNumRecord: 99,
        finalizesAt: 1,
        stateHash: constants.HashZero, // not realistic, but OK for purpose of this test
        outcomeHash: vOutcomeHash,
      })
    ).wait();

    // prepare an appropriate ledger channel outcome and finalize

    const lAllocations: Allocation[] = [
      {
        destination: Alice,
        amount: BigNumber.from(10).toHexString(),
        allocationType: AllocationType.simple,
        metadata: '0x',
      },
      {
        destination: Irene,
        amount: BigNumber.from(10).toHexString(),
        allocationType: AllocationType.simple,
        metadata: '0x',
      },
      {
        destination: targetId,
        amount: BigNumber.from(10).toHexString(),
        allocationType: AllocationType.guarantee,
        metadata: encodeGuaranteeData({ left: Alice, right: Irene }),
      },
    ];

    const lOutcome: Outcome = [
      {
        asset: MAGIC_ADDRESS_INDICATING_ETH,
        allocations: lAllocations,
        assetMetadata: { assetType: 0, metadata: '0x' },
      },
    ];
    const lOutcomeHash = hashOutcome(lOutcome);
    await (
      await testNitroAdjudicator.setStatusFromChannelData(sourceId, {
        turnNumRecord: 99,
        finalizesAt: 1,
        stateHash: constants.HashZero, // not realistic, but OK for purpose of this test
        outcomeHash: lOutcomeHash,
      })
    ).wait();

    // call reclaim

    const tx = testNitroAdjudicator.reclaim({
      sourceChannelId: sourceId,
      sourceStateHash: constants.HashZero,
      sourceOutcomeBytes: encodeOutcome(lOutcome),
      sourceAssetIndex: 0, // TODO: introduce test cases with multiple-asset Source and Targets
      indexOfTargetInSource: 2,
      targetStateHash: constants.HashZero,
      targetOutcomeBytes: encodeOutcome(vOutcome),
      targetAssetIndex: 0,
    });

    // Extract logs
    const { events: eventsFromTx } = await (await tx).wait();

    // Compile event expectations

    // Check that each expectedEvent is contained as a subset of the properies of each *corresponding* event: i.e. the order matters!
    const expectedEvents = [
      {
        event: 'Reclaimed',
        args: {
          channelId: sourceId,
          assetIndex: BigNumber.from(0),
        },
      },
    ];

    for (const [index, expectedEvent] of expectedEvents.entries()) {
      const actualEvent = eventsFromTx[index];

      // Assert the 'event' field
      expect(actualEvent.event).to.equal(expectedEvent.event);

      // Assert each field in 'args'
      for (const [key, value] of Object.entries(expectedEvent.args)) {
        expect(actualEvent.args[key]).to.deep.equal(value);
      }
    }

    // assert on updated ledger channel

    // Check new outcomeHash
    const allocationAfter: Allocation[] = [
      {
        destination: Alice,
        amount: BigNumber.from(17).toHexString(),
        allocationType: AllocationType.simple,
        metadata: '0x',
      },
      {
        destination: Irene,
        amount: BigNumber.from(13).toHexString(),
        allocationType: AllocationType.simple,
        metadata: '0x',
      },
    ];

    const outcomeAfter: Outcome = [
      {
        asset: MAGIC_ADDRESS_INDICATING_ETH,
        allocations: allocationAfter,
        assetMetadata: { assetType: 0, metadata: '0x' },
      },
    ];
    const expectedStatusAfter = channelDataToStatus({
      turnNumRecord: 99,
      finalizesAt: 1,
      // stateHash will be set to HashZero by this helper fn
      // if state property of this object is undefined
      outcome: outcomeAfter,
    });

    expect(await testNitroAdjudicator.statusOf(sourceId)).to.equal(expectedStatusAfter);

    // assert that virtual channel did not change.

    expect(await testNitroAdjudicator.statusOf(targetId)).to.equal(
      channelDataToStatus({
        turnNumRecord: 99,
        finalizesAt: 1,
        outcome: vOutcome,
      }),
    );
  });
});
