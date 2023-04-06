import { ethers } from "hardhat";
import { Contract, Signer } from "ethers";
import { expect } from "chai";

describe("TokenVesting", function () {
  let TokenVesting: any, ERC20: any;
  let tokenVesting: Contract, erc20: Contract;
  let owner: Signer, beneficiary: Signer;
  let ownerAddress: string, beneficiaryAddress: string;

  beforeEach(async function () {
    [owner, beneficiary] = await ethers.getSigners();
    ownerAddress = await owner.getAddress();
    beneficiaryAddress = await beneficiary.getAddress();

    ERC20 = await ethers.getContractFactory("ERC20");
    erc20 = await ERC20.deploy("TestToken", "TTK");
    await erc20.deployed();

    TokenVesting = await ethers.getContractFactory("TokenVesting");
    tokenVesting = await TokenVesting.deploy(erc20.address);
    await tokenVesting.deployed();

    // Transfer tokens to TokenVesting contract
    await erc20.transfer(tokenVesting.address, ethers.utils.parseUnits("1000", 18));
  });

  it("should deploy the TokenVesting contract with the correct token address", async function () {
    expect(await tokenVesting.token()).to.equal(erc20.address);
  });

  it("should add a new vesting schedule for a beneficiary", async function () {
    await expect(tokenVesting.connect(owner).addSchedule(beneficiaryAddress, ethers.utils.parseUnits("100", 18), 1627376512, 100))
      .to.emit(tokenVesting, "ScheduleAdded")
      .withArgs(beneficiaryAddress, ethers.utils.parseUnits("100", 18), 1627376512, 100);
  });

  it("should not add a new vesting schedule if not the owner", async function () {
    await expect(tokenVesting.connect(beneficiary).addSchedule(beneficiaryAddress, ethers.utils.parseUnits("100", 18), 1627376512, 100)).to.be.revertedWith(
      "Ownable: caller is not the owner"
    );
  });

  it("should delete a vesting schedule for a beneficiary", async function () {
    await tokenVesting.connect(owner).addSchedule(beneficiaryAddress, ethers.utils.parseUnits("100", 18), 1627376512, 100);
    await expect(tokenVesting.connect(owner).deleteSchedule(beneficiaryAddress, 0))
      .to.emit(tokenVesting, "ScheduleDeleted")
      .withArgs(beneficiaryAddress, 0);
  });

  it("should not delete a vesting schedule if not the owner", async function () {
    await tokenVesting.connect(owner).addSchedule(beneficiaryAddress, ethers.utils.parseUnits("100", 18), 1627376512, 100);
    await expect(tokenVesting.connect(beneficiary).deleteSchedule(beneficiaryAddress, 0)).to.be.revertedWith("Ownable: caller is not the owner");
  });
