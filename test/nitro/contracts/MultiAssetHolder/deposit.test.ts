import { BigNumber, Event, Wallet, utils } from 'ethers';
import { ethers } from 'hardhat';
import { afterEach, before, describe, it } from 'mocha';
import { expect } from 'chai';

import { getChannelId } from '../../../../src/nitro/contract/channel';
import { setupContract } from '../../test-helpers';
import { MAGIC_ADDRESS_INDICATING_ETH } from '../../../../src/nitro/transactions';
import { getRandomNonce } from '../../../../src/nitro/helpers';

import type {
  BadToken,
  CountingApp,
  TESTNitroAdjudicator,
  Token,
} from '../../../../typechain-types';
import type { Result } from 'ethers/lib/utils';

const { AddressZero } = ethers.constants;

let countingApp: CountingApp;
let testNitroAdjudicator: TESTNitroAdjudicator;
let token: Token;
let badToken: BadToken;

const ETH = MAGIC_ADDRESS_INDICATING_ETH;
let ERC20: string;
let BadERC20: string;

let signer0Address: string;
const participants: string[] = [];
const challengeDuration = 0x10_00;

// Populate destinations array
for (let i = 0; i < 3; i++) {
  participants[i] = Wallet.createRandom({ extraEntropy: utils.id('erc20-deposit-test') }).address;
}

before(async () => {
  signer0Address = await ethers.provider.getSigner(0).getAddress();
  countingApp = await setupContract<CountingApp>('CountingApp');
  testNitroAdjudicator = await setupContract<TESTNitroAdjudicator>('TESTNitroAdjudicator');
  token = await setupContract<Token>('Token', signer0Address);
  badToken = await setupContract<BadToken>('BadToken', signer0Address);

  ERC20 = token.address;
  BadERC20 = badToken.address;
});

