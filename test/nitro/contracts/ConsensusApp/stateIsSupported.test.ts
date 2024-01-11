import { BigNumber, Wallet, ethers } from 'ethers';
import { before, describe, it } from 'mocha';
import { expect } from 'chai';

import { bindSignaturesWithSignedByBitfield, signState } from '../../../../src/nitro';
import {
  FixedPart,
  RecoveredVariablePart,
  State,
  VariablePart,
  getFixedPart,
  getVariablePart,
} from '../../../../src/nitro/contract/state';
import { expectSupportedState } from '../../tx-expect-wrappers';
import { generateParticipants, setupContract } from '../../test-helpers';

import type { ConsensusApp } from '../../../../typechain-types';

const { HashZero } = ethers.constants;

let consensusApp: ConsensusApp;
let state: State;
let fixedPart: FixedPart;
let variablePart: VariablePart;
let sigs: ethers.Signature[];
let candidate: RecoveredVariablePart;

before(async () => {
  consensusApp = await setupContract<ConsensusApp>('ConsensusApp');

  const nParticipants = 3;
  const { wallets, participants } = generateParticipants(nParticipants);

  state = {
    turnNum: 5,
    isFinal: false,
    channelNonce: BigNumber.from(8).toHexString(),
    participants,
    challengeDuration: 0x1_00,
    outcome: [],
    appData: HashZero,
    appDefinition: consensusApp.address,
  };

  // Sign the states
  sigs = wallets.map((w: Wallet) => signState(state, w.privateKey).signature);

  fixedPart = getFixedPart(state);
  variablePart = getVariablePart(state);

  candidate = bindSignaturesWithSignedByBitfield([variablePart], sigs, [0, 0, 0])[0];
});

describe('stateIsSupported', () => {
  it('A single state signed by everyone is considered supported', async () => {
    // expect.assertions(3);
    await expectSupportedState(() => consensusApp.stateIsSupported(fixedPart, [], candidate));
  });

  it('Submitting more than one state does NOT constitute a support proof', async () => {
    await expect(consensusApp.stateIsSupported(fixedPart, [candidate, candidate], candidate)).to.be
      .reverted;
  });

  it('A single state signed by less than everyone is NOT considered supported', async () => {
    const candidate: RecoveredVariablePart = {
      variablePart,
      signedBy: BigNumber.from(0b011).toHexString(),
    };
    await expect(consensusApp.stateIsSupported(fixedPart, [], candidate)).to.be.reverted;
  });
});
