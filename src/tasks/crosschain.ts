import '@nomicfoundation/hardhat-toolbox';
import { task } from 'hardhat/config';
import fetch from 'node-fetch';
import { BridgeClient } from '@maticnetwork/maticjs';

import type {
  YellowToken,
  YellowTokenChildTunnel,
  YellowTokenRootTunnel,
} from '../../typechain-types';

type EthereumNetwork = 'ethereum' | 'goerli';
type PolygonNetwork = 'matic' | 'mumbai';

const ethereumCheckpointManagerAddress = '0x86e4dc95c7fbdbf52e33d563bbdb00823894c287';
const ethereumFxRootAddress = '0xfe5e5D361b2ad62c541bAb87C45a0B9B018389a2';

const goerliCheckpointManagerAddress = '0x2890bA17EfE978480615e330ecB65333b880928e';
const goerliFxRootAddress = '0x3d1d3E34f7fB6D26245E6640E1c50710eFFf15bA';

const polygonFxChildAddress = '0x8397259c983751DAf40400790063935a11afa28a';

const mumbaiFxChildAddress = '0xCf73231F28B7331BBe3124B907840A94851f9f11';

interface SetupTunnelsArgs {
  rootTokenAddress: string;
  childTokenAddress: string;
  isMainnet?: boolean;
}

task(
  'setupTunnels',
  'Deploy Yellow Token Root (Ethereum) and Child (Polygon) Tunnel contracts, link them together and with Yellow Token on Ethereum and Polygon.',
)
  .addParam('rootTokenAddress', 'The address of the Yellow token on Ethereum')
  .addParam('childTokenAddress', 'The address of the Yellow token on Polygon')
  .addOptionalParam<boolean>('isMainnet', 'true if mainnet, false if testnet (default)')
  .setAction(async (taskArgs: SetupTunnelsArgs, hre) => {
    const { ethers } = hre;
    const { rootTokenAddress, childTokenAddress } = taskArgs;

    const isMainnet = taskArgs.isMainnet ?? false;
    const ethereumNetwork = isMainnet ? 'ethereum' : 'goerli';
    const polygonNetwork = isMainnet ? 'matic' : 'mumbai';

    console.log(`Running on ${ethereumNetwork} and ${polygonNetwork}`);

    // -------- Deploy Yellow Token Root Tunnel --------

    hre.changeNetwork(ethereumNetwork);

    const checkpointManagerAddress = isMainnet
      ? ethereumCheckpointManagerAddress
      : goerliCheckpointManagerAddress;

    const fxRootAddress = isMainnet ? ethereumFxRootAddress : goerliFxRootAddress;

    console.log('Deploying YellowTokenRootTunnel...');

    const YellowTokenRootTunnelFactory = await ethers.getContractFactory('YellowTokenRootTunnel');
    const YellowTokenRootTunnel = await YellowTokenRootTunnelFactory.deploy(
      childTokenAddress,
      rootTokenAddress,
      checkpointManagerAddress,
      fxRootAddress,
    );
    await YellowTokenRootTunnel.deployed();

    console.log(`YellowTokenRootTunnel deployed to: ${YellowTokenRootTunnel.address}`);

    // -------- Grant MINTER_ROLE on RootYellowToken to YellowTokenRootTunnel --------

    const RootYellowToken = (await ethers.getContractAt(
      'YellowToken',
      rootTokenAddress,
    )) as YellowToken;
    await RootYellowToken.grantRole(
      await RootYellowToken.MINTER_ROLE(),
      YellowTokenRootTunnel.address,
    );

    console.log(
      `MINTER_ROLE on RootYellowToken (${rootTokenAddress}) granted to YellowTokenRootTunnel(${YellowTokenRootTunnel.address})\n`,
    );

    // -------- Deploy Yellow Token Child Tunnel --------

    hre.changeNetwork(polygonNetwork);

    const fxChildAddress = isMainnet ? polygonFxChildAddress : mumbaiFxChildAddress;

    console.log('Deploying YellowTokenChildTunnel...');

    const YellowTokenChildTunnelFactory = await ethers.getContractFactory('YellowTokenChildTunnel');
    const YellowTokenChildTunnel = await YellowTokenChildTunnelFactory.deploy(
      childTokenAddress,
      rootTokenAddress,
      fxChildAddress,
    );
    await YellowTokenChildTunnel.deployed();

    console.log(`YellowTokenChildTunnel deployed to: ${YellowTokenChildTunnel.address}\n`);

    // -------- Link Yellow Token Root and Child Tunnels --------

    hre.changeNetwork(ethereumNetwork);

    await YellowTokenRootTunnel.setFxChildTunnel(YellowTokenChildTunnel.address);

    console.log(
      `YellowTokenChildTunnel(${YellowTokenChildTunnel.address}) was set on YellowTokenRootTunnel(${YellowTokenRootTunnel.address})`,
    );

    hre.changeNetwork(polygonNetwork);

    await YellowTokenChildTunnel.setFxRootTunnel(YellowTokenRootTunnel.address);

    console.log(
      `YellowTokenRootTunnel(${YellowTokenRootTunnel.address}) was set on YellowTokenChildTunnel(${YellowTokenChildTunnel.address})`,
    );

    console.log('Done!');
  });

