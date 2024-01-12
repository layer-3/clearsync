import { BigNumber, constants, ethers, utils } from 'ethers';
import ExitFormat, { AllocationType } from '@statechannels/exit-format';

import { parseEventResult } from '../ethers-utils';

import { decodeGuaranteeData } from './outcome';

/**
 * Holds rich information about a Deposited event emitted on chain
 */
export interface DepositedEvent {
  destination: string; // The channel that funds were deposited in to.
  destinationHoldings: BigNumber; // The amount the holdings were updated to.
}

/**
 * Extracts a DepositedEvent from a suitable eventResult
 * @param eventResult
 * @returns a DepositedEvent
 */
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function getDepositedEvent(eventResult: ethers.Event[]): DepositedEvent {
  const { destination, destinationHoldings } = parseEventResult(eventResult) as {
    destination: string;
    destinationHoldings: string;
  };
  return {
    destination,
    destinationHoldings: BigNumber.from(destinationHoldings),
  };
}

/**
 * Extracts a 20 byte address hex string from a 32 byte hex string by slicing
 * @param bytes32 a hex string
 * @returns a copy of the hex string with the first 12 bytes removed (after the "0x")
 */
export function convertBytes32ToAddress(bytes32: string): string {
  const normalized = utils.hexZeroPad(bytes32, 32);
  return utils.getAddress(`0x${normalized.slice(-40)}`);
}

/**
 * Left pads a 20 byte address hex string with zeros until it is a 32 byte hex string
 * e.g.,
 * 0x9546E319878D2ca7a21b481F873681DF344E0Df8 becomes
 * 0x0000000000000000000000009546E319878D2ca7a21b481F873681DF344E0Df8
 * @param address 20 byte hex string
 * @returns 32 byte padded hex string
 */
export function convertAddressToBytes32(address: string): string {
  const normalizedAddress = BigNumber.from(address).toHexString();
  if (!utils.isAddress(address)) {
    throw new Error(`Input is not a valid Ethereum address.`);
  }

  // We pad to 66 = (32*2) + 2('0x')
  return utils.hexZeroPad(normalizedAddress, 32);
}

/**
 *
 * Emulates solidity code. TODO replace with PureEVM implementation?
 * @param initialHoldings
 * @param allocation
 * @param indices
 */
export function computeReclaimEffects(
  sourceAllocations: ExitFormat.Allocation[], // we must index this with a JS number that is less than 2**32 - 1
  targetAllocations: ExitFormat.Allocation[], // we must index this with a JS number that is less than 2**32 - 1
  indexOfTargetInSource: number,
): ExitFormat.Allocation[] {
  const newSourceAllocations: ExitFormat.Allocation[] = []; // will be one slot shorter than sourceAllocations
  const guarantee = sourceAllocations[indexOfTargetInSource];

  if ((guarantee.allocationType as AllocationType) != AllocationType.guarantee) {
    throw new Error('not a guarantee');
  }

  const { left, right } = decodeGuaranteeData(guarantee.metadata);

  let foundTarget = false;
  let foundLeft = false;
  let foundRight = false;

  let totalReclaimed = BigNumber.from(0);

  let k = 0;
  for (const [i, sourceAllocation] of sourceAllocations.entries()) {
    if (i == indexOfTargetInSource) {
      foundTarget = true;
      continue;
    }
    newSourceAllocations[k] = {
      destination: sourceAllocation.destination,
      amount: sourceAllocation.amount,
      allocationType: sourceAllocation.allocationType,
      metadata: sourceAllocation.metadata,
    };

    // copy each element except the indexOfTargetInSource element
    if (!foundLeft && sourceAllocation.destination.toLowerCase() == left.toLowerCase()) {
      newSourceAllocations[k].amount = BigNumber.from(sourceAllocation.amount)
        .add(targetAllocations[0].amount)
        .toHexString();
      totalReclaimed = totalReclaimed.add(targetAllocations[0].amount);
      foundLeft = true;
    }
    if (!foundRight && sourceAllocation.destination.toLowerCase() == right.toLowerCase()) {
      newSourceAllocations[k].amount = BigNumber.from(sourceAllocation.amount)
        .add(targetAllocations[1].amount)
        .toHexString();
      totalReclaimed = totalReclaimed.add(targetAllocations[1].amount);
      foundRight = true;
    }
    k++;
  }

  if (!foundTarget) {
    throw new Error('could not find target');
  }

  if (!foundLeft) {
    throw new Error('could not find left');
  }

  if (!foundRight) {
    throw new Error('could not find right');
  }

  if (!totalReclaimed.eq(guarantee.amount)) {
    throw new Error('totalReclaimed!=guarantee.amount');
  }

  return newSourceAllocations;
}

/**
 *
 * Emulates solidity code. TODO: replace with PureEVM implementation?
 * @param initialHoldings
 * @param allocation
 * @param indices
 */
export function computeTransferEffectsAndInteractions(
  initialHoldings: string,
  allocations: ExitFormat.Allocation[], // we must index this with a JS number that is less than 2**32 - 1
  indices: number[],
): {
  newAllocations: ExitFormat.Allocation[];
  allocatesOnlyZeros: boolean;
  exitAllocations: ExitFormat.Allocation[];
  totalPayouts: string;
} {
  let totalPayouts = BigNumber.from(0);
  const newAllocations: ExitFormat.Allocation[] = [];
  const exitAllocations = Array.from({
    length: indices.length > 0 ? indices.length : allocations.length,
  }).fill({
    destination: constants.HashZero,
    amount: '0x00',
    metadata: '0x',
    allocationType: 0,
  }) as ExitFormat.Allocation[];
  let allocatesOnlyZeros = true;
  let surplus = BigNumber.from(initialHoldings);
  let k = 0;

  for (const [i, allocation] of allocations.entries()) {
    newAllocations.push({
      destination: allocation.destination,
      amount: BigNumber.from(0).toHexString(),
      metadata: allocation.metadata,
      allocationType: allocation.allocationType,
    });
    const affordsForDestination = min(BigNumber.from(allocation.amount), surplus);
    if (indices.length === 0 || (k < indices.length && indices[k] === i)) {
      newAllocations[i].amount = BigNumber.from(allocation.amount)
        .sub(affordsForDestination)
        .toHexString();
      exitAllocations[k] = {
        destination: allocation.destination,
        amount: affordsForDestination.toHexString(),
        metadata: allocation.metadata,
        allocationType: allocation.allocationType,
      };
      totalPayouts = totalPayouts.add(affordsForDestination);
      ++k;
    } else {
      newAllocations[i].amount = allocation.amount;
    }
    if (!BigNumber.from(newAllocations[i].amount).isZero()) allocatesOnlyZeros = false;
    surplus = surplus.sub(affordsForDestination);
  }

  return {
    newAllocations,
    allocatesOnlyZeros,
    exitAllocations,
    totalPayouts: totalPayouts.toHexString(),
  };
}

function min(a: BigNumber, b: BigNumber): BigNumber {
  return a.gt(b) ? b : a;
}
