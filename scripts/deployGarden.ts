import { ethers, upgrades } from 'hardhat';

import type { Garden } from '../typechain-types';

async function main(): Promise<void> {
  const duckiesAddress = process.env.DUCKIES_ADDRESS ?? '';
  const issuerAddress = process.env.ISSUER_ADDRESS ?? '';

  const GardenFactory = await ethers.getContractFactory('Garden');
  console.log('Deploying Garden with Duckies address:', duckiesAddress);
  const Garden = (await upgrades.deployProxy(GardenFactory, [duckiesAddress], {
    kind: 'uups',
  })) as Garden;

  await Garden.deployed();
  console.log(`Garden deployed to ${Garden.address}`);

  console.log('Setting Issuer to', issuerAddress);
  await Garden.setIssuer(issuerAddress);
  console.log('Issuer set to', issuerAddress);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
