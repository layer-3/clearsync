import { BigNumber, BytesLike, Contract, constants } from 'ethers';
import { Allocation, AllocationType } from '@statechannels/exit-format';
import { describe, it } from 'mocha';

import { setupContract } from '../../test-helpers';
import type { TESTNitroAdjudicator } from '../../../../typechain-types';
import { computeReclaimEffects } from '../../../../src/nitro/contract/multi-asset-holder';
import { encodeGuaranteeData } from '../../../../src/nitro/contract/outcome';
import { assert } from 'chai';

let testNitroAdjudicator: Contract & TESTNitroAdjudicator;

const Alice = '0x000000000000000000000000000000000000000000000000000000000000000a';
const Bob = '0x000000000000000000000000000000000000000000000000000000000000000b';

interface TestCaseInputs {
  sourceAllocations: Allocation[];
  targetAllocations: Allocation[];
  indexOfTargetInSource: number;
}

interface TestCaseOutputs {
  newSourceAllocations: Allocation[];
}
interface TestCase {
  inputs: TestCaseInputs;
  outputs: TestCaseOutputs;
}

interface AllocationT {
  destination: string;
  amount: BigNumber;
  allocationType: number;
  metadata: BytesLike;
}

const testCases: TestCase[] = [
  {
    inputs: {
      indexOfTargetInSource: 2,
      sourceAllocations: [
        {
          destination: Alice,
          amount: '0x02',
          allocationType: AllocationType.simple,
          metadata: '0x',
        },
        {
          destination: Bob,
          amount: '0x02',
          allocationType: AllocationType.simple,
          metadata: '0x',
        },
        {
          destination: constants.HashZero,
          amount: '0x06',
          allocationType: AllocationType.guarantee,
          metadata: encodeGuaranteeData({ left: Alice, right: Bob }),
        },
      ],
      targetAllocations: [
        {
          destination: Alice,
          amount: '0x01',
          allocationType: AllocationType.simple,
          metadata: '0x',
        },
        {
          destination: Bob,
          amount: '0x05',
          allocationType: AllocationType.simple,
          metadata: '0x',
        },
      ],
    },
    outputs: {
      newSourceAllocations: [
        {
          destination: Alice,
          amount: '0x03',
          allocationType: AllocationType.simple,
          metadata: '0x',
        },
        {
          destination: Bob,
          amount: '0x07',
          allocationType: AllocationType.simple,
          metadata: '0x',
        },
      ],
    },
  },
];

before(async () => {
  testNitroAdjudicator = await setupContract<TESTNitroAdjudicator>('TESTNitroAdjudicator');
});

describe('computeReclaimEffects', () => {
  testCases.forEach((testCase: TestCase) => {
    it('off chain method matches expectation', () => {
      const offChainNewSourceAllocations = computeReclaimEffects(
        testCase.inputs.sourceAllocations,
        testCase.inputs.targetAllocations,
        testCase.inputs.indexOfTargetInSource,
      );

      assert.deepEqual(offChainNewSourceAllocations, testCase.outputs.newSourceAllocations);
    });

    it('on chain method matches expectation', async () => {
      const onChainNewSourceAllocations = await testNitroAdjudicator.compute_reclaim_effects(
        testCase.inputs.sourceAllocations,
        testCase.inputs.targetAllocations,
        testCase.inputs.indexOfTargetInSource,
      );

      const res = onChainNewSourceAllocations.map(convertAmountToHexString);
      console.log(res, testCase.outputs.newSourceAllocations);

      assert.deepEqual(res, testCase.outputs.newSourceAllocations);
    });
  });

  const convertAmountToHexString = (a: AllocationT) => ({ ...a, amount: a.amount.toHexString() });
});
