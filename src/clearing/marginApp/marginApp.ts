import { encodeSignedMarginCall, encodeSignedSettlementRequest } from './encode';
import { signChannelIdAndMarginCall, signChannelIdAndSettlementRequest } from './signatures';

import type { MarginCall, SettlementRequest } from './types';
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

export async function settlementRequestAppData(
  channelId: string,
  settlementRequest: SettlementRequest,
  signers: SignerWithAddress[],
): Promise<string> {
  return encodeSignedSettlementRequest({
    settlementRequest,
    sigs: (await signChannelIdAndSettlementRequest(signers, channelId, settlementRequest)) as [
      Signature,
      Signature,
    ],
  });
}
