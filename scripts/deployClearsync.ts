import { readFileSync, writeFileSync } from 'node:fs';

import { ethers } from 'hardhat';
import { type Signer, utils } from 'ethers';

import type { TestERC20 } from '../typechain-types';

const INFO_OUTPUT_FILEPATH = (stage: string): string => `./config/${stage}.info.json`;
const CONFIG_FILEPATH = (stage: string): string => `./config/${stage}.config.json`;
const MINT_AMOUNT = 1_000_000;

async function main(): Promise<void> {
  const [Deployer, stage] = await setup();

  console.log('deployer:', await Deployer.getAddress());

  const { chainId } = await ethers.provider.getNetwork();
  console.log('working on chainId:', chainId);

  // create an empty file
  writeFileSync(INFO_OUTPUT_FILEPATH(stage), '{}');

  await deployYNContracts(Deployer, INFO_OUTPUT_FILEPATH(stage));

  if (stage == 'testnet') {
    const config: Config = await readConfig(CONFIG_FILEPATH(stage));
    await deployAndMintTokens(Deployer, config, INFO_OUTPUT_FILEPATH(stage));
  }
}

type stage = 'testnet' | 'canarynet' | 'mainnet';

async function setup(): Promise<[Signer, stage]> {
  const mnemonic = process.env.MNEMONIC;

  let Deployer: Signer;

  if (!mnemonic) {
    // first Signer is deployer by default
    console.log('no MNEMONIC env var, defaulting to first network account');
    [Deployer] = await ethers.getSigners();
  } else if (ethers.utils.isValidMnemonic(mnemonic)) {
    Deployer = ethers.Wallet.fromMnemonic(mnemonic).connect(ethers.provider);
  } else {
    throw new Error('invalid MNEMONIC');
  }

  const stageStr = process.env.STAGE;
  let stage: stage;

  if (!stageStr) {
    console.log('no STAGE env var, defaulting to testnet');
    stage = 'testnet';
  }

  if (stageStr && !['testnet', 'canarynet', 'mainnet'].includes(stageStr)) {
    throw new Error(`invalid STAGE env var: ${stageStr}`);
  }

  stage = stageStr as stage;

  console.log('deploying to stage:', stage);

  return [Deployer, stage];
}

interface TokenConfig {
  name: string;
  symbol: string;
  decimals: number;
}

interface Config {
  allocationAddresses: string[];
  tokens: TokenConfig[];
}

async function readConfig(filepath: string): Promise<Config> {
  const config = JSON.parse(readFileSync(filepath, 'utf8')) as Config;

  for (const [i, token] of config.tokens.entries()) {
    if (!token.name || !token.symbol || !token.decimals) {
      throw new Error(`invalid config for token (${i}): ${JSON.stringify(token)}`);
    }
  }
  console.log(`${config.tokens.length} tokens read from config`);

  if (config.allocationAddresses.length > 0) {
    for (const address of config.allocationAddresses) {
      if (!ethers.utils.isAddress(address)) {
        throw new Error(`invalid allocationAddress: ${address}`);
      }
    }
    console.log(`${config.allocationAddresses.length} allocationAddresses read from config`);
  } else {
    // if no allocation addresses, take 10 network accounts
    console.log(`no allocationAddresses found in the config, defaulting to 10 network accounts`);
    let accounts = await ethers.getSigners();
    if (accounts.length > 10) {
      accounts = accounts.slice(0, 10);
    }

    console.log(`using ${accounts.length} found signer accounts in the network`);

    config.allocationAddresses = await Promise.all(
      accounts.map(async (account) => await account.getAddress()),
    );
  }

  return config;
}

interface Info {
  deployedContracts: YNDeployedContracts;
  tokenList: TokenList;
}

interface YNDeployedContracts {
  adjudicator: string;
  clearingApp: string;
  escrowApp: string;
}

async function deployYNContracts(deployer: Signer, filepath: string): Promise<void> {
  const AdjudicatorFactory = await ethers.getContractFactory('YellowAdjudicator');
  const Adjudicator = await AdjudicatorFactory.connect(deployer).deploy();
  await Adjudicator.deployed();

  const ClearginAppFactory = await ethers.getContractFactory('ConsensusApp'); // TODO: change with ClearingApp when ready
  const ClearginApp = await ClearginAppFactory.connect(deployer).deploy();
  await ClearginApp.deployed();

  const EscrowAppFactory = await ethers.getContractFactory('ConsensusApp'); // TODO: change with EscrowApp when ready
  const EscrowApp = await EscrowAppFactory.connect(deployer).deploy();
  await EscrowApp.deployed();

  const deployedContracts: YNDeployedContracts = {
    adjudicator: Adjudicator.address,
    clearingApp: ClearginApp.address,
    escrowApp: EscrowApp.address,
  };

  // read, modify and write info file
  const info: Info = JSON.parse(readFileSync(filepath, 'utf8')) as Info;
  info.deployedContracts = deployedContracts;
  const json = JSON.stringify(info);
  writeFileSync(filepath, json);
  console.log('contracts deployed, addresses written to', filepath);
}

interface Token {
  chainId: number;
  address: string;
  name: string;
  symbol: string;
  decimals: number;
}

interface TokenList {
  name: string;
  timestamp: string;
  tokens: Token[];
}

async function deployAndMintTokens(
  deployer: Signer,
  config: Config,
  filepath: string,
): Promise<void> {
  if (config.tokens.length === 0) {
    console.log('no tokens to deploy');
    return;
  }

  const { chainId } = await ethers.provider.getNetwork();

  // deploy tokens
  const TokenFactory = await ethers.getContractFactory('TestERC20');
  const tokens: Token[] = [];

  for (const tokenConfig of config.tokens) {
    const Token = (await TokenFactory.connect(deployer).deploy(
      tokenConfig.name,
      tokenConfig.symbol,
      tokenConfig.decimals,
      utils.parseUnits(String(Number.MAX_SAFE_INTEGER), tokenConfig.decimals),
    )) as TestERC20;
    await Token.deployed();

    const token = {
      chainId: chainId,
      address: Token.address,
      name: tokenConfig.name,
      symbol: tokenConfig.symbol,
      decimals: tokenConfig.decimals,
    };

    tokens.push(token);

    // mint tokens to allocationAddresses
    for (const address of config.allocationAddresses) {
      await Token.connect(deployer).mint(
        address,
        utils.parseUnits(String(MINT_AMOUNT), tokenConfig.decimals),
      );
    }
  }

  // read, modify and write info file
  const info: Info = JSON.parse(readFileSync(filepath, 'utf8')) as Info;
  info.tokenList = {
    name: 'Yellow Network test tokens',
    timestamp: new Date().toISOString(),
    tokens: tokens,
  };
  const json = JSON.stringify(info);
  writeFileSync(filepath, json);
  console.log(
    `${config.tokens.length} tokens deployed, balances minted to ${config.allocationAddresses.length} addresses`,
  );
  console.log('tokenList written to', filepath);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
