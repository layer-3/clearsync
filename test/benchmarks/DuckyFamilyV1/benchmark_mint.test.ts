import { utils } from 'ethers';

import { setup } from '../../duckies/games/DuckyFamily/setup';
import {
  Collections,
  GeneDistrTypes,
  collectionsGeneDistributionTypes,
  collectionsGeneValuesNum,
} from '../../duckies/games/DuckyFamily/config';

import type {
  DuckyFamilyV1,
  DuckyGenomeTestConsumer,
  TESTDuckyFamilyV1,
} from '../../../typechain-types';

const SEED = utils.id('seed');
const BIT_SLICE = '0xaabbcc';

describe('Benchmark DuckyFamilyV1 minting', () => {
  let DuckyGenome: DuckyGenomeTestConsumer;
  let Game: TESTDuckyFamilyV1;
  let GameAsSomeone: DuckyFamilyV1;

  const generateGenome = async (collectionId: Collections): Promise<bigint> => {
    const tx = await Game.generateGenome(collectionId);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  const generateAndSetGenes = async (
    genome: bigint,
    collectionId: Collections,
    seed: string,
  ): Promise<bigint> => {
    const geneValuesNum = [...collectionsGeneValuesNum[collectionId]];
    const geneDistrTypes = collectionsGeneDistributionTypes[collectionId];
    const tx = await DuckyGenome.generateAndSetGenes(
      genome,
      collectionId,
      geneValuesNum,
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
    seed: string,
  ): Promise<bigint> => {
    const tx = await DuckyGenome.generateAndSetGene(genome, geneIx, geneValuesNum, distrType, seed);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  beforeEach(async () => {
    ({ DuckyGenome, Game, GameAsSomeone } = await setup());
    await Game.setMintPrice(1);
  });

  it('mint', async () => {
    await GameAsSomeone.mintPack(1);
    await GameAsSomeone.mintPack(1);
  });

  it('generateGenome', async () => {
    await generateGenome(Collections.Duckling);
  });

  it('generateAndSetGenes', async () => {
    await generateAndSetGenes(0n, Collections.Duckling, SEED);
  });

  it('generateAndSetGene even', async () => {
    await generateAndSetGene(0n, 0, 2, GeneDistrTypes.Even, BIT_SLICE);
  });

  it('generateAndSetGene uneven', async () => {
    await generateAndSetGene(0n, 0, 2, GeneDistrTypes.Uneven, BIT_SLICE);
  });

  it('generateUnevenGeneValue', async () => {
    await DuckyGenome.generateUnevenGeneValue(32, BIT_SLICE);
  });
});
