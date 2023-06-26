import '@nomicfoundation/hardhat-toolbox';
import { task } from 'hardhat/config';
import { Contract, ethers } from 'ethers';

import type { YellowToken, YellowTokenRootTunnel } from '../../../typechain-types';

// on child chain
const STATE_RECEIVER_ADDRESS = '0x0000000000000000000000000000000000001001';
const STATE_RECEIVER_ABI = [
  {
    inputs: [],
    name: 'lastStateId',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
];

interface WithdrawToChildArgs {
  amount: number;
  rootTokenAddress: string;
  rootTunnelAddress: string;
  to?: string;
  isMainnet?: boolean;
}

task('withdrawToChild', 'Withdraw Yellow tokens from Ethereum to Polygon')
  .addParam('amount', 'Amount of Yellow tokens to withdraw')
  .addParam('rootTokenAddress', 'The address of the Yellow token on Ethereum')
  .addParam('rootTunnelAddress', 'The address of the Yellow token on Polygon')
  .addOptionalParam('to', 'The address to withdraw to (default: your address)')
  .addOptionalParam('isMainnet', 'true if mainnet, false if testnet (default)')
  .setAction(async (taskArgs: WithdrawToChildArgs, hre) => {
    const { ethers } = hre;
    const { rootTokenAddress, rootTunnelAddress } = taskArgs;

    const Account = (await ethers.getSigners())[0];

    const isMainnet = taskArgs.isMainnet ?? false;
    const polygonNetwork = isMainnet ? 'matic' : 'mumbai';
    const ethereumNetwork = isMainnet ? 'ethereum' : 'goerli';

    console.log(`Running on ${ethereumNetwork}\n`);

    // -------- Withdraw Yellow Tokens to Child Tunnel --------

    hre.changeNetwork(ethereumNetwork);

    const YellowTokenRootTunnel = (await ethers.getContractAt(
      'YellowTokenRootTunnel',
      rootTunnelAddress,
    )) as YellowTokenRootTunnel;

    const YellowToken = (await ethers.getContractAt(
      'YellowToken',
      rootTokenAddress,
    )) as YellowToken;

    const amount = taskArgs.amount * 10 ** (await YellowToken.decimals());

    // -------- Check allowance --------

    const approvalBN = await YellowToken.allowance(Account.address, YellowTokenRootTunnel.address);
    const approval = approvalBN.toNumber();

    console.log('Checking allowance...');
    if (approval < amount) {
      console.log(`Allowance not enough (${approval} / ${amount}), approving more...`);
      await (await YellowToken.approve(YellowTokenRootTunnel.address, amount - approval)).wait();
    }
    console.log('Allowance is enough\n');

    // -------- Withdraw --------

    const to = taskArgs.to ?? Account.address;

    console.log('Withdrawing...');
    const tx = await YellowTokenRootTunnel.withdrawTo(to, amount, '0x');
    const receipt = await tx.wait();
    const withdrawBlockNum = receipt.blockNumber;
    console.log(
      `Withdrawn! txHash: ${receipt.transactionHash}, block number: ${withdrawBlockNum}\n`,
    );
    console.log(
      'Now waiting for Polygon StateSync to pick up the transaction. This can take from 22 to 30 mins...\n',
    );

    // -------- StateSync check --------

    // StateSynced event signature
    const eventSignature = `0x103fed9db65eac19c4d870f49ab7520fe03b99f1838e5996caf47e9e43308392`;
    const targetLog = receipt.logs.find((q) => q.topics[0] === eventSignature);
    if (!targetLog) {
      throw new Error('StateSynced event not found');
    }

    const rootTxStateId = ethers.utils.defaultAbiCoder.decode(
      ['uint256'],
      targetLog.topics[1],
    )[0] as number;

    hre.changeNetwork(polygonNetwork);

    const StateReceiver = await ethers.getContractAt(STATE_RECEIVER_ABI, STATE_RECEIVER_ADDRESS);
    await waitStateSynced(StateReceiver, rootTxStateId);

    console.log('Done!');
  });

async function waitStateSynced(StateReceiver: Contract, rootTxStateId: number): Promise<void> {
  let synced = false;
  let timeStr = new Date().toLocaleTimeString();

  while (!synced) {
    synced = await checkStateSynced(StateReceiver, rootTxStateId);
    timeStr = new Date().toLocaleTimeString();
    if (!synced) {
      console.log(`[${timeStr}] State not synced, will check again in 1 minute...`);
      await new Promise((resolve) => setTimeout(resolve, 60_000));
    }
  }

  console.log(`[${timeStr}] State synced!`);
}

async function checkStateSynced(StateReceiver: Contract, rootTxStateId: number): Promise<boolean> {
  const childLastStateId = (await StateReceiver.lastStateId()) as number;
  return ethers.BigNumber.from(childLastStateId).gte(rootTxStateId);
}
