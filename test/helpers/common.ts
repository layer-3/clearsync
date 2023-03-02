import { utils } from 'ethers';

export function ACCOUNT_MISSING_ROLE(account: string, role: string): string {
  return `AccessControl: account ${utils.hexlify(account)} is missing role ${role}`;
}

export function randomHex(size: number): string {
  return Array.from({ length: size }, () => Math.floor(Math.random() * 16).toString(16)).join('');
}

export function randomBytes32(): string {
  return '0x' + randomHex(64);
}
