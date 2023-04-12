import { assert, expect } from 'chai';
import { ethers } from 'hardhat';

import type { GenomeTestConsumer } from '../../../typechain-types';

/*
 * Genes are written right to left, meaning first gene is written on the right of the genome
 *
 * 00000001|00000010|00000011
 *   Body    Head     Rarity
 */

describe('Genome', () => {
  let GenomeConsumer: GenomeTestConsumer;

  before(async () => {
    const GenomeConsumerFactory = await ethers.getContractFactory('GenomeTestConsumer');
    GenomeConsumer = (await GenomeConsumerFactory.deploy()) as GenomeTestConsumer;
    await GenomeConsumer.deployed();
  });

  describe('getFlags', () => {
    it('return correct flags on 30st byte', async () => {
      expect(await GenomeConsumer.getFlags(0b1n << 240n)).to.equal(0b1);
      expect(await GenomeConsumer.getFlags(0b10n << 240n)).to.equal(0b10);
      expect(await GenomeConsumer.getFlags(0b1010n << 240n)).to.equal(0b1010);
      expect(await GenomeConsumer.getFlags(0b1100_1010n << 240n)).to.equal(0b1100_1010);
    });
  });

  describe('getFlag', () => {
    it('return correct true flag', async () => {
      expect(await GenomeConsumer.getFlag(0b1n << 240n, 0b1)).to.be.true;
      expect(await GenomeConsumer.getFlag(0b101n << 240n, 0b1)).to.be.true;
      expect(await GenomeConsumer.getFlag(0b101n << 240n, 0b100)).to.be.true;
      expect(await GenomeConsumer.getFlag(0b1101n << 240n, 0b1000)).to.be.true;
      expect(await GenomeConsumer.getFlag(0b1010_1000n << 240n, 0b10_0000)).to.be.true;
      expect(await GenomeConsumer.getFlag(0b1000_0000n << 240n, 0b1000_0000)).to.be.true;
    });

    it('return correct false flag', async () => {
      expect(await GenomeConsumer.getFlag(0b0n << 240n, 0b1)).to.be.false;
      expect(await GenomeConsumer.getFlag(0b101n << 240n, 0b10)).to.be.false;
      expect(await GenomeConsumer.getFlag(0b101n << 240n, 0b1000)).to.be.false;
      expect(await GenomeConsumer.getFlag(0b1101n << 240n, 0b10)).to.be.false;
      expect(await GenomeConsumer.getFlag(0b1010_1000n << 240n, 0b100_0000)).to.be.false;
      expect(await GenomeConsumer.getFlag(0b1000_0000n << 240n, 0b10)).to.be.false;
    });
  });

  describe('setFlag', () => {
    it('can set flag to true', async () => {
      expect(await GenomeConsumer.setFlag(0, 0b1, true)).to.equal(0b1n << 240n);
      expect(await GenomeConsumer.setFlag(0, 0b10, true)).to.equal(0b10n << 240n);
      expect(await GenomeConsumer.setFlag(0, 0b100, true)).to.equal(0b100n << 240n);
      expect(await GenomeConsumer.setFlag(0, 0b1000, true)).to.equal(0b1000n << 240n);
      expect(await GenomeConsumer.setFlag(0, 0b1000_0000, true)).to.equal(0b1000_0000n << 240n);
    });

    it('can set flat to false', async () => {
      expect(await GenomeConsumer.setFlag(0, 0b1, false)).to.equal(0);
      expect(await GenomeConsumer.setFlag(0, 0b10, false)).to.equal(0);
      expect(await GenomeConsumer.setFlag(0, 0b100, false)).to.equal(0);
      expect(await GenomeConsumer.setFlag(0, 0b1000, false)).to.equal(0);
      expect(await GenomeConsumer.setFlag(0, 0b1000_0000, false)).to.equal(0);
    });

    it('can overwrite existing flag', async () => {
      expect(await GenomeConsumer.setFlag(0b101n << 240n, 0b1, false)).to.equal(0b100n << 240n);
      expect(await GenomeConsumer.setFlag(0b101n << 240n, 0b10, false)).to.equal(0b101n << 240n);
      expect(await GenomeConsumer.setFlag(0b101n << 240n, 0b100, false)).to.equal(0b001n << 240n);
      expect(await GenomeConsumer.setFlag(0b101n << 240n, 0b10, false)).to.equal(0b101n << 240n);
      expect(await GenomeConsumer.setFlag(0b1111n << 240n, 0b1000, false)).to.equal(
        0b0111n << 240n,
      );

      expect(await GenomeConsumer.setFlag(0b1n << 240n, 0b1, true)).to.equal(0b1n << 240n);
      expect(await GenomeConsumer.setFlag(0b101n << 240n, 0b10, true)).to.equal(0b111n << 240n);
      expect(await GenomeConsumer.setFlag(0b101n << 240n, 0b100, true)).to.equal(0b101n << 240n);
      expect(await GenomeConsumer.setFlag(0b101n << 240n, 0b1000, true)).to.equal(0b1101n << 240n);
    });
  });

  describe('setGene', () => {
    it('can set 0 gene value', async () => {
      const genome = 0b0;
      expect(await GenomeConsumer.setGene(genome, 0, 1)).to.equal(0b1);
      expect(await GenomeConsumer.setGene(genome, 0, 2)).to.equal(0b10);
      expect(await GenomeConsumer.setGene(genome, 0, 4)).to.equal(0b100);
    });

    it('can set 1st gene value', async () => {
      expect(await GenomeConsumer.setGene(0, 1, 1)).to.equal(0b1_0000_0000);

      const genome = 0b1;
      expect(await GenomeConsumer.setGene(genome, 1, 1)).to.equal(0b1_0000_0001);
      expect(await GenomeConsumer.setGene(genome, 1, 2)).to.equal(0b10_0000_0001);
      expect(await GenomeConsumer.setGene(genome, 1, 4)).to.equal(0b100_0000_0001);
    });

    it('can set 2nd gene value', async () => {
      expect(await GenomeConsumer.setGene(0, 2, 1)).to.equal(0b1_0000_0000_0000_0000);

      const genome = 0b1;
      expect(await GenomeConsumer.setGene(genome, 2, 1)).to.equal(0b1_0000_0000_0000_0001);
      expect(await GenomeConsumer.setGene(genome, 2, 2)).to.equal(0b10_0000_0000_0000_0001);
      expect(await GenomeConsumer.setGene(genome, 2, 4)).to.equal(0b100_0000_0000_0000_0001);
    });

    it('can set 5th gene value', async () => {
      expect(await GenomeConsumer.setGene(0, 5, 1)).to.equal(
        0b1_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000,
      );

      const genome = 0b1_0000_0001;
      expect(await GenomeConsumer.setGene(genome, 5, 1)).to.equal(
        0b1_0000_0000_0000_0000_0000_0000_0000_0001_0000_0001,
      );
      expect(await GenomeConsumer.setGene(genome, 5, 2)).to.equal(
        0b10_0000_0000_0000_0000_0000_0000_0000_0001_0000_0001,
      );
      expect(await GenomeConsumer.setGene(genome, 5, 4)).to.equal(
        0b100_0000_0000_0000_0000_0000_0000_0000_0001_0000_0001,
      );
    });

    it('revert on setting 32nd gene value', async () => {
      await expect(GenomeConsumer.setGene(0, 32, 1)).to.be.reverted;
    });

    it('revert on setting value bigger than 255', async () => {
      try {
        await GenomeConsumer.setGene(0, 0, 256);
        assert(false);
      } catch {
        assert(true);
      }
    });
  });

  describe('getGene', () => {
    it('can get 0 gene value', async () => {
      expect(await GenomeConsumer.getGene(0b0, 0)).to.equal(0);

      expect(await GenomeConsumer.getGene(0b1, 0)).to.equal(1);
      expect(await GenomeConsumer.getGene(0b10, 0)).to.equal(2);
      expect(await GenomeConsumer.getGene(0b11, 0)).to.equal(3);
      expect(await GenomeConsumer.getGene(0b1111_1111, 0)).to.equal(255);
    });

    it('can get 1st gene value', async () => {
      expect(await GenomeConsumer.getGene(0b0_0001_0001, 1)).to.equal(0);

      expect(await GenomeConsumer.getGene(0b1_0001_0001, 1)).to.equal(1);
      expect(await GenomeConsumer.getGene(0b10_0001_0001, 1)).to.equal(2);
      expect(await GenomeConsumer.getGene(0b11_0001_0001, 1)).to.equal(3);
      expect(await GenomeConsumer.getGene(0b1111_1111_0001_0001, 1)).to.equal(255);
    });

    it('can get 2nd gene value', async () => {
      expect(await GenomeConsumer.getGene(0b0_0000_0001_0000_0001, 2)).to.equal(0);

      expect(await GenomeConsumer.getGene(0b1_0000_0001_0000_0001, 2)).to.equal(1);
      expect(await GenomeConsumer.getGene(0b10_0000_0001_0000_0001, 2)).to.equal(2);
      expect(await GenomeConsumer.getGene(0b11_0000_0001_0000_0001, 2)).to.equal(3);
      expect(await GenomeConsumer.getGene(0b1111_1111_0000_0001_0000_0001, 2)).to.equal(255);
    });

    it('can get 4th gene value', async () => {
      expect(await GenomeConsumer.getGene(0b0_0000_0001_0000_0001_0000_0001_0000_0001, 4)).to.equal(
        0,
      );

      expect(await GenomeConsumer.getGene(0b1_0000_0001_0000_0001_0000_0001_0000_0001, 4)).to.equal(
        1,
      );
      expect(
        await GenomeConsumer.getGene(0b10_0000_0001_0000_0001_0000_0001_0000_0001, 4),
      ).to.equal(2);
      expect(
        await GenomeConsumer.getGene(0b11_0000_0001_0000_0001_0000_0001_0000_0001, 4),
      ).to.equal(3);
      expect(
        await GenomeConsumer.getGene(0b1111_1111_0000_0001_0000_0001_0000_0001_0000_0001, 4),
      ).to.equal(255);
    });

    it('revert on getting 32nd gene value', async () => {
      await expect(GenomeConsumer.getGene(0, 32)).to.be.reverted;
    });
  });

  describe('maxGene', () => {
    const arr: number[] = Array.from({ length: 5 });

    it('gene 0, all 0', async () => {
      expect(await GenomeConsumer.maxGene(arr.fill(0b0), 0)).to.equal(0);
    });

    it('gene 0, all 0b11', async () => {
      expect(await GenomeConsumer.maxGene(arr.fill(0b11), 0)).to.equal(0b11);
    });

    it('gene 1, all 1', async () => {
      expect(await GenomeConsumer.maxGene(arr.fill(0b1_0000_0001_0000_0001), 1)).to.equal(1);
    });

    it('gene 0, ascending order', async () => {
      expect(await GenomeConsumer.maxGene([0b0, 0b1, 0b10, 0b10, 0b11], 0)).to.equal(0b11);
    });

    it('gene 0, random order', async () => {
      expect(await GenomeConsumer.maxGene([0b0, 0b11, 0b10, 0b10, 0b0], 0)).to.equal(0b11);
    });

    it('gene 1, random order', async () => {
      expect(
        await GenomeConsumer.maxGene(
          [
            0b1_0000_0011_0000_0001, 0b1_0000_0010_0000_0001, 0b1_0000_0001_0000_0001,
            0b1_0000_0010_0000_0001, 0b1,
          ],
          1,
        ),
      ).to.equal(0b11);
    });
  });

  describe('geneValuesAreEqual', () => {
    it('true when all equal, gene 0', async () => {
      expect(
        await GenomeConsumer.geneValuesAreEqual(
          [0b11, 0b11, 0b1_0000_0011, 0b1_0000_0001_0000_0011, 0b11],
          0,
        ),
      ).to.be.true;
    });

    it('true when all equal, gene 2', async () => {
      expect(
        await GenomeConsumer.geneValuesAreEqual(
          [
            0b101_0000_0001_1000_0000, 0b101_0000_0010_0100_0000, 0b101_0000_0100_0010_0000,
            0b101_0000_1000_0001_0000, 0b101_0001_0000_0000_1000,
          ],
          2,
        ),
      ).to.be.true;
    });

    it('false when not all equal, gene 0', async () => {
      expect(await GenomeConsumer.geneValuesAreEqual([0b0, 0b10, 0b1, 0b1_0000_0000, 0b0], 0)).to.be
        .false;
    });

    it('false when not all equal, gene 0', async () => {
      expect(await GenomeConsumer.geneValuesAreEqual([0b1, 0b0, 0b0, 0b11, 0b1_0000_0010], 0)).to.be
        .false;
    });

    it('false when not all equal, gene 2', async () => {
      expect(
        await GenomeConsumer.geneValuesAreEqual(
          [
            0b101_0000_0001_0000_0000, 0b1, 0b11_0000_0001_0000_0000, 0b1_0000_0001_0000_0000,
            0b110_0000_0001_0000_0000,
          ],
          2,
        ),
      ).to.be.false;
    });
  });

  describe('geneValuesAreUnique', () => {
    it('true when all are unique, gene 0', async () => {
      expect(await GenomeConsumer.geneValuesAreUnique([0b0, 0b1, 0b10, 0b11, 0b100], 0)).to.be.true;
    });

    it('true when all are unique, gene 1', async () => {
      expect(
        await GenomeConsumer.geneValuesAreUnique(
          [0b1_0000_0000, 0b11_0000_0000, 0b111_0000_0000, 0b1111_0000_0000, 0b10_0000_0000],
          1,
        ),
      ).to.be.true;
    });

    it('false when not all are unique, gene 0', async () => {
      expect(await GenomeConsumer.geneValuesAreUnique([0b0, 0b1, 0b10, 0b11, 0b0], 0)).to.be.false;
    });

    it('true when all are unique, gene 0', async () => {
      expect(await GenomeConsumer.geneValuesAreUnique([0b1, 0b10, 0b11, 0b11, 0b100], 0)).to.be
        .false;
    });

    it('false when not all are unique, gene 1', async () => {
      expect(
        await GenomeConsumer.geneValuesAreUnique(
          [0b1_1000_0000, 0b11_0100_0000, 0b11_0010_0000, 0b1111_0001_0000, 0b10_0000_1000],
          1,
        ),
      ).to.be.false;
    });
  });
});
