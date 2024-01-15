import { BigNumber } from 'ethers';
import { expect } from 'chai';

import { getRandomNonce, replaceAddressesAndBigNumberify } from '../../../src/nitro/helpers';

const addresses = {
  // Channels
  C: '0xCHANNEL',
  X: '0xANOTHERCHANNEL',
  // Externals
  A: '0x000EXTERNAL',
  B: '0x000ANOTHEREXTERNAL',
  ETH: '0xETH',
  TOK: '0xTOK',
};

const singleAsset = { C: 1, X: 2 };
const singleAssetReplaced = {
  '0xCHANNEL': BigNumber.from(1),
  '0xANOTHERCHANNEL': BigNumber.from(2),
};

const multiAsset = { ETH: { C: 3 }, TOK: { X: 4 } };
const multiAssetReplaced = {
  '0xETH': { '0xCHANNEL': BigNumber.from(3) },
  '0xTOK': { '0xANOTHERCHANNEL': BigNumber.from(4) },
};

describe('replaceAddressesAndBigNumberify', () => {
  it('replaces without recursion', () => {
    expect(replaceAddressesAndBigNumberify(singleAsset, addresses)).deep.equal(singleAssetReplaced);
  });
  it('replaces with one level of recursion', () => {
    expect(replaceAddressesAndBigNumberify(multiAsset, addresses)).deep.equal(multiAssetReplaced);
  });
});

describe('getRandomNonce', () => {
  it('generates hex strings representing 64 bit integers', () => {
    const result = getRandomNonce('StrictTurnTaking');
    expect(BigNumber.from(result).lt(BigNumber.from('0xffffffffffffffff'))).to.be.true;
  });
});
