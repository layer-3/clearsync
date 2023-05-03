import { expect } from 'chai';

import {
  Collections,
  DucklingGenes,
  FLOCK_SIZE,
  GeneDistrTypes,
  Rarities,
  ZombeakGenes,
  collectionsGeneValuesNum,
  geneInheritanceChances,
  geneMutationChance,
  raritiesNum,
} from './config';
import { randomGenome, randomGenomes } from './helpers';
import { setup } from './setup';

import type { DuckyGenomeTestConsumer } from '../../../../typechain-types';

const BIT_SLICE = '0xaabbcc';

describe('DuckyGenome melding', () => {
  let DuckyGenome: DuckyGenomeTestConsumer;

  const isCollectionMutating = async (
    rarity: Rarities,
    mutationChances: number[],
    bitSlice: string = BIT_SLICE,
  ): Promise<boolean> => {
    const tx = await DuckyGenome.isCollectionMutating(rarity, mutationChances, bitSlice);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'BoolReturned');
    return event?.args?.returnedBool as boolean;
  };

  const meldGenes = async (
    genomes: bigint[],
    gene: number,
    maxGeneValue: number,
    geneDistrType: GeneDistrTypes,
    mutationChance: number[] = geneMutationChance,
    inheritanceChances: number[] = geneInheritanceChances,
    bitSlice: string = BIT_SLICE,
  ): Promise<number> => {
    const tx = await DuckyGenome.meldGenes(
      genomes,
      gene,
      maxGeneValue,
      geneDistrType,
      mutationChance,
      inheritanceChances,
      bitSlice,
    );
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GeneReturned');
    // gene is already a number
    return event?.args?.gene as number;
  };

  beforeEach(async () => {
    ({ DuckyGenome } = await setup());
  });

  describe('requireGenomesSatisfyMelding', () => {
    it('success on Common Duckling', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        amount: FLOCK_SIZE,
        [DucklingGenes.Rarity]: Rarities.Common,
        [DucklingGenes.Color]: 0,
      });

      await DuckyGenome.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Rare Duckling', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        amount: FLOCK_SIZE,
        [DucklingGenes.Rarity]: Rarities.Rare,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 0,
      });

      await DuckyGenome.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Epic Duckling', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        [DucklingGenes.Rarity]: Rarities.Epic,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 0,
      });

      await DuckyGenome.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Legendary Duckling', async () => {
      const genomes = [];

      for (let i = 0; i < FLOCK_SIZE; i++) {
        genomes.push(
          randomGenome(Collections.Duckling, {
            [DucklingGenes.Rarity]: Rarities.Legendary,
            [DucklingGenes.Color]: 0,
            // all families
            [DucklingGenes.Family]: i,
          }),
        );
      }

      await DuckyGenome.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Common Zombeak', async () => {
      const genomes = randomGenomes(Collections.Zombeak, {
        amount: FLOCK_SIZE,
        [ZombeakGenes.Rarity]: Rarities.Common,
        [ZombeakGenes.Color]: 0,
      });

      await DuckyGenome.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Rare Zombeak', async () => {
      const genomes = randomGenomes(Collections.Zombeak, {
        amount: FLOCK_SIZE,
        [ZombeakGenes.Rarity]: Rarities.Rare,
        [ZombeakGenes.Color]: 0,
        [ZombeakGenes.Family]: 0,
      });

      await DuckyGenome.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Epic Zombeak', async () => {
      const genomes = randomGenomes(Collections.Zombeak, {
        amount: FLOCK_SIZE,
        [ZombeakGenes.Rarity]: Rarities.Epic,
        [ZombeakGenes.Color]: 0,
        [ZombeakGenes.Family]: 0,
      });

      await DuckyGenome.requireGenomesSatisfyMelding(genomes);
    });

    it('revert on different collections', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        amount: FLOCK_SIZE - 1,
        [DucklingGenes.Rarity]: Rarities.Common,
        [DucklingGenes.Color]: 0,
      });

      genomes.push(
        randomGenome(Collections.Zombeak, {
          [ZombeakGenes.Rarity]: Rarities.Common,
          [ZombeakGenes.Color]: 0,
        }),
      );

      await expect(DuckyGenome.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(DuckyGenome, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });

    it('revert on different rarities', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        amount: FLOCK_SIZE - 1,
        [DucklingGenes.Rarity]: Rarities.Common,
        [DucklingGenes.Color]: 0,
      });

      genomes.push(
        randomGenome(Collections.Duckling, {
          [DucklingGenes.Rarity]: Rarities.Rare,
          [DucklingGenes.Color]: 0,
        }),
      );

      await expect(DuckyGenome.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(DuckyGenome, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });

    it('revert on Mythic', async () => {
      const genomes = randomGenomes(Collections.Mythic);

      await expect(DuckyGenome.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(DuckyGenome, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });

    it('revert on Legendary Zombeak', async () => {
      const genomes = randomGenomes(Collections.Zombeak, {
        amount: FLOCK_SIZE,
        [ZombeakGenes.Rarity]: Rarities.Legendary,
        [ZombeakGenes.Color]: 0,
        [ZombeakGenes.Family]: 0,
      });

      await expect(DuckyGenome.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(DuckyGenome, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });

    it('revert on Legendaries having different color', async () => {
      const genomes = [];

      for (let i = 0; i < FLOCK_SIZE; i++) {
        genomes.push(
          randomGenome(Collections.Duckling, {
            [ZombeakGenes.Rarity]: Rarities.Legendary,
            // different colors
            [ZombeakGenes.Color]: i,
            [ZombeakGenes.Family]: i,
          }),
        );
      }

      await expect(DuckyGenome.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(DuckyGenome, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });

    it('revert on Legendaries having repeated families', async () => {
      const genomes = [];

      for (let i = 0; i < FLOCK_SIZE; i++) {
        genomes.push(
          randomGenome(Collections.Duckling, {
            [ZombeakGenes.Rarity]: Rarities.Legendary,
            [ZombeakGenes.Color]: 0,
            // repeated families
            [ZombeakGenes.Family]: 0,
          }),
        );
      }

      await expect(DuckyGenome.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(DuckyGenome, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });

    it('revert on Epic having different color', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        amount: FLOCK_SIZE,
        [DucklingGenes.Rarity]: Rarities.Epic,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 0,
      });

      // different color
      genomes[0] = randomGenome(Collections.Duckling, {
        [DucklingGenes.Rarity]: Rarities.Epic,
        [DucklingGenes.Color]: 1,
        [DucklingGenes.Family]: 0,
      });

      await expect(DuckyGenome.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(DuckyGenome, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });

    it('revert on Epic having different family', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        amount: FLOCK_SIZE,
        [DucklingGenes.Rarity]: Rarities.Epic,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 0,
      });

      // different family
      genomes[0] = randomGenome(Collections.Duckling, {
        [DucklingGenes.Rarity]: Rarities.Epic,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 1,
      });

      await expect(DuckyGenome.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(DuckyGenome, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });

    it('revert on Rare having different color', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        amount: FLOCK_SIZE,
        [DucklingGenes.Rarity]: Rarities.Rare,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 0,
      });

      // different color
      genomes[0] = randomGenome(Collections.Duckling, {
        [DucklingGenes.Rarity]: Rarities.Rare,
        [DucklingGenes.Color]: 1,
        [DucklingGenes.Family]: 0,
      });

      await expect(DuckyGenome.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(DuckyGenome, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });

    it('revert on Rare having different family', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        amount: FLOCK_SIZE,
        [DucklingGenes.Rarity]: Rarities.Rare,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 0,
      });

      // different family
      genomes[0] = randomGenome(Collections.Duckling, {
        [DucklingGenes.Rarity]: Rarities.Rare,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 1,
      });

      await expect(DuckyGenome.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(DuckyGenome, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });
  });

  describe('isCollectionMutating', () => {
    describe('all collections do mutate', () => {
      const isCollectionMutating_always = async (rarity: Rarities): Promise<boolean> => {
        return isCollectionMutating(rarity, [1000, 1000, 1000, 1000]);
      };

      it('Common can mutate', async () => {
        expect(await isCollectionMutating_always(Rarities.Common)).to.be.true;
      });

      it('Rare can mutate', async () => {
        expect(await isCollectionMutating_always(Rarities.Rare)).to.be.true;
      });

      it('Epic can mutate', async () => {
        expect(await isCollectionMutating_always(Rarities.Epic)).to.be.true;
      });

      it('Legendary can mutate', async () => {
        expect(await isCollectionMutating_always(Rarities.Legendary)).to.be.true;
      });
    });

    describe('all collections do not mutate', () => {
      const isCollectionMutating_never = async (rarity: Rarities): Promise<boolean> => {
        return isCollectionMutating(rarity, [0, 0, 0, 0]);
      };

      it('Common can not mutate', async () => {
        expect(await isCollectionMutating_never(Rarities.Common)).to.be.false;
      });

      it('Rare can not mutate', async () => {
        expect(await isCollectionMutating_never(Rarities.Rare)).to.be.false;
      });

      it('Epic can not mutate', async () => {
        expect(await isCollectionMutating_never(Rarities.Epic)).to.be.false;
      });

      it('Legendary can not mutate', async () => {
        expect(await isCollectionMutating_never(Rarities.Legendary)).to.be.false;
      });
    });

    it('revert when rarity is out of bounds', async () => {
      await expect(DuckyGenome.isCollectionMutating(raritiesNum, [1, 1, 1, 1], BIT_SLICE)).to.be
        .reverted;
    });
  });

  // eslint-disable-next-line sonarjs/cognitive-complexity
  describe('meldGenes', () => {
    const geneValuesNum = collectionsGeneValuesNum[Collections.Duckling];

    describe('when mutation is 100%', () => {
      const meldGenes_alwaysMutate = async (
        genomes: bigint[],
        gene: number,
        maxGeneValue: number,
        geneDistrType: GeneDistrTypes,
        bitSlice: string = BIT_SLICE,
      ): Promise<number> => {
        return meldGenes(
          genomes,
          gene,
          maxGeneValue,
          geneDistrType,
          [0, 1000],
          geneInheritanceChances,
          bitSlice,
        );
      };

      it('uneven gene gets mutated', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          // uneven gene
          [DucklingGenes.Head]: 0,
        });

        const meldedGene = await meldGenes_alwaysMutate(
          genomes,
          DucklingGenes.Head,
          geneValuesNum[DucklingGenes.Head],
          GeneDistrTypes.Uneven,
        );
        expect(meldedGene).not.to.equal(0);
      });

      it('when mutated, the value is max + 1', async () => {
        const genomes = [];
        for (let i = 0; i < FLOCK_SIZE; i++) {
          const genome = randomGenome(Collections.Duckling, {
            [DucklingGenes.Head]: i,
          });
          genomes.push(genome);
        }

        const maxGeneValue = FLOCK_SIZE - 1;

        const meldedGene = await meldGenes_alwaysMutate(
          genomes,
          DucklingGenes.Head,
          geneValuesNum[DucklingGenes.Head],
          GeneDistrTypes.Uneven,
        );
        expect(meldedGene).to.equal(maxGeneValue + 1);
      });

      it('even gene does not mutate', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          // even gene
          [DucklingGenes.FirstName]: 0,
        });

        const meldedGene = await meldGenes_alwaysMutate(
          genomes,
          DucklingGenes.FirstName,
          geneValuesNum[DucklingGenes.FirstName],
          GeneDistrTypes.Even,
        );
        expect(meldedGene).to.equal(0);
      });
    });

    describe('when mutation is 0%', () => {
      const meldGenes_neverMutate = async (
        genomes: bigint[],
        gene: number,
        maxGeneValue: number,
        geneDistrType: GeneDistrTypes,
        inheritanceChances: number[] = geneInheritanceChances,
        bitSlice: string = BIT_SLICE,
      ): Promise<number> => {
        return meldGenes(
          genomes,
          gene,
          maxGeneValue,
          geneDistrType,
          [1000, 0],
          inheritanceChances,
          bitSlice,
        );
      };

      it('value for uneven gene is got from parents', async () => {
        const geneValue = 3;
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          // uneven gene
          [DucklingGenes.Head]: geneValue,
        });

        const meldedGene = await meldGenes_neverMutate(
          genomes,
          DucklingGenes.Head,
          geneValuesNum[DucklingGenes.Head],
          GeneDistrTypes.Uneven,
        );
        expect(meldedGene).to.equal(geneValue);
      });

      it('value for even gene is got from parents', async () => {
        const geneValue = 3;

        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          // even gene
          [DucklingGenes.FirstName]: geneValue,
        });

        const meldedGene = await meldGenes_neverMutate(
          genomes,
          DucklingGenes.FirstName,
          geneValuesNum[DucklingGenes.FirstName],
          GeneDistrTypes.Even,
        );
        expect(meldedGene).to.equal(geneValue);
      });

      it('can inherit from 1st parent', async () => {
        const genomes = [];
        for (let i = 0; i < FLOCK_SIZE; i++) {
          const genome = randomGenome(Collections.Duckling, {
            [DucklingGenes.Head]: i,
          });
          genomes.push(genome);
        }

        const meldedGene = await meldGenes_neverMutate(
          genomes,
          DucklingGenes.Head,
          geneValuesNum[DucklingGenes.Head],
          GeneDistrTypes.Uneven,
          [1000, 0, 0, 0, 0],
        );

        expect(meldedGene).to.equal(0);
      });

      it('can inherit from 2nd parent', async () => {
        const genomes = [];
        for (let i = 0; i < FLOCK_SIZE; i++) {
          const genome = randomGenome(Collections.Duckling, {
            [DucklingGenes.Head]: i,
          });
          genomes.push(genome);
        }

        const meldedGene = await meldGenes_neverMutate(
          genomes,
          DucklingGenes.Head,
          geneValuesNum[DucklingGenes.Head],
          GeneDistrTypes.Uneven,
          [0, 1000, 0, 0, 0],
        );

        expect(meldedGene).to.equal(1);
      });

      it('can inherit from 3rd parent', async () => {
        const genomes = [];
        for (let i = 0; i < FLOCK_SIZE; i++) {
          const genome = randomGenome(Collections.Duckling, {
            [DucklingGenes.Head]: i,
          });
          genomes.push(genome);
        }

        const meldedGene = await meldGenes_neverMutate(
          genomes,
          DucklingGenes.Head,
          geneValuesNum[DucklingGenes.Head],
          GeneDistrTypes.Uneven,
          [0, 0, 1000, 0, 0],
        );

        expect(meldedGene).to.equal(2);
      });

      it('can inherit from 4th parent', async () => {
        const genomes = [];
        for (let i = 0; i < FLOCK_SIZE; i++) {
          const genome = randomGenome(Collections.Duckling, {
            [DucklingGenes.Head]: i,
          });
          genomes.push(genome);
        }

        const meldedGene = await meldGenes_neverMutate(
          genomes,
          DucklingGenes.Head,
          geneValuesNum[DucklingGenes.Head],
          GeneDistrTypes.Uneven,
          [0, 0, 0, 1000, 0],
        );

        expect(meldedGene).to.equal(3);
      });

      it('can inherit from 5th parent', async () => {
        const genomes = [];
        for (let i = 0; i < FLOCK_SIZE; i++) {
          const genome = randomGenome(Collections.Duckling, {
            [DucklingGenes.Head]: i,
          });
          genomes.push(genome);
        }

        const meldedGene = await meldGenes_neverMutate(
          genomes,
          DucklingGenes.Head,
          geneValuesNum[DucklingGenes.Head],
          GeneDistrTypes.Uneven,
          [0, 0, 0, 0, 1000],
        );

        expect(meldedGene).to.equal(4);
      });
    });
  });
});
