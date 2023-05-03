import { assert, expect } from 'chai';
import { ethers } from 'hardhat';

import {
  Collections,
  DucklingGenes,
  FLOCK_SIZE,
  Rarities,
  ZombeakGenes,
  collectionGeneIdx,
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

  const meldGenomes = async (genomes: bigint[]): Promise<bigint> => {
    const tx = await Game.meldGenomes(genomes);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  beforeEach(async () => {
    ({ Someone, GenomeSetter, Duckies, Ducklings, Game, GameAsSomeone } = await setup());

    mintTo = setupMintTo(Ducklings.connect(GenomeSetter));
    generateAndMintGenomes = setupGenerateAndMintGenomes(mintTo, Someone.address);
  });

  describe('meldGenomes', () => {
    const collectionAndRarity = (genome_: bigint): [Collections, Rarities] => {
      const genome = new Genome(genome_);
      return [genome.getGene(collectionGeneIdx), genome.getGene(rarityGeneIdx)];
    };

    describe('defaulted values are randomized after melding Ducklings', () => {
      beforeEach(async () => {
        await Game.setCollectionMutationChances([0, 0, 0, 0]);
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

        // TODO: remove this when we have a better way to test randomness
        await Promise.all(
          Array.from({ length: 100 }, async () => {
            const _genome = await meldGenomes(genomes);
            const genome = new Genome(_genome);
            expect(genome.getGene(DucklingGenes.Body)).greaterThan(0);
            // 10 - number of all bodies
            expect(genome.getGene(DucklingGenes.Body)).lessThanOrEqual(10);
          }),
        );
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
