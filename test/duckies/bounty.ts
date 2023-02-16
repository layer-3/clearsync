import { utils } from 'ethers';

import { randomBytes32 } from '../helpers/payload';

import type { ParamType } from 'ethers/lib/utils';

export interface Bounty {
  amount: number;
  tokenAddress: string;
  beneficiary: string;
  isPaidToReferrers: boolean;
  referrer: string;
  expire: number;
  chainId: number;
  bountyCodeHash: string;
}

export const BountyTy = {
  name: 'bounty',
  type: 'tuple',
  components: [
    {
      name: 'amount',
      type: 'uint256',
    },
    {
      name: 'tokenAddress',
      type: 'address',
    },
    {
      name: 'beneficiary',
      type: 'address',
    },
    {
      name: 'isPaidToReferrers',
      type: 'bool',
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
      name: 'bountyCodeHash',
      type: 'bytes32',
    },
  ],
} as ParamType;

export function encodeBounty(bounty: Bounty): string {
  return utils.defaultAbiCoder.encode([BountyTy], [bounty]);
}

export function setBountyField<BountyKey extends keyof Bounty>(
  bounty: Bounty,
  field: BountyKey,
  value: Bounty[BountyKey],
): Bounty {
  return {
    ...bounty,
    [field]: value,
  };
}

export function setRandomBountyCodeHash(bounty: Bounty): Bounty {
  return { ...bounty, bountyCodeHash: randomBytes32() };
}
