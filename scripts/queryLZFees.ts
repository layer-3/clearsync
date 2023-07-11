import hre, { ethers } from 'hardhat';

import { mainnetChainIdToLZInfo, testnetChainIdToLZInfo } from '../src/crosschain/lz-networks';
import { networkIdToName } from '../src/networks';

import type { ILayerZeroEndpoint } from '../typechain-types';

async function main(): Promise<void> {
  const [Sender] = await ethers.getSigners();

  const isMainnet = process.env.MAINNET === 'true';

  const UAAddress = '0x000000000000000000000000000000000000dEaD';
  const payload = Sender.address + '420000000000';
  const payInZRO = false;
  const adapterParams = '0x';

  const chainIdToLZInfo = isMainnet ? mainnetChainIdToLZInfo : testnetChainIdToLZInfo;

  const totalFees = [];

  for (const [chainId, lzInfo] of chainIdToLZInfo.entries()) {
    const { lzChainId: srcLZChainId, endpointAddress } = lzInfo;

    const networkName = networkIdToName.get(chainId);
    if (!networkName) {
      throw new Error(`Unknown networkId ${chainId}`);
    }

    console.log(
      `estimating fees for ${networkName} (chainId ${chainId}, LZ chainId ${srcLZChainId})`,
    );

    const feesForSrcChain = [];

    // change network
    hre.changeNetwork(networkName);

    // initialize Endpoint
    const lzEndpoint = (await ethers.getContractAt(
      'ILayerZeroEndpoint',
      endpointAddress,
    )) as ILayerZeroEndpoint;

    // fetch fees for each destination chain
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    for (const [__, lzInfo] of chainIdToLZInfo.entries()) {
      const { lzChainId: dstLZChainId } = lzInfo;

      if (srcLZChainId === dstLZChainId) {
        feesForSrcChain.push(0);
        continue;
      }

      let nativeFee = -1;

      try {
        // estimate fee
        const [nativeFeeBN] = await lzEndpoint.estimateFees(
          dstLZChainId,
          UAAddress,
          payload,
          payInZRO,
          adapterParams,
        );

        // convert to number
        const feeStr = nativeFeeBN.toString();
        const notRoundedPart = String(Number(feeStr.slice(0, 3)) + 1);
        const feeRoundedStr = notRoundedPart + '0'.repeat(feeStr.length - notRoundedPart.length);
        const feeRoundedBN = ethers.BigNumber.from(feeRoundedStr);
        nativeFee = Number(ethers.utils.formatEther(feeRoundedBN));
      } catch {
        console.log(`error estimating fees for chain: ${dstLZChainId}...\n`);
      }

      feesForSrcChain.push(nativeFee);
    }

    totalFees.push(feesForSrcChain);
  }

  console.log(totalFees);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
