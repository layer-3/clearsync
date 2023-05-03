import { assert } from 'chai';
import { utils } from 'ethers';

import {
  GenerateAndMintGenomesFunctT,
  MintToFuncT,
  randomGenome,
  randomGenomes,
  setupGenerateAndMintGenomes,
  setupMintTo,
} from '../../duckies/games/DuckyFamily/helpers';
import {
  Collections,
  DucklingGenes,
  FLOCK_SIZE,
  GeneDistrTypes,
  Rarities,
  ZombeakGenes,
  collectionMutationChances,
  collectionsGeneValuesNum,
  geneInheritanceChances,
  geneMutationChance,
} from '../../duckies/games/DuckyFamily/config';
import { setup } from '../../duckies/games/DuckyFamily/setup';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type {
  DucklingsV1,
  DuckyFamilyV1,
  DuckyGenomeTestConsumer,
  TESTDuckyFamilyV1,
} from '../../../typechain-types';

const SEED = utils.id('seed');
const BIT_SLICE = '0xaabbcc';

describe('Benchmark DuckyFamilyV1 melding', () => {
  let Someone: SignerWithAddress;
  let GenomeSetter: SignerWithAddress;

  let Ducklings: DucklingsV1;
  let DuckyGenome: DuckyGenomeTestConsumer;
  let Game: TESTDuckyFamilyV1;

  let GameAsSomeone: DuckyFamilyV1;

  let mintTo: MintToFuncT;
  let generateAndMintGenomes: GenerateAndMintGenomesFunctT;

  const isCollectionMutating = async (
    rarity: Rarities,
    bitSlice: string = BIT_SLICE,
  ): Promise<boolean> => {
    const tx = await DuckyGenome.isCollectionMutating(rarity, collectionMutationChances, bitSlice);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'BoolReturned');
    return event?.args?.returnedBool as boolean;
  };

  const meldGenomes = async (genomes: bigint[], seed: string = SEED): Promise<bigint> => {
    const tx = await Game.meldGenomes(genomes, seed);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  const meldGenes = async (
    genomes: bigint[],
    gene: number,
    maxGeneValue: number,
    geneDistrType: GeneDistrTypes,
    seed: string,
  ): Promise<number> => {
    const tx = await DuckyGenome.meldGenes(
      genomes,
      gene,
      maxGeneValue,
      geneDistrType,
      geneMutationChance,
      geneInheritanceChances,
      seed,
    );
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GeneReturned');
    // gene is already a number
    return event?.args?.gene as number;
  };

  beforeEach(async () => {
    ({ Someone, GenomeSetter, Ducklings, DuckyGenome, Game, GameAsSomeone } = await setup());

    mintTo = setupMintTo(Ducklings.connect(GenomeSetter));
    generateAndMintGenomes = setupGenerateAndMintGenomes(mintTo, Someone.address);
  });

  it('meld Common Ducklings', async () => {
    const { tokenIds } = await generateAndMintGenomes(Collections.Duckling, {
      [DucklingGenes.Rarity]: Rarities.Common,
      [DucklingGenes.Color]: 0,
    });

    await GameAsSomeone.meldFlock(tokenIds);
  });

  it('meld Common Zombeaks', async () => {
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

  it('meldGenomes', async () => {
    await Game.setCollectionMutationChances([0, 0, 0, 0]);

    const genomes = randomGenomes(Collections.Duckling, {
      amount: FLOCK_SIZE,
      [DucklingGenes.Rarity]: Rarities.Common,
      [DucklingGenes.Color]: 0,
    });

    await meldGenomes(genomes);
  });

  it('meldGenomes Legendary Ducklings', async () => {
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

    await meldGenomes(genomes);
  });

  it('isCollectionMutating', async () => {
    await isCollectionMutating(Rarities.Common, BIT_SLICE);
  });

  it('meldGenes', async () => {
    const geneValuesNum = collectionsGeneValuesNum[Collections.Duckling];
    const genomes = [];
    for (let i = 0; i < FLOCK_SIZE; i++) {
      const genome = randomGenome(Collections.Duckling, {
        [DucklingGenes.Head]: i,
      });
      genomes.push(genome);
    }

    await meldGenes(
      genomes,
      DucklingGenes.Head,
      geneValuesNum[DucklingGenes.Head],
      GeneDistrTypes.Uneven,
      BIT_SLICE,
    );
  });
});
