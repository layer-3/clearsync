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

  describe.only('_calcUniqIdGenerationParams', () => {
    const maxUniqId = 59;

    it('mythic id range overlaps with left dispersion border', async () => {
      expect(await Game.calcUniqIdGenerationParams(0, maxUniqId)).to.deep.equal([0, 5]);
      expect(await Game.calcUniqIdGenerationParams(2, maxUniqId)).to.deep.equal([0, 7]);
      expect(await Game.calcUniqIdGenerationParams(4, maxUniqId)).to.deep.equal([0, 9]);
      expect(await Game.calcUniqIdGenerationParams(5, maxUniqId)).to.deep.equal([0, 10]);
    });

    it('mythic id range overlaps with right dispersion border', async () => {
      expect(await Game.calcUniqIdGenerationParams(55, maxUniqId)).to.deep.equal([50, 10]);
      expect(await Game.calcUniqIdGenerationParams(56, maxUniqId)).to.deep.equal([51, 9]);
      expect(await Game.calcUniqIdGenerationParams(57, maxUniqId)).to.deep.equal([52, 8]);
      expect(await Game.calcUniqIdGenerationParams(59, maxUniqId)).to.deep.equal([54, 6]);
    });

    it('mythic id range does not overlap with dispersion borders', async () => {
      expect(await Game.calcUniqIdGenerationParams(6, maxUniqId)).to.deep.equal([1, 10]);
      expect(await Game.calcUniqIdGenerationParams(7, maxUniqId)).to.deep.equal([2, 10]);
      expect(await Game.calcUniqIdGenerationParams(8, maxUniqId)).to.deep.equal([3, 10]);
      expect(await Game.calcUniqIdGenerationParams(54, maxUniqId)).to.deep.equal([49, 10]);
    });
  });
});
