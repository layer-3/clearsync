import {BigNumber} from 'ethers';

export function addHex(a: string, b: string): string {
  return BigNumber.from(a).add(BigNumber.from(b)).toHexString();
}
export function subHex(a: string, b: string): string {
  return BigNumber.from(a).sub(BigNumber.from(b)).toHexString();
}

export function eqHex(a: string, b: string): boolean {
  return BigNumber.from(a).eq(b);
}

export function toHex(a: number): string {
  return BigNumber.from(a).toHexString();
}

export function fromHex(a: string): number {
  return BigNumber.from(a).toNumber();
}
