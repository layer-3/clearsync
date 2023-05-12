import { task, types } from 'hardhat/config';
import { BatchTransfer__factory, ERC20__factory } from '../../typechain-types';

import fs from 'fs';
import readline from 'readline';

const MAX_BATCH_SIZE = 500;
const DEFAULT_INTERVAL = 10; // minutes
const addressRegex = /^0x[0-9a-fA-F]{40}$/;

interface TaskArgs {
  addressesPath: string;
  tokenAddress: string;
  batcherAddress: string;
  amount: number;
  minBatchSize?: number;
  maxBatchSize?: number;
  minInterval?: number;
  maxInterval?: number;
}

task('sendBatchTransfer', 'Send batch transfer')
  .addParam('addressesPath', 'The path to the file with addresses')
  .addParam('tokenAddress', 'The token address')
  .addParam('batcherAddress', 'The batcher address')
  .addParam('amount', 'The amount to send')
  .addOptionalParam('minBatchSize', 'The minimum batch size', undefined, types.int)
  .addOptionalParam('maxBatchSize', 'The maximum batch size', undefined, types.int)
  .addOptionalParam('minInterval', 'The minimum interval between batches (minutes)', undefined, types.int)
  .addOptionalParam('maxInterval', 'The maximum interval between batches (minutes)', undefined, types.int)
  .setAction(async (taskArgs: TaskArgs, hre) => {
    const { addressesPath, tokenAddress, batcherAddress, amount } = taskArgs;
    const minBatchSize = taskArgs.minBatchSize ?? MAX_BATCH_SIZE;
    const maxBatchSize = taskArgs.maxBatchSize ?? MAX_BATCH_SIZE;
    const minInterval = taskArgs.minInterval ?? DEFAULT_INTERVAL;
    const maxInterval = taskArgs.maxInterval ?? DEFAULT_INTERVAL;

    if (minBatchSize > maxBatchSize) {
      throw new Error(`minBatchSize must be less than or equal to maxBatchSize: ${minBatchSize} > ${maxBatchSize}`);
    } else if (minInterval > maxInterval) {
      throw new Error(`minInterval must be less than or equal to maxInterval: ${minInterval} > ${maxInterval}`);
    } else if (maxBatchSize > MAX_BATCH_SIZE) {
      throw new Error(`maxBatchSize must be less than or equal to ${MAX_BATCH_SIZE}`);
    }

    const [sender] = await hre.ethers.getSigners();
  
    console.log('Sending airdrops from address:', sender.address);
    const balanceBN = await sender.getBalance();
    console.log('Native balance:', balanceBN.toString());
  
    const addresses = await parseAddressesFile(addressesPath);
    const quantity = addresses.length;
    console.log(`Processed file with ${quantity} addresses`);
  
    const token = ERC20__factory.connect(tokenAddress, sender);
    const decimals = await token.decimals();
    const amountFormatted = amount * 10 ** decimals;
    console.log(`Sending ${amountFormatted} tokens to each address`);

    let i = 0;
    while (i < addresses.length) {
      let batchSize = Math.floor(Math.random() * (maxBatchSize - minBatchSize + 1) + minBatchSize);
      if (i + batchSize > addresses.length) {
        batchSize = addresses.length - i;
      }
      console.log(`Sending batch of ${batchSize} addresses...`)

      const batchAddresses = addresses.slice(i, i + batchSize);
  
      const contract = BatchTransfer__factory.connect(batcherAddress, sender);
      const tx = await contract.batchTransfer(tokenAddress, batchAddresses, amountFormatted);
      console.log(`${i+1}. Transaction hash: ${tx.hash}`);
  
      const interval = Math.floor(Math.random() * (maxInterval - minInterval + 1) + minInterval);
      console.log(`Waiting for ${interval} minutes...`);
      await new Promise(resolve => setTimeout(resolve, interval * 60 * 1000));
  
      i += batchSize;
    }
  });

async function parseAddressesFile(path: string): Promise<string[]> {
  return new Promise((resolve, _) => {
    let addresses: string[] = [];

    const reader = fs.createReadStream(path);
    const rl = readline.createInterface({
      input: reader,
      crlfDelay: Infinity
    });

    rl.on('line', (address) => {
      if (!addressRegex.test(address)) {
        console.log(`Invalid address: ${address}`);
        return
      }

      if (address === '0x0000000000000000000000000000000000000000') return

      addresses.push(address);
    }).on('close', () => {
      resolve(addresses);
    });
  });
}