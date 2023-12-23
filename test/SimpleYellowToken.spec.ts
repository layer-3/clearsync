import { loadFixture } from '@nomicfoundation/hardhat-network-helpers';
import { expect } from 'chai';
import { ethers } from 'hardhat';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { SimpleYellowToken } from '../typechain-types';

const DECIMALS = 8;
const TOKEN_SUPPLY_CAP = 10_000_000_000;

describe('SimpleYellowToken', function () {
  let Deployer: SignerWithAddress;
  let Token: SimpleYellowToken;

  async function deployTokenFixture(): Promise<{ Token: SimpleYellowToken }> {
    const TokenFactory = await ethers.getContractFactory('SimpleYellowToken');
    const Token = (await TokenFactory.connect(Deployer).deploy(
      'Canary',
      'CANARY',
      TOKEN_SUPPLY_CAP,
    )) as SimpleYellowToken;

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

    it('Correct supply cap', async () => {
      expect(await Token.cap()).to.equal(TOKEN_SUPPLY_CAP);
    });

    it('Supply cap minted to deployer', async () => {
      expect(await Token.balanceOf(Deployer.address)).to.equal(TOKEN_SUPPLY_CAP);
    });
  });
});
