# Deploy Clearsync contracts

This script performs several operations:

- deploy `YellowAdjudicator`, `ClearingApp`, `EscrowApp` and write their addresses
- deploy specified ERC20 smart contracts and write their addresses
- mint some amount of tokens to specified addresses

## Configuration

The script expects a configuration file `clearsync/scripts/config.json`:

```ts
{
  "allocationAddresses": string[],
  "tokens": {
    [
      {
        "name": string,
        "symbol": string,
        "decimals": number,
      }
    ]
  }
}
```

`allocationAddresses` is a list of addresses to mint tokens to.

`tokens` is a list of ERC20 tokens to deploy and mint.

## Usage

```bash
DEPLOYER_PRIV_KEY="0x..." npx hardhat run scripts/deployClearsync.ts --network <network_name>
```

> `<network_name>` name must be specified in `hardhat.config.ts` file.

When running the script, you will see `deployerAddress` output to the console alongside with `chainId` the hardhat runner has connected to.

## Output

After deploying `YellowAdjudicator`, `ClearingApp` and `EscrowApp`, the script will save their addresses in a json file of a format:

```ts
{
  adjudicator: string;
  clearingApp: string;
  escrowApp: string;
}
```

After deploying and minting tokens, the script will save information about them in a `Uniswap token list` format:

```ts
interface Token {
  chainId: number;
  address: string;
  name: string;
  symbol: string;
  decimals: number;
}

interface TokenList {
  name: string;
  timestamp: string;
  tokens: Token[];
}
```
