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

// flag
export const flagsGeneIdx = 30;
export enum Flags {
  isTransferable = 1,
}

// magic number
export const magicNumberGeneIdx = 31;
export const baseMagicNumber = 209;
export const mythicMagicNumber = 210;

export const generativeGenesOffset = 2;

export enum GeneralGenes {
  Collection,
  Flags = 30,
  MagicNumber = 31,
}

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
  Flags = 30,
  MagicNumber = 31,
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
  Flags = 30,
  MagicNumber = 31,
}

export enum MythicGenes {
  Collection,
  UniqId,
  Temper,
  Skill,
  Habitat,
  Breed,
  Birthplace,
  Quirk,
  FavoriteFood,
  FavoriteColor,
  Flags = 30,
  MagicNumber = 31,
}

export const mythicAmount = 60;
export const MAX_PECULIARITY = 145;
export const MYTHIC_DISPERSION = 5;

export type Genes = DucklingGenes | ZombeakGenes | MythicGenes | GeneralGenes;
export type GeneIdxToValue = { [key in Genes]: number };
export type GeneToValue =
  | keyof typeof DucklingGenes
  | keyof typeof ZombeakGenes
  | keyof typeof MythicGenes
  | keyof typeof GeneralGenes;

export const collectionsGeneValuesNum = [
  // Duckling genes: (Collection, Rarity), Color, Family, Body, Head, Eyes, Beak, Wings, FirstName, Temper, Skill, Habitat, Breed
  [4, 5, 10, 25, 30, 14, 10, 36, 16, 12, 5, 28],
  // Zombeak genes: (Collection, Rarity), Color, Family, Body, Head, Eyes, Beak, Wings, FirstName, Temper, Skill, Habitat, Breed
  [2, 3, 7, 6, 9, 7, 10, 36, 16, 12, 5, 28],
  // Mythic genes: (Collection, UniqId), Temper, Skill, Habitat, Breed, Birthplace, Quirk, Favorite Food, Favorite Color
  [16, 12, 5, 28, 5, 10, 8, 4],
] as const;

export const collectionsGeneDistributionTypes = [
  2940, // reverse(001111101101) = 101101111100
  2940, // reverse(001111101101) = 101101111100
  107, // reverse(11010110) = 01101011
] as const;

export const MAX_PACK_SIZE = 50;
export const FLOCK_SIZE = 5;

export const collectionMutationChances = [150, 100, 50, 10];
export const geneMutationChance = [955, 45];
export const geneInheritanceChances = [400, 300, 150, 100, 50];
