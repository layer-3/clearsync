import { describe, it } from 'mocha';

import { convertBytes32ToAddress } from '../../../../src/nitro/contract/multi-asset-holder';
import { expect } from 'chai';

describe('convertBytes32ToAddress', () => {
  const testCases = [
    {
      bytes32: '0x0000000000000000000000000000000000000000000000000000000000000000',
      address: '0x0000000000000000000000000000000000000000',
    },
    {
      bytes32: '0x000000000000000000000000aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
      address: '0xaAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa',
    },
    {
      bytes32: '0x000000000000000000000000aAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa',
      address: '0xaAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa',
    },
    {
      bytes32: '0x000000000000000000000000000000000000000000000000000000000000000a',
      address: '0x000000000000000000000000000000000000000A',
    },
    {
      bytes32: '0x000000000000000000000000000000000000000000000000000000000000000A',
      address: '0x000000000000000000000000000000000000000A',
    },
  ];

  testCases.forEach((tc) =>
    it(`${tc.bytes32} -- ${tc.address}`, () => {
      expect(convertBytes32ToAddress(tc.bytes32)).to.equal(tc.address);
    }),
  );
});
