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

  const DucklingsNFTFactory = await ethers.getContractFactory('DucklingsNFT');
  const DucklingsNFT = await upgrades.deployProxy(
    DucklingsNFTFactory,
    [process.env.DEPLOYED_SMART_CONTRACT],
    {
      kind: 'uups',
    },
  );

  await DucklingsNFT.deployed();

  console.log('DUCKLINGS NFT smart-contract is deployed to:', DucklingsNFT.address);
  await DucklingsNFT.setAPIBaseURL(
    'https://www.ynet-cat.uat.opendax.app/api/v3/public/nft/metadata/',
  );
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
