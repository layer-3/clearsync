import { expect } from 'chai';

import {
  Collections,
  GeneDistrTypes,
  MAX_PECULIARITY,
  collectionsGeneValuesNum,
  generativeGenesOffset,
} from './config';
import { Genome } from './genome';
import { setup } from './setup';

import type { TESTDuckyFamilyV1 } from '../../../../typechain-types';

describe('DuckyGameV1 helpers', () => {
  let Game: TESTDuckyFamilyV1;

  beforeEach(async () => {
    ({ Game } = await setup());
  });
  it('_getDistributionType', async () => {
    expect(await Game.getDistributionType(0b0, 0)).to.equal(GeneDistrTypes.Even);
    expect(await Game.getDistributionType(0b1, 0)).to.equal(GeneDistrTypes.Uneven);
    expect(await Game.getDistributionType(0b10, 1)).to.equal(GeneDistrTypes.Uneven);
    expect(await Game.getDistributionType(0b010, 2)).to.equal(GeneDistrTypes.Even);
    expect(await Game.getDistributionType(0b0010_0010, 7)).to.equal(GeneDistrTypes.Even);
    expect(await Game.getDistributionType(0b100_0010_0010, 10)).to.equal(GeneDistrTypes.Uneven);
  });

  it('_calcMaxPeculiarity', async () => {
    expect(await Game.calcMaxPeculiarity()).to.equal(MAX_PECULIARITY);
  });

  it('_calcPeculiarity', async () => {
    const geneValues = Array.from({
      length: collectionsGeneValuesNum[Collections.Duckling].length + generativeGenesOffset,
    }).fill(1) as number[];
    expect(await Game.calcPeculiarity(Genome.from(geneValues).genome)).to.equal(8); // 8 uneven genes
    expect(
      await Game.calcPeculiarity(
        Genome.from([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14]).genome,
      ),
    ).to.equal(5 + 6 + 7 + 8 + 9 + 11 + 12 + 14); // 001111101101
  });
});
