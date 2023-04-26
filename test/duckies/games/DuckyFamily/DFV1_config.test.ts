import { expect } from 'chai';
import { utils } from 'ethers';

import { ACCOUNT_MISSING_ROLE } from '../../../helpers/common';

import { ADMIN_ROLE, MAINTAINER_ROLE, setup } from './setup';
import {
  Collections,
  collectionsGeneDistributionTypes,
  collectionsGeneValuesNum,
  mythicAmount,
} from './config';

import type { DuckyFamilyV1, TESTDuckyFamilyV1, YellowToken } from '../../../../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

describe('DuckyFamilyV1 config', () => {
  let Someone: SignerWithAddress;
  let Game: TESTDuckyFamilyV1;
  let Duckies: YellowToken;

  let GameAsMaintainer: DuckyFamilyV1;
  let GameAsSomeone: DuckyFamilyV1;

  let duckiesDecMultiplies: number;

  beforeEach(async () => {
    ({ Someone, Game, GameAsMaintainer, GameAsSomeone, Duckies } = await setup());
    duckiesDecMultiplies = 10 ** (await Duckies.decimals());
  });

  type MeldPrices = [number, number, number, number];

  const MINT_PRICE = 50;
  const MELD_PRICES = [100, 200, 500, 1000];

  const CUSTOM_MINT_PRICE = 5;
  const CUSTOM_MELD_PRICES = [10, 20, 50, 100];

  describe('setPepper', () => {
    const pepper = utils.id('pepper');

    it('maintainer can set pepper', async () => {
      await GameAsMaintainer.setPepper(pepper);
    });

    it('revert on not maintainer set pepper', async () => {
      await expect(GameAsSomeone.setPepper(pepper)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, MAINTAINER_ROLE),
      );
    });
  });

  describe('mintPrice', () => {
    it('returns correct value', async () => {
      expect(await Game.getMintPrice()).to.deep.equal(MINT_PRICE * duckiesDecMultiplies);
    });
  });

  describe('setMintPrice', () => {
    it('maintainer can set mint price', async () => {
      await GameAsMaintainer.setMintPrice(CUSTOM_MINT_PRICE);
      expect(await Game.getMintPrice()).to.deep.equal(CUSTOM_MINT_PRICE * duckiesDecMultiplies);
    });

    it('revert on not maintainer set mint price', async () => {
      await expect(GameAsSomeone.setMintPrice(MINT_PRICE)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, MAINTAINER_ROLE),
      );
    });
  });

  describe('getMeldPrices', () => {
    it('returns correct values', async () => {
      const duckiesDecMultiplier = 10 ** (await Duckies.decimals());
      expect(await Game.getMeldPrices()).to.deep.equal(
        MELD_PRICES.map((v) => v * duckiesDecMultiplier),
      );
    });
  });

  describe('setMeldPrices', () => {
    it('maintainer can set meld price', async () => {
      await GameAsMaintainer.setMeldPrices(CUSTOM_MELD_PRICES as MeldPrices);
      const duckiesDecMultiplier = 10 ** (await Duckies.decimals());
      expect(await Game.getMeldPrices()).to.deep.equal(
        CUSTOM_MELD_PRICES.map((v) => v * duckiesDecMultiplier),
      );
    });

    it('revert on not maintainer set meld price', async () => {
      await expect(GameAsSomeone.setMeldPrices(MELD_PRICES as MeldPrices)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, MAINTAINER_ROLE),
      );
    });
  });

  describe('getCollectionsGeneValues', () => {
    it('returns correct values', async () => {
      const [contractCollectionsGeneValues, contractMythicAmount] =
        await Game.getCollectionsGeneValues();
      expect(contractCollectionsGeneValues).to.deep.equal(collectionsGeneValuesNum);
      expect(contractMythicAmount).to.deep.equal(mythicAmount);
    });
  });

  describe('getCollectionsGeneDistributionTypes', () => {
    it('returns correct values', async () => {
      const contractGeneDistrTypes = await Game.getCollectionsGeneDistributionTypes();
      expect(contractGeneDistrTypes).to.deep.equal(collectionsGeneDistributionTypes);
    });
  });

  describe('setDucklingGeneValues', () => {
    it('admin can set gene values', async () => {
      const newDucklingGeneValues = [1, 2, 3, 4, 42];
      await Game.setDucklingGeneValues(newDucklingGeneValues);
      const [contractGeneValues] = await Game.getCollectionsGeneValues();
      expect(contractGeneValues[Collections.Duckling]).to.deep.equal(newDucklingGeneValues);
    });

    it('revert on not admin set gene values', async () => {
      const newDucklingGeneValues = [1, 2, 3, 4, 42];
      await expect(GameAsSomeone.setDucklingGeneValues(newDucklingGeneValues)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE),
      );
    });
  });

  describe('setDucklingGeneDistributionTypes', () => {
    it('admin can set gene distribution types', async () => {
      const newDucklingGeneDistrTypes = 42;
      await Game.setDucklingGeneDistributionTypes(newDucklingGeneDistrTypes);
      const contractGeneDistrTypes = await Game.getCollectionsGeneDistributionTypes();
      expect(contractGeneDistrTypes[Collections.Duckling]).to.equal(newDucklingGeneDistrTypes);
    });

    it('revert on not admin set gene distribution types', async () => {
      const newDucklingGeneDistrTypes = 42;
      await expect(
        GameAsSomeone.setDucklingGeneDistributionTypes(newDucklingGeneDistrTypes),
      ).to.be.revertedWith(ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE));
    });
  });

  describe('setZombeakGeneValues', () => {
    it('admin can set gene values', async () => {
      const newZombeakGeneValues = [1, 2, 3, 4, 42];
      await Game.setZombeakGeneValues(newZombeakGeneValues);
      const [contractGeneValues] = await Game.getCollectionsGeneValues();
      expect(contractGeneValues[Collections.Zombeak]).to.deep.equal(newZombeakGeneValues);
    });

    it('revert on not admin set gene values', async () => {
      const newZombeakGeneValues = [1, 2, 3, 4, 42];
      await expect(GameAsSomeone.setZombeakGeneValues(newZombeakGeneValues)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE),
      );
    });
  });

  describe('setZombeakGeneDistributionTypes', () => {
    it('admin can set gene distribution types', async () => {
      const newZombeakGeneDistrTypes = 42;
      await Game.setZombeakGeneDistributionTypes(newZombeakGeneDistrTypes);
      const contractGeneDistrTypes = await Game.getCollectionsGeneDistributionTypes();
      expect(contractGeneDistrTypes[Collections.Zombeak]).to.equal(newZombeakGeneDistrTypes);
    });

    it('revert on not admin set gene distribution types', async () => {
      const newZombeakGeneDistrTypes = 42;
      await expect(
        GameAsSomeone.setZombeakGeneDistributionTypes(newZombeakGeneDistrTypes),
      ).to.be.revertedWith(ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE));
    });
  });

  describe('setMythicAmount', () => {
    it('admin can set mythic amount', async () => {
      const newMythicAmount = 42;
      await Game.setMythicAmount(newMythicAmount);
      const [, contractMythicAmount] = await Game.getCollectionsGeneValues();
      expect(contractMythicAmount).to.equal(newMythicAmount);
    });

    it('revert on not admin set mythic amount', async () => {
      const newMythicAmount = 42;
      await expect(GameAsSomeone.setMythicAmount(newMythicAmount)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE),
      );
    });
  });

  describe('setMythicGeneValues', () => {
    it('admin can set gene values', async () => {
      const newMythicGeneValues = [1, 2, 3, 4, 42];
      await Game.setMythicGeneValues(newMythicGeneValues);
      const [contractGeneValues] = await Game.getCollectionsGeneValues();
      expect(contractGeneValues[Collections.Mythic]).to.deep.equal(newMythicGeneValues);
    });

    it('revert on not admin set gene values', async () => {
      const newMythicGeneValues = [1, 2, 3, 4, 42];
      await expect(GameAsSomeone.setMythicGeneValues(newMythicGeneValues)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE),
      );
    });
  });

  describe('setMythicGeneDistributionTypes', () => {
    it('admin can set gene distribution types', async () => {
      const newMythicGeneDistrTypes = 42;
      await Game.setMythicGeneDistributionTypes(newMythicGeneDistrTypes);
      const contractGeneDistrTypes = await Game.getCollectionsGeneDistributionTypes();
      expect(contractGeneDistrTypes[Collections.Mythic]).to.equal(newMythicGeneDistrTypes);
    });

    it('revert on not admin set gene distribution types', async () => {
      const newMythicGeneDistrTypes = 42;
      await expect(
        GameAsSomeone.setMythicGeneDistributionTypes(newMythicGeneDistrTypes),
      ).to.be.revertedWith(ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE));
    });
  });
});
