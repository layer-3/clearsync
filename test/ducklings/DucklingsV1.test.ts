import { ethers, upgrades } from 'hardhat';
import { ContractTransaction, constants, utils } from 'ethers';
import { assert, expect } from 'chai';

import { ACCOUNT_MISSING_ROLE } from '../helpers/common';
import { connectGroup } from '../helpers/connect';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { DucklingsV1, TESTDucklingsV2 } from '../../typechain-types';

async function expectTokenExists(Ducklings: DucklingsV1, tokenId: number): Promise<void> {
  expect(await Ducklings.ownerOf(tokenId)).not.to.equal(AddressZero);
}

async function expectTokenNotExists(Ducklings: DucklingsV1, tokenId: number): Promise<void> {
  await expect(Ducklings.ownerOf(tokenId)).to.be.revertedWith(INVALID_TOKEN_ID);
}

async function expectDucklingHasGenome(
  Ducklings: DucklingsV1,
  tokenId: number,
  genome: bigint,
): Promise<void> {
  const Duckling = await Ducklings.tokenToDuckling(tokenId);
  expect(Duckling.genome).to.equal(genome);
}

const AddressZero = constants.AddressZero;

const INVALID_TOKEN_ID = 'ERC721: invalid token ID';
const CUSTOM_INVALID_TOKEN_ID = 'InvalidTokenId';

const ADMIN_ROLE = constants.HashZero;
const UPGRADER_ROLE = utils.id('UPGRADER_ROLE');
const MAINTAINER_ROLE = utils.id('MAINTAINER_ROLE');
const GAME_ROLE = utils.id('GAME_ROLE');

const API_BASE_URL = 'test-url.com';

const GENOME = 42n;
const GENOME_2 = 422n;

