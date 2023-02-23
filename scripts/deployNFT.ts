// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.

import { ethers, upgrades } from 'hardhat';

async function main() {
  const [deployer] = await ethers.getSigners();

  console.log('Deploying contracts with the account:', deployer.address);
  console.log('Account balance:', (await deployer.getBalance()).toString());

  const DuckiesNFTFactory = await ethers.getContractFactory('DuckiesNFT');
  const DuckiesNFT = await upgrades.deployProxy(
    DuckiesNFTFactory,
    [process.env.DEPLOYED_SMART_CONTRACT],
    {
      kind: 'uups',
    },
  );

  await DuckiesNFT.deployed();

  console.log('DUCKIES NFT smart-contract is deployed to:', DuckiesNFT.address);
  await DuckiesNFT.setAPIBaseURL(
    'https://www.ynet-cat.uat.opendax.app/api/v3/public/nft/metadata/',
  );
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
