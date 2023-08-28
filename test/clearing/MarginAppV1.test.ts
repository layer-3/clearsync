import { ethers } from 'hardhat';
import { expect } from 'chai';
import {
  FixedPart,
  RecoveredVariablePart,
  State,
  VariablePart,
  getFixedPart,
  getVariablePart,
} from '@statechannels/nitro-protocol/dist/src/contract/state';
import { Outcome, convertAddressToBytes32, getChannelId } from '@statechannels/nitro-protocol';

import {
  SIGNED_BY_NO_ONE,
  signChannelIdAndMarginCall,
  signChannelIdAndSwapCall,
  signedBy,
} from '../helpers/MarginApp/signatures';
import { singleAssetOutcome } from '../helpers/MarginApp/outcome';
import { encodeSignedMarginCall, encodeSignedSwapCall } from '../helpers/MarginApp/encode';
import { marginCallAppData, swapCallAppData } from '../helpers/MarginApp/marginApp';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { MarginAppV1 } from '../../typechain-types';
import type {
  MarginCall,
  SignedMarginCall,
  SignedSwapCall,
  SwapCall,
} from '../helpers/MarginApp/types';
import type { Signature } from 'ethers';

// TODO: suggest adding imports from explicit filepath as exports in nitro-protocol lib

const NOT_UNANIMOUS_PROOF_0 = '!unanimous; |proof|=0';
const NOT_FINAL_TURN_NUM_BIGG_3_PROOF_0 = '!final; turnNum>=3 && |proof|=0';
const POSTFUND_TURN_NUM_NOT_EQUAL_1 = 'postfund.turnNum != 1';
const POSTFUND_NOT_UNANIMOUS = 'postfund !unanimous';
const MARGIN_CALL_TURN_NUM_LESS_2 = 'marginCall.turnNum < 2';
const NO_IDENTITY_PROOF_ON_MARGIN_CALL = 'no identity proof on margin call';
const MARGIN_VERSION_NOT_EQUAL_TURN_NUM = 'marginCall.version != turnNum';
const INCORRECT_LEADER_MARGIN = 'incorrect leader margin';
const INCORRECT_FOLLOWER_MARGIN = 'incorrect follower margin';
const TOTAL_ALLOCATED_CANNOT_CHANGE = 'total allocated cannot change';
const INCORRECT_NUMBER_OF_ASSETS = 'incorrect number of assets';
const INCORRECT_NUMBER_OF_ALLOCATIONS = 'incorrect number of allocations';
const DESTINATIONS_CANNOT_CHANGE = 'destinations cannot change';
const INVALID_LEADER_SIGNATURE = 'invalid leader signature';
const INVALID_FOLLOWER_SIGNATURE = 'invalid follower signature';
const NO_IDENTITY_PROOF_ON_SWAP_CALL = 'no identity proof on swap call';
const SWAP_CALL_TURN_NUM_LESS_3 = 'swapCall.turnNum < 3';
const FIRST_BROKER_NOT_LEADER = '1st broker not leader';
const SECOND_BROKER_NOT_FOLLOWER = '2nd broker not follower';
const SWAP_NOT_DIRECT_SUCCESSOR_OF_MARGIN = 'swapCall not direct successor of marginCall';
const SWAP_VERSION_NOT_EQUAL_TURN_NUM = 'swapCall.version != turnNum';

const brokerAMargin = 100;
const brokerBMargin = 100;