describe('DucklingsV1', () => {
  let Admin: SignerWithAddress;
  let Upgrader: SignerWithAddress;
  let Maintainer: SignerWithAddress;
  let Someone: SignerWithAddress;
  let Someother: SignerWithAddress;
  let Game: SignerWithAddress;

  let Ducklings: DucklingsV1;
  let DucklingsAsSomeone: DucklingsV1;
  let DucklingsAsGame: DucklingsV1;

  const mintTo = async (
    to: string,
    genome: bigint,
    isTransferable?: boolean,
  ): Promise<ContractTransaction> => {
    return await Ducklings.connect(Game).mintTo(to, genome, isTransferable ?? true);
  };

  const mintBatchTo = async (
    to: string,
    genomes: bigint[],
    isTransferable?: boolean,
  ): Promise<ContractTransaction> => {
    return await Ducklings.connect(Game).mintBatchTo(to, genomes, isTransferable ?? true);
  };

  before(async () => {
    [Admin, Upgrader, Maintainer, Someone, Someother, Game] = await ethers.getSigners();
  });

  beforeEach(async () => {
    const DucklingsFactory = await ethers.getContractFactory('DucklingsV1');
    Ducklings = (await upgrades.deployProxy(DucklingsFactory, [], { kind: 'uups' })) as DucklingsV1;
    await Ducklings.deployed();

    await Ducklings.grantRole(UPGRADER_ROLE, Upgrader.address);
    await Ducklings.grantRole(MAINTAINER_ROLE, Maintainer.address);
    await Ducklings.grantRole(GAME_ROLE, Game.address);

    [DucklingsAsSomeone, DucklingsAsGame] = connectGroup(Ducklings, [Someone, Game]);
  });

  describe('deployment', () => {
    it('deployer has correct roles', async () => {
      expect(await Ducklings.hasRole(ADMIN_ROLE, Admin.address));
      expect(await Ducklings.hasRole(UPGRADER_ROLE, Admin.address));
      expect(await Ducklings.hasRole(MAINTAINER_ROLE, Admin.address));
    });

    it('deployer is Royalty collector', async () => {
      expect(await Ducklings.getRoyaltyCollector()).to.equal(Admin.address);
    });

    // TODO:
    it('has correct Royalty fee');

    it('NFT has correct name', async () => {
      expect(await Ducklings.name()).to.equal('Yellow Ducklings NFT Collection');
    });

    it('NFT has correct symbol', async () => {
      expect(await Ducklings.symbol()).to.equal('YDNC');
    });
  });

  describe('APIBaseURL', () => {
    it('Maintainer can set APIBaseURL', async () => {
      await Ducklings.connect(Maintainer).setAPIBaseURL(API_BASE_URL);
      expect(await Ducklings.apiBaseURL()).to.equal;
    });

    it('revert on not Maintainer setting APIBaseURL', async () => {
      await expect(DucklingsAsSomeone.setAPIBaseURL(API_BASE_URL)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, MAINTAINER_ROLE),
      );
    });
  });

  describe('IDucklings', () => {
    describe('is owner of', () => {
      it('return true for owner of 1 NFT', async () => {
        await mintTo(Someone.address, GENOME);
        expect(await Ducklings['isOwnerOf(address,uint256)'](Someone.address, 0)).to.be.true;
      });

      it('return false for not owner of 1 NFT', async () => {
        await mintTo(Someone.address, GENOME);
        expect(await Ducklings['isOwnerOf(address,uint256)'](Someother.address, 0)).to.be.false;
      });

      it('return true for owner of several NFTs', async () => {
        await mintTo(Someone.address, GENOME);
        await mintTo(Someone.address, GENOME_2);

        expect(await Ducklings['isOwnerOf(address,uint256[])'](Someone.address, [0, 1])).to.be.true;
      });

      it('return false for not owner of at least 1 of several NFTs', async () => {
        await mintTo(Someone.address, GENOME);
        await mintTo(Someother.address, GENOME_2);

        expect(await Ducklings['isOwnerOf(address,uint256[])'](Someone.address, [0, 1])).to.be
          .false;
      });

      it('return false for not owner of all NFTs', async () => {
        await mintTo(Someone.address, GENOME);
        await mintTo(Someone.address, GENOME_2);

        expect(await Ducklings['isOwnerOf(address,uint256[])'](Someother.address, [0, 1])).to.be
          .false;
      });
    });

    describe('get genome', () => {
      it('return correct genome given existing token id', async () => {
        await mintTo(Someone.address, GENOME);
        expect(await Ducklings.getGenome(0)).to.equal(GENOME);
      });

      it('revert when token does not exist', async () => {
        await expect(Ducklings.getGenome(0))
          .to.be.revertedWithCustomError(Ducklings, CUSTOM_INVALID_TOKEN_ID)
          .withArgs(0);
      });

      it('return correct genomes given array of token ids', async () => {
        await mintTo(Someone.address, GENOME);
        await mintTo(Someone.address, GENOME_2);
        expect(await Ducklings.getGenomes([0, 1])).to.deep.equal([GENOME, GENOME_2]);
      });

      it('revert when at least 1 token does not exist', async () => {
        await mintTo(Someone.address, GENOME);
        await expect(Ducklings.getGenomes([0, 1]))
          .to.be.revertedWithCustomError(Ducklings, CUSTOM_INVALID_TOKEN_ID)
          .withArgs(1);
      });
    });

    describe('transferable', () => {
      it('success on Game set transferable', async () => {
        await mintTo(Someone.address, GENOME);
        expect(await DucklingsAsGame.isTokenTransferable(0)).to.be.true;
      });

      it('revert on not Game set tranferable', async () => {
        await mintTo(Someone.address, GENOME);
        await expect(DucklingsAsSomeone.setTransferable(0, false)).to.be.revertedWith(
          ACCOUNT_MISSING_ROLE(Someone.address, GAME_ROLE),
        );
      });

      it('revert on set unexisting token tranferable', async () => {
        await expect(DucklingsAsGame.setTransferable(0, false))
          .to.be.revertedWithCustomError(Ducklings, 'InvalidTokenId')
          .withArgs(0);
      });

      it('revert on transfering untransferable token', async () => {
        await mintTo(Someone.address, GENOME);
        await DucklingsAsGame.setTransferable(0, false);

        await expect(DucklingsAsSomeone.transferFrom(Someone.address, Someother.address, 0))
          .to.be.revertedWithCustomError(Ducklings, 'TokenNotTransferable')
          .withArgs(0);
      });

      it('can burn not transferable token', async () => {
        await mintTo(Someone.address, GENOME);
        await DucklingsAsGame.setTransferable(0, false);

        try {
          await DucklingsAsGame.burn(0);
          assert(true);
        } catch {
          assert(false);
        }
      });

      it('emit event', async () => {
        await mintTo(Someone.address, GENOME);
        await expect(DucklingsAsGame.setTransferable(0, false))
          .to.emit(Ducklings, 'TransferableSet')
          .withArgs(0, false);
      });
    });

    describe('mintTo', () => {
      it('Game can mint', async () => {
        try {
          await mintTo(Someone.address, GENOME);
          assert(true);
        } catch {
          assert(false);
        }
      });

      it('revert on not Game minting', async () => {
        await expect(DucklingsAsSomeone.mintTo(Someone.address, GENOME, true)).to.be.rejectedWith(
          ACCOUNT_MISSING_ROLE(Someone.address, GAME_ROLE),
        );
      });

      it('genome is set to NFT', async () => {
        await mintTo(Someone.address, GENOME);
        await expectDucklingHasGenome(Ducklings, 0, GENOME);
      });

      it('NFT birthdate is block timestamp', async () => {
        await mintTo(Someone.address, GENOME);
        const Duckling = await Ducklings.tokenToDuckling(0);

        const latestBlock = await ethers.provider.getBlock('latest');

        expect(Duckling.birthdate).to.equal(latestBlock.timestamp);
      });

      it('NFT id is incremental', async () => {
        await mintTo(Someone.address, GENOME);
        await expectTokenExists(Ducklings, 0);
        await expectTokenNotExists(Ducklings, 1);

        await expectDucklingHasGenome(Ducklings, 0, GENOME);

        await mintTo(Someone.address, GENOME_2);
        await expectTokenExists(Ducklings, 0);
        await expectTokenExists(Ducklings, 1);
        await expectTokenNotExists(Ducklings, 2);

        await expectDucklingHasGenome(Ducklings, 1, GENOME_2);
      });

      it('Mint event is emitted', async () => {
        const { chainId } = await ethers.provider.getNetwork();
        const isTokenTransferable = true;
        const tx = await mintTo(Someone.address, GENOME, isTokenTransferable);
        const latestBlock = await ethers.provider.getBlock('latest');

        await expect(tx)
          .to.emit(Ducklings, 'Minted')
          .withArgs(
            Someone.address,
            0,
            isTokenTransferable,
            GENOME,
            latestBlock.timestamp,
            chainId,
          );
      });

      it('TransferableSet event is emitted', async () => {
        const isTokenTransferable = true;
        const tx = await mintTo(Someone.address, GENOME, isTokenTransferable);

        await expect(tx).to.emit(Ducklings, 'TransferableSet').withArgs(0, isTokenTransferable);
      });
    });

    describe('mintBatchTo', () => {
      it('Game can mint', async () => {
        try {
          await mintBatchTo(Someone.address, [GENOME, GENOME_2]);
          assert(true);
        } catch {
          assert(false);
        }
      });

      it('revert on not Game minting', async () => {
        await expect(
          DucklingsAsSomeone.mintBatchTo(Someone.address, [GENOME, GENOME_2], true),
        ).to.be.rejectedWith(ACCOUNT_MISSING_ROLE(Someone.address, GAME_ROLE));
      });

      it('genomes are set to NFTs', async () => {
        await mintBatchTo(Someone.address, [GENOME, GENOME_2]);
        await expectDucklingHasGenome(Ducklings, 0, GENOME);
        await expectDucklingHasGenome(Ducklings, 1, GENOME_2);
      });

      it('NFTs birthdate is block timestamp', async () => {
        await mintBatchTo(Someone.address, [GENOME, GENOME_2]);
        const Duckling1 = await Ducklings.tokenToDuckling(0);
        const Duckling2 = await Ducklings.tokenToDuckling(0);

        const latestBlock = await ethers.provider.getBlock('latest');

        expect(Duckling1.birthdate).to.equal(latestBlock.timestamp);
        expect(Duckling2.birthdate).to.equal(latestBlock.timestamp);
      });

      it('NFTs id are incremental', async () => {
        await mintBatchTo(Someone.address, [GENOME, GENOME_2]);
        await expectTokenExists(Ducklings, 0);
        await expectTokenExists(Ducklings, 1);
        await expectTokenNotExists(Ducklings, 2);

        await expectDucklingHasGenome(Ducklings, 0, GENOME);
        await expectDucklingHasGenome(Ducklings, 1, GENOME_2);

        // mint once more
        await mintTo(Someone.address, GENOME_2);
        await expectTokenExists(Ducklings, 0);
        await expectTokenExists(Ducklings, 1);
        await expectTokenExists(Ducklings, 2);
        await expectTokenNotExists(Ducklings, 3);

        await expectDucklingHasGenome(Ducklings, 2, GENOME_2);
      });

      it('Mint events are emitted', async () => {
        const { chainId } = await ethers.provider.getNetwork();
        const isTokenTransferable = true;
        const tx = await mintBatchTo(Someone.address, [GENOME, GENOME_2], isTokenTransferable);
        const latestBlock = await ethers.provider.getBlock('latest');

        await expect(tx)
          .to.emit(Ducklings, 'Minted')
          .withArgs(
            Someone.address,
            0,
            isTokenTransferable,
            GENOME,
            latestBlock.timestamp,
            chainId,
          );

        await expect(tx)
          .to.emit(Ducklings, 'Minted')
          .withArgs(
            Someone.address,
            1,
            isTokenTransferable,
            GENOME_2,
            latestBlock.timestamp,
            chainId,
          );
      });

      it('TransferableSet events are emitted', async () => {
        const isTokenTransferable = true;
        const tx = await mintBatchTo(Someone.address, [GENOME, GENOME_2], isTokenTransferable);

        await expect(tx).to.emit(Ducklings, 'TransferableSet').withArgs(0, isTokenTransferable);
        await expect(tx).to.emit(Ducklings, 'TransferableSet').withArgs(1, isTokenTransferable);
      });
    });

    describe('burn', () => {
      it('Game can burn 1 token', async () => {
        await mintTo(Someone.address, GENOME);
        await DucklingsAsGame.burn(0);
        await expectTokenNotExists(Ducklings, 0);
      });

      it('revert on not Game burning', async () => {
        await expect(DucklingsAsSomeone.burn(0)).to.be.revertedWith(
          ACCOUNT_MISSING_ROLE(Someone.address, GAME_ROLE),
        );
      });

      it('revert on Game burning unexisting token', async () => {
        await expect(DucklingsAsGame.burn(0)).to.be.revertedWith(INVALID_TOKEN_ID);
      });
    });

    describe('burnBatch', () => {
      it('Game can burn several tokens of the same owner', async () => {
        await mintBatchTo(Someone.address, [GENOME, GENOME_2]);
        await DucklingsAsGame.burnBatch([0, 1]);
        await expectTokenNotExists(Ducklings, 0);
        await expectTokenNotExists(Ducklings, 1);
      });

      it('revert on not Game burning', async () => {
        await expect(DucklingsAsSomeone.burnBatch([0, 1])).to.be.revertedWith(
          ACCOUNT_MISSING_ROLE(Someone.address, GAME_ROLE),
        );
      });

      it('Game can burn several tokens of the different owners', async () => {
        await mintTo(Someone.address, GENOME);
        await mintTo(Someother.address, GENOME_2);
        await DucklingsAsGame.burnBatch([0, 1]);
        await expectTokenNotExists(Ducklings, 0);
        await expectTokenNotExists(Ducklings, 1);
      });

      it('revert on Game burning tokens with one unexisting', async () => {
        await mintTo(Someone.address, GENOME);
        await expect(DucklingsAsGame.burnBatch([0, 1])).to.be.revertedWith(INVALID_TOKEN_ID);
      });

      it('revert on Game burning unexisting tokens', async () => {
        await expect(DucklingsAsGame.burnBatch([0, 1])).to.be.revertedWith(INVALID_TOKEN_ID);
      });
    });
  });

  describe('ERC721', () => {
    it('return correct tokenURI', async () => {
      await Ducklings.setAPIBaseURL(API_BASE_URL);
      await mintTo(Someone.address, GENOME);

      const latestBlock = await ethers.provider.getBlock('latest');

      expect(await Ducklings.tokenURI(0)).to.equal(
        `${API_BASE_URL}${GENOME}-${latestBlock.timestamp}`,
      );
    });
  });

  describe('ERC2981 Royalties', () => {
    it('Admin can set Royalty collector', async () => {
      await Ducklings.setRoyaltyCollector(Someone.address);
      expect(await Ducklings.getRoyaltyCollector()).to.equal(Someone.address);
    });

    it('revert on not Admin setting Royalty collector', async () => {
      await expect(DucklingsAsSomeone.setRoyaltyCollector(Someother.address)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, ADMIN_ROLE),
      );
    });
  });

  describe('upgrade', () => {
    let DucklingsV2: TESTDucklingsV2;

    beforeEach(async () => {
      const TESTDucklingsV2Factory = await ethers.getContractFactory('TESTDucklingsV2');
      DucklingsV2 = (await TESTDucklingsV2Factory.deploy()) as TESTDucklingsV2;
      await DucklingsV2.deployed();
    });

    it('Upgrader can upgrade', async () => {
      await Ducklings.connect(Upgrader).upgradeTo(DucklingsV2.address);
    });

    it('revert on not Upgrader upgrading', async () => {
      await expect(DucklingsAsSomeone.upgradeTo(DucklingsV2.address)).to.be.revertedWith(
        ACCOUNT_MISSING_ROLE(Someone.address, UPGRADER_ROLE),
      );
    });
  });
});