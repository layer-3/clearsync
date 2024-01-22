import 'dotenv/config';

import '@nomicfoundation/hardhat-toolbox';
import '@nomiclabs/hardhat-ethers';
import type { HardhatUserConfig } from 'hardhat/config';
import 'solidity-docgen';
import 'hardhat-dependency-compiler';

import './src/tasks/accounts';
import './src/tasks/vesting';

let accounts;

if (process.env.MNEMONIC) {
  accounts = {
    mnemonic: process.env.MNEMONIC,
    initialIndex: 0,
    count: 10,
  };
} else if (process.env.PRIVATE_KEY) {
  accounts = [process.env.PRIVATE_KEY];
}

const ETHERSCAN_API_KEY = process.env.ETHERSCAN_API_KEY ?? '';
const POLYGONSCAN_API_KEY = process.env.POLYGONSCAN_API_KEY ?? '';
const BSCSCAN_API_KEY = process.env.BSCSCAN_API_KEY ?? '';

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: '0.8.22',
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    ],
  },
  typechain: {
    outDir: 'typechain-types',
    target: 'ethers-v5',
  },
  networks: {
    hardhat: {
      accounts: {
        mnemonic: 'bring tumble anger wild frame you famous usage ramp federal captain company',
        count: 100,
      },
    },
    localdev: {
      url: process.env.LOCALDEV_URL ?? 'http://localhost:8545',
      accounts: {
        mnemonic: process.env.MNEMONIC ?? '',
        count: 10,
      },
    },
    kayenuat: {
      url: process.env.KAYENUAT_URL ?? '',
      accounts: {
        mnemonic: process.env.MNEMONIC ?? '',
        count: 10,
      },
    },
    ethereum: {
      url: process.env.ETHEREUM_URL ?? '',
      accounts,
    },
    goerli: {
      url: process.env.GOERLI_URL ?? '',
      accounts,
    },
    sepolia: {
      url: process.env.SEPOLIA_URL ?? '',
      accounts,
    },
    matic: {
      url: process.env.POLYGON_URL ?? '',
      gasPrice: 2e11,
      accounts,
    },
    mumbai: {
      url: process.env.MUMBAI_URL ?? '',
      accounts,
    },
    bsc: {
      url: process.env.BSC_URL ?? '',
      accounts,
    },
    'bsc-testnet': {
      url: process.env.BSC_TESTNET_URL ?? '',
      accounts,
    },
    avalanche: {
      url: process.env.AVALANCHE_URL ?? '',
      accounts,
    },
    'avalanche-fuji': {
      url: process.env.AVALANCHE_FUJI_URL ?? '',
      accounts,
    },
    arbitrum: {
      url: process.env.ARBITRUM_URL ?? '',
      accounts,
    },
    'arbitrum-goerli': {
      url: process.env.ARBITRUM_GOERLI_URL ?? '',
      accounts,
    },
    optimism: {
      url: process.env.OPTIMISM_URL ?? '',
      accounts,
    },
    'optimism-goerli': {
      url: process.env.OPTIMISM_GOERLI_URL ?? '',
      accounts,
    },
    fantom: {
      url: process.env.FANTOM_URL ?? '',
      accounts,
    },
    'fantom-testnet': {
      url: process.env.FANTOM_TESTNET_URL ?? '',
      accounts,
    },
    zkEVM: {
      url: process.env.ZKEVM_URL ?? '',
      accounts,
    },
    'zkEVM-testnet': {
      url: process.env.ZKEVM_TESTNET_URL ?? '',
      accounts,
    },
    zkSync: {
      url: process.env.ZKSYNC_URL ?? '',
      accounts,
    },
    'zkSync-testnet': {
      url: process.env.ZKSYNC_TESTNET_URL ?? '',
      accounts,
    },
    linea: {
      url: process.env.LINEA_URL ?? '',
      accounts,
    },
    'linea-testnet': {
      url: process.env.LINEA_TESTNET_URL ?? '',
      accounts,
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
      sepolia: ETHERSCAN_API_KEY,
      polygon: POLYGONSCAN_API_KEY,
      polygonMumbai: POLYGONSCAN_API_KEY,
      bsc: BSCSCAN_API_KEY,
      bscTestnet: BSCSCAN_API_KEY,
    },
    customChains: [
      {
        network: 'kayenuat',
        chainId: 828_282,
        urls: {
          apiURL: 'https://ethnodescan.v4.uat.opendax.app/api/',
          browserURL: 'https://ethnodescan.v4.uat.opendax.app/',
        },
      },
    ],
  },
  dependencyCompiler: {
    paths: [
      '@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol',
      '@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol',
      '@uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol',
      '@syncswap/core-contracts/contracts/interfaces/factory/IBasePoolFactory.sol',
      '@syncswap/core-contracts/contracts/interfaces/IRouter.sol',
    ],
  },
};

export default config;
