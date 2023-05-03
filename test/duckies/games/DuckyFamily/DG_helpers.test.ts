import { expect } from 'chai';

import {
  Collections,
  GeneDistrTypes,
  MAX_PECULIARITY,
  MYTHIC_DISPERSION,
  collectionsGeneDistributionTypes,
  collectionsGeneValuesNum,
} from './config';
import { Genome } from './genome';
import { setup } from './setup';
import { bytes3, reverse } from './helpers';

import type { DuckyGenomeTestConsumer } from '../../../../typechain-types';

describe('DuckyGenome helpers', () => {
  let DuckyGenome: DuckyGenomeTestConsumer;

  const generateUnevenGeneValue = async (valuesNum: number, bitSlice: string): Promise<number> => {
    const tx = await DuckyGenome.generateUnevenGeneValue(valuesNum, bitSlice);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'Uint8Returned');
    return event?.args?.returnedUint8 as number;
  };

  beforeEach(async () => {
    ({ DuckyGenome } = await setup());
  });
  it('_getDistributionType', async () => {
    expect(await DuckyGenome.getDistributionType(0b0, 0)).to.equal(GeneDistrTypes.Even);
    expect(await DuckyGenome.getDistributionType(0b1, 0)).to.equal(GeneDistrTypes.Uneven);
    expect(await DuckyGenome.getDistributionType(0b10, 1)).to.equal(GeneDistrTypes.Uneven);
    expect(await DuckyGenome.getDistributionType(0b010, 2)).to.equal(GeneDistrTypes.Even);
    expect(await DuckyGenome.getDistributionType(0b0010_0010, 7)).to.equal(GeneDistrTypes.Even);
    expect(await DuckyGenome.getDistributionType(0b100_0010_0010, 10)).to.equal(
      GeneDistrTypes.Uneven,
    );
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

  it('_calcConfigPeculiarity', async () => {
    expect(await DuckyGenome.calcConfigPeculiarity([1, 1, 1], reverse(0b111))).to.equal(3);
    expect(await DuckyGenome.calcConfigPeculiarity([1, 2, 3], reverse(0b111))).to.equal(6);
    expect(await DuckyGenome.calcConfigPeculiarity([1, 1, 1, 1], reverse(0b1010))).to.equal(2);
    expect(await DuckyGenome.calcConfigPeculiarity([7, 6, 5, 4, 3], reverse(0b1_1001))).to.equal(
      16,
    );

    const ducklingGeneValuesNum = [...collectionsGeneValuesNum[Collections.Duckling]];
    const ducklingGeneDistrTypes = collectionsGeneDistributionTypes[Collections.Duckling];
    expect(
      await DuckyGenome.calcConfigPeculiarity(ducklingGeneValuesNum, ducklingGeneDistrTypes),
    ).to.equal(MAX_PECULIARITY);
  });

  it('_calcPeculiarity', async () => {
    // NOTE: only genes after `generativeGenesOffset` are taken into account
    expect(
      //                                              \/ generativeGenesOffset
      await DuckyGenome.calcPeculiarity(Genome.from([0, 0, 1, 1, 1]).genome, 3, reverse(0b111)),
    ).to.equal(3);
    expect(
      await DuckyGenome.calcPeculiarity(Genome.from([0, 0, 1, 1, 1]).genome, 4, reverse(0b111)),
    ).to.equal(3);
    expect(
      await DuckyGenome.calcPeculiarity(Genome.from([0, 0, 1, 1, 1]).genome, 3, reverse(0b101)),
    ).to.equal(2);
    expect(
      await DuckyGenome.calcPeculiarity(
        Genome.from([0, 0, 1, 2, 3, 4, 5, 6]).genome,
        6,
        reverse(0b10_1011),
      ),
    ).to.equal(15);
  });

  describe('_calcUniqIdGenerationParams', () => {
    const maxUniqId = 59;

    it('mythic id range overlaps with left dispersion border', async () => {
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(0, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([0, 5]);
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(2, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([0, 7]);
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(4, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([0, 9]);
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(5, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([0, 10]);
    });

    it('mythic id range overlaps with right dispersion border', async () => {
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(55, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([50, 10]);
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(56, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([51, 9]);
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(57, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([52, 8]);
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(59, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([54, 6]);
    });

    it('mythic id range does not overlap with dispersion borders', async () => {
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(6, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([1, 10]);
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(7, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([2, 10]);
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(8, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([3, 10]);
      expect(
        await DuckyGenome.calcUniqIdGenerationParams(54, maxUniqId, MYTHIC_DISPERSION),
      ).to.deep.equal([49, 10]);
    });
  });
});
