import { utils } from 'ethers';

import type { ParamType } from 'ethers/lib/utils';

export const REFERRAL_MAX_DEPTH = 5;

export enum VoucherAction {
  MintPack,
  MeldFlock,
}

export interface MintParams {
  to: string;
  size: number;
  isTransferable: boolean;
}

export interface MeldParams {
  owner: string;
  tokenIds: number[];
  isTransferable: boolean;
}

export const MintParamsABI = {
  name: 'mintParams',
  type: 'tuple',
  components: [
    {
      name: 'to',
      type: 'address',
    },
    {
      name: 'size',
      type: 'uint256',
    },
    {
      name: 'isTransferable',
      type: 'bool',
    },
  ],
} as ParamType;

export const MeldParamsABI = {
  name: 'meldParams',
  type: 'tuple',
  components: [
    {
      name: 'owner',
      type: 'address',
    },
    {
      name: 'tokenIds',
      type: 'uint256[]',
    },
    {
      name: 'isTransferable',
      type: 'bool',
    },
  ],
} as ParamType;

export function encodeMintParams(mintParams: MintParams): string {
  return utils.defaultAbiCoder.encode([MintParamsABI], [mintParams]);
}

export function encodeMeldParams(meldParams: MeldParams): string {
  return utils.defaultAbiCoder.encode([MeldParamsABI], [meldParams]);
}
