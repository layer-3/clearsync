import { task } from 'hardhat/config';

import type { YellowToken } from '../../../typechain-types';

const ethereumCheckpointManagerAddress = '0x86e4dc95c7fbdbf52e33d563bbdb00823894c287';
const ethereumFxRootAddress = '0xfe5e5D361b2ad62c541bAb87C45a0B9B018389a2';

const goerliCheckpointManagerAddress = '0x2890bA17EfE978480615e330ecB65333b880928e';
const goerliFxRootAddress = '0x3d1d3E34f7fB6D26245E6640E1c50710eFFf15bA';

const polygonFxChildAddress = '0x8397259c983751DAf40400790063935a11afa28a';

const mumbaiFxChildAddress = '0xCf73231F28B7331BBe3124B907840A94851f9f11';

interface SetupTunnelsArgs {
  rootTokenAddress: string;
  childTokenAddress: string;
  isMainnet?: boolean;
}

task(
  'setupTunnels',
  'Deploy Yellow Token Root (Ethereum) and Child (Polygon) Tunnel contracts, link them together and with Yellow Token on Ethereum and Polygon.',
)
  .addParam('rootTokenAddress', 'The address of the Yellow token on Ethereum')
  .addParam('childTokenAddress', 'The address of the Yellow token on Polygon')
  .addOptionalParam<boolean>('isMainnet', 'true if mainnet, false if testnet (default)')
  .setAction(async (taskArgs: SetupTunnelsArgs, hre) => {
    const { ethers } = hre;
    const { rootTokenAddress, childTokenAddress } = taskArgs;

    const isMainnet = taskArgs.isMainnet ?? false;
    const ethereumNetwork = isMainnet ? 'ethereum' : 'goerli';
    const polygonNetwork = isMainnet ? 'matic' : 'mumbai';

    console.log(`Running on ${ethereumNetwork} and ${polygonNetwork}`);

    // -------- Deploy Yellow Token Root Tunnel --------

    hre.changeNetwork(ethereumNetwork);

    const checkpointManagerAddress = isMainnet
      ? ethereumCheckpointManagerAddress
      : goerliCheckpointManagerAddress;

    const fxRootAddress = isMainnet ? ethereumFxRootAddress : goerliFxRootAddress;

    console.log('Deploying YellowTokenRootTunnel...');

    const YellowTokenRootTunnelFactory = await ethers.getContractFactory('YellowTokenRootTunnel');
    const YellowTokenRootTunnel = await YellowTokenRootTunnelFactory.deploy(
      childTokenAddress,
      rootTokenAddress,
      checkpointManagerAddress,
      fxRootAddress,
    );
    await YellowTokenRootTunnel.deployed();

    console.log(`YellowTokenRootTunnel deployed to: ${YellowTokenRootTunnel.address}`);

    // -------- Grant MINTER_ROLE on RootYellowToken to YellowTokenRootTunnel --------

    const RootYellowToken = (await ethers.getContractAt(
      'YellowToken',
      rootTokenAddress,
    )) as YellowToken;
    await RootYellowToken.grantRole(
      await RootYellowToken.MINTER_ROLE(),
      YellowTokenRootTunnel.address,
    );

    console.log(
      `MINTER_ROLE on RootYellowToken (${rootTokenAddress}) granted to YellowTokenRootTunnel(${YellowTokenRootTunnel.address})\n`,
    );

    // -------- Deploy Yellow Token Child Tunnel --------

    hre.changeNetwork(polygonNetwork);

    const fxChildAddress = isMainnet ? polygonFxChildAddress : mumbaiFxChildAddress;

    console.log('Deploying YellowTokenChildTunnel...');

    const YellowTokenChildTunnelFactory = await ethers.getContractFactory('YellowTokenChildTunnel');
    const YellowTokenChildTunnel = await YellowTokenChildTunnelFactory.deploy(
      childTokenAddress,
      rootTokenAddress,
      fxChildAddress,
    );
    await YellowTokenChildTunnel.deployed();

    console.log(`YellowTokenChildTunnel deployed to: ${YellowTokenChildTunnel.address}\n`);

    // -------- Link Yellow Token Root and Child Tunnels --------

    hre.changeNetwork(ethereumNetwork);

    await YellowTokenRootTunnel.setFxChildTunnel(YellowTokenChildTunnel.address);

    console.log(
      `YellowTokenChildTunnel(${YellowTokenChildTunnel.address}) was set on YellowTokenRootTunnel(${YellowTokenRootTunnel.address})`,
    );

    hre.changeNetwork(polygonNetwork);

    await YellowTokenChildTunnel.setFxRootTunnel(YellowTokenRootTunnel.address);

    console.log(
      `YellowTokenRootTunnel(${YellowTokenRootTunnel.address}) was set on YellowTokenChildTunnel(${YellowTokenChildTunnel.address})`,
    );

    console.log('Done!');
  });
