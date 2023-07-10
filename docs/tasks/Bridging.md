# Bridging

## Deploy Bridges

Firstly, a Bridge needs to be deployed on each network. There are two types of Bridges: Root and Child. Root Bridge is deployed on the chain where the Token originates from. Child Bridge is deployed on the chain where Token is bridged to.

> Be sure to deploy the Token on the other network before deploying the Bridge.

To deploy a Bridge, run the following command:

```bash
npx hardhat deployBridge --endpoint-address <endpoint_address> --token-address <token-address> [--is-root <true|false>] --network <network_name>
```

Endpoint addresses for each supported network can be found in the LayerZero docs: [testnet](https://layerzero.gitbook.io/docs/technical-reference/testnet/testnet-addresses) and [mainnet](https://layerzero.gitbook.io/docs/technical-reference/mainnet/supported-chain-ids).

If the Bridge is a Child Bridge, it also needs a permission to mint tokens. This permission will be granted to the Bridge during the deployment given the deployer can grant Token roles.

By default, the Bridge being deployed is a Child Bridge. To deploy a Root Bridge, you need to pass `--is-root true` parameter.

## Connect bridges

To connect Bridges, you need to run the following command for each Bridge:

```bash
npx hardhat addTrustedRemote --bridge-address <bridge_address> --remote-chain-id <remote_chain_id> --remote-address <remote_bridge_address> --network <network_name>
```

ChainId values are not related to EVM ids. Since LayerZero will span EVM & non-EVM chains the chainId are proprietary to LZ Endpoints.

ChainIds for each supported network can be found in the LayerZero docs: [testnet](https://layerzero.gitbook.io/docs/technical-reference/testnet/testnet-addresses) and [mainnet](https://layerzero.gitbook.io/docs/technical-reference/mainnet/supported-chain-ids).

## Bridge tokens

To bridge tokens, you need to run the following command:

```bash
npx hardhat bridgeToken [--receiver <receiver_address>] --amount <amount>  --bridge-address <bridge_address> --remote-chain-id <remote_chain_id> --network <network_name>
```

By default, `receiver` is the sender of this transaction. If you want to bridge tokens to another address, you need to specify `--receiver` parameter.

`amount` should be specified without Token decimals. E.g. if a Token has 8 decimals, to bridge 1000 tokens, you need to pass `1000`.

ChainIds for each supported network can be found in the LayerZero docs: [testnet](https://layerzero.gitbook.io/docs/technical-reference/testnet/testnet-addresses) and [mainnet](https://layerzero.gitbook.io/docs/technical-reference/mainnet/supported-chain-ids).

The task will check your allowance and approve the Bridge to spend your tokens if needed.

The task will calculate native fees needed for bridging and will fail if you don't have enough native tokens to pay for the fees.
