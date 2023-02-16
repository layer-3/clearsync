import { expect } from 'chai';
import { ethers, upgrades } from 'hardhat';
import { constants, utils } from 'ethers';

import { connectGroup } from '../connectContract';
import { randomBytes32 } from '../helpers/payload';

import { signBounty } from './signatures';

import type { Garden, TestERC20 } from '../../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { Bounty } from './bounty';
import { ACCOUNT_MISSING_ROLE } from '../helpers/common';

const INSUF_TOKEN_BALANCE = 'InsufficientTokenBalance';

const TOKEN_CAP = 100_000_000_000;
const GARDEN_DEPOSITED_DUCKIES = 10_000_000_000;
const GARDEN_DEPOSITED_PARTNER_TOKEN = 10_000_000_000;
const AMOUNT = 100;

const ADMIN_ROLE = constants.HashZero;
const UPGRADER_ROLE = utils.id('UPGRADER_ROLE');

describe('Garden', () => {
  let Duckies: TestERC20;
  let PartnerToken: TestERC20;
  let Garden: Garden;

  let GardenAdmin: SignerWithAddress;
  let Issuer: SignerWithAddress;
  let DuckiesAdmin: SignerWithAddress;
  let Someone: SignerWithAddress;
  let Someother: SignerWithAddress;

  let GardenAsAdmin: Garden;
  let GardenAsSomeone: Garden;

  const BountyBase: Bounty = {
    amount: AMOUNT,
    tokenAddress: '',
    beneficiary: '',
    isPaidToReferrers: false,
    referrer: constants.AddressZero,
    expire: 0,
    chainId: 31_337,
    bountyUID: '',
  };

  let SomeoneDuckiesBounty: Bounty;
  let SomeonePartnerTokenBounty: Bounty;

  before(async () => {
    [GardenAdmin, Issuer, DuckiesAdmin, Someone, Someother] = await ethers.getSigners();
  });

  beforeEach(async () => {
    const TestERC20Factory = await ethers.getContractFactory('TestERC20');
    Duckies = (await TestERC20Factory.deploy(
      'Duckies',
      'DUCKIES',
      TOKEN_CAP,
    )) as unknown as TestERC20;
    await Duckies.deployed();

    PartnerToken = (await TestERC20Factory.deploy(
      'Partner',
      'PARTNER',
      TOKEN_CAP,
    )) as unknown as TestERC20;
    await PartnerToken.deployed();

    const GardenFactory = await ethers.getContractFactory('Garden', GardenAdmin);
    Garden = (await upgrades.deployProxy(GardenFactory, [Duckies.address], {
      kind: 'uups',
    })) as unknown as Garden;
    await Garden.deployed();

    [GardenAsAdmin, GardenAsSomeone] = connectGroup(Garden, [GardenAdmin, Someone]);

    await GardenAsAdmin.setIssuer(Issuer.address);

    await Duckies.mint(Garden.address, GARDEN_DEPOSITED_DUCKIES);
    await PartnerToken.mint(Garden.address, GARDEN_DEPOSITED_PARTNER_TOKEN);

    BountyBase.expire = Math.round(Date.now() / 1000) + 600; // 10 mins from now
    BountyBase.bountyUID = randomBytes32();

    SomeoneDuckiesBounty = {
      ...BountyBase,
      beneficiary: Someone.address,
      tokenAddress: Duckies.address,
    };

    SomeonePartnerTokenBounty = {
      ...BountyBase,
      beneficiary: Someone.address,
      tokenAddress: PartnerToken.address,
    };
  });

  describe('initialize', () => {
    it('deployer is admin', async () => {
      expect(await Garden.hasRole(ADMIN_ROLE, GardenAdmin.address)).to.be.true;
    });

    it('deployer is upgrader', async () => {
      expect(await Garden.hasRole(UPGRADER_ROLE, GardenAdmin.address)).to.be.true;
    });

    it('issuer not set', async () => {
      const GardenFactory = await ethers.getContractFactory('Garden', GardenAdmin);
      Garden = (await upgrades.deployProxy(GardenFactory, [Duckies.address], {
        kind: 'uups',
      })) as unknown as Garden;
      await Garden.deployed();

      expect(await Garden.getIssuer()).to.equal(constants.AddressZero);
    });
  });

  describe('issuer', () => {
    it('admin can set issuer', async () => {
      await GardenAsAdmin.setIssuer(Someone.address);
      expect(await Garden.getIssuer()).to.equal(Someone.address);
    });

    it('revert on someone set issuer', async () => {
      await expect(GardenAsSomeone.setIssuer(Someother.address)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE),
      );
    });
  });

  describe('transferTokenBalanceToPartner', () => {
    it('admin can transfer token balance to partner', async () => {
      await GardenAsAdmin.transferTokenBalanceToPartner(PartnerToken.address, Someone.address);
      expect(await PartnerToken.balanceOf(Someone.address)).to.equal(
        GARDEN_DEPOSITED_PARTNER_TOKEN,
      );
      expect(await PartnerToken.balanceOf(Garden.address)).to.equal(0);
    });

    it('revert on someone transfer token balance to partner', async () => {
      await expect(
        GardenAsSomeone.transferTokenBalanceToPartner(PartnerToken.address, Someone.address),
      ).to.be.revertedWith(ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE));
    });

    it('revert on admin transfer partner token if partner token balance is 0', async () => {
      // withdraw PartnerToken
      await GardenAsAdmin.transferTokenBalanceToPartner(PartnerToken.address, Someone.address);

      await expect(
        GardenAsAdmin.transferTokenBalanceToPartner(PartnerToken.address, Someone.address),
      )
        .to.revertedWithCustomError(Garden, INSUF_TOKEN_BALANCE)
        .withArgs(PartnerToken.address, 1, 0);
    });
  });

  describe('payouts', () => {});

  describe('halving', () => {});

  describe('claim bounty', () => {
    it('successfuly claim bounty in Duckies', async () => {
      await GardenAsSomeone.claimBounty(
        SomeoneDuckiesBounty,
        signBounty(SomeoneDuckiesBounty, Issuer),
      );
      expect(await Duckies.balanceOf(Someone.address)).to.equal(AMOUNT);
    });

    it('successfuly claim bounty in Partner Token', async () => {
      await GardenAsSomeone.claimBounty(
        SomeonePartnerTokenBounty,
        signBounty(SomeonePartnerTokenBounty, Issuer),
      );
      expect(await PartnerToken.balanceOf(Someone.address)).to.equal(AMOUNT);
    });
  });

  describe('claim bounties', () => {});

  // TODO: add events
});
