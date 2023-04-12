import { expect } from 'chai';

import { randomGenome, randomMaxNum } from './helpers';
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
