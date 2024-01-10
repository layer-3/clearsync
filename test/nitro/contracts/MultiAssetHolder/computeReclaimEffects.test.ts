import { Contract, constants } from 'ethers';
import { Allocation, AllocationType } from '@statechannels/exit-format';
import { describe, it } from 'mocha';
import { expect } from 'chai';

import { convertToStruct, setupContract } from '../../test-helpers';
import { computeReclaimEffects } from '../../../../src/nitro/contract/multi-asset-holder';
import { encodeGuaranteeData } from '../../../../src/nitro/contract/outcome';

import type { TESTNitroAdjudicator } from '../../../../typechain-types';

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
  for (const testCase of testCases) {
    it('off chain method matches expectation', () => {
      const offChainNewSourceAllocations = computeReclaimEffects(
        testCase.inputs.sourceAllocations,
        testCase.inputs.targetAllocations,
        testCase.inputs.indexOfTargetInSource,
      );

      expect(offChainNewSourceAllocations).to.deep.equal(testCase.outputs.newSourceAllocations);
    });

    it('on chain method matches expectation', async () => {
      const onChainNewSourceAllocations = await testNitroAdjudicator.compute_reclaim_effects(
        testCase.inputs.sourceAllocations,
        testCase.inputs.targetAllocations,
        testCase.inputs.indexOfTargetInSource,
      );

      const convertedOnChainAllocs = onChainNewSourceAllocations.map((alloc) =>
        convertToStruct(alloc),
      );
      expect(convertedOnChainAllocs).to.deep.equal(testCase.outputs.newSourceAllocations);
    });
  }
});
