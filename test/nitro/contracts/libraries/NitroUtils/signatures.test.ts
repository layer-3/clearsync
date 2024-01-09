import { Contract, Wallet } from 'ethers';
import { ethers } from 'hardhat';
import { expect } from 'chai';
import { describe, before, it } from 'mocha';

const { arrayify, id } = ethers.utils;

import { setupContract } from '../../../test-helpers';
import { sign } from '../../../../../src/nitro/signatures';
import type { TESTNitroUtils } from '../../../../../typechain-types';

let nitroUtils: Contract & TESTNitroUtils;

before(async () => {
  nitroUtils = await setupContract<TESTNitroUtils>('TESTNitroUtils');
});

describe('_recoverSigner', () => {
  it('recovers the signer correctly', async () => {
    // Following https://docs.ethers.io/ethers.js/html/cookbook-signing.html
    const privateKey = '0x0123456789012345678901234567890123456789012345678901234567890123';
    const wallet = new Wallet(privateKey);
    const msgHash = id('Hello World');
    const msgHashBytes = arrayify(msgHash);
    const sig = await sign(wallet, msgHashBytes);
    expect(await nitroUtils.recoverSigner(msgHash, sig)).to.equal(wallet.address);
  });
});

describe('isClaimedSignedBy', () => {
  // prettier-ignore
  it('returns true when a participant bit is set', async () => {
    expect(await nitroUtils.isClaimedSignedBy(0b101     ,0)).to.equal(true);
    expect(await nitroUtils.isClaimedSignedBy(0b101     ,2)).to.equal(true);
    expect(await nitroUtils.isClaimedSignedBy(0b001     ,0)).to.equal(true);
    expect(await nitroUtils.isClaimedSignedBy(0b10000000,7)).to.equal(true);
    expect(await nitroUtils.isClaimedSignedBy(8         ,3)).to.equal(true);
  });
  // prettier-ignore
  it('returns false when a participant bit is not set', async () => {
    expect(await nitroUtils.isClaimedSignedBy(0b101     ,1)).to.equal(false);
    expect(await nitroUtils.isClaimedSignedBy(0b001     ,3)).to.equal(false);
    expect(await nitroUtils.isClaimedSignedBy(0b001     ,2)).to.equal(false);
    expect(await nitroUtils.isClaimedSignedBy(0b001     ,1)).to.equal(false);
  });
});

describe('isClaimedSignedOnlyBy', () => {
  // prettier-ignore
  it('returns true when only that participant bit is set', async () => {
    expect(await nitroUtils.isClaimedSignedOnlyBy(0b001     ,0)).to.equal(true);
    expect(await nitroUtils.isClaimedSignedOnlyBy(0b10000000,7)).to.equal(true);
    expect(await nitroUtils.isClaimedSignedOnlyBy(8         ,3)).to.equal(true);
  });
  // prettier-ignore
  it('returns false when that participant bit is not set', async () => {
    expect(await nitroUtils.isClaimedSignedOnlyBy(0b011     ,0)).to.equal(false);
    expect(await nitroUtils.isClaimedSignedOnlyBy(0b10010000,7)).to.equal(false);
    expect(await nitroUtils.isClaimedSignedOnlyBy(9         ,3)).to.equal(false);
    expect(await nitroUtils.isClaimedSignedOnlyBy(0b101     ,0)).to.equal(false);
    expect(await nitroUtils.isClaimedSignedOnlyBy(0b101     ,2)).to.equal(false);
    expect(await nitroUtils.isClaimedSignedOnlyBy(0b101     ,1)).to.equal(false);
    expect(await nitroUtils.isClaimedSignedOnlyBy(0b001     ,3)).to.equal(false);
    expect(await nitroUtils.isClaimedSignedOnlyBy(0b001     ,2)).to.equal(false);
    expect(await nitroUtils.isClaimedSignedOnlyBy(0b001     ,1)).to.equal(false);
  });
});

describe('getClaimedSignersNum', () => {
  // prettier-ignore
  it('counts the number of signers correctly', async () => {
    expect(await nitroUtils.getClaimedSignersNum(0b001)).to.equal(1)
    expect(await nitroUtils.getClaimedSignersNum(0b011)).to.equal(2)
    expect(await nitroUtils.getClaimedSignersNum(0b101)).to.equal(2)
    expect(await nitroUtils.getClaimedSignersNum(0b111)).to.equal(3)
    expect(await nitroUtils.getClaimedSignersNum(0b000)).to.equal(0)
  });
});

describe('getClaimedSignersIndices', () => {
  // prettier-ignore
  it('returns the correct indices', async () => {
    expect(await nitroUtils.getClaimedSignersIndices(0b001)).to.deep.equal([0])
    expect(await nitroUtils.getClaimedSignersIndices(0b011)).to.deep.equal([0,1])
    expect(await nitroUtils.getClaimedSignersIndices(0b101)).to.deep.equal([0,2])
    expect(await nitroUtils.getClaimedSignersIndices(0b111)).to.deep.equal([0,1,2])
    expect(await nitroUtils.getClaimedSignersIndices(0b000)).to.deep.equal([])
  });
});
