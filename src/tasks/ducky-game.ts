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
