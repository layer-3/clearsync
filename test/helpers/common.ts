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

export function encodeError(errorSignature: string, ...args: unknown[]): string {
  const leftParenthesisIdx = errorSignature.indexOf('(');
  const rightParenthesisIdx = errorSignature.indexOf(')');

  if (
    leftParenthesisIdx > rightParenthesisIdx ||
    leftParenthesisIdx == -1 ||
    rightParenthesisIdx == -1
  ) {
    throw new Error('Error signature must contain parenthesis');
  }

  const errorSelector = utils.id(errorSignature).slice(0, 10);
  const types = errorSignature.slice(leftParenthesisIdx + 1, rightParenthesisIdx).split(',');
  const encodedParams = utils.defaultAbiCoder.encode(types, args);
  return errorSelector + encodedParams.slice(2);
}
