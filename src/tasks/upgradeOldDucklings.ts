import '@nomicfoundation/hardhat-toolbox';
import { task } from 'hardhat/config';

import type { OldDucklingsV2 } from '../../typechain-types';

interface UpgradeArgs {
  ducklings: string;
}

task('upgradeOldDucklings', 'Perform an upgrade for OldDucklings contract and reinitialize it')
  .addParam('ducklings', 'The address of Ducklings contract')
  .setAction(async (taskArgs: UpgradeArgs, { ethers, upgrades }) => {
    console.log('Upgrading OldDucklingsV1 to OldDucklingsV2...');
    const OldDucklingsV2Factory = await ethers.getContractFactory('OldDucklingsV2');
    const OldDucklingsV2 = (await upgrades.upgradeProxy(taskArgs.ducklings, OldDucklingsV2Factory, {
      kind: 'uups',
    })) as OldDucklingsV2;
    console.log('OldDucklingsV1 was upgraded to OldDucklingsV2');

    console.log('Reinitializing OldDucklingsV2...');
    await OldDucklingsV2.initializeV2();
    console.log('OldDucklingsV2 was reinitialized');
  });
