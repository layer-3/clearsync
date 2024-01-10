import { expect } from 'chai';

import { setupContract } from '../../test-helpers';
import type { TESTNitroAdjudicator } from '../../../../typechain-types';
import { before } from 'mocha';

let testNitroAdjudicator: TESTNitroAdjudicator;
before(async () => {
  testNitroAdjudicator = await setupContract<TESTNitroAdjudicator>('TESTNitroAdjudicator');
});

describe('isExternalDestination', () => {
  it('verifies an external destination', async () => {
    const zerosPaddedExternalDestination =
      '0x' + 'eb89373c708B40fAeFA76e46cda92f801FAFa288'.padStart(64, '0');
    expect(
      await testNitroAdjudicator.isExternalDestination(zerosPaddedExternalDestination),
    ).to.equal(true);
  });
  it('rejects a non-external-address', async () => {
    const onesPaddedExternalDestination =
      '0x' + 'eb89373c708B40fAeFA76e46cda92f801FAFa288'.padStart(64, '1');
    expect(
      await testNitroAdjudicator.isExternalDestination(onesPaddedExternalDestination),
    ).to.equal(false);
  });
});
