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

  const generateUnevenGeneValue = async (valuesNum: number, bitSlice: string): Promise<number> => {
    const tx = await Game.generateUnevenGeneValue(valuesNum, bitSlice);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'Uint8Returned');
    return event?.args?.returnedUint8 as number;
  };

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
    // value 0: [883600, 1000000]                            \/ account for +1 shift in the function
    expect(await generateUnevenGeneValue(32, bytes3(883_600 - 1))).to.equal(0);
    expect(await generateUnevenGeneValue(32, bytes3(969_696))).to.equal(0);
    expect(await generateUnevenGeneValue(32, bytes3(1_000_000 - 1))).to.equal(0);

    // value 1: [779689, 883599]
    expect(await generateUnevenGeneValue(32, bytes3(779_689 - 1))).to.equal(1);
    expect(await generateUnevenGeneValue(32, bytes3(805_696))).to.equal(1);
    expect(await generateUnevenGeneValue(32, bytes3(883_599 - 1))).to.equal(1);

    // value 2: [687241, 779688]
    expect(await generateUnevenGeneValue(32, bytes3(687_241 - 1))).to.equal(2);
    expect(await generateUnevenGeneValue(32, bytes3(750_750))).to.equal(2);
    expect(await generateUnevenGeneValue(32, bytes3(779_688 - 1))).to.equal(2);

    // value 5: [469225, 532899]
    expect(await generateUnevenGeneValue(32, bytes3(469_225 - 1))).to.equal(5);
    expect(await generateUnevenGeneValue(32, bytes3(500_000))).to.equal(5);
    expect(await generateUnevenGeneValue(32, bytes3(532_899 - 1))).to.equal(5);

    // value 17: [78961, 94248]
    expect(await generateUnevenGeneValue(32, bytes3(78_961 - 1))).to.equal(17);
    expect(await generateUnevenGeneValue(32, bytes3(85_858))).to.equal(17);
    expect(await generateUnevenGeneValue(32, bytes3(94_248 - 1))).to.equal(17);

    // value 29: [1089, 2499]
    expect(await generateUnevenGeneValue(32, bytes3(1089 - 1))).to.equal(29);
    expect(await generateUnevenGeneValue(32, bytes3(1750))).to.equal(29);
    expect(await generateUnevenGeneValue(32, bytes3(2499 - 1))).to.equal(29);

    // value 30: [256, 1088]
    expect(await generateUnevenGeneValue(32, bytes3(256 - 1))).to.equal(30);
    expect(await generateUnevenGeneValue(32, bytes3(555))).to.equal(30);
    expect(await generateUnevenGeneValue(32, bytes3(1088 - 1))).to.equal(30);

    // value 31: [1, 255]
    // eslint-disable-next-line sonarjs/no-identical-expressions
    expect(await generateUnevenGeneValue(32, bytes3(1 - 1))).to.equal(31);
    expect(await generateUnevenGeneValue(32, bytes3(69))).to.equal(31);
    expect(await generateUnevenGeneValue(32, bytes3(255 - 1))).to.equal(31);
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
