import { BigNumber, Wallet, utils } from 'ethers';
import { ethers } from 'hardhat';
import { before, describe, it } from 'mocha';
import { expect } from 'chai';

import { expectRevert } from '../../../helpers/expect-revert';
import { MAGIC_ADDRESS_INDICATING_ETH, getChannelId, getRandomNonce } from '../../../../src/nitro';
import { setupContract } from '../../test-helpers';

import type {
  BadToken,
  BatchOperator,
  ConsensusApp,
  NitroAdjudicator,
  Token,
} from '../../../../typechain-types';

let consensusApp: ConsensusApp;
let nitroAdjudicator: NitroAdjudicator;
let batchOperator: BatchOperator;
let token: Token;
let badToken: BadToken;

const ETH = MAGIC_ADDRESS_INDICATING_ETH;
let ERC20: string; // = token.address;
let BadERC20: string; // = badToken.address;
let signerAddress: string;

const batchSize = 3;
const counterparties: string[] = [];
for (let i = 0; i < batchSize; i++) {
  counterparties[i] = Wallet.createRandom({
    extraEntropy: utils.id('multi-asset-holder-deposit-test'),
  }).address;
}

before(async () => {
  signerAddress = await ethers.provider.getSigner(0).getAddress();

  consensusApp = await setupContract<ConsensusApp>('ConsensusApp');
  nitroAdjudicator = await setupContract<NitroAdjudicator>('NitroAdjudicator');
  batchOperator = await setupContract<BatchOperator>('BatchOperator', nitroAdjudicator.address);
  token = await setupContract<Token>('Token', '0x6B8B2958795a5E9c00A2E8D4B0b90b870cbAB4b6');
  badToken = await setupContract<BadToken>(
    'BadToken',
    '0x6B8B2958795a5E9c00A2E8D4B0b90b870cbAB4b6',
  );
});

interface testParams {
  description: string;
  assetId: string;
  expectedHelds: number[];
  amounts: number[];
  heldAfters: number[];
  reasonString: string;
}

function sum(x: BigNumber[]): BigNumber {
  return x.reduce((s, n) => s.add(n));
}

