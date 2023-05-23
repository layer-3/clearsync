// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

/**
 * @title VestingVault
 * @dev A token vesting contract that supports multiple vesting schedules for each beneficiary.
 */
contract VestingVault is Ownable {
	// The vesting schedule structure
	struct Schedule {
		uint256 amount;
		uint256 start;
		uint256 duration;
		uint256 released;
	}

	// The ERC20 token being vested
	IERC20 public token;
	// Mapping of beneficiary address to an array of vesting schedules
	mapping(address => Schedule[]) public beneficiarySchedules;

	// Events
	event ScheduleAdded(
		address indexed beneficiary,
		uint256 amount,
		uint256 start,
		uint256 duration
	);
	event ScheduleDeleted(address indexed beneficiary, uint256 index);
	event TokensReleased(address indexed beneficiary, uint256 amount);

	error InvalidTokenAddress(address tokenAddress);
	error InvalidSchedule(Schedule schedule);
	error NoScheduleForBeneficiary(address beneficiary, uint256 index);
	error UnableToRelease(address beneficiary);

	/**
	 * @dev Initializes the contract with the given ERC20 token.
	 * @param _token The address of the ERC20 token.
	 */
	constructor(IERC20 _token) {
		if (address(_token) == address(0)) revert InvalidTokenAddress(address(_token));
		token = _token;
	}

	/**
	 * @dev Adds a vesting schedule for a beneficiary.
	 * Can only be called by the contract owner.
	 * @param _beneficiary The address of the beneficiary.
	 * @param _amount The total amount of tokens to be vested.
	 * @param _start The start timestamp for the vesting schedule.
	 * @param _duration The duration of the vesting period in seconds.
	 */
	function addSchedule(
		address _beneficiary,
		uint256 _amount,
		uint256 _start,
		uint256 _duration
	) public onlyOwner {
		Schedule memory newSchedule = Schedule({
			amount: _amount,
			start: _start,
			duration: _duration,
			released: 0
		});

		if (
			_beneficiary == address(0) ||
			_amount == 0 ||
			_start <= block.timestamp ||
			_duration == 0
		) revert InvalidSchedule(newSchedule);

		beneficiarySchedules[_beneficiary].push(newSchedule);

		emit ScheduleAdded(_beneficiary, _amount, _start, _duration);
	}

	/**
	 * @dev Deletes a vesting schedule for a beneficiary.
	 * Can only be called by the contract owner.
	 * @param _beneficiary The address of the beneficiary.
	 * @param _index The index of the vesting schedule to be deleted.
	 */
	function deleteSchedule(address _beneficiary, uint256 _index) public onlyOwner {
		Schedule[] storage schedules = beneficiarySchedules[_beneficiary];
		if (_index >= schedules.length) revert NoScheduleForBeneficiary(_beneficiary, _index);

		schedules[_index] = schedules[schedules.length - 1];
		schedules.pop();

		emit ScheduleDeleted(_beneficiary, _index);
	}

	/**
	 * @dev Releases vested tokens for the calling beneficiary.
	 * Can only be called by a beneficiary.
	 */
	function claim() public {
		uint256 totalUnreleasedAmount = 0;
		Schedule[] storage schedules = beneficiarySchedules[msg.sender];

		if (schedules.length == 0) revert UnableToRelease(msg.sender);

		for (uint256 i = 0; i < schedules.length; i++) {
			Schedule storage schedule = schedules[i];

			uint256 elapsedTime = block.timestamp - schedule.start;
			uint256 vestedAmount = (schedule.amount * elapsedTime) / schedule.duration;
			uint256 unreleasedAmount = vestedAmount - schedule.released;

			if (unreleasedAmount > 0) {
				schedule.released += unreleasedAmount;
				totalUnreleasedAmount += unreleasedAmount;
			}
		}

		if (totalUnreleasedAmount == 0) revert UnableToRelease(msg.sender);

		token.transfer(msg.sender, totalUnreleasedAmount);

		emit TokensReleased(msg.sender, totalUnreleasedAmount);
	}
}
