import { Contract, BigNumber } from 'ethers';
import type { ParamType } from 'ethers/lib/utils';
import { ethers } from 'hardhat';

import { expectRevert } from '../../../helpers/expect-revert';
import { computeOutcome, convertAddressToBytes32 } from '../../../../src/nitro';
import {
  FixedPart,
  getFixedPart,
  getVariablePart,
  RecoveredVariablePart,
  State,
} from '../../../../src/nitro/contract/state';
import { generateParticipants, setupContract } from '../../test-helpers';
import { expectUnsupportedState } from '../../tx-expect-wrappers';
import type { InterestBearingApp } from '../../../../typechain-types';

let interestBearingApp: Contract;
let baseState: State;
let fixedPart: FixedPart;
let signedByBorrower: RecoveredVariablePart;
let signedByLender: RecoveredVariablePart;
let signedByBoth: RecoveredVariablePart;

const { participants } = generateParticipants(2);
const challengeDuration = 0x100;
const MAGIC_NATIVE_ASSET_ADDRESS = '0x0000000000000000000000000000000000000000';

const merchant = convertAddressToBytes32(participants[0]);
const intermediary = convertAddressToBytes32(participants[1]);

interface Funds {
  asset: string[]; // asset token address
  amount: number[]; // amount of each asset with shared index
}

interface InterestBearingAppData {
  interestPerBlockDivisor: number;
  blocknumber: number;
  principal: Funds;
  collectedInterest: Funds;
}
const interestBearingAppDataTy: ParamType = {
  type: 'tuple',
  components: [
    { type: 'uint256', name: 'interestPerBlockDivisor' },
    { type: 'uint256', name: 'blocknumber' },
    {
      type: 'tuple',
      name: 'principal',
      components: [
        { type: 'address[]', name: 'asset' },
        { type: 'uint256[]', name: 'amount' },
      ],
    },
    {
      type: 'tuple',
      name: 'collectedInterest',
      components: [
        { type: 'address[]', name: 'asset' },
        { type: 'uint256[]', name: 'amount' },
      ],
    },
  ],
} as ParamType;

function appDataABIEncode(appData: InterestBearingAppData): string {
  return ethers.utils.defaultAbiCoder.encode(
    // ['uint128', 'uint128', 'uint256', 'tuple(address[], uint256[])', 'tuple(address[], uint256[])'],
    [interestBearingAppDataTy],
    [appData],
  );
}

const initialOutcome = computeOutcome({
  [MAGIC_NATIVE_ASSET_ADDRESS]: { [merchant]: 1000, [intermediary]: 0 },
});

const baseAppData: InterestBearingAppData = {
  interestPerBlockDivisor: 1000, // 0.1% block percentage yield (1/1000)
  blocknumber: 1,
  principal: {
    asset: [MAGIC_NATIVE_ASSET_ADDRESS],
    amount: [1000],
  },
  collectedInterest: {
    asset: [MAGIC_NATIVE_ASSET_ADDRESS],
    amount: [0],
  },
};

before(async () => {
  interestBearingApp = await setupContract<InterestBearingApp>('InterestBearingApp');

  baseState = {
    turnNum: 0,
    isFinal: false, // intermediary wants to force finalization
    channelNonce: '0x8',
    participants,
    challengeDuration,
    outcome: initialOutcome,
    appData: appDataABIEncode(baseAppData),
    appDefinition: interestBearingApp.address,
  };

  const variablePart = getVariablePart(baseState);
  fixedPart = getFixedPart(baseState);

  signedByBorrower = {
    variablePart,
    signedBy: BigNumber.from(0b10).toHexString(),
  };
  signedByLender = {
    variablePart,
    signedBy: BigNumber.from(0b01).toHexString(),
  };
  signedByBoth = {
    variablePart,
    signedBy: BigNumber.from(0b11).toHexString(),
  };
});

