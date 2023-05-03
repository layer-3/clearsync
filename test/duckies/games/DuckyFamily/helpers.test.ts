import { expect } from 'chai';

import { bytes3, randomGenome, randomMaxNum, reverse } from './helpers';
import {
  Collections,
  DucklingGenes,
  MythicGenes,
  ZombeakGenes,
  collectionsGeneValuesNum,
  generativeGenesOffset,
  mythicAmount,
} from './config';
import { Genome } from './genome';

describe('helpers', () => {
  it('randomMaxNum', () => {
    const maxNum = 5;
    for (let i = 0; i < maxNum * 3; i++) {
      const num = randomMaxNum(maxNum);
      expect(num).to.be.within(0, maxNum);
    }

    expect(randomMaxNum(0)).to.equal(0);
  });

  it('bytes3', () => {
    expect(bytes3(0)).to.equal('0x000000');
    expect(bytes3(1)).to.equal('0x000001');
    expect(bytes3(15)).to.equal('0x00000f');
    expect(bytes3(16)).to.equal('0x000010');
    expect(bytes3(42)).to.equal('0x00002a');
    expect(bytes3(255)).to.equal('0x0000ff');
    expect(bytes3(256)).to.equal('0x000100');
    expect(bytes3(424)).to.equal('0x0001a8');
    expect(bytes3(4095)).to.equal('0x000fff');
    expect(bytes3(4096)).to.equal('0x001000');
    expect(bytes3(4242)).to.equal('0x001092');
  });

  it('reverse', () => {
    expect(reverse(0b0)).to.equal(0b0);
    expect(reverse(0b1)).to.equal(0b1);
    expect(reverse(0b10)).to.equal(0b01);
    expect(reverse(0b11)).to.equal(0b11);
    expect(reverse(0b100)).to.equal(0b001);
    expect(reverse(0b101)).to.equal(0b101);
    expect(reverse(0b110)).to.equal(0b011);
    expect(reverse(0b1010)).to.equal(0b0101);
    expect(reverse(0b11_0001_1101)).to.equal(0b10_1110_0011);
  });

  describe('randomGenome', () => {
    it('duckling', () => {
      const ducklingGeneValuesNum = collectionsGeneValuesNum[Collections.Duckling];

      for (let i = 0; i < 15; i++) {
        const genome = new Genome(randomGenome(Collections.Duckling));
        expect(genome.getGene(DucklingGenes.Collection)).to.equal(Collections.Duckling);

        for (const [i, geneValues] of ducklingGeneValuesNum.entries()) {
          expect(genome.getGene(i + generativeGenesOffset)).to.be.within(0, geneValues);
        }
      }
    });

    it('zombeak', () => {
      const zombeakGeneValuesNum = collectionsGeneValuesNum[Collections.Zombeak];

      for (let i = 0; i < 15; i++) {
        const genome = new Genome(randomGenome(Collections.Zombeak));
        expect(genome.getGene(ZombeakGenes.Collection)).to.equal(Collections.Zombeak);

        for (const [i, geneValues] of zombeakGeneValuesNum.entries()) {
          expect(genome.getGene(i + generativeGenesOffset)).to.be.within(0, geneValues);
        }
      }
    });

    it('mythic', () => {
      for (let i = 0; i < 15; i++) {
        const genome = new Genome(randomGenome(Collections.Mythic));
        expect(genome.getGene(MythicGenes.Collection)).to.equal(Collections.Mythic);
        expect(genome.getGene(MythicGenes.UniqId)).to.be.within(0, mythicAmount);
      }
    });

    it('setting genes', () => {
      for (let i = 0; i < 15; i++) {
        const genome = new Genome(
          randomGenome(Collections.Duckling, {
            [DucklingGenes.Rarity]: 0,
            [DucklingGenes.Color]: 1,
            [DucklingGenes.Family]: 2,
          }),
        );

        expect(genome.getGene(0)).to.equal(Collections.Duckling);
        expect(genome.getGene(DucklingGenes.Rarity)).to.equal(0);
        expect(genome.getGene(DucklingGenes.Color)).to.equal(1);
        expect(genome.getGene(DucklingGenes.Family)).to.equal(2);
      }
    });
  });
});
