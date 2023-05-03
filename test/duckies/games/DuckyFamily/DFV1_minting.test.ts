import { expect } from 'chai';
import { ethers } from 'hardhat';
import { anyUint } from '@nomicfoundation/hardhat-chai-matchers/withArgs';
import { utils } from 'ethers';

import { setup } from './setup';
import {
  Collections,
  MAX_PACK_SIZE,
  MAX_PECULIARITY,
  MYTHIC_DISPERSION,
  MythicGenes,
  baseMagicNumber,
  collectionGeneIdx,
  magicNumberGeneIdx,
  mythicAmount,
  mythicMagicNumber,
  rarityGeneIdx,
} from './config';
import { Genome } from './genome';

import type {
  DucklingsV1,
  DuckyFamilyV1,
  TESTDuckyFamilyV1,
  YellowToken,
} from '../../../../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

const SEED = utils.id('seed');

describe('DuckyFamilyV1 minting', () => {
  let Someone: SignerWithAddress;

  let Duckies: YellowToken;
  let Ducklings: DucklingsV1;
  let Game: TESTDuckyFamilyV1;
  let GameAsSomeone: DuckyFamilyV1;

  beforeEach(async () => {
    ({ Someone, Duckies, Ducklings, Game, GameAsSomeone } = await setup());
  });

  const generateGenome = async (
    collectionId: Collections,
    seed: string = SEED,
  ): Promise<bigint> => {
    const tx = await Game.generateGenome(collectionId, seed);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  const generateMythicGenome = async (genomes: bigint[], seed: string = SEED): Promise<bigint> => {
    const tx = await Game.generateMythicGenome(genomes, MAX_PECULIARITY, mythicAmount, seed);
    const receipt = await tx.wait();
    const event = receipt.events?.find((e) => e.event === 'GenomeReturned');
    return event?.args?.genome.toBigInt() as bigint;
  };

  describe('generateGenome', () => {
    beforeEach(async () => {
      await Game.setRarityChances([1000, 0, 0, 0]);
    });

    it('set Duckling collectionId and rarity', async () => {
      const _genome = await generateGenome(Collections.Duckling);
      const genome = new Genome(_genome);
      expect(genome.getGene(collectionGeneIdx)).to.equal(Collections.Duckling);
      expect(genome.getGene(rarityGeneIdx)).to.equal(0);
    });

    it('set Zombeak collectionId and rarity', async () => {
      const _genome = await generateGenome(Collections.Zombeak);
      const genome = new Genome(_genome);
      expect(genome.getGene(collectionGeneIdx)).to.equal(Collections.Zombeak);
      expect(genome.getGene(rarityGeneIdx)).to.equal(0);
    });

    it('set correct magic number for Duckling', async () => {
      const _genome = await generateGenome(Collections.Duckling);
      const genome = new Genome(_genome);
      expect(genome.getGene(magicNumberGeneIdx)).to.equal(baseMagicNumber);
    });

    it('set correct magic number for Zombeak', async () => {
      const _genome = await generateGenome(Collections.Zombeak);
      const genome = new Genome(_genome);
      expect(genome.getGene(magicNumberGeneIdx)).to.equal(baseMagicNumber);
    });

    it('revert on collectionId not Duckling or Zombeak', async () => {
      const notDucklingOrZombeak = 2;
      await expect(Game.generateGenome(2, SEED))
        .to.be.revertedWithCustomError(Game, 'MintingRulesViolated')
        .withArgs(notDucklingOrZombeak, 1);
    });
  });

  describe('generateMythicGenome', () => {
    const zeroedGenome = new Genome(0).genome;

    type generateGenomeArgsType = [bigint, bigint, bigint, bigint, bigint];
    const maxUniqId = mythicAmount - 1;

    const genomesWithTheoUniqId = (theoUniqId: number): generateGenomeArgsType => {
      const peculiarity = Math.ceil((theoUniqId / maxUniqId) * MAX_PECULIARITY);
      const genome = Genome.from([0, 0, 0, 0, peculiarity]).genome;
      return Array.from({ length: 5 }).fill(genome) as generateGenomeArgsType;
    };

    it('correct range for zero peculiarity', async () => {
      const zeroedGenomes = Array.from({ length: 5 }).fill(zeroedGenome);
      const _genome = await generateMythicGenome(zeroedGenomes as generateGenomeArgsType);
      const genome = new Genome(_genome);
      expect(genome.getGene(MythicGenes.UniqId)).to.be.within(0, MYTHIC_DISPERSION);
    });

    it('correct range for low UniqId', async () => {
      const theoUniqId = 2;
      const _genome = await generateMythicGenome(genomesWithTheoUniqId(theoUniqId));
      const genome = new Genome(_genome);
      expect(genome.getGene(MythicGenes.UniqId)).to.be.within(0, theoUniqId + MYTHIC_DISPERSION);
    });

    it('correct range for medium UniqId', async () => {
      const theoUniqId = 30;
      const _genome = await generateMythicGenome(genomesWithTheoUniqId(theoUniqId));
      const genome = new Genome(_genome);
      expect(genome.getGene(MythicGenes.UniqId)).to.be.within(
        theoUniqId - MYTHIC_DISPERSION,
        theoUniqId + MYTHIC_DISPERSION,
      );
    });

    it('correct range for high UniqId', async () => {
      const theoUniqId = maxUniqId - 2;
      const _genome = await generateMythicGenome(genomesWithTheoUniqId(theoUniqId));
      const genome = new Genome(_genome);
      expect(genome.getGene(MythicGenes.UniqId)).to.be.within(
        theoUniqId - MYTHIC_DISPERSION,
        mythicAmount,
      );
    });

    it('correct range for max peruliarity', async () => {
      const theoUniqId = maxUniqId;
      const _genome = await generateMythicGenome(genomesWithTheoUniqId(theoUniqId));
      const genome = new Genome(_genome);
      expect(genome.getGene(MythicGenes.UniqId)).to.be.within(
        theoUniqId - MYTHIC_DISPERSION,
        mythicAmount,
      );
    });

    it('set correct magic number', async () => {
      const _genome = await generateMythicGenome(genomesWithTheoUniqId(0));
      const genome = new Genome(_genome);
      expect(genome.getGene(magicNumberGeneIdx)).to.equal(mythicMagicNumber);
    });
  });

  describe('mintPack', () => {
    beforeEach(async () => {
      await Game.setMintPrice(1);
    });

    it('duckies are paid for mint', async () => {
      const amount = 10;
      await expect(GameAsSomeone.mintPack(amount)).to.changeTokenBalance(
        Duckies,
        Someone,
        -amount * 10 ** (await Duckies.decimals()),
      );
    });

    it('correct amount of tokens is minted', async () => {
      const amount = 10;
      await GameAsSomeone.mintPack(amount);

      expect(await Ducklings.balanceOf(Someone.address)).to.equal(amount);
    });

    it('FLAG_TRANSFERABLE is set', async () => {
      const amount = 1;
      await GameAsSomeone.mintPack(amount);
      expect(await Ducklings.isTransferable(0)).to.be.true;
    });

    it('revert on amount == 0', async () => {
      const amount = 0;
      await expect(GameAsSomeone.mintPack(amount))
        .to.be.revertedWithCustomError(Game, 'MintingRulesViolated')
        .withArgs(Collections.Duckling, amount);
    });

    it('revert on amount > MAX_PACK_SIZE', async () => {
      const amount = MAX_PACK_SIZE + 1;
      await expect(GameAsSomeone.mintPack(amount))
        .to.be.revertedWithCustomError(Game, 'MintingRulesViolated')
        .withArgs(Collections.Duckling, amount);
    });

    it('event is emitted for every token', async () => {
      const amount = 10;
      const chainId = await ethers.provider.getNetwork().then((network) => network.chainId);
      const tx = await GameAsSomeone.mintPack(amount);
      const receipt = await tx.wait();
      const { timestamp } = await ethers.provider.getBlock(receipt.blockNumber);

      for (let i = 0; i < amount; i++) {
        await expect(tx)
          .to.emit(Ducklings, 'Minted')
          .withArgs(Someone.address, i, anyUint, timestamp, chainId);
      }
    });
  });
});
