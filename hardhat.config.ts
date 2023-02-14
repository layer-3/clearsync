import 'dotenv/config';

import '@nomicfoundation/hardhat-toolbox';
import '@nomiclabs/hardhat-ethers';
import 'solidity-docgen';
import type { HardhatUserConfig } from 'hardhat/config';

const ACCOUNTS = process.env.PRIVATE_KEY === undefined ? [] : [process.env.PRIVATE_KEY];
const ETHERSCAN_API_KEY = process.env.ETHERSCAN_API_KEY ?? '';
const POLYGONSCAN_API_KEY = process.env.POLYGONSCAN_API_KEY ?? '';

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: '0.8.16',
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
      {
        version: '0.8.17',
        settings: {
          optimizer: {
            enabled: true,
            runs: 20_000,
          },
          viaIR: true,
        },
      },
    ],
  },
  networks: {
    ethereum: {
      url: process.env.ETHEREUM_URL ?? '',
      accounts: ACCOUNTS,
    },
    goerli: {
      url: process.env.GOERLI_URL ?? '',
      accounts: ACCOUNTS,
    },
    polygon: {
      url: process.env.POLYGON_URL ?? '',
      accounts: ACCOUNTS,
    },
    mumbai: {
      url: process.env.MUMBAI_URL ?? '',
      accounts: ACCOUNTS,
    },
    generic: {
      url: process.env.GENERIC_URL ?? '',
      chainId: Number.parseInt(process.env.GENERIC_CHAIN_ID ?? '0'),
      gasPrice: process.env.GENERIC_GAS_PRICE
        ? Number.parseInt(process.env.GENERIC_GAS_PRICE)
        : 'auto',
      accounts: ACCOUNTS,
    },
  },
  docgen: {
    output: 'docs',
    pages: 'files',
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: 'USD',
  },
  etherscan: {
    apiKey: {
      mainnet: ETHERSCAN_API_KEY,
      goerli: ETHERSCAN_API_KEY,
      polygon: POLYGONSCAN_API_KEY,
      polygonMumbai: POLYGONSCAN_API_KEY,
    },
  },
};

export default config;
