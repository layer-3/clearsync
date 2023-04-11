import {
  Collections,
  Genes,
  MythicGenes,
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

export type RandomGenomeConfig = { [key in Genes]?: number };

export const randomMaxNum = (maxNum: number): number => Math.floor(Math.random() * (maxNum + 1));

export function randomGenome(collectionId: Collections, config?: RandomGenomeConfig): bigint {
  const genome = new Genome();
  genome.setGene(collectionGeneIdx, collectionId);

  if (collectionId == Collections.Mythic) {
    if (config?.[MythicGenes.UniqId]) {
      return genome.setGene(MythicGenes.UniqId, config[MythicGenes.UniqId]).genome;
    } else {
      return genome.randomizeGene(MythicGenes.UniqId, mythicAmount).genome;
    }
  }

  genome.setGene(rarityGeneIdx, randomMaxNum(raritiesNum - 1));

  const geneValuesNum = collectionsGeneValuesNum[collectionId];

  for (const [i, geneValues] of geneValuesNum.entries()) {
    let geneValue = 0;
    if (config?.[i as Genes] === undefined) {
      geneValue = randomMaxNum(geneValues);
      genome.setGene(i + generativeGenesOffset, geneValue);
    } else {
      geneValue = config[i as Genes] as unknown as number;
      genome.setGene(i, geneValue);
    }
  }

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

export type MintToFuncT = (
  to: string,
  genome: bigint,
  isTransferable?: boolean,
) => Promise<ContractTransaction>;

export const setupMintTo = (DucklingsAsGame: DucklingsV1): MintToFuncT => {
  return async (
    to: string,
    genome: bigint,
    isTransferable?: boolean,
  ): Promise<ContractTransaction> => {
    return await DucklingsAsGame.mintTo(to, genome, isTransferable ?? true);
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
