import { assert } from 'chai';

import { Collections, DucklingGenes, Rarities, ZombeakGenes } from './config';
import { RandomGenomeConfig, randomGenome } from './helpers';
import { setup } from './setup';

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

  beforeEach(async () => {
    ({ Someone, GenomeSetter, Duckies, Ducklings, Game } = await setup());
  });

  interface TokenIdsAndGenomes {
    tokenIds: number[];
    genomes: bigint[];
  }

  const generateAndMintGenomesForMelding = async (
    collection: Collections,
    config?: RandomGenomeConfig,
  ): Promise<TokenIdsAndGenomes> => {
    const tokenIds: number[] = [];
    const genomes: bigint[] = [];

    for (let i = 0; i < 5; i++) {
      const tokenIdAndGenome = await extractMintedTokenId(
        await mintTo(Someone.address, randomGenome(collection, config)),
      );
      tokenIds.push(tokenIdAndGenome.tokenId);
      genomes.push(tokenIdAndGenome.genome);
    }

    return { tokenIds, genomes };
  };

  describe('meldGenes', () => {
    it('can meld', async () => {
      await mintTo(
        Someone.address,
        // eslint-disable-next-line unicorn/numeric-separators-style
        182700775082802730930410854023168n,
      );

      await mintTo(
        Someone.address,
        // eslint-disable-next-line unicorn/numeric-separators-style
        60926767771370839915004195766272n,
      );

      await mintTo(
        Someone.address,
        // eslint-disable-next-line unicorn/numeric-separators-style
        121932763563511447839369064611840n,
      );

      await mintTo(
        Someone.address,
        // eslint-disable-next-line unicorn/numeric-separators-style
        61164767845294952445087173574656n,
      );

      await mintTo(
        Someone.address,
        // eslint-disable-next-line unicorn/numeric-separators-style
        162419591386637366713636064854016n,
      );

      await Game.connect(Someone).meldFlock([0, 1, 2, 3, 4]);
    });

    it('uneven gene can mutate');

    it('even gene can not mutate');

    it('can inherit from all parents');
  });

  describe('isCollectionMutating', () => {
    it('<= Rare can collection mutate');

    it('legendary can not mutate');
  });

  describe('requireGenomesSatisfyMelding', () => {
    it('success on Common Duckling', async () => {
      const tokenIdsAndGenomes = await generateAndMintGenomesForMelding(Collections.Duckling, {
        [DucklingGenes.Rarity]: Rarities.Common,
        [DucklingGenes.Color]: 0,
      });

      await Game.requireGenomesSatisfyMelding(tokenIdsAndGenomes.genomes);
    });

    it('success on Rare Duckling', async () => {
      const tokenIdsAndGenomes = await generateAndMintGenomesForMelding(Collections.Duckling, {
        [DucklingGenes.Rarity]: Rarities.Rare,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 0,
      });

      await Game.requireGenomesSatisfyMelding(tokenIdsAndGenomes.genomes);
    });

    it('success on Epic Duckling', async () => {
      const tokenIdsAndGenomes = await generateAndMintGenomesForMelding(Collections.Duckling, {
        [DucklingGenes.Rarity]: Rarities.Epic,
        [DucklingGenes.Color]: 0,
        [DucklingGenes.Family]: 0,
      });

      await Game.requireGenomesSatisfyMelding(tokenIdsAndGenomes.genomes);
    });

    it('success on Legendary Duckling', async () => {
      const genomes: bigint[] = [];
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
          genomes.push(tokenIdAndGenome.genome);
        }
      })();

      await Game.requireGenomesSatisfyMelding(genomes);
    });

    it('success on Common Zombeak', async () => {
      const tokenIdsAndGenomes = await generateAndMintGenomesForMelding(Collections.Zombeak, {
        [ZombeakGenes.Rarity]: Rarities.Common,
        [ZombeakGenes.Color]: 0,
      });

      await Game.requireGenomesSatisfyMelding(tokenIdsAndGenomes.genomes);
    });

    it('success on Rare Zombeak', async () => {
      const tokenIdsAndGenomes = await generateAndMintGenomesForMelding(Collections.Zombeak, {
        [ZombeakGenes.Rarity]: Rarities.Rare,
        [ZombeakGenes.Color]: 0,
        [ZombeakGenes.Family]: 0,
      });

      await Game.requireGenomesSatisfyMelding(tokenIdsAndGenomes.genomes);
    });

    it('success on Epic Zombeak', async () => {
      const tokenIdsAndGenomes = await generateAndMintGenomesForMelding(Collections.Zombeak, {
        [ZombeakGenes.Rarity]: Rarities.Epic,
        [ZombeakGenes.Color]: 0,
        [ZombeakGenes.Family]: 0,
      });

      await Game.requireGenomesSatisfyMelding(tokenIdsAndGenomes.genomes);
    });

    it('revert on different collections');

    it('revert on different rarities');

    it('revert on Mythic');

    it('revert on legendary Zombeak');

    it('revert on legendaries having different color');

    it('revert on legendaries having repeated families');

    it('revert on not legendary having different color and different family');
  });

  // eslint-disable-next-line sonarjs/cognitive-complexity
  describe('meldFlock', () => {
    describe.skip('Duckling', () => {
      it('success on melding Common Ducklings', async () => {
        const tokenIdsAndGenomes = await generateAndMintGenomesForMelding(Collections.Duckling, {
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
        const tokenIdsAndGenomes = await generateAndMintGenomesForMelding(Collections.Duckling, {
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
        const tokenIdsAndGenomes = await generateAndMintGenomesForMelding(Collections.Duckling, {
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
