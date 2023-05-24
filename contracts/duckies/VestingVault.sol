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
		uint256 releasedAmount;
		uint64 start;
		uint64 duration;
	}

	// The ERC20 token being vested
	IERC20 public token;
	// Mapping of beneficiary address to an array of vesting schedules
	mapping(address => Schedule[]) internal _beneficiarySchedules;

	// Events
	event ScheduleAdded(
		address indexed beneficiary,
		uint256 amount,
		uint256 start,
		uint256 duration
	);
	event ScheduleDeleted(address indexed beneficiary, uint256 index);
	event TokensClaimed(address indexed beneficiary, uint256 amount);

	error InvalidTokenAddress(address tokenAddress);
	error InvalidSchedule(Schedule schedule);
	error NoScheduleForBeneficiary(address beneficiary, uint256 index);
	error UnableToClaim(address beneficiary);

	/**
	 * @dev Initializes the contract with the given ERC20 token.
	 * @param token_ The address of the ERC20 token.
	 */
	constructor(IERC20 token_) {
		if (address(token_) == address(0)) revert InvalidTokenAddress(address(token_));
		token = token_;
	}

	function beneficiarySchedules(address beneficiary) public view returns (Schedule[] memory) {
		return _beneficiarySchedules[beneficiary];
	}

	function beneficiarySchedule(
		address beneficiary,
		uint256 index
	) public view returns (Schedule memory) {
		Schedule[] memory schedules = _beneficiarySchedules[beneficiary];
		if (index >= schedules.length) revert NoScheduleForBeneficiary(beneficiary, index);
		return schedules[index];
	}

	/**
	 * @dev Adds a vesting schedule for a beneficiary.
	 * Can only be called by the contract owner.
	 * @param beneficiary The address of the beneficiary.
	 * @param amount The total amount of tokens to be vested.
	 * @param start The start timestamp for the vesting schedule.
	 * @param duration The duration of the vesting period in seconds.
	 */
	function addSchedule(
		address beneficiary,
		uint256 amount,
		uint64 start,
		uint64 duration
	) public onlyOwner {
		Schedule memory newSchedule = Schedule({
			amount: amount,
			releasedAmount: 0,
			start: start,
			duration: duration
		});

		if (beneficiary == address(0) || amount == 0 || start <= block.timestamp || duration == 0)
			revert InvalidSchedule(newSchedule);

		_beneficiarySchedules[beneficiary].push(newSchedule);

		emit ScheduleAdded(beneficiary, amount, start, duration);
	}

	/**
	 * @dev Deletes a vesting schedule for a beneficiary.
	 * Can only be called by the contract owner.
	 * @param beneficiary The address of the beneficiary.
	 * @param index The index of the vesting schedule to be deleted.
	 */
	function deleteSchedule(address beneficiary, uint256 index) public onlyOwner {
		_deleteSchedule(beneficiary, index);
	}

	function _deleteSchedule(address beneficiary, uint256 index) internal {
		Schedule[] storage schedules = _beneficiarySchedules[beneficiary];
		if (index >= schedules.length) revert NoScheduleForBeneficiary(beneficiary, index);

		schedules[index] = schedules[schedules.length - 1];
		schedules.pop();

		emit ScheduleDeleted(beneficiary, index);
	}

	/**
	 * @dev Releases vested tokens for the calling beneficiary.
	 * Can only be called by a beneficiary.
	 */
	function claim() public {
		uint256 totalUnreleasedAmount = 0;
		Schedule[] storage schedules = _beneficiarySchedules[msg.sender];

		if (schedules.length == 0) revert UnableToClaim(msg.sender);

		// amount of fully paid schedules is not known beforehand, but it's always <= schedules.length
		uint256[] memory fullyPaidScheduleIndices = new uint256[](schedules.length);
		uint256 numberOfFullyPaidSchedules = 0;

		for (uint256 i = 0; i < schedules.length; i++) {
			Schedule storage schedule = schedules[i];

			if (schedule.start > block.timestamp) continue;

			uint256 elapsedTime = block.timestamp - schedule.start;
			uint256 vestedAmount = 0;
			if (elapsedTime >= schedule.duration) {
				vestedAmount = schedule.amount;
				fullyPaidScheduleIndices[numberOfFullyPaidSchedules] = i;
				numberOfFullyPaidSchedules++;
			} else {
				vestedAmount = (schedule.amount * elapsedTime) / schedule.duration;
			}

			uint256 unreleasedAmount = vestedAmount - schedule.releasedAmount;

			if (unreleasedAmount > 0) {
				schedule.releasedAmount = vestedAmount;
				totalUnreleasedAmount += unreleasedAmount;
			}
		}

		if (totalUnreleasedAmount == 0) revert UnableToClaim(msg.sender);

		// delete fully paid schedules
		// traverse indices in descending order to avoid shifting indices when deleting
		for (uint256 i = numberOfFullyPaidSchedules; i > 0; i--) {
			_deleteSchedule(msg.sender, fullyPaidScheduleIndices[i - 1]);
		}

		token.transfer(msg.sender, totalUnreleasedAmount);

		emit TokensClaimed(msg.sender, totalUnreleasedAmount);
	}
}
