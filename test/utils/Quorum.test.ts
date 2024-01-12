import { ethers } from 'hardhat';
import { expect } from 'chai';

import { connectGroup } from '../helpers/connect';

import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { Quorum, Quorum__factory } from '../../typechain-types';

const ZERO_QUORUM = 'Zero quorum';
const QUORUM_TOO_LARGE = 'Quorum too large';
const INVALID_ADDRESS = 'Invalid address';
const DUPLICATE_VALIDATORS = 'Duplicate validators';

describe('Quorum', () => {
  const quorum = 5;
  let Validators: SignerWithAddress[];
  let validators: string[];

  let ValidatorA: SignerWithAddress;

  let QuorumFactory: Quorum__factory;
  let Quorum: Quorum;
  let QuorumAsValidatorA: Quorum;

  beforeEach(async () => {
    [...Validators] = await ethers.getSigners();
    validators = Validators.map((V) => V.address);
    ValidatorA = Validators[0];

    QuorumFactory = (await ethers.getContractFactory('Quorum')) as Quorum__factory;
    Quorum = await QuorumFactory.deploy(validators, quorum);
    await Quorum.deployed();

    [QuorumAsValidatorA] = connectGroup(Quorum, [ValidatorA]);
  });

  describe('deployment', () => {
    it('correct validators and quorum', async () => {
      expect(await QuorumAsValidatorA.getValidators()).to.deep.equal(validators);
      expect(await QuorumAsValidatorA.getQuorum()).to.equal(quorum);
    });

    it('revert when quorum = 0', async () => {
      await expect(QuorumFactory.deploy(validators, 0)).to.be.revertedWith(ZERO_QUORUM);
    });

    it('revert when quorum > |validators|', async () => {
      await expect(QuorumFactory.deploy(validators, validators.length)).to.be.revertedWith(
        QUORUM_TOO_LARGE,
      );
    });

    it('revert when validators include address(0)', async () => {
      validators[3] = ethers.constants.AddressZero;
      await expect(QuorumFactory.deploy(validators, quorum)).to.be.revertedWith(INVALID_ADDRESS);
    });

    it('revert when validators have duplicates', async () => {
      validators[4] = validators[3];
      await expect(QuorumFactory.deploy(validators, quorum)).to.be.revertedWith(
        DUPLICATE_VALIDATORS,
      );
    });
  });

  describe('addValidator', () => {
    //
  });

  describe('removeValidator', () => {
    //
  });

  describe('setQuorum', () => {
    //
  });

  describe('setQuorumConfiguration', () => {
    //
  });

  describe('requireQuorum', () => {
    //
  });
});
