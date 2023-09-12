import { ethers } from 'hardhat';
import { expect } from 'chai';
import {
  increaseTo,
  setNextBlockTimestamp,
} from '@nomicfoundation/hardhat-network-helpers/dist/src/helpers/time';
import { SnapshotRestorer, takeSnapshot } from '@nomicfoundation/hardhat-network-helpers';
import { constants } from 'ethers';

import { connectGroup } from '../helpers/connect';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { TestERC20, VestingVault } from '../../typechain-types';

const TOKEN_CAP = 100_000_000_000_000;
const TOKEN_DECIMALS = 8;
const NOW = Math.floor(Date.now() / 1000);
const TIME_DIFF = 600 * 10; // 10 minutes
const VESTING_1_START = NOW + TIME_DIFF;

const VESTING_1_AMOUNT = ethers.utils.parseUnits('100', TOKEN_DECIMALS);
const VESTING_DURATION = 60 * 60 * 24 * 10; // 10 days

const VESTING_2_AMOUNT = VESTING_1_AMOUNT.mul(2);
const VESTING_2_START_SHIFT = VESTING_DURATION / 2;
const VESTING_2_START = VESTING_1_START + VESTING_2_START_SHIFT;

describe('Vesting', function () {
  let Owner: SignerWithAddress, Beneficiary: SignerWithAddress, Someone: SignerWithAddress;
  let BeneficiaryAddress: string;

  let Vesting: VestingVault,
    VestingAsOwner: VestingVault,
    VestingAsBeneficiary: VestingVault,
    VestingAsSomeone: VestingVault,
    ERC20: TestERC20;

  // each test to start with current timestamp
  let snapshot: SnapshotRestorer;

  before(async () => {
    [Owner, Beneficiary, Someone] = await ethers.getSigners();
    BeneficiaryAddress = await Beneficiary.getAddress();
  });

  beforeEach(async function () {
    snapshot = await takeSnapshot();

    const TestERC20Factory = await ethers.getContractFactory('TestERC20');
    ERC20 = (await TestERC20Factory.deploy('TestToken', 'TTK', 8, TOKEN_CAP)) as TestERC20;
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

  afterEach(async function () {
    await snapshot.restore();
  });

  describe('deployment', () => {
    it('has correct token address', async function () {
      expect(await Vesting.token()).to.equal(ERC20.address);
    });

    it('revert when token address is zero', async function () {
      const VestingVaultFactory = await ethers.getContractFactory('VestingVault');
      await expect(VestingVaultFactory.deploy(constants.AddressZero)).to.be.reverted;
    });
  });

  describe('addSchedule', () => {
    it('success when Owner adds new schedule', async function () {
      await expect(
        VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          VESTING_1_AMOUNT,
          VESTING_1_START,
          VESTING_DURATION,
        ),
      )
        .to.emit(Vesting, 'ScheduleAdded')
        .withArgs(BeneficiaryAddress, VESTING_1_AMOUNT, VESTING_1_START, VESTING_DURATION);
    });

    it('revert when not Owner adds new schedule', async function () {
      await expect(
        VestingAsSomeone.addSchedule(
          BeneficiaryAddress,
          VESTING_1_AMOUNT,
          VESTING_1_START,
          VESTING_DURATION,
        ),
      ).to.be.revertedWith('Ownable: caller is not the owner');
    });

    it('revert when adding schedule for zero address', async function () {
      await expect(
        VestingAsOwner.addSchedule(
          ethers.constants.AddressZero,
          VESTING_1_AMOUNT,
          VESTING_1_START,
          VESTING_DURATION,
        ),
      ).to.be.revertedWithCustomError(Vesting, 'InvalidSchedule');
    });

    it('revert when adding schedule with zero amount', async function () {
      await expect(
        VestingAsOwner.addSchedule(BeneficiaryAddress, 0, VESTING_1_START, VESTING_DURATION),
      ).to.be.revertedWithCustomError(Vesting, 'InvalidSchedule');
    });

    it('revert when adding schedule with start in the past', async function () {
      await expect(
        VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          VESTING_1_AMOUNT,
          NOW - 42,
          VESTING_DURATION,
        ),
      ).to.be.revertedWithCustomError(Vesting, 'InvalidSchedule');
    });

    it('revert when adding schedule with zero duration', async function () {
      await expect(
        VestingAsOwner.addSchedule(BeneficiaryAddress, VESTING_1_AMOUNT, VESTING_1_START, 0),
      ).to.be.revertedWithCustomError(Vesting, 'InvalidSchedule');
    });

    it('event is emitted', async function () {
      await expect(
        VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          VESTING_1_AMOUNT,
          VESTING_1_START,
          VESTING_DURATION,
        ),
      )
        .to.emit(Vesting, 'ScheduleAdded')
        .withArgs(BeneficiaryAddress, VESTING_1_AMOUNT, VESTING_1_START, VESTING_DURATION);
    });
  });

  describe('deleteSchedule', () => {
    it('success when Owner removes schedule', async function () {
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        VESTING_1_AMOUNT,
        VESTING_1_START,
        VESTING_DURATION,
      );
      await expect(Vesting.connect(Owner).deleteSchedule(BeneficiaryAddress, 0))
        .to.emit(Vesting, 'ScheduleDeleted')
        .withArgs(BeneficiaryAddress, 0);
    });

    it('revert when not Owner removes schedule', async function () {
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        VESTING_1_AMOUNT,
        VESTING_1_START,
        VESTING_DURATION,
      );
      await expect(VestingAsSomeone.deleteSchedule(BeneficiaryAddress, 0)).to.be.revertedWith(
        'Ownable: caller is not the owner',
      );
    });

    it('revert when Owner removes non-existent schedule', async function () {
      const index = 0;
      await expect(VestingAsOwner.deleteSchedule(BeneficiaryAddress, index))
        .to.be.revertedWithCustomError(Vesting, 'NoScheduleForBeneficiary')
        .withArgs(BeneficiaryAddress, index);
    });

    it('event is emitted', async function () {
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        VESTING_1_AMOUNT,
        VESTING_1_START,
        VESTING_DURATION,
      );
      await expect(VestingAsOwner.deleteSchedule(BeneficiaryAddress, 0))
        .to.emit(Vesting, 'ScheduleDeleted')
        .withArgs(BeneficiaryAddress, 0);
    });
  });

  describe('beneficiarySchedules', () => {
    it('returns correct schedules', async function () {
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        VESTING_1_AMOUNT,
        VESTING_1_START,
        VESTING_DURATION,
      );
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        VESTING_2_AMOUNT,
        VESTING_2_START,
        VESTING_DURATION,
      );

      const schedules = await VestingAsOwner.beneficiarySchedules(BeneficiaryAddress);
      expect(schedules).to.have.lengthOf(2);
      expect(schedules[0].amount).to.equal(VESTING_1_AMOUNT);
      expect(schedules[0].start).to.equal(VESTING_1_START);
      expect(schedules[0].duration).to.equal(VESTING_DURATION);
      expect(schedules[1].amount).to.equal(VESTING_2_AMOUNT);
      expect(schedules[1].start).to.equal(VESTING_2_START);
      expect(schedules[1].duration).to.equal(VESTING_DURATION);
    });

    it('returns empty array when no schedules', async function () {
      const schedules = await VestingAsOwner.beneficiarySchedules(BeneficiaryAddress);
      expect(schedules).to.have.lengthOf(0);
    });
  });

  describe('beneficiarySchedule', () => {
    it('returns correct schedule', async function () {
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        VESTING_1_AMOUNT,
        VESTING_1_START,
        VESTING_DURATION,
      );

      const schedule = await VestingAsOwner.beneficiarySchedule(BeneficiaryAddress, 0);
      expect(schedule.amount).to.equal(VESTING_1_AMOUNT);
      expect(schedule.start).to.equal(VESTING_1_START);
      expect(schedule.duration).to.equal(VESTING_DURATION);
    });

    it('revert when no schedule', async function () {
      await expect(VestingAsOwner.beneficiarySchedule(BeneficiaryAddress, 0))
        .to.be.revertedWithCustomError(Vesting, 'NoScheduleForBeneficiary')
        .withArgs(BeneficiaryAddress, 0);
    });
  });

  describe('scheduleClaimable', () => {
    beforeEach(async function () {
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        VESTING_1_AMOUNT,
        VESTING_1_START,
        VESTING_DURATION,
      );
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        VESTING_2_AMOUNT,
        VESTING_2_START,
        VESTING_DURATION,
      );
    });

    it('returns correct amount for schedule 1 after period ends', async function () {
      await increaseTo(VESTING_1_START + VESTING_DURATION);
      expect(await Vesting.scheduleClaimable(BeneficiaryAddress, 0)).to.equal(VESTING_1_AMOUNT);
    });

    it('returns correct amount for schedule 2 after period ends', async function () {
      await increaseTo(VESTING_2_START + VESTING_DURATION);
      expect(await Vesting.scheduleClaimable(BeneficiaryAddress, 1)).to.equal(VESTING_2_AMOUNT);
    });

    it('returns correct amount for schedule 1 before perion ends', async function () {
      await increaseTo(VESTING_1_START + VESTING_DURATION / 2);
      expect(await Vesting.scheduleClaimable(BeneficiaryAddress, 0)).to.equal(
        VESTING_1_AMOUNT.div(2),
      );
    });

    it('returns zero after tokens claimed before period ends', async function () {
      await increaseTo(VESTING_1_START + VESTING_DURATION / 2);
      await VestingAsBeneficiary.claim();
      expect(await Vesting.scheduleClaimable(BeneficiaryAddress, 0)).to.equal(0);
    });

    it('revert after schedule claimed', async function () {
      await increaseTo(VESTING_1_START + VESTING_DURATION);
      await VestingAsBeneficiary.claim();
      expect(await Vesting.scheduleClaimable(BeneficiaryAddress, 0)).to.equal(0);
    });

    it('revert when schedule does not exist', async function () {
      await expect(Vesting.scheduleClaimable(BeneficiaryAddress, 2))
        .to.be.revertedWithCustomError(Vesting, 'NoScheduleForBeneficiary')
        .withArgs(BeneficiaryAddress, 2);
    });
  });

  describe('totalClaimable', () => {
    beforeEach(async function () {
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        VESTING_1_AMOUNT,
        VESTING_1_START,
        VESTING_DURATION,
      );
      await VestingAsOwner.addSchedule(
        BeneficiaryAddress,
        VESTING_2_AMOUNT,
        VESTING_2_START,
        VESTING_DURATION,
      );
    });

    it('correct amount after all periods end', async function () {
      await increaseTo(VESTING_2_START + VESTING_DURATION);
      expect(await Vesting.totalClaimable(BeneficiaryAddress)).to.equal(
        VESTING_1_AMOUNT.add(VESTING_2_AMOUNT),
      );
    });

    it('correct amount before all periods end', async function () {
      await increaseTo(VESTING_1_START + (VESTING_DURATION * 3) / 4);
      expect(await Vesting.totalClaimable(BeneficiaryAddress)).to.equal(
        VESTING_1_AMOUNT.mul(3).div(4).add(VESTING_2_AMOUNT.div(4)),
      );
    });

    it('zero after tokens claimed before period ends', async function () {
      await increaseTo(VESTING_1_START + (VESTING_DURATION * 3) / 4);
      await VestingAsBeneficiary.claim();
      expect(await Vesting.totalClaimable(BeneficiaryAddress)).to.equal(0);
    });

    it('zero after all tokens claimed', async function () {
      await increaseTo(VESTING_2_START + VESTING_DURATION);
      await VestingAsBeneficiary.claim();
      expect(await Vesting.totalClaimable(BeneficiaryAddress)).to.equal(0);
    });

    it('zero when no schedules for Beneficiary', async function () {
      expect(await Vesting.totalClaimable(Someone.address)).to.equal(0);
    });
  });

  describe('claim', () => {
    describe('one schedule', () => {
      beforeEach(async function () {
        await VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          VESTING_1_AMOUNT,
          VESTING_1_START,
          VESTING_DURATION,
        );
      });

      it('claim all tokens after vesting period ends', async function () {
        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(0);

        await setNextBlockTimestamp(VESTING_1_START + VESTING_DURATION);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(VESTING_1_AMOUNT);
      });

      it('claim part of the tokens before vesting period ends', async function () {
        const timeDiff = 60 * 60 * 24 * 5;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          VESTING_1_AMOUNT.mul(timeDiff).div(VESTING_DURATION),
        );
      });

      it('claim consequently', async function () {
        const timeDiff1 = 60 * 60 * 24 * 2;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff1);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          VESTING_1_AMOUNT.mul(timeDiff1).div(VESTING_DURATION),
        );

        const timeDiff2 = 60 * 60 * 24 * 4;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff2);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          VESTING_1_AMOUNT.mul(timeDiff2).div(VESTING_DURATION),
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
        await setNextBlockTimestamp(VESTING_1_START + VESTING_DURATION);
        await VestingAsBeneficiary.claim();

        await expect(VestingAsBeneficiary.claim())
          .to.be.revertedWithCustomError(Vesting, 'UnableToClaim')
          .withArgs(BeneficiaryAddress);
      });

      it('deletes schedule when all tokens claimed', async () => {
        await setNextBlockTimestamp(VESTING_1_START + VESTING_DURATION);
        await VestingAsBeneficiary.claim();

        await expect(VestingAsOwner.beneficiarySchedule(BeneficiaryAddress, 0))
          .to.be.revertedWithCustomError(Vesting, 'NoScheduleForBeneficiary')
          .withArgs(BeneficiaryAddress, 0);
      });

      it('event is emitted', async function () {
        await setNextBlockTimestamp(VESTING_1_START + VESTING_DURATION);
        await expect(VestingAsBeneficiary.claim())
          .to.emit(Vesting, 'TokensClaimed')
          .withArgs(BeneficiaryAddress, VESTING_1_AMOUNT);
      });
    });

    describe('multiple schedules', () => {
      // starts and vesting periods are the same
      beforeEach(async function () {
        await VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          VESTING_1_AMOUNT,
          VESTING_1_START,
          VESTING_DURATION,
        );
        await VestingAsOwner.addSchedule(
          BeneficiaryAddress,
          VESTING_2_AMOUNT,
          VESTING_2_START,
          VESTING_DURATION,
        );
      });

      it('claim all tokens after multiple schedules end', async () => {
        // after both schedules end
        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(0);

        await setNextBlockTimestamp(VESTING_2_START + VESTING_DURATION);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          VESTING_1_AMOUNT.add(VESTING_2_AMOUNT),
        );
      });

      it('claim part of the tokens before any schedule ends', async () => {
        // when both schedules are active: 1st schedule is 75% done, 2nd schedule is 25% done
        const timeDiff = VESTING_2_START_SHIFT * 1.5;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff);
        await VestingAsBeneficiary.claim();

        const expectedBalance = VESTING_1_AMOUNT.mul(timeDiff)
          .div(VESTING_DURATION)
          .add(
            VESTING_2_AMOUNT.mul(VESTING_1_START + timeDiff - VESTING_2_START).div(
              VESTING_DURATION,
            ),
          );

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(expectedBalance);
      });

      it('claim consequently', async () => {
        const timeDiff1 = VESTING_2_START_SHIFT - 100;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff1);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          VESTING_1_AMOUNT.mul(timeDiff1).div(VESTING_DURATION),
        );

        const timeDiff2 = VESTING_2_START_SHIFT + 4200;
        await setNextBlockTimestamp(VESTING_1_START + timeDiff2);
        await VestingAsBeneficiary.claim();

        expect(await ERC20.balanceOf(BeneficiaryAddress)).to.equal(
          VESTING_1_AMOUNT.mul(timeDiff2)
            .div(VESTING_DURATION)
            .add(VESTING_2_AMOUNT.mul(4200).div(VESTING_DURATION)),
        );
      });

      it('revert when neither vesting has not started', async function () {
        await expect(VestingAsBeneficiary.claim())
          .to.be.revertedWithCustomError(Vesting, 'UnableToClaim')
          .withArgs(BeneficiaryAddress);
      });

      it('revert when no tokens to claim', async function () {
        await setNextBlockTimestamp(VESTING_2_START + VESTING_DURATION);
        await VestingAsBeneficiary.claim();

        await expect(VestingAsBeneficiary.claim())
          .to.be.revertedWithCustomError(Vesting, 'UnableToClaim')
          .withArgs(BeneficiaryAddress);
      });

      it('deletes one schedule when all its tokens claimed', async () => {
        await setNextBlockTimestamp(VESTING_1_START + VESTING_DURATION);
        await VestingAsBeneficiary.claim();

        const schedulesLeft = await VestingAsOwner.beneficiarySchedules(BeneficiaryAddress);
        expect(schedulesLeft).to.have.lengthOf(1);
      });

      it('deletes both schedules when all tokens claimed', async () => {
        await setNextBlockTimestamp(VESTING_2_START + VESTING_DURATION);
        await VestingAsBeneficiary.claim();

        await expect(VestingAsOwner.beneficiarySchedule(BeneficiaryAddress, 0))
          .to.be.revertedWithCustomError(Vesting, 'NoScheduleForBeneficiary')
          .withArgs(BeneficiaryAddress, 0);

        await expect(VestingAsOwner.beneficiarySchedule(BeneficiaryAddress, 1))
          .to.be.revertedWithCustomError(Vesting, 'NoScheduleForBeneficiary')
          .withArgs(BeneficiaryAddress, 1);
      });

      it('event is emitted', async function () {
        await setNextBlockTimestamp(VESTING_2_START + VESTING_DURATION);
        await expect(VestingAsBeneficiary.claim())
          .to.emit(Vesting, 'TokensClaimed')
          .withArgs(BeneficiaryAddress, VESTING_1_AMOUNT.add(VESTING_2_AMOUNT));
      });
    });
  });
});
