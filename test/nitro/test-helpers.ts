import { BigNumber, Contract, Wallet, providers } from 'ethers';
import { ethers } from 'hardhat';
import { expect } from 'chai';
import { Allocation, AllocationType, AssetMetadata } from '@statechannels/exit-format';

import { channelDataToStatus } from '../../src/nitro/contract/channel-storage';

import type { LogDescription } from '@ethersproject/abi';
import type {
  ChallengeClearedEvent,
  ChallengeRegisteredStruct,
} from '../../src/nitro/contract/challenge';
import type { Outcome } from '../../src/nitro/contract/outcome';
import type { Bytes32, OutcomeShortHand, VariablePart } from '../../src/nitro';

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
  ...args: any[]
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
  const participants = new Array(n).fill('');
  const wallets = new Array(n);

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
  state,
): Bytes32 =>
  channelDataToStatus({
    turnNumRecord,
    finalizesAt,
    outcome,
    state,
  });

export const parseOutcomeEventResult = (eventOutcomeResult: any[]): Outcome => {
  eventOutcomeResult = [...eventOutcomeResult];
  const outcome: Outcome = [];

  if (eventOutcomeResult.length === 0) {
    return outcome;
  }

  eventOutcomeResult.forEach((eventSingleAssetExit: any[]) => {
    const asset: string = eventSingleAssetExit[0];
    const assetMetadata: AssetMetadata = {
      assetType: eventSingleAssetExit[1].assetType,
      metadata: eventSingleAssetExit[1].metadata,
    };
    const eventAllocations: any[] = [...eventSingleAssetExit[2]];
    const allocations: Allocation[] = [];

    if (eventAllocations.length > 0) {
      eventAllocations.forEach((eventAllocation: any[]) => {
        const destination: string = eventAllocation[0];
        const amount: string = BigNumber.from(eventAllocation[1]._hex).toString();
        const allocationType: number = eventAllocation[2];
        const metadata: string = eventAllocation[3];

        allocations.push({ destination, amount, allocationType, metadata });
      });
    }

    outcome.push({ asset, assetMetadata, allocations });
  });

  return outcome;
};

export const parseVariablePartEventResult = (vpEventResult: any[]): VariablePart => {
  vpEventResult = [...vpEventResult];
  return {
    outcome: parseOutcomeEventResult(vpEventResult[0]),
    appData: vpEventResult[1],
    turnNum: vpEventResult[2],
    isFinal: vpEventResult[3],
  };
};

export const newChallengeRegisteredEvent = (
  contract: ethers.Contract,
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
  contract: ethers.Contract,
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

export const newConcludedEvent = (
  contract: ethers.Contract,
  channelId: string,
): Promise<[Bytes32]> => {
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
  contract: ethers.Contract,
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
const genRanHex = (size: number) =>
  [...new Array(size)].map(() => Math.floor(Math.random() * 16).toString(16)).join('');

export const randomChannelId = (): Bytes32 => '0x' + genRanHex(64);
export const randomExternalDestination = (): Bytes32 => '0x' + genRanHex(40).padStart(64, '0');

export async function sendTransaction(
  provider: ethers.providers.JsonRpcProvider,
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

export function compileEventsFromLogs(logs: any[], contractsArray: Contract[]): Event[] {
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

// Sets the holdings defined in the multipleHoldings object. Requires an array of the relevant contracts to be passed in.
export function resetMultipleHoldings(
  multipleHoldings: OutcomeShortHand,
  contractsArray: Contract[],
): void {
  for (const assetHolder of Object.keys(multipleHoldings)) {
    const holdings = multipleHoldings[assetHolder];
    Object.keys(holdings).forEach(async (destination) => {
      const amount = holdings[destination];
      contractsArray.forEach(async (contract) => {
        if (contract.address === assetHolder) {
          await (await contract.setHoldings(destination, amount)).wait();
          const res = await contract.holdings(destination);
          expect(res.eq(amount)).to.equal(true);
        }
      });
    });
  }
}

// Check the holdings defined in the multipleHoldings object. Requires an array of the relevant contracts to be passed in.
export function checkMultipleHoldings(
  multipleHoldings: OutcomeShortHand,
  contractsArray: Contract[],
): void {
  for (const assetHolder of Object.keys(multipleHoldings)) {
    const holdings = multipleHoldings[assetHolder];
    Object.keys(holdings).forEach(async (destination) => {
      const amount = holdings[destination];
      contractsArray.forEach(async (contract) => {
        if (contract.address === assetHolder) {
          const res = await contract.holdings(destination);
          expect(res.eq(amount)).to.equal(true);
        }
      });
    });
  }
}

export const largeOutcome = (
  numAllocationItems: number,
  asset: string = ethers.Wallet.createRandom().address,
): Outcome => {
  const randomDestination = '0x8595a84df2d81430f6213ece3d8519c77daf98f04fe54e253a2caeef4d2add39';
  return numAllocationItems > 0
    ? [
        {
          allocations: new Array(numAllocationItems).fill({
            destination: randomDestination,
            amount: '0x01',
            allocationType: AllocationType.simple,
            metadata: '0x',
          }),
          asset,
          assetMetadata: { assetType: 0, metadata: '0x' },
        },
      ]
    : [];
};