describe('stateIsSupported', () => {
  it('accepts unanimous states', async () => {
    await interestBearingApp.stateIsSupported(fixedPart, [], signedByBoth);
  });

  it('accepts legitimate interest calculations', async () => {
    // construct a proof+candidate test case with fair interest calculation, assert passing
    // test case:
    // - appdata interest rate is 0.1% per block
    // - initial outcome is 1000:0, with 1000 principal
    // - 1 block passes
    // - challenge outcome is 999:1

    const currentBlockNumber = await ethers.provider.getBlockNumber();
    const supportproofAppData = {
      ...baseAppData,
      blocknumber: currentBlockNumber,
    };
    const supportproofState: State = {
      ...baseState,
      appData: appDataABIEncode(supportproofAppData),
    };
    const pfStateSignedByBoth: RecoveredVariablePart = {
      variablePart: getVariablePart(supportproofState),
      signedBy: BigNumber.from(0b11).toHexString(),
    };

    // mine a block
    ethers.provider.send('evm_mine', []);

    const challengeState: State = {
      ...baseState,
      turnNum: baseState.turnNum + 1,
      outcome: computeOutcome({
        [MAGIC_NATIVE_ASSET_ADDRESS]: { [merchant]: 999, [intermediary]: 1 }, // intermediary picks up 0.1% of the principal
      }),
    };
    const challengeWithLenderSignature: RecoveredVariablePart = {
      variablePart: getVariablePart(challengeState),
      signedBy: BigNumber.from(0b01).toHexString(),
    };

    await interestBearingApp.stateIsSupported(
      fixedPart,
      [pfStateSignedByBoth],
      challengeWithLenderSignature,
    );
  });

  it('rejects excessive interest calculations', async () => {
    // construct proof+candidate test case with unfair interest calculation, assert failure.
    // test case:
    //  - the appData's interest rate is 0.1% per block
    //  - the initial outcome is 1000:0
    //  - the chain is advanced by 1 block
    //  - challenge outcome is 998:2. Fraud!

    const currentBlockNumber = await ethers.provider.getBlockNumber();
    const supportproofAppData = {
      ...baseAppData,
      blocknumber: currentBlockNumber,
    };

    const supportproofState: State = {
      ...baseState,
      appData: appDataABIEncode(supportproofAppData),
    };
    const pfStateSignedByBoth: RecoveredVariablePart = {
      variablePart: getVariablePart(supportproofState),
      signedBy: BigNumber.from(0b11).toHexString(),
    };

    // advance one block
    ethers.provider.send('evm_mine', []);

    const challengeState: State = {
      ...baseState,
      appData: appDataABIEncode(supportproofAppData),
      turnNum: baseState.turnNum + 1,
      outcome: computeOutcome({
        [MAGIC_NATIVE_ASSET_ADDRESS]: { [merchant]: 998, [intermediary]: 2 }, // 998 is unfair: should be 999
      }),
    };
    const updatedWithIntermediarySignature: RecoveredVariablePart = {
      variablePart: getVariablePart(challengeState),
      signedBy: BigNumber.from(0b01).toHexString(),
    };
    await expectRevert(
      () =>
        interestBearingApp.stateIsSupported(
          fixedPart,
          [pfStateSignedByBoth],
          updatedWithIntermediarySignature,
        ),
      'earned<claimed',
    );
  });

  it('rejects unilateral unsupported candidates', async () => {
    // construct challenges with unsupported candidates signed by:
    // - only the intermediary
    // - only the merchant
    // assert failure

    await expectRevert(
      () => interestBearingApp.stateIsSupported(fixedPart, [], signedByBorrower),
      '!unanimous',
    );
    await expectRevert(
      () => interestBearingApp.stateIsSupported(fixedPart, [], signedByLender),
      '!unanimous',
    );
  });

  it('rejects unilateral support proof states', async () => {
    // construct challenges where the proof state is signed by
    //  - only the intermediary
    //  - only the merchant
    // assert failure.

    await expectRevert(
      () => interestBearingApp.stateIsSupported(fixedPart, [signedByBorrower], signedByLender),
      '!unanimous',
    );

    await expectRevert(
      () => interestBearingApp.stateIsSupported(fixedPart, [signedByLender], signedByBorrower),
      '!unanimous',
    );
  });

  it('rejects too-long proofs', async () => {
    // construct a challenge with two proof states, assert failure.

    await expectUnsupportedState(
      () =>
        interestBearingApp.stateIsSupported(
          fixedPart,
          [signedByBorrower, signedByLender], // two proof states - should fail
          signedByLender,
        ),
      '|proof| > 1',
    );
  });
});
