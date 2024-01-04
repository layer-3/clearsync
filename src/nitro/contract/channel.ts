import {utils} from 'ethers';

import {FixedPart} from './state';
import {Bytes32} from './types';

/**
 * Determines if the supplied 32 byte hex string represents an external destination (meaning funds will be paid _out_ of the adjudicator on chain)
 * @param bytes32 a destination
 * @returns true if the destination has 12 leading bytes as zero, false otherwise
 */
export function isExternalDestination(bytes32: Bytes32): boolean {
  return /^0x(0{24})([a-fA-F0-9]{40})$/.test(bytes32);
}

/**
 * Computes the unique id for the supplied channel
 * @param channel Parameters which determine the id
 * @returns a 32 byte hex string representing the id
 */
export function getChannelId(fixedPart: FixedPart): Bytes32 {
  const {participants, channelNonce, appDefinition, challengeDuration} = fixedPart;
  const channelId = utils.keccak256(
    utils.defaultAbiCoder.encode(
      ['address[]', 'uint256', 'address', 'uint48'],
      [participants, channelNonce, appDefinition, challengeDuration]
    )
  );
  if (isExternalDestination(channelId))
    throw Error('This channel would have an external destination as an id');
  return channelId;
}
