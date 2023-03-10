import { ethers } from 'hardhat';
import type { Collection } from '../src/duckies/ducklings/collection';
import type { Ducklings } from '../typechain-types';

// data must be set up by hand before invoking the script
async function main(): Promise<void> {
  const ducklingsAddress = '';
  const Ducklings = (await ethers.getContractAt('Ducklings', ducklingsAddress)) as Ducklings;

  const collection: Collection = {
    availableBefore: Math.round(new Date('2023-11-02T12:00:00').getTime() / 1000), // type(uint64).max if collection is indefinite, timestamp otherwise
    isMeldable: true,
    traitWeights: [
      // Class
      [74, 20, 5, 1],
      // Body
      [0, 17, 16, 14, 14, 12, 10, 9, 8],
      // Head
      [0, 9, 9, 7, 6, 6, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 3, 3],
      // Background
      [4],
      // Element
      [5],
      // Eyes
      [8, 7, 6, 6, 6, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 3, 3, 3, 2, 2],
      // Beak
      [15, 14, 14, 11, 10, 9, 8, 7, 7, 6],
      // Wings
      [20, 19, 15, 13, 12, 11, 10],
      // First name
      [32],
      // Last name
      [17],
      // Temper
      [16],
      // Peculiarity - for SuperLegendary
      [100],
    ],
  };

  await Ducklings.addCollection(collection);

  console.log(`Collection added`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
