import { Signature, utils } from 'ethers';

import { encodeChannelIdAndMarginCall, encodeChannelIdAndSettlementRequest } from './encode';

import type { MarginCall, SettlementRequest } from './types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

export async function signEncoded(
  signer: SignerWithAddress,
  encodedData: string,
): Promise<Signature> {
  return utils.splitSignature(
    await signer.signMessage(utils.arrayify(utils.keccak256(encodedData))),
  );
}

export async function signChannelIdAndMarginCall(
  signers: SignerWithAddress[],
  channelId: string,
  marginCall: MarginCall,
): Promise<Signature[]> {
  return Promise.all(
    signers.map((signer) =>
      signEncoded(signer, encodeChannelIdAndMarginCall(channelId, marginCall)),
    ),
  );
}

export async function signChannelIdAndSettlementRequest(
  signers: SignerWithAddress[],
  channelId: string,
  settlementRequest: SettlementRequest,
): Promise<Signature[]> {
  return Promise.all(
    signers.map((signer) =>
      signEncoded(signer, encodeChannelIdAndSettlementRequest(channelId, settlementRequest)),
    ),
  );
}

export const SIGNED_BY_NO_ONE = '0';

export function signedBy(index: number | number[]): string {
  if (Array.isArray(index)) {
    let res = 0;
    for (const idx of index) res += _signedBy(idx);
    return res.toString();
  } else {
    return _signedBy(index).toString();
  }
}

function _signedBy(index: number): number {
  return 2 ** index;
}
