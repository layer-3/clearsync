import { expect } from 'chai';

import { ACCOUNT_MISSING_ROLE } from '../../../helpers/common';

import { MAINTAINER_ROLE, setup } from './setup';

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

  describe('mintPrice', () => {
    it('returns correct value', async () => {
      expect(await Game.mintPrice()).to.deep.equal(MINT_PRICE * duckiesDecMultiplies);
    });

    describe('setMintPrice', () => {
      it('maintainer can set mint price', async () => {
        await GameAsMaintainer.setMintPrice(CUSTOM_MINT_PRICE);
        expect(await Game.mintPrice()).to.deep.equal(CUSTOM_MINT_PRICE * duckiesDecMultiplies);
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
      it('returns correct values');
    });

    describe('getCollectionsGeneDistributionTypes', () => {
      it('returns correct values');
    });

    describe('setDucklingGeneValues', () => {
      it('admin can set gene values');

      it('revert on not admin set gene values');
    });

    describe('setDucklingGeneDistributionTypes', () => {
      it('admin can set gene distribution types');

      it('revert on not admin set gene distribution types');
    });

    describe('setZombeakGeneValues', () => {
      it('admin can set gene values');

      it('revert on not admin set gene values');
    });

    describe('setZombeakGeneDistributionTypes', () => {
      it('admin can set gene distribution types');

      it('revert on not admin set gene distribution types');
    });

    describe('setMythicAmount', () => {
      it('admin can set mythic amount');

      it('revert on not admin set mythic amount');
    });
  });
});
