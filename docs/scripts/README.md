# Scripts documentation

Every command is invoked from the root of this repo, unless specified otherwise.

## Set .env file

To be able to use scripts, you need to set `.env` file which must contain RPC endpoints and blockscan API keys for each network you want to use.

E.g. for Polygon, you need to set:

```bash
POLYGON_URL=https://polygon-mainnet.g.alchemy.com/v2/YOUR_API_KEY
POLYGONSCAN_API_KEY=YOUR_API_KEY
```

Also, you can put your private key, so that you don't have to prepend it before every command:

```bash
PRIVATE_KEY=YOUR_PRIVATE_KEY
```

Otherwise, to run scripts, you must prepend your private key to every command (except for verifications):

```bash
PRIVATE_KEY=YOUR_PRIVATE_KEY ...command
```

## Deploy any contract

You can deploy any contract with source code in this repo with the following command:

```bash
NAME=<contract_name> ARGS=<constructor_args_separated_by_comma> npx hardhat run scripts/deployContract.ts --network <network_name>
```

Note, that to use a network with `network_name`, it should be correctly set up in the `hardhat.config.ts` file.

## Verify any contract

You can verify any contract with source code in this repo with the following command:

```bash
npx hardhat verify <contract_address> --contract '<contract_path>:<contract_name>' <constructor_args_separated_by_space> --network <network_name>
```

If the contract being verified is a proxy, to be able to interact with it using blockscans, you need to verify it with implementation contract source code.
