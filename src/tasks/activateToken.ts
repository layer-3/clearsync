import { task } from 'hardhat/config';
import '@nomicfoundation/hardhat-toolbox';

import type { Token } from '../../typechain-types';

interface ActivateTokenArgs {
  tokenAddress: string;
  premint: number;
  premintTo: string;
}

task('activateToken', 'Activates Token (Canary and Yellow)')
  .addParam('tokenAddress', 'The address of Token')
  .addParam('premint', 'Amount of Tokens to premint')
  .addParam('premintTo', 'Address of an account to premint Tokens to')
  .setAction(async (taskArgs: ActivateTokenArgs, { ethers }) => {
    const Token = (await ethers.getContractAt('Token', taskArgs.tokenAddress)) as Token;
    await Token.activate(taskArgs.premint, taskArgs.premintTo);

    console.log(
      `Token was activated with \`premint\`=${taskArgs.premint} and \`account\`=${taskArgs.premintTo}`,
    );
  });
