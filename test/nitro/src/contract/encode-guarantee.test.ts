import { describe, it } from 'mocha';
import { expect } from 'chai';

import {
  Guarantee,
  decodeGuaranteeData,
  encodeGuaranteeData,
} from '../../../../src/nitro/contract/outcome';

import type { BytesLike } from 'ethers';

describe('outcome', () => {
  describe('encoding and decoding', () => {
    const testCases = [
      {
        description: 'Encodes and decodes guarantee',
        encodeFunction: encodeGuaranteeData,
        decodeFunction: decodeGuaranteeData,
        data: {
          left: '0x14bcc435f49d130d189737f9762feb25c44ef5b886bef833e31a702af6be4748',
          right: '0x14bcc435f49d130d189736bef833e31a702af6be47487f9762feb25c44ef5b88',
        },
      },
    ];

    for (const tc of testCases) it(tc.description, () => {
        const { encodeFunction, decodeFunction, data } = tc as {
          encodeFunction: (g: Guarantee) => BytesLike;
          decodeFunction: (b: BytesLike) => Guarantee;
          data: Guarantee;
        };
        const encodedData = encodeFunction(data);
        const decodedData = decodeFunction(encodedData);
        expect(decodedData).to.deep.equal(data);
      })
    ;
  });
});
