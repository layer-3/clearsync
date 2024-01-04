import {constants, providers, Signature} from 'ethers';

import {State} from './contract/state';
import * as forceMoveTrans from './contract/transaction-creators/force-move';
import * as multiAssetHolderTrans from './contract/transaction-creators/multi-asset-holder';
import * as nitroAdjudicatorTrans from './contract/transaction-creators/nitro-adjudicator';
import {getStateSignerAddress, SignedState} from './signatures';

// CONSTANTS
export const MAGIC_ADDRESS_INDICATING_ETH = constants.AddressZero;
export const MAX_TX_DATA_SIZE = 128 * 1024; // (bytes) https://github.com/ethereum/go-ethereum/blob/f59ed3565d18c1fa676fd34f4fd41ecccad707e8/core/tx_pool.go#L51
export const NITRO_MAX_GAS = 6_000_000; // should be below the block gas limit, which can change over time and is generally different on mainnet, testnet and ganache.
// sampling some recent blocks on 26/11/2020:
// mainnet  12505858
// ropsten:  8000029
// rinkeby: 10000000
// ganache:  6721975 (hardcoded but can be overriden via config)

/**
 * Marshalls the supplied signedStates into an ethereum transaction for the challenge method. Automatically computes signedVariableParts, etc.
 * @param signedStates an array of signed states
 * @param challengePrivateKey private key of the challenger
 * @returns An ethers TransactionRequest. This can be launched with `await signer.sendTransaction({to: adjudicator.address, ...txRequest}`)
 */
export function createChallengeTransaction(
  signedStates: SignedState[],
  challengePrivateKey: string
): providers.TransactionRequest {
  const {states, signatures, whoSignedWhat} = createSignatureArguments(signedStates);

  return forceMoveTrans.createChallengeTransaction(
    states,
    signatures,
    whoSignedWhat,
    challengePrivateKey
  );
}

/**
 * Marshalls the supplied signedStates into an ethereum transaction for the checkpoint method. Automatically computes signedVariableParts, etc.
 * @param signedStates an array of signed states
 * @returns An ethers TransactionRequest. This can be launched with `await signer.sendTransaction({to: adjudicator.address, ...txRequest}`)
 */
export function createCheckpointTransaction(
  signedStates: SignedState[]
): providers.TransactionRequest {
  const {states, signatures, whoSignedWhat} = createSignatureArguments(signedStates);
  return forceMoveTrans.createCheckpointTransaction({
    states,
    signatures,
    whoSignedWhat,
  });
}

/**
 * Marshalls the supplied signedStates into an ethereum transaction for the conclude method. Automatically computes signedVariableParts, etc.
 * @param conclusionProof an array of signed states
 * @returns An ethers TransactionRequest. This can be launched with `await signer.sendTransaction({to: adjudicator.address, ...txRequest}`)
 */
export function createConcludeTransaction(
  conclusionProof: SignedState[]
): providers.TransactionRequest {
  const {states, signatures, whoSignedWhat} = createSignatureArguments(conclusionProof);
  return forceMoveTrans.createConcludeTransaction(states, signatures, whoSignedWhat);
}

/**
 * Marshalls the supplied destination, expectedHeld and amount into an ethereum transaction for the deposit ETH method.
 * @param destination The channelId to deposit into
 * @param expectedHeld The amount you expect to have already been deposited
 * @param amount The amount you intend to deposit
 * @returns the transaction request
 */
export function createETHDepositTransaction(
  destination: string,
  expectedHeld: string,
  amount: string
): providers.TransactionRequest {
  return multiAssetHolderTrans.createETHDepositTransaction(destination, expectedHeld, amount);
}

/**
 * Marshalls the supplied tokenAddress, destination, expectedHeld and amount into an ethereum transaction for the deposit ERC20 token method.
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
): providers.TransactionRequest {
  return multiAssetHolderTrans.createERC20DepositTransaction(
    tokenAddress,
    destination,
    expectedHeld,
    amount
  );
}

/**
 * Marshalls the supplied signedStates into an ethereum transaction for the conclude and transfer all assets method. Automatically computes signedVariableParts, etc.
 * @param signedStates an array of signed states
 * @returns An ethers TransactionRequest. This can be launched with `await signer.sendTransaction({to: adjudicator.address, ...txRequest}`)
 */
export function createConcludeAndTransferAllAssetsTransaction(
  signedStates: SignedState[]
): providers.TransactionRequest {
  const {states, signatures, whoSignedWhat} = createSignatureArguments(signedStates);
  return nitroAdjudicatorTrans.createConcludeAndTransferAllAssetsTransaction(
    states,
    signatures,
    whoSignedWhat
  );
}

/**
 * Marshalls the supplied signedStates into an ethereum transaction for the transfer all assets method. Automatically computes signedVariableParts, etc.
 * @param state last state in the channel
 * @returns An ethers TransactionRequest. This can be launched with `await signer.sendTransaction({to: adjudicator.address, ...txRequest}`)
 */
export function createTransferAllAssetsTransaction(state: State): providers.TransactionRequest {
  return nitroAdjudicatorTrans.createTransferAllAssetsTransaction(state);
}

/**
 * Marshalls the supplied signedStates into the signature arguments used in most on chain methods.
 * Currently we assume each signedState is a unique combination of state/signature
 * So if multiple participants sign a state we expect a SignedState for each participant
 * @param signedStates an array of signed states
 * @returns Object with (states, signatures, whosignedWhat)
 */
export function createSignatureArguments(signedStates: SignedState[]): {
  states: State[];
  signatures: Signature[];
  whoSignedWhat: number[];
} {
  const {participants} = signedStates[0].state;
  const states = [];
  const whoSignedWhat = new Array<number>(participants.length);

  // Get a list of all unique signed states.
  const uniqueSignedStates = signedStates.filter((s, i, a) => a.indexOf(s) === i);
  // We may receive multiple Signed States which have the same state and different signatures
  // so we get a list of unique states ignoring their signatures
  // which allows us to create a single state with multiple signatures
  const uniqueStates = uniqueSignedStates.map(s => s.state).filter((s, i, a) => a.indexOf(s) === i);
  const signatures = new Array<Signature>(uniqueStates.length);
  for (let i = 0; i < uniqueStates.length; i++) {
    states.push(uniqueStates[i]);
    // Get a list of all signed states that have the state
    const signedStatesForUniqueState = uniqueSignedStates.filter(s => s.state === uniqueStates[i]);
    // Iterate through the signatures and set signatures/whoSignedWhawt
    for (const ss of signedStatesForUniqueState) {
      const participantIndex = participants.indexOf(getStateSignerAddress(ss));

      signatures[participantIndex] = ss.signature;
      whoSignedWhat[participantIndex] = i;
    }
  }

  return {
    states,
    signatures,
    whoSignedWhat,
  };
}
