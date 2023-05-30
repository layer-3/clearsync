import '@nomicfoundation/hardhat-toolbox';
import { task } from 'hardhat/config';

import type { DuckyFamilyV1 } from '../../typechain-types';

interface SetMintPriceArgs {
  duckyFamily: string;
  price: number;
}

task('setMintPrice', 'Set the price of minting one Duckling')
  .addParam('duckyFamily', 'Address of DuckyFamily contract')
  .addParam('price', 'Price for minting one Duckling in Duckies')
  .setAction(async (taskArgs: SetMintPriceArgs, { ethers }) => {
    console.log(
      'Working on chain#',
      await ethers.provider.getNetwork().then((network) => network.chainId),
    );

    const DuckyFamily = (await ethers.getContractAt(
      'DuckyFamilyV1',
      taskArgs.duckyFamily,
    )) as DuckyFamilyV1;

    await DuckyFamily.setMintPrice(taskArgs.price);
    console.log(`DuckyFamily: \`mintPrice\` was set to ${taskArgs.price} Duckies`);
  });
