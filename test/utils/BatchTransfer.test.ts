import { loadFixture } from '@nomicfoundation/hardhat-network-helpers';
import { expect } from 'chai';
import { ethers } from 'hardhat';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { BatchTransfer, TestERC20 } from '../../typechain-types';

describe('BatchTransfer', function () {
  const TOKEN_CAP = 100_000_000_000;
  const BATCHER_BALANCE = 10000;

  let TokenAdmin: SignerWithAddress;
  let BatchTransferOwner: SignerWithAddress;
  let Receivers: SignerWithAddress[];

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

    await Token.mint(Batcher.address, BATCHER_BALANCE);
  });

  it('Successful transfers', async () => {
    const amount = 20;
    const receivers = Receivers.slice(0, 20).map(r => r.address);
    await Batcher.batchTransfer(Token.address, receivers, amount);

    expect(await Token.balanceOf(Batcher.address)).to.equal(BATCHER_BALANCE - amount * receivers.length);
    
    for (const receiver of receivers) {
      expect(await Token.balanceOf(receiver)).to.equal(amount);
    }
  });

  it('Insufficient balance', async () => {
    const amount = 2000;
    const receivers = Receivers.slice(0, 20).map(r => r.address);
    await expect(Batcher.batchTransfer(Token.address, receivers, amount)).to.be.revertedWith('Contract has insufficient balance.');

    expect(await Token.balanceOf(Batcher.address)).to.equal(BATCHER_BALANCE);
    
    for (const receiver of receivers) {
      expect(await Token.balanceOf(receiver)).to.equal(0);
    }
  });

  it('Successful withdrawal', async () => {
    await Batcher.withdraw(Token.address);
    
    expect(await Token.balanceOf(Batcher.address)).to.equal(0);
    expect(await Token.balanceOf(BatchTransferOwner.address)).to.equal(BATCHER_BALANCE);
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
