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
import { bytes3 } from './helpers';

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

  it('_generateUnevenGeneValue', async () => {
    // value 0: [961, 1023]
    expect(await Game.generateUnevenGeneValue(32, bytes3(961))).to.equal(0);
    expect(await Game.generateUnevenGeneValue(32, bytes3(1000))).to.equal(0);
    expect(await Game.generateUnevenGeneValue(32, bytes3(1023))).to.equal(0);

    // value 1: [900, 960]
    expect(await Game.generateUnevenGeneValue(32, bytes3(900))).to.equal(1);
    expect(await Game.generateUnevenGeneValue(32, bytes3(942))).to.equal(1);
    expect(await Game.generateUnevenGeneValue(32, bytes3(960))).to.equal(1);

    // value 2: [841, 899]
    expect(await Game.generateUnevenGeneValue(32, bytes3(841))).to.equal(2);
    expect(await Game.generateUnevenGeneValue(32, bytes3(869))).to.equal(2);
    expect(await Game.generateUnevenGeneValue(32, bytes3(899))).to.equal(2);

    // value 5: [676, 728]
    expect(await Game.generateUnevenGeneValue(32, bytes3(676))).to.equal(5);
    expect(await Game.generateUnevenGeneValue(32, bytes3(700))).to.equal(5);
    expect(await Game.generateUnevenGeneValue(32, bytes3(728))).to.equal(5);

    // value 17: [196, 224]
    expect(await Game.generateUnevenGeneValue(32, bytes3(196))).to.equal(17);
    expect(await Game.generateUnevenGeneValue(32, bytes3(210))).to.equal(17);
    expect(await Game.generateUnevenGeneValue(32, bytes3(224))).to.equal(17);

    // value 29: [4, 8]
    expect(await Game.generateUnevenGeneValue(32, bytes3(4))).to.equal(29);
    expect(await Game.generateUnevenGeneValue(32, bytes3(6))).to.equal(29);
    expect(await Game.generateUnevenGeneValue(32, bytes3(8))).to.equal(29);

    // value 30: [1, 3]
    expect(await Game.generateUnevenGeneValue(32, bytes3(1))).to.equal(30);
    expect(await Game.generateUnevenGeneValue(32, bytes3(2))).to.equal(30);
    expect(await Game.generateUnevenGeneValue(32, bytes3(3))).to.equal(30);

    // value 31: [0, 0]
    expect(await Game.generateUnevenGeneValue(32, bytes3(0))).to.equal(31);
  });

  it('_floorSqrt', async () => {
    expect(await Game.floorSqrt(0)).to.equal(0);
    expect(await Game.floorSqrt(1)).to.equal(1);
    expect(await Game.floorSqrt(2)).to.equal(1);
    expect(await Game.floorSqrt(3)).to.equal(1);
    expect(await Game.floorSqrt(4)).to.equal(2);
    expect(await Game.floorSqrt(5)).to.equal(2);
    expect(await Game.floorSqrt(6)).to.equal(2);
    expect(await Game.floorSqrt(7)).to.equal(2);
    expect(await Game.floorSqrt(8)).to.equal(2);
    expect(await Game.floorSqrt(9)).to.equal(3);
    expect(await Game.floorSqrt(15)).to.equal(3);
    expect(await Game.floorSqrt(16)).to.equal(4);
    expect(await Game.floorSqrt(24)).to.equal(4);
    expect(await Game.floorSqrt(25)).to.equal(5);
    expect(await Game.floorSqrt(35)).to.equal(5);
    expect(await Game.floorSqrt(36)).to.equal(6);
    expect(await Game.floorSqrt(48)).to.equal(6);
    expect(await Game.floorSqrt(49)).to.equal(7);
    expect(await Game.floorSqrt(63)).to.equal(7);
    expect(await Game.floorSqrt(64)).to.equal(8);
    expect(await Game.floorSqrt(80)).to.equal(8);
    expect(await Game.floorSqrt(81)).to.equal(9);
    expect(await Game.floorSqrt(99)).to.equal(9);
    expect(await Game.floorSqrt(100)).to.equal(10);
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

  describe('_calcUniqIdGenerationParams', () => {
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
