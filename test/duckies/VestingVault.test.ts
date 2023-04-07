import { ethers } from "hardhat";
import { expect } from "chai";

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { TestERC20, VestingVault } from '../../typechain-types';


const TOKEN_CAP = 100_000_000_000_000;
const TOKEN_DECIMALS = 8;

describe("TokenVesting", function () {
  let TokenVesting: VestingVault, ERC20: TestERC20;
  let owner: SignerWithAddress, beneficiary: SignerWithAddress;
  let ownerAddress: string, beneficiaryAddress: string;

  before(async () => {
    [owner, beneficiary] = await ethers.getSigners();
    ownerAddress = await owner.getAddress();
    beneficiaryAddress = await beneficiary.getAddress();
  });

  beforeEach(async function () {
    const TestERC20Factory = await ethers.getContractFactory("TestERC20");
    ERC20 = (await TestERC20Factory.deploy("TestToken", "TTK", TOKEN_CAP)) as TestERC20;
    await ERC20.deployed();

    const VestingVaultFactory = await ethers.getContractFactory("VestingVault");
    TokenVesting = (await VestingVaultFactory.deploy(ERC20.address)) as VestingVault;
    await TokenVesting.deployed();

    // Transfer tokens to TokenVesting contract
    await ERC20.mint(TokenVesting.address, ethers.utils.parseUnits("1000", TOKEN_DECIMALS));
  });

  it("should deploy the TokenVesting contract with the correct token address", async function () {
    expect(await TokenVesting.token()).to.equal(ERC20.address);
  });

  it("should add a new vesting schedule for a beneficiary", async function () {
    await expect(TokenVesting.connect(owner).addSchedule(beneficiaryAddress, ethers.utils.parseUnits("100", TOKEN_DECIMALS), 1627376512, 100))
      .to.emit(TokenVesting, "ScheduleAdded")
      .withArgs(beneficiaryAddress, ethers.utils.parseUnits("100", TOKEN_DECIMALS), 1627376512, 100);
  });

  it("should not add a new vesting schedule if not the owner", async function () {
    await expect(TokenVesting.connect(beneficiary).addSchedule(beneficiaryAddress, ethers.utils.parseUnits("100", TOKEN_DECIMALS), 1627376512, 100)).to.be.revertedWith(
      "Ownable: caller is not the owner"
    );
  });

  it("should delete a vesting schedule for a beneficiary", async function () {
    await TokenVesting.connect(owner).addSchedule(beneficiaryAddress, ethers.utils.parseUnits("100", TOKEN_DECIMALS), 1627376512, 100);
    await expect(TokenVesting.connect(owner).deleteSchedule(beneficiaryAddress, 0))
      .to.emit(TokenVesting, "ScheduleDeleted")
      .withArgs(beneficiaryAddress, 0);
  });

  it("should not delete a vesting schedule if not the owner", async function () {
    await TokenVesting.connect(owner).addSchedule(beneficiaryAddress, ethers.utils.parseUnits("100", TOKEN_DECIMALS), 1627376512, 100);
    await expect(TokenVesting.connect(beneficiary).deleteSchedule(beneficiaryAddress, 0)).to.be.revertedWith("Ownable: caller is not the owner");
  });
})
