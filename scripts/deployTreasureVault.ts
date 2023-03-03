import { ethers, upgrades } from 'hardhat';

import type { TreasureVault } from '../typechain-types';

async function main(): Promise<void> {
  const [deployer] = await ethers.getSigners();

  console.log('Deploying contracts with the account:', deployer.address);
  const balanceBN = await deployer.getBalance();
  console.log('Account balance:', balanceBN.toString());

  //FIXME: automate probably address management and params
  const issuerAddress = process.env.ISSUER_ADDRESS ?? '';

  const TreasureVaultFactory = await ethers.getContractFactory('TreasureVault');
  const TreasureVault = (await upgrades.deployProxy(TreasureVaultFactory, [], {
    kind: 'uups',
  })) as TreasureVault;

  await TreasureVault.deployed();
  console.log(`TreasureVault deployed to ${TreasureVault.address}`);

  console.log('Setting Issuer to', issuerAddress);
  await TreasureVault.setIssuer(issuerAddress);
  console.log('Issuer set to', issuerAddress);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
