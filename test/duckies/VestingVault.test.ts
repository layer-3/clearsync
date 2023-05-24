import { ethers } from 'hardhat';
import { expect } from 'chai';
import { setNextBlockTimestamp } from '@nomicfoundation/hardhat-network-helpers/dist/src/helpers/time';
import { SnapshotRestorer, takeSnapshot } from '@nomicfoundation/hardhat-network-helpers';

import { connectGroup } from '../helpers/connect';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { TestERC20, VestingVault } from '../../typechain-types';

const TOKEN_CAP = 100_000_000_000_000;
const TOKEN_DECIMALS = 8;
const NOW = Math.floor(Date.now() / 1000);
const IN_FUTURE = NOW + 60 * 10;
const IN_PAST = NOW - 60 * 10;

describe('Vesting', function () {
  let Owner: SignerWithAddress, Beneficiary: SignerWithAddress, Someone: SignerWithAddress;
  let BeneficiaryAddress: string;

  let Vesting: VestingVault,
    VestingAsOwner: VestingVault,
    VestingAsBeneficiary: VestingVault,
    VestingAsSomeone: VestingVault,
    ERC20: TestERC20;

  before(async () => {
    [Owner, Beneficiary, Someone] = await ethers.getSigners();
    BeneficiaryAddress = await Beneficiary.getAddress();
  });

  beforeEach(async function () {
    const TestERC20Factory = await ethers.getContractFactory('TestERC20');
    ERC20 = (await TestERC20Factory.deploy('TestToken', 'TTK', TOKEN_CAP)) as TestERC20;
    await ERC20.deployed();

    const VestingVaultFactory = await ethers.getContractFactory('VestingVault');
    Vesting = (await VestingVaultFactory.deploy(ERC20.address)) as VestingVault;
    await Vesting.deployed();

    [VestingAsOwner, VestingAsBeneficiary, VestingAsSomeone] = connectGroup(Vesting, [
      Owner,
      Beneficiary,
      Someone,
    ]);

    // Transfer tokens to Vesting contract
    await ERC20.mint(Vesting.address, ethers.utils.parseUnits('1000', TOKEN_DECIMALS));
  });

  describe('deployment', () => {
    it('has correct token address', async function () {
      expect(await Vesting.token()).to.equal(ERC20.address);
    });
  });

  describe('addSchedule', () => {
    it('success when Owner adds new schedule', async function () {
      await expect(
        VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          IN_FUTURE,
          100,
        ),
      )
        .to.emit(Vesting, 'ScheduleAdded')
        .withArgs(
          BeneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          IN_FUTURE,
          100,
        );
    });

    it('revert when not Owner adds new schedule', async function () {
      await expect(
        VestingAsSomeone.addSchedule(
          BeneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          IN_FUTURE,
          100,
        ),
      ).to.be.revertedWith('Ownable: caller is not the Owner');
    });

    it('revert when adding schedule with zero amount', async function () {
      await expect(
        VestingAsOwner.addSchedule(BeneficiaryAddress, 0, IN_FUTURE, 100),
      ).to.be.revertedWithCustomError(Vesting, 'InvalidSchedule');
    });

    it('revert when adding schedule with zero duration', async function () {
      await expect(
        VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          0,
          100,
        ),
      ).to.be.revertedWithCustomError(Vesting, 'InvalidSchedule');
    });

    it('revert when adding schedule with start in the past', async function () {
      await expect(
        VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          IN_PAST,
          100,
        ),
      ).to.be.revertedWithCustomError(Vesting, 'InvalidSchedule');
    });
  });

  describe('deleteSchedule', () => {
    it('success when Owner removes schedule', async function () {
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        ethers.utils.parseUnits('100', TOKEN_DECIMALS),
        IN_FUTURE,
        100,
      );
      await expect(Vesting.connect(Owner).deleteSchedule(BeneficiaryAddress, 0))
        .to.emit(Vesting, 'ScheduleDeleted')
        .withArgs(BeneficiaryAddress, 0);
    });

    it('revert when not Owner removes schedule', async function () {
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        ethers.utils.parseUnits('100', TOKEN_DECIMALS),
        IN_FUTURE,
        100,
      );
      await expect(VestingAsSomeone.deleteSchedule(BeneficiaryAddress, 0)).to.be.revertedWith(
        'Ownable: caller is not the Owner',
      );
    });

    it('revert when Owner removes non-existent schedule', async function () {
      const index = 0;
      await expect(VestingAsOwner.deleteSchedule(BeneficiaryAddress, index))
        .to.be.revertedWithCustomError(Vesting, 'NoScheduleForBeneficiary')
        .withArgs(BeneficiaryAddress, index);
    });
  });

  describe.only('claim', () => {
    const vestingAmount = ethers.utils.parseUnits('100', TOKEN_DECIMALS);
    const vestingPeriod = 60 * 60 * 24 * 10;

    let snapshot: SnapshotRestorer;

    beforeEach(async function () {
      snapshot = await takeSnapshot();
    });

    afterEach(async function () {
      await snapshot.restore();
    });

    describe('one schedule', () => {
      beforeEach(async function () {
        await VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          vestingAmount,
          IN_FUTURE,
          vestingPeriod,
        );
      });

      it('claim all tokens after vesting period ends', async function () {
        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(0);

        await setNextBlockTimestamp(IN_FUTURE + vestingPeriod);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(vestingAmount);
      });

      it('claim part of the tokens before vesting period ends', async function () {
        const timeDiff = 60 * 60 * 24 * 5;
        await setNextBlockTimestamp(IN_FUTURE + timeDiff);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          vestingAmount.mul(timeDiff).div(vestingPeriod),
        );
      });

      it('claim consequently', async function () {
        const timeDiff1 = 60 * 60 * 24 * 2;
        await setNextBlockTimestamp(IN_FUTURE + timeDiff1);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          vestingAmount.mul(timeDiff1).div(vestingPeriod),
        );

        const timeDiff2 = 60 * 60 * 24 * 4;
        await setNextBlockTimestamp(IN_FUTURE + timeDiff2);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          vestingAmount.mul(timeDiff2).div(vestingPeriod),
        );
      });

      it('revert when no schedule for Beneficiary', async function () {
        await expect(VestingAsSomeone.claim())
          .to.be.revertedWithCustomError(Vesting, 'UnableToClaim')
          .withArgs(Someone.address);
      });

      it('revert when vesting has not started', async function () {
        await expect(VestingAsBeneficiary.claim())
          .to.be.revertedWithCustomError(Vesting, 'UnableToClaim')
          .withArgs(BeneficiaryAddress);
      });

      it('revert when no tokens to claim', async function () {
        await setNextBlockTimestamp(IN_FUTURE + vestingPeriod);
        await VestingAsBeneficiary.claim();

        await expect(VestingAsBeneficiary.claim())
          .to.be.revertedWithCustomError(Vesting, 'UnableToClaim')
          .withArgs(BeneficiaryAddress);
      });
    });

    describe('multiple schedules', () => {
      it.skip('claim from multiple schedules');
    });
  });
});
