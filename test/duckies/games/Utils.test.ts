import { ethers } from 'hardhat';
import { expect } from 'chai';
import { utils } from 'ethers';

import { randomBytes32 } from '../../helpers/common';

import type { UtilsTestConsumer } from '../../../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

const INVALID_WEIGHTS = 'InvalidWeights';

const SEED_TIME_OF_LIFE = 32;

const MAX_NUM_RUNS = 500;
const MAX_NUM = 5;

const WEIGHTS_RUNS = 1000;
const WEIGHTS = [5, 3, 2];
const WEIGHTS_SUM = WEIGHTS.reduce((acc, curr) => acc + curr, 0);
const PRECISION_MULTIPLIER = 0.5;

describe('Utils', () => {
  let Someone: SignerWithAddress;
  let Someother: SignerWithAddress;

  let Utils: UtilsTestConsumer;

  before(async () => {
    [Someone, Someother] = await ethers.getSigners();

    const UtilsTestFactory = await ethers.getContractFactory('UtilsTestConsumer');
    Utils = (await UtilsTestFactory.deploy()) as UtilsTestConsumer;
    await Utils.deployed();
  });

  describe('max', () => {
    const resultsDistibution = Array.from({ length: MAX_NUM }).fill(0) as number[];

    before(async () => {
      let [bitSlice, newSeed] = ['', ''];
      for (let i = 0; i < MAX_NUM_RUNS; i++) {
        if (i % SEED_TIME_OF_LIFE == 0) {
          newSeed = randomBytes32();
        }

        [bitSlice, newSeed] = await Utils.shiftSeedSlice(newSeed);
        const randomBN = await Utils.max(bitSlice, MAX_NUM);
        resultsDistibution[randomBN.toNumber()]++;
      }
    });

    it('generated number is < supplied', () => {
      expect(resultsDistibution.length == MAX_NUM);
    });

    it('correct distribution', () => {
      const sum = resultsDistibution.reduce((acc, curr) => acc + curr, 0);
      const left = (1 / MAX_NUM) * (1 - PRECISION_MULTIPLIER);
      const right = (1 / MAX_NUM) * (1 + PRECISION_MULTIPLIER);

      for (const num of resultsDistibution) {
        const freq = num / sum;
        const thFreq = 1 / MAX_NUM;
        console.log('Frequency: actual', freq, ', theoretical:', thFreq);
        expect(freq).to.be.greaterThanOrEqual(left);
        expect(freq).to.be.lessThanOrEqual(right);
      }
    });

    it('generates same number for same seed', async () => {
      const bigMaxNum = 424_242;
      const seed = '0xaabbcc';
      const randomBN = await Utils.max(seed, bigMaxNum);
      const randomBN2 = await Utils.max(seed, bigMaxNum);
      expect(randomBN).to.be.equal(randomBN2);
    });
  });

  describe('randomWeightedNumber', () => {
    describe('success', () => {
      const resultsDistibution = Array.from({ length: MAX_NUM }).fill(0) as number[];

      before(async () => {
        let [bitSlice, newSeed] = ['', ''];
        for (let i = 0; i < WEIGHTS_RUNS; i++) {
          if (i % SEED_TIME_OF_LIFE == 0) {
            newSeed = randomBytes32();
          }

          [bitSlice, newSeed] = await Utils.shiftSeedSlice(newSeed);
          const randomBN = await Utils.randomWeightedNumber(WEIGHTS, bitSlice);
          resultsDistibution[randomBN.toNumber()]++;
        }
      });

      it('generated number is in correct range', () => {
        expect(resultsDistibution.length == WEIGHTS.length);
      });

      it('correct distribution', () => {
        const sum = resultsDistibution.reduce((acc, curr) => acc + curr, 0);
        for (const [i, WEIGHT] of WEIGHTS.entries()) {
          const freq = resultsDistibution[i] / sum;
          const thFreq = WEIGHT / WEIGHTS_SUM;
          const left = thFreq * (1 - PRECISION_MULTIPLIER);
          const right = thFreq * (1 + PRECISION_MULTIPLIER);

          console.log('Frequency: actual', freq, ', theoretical:', thFreq);
          expect(freq).to.be.greaterThanOrEqual(left);
          expect(freq).to.be.lessThanOrEqual(right);
        }
      });
    });

    describe('revert', () => {
      it('when empty weights array', async () => {
        await expect(Utils.randomWeightedNumber([], '0xaabbcc'))
          .to.be.revertedWithCustomError(Utils, INVALID_WEIGHTS)
          .withArgs([]);
      });
    });
  });

  describe('_requireCorrectSigner', () => {
    const coder = ethers.utils.defaultAbiCoder;

    it('success on correct signer', async () => {
      const message = '0x42';
      const encodedMessage = coder.encode(['string'], [message]);
      const signedMessage = await Someone.signMessage(
        utils.arrayify(utils.keccak256(encodedMessage)),
      );

      await Utils.requireCorrectSigner(encodedMessage, signedMessage, Someone.address);
    });

    it('revert on incorrect signer', async () => {
      const message = '0x42';
      const encodedMessage = coder.encode(['string'], [message]);
      const signedMessage = await Someone.signMessage(
        utils.arrayify(utils.keccak256(encodedMessage)),
      );

      await expect(Utils.requireCorrectSigner(encodedMessage, signedMessage, Someother.address))
        .to.be.revertedWithCustomError(Utils, 'IncorrectSigner')
        .withArgs(Someother.address, Someone.address);
    });
  });
});