describe('deposit_batch', () => {
  const testCases = [
    {
      description: 'Deposits Eth to Multiple Channels (expectedHeld = 0)',
      assetId: ETH,
      expectedHelds: [0, 0, 0],
      amounts: [1, 2, 3],
      heldAfters: [1, 2, 3],
      reasonString: '',
    },
    {
      description: 'Deposits Eth to Multiple Channels (expectedHeld = 1)',
      assetId: ETH,
      expectedHelds: [1, 1, 1],
      amounts: [2, 2, 2],
      heldAfters: [3, 3, 3],
      reasonString: '',
    },
    {
      description: 'Deposits Eth to Multiple Channels (mixed expectedHeld)',
      assetId: ETH,
      expectedHelds: [0, 1, 2],
      amounts: [1, 1, 1],
      heldAfters: [1, 2, 3],
      reasonString: '',
    },
    {
      description:
        'Reverts deposit of Eth to Multiple Channels (mismatched expectedHeld, zero expected)',
      assetId: ETH,
      expectedHelds: [0, 0, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 1, 1],
      reasonString: 'unexpectedHeld',
    },
    {
      description:
        'Reverts deposit of Eth to Multiple Channels (mismatched expectedHeld, nonzero expected)',
      assetId: ETH,
      expectedHelds: [1, 1, 1],
      amounts: [1, 1, 1],
      heldAfters: [2, 2, 2],
      reasonString: 'unexpectedHeld',
    },
    {
      description: 'Deposits Tokens to Multiple Channels (expectedHeld = 0)',
      assetId: ERC20,
      expectedHelds: [0, 0, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 1, 1],
      reasonString: '',
    },
    {
      description: 'Deposits Tokens to Multiple Channels (expectedHeld = 1)',
      assetId: ERC20,
      expectedHelds: [1, 1, 1],
      amounts: [1, 1, 1],
      heldAfters: [2, 2, 2],
      reasonString: '',
    },
    {
      description: 'Deposits Tokens to Multiple Channels (mixed expectedHeld)',
      assetId: ERC20,
      expectedHelds: [0, 1, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 2, 1],
      reasonString: '',
    },
    {
      description:
        'Reverts deposit of Tokens to Multiple Channels (mismatched expectedHeld, zero expected)',
      assetId: ERC20,
      expectedHelds: [0, 0, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 1, 1],
      reasonString: 'unexpectedHeld',
    },
    {
      description:
        'Reverts deposit of Tokens to Multiple Channels (mismatched expectedHeld, nonzero expected)',
      assetId: ERC20,
      expectedHelds: [1, 1, 1],
      amounts: [1, 1, 1],
      heldAfters: [2, 2, 2],
      reasonString: 'unexpectedHeld',
    },
    {
      description: 'Deposits BadToken to Multiple Channels (expectedHeld = 0)',
      assetId: BadERC20,
      expectedHelds: [0, 0, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 1, 1],
      reasonString: '',
    },
    {
      description: 'Reverts if input lengths do not match',
      assetId: ETH,
      expectedHelds: [0, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 1, 1],
      reasonString: 'Array lengths must match',
    },
  ];

  for (const tc of testCases) it(tc.description, async () => {
      const { description, assetId, expectedHelds, amounts, heldAfters, reasonString } =
        tc as testParams;
      ///////////////////////////////////////
      //
      // Construct deposit_batch parameters
      //
      ///////////////////////////////////////
      const channelIds = counterparties.map((counterparty) =>
        getChannelId({
          channelNonce: getRandomNonce(description),
          participants: [signerAddress, counterparty],
          appDefinition: consensusApp.address,
          challengeDuration: 100,
        }),
      );
      const expectedHeldsBN = expectedHelds.map((x) => BigNumber.from(x));
      const amountsBN = amounts.map((x) => BigNumber.from(x));
      const heldAftersBN = heldAfters.map((x) => BigNumber.from(x));
      const totalValue = sum(amountsBN);

      if (assetId === ERC20) {
        await (
          await token.increaseAllowance(batchOperator.address, totalValue.add(totalValue))
        ).wait();
      }

      if (assetId === BadERC20) {
        await (
          await badToken.increaseAllowance(batchOperator.address, totalValue.add(totalValue))
        ).wait();
      }

      ///////////////////////////////////////
      //
      // Set up preexisting holdings (if any)
      //
      ///////////////////////////////////////

      await Promise.all(
        expectedHeldsBN.map(async (expected, i) => {
          const channelID = channelIds[i];
          // apply incorrect amount if unexpectedHeld reasonString is set
          const value = reasonString == 'held != expectedHeld' ? expected.add(1) : expected;

          if (assetId === ERC20) {
            await (await token.increaseAllowance(nitroAdjudicator.address, value)).wait();
          }
          if (assetId === BadERC20) {
            await (await badToken.increaseAllowance(nitroAdjudicator.address, value)).wait();
          }

          const { events } = await (
            await nitroAdjudicator.deposit(assetId, channelID, 0, value, {
              value: assetId === ETH ? value : 0,
            })
          ).wait();
          expect(events).not.to.equal(undefined);
        }),
      );

      ///////////////////////////////////////
      //
      // Execute deposit
      //
      ///////////////////////////////////////

      const tx =
        assetId === ETH
          ? batchOperator.deposit_batch_eth(channelIds, expectedHeldsBN, amountsBN, {
              value: totalValue,
            })
          : batchOperator.deposit_batch_erc20(
              assetId,
              channelIds,
              expectedHeldsBN,
              amountsBN,
              totalValue,
            );

      ///////////////////////////////////////
      //
      // Check postconditions
      //
      ///////////////////////////////////////
      if (reasonString == '') {
        await (await tx).wait();

        for (const [i, channelId] of channelIds.entries()) {
          const expectedHoldings = heldAftersBN[i];

          const holdings = await nitroAdjudicator.holdings(assetId, channelId);
          expect(holdings).to.equal(expectedHoldings);
        }
      } else {
        await expectRevert(() => tx, reasonString);
      }
    })
  ;
});
