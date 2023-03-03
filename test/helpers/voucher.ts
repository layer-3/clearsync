import { Signer, utils } from 'ethers';

import { signEncoded } from './sign';

import type { ParamType } from 'ethers/lib/utils';

export interface Voucher {
  target: string;
  action: number;
  beneficiary: string;
  referrer: string;
  expire: number;
  chainId: number;
  voucherCodeHash: string;
  encodedParams: string;
}

export const VoucherABI = {
  name: 'voucher',
  type: 'tuple',
  components: [
    {
      name: 'target',
      type: 'address',
    },
    {
      name: 'action',
      type: 'uint8',
    },
    {
      name: 'beneficiary',
      type: 'address',
    },
    {
      name: 'referrer',
      type: 'address',
    },
    {
      name: 'expire',
      type: 'uint64',
    },
    {
      name: 'chainId',
      type: 'uint32',
    },
    {
      name: 'voucherCodeHash',
      type: 'bytes32',
    },
    {
      name: 'encodedParams',
      type: 'bytes',
    },
  ],
} as ParamType;

export const VouchersABI = {
  type: 'tuple[]',
  components: VoucherABI.components,
} as ParamType;

export function encodeVoucher(voucher: Voucher): string {
  return utils.defaultAbiCoder.encode([VoucherABI], [voucher]);
}

export function encodeVouchers(vouchers: Voucher[]): string {
  return utils.defaultAbiCoder.encode([VouchersABI], [vouchers]);
}

export async function signVoucher(voucher: Voucher, signer: Signer): Promise<string> {
  return signEncoded(encodeVoucher(voucher), signer);
}

export async function signVouchers(vouchers: Voucher[], signer: Signer): Promise<string> {
  return signEncoded(encodeVouchers(vouchers), signer);
}
