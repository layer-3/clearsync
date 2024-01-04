/**
 * @see https://docs.statechannels.org/docs/protocol-tutorial/understand-channel-storage/#channel-modes
 */
export type ChannelMode = 'Open' | 'Challenge' | 'Finalized';

/**
 * Takes in when the channel finalizes and the current timestamp and determines the `ChannelMode`
 * @see https://docs.statechannels.org/docs/protocol-tutorial/understand-channel-storage/#channel-modes
 * @param finalizesAt The timestamp when the channel finalizes. 0 indicates no finalization.
 * @param now The current ethereum block timestamp
 */
export function getChannelMode(finalizesAt: number, now: number): ChannelMode {
  if (finalizesAt == 0) {
    return 'Open';
  } else if (finalizesAt <= now) {
    return 'Finalized';
  } else {
    return 'Challenge';
  }
}
