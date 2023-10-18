# Deploy Clearsync contracts

This script performs several operations:

- deploy `YellowAdjudicator`, `ClearingApp`, `EscrowApp` and write their addresses
- deploy specified ERC20 smart contracts and write their addresses
- mint some amount of tokens to specified addresses

## Configuration

The scripts expects two environment variables:

- `MNEMONIC` - mnemonic phrase of the deployer account
- `STAGE` - stage of the deployment, one of `testnet`, `canarynet`, `mainnet`

> Note: in `mainnet` stage only the `YellowAdjudicator`, `ClearingApp`, `EscrowApp` are deployed, and not the tokens.

The script also expects a configuration file `clearsync/config/<stage>.config.json`:

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
DEPLOYER_PRIV_KEY="0x..." STAGE=<stage> npx hardhat run scripts/deployClearsync.ts --network <network_name>
```

> `<network_name>` name must be specified in `hardhat.config.ts` file.

When running the script, you will see `deployerAddress` output to the console alongside with `chainId` the hardhat runner has connected to.

## Output

The result of deploying contracts and tokens is saved to `config/<stage>.info.json` file, and is in a format:

```ts
interface Info {
  deployedContracts: DeployedContracts;
  tokenList: TokenList;
}

interface DeployedContracts {
  adjudicator: string;
  clearingApp: string;
  escrowApp: string;
}

interface TokenList {
  name: string;
  timestamp: string;
  tokens: Token[];
}

interface Token {
  chainId: number;
  address: string;
  name: string;
  symbol: string;
  decimals: number;
}
```
