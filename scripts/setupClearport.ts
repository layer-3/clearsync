import { readFileSync, writeFileSync } from 'node:fs';

import { ethers } from 'hardhat';
import { type Signer, utils } from 'ethers';

import type { TestERC20 } from '../typechain-types';

const DEPLOYED_CONTRACTS_FILEPATH = './scripts/deployedContracts.json';
const CONFIG_FILEPATH = './scripts/config.json';
const MINT_AMOUNT = 1_000_000;
const TOKEN_LIST_FILEPATH = './scripts/tokenList.json';

async function main(): Promise<void> {
  const deployerPrivKey = process.env.DEPLOYER_PRIV_KEY;

  if (!deployerPrivKey || ethers.utils.isHexString(deployerPrivKey, 64)) {
    throw new Error('invalid DEPLOYER_PRIV_KEY');
  }

  const deployer = new ethers.Wallet(deployerPrivKey, ethers.provider);
  const deployerSigner = deployer.connect(ethers.provider);

  console.log('deployer:', deployer.address);

  const { chainId } = await ethers.provider.getNetwork();

  console.log('working on chainId:', chainId);

  await deployYNContracts(deployerSigner, DEPLOYED_CONTRACTS_FILEPATH);

  const config: Config = readConfig(CONFIG_FILEPATH);
  await deployAndMintTokens(deployerSigner, config, TOKEN_LIST_FILEPATH);
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

function readConfig(filepath: string): Config {
  const config = JSON.parse(readFileSync(filepath, 'utf8')) as Config;

  for (const [i, token] of config.tokens.entries()) {
    if (!token.name || !token.symbol || !token.decimals) {
      throw new Error(`invalid config for token (${i}): ${JSON.stringify(token)}`);
    }
  }
  console.log(`${config.tokens.length} tokens read from config`);

  for (const address of config.allocationAddresses) {
    if (!ethers.utils.isAddress(address)) {
      throw new Error(`invalid allocationAddress: ${address}`);
    }
  }
  console.log(`${config.allocationAddresses.length} allocationAddresses read from config`);

  return config;
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

  // write deployedContracts to json file
  const json = JSON.stringify(deployedContracts);
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

  // write token addresses to json file
  const tokenList: TokenList = {
    name: 'Yellow Network test tokens',
    timestamp: new Date().toISOString(),
    tokens: tokens,
  };
  const json = JSON.stringify(tokenList);
  writeFileSync(filepath, json);
  console.log('tokens deployed, balances minted');
  console.log('tokenList written to', filepath);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
