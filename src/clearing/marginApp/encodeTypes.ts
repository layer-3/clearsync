import type { ParamType } from 'ethers/lib/utils';

export const signaturesTy = {
  name: 'sigs',
  type: 'tuple[2]',
  components: [
    { name: 'v', type: 'uint8' },
    { name: 'r', type: 'bytes32' },
    { name: 's', type: 'bytes32' },
  ],
} as ParamType;

export const marginCallTy = {
  name: 'marginCall',
  type: 'tuple',
  components: [
    { name: 'version', type: 'uint64' },
    { name: 'margin', type: 'uint256[2]' },
  ],
} as ParamType;

export const signedMarginCallTy = {
  type: 'tuple',
  components: [marginCallTy, signaturesTy],
} as ParamType;

export const LiabilityTy = {
  type: 'tuple',
  components: [
    { name: 'token', type: 'address' },
    { name: 'amount', type: 'uint256' },
  ],
} as ParamType;

export const swapCallTy = {
  name: 'swapCall',
  type: 'tuple',
  components: [
    { name: 'brokers', type: 'address[2]' },
    { ...LiabilityTy, type: 'tuple[][2]', name: 'swaps' },
    { name: 'version', type: 'uint64' },
    { name: 'expire', type: 'uint64' },
    { name: 'chainId', type: 'uint256' },
    { ...marginCallTy, name: 'adjustedMargin' },
  ],
} as ParamType;

export const signedSwapCallTy = {
  type: 'tuple',
  components: [swapCallTy, signaturesTy],
} as ParamType;
