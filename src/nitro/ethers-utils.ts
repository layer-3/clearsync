import type { ethers } from 'ethers';

// The last value in a result from an ethers event emission (i.e., Contract.on(<filter>, <result>))
// is an object with keys as the names of the fields emitted by the event.
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function parseEventResult(result: ethers.Event[]): Record<string, unknown> {
  if (result.length === 0) {
    throw new Error('No event result provided');
  }

  return result.at(-1)?.args as Record<string, unknown>;
}
