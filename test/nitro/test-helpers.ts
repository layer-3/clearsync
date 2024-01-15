import { BigNumber, Contract, Wallet, providers } from 'ethers';
import { ethers } from 'hardhat';
import {
  Allocation,
  AllocationType,
  AssetMetadata,
  SingleAssetExit,
} from '@statechannels/exit-format';

import { channelDataToStatus } from '../../src/nitro/contract/channel-storage';

import type { LogDescription } from '@ethersproject/abi';
import type {
  ChallengeClearedEvent,
  ChallengeRegisteredStruct,
} from '../../src/nitro/contract/challenge';
import type { Outcome } from '../../src/nitro/contract/outcome';
import type { Bytes32, State, VariablePart } from '../../src/nitro';

/**
 * Deploys a given contract using the first available signer from ethers.
 *
 * This function is generic and can be used to deploy any contract for which a factory can be created.
 * It retrieves the first signer from the ethers provider to use as the deployer of the contract.
 *
 * @param contractName The name of the contract to deploy. This should match the name used in the contract's declaration.
 * @returns A promise that resolves to a deployed contract instance. The type of the contract is specified by the generic type parameter <T>.
 * @typeparam T The type of the contract to be deployed. This should correspond to the type of the contract instance expected.
 *
 * Usage example:
 * const contract: MyContractType = await deployContract<MyContractType>('MyContractName');
 */
export async function setupContract<T extends Contract>(
  contractName: string,
  ...args: unknown[]
): Promise<T> {
  const [Deployer] = await ethers.getSigners();
  const factory = await ethers.getContractFactory(contractName);
  const contract = (await factory.connect(Deployer).deploy(...args)) as T;
  await contract.deployed();
  return contract;
}

export const nonParticipant = ethers.Wallet.createRandom();

/**
 * Generate `n` wallets and return them alongside with array of their addresses.
 * @param n Number of wallets to create.
 * @returns N wallets with their addresses.
 */
export function generateParticipants(n: number): {
  wallets: Wallet[];
  participants: string[];
} {
  const participants = Array.from({ length: n }).fill('') as string[];
  const wallets: Wallet[] = Array.from({ length: n });

  for (let i = 0; i < n; i++) {
    wallets[i] = Wallet.createRandom();
    participants[i] = wallets[i].address;
  }

  return { participants, wallets };
}

export const clearedChallengeFingerprint = (turnNumRecord = 5): Bytes32 =>
  channelDataToStatus({
    turnNumRecord,
    finalizesAt: 0,
  });

export const ongoingChallengeFingerprint = (turnNumRecord = 5): Bytes32 =>
  channelDataToStatus({
    turnNumRecord,
    finalizesAt: 1e12,
    outcome: [],
  });

export const finalizedFingerprint = (
  turnNumRecord = 5,
  finalizesAt = 1,
  outcome: Outcome = [],
  state?: State,
): Bytes32 =>
  channelDataToStatus({
    turnNumRecord,
    finalizesAt,
    outcome,
    state,
  });

// See https://github.com/ethers-io/ethers.js/discussions/2429
// convertToStruct takes an array-ish type (like those returned from on-chain calls) and converts it to an object type
export const convertToStruct = <A extends unknown[]>(arr: A): ExtractPropsFromArray<A> => {
  const keys = Object.keys(arr).filter((key) => Number.isNaN(Number(key)));
  const result: Record<string, unknown> = {};

  for (const [index, item] of arr.entries()) {
    result[keys[index]] = item;
  }

  return result as A;
};

// This is to remove unnecessary properties from the output type
export type ExtractPropsFromArray<T> = Omit<T, keyof unknown[] | `${number}`>;

export const parseOutcomeEventResult = (eventOutcomeResult: unknown[]): Outcome => {
  const outcome: Outcome = [];

  if (eventOutcomeResult.length === 0) {
    return outcome;
  }

  for (const eventSingleAssetExit of eventOutcomeResult) {
    const singleAssetOutcome = convertToStruct(
      eventSingleAssetExit as unknown[],
    ) as SingleAssetExit;
    singleAssetOutcome.assetMetadata = convertToStruct(
      singleAssetOutcome.assetMetadata as unknown as unknown[],
    ) as AssetMetadata;

    const allocations: Allocation[] = [];

    if (singleAssetOutcome.allocations.length > 0) {
      for (const eventAllocation of singleAssetOutcome.allocations) {
        const allocation = convertToStruct(eventAllocation as unknown as unknown[]) as Allocation;
        allocations.push(allocation);
      }
    }
    singleAssetOutcome.allocations = allocations;
    outcome.push(singleAssetOutcome);
  }
  return outcome;
};

