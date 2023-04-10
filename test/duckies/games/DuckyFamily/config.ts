export enum Collections {
  Duckling = 0,
  Zombeak,
  Mythic,
}

export const raritiesNum = 4;
export enum Rarities {
  Common,
  Rare,
  Epic,
  Legendary,
}

export enum GeneDistrTypes {
  Even,
  Uneven,
}

export const collectionGeneIdx = 0;
export const rarityGeneIdx = 1;

export const generativeGenesOffset = 2;

export enum DucklingGenes {
  Collection,
  Rarity,
  Color,
  Family,
  Body,
  Head,
  Eyes,
  Beak,
  Wings,
  FirstName,
  Temper,
  Skill,
  Habitat,
  Breed,
}

export enum ZombeakGenes {
  Collection,
  Rarity,
  Color,
  Family,
  Body,
  Head,
  Eyes,
  Beak,
  Wings,
  FirstName,
  Temper,
  Skill,
  Habitat,
  Breed,
}

export enum MythicGenes {
  Collection,
  UniqId,
}

export const mythicAmount = 59;
export const MAX_PECULIARITY = 145;
export const MYTHIC_DISPERSION = 5;

export type Genes = DucklingGenes | ZombeakGenes | MythicGenes;

export const collectionsGeneValuesNum = [
  // Duckling genes: (Collection, Rarity), Color, Family, Body, Head, Eyes, Beak, Wings, FirstName, Temper, Skill, Habitat, Breed
  [4, 5, 10, 25, 30, 14, 10, 36, 16, 12, 5, 28],
  // Zombeak genes: (Collection, Rarity), Color, Family, Body, Head, Eyes, Beak, Wings, FirstName, Temper, Skill, Habitat, Breed
  [2, 3, 7, 6, 9, 7, 10, 36, 16, 12, 5, 28],
] as const;

export const collectionsGeneDistributionTypes = [
  2940, // reverse(001111101101) = 101101111100
  2940, // reverse(001111101101) = 101101111100
] as const;

export const FLOCK_SIZE = 5;
