import { expect } from 'chai';
import { ethers } from 'hardhat';
import { constants, utils } from 'ethers';

import { ACCOUNT_MISSING_ROLE, encodeError } from '../helpers/common';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { LZEndpointMock, TokenBridge, YellowToken } from '../../typechain-types';

const ZRO_PAYMENT_ADDRESS = constants.AddressZero;
const ADAPTER_PARAMS = '0x';

const ADMIN_ROLE = constants.HashZero;
const BRIDGER_ROLE = utils.id('BRIDGER_ROLE');
const BRIDGING_UNAUTHORIZED = 'BridgingUnauthorized';

/*
describe('TokenBridge', function () {
  const chainId = 123;

  const startBalance = 1000;
  const amount = 100;

  let Admin: SignerWithAddress;
  let Bridger: SignerWithAddress;
  let OtherBridger: SignerWithAddress;
  let RootTokenBridger: SignerWithAddress;
  let Someone: SignerWithAddress;

  let LZEndpointMock: LZEndpointMock;

  let RootBridge: TokenBridge;
  let ChildBridge: TokenBridge;

  let RootToken: YellowToken;
  let ChildToken: YellowToken;

  beforeEach(async () => {
    [Admin, Bridger, OtherBridger, RootTokenBridger, Someone] = await ethers.getSigners();

    // create a LayerZero Endpoint mock for testing
    const LZEndpointMockFactory = await ethers.getContractFactory('LZEndpointMock');
    LZEndpointMock = (await LZEndpointMockFactory.deploy(chainId)) as LZEndpointMock;

    // create Root and Child TokenBridges instances
    const TokenBridgeFactory = await ethers.getContractFactory('TokenBridge');
    RootBridge = (await TokenBridgeFactory.deploy(LZEndpointMock.address)) as TokenBridge;
    ChildBridge = (await TokenBridgeFactory.deploy(LZEndpointMock.address)) as TokenBridge;

    // set each contracts source address so it can send to each other
    await RootBridge.setTrustedRemoteAddress(chainId, ChildBridge.address);
    await ChildBridge.setTrustedRemoteAddress(chainId, RootBridge.address);

    // grant BridgerRole to Bridger and OtherBridger
    await RootBridge.grantRole(BRIDGER_ROLE, Bridger.address);
    await RootBridge.grantRole(BRIDGER_ROLE, OtherBridger.address);
    await ChildBridge.grantRole(BRIDGER_ROLE, Bridger.address);
    await ChildBridge.grantRole(BRIDGER_ROLE, OtherBridger.address);

    await LZEndpointMock.setDestLzEndpoint(RootBridge.address, LZEndpointMock.address);
    await LZEndpointMock.setDestLzEndpoint(ChildBridge.address, LZEndpointMock.address);

    // create Root and Child tokens
    const YellowTokenFactory = await ethers.getContractFactory('YellowToken');
    RootToken = (await YellowTokenFactory.deploy('RootToken', 'RTN', 1_000_000)) as YellowToken;
    ChildToken = (await YellowTokenFactory.deploy('ChildToken', 'CTN', 1_000_000)) as YellowToken;

    // activate tokens
    await RootToken.activate(startBalance, Bridger.address);
    // NOTE: burn ChildToken as it should appear only after bridging
    await ChildToken.activate(startBalance, Bridger.address);
    await ChildToken.connect(Bridger).burn(startBalance);

    const encodedRootTokenAddress = utils.defaultAbiCoder.encode(['address'], [RootToken.address]);

    // grant RootTokenBridger corresponding role
    await RootBridge.grantRole(utils.keccak256(encodedRootTokenAddress), RootTokenBridger.address);
    await ChildBridge.grantRole(utils.id(ChildToken.address), RootTokenBridger.address);
  });

  describe('view', () => {
    describe('getDstToken', () => {
      it('return correct dst token', async () => {
        await RootBridge.addToken(RootToken.address, true);
        await RootBridge.setDstToken(RootToken.address, chainId, ChildToken.address);

        expect(await RootBridge.getDstToken(RootToken.address, chainId)).to.equal(
          ChildToken.address,
        );
      });

      it('revert if token is not supported', async () => {
        await expect(
          RootBridge.getDstToken(RootToken.address, chainId),
        ).to.be.revertedWithCustomError(RootBridge, 'TokenNotSupported');
      });
    });

    describe('estimateFees', () => {
      it('can estimate fees', async () => {
        await RootBridge.addToken(RootToken.address, true);
        await RootBridge.setDstToken(RootToken.address, chainId, ChildToken.address);

        await ChildBridge.addToken(ChildToken.address, false);

        await expect(
          RootBridge.estimateFees(
            chainId,
            RootToken.address,
            Admin.address,
            amount,
            false,
            ADAPTER_PARAMS,
          ),
        ).to.not.be.reverted;
      });

      it('does not revert if not correct inputs', async () => {
        await expect(
          RootBridge.estimateFees(
            chainId,
            RootToken.address,
            Admin.address,
            amount,
            false,
            ADAPTER_PARAMS,
          ),
        ).to.not.be.reverted;
      });
    });
  });

  describe('modification', () => {
    describe('addToken', () => {
      it('admin can add token', async () => {
        await RootBridge.connect(Admin).addToken(RootToken.address, true);

        const [isSupported] = await RootBridge.tokensLookup(RootToken.address);
        expect(isSupported).to.be.true;
      });

      it('revert on not admin adding token', async () => {
        await expect(
          RootBridge.connect(Someone).addToken(RootToken.address, true),
        ).to.be.revertedWith(ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE));
      });

      it('revert on adding address(0)', async () => {
        await expect(RootBridge.addToken(constants.AddressZero, true))
          .to.be.revertedWithCustomError(RootBridge, 'InvalidToken')
          .withArgs(constants.AddressZero);
      });

      it('revert on adding already supported token', async () => {
        await RootBridge.connect(Admin).addToken(RootToken.address, true);

        await expect(RootBridge.addToken(RootToken.address, true))
          .to.be.revertedWithCustomError(RootBridge, 'TokenAlreadySupported')
          .withArgs(RootToken.address);
      });

      it('event is emitted', async () => {
        const isRoot = true;
        await expect(RootBridge.addToken(RootToken.address, isRoot))
          .to.emit(RootBridge, 'TokenAdded')
          .withArgs(RootToken.address, isRoot);
      });
    });

    describe('removeToken', () => {
      beforeEach(async () => {
        await RootBridge.addToken(RootToken.address, true);
      });

      it('admin can remove token', async () => {
        await RootBridge.removeToken(RootToken.address);

        const [isSupported] = await RootBridge.tokensLookup(RootToken.address);
        expect(isSupported).to.be.false;
      });

      it('revert on not admin removing token', async () => {
        await expect(RootBridge.connect(Someone).removeToken(RootToken.address)).to.be.revertedWith(
          ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE),
        );
      });

      it('revert on removing not supported token', async () => {
        await expect(RootBridge.removeToken(ChildToken.address))
          .to.be.revertedWithCustomError(RootBridge, 'TokenNotSupported')
          .withArgs(ChildToken.address);
      });

      it('event is emitted', async () => {
        await expect(RootBridge.removeToken(RootToken.address))
          .to.emit(RootBridge, 'TokenRemoved')
          .withArgs(RootToken.address);
      });
    });

    describe('setDstToken', () => {
      beforeEach(async () => {
        await RootBridge.addToken(RootToken.address, true);
      });

      it('adming can add non-zero dst token', async () => {
        await RootBridge.connect(Admin).setDstToken(RootToken.address, chainId, ChildToken.address);
        expect(await RootBridge.getDstToken(RootToken.address, chainId)).to.equal(
          ChildToken.address,
        );
      });

      it('adming can add zero dst token', async () => {
        await RootBridge.connect(Admin).setDstToken(RootToken.address, chainId, ChildToken.address);
        expect(await RootBridge.getDstToken(RootToken.address, chainId)).to.equal(
          ChildToken.address,
        );

        await RootBridge.connect(Admin).setDstToken(
          RootToken.address,
          chainId,
          constants.AddressZero,
        );
        expect(await RootBridge.getDstToken(RootToken.address, chainId)).to.equal(
          constants.AddressZero,
        );
      });

      it('revert on not admin setting dst token', async () => {
        await expect(
          RootBridge.connect(Someone).setDstToken(RootToken.address, chainId, ChildToken.address),
        ).to.be.revertedWith(ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE));
      });

      it('revert on adding dst token to not supported token', async () => {
        await expect(RootBridge.setDstToken(ChildToken.address, chainId, RootToken.address))
          .to.be.revertedWithCustomError(RootBridge, 'TokenNotSupported')
          .withArgs(ChildToken.address);
      });

      it('event is emitted', async () => {
        await expect(RootBridge.setDstToken(RootToken.address, chainId, ChildToken.address))
          .to.emit(RootBridge, 'DstTokenSet')
          .withArgs(RootToken.address, chainId, ChildToken.address);
      });
    });
  });

  describe('bridge', () => {
    beforeEach(async () => {
      // add Root and Child Tokens
      await RootBridge.addToken(RootToken.address, true);
      await ChildBridge.addToken(ChildToken.address, false);

      await RootBridge.setDstToken(RootToken.address, chainId, ChildToken.address);
      await ChildBridge.setDstToken(ChildToken.address, chainId, RootToken.address);

      // grant MinterRole to ChildBridge on ChildToken
      await ChildToken.grantRole(await ChildToken.MINTER_ROLE(), ChildBridge.address);
    });

    describe('success', () => {
      it('bridger role can bridge', async () => {
        // initial balances
        expect(await RootToken.balanceOf(Bridger.address)).to.be.equal(startBalance);
        expect(await ChildToken.balanceOf(Bridger.address)).to.be.equal(0);

        // approve tokens for RootBridge
        await RootToken.connect(Bridger).approve(RootBridge.address, amount);

        // bridge out
        await RootBridge.connect(Bridger).bridge(
          chainId,
          RootToken.address,
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
        expect(await ChildToken.balanceOf(Bridger.address)).to.be.equal(amount);
        expect(await ChildToken.balanceOf(ChildBridge.address)).to.be.equal(0);
      });

      it('briding lock tokens on RootBridge, does not mint to ChildBridge', async () => {
        // approve tokens for RootBridge
        await RootToken.connect(Bridger).approve(RootBridge.address, amount);

        // bridge out
        await RootBridge.connect(Bridger).bridge(
          chainId,
          RootToken.address,
          Bridger.address,
          amount,
          ZRO_PAYMENT_ADDRESS,
          ADAPTER_PARAMS,
          {
            value: ethers.utils.parseEther('0.5'),
          },
        );

        // locked on RootBridge
        expect(await RootToken.balanceOf(RootBridge.address)).to.be.equal(amount);
        // no balance on ChildBridge
        expect(await ChildToken.balanceOf(ChildBridge.address)).to.be.equal(0);
      });

      it('token bridger can bridge their token', async () => {
        // RootTokenBridger does not have a BRIDGER_ROLE, but a `keccak256(RootToken.address)` role

        await RootToken.mint(RootTokenBridger.address, startBalance);

        // initial balances
        expect(
          await RootToken.connect(RootTokenBridger).balanceOf(RootTokenBridger.address),
        ).to.be.equal(startBalance);

        // approve tokens for RootBridge
        await RootToken.connect(RootTokenBridger).approve(RootBridge.address, amount);

        // bridge out
        await RootBridge.connect(RootTokenBridger).bridge(
          chainId,
          RootToken.address,
          Bridger.address,
          amount,
          ZRO_PAYMENT_ADDRESS,
          ADAPTER_PARAMS,
          {
            value: ethers.utils.parseEther('0.5'),
          },
        );

        // decreased for RootTokenBridger
        expect(await RootToken.balanceOf(RootTokenBridger.address)).to.be.equal(
          startBalance - amount,
        );
        // increased for Bridger
        expect(await ChildToken.balanceOf(Bridger.address)).to.be.equal(amount);
      });

      it('bridge out Child: tokens are burned on Child and unlocked on Root', async () => {
        // mint start balance on ChildToken to ease testing
        await ChildToken.mint(Bridger.address, startBalance);

        // approve tokens for ChildBridge
        await ChildToken.connect(Bridger).approve(ChildBridge.address, amount);

        // bridge back
        await ChildBridge.connect(Bridger).bridge(
          chainId,
          ChildToken.address,
          Bridger.address,
          amount,
          ZRO_PAYMENT_ADDRESS,
          ADAPTER_PARAMS,
          {
            value: ethers.utils.parseEther('0.5'),
          },
        );

        // increased for Bridger
        expect(await RootToken.balanceOf(Bridger.address)).to.be.equal(startBalance);
        // decreased for Bridger
        expect(await ChildToken.balanceOf(Bridger.address)).to.be.equal(startBalance - amount);
      });

      it('can bridge to other account', async () => {
        // approve tokens for RootBridge
        await RootToken.connect(Bridger).approve(RootBridge.address, amount);

        // bridge out
        await RootBridge.connect(Bridger).bridge(
          chainId,
          RootToken.address,
          OtherBridger.address,
          amount,
          ZRO_PAYMENT_ADDRESS,
          ADAPTER_PARAMS,
          {
            value: ethers.utils.parseEther('0.5'),
          },
        );

        // decreased for Bridger
        expect(await RootToken.balanceOf(Bridger.address)).to.be.equal(startBalance - amount);
        // increased for OtherBridger
        expect(await ChildToken.balanceOf(OtherBridger.address)).to.be.equal(amount);
      });

      it('events are emitted', async () => {
        // approve tokens for RootBridge
        await RootToken.connect(Bridger).approve(RootBridge.address, amount);

        // bridge out
        await expect(
          RootBridge.connect(Bridger).bridge(
            chainId,
            RootToken.address,
            Bridger.address,
            amount,
            ZRO_PAYMENT_ADDRESS,
            ADAPTER_PARAMS,
            {
              value: ethers.utils.parseEther('0.5'),
            },
          ),
        )
          .to.emit(RootBridge, 'BridgeOut')
          .withArgs(chainId, 1, RootToken.address, Bridger.address, amount)
          .to.emit(ChildBridge, 'BridgeIn')
          .withArgs(chainId, 1, ChildToken.address, Bridger.address, amount);

        // approve tokens for ChildBridge
        await ChildToken.connect(Bridger).approve(ChildBridge.address, amount);

        // bridge back
        await expect(
          ChildBridge.connect(Bridger).bridge(
            chainId,
            ChildToken.address,
            Bridger.address,
            amount,
            ZRO_PAYMENT_ADDRESS,
            ADAPTER_PARAMS,
            {
              value: ethers.utils.parseEther('0.5'),
            },
          ),
        )
          .to.emit(ChildBridge, 'BridgeOut')
          .withArgs(chainId, 1, ChildToken.address, Bridger.address, amount)
          .to.emit(RootBridge, 'BridgeIn')
          .withArgs(chainId, 1, RootToken.address, Bridger.address, amount);
      });
    });

    describe('revert', () => {
      it('revert when bridging not supported token', async () => {
        await RootBridge.removeToken(RootToken.address);

        // bridge out
        await expect(
          RootBridge.connect(Bridger).bridge(
            chainId,
            RootToken.address,
            Bridger.address,
            amount + 1,
            ZRO_PAYMENT_ADDRESS,
            ADAPTER_PARAMS,
            {
              value: ethers.utils.parseEther('0.5'),
            },
          ),
        )
          .to.be.revertedWithCustomError(RootBridge, 'TokenNotSupported')
          .withArgs(RootToken.address);
      });

      it('revert when no dst token set', async () => {
        await RootBridge.setDstToken(RootToken.address, chainId, constants.AddressZero);

        // bridge out
        await expect(
          RootBridge.connect(Bridger).bridge(
            chainId,
            RootToken.address,
            Bridger.address,
            amount + 1,
            ZRO_PAYMENT_ADDRESS,
            ADAPTER_PARAMS,
            {
              value: ethers.utils.parseEther('0.5'),
            },
          ),
        )
          .to.be.revertedWithCustomError(RootBridge, 'NoDstToken')
          .withArgs(RootToken.address, chainId);
      });

      it('revert when bridging more than allowed', async () => {
        // approve tokens for RootBridge
        await RootToken.connect(Bridger).approve(RootBridge.address, amount);

        // bridge out
        await expect(
          RootBridge.bridge(
            chainId,
            RootToken.address,
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

      it('revert when bridging by not bridger role', async () => {
        // approve tokens for RootBridge
        await RootToken.connect(Someone).approve(RootBridge.address, amount);

        // bridge out
        await expect(
          RootBridge.connect(Someone).bridge(
            chainId,
            RootToken.address,
            Bridger.address,
            amount,
            ZRO_PAYMENT_ADDRESS,
            ADAPTER_PARAMS,
            {
              value: ethers.utils.parseEther('0.5'),
            },
          ),
        )
          .to.be.revertedWithCustomError(RootBridge, BRIDGING_UNAUTHORIZED)
          .withArgs(Someone.address, RootToken.address);
      });

      it('revert when bridging by token bridger of other token', async () => {
        const encodedChildTokenAddress = utils.defaultAbiCoder.encode(
          ['address'],
          [ChildToken.address],
        );
        await RootToken.grantRole(utils.keccak256(encodedChildTokenAddress), Someone.address);

        // approve tokens for RootBridge
        await RootToken.connect(Someone).approve(RootBridge.address, amount);

        // bridge out
        await expect(
          RootBridge.connect(Someone).bridge(
            chainId,
            RootToken.address,
            Bridger.address,
            amount,
            ZRO_PAYMENT_ADDRESS,
            ADAPTER_PARAMS,
            {
              value: ethers.utils.parseEther('0.5'),
            },
          ),
        )
          .to.be.revertedWithCustomError(RootBridge, BRIDGING_UNAUTHORIZED)
          .withArgs(Someone.address, RootToken.address);
      });

      it('revert when bridging to not bridger role', async () => {
        // approve tokens for RootBridge
        await RootToken.connect(Bridger).approve(RootBridge.address, amount);

        const path = (RootBridge.address + ChildBridge.address.slice(2)).toLowerCase();

        const nonce = 1;

        const payload = utils.defaultAbiCoder.encode(
          ['address', 'address', 'uint256'],
          [ChildToken.address, Someone.address, amount],
        );

        const encodedError = encodeError(
          'BridgingUnauthorized(address,address)',
          Someone.address,
          ChildToken.address,
        );

        // bridge out; the tx is not reverted, but rather stored in the application state
        await expect(
          RootBridge.connect(Bridger).bridge(
            chainId,
            RootToken.address,
            Someone.address,
            amount,
            ZRO_PAYMENT_ADDRESS,
            ADAPTER_PARAMS,
            {
              value: ethers.utils.parseEther('0.5'),
            },
          ), // Note: event is emitted on the ChildBridge
        )
          .to.emit(ChildBridge, 'MessageFailed')
          .withArgs(chainId, path, nonce, payload, encodedError);
      });

      it('revert when bridging to token bridger of other token', async () => {
        const encodedRootTokenAddress = utils.defaultAbiCoder.encode(
          ['address'],
          [RootToken.address],
        );
        await RootToken.grantRole(utils.keccak256(encodedRootTokenAddress), Someone.address);

        // approve tokens for RootBridge
        await RootToken.connect(Bridger).approve(RootBridge.address, amount);

        const path = (RootBridge.address + ChildBridge.address.slice(2)).toLowerCase();

        const nonce = 1;

        const payload = utils.defaultAbiCoder.encode(
          ['address', 'address', 'uint256'],
          [ChildToken.address, Someone.address, amount],
        );

        const encodedError = encodeError(
          'BridgingUnauthorized(address,address)',
          Someone.address,
          ChildToken.address,
        );

        // bridge out
        await expect(
          RootBridge.connect(Bridger).bridge(
            chainId,
            RootToken.address,
            Someone.address,
            amount,
            ZRO_PAYMENT_ADDRESS,
            ADAPTER_PARAMS,
            {
              value: ethers.utils.parseEther('0.5'),
            },
          ),
        )
          .to.emit(ChildBridge, 'MessageFailed')
          .withArgs(chainId, path, nonce, payload, encodedError);
      });
    });
  });
});
*/
