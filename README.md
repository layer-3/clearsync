# Yellow Network Smart Contracts

This repository contains the smart contracts of the [Yellow Network](https://www.yellow.org).

## What is Yellow Network?

Yellow Network is a Layer-3 peer-to-peer network that uses state channel technology to securely exchange liquidity & facilitate trading, clearing, settlement, and compliance.

## DUCKIES as a canary network of YELLOW

### What is a canary network?

Canary networks inherited their name after the historical usage of canaries in the coal mines. Like brave little birds being great early detectors, Duckies network acts as a pioneering development playground where Yellow experiments with technologies and gathers real users‘ feedback in fast-paced, dynamic conditions.
Duckies network is a living platform that tests next-gen tech and functionality under the most realistic conditions before they are integrated into Yellow Network.

## Smart Contracts

### contracts/Token.sol

This smart contract is the ERC20 used by both YELLOW and DUCKIES tokens. The YELLOW token is collateral to open a state channel with another network entity. Additionally, it is used to pay the settlement fees on the network.

### contracts/duckies/

This directory contains DUCKIES-specific smart contracts. It includes specific features included in the Duckies project only. Specific features are related to the Duckies game, designed to stimulate the YELLOW community while the Yellow network is under development.
