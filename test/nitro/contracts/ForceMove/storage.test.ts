import { ethers } from 'ethers';
import { before, describe, it } from 'mocha';
import { expect } from 'chai';

import {
  ChannelData,
  channelDataToStatus,
  parseStatus,
} from '../../../../src/nitro/contract/channel-storage';
import { randomChannelId, setupContract } from '../../test-helpers';

import type { TESTForceMove } from '../../../../typechain-types';

let forceMove: TESTForceMove;
before(async () => {
  forceMove = await setupContract<TESTForceMove>('TESTForceMove');
});

const zeroData = {
  stateHash: ethers.constants.HashZero,
  outcomeHash: ethers.constants.HashZero,
};
describe('storage', () => {
  const testCases = [
    {
      turnNumRecord: 0x42,
      finalizesAt: 0x90_01,
    },
    {
      turnNumRecord: 123_456,
      finalizesAt: 789,
    },
  ];

  for (const tc of testCases)
    it('Statusing and data retrieval', async () => {
      const storage = tc;
      const blockchainStorage = { ...storage, ...zeroData };
      const blockchainStatus = await forceMove.generateStatus(blockchainStorage);

      const clientStatus = channelDataToStatus(storage as unknown as ChannelData);

      const expected = { ...storage, fingerprint: '0x' + clientStatus.slice(2 + 24) };

      expect(clientStatus).to.equal(blockchainStatus);
      expect(await forceMove.matchesStatus(blockchainStorage, blockchainStatus)).to.equal(true);
      expect(await forceMove.matchesStatus(blockchainStorage, clientStatus)).to.equal(true);

      let event = parseStatus(clientStatus);
      for (const [key, value] of Object.entries(expected)) {
        expect(event[key as keyof typeof event]).to.equal(value);
      }

      // Testing getData is a little more laborious
      const setStatusTx = await forceMove.setStatusFromChannelData(
        ethers.constants.HashZero,
        blockchainStorage,
      );
      await setStatusTx.wait();

      const {
        turnNumRecord,
        finalizesAt,
        fingerprint: f,
      } = await forceMove.unpackStatus(ethers.constants.HashZero);

      event = { turnNumRecord, finalizesAt, fingerprint: f._hex };
      for (const [key, value] of Object.entries(expected)) {
        expect(event[key as keyof typeof event]).to.equal(value);
      }
    });
});

describe('_requireChannelOpen', () => {
  let channelId: string;
  beforeEach(() => {
    channelId = randomChannelId();
  });

  it('works when the slot is empty', async () => {
    expect(await forceMove.statusOf(channelId)).to.equal(ethers.constants.HashZero);
    await forceMove.requireChannelOpen(channelId);
  });

  const testCases = [
    {
      result: 'reverts',
      turnNumRecord: 42,
      finalizesAt: 1e12,
    },
    {
      result: 'reverts',
      turnNumRecord: 42,
      finalizesAt: 0x90_01,
    },
    {
      result: 'works',
      turnNumRecord: 123,
      finalizesAt: '0x00',
    },
    {
      result: 'works',
      turnNumRecord: 0xa_bc,
      finalizesAt: '0x00',
    },
    {
      result: 'works',
      turnNumRecord: 1,
      finalizesAt: '0x00',
    },
    {
      result: 'works',
      turnNumRecord: 0,
      finalizesAt: '0x00',
    },
  ];

  for (const tc of testCases)
    it(`${tc.result} with turnNumRecord: ${tc.turnNumRecord}, finalizesAt: ${tc.finalizesAt}`, async () => {
      const blockchainStorage = {
        turnNumRecord: tc.turnNumRecord,
        finalizesAt: tc.finalizesAt,
        ...zeroData,
      };

      const setStatusTx = await forceMove.setStatusFromChannelData(channelId, blockchainStorage);
      await setStatusTx.wait();
      expect(await forceMove.statusOf(channelId)).to.equal(
        channelDataToStatus(blockchainStorage as ChannelData),
      );

      const tx = forceMove.requireChannelOpen(channelId);
      tc.result === 'reverts' ? await expect(tx).to.be.revertedWith('Channel not open.') : await tx;
    });
});
