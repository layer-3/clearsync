import '@nomicfoundation/hardhat-toolbox';
import { task } from 'hardhat/config';

import type { DucklingsV1, DuckyFamilyV1 } from '../../typechain-types';

interface SetupNFTsArgs {
  ducklings: string;
  duckyFamily: string;
  apiBaseUrl: string;
  issuer: string;
}

const GAME_ROLE = '0x6a64baf327d646d1bca72653e2a075d15fd6ac6d8cbd7f6ee03fc55875e0fa88';

task('setupNFTs', 'Set apiBaseURL, grant role to game and sets issuer')
  .addParam('ducklings', 'The address of Ducklings contract')
  .addParam('duckyFamily', 'The address of DuckyFamily contract')
  .addParam('apiBaseUrl', 'ApiBaseURL to set to Ducklings contract')
  .addParam('issuer', 'Issuer address to set to DuckyFamily contract')
  .setAction(async (taskArgs: SetupNFTsArgs, { ethers }) => {
    const Ducklings = (await ethers.getContractAt(
      'DucklingsV1',
      taskArgs.ducklings,
    )) as DucklingsV1;
    await Ducklings.setAPIBaseURL(taskArgs.apiBaseUrl);
    console.log(`Ducklings: \`apiBaseURL\` was set to ${taskArgs.apiBaseUrl}`);

    const DuckyFamily = (await ethers.getContractAt(
      'DuckyFamilyV1',
      taskArgs.duckyFamily,
    )) as DuckyFamilyV1;
    await DuckyFamily.setIssuer(taskArgs.issuer);
    console.log(`DuckyFamily: \`issuer\` was set to ${taskArgs.issuer}`);

    await Ducklings.grantRole(GAME_ROLE, DuckyFamily.address);
    console.log(`Ducklings: \`GAME_ROLE\` was granted to DuckyFamily (${DuckyFamily.address})`);
  });
