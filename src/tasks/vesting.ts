import '@nomicfoundation/hardhat-toolbox';
import { readFileSync } from 'node:fs';
import { join } from 'node:path';

import { task } from 'hardhat/config';

import type { VestingVault } from '../../typechain-types';

interface AddScheduleArgs {
  contract: string;
  beneficiary: string;
  amount: number;
  start: number;
  duration: number;
}

task('addSchedule', 'Add a vesting schedule to the Vesting contract')
  .addParam('contract', 'The address of Vesting contract')
  .addParam('beneficiary', 'The address to vest tokens to')
  .addParam('amount', 'Amount of tokens to vest')
  .addParam('start', 'Start time (UNIX timestamp in seconds) of the vesting schedule')
  .addParam('duration', 'Duration of the vesting schedule (in seconds)')
  .setAction(async (taskArgs: AddScheduleArgs, { ethers }) => {
    const Vesting = (await ethers.getContractAt('VestingVault', taskArgs.contract)) as VestingVault;

    await Vesting.addSchedule(
      taskArgs.beneficiary,
      taskArgs.amount,
      taskArgs.start,
      taskArgs.duration,
    );

    console.log(
      `Vesting: Added schedule for ${taskArgs.beneficiary} {amount: ${taskArgs.amount}, start: ${taskArgs.start}, duration: ${taskArgs.duration}}`,
    );
  });

interface AddSchedulesArgs {
  contract: string;
  src: string;
}

task('addSchedules', 'Read from the file and add vesting schedules to the Vesting contract')
  .addParam('contract', 'The address of Vesting contract')
  .addParam('src', 'Location of a .csv file with vesting schedules')
  .setAction(async (taskArgs: AddSchedulesArgs, { ethers }) => {
    const Vesting = (await ethers.getContractAt('VestingVault', taskArgs.contract)) as VestingVault;

    const filePath = join(process.cwd(), taskArgs.src);
    const rawFileContent = readFileSync(filePath, 'utf8');
    const fileContent = rawFileContent.trim();
    const lines = fileContent.split('\n');

    console.log(`Vesting: Adding ${lines.length} schedules...\n`);

    for (const [idx, line] of lines.entries()) {
      const [beneficiary, amount, start, duration] = line.split(',');
      await Vesting.addSchedule(
        beneficiary,
        ethers.utils.parseEther(amount),
        Number(start),
        Number(duration),
      );
      console.log(`[${idx + 1}/${lines.length} done]`);
      console.log(
        `Added schedule for ${beneficiary} {amount: ${amount}, start: ${start}, duration: ${duration}}\n`,
      );
    }

    console.log(`Vesting: Added ${lines.length} schedules`);
  });

interface DeleteScheduleArgs {
  contract: string;
  beneficiary: string;
  index: number;
}

task('deleteSchedule', 'Delete a vesting schedule at a specified index from the Vesting contract')
  .addParam('contract', 'The address of Vesting contract')
  .addParam('beneficiary', 'The address to vest tokens to')
  .addParam('index', 'Index of the schedule to delete')
  .setAction(async (taskArgs: DeleteScheduleArgs, { ethers }) => {
    const Vesting = (await ethers.getContractAt('VestingVault', taskArgs.contract)) as VestingVault;

    await Vesting.deleteSchedule(taskArgs.beneficiary, taskArgs.index);

    console.log(
      `Vesting: schedule for ${taskArgs.beneficiary} at index ${taskArgs.index} was deleted`,
    );
  });
