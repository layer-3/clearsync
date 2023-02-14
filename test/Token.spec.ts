import { expect } from 'chai';
import { loadFixture, time } from '@nomicfoundation/hardhat-network-helpers';
import { ethers } from 'hardhat';
import { utils } from 'ethers';

import { connectGroup } from './helpers/connectContract';
import { ACCOUNT_MISSING_ROLE } from './helpers/common';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { Token } from '../typechain-types';

const ADMIN_ROLE = ethers.constants.HashZero;
const MINTER_ROLE = utils.id('MINTER_ROLE');
const COMPLIANCE_ROLE = utils.id('COMPLIANCE_ROLE');
const BLACKLISTED_ROLE = utils.id('BLACKLISTED_ROLE');

const DECIMALS = 8;
const TOKEN_SUPPLY_CAP = 10_000_000_000;
const PREMINT = 10_000_000_000;

const ACTIVATED_EVENT = 'Activated';
const BLACKLISTED_EVENT = 'Blacklisted';
const BLACKLISTED_REMOVED_EVENT = 'BlacklistedRemoved';
const BLACKLISTED_BURNT_EVENT = 'BlacklistedBurnt';

// activate
const ALREADY_ACTIVATED = 'Already activated';
const ZERO_PREMINT = 'Zero premint';
const PREMINT_EXCEEDS_CAP = 'Premint exceeds cap';

// mint
const NOT_ACTIVATED = 'Not activated';
const MINT_EXCEEDS_CAP = 'Mint exceeds cap';

// blacklist
const IS_BLACKLISTED = 'Account is blacklisted';

// erc20
const INSUFFICIENT_ALLOWANCE = 'ERC20: insufficient allowance';

