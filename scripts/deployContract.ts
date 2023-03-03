import { ethers } from 'hardhat';

async function main(): Promise<void> {
  const [deployer] = await ethers.getSigners();

  console.log('Deploying contracts with the account:', deployer.address);
  const balanceBN = await deployer.getBalance();
  console.log('Account balance:', balanceBN.toString());

  const contractName = process.env.NAME ?? '';
  let args: unknown[] = [];
  if (process.env.ARGS) {
    args = process.env.ARGS.split(',').map((v) => v.trim());
    console.log(`args:`, args);
  }

  const ContractFactory = await ethers.getContractFactory(contractName);
  const Contract = await ContractFactory.deploy(...args);

  await Contract.deployed();

  console.log(`${contractName} deployed to ${Contract.address} with args ${args.toString()}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
