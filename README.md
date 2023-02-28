
# Yellow Clearing Network

This repository contains the smart contracts of the [Yellow Network](https://www.yellow.org).

Yellow Network is a Layer-3 peer-to-peer network that uses [State Channels](https://statechannels.org/) technology to scale and facilitate trading, clearing and settlement.
## License

[GPL-3.0](https://spdx.org/licenses/GPL-3.0-or-later.html)

ClearSync is licensed under the GNU General Public License v3.0

## Documentation

 - [Project Wiki](https://docs.yellow.org)
 - [Technical Paper](docs/clearsync-paper.md)

### Architecture overview
### Project purpose
### Business use cases
### Performance requirements
## Features

- Open and Closing Trading channels
    - Deposit fee using YELLOW / DUCKIES
    - Deposit Collateral in USDT, USDC, DAI, WETH, WBTC
- Example Price feed from major providers
- Example Risk Management module
- Sending High-frequency Margin Request off-chain
- Ability to request a settlement
    - Using our provided JointCustody
    - Using other HTLC and escrow smart-contracts
    - Using third party custodian
- Ability to challenge a Margin state to unlock the collateral


## Usage

```bash
npx hardhat help
npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.ts
```

