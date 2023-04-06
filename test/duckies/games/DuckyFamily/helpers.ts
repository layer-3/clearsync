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
