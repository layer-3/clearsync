import { ethers } from 'hardhat';

async function main(): Promise<void> {
  const TokenFactory = await ethers.getContractFactory('Token');
  const Token = await TokenFactory.deploy('Canary', 'CANARY', 8);

  await Token.deployed();

  console.log(`Canary deployed to ${Token.address}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
