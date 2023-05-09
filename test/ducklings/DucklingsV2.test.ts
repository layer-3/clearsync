import { ethers, upgrades } from 'hardhat';
import { constants, utils } from 'ethers';
import { expect } from 'chai';

import { connect } from '../helpers/connect';
import { ACCOUNT_MISSING_ROLE } from '../helpers/common';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { DucklingsV1, DucklingsV2 } from '../../typechain-types';

const ADMIN_ROLE = constants.HashZero;
const UPGRADER_ROLE = utils.id('UPGRADER_ROLE');
const MAINTAINER_ROLE = utils.id('MAINTAINER_ROLE');
const GAME_ROLE = utils.id('GAME_ROLE');

describe('DucklingsV2', () => {
  let Upgrader: SignerWithAddress;
  let Maintainer: SignerWithAddress;
  let Someone: SignerWithAddress;
  let Game: SignerWithAddress;

  let DucklingsV1: DucklingsV1;
  let DucklingsV2: DucklingsV2;
  let DucklingsV2AsSomeone: DucklingsV2;

  const upgrade = async (Ducklings: DucklingsV1): Promise<DucklingsV2> => {
    const DucklingsFactory = await ethers.getContractFactory('DucklingsV2');
    return upgrades.upgradeProxy(Ducklings.address, DucklingsFactory) as Promise<DucklingsV2>;
  };

  before(async () => {
    [Upgrader, Maintainer, Someone, Game] = await ethers.getSigners();
  });

  beforeEach(async () => {
    const DucklingsFactory = await ethers.getContractFactory('DucklingsV1');
    DucklingsV1 = (await upgrades.deployProxy(DucklingsFactory, [], {
      kind: 'uups',
    })) as DucklingsV1;
    await DucklingsV1.deployed();

    await DucklingsV1.grantRole(UPGRADER_ROLE, Upgrader.address);
    await DucklingsV1.grantRole(MAINTAINER_ROLE, Maintainer.address);
    await DucklingsV1.grantRole(GAME_ROLE, Game.address);

    DucklingsV2 = await upgrade(DucklingsV1);
    DucklingsV2AsSomeone = connect(DucklingsV2, Someone);
  });

  describe('destruct', () => {
    it('Admin can destruct', async () => {
      await expect(DucklingsV2.destruct()).to.not.be.reverted;
    });

    it('revert on any call to destructed Ducklings', async () => {
      await DucklingsV2.destruct();
      await expect(DucklingsV2.getRoyaltyCollector()).to.be.reverted;
    });

    it('revert on not Admin destruct', async () => {
      await expect(DucklingsV2AsSomeone.destruct()).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE),
      );
    });
  });

  describe('collection', () => {
    it('collection name changes after reinitialization', async () => {
      // reinitialize
      await DucklingsV2.initializeV2();

      expect(await DucklingsV2.name()).to.equal('(Deleted) Yellow Ducklings');
      expect(await DucklingsV2.name()).to.equal('(Deleted) Yellow Ducklings');
    });

    it('collection symbol changes after reinitialization', async () => {
      // reinitialize
      await DucklingsV2.initializeV2();

      expect(await DucklingsV2.symbol()).to.equal('DELETED');
    });
  });
});
