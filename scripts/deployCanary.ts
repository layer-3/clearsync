import { ethers } from 'hardhat';

async function main(): Promise<void> {
  const [deployer] = await ethers.getSigners();

  console.log('Deploying contracts with the account:', deployer.address);
  const balanceBN = await deployer.getBalance();
  console.log('Account balance:', balanceBN.toString());

  const TokenFactory = await ethers.getContractFactory('YellowToken');
  const tx = TokenFactory.getDeployTransaction('Yellow Duckies', 'DUCKIES', 1_000_000_000n * 10n ** 8n);
  ethers.getDefaultProvider().estimateGas(tx).then((gasEstimate) => { console.log('Gas estimate:', gasEstimate.toString()) });
  const Token = await TokenFactory.deploy('Yellow Duckies', 'DUCKIES', 1_000_000_000n * 10n ** 8n);

  await Token.deployed();

  console.log(`Canary deployed to ${Token.address}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
