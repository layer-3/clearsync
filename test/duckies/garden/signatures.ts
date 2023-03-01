import { utils } from 'ethers';

import { Bounty, encodeBounty } from './bounty';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

export async function signEncoded(encodedData: string, signer: SignerWithAddress): Promise<string> {
  return await signer.signMessage(utils.arrayify(utils.keccak256(encodedData)));
}

export async function signBounty(bounty: Bounty, signer: SignerWithAddress): Promise<string> {
  return await signEncoded(encodeBounty(bounty), signer);
}
