import { expect } from 'chai';
import { utils } from 'ethers';

import { setup } from './setup';
import {
  Collections,
  DucklingGenes,
  GeneDistrTypes,
  Rarities,
  collectionGeneIdx,
  collectionsGeneDistributionTypes,
  collectionsGeneValuesNum,
  generativeGenesOffset,
  rarityGeneIdx,
} from './config';
import { Genome } from './genome';

import type { DuckyGenomeTestConsumer } from '../../../../typechain-types';

const SEED = utils.id('seed');
const BIT_SLICE = '0xaabbcc';

describe('DuckyGenome minting', () => {
  let DuckyGenome: DuckyGenomeTestConsumer;

  beforeEach(async () => {
    ({ DuckyGenome } = await setup());
  });

  const generateAndSetGenes = async (
    genome: bigint,
    collectionId: Collections,
    geneValues: number[],
    geneDistrTypes: number,
    seed: string = SEED,
  ): Promise<bigint> => {
    const tx = await DuckyGenome.generateAndSetGenes(
      genome,
      collectionId,
      geneValues,
      geneDistrTypes,
      seed,
    );
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  const generateAndSetGene = async (
    genome: bigint,
    geneIx: number,
    geneValuesNum: number,
    distrType: GeneDistrTypes,
    bitSlice: string = BIT_SLICE,
  ): Promise<bigint> => {
    const tx = await DuckyGenome.generateAndSetGene(
      genome,
      geneIx,
      geneValuesNum,
      distrType,
      bitSlice,
    );
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  // does not include flags and magic number
  // eslint-disable-next-line sonarjs/cognitive-complexity
  describe('generateAndSetGenes', () => {
    describe('Duckling', () => {
      const baseDucklingGenome = new Genome(0).setGene(collectionGeneIdx, Collections.Duckling);

      const geneValuesNum = [...collectionsGeneValuesNum[Collections.Duckling]];
      const geneDistrTypes = collectionsGeneDistributionTypes[Collections.Duckling];

      const generateAndSetDucklingGenes = async (genome: bigint): Promise<bigint> => {
        return await generateAndSetGenes(
          genome,
          Collections.Duckling,
          geneValuesNum,
          geneDistrTypes,
          SEED,
        );
      };

      it('has correct numbers of genes', async () => {
        const ducklingGenome = baseDucklingGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const genome = await generateAndSetDucklingGenes(ducklingGenome);

        // as not default values start from 1 and the last gene is not default,
        // the number of genes is equal to number of bytes in genome returned
        const numberOfGenes = Math.ceil(genome.toString(2).length / 8);
        expect(numberOfGenes).to.equal(generativeGenesOffset + geneValuesNum.length);
      });

      it('sets default values for Common', async () => {
        const ducklingGenome = baseDucklingGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const _genome = await generateAndSetDucklingGenes(ducklingGenome);
        const genome = new Genome(_genome);

        expect(genome.getGene(DucklingGenes.Body)).to.equal(0);
        expect(genome.getGene(DucklingGenes.Head)).to.equal(0);
      });

      it('sets default values for Rare', async () => {
        const ducklingGenome = baseDucklingGenome.setGene(rarityGeneIdx, Rarities.Rare).genome;
        const _genome = await generateAndSetDucklingGenes(ducklingGenome);
        const genome = new Genome(_genome);

        expect(genome.getGene(DucklingGenes.Body)).to.not.equal(0);
        expect(genome.getGene(DucklingGenes.Head)).to.equal(0);
      });

      it('not defaulted values start at 1', async () => {
        const ducklingGenome = baseDucklingGenome.setGene(rarityGeneIdx, Rarities.Epic).genome;
        const _genome = await generateAndSetDucklingGenes(ducklingGenome);
        const genome = new Genome(_genome);

        for (let i = 0; i < geneValuesNum.length; i++) {
          const gene = genome.getGene(generativeGenesOffset + i);
          expect(gene).to.be.greaterThan(0);
        }
      });

      it('does not exceed max gene values', async () => {
        const ducklingGenome = baseDucklingGenome.setGene(rarityGeneIdx, Rarities.Epic).genome;
        const _genome = await generateAndSetDucklingGenes(ducklingGenome);
        const genome = new Genome(_genome);

        for (const [i, valuesNum] of geneValuesNum.entries()) {
          const gene = genome.getGene(generativeGenesOffset + i);
          expect(gene).to.be.within(1, valuesNum);
        }
      });
    });

    describe('Zombeak', () => {
      const baseZombeakGenome = new Genome(0).setGene(collectionGeneIdx, Collections.Zombeak);

      const geneValuesNum = [...collectionsGeneValuesNum[Collections.Zombeak]];
      const geneDistrTypes = collectionsGeneDistributionTypes[Collections.Zombeak];

      const generateAndSetZombeakGenes = async (genome: bigint): Promise<bigint> => {
        return await generateAndSetGenes(
          genome,
          Collections.Zombeak,
          geneValuesNum,
          geneDistrTypes,
          SEED,
        );
      };

      it('has correct numbers of genes', async () => {
        const zombeakGenome = baseZombeakGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const genome = await generateAndSetZombeakGenes(zombeakGenome);

        // as not default values start from 1 and the last gene is not default,
        // the number of genes is equal to number of bytes in genome returned
        const numberOfGenes = Math.ceil(genome.toString(2).length / 8);
        expect(numberOfGenes).to.equal(generativeGenesOffset + geneValuesNum.length);
      });

      it('not defaulted values start at 1', async () => {
        const zombeakGenome = baseZombeakGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const _genome = await generateAndSetZombeakGenes(zombeakGenome);
        const genome = new Genome(_genome);

        for (let i = 0; i < geneValuesNum.length; i++) {
          const gene = genome.getGene(generativeGenesOffset + i);
          expect(gene).to.be.greaterThan(0);
        }
      });

      it('does not exceed max gene values', async () => {
        const zombeakGenome = baseZombeakGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const _genome = await generateAndSetZombeakGenes(zombeakGenome);
        const genome = new Genome(_genome);

        for (const [i, valuesNum] of geneValuesNum.entries()) {
          const gene = genome.getGene(generativeGenesOffset + i);
          expect(gene).to.be.within(1, valuesNum);
        }
      });
    });

    describe('Mythic', () => {
      const baseMythicGenome = new Genome(0).setGene(collectionGeneIdx, Collections.Mythic);

      const geneValuesNum = [...collectionsGeneValuesNum[Collections.Mythic]];
      const geneDistrTypes = collectionsGeneDistributionTypes[Collections.Mythic];

      const generateAndSetMythicGenes = async (genome: bigint): Promise<bigint> => {
        return await generateAndSetGenes(
          genome,
          Collections.Mythic,
          geneValuesNum,
          geneDistrTypes,
          SEED,
        );
      };

      it('has correct numbers of genes', async () => {
        const mythicGenome = baseMythicGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const genome = await generateAndSetMythicGenes(mythicGenome);

        // as not default values start from 1 and the last gene is not default,
        // the number of genes is equal to number of bytes in genome returned
        const numberOfGenes = Math.ceil(genome.toString(2).length / 8);
        expect(numberOfGenes).to.equal(generativeGenesOffset + geneValuesNum.length);
      });

      it('not defaulted values start at 1', async () => {
        const mythicGenome = baseMythicGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const _genome = await generateAndSetMythicGenes(mythicGenome);
        const genome = new Genome(_genome);

        for (let i = 0; i < geneValuesNum.length; i++) {
          const gene = genome.getGene(generativeGenesOffset + i);
          expect(gene).to.be.greaterThan(0);
        }
      });

      it('does not exceed max gene values', async () => {
        const mythicGenome = baseMythicGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const _genome = await generateAndSetMythicGenes(mythicGenome);
        const genome = new Genome(_genome);

        for (const [i, valuesNum] of geneValuesNum.entries()) {
          const gene = genome.getGene(generativeGenesOffset + i);
          expect(gene).to.be.within(1, valuesNum);
        }
      });
    });
  });

  describe('generateAndSetGene', () => {
    it('values start at 1', async () => {
      const geneIdx = 0;
      const geneValuesNum = 15;

      const _genome = await generateAndSetGene(
        BigInt(0),
        geneIdx,
        geneValuesNum,
        GeneDistrTypes.Even,
      );
      const genome = new Genome(_genome);
      expect(genome.getGene(geneIdx)).to.be.greaterThan(0);
    });
  });
});
