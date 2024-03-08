# UserOp Golang Library

## Overview

The UserOp library is a Golang package designed to simplify the creation and interaction with user operations in a
decentralized application. It provides functionalities to work with smart walets, create user operations, and send them
to the client bundler for execution.

## Features

- **Smart Wallet Support**: Enables the use of smart wallets for transactions, including deploying new wallets and
  performing transactions through them.
- **Multi-Call Operations**: Supports bundling multiple operations into a single userOp, reducing the need for multiple
  transactions.
- **Gas Optimization**: Offers configurable gas pricing strategies to optimize transaction costs.
- **Paymaster Integration**: Integrates with paymasters to enable gasless transactions, where transaction fees can be
  paid using ERC20 tokens.
- **Infrastructure Flexibility**: Allows configuration for different infrastructure providers by setting the provider
  and bundler URLs.
- **Account abstraction interoperability**: Provides a unified interface for interacting with smart wallets, regardless
  of the underlying smart wallet implementation.

### Currently supported Smart Contract providers

#### Smart Wallets

- [Kernel v2.2](https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol)
- [Biconomy v2.0](https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/SmartAccount.sol)
- [SimpleAccount v0.6](https://github.com/eth-infinitism/account-abstraction/blob/releases/v0.6/contracts/samples/SimpleAccount.sol) (
  not finished)

#### Paymasters

