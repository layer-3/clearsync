import { Wallet } from 'ethers';
import { ethers } from 'hardhat';
import { beforeEach, describe, it } from 'mocha';
import { expect } from 'chai';

import { State, getFixedPart, getRandomNonce, getVariablePart } from '../../../../src/nitro';
import { setupContract } from '../../test-helpers';

import type { FixedPart, SignedVariablePart } from '../../../../src/nitro/contract/state';
import type { NitroAdjudicator } from '../../../../typechain-types';

let nitroAdjudicator: NitroAdjudicator;

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

    await expect(nitroAdjudicator.stateIsSupported(fixedPart, [from], to)).to.be.reverted;
  });
});
