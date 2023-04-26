import { ethers } from 'hardhat';
import { expect } from 'chai';

import type { ContractTransaction } from 'ethers';
import type { RandomTestConsumer } from '../../../typechain-types';

const INVALID_WEIGHTS = 'InvalidWeights';

const SEED_TIME_OF_LIFE = 32;

const MAX_NUM_RUNS = 500;
const MAX_NUM = 5;

const WEIGHTS_RUNS = 1000;
const WEIGHTS = [5, 3, 2];
const WEIGHTS_SUM = WEIGHTS.reduce((acc, curr) => acc + curr, 0);
const PRECISION_MULTIPLIER = 0.5;

describe('Random', () => {
  let RandomConsumer: RandomTestConsumer;

  before(async () => {
    const RandomConsumerFactory = await ethers.getContractFactory('RandomTestConsumer');
    RandomConsumer = (await RandomConsumerFactory.deploy()) as RandomTestConsumer;
    await RandomConsumer.deployed();
  });

  const extractSeed = async (txPromise: Promise<ContractTransaction>): Promise<string> => {
    const tx = await txPromise;
    const receipt = await tx.wait();
    return receipt.events?.[0].args?.[0] as string;
  };

  describe('random', () => {
    const resultsDistibution = Array.from({ length: MAX_NUM }).fill(0) as number[];

    before(async () => {
      let [seedChunk, newSeed] = ['', ''];
      for (let i = 0; i < MAX_NUM_RUNS; i++) {
        if (i % SEED_TIME_OF_LIFE == 0) {
          newSeed = await extractSeed(RandomConsumer.randomSeed());
        }

        [seedChunk, newSeed] = await RandomConsumer.rotateSeedChunk(newSeed);
        const randomBN = await RandomConsumer.random(MAX_NUM, seedChunk);
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
      const randomBN = await RandomConsumer.random(bigMaxNum, seed);
      const randomBN2 = await RandomConsumer.random(bigMaxNum, seed);
      expect(randomBN).to.be.equal(randomBN2);
    });
  });

  describe('randomWeightedNumber', () => {
    describe('success', () => {
      const resultsDistibution = Array.from({ length: MAX_NUM }).fill(0) as number[];

      before(async () => {
        let [seedChunk, newSeed] = ['', ''];
        for (let i = 0; i < WEIGHTS_RUNS; i++) {
          if (i % SEED_TIME_OF_LIFE == 0) {
            newSeed = await extractSeed(RandomConsumer.randomSeed());
          }

          [seedChunk, newSeed] = await RandomConsumer.rotateSeedChunk(newSeed);
          const randomBN = await RandomConsumer.randomWeightedNumber(WEIGHTS, seedChunk);
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
        await expect(RandomConsumer.randomWeightedNumber([], '0xaabbcc'))
          .to.be.revertedWithCustomError(RandomConsumer, INVALID_WEIGHTS)
          .withArgs([]);
      });
    });
  });
});
