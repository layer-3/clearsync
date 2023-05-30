import '@nomicfoundation/hardhat-toolbox';
import { task } from 'hardhat/config';

import type { DuckyFamilyV1 } from '../../typechain-types';

interface SetMeldPricesArgs {
  duckyFamily: string;
  prices: string;
}

task('setMeldPrices', 'Set the price for melding Ducklings of each rarity')
  .addParam('duckyFamily', 'Address of DuckyFamily contract')
  .addParam(
    'prices',
    'Price for melding Duckling of each rarity (Common, Rare, Epic, Legendary), separated by comma',
  )
  .setAction(async (taskArgs: SetMeldPricesArgs, { ethers }) => {
    console.log(
      'Working on chain#',
      await ethers.provider.getNetwork().then((network) => network.chainId),
    );

    const DuckyFamily = (await ethers.getContractAt(
      'DuckyFamilyV1',
      taskArgs.duckyFamily,
    )) as DuckyFamilyV1;

    const prices = taskArgs.prices.split(',').map(Number) as [number, number, number, number];

    await DuckyFamily.setMeldPrices(prices);
    console.log(
      `DuckyFamily: \`meldPrices\` were set to
      Common: ${prices[0]}
      Rare: ${prices[1]}
      Epic: ${prices[2]}
      Legendary: ${prices[3]}
      Duckies`,
    );
  });
