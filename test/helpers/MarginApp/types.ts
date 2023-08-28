import type { Signature } from '@ethersproject/bytes';

export interface MarginCall {
  version: number;
  margin: [number, number];
}

export interface SignedMarginCall {
  marginCall: MarginCall;
  sigs: [Signature, Signature];
}

export interface Liability {
  token: string;
  amount: number;
}

export interface SwapCall {
  brokers: [string, string];
  swaps: [Liability[], Liability[]];
  version: number;
  expire: number;
  chainId: number;
  adjustedMargin: MarginCall;
}

export interface SignedSwapCall {
  swapCall: SwapCall;
  sigs: [Signature, Signature];
}
