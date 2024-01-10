import { generateParticipants, setupContract } from '../../test-helpers';
import { expect } from 'chai';
import type { TESTForceMove } from '../../../../typechain-types';
import type { Contract } from 'ethers';

let forceMove: Contract & TESTForceMove;
let suspect: string;
let addresses: string[];

before(async () => {
  forceMove = await setupContract<TESTForceMove>('TESTForceMove');

  const nParticipants = 4;
  const { participants } = generateParticipants(nParticipants);
  suspect = participants[0];
  addresses = participants.slice(1);
});

describe('_isAddressInArray', () => {
  it('verifies absence of suspect', async () => {
    expect(await forceMove.isAddressInArray(suspect, addresses)).to.equal(false);
  });
  it('finds an address hiding in an array', async () => {
    addresses[1] = suspect;
    expect(await forceMove.isAddressInArray(suspect, addresses)).to.equal(true);
  });
});
