import { ethers, upgrades } from 'hardhat';
import { constants, utils } from 'ethers';
import { expect } from 'chai';

import { connect } from '../helpers/connect';
import { ACCOUNT_MISSING_ROLE } from '../helpers/common';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { OldDucklingsV1, OldDucklingsV2 } from '../../typechain-types';

const ADMIN_ROLE = constants.HashZero;
const UPGRADER_ROLE = utils.id('UPGRADER_ROLE');
const MAINTAINER_ROLE = utils.id('MAINTAINER_ROLE');
const GAME_ROLE = utils.id('GAME_ROLE');

describe('OldDucklingsV2', () => {
  let Upgrader: SignerWithAddress;
  let Maintainer: SignerWithAddress;
  let Someone: SignerWithAddress;
  let Game: SignerWithAddress;

  let OldDucklingsV1: OldDucklingsV1;
  let OldDucklingsV2: OldDucklingsV2;
  let OldDucklingsV2AsSomeone: OldDucklingsV2;

  const upgrade = async (Ducklings: OldDucklingsV1): Promise<OldDucklingsV2> => {
    const DucklingsFactory = await ethers.getContractFactory('OldDucklingsV2');
    return upgrades.upgradeProxy(Ducklings.address, DucklingsFactory) as Promise<OldDucklingsV2>;
  };

  before(async () => {
    [Upgrader, Maintainer, Someone, Game] = await ethers.getSigners();
  });

  beforeEach(async () => {
    const OldDucklingsFactory = await ethers.getContractFactory('OldDucklingsV1');
    OldDucklingsV1 = (await upgrades.deployProxy(OldDucklingsFactory, [], {
      kind: 'uups',
    })) as OldDucklingsV1;
    await OldDucklingsV1.deployed();

    await OldDucklingsV1.grantRole(UPGRADER_ROLE, Upgrader.address);
    await OldDucklingsV1.grantRole(MAINTAINER_ROLE, Maintainer.address);
    await OldDucklingsV1.grantRole(GAME_ROLE, Game.address);

    OldDucklingsV2 = await upgrade(OldDucklingsV1);
    OldDucklingsV2AsSomeone = connect(OldDucklingsV2, Someone);
  });

  describe('destruct', () => {
    it('Admin can destruct', async () => {
      await expect(OldDucklingsV2.destruct()).to.not.be.reverted;
    });

    it('revert on any call to destructed Ducklings', async () => {
      await OldDucklingsV2.destruct();
      await expect(OldDucklingsV2.getRoyaltyCollector()).to.be.reverted;
    });

    it('revert on not Admin destruct', async () => {
      await expect(OldDucklingsV2AsSomeone.destruct()).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE),
      );
    });
  });

  describe('collection', () => {
    it('collection name changes after reinitialization', async () => {
      // reinitialize
      await OldDucklingsV2.initializeV2();

      expect(await OldDucklingsV2.name()).to.equal('(Deleted) Yellow Ducklings');
      expect(await OldDucklingsV2.name()).to.equal('(Deleted) Yellow Ducklings');
    });

    it('collection symbol changes after reinitialization', async () => {
      // reinitialize
      await OldDucklingsV2.initializeV2();

      expect(await OldDucklingsV2.symbol()).to.equal('DELETED');
    });
  });
});
