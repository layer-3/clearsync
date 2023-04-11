import { expect } from 'chai';
import { constants } from 'ethers';

import { ACCOUNT_MISSING_ROLE, randomBytes32 } from '../../../helpers/common';
import { Voucher, signVoucher, signVouchers } from '../../../helpers/voucher';

import { ADMIN_ROLE, setup } from './setup';
import {
  MeldParams,
  MintParams,
  VoucherAction,
  encodeMeldParams,
  encodeMintParams,
} from './voucher';
import { Collections, DucklingGenes, FLOCK_SIZE, MAX_PACK_SIZE, Rarities } from './config';
import {
  GenerateAndMintGenomesFunctT,
  MintToFuncT,
  setupGenerateAndMintGenomes,
  setupMintTo,
} from './helpers';

import type { DucklingsV1, TESTDuckyFamilyV1, YellowToken } from '../../../../typechain-types';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';

describe('DuckyFamilyV1 vouchers', () => {
  let Signer: SignerWithAddress;
  let Someone: SignerWithAddress;
  let Someother: SignerWithAddress;
  let GenomeSetter: SignerWithAddress;

  let Duckies: YellowToken;
  let Ducklings: DucklingsV1;
  let Game: TESTDuckyFamilyV1;
  let GameAsSomeone: TESTDuckyFamilyV1;

  let mintTo: MintToFuncT;
  let generateAndMintGenomes: GenerateAndMintGenomesFunctT;

  const VoucherBase: Voucher = {
    target: '',
    action: 0,
    beneficiary: '',
    referrer: constants.AddressZero,
    expire: 0,
    chainId: 31_337,
    voucherCodeHash: '',
    encodedParams: '0x',
  };

  beforeEach(async () => {
    ({ Signer, Someone, Someother, GenomeSetter, Duckies, Ducklings, Game, GameAsSomeone } =
      await setup());

    await Game.setIssuer(Signer.address);

    mintTo = setupMintTo(Ducklings.connect(GenomeSetter));
    generateAndMintGenomes = setupGenerateAndMintGenomes(mintTo, Someone.address);

    VoucherBase.voucherCodeHash = randomBytes32();
    VoucherBase.target = Game.address;
    VoucherBase.expire = Math.round(Date.now() / 1000) + 600; // 10 mins from now
  });

  describe('issuer', () => {
    it('admin can set issuer', async () => {
      await Game.setIssuer(Someone.address);
      expect(await Game.issuer()).to.equal(Someone.address);
    });

    it('revert on not admin set issuer', async () => {
      await expect(GameAsSomeone.setIssuer(Someone.address)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE),
      );
    });
  });

  describe('useVoucher', () => {
    describe('general revert', () => {
      let SomeoneVoucher: Voucher;
      let voucherSig: string;

      beforeEach(async () => {
        SomeoneVoucher = {
          ...VoucherBase,
          beneficiary: Someone.address,
        };

        voucherSig = await signVoucher(SomeoneVoucher, Signer);
      });

      it('revert on incorrect signer', async () => {
        const someoneVoucherSig = await signVoucher(SomeoneVoucher, Someone);
        await expect(GameAsSomeone.useVoucher(SomeoneVoucher, someoneVoucherSig))
          .to.be.revertedWithCustomError(Game, 'IncorrectSigner')
          .withArgs(Signer.address, Someone.address);
      });

      it('revert on target address != contract address', async () => {
        SomeoneVoucher.target = Someone.address;
        voucherSig = await signVoucher(SomeoneVoucher, Signer);
        await expect(
          GameAsSomeone.useVoucher(SomeoneVoucher, voucherSig),
        ).to.be.revertedWithCustomError(Game, 'InvalidVoucher');
        // .withArgs(SomeoneVoucher); <- seems like ethers can not parse JS Voucher properly
      });

      it('revert on beneficiary != sender', async () => {
        SomeoneVoucher.beneficiary = Someother.address;
        voucherSig = await signVoucher(SomeoneVoucher, Signer);
        await expect(
          GameAsSomeone.useVoucher(SomeoneVoucher, voucherSig),
        ).to.be.revertedWithCustomError(Game, 'InvalidVoucher');
      });

      it('revert on expired voucher', async () => {
        SomeoneVoucher.expire = Math.round(Date.now() / 1000) - 600; // 10 mins ago
        voucherSig = await signVoucher(SomeoneVoucher, Signer);
        await expect(
          GameAsSomeone.useVoucher(SomeoneVoucher, voucherSig),
        ).to.be.revertedWithCustomError(Game, 'InvalidVoucher');
      });

      it('revert on wrong chainId', async () => {
        SomeoneVoucher.chainId = 42;
        voucherSig = await signVoucher(SomeoneVoucher, Signer);
        await expect(
          GameAsSomeone.useVoucher(SomeoneVoucher, voucherSig),
        ).to.be.revertedWithCustomError(Game, 'InvalidVoucher');
      });

      it('revert on invalid voucher action', async () => {
        SomeoneVoucher.action = 42;
        voucherSig = await signVoucher(SomeoneVoucher, Signer);
        await expect(
          GameAsSomeone.useVoucher(SomeoneVoucher, voucherSig),
        ).to.be.revertedWithCustomError(Game, 'InvalidVoucher');
      });
    });

    describe('mint voucher', () => {
      let SomeoneMintVoucher: Voucher;
      let mintParams: MintParams;
      let voucherSig: string;

      beforeEach(async () => {
        mintParams = {
          to: Someone.address,
          size: 1,
          isTransferable: false,
        };

        SomeoneMintVoucher = {
          ...VoucherBase,
          action: VoucherAction.MintPack,
          beneficiary: Someone.address,
          encodedParams: encodeMintParams(mintParams),
        };

        voucherSig = await signVoucher(SomeoneMintVoucher, Signer);
      });

      // TODO: add transferability tests
      it('successfuly use mint voucher', async () => {
        await GameAsSomeone.useVoucher(SomeoneMintVoucher, voucherSig);
        expect(await Ducklings.balanceOf(Someone.address)).to.equal(1);
      });

      it('duckies are not paid for mint', async () => {
        await expect(
          GameAsSomeone.useVoucher(SomeoneMintVoucher, voucherSig),
        ).to.changeTokenBalance(Duckies, Someone, 0);
      });

      it('revert on using same voucher for second time', async () => {
        await GameAsSomeone.useVoucher(SomeoneMintVoucher, voucherSig);
        await expect(GameAsSomeone.useVoucher(SomeoneMintVoucher, voucherSig))
          .to.be.revertedWithCustomError(Game, 'VoucherAlreadyUsed')
          .withArgs(SomeoneMintVoucher.voucherCodeHash);
      });

      it('revert on to == address(0)', async () => {
        SomeoneMintVoucher.encodedParams = encodeMintParams({
          ...mintParams,
          to: constants.AddressZero,
        });
        voucherSig = await signVoucher(SomeoneMintVoucher, Signer);
        await expect(
          GameAsSomeone.useVoucher(SomeoneMintVoucher, voucherSig),
        ).to.be.revertedWithCustomError(Game, 'InvalidMintParams');
      });

      it('revert on size == 0', async () => {
        SomeoneMintVoucher.encodedParams = encodeMintParams({
          ...mintParams,
          size: 0,
        });
        voucherSig = await signVoucher(SomeoneMintVoucher, Signer);
        await expect(
          GameAsSomeone.useVoucher(SomeoneMintVoucher, voucherSig),
        ).to.be.revertedWithCustomError(Game, 'InvalidMintParams');
      });

      it('revert on size > MAX_PACK_SIZE', async () => {
        SomeoneMintVoucher.encodedParams = encodeMintParams({
          ...mintParams,
          size: MAX_PACK_SIZE + 1,
        });
        voucherSig = await signVoucher(SomeoneMintVoucher, Signer);
        await expect(
          GameAsSomeone.useVoucher(SomeoneMintVoucher, voucherSig),
        ).to.be.revertedWithCustomError(Game, 'InvalidMintParams');
      });

      it('event emitted', async () => {
        await expect(GameAsSomeone.useVoucher(SomeoneMintVoucher, voucherSig))
          .to.emit(Game, 'VoucherUsed')
          .withArgs(
            SomeoneMintVoucher.beneficiary,
            SomeoneMintVoucher.action,
            SomeoneMintVoucher.voucherCodeHash,
            SomeoneMintVoucher.chainId,
          );
      });
    });

    describe('meld voucher', () => {
      let SomeoneMeldVoucher: Voucher;
      let meldParams: MeldParams;
      let voucherSig: string;

      beforeEach(async () => {
        const { tokenIds } = await generateAndMintGenomes(Collections.Duckling, {
          amount: 5,
          [DucklingGenes.Rarity]: Rarities.Common,
          [DucklingGenes.Color]: 0,
        });

        meldParams = {
          owner: Someone.address,
          tokenIds,
          isTransferable: false,
        };

        SomeoneMeldVoucher = {
          ...VoucherBase,
          action: VoucherAction.MeldFlock,
          beneficiary: Someone.address,
          encodedParams: encodeMeldParams(meldParams),
        };

        voucherSig = await signVoucher(SomeoneMeldVoucher, Signer);
      });

      it('successfuly use meld voucher', async () => {
        const meldedTokenId = 5;
        await GameAsSomeone.useVoucher(SomeoneMeldVoucher, voucherSig);
        expect(await Ducklings.balanceOf(Someone.address)).to.equal(1);
        expect(await Ducklings.ownerOf(meldedTokenId)).to.equal(Someone.address);
      });

      it('duckies are not paid for meld', async () => {
        await expect(
          GameAsSomeone.useVoucher(SomeoneMeldVoucher, voucherSig),
        ).to.changeTokenBalance(Duckies, Someone, 0);
      });

      it('revert on using same voucher for second time', async () => {
        await GameAsSomeone.useVoucher(SomeoneMeldVoucher, voucherSig);
        await expect(GameAsSomeone.useVoucher(SomeoneMeldVoucher, voucherSig))
          .to.be.revertedWithCustomError(Game, 'VoucherAlreadyUsed')
          .withArgs(SomeoneMeldVoucher.voucherCodeHash);
      });

      it('revert on owner == address(0)', async () => {
        SomeoneMeldVoucher.encodedParams = encodeMeldParams({
          ...meldParams,
          owner: constants.AddressZero,
        });
        voucherSig = await signVoucher(SomeoneMeldVoucher, Signer);
        await expect(
          GameAsSomeone.useVoucher(SomeoneMeldVoucher, voucherSig),
        ).to.be.revertedWithCustomError(Game, 'InvalidMeldParams');
      });

      it('revert on number of tokens != FLOCK_SIZE', async () => {
        SomeoneMeldVoucher.encodedParams = encodeMeldParams({
          ...meldParams,
          tokenIds: meldParams.tokenIds.slice(0, FLOCK_SIZE - 1),
        });
        voucherSig = await signVoucher(SomeoneMeldVoucher, Signer);
        await expect(
          GameAsSomeone.useVoucher(SomeoneMeldVoucher, voucherSig),
        ).to.be.revertedWithCustomError(Game, 'InvalidMeldParams');
      });

      it('event is emitted', async () => {
        await expect(GameAsSomeone.useVoucher(SomeoneMeldVoucher, voucherSig))
          .to.emit(Game, 'VoucherUsed')
          .withArgs(
            SomeoneMeldVoucher.beneficiary,
            SomeoneMeldVoucher.action,
            SomeoneMeldVoucher.voucherCodeHash,
            SomeoneMeldVoucher.chainId,
          );
      });
    });
  });

  describe('useVouchers', () => {
    describe('mint', () => {
      let SomeoneMintVoucherA: Voucher;
      let SomeoneMintVoucherB: Voucher;
      let mintVouchersSig: string;

      const mintASize = 1;
      const mintBSize = 2;

      beforeEach(async () => {
        const mintParams = {
          to: Someone.address,
          size: mintASize,
          isTransferable: false,
        };

        SomeoneMintVoucherA = {
          ...VoucherBase,
          action: VoucherAction.MintPack,
          beneficiary: Someone.address,
          encodedParams: encodeMintParams(mintParams),
        };

        mintParams.size = mintBSize;

        SomeoneMintVoucherB = {
          ...VoucherBase,
          action: VoucherAction.MintPack,
          beneficiary: Someone.address,
          voucherCodeHash: randomBytes32(),
          encodedParams: encodeMintParams(mintParams),
        };

        mintVouchersSig = await signVouchers([SomeoneMintVoucherA, SomeoneMintVoucherB], Signer);
      });

      it('can use several mint vouchers', async () => {
        await GameAsSomeone.useVouchers(
          [SomeoneMintVoucherA, SomeoneMintVoucherB],
          mintVouchersSig,
        );
        expect(await Ducklings.balanceOf(Someone.address)).to.equal(mintASize + mintBSize);
      });

      it('revert on incorrect signer', async () => {
        const wrongSig = await signVouchers([SomeoneMintVoucherA, SomeoneMintVoucherB], Someone);
        await expect(
          GameAsSomeone.useVouchers([SomeoneMintVoucherA, SomeoneMintVoucherB], wrongSig),
        )
          .to.be.revertedWithCustomError(Game, 'IncorrectSigner')
          .withArgs(Signer.address, Someone.address);
      });
    });

    describe('meld', () => {
      let meldParams: MeldParams;
      let SomeoneMeldVoucherA: Voucher;
      let SomeoneMeldVoucherB: Voucher;
      let meldVouchersSig: string;

      beforeEach(async () => {
        const { tokenIds: tokenIdsA } = await generateAndMintGenomes(Collections.Duckling, {
          amount: 5,
          [DucklingGenes.Rarity]: Rarities.Common,
          [DucklingGenes.Color]: 0,
        });

        meldParams = {
          owner: Someone.address,
          tokenIds: tokenIdsA,
          isTransferable: false,
        };

        SomeoneMeldVoucherA = {
          ...VoucherBase,
          action: VoucherAction.MeldFlock,
          beneficiary: Someone.address,
          encodedParams: encodeMeldParams(meldParams),
        };

        const { tokenIds: tokenIdsB } = await generateAndMintGenomes(Collections.Duckling, {
          amount: 5,
          [DucklingGenes.Rarity]: Rarities.Rare,
          [DucklingGenes.Color]: 0,
          [DucklingGenes.Family]: 0,
        });

        meldParams.tokenIds = tokenIdsB;

        SomeoneMeldVoucherB = {
          ...VoucherBase,
          action: VoucherAction.MeldFlock,
          beneficiary: Someone.address,
          voucherCodeHash: randomBytes32(),
          encodedParams: encodeMeldParams(meldParams),
        };

        meldVouchersSig = await signVouchers([SomeoneMeldVoucherA, SomeoneMeldVoucherB], Signer);
      });

      it('can use several meld vouchers', async () => {
        await GameAsSomeone.useVouchers(
          [SomeoneMeldVoucherA, SomeoneMeldVoucherB],
          meldVouchersSig,
        );

        const vouchersAmount = 2;
        const meldedFromAId = 10;
        const meldedFromBId = 11;
        expect(await Ducklings.balanceOf(Someone.address)).to.equal(vouchersAmount);
        expect(await Ducklings.ownerOf(meldedFromAId)).to.equal(Someone.address);
        expect(await Ducklings.ownerOf(meldedFromBId)).to.equal(Someone.address);
      });

      it('revert on incorrect signer', async () => {
        const wrongSig = await signVouchers([SomeoneMeldVoucherA, SomeoneMeldVoucherB], Someone);
        await expect(
          GameAsSomeone.useVouchers([SomeoneMeldVoucherA, SomeoneMeldVoucherB], wrongSig),
        )
          .to.be.revertedWithCustomError(Game, 'IncorrectSigner')
          .withArgs(Signer.address, Someone.address);
      });
    });
  });
});
