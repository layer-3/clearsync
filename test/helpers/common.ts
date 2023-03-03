import { Signer, utils } from 'ethers';
import type { ParamType } from 'ethers/lib/utils';

export function ABIEncode(ABIType: ParamType, data: unknown): string {
  return utils.defaultAbiCoder.encode([ABIType], [data]);
}

export async function signEncoded(encodedData: string, signer: Signer): Promise<string> {
  return await signer.signMessage(utils.arrayify(utils.keccak256(encodedData)));
}

export function ACCOUNT_MISSING_ROLE(account: string, role: string): string {
  return `AccessControl: account ${utils.hexlify(account)} is missing role ${role}`;
}

export function randomHex(size: number): string {
  return Array.from({ length: size }, () => Math.floor(Math.random() * 16).toString(16)).join('');
}

export function randomBytes32(): string {
  return '0x' + randomHex(64);
}