// eslint-disable-next-line sonarjs/cognitive-complexity
describe('deposit', () => {
  let channelNonce = getRandomNonce('deposit');

  // NOTE: getters allow to use values that are created in "beforeAll" hook inside test cases
  // This is needed because mocha executes the `describe` callback,
  // then hooks (including `beforeAll`) and only then test cases (`it` callbacks)
  const ETH_getter = (): string => ETH;
  const ERC20_getter = (): string => ERC20;
  const BadERC20_getter = (): string => BadERC20;

  const testCases = [
    {
      description: 'Deposits Tokens (expectedHeld = 0)',
      assetGetter: ERC20_getter,
      held: 0,
      expectedHeld: 0,
      amount: 1,
      heldAfter: 1,
      reasonString: undefined,
    },
    {
      description: 'Deposits Tokens (expectedHeld = 1)',
      assetGetter: ERC20_getter,
      held: 1,
      expectedHeld: 1,
      amount: 1,
      heldAfter: 2,
      reasonString: undefined,
    },
    {
      description: 'Reverts deposit of Tokens (expectedHeld > holdings)',
      assetGetter: ERC20_getter,
      held: 0,
      expectedHeld: 1,
      amount: 2,
      heldAfter: 0,
      reasonString: 'held != expectedHeld',
    },
    {
      description: 'Reverts deposit of Tokens (expectedHeld < holdings)',
      assetGetter: ERC20_getter,
      held: 1,
      expectedHeld: 0,
      amount: 2,
      heldAfter: 2,
      reasonString: 'held != expectedHeld',
    },
    {
      description: 'Deposits ETH (msg.value = amount , expectedHeld = 0)',
      assetGetter: ETH_getter,
      held: 0,
      expectedHeld: 0,
      amount: 1,
      heldAfter: 1,
      reasonString: undefined,
    },
    {
      description: 'Deposits ETH (msg.value = amount , expectedHeld = 1)',
      assetGetter: ETH_getter,
      held: 1,
      expectedHeld: 1,
      amount: 1,
      heldAfter: 2,
      reasonString: undefined,
    },
    {
      description: 'Reverts deposit of ETH (msg.value = amount, expectedHeld > holdings)',
      assetGetter: ETH_getter,
      held: 0,
      expectedHeld: 1,
      amount: 2,
      heldAfter: 0,
      reasonString: 'held != expectedHeld',
    },
    {
      description: 'Reverts deposit of ETH (msg.value = amount, expectedHeld < holdings)',
      assetGetter: ETH_getter,
      held: 1,
      expectedHeld: 0,
      amount: 2,
      heldAfter: 2,
      reasonString: 'held != expectedHeld',
    },
    {
      description: 'Reverts deposit of ETH (msg.value != amount)',
      assetGetter: ETH_getter,
      held: 0,
      expectedHeld: 0,
      amount: 1,
      heldAfter: 1,
      reasonString: 'Incorrect msg.value for deposit',
    },
    {
      description: 'Deposits a Bad token (expectedHeld = 0)',
      assetGetter: BadERC20_getter,
      held: 0,
      expectedHeld: 0,
      amount: 1,
      heldAfter: 1,
      reasonString: undefined,
    },
  ];

  afterEach(() => {
    channelNonce = BigNumber.from(channelNonce).add(1).toHexString();
  });

  for (const tc of testCases) {
    it(tc.description, async () => {
      const held = BigNumber.from(tc.held);
      const expectedHeld = BigNumber.from(tc.expectedHeld);
      const amount = BigNumber.from(tc.amount);
      const heldAfter = BigNumber.from(tc.heldAfter);
      const asset = tc.assetGetter();
      const reasonString = tc.reasonString;
      const destination = getChannelId({
        channelNonce,
        participants,
        appDefinition: countingApp.address,
        challengeDuration,
      });

      if (asset === ERC20) {
        // Check msg.sender has enough tokens
        const balance = await token.balanceOf(signer0Address);
        expect(balance.gte(held.add(amount))).to.be.true;

        // Increase allowance
        const tx = await token.approve(testNitroAdjudicator.address, held.add(amount));
        await tx.wait(); // Approve enough for setup and main test

        // Check allowance updated
        const allowance = BigNumber.from(
          await token.allowance(signer0Address, testNitroAdjudicator.address),
        );
        expect(allowance.sub(amount).sub(held).gte(0)).to.be.true;
      }

      if (asset === BadERC20) {
        // Check msg.sender has enough tokens
        const balance = await badToken.balanceOf(signer0Address);
        expect(balance.gte(held.add(amount))).to.be.true;

        // Increase allowance
        const tx = await badToken.approve(testNitroAdjudicator.address, held.add(amount));
        await tx.wait(); // Approve enough for setup and main test

        // Check allowance updated
        const allowance = BigNumber.from(
          await badToken.allowance(signer0Address, testNitroAdjudicator.address),
        );
        expect(allowance.sub(amount).sub(held).gte(0)).to.be.true;
      }

      if (held.gt(0)) {
        // Set holdings by depositing in the 'safest' way
        const tx = await testNitroAdjudicator.deposit(asset, destination, 0, held, {
          value: asset === ETH ? held : 0,
        });
        const { events } = await tx.wait();

        expect(events).to.not.be.undefined;
        if (events === undefined) {
          return;
        }

        expect(await testNitroAdjudicator.holdings(asset, destination)).to.equal(held);
        if (asset === ERC20 || asset == BadERC20) {
          const { data: amountTransferred } = getTransferEvent(events);
          expect(held.eq(amountTransferred)).to.be.true;
        }
      }

      const balanceBefore = await getBalance(asset, signer0Address);

      const pendingTx = testNitroAdjudicator.deposit(asset, destination, expectedHeld, amount, {
        value: asset === ETH && reasonString != 'Incorrect msg.value for deposit' ? amount : 0,
      });

      if (reasonString) {
        await expect(pendingTx).to.be.revertedWith(reasonString);
      } else {
        const tx = await pendingTx;
        const { events } = await tx.wait();
        expect(events).to.not.be.undefined;
        if (events === undefined) {
          return;
        }
        const depositedEvent = getDepositedEvent(events);
        const expectedEvent = {
          destination,
          destinationHoldings: heldAfter,
        };
        for (const [key, value] of Object.entries(expectedEvent)) {
          expect(depositedEvent[key as keyof Event]).to.deep.equal(value);
        }

        if (asset == ERC20 || asset == BadERC20) {
          const amountTransferred = BigNumber.from(getTransferEvent(events).data);
          expect(heldAfter.sub(held).eq(amountTransferred)).to.be.true;
          const balanceAfter = await getBalance(asset, signer0Address);
          expect(balanceAfter.eq(balanceBefore.sub(heldAfter.sub(held)))).to.be.true;
        }

        const allocatedAmount = await testNitroAdjudicator.holdings(asset, destination);
        expect(allocatedAmount).to.equal(heldAfter);
      }
    });
  }
});

const getDepositedEvent = (events: Event[]): Result => {
  const event = events.find(({ event }) => event === 'Deposited');
  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  return event!.args!;
};

const getTransferEvent = (events: Event[]): Event =>
  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  events.find(({ topics }) => topics[0] === token.filters.Transfer(AddressZero).topics![0])!;

async function getBalance(asset: string, address: string): Promise<BigNumber> {
  switch (asset) {
    case ETH: {
      return BigNumber.from(await ethers.provider.getBalance(address));
    }
    case ERC20: {
      return BigNumber.from(await token.balanceOf(address));
    }
    case BadERC20: {
      return BigNumber.from(await badToken.balanceOf(address));
    }
  }
  throw new Error('unrecognized asset');
}
