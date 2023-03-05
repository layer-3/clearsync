import '@nomicfoundation/hardhat-toolbox';
import { task } from 'hardhat/config';

import type { YellowToken } from '../../typechain-types';

interface ActivateTokenArgs {
  tokenAddress: string;
  premint: number;
  premintTo: string;
}

task('activate', 'Activates Token (Canary and Yellow)')
  .addParam('tokenAddress', 'The address of Token')
  .addParam('premint', 'Amount of Tokens to premint')
  .addParam('premintTo', 'Address of an account to premint Tokens to')
  .setAction(async (taskArgs: ActivateTokenArgs, { ethers }) => {
    const Token = (await ethers.getContractAt('YellowToken', taskArgs.tokenAddress)) as YellowToken;
    await Token.activate(taskArgs.premint, taskArgs.premintTo);

    console.log(
      `Token was activated with \`premint\`=${taskArgs.premint} and \`account\`=${taskArgs.premintTo}`,
    );
  });
