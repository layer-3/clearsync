import { expect } from 'chai';
import { ethers, upgrades } from 'hardhat';
import { constants } from 'ethers';

import { connectGroup } from '../connectContract';
import { randomBytes32 } from '../helpers/payload';

import { signBounty } from './signatures';

import type { Garden, TestERC20 } from '../../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { Bounty } from './bounty';

const TOKEN_CAP = 100_000_000_000;
const GARDEN_DEPOSITED_DUCKIES = 10_000_000_000;
const GARDEN_DEPOSITED_PARTNER_TOKEN = 10_000_000_000;
const AMOUNT = 100;

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

  describe('initialize');

  describe('issuer');

  describe('transferTokenBalanceToPartner');

  describe('payouts');

  describe('halving');

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

  describe('claim bounties');
});
