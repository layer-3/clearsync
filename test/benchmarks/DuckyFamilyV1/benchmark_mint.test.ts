import { setup } from '../../duckies/games/DuckyFamily/setup';
import { Collections, GeneDistrTypes } from '../../duckies/games/DuckyFamily/config';

import type { DuckyFamilyV1, TESTDuckyFamilyV1 } from '../../../typechain-types';

const seed = '0xaabbcc';

describe('Benchmark DuckyFamilyV1 minting', () => {
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
    const tx = await Game.generateAndSetGenes(genome, collectionId, seed);
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
    const tx = await Game.generateAndSetGene(genome, geneIx, geneValuesNum, distrType, seed);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  beforeEach(async () => {
    ({ Game, GameAsSomeone } = await setup());
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
    await generateAndSetGenes(0n, Collections.Duckling, seed);
  });

  it('generateAndSetGene even', async () => {
    await generateAndSetGene(0n, 0, 2, GeneDistrTypes.Even, seed);
  });

  it('generateAndSetGene uneven', async () => {
    await generateAndSetGene(0n, 0, 2, GeneDistrTypes.Uneven, seed);
  });

  it('generateUnevenGeneValue', async () => {
    await Game.generateUnevenGeneValue(32, '0xaabbcc');
  });
});
