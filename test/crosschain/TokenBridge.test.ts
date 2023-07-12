import { expect } from 'chai';
import { ethers } from 'hardhat';
import { constants, utils } from 'ethers';

import { ACCOUNT_MISSING_ROLE } from '../helpers/common';

import type { LZEndpointMock, TokenBridge, YellowToken } from '../../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

const ZRO_PAYMENT_ADDRESS = constants.AddressZero;
const ADAPTER_PARAMS = '0x';

const BRIDGER_ROLE = utils.id('BRIDGER_ROLE');

describe('TokenBridge', function () {
  const chainId = 123;

  const startBalance = 1000;
  const amount = 100;

  let Bridger: SignerWithAddress;
  let OtherBridger: SignerWithAddress;
  let Someone: SignerWithAddress;

  let LZEndpointMock: LZEndpointMock;

  let RootToken: YellowToken;
  let ChildToken: YellowToken;

  let RootBridge: TokenBridge;
  let ChildBridge: TokenBridge;

  beforeEach(async () => {
    [Bridger, OtherBridger, Someone] = await ethers.getSigners();

    // create a LayerZero Endpoint mock for testing
    const LZEndpointMockFactory = await ethers.getContractFactory('LZEndpointMock');
    LZEndpointMock = (await LZEndpointMockFactory.deploy(chainId)) as LZEndpointMock;

    // create Root and Child tokens
    const YellowTokenFactory = await ethers.getContractFactory('YellowToken');
    RootToken = (await YellowTokenFactory.deploy('RootToken', 'RTN', 1_000_000)) as YellowToken;
    ChildToken = (await YellowTokenFactory.deploy('ChildToken', 'CTN', 1_000_000)) as YellowToken;

    // activate tokens
    await RootToken.activate(startBalance, Bridger.address);
    // NOTE: active ChildToken with non-zero premint to ease testing of bridging to Root
    await ChildToken.activate(startBalance, Bridger.address);

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

    // grant BridgerRole to OtherBridger
    await RootBridge.grantRole(BRIDGER_ROLE, OtherBridger.address);
    await ChildBridge.grantRole(BRIDGER_ROLE, OtherBridger.address);
  });

  describe('success', () => {
    it('bridge out Root: tokens are locked on Root and minted on Child', async () => {
      // initial balances
      expect(await RootToken.balanceOf(Bridger.address)).to.be.equal(startBalance);
      expect(await ChildToken.balanceOf(Bridger.address)).to.be.equal(startBalance);

      // approve tokens for RootBridge
      await RootToken.connect(Bridger).approve(RootBridge.address, amount);

      // bridge out
      await RootBridge.bridge(
        chainId,
        Bridger.address,
        amount,
        ZRO_PAYMENT_ADDRESS,
        ADAPTER_PARAMS,
        {
          value: ethers.utils.parseEther('0.5'),
        },
      );

      // decreased for Bridger, locked on RootBridge
      expect(await RootToken.balanceOf(Bridger.address)).to.be.equal(startBalance - amount);
      expect(await RootToken.balanceOf(RootBridge.address)).to.be.equal(amount);

      // increased for Bridger, no balance on ChildBridge
      expect(await ChildToken.balanceOf(Bridger.address)).to.be.equal(startBalance + amount);
      expect(await ChildToken.balanceOf(ChildBridge.address)).to.be.equal(0);
    });

    it('bridge out Child: tokens are burned on Child and unlocked on Root', async () => {
      // approve tokens for ChildBridge
      await ChildToken.connect(Bridger).approve(ChildBridge.address, amount);

      // bridge back
      await ChildBridge.bridge(
        chainId,
        Bridger.address,
        amount,
        ZRO_PAYMENT_ADDRESS,
        ADAPTER_PARAMS,
        {
          value: ethers.utils.parseEther('0.5'),
        },
      );

      // increased for Bridger, no balance on RootBridge
      expect(await RootToken.balanceOf(Bridger.address)).to.be.equal(startBalance);
      expect(await RootToken.balanceOf(RootBridge.address)).to.be.equal(0);

      // decreased for Bridger, no balance on ChildBridge
      expect(await ChildToken.balanceOf(Bridger.address)).to.be.equal(startBalance - amount);
      expect(await ChildToken.balanceOf(ChildBridge.address)).to.be.equal(0);
    });

    it('can bridge to other account', async () => {
      // approve tokens for RootBridge
      await RootToken.connect(Bridger).approve(RootBridge.address, amount);

      // bridge out
      await RootBridge.bridge(
        chainId,
        OtherBridger.address,
        amount,
        ZRO_PAYMENT_ADDRESS,
        ADAPTER_PARAMS,
        {
          value: ethers.utils.parseEther('0.5'),
        },
      );

      // decreased for Bridger, locked on RootBridge
      expect(await RootToken.balanceOf(Bridger.address)).to.be.equal(startBalance - amount);
      expect(await RootToken.balanceOf(RootBridge.address)).to.be.equal(amount);

      // increased for Bridger, no balance on ChildBridge
      expect(await ChildToken.balanceOf(OtherBridger.address)).to.be.equal(amount);
      expect(await ChildToken.balanceOf(ChildBridge.address)).to.be.equal(0);
    });

    it('events are emitted', async () => {
      // approve tokens for RootBridge
      await RootToken.connect(Bridger).approve(RootBridge.address, amount);

      // bridge out
      await expect(
        RootBridge.bridge(chainId, Bridger.address, amount, ZRO_PAYMENT_ADDRESS, ADAPTER_PARAMS, {
          value: ethers.utils.parseEther('0.5'),
        }),
      )
        .to.emit(RootBridge, 'BridgeOut')
        .withArgs(chainId, 1, Bridger.address, amount)
        .to.emit(ChildBridge, 'BridgeIn')
        .withArgs(chainId, 1, Bridger.address, amount);

      // approve tokens for ChildBridge
      await ChildToken.connect(Bridger).approve(ChildBridge.address, amount);

      // bridge back
      await expect(
        ChildBridge.bridge(chainId, Bridger.address, amount, ZRO_PAYMENT_ADDRESS, ADAPTER_PARAMS, {
          value: ethers.utils.parseEther('0.5'),
        }),
      )
        .to.emit(ChildBridge, 'BridgeOut')
        .withArgs(chainId, 1, Bridger.address, amount)
        .to.emit(RootBridge, 'BridgeIn')
        .withArgs(chainId, 1, Bridger.address, amount);
    });
  });

  describe('revert', () => {
    it('revert when bridging more than allowed', async () => {
      // approve tokens for RootBridge
      await RootToken.connect(Bridger).approve(RootBridge.address, amount);

      // bridge out
      await expect(
        RootBridge.bridge(
          chainId,
          Bridger.address,
          amount + 1,
          ZRO_PAYMENT_ADDRESS,
          ADAPTER_PARAMS,
          {
            value: ethers.utils.parseEther('0.5'),
          },
        ),
      ).to.be.reverted;
    });

    it('revert when bridging by not bridger', async () => {
      // approve tokens for RootBridge
      await RootToken.connect(Someone).approve(RootBridge.address, amount);

      // bridge out
      await expect(
        RootBridge.connect(Someone).bridge(
          chainId,
          Bridger.address,
          amount,
          ZRO_PAYMENT_ADDRESS,
          ADAPTER_PARAMS,
          {
            value: ethers.utils.parseEther('0.5'),
          },
        ),
      ).to.be.revertedWith(ACCOUNT_MISSING_ROLE(Someone.address, BRIDGER_ROLE));
    });

    it('revert when bridging to not bridger', async () => {
      // approve tokens for RootBridge
      await RootToken.connect(Bridger).approve(RootBridge.address, amount);

      // bridge out; the tx is not reverted, but rather stored in the application state
      await expect(
        RootBridge.connect(Bridger).bridge(
          chainId,
          Someone.address,
          amount,
          ZRO_PAYMENT_ADDRESS,
          ADAPTER_PARAMS,
          {
            value: ethers.utils.parseEther('0.5'),
          },
        ), // Note: event is emitted on the ChildBridge
      ).to.emit(ChildBridge, 'MessageFailed');
    });
  });
});
