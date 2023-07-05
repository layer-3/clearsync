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

## Deploy proxy contract

You can deploy a proxy contract with the following command:

```bash
NAME=<contract_name> ARGS=<constructor_args_separated_by_comma> npx hardhat run scripts/deployProxy.ts --network <network_name>
```

## Upgrade proxy contract

You can upgrade a proxy contract with the following command:

```bash
UPGRADE_TO_NAME=<contract_name_to_upgrade_to> ADDRESS=<proxy_address> npx hardhat run scripts/upgradeProxy.ts --network <network_name>
```

## Verify any contract

You can verify any contract with source code in this repo with the following command:

```bash
npx hardhat verify <contract_address> --contract '<contract_path>:<contract_name>' <constructor_args_separated_by_space> --network <network_name>
```

If the contract being verified is a proxy, to be able to interact with it using blockscans, you need to verify it with implementation contract source code.

## Duckies and Ducklings

### Deploy and verify Ducklings

Deploy:

```bash
NAME=DucklingsV1 npx hardhat run scripts/deployProxy.ts --network polygon
```

In the console, you will see:

```sh
Deploying contracts with the account: YOUR_ADDRESS
Account balance: YOUR_BALANCE
DucklingsV1 deployed to DUCKLINGSV1_ADDRESS with args ''
```

Verify:

```bash
npx hardhat verify DUCKLINGSV1_ADDRESS --contract 'contracts/duckies/ducklings/DucklingsV1.sol:DucklingsV1' --network polygon
```

This command should verify both the Proxy and the Implementation contract.

### Deploy and verify DuckyFamily

Deploy:

```bash
NAME=DuckyFamilyV1 ARGS=DUCKIES_ADDRESS,DUCKLINGSV1_ADDRESS,TREASURE_VAULT_ADDRESS npx hardhat run scripts/deployContract.ts --network polygon
```

Which will result in DUCKY_FAMILYV1_ADDRESS.

Verify:

```bash
npx hardhat verify DUCKY_FAMILYV1_ADDRESS --contract 'contracts/duckies/games/DuckyFamily/DuckyFamilyV1.sol:DuckyFamilyV1' DUCKIES_ADDRESS DUCKLINGSV1_ADDRESS TREASURE_VAULT_ADDRESS --network polygon
```

### Setup and connect DuckyFamily to Ducklings

```bash
npx hardhat setupNFTs  --ducklings DUCKLINGSV1_ADDRESS  --ducky-family DUCKY_FAMILYV1_ADDRESS --api-base-url 'https://www.yellow.org/api/v3/public/nft/metadata/' --issuer 0xC5F825188ad49b13a8c5116FfDab7121b1CEf595 --network polygon`
```
