import { expect } from 'chai';

import { getSignedBy, getSignersIndices, getSignersNum } from '../../../src/nitro/bitfield-utils';

describe('bitfield utils', () => {
  it('getSignersNum', () => {
    expect(getSignersNum('0')).to.equal(0);
    expect(getSignersNum('1')).to.equal(1);
    expect(getSignersNum('2')).to.equal(1);
    expect(getSignersNum('3')).to.equal(2);
    expect(getSignersNum('5')).to.equal(2);
    expect(getSignersNum('100')).to.equal(3);
    expect(getSignersNum('1000')).to.equal(6);
  });

  it('getSignersIndices', () => {
    expect(getSignersIndices('0').length).to.equal(0);
    expect(getSignersIndices('1')).to.deep.equal([0]);
    expect(getSignersIndices('2')).to.deep.equal([1]);
    expect(getSignersIndices('3')).to.deep.equal([0, 1]);
    expect(getSignersIndices('5')).to.deep.equal([0, 2]);
    expect(getSignersIndices('100')).to.deep.equal([2, 5, 6]);
    expect(getSignersIndices('1000')).to.deep.equal([3, 5, 6, 7, 8, 9]);
  });

  it('getSignedBy', () => {
    expect(getSignedBy(0)).to.equal('1');
    expect(getSignedBy(1)).to.equal('2');
    expect(getSignedBy(5)).to.equal('32');

    expect(getSignedBy([])).to.equal('0');
    expect(getSignedBy([0])).to.equal('1');
    expect(getSignedBy([1])).to.equal('2');
    expect(getSignedBy([0, 1])).to.equal('3');
    expect(getSignedBy([0, 2])).to.equal('5');
    expect(getSignedBy([2, 5, 6])).to.equal('100');
    expect(getSignedBy([3, 5, 6, 7, 8, 9])).to.equal('1000');
  });
});
