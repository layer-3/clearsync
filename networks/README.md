# Networks configs

This directory contains the network configurations for the different networks that are used by several components, including (the old) Terminal, Yellow.com (Yellow Vault) and the Pathfinder.

## Config structure

Config for each network resides in a separate directory, named after the network's id.

> Note: networks id reference: [chainlist.org](https://chainlist.org/).

Each network config contains:

- `assets.json` - [Yellow Vault] list of supported assets
- `mapping.json` - [Terminal] mapping for asset symbols that represent the same asset
- `markets.json` - [Terminal] list of supported markets
- `peers.json` - [Terminal] list of peers to connect to
- `wallet.json` [Yellow Vault, Pathfinder] wallet configuration

### assets.json

```ts
{
  "tokens": [
    {
        "address": Address,
        "name": string,
        "symbol": string,
        "decimals": int,
        "precision": int,
        "logoURI": string, // link to logo
        "extensions":
        {
          "allow_locking": boolean, // is asset lockable on Yellow Vault ?
          "coingecko_api_id": string,
          "locking_multiplier": float // locking leaderboard points multiplier for each $ of asset locked (1.5 means 1.5 points per $ locked)
        }
      },
  ]
}
```

### wallet.json

```ts
type UserOpType = "withdrawal", "swap", "lock", "unlock", "daily_claim", "daily_tap_reward", "mint", "other";
type UserOpFeeType = "native" | "erc20" | "sponsored";

type FeeTokenConfig = {
  paymasterAddress: Address;
  feeTokenBalanceSlot: number; // specific for each Token. See https://docs.soliditylang.org/en/v0.8.25/internals/layout_in_storage.html
  feeTokenAllowanceSlot: number; // specific for each Token. See https://docs.soliditylang.org/en/v0.8.25/internals/layout_in_storage.html
};

type WalletConfig = {
  erc20Paymasters: {
    [feeTokenAddress: Address]: FeeTokenConfig;
  };
  liteVaultAddress: Address;
  callTypes: {
    [key in UserOpType]: {
      allowedFeeTypes: UserOpFeeType[];
    };
  };
  trustedAddresses: Address[]; // if a trusted address is specified as `to` in a userOp, there is no call type rules check
};
```

## Adding a new network

When adding a new network, you need to:

1. Create a new directory with the network id
2. Add the required config files (at least `assets.json` and `wallet.json`)
3. Configure the network is supported by the Pathfinder (see its `Adding Newtork` docs)
4. Configure the network is supported by Yellow.com (Yellow Vault) (see its `Adding Newtork` docs)
5. Update the env on Yellow.com (Yellow Vault) that points to networks config (use a link that points to a commit, not branch)
