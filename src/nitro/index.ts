import FULLNitroAdjudicatorArtifact from '../../artifacts/contracts/nitro/NitroAdjudicator.sol/NitroAdjudicator.json';
import FULLConsensusAppArtifact from '../../artifacts/contracts/nitro/ConsensusApp.sol/ConsensusApp.json';
import FULLVirtualPaymentAppArtifact from '../../artifacts/contracts/nitro/VirtualPaymentApp.sol/VirtualPaymentApp.json';

export const ContractArtifacts = {
  NitroAdjudicatorArtifact: FULLNitroAdjudicatorArtifact,
  ConsensusAppArtifact: FULLConsensusAppArtifact,
  VirtualPaymentAppArtifact: FULLVirtualPaymentAppArtifact,
};

export {
  DepositedEvent,
  getDepositedEvent,
  convertBytes32ToAddress,
  convertAddressToBytes32,
} from './contract/multi-asset-holder';
export {
  getChallengeRegisteredEvent,
  getChallengeClearedEvent,
  ChallengeRegisteredEvent,
} from './contract/challenge';
export {getChannelId, isExternalDestination} from './contract/channel';
export {
  validTransition,
  ForceMoveAppContractInterface,
  createValidTransitionTransaction,
} from './contract/force-move-app';
export {encodeOutcome, decodeOutcome, Outcome, AssetOutcome, hashOutcome} from './contract/outcome';
export {channelDataToStatus} from './contract/channel-storage';
export {getSignedBy} from './bitfield-utils';

export {
  State,
  FixedPart,
  VariablePart,
  RecoveredVariablePart,
  getVariablePart,
  getFixedPart,
  hashState,
} from './contract/state';

export * from './signatures';
export * from './transactions';
export * from './contract/vouchers';

// types
export {Uint256, Bytes32} from './contract/types';

export * from './channel-mode';

export {OutcomeShortHand, AssetOutcomeShortHand, computeOutcome, getRandomNonce} from './helpers';
