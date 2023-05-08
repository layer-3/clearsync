import '@nomicfoundation/hardhat-toolbox';
import { task } from 'hardhat/config';

import type { OldDucklingsV2 } from '../../typechain-types';

interface UpgradeAndDestroyArgs {
  ducklings: string;
}

task('upgradeAndDestruct', 'Perform an upgrade for Ducklings contract, selfdestruct it afwerwards')
  .addParam('ducklings', 'The address of Ducklings contract')
  .setAction(async (taskArgs: UpgradeAndDestroyArgs, { ethers, upgrades }) => {
    console.log('Upgrading OldDucklingsV1 to OldDucklingsV2...');
    const OldDucklingsV2Factory = await ethers.getContractFactory('OldDucklingsV2');
    const OldDucklingsV2 = (await upgrades.upgradeProxy(taskArgs.ducklings, OldDucklingsV2Factory, {
      kind: 'uups',
    })) as OldDucklingsV2;
    console.log('OldDucklingsV1 was upgraded to OldDucklingsV2');

    console.log('Destructing OldDucklingsV2...');
    await OldDucklingsV2.destruct();
    console.log('OldDucklingsV2 was destructed');
  });