describe('Token', function () {
  let TokenAdmin: SignerWithAddress;
  let Minter: SignerWithAddress;
  let Compliance: SignerWithAddress;
  let Blacklisted: SignerWithAddress;
  let User: SignerWithAddress;
  let Someone: SignerWithAddress;
  let PremintReceiver: SignerWithAddress;

  let Token: Token;

  let TokenAsAdmin: Token;
  let TokenAsMinter: Token;
  let TokenAsCompliance: Token;
  let TokenAsBlacklisted: Token;
  let TokenAsUser: Token;
  let TokenAsSomeone: Token;

  async function deployTokenFixture(): Promise<{ Token: Token }> {
    const TokenFactory = await ethers.getContractFactory('Token');
    const Token = (await TokenFactory.connect(TokenAdmin).deploy(
      'Canary',
      'CANARY',
      TOKEN_SUPPLY_CAP,
    )) as Token;

    return { Token };
  }

  before(async () => {
    [TokenAdmin, Minter, Compliance, Blacklisted, User, Someone, PremintReceiver] =
      await ethers.getSigners();
  });

  beforeEach(async () => {
    ({ Token } = await loadFixture(deployTokenFixture));

    await Token.grantRole(MINTER_ROLE, Minter.address);
    await Token.grantRole(COMPLIANCE_ROLE, Compliance.address);

    [
      TokenAsAdmin,
      TokenAsMinter,
      TokenAsCompliance,
      TokenAsBlacklisted,
      TokenAsUser,
      TokenAsSomeone,
    ] = connectGroup(Token, [TokenAdmin, Minter, Compliance, Blacklisted, User, Someone]);
  });

  describe('Deployment', () => {
    it('Correct name and symbol', async () => {
      expect(await Token.name()).to.equal('Canary');
      expect(await Token.symbol()).to.equal('CANARY');
    });

    it('Correct decimals', async () => {
      expect(await Token.decimals()).to.equal(DECIMALS);
    });

    it('Correct supply cap', async () => {
      expect(await Token.cap()).to.equal(TOKEN_SUPPLY_CAP);
    });

    it('activatedAt == 0', async () => {
      expect(await Token.activatedAt()).to.equal(0);
    });

    it('Deployer granted roles', async () => {
      expect(await Token.hasRole(ADMIN_ROLE, TokenAdmin.address)).to.be.true;
      expect(await Token.hasRole(MINTER_ROLE, TokenAdmin.address)).to.be.true;
      expect(await Token.hasRole(COMPLIANCE_ROLE, TokenAdmin.address)).to.be.true;
    });
  });

  describe('Activation', () => {
    it('Admin can activate', async () => {
      await TokenAsAdmin.activate(PREMINT, PremintReceiver.address);
      expect(await Token.balanceOf(PremintReceiver.address)).to.equal(PREMINT);
    });

    it('Activation time is saved', async () => {
      await TokenAsAdmin.activate(PREMINT, PremintReceiver.address);
      const activatedAt = await time.latest();
      expect(await Token.activatedAt()).to.equal(activatedAt);
    });

    it('Revert on not admin activate', async () => {
      await expect(TokenAsUser.activate(PREMINT, PremintReceiver.address)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(User.address, ADMIN_ROLE),
      );
    });

    it('Revert on second activation', async () => {
      await TokenAsAdmin.activate(PREMINT, PremintReceiver.address);
      const activatedAt = await time.latest();
      expect(await Token.activatedAt()).to.equal(activatedAt);

      await expect(TokenAsAdmin.activate(PREMINT, PremintReceiver.address)).to.be.revertedWith(
        ALREADY_ACTIVATED,
      );
    });

    it('Revert on premint 0', async () => {
      await expect(TokenAsAdmin.activate(0, PremintReceiver.address)).to.be.revertedWith(
        ZERO_PREMINT,
      );
    });

    it('Revert on premint > cap', async () => {
      await expect(
        TokenAsAdmin.activate(TOKEN_SUPPLY_CAP + 1, PremintReceiver.address),
      ).to.be.revertedWith(PREMINT_EXCEEDS_CAP);
    });

    it('Emit event', async () => {
      await expect(TokenAsAdmin.activate(PREMINT, PremintReceiver.address))
        .to.emit(Token, ACTIVATED_EVENT)
        .withArgs(PREMINT);
    });
  });

  describe('Mint', () => {
    const mintAmount = 100;

    beforeEach(async () => {
      // NOTE: premint is < supply cap here to allow manual minting further in tests
      await TokenAsAdmin.activate(1, PremintReceiver.address);
    });

    it('Minter can mint', async () => {
      await TokenAsMinter.mint(User.address, mintAmount);
      expect(await Token.balanceOf(User.address)).to.equal(mintAmount);
    });

    it('Revert on not admin activate', async () => {
      await expect(TokenAsUser.mint(Someone.address, mintAmount)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(User.address, MINTER_ROLE),
      );
    });

    it('Revert on not activated', async () => {
      const { Token: NotActivatedToken } = await loadFixture(deployTokenFixture);

      await expect(
        NotActivatedToken.connect(TokenAdmin).mint(User.address, mintAmount),
      ).to.be.revertedWith(NOT_ACTIVATED);
    });

    it('Revert if exceeds supply cap', async () => {
      await expect(TokenAsMinter.mint(Someone.address, TOKEN_SUPPLY_CAP + 1)).to.be.revertedWith(
        MINT_EXCEEDS_CAP,
      );
    });
  });

  describe('Burn', () => {
    const amount = 1000;

    beforeEach(async () => {
      await TokenAsAdmin.activate(1, PremintReceiver.address);
      await TokenAsAdmin.mint(User.address, amount);
    });

    it('User can burn their funds', async () => {
      await TokenAsUser.burn(amount);
      expect(await Token.balanceOf(User.address)).to.equal(0);
    });

    it('Revert on burning others funds', async () => {
      await expect(TokenAsSomeone.burnFrom(User.address, amount)).to.be.revertedWith(
        INSUFFICIENT_ALLOWANCE,
      );
    });

    it('User can burn others funds if allowance got', async () => {
      await TokenAsUser.approve(Someone.address, amount);
      await TokenAsSomeone.burnFrom(User.address, amount);
      expect(await Token.balanceOf(User.address)).to.equal(0);
    });
  });

  describe('Blacklist', () => {
    const amount = 100;

    beforeEach(async () => {
      // NOTE: premint is < supply cap here to allow manual minting further in tests
      await TokenAsAdmin.activate(1, PremintReceiver.address);
      await TokenAsCompliance.blacklist(Blacklisted.address);
    });

    it('Compliance can blacklist', async () => {
      await TokenAsCompliance.blacklist(User.address);
      expect(await Token.hasRole(BLACKLISTED_ROLE, User.address)).to.be.true;
    });

    it('Revert on not compliance blacklist', async () => {
      await expect(TokenAsUser.blacklist(User.address)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(User.address, COMPLIANCE_ROLE),
      );
    });

    it('Compliance can unblock blacklisted', async () => {
      await TokenAsCompliance.removeBlacklisted(Blacklisted.address);
      expect(await Token.hasRole(BLACKLISTED_ROLE, Blacklisted.address)).to.be.false;
    });

    it('Revern on not compliance unblock blacklisted', async () => {
      await expect(TokenAsUser.removeBlacklisted(Blacklisted.address)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(User.address, COMPLIANCE_ROLE),
      );
    });

    it('Blacklisted can not transfer', async () => {
      await expect(TokenAsBlacklisted.transfer(User.address, amount)).to.be.revertedWith(
        IS_BLACKLISTED,
      );
    });

    it('Blacklisted can not be transfered from', async () => {
      await expect(
        TokenAsUser.transferFrom(Blacklisted.address, User.address, amount),
      ).to.be.revertedWith(IS_BLACKLISTED);
    });

    it('Blacklisted funds can be burnt by compliance', async () => {
      await TokenAsAdmin.mint(Blacklisted.address, amount);
      expect(await Token.balanceOf(Blacklisted.address)).to.equal(amount);

      await TokenAsCompliance.burnBlacklisted(Blacklisted.address);
      expect(await Token.balanceOf(Blacklisted.address)).to.equal(0);
    });

    it('Revert on blacklisted funds burnt by not compliance', async () => {
      await expect(TokenAsUser.burnBlacklisted(Blacklisted.address)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(User.address, COMPLIANCE_ROLE),
      );
    });

    it('Blacklist emit event', async () => {
      await expect(TokenAsCompliance.blacklist(User.address))
        .to.emit(Token, BLACKLISTED_EVENT)
        .withArgs(User.address);
    });

    it('Remove blacklisted emit event', async () => {
      await expect(TokenAsCompliance.removeBlacklisted(Blacklisted.address))
        .to.emit(Token, BLACKLISTED_REMOVED_EVENT)
        .withArgs(Blacklisted.address);
    });

    it('Burn blacklisted emit event', async () => {
      await TokenAsAdmin.mint(Blacklisted.address, amount);

      await expect(TokenAsCompliance.burnBlacklisted(Blacklisted.address))
        .to.emit(Token, BLACKLISTED_BURNT_EVENT)
        .withArgs(Blacklisted.address, amount);
    });
  });

  describe('Transfer', () => {
    const amount = 100;

    beforeEach(async () => {
      // NOTE: premint is < supply cap here to allow manual minting further in tests
      await TokenAsAdmin.activate(1, PremintReceiver.address);
    });

    it('can successfully transfer', async () => {
      await TokenAsAdmin.mint(User.address, amount);

      await TokenAsUser.transfer(Someone.address, amount);
      expect(await Token.balanceOf(Someone.address)).to.equal(amount);
    });

    it('can successfully transfer from', async () => {
      await TokenAsAdmin.mint(User.address, amount);
      await TokenAsUser.approve(Someone.address, amount);

      await TokenAsSomeone.transferFrom(User.address, Someone.address, amount);
      expect(await Token.balanceOf(Someone.address)).to.equal(amount);
    });
  });
});
