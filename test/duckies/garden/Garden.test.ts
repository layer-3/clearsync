import { expect } from 'chai';
import { ethers, upgrades } from 'hardhat';
import { constants, utils } from 'ethers';

import { connectGroup } from '../../helpers/connect';
import { ACCOUNT_MISSING_ROLE, randomBytes32 } from '../../helpers/common';
import {
  RewardParams,
  VoucherAction,
  encodeRewardParams,
} from '../../../src/contracts/garden/garden';
import { Voucher, signVoucher, signVouchers } from '../../../src/contracts/garden/voucher';

import type { Garden, TestERC20 } from '../../../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

const CIRCULAR_REFS = 'CircularReferrers';
const VOUCHER_USED = 'VoucherAlreadyUsed';
const INVALID_VOUCHER = 'InvalidVoucher';
const INVALID_REWARD = 'InvalidRewardParams';
const INSUF_TOKEN_BALANCE = 'InsufficientTokenBalance';
const INCORRECT_SIGNER = 'IncorrectSigner';

const TOKEN_CAP = 100_000_000_000;
const TOKEN_DEPOSITED_TO_GARDEN = 10_000_000_000;
const AMOUNT = 100;

const REFERRAL_PAYOUT_DIVIDER = 100;
const COMMISSIONS = [50, 40, 30, 20, 10];

const ADMIN_ROLE = constants.HashZero;
const UPGRADER_ROLE = utils.id('UPGRADER_ROLE');

