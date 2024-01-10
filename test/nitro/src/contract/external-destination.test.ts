import { describe, it } from 'mocha';
import { expect } from 'chai';

import { isExternalDestination } from '../../../../src/nitro/contract/channel';

describe('isExternalDestination', () => {
  const testCases = [
    {
      bytes32: '0x0',
      result: false,
    },
    {
      bytes32: '0x0000000000000000000000002F0E2cB3c2c98E6AfB89A8c50cbEF0cB6B3DC35c',
      result: true,
    },
    {
      bytes32: '0x0000000040000000000000002F0E2cB3c2c98E6AfB89A8c50cbEF0cB6B3DC35c',
      result: false,
    },
  ];

  for (const tc of testCases) it(`${tc.bytes32} -- ${tc.result}`, () => {
      const { bytes32, result } = tc as { bytes32: string; result: boolean };
      expect(isExternalDestination(bytes32)).to.equal(result);
    })
  ;
});
