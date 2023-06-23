import { task } from 'hardhat/config';
import fetch from 'node-fetch';

import type {
  YellowToken,
  YellowTokenChildTunnel,
  YellowTokenRootTunnel,
} from '../../../typechain-types';

type PolygonNetwork = 'matic' | 'mumbai';

const isCheckpointed = async (network: PolygonNetwork, blockNum: number): Promise<boolean> => {
  const response = await fetch(
    `https://proof-generator.polygon.technology/api/v1/${network}/block-included/${blockNum}`,
  );
  const json = (await response.json()) as { message: string };
  return json.message === 'success';
};

const waitCheckpointed = async (network: PolygonNetwork, blockNum: number): Promise<void> => {
  let isCheckpointedFlag = false;
  let timeStr = new Date().toLocaleTimeString();

  while (!isCheckpointedFlag) {
    isCheckpointedFlag = await isCheckpointed(network, blockNum);
    timeStr = new Date().toLocaleTimeString();
    if (!isCheckpointedFlag) {
      console.log(`[${timeStr}] Checkpoint status: not included, will check again in 1 min...`);
      await new Promise((resolve) => setTimeout(resolve, 60_000));
    }
  }

  console.log(`[${timeStr}] Deposit transaction was included into a checkpoint!\n`);
};

const getDepositPayload = async (network: PolygonNetwork, txHash: string): Promise<string> => {
  const response = await fetch(
    `https://proof-generator.polygon.technology/api/v1/${network}/exit-payload/${txHash}?eventSignature=0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036`,
  );

  const json = (await response.json()) as { message: string; result: string };

  if (json.message !== 'Payload generation success') {
    throw new Error(`Error while getting deposit data: ${json.result}`);
  }

  return json.result;
};

interface DepositToChildArgs {
  amount: number;
  childTokenAddress: string;
  childTunnelAddress: string;
  isMainnet?: boolean;
}

task('depositToRoot', 'Deposit Yellow tokens from Polygon to Ethereum')
  .addParam('amount', 'Amount of Yellow tokens to deposit')
  .addParam('childTokenAddress', 'The address of the Yellow token on Ethereum')
  .addParam('childTunnelAddress', 'The address of the Yellow token on Polygon')
  .addOptionalParam('isMainnet', 'true if mainnet, false if testnet (default)')
  .setAction(async (taskArgs: DepositToChildArgs, hre) => {
    const { ethers } = hre;
    const { childTokenAddress, childTunnelAddress } = taskArgs;

    const Account = (await ethers.getSigners())[0];

    const isMainnet = taskArgs.isMainnet ?? false;
    const polygonNetwork = isMainnet ? 'matic' : 'mumbai';
    const ethereumNetwork = isMainnet ? 'ethereum' : 'goerli';

    console.log(`Running on ${polygonNetwork}\n`);

    // -------- Deposit Yellow Tokens to Root Tunnel --------

    hre.changeNetwork(polygonNetwork);

    const YellowTokenChildTunnel = (await ethers.getContractAt(
      'YellowTokenChildTunnel',
      childTunnelAddress,
    )) as YellowTokenChildTunnel;

    const YellowToken = (await ethers.getContractAt(
      'YellowToken',
      childTokenAddress,
    )) as YellowToken;

    const amount = taskArgs.amount * 10 ** (await YellowToken.decimals());

    // -------- Check allowance --------

    const approvalBN = await YellowToken.allowance(Account.address, YellowTokenChildTunnel.address);
    const approval = approvalBN.toNumber();

    console.log('Checking allowance...');
    if (approval < amount) {
      console.log(`Allowance not enough (${approval} / ${amount}), approving more...`);
      await (await YellowToken.approve(YellowTokenChildTunnel.address, amount - approval)).wait();
    }
    console.log('Allowance is enough\n');

    // -------- Deposit --------

    console.log('Depositing...');
    const tx = await YellowTokenChildTunnel.deposit(amount);
    const receipt = await tx.wait();
    const depositBlockNum = receipt.blockNumber;

    console.log(
      `Deposited!\n` +
        `[${polygonNetwork}] txHash: ${receipt.transactionHash}, block number: ${depositBlockNum}\n`,
    );

    // -------- Wait for checkpoint --------

    console.log('Waiting for deposit transaction to be included into a checkpoint...');
    console.log('This could take up from 10 min to 2 hrs!\n');

    await waitCheckpointed(polygonNetwork, depositBlockNum);

    // -------- Receive deposit on Root Tunnel --------

    const depositPayload = await getDepositPayload(polygonNetwork, receipt.transactionHash);

    console.log(`Deposit payload: ${depositPayload}\n`);

    hre.changeNetwork(ethereumNetwork);

    const YellowTokenRootTunnel = (await ethers.getContractAt(
      'YellowTokenRootTunnel',
      await YellowTokenChildTunnel.fxRootTunnel(),
    )) as YellowTokenRootTunnel;

    console.log('Receiving deposit on Root Tunnel...');

    const tx2 = await YellowTokenRootTunnel.receiveMessage(depositPayload);
    const receipt2 = await tx2.wait();
    console.log(
      `Deposit received!\n` +
        `[${ethereumNetwork}] txHash: ${receipt2.transactionHash}, block number: ${receipt2.blockNumber}\n`,
    );
  });
