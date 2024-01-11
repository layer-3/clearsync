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

const ERR_NOT_EXPECTED_HELD = 'held != expectedHeld';

let consensusApp: ConsensusApp;
let nitroAdjudicator: NitroAdjudicator;
let batchOperator: BatchOperator;
let token: Token;
let badToken: BadToken;

const ETH = MAGIC_ADDRESS_INDICATING_ETH;
let ERC20: string;
let BadERC20: string;
let participant: string;

const batchSize = 3;
const counterparties: string[] = [];
for (let i = 0; i < batchSize; i++) {
  counterparties[i] = Wallet.createRandom({
    extraEntropy: utils.id('multi-asset-holder-deposit-test'),
  }).address;
}

before(async () => {
  participant = await ethers.provider.getSigner(0).getAddress();

  consensusApp = await setupContract<ConsensusApp>('ConsensusApp');
  nitroAdjudicator = await setupContract<NitroAdjudicator>('NitroAdjudicator');
  batchOperator = await setupContract<BatchOperator>('BatchOperator', nitroAdjudicator.address);
  token = await setupContract<Token>('Token', participant);
  badToken = await setupContract<BadToken>('BadToken', participant);

  ERC20 = token.address;
  BadERC20 = badToken.address;
});

interface testParams {
  description: string;
  assetGetter: () => string;
  expectedHelds: number[];
  amounts: number[];
  heldAfters: number[];
  reasonString: string;
}

