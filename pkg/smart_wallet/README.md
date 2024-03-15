# SmartWallet GoLang library

## Overview

SmartWallet is a golang library that provides helper functions to user operations.

## Features

This package introduces the following types and helper functionality:

- `Config` - A struct that holds the configuration for the SmartWallet.
- `Type` - an enum that represents the type of the wallet. Currently only ZeroDev `Kernel`, `Biconomy` and eth-infinitism `SimpleWallet` are supported.
- `IsAccountDeployed(swAddress)` - a function that checks if the smart wallet is deployed.
- `GetAccountAddress(owner, index)` - a function that calculates the address of the smart wallet.
- `GetInitCode(smartWalletConfig)` - a function that returns the init code of the smart wallet.
