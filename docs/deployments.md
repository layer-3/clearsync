# Yellow Network Smart Contract Deployments

Yellow Network may be conducting testing and operations on different networks, and it is crucial to keep track of the deployed contracts and their addresses.

## Mainnet

### Ethereum

Chain id: 1.

Last updated: January 17, 2024.

| Description                | Contract Name | Address                                                                                                                  | Git SHA and Repository                                                                                                     |
| -------------------------- | ------------- | ------------------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------- |
| Yellow Duckies ERC20 Token | YellowToken   | `0x90b7E285ab6cf4e3A2487669dba3E339dB8a3320`[↗](https://etherscan.io/address/0x90b7E285ab6cf4e3A2487669dba3E339dB8a3320) | `c197beb`[↗](https://github.com/layer-3/clearsync/blob/c197bebe236ba3134ca2de8c0ac6fa08c2550430/contracts/YellowToken.sol) |

### Polygon

Chain id: 137.

Last updated: April 5, 2024.

| Description                     | Contract Name     | Address                                                                                                                     | Git SHA and Repository                                                                                                                       |
| ------------------------------- | ----------------- | --------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| Yellow Duckies ERC20 Token      | YellowToken       | `0x18e73A5333984549484348A94f4D219f4faB7b81`[↗](https://polygonscan.com/address/0x18e73A5333984549484348A94f4D219f4faB7b81) | `7b9389c`[↗](https://github.com/layer-3/duckies/blob/7b9389c83a540870369b82d0a1144d61f50f9c35/contracts/YellowToken.sol)                     |
| Yellow Ducklings NFT            | DucklingsV1       | `0x435b74f6DC4A0723CA19e4dD2AC8Aa1361c7B0f0`[↗](https://polygonscan.com/address/0x435b74f6DC4A0723CA19e4dD2AC8Aa1361c7B0f0) | `7b9389c`[↗](https://github.com/layer-3/duckies/blob/7b9389c83a540870369b82d0a1144d61f50f9c35/contracts/ducklings/DucklingsV1.sol)           |
| Yellow Ducklings NFT Game       | DuckyFamilyV1     | `0xb66bf78cad7cbab51988ddc792652cbabdff7675`[↗](https://polygonscan.com/address/0xb66bf78cad7cbab51988ddc792652cbabdff7675) | `7b9389c`[↗](https://github.com/layer-3/duckies/blob/7b9389c83a540870369b82d0a1144d61f50f9c35/contracts/games/DuckyFamily/DuckyFamilyV1.sol) |
| Yellow Duckies Treasury         | TreasureVault     | `0x68d1E3F802058Ce517e9ba871Ab182299E74D852`[↗](https://polygonscan.com/address/0x68d1E3F802058Ce517e9ba871Ab182299E74D852) | `7b9389c`[↗](https://github.com/layer-3/duckies/blob/7b9389c83a540870369b82d0a1144d61f50f9c35/contracts/TreasureVault.sol)                   |
| Yellow Network Adjudicator      | YellowAdjudicator | `0xf81A43EBA92538B0323fCDb1A040F2183B352Ca3`[↗](https://polygonscan.com/address/0xf81A43EBA92538B0323fCDb1A040F2183B352Ca3) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/YellowAdjudicator.sol)    |
| State channels Margin App       | ConsensusApp      | `0xd3f6EA0DCe26E7836fB309dcfcf506e44524B2A5`[↗](https://polygonscan.com/address/0xd3f6EA0DCe26E7836fB309dcfcf506e44524B2A5) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)          |
| State channels Escrow App       | ConsensusApp      | `0x59735037AC294641F8CE51d68D6C45a500B8e645`[↗](https://polygonscan.com/address/0x59735037AC294641F8CE51d68D6C45a500B8e645) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)          |
| Modified Solady Create3 Factory | SoladyCreate3     | `0x6c0bbB08ea7926E378ff2068af696E613E0B0cBb`[↗](https://polygonscan.com/address/0x6c0bbB08ea7926E378ff2068af696E613E0B0cBb) | `9b31a70`[↗](https://github.com/layer-3/chaintools/blob/9b31a7018b5082c9472ac99b4475270b4ae8b689/contracts/SoladyCreate3.sol)                |

### Linea

Chain id: 59144.

#### Tokens Bridged

Last updated: January 17, 2024.

Bridged Token Smart Contract: [`BridgedToken.sol`](https://github.com/Consensys/linea-contracts/blob/3cf85529fd4539eb06ba998030c37e47f98c528a/contracts/tokenBridge/BridgedToken.sol)

| Description                | Address                                                                                                                   |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| Yellow Duckies ERC20 Token | `0x796000FAd0d00B003B9dd8e531BA90cff39E01E0`[↗](https://lineascan.build/token/0x796000FAd0d00B003B9dd8e531BA90cff39E01E0) |

#### Contracts Deployed

Last updated: January 31, 2024.

| Description                | Contract Name     | Address                                                                                                                     | Git SHA and Repository                                                                                                                    |
| -------------------------- | ----------------- | --------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| Yellow Network Adjudicator | YellowAdjudicator | `0x0ba4d7cb34ff4b4a60338a0234576f52d1631182`[↗](https://lineascan.build/address/0x0ba4d7cb34ff4b4a60338a0234576f52d1631182) | `58f7c97`[↗](https://github.com/layer-3/clearsync/blob/58f7c97934585b72ba05eceb2d8737f76d51dece/contracts/clearing/YellowAdjudicator.sol) |
| State channels Margin App  | ConsensusApp      | `0x6178d14644d29c389b9fdb3b0d25dbdc7c428cad`[↗](https://lineascan.build/address/0x6178d14644d29c389b9fdb3b0d25dbdc7c428cad) | `58f7c97`[↗](https://github.com/layer-3/clearsync/blob/58f7c97934585b72ba05eceb2d8737f76d51dece/contracts/clearing/ConsenusApp.sol)       |
| State channels Escrow App  | ConsensusApp      | `0xa230bc7f76351dfbf97064a16e0b1a9e141cbf9c`[↗](https://lineascan.build/address/0xa230bc7f76351dfbf97064a16e0b1a9e141cbf9c) | `58f7c97`[↗](https://github.com/layer-3/clearsync/blob/58f7c97934585b72ba05eceb2d8737f76d51dece/contracts/clearing/ConsenusApp.sol)       |

## Testnet

### Sepolia

Chain id: 11155111.

Last updated: January 16, 2024.

| Description                     | Contract Name     | Address                                                                                                                          | Git SHA and Repository                                                                                                                    |
| ------------------------------- | ----------------- | -------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| Yellow Network Adjudicator      | YellowAdjudicator | `0x47871f064d0b2ABf9190275C4D69f466C98fBD77`[↗](https://sepolia.etherscan.io/address/0xf81a43eba92538b0323fcdb1a040f2183b352ca3) | `418c724`[↗](https://github.com/layer-3/clearsync/blob/418c72494d451b5c4593b602993b71d95517b83f/contracts/clearing/YellowAdjudicator.sol) |
| State channels Margin App       | ConsensusApp      | `0xa6F5563CD2D38a0c1F2D41DF7Eff7181bf3c6a7e`[↗](https://sepolia.etherscan.io/address/0xa6f5563cd2d38a0c1f2d41df7eff7181bf3c6a7e) | `418c724`[↗](https://github.com/layer-3/clearsync/blob/418c72494d451b5c4593b602993b71d95517b83f/contracts/clearing/ConsenusApp.sol)       |
| State channels Escrow App       | ConsensusApp      | `0xcccb67333fEefb04e85521fF0c219Cdb12539b84`[↗](https://sepolia.etherscan.io/address/0xcccb67333feefb04e85521ff0c219cdb12539b84) | `418c724`[↗](https://github.com/layer-3/clearsync/blob/418c72494d451b5c4593b602993b71d95517b83f/contracts/clearing/ConsenusApp.sol)       |
| USD Tether ERC20 Token          | TestERC20         | `0x98e255A2D9e36d5174a7787aBA7053e60F47Fc08`[↗](https://sepolia.etherscan.io/address/0x98e255a2d9e36d5174a7787aba7053e60f47fc08) | `418c724`[↗](https://github.com/layer-3/clearsync/blob/418c72494d451b5c4593b602993b71d95517b83f/contracts/test/TestERC20.sol)             |
| USD Coin ERC20 Token            | TestERC20         | `0x42b757f0B530cb44139ceDd9F0C47249175CBC7E`[↗](https://sepolia.etherscan.io/address/0x42b757f0b530cb44139cedd9f0c47249175cbc7e) | `418c724`[↗](https://github.com/layer-3/clearsync/blob/418c72494d451b5c4593b602993b71d95517b83f/contracts/test/TestERC20.sol)             |
| Shiba Inu ERC20 Token           | TestERC20         | `0x3e594179ad7E013f817bCddF310a7e75b2b069a9`[↗](https://sepolia.etherscan.io/address/0x3e594179ad7e013f817bcddf310a7e75b2b069a9) | `418c724`[↗](https://github.com/layer-3/clearsync/blob/418c72494d451b5c4593b602993b71d95517b83f/contracts/test/TestERC20.sol)             |
| Wrapped Bitcoin ERC20 Token     | TestERC20         | `0x9cEF6720Ba49c8C94Df1CfA0D713828B7B9fAEB1`[↗](https://sepolia.etherscan.io/address/0x9cef6720ba49c8c94df1cfa0d713828b7b9faeb1) | `418c724`[↗](https://github.com/layer-3/clearsync/blob/418c72494d451b5c4593b602993b71d95517b83f/contracts/test/TestERC20.sol)             |
| Wrapped Ether ERC20 Token       | TestERC20         | `0x2Ff553CA05b647b0e352fe25828BB754a35Ff7dE`[↗](https://sepolia.etherscan.io/address/0x2ff553ca05b647b0e352fe25828bb754a35ff7de) | `418c724`[↗](https://github.com/layer-3/clearsync/blob/418c72494d451b5c4593b602993b71d95517b83f/contracts/test/TestERC20.sol)             |
| Wrapped Matic ERC20 Token       | TestERC20         | `0x4e519DD7b0137D3D7a3Fe8EA2aC38E9c598230DB`[↗](https://sepolia.etherscan.io/address/0x4e519dd7b0137d3d7a3fe8ea2ac38e9c598230db) | `418c724`[↗](https://github.com/layer-3/clearsync/blob/418c72494d451b5c4593b602993b71d95517b83f/contracts/test/TestERC20.sol)             |
| Yellow Duckies ERC20 Token      | TestERC20         | `0x63FD175d3215779deBA7532fC660fA0E10c18676`[↗](https://sepolia.etherscan.io/address/0x63fd175d3215779deba7532fc660fa0e10c18676) | `418c724`[↗](https://github.com/layer-3/clearsync/blob/418c72494d451b5c4593b602993b71d95517b83f/contracts/test/TestERC20.sol)             |
| Modified Solady Create3 Factory | SoladyCreate3     | `0x6c0bbB08ea7926E378ff2068af696E613E0B0cBb`[↗](https://sepolia.etherscan.io/address/0x6c0bbB08ea7926E378ff2068af696E613E0B0cBb) | `9b31a70`[↗](https://github.com/layer-3/chaintools/blob/9b31a7018b5082c9472ac99b4475270b4ae8b689/contracts/SoladyCreate3.sol)             |

### Amoy

Chain id: 80002.

Last updated: April 5, 2024.

| Description                     | Contract Name     | Address                                                                                                                         | Git SHA and Repository                                                                                                                    |
| ------------------------------- | ----------------- | ------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| Yellow Network Adjudicator      | YellowAdjudicator | `0xf81A43EBA92538B0323fCDb1A040F2183B352Ca3`[↗](https://www.oklink.com/amoy/address/0xf81A43EBA92538B0323fCDb1A040F2183B352Ca3) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/YellowAdjudicator.sol) |
| State channels Margin App       | ConsensusApp      | `0xd3f6EA0DCe26E7836fB309dcfcf506e44524B2A5`[↗](https://www.oklink.com/amoy/address/0xd3f6EA0DCe26E7836fB309dcfcf506e44524B2A5) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| State channels Escrow App       | ConsensusApp      | `0x59735037AC294641F8CE51d68D6C45a500B8e645`[↗](https://www.oklink.com/amoy/address/0x59735037AC294641F8CE51d68D6C45a500B8e645) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| Yellow Duckies ERC20 Token      | TestERC20         | `0x38871D722689f6d799B2A1FE93bE096E98e00986`[↗](https://www.oklink.com/amoy/address/0x38871D722689f6d799B2A1FE93bE096E98e00986) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/test/TestERC20.sol)             |
| USD Tether ERC20 Token          | TestERC20         | `0xE22e3af98FdF9d6a429150D92d446fcEeEA967fF`[↗](https://www.oklink.com/amoy/address/0xE22e3af98FdF9d6a429150D92d446fcEeEA967fF) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/test/TestERC20.sol)             |
| USD Coin ERC20 Token            | TestERC20         | `0x8d48ba6D6ABD283E672B917cdfBd6222DD1b80dB`[↗](https://www.oklink.com/amoy/address/0x8d48ba6D6ABD283E672B917cdfBd6222DD1b80dB) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/test/TestERC20.sol)             |
| Lube ERC20 Token                | TestERC20         | `0x57634f3202E291e37801C426408b86a1eD0dB23d`[↗](https://www.oklink.com/amoy/address/0x57634f3202E291e37801C426408b86a1eD0dB23d) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/test/TestERC20.sol)             |
| Wrapped Bitcoin ERC20 Token     | TestERC20         | `0xE11E73Ce5f7c45580e6B4A63069C18b0A5d70CA2`[↗](https://www.oklink.com/amoy/address/0xE11E73Ce5f7c45580e6B4A63069C18b0A5d70CA2) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/test/TestERC20.sol)             |
| Wrapped Ether ERC20 Token       | TestERC20         | `0xEB31ce20F5c72C568e82D0183aa8Da3B9e5e177f`[↗](https://www.oklink.com/amoy/address/0xEB31ce20F5c72C568e82D0183aa8Da3B9e5e177f) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/test/TestERC20.sol)             |
| Linda ERC20 Token               | TestERC20         | `0xd15833D5184Ca08a8a26b923754EaA654F5Bb85F`[↗](https://www.oklink.com/amoy/address/0xd15833D5184Ca08a8a26b923754EaA654F5Bb85F) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/test/TestERC20.sol)             |
| Modified Solady Create3 Factory | SoladyCreate3     | `0x6c0bbB08ea7926E378ff2068af696E613E0B0cBb`[↗](https://www.oklink.com/amoy/address/0x6c0bbB08ea7926E378ff2068af696E613E0B0cBb) | `9b31a70`[↗](https://github.com/layer-3/chaintools/blob/9b31a7018b5082c9472ac99b4475270b4ae8b689/contracts/SoladyCreate3.sol)             |

### Linea-sepolia

Chain id: 59141.

Last updated: June 7, 2024.

| Description                     | Contract Name     | Address                                                                                                                             | Git SHA and Repository                                                                                                                    |
| ------------------------------- | ----------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| Yellow Network Adjudicator      | YellowAdjudicator | `0xf81A43EBA92538B0323fCDb1A040F2183B352Ca3`[↗](https://sepolia.lineascan.build/address/0xf81A43EBA92538B0323fCDb1A040F2183B352Ca3) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/YellowAdjudicator.sol) |
| State channels Margin App       | ConsensusApp      | `0xd3f6EA0DCe26E7836fB309dcfcf506e44524B2A5`[↗](https://sepolia.lineascan.build/address/0xd3f6EA0DCe26E7836fB309dcfcf506e44524B2A5) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| State channels Escrow App       | ConsensusApp      | `0x59735037AC294641F8CE51d68D6C45a500B8e645`[↗](https://sepolia.lineascan.build/address/0x59735037AC294641F8CE51d68D6C45a500B8e645) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| Yellow Duckies ERC20 Token      | TestERC20         | `0x38871D722689f6d799B2A1FE93bE096E98e00986`[↗](https://sepolia.lineascan.build/address/0x38871D722689f6d799B2A1FE93bE096E98e00986) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| USD Tether ERC20 Token          | TestERC20         | `0xE22e3af98FdF9d6a429150D92d446fcEeEA967fF`[↗](https://sepolia.lineascan.build/address/0xE22e3af98FdF9d6a429150D92d446fcEeEA967fF) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| USD Coin ERC20 Token            | TestERC20         | `0x8d48ba6D6ABD283E672B917cdfBd6222DD1b80dB`[↗](https://sepolia.lineascan.build/address/0x8d48ba6D6ABD283E672B917cdfBd6222DD1b80dB) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| Lube ERC20 Token                | TestERC20         | `0x57634f3202E291e37801C426408b86a1eD0dB23d`[↗](https://sepolia.lineascan.build/address/0x57634f3202E291e37801C426408b86a1eD0dB23d) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| Wrapped Bitcoin ERC20 Token     | TestERC20         | `0xE11E73Ce5f7c45580e6B4A63069C18b0A5d70CA2`[↗](https://sepolia.lineascan.build/address/0xE11E73Ce5f7c45580e6B4A63069C18b0A5d70CA2) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| Wrapped Ether ERC20 Token       | TestERC20         | `0xEB31ce20F5c72C568e82D0183aa8Da3B9e5e177f`[↗](https://sepolia.lineascan.build/address/0xEB31ce20F5c72C568e82D0183aa8Da3B9e5e177f) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| Linda ERC20 Token               | TestERC20         | `0xd15833D5184Ca08a8a26b923754EaA654F5Bb85F`[↗](https://sepolia.lineascan.build/address/0xd15833D5184Ca08a8a26b923754EaA654F5Bb85F) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| BNB ERC20 Token                 | TestERC20         | `0xd14c1883B0554c7Ca43797D8F06618922DEE973A`[↗](https://sepolia.lineascan.build/address/0xd14c1883B0554c7Ca43797D8F06618922DEE973A) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| Pepe ERC20 Token                | TestERC20         | `0xC981Cd9be7a334e8F5ef24F245BbF18b181B8fa9`[↗](https://sepolia.lineascan.build/address/0xC981Cd9be7a334e8F5ef24F245BbF18b181B8fa9) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| Badger ERC20 Token              | TestERC20         | `0x2A7a211fB5860C63f2F53c8CB1FE400aBe222b92`[↗](https://sepolia.lineascan.build/address/0x2A7a211fB5860C63f2F53c8CB1FE400aBe222b92) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| FLOKI ERC20 Token               | TestERC20         | `0xEC077c58C614707D4c09473C8ed1B359F7D2C935`[↗](https://sepolia.lineascan.build/address/0xEC077c58C614707D4c09473C8ed1B359F7D2C935) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| SHIBA INU ERC20 Token           | TestERC20         | `0x0fe19c1da3a8af033f11A349631Ea60F8244e6Ce`[↗](https://sepolia.lineascan.build/address/0x0fe19c1da3a8af033f11A349631Ea60F8244e6Ce) | `2c6129a`[↗](https://github.com/layer-3/clearsync/blob/2c6129a1831778972c4f07766ec69e25674d2438/contracts/clearing/ConsenusApp.sol)       |
| Modified Solady Create3 Factory | SoladyCreate3     | `0x6c0bbB08ea7926E378ff2068af696E613E0B0cBb`[↗](https://sepolia.lineascan.build/address/0x6c0bbB08ea7926E378ff2068af696E613E0B0cBb) | `9b31a70`[↗](https://github.com/layer-3/chaintools/blob/9b31a7018b5082c9472ac99b4475270b4ae8b689/contracts/SoladyCreate3.sol)             |
