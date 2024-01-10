import { Contract, ethers } from 'ethers';
import { describe, before, it } from 'mocha';

import { expectRevert } from '../../../helpers/expect-revert';
import {
  ChannelData,
  channelDataToStatus,
  parseStatus,
} from '../../../../src/nitro/contract/channel-storage';
import { randomChannelId, setupContract } from '../../test-helpers';
import type { TESTForceMove } from '../../../../typechain-types';
import { expect } from 'chai';

let forceMove: Contract & TESTForceMove;
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
      finalizesAt: 0x9001,
    },
    {
      turnNumRecord: 123456,
      finalizesAt: 789,
    },
  ];

  testCases.forEach((tc) =>
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
      Object.entries(expected).forEach(([key, value]) => {
        expect(event[key]).to.equal(value);
      });

      // Testing getData is a little more laborious
      await (
        await forceMove.setStatusFromChannelData(ethers.constants.HashZero, blockchainStorage)
      ).wait();
      const {
        turnNumRecord,
        finalizesAt,
        fingerprint: f,
      } = await forceMove.unpackStatus(ethers.constants.HashZero);

      event = { turnNumRecord, finalizesAt, fingerprint: f._hex };
      Object.entries(expected).forEach(([key, value]) => {
        expect(event[key]).to.equal(value);
      });
    }),
  );
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
      finalizesAt: 0x9001,
    },
    {
      result: 'works',
      turnNumRecord: 123,
      finalizesAt: '0x00',
    },
    {
      result: 'works',
      turnNumRecord: 0xabc,
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

  testCases.forEach((tc) =>
    it(`${tc.result} with turnNumRecord: ${tc.turnNumRecord}, finalizesAt: ${tc.finalizesAt}`, async () => {
      const blockchainStorage = {
        turnNumRecord: tc.turnNumRecord,
        finalizesAt: tc.finalizesAt,
        ...zeroData,
      };

      await (await forceMove.setStatusFromChannelData(channelId, blockchainStorage)).wait();
      expect(await forceMove.statusOf(channelId)).to.equal(
        channelDataToStatus(blockchainStorage as ChannelData),
      );

      const tx = forceMove.requireChannelOpen(channelId);
      // eslint-disable-next-line no-unused-expressions
      tc.result === 'reverts' ? await expectRevert(() => tx, 'Channel not open.') : await tx;
    }),
  );
});
