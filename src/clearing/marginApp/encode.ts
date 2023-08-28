import { utils } from 'ethers';

import { marginCallTy, signedMarginCallTy, signedSwapCallTy, swapCallTy } from './encodeTypes';

import type { MarginCall, SignedMarginCall, SignedSwapCall, SwapCall } from './types';

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

export function encodeSwapCall(swapCall: SwapCall): string {
  return utils.defaultAbiCoder.encode([swapCallTy], [swapCall]);
}

export function encodeChannelIdAndSwapCall(channelId: string, swapCall: SwapCall): string {
  return utils.defaultAbiCoder.encode(['bytes32', swapCallTy], [channelId, swapCall]);
}

export function encodeSignedSwapCall(signedSwapCall: SignedSwapCall): string {
  return utils.defaultAbiCoder.encode([signedSwapCallTy], [signedSwapCall]);
}

export function encodeChannelIdAndSignedSwapCall(
  channelId: string,
  signedSwapCall: SignedSwapCall,
): string {
  return utils.defaultAbiCoder.encode(['bytes32', signedSwapCallTy], [channelId, signedSwapCall]);
}