export const parseVariablePartEventResult = (vpEventResult: unknown): VariablePart => {
  const vp = convertToStruct(vpEventResult as unknown[]) as VariablePart;
  vp.outcome = parseOutcomeEventResult(vp.outcome);

  return vp;
};

export const newChallengeRegisteredEvent = (
  contract: Contract,
  channelId: string,
): Promise<ChallengeRegisteredStruct[keyof ChallengeRegisteredStruct]> => {
  const filter = contract.filters.ChallengeRegistered(channelId);
  return new Promise((resolve) => {
    contract.on(
      filter,
      (
        eventChannelIdArg,
        eventTurnNumRecordArg,
        eventFinalizesAtArg,
        eventChallengerArg,
        eventIsFinalArg,
        eventFixedPartArg,
        eventChallengeVariablePartArg,
      ) => {
        contract.removeAllListeners(filter);
        resolve([
          eventChannelIdArg,
          eventTurnNumRecordArg,
          eventFinalizesAtArg,
          eventChallengerArg,
          eventIsFinalArg,
          eventFixedPartArg,
          eventChallengeVariablePartArg,
        ]);
      },
    );
  });
};

export const newChallengeClearedEvent = (
  contract: Contract,
  channelId: string,
): Promise<ChallengeClearedEvent[keyof ChallengeClearedEvent]> => {
  const filter = contract.filters.ChallengeCleared(channelId);
  return new Promise((resolve) => {
    contract.on(filter, (eventChannelId, eventTurnNumRecord) => {
      // Match event for this channel only
      contract.removeAllListeners(filter);
      resolve([eventChannelId, eventTurnNumRecord]);
    });
  });
};

export const newConcludedEvent = (contract: Contract, channelId: string): Promise<[Bytes32]> => {
  const filter = contract.filters.Concluded(channelId);
  return new Promise((resolve) => {
    contract.on(filter, () => {
      // Match event for this channel only
      contract.removeAllListeners(filter);
      resolve([channelId]);
    });
  });
};

export const newDepositedEvent = (
  contract: Contract,
  destination: string,
): Promise<[string, BigNumber]> => {
  const filter = contract.filters.Deposited(destination);
  return new Promise((resolve) => {
    contract.on(filter, (eventDestination: string, amountHeld: BigNumber) => {
      // Match event for this destination only
      contract.removeAllListeners(filter);
      resolve([eventDestination, amountHeld]);
    });
  });
};

// Copied from https://stackoverflow.com/questions/58325771/how-to-generate-random-hex-string-in-javascript
const genRanHex = (size: number): string =>
  Array.from({ length: size }, () => Math.floor(Math.random() * 16).toString(16)).join('');

export const randomChannelId = (): Bytes32 => '0x' + genRanHex(64);
export const randomExternalDestination = (): Bytes32 => '0x' + genRanHex(40).padStart(64, '0');

export async function sendTransaction(
  provider: providers.JsonRpcProvider,
  contractAddress: string,
  transaction: providers.TransactionRequest,
): Promise<providers.TransactionReceipt> {
  const signer = provider.getSigner();
  const response = await signer.sendTransaction({ to: contractAddress, ...transaction });
  return await response.wait();
}

interface Event extends LogDescription {
  contract: string;
}

interface Log {
  topics: string[];
  data: string;
  address: string;
}

export function compileEventsFromLogs(logs: Log[], contractsArray: Contract[]): Event[] {
  const events: Event[] = [];
  for (const log of logs) {
    for (const contract of contractsArray) {
      if (log.address === contract.address) {
        events.push({ ...contract.interface.parseLog(log), contract: log.address });
      }
    }
  }
  return events;
}

export const largeOutcome = (
  numAllocationItems: number,
  asset: string = ethers.Wallet.createRandom().address,
): Outcome => {
  const randomDestination = '0x8595a84df2d81430f6213ece3d8519c77daf98f04fe54e253a2caeef4d2add39';
  return numAllocationItems > 0
    ? [
        {
          allocations: Array.from({ length: numAllocationItems }).fill({
            destination: randomDestination,
            amount: '0x01',
            allocationType: AllocationType.simple,
            metadata: '0x',
          }) as Allocation[],
          asset,
          assetMetadata: { assetType: 0, metadata: '0x' },
        },
      ]
    : [];
};
