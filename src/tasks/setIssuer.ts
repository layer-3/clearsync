import { task } from 'hardhat/config';

import type { TreasureVault } from '../../typechain-types';

interface TaskArgs {
  contract: string;
  account: string;
}

// FIXME: ts errors and test it
task('setIssuer', 'Set TokenContract Reward Issuer')
  .addParam('contract', 'The smart-contract address')
  .addParam('account', "The issuer account's address")
  .setAction(async (taskArgs: TaskArgs, hre) => {
    const [deployer] = await hre.ethers.getSigners();
    const TreasureVault = (await hre.ethers.getContractAt(
      'TreasureVault',
      taskArgs.contract,
      deployer,
    )) as unknown as TreasureVault;

    console.log('Setting issuer with the account:', deployer.address);
    console.log('Account balance:', (await deployer.getBalance()).toString());

    const prevIssuer = await TreasureVault.issuer();
    console.log(`Previous issuer was ${prevIssuer}, Updating...`);

    await TreasureVault.setIssuer(taskArgs.account);
    const newIssuer = await TreasureVault.issuer();
    console.log(`New issuer is ${newIssuer}`);
  });
