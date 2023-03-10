export type TraitWeights = number[][];

export interface Collection {
  availableBefore: number;
  isMeldable: boolean;
  traitWeights: TraitWeights;
}

export const ZOMBEAKS_COLLECTION_ID = 0;
