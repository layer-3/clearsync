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
const TIME_DIFF = 60 * 10; // 10 minutes
const VESTING_1_START = NOW + TIME_DIFF;

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
          VESTING_1_START,
          100,
        ),
      )
        .to.emit(Vesting, 'ScheduleAdded')
        .withArgs(
          BeneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          VESTING_1_START,
          100,
        );
    });

    it('revert when not Owner adds new schedule', async function () {
      await expect(
        VestingAsSomeone.addSchedule(
          BeneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          VESTING_1_START,
          100,
        ),
      ).to.be.revertedWith('Ownable: caller is not the Owner');
    });

    it('revert when adding schedule with zero amount', async function () {
      await expect(
        VestingAsOwner.addSchedule(BeneficiaryAddress, 0, VESTING_1_START, 100),
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
          NOW - 42,
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
        VESTING_1_START,
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
        VESTING_1_START,
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

  describe('claim', () => {
    const vestingAmount = ethers.utils.parseUnits('100', TOKEN_DECIMALS);
    const vestingDuration = 60 * 60 * 24 * 10; // 10 days

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
          VESTING_1_START,
          vestingDuration,
        );
      });

      it('claim all tokens after vesting period ends', async function () {
        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(0);

        await setNextBlockTimestamp(VESTING_1_START + vestingDuration);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(vestingAmount);
      });

      it('claim part of the tokens before vesting period ends', async function () {
        const timeDiff = 60 * 60 * 24 * 5;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          vestingAmount.mul(timeDiff).div(vestingDuration),
        );
      });

      it('claim consequently', async function () {
        const timeDiff1 = 60 * 60 * 24 * 2;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff1);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          vestingAmount.mul(timeDiff1).div(vestingDuration),
        );

        const timeDiff2 = 60 * 60 * 24 * 4;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff2);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          vestingAmount.mul(timeDiff2).div(vestingDuration),
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
        await setNextBlockTimestamp(VESTING_1_START + vestingDuration);
        await VestingAsBeneficiary.claim();

        await expect(VestingAsBeneficiary.claim())
          .to.be.revertedWithCustomError(Vesting, 'UnableToClaim')
          .withArgs(BeneficiaryAddress);
      });
    });

    describe.only('multiple schedules', () => {
      const VESTING_AMOUNT_2 = vestingAmount.mul(2);
      const VESTING_2_START_SHIFT = vestingDuration / 2;
      const VESTING_2_START = VESTING_1_START + VESTING_2_START_SHIFT;

      // starts and vesting periods are the same
      beforeEach(async function () {
        await VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          vestingAmount,
          VESTING_1_START,
          vestingDuration,
        );
        await VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          VESTING_AMOUNT_2,
          VESTING_2_START,
          vestingDuration,
        );
      });

      it('claim all tokens after multiple schedules end', async () => {
        // after both schedules end
        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(0);

        await setNextBlockTimestamp(VESTING_2_START + vestingDuration);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          vestingAmount.add(VESTING_AMOUNT_2),
        );
      });

      it('claim part of the tokens before any schedule ends', async () => {
        // when both schedules are active: 1st schedule is 75% done, 2nd schedule is 25% done
        const timeDiff = VESTING_2_START_SHIFT * 1.5;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff);
        await VestingAsBeneficiary.claim();

        const expectedBalance = vestingAmount
          .mul(timeDiff)
          .div(vestingDuration)
          .add(
            VESTING_AMOUNT_2.mul(VESTING_1_START + timeDiff - VESTING_2_START).div(vestingDuration),
          );

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(expectedBalance);
      });

      it('claim consequently', async () => {
        const timeDiff1 = VESTING_2_START_SHIFT - 100;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff1);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          vestingAmount.mul(timeDiff1).div(vestingDuration),
        );

        const timeDiff2 = VESTING_2_START_SHIFT + 4200;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff2);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          vestingAmount
            .mul(timeDiff2)
            .div(vestingDuration)
            .add(VESTING_AMOUNT_2.mul(4200).div(vestingDuration)),
        );
      });

      it('revert when neither vesting has not started', async function () {
        await expect(VestingAsBeneficiary.claim())
          .to.be.revertedWithCustomError(Vesting, 'UnableToClaim')
          .withArgs(BeneficiaryAddress);
      });

      it('revert when no tokens to claim', async function () {
        await setNextBlockTimestamp(VESTING_2_START + vestingDuration);
        await VestingAsBeneficiary.claim();

        await expect(VestingAsBeneficiary.claim())
          .to.be.revertedWithCustomError(Vesting, 'UnableToClaim')
          .withArgs(BeneficiaryAddress);
      });
    });
  });
});
