import '@nomicfoundation/hardhat-toolbox';
import { task } from 'hardhat/config';

import type { DucklingsV2 } from '../../typechain-types';

interface UpgradeArgs {
  ducklings: string;
}

task('upgradeDucklings', 'Perform an upgrade for Ducklings contract and reinitialize it')
  .addParam('ducklings', 'The address of Ducklings contract')
  .setAction(async (taskArgs: UpgradeArgs, { ethers, upgrades }) => {
    console.log('Upgrading DucklingsV1 to DucklingsV2...');
    const DucklingsV2Factory = await ethers.getContractFactory('DucklingsV2');
    const DucklingsV2 = (await upgrades.upgradeProxy(taskArgs.ducklings, DucklingsV2Factory, {
      kind: 'uups',
    })) as DucklingsV2;
    console.log('DucklingsV1 was upgraded to DucklingsV2');

    console.log('Reinitializing DucklingsV2...');
    await DucklingsV2.initializeV2();
    console.log('DucklingsV2 was reinitialized');
  });
