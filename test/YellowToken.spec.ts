import { loadFixture } from '@nomicfoundation/hardhat-network-helpers';
import { expect } from 'chai';
import { ethers } from 'hardhat';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { YellowToken } from '../typechain-types';

const DECIMALS = 8;
const TOKEN_SUPPLY = 10_000_000_000;

describe('YellowToken', function () {
  let Deployer: SignerWithAddress;
  let Token: YellowToken;

  async function deployTokenFixture(): Promise<{ Token: YellowToken }> {
    const TokenFactory = await ethers.getContractFactory('YellowToken');
    const Token = (await TokenFactory.connect(Deployer).deploy(
      'Canary',
      'CANARY',
      TOKEN_SUPPLY,
    )) as YellowToken;

    return { Token };
  }

  before(async () => {
    [Deployer] = await ethers.getSigners();
  });

  beforeEach(async () => {
    ({ Token } = await loadFixture(deployTokenFixture));
  });

  describe('Deployment', () => {
    it('Correct name and symbol', async () => {
      expect(await Token.name()).to.equal('Canary');
      expect(await Token.symbol()).to.equal('CANARY');
    });

    it('Correct decimals', async () => {
      expect(await Token.decimals()).to.equal(DECIMALS);
    });

    it('Supply minted to deployer', async () => {
      expect(await Token.balanceOf(Deployer.address)).to.equal(TOKEN_SUPPLY);
    });
  });
});
