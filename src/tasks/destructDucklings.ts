import '@nomicfoundation/hardhat-toolbox';
import { task } from 'hardhat/config';

import type { DucklingsV2 } from '../../typechain-types';

interface DestructArgs {
  ducklings: string;
}

task('destructDucklings', 'Destruct Ducklings contract')
  .addParam('ducklings', 'The address of Ducklings contract')
  .setAction(async (taskArgs: DestructArgs, { ethers }) => {
    const DucklingsV2 = (await ethers.getContractAt(
      'DucklingsV2',
      taskArgs.ducklings,
    )) as DucklingsV2;

    console.log('Destructing DucklingsV2...');
    await DucklingsV2.destruct();
    console.log('DucklingsV2 was destructed');
  });
