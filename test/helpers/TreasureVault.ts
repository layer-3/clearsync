import { utils } from 'ethers';

import type { ParamType } from 'ethers/lib/utils';

export const REFERRAL_MAX_DEPTH = 5;

export enum VoucherAction {
  Reward,
}

export interface RewardParams {
  token: string;
  amount: number;
  commissions: number[];
}

export const RewardParamsABI = {
  name: 'rewardParams',
  type: 'tuple',
  components: [
    {
      name: 'token',
      type: 'address',
    },
    {
      name: 'amount',
      type: 'uint256',
    },
    {
      name: 'commissions',
      type: 'uint8[5]',
    },
  ],
} as ParamType;

export function encodeRewardParams(rewardParams: RewardParams): string {
  return utils.defaultAbiCoder.encode([RewardParamsABI], [rewardParams]);
}
