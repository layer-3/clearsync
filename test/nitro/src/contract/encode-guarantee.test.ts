import type { BytesLike } from 'ethers';
import { describe, it } from 'mocha';

import {
  encodeGuaranteeData,
  decodeGuaranteeData,
  Guarantee,
} from '../../../../src/nitro/contract/outcome';
import { expect } from 'chai';

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

    testCases.forEach((tc) =>
      it(tc.description, () => {
        const { encodeFunction, decodeFunction, data } = tc as {
          encodeFunction: (g: Guarantee) => BytesLike;
          decodeFunction: (b: BytesLike) => Guarantee;
          data: Guarantee;
        };
        const encodedData = encodeFunction(data);
        const decodedData = decodeFunction(encodedData);
        expect(decodedData).to.deep.equal(data);
      }),
    );
  });
});
