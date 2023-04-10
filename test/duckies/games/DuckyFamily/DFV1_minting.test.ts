import { expect } from 'chai';

import { setup } from './setup';
import {
  Collections,
  DucklingGenes,
  MAX_PECULIARITY,
  MYTHIC_DISPERSION,
  MythicGenes,
  Rarities,
  collectionGeneIdx,
  collectionsGeneValuesNum,
  generativeGenesOffset,
  mythicAmount,
  rarityGeneIdx,
} from './config';
import { Genome } from './genome';

import type { TESTDuckyFamilyV1 } from '../../../../typechain-types';

describe('DuckyFamilyV1 minting', () => {
  let Game: TESTDuckyFamilyV1;

  beforeEach(async () => {
    ({ Game } = await setup());
  });

  const generateGenome = async (collectionId: Collections): Promise<bigint> => {
    const tx = await Game.generateGenome(collectionId);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  const generateAndSetGenes = async (
    genome: bigint,
    collectionId: Collections,
  ): Promise<bigint> => {
    const tx = await Game.generateAndSetGenes(genome, collectionId);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  const generateMythicGenome = async (genomes: bigint[]): Promise<bigint> => {
    const tx = await Game.generateMythicGenome(genomes);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  describe('generateGenome', () => {
    beforeEach(async () => {
      await Game.setRarityChances([1000, 0, 0, 0]);
    });

    it('set Duckling collectionId and rarity', async () => {
      const _genome = await generateGenome(Collections.Duckling);
      const genome = new Genome(_genome);
      expect(genome.getGene(collectionGeneIdx)).to.equal(Collections.Duckling);
      expect(genome.getGene(rarityGeneIdx)).to.equal(0);
    });

    it('set Zombeak collectionId and rarity', async () => {
      const _genome = await generateGenome(Collections.Zombeak);
      const genome = new Genome(_genome);
      expect(genome.getGene(collectionGeneIdx)).to.equal(Collections.Zombeak);
      expect(genome.getGene(rarityGeneIdx)).to.equal(0);
    });

    it('revert on collectionId not Duckling or Zombeak', async () => {
      const notDucklingOrZombeak = 2;
      await expect(Game.generateGenome(2))
        .to.be.revertedWithCustomError(Game, 'MintingRulesViolated')
        .withArgs(notDucklingOrZombeak, 1);
    });
  });

  describe('generateAndSetGenes', () => {
    describe('Duckling', () => {
      const baseDucklingGenome = new Genome(0).setGene(collectionGeneIdx, Collections.Duckling);

      const geneValuesNum = collectionsGeneValuesNum[Collections.Duckling];

      it('has correct numbers of genes', async () => {
        const ducklingGenome = baseDucklingGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const genome = await generateAndSetGenes(ducklingGenome, Collections.Duckling);

        // as not default values start from 1 and the last gene is not default,
        // the number of genes is equal to number of bytes in genome returned
        const numberOfGenes = Math.ceil(genome.toString(2).length / 8);
        expect(numberOfGenes).to.equal(generativeGenesOffset + geneValuesNum.length);
      });

      it('sets default values for Common', async () => {
        const ducklingGenome = baseDucklingGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const _genome = await generateAndSetGenes(ducklingGenome, Collections.Duckling);
        const genome = new Genome(_genome);

        expect(genome.getGene(DucklingGenes.Body)).to.equal(0);
        expect(genome.getGene(DucklingGenes.Head)).to.equal(0);
      });

      it('sets default values for Rare', async () => {
        const ducklingGenome = baseDucklingGenome.setGene(rarityGeneIdx, Rarities.Rare).genome;
        const _genome = await generateAndSetGenes(ducklingGenome, Collections.Duckling);
        const genome = new Genome(_genome);

        expect(genome.getGene(DucklingGenes.Body)).to.not.equal(0);
        expect(genome.getGene(DucklingGenes.Head)).to.equal(0);
      });

      it('not defaulted values start at 1', async () => {
        const ducklingGenome = baseDucklingGenome.setGene(rarityGeneIdx, Rarities.Epic).genome;
        const _genome = await generateAndSetGenes(ducklingGenome, Collections.Duckling);
        const genome = new Genome(_genome);

        for (let i = 0; i < geneValuesNum.length; i++) {
          const gene = genome.getGene(generativeGenesOffset + i);
          expect(gene).to.be.greaterThan(0);
        }
      });

      it('does not exceed max gene values', async () => {
        const ducklingGenome = baseDucklingGenome.setGene(rarityGeneIdx, Rarities.Epic).genome;
        const _genome = await generateAndSetGenes(ducklingGenome, Collections.Duckling);
        const genome = new Genome(_genome);

        for (const [i, valuesNum] of geneValuesNum.entries()) {
          const gene = genome.getGene(generativeGenesOffset + i);
          expect(gene).to.be.within(1, valuesNum);
        }
      });
    });

    describe('Zombeak', () => {
      const baseZombeakGenome = new Genome(0).setGene(collectionGeneIdx, Collections.Zombeak);
      const geneValuesNum = collectionsGeneValuesNum[Collections.Zombeak];

      it('has correct numbers of genes', async () => {
        const zombeakGenome = baseZombeakGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const genome = await generateAndSetGenes(zombeakGenome, Collections.Zombeak);

        // as not default values start from 1 and the last gene is not default,
        // the number of genes is equal to number of bytes in genome returned
        const numberOfGenes = Math.ceil(genome.toString(2).length / 8);
        expect(numberOfGenes).to.equal(generativeGenesOffset + geneValuesNum.length);
      });

      it('not defaulted values start at 1', async () => {
        const zombeakGenome = baseZombeakGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const _genome = await generateAndSetGenes(zombeakGenome, Collections.Zombeak);
        const genome = new Genome(_genome);

        for (let i = 0; i < geneValuesNum.length; i++) {
          const gene = genome.getGene(generativeGenesOffset + i);
          expect(gene).to.be.greaterThan(0);
        }
      });

      it('does not exceed max gene values', async () => {
        const zombeakGenome = baseZombeakGenome.setGene(rarityGeneIdx, Rarities.Common).genome;
        const _genome = await generateAndSetGenes(zombeakGenome, Collections.Zombeak);
        const genome = new Genome(_genome);

        for (const [i, valuesNum] of geneValuesNum.entries()) {
          const gene = genome.getGene(generativeGenesOffset + i);
          expect(gene).to.be.within(1, valuesNum);
        }
      });
    });
  });

  describe.only('generateMythicGenome', () => {
    const zeroedGenome = new Genome(0).genome;

    type generateGenomeArgsType = [bigint, bigint, bigint, bigint, bigint];
    const maxUniqId = mythicAmount - 1;

    const genomesWithTheoUniqId = (theoUniqId: number): generateGenomeArgsType => {
      const peculiarity = Math.ceil((theoUniqId / maxUniqId) * MAX_PECULIARITY);
      const genome = Genome.from([0, 0, 0, 0, peculiarity]).genome;
      return Array.from({ length: 5 }).fill(genome) as generateGenomeArgsType;
    };

    it('correct range for zero peculiarity', async () => {
      const zeroedGenomes = Array.from({ length: 5 }).fill(zeroedGenome);
      const _genome = await generateMythicGenome(zeroedGenomes as generateGenomeArgsType);
      const genome = new Genome(_genome);
      expect(genome.getGene(MythicGenes.UniqId)).to.be.within(0, MYTHIC_DISPERSION);
    });

    it('correct range for low UniqId', async () => {
      const theoUniqId = 2;
      const _genome = await generateMythicGenome(genomesWithTheoUniqId(theoUniqId));
      const genome = new Genome(_genome);
      expect(genome.getGene(MythicGenes.UniqId)).to.be.within(0, theoUniqId + MYTHIC_DISPERSION);
    });

    it('correct range for medium UniqId', async () => {
      const theoUniqId = 30;
      const _genome = await generateMythicGenome(genomesWithTheoUniqId(theoUniqId));
      const genome = new Genome(_genome);
      expect(genome.getGene(MythicGenes.UniqId)).to.be.within(
        theoUniqId - MYTHIC_DISPERSION,
        theoUniqId + MYTHIC_DISPERSION,
      );
    });

    it('correct range for high UniqId', async () => {
      const theoUniqId = maxUniqId - 2;
      const _genome = await generateMythicGenome(genomesWithTheoUniqId(theoUniqId));
      const genome = new Genome(_genome);
      expect(genome.getGene(MythicGenes.UniqId)).to.be.within(
        theoUniqId - MYTHIC_DISPERSION,
        mythicAmount,
      );
    });

    it('correct range for max peruliarity', async () => {
      const theoUniqId = maxUniqId;
      const _genome = await generateMythicGenome(genomesWithTheoUniqId(theoUniqId));
      const genome = new Genome(_genome);
      expect(genome.getGene(MythicGenes.UniqId)).to.be.within(
        theoUniqId - MYTHIC_DISPERSION,
        mythicAmount,
      );
    });
  });

  describe('mintPack', () => {
    it('duckies are paid for mint');

    it('correct amount of tokens is minted');

    it('revert on amount == 0');

    it('revert on amount > MAX_PACK_SIZE');

    it('event is emitted');
  });
});
