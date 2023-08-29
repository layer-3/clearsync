import { utils } from 'ethers';

import {
  marginCallTy,
  settlementRequestTy,
  signedMarginCallTy,
  signedSettlementRequestTy,
} from './encodeTypes';

import type {
  MarginCall,
  SettlementRequest,
  SignedMarginCall,
  SignedSettlementRequest,
} from './types';

export function encodeMarginCall(marginCall: MarginCall): string {
  return utils.defaultAbiCoder.encode([marginCallTy], [marginCall]);
}

export function encodeChannelIdAndMarginCall(channelId: string, marginCall: MarginCall): string {
  return utils.defaultAbiCoder.encode(['bytes32', marginCallTy], [channelId, marginCall]);
}

export function encodeSignedMarginCall(signedMarginCall: SignedMarginCall): string {
  return utils.defaultAbiCoder.encode([signedMarginCallTy], [signedMarginCall]);
}

export function encodeChannelIdAndSignedMarginCall(
  channelId: string,
  signedMarginCall: SignedMarginCall,
): string {
  return utils.defaultAbiCoder.encode(
    ['bytes32', signedMarginCallTy],
    [channelId, signedMarginCall],
  );
}

export function encodeSettlementRequest(settlementRequest: SettlementRequest): string {
  return utils.defaultAbiCoder.encode([settlementRequestTy], [settlementRequest]);
}

export function encodeChannelIdAndSettlementRequest(
  channelId: string,
  settlementRequest: SettlementRequest,
): string {
  return utils.defaultAbiCoder.encode(
    ['bytes32', settlementRequestTy],
    [channelId, settlementRequest],
  );
}

export function encodeSignedSettlementRequest(
  signedSettlementRequest: SignedSettlementRequest,
): string {
  return utils.defaultAbiCoder.encode([signedSettlementRequestTy], [signedSettlementRequest]);
}

export function encodeChannelIdAndSignedSettlementRequest(
  channelId: string,
  signedSettlementRequest: SignedSettlementRequest,
): string {
  return utils.defaultAbiCoder.encode(
    ['bytes32', signedSettlementRequestTy],
    [channelId, signedSettlementRequest],
  );
}
