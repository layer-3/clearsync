import { assert, expect } from 'chai';
import { ethers } from 'hardhat';

import {
  Collections,
  DucklingGenes,
  FLOCK_SIZE,
  GeneDistrTypes,
  Rarities,
  ZombeakGenes,
  collectionGeneIdx,
  collectionsGeneValuesNum,
  raritiesNum,
  rarityGeneIdx,
} from './config';
import {
  GenerateAndMintGenomesFunctT,
  MintToFuncT,
  extractMintedTokenId,
  randomGenome,
  randomGenomes,
  setupGenerateAndMintGenomes,
  setupMintTo,
} from './helpers';
import { setup } from './setup';
import { Genome } from './genome';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type {
  DucklingsV1,
  DuckyFamilyV1,
  TESTDuckyFamilyV1,
  YellowToken,
} from '../../../../typechain-types';

describe('DuckyFamilyV1 melding', () => {
  let Someone: SignerWithAddress;
  let GenomeSetter: SignerWithAddress;

  let Duckies: YellowToken;
  let Ducklings: DucklingsV1;
  let Game: TESTDuckyFamilyV1;

  let GameAsSomeone: DuckyFamilyV1;

  let mintTo: MintToFuncT;
  let generateAndMintGenomes: GenerateAndMintGenomesFunctT;

  const isCollectionMutating = async (rarity: Rarities): Promise<boolean> => {
    const tx = await Game.isCollectionMutating(rarity);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'BoolReturned');
    return event?.args?.returnedBool as boolean;
  };

  const meldGenomes = async (genomes: bigint[]): Promise<bigint> => {
    const tx = await Game.meldGenomes(genomes);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  const meldGenes = async (
    genomes: bigint[],
    gene: number,
    maxGeneValue: number,
    geneDistrType: GeneDistrTypes,
  ): Promise<number> => {
    const tx = await Game.meldGenes(genomes, gene, maxGeneValue, geneDistrType);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GeneReturned');
    // gene is already a number
    return event?.args?.gene as number;
  };

  beforeEach(async () => {
    ({ Someone, GenomeSetter, Duckies, Ducklings, Game, GameAsSomeone } = await setup());

    mintTo = setupMintTo(Ducklings.connect(GenomeSetter));
    generateAndMintGenomes = setupGenerateAndMintGenomes(mintTo, Someone.address);
  });

  describe('requireGenomesSatisfyMelding', () => {
    it('success on Common Duckling', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        amount: FLOCK_SIZE,
        [DucklingGenes.Rarity]: Rarities.Common,
        [DucklingGenes.Color]: 0,
      });

      await Game.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Rare Duckling', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        amount: FLOCK_SIZE,
        [DucklingGenes.Rarity]: Rarities.Rare,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 0,
      });

      await Game.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Epic Duckling', async () => {
      const genomes = randomGenomes(Collections.Duckling, {
        [DucklingGenes.Rarity]: Rarities.Epic,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 0,
      });

      await Game.requireGenomesSatisfyMelding(genomes);
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

      await Game.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Common Zombeak', async () => {
      const genomes = randomGenomes(Collections.Zombeak, {
        amount: FLOCK_SIZE,
        [ZombeakGenes.Rarity]: Rarities.Common,
        [ZombeakGenes.Color]: 0,
      });

      await Game.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Rare Zombeak', async () => {
      const genomes = randomGenomes(Collections.Zombeak, {
        amount: FLOCK_SIZE,
        [ZombeakGenes.Rarity]: Rarities.Rare,
        [ZombeakGenes.Color]: 0,
        [ZombeakGenes.Family]: 0,
      });

      await Game.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Epic Zombeak', async () => {
      const genomes = randomGenomes(Collections.Zombeak, {
        amount: FLOCK_SIZE,
        [ZombeakGenes.Rarity]: Rarities.Epic,
        [ZombeakGenes.Color]: 0,
        [ZombeakGenes.Family]: 0,
      });

      await Game.requireGenomesSatisfyMelding(genomes);
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

      await expect(Game.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(Game, 'IncorrectGenomesForMelding')
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

      await expect(Game.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(Game, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });

    it('revert on Mythic', async () => {
      const genomes = randomGenomes(Collections.Mythic);

      await expect(Game.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(Game, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });

    it('revert on Legendary Zombeak', async () => {
      const genomes = randomGenomes(Collections.Zombeak, {
        amount: FLOCK_SIZE,
        [ZombeakGenes.Rarity]: Rarities.Legendary,
        [ZombeakGenes.Color]: 0,
        [ZombeakGenes.Family]: 0,
      });

      await expect(Game.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(Game, 'IncorrectGenomesForMelding')
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

      await expect(Game.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(Game, 'IncorrectGenomesForMelding')
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

      await expect(Game.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(Game, 'IncorrectGenomesForMelding')
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

      await expect(Game.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(Game, 'IncorrectGenomesForMelding')
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

      await expect(Game.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(Game, 'IncorrectGenomesForMelding')
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

      await expect(Game.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(Game, 'IncorrectGenomesForMelding')
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

      await expect(Game.requireGenomesSatisfyMelding(genomes))
        .to.be.revertedWithCustomError(Game, 'IncorrectGenomesForMelding')
        .withArgs(genomes);
    });
  });

  describe('isCollectionMutating', () => {
    describe('all collections do mutate', () => {
      beforeEach(async () => {
        await Game.setCollectionMutationChances([1000, 1000, 1000, 1000]); // per mil chances
      });

      it('Common can mutate', async () => {
        expect(await isCollectionMutating(Rarities.Common)).to.be.true;
      });

      it('Rare can mutate', async () => {
        expect(await isCollectionMutating(Rarities.Rare)).to.be.true;
      });

      it('Epic can mutate', async () => {
        expect(await isCollectionMutating(Rarities.Epic)).to.be.true;
      });

      it('Legendary can mutate', async () => {
        expect(await isCollectionMutating(Rarities.Legendary)).to.be.true;
      });
    });

    describe('all collections do not mutate', () => {
      beforeEach(async () => {
        await Game.setCollectionMutationChances([0, 0, 0, 0]);
      });

      it('Common can not mutate', async () => {
        expect(await isCollectionMutating(Rarities.Common)).to.be.false;
      });

      it('Rare can not mutate', async () => {
        expect(await isCollectionMutating(Rarities.Rare)).to.be.false;
      });

      it('Epic can not mutate', async () => {
        expect(await isCollectionMutating(Rarities.Epic)).to.be.false;
      });

      it('Legendary can not mutate', async () => {
        expect(await isCollectionMutating(Rarities.Legendary)).to.be.false;
      });
    });

    it('revert when rarity is out of bounds', async () => {
      await expect(Game.isCollectionMutating(raritiesNum)).to.be.reverted;
    });
  });

  describe('meldGenomes', () => {
    const collectionAndRarity = (genome_: bigint): [Collections, Rarities] => {
      const genome = new Genome(genome_);
      return [genome.getGene(collectionGeneIdx), genome.getGene(rarityGeneIdx)];
    };

    describe('defaulted values are randomized after melding Ducklings', () => {
      beforeEach(async () => {
        await Game.setGeneMutationChance(0);
      });

      it('Body is randomized after melding Common Ducklings', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          [DucklingGenes.Rarity]: Rarities.Common,
          [DucklingGenes.Color]: 0,
          // set default Body
          [DucklingGenes.Body]: 0,
        });

        const _genome = await meldGenomes(genomes);
        const genome = new Genome(_genome);
        expect(genome.getGene(DucklingGenes.Body)).to.not.equal(0);
      });

      it('Head is randomized after melding Rare Ducklings', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          [DucklingGenes.Rarity]: Rarities.Rare,
          [DucklingGenes.Color]: 0,
          [DucklingGenes.Family]: 0,
          // set default Head
          [DucklingGenes.Head]: 0,
        });

        const _genome = await meldGenomes(genomes);
        const genome = new Genome(_genome);
        expect(genome.getGene(DucklingGenes.Head)).to.not.equal(0);
      });

      it('Body is not randomized after melding Rare Ducklings', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          [DucklingGenes.Rarity]: Rarities.Rare,
          [DucklingGenes.Color]: 0,
          [DucklingGenes.Family]: 0,
          // set default Body
          [DucklingGenes.Body]: 0,
        });

        const _genome = await meldGenomes(genomes);
        const genome = new Genome(_genome);
        expect(genome.getGene(DucklingGenes.Body)).to.equal(0);
      });

      it('Head is not randomized after melding Common Ducklings', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          [DucklingGenes.Rarity]: Rarities.Common,
          [DucklingGenes.Color]: 0,
          [DucklingGenes.Family]: 0,
          // set default Head
          [DucklingGenes.Head]: 0,
        });

        const _genome = await meldGenomes(genomes);
        const genome = new Genome(_genome);
        expect(genome.getGene(DucklingGenes.Head)).to.equal(0);
      });

      it('Body is not randomized after melding Common Zombeaks', async () => {
        const genomes = randomGenomes(Collections.Zombeak, {
          amount: FLOCK_SIZE,
          [ZombeakGenes.Rarity]: Rarities.Common,
          [ZombeakGenes.Color]: 0,
          // set default Body
          [ZombeakGenes.Body]: 0,
        });

        const _genome = await meldGenomes(genomes);
        const genome = new Genome(_genome);
        expect(genome.getGene(ZombeakGenes.Body)).to.equal(0);
      });

      it('Head is not randomized after melding Rare Zombeaks', async () => {
        const genomes = randomGenomes(Collections.Zombeak, {
          amount: FLOCK_SIZE,
          [ZombeakGenes.Rarity]: Rarities.Rare,
          [ZombeakGenes.Color]: 0,
          [ZombeakGenes.Family]: 0,
          // set default Head
          [ZombeakGenes.Head]: 0,
        });

        const _genome = await meldGenomes(genomes);
        const genome = new Genome(_genome);
        expect(genome.getGene(ZombeakGenes.Head)).to.equal(0);
      });
    });

    describe('same collection, increased rarity', () => {
      beforeEach(async () => {
        // disable mutations
        await Game.setCollectionMutationChances([0, 0, 0, 0]);
      });

      it('when melding Common Ducklings', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          [DucklingGenes.Rarity]: Rarities.Common,
          [DucklingGenes.Color]: 0,
        });

        const [collection, rarity] = collectionAndRarity(await meldGenomes(genomes));

        expect(collection).to.equal(Collections.Duckling);
        expect(rarity).to.equal(Rarities.Rare);
      });

      it('when melding Rare Ducklings', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          [DucklingGenes.Rarity]: Rarities.Rare,
          [DucklingGenes.Color]: 0,
          [DucklingGenes.Family]: 0,
        });

        const [collection, rarity] = collectionAndRarity(await meldGenomes(genomes));

        expect(collection).to.equal(Collections.Duckling);
        expect(rarity).to.equal(Rarities.Epic);
      });

      it('when melding Epic Ducklings', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          [DucklingGenes.Rarity]: Rarities.Epic,
          [DucklingGenes.Color]: 0,
          [DucklingGenes.Family]: 0,
        });

        const [collection, rarity] = collectionAndRarity(await meldGenomes(genomes));

        expect(collection).to.equal(Collections.Duckling);
        expect(rarity).to.equal(Rarities.Legendary);
      });

      it('Mythic when melding Legendary Ducklings', async () => {
        const genomes = [];

        for (let i = 0; i < FLOCK_SIZE; i++) {
          genomes.push(
            randomGenome(Collections.Duckling, {
              [ZombeakGenes.Rarity]: Rarities.Legendary,
              // different colors
              [ZombeakGenes.Color]: 0,
              [ZombeakGenes.Family]: i,
            }),
          );
        }

        const [collection] = collectionAndRarity(await meldGenomes(genomes));
        expect(collection).to.equal(Collections.Mythic);
      });

      it('when melding Common Zombeaks', async () => {
        const genomes = randomGenomes(Collections.Zombeak, {
          amount: FLOCK_SIZE,
          [ZombeakGenes.Rarity]: Rarities.Common,
          [ZombeakGenes.Color]: 0,
        });

        const [collection, rarity] = collectionAndRarity(await meldGenomes(genomes));

        expect(collection).to.equal(Collections.Zombeak);
        expect(rarity).to.equal(Rarities.Rare);
      });

      it('when melding Rare Zombeaks', async () => {
        const genomes = randomGenomes(Collections.Zombeak, {
          amount: FLOCK_SIZE,
          [ZombeakGenes.Rarity]: Rarities.Rare,
          [ZombeakGenes.Color]: 0,
          [ZombeakGenes.Family]: 0,
        });

        const [collection, rarity] = collectionAndRarity(await meldGenomes(genomes));

        expect(collection).to.equal(Collections.Zombeak);
        expect(rarity).to.equal(Rarities.Epic);
      });

      it('when melding Epic Zombeaks', async () => {
        const genomes = randomGenomes(Collections.Zombeak, {
          amount: FLOCK_SIZE,
          [ZombeakGenes.Rarity]: Rarities.Epic,
          [ZombeakGenes.Color]: 0,
          [ZombeakGenes.Family]: 0,
        });

        const [collection, rarity] = collectionAndRarity(await meldGenomes(genomes));

        expect(collection).to.equal(Collections.Zombeak);
        expect(rarity).to.equal(Rarities.Legendary);
      });
    });

    describe('mutated into zombeak with the same rarity', () => {
      beforeEach(async () => {
        // 100% mutations
        await Game.setCollectionMutationChances([1000, 1000, 1000, 1000]);
      });

      it('when melding Common Ducklings', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          [DucklingGenes.Rarity]: Rarities.Common,
          [DucklingGenes.Color]: 0,
        });

        const [collection, rarity] = collectionAndRarity(await meldGenomes(genomes));

        expect(collection).to.equal(Collections.Zombeak);
        expect(rarity).to.equal(Rarities.Common);
      });

      it('when melding Rare Ducklings', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          [DucklingGenes.Rarity]: Rarities.Rare,
          [DucklingGenes.Color]: 0,
          [DucklingGenes.Family]: 0,
        });

        const [collection, rarity] = collectionAndRarity(await meldGenomes(genomes));

        expect(collection).to.equal(Collections.Zombeak);
        expect(rarity).to.equal(Rarities.Rare);
      });

      it('when melding Epic Ducklings', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          [DucklingGenes.Rarity]: Rarities.Epic,
          [DucklingGenes.Color]: 0,
          [DucklingGenes.Family]: 0,
        });

        const [collection, rarity] = collectionAndRarity(await meldGenomes(genomes));

        expect(collection).to.equal(Collections.Zombeak);
        expect(rarity).to.equal(Rarities.Epic);
      });

      it('when melding Legendary Ducklings', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          [DucklingGenes.Rarity]: Rarities.Legendary,
          [DucklingGenes.Color]: 0,
          [DucklingGenes.Family]: 0,
        });

        const [collection, rarity] = collectionAndRarity(await meldGenomes(genomes));

        expect(collection).to.equal(Collections.Zombeak);
        expect(rarity).to.equal(Rarities.Legendary);
      });
    });
  });

  describe('meldGenes', () => {
    const geneValuesNum = collectionsGeneValuesNum[Collections.Duckling];

    describe('when mutation is 100%', () => {
      beforeEach(async () => {
        await Game.setGeneMutationChance(1000);
      });

      it('uneven gene gets mutated', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: 5,
          // uneven gene
          [DucklingGenes.Head]: 0,
        });

        const meldedGene = await meldGenes(
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

        const meldedGene = await meldGenes(
          genomes,
          DucklingGenes.Head,
          geneValuesNum[DucklingGenes.Head],
          GeneDistrTypes.Uneven,
        );
        expect(meldedGene).to.equal(maxGeneValue + 1);
      });

      it('even gene does not mutate', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: 5,
          // even gene
          [DucklingGenes.FirstName]: 0,
        });

        const meldedGene = await meldGenes(
          genomes,
          DucklingGenes.FirstName,
          geneValuesNum[DucklingGenes.FirstName],
          GeneDistrTypes.Even,
        );
        expect(meldedGene).to.equal(0);
      });
    });

    describe('when mutation is 0%', () => {
      beforeEach(async () => {
        await Game.setGeneMutationChance(0);
      });

      it('value for uneven gene is got from parents', async () => {
        const genomes = [];
        for (let i = 0; i < FLOCK_SIZE; i++) {
          const genome = randomGenome(Collections.Duckling, {
            [DucklingGenes.Head]: i,
          });
          genomes.push(genome);
        }

        const meldedGene = await meldGenes(
          genomes,
          DucklingGenes.Head,
          geneValuesNum[DucklingGenes.Head],
          GeneDistrTypes.Uneven,
        );
        expect(meldedGene).to.be.within(0, FLOCK_SIZE - 1);
      });

      it('value for even gene is got from parents', async () => {
        const genomes = [];
        for (let i = 0; i < FLOCK_SIZE; i++) {
          const genome = randomGenome(Collections.Duckling, {
            [DucklingGenes.FirstName]: i,
          });
          genomes.push(genome);
        }

        const meldedGene = await meldGenes(
          genomes,
          DucklingGenes.FirstName,
          geneValuesNum[DucklingGenes.FirstName],
          GeneDistrTypes.Even,
        );
        expect(meldedGene).to.be.within(0, FLOCK_SIZE - 1);
      });

      it('uneven gene: value does not change if all parents have the same value', async () => {
        const genomes = randomGenomes(Collections.Duckling, {
          amount: FLOCK_SIZE,
          // uneven gene
          [DucklingGenes.Head]: 0,
        });

        const meldedGene = await meldGenes(
          genomes,
          DucklingGenes.Head,
          geneValuesNum[DucklingGenes.Head],
          GeneDistrTypes.Uneven,
        );
        expect(meldedGene).to.equal(0);
      });
    });
  });

  // eslint-disable-next-line sonarjs/cognitive-complexity
  describe('meldFlock', () => {
    beforeEach(async () => {
      await Duckies.connect(Someone).increaseAllowance(Game.address, 100_000_000_000);
    });

    describe('Duckling', () => {
      it('success on melding Common Ducklings', async () => {
        const { tokenIds } = await generateAndMintGenomes(Collections.Duckling, {
          [DucklingGenes.Rarity]: Rarities.Common,
          [DucklingGenes.Color]: 0,
        });

        await GameAsSomeone.meldFlock(tokenIds);
      });

      it('success on melding Rare Ducklings', async () => {
        const { tokenIds } = await generateAndMintGenomes(Collections.Duckling, {
          [DucklingGenes.Rarity]: Rarities.Rare,
          [DucklingGenes.Color]: 1,
          [DucklingGenes.Family]: 1,
        });

        try {
          await GameAsSomeone.meldFlock(tokenIds);
          assert(true);
        } catch {
          assert(false);
        }
      });

      it('success on melding Epic Ducklings', async () => {
        const { tokenIds } = await generateAndMintGenomes(Collections.Duckling, {
          [DucklingGenes.Rarity]: Rarities.Epic,
          [DucklingGenes.Color]: 1,
          [DucklingGenes.Family]: 1,
        });

        try {
          await GameAsSomeone.meldFlock(tokenIds);
          assert(true);
        } catch {
          assert(false);
        }
      });
    });

    describe('Zombeak', () => {
      it('success on melding Common Zombeaks', async () => {
        const { tokenIds } = await generateAndMintGenomes(Collections.Zombeak, {
          [ZombeakGenes.Rarity]: Rarities.Common,
          [ZombeakGenes.Color]: 0,
        });

        try {
          await GameAsSomeone.meldFlock(tokenIds);
          assert(true);
        } catch {
          assert(false);
        }
      });

      it('success on melding Rare Zombeaks', async () => {
        const { tokenIds } = await generateAndMintGenomes(Collections.Zombeak, {
          [ZombeakGenes.Rarity]: Rarities.Rare,
          [ZombeakGenes.Color]: 0,
          [ZombeakGenes.Family]: 0,
        });

        try {
          await GameAsSomeone.meldFlock(tokenIds);
          assert(true);
        } catch {
          assert(false);
        }
      });

      it('success on melding Epic Zombeaks', async () => {
        const { tokenIds } = await generateAndMintGenomes(Collections.Zombeak, {
          [ZombeakGenes.Rarity]: Rarities.Epic,
          [ZombeakGenes.Color]: 0,
          [ZombeakGenes.Family]: 0,
        });

        try {
          await GameAsSomeone.meldFlock(tokenIds);
          assert(true);
        } catch {
          assert(false);
        }
      });
    });

    describe('Mythic', () => {
      it('success on melding Legendary Ducklings into Mythic', async () => {
        const meldingTokenIds: number[] = [];
        await (async () => {
          for (let i = 0; i < 5; i++) {
            const tokenIdAndGenome = await extractMintedTokenId(
              await mintTo(
                Someone.address,
                randomGenome(Collections.Duckling, {
                  [DucklingGenes.Rarity]: Rarities.Legendary,
                  [DucklingGenes.Color]: 1,
                  [DucklingGenes.Family]: i,
                }),
              ),
            );
            meldingTokenIds.push(tokenIdAndGenome.tokenId);
          }
        })();

        try {
          await GameAsSomeone.meldFlock(meldingTokenIds);
          assert(true);
        } catch {
          assert(false);
        }
      });
    });

    it('event is emitted', async () => {
      const { tokenIds } = await generateAndMintGenomes(Collections.Duckling, {
        [DucklingGenes.Rarity]: Rarities.Common,
        [DucklingGenes.Color]: 0,
      });

      const meldedTokenId = FLOCK_SIZE;
      const chainId = await ethers.provider.getNetwork().then((network) => network.chainId);

      await expect(GameAsSomeone.meldFlock(tokenIds))
        .to.emit(Game, 'Melded')
        .withArgs(Someone.address, tokenIds, meldedTokenId, chainId);
    });
  });
});
