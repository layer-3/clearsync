import { ethers } from 'hardhat';
import { expect } from 'chai';

import { connectGroup } from '../helpers/connect';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { TestERC20, VestingVault } from '../../typechain-types';

const TOKEN_CAP = 100_000_000_000_000;
const TOKEN_DECIMALS = 8;
const NOW = Math.floor(Date.now() / 1000);
const IN_FUTURE = NOW + 60 * 10;
const IN_PAST = NOW - 60 * 10;

describe('Vesting', function () {
  let owner: SignerWithAddress, beneficiary: SignerWithAddress, someone: SignerWithAddress;
  let beneficiaryAddress: string;

  let Vesting: VestingVault,
    VestingAsOwner: VestingVault,
    VestingAsBeneficiary: VestingVault,
    VestingAsSomeone: VestingVault,
    ERC20: TestERC20;

  before(async () => {
    [owner, beneficiary, someone] = await ethers.getSigners();
    beneficiaryAddress = await beneficiary.getAddress();
  });

  beforeEach(async function () {
    const TestERC20Factory = await ethers.getContractFactory('TestERC20');
    ERC20 = (await TestERC20Factory.deploy('TestToken', 'TTK', TOKEN_CAP)) as TestERC20;
    await ERC20.deployed();

    const VestingVaultFactory = await ethers.getContractFactory('VestingVault');
    Vesting = (await VestingVaultFactory.deploy(ERC20.address)) as VestingVault;
    await Vesting.deployed();

    [VestingAsOwner, VestingAsBeneficiary, VestingAsSomeone] = connectGroup(Vesting, [
      owner,
      beneficiary,
      someone,
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
    it('success when owner adds new schedule', async function () {
      await expect(
        VestingAsOwner.addSchedule(
          beneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          IN_FUTURE,
          100,
        ),
      )
        .to.emit(Vesting, 'ScheduleAdded')
        .withArgs(
          beneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          IN_FUTURE,
          100,
        );
    });

    it('revert when not owner adds new schedule', async function () {
      await expect(
        VestingAsSomeone.addSchedule(
          beneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          IN_FUTURE,
          100,
        ),
      ).to.be.revertedWith('Ownable: caller is not the owner');
    });

    it('revert when adding schedule with zero amount', async function () {
      await expect(
        VestingAsOwner.addSchedule(beneficiaryAddress, 0, IN_FUTURE, 100),
      ).to.be.revertedWithCustomError(Vesting, 'InvalidSchedule');
    });

    it('revert when adding schedule with zero duration', async function () {
      await expect(
        VestingAsOwner.addSchedule(
          beneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          0,
          100,
        ),
      ).to.be.revertedWithCustomError(Vesting, 'InvalidSchedule');
    });

    it('revert when adding schedule with start in the past', async function () {
      await expect(
        VestingAsOwner.addSchedule(
          beneficiaryAddress,
          ethers.utils.parseUnits('100', TOKEN_DECIMALS),
          IN_PAST,
          100,
        ),
      ).to.be.revertedWithCustomError(Vesting, 'InvalidSchedule');
    });
  });

  describe('deleteSchedule', () => {
    it('success when owner removes schedule', async function () {
      await VestingAsOwner.addSchedule(
        beneficiaryAddress,
        ethers.utils.parseUnits('100', TOKEN_DECIMALS),
        IN_FUTURE,
        100,
      );
      await expect(Vesting.connect(owner).deleteSchedule(beneficiaryAddress, 0))
        .to.emit(Vesting, 'ScheduleDeleted')
        .withArgs(beneficiaryAddress, 0);
    });

    it('revert when not owner removes schedule', async function () {
      await VestingAsOwner.addSchedule(
        beneficiaryAddress,
        ethers.utils.parseUnits('100', TOKEN_DECIMALS),
        IN_FUTURE,
        100,
      );
      await expect(VestingAsSomeone.deleteSchedule(beneficiaryAddress, 0)).to.be.revertedWith(
        'Ownable: caller is not the owner',
      );
    });

    it('revert when owner removes non-existent schedule', async function () {
      const index = 0;
      await expect(VestingAsOwner.deleteSchedule(beneficiaryAddress, index))
        .to.be.revertedWithCustomError(Vesting, 'NoScheduleForBeneficiary')
        .withArgs(beneficiaryAddress, index);
    });
  });

  // describe('claim', () => {});
});
