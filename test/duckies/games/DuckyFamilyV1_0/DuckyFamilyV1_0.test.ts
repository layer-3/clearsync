import { utils } from 'ethers';
import { ethers, upgrades } from 'hardhat';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type {
  Ducklings,
  DuckyFamilyV1_0,
  TreasureVault,
  YellowToken,
} from '../../../../typechain-types';

const GAME_ROLE = utils.id('GAME_ROLE');

describe('DuckyFamilyV1_0', () => {
  let Admin: SignerWithAddress;
  let Maintainer: SignerWithAddress;
  let Someone: SignerWithAddress;
  let Someother: SignerWithAddress;

  let Duckies: YellowToken;
  let Ducklings: Ducklings;
  let TreasureVault: TreasureVault;
  let DuckyFamily: DuckyFamilyV1_0;

  before(async () => {
    [Admin, Maintainer, Someone, Someother] = await ethers.getSigners();

    const DuckiesFactory = await ethers.getContractFactory('YellowToken');
    Duckies = (await DuckiesFactory.deploy('Duckies', 'DUCKIES', 1_000_000 * 10e8)) as YellowToken;
    await Duckies.deployed();

    const DucklingsFactory = await ethers.getContractFactory('Ducklings');
    Ducklings = (await upgrades.deployProxy(DucklingsFactory, [], { kind: 'uups' })) as Ducklings;
    await Ducklings.deployed();

    const TreasureVaultFactory = await ethers.getContractFactory('TreasureVault');
    TreasureVault = (await upgrades.deployProxy(TreasureVaultFactory, [], {
      kind: 'uups',
    })) as TreasureVault;
    await TreasureVault.deployed();

    const DuckyFamilyFactory = await ethers.getContractFactory('DuckyFamilyV1_0');
    DuckyFamily = (await upgrades.deployProxy(DuckyFamilyFactory, [
      Duckies.address,
      Ducklings.address,
      TreasureVault.address,
    ])) as DuckyFamilyV1_0;
    await DuckyFamily.deployed();

    await Ducklings.grantRole(GAME_ROLE, DuckyFamily.address);
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
      it('maintainer can set mint price');

      it('revert on not maintainer set mint price');

      it('maintainer can set meld price');

      it('revert on not maintainer set meld price');
    });

    describe('helpers', () => {
      it('get correct distribution type');
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
        it('uneven gene can mutate');

        it('even gene can not mutate');

        it('can inherit from all parents');
      });

      describe('isCollectionMutating', () => {
        it('<= Rare can collection mutate');

        it('legendary can not mutate');
      });

      describe('revert on melding rules not satisfied', () => {
        it('on different collections');

        it('on different rarities');

        it('on melding Mythic');

        it('on melding legendary Zombeak');

        it('on legendaries having different color');

        it('on legendaries having repeated families');

        it('on not legendary having different color and different family');
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
