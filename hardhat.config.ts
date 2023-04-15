import 'dotenv/config';

import '@nomicfoundation/hardhat-toolbox';
import '@nomiclabs/hardhat-ethers';
import '@openzeppelin/hardhat-upgrades';
import type { HardhatUserConfig } from 'hardhat/config';
import 'solidity-docgen';

import './src/tasks/accounts';
import './src/tasks/activate';
import './src/tasks/forceImport';
import './src/tasks/setupNFTs';

const accounts = {
  mnemonic:
    process.env.SIGNER_MNEMONIC ??
    'clown border solid resource camp medal angle success achieve impulse beach inherit busy track hazard',
};

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: '0.8.18',
        settings: {
          viaIR: true,
          optimizer: {
            enabled: true,
            details: {
              yulDetails: {
                optimizerSteps: 'u',
              },
            },
          },
        },
      },
    ],
  },
  networks: {
    hardhat: {
      accounts: {
        count: 10,
      },
    },
    mumbai: {
      url: 'https://wandering-aged-tree.matic-testnet.quiknode.pro/a1e69e9f8661279922044553d45860b09aa4765e/',
      accounts,
    },
    polygon: {
      url: 'https://frequent-icy-feather.matic.quiknode.pro/d2a51d3b849ba555c8f56e4ded259f70d9ae724e/',
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
      polygon: process.env.POLYGONSCAN_API_KEY,
    },
  },
};

export default config;
