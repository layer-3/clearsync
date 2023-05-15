import { loadFixture } from '@nomicfoundation/hardhat-network-helpers';
import { expect } from 'chai';
import { ethers } from 'hardhat';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { BatchTransfer, TestERC20 } from '../../typechain-types';
import type { BigNumber } from 'ethers';

describe('BatchTransfer', function () {
  const TOKEN_CAP = 100_000_000_000;
  const BATCHER_BALANCE = 10000;

  let TokenAdmin: SignerWithAddress;
  let BatchTransferOwner: SignerWithAddress;
  let Receivers: SignerWithAddress[];
  let nativeBalances: BigNumber[];
  let tokenBalances: BigNumber[];

  let Token:TestERC20;
  let Batcher: BatchTransfer;

  async function deployTokenFixture(): Promise<{ Token:TestERC20, Batcher: BatchTransfer }> {
    const TokenFactory = await ethers.getContractFactory('TestERC20');
    const Token = (await TokenFactory.connect(TokenAdmin).deploy(
      'Test',
      'TEST',
      TOKEN_CAP,
    )) as TestERC20;

    const BatcherFactory = await ethers.getContractFactory('BatchTransfer');
    const Batcher = (await BatcherFactory.connect(BatchTransferOwner).deploy()) as BatchTransfer;

    return { Token, Batcher };
  }

  before(async () => {
    [TokenAdmin, BatchTransferOwner, ...Receivers] =
      await ethers.getSigners();
  });

  beforeEach(async () => {
    ({ Token, Batcher } = await loadFixture(deployTokenFixture));
    Token = Token.connect(TokenAdmin);
    Batcher = Batcher.connect(BatchTransferOwner);
    nativeBalances = await Promise.all(Receivers.map(r => r.getBalance()));
    tokenBalances = await Promise.all(Receivers.map(r => Token.balanceOf(r.address)));

    await Token.mint(Batcher.address, BATCHER_BALANCE);
    await TokenAdmin.sendTransaction({
      to: Batcher.address,
      value: BATCHER_BALANCE,
    });
  });

  it('Successful ERC20 transfers', async () => {
    const amount = 20;
    const receivers = Receivers.slice(0, 20).map(r => r.address);
    await Batcher.batchTransfer(Token.address, receivers, amount);

    expect(await Token.balanceOf(Batcher.address)).to.equal(BATCHER_BALANCE - amount * receivers.length);
    
    for (let i = 0; i < receivers.length; i++) {
      expect(await Token.balanceOf(receivers[i])).to.equal(tokenBalances[i].add(amount));
    }
  });

  it('Successful native token transfers', async () => {
    const amount = 20;
    const receivers = Receivers.slice(0, 20).map(r => r.address);
    await Batcher.batchTransfer(ethers.constants.AddressZero, receivers, amount);

    expect(await Batcher.provider.getBalance(Batcher.address)).to.equal(BATCHER_BALANCE - amount * receivers.length);
    
    for (let i = 0; i < receivers.length; i++) {
      expect(await Receivers[i].getBalance()).to.equal(nativeBalances[i].add(amount));
    }
  });

  it('Insufficient ERC20 balance', async () => {
    const amount = 2000;
    const receivers = Receivers.slice(0, 20).map(r => r.address);
    await expect(Batcher.batchTransfer(Token.address, receivers, amount)).to.be.revertedWith('Contract has insufficient balance.');

    expect(await Token.balanceOf(Batcher.address)).to.equal(BATCHER_BALANCE);
    
    for (let i = 0; i < receivers.length; i++) {
      expect(await Token.balanceOf(receivers[i])).to.equal(tokenBalances[i]);
    }
  });

  it('Insufficient native token balance', async () => {
    const amount = 2000;
    const receivers = Receivers.slice(0, 20).map(r => r.address);
    await expect(Batcher.batchTransfer(ethers.constants.AddressZero, receivers, amount)).to.be.revertedWith('Contract has insufficient balance.');

    expect(await Batcher.provider.getBalance(Batcher.address)).to.equal(BATCHER_BALANCE);
    
    for (let i = 0; i < receivers.length; i++) {
      expect(await Receivers[i].getBalance()).to.equal(nativeBalances[i]);
    }
  });

  it('Successful ERC20 withdrawal', async () => {
    await Batcher.withdraw(Token.address);
    
    expect(await Token.balanceOf(Batcher.address)).to.equal(0);
    expect(await Token.balanceOf(BatchTransferOwner.address)).to.equal(BATCHER_BALANCE);
  });

  it('Successful native token withdrawal', async () => {
    const ownerBalance = await BatchTransferOwner.getBalance();
    const batcherBalance = await Batcher.provider.getBalance(Batcher.address);
    
    const { gasLimit, gasPrice } = await Batcher.withdraw(ethers.constants.AddressZero);

    expect(await Batcher.provider.getBalance(Batcher.address)).to.equal(0);
    expect(await BatchTransferOwner.getBalance()).to.be.greaterThan(
      ownerBalance.add(batcherBalance).sub(gasLimit.mul(gasPrice?.toNumber() ?? 0))
    );
  });

  it('Unknown token', async () => {
    const TokenFactory = await ethers.getContractFactory('TestERC20');
    const UnknownToken = (await TokenFactory.connect(TokenAdmin).deploy(
      'Test',
      'TEST',
      TOKEN_CAP,
    )) as TestERC20;

    await expect(Batcher.withdraw(UnknownToken.address)).to.be.revertedWith('Contract has no balance of such token.');
  });
});
