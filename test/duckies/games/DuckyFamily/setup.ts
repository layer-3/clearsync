import { constants, utils } from 'ethers';
import { ethers, upgrades } from 'hardhat';

import type {
  DucklingsV1,
  TESTDuckyFamilyV1,
  TreasureVault,
  YellowToken,
} from '../../../../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

export const ADMIN_ROLE = constants.HashZero;
export const GAME_ROLE = utils.id('GAME_ROLE');
export const MAINTAINER_ROLE = utils.id('MAINTAINER_ROLE');

interface Config {
  Duckies: YellowToken;
  Ducklings: DucklingsV1;
  TreasureVault: TreasureVault;
  Game: TESTDuckyFamilyV1;
  GameAsMaintainer: TESTDuckyFamilyV1;
  GameAsSomeone: TESTDuckyFamilyV1;
  GameAsSomeother: TESTDuckyFamilyV1;
  Admin: SignerWithAddress;
  Maintainer: SignerWithAddress;
  Someone: SignerWithAddress;
  Someother: SignerWithAddress;
  GenomeSetter: SignerWithAddress;
}

export async function setup(): Promise<Config> {
  const [Admin, Maintainer, Someone, Someother, GenomeSetter] = await ethers.getSigners();

  const DuckiesFactory = await ethers.getContractFactory('YellowToken');
  const Duckies = (await DuckiesFactory.deploy(
    'Duckies',
    'DUCKIES',
    1_000_000 * 10e8,
  )) as YellowToken;
  await Duckies.deployed();

  await Duckies.activate(100_000_000_000_000, Admin.address);
  await Duckies.mint(Someone.address, 100_000_000_000_000);
  await Duckies.mint(Someother.address, 100_000_000_000_000);

  const DucklingsFactory = await ethers.getContractFactory('DucklingsV1');
  const Ducklings = (await upgrades.deployProxy(DucklingsFactory, [], {
    kind: 'uups',
  })) as DucklingsV1;
  await Ducklings.deployed();

  const TreasureVaultFactory = await ethers.getContractFactory('TreasureVault');
  const TreasureVault = (await upgrades.deployProxy(TreasureVaultFactory, [], {
    kind: 'uups',
  })) as TreasureVault;
  await TreasureVault.deployed();

  const DuckyFamilyFactory = await ethers.getContractFactory('TESTDuckyFamilyV1');
  const Game = (await DuckyFamilyFactory.deploy(
    Duckies.address,
    Ducklings.address,
    TreasureVault.address,
  )) as TESTDuckyFamilyV1;
  await Game.deployed();

  await Duckies.connect(Someone).increaseAllowance(Game.address, 10_000_000_000);
  await Duckies.connect(Someother).increaseAllowance(Game.address, 10_000_000_000);

  await Ducklings.grantRole(GAME_ROLE, Game.address);
  await Ducklings.grantRole(GAME_ROLE, GenomeSetter.address);

  await Game.grantRole(MAINTAINER_ROLE, Maintainer.address);

  return {
    Duckies,
    Ducklings,
    TreasureVault,
    Game,
    GameAsMaintainer: Game.connect(Maintainer),
    GameAsSomeone: Game.connect(Someone),
    GameAsSomeother: Game.connect(Someother),
    Admin,
    Maintainer,
    Someone,
    Someother,
    GenomeSetter,
  };
}
