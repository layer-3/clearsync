import { ethers } from 'hardhat';

async function main(): Promise<void> {
  const privKey = '004d5f8564328d8120fbe34d9da9438a87d16550745e2006e2e376488cbbeb6e';
  const wallet = new ethers.Wallet(privKey, ethers.provider);
  const signer = wallet.connect(ethers.provider);

  const bnBalance = await signer.getBalance();

  console.log('balance:', bnBalance.toString());
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
