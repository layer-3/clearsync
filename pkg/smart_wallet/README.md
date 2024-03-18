# SmartWallet GoLang library

## Overview

SmartWallet is a golang library that provides helper functions to user operations.

## Features

This package introduces the following types and helper functionality:

- `Config` - A struct that holds the configuration for the SmartWallet.
- `Type` - an enum that represents the type of the wallet. Currently only ZeroDev [`Kernel v2.2`](https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol), [`Biconomy v2.0`](https://github.com/bcnmy/scw-contracts/blob/v2-deployments/contracts/smart-account/SmartAccount.sol) and eth-infinitism [`SimpleAccount v0.6` (still in progress)](https://github.com/eth-infinitism/account-abstraction/blob/releases/v0.6/contracts/samples/SimpleAccount.sol) are supported.
- `IsAccountDeployed(swAddress)` - a function that checks if the smart wallet is deployed.
- `GetAccountAddress(owner, index)` - a function that calculates the address of the smart wallet.
- `GetInitCode(smartWalletConfig)` - a function that returns the init code of the smart wallet.
