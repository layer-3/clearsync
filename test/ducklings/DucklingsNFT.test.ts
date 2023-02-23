import { expect } from 'chai';
import { ethers, upgrades } from 'hardhat';

import type { DucklingsNFT } from '../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

const API_BASE_URL = 'https://www.ynet-cat.uat.opendax.app/api/v3/public/nft/metadata/';

describe('DucklingsNFT', function () {
  let DucklingsNFT: DucklingsNFT;
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

    const DucklingsNFTFactory = await ethers.getContractFactory('DucklingsNFT');
    DucklingsNFT = (await upgrades.deployProxy(DucklingsNFTFactory, [])) as unknown as DucklingsNFT;
    await DucklingsNFT.deployed();
    await DucklingsNFT.setAPIBaseURL(API_BASE_URL);
  });

  describe('minting', () => {
    it.skip('correct tokenURI', async () => {
      await DucklingsNFT.mintPack(1);
      const duckling = (await DucklingsNFT.tokenIdToDuckie(0)) as unknown;
      expect(await DucklingsNFT.tokenURI(0)).to.equal(API_BASE_URL + duckling.gene.toString());
    });
  });
});
