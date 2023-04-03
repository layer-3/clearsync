import { ethers, upgrades } from 'hardhat';

async function main(): Promise<void> {
  const [deployer] = await ethers.getSigners();

  console.log('Upgrading smart contract with the account:', deployer.address);
  const balanceBN = await deployer.getBalance();
  console.log('Account balance:', balanceBN.toString());

  const upgradeToName = process.env.UPGRADE_TO_NAME ?? '';

  const Upgrade_To_Factory = await ethers.getContractFactory(upgradeToName);
  const address = process.env.ADDRESS ?? '';

  await upgrades.upgradeProxy(address, Upgrade_To_Factory, {
    kind: 'uups',
  });

  console.log(`Smart contract at ${address} was upgraded to ${upgradeToName}`);
}

main().catch((error) => {
  console.error(error);

  process.exitCode = 1;
});
