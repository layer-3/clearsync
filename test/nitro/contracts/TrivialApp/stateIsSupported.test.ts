import { Wallet, utils } from 'ethers';
import { ethers } from 'hardhat';

import { setupContract } from '../../test-helpers';
import {
  type FixedPart,
  type RecoveredVariablePart,
  State,
  type VariablePart,
  getFixedPart,
  getVariablePart,
} from '../../../../src/nitro/contract/state';
import { expectSupportedState } from '../../tx-expect-wrappers';
import { getRandomNonce } from '../../../../src/nitro/helpers';

import type { TrivialApp } from '../../../../typechain-types';

let trivialApp: TrivialApp;

function computeSaltedHash(salt: string, num: number): string {
  return utils.solidityKeccak256(['bytes32', 'uint256'], [salt, num]);
}

function getRandomRecoveredVariablePart(): RecoveredVariablePart {
  const randomNum = Math.floor(Math.random() * 100);
  const salt = ethers.constants.MaxUint256.toHexString();
  const hash = computeSaltedHash(salt, randomNum);

  const recoveredVariablePart: RecoveredVariablePart = {
    variablePart: {
      outcome: [],
      appData: hash,
      turnNum: 1,
      isFinal: false,
    },
    signedBy: '0',
  };
  return recoveredVariablePart;
}

function getMockedFixedPart(): FixedPart {
  const fixedPart: FixedPart = {
    participants: [Wallet.createRandom().address, Wallet.createRandom().address],
    channelNonce: '0x0',
    appDefinition: trivialApp.address,
    challengeDuration: 0,
  };
  return fixedPart;
}

function mockSigs(vp: VariablePart): RecoveredVariablePart {
  return {
    variablePart: vp,
    signedBy: '0',
  };
}

before(async () => {
  trivialApp = await setupContract('TrivialApp');
});

describe('stateIsSupported', () => {
  it('Transitions between random VariableParts are valid', async () => {
    // expect.assertions(15);
    for (let i = 0; i < 5; i++) {
      const from: RecoveredVariablePart = getRandomRecoveredVariablePart();
      const to: RecoveredVariablePart = getRandomRecoveredVariablePart();

      await expectSupportedState(() =>
        trivialApp.stateIsSupported(getMockedFixedPart(), [from], to),
      );
    }
  });

  it('Transitions between States with mocked-up data are valid', async () => {
    const fromState: State = {
      participants: [Wallet.createRandom().address, Wallet.createRandom().address],
      channelNonce: getRandomNonce('trivialApp'),
      outcome: [],
      turnNum: 1,
      isFinal: false,
      challengeDuration: 0x0,
      appDefinition: trivialApp.address,
      appData: '0x00',
    };
    const toState: State = { ...fromState, turnNum: 2 };

    const from: RecoveredVariablePart = mockSigs(getVariablePart(fromState));
    const to: RecoveredVariablePart = mockSigs(getVariablePart(toState));

    await expectSupportedState(() =>
      trivialApp.stateIsSupported(getFixedPart(fromState), [from], to),
    );
  });
});
