# How to deploy Ducklings, DuckyFamily

Every command is invoked from the root of this repo, unless specified otherwise.

## Set .env file

To be able to use scripts, you need to set `.env` file which must contain Polygon RPC endpoint and Polygonscan API key:

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

## Deploy and verify Ducklings

Deploy:
`NAME=DucklingsV1 npx hardhat run scripts/deployProxy.ts --network polygon`

In the console, you will see:

```sh
Deploying contracts with the account: YOUR_ADDRESS
Account balance: YOUR_BALANCE
DucklingsV1 deployed to DUCKLINGSV1_ADDRESS with args ''
```

Verify:
`npx hardhat verify DUCKLINGSV1_ADDRESS --contract 'contracts/duckies/ducklings/DucklingsV1.sol:DucklingsV1' --network polygon`

This command should verify both the Proxy and the Implementation contract.

## Deploy and verify DuckyFamily

Deploy:
`NAME=DuckyFamilyV1 ARGS=DUCKIES_ADDRESS,DUCKLINGSV1_ADDRESS,TREASURE_VAULT_ADDRESS npx hardhat run scripts/deployContract.ts --network polygon`

Which will result in DUCKY_FAMILYV1_ADDRESS.

Verify:
`npx hardhat verify DUCKY_FAMILYV1_ADDRESS --contract 'contracts/duckies/games/DuckyFamily/DuckyFamilyV1.sol:DuckyFamilyV1' DUCKIES_ADDRESS DUCKLINGSV1_ADDRESS TREASURE_VAULT_ADDRESS --network polygon`

## Setup and connect DuckyFamily to Ducklings

`npx hardhat setupNFTs  --ducklings DUCKLINGSV1_ADDRESS  --ducky-family DUCKY_FAMILYV1_ADDRESS --api-base-url 'https://www.yellow.org/api/v3/nft/metadata/' --issuer 0xC5F825188ad49b13a8c5116FfDab7121b1CEf595 --network polygon`
