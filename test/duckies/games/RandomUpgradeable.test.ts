import { ethers } from 'hardhat';
import { expect } from 'chai';

import type { BigNumber, ContractTransaction } from 'ethers';
import type { RandomTestConsumer } from '../../../typechain-types';

const INVALID_WEIGHTS = 'InvalidWeights';

const MAX_NUM_RUNS = 500;
const MAX_NUM = 5;

const WEIGHTS_RUNS = 1000;
const WEIGHTS = [5, 3, 2];
const WEIGHTS_SUM = WEIGHTS.reduce((acc, curr) => acc + curr, 0);
const PRECISION_MULTIPLIER = 0.5;

describe('RandomUpgradeable', () => {
  let RandomConsumer: RandomTestConsumer;

  before(async () => {
    const RandomConsumerFactory = await ethers.getContractFactory('RandomTestConsumer');
    RandomConsumer = (await RandomConsumerFactory.deploy()) as RandomTestConsumer;
    await RandomConsumer.deployed();
  });

  const extractNumFromEvent = async (tx: ContractTransaction): Promise<number> => {
    const receipt = await tx.wait();
    return (receipt.events?.[0].args?.[0] as BigNumber).toNumber();
  };

  describe('randomMaxNumber', () => {
    const resultsDistibution = Array.from({ length: MAX_NUM }).fill(0) as number[];

    before(async () => {
      for (let i = 0; i < MAX_NUM_RUNS; i++) {
        resultsDistibution[
          await extractNumFromEvent(await RandomConsumer.randomMaxNumber(MAX_NUM))
        ]++;
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
  });

  describe('randomWeightedNumber', () => {
    describe('success', () => {
      const resultsDistibution = Array.from({ length: MAX_NUM }).fill(0) as number[];

      before(async () => {
        for (let i = 0; i < WEIGHTS_RUNS; i++) {
          resultsDistibution[
            await extractNumFromEvent(await RandomConsumer.randomWeightedNumber(WEIGHTS))
          ]++;
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
        await expect(RandomConsumer.randomWeightedNumber([]))
          .to.be.revertedWithCustomError(RandomConsumer, INVALID_WEIGHTS)
          .withArgs([]);
      });
    });
  });
});
