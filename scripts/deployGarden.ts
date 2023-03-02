import { ethers, upgrades } from 'hardhat';

import type { Garden } from '../typechain-types';

async function main(): Promise<void> {
  const issuerAddress = process.env.ISSUER_ADDRESS ?? '';

  const GardenFactory = await ethers.getContractFactory('Garden');
  const Garden = (await upgrades.deployProxy(GardenFactory, [], {
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