const isCheckpointed = async (network: PolygonNetwork, blockNum: number): Promise<boolean> => {
  const response = await fetch(
    `https://proof-generator.polygon.technology/api/v1/${network}/block-included/${blockNum}`,
  );
  const json = (await response.json()) as { message: string };
  return json.message === 'success';
};

const waitCheckpointed = async (network: PolygonNetwork, blockNum: number): Promise<void> => {
  let isCheckpointedFlag = await isCheckpointed(network, blockNum);
  let timeStr = new Date().toLocaleTimeString();

  if (!isCheckpointedFlag) {
    console.log(`[${timeStr}] Checkpoint status: not included, will check again in 1 min...`);

    // eslint-disable-next-line @typescript-eslint/no-misused-promises
    const interval = setInterval(async (): Promise<void> => {
      isCheckpointedFlag = await isCheckpointed(network, blockNum);

      timeStr = new Date().toLocaleTimeString();
      if (isCheckpointedFlag) {
        clearInterval(interval);
      } else {
        console.log(`[${timeStr}] Checkpoint status: not included, will check again in 1 min...`);
      }
    }, 60_000);

    while (!isCheckpointedFlag) {
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

    // // -------- Check allowance --------

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

    console.log(`Running on ${polygonNetwork}\n`);

    // -------- Withdraw Yellow Tokens to Child Tunnel --------

    hre.changeNetwork(polygonNetwork);

    // const YellowTokenRootTunnel = (await ethers.getContractAt(
    //   'YellowTokenRootTunnel',
    //   rootTunnelAddress,
    // )) as YellowTokenRootTunnel;

    // const YellowToken = (await ethers.getContractAt(
    //   'YellowToken',
    //   rootTokenAddress,
    // )) as YellowToken;

    // const amount = taskArgs.amount * 10 ** (await YellowToken.decimals());

    // // -------- Check allowance --------

    // const approvalBN = await YellowToken.allowance(Account.address, YellowTokenRootTunnel.address);
    // const approval = approvalBN.toNumber();

    // console.log('Checking allowance...');
    // if (approval < amount) {
    //   console.log(`Allowance not enough (${approval} / ${amount}), approving more...`);
    //   await (await YellowToken.approve(YellowTokenRootTunnel.address, amount - approval)).wait();
    // }
    // console.log('Allowance is enough\n');

    // // -------- Withdraw --------

    // const to = taskArgs.to ?? Account.address;

    // console.log('Withdrawing...');
    // const tx = await YellowTokenRootTunnel.withdrawTo(to, amount, '0x');
    // const receipt = await tx.wait();
    // const withdrawBlockNum = receipt.blockNumber;
    // console.log(
    //   `Withdrown! txHash: ${receipt.transactionHash}, block number: ${withdrawBlockNum}\n`,
    // );
    // console.log(
    //   'Now waiting for Polygon StateSync to pick up the transaction. This can take from 22 to 30 mins...\n',
    // );

    // // TODO: implement a check for StateSync

    // -------- StateSync check --------

    const bridgeClient = new BridgeClient();
    console.log(bridgeClient);
    console.log(
      await bridgeClient.isDeposited(
        '0xd04359b72d0289b8dd718e96c9f758b32d646d2c2cfff796383e24d326e035e4',
      ),
    );
  });
