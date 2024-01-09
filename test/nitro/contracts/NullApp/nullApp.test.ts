import { Contract, Wallet } from 'ethers';
import { ethers } from 'hardhat';
import { describe, beforeEach, it } from 'mocha';

import { expectRevert } from '../../../helpers/expect-revert';
import { getVariablePart, State, getFixedPart, getRandomNonce } from '../../../../src/nitro';
import type { FixedPart, SignedVariablePart } from '../../../../src/nitro/contract/state';
import type { NitroAdjudicator } from '../../../../typechain-types';
import { setupContract } from '../../test-helpers';

let nitroAdjudicator: Contract;

describe('null app', () => {
  beforeEach(async () => {
    nitroAdjudicator = await setupContract<NitroAdjudicator>('NitroAdjudicator');
  });

  it('should revert when stateIsSupported is called', async () => {
    const fromState: State = {
      participants: [Wallet.createRandom().address, Wallet.createRandom().address],
      channelNonce: getRandomNonce('nullApp'),
      outcome: [],
      turnNum: 1,
      isFinal: false,
      challengeDuration: 0x0,
      appDefinition: ethers.constants.AddressZero,
      appData: '0x00',
    };
    const toState: State = { ...fromState, turnNum: 2 };

    const fixedPart: FixedPart = getFixedPart(fromState);
    const from: SignedVariablePart = {
      variablePart: getVariablePart(fromState),
      sigs: [],
    };
    const to: SignedVariablePart = {
      variablePart: getVariablePart(toState),
      sigs: [],
    };

    await expectRevert(async () => {
      await nitroAdjudicator.stateIsSupported(fixedPart, [from], to);
    }, 'call revert exception');
  });
});
