import '@nomicfoundation/hardhat-toolbox';
import { task, types } from 'hardhat/config';
import { formatEther } from 'ethers/lib/utils';
import { constants } from 'ethers';

import type { ERC20, TokenBridge, YellowToken } from '../../typechain-types';

interface DeployBridgeArgs {
  endpointAddress: string;
  tokenAddress: string;
  isRoot?: boolean;
}

task('deployBridge', 'Deploys Root Token Bridge')
  .addParam('endpointAddress', 'The address of LZ Endpoint on the chain')
  .addParam('tokenAddress', 'The address of the Token to bridge')
  .addOptionalParam(
    'isRoot',
    'Whether the bridge is root or child (default - child)',
    false,
    types.boolean,
  )
  .setAction(async (taskArgs: DeployBridgeArgs, { ethers }) => {
    const TokenBridgeFactory = await ethers.getContractFactory('TokenBridge');
    const TokenBridge = await TokenBridgeFactory.deploy(
      taskArgs.endpointAddress,
      taskArgs.tokenAddress,
      taskArgs.isRoot,
    );

    await TokenBridge.deployed();

    console.log(`Token Bridge was deployed at ${TokenBridge.address}`);

    if (!taskArgs.isRoot) {
      console.log('\nGranting minter role to the bridge');

      const Token = (await ethers.getContractAt(
        'YellowToken',
        taskArgs.tokenAddress,
      )) as YellowToken;

      const tx = await Token.grantRole(await Token.MINTER_ROLE(), TokenBridge.address);
      await tx.wait();

      console.log('Minter role was granted to the bridge');
    }
  });

// TODO: change to task connecting both bridges. A mapping between LZ chain ids to network names is needed
interface AddTrustedRemoteAddressArgs {
  bridgeAddress: string;
  remoteChainId: number;
  remoteAddress: string;
}

task('addTrustedRemote', 'Adds a trusted remote address to the bridge')
  .addParam('bridgeAddress', 'The address of the Token Bridge')
  .addParam('remoteChainId', 'The chainId of the remote bridge')
  .addParam('remoteAddress', 'The address of the remote bridge')
  .setAction(async (taskArgs: AddTrustedRemoteAddressArgs, { ethers }) => {
    const TokenBridge = (await ethers.getContractAt(
      'TokenBridge',
      taskArgs.bridgeAddress,
    )) as TokenBridge;

    await TokenBridge.setTrustedRemoteAddress(taskArgs.remoteChainId, taskArgs.remoteAddress);

    console.log(
      `Added trusted remote address ${taskArgs.remoteAddress} from chain ${taskArgs.remoteChainId}`,
    );
  });

interface BridgeTokenArgs {
  receiver: string;
  amount: number;
  bridgeAddress: string;
  remoteChainId: number;
}

task(
  'bridgeToken',
  'Bridges token (registered at the bridge) from active chain to the supplied chain',
)
  .addOptionalParam(
    'receiver',
    'The address to receive tokens on the other chain (default - sender address)',
    '',
    types.string,
  )
  .addParam('amount', 'The amount of tokens to bridge (without decimals)')
  .addParam('bridgeAddress', 'The address of the Token Bridge')
  .addParam('remoteChainId', 'The chainId of the remote bridge')
  .setAction(async (taskArgs: BridgeTokenArgs, { ethers }) => {
    const [Signer] = await ethers.getSigners();

    // transform task args
    if (!taskArgs.receiver) {
      taskArgs.receiver = Signer.address;
    }

    const TokenBridge = (await ethers.getContractAt(
      'TokenBridge',
      taskArgs.bridgeAddress,
    )) as TokenBridge;

    const tokenAddress = await TokenBridge.tokenContract();
    const Token = (await ethers.getContractAt('ERC20', tokenAddress)) as ERC20;
    const amount = taskArgs.amount * 10 ** (await Token.decimals());

    const receiver = taskArgs.receiver;
    const remoteChainId = taskArgs.remoteChainId;

    // check allowance
    console.log('Checking allowance');
    const allowance = await Token.allowance(Signer.address, TokenBridge.address);

    if (allowance.lt(amount)) {
      console.log('Allowance not enough! Approving more tokens for the bridge');
      await Token.approve(TokenBridge.address, amount);
    }

    console.log('Allowance is enough!\n');

    // calculate fee
    const [nativeFeeBN] = await TokenBridge.estimateFees(
      remoteChainId,
      receiver,
      amount,
      false,
      '0x',
    );

    console.log(
      `${formatEther(
        nativeFeeBN,
      )} of native fees is needed to bridge ${amount} to ${receiver} on chain ${remoteChainId}\n`,
    );

    // check native balance
    console.log('Checking native balance');
    const nativeBalance = await Signer.getBalance();
    if (nativeBalance.lt(nativeFeeBN)) {
      throw new Error('Native balance not enough!');
    }
    console.log('Native balance is enough!\n');

    // bridge tokens
    console.log('Bridging tokens');
    const tx = await TokenBridge.bridge(
      remoteChainId,
      receiver,
      amount,
      constants.AddressZero,
      '0x',
      {
        value: nativeFeeBN,
      },
    );
    await tx.wait();

    console.log(`'Bridge token' transaction was mined: ${tx.hash}`);
    console.log('Wait for configured number of blocks for transfer to be finalized');
  });