describe('MarginAppV1', () => {
  let BrokerA: SignerWithAddress;
  let BrokerB: SignerWithAddress;
  let Intermediary: SignerWithAddress;
  let BrokerC: SignerWithAddress;
  let BrokerD: SignerWithAddress;

  let brokerADestination: string;
  let brokerBDestination: string;
  let intermediaryDestination: string;

  let testToken1Address: string;
  let testToken2Address: string;

  let MarginAppV1: MarginAppV1;

  // only one MarginApp instance as it is stateless
  before(async () => {
    const MarginAppV1Factory = await ethers.getContractFactory('MarginAppV1');
    MarginAppV1 = (await MarginAppV1Factory.deploy()) as MarginAppV1;
    await MarginAppV1.deployed();

    [BrokerA, BrokerB, Intermediary, BrokerC, BrokerD] = await ethers.getSigners();

    brokerADestination = convertAddressToBytes32(BrokerA.address);
    brokerBDestination = convertAddressToBytes32(BrokerB.address);
    intermediaryDestination = convertAddressToBytes32(Intermediary.address);

    testToken1Address = ethers.Wallet.createRandom().address;
    testToken2Address = ethers.Wallet.createRandom().address;
  });

  let outcome: Outcome;
  let baseStateAIB: State;
  let fixedPartAIB: FixedPart;
  let channelIdAIB: string;

  beforeEach(() => {
    outcome = singleAssetOutcome(testToken1Address, [
      [brokerADestination, brokerAMargin],
      [brokerBDestination, brokerBMargin],
    ]);

    baseStateAIB = {
      turnNum: 0,
      isFinal: false,
      chainId: '31113',
      channelNonce: '0x0',
      participants: [BrokerA.address, Intermediary.address, BrokerB.address],
      challengeDuration: 100,
      outcome: outcome,
      appData: '0x',
      appDefinition: MarginAppV1.address,
    };

    fixedPartAIB = getFixedPart(baseStateAIB);

    channelIdAIB = getChannelId(fixedPartAIB);
  });

  describe('prefund', () => {
    let preFundCandidate: RecoveredVariablePart;

    beforeEach(() => {
      preFundCandidate = {
        signedBy: signedBy([0, 1, 2]),
        variablePart: getVariablePart(baseStateAIB),
      };
    });

    it('succeed when unanimously signed and turnNum = 0', async () => {
      await MarginAppV1.requireStateSupported(fixedPartAIB, [], preFundCandidate);
    });

    it('revert when not unanimously signed', async () => {
      preFundCandidate.signedBy = signedBy([1, 2]);
      await expect(
        MarginAppV1.requireStateSupported(fixedPartAIB, [], preFundCandidate),
      ).to.be.revertedWith(NOT_UNANIMOUS_PROOF_0);
    });

    it('revert when turnNum != 0', async () => {
      preFundCandidate.variablePart.turnNum = 42;
      await expect(
        MarginAppV1.requireStateSupported(fixedPartAIB, [], preFundCandidate),
      ).to.be.revertedWith(NOT_FINAL_TURN_NUM_BIGG_3_PROOF_0);
    });
  });

  describe('postfund', () => {
    let postFundCandidate: RecoveredVariablePart;

    beforeEach(() => {
      postFundCandidate = {
        signedBy: signedBy([0, 1, 2]),
        variablePart: { ...getVariablePart(baseStateAIB), turnNum: 1 },
      };
    });
    it('succeed when unanimously signed and turnNum = 1', async () => {
      await MarginAppV1.requireStateSupported(fixedPartAIB, [], postFundCandidate);
    });

    it('revert when not unanimously signed', async () => {
      postFundCandidate.signedBy = signedBy([0, 2]);

      await expect(
        MarginAppV1.requireStateSupported(fixedPartAIB, [], postFundCandidate),
      ).to.be.revertedWith(NOT_UNANIMOUS_PROOF_0);
    });

    it('revert when turnNum != 1', async () => {
      postFundCandidate.variablePart.turnNum = 42;
      await expect(
        MarginAppV1.requireStateSupported(fixedPartAIB, [], postFundCandidate),
      ).to.be.revertedWith(NOT_FINAL_TURN_NUM_BIGG_3_PROOF_0);
    });
  });

  describe('margin call', () => {
    const delta = 5;
    const brokerAChangedMargin = brokerAMargin + delta;
    const brokerBChangedMargin = brokerBMargin - delta;

    let postFundVariablePart: VariablePart;
    let marginCallVariablePart: VariablePart;

    let recoveredPostFundState: RecoveredVariablePart;

    let marginCall: MarginCall;
    let signedMarginCall: SignedMarginCall;
    let marginCallCandidate: RecoveredVariablePart;

    beforeEach(async () => {
      postFundVariablePart = getVariablePart(baseStateAIB);
      postFundVariablePart.turnNum = 1;

      recoveredPostFundState = {
        signedBy: signedBy([0, 1, 2]),
        variablePart: postFundVariablePart,
      };

      marginCall = {
        version: 2,
        margin: [brokerAChangedMargin, brokerBChangedMargin],
      };

      signedMarginCall = {
        marginCall,
        sigs: (await signChannelIdAndMarginCall([BrokerA, BrokerB], channelIdAIB, marginCall)) as [
          Signature,
          Signature,
        ],
      };

      marginCallVariablePart = getVariablePart(baseStateAIB);
      marginCallVariablePart.turnNum = 2;
      marginCallVariablePart.appData = encodeSignedMarginCall(signedMarginCall);
      marginCallVariablePart.outcome = singleAssetOutcome(testToken1Address, [
        [brokerADestination, brokerAChangedMargin],
        [brokerBDestination, brokerBChangedMargin],
      ]);

      marginCallCandidate = {
        signedBy: signedBy(2),
        variablePart: marginCallVariablePart,
      };
    });

    describe('succeed', () => {
      it('when supplied first margin', async () => {
        await MarginAppV1.requireStateSupported(
          fixedPartAIB,
          [recoveredPostFundState],
          marginCallCandidate,
        );
      });

      it('when supplied margin with higher version', async () => {
        marginCall.version = 42;
        marginCallCandidate.variablePart.turnNum = 42;
        marginCallCandidate.variablePart.appData = await marginCallAppData(
          channelIdAIB,
          marginCall,
          [BrokerA, BrokerB],
        );

        await MarginAppV1.requireStateSupported(
          fixedPartAIB,
          [recoveredPostFundState],
          marginCallCandidate,
        );
      });
    });

    describe('revert', () => {
      describe('postfund', () => {
        it('when postfund not supplied', async () => {
          await expect(
            MarginAppV1.requireStateSupported(fixedPartAIB, [], marginCallCandidate),
          ).to.be.revertedWith(NOT_UNANIMOUS_PROOF_0);
        });

        it('when postfund.turnNum != 1', async () => {
          recoveredPostFundState.variablePart.turnNum = 0;

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(POSTFUND_TURN_NUM_NOT_EQUAL_1);
        });

        it('when postfund !unanimous', async () => {
          recoveredPostFundState.signedBy = signedBy(0);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(POSTFUND_NOT_UNANIMOUS);
        });
      });

      describe('margin call', () => {
        it('when not signed', async () => {
          marginCallCandidate.signedBy = SIGNED_BY_NO_ONE;

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(NO_IDENTITY_PROOF_ON_MARGIN_CALL);
        });

        it('when turnNum != 2+', async () => {
          marginCallCandidate.variablePart.turnNum = 1;

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(MARGIN_CALL_TURN_NUM_LESS_2);
        });

        it('when version != turnNum = 2', async () => {
          marginCall.version = 42;
          marginCallCandidate.variablePart.appData = await marginCallAppData(
            channelIdAIB,
            marginCall,
            [BrokerA, BrokerB],
          );

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(MARGIN_VERSION_NOT_EQUAL_TURN_NUM);
        });

        it('when turnNum != version = 2', async () => {
          marginCallCandidate.variablePart.turnNum = 42;

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(MARGIN_VERSION_NOT_EQUAL_TURN_NUM);
        });

        it('when outcome != leader margin specified', async () => {
          marginCallCandidate.variablePart.outcome = singleAssetOutcome(testToken1Address, [
            [brokerADestination, brokerAMargin],
            [brokerBDestination, brokerBChangedMargin],
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(INCORRECT_LEADER_MARGIN);
        });

        it('when outcome != follower margin specified', async () => {
          marginCallCandidate.variablePart.outcome = singleAssetOutcome(testToken1Address, [
            [brokerADestination, brokerAChangedMargin],
            [brokerBDestination, brokerBMargin],
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(INCORRECT_FOLLOWER_MARGIN);
        });

        it('when margin signed by not leader', async () => {
          marginCallCandidate.variablePart.appData = await marginCallAppData(
            channelIdAIB,
            marginCall,
            [BrokerA, BrokerC],
          );

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(INVALID_FOLLOWER_SIGNATURE);
        });

        it('when margin signed by not follower', async () => {
          marginCallCandidate.variablePart.appData = await marginCallAppData(
            channelIdAIB,
            marginCall,
            [BrokerA, BrokerC],
          );

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(INVALID_FOLLOWER_SIGNATURE);
        });

        it('when garbage encoded', async () => {
          marginCallCandidate.variablePart.appData = '0xdeadbeef';
          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.reverted;
        });
      });

      describe('outcome', () => {
        it('when outcome sum has changed', async () => {
          marginCall.margin = [brokerAChangedMargin + 1, brokerBChangedMargin + 1];

          marginCallCandidate.variablePart.appData = await marginCallAppData(
            channelIdAIB,
            marginCall,
            [BrokerA, BrokerB],
          );
          marginCallCandidate.variablePart.outcome = singleAssetOutcome(testToken1Address, [
            [brokerADestination, brokerAChangedMargin + 1],
            [brokerBDestination, brokerBChangedMargin + 1],
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(TOTAL_ALLOCATED_CANNOT_CHANGE);
        });

        it('when outcome with several assets', async () => {
          marginCallCandidate.variablePart.outcome = [
            ...singleAssetOutcome(testToken1Address, [
              [brokerADestination, brokerAChangedMargin],
              [brokerBDestination, brokerBChangedMargin],
            ]),
            ...singleAssetOutcome(testToken2Address, [
              [brokerADestination, brokerAChangedMargin],
              [brokerBDestination, brokerBChangedMargin],
            ]),
          ];

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(INCORRECT_NUMBER_OF_ASSETS);
        });

        it('when more than 2 allocations', async () => {
          marginCallCandidate.variablePart.outcome = singleAssetOutcome(testToken1Address, [
            [brokerADestination, brokerAChangedMargin],
            [brokerBDestination, brokerBChangedMargin],
            [intermediaryDestination, brokerAChangedMargin],
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(INCORRECT_NUMBER_OF_ALLOCATIONS);
        });

        it('when destination changed', async () => {
          marginCallCandidate.variablePart.outcome = singleAssetOutcome(testToken1Address, [
            [brokerADestination, brokerAChangedMargin],
            [intermediaryDestination, brokerBChangedMargin],
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              marginCallCandidate,
            ),
          ).to.be.revertedWith(DESTINATIONS_CANNOT_CHANGE);
        });
      });
    });
  });

  describe('swap call', () => {
    const delta = 5;
    const brokerAChangedMargin = brokerAMargin + delta;
    const brokerBChangedMargin = brokerBMargin - delta;

    let postFundVariablePart: VariablePart;
    let marginCallVariablePart: VariablePart;
    let swapCallVariablePart: VariablePart;

    let recoveredPostFundState: RecoveredVariablePart;

    let marginCall: MarginCall;
    let signedMarginCall: SignedMarginCall;
    let recoveredMarginCallState: RecoveredVariablePart;

    let adjustedMargin: MarginCall;
    let swapCall: SwapCall;
    let signedSwapCall: SignedSwapCall;
    let swapCallCandidate: RecoveredVariablePart;

    beforeEach(async () => {
      postFundVariablePart = getVariablePart(baseStateAIB);
      postFundVariablePart.turnNum = 1;

      recoveredPostFundState = {
        signedBy: signedBy([0, 1, 2]),
        variablePart: postFundVariablePart,
      };

      marginCall = {
        version: 2,
        margin: [brokerAChangedMargin, brokerBChangedMargin],
      };

      signedMarginCall = {
        marginCall,
        sigs: (await signChannelIdAndMarginCall([BrokerA, BrokerB], channelIdAIB, marginCall)) as [
          Signature,
          Signature,
        ],
      };

      marginCallVariablePart = getVariablePart(baseStateAIB);
      marginCallVariablePart.turnNum = 2;
      marginCallVariablePart.appData = encodeSignedMarginCall(signedMarginCall);
      marginCallVariablePart.outcome = singleAssetOutcome(testToken1Address, [
        [brokerADestination, brokerAChangedMargin],
        [brokerBDestination, brokerBChangedMargin],
      ]);

      recoveredMarginCallState = {
        signedBy: signedBy(2),
        variablePart: marginCallVariablePart,
      };

      adjustedMargin = {
        version: 3,
        margin: [brokerAMargin, brokerBMargin],
      };

      swapCall = {
        brokers: [BrokerA.address, BrokerB.address],
        // TODO: do we need any checks on swaps?
        swaps: [
          [{ token: ethers.Wallet.createRandom().address, amount: 42 }],
          [{ token: ethers.Wallet.createRandom().address, amount: 42 }],
        ],
        version: 3,
        expire: Math.round(Date.now() / 1000) + 3600,
        chainId: 31_113,
        adjustedMargin,
      };

      signedSwapCall = {
        swapCall,
        sigs: (await signChannelIdAndSwapCall([BrokerA, BrokerB], channelIdAIB, swapCall)) as [
          Signature,
          Signature,
        ],
      };

      swapCallVariablePart = getVariablePart(baseStateAIB);
      swapCallVariablePart.turnNum = 3;
      swapCallVariablePart.appData = encodeSignedSwapCall(signedSwapCall);
      swapCallVariablePart.outcome = singleAssetOutcome(testToken1Address, [
        [brokerADestination, brokerAMargin],
        [brokerBDestination, brokerBMargin],
      ]);

      swapCallCandidate = {
        signedBy: signedBy(0),
        variablePart: swapCallVariablePart,
      };
    });
    describe('succeed', () => {
      it('when supplied first swap call', async () => {
        await MarginAppV1.requireStateSupported(
          fixedPartAIB,
          [recoveredPostFundState, recoveredMarginCallState],
          swapCallCandidate,
        );
      });

      it('when supplied swap call with higher version', async () => {
        marginCall.version = 42;
        recoveredMarginCallState.variablePart.turnNum = 42;
        recoveredMarginCallState.variablePart.appData = await marginCallAppData(
          channelIdAIB,
          marginCall,
          [BrokerA, BrokerB],
        );

        adjustedMargin.version = 43;

        swapCall.version = 43;
        swapCallCandidate.variablePart.turnNum = 43;
        swapCallCandidate.variablePart.appData = await swapCallAppData(channelIdAIB, swapCall, [
          BrokerA,
          BrokerB,
        ]);

        await MarginAppV1.requireStateSupported(
          fixedPartAIB,
          [recoveredPostFundState, recoveredMarginCallState],
          swapCallCandidate,
        );
      });
    });

    describe('revert', () => {
      describe('postfund', () => {
        it('when postfund not supplied', async () => {
          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.reverted;
        });

        it('when postfund.turnNum != 1', async () => {
          recoveredPostFundState.variablePart.turnNum = 42;

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(POSTFUND_TURN_NUM_NOT_EQUAL_1);
        });

        it('when postfund not unanimously signed', async () => {
          recoveredPostFundState.signedBy = signedBy(0);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(POSTFUND_NOT_UNANIMOUS);
        });
      });

      describe('pre swap margin call', () => {
        it('when margin call not supplied', async () => {
          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState],
              swapCallCandidate,
            ),
          ).to.be.reverted;
        });

        it('when not signed', async () => {
          swapCallCandidate.signedBy = SIGNED_BY_NO_ONE;

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(NO_IDENTITY_PROOF_ON_SWAP_CALL);
        });

        it('when turnNum != 2+', async () => {
          recoveredMarginCallState.variablePart.turnNum = 1;

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(MARGIN_CALL_TURN_NUM_LESS_2);
        });

        it('when version != turnNum', async () => {
          // .version = 2
          recoveredMarginCallState.variablePart.turnNum = 3;

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(MARGIN_VERSION_NOT_EQUAL_TURN_NUM);
        });

        it('when outcome != leader margin specified', async () => {
          recoveredMarginCallState.variablePart.outcome = singleAssetOutcome(testToken1Address, [
            [brokerADestination, brokerAMargin],
            [brokerBDestination, brokerBChangedMargin],
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(INCORRECT_LEADER_MARGIN);
        });

        it('when outcome != follower margin specified', async () => {
          recoveredMarginCallState.variablePart.outcome = singleAssetOutcome(testToken1Address, [
            [brokerADestination, brokerAChangedMargin],
            [brokerBDestination, brokerBMargin],
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(INCORRECT_FOLLOWER_MARGIN);
        });

        it('when margin signed by not leader', async () => {
          recoveredMarginCallState.variablePart.appData = await marginCallAppData(
            channelIdAIB,
            marginCall,
            [BrokerC, BrokerB],
          );

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(INVALID_LEADER_SIGNATURE);
        });

        it('when margin signed by not follower', async () => {
          recoveredMarginCallState.variablePart.appData = await marginCallAppData(
            channelIdAIB,
            marginCall,
            [BrokerA, BrokerD],
          );

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(INVALID_FOLLOWER_SIGNATURE);
        });

        it('when garbage encoded', async () => {
          recoveredMarginCallState.variablePart.appData = '0xdeadbeef';

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.reverted;
        });

        describe('outcome', () => {
          it('when outcome sum has changed', async () => {
            marginCall.margin = [brokerAChangedMargin + 1, brokerBChangedMargin + 1];

            recoveredMarginCallState.variablePart.appData = await marginCallAppData(
              channelIdAIB,
              marginCall,
              [BrokerA, BrokerB],
            );
            recoveredMarginCallState.variablePart.outcome = singleAssetOutcome(testToken1Address, [
              [brokerADestination, brokerAChangedMargin + 1],
              [brokerBDestination, brokerBChangedMargin + 1],
            ]);

            await expect(
              MarginAppV1.requireStateSupported(
                fixedPartAIB,
                [recoveredPostFundState, recoveredMarginCallState],
                swapCallCandidate,
              ),
            ).to.be.revertedWith(TOTAL_ALLOCATED_CANNOT_CHANGE);
          });

          it('when outcome with several assets', async () => {
            recoveredMarginCallState.variablePart.outcome = [
              ...singleAssetOutcome(testToken1Address, [
                [brokerADestination, brokerAChangedMargin],
                [brokerBDestination, brokerBChangedMargin],
              ]),
              ...singleAssetOutcome(testToken2Address, [
                [brokerADestination, brokerAChangedMargin],
                [brokerBDestination, brokerBChangedMargin],
              ]),
            ];

            await expect(
              MarginAppV1.requireStateSupported(
                fixedPartAIB,
                [recoveredPostFundState, recoveredMarginCallState],
                swapCallCandidate,
              ),
            ).to.be.revertedWith(INCORRECT_NUMBER_OF_ASSETS);
          });

          it('when more than 2 allocations', async () => {
            recoveredMarginCallState.variablePart.outcome = singleAssetOutcome(testToken1Address, [
              [brokerADestination, brokerAChangedMargin],
              [brokerBDestination, brokerBChangedMargin],
              [intermediaryDestination, brokerAChangedMargin],
            ]);

            await expect(
              MarginAppV1.requireStateSupported(
                fixedPartAIB,
                [recoveredPostFundState, recoveredMarginCallState],
                swapCallCandidate,
              ),
            ).to.be.revertedWith(INCORRECT_NUMBER_OF_ALLOCATIONS);
          });

          it('when destination changed', async () => {
            recoveredMarginCallState.variablePart.outcome = singleAssetOutcome(testToken1Address, [
              [brokerADestination, brokerAChangedMargin],
              [intermediaryDestination, brokerBChangedMargin],
            ]);

            await expect(
              MarginAppV1.requireStateSupported(
                fixedPartAIB,
                [recoveredPostFundState, recoveredMarginCallState],
                swapCallCandidate,
              ),
            ).to.be.revertedWith(DESTINATIONS_CANNOT_CHANGE);
          });
        });
      });

      describe('swap call', () => {
        it('when not signed', async () => {
          swapCallCandidate.signedBy = SIGNED_BY_NO_ONE;

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(NO_IDENTITY_PROOF_ON_SWAP_CALL);
        });

        it('when turnNum != 3+', async () => {
          swapCallCandidate.variablePart.turnNum = 2;

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(SWAP_CALL_TURN_NUM_LESS_3);
        });

        it('when first broker is not participant', async () => {
          swapCall.brokers = [BrokerC.address, BrokerB.address];
          swapCallCandidate.variablePart.appData = await swapCallAppData(channelIdAIB, swapCall, [
            BrokerA,
            BrokerB,
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(FIRST_BROKER_NOT_LEADER);
        });

        it('when second broker is not participant', async () => {
          swapCall.brokers = [BrokerA.address, BrokerD.address];
          swapCallCandidate.variablePart.appData = await swapCallAppData(channelIdAIB, swapCall, [
            BrokerA,
            BrokerB,
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(SECOND_BROKER_NOT_FOLLOWER);
        });

        it('when swap call is not a direct successor of margin call', async () => {
          // .version = 3
          swapCallCandidate.variablePart.turnNum = 4;

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(SWAP_NOT_DIRECT_SUCCESSOR_OF_MARGIN);
        });

        it('when version != turnNum', async () => {
          swapCall.version = 4;
          swapCallCandidate.variablePart.appData = await swapCallAppData(channelIdAIB, swapCall, [
            BrokerA,
            BrokerB,
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(SWAP_VERSION_NOT_EQUAL_TURN_NUM);
        });

        it('when swap signed by not leader', async () => {
          swapCallCandidate.variablePart.appData = await swapCallAppData(channelIdAIB, swapCall, [
            BrokerC,
            BrokerB,
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(INVALID_LEADER_SIGNATURE);
        });

        it('when swap signed by not follower', async () => {
          swapCallCandidate.variablePart.appData = await swapCallAppData(channelIdAIB, swapCall, [
            BrokerA,
            BrokerD,
          ]);

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.revertedWith(INVALID_FOLLOWER_SIGNATURE);
        });

        it('when garbage encoded', async () => {
          swapCallCandidate.variablePart.appData = '0xdeadbeef';

          await expect(
            MarginAppV1.requireStateSupported(
              fixedPartAIB,
              [recoveredPostFundState, recoveredMarginCallState],
              swapCallCandidate,
            ),
          ).to.be.reverted;
        });

        describe('adjusted margin', () => {
          it('when margin.version != swapCall.version', async () => {
            swapCall.adjustedMargin.version = 42;
            swapCallCandidate.variablePart.appData = await swapCallAppData(channelIdAIB, swapCall, [
              BrokerA,
              BrokerB,
            ]);

            await expect(
              MarginAppV1.requireStateSupported(
                fixedPartAIB,
                [recoveredPostFundState, recoveredMarginCallState],
                swapCallCandidate,
              ),
            ).to.be.revertedWith(MARGIN_VERSION_NOT_EQUAL_TURN_NUM);
          });

          it('when outcome != adjusted leader margin specified', async () => {
            swapCallCandidate.variablePart.outcome = singleAssetOutcome(testToken1Address, [
              [brokerADestination, brokerAChangedMargin],
              [brokerBDestination, brokerBMargin],
            ]);

            await expect(
              MarginAppV1.requireStateSupported(
                fixedPartAIB,
                [recoveredPostFundState, recoveredMarginCallState],
                swapCallCandidate,
              ),
            ).to.be.revertedWith(INCORRECT_LEADER_MARGIN);
          });

          it('when outcome != adjusted follower margin specified', async () => {
            swapCallCandidate.variablePart.outcome = singleAssetOutcome(testToken1Address, [
              [brokerADestination, brokerAMargin],
              [brokerBDestination, brokerBChangedMargin],
            ]);

            await expect(
              MarginAppV1.requireStateSupported(
                fixedPartAIB,
                [recoveredPostFundState, recoveredMarginCallState],
                swapCallCandidate,
              ),
            ).to.be.revertedWith(INCORRECT_FOLLOWER_MARGIN);
          });
        });

        describe('outcome', () => {
          it('when outcome sum has changed', async () => {
            swapCall.adjustedMargin.margin = [brokerAMargin + 1, brokerBMargin + 1];

            swapCallCandidate.variablePart.appData = await swapCallAppData(channelIdAIB, swapCall, [
              BrokerA,
              BrokerB,
            ]);
            swapCallCandidate.variablePart.outcome = singleAssetOutcome(testToken1Address, [
              [brokerADestination, brokerAMargin + 1],
              [brokerBDestination, brokerBMargin + 1],
            ]);

            await expect(
              MarginAppV1.requireStateSupported(
                fixedPartAIB,
                [recoveredPostFundState, recoveredMarginCallState],
                swapCallCandidate,
              ),
            ).to.be.revertedWith(TOTAL_ALLOCATED_CANNOT_CHANGE);
          });

          it('when outcome with several assets', async () => {
            swapCallCandidate.variablePart.outcome = [
              ...singleAssetOutcome(testToken1Address, [
                [brokerADestination, brokerAMargin],
                [brokerBDestination, brokerBMargin],
              ]),
              ...singleAssetOutcome(testToken2Address, [
                [brokerADestination, brokerAMargin],
                [brokerBDestination, brokerBMargin],
              ]),
            ];

            await expect(
              MarginAppV1.requireStateSupported(
                fixedPartAIB,
                [recoveredPostFundState, recoveredMarginCallState],
                swapCallCandidate,
              ),
            ).to.be.revertedWith(INCORRECT_NUMBER_OF_ASSETS);
          });

          it('when more than 2 allocations', async () => {
            swapCallCandidate.variablePart.outcome = singleAssetOutcome(testToken1Address, [
              [brokerADestination, brokerAMargin],
              [brokerBDestination, brokerBMargin],
              [intermediaryDestination, brokerAMargin],
            ]);

            await expect(
              MarginAppV1.requireStateSupported(
                fixedPartAIB,
                [recoveredPostFundState, recoveredMarginCallState],
                swapCallCandidate,
              ),
            ).to.be.revertedWith(INCORRECT_NUMBER_OF_ALLOCATIONS);
          });

          it('when destination changed', async () => {
            swapCallCandidate.variablePart.outcome = singleAssetOutcome(testToken1Address, [
              [brokerADestination, brokerAMargin],
              [intermediaryDestination, brokerBMargin],
            ]);

            await expect(
              MarginAppV1.requireStateSupported(
                fixedPartAIB,
                [recoveredPostFundState, recoveredMarginCallState],
                swapCallCandidate,
              ),
            ).to.be.revertedWith(DESTINATIONS_CANNOT_CHANGE);
          });
        });
      });
    });
  });
});
