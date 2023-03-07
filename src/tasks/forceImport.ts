import { task } from 'hardhat/config';

interface ForceImportArgs {
  name: string;
  address: string;
}

task('forceImport', 'Forcefully imports openzeppelin-upgradeable artifacts')
  .addParam('name', 'Name of the contract to import artifacts of')
  .addParam('address', 'Address of the contract to import artifacts of')
  .setAction(async (taskArgs: ForceImportArgs, { ethers, upgrades }) => {
    const Contract = await ethers.getContractFactory(taskArgs.name);

    await upgrades.forceImport(taskArgs.address, Contract, { kind: 'uups' });

    console.log('Force import is finished');
  });