// eslint-disable-next-line sonarjs/cognitive-complexity
describe('deposit_batch', () => {
  // NOTE: getters allow to use values that are created in "beforeAll" hook inside test cases
  // This is needed because mocha executes the `describe` callback,
  // then hooks (including `beforeAll`) and only then test cases (`it` callbacks)
  const ETH_getter = (): string => ETH;
  const ERC20_getter = (): string => ERC20;
  const BadERC20_getter = (): string => BadERC20;

  const testCases = [
    {
      description: 'Deposits Eth to Multiple Channels (expectedHeld = 0)',
      assetGetter: ETH_getter,
      expectedHelds: [0, 0, 0],
      amounts: [1, 2, 3],
      heldAfters: [1, 2, 3],
      reasonString: '',
    },
    {
      description: 'Deposits Eth to Multiple Channels (expectedHeld = 1)',
      assetGetter: ETH_getter,
      expectedHelds: [1, 1, 1],
      amounts: [2, 2, 2],
      heldAfters: [3, 3, 3],
      reasonString: '',
    },
    {
      description: 'Deposits Eth to Multiple Channels (mixed expectedHeld)',
      assetGetter: ETH_getter,
      expectedHelds: [0, 1, 2],
      amounts: [1, 1, 1],
      heldAfters: [1, 2, 3],
      reasonString: '',
    },
    {
      description:
        'Reverts deposit of Eth to Multiple Channels (mismatched expectedHeld, zero expected)',
      assetGetter: ETH_getter,
      expectedHelds: [0, 0, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 1, 1],
      reasonString: ERR_NOT_EXPECTED_HELD,
    },
    {
      description:
        'Reverts deposit of Eth to Multiple Channels (mismatched expectedHeld, nonzero expected)',
      assetGetter: ETH_getter,
      expectedHelds: [1, 1, 1],
      amounts: [1, 1, 1],
      heldAfters: [2, 2, 2],
      reasonString: ERR_NOT_EXPECTED_HELD,
    },
    {
      description: 'Deposits Tokens to Multiple Channels (expectedHeld = 0)',
      assetGetter: ERC20_getter,
      expectedHelds: [0, 0, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 1, 1],
      reasonString: '',
    },
    {
      description: 'Deposits Tokens to Multiple Channels (expectedHeld = 1)',
      assetGetter: ERC20_getter,
      expectedHelds: [1, 1, 1],
      amounts: [1, 1, 1],
      heldAfters: [2, 2, 2],
      reasonString: '',
    },
    {
      description: 'Deposits Tokens to Multiple Channels (mixed expectedHeld)',
      assetGetter: ERC20_getter,
      expectedHelds: [0, 1, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 2, 1],
      reasonString: '',
    },
    {
      description:
        'Reverts deposit of Tokens to Multiple Channels (mismatched expectedHeld, zero expected)',
      assetGetter: ERC20_getter,
      expectedHelds: [0, 0, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 1, 1],
      reasonString: ERR_NOT_EXPECTED_HELD,
    },
    {
      description:
        'Reverts deposit of Tokens to Multiple Channels (mismatched expectedHeld, nonzero expected)',
      assetGetter: ERC20_getter,
      expectedHelds: [1, 1, 1],
      amounts: [1, 1, 1],
      heldAfters: [2, 2, 2],
      reasonString: ERR_NOT_EXPECTED_HELD,
    },
    {
      description: 'Deposits BadToken to Multiple Channels (expectedHeld = 0)',
      assetGetter: BadERC20_getter,
      expectedHelds: [0, 0, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 1, 1],
      reasonString: '',
    },
    {
      description: 'Reverts if input lengths do not match',
      assetGetter: ETH_getter,
      expectedHelds: [0, 0],
      amounts: [1, 1, 1],
      heldAfters: [1, 1, 1],
      reasonString: 'Array lengths must match',
    },
  ];

  for (const tc of testCases)
    it(tc.description, async () => {
      const { description, assetGetter, expectedHelds, amounts, heldAfters, reasonString } =
        tc as testParams;
      const asset: string = assetGetter();

      // Construct deposit_batch parameters
      const channelIds = counterparties.map((counterparty) =>
        getChannelId({
          channelNonce: getRandomNonce(description),
          participants: [participant, counterparty],
          appDefinition: consensusApp.address,
          challengeDuration: 100,
        }),
      );
      const expectedHeldsBN = expectedHelds.map((x) => BigNumber.from(x));
      const amountsBN = amounts.map((x) => BigNumber.from(x));
      const heldAftersBN = heldAfters.map((x) => BigNumber.from(x));
      const totalValue = BigNumber.from(amounts.reduce((s, n) => s + n));

      if (asset === ERC20) {
        const tx = await token.increaseAllowance(batchOperator.address, totalValue.add(totalValue));
        await tx.wait();
      }

      if (asset === BadERC20) {
        const tx = await badToken.increaseAllowance(
          batchOperator.address,
          totalValue.add(totalValue),
        );
        await tx.wait();
      }

      // Set up preexisting holdings (if any)
      await Promise.all(
        expectedHeldsBN.map(async (expected, i) => {
          const channelID = channelIds[i];
          // apply incorrect amount if unexpectedHeld reasonString is set
          const value = reasonString == ERR_NOT_EXPECTED_HELD ? expected.add(1) : expected;

          if (asset === ERC20) {
            const tx = await token.increaseAllowance(nitroAdjudicator.address, value);
            await tx.wait();
          }
          if (asset === BadERC20) {
            const tx = await badToken.increaseAllowance(nitroAdjudicator.address, value);
            await tx.wait();
          }

          const tx = await nitroAdjudicator.deposit(asset, channelID, 0, value, {
            value: asset === ETH ? value : 0,
          });
          const { events } = await tx.wait();
          expect(events).not.to.equal(undefined);
        }),
      );

      // Execute deposit
      const txPromise =
        asset === ETH
          ? batchOperator.deposit_batch_eth(channelIds, expectedHeldsBN, amountsBN, {
              value: totalValue,
            })
          : batchOperator.deposit_batch_erc20(
              asset,
              channelIds,
              expectedHeldsBN,
              amountsBN,
              totalValue,
            );

      // Check post-conditions
      if (reasonString == '') {
        const tx = await txPromise;
        await tx.wait();

        for (const [i, channelId] of channelIds.entries()) {
          const expectedHoldings = heldAftersBN[i];

          const holdings = await nitroAdjudicator.holdings(asset, channelId);
          expect(holdings).to.equal(expectedHoldings);
        }
      } else {
        await expectRevert(() => txPromise, reasonString);
      }
    });
});
