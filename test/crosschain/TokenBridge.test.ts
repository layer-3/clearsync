import { expect } from 'chai';
import { ethers } from 'hardhat';
import { constants } from 'ethers';

import type { LZEndpointMock, TokenBridge, YellowToken } from '../../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

const ZRO_PAYMENT_ADDRESS = constants.AddressZero;
const ADAPTER_PARAMS = '0x';

describe('TokenBridge', function () {
  const chainId = 123;

  const startRootBalance = 1000;
  const amount = 100;

  let Someone: SignerWithAddress;
  let Someother: SignerWithAddress;

  let LZEndpointMock: LZEndpointMock;

  let RootToken: YellowToken;
  let ChildToken: YellowToken;

  let RootBridge: TokenBridge;
  let ChildBridge: TokenBridge;

  beforeEach(async () => {
    [Someone, Someother] = await ethers.getSigners();

    // create a LayerZero Endpoint mock for testing
    const LZEndpointMockFactory = await ethers.getContractFactory('LZEndpointMock');
    LZEndpointMock = (await LZEndpointMockFactory.deploy(chainId)) as LZEndpointMock;

    // create Root and Child tokens
    const YellowTokenFactory = await ethers.getContractFactory('YellowToken');
    RootToken = (await YellowTokenFactory.deploy('RootToken', 'RTN', 1_000_000)) as YellowToken;
    ChildToken = (await YellowTokenFactory.deploy('ChildToken', 'CTN', 1_000_000)) as YellowToken;

    // activate tokens
    await RootToken.activate(startRootBalance, Someone.address);
    await ChildToken.activate(startRootBalance, Someone.address);
    // clean ChildToken balance
    await ChildToken.connect(Someone).burn(startRootBalance);

    // create Root and Child TokenBridges instances
    const TokenBridgeFactory = await ethers.getContractFactory('TokenBridge');
    RootBridge = (await TokenBridgeFactory.deploy(
      LZEndpointMock.address,
      RootToken.address,
      true,
    )) as TokenBridge;
    ChildBridge = (await TokenBridgeFactory.deploy(
      LZEndpointMock.address,
      ChildToken.address,
      false,
    )) as TokenBridge;

    // grant MinterRole to ChildBridge on ChildToken
    await ChildToken.grantRole(await ChildToken.MINTER_ROLE(), ChildBridge.address);

    await LZEndpointMock.setDestLzEndpoint(RootBridge.address, LZEndpointMock.address);
    await LZEndpointMock.setDestLzEndpoint(ChildBridge.address, LZEndpointMock.address);

    // set each contracts source address so it can send to each other
    await RootBridge.setTrustedRemoteAddress(chainId, ChildBridge.address);
    await ChildBridge.setTrustedRemoteAddress(chainId, RootBridge.address);
  });

  it('bridge out Root: tokens are locked on Root and minted on Child', async () => {
    // initial balances
    expect(await RootToken.balanceOf(Someone.address)).to.be.equal(startRootBalance);
    expect(await ChildToken.balanceOf(Someone.address)).to.be.equal(0);

    // approve tokens for RootBridge
    await RootToken.connect(Someone).approve(RootBridge.address, amount);

    // bridge out
    await RootBridge.bridge(chainId, Someone.address, amount, ZRO_PAYMENT_ADDRESS, ADAPTER_PARAMS, {
      value: ethers.utils.parseEther('0.5'),
    });

    // decreased for Someone, locked on RootBridge
    expect(await RootToken.balanceOf(Someone.address)).to.be.equal(startRootBalance - amount);
    expect(await RootToken.balanceOf(RootBridge.address)).to.be.equal(amount);

    // increased for Someone, no balance on ChildBridge
    expect(await ChildToken.balanceOf(Someone.address)).to.be.equal(amount);
    expect(await ChildToken.balanceOf(ChildBridge.address)).to.be.equal(0);
  });

  it('bridge out Child: tokens are burned on Child and unlocked on Root', async () => {
    // approve tokens for RootBridge
    await RootToken.connect(Someone).approve(RootBridge.address, amount);

    // bridge out
    await RootBridge.bridge(chainId, Someone.address, amount, ZRO_PAYMENT_ADDRESS, ADAPTER_PARAMS, {
      value: ethers.utils.parseEther('0.5'),
    });

    // approve tokens for ChildBridge
    await ChildToken.connect(Someone).approve(ChildBridge.address, amount);

    // bridge back
    await ChildBridge.bridge(
      chainId,
      Someone.address,
      amount,
      ZRO_PAYMENT_ADDRESS,
      ADAPTER_PARAMS,
      {
        value: ethers.utils.parseEther('0.5'),
      },
    );

    // increased for Someone, no balance on RootBridge
    expect(await RootToken.balanceOf(Someone.address)).to.be.equal(startRootBalance);
    expect(await RootToken.balanceOf(RootBridge.address)).to.be.equal(0);

    // decreased for Someone, no balance on ChildBridge
    expect(await ChildToken.balanceOf(Someone.address)).to.be.equal(0);
    expect(await ChildToken.balanceOf(ChildBridge.address)).to.be.equal(0);
  });

  it('can bridge to other account', async () => {
    // approve tokens for RootBridge
    await RootToken.connect(Someone).approve(RootBridge.address, amount);

    // bridge out
    await RootBridge.bridge(
      chainId,
      Someother.address,
      amount,
      ZRO_PAYMENT_ADDRESS,
      ADAPTER_PARAMS,
      {
        value: ethers.utils.parseEther('0.5'),
      },
    );

    // decreased for Someone, locked on RootBridge
    expect(await RootToken.balanceOf(Someone.address)).to.be.equal(startRootBalance - amount);
    expect(await RootToken.balanceOf(RootBridge.address)).to.be.equal(amount);

    // increased for Someone, no balance on ChildBridge
    expect(await ChildToken.balanceOf(Someother.address)).to.be.equal(amount);
    expect(await ChildToken.balanceOf(ChildBridge.address)).to.be.equal(0);
  });

  it('revert when bridging more than allowed', async () => {
    // approve tokens for RootBridge
    await RootToken.connect(Someone).approve(RootBridge.address, amount);

    // bridge out
    await expect(
      RootBridge.bridge(chainId, Someone.address, amount + 1, ZRO_PAYMENT_ADDRESS, ADAPTER_PARAMS, {
        value: ethers.utils.parseEther('0.5'),
      }),
    ).to.be.reverted;
  });
});
