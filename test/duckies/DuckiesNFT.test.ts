import { expect } from 'chai';
import { ethers, upgrades } from 'hardhat';

import type { DuckiesNFT } from '../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

const API_BASE_URL = 'https://www.ynet-cat.uat.opendax.app/api/v3/public/nft/metadata/';

describe('NewDuckiesNFT', function () {
  let DuckiesNFT: DuckiesNFT;
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

    const DuckiesNFTFactory = await ethers.getContractFactory('DuckiesNFT');
    DuckiesNFT = (await upgrades.deployProxy(DuckiesNFTFactory, [])) as unknown as DuckiesNFT;
    await DuckiesNFT.deployed();
    await DuckiesNFT.setAPIBaseURL(API_BASE_URL);
  });

  describe('minting', () => {
    it.skip('correct tokenURI', async () => {
      await DuckiesNFT.mintPack(1);
      const duckie = (await DuckiesNFT.tokenIdToDuckie(0)) as unknown;
      expect(await DuckiesNFT.tokenURI(0)).to.equal(API_BASE_URL + duckie.gene.toString());
    });
  });
});
