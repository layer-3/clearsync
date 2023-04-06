import { utils } from 'ethers';
import { ethers, upgrades } from 'hardhat';
import { assert, expect } from 'chai';

import { connectGroup } from '../../../helpers/connect';
import { ACCOUNT_MISSING_ROLE } from '../../../helpers/common';

import {
  Collections,
  DucklingGenes,
  GeneDistrTypes,
  MAX_PECULIARITY,
  Rarities,
  collectionsGeneValuesNum,
  generativeGenesOffset,
} from './config';
import { Genome } from './genome';
import { randomGenome } from './helpers';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type {
  Ducklings,
  DuckyFamilyV1,
  TESTDuckyFamilyV1,
  TreasureVault,
  YellowToken,
} from '../../../../typechain';

const GAME_ROLE = utils.id('GAME_ROLE');
const MAINTAINER_ROLE = utils.id('MAINTAINER_ROLE');

describe('DuckyFamilyV1_0', () => {
  let Admin: SignerWithAddress;
  let Maintainer: SignerWithAddress;
  let Someone: SignerWithAddress;
  let Someother: SignerWithAddress;
  let GenomeSetter: SignerWithAddress;

  let Duckies: YellowToken;
  let Ducklings: Ducklings;
  let TreasureVault: TreasureVault;
  let Game: TESTDuckyFamilyV1;

  let GameAsMaintainer: DuckyFamilyV1;
  let GameAsSomeone: DuckyFamilyV1;

  const mintTo = async (to: string, genome: bigint, isTransferable?: boolean): Promise<void> => {
    await Ducklings.connect(GenomeSetter).mintTo(to, genome, isTransferable ?? true);
  };

  beforeEach(async () => {
    [Admin, Maintainer, Someone, Someother, GenomeSetter] = await ethers.getSigners();

    const DuckiesFactory = await ethers.getContractFactory('YellowToken');
    Duckies = (await DuckiesFactory.deploy('Duckies', 'DUCKIES', 1_000_000 * 10e8)) as YellowToken;
    await Duckies.deployed();

    await Duckies.activate(100_000_000_000_000, Admin.address);
    await Duckies.mint(Someone.address, 100_000_000_000_000);
    await Duckies.mint(Someother.address, 100_000_000_000_000);

    const DucklingsFactory = await ethers.getContractFactory('Ducklings');
    Ducklings = (await upgrades.deployProxy(DucklingsFactory, [], { kind: 'uups' })) as Ducklings;
    await Ducklings.deployed();

    const TreasureVaultFactory = await ethers.getContractFactory('TreasureVault');
    TreasureVault = (await upgrades.deployProxy(TreasureVaultFactory, [], {
      kind: 'uups',
    })) as TreasureVault;
    await TreasureVault.deployed();

    const DuckyFamilyFactory = await ethers.getContractFactory('TESTDuckyFamilyV1');
    Game = (await DuckyFamilyFactory.deploy(
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

    [GameAsMaintainer, GameAsSomeone] = connectGroup(Game, [Maintainer, Someone]);
  });

  describe('vouchers', () => {
    describe('issuer', () => {
      it('admin can set issuer');

      it('revert on not admin set issuer');
    });

    describe('useVoucher', () => {
      describe('general revert', () => {
        it('revert on incorrect signer');

        it('revert on using same voucher for second time');

        it('revert on target address != contract address');

        it('revert on beneficiary != sender');

        it('revert on expired voucher');

        it('revert on wrong chainId');

        it('revert on invalid voucher action');
      });

      describe('mint voucher', () => {
        it('successfuly use mint voucher');

        it('duckies are not paid for mint');

        it('revert on to == address(0)');

        it('revert on size == 0');

        it('revert on size > MAX_PACK_SIZE');

        it('event emitted');
      });

      describe('meld voucher', () => {
        it('successfuly use meld voucher');

        it('duckies are not paid for meld');

        it('revert on owner == address(0)');

        it('revert on number of tokens != FLOCK_SIZE');

        it('event is emitted');
      });
    });

    describe('useVouchers', () => {
      it('can use several mint vouchers');

      it('can use several meld vouchers');

      it('revert on incorrect signer');
    });
  });

  describe('NFT game', () => {
    describe('prices', () => {
      const MINT_PRICE = 5;
      const MELD_PRICES = [10, 20, 50, 100];

      it('maintainer can set mint price', async () => {
        await GameAsMaintainer.setMintPrice(MINT_PRICE);
        expect(await Game.mintPrice()).to.equal(MINT_PRICE);
      });

      it('revert on not maintainer set mint price', async () => {
        await expect(GameAsSomeone.setMintPrice(MINT_PRICE)).to.be.revertedWith(
          ACCOUNT_MISSING_ROLE(Someone.address, MAINTAINER_ROLE),
        );
      });

      it('maintainer can set meld price', async () => {
        await GameAsMaintainer.setMeldPrices(MELD_PRICES);
        expect(await Game.getMeldPrices()).to.deep.equal(MELD_PRICES);
      });

      it('revert on not maintainer set meld price', async () => {
        await expect(GameAsSomeone.setMeldPrices(MELD_PRICES)).to.be.revertedWith(
          ACCOUNT_MISSING_ROLE(Someone.address, MAINTAINER_ROLE),
        );
      });
    });

    describe('helpers', () => {
      it('_getDistributionType', async () => {
        expect(await Game.getDistributionType(0b0, 0)).to.equal(GeneDistrTypes.Even);
        expect(await Game.getDistributionType(0b1, 0)).to.equal(GeneDistrTypes.Uneven);
        expect(await Game.getDistributionType(0b10, 1)).to.equal(GeneDistrTypes.Uneven);
        expect(await Game.getDistributionType(0b010, 2)).to.equal(GeneDistrTypes.Even);
        expect(await Game.getDistributionType(0b0010_0010, 7)).to.equal(GeneDistrTypes.Even);
        expect(await Game.getDistributionType(0b100_0010_0010, 10)).to.equal(GeneDistrTypes.Uneven);
      });

      it('_calcMaxPeculiarity', async () => {
        expect(await Game.calcMaxPeculiarity()).to.equal(MAX_PECULIARITY);
      });

      it('_calcPeculiarity', async () => {
        const geneValues = Array.from({
          length: collectionsGeneValuesNum[Collections.Duckling].length + generativeGenesOffset,
        }).fill(1) as number[];
        expect(await Game.calcPeculiarity(Genome.from(geneValues).genome)).to.equal(8); // 8 uneven genes
        expect(
          await Game.calcPeculiarity(
            Genome.from([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14]).genome,
          ),
        ).to.equal(5 + 6 + 7 + 8 + 9 + 11 + 12 + 14); // 001111101101
      });
    });

    describe('minting', () => {
      describe('generateGenome', () => {
        it('duckling has correct structure');

        it('zombeak has correct structure');

        it('mythic has correct structure');
      });

      describe('generateAndSetGenes', () => {
        it('has correct numbers of genes');

        it('does not exceed max gene values');
      });

      describe('mintPack', () => {
        it('duckies are paid for mint');

        it('correct amount of tokens is minted');

        it('revert on amount == 0');

        it('revert on amount > MAX_PACK_SIZE');

        it('event is emitted');
      });
    });

    describe('melding', () => {
      describe('meldGenes', () => {
        it('can meld', async () => {
          await mintTo(
            Someone.address,
            // eslint-disable-next-line unicorn/numeric-separators-style
            182700775082802730930410854023168n,
          );

          await mintTo(
            Someone.address,
            // eslint-disable-next-line unicorn/numeric-separators-style
            60926767771370839915004195766272n,
          );

          await mintTo(
            Someone.address,
            // eslint-disable-next-line unicorn/numeric-separators-style
            121932763563511447839369064611840n,
          );

          await mintTo(
            Someone.address,
            // eslint-disable-next-line unicorn/numeric-separators-style
            61164767845294952445087173574656n,
          );

          await mintTo(
            Someone.address,
            // eslint-disable-next-line unicorn/numeric-separators-style
            162419591386637366713636064854016n,
          );

          await Game.connect(Someone).meldFlock([0, 1, 2, 3, 4]);
        });

        it('uneven gene can mutate');

        it('even gene can not mutate');

        it('can inherit from all parents');
      });

      describe('isCollectionMutating', () => {
        it('<= Rare can collection mutate');

        it('legendary can not mutate');
      });

      describe('requireGenomesSatisfyMelding', () => {
        it.only('success on melding Common Ducklings', async () => {
          await (async () => {
            for (let i = 0; i < 5; i++)
              await mintTo(
                Someone.address,
                randomGenome(Collections.Duckling, {
                  [DucklingGenes.Rarity]: Rarities.Common,
                  [DucklingGenes.Color]: 0,
                }),
              );
          })();

          try {
            await Duckies.connect(Someone).increaseAllowance(Game.address, 10_000_000_000);
            await GameAsSomeone.meldFlock([0, 1, 2, 3, 4]);
            assert(true);
          } catch {
            assert(false);
          }
        });

        it('success on melding Rare Ducklings');

        it('success on melding Epic Ducklings');

        it('success on melding Legendary Ducklings');

        it('success on melding Common Zombeak');

        it('success on melding Rare Zombeak');

        it('success on melding Epic Zombeak');

        it('revert on melding different collections');

        it('revert on melding different rarities');

        it('revert on melding melding Mythic');

        it('revert on melding melding legendary Zombeak');

        it('revert on melding legendaries having different color');

        it('revert on melding legendaries having repeated families');

        it('revert on melding not legendary having different color and different family');
      });

      describe('meldFlock', () => {
        describe('Duckling', () => {
          it('can meld into rare');

          it('can meld into epic');

          it('can meld into legendary');
        });

        describe('Duckling', () => {
          it('can meld into rare');

          it('can meld into epic');

          it('can meld into legendary');
        });

        describe('Mythic', () => {
          it('can meld into Mythic');

          it('Mythic id increments');

          it('revert on all Mythic minted');
        });

        it('event is emitted');
      });
    });
  });
});
