import { utils } from 'ethers';

export function ACCOUNT_MISSING_ROLE(account: string, role: string): string {
  return `AccessControl: account ${utils.hexlify(account)} is missing role ${role}`;
}
