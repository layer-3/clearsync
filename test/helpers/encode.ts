import { utils } from 'ethers';

import type { ParamType } from 'ethers/lib/utils';

export function ABIEncode(ABIType: ParamType, data: unknown): string {
  return utils.defaultAbiCoder.encode([ABIType], [data]);
}
