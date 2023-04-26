import { assert } from 'chai';

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
  collectionsGeneValuesNum,
} from '../../duckies/games/DuckyFamily/config';
import { setup } from '../../duckies/games/DuckyFamily/setup';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { DucklingsV1, DuckyFamilyV1, TESTDuckyFamilyV1 } from '../../../typechain-types';

const seed = '0xaabbcc';

describe('Benchmark DuckyFamilyV1 melding', () => {
  let Someone: SignerWithAddress;
  let GenomeSetter: SignerWithAddress;

  let Ducklings: DucklingsV1;
  let Game: TESTDuckyFamilyV1;

  let GameAsSomeone: DuckyFamilyV1;

  let mintTo: MintToFuncT;
  let generateAndMintGenomes: GenerateAndMintGenomesFunctT;

  const isCollectionMutating = async (rarity: Rarities, seed: string): Promise<boolean> => {
    const tx = await Game.isCollectionMutating(rarity, seed);
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
    seed: string,
  ): Promise<number> => {
    const tx = await Game.meldGenes(genomes, gene, maxGeneValue, geneDistrType, seed);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GeneReturned');
    // gene is already a number
    return event?.args?.gene as number;
  };

  beforeEach(async () => {
    ({ Someone, GenomeSetter, Ducklings, Game, GameAsSomeone } = await setup());

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

  it('isCollectionMutating', async () => {
    await isCollectionMutating(Rarities.Common, seed);
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
      seed,
    );
  });
});
