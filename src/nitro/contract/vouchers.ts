import {BigNumberish, Signature, Wallet} from 'ethers';
import {defaultAbiCoder, keccak256, ParamType} from 'ethers/lib/utils';

import {sign} from '../signatures';

import {Bytes32} from './types';

export interface Voucher {
  channelId: Bytes32;
  amount: string;
}

export interface VoucherAmountAndSignature {
  amount: string;
  signature: Signature; // signature on (channelId,amount)
}

const voucherTy = {
  type: 'tuple',
  components: [
    {name: 'channelId', type: 'bytes32'},
    {
      name: 'amount',
      type: 'uint256',
    },
  ],
} as ParamType;

export async function signVoucher(voucher: Voucher, wallet: Wallet): Promise<Signature> {
  return sign(wallet, keccak256(defaultAbiCoder.encode([voucherTy], [voucher])));
}

const voucherAmountAndSignatureTy = {
  type: 'tuple',
  components: [
    {
      name: 'amount',
      type: 'uint256',
    },
    {
      type: 'tuple',
      name: 'signature',
      components: [
        {name: 'v', type: 'uint8'},
        {name: 'r', type: 'bytes32'},
        {name: 's', type: 'bytes32'},
      ],
    } as ParamType,
  ],
} as ParamType;

export function encodeVoucherAmountAndSignature(amount: BigNumberish, signature: Signature) {
  return defaultAbiCoder.encode([voucherAmountAndSignatureTy], [{amount, signature}]);
}