- [Pimlico ERC20 Paymaster](https://github.com/pimlicolabs/erc20-paymaster-contracts/blob/44f7ff3e144ed0d4b893886d6b5d5586547a8ce7/src/PimlicoERC20Paymaster.sol)
- [Pimlico Verifying Paymaster](https://docs.pimlico.io/paymaster/verifying-paymaster)
- [Biconomy ERC20 Paymaster](https://docs.biconomy.io/Paymaster/methods#usage-mode-erc20) (not tested)
- [Biconomy Sponsoring Paymaster](https://docs.biconomy.io/Paymaster/methods#usage-mode-sponsored) (not tested)
- [Kernel v0.7 ERC20 Paymaster](https://github.com/eth-infinitism/account-abstraction/blob/releases/v0.7/contracts/samples/TokenPaymaster.sol) (
  to be added)

#### Signers

- [Biconomy ECDSAValidationModule](https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/SmartAccount.sol#L337)
- [Kernel ECDSAValidator](https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol#L127)
- [Kernel use / enable Validator](https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol#L126) (
  to be added)
- [Kernel SessionKeyValidator](https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/validator/SessionKeyValidator.sol#L138) (
  to be added)
- [SimpleAccount owner validation](https://github.com/eth-infinitism/account-abstraction/blob/releases/v0.6/contracts/samples/SimpleAccount.sol#L93-L99) (
  to be added)

## Installation

To use this library in your Golang project, you can install it using the following command:

```bash
go get github.com/layer-3/clearsync/pkg/userop
```

## Usage

### UserOp Client

All the functionalities provided by the UserOp library are accessed through the UserOp client.
The client is responsible for checking smart account information, creating user operations and sending them to the
client bundler.

UserOp client implements the following interface:

```go
type Client interface {
  IsAccountDeployed(ctx context.Context, owner common.Address, index decimal.Decimal) (bool, error)
  GetAccountAddress(ctx context.Context, owner common.Address, index decimal.Decimal) (common.Address, error)
  NewUserOp(
    ctx context.Context,
    sender common.Address,
    signer Signer,
    calls []Call,
    walletDeploymentOpts *WalletDeploymentOpts,
    gasLimitOverrides *GasLimitOverrides,
  ) (UserOperation, error)
  SendUserOp(ctx context.Context, op UserOperation) (done <-chan Receipt, err error)
}
```

### Configuration

The UserOp client requires a configuration struct to be created. The configuration struct allows you to specify the
provider, bundler, wallet, paymaster and other details.

Below is an example of a configuration struct (`testing/config.example.go`):

```go
var (
  exampleConfig = userop.ClientConfig{
    ProviderURL: *must(url.Parse("https://YOUR_PROVIDER_URL")),
    BundlerURL:  *must(url.Parse("https://YOUR_BUNDLER_URL")),
    EntryPoint:  common.HexToAddress("ENTRY_POINT_ADDRESS"),
    SmartWallet: userop.SmartWalletConfig{
      // Example of a Kernel Smart Wallet config with Kernel v2.2.
      Type: &userop.SmartWalletKernel,
      Factory: common.HexToAddress("0x5de4839a76cf55d0c90e2061ef4386d962E15ae3"), // Zerodev Kernel factory address:
      Logic:          common.HexToAddress("0x0DA6a956B9488eD4dd761E59f52FDc6c8068E6B5"), // Zerodev Kernel implementation (logic) address:
      ECDSAValidator: common.HexToAddress("0xd9AB5096a832b9ce79914329DAEE236f8Eea0390"),
    },
    Paymaster: userop.PaymasterConfig{
      // Example of a Pimlico USDC.e ERC20 Paymaster config.
      Type:    &userop.PaymasterPimlicoERC20,
      URL:     url.URL{}, // Pimlico ERC20 Paymaster does not require a URL.
      Address: common.HexToAddress("0xa683b47e447De6c8A007d9e294e87B6Db333Eb18"),
      PimlicoERC20: userop.PimlicoERC20Config{
        VerificationGasOverhead: decimal.RequireFromString("10000"), // verification gas overhead to be added to user op verification gas limit
      },
    },
    Gas: userop.GasConfig{
      // These are default values.
      MaxPriorityFeePerGasMultiplier: decimal.RequireFromString("1.13"),
      MaxFeePerGasMultiplier:         decimal.RequireFromString("2"),
    },
  }

  // wallet deployment options are used when creating a new user op for the smart wallet that is not deployed yet
  walletDeploymentOpts = &userop.WalletDeploymentOpts{
    Owner: common.HexToAddress("YOUR_OWNER_ADDRESS"),
    Index: decimal.NewFromInt(0),
  }

  // You can set either of gas limits when creating an user op to override the bundler's estimation.
  // Or you can set all of them to disable the bundler's estimation.
  gasLimitOverrides = &userop.GasLimitOverrides{
    CallGasLimit:         big.NewInt(42),
    VerificationGasLimit: big.NewInt(42),
    PreVerificationGas:   big.NewInt(42),
  }

  // signer is used to sign the user op upon creation
  exampleSigner = userop.SignerForKernel(must(crypto.HexToECDSA(
  "0xYOUR_PRIVATE_KEY")))
)
```

### Creating a UserOp Client

To create a UserOp client, you need to provide a configuration struct. If you want to change any of configuration
components, you should modify the configuration and create a new client.

```go
import "github.com/layer-3/clearsync/pkg/userop"

// Create a UserOp client
client, err := userop.NewClient(exampleConfig)
if err != nil {
    panic(errors.New("Error creating UserOp client:", err))
}
```

### Smart Account Operations

The UserOp client provides the following smart account-related functionalities:

#### Check if Smart Account is Deployed

Use `client.IsAccountDeployed` to check if a smart account is deployed for a given owner and index.

```go
// IsAccountDeployed checks whether the smart wallet for the specified owner EOA and index is deployed.
//
// Parameters:
//   - owner - is the EOA address of the smart wallet owner.
//   - index - is the index of the smart wallet, 0 by default. SW index allows to deploy multiple smart wallets for the same owner.
//
// Returns:
//   - bool - true if the smart wallet is deployed, false if not
//   - error - if failed to check.
func (c *backend) IsAccountDeployed(ctx context.Context, owner common.Address, index decimal.Decimal) (bool, error)
```

Usage example:

```go
ownerEOA := "0xEOAownerAddress"
accountIndex := 0 // Index of the smart wallet, 0 by default. EOA can have multiple smart wallets.

ctx := context.Background()
deployed, err := client.IsAccountDeployed(ctx, ownerEOA, accountIndex)
if err != nil {
    log.Fatal("Error checking if Smart Account is deployed:", err)
}
fmt.Printf("Smart Account for owner %s and index %d is deployed: %t\n", ownerEOA, accountIndex, deployed)
```

#### Calculate Smart Account Address

Use `client.GetAccountAddress` to calculate the address of a smart account for a given owner and index.

```go
// GetAccountAddress returns the address of the smart wallet for the specified owner EOA and index.
//
// Parameters:
//   - owner - is the EOA address of the smart wallet owner.
//   - index - is the index of the smart wallet, 0 by default. SW index allows to deploy multiple smart wallets for the same owner.
//
// Returns:
//   - common.Address - an address of the smart wallet
//   - error - if failed to calculate it.
func (c *backend) GetAccountAddress(ctx context.Context, owner common.Address, index decimal.Decimal) (common.Address, error)
```

Usage example:

```go
ownerEOA := "0xEOAownerAddress"
accountIndex := 0

ctx := context.Background()
address, err := client.GetAccountAddress(ctx, ownerEOA, accountIndex)
if err != nil {
    log.Fatal("Error calculating Smart Account address:", err)
}
fmt.Printf("Smart Account address for owner %s and index %d is %s\n", ownerEOA, accountIndex, address)
```

### User Operation Creation and Execution

#### Create User Operation

Use `client.NewUserOp` to create a new user operation.

```go
// NewUserOp builds a new UserOperation and fills all the fields.
//
// Parameters:
//   - ctx - is the context of the operation.
//   - smartWallet - is the address of the smart wallet that will execute the user operation.
//   - signer - is the signer function that will sign the user operation.
//   - calls - is the list of calls to be executed in the user operation.
//   - walletDeploymentOpts - are the options for the smart wallet deployment. Can be nil if the smart wallet is already deployed.
//
// Returns:
//   - UserOperation - user operation with all fields filled in.
//   - error - if failed to build the user operation.
func (c *backend) NewUserOp(
  ctx context.Context,
  smartWallet common.Address,
  signer Signer,
  calls []Call,
  walletDeploymentOpts *WalletDeploymentOpts,
) (UserOperation, error)
```

Usage example:

```go
sender := common.HexToAddress("0xsmartWalletAddress")
call := userop.Call{
    To:      common.HexToAddress("0xtoAddress"),
    Value:   1000000000000000000, // 1 ETH in wei
}

ctx := context.Background()
userOp, err := client.NewUserOp(ctx, sender, signer, userop.Call[]{call}, nil)
if err != nil {
    log.Fatal("Error creating User Operation:", err)
}
fmt.Printf("User Operation created: %v\n", userOp)
```

#### Send User Operation to Bundler

Use `client.SendUserOp` to send a user operation to the client bundler for execution.

```go
type Receipt struct {
  UserOpHash    common.Hash
  TxHash        common.Hash
  Sender        common.Address
  Nonce         decimal.Decimal
  Success       bool
  ActualGasCost decimal.Decimal
  ActualGasUsed decimal.Decimal
  RevertData    []byte // non-empty if Success is false and EntryPoint was able to catch revert reason.
}

// SendUserOp submits a user operation to a bundler and returns a channel to await for the userOp receipt.
//
// Parameters:
//   - ctx - is the context of the operation.
//   - op - is the user operation to be sent.
//
// Returns:
//   - <-chan Receipt - a channel to await for the userOp receipt.
//   - error - if failed to send the user operation
func SendUserOp(ctx context.Context, op UserOperation) (done <-chan Receipt, err error)
```

Usage example:

```go
ctx := context.Background()
receiptChannel, err := client.SendUserOp(ctx, userOp)
if err != nil {
    log.Fatal("Error sending User Operation to bundler:", err)
}

// Await the receipt from the channel
result := <-receiptChannel
fmt.Printf("User Operation result: %v\n", result)
```

## Contributing

Feel free to contribute by opening issues or submitting pull requests to this repository.

## License

This UserOp library is licensed under the MIT License.
