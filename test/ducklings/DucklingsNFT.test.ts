import { expect } from 'chai';
import { ethers, upgrades } from 'hardhat';

import type { Ducklings } from '../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

const API_BASE_URL = 'https://www.ynet-cat.uat.opendax.app/api/v3/public/nft/metadata/';

describe('Ducklings', function () {
  let Ducklings: Ducklings;
  let Owner: SignerWithAddress;
  let Signer: SignerWithAddress;

  before(async () => {
    [Owner, Signer] = await ethers.getSigners();
  });

  beforeEach(async () => {
    // const Duckies = await ethers.getContractFactory('DuckiesV3');
    // const duckies = await upgrades.deployProxy(Duckies, []);
    // const duckiesContractAddress = (await duckies.deployed()).address;
    // await duckies.setIssuer(signer.address);

    const DucklingsFactory = await ethers.getContractFactory('DucklingsNFT');
    Ducklings = (await upgrades.deployProxy(DucklingsFactory, [])) as unknown as Ducklings;
    await Ducklings.deployed();
    await Ducklings.setAPIBaseURL(API_BASE_URL);
  });

  describe('minting', () => {
    it.skip('correct tokenURI', async () => {
      await Ducklings.mintPack(1);
      const duckling = (await Ducklings.tokenIdToDuckie(0)) as unknown;
      expect(await Ducklings.tokenURI(0)).to.equal(API_BASE_URL + duckling.gene.toString());
    });
  });
});
