import {constants, utils, BigNumber} from 'ethers';

import {hashOutcome, Outcome} from './outcome';
import {hashState, State} from './state';
import {Bytes, Bytes32, Uint48} from './types';

export interface ChannelData {
  turnNumRecord: Uint48;
  finalizesAt: Uint48;
  state?: State;
  outcome?: Outcome;
}
interface FingerprintPreimage {
  stateHash: Bytes32;
  outcomeHash: Bytes32;
}
const FINGERPRINT_PREIMAGE_TYPE = `tuple(
  bytes32 stateHash,
  bytes32 outcomeHash
)`;

/**
 * Computes the on chain status from the supplied channelData
 * @param channelData
 * @returns the 32 byte "status" word that may be stored on chain for this channel
 */
export function channelDataToStatus(channelData: ChannelData): Bytes32 {
  const {turnNumRecord, finalizesAt} = channelData;
  const hash = utils.keccak256(encodeFingerprintPreimage(channelData));
  const fingerprint = utils.hexDataSlice(hash, 12);

  const status =
    '0x' +
    utils.hexZeroPad(utils.hexlify(turnNumRecord), 6).slice(2) +
    utils.hexZeroPad(utils.hexlify(finalizesAt), 6).slice(2) +
    fingerprint.slice(2);

  return status;
}

export function parseStatus(status: Bytes32): {
  turnNumRecord: number;
  finalizesAt: number;
  fingerprint: Bytes;
} {
  validateHexString(status);

  //
  let cursor = 2;
  const turnNumRecord = '0x' + status.slice(cursor, (cursor += 12));
  const finalizesAt = '0x' + status.slice(cursor, (cursor += 12));
  const fingerprint = '0x' + status.slice(cursor);

  return {
    turnNumRecord: asNumber(turnNumRecord),
    finalizesAt: asNumber(finalizesAt),
    fingerprint,
  };
}
const asNumber: (s: string) => number = s => BigNumber.from(s).toNumber();

function getFingerprintPreimage({finalizesAt, state, outcome}: ChannelData): FingerprintPreimage {
  /*
  When the channel is not open, it is still possible for the state to be missing, indicating that the channel is finalized.
  It is currently up to the caller to ensure this.
  */
  const isOpen = finalizesAt === 0;

  if (isOpen && (outcome || state)) {
    console.warn(`Invalid open channel storage: ${JSON.stringify(outcome || state)}`);
  }

  const stateHash = isOpen || !state ? constants.HashZero : hashState(state);
  const outcomeHash = isOpen || !outcome ? constants.HashZero : hashOutcome(outcome);

  return {stateHash, outcomeHash};
}

export function encodeFingerprintPreimage(data: ChannelData): Bytes {
  return utils.defaultAbiCoder.encode([FINGERPRINT_PREIMAGE_TYPE], [getFingerprintPreimage(data)]);
}

function validateHexString(hexString: string) {
  if (!utils.isHexString(hexString)) {
    throw new Error(`Not a hex string: ${hexString}`);
  }
  if (hexString.length !== 66) {
    throw new Error(`Incorrect length: ${hexString.length}`);
  }
}
