import { Signer, utils } from 'ethers';

export async function signEncoded(encodedData: string, signer: Signer): Promise<string> {
  return await signer.signMessage(utils.arrayify(utils.keccak256(encodedData)));
}
