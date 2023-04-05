import {
  Collections,
  Genes,
  MythicGenes,
  collectionGeneIdx,
  collectionsGeneValuesNum,
  mythicAmount,
  raritiesNum,
  rarityGeneIdx,
} from './config';
import { Genome } from './genome';

type RandomGenomeConfig = Record<Genes, number>;

const randomMaxNum = (maxNum: number): number => Math.round(Math.random() * maxNum);

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
    let geneValue;
    if (config?.[i as Genes]) {
      geneValue = config[i as Genes];
    } else {
      geneValue = randomMaxNum(geneValues);
    }

    genome.setGene(i, geneValue);
  }

  return genome.genome;
}
