# Bridging Token

## Brindging Checklist

1. Deploy ChildToken

   `NAME=YellowToken ARGS=<name>,<symbol>,<supply_cap> npx hardhat run scripts/deployContract.ts`

2. Activate ChildToken

   `npx hardhat activate --token-address <token_address> --premint <premint_amount> --premint-to <premint_to_address> --network <network_name>`

   > Note: As this is child token, by convention it should not have any tokens minted before brindging enabled.
   > For this reason, specify `--premint` as small as possible (However, with current restrictions, it should be greater than 0)

3. Deploy RootBridge

   `npx hardhat deployBridge --endpoint-address <endpoint_address> --token-address <token_address> --is-root true --network <network_name>`

   Endpoint addresses for each supported network can be found in the LayerZero docs: [testnet](https://layerzero.gitbook.io/docs/technical-reference/testnet/testnet-addresses) and [mainnet](https://layerzero.gitbook.io/docs/technical-reference/mainnet/supported-chain-ids).

   > Note: this task will also grant `MINTER_ROLE` to the bridge contract if the executor of the task is an Admin of the token contract.

4. Deploy ChildBridge

   `npx hardhat deployBridge --endpoint-address <endpoint_address> --token-address <token_address> --network <network_name>`

   Endpoint addresses for each supported network can be found in the LayerZero docs: [testnet](https://layerzero.gitbook.io/docs/technical-reference/testnet/testnet-addresses) and [mainnet](https://layerzero.gitbook.io/docs/technical-reference/mainnet/supported-chain-ids).

5. Connect Bridges

   `npx hardhat addTrustedRemote --bridge-address <bridge_address> --remote-chain-id <remote_chain_id> --remote-address <remote_bridge_address> --network <network_name>`

   ChainId values are not related to EVM ids. Since LayerZero will span EVM & non-EVM chains the chainId are proprietary to LZ Endpoints.

   ChainIds for each supported network can be found in the LayerZero docs: [testnet](https://layerzero.gitbook.io/docs/technical-reference/testnet/testnet-addresses) and [mainnet](https://layerzero.gitbook.io/docs/technical-reference/mainnet/supported-chain-ids).

6. Bridge tokens

   `npx hardhat bridgeToken [--receiver <receiver_address>] --amount <amount>  --bridge-address <bridge_address> --remote-chain-id <remote_chain_id> --network <network_name>`

   > Note: `amount` should be specified without Token decimals. E.g. if a Token has 8 decimals, to bridge 1000 tokens, you need to pass `1000`.

   ChainIds for each supported network can be found in the LayerZero docs: [testnet](https://layerzero.gitbook.io/docs/technical-reference/testnet/testnet-addresses) and [mainnet](https://layerzero.gitbook.io/docs/technical-reference/mainnet/supported-chain-ids).

7. Wait for the bridge to confirm the transaction

   You can check the status of the bridge transaction by viewing it by specifying tx hash on the LayerZero Explorer: [testnet](https://testnet.layerzeroscan.com/) and [mainnet](hhttps://layerzeroscan.com/).
