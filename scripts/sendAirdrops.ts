import { ethers } from 'hardhat';
import { BatchTransfer__factory, ERC20__factory } from '../typechain-types';

import fs from 'fs';
import { parse } from 'csv-parse';

const BATCH_SIZE = 500;

interface AddressesToAmounts {
  addresses: string[];
  amounts: number[];
}

async function main(): Promise<void> {
  const csvPath = process.env.CSV_PATH ?? '';
  const tokenAddress = process.env.TOKEN_ADDRESS ?? '';
  const batcherAddress = process.env.BATCHER_ADDRESS ?? '';
  const withDecimals = process.env.WITH_DECIMALS === 'true' ? true : false;

  const [sender] = await ethers.getSigners();

  console.log('Sending airdrops from address:', sender.address);
  const balanceBN = await sender.getBalance();
  console.log('Native balance:', balanceBN.toString());

  let {addresses, amounts} = await parseCsv(csvPath);
  const quantity = addresses.length;
  console.log(`Processed CSV file with ${quantity} addresses`);

  const batchCount = Math.ceil(quantity / BATCH_SIZE);
  console.log(`Batch count: ${batchCount}`);

  const totalAmount = amounts.reduce((total: number, amount: number) => {
    total += amount;
    return total;
  }, 0);

  const token = ERC20__factory.connect(tokenAddress, sender);

  const decimals = await token.decimals();
  if (withDecimals) amounts = amounts.map((amount) => amount * 10 ** decimals);

  const tokenBalance = await token.balanceOf(sender.address);
  console.log(`Token balance: ${tokenBalance.toString()}`);

  if (tokenBalance.lt(totalAmount)) {
    console.log(`Not enough tokens to send ${totalAmount} tokens`);
    return;
  }

  const allowance = await token.allowance(sender.address, batcherAddress);
  if (allowance.lt(totalAmount)) {
    const tx = await token.approve(batcherAddress, totalAmount);
    await tx.wait();
  }

  for (let i = 0; i < batchCount; i++) {
    const start = i * BATCH_SIZE;
    const end = start + BATCH_SIZE;
    const batchAddresses = addresses.slice(start, end);
    const batchAmounts = amounts.slice(start, end);

    const batchTotalAmount = amounts.reduce((total: number, amount: number) => {
      total += amount;
      return total;
    }, 0);

    const contract = BatchTransfer__factory.connect(batcherAddress, sender);
    const tx = await contract.batchTransferUniqueAmounts(tokenAddress, batchAddresses, batchAmounts, batchTotalAmount)
    console.log(`${i+1}. Transaction hash: ${tx.hash}`);
  }
}

main().catch((error) => {
  console.error(error);

  process.exitCode = 1;
});

async function parseCsv(path: string): Promise<AddressesToAmounts> {
  return new Promise((resolve, reject) => {
    let addresses: string[] = [];
    let amounts: any[] = [];

    const parser = fs.createReadStream(path)
      .pipe(parse({ delimiter: ",", from_line: 2 }));

    parser.on("data", function (row: any[]) {
      if (row.length < 2) return

      const address = row[0] as string;
      if (address === '0x0000000000000000000000000000000000000000') return

      const amount = parseInt(row[1], 16);
      if (isNaN(amount) || amount === 0) return

      addresses.push(address);
      amounts.push(amount);
    }).on("end", function () {
      resolve({addresses, amounts});
    }).on("error", function (error) {
      reject(error);
    });
  });
}