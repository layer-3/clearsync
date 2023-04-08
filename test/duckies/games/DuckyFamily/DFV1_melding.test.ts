import { assert, expect } from 'chai';

import {
  Collections,
  DucklingGenes,
  FLOCK_SIZE,
  Rarities,
  ZombeakGenes,
  collectionGeneIdx,
  raritiesNum,
  rarityGeneIdx,
} from './config';
import { RandomGenomeConfig, randomGenome, randomGenomes } from './helpers';
import { setup } from './setup';
import { Genome } from './genome';

import type { ContractTransaction } from 'ethers';
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

  const mintTo = async (
    to: string,
    genome: bigint,
    isTransferable?: boolean,
  ): Promise<ContractTransaction> => {
    return await Ducklings.connect(GenomeSetter).mintTo(to, genome, isTransferable ?? true);
  };

  interface TokenIdAndGenome {
    tokenId: number;
    genome: bigint;
  }

  const extractMintedTokenId = async (tx: ContractTransaction): Promise<TokenIdAndGenome> => {
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'Minted');
    const tokenId = event?.args?.tokenId.toNumber() as number;
    const genome = event?.args?.genome.toBigInt() as bigint;
    return { tokenId, genome };
  };

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

  beforeEach(async () => {
    ({ Someone, GenomeSetter, Duckies, Ducklings, Game } = await setup());
  });

  interface TokenIdsAndGenomes {
    tokenIds: number[];
    genomes: bigint[];
  }

  const generateAndMintGenomes = async (
    collection: Collections,
    config?: RandomGenomeConfig & { amount?: number },
  ): Promise<TokenIdsAndGenomes> => {
    const tokenIds: number[] = [];
    const genomes: bigint[] = [];

    const amount = config?.amount ?? 5;

    for (let i = 0; i < amount; i++) {
      const tokenIdAndGenome = await extractMintedTokenId(
        await mintTo(Someone.address, randomGenome(collection, config)),
      );
      tokenIds.push(tokenIdAndGenome.tokenId);
      genomes.push(tokenIdAndGenome.genome);
    }

    return { tokenIds, genomes };
  };

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
    it('can meld');

    it('uneven gene can mutate');

    it('even gene can not mutate');

    it('can inherit from all parents');
  });

  // eslint-disable-next-line sonarjs/cognitive-complexity
  describe('meldFlock', () => {
    describe.skip('Duckling', () => {
      it('success on melding Common Ducklings', async () => {
        const tokenIdsAndGenomes = await generateAndMintGenomes(Collections.Duckling, {
          [DucklingGenes.Rarity]: Rarities.Common,
          [DucklingGenes.Color]: 0,
        });

        try {
          await Duckies.connect(Someone).increaseAllowance(Game.address, 10_000_000_000);
          await GameAsSomeone.meldFlock(tokenIdsAndGenomes.tokenIds);
          assert(true);
        } catch {
          assert(false);
        }
      });

      it('success on melding Rare Ducklings', async () => {
        const tokenIdsAndGenomes = await generateAndMintGenomes(Collections.Duckling, {
          [DucklingGenes.Rarity]: Rarities.Rare,
          [DucklingGenes.Color]: 1,
          [DucklingGenes.Family]: 1,
        });

        try {
          await Duckies.connect(Someone).increaseAllowance(Game.address, 10_000_000_000);
          await GameAsSomeone.meldFlock(tokenIdsAndGenomes.tokenIds);
          assert(true);
        } catch {
          assert(false);
        }
      });

      it('success on melding Epic Ducklings', async () => {
        const tokenIdsAndGenomes = await generateAndMintGenomes(Collections.Duckling, {
          [DucklingGenes.Rarity]: Rarities.Epic,
          [DucklingGenes.Color]: 1,
          [DucklingGenes.Family]: 1,
        });

        try {
          await Duckies.connect(Someone).increaseAllowance(Game.address, 10_000_000_000);
          await GameAsSomeone.meldFlock(tokenIdsAndGenomes.tokenIds);
          assert(true);
        } catch {
          assert(false);
        }
      });

      it('success on melding Legendary Ducklings', async () => {
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
          await Duckies.connect(Someone).increaseAllowance(Game.address, 10_000_000_000);
          await GameAsSomeone.meldFlock(meldingTokenIds);
          assert(true);
        } catch {
          assert(false);
        }
      });
    });

    describe('Zombeak', () => {
      it('can meld into rare');

      it('can meld into epic');

      it('can meld into legendary');
    });

    describe('Mythic', () => {
      it('can meld into Mythic');

      it('Mythic id increments');

      it('revert on all Mythic minted');
    });

    it('event is emitted');
  });
});
