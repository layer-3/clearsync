import { Contract, Wallet } from 'ethers';
import { ethers } from 'hardhat';
import { describe, beforeEach, it } from 'mocha';

import { expectRevert } from '../../../helpers/expect-revert';
import { getVariablePart, State, getFixedPart, getRandomNonce } from '../../../../src/nitro';
import type { FixedPart, SignedVariablePart } from '../../../../src/nitro/contract/state';
import type { NitroAdjudicator } from '../../../../typechain-types';

let nitroAdjudicator: Contract;

describe('null app', () => {
  beforeEach(async () => {
    let [Deployer] = await ethers.getSigners();
    const NitroAdjudicatorFactory = await ethers.getContractFactory('NitroAdjudicator');
    nitroAdjudicator = (await NitroAdjudicatorFactory.connect(
      Deployer,
    ).deploy()) as NitroAdjudicator;
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
