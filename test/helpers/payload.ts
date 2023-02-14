import { Contract, utils } from 'ethers';

import type { ParamType } from 'ethers/lib/utils';
import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { Message, NFTMeldMessage, NFTMintMessage, NFTsMintMessage } from './types';

export interface TestContext {
  duckies: Contract;
  owner: SignerWithAddress;
  user: SignerWithAddress;
  referer: SignerWithAddress;
  others: SignerWithAddress[];
}

export function hashMessage(message: Message): string {
  return utils.keccak256(
    utils.defaultAbiCoder.encode(
      [
        {
          type: 'tuple',
          components: [
            { name: 'bounty_hash', type: 'bytes32' },
            { name: 'beneficiary', type: 'address' },
            { name: 'referrer', type: 'address' },
            { name: 'amount', type: 'uint32' },
            { name: 'expire', type: 'uint64' },
            { name: 'chain_id', type: 'uint32' },
          ],
        } as ParamType,
      ],
      [message],
    ),
  );
}

export function hashNFTMintMessage(nftMintMessage: NFTMintMessage): string {
  return utils.keccak256(
    utils.defaultAbiCoder.encode(
      [
        {
          type: 'tuple',
          components: [
            { name: 'mint_nft_hash', type: 'bytes32' },
            { name: 'token_uri', type: 'string' },
            { name: 'beneficiary', type: 'address' },
            { name: 'rate', type: 'uint64' },
            { name: 'expire', type: 'uint64' },
            { name: 'chain_id', type: 'uint32' },
          ],
        } as ParamType,
      ],
      [nftMintMessage],
    ),
  );
}

export function hashNFTsMintMessage(nftMintMessage: NFTsMintMessage): string {
  return utils.keccak256(
    utils.defaultAbiCoder.encode(
      [
        {
          components: [
            {
              components: [
                {
                  name: 'is_common',
                  type: 'bool',
                },
                {
                  name: 'token_uri',
                  type: 'string',
                },
              ],
              name: 'messages',
              type: 'tuple[]',
            },
            {
              name: 'rate',
              type: 'uint64',
            },
            {
              name: 'beneficiary',
              type: 'address',
            },
            {
              name: 'expire',
              type: 'uint64',
            },
            {
              name: 'chain_id',
              type: 'uint32',
            },
          ],
          type: 'tuple',
        } as ParamType,
      ],
      [nftMintMessage],
    ),
  );
}

export function hashNFTMeldMessage(nftMeldMessage: NFTMeldMessage): string {
  return utils.keccak256(
    utils.defaultAbiCoder.encode(
      [
        {
          components: [
            {
              name: 'token_ids',
              type: 'uint256[5]',
            },
            {
              components: [
                {
                  name: 'is_common',
                  type: 'bool',
                },
                {
                  name: 'token_uri',
                  type: 'string',
                },
              ],
              name: 'melded_nft_data',
              type: 'tuple',
            },
            {
              name: 'beneficiary',
              type: 'address',
            },
            {
              name: 'expire',
              type: 'uint64',
            },
            {
              name: 'chain_id',
              type: 'uint32',
            },
          ],
          type: 'tuple',
        } as ParamType,
      ],
      [nftMeldMessage],
    ),
  );
}

export function convertToBountyHash(uuid: string, bountyID: string): string {
  const valueToConvert = `${uuid}:${bountyID}`;
  return utils.solidityKeccak256(['string'], [valueToConvert]);
}

export function convertToMintNFTHash(uuid: string, tokenURIIPFS: string): string {
  const valueToConvert = `${uuid}:${tokenURIIPFS}`;
  return utils.solidityKeccak256(['string'], [valueToConvert]);
}

export function packedMessage(message: string): string {
  return utils.solidityKeccak256(
    ['string', 'bytes32'],
    // eslint-disable-next-line unicorn/no-hex-escape
    ['\x19Ethereum Signed Message:\n32', message],
  );
}

export function randomHex(size: number): string {
  return Array.from({ length: size }, () => Math.floor(Math.random() * 16).toString(16)).join('');
}

export function randomBytes32(): string {
  return '0x' + randomHex(64);
}
