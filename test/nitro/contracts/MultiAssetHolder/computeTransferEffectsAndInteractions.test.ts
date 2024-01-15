import { BigNumber } from 'ethers';
import shuffle from 'lodash.shuffle';
import { Allocation, AllocationType } from '@statechannels/exit-format';
import { expect } from 'chai';

import { convertToStruct, randomExternalDestination, setupContract } from '../../test-helpers';
import { computeTransferEffectsAndInteractions } from '../../../../src/nitro/contract/multi-asset-holder';

import type { TESTNitroAdjudicator } from '../../../../typechain-types';

let testNitroAdjudicator: TESTNitroAdjudicator;

const randomAllocations = (numAllocations: number): Allocation[] => {
  return numAllocations > 0
    ? Array.from({ length: numAllocations }, () => ({
        destination: randomExternalDestination(),
        amount: BigNumber.from(Math.ceil(Math.random() * 10)).toHexString(),
        metadata: '0x',
        allocationType: AllocationType.simple,
      }))
    : [];
};

const heldBefore = 100;
const heldBeforeHex = BigNumber.from(heldBefore).toHexString();
const allocation = randomAllocations(Math.floor(Math.random() * 20));
const indices = shuffle([...Array.from({ length: allocation.length }).keys()]); // [0, 1, 2, 3,...] but shuffled
// TODO -- does it make sense to test with indices that don't increase when the chain requires that they do?

before(async () => {
  testNitroAdjudicator = await setupContract<TESTNitroAdjudicator>('TESTNitroAdjudicator');
});

describe('MultiAssetHolder.compute_transfer_effects_and_interactions', () => {
  it(`matches on chain method for input \n heldBefore: ${heldBefore}, \n allocationLen: ${
    allocation.length
  }, \n indices: ${indices.reduce((acc, i) => acc + i.toString() + ',', '')}`, async () => {
    // check local function works as expected
    const locallyComputedNewAllocation = computeTransferEffectsAndInteractions(
      heldBeforeHex,
      allocation,
      indices,
    );

    const onChainResult = await testNitroAdjudicator.compute_transfer_effects_and_interactions(
      heldBefore,
      allocation,
      indices,
    );

    const convertedOnChainResult = {
      newAllocations: onChainResult.newAllocations.map((a) => convertToStruct(a)),
      allocatesOnlyZeros: onChainResult.allocatesOnlyZeros,
      exitAllocations: onChainResult.exitAllocations.map((a) => convertToStruct(a)),
      totalPayouts: onChainResult.totalPayouts,
    };

    expect(convertedOnChainResult).to.exist;

    expect(convertedOnChainResult.newAllocations).to.deep.equal(
      locallyComputedNewAllocation.newAllocations.map((a) => ({
        ...a,
        amount: BigNumber.from(a.amount),
      })),
    );

    expect(convertedOnChainResult.allocatesOnlyZeros).to.equal(
      locallyComputedNewAllocation.allocatesOnlyZeros,
    );

    expect(convertedOnChainResult.exitAllocations).to.deep.equal(
      locallyComputedNewAllocation.exitAllocations.map((a) => ({
        ...a,
        amount: BigNumber.from(a.amount),
      })),
    );

    expect(convertedOnChainResult.totalPayouts.toString()).to.equal(
      BigNumber.from(locallyComputedNewAllocation.totalPayouts).toString(),
    );
  });
});
