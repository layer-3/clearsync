import { encodeSignedMarginCall, encodeSignedSwapCall } from './encode';
import { signChannelIdAndMarginCall, signChannelIdAndSwapCall } from './signatures';

import type { MarginCall, SwapCall } from './types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { Signature } from '@ethersproject/bytes';

export async function marginCallAppData(
  channelId: string,
  marginCall: MarginCall,
  signers: SignerWithAddress[],
): Promise<string> {
  return encodeSignedMarginCall({
    marginCall,
    sigs: (await signChannelIdAndMarginCall(signers, channelId, marginCall)) as [
      Signature,
      Signature,
    ],
  });
}

export async function swapCallAppData(
  channelId: string,
  swapCall: SwapCall,
  signers: SignerWithAddress[],
): Promise<string> {
  return encodeSignedSwapCall({
    swapCall,
    sigs: (await signChannelIdAndSwapCall(signers, channelId, swapCall)) as [Signature, Signature],
  });
}
