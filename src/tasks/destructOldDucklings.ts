import '@nomicfoundation/hardhat-toolbox';
import { task } from 'hardhat/config';

import type { OldDucklingsV2 } from '../../typechain-types';

interface DestructArgs {
  ducklings: string;
}

task('destructOldDucklings', 'Destruct OldDucklings contract')
  .addParam('ducklings', 'The address of Ducklings contract')
  .setAction(async (taskArgs: DestructArgs, { ethers }) => {
    const OldDucklingsV2 = (await ethers.getContractAt(
      'OldDucklingsV2',
      taskArgs.ducklings,
    )) as OldDucklingsV2;

    console.log('Destructing OldDucklingsV2...');
    await OldDucklingsV2.destruct();
    console.log('OldDucklingsV2 was destructed');
  });
