import {
  Collections,
  DucklingGenes,
  GeneToValue,
  GeneralGenes,
  Genes,
  MythicGenes,
  ZombeakGenes,
  baseMagicNumber,
  collectionGeneIdx,
  collectionsGeneValuesNum,
  generativeGenesOffset,
  mythicAmount,
  raritiesNum,
  rarityGeneIdx,
} from './config';
import { Genome } from './genome';

import type { DucklingsV1 } from '../../../../typechain-types';
import type { ContractTransaction } from 'ethers';

export type GenesConfig = { [key in Genes]?: number };
// only isTransferable flag is supported for now
export type RandomGenomeConfig = GenesConfig & { isTransferable?: boolean };

export const randomMaxNum = (maxNum: number): number => Math.floor(Math.random() * maxNum);

export const bytes3 = (rem: number): string => {
  const bytes3 = rem.toString(16).padStart(6, '0');
  return `0x${bytes3}`;
};

export const reverse = (num: number): number => {
  const bin = num.toString(2);
  const reversed = [...bin].reverse().join('');
  return Number.parseInt(reversed, 2);
};

export function randomGenome(collectionId: Collections, config?: RandomGenomeConfig): bigint {
  const genome = new Genome();
  genome.setGene(collectionGeneIdx, collectionId);

  if (collectionId == Collections.Mythic) {
    if (config?.[MythicGenes.UniqId]) {
      genome.setGene(MythicGenes.UniqId, config[MythicGenes.UniqId]).genome;
    } else {
      genome.randomizeGene(MythicGenes.UniqId, mythicAmount).genome;
    }
  } else {
    const rarity = config?.[rarityGeneIdx] ?? randomMaxNum(raritiesNum - 1);
    genome.setGene(rarityGeneIdx, rarity);
  }

  const geneValuesNum = collectionsGeneValuesNum[collectionId];

  for (const [i, geneValues] of geneValuesNum.entries()) {
    if (config?.[(i + generativeGenesOffset) as Genes] === undefined) {
      const geneValue = randomMaxNum(geneValues) + 1;
      genome.setGene(i + generativeGenesOffset, geneValue);
    } else {
      genome.setGene(
        i + generativeGenesOffset,
        // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
        config[(i + generativeGenesOffset) as Genes]!,
      );
    }
  }

  let flags = config?.[GeneralGenes.Flags] ?? 0;

  // only isTransferable flag is supported for now
  if (config?.isTransferable) {
    flags = 1;
  }

  genome.setGene(GeneralGenes.Flags, flags);

  // default is base magic number as this function generates only Duckling and Zombeak genomes
  const magicNumber = config?.[GeneralGenes.MagicNumber] ?? baseMagicNumber;
  genome.setGene(GeneralGenes.MagicNumber, magicNumber);

  return genome.genome;
}

export function randomGenomes(
  collectionId: Collections,
  config?: RandomGenomeConfig & { amount?: number },
): bigint[] {
  const amount = config?.amount ?? 1;
  return Array.from({ length: amount })
    .fill(0)
    .map(() => randomGenome(collectionId, config));
}

export function parseGenome(genome_: bigint): GeneToValue {
  const genome = new Genome(genome_);

  let NFTGenes;
  // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment, @typescript-eslint/no-explicit-any
  const NFTGenome = {} as any; // use any as we are building this object dynamically

  const collection = genome.getGene(collectionGeneIdx);

  switch (collection) {
    case Collections.Duckling: {
      NFTGenes = DucklingGenes;
      break;
    }
    case Collections.Zombeak: {
      NFTGenes = ZombeakGenes;
      break;
    }
    case Collections.Mythic: {
      NFTGenes = MythicGenes;
      break;
    }
    default: {
      throw new Error('Unknown collection');
    }
  }

  NFTGenome[NFTGenes[GeneralGenes.Collection]] = collection;

  for (const value in NFTGenes) {
    if (Number.isNaN(Number(value))) continue;
    const numValue = Number(value);
    NFTGenome[NFTGenes[numValue]] = genome.getGene(numValue as Genes);
  }

  return NFTGenome as GeneToValue;
}

export type MintToFuncT = (
  to: string,
  genome: bigint,
  isTransferable?: boolean,
) => Promise<ContractTransaction>;

export const setupMintTo = (DucklingsAsGame: DucklingsV1): MintToFuncT => {
  return async (to: string, genome: bigint): Promise<ContractTransaction> => {
    return await DucklingsAsGame.mintTo(to, genome);
  };
};

interface TokenIdAndGenome {
  tokenId: number;
  genome: bigint;
}

export const extractMintedTokenId = async (tx: ContractTransaction): Promise<TokenIdAndGenome> => {
  const receipt = await tx.wait();
  const event = receipt.events?.find((e) => e.event === 'Minted');
  const tokenId = event?.args?.tokenId.toNumber() as number;
  const genome = event?.args?.genome.toBigInt() as bigint;
  return { tokenId, genome };
};

interface TokenIdsAndGenomes {
  tokenIds: number[];
  genomes: bigint[];
}

export type GenerateAndMintGenomesFunctT = (
  collection: Collections,
  config?: RandomGenomeConfig & { amount?: number },
) => Promise<TokenIdsAndGenomes>;

export const setupGenerateAndMintGenomes = (
  mintTo: MintToFuncT,
  to: string,
): GenerateAndMintGenomesFunctT => {
  return async (
    collection: Collections,
    config?: RandomGenomeConfig & { amount?: number },
  ): Promise<TokenIdsAndGenomes> => {
    const tokenIds: number[] = [];
    const genomes: bigint[] = [];

    const amount = config?.amount ?? 5;

    for (let i = 0; i < amount; i++) {
      const tokenIdAndGenome = await extractMintedTokenId(
        await mintTo(to, randomGenome(collection, config)),
      );
      tokenIds.push(tokenIdAndGenome.tokenId);
      genomes.push(tokenIdAndGenome.genome);
    }

    return { tokenIds, genomes };
  };
};