describe('Garden', () => {
  let Token: TestERC20;
  let Garden: Garden;

  let GardenAdmin: SignerWithAddress;
  let Issuer: SignerWithAddress;
  let Someone: SignerWithAddress;
  let Someother: SignerWithAddress;
  let Referrer: SignerWithAddress;

  let GardenAsAdmin: Garden;
  let GardenAsSomeone: Garden;

  const VoucherBase: Voucher = {
    target: '',
    action: 0,
    beneficiary: '',
    referrer: constants.AddressZero,
    expire: 0,
    chainId: 31_337,
    voucherCodeHash: randomBytes32(),
    encodedParams: '0x',
  };

  let someoneVoucher: Voucher;

  before(async () => {
    [GardenAdmin, Issuer, Someone, Someother, Referrer] = await ethers.getSigners();
  });

  beforeEach(async () => {
    const TestERC20Factory = await ethers.getContractFactory('TestERC20');
    Token = (await TestERC20Factory.deploy(
      'Partner',
      'PARTNER',
      TOKEN_CAP,
    )) as unknown as TestERC20;
    await Token.deployed();

    const GardenFactory = await ethers.getContractFactory('Garden', GardenAdmin);
    Garden = (await upgrades.deployProxy(GardenFactory, [], {
      kind: 'uups',
    })) as unknown as Garden;
    await Garden.deployed();

    [GardenAsAdmin, GardenAsSomeone] = connectGroup(Garden, [GardenAdmin, Someone]);

    await GardenAsAdmin.setIssuer(Issuer.address);

    await Token.mint(Garden.address, TOKEN_DEPOSITED_TO_GARDEN);

    VoucherBase.target = Garden.address;
    VoucherBase.expire = Math.round(Date.now() / 1000) + 600; // 10 mins from now

    someoneVoucher = {
      ...VoucherBase,
      beneficiary: Someone.address,
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
      Garden = (await upgrades.deployProxy(GardenFactory, [], {
        kind: 'uups',
      })) as unknown as Garden;
      await Garden.deployed();

      expect(await Garden.issuer()).to.equal(constants.AddressZero);
    });
  });

  describe('issuer', () => {
    it('admin can set issuer', async () => {
      await GardenAsAdmin.setIssuer(Someone.address);
      expect(await Garden.issuer()).to.equal(Someone.address);
    });

    it('revert on someone set issuer', async () => {
      await expect(GardenAsSomeone.setIssuer(Someother.address)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE),
      );
    });
  });

  describe('transferTokenBalanceToPartner', () => {
    it('admin can transfer token balance to partner', async () => {
      await GardenAsAdmin.transferTokenBalanceToPartner(Token.address, Someone.address);
      expect(await Token.balanceOf(Someone.address)).to.equal(TOKEN_DEPOSITED_TO_GARDEN);
      expect(await Token.balanceOf(Garden.address)).to.equal(0);
    });

    it('revert on someone transfer token balance to partner', async () => {
      await expect(
        GardenAsSomeone.transferTokenBalanceToPartner(Token.address, Someone.address),
      ).to.be.revertedWith(ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE));
    });

    it('revert on admin transfer partner token if partner token balance is 0', async () => {
      // withdraw Token
      await GardenAsAdmin.transferTokenBalanceToPartner(Token.address, Someone.address);

      await expect(GardenAsAdmin.transferTokenBalanceToPartner(Token.address, Someone.address))
        .to.revertedWithCustomError(Garden, INSUF_TOKEN_BALANCE)
        .withArgs(Token.address, 1, 0);
    });
  });

  describe('use voucher', () => {
    describe('reward', () => {
      let rewardParams: RewardParams;
      let tokenRewardVoucher: Voucher;
      let someoneTokenRewardVoucher: Voucher;
      let voucherSig: string;

      beforeEach(async () => {
        rewardParams = {
          token: Token.address,
          amount: AMOUNT,
          commissions: COMMISSIONS,
        };

        tokenRewardVoucher = {
          ...VoucherBase,
          encodedParams: encodeRewardParams(rewardParams),
        };

        someoneTokenRewardVoucher = {
          ...tokenRewardVoucher,
          beneficiary: Someone.address,
        };

        voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);
      });

      describe('success', () => {
        it('successfully transfer token', async () => {
          expect(await Token.balanceOf(Someone.address)).to.equal(0);
          await GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig);
          expect(await Token.balanceOf(Someone.address)).to.equal(AMOUNT);
        });

        it('successfully register referrer & pay commissions', async () => {
          someoneTokenRewardVoucher.referrer = Referrer.address;
          voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);

          await GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig);
          expect(await Token.balanceOf(Referrer.address)).to.equal(
            (AMOUNT * COMMISSIONS[0]) / REFERRAL_PAYOUT_DIVIDER,
          );
        });

        it('paid only to existing referrers even if commission is set to more levels', async () => {
          someoneTokenRewardVoucher.referrer = Referrer.address;
          voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);

          await GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig);

          // payed = to SOMEONE + to REFERRER
          const payed = AMOUNT + (AMOUNT * COMMISSIONS[0]) / REFERRAL_PAYOUT_DIVIDER;
          expect(await Token.balanceOf(Garden.address)).to.equal(TOKEN_DEPOSITED_TO_GARDEN - payed);
        });

        it('event emitted on successfully used voucher', async () => {
          await expect(GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig))
            .to.emit(Garden, 'VoucherUsed')
            .withArgs(
              Someone.address,
              VoucherAction.Reward,
              someoneTokenRewardVoucher.voucherCodeHash,
              someoneTokenRewardVoucher.chainId,
            );
        });

        it('even emitted on referrer register', async () => {
          someoneTokenRewardVoucher.referrer = Referrer.address;
          voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);

          await expect(GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig))
            .to.emit(Garden, 'AffiliateRegistered')
            .withArgs(Someone.address, Referrer.address);
        });
      });

      describe('revert', () => {
        it('insufficient reward token balance', async () => {
          // withdraw Token balance
          await GardenAsAdmin.transferTokenBalanceToPartner(Token.address, Someone.address);

          await expect(GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig))
            .to.be.revertedWithCustomError(Garden, INSUF_TOKEN_BALANCE)
            .withArgs(Token.address, AMOUNT, 0);
        });

        describe('invalid voucher', () => {
          it('revert on signed by not issuer', async () => {
            voucherSig = await signVoucher(someoneTokenRewardVoucher, Someone);

            await expect(GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig))
              .to.be.revertedWithCustomError(Garden, INCORRECT_SIGNER)
              .withArgs(Issuer.address, Someone.address);
          });

          it('revert on target not Garden address', async () => {
            someoneTokenRewardVoucher.target = Someone.address;
            voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);

            await expect(
              GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig),
            ).to.be.revertedWithCustomError(Garden, INVALID_VOUCHER);
            //   .withArgs(someoneTokenRewardVoucher); <- seems like ethers can not parse JS Voucher properly
          });

          it('revert on action not in Actions', async () => {
            someoneTokenRewardVoucher.action = 42;
            voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);
            await expect(
              GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig),
            ).to.be.revertedWithCustomError(Garden, INVALID_VOUCHER);
          });

          it('revert on beneficiary not caller', async () => {
            someoneTokenRewardVoucher.beneficiary = Someother.address;
            voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);

            await expect(
              GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig),
            ).to.be.revertedWithCustomError(Garden, INVALID_VOUCHER);
          });

          it('revert on 1st level circular referrer', async () => {
            someoneTokenRewardVoucher.referrer = Someone.address;
            voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);

            await expect(GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig))
              .to.be.revertedWithCustomError(Garden, CIRCULAR_REFS)
              .withArgs(Someone.address, Someone.address);
          });

          it('revert on 3st level circular referrer', async () => {
            //         3            1           2
            // Someone -> Someother -> Referrer -> Someone

            // 1
            tokenRewardVoucher.beneficiary = Someother.address;
            tokenRewardVoucher.referrer = Referrer.address;
            voucherSig = await signVoucher(tokenRewardVoucher, Issuer);
            await Garden.connect(Someother).useVoucher(tokenRewardVoucher, voucherSig);

            // 2
            tokenRewardVoucher.beneficiary = Referrer.address;
            tokenRewardVoucher.referrer = Someone.address;
            tokenRewardVoucher.voucherCodeHash = randomBytes32();
            voucherSig = await signVoucher(tokenRewardVoucher, Issuer);
            await Garden.connect(Referrer).useVoucher(tokenRewardVoucher, voucherSig);

            // 3
            tokenRewardVoucher.beneficiary = Someone.address;
            tokenRewardVoucher.referrer = Someother.address;
            tokenRewardVoucher.voucherCodeHash = randomBytes32();
            voucherSig = await signVoucher(tokenRewardVoucher, Issuer);
            await expect(GardenAsSomeone.useVoucher(tokenRewardVoucher, voucherSig))
              .to.be.revertedWithCustomError(Garden, CIRCULAR_REFS)
              .withArgs(Someone.address, Someother.address);
          });

          it('revert on 5st level circular referrer', async () => {
            //         5            1              2         3           4
            // Someone -> Someother -> GardenAdmin -> Issuer -> Referrer -> Someone

            // 1
            tokenRewardVoucher.beneficiary = Someother.address;
            tokenRewardVoucher.referrer = GardenAdmin.address;
            voucherSig = await signVoucher(tokenRewardVoucher, Issuer);
            await Garden.connect(Someother).useVoucher(tokenRewardVoucher, voucherSig);

            // 2
            tokenRewardVoucher.beneficiary = GardenAdmin.address;
            tokenRewardVoucher.referrer = Issuer.address;
            tokenRewardVoucher.voucherCodeHash = randomBytes32();
            voucherSig = await signVoucher(tokenRewardVoucher, Issuer);
            await Garden.connect(GardenAdmin).useVoucher(tokenRewardVoucher, voucherSig);

            // 3
            tokenRewardVoucher.beneficiary = Issuer.address;
            tokenRewardVoucher.referrer = Referrer.address;
            tokenRewardVoucher.voucherCodeHash = randomBytes32();
            voucherSig = await signVoucher(tokenRewardVoucher, Issuer);
            await Garden.connect(Issuer).useVoucher(tokenRewardVoucher, voucherSig);

            // 4
            tokenRewardVoucher.beneficiary = Referrer.address;
            tokenRewardVoucher.referrer = Someone.address;
            tokenRewardVoucher.voucherCodeHash = randomBytes32();
            voucherSig = await signVoucher(tokenRewardVoucher, Issuer);
            await Garden.connect(Referrer).useVoucher(tokenRewardVoucher, voucherSig);

            // 5
            tokenRewardVoucher.beneficiary = Someone.address;
            tokenRewardVoucher.referrer = Someother.address;
            tokenRewardVoucher.voucherCodeHash = randomBytes32();
            voucherSig = await signVoucher(tokenRewardVoucher, Issuer);
            await expect(GardenAsSomeone.useVoucher(tokenRewardVoucher, voucherSig))
              .to.be.revertedWithCustomError(Garden, CIRCULAR_REFS)
              .withArgs(Someone.address, Someother.address);
          });

          it('revert on expire < now', async () => {
            someoneTokenRewardVoucher.expire = Math.round(Date.now() / 1000) - 60; // 1 min ago
            voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);

            await expect(
              GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig),
            ).to.be.revertedWithCustomError(Garden, INVALID_VOUCHER);
          });

          it('revert on incorrect chainId', async () => {
            someoneTokenRewardVoucher.chainId = 4242;
            voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);

            await expect(
              GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig),
            ).to.be.revertedWithCustomError(Garden, INVALID_VOUCHER);
          });

          it('revert on already used voucherCodeHash', async () => {
            await GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig);

            await expect(
              GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig),
            ).to.be.revertedWithCustomError(Garden, VOUCHER_USED);
          });
        });

        describe('invalid encodedParams', () => {
          it('revert on zero address reward token', async () => {
            rewardParams.token = constants.AddressZero;
            someoneTokenRewardVoucher.encodedParams = encodeRewardParams(rewardParams);
            voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);

            await expect(
              GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig),
            ).to.be.revertedWithCustomError(Garden, INVALID_REWARD);
          });

          it('revert on zero amount reward token', async () => {
            rewardParams.amount = 0;
            someoneTokenRewardVoucher.encodedParams = encodeRewardParams(rewardParams);
            voucherSig = await signVoucher(someoneTokenRewardVoucher, Issuer);

            await expect(
              GardenAsSomeone.useVoucher(someoneTokenRewardVoucher, voucherSig),
            ).to.be.revertedWithCustomError(Garden, INVALID_REWARD);
          });
        });
      });
    });
  });

  describe('use vouchers', () => {
    let Token2: TestERC20;

    let rewardParams: RewardParams;
    let reward2Params: RewardParams;
    let someoneTokenRewardVoucher: Voucher;
    let someoneToken2RewardVoucher: Voucher;

    let vouchersSig: string;

    beforeEach(async () => {
      const TestERC20Factory = await ethers.getContractFactory('TestERC20');
      Token2 = (await TestERC20Factory.deploy(
        'Partner2',
        'PARTNER2',
        TOKEN_CAP,
      )) as unknown as TestERC20;
      await Token2.deployed();

      await Token2.mint(Garden.address, TOKEN_DEPOSITED_TO_GARDEN);

      rewardParams = {
        token: Token.address,
        amount: AMOUNT,
        commissions: COMMISSIONS,
      };

      someoneTokenRewardVoucher = {
        ...VoucherBase,
        beneficiary: Someone.address,
        encodedParams: encodeRewardParams(rewardParams),
      };

      reward2Params = { ...rewardParams, token: Token2.address };

      someoneToken2RewardVoucher = {
        ...VoucherBase,
        beneficiary: Someone.address,
        voucherCodeHash: randomBytes32(),
        encodedParams: encodeRewardParams(reward2Params),
      };

      vouchersSig = await signVouchers(
        [someoneTokenRewardVoucher, someoneToken2RewardVoucher],
        Issuer,
      );
    });

    it('successfully transfer token', async () => {
      await GardenAsSomeone.useVouchers(
        [someoneTokenRewardVoucher, someoneToken2RewardVoucher],
        vouchersSig,
      );

      expect(await Token.balanceOf(Someone.address)).to.equal(AMOUNT);
      expect(await Token2.balanceOf(Someone.address)).to.equal(AMOUNT);
    });

    it('emit event for each voucher', async () => {
      await expect(
        GardenAsSomeone.useVouchers(
          [someoneTokenRewardVoucher, someoneToken2RewardVoucher],
          vouchersSig,
        ),
      )
        .to.emit(Garden, 'VoucherUsed')
        .withArgs(
          Someone.address,
          VoucherAction.Reward,
          someoneTokenRewardVoucher.voucherCodeHash,
          someoneTokenRewardVoucher.chainId,
        )
        .to.emit(Garden, 'VoucherUsed')
        .withArgs(
          Someone.address,
          VoucherAction.Reward,
          someoneToken2RewardVoucher.voucherCodeHash,
          someoneToken2RewardVoucher.chainId,
        );
    });
  });
});
