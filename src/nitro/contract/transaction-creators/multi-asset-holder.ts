import {ethers, constants} from 'ethers';

import NitroAdjudicatorArtifact from '../../../artifacts/contracts/NitroAdjudicator.sol/NitroAdjudicator.json';

export const NitroAdjudicatorContractInterface = new ethers.utils.Interface(
  NitroAdjudicatorArtifact.abi
);

/**
 * Crafts an ethers TransactionRequest targeting the adjudicator's deposit method, specifically for ETH deposits
 * @param destination The channelId to deposit into
 * @param expectedHeld The amount you expect to have already been deposited
 * @param amount The amount you intend to deposit
 * @returns the transaction request
 */
export function createETHDepositTransaction(
  destination: string,
  expectedHeld: string,
  amount: string
): ethers.providers.TransactionRequest {
  const data = NitroAdjudicatorContractInterface.encodeFunctionData('deposit', [
    constants.AddressZero, // Magic constant indicating ETH
    destination,
    expectedHeld,
    amount,
  ]);
  return {data};
}

/**
 * Crafts an ethers TransactionRequest targeting the adjudicator's deposit method, specifically for ERC20 deposits
 * @param tokenAddress The ERC20 token contract address
 * @param destination The channelId to deposit into
 * @param expectedHeld The amount you expect to have already been deposited
 * @param amount The amount you intend to deposit
 * @returns the transaction request
 */
export function createERC20DepositTransaction(
  tokenAddress: string,
  destination: string,
  expectedHeld: string,
  amount: string
): ethers.providers.TransactionRequest {
  const data = NitroAdjudicatorContractInterface.encodeFunctionData('deposit', [
    tokenAddress,
    destination,
    expectedHeld,
    amount,
  ]);
  return {data};
}
