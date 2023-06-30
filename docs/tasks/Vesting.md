# Vesting tasks

Before adding vesting schedules, you need to deploy the `Vesting` contract, see [Deploy a contract](../../scripts/README.md#deploy-any-contract) section.

## Add vesting schedule

To add a vesting schedule, you need to run the following command:

```bash
npx hardhat addSchedule --amount <amount> --beneficiary <beneficiary_address> --contract <vesting_address> --duration <vesting_duration_in_seconds> --start <start_timestamp_in_seconds> --network <network_name>
```

Note: amount should account for decimals in the Token, e.g. if a Token has 8 decimals, to create a schedule for 1000 tokens, you need to pass `1000_0000_0000`.

## Add multiple schedules

You can add multiple schedules, by running

```bash
npx hardhat addSchedules --contract <vesting_address> --src <path_to_csv_file> --network <network_name>
```

The `.csv` file should not have a header and have the following format:

```csv
beneficiary,amount,start,duration
```

Note: path to a `.csv` file is specified relatively to the root of this repo.

## Delete a schedule

To delete a vesting schedule, you need to run the following command:

```bash
npx hardhat deleteSchedule --beneficiary <beneficiary_address> --contract <vesting_address> --index <schedule_index_for_beneficiary> --network <network_name>
```
