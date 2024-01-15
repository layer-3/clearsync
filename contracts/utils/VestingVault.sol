// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

/**
 * @title VestingVault
 * @notice A linear token vesting contract that supports multiple vesting schedules for each beneficiary.
 */
contract VestingVault is Ownable {
	// The vesting schedule structure
	struct Schedule {
		uint256 amount;
		uint256 releasedAmount;
		uint64 start;
		uint64 duration;
	}

	// ====== Variables ======

	// The ERC20 token being vested
	IERC20 public token;
	// Mapping of beneficiary address to an array of vesting schedules
	mapping(address => Schedule[]) internal _beneficiarySchedules;

	// ====== Events ======

	/**
	 * @notice Emitted when a new vesting schedule is added.
	 * @dev Emitted when a new vesting schedule is added.
	 * @param beneficiary The address of the beneficiary.
	 * @param amount The total amount of tokens to be vested.
	 * @param start The start timestamp for the vesting schedule.
	 * @param duration The duration of the vesting period in seconds.
	 */
	event ScheduleAdded(
		address indexed beneficiary,
		uint256 amount,
		uint256 start,
		uint256 duration
	);

	/**
	 * @notice Emitted when a vesting schedule is deleted.
	 * @dev Emitted when a vesting schedule is deleted.
	 * @param beneficiary The address of the beneficiary.
	 * @param index The index of the vesting schedule that was deleted.
	 */
	event ScheduleDeleted(address indexed beneficiary, uint256 index);

	/**
	 * @notice Emitted when tokens are claimed by a beneficiary.
	 * @dev Emitted when tokens are claimed by a beneficiary.
	 * @param beneficiary The address of the beneficiary.
	 * @param amount The amount of tokens claimed.
	 */
	event TokensClaimed(address indexed beneficiary, uint256 amount);

	// ====== Errors ======

	/**
	 * @dev Is used when token address is zero.
	 */
	error InvalidTokenAddress(address tokenAddress);

	/**
	 * @dev Is used when a vesting schedule does not comply with the rules.
	 */
	error InvalidSchedule(Schedule schedule);

	/**
	 * @dev Is used when a beneficiary does not have a vesting schedule at the given index.
	 */
	error NoScheduleForBeneficiary(address beneficiary, uint256 index);

	/**
	 * @dev Is used when a beneficiary tries to claim tokens but has no vesting schedules
	 * or no tokens have been vested since last claim.
	 */
	error UnableToClaim(address beneficiary);

	// ====== Constructor ======

	/**
	 * @notice Initializes the contract with the given ERC20 token.
	 * @dev Initializes the contract with the given ERC20 token.
	 * @param token_ The address of the ERC20 token.
	 */
	constructor(IERC20 token_) Ownable(msg.sender) {
		if (address(token_) == address(0)) revert InvalidTokenAddress(address(token_));
		token = token_;
	}

	// ====== Owner functions ======

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
	 * @dev Deletes a vesting schedule for a `beneficiary`.
	 * Can only be called by the contract owner.
	 * @param beneficiary The address of the beneficiary.
	 * @param index The index of the vesting schedule to be deleted.
	 */
	function deleteSchedule(address beneficiary, uint256 index) public onlyOwner {
		_deleteSchedule(beneficiary, index);
	}

	// ====== View functions ======

	/**
	 * @notice Returns the vesting schedules for a `beneficiary`.
	 * @dev Returns the vesting schedules for a `beneficiary`.
	 * @param beneficiary The address of the beneficiary.
	 * @return The vesting schedules for the beneficiary.
	 */
	function beneficiarySchedules(address beneficiary) public view returns (Schedule[] memory) {
		return _beneficiarySchedules[beneficiary];
	}

	/**
	 * @notice Returns the vesting schedule for a `beneficiary` at `index`.
	 * @dev Reverts if the `beneficiary` does not have a vesting schedule at `index`.
	 * @param beneficiary The address of the beneficiary.
	 * @param index The index of the vesting schedule.
	 * @return The vesting schedule for the beneficiary at the given index.
	 */
	function beneficiarySchedule(
		address beneficiary,
		uint256 index
	) public view returns (Schedule memory) {
		Schedule[] memory schedules = _beneficiarySchedules[beneficiary];
		if (index >= schedules.length) revert NoScheduleForBeneficiary(beneficiary, index);
		return schedules[index];
	}

	/**
	 * @notice Amount of token that can be claimed from a vesting schedule at `index` of `beneficiary`.
	 * @dev Reverts if the `beneficiary` does not have a vesting schedule at `index`.
	 * @param beneficiary The address of the beneficiary.
	 * @param index The index of the vesting schedule.
	 * @return claimable The amount of token that can be claimed from a vesting schedule at `index` of `beneficiary`.
	 */
	function scheduleClaimable(
		address beneficiary,
		uint256 index
	) public view returns (uint256 claimable) {
		if (index >= _beneficiarySchedules[beneficiary].length)
			revert NoScheduleForBeneficiary(beneficiary, index);

		(claimable, ) = _scheduleClaimable(_beneficiarySchedules[beneficiary][index]);
	}

	/**
	 * @notice Total amount of token that can be claimed by a `beneficiary`.
	 * @dev Total amount of token that can be claimed by a `beneficiary`.
	 * @param beneficiary The address of the beneficiary.
	 * @return claimable The total amount of token that can be claimed by a `beneficiary`.
	 */
	function totalClaimable(address beneficiary) public view returns (uint256 claimable) {
		Schedule[] memory schedules = _beneficiarySchedules[beneficiary];

		for (uint256 i = 0; i < schedules.length; i++) {
			(uint256 scheduleClaimable_, ) = _scheduleClaimable(schedules[i]);
			claimable += scheduleClaimable_;
		}
	}

	// ====== Modifying functions ======

	/**
	 * @notice Releases vested tokens for the calling beneficiary.
	 * @dev Reverts if the beneficiary does not have any vesting schedules or no tokens have been vested since last claim.
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

			(uint256 claimable, bool fullyPaid) = _scheduleClaimable(schedule);

			if (claimable > 0) {
				schedule.releasedAmount += claimable;
				totalUnreleasedAmount += claimable;
			}

			if (fullyPaid) {
				fullyPaidScheduleIndices[numberOfFullyPaidSchedules] = i;
				numberOfFullyPaidSchedules++;
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

	// ====== Internal functions ======

	/**
	 * @notice Deletes a vesting schedule for a `beneficiary` at `index`. Internal function.
	 * @dev Reverts if the `beneficiary` does not have a vesting schedule at `index`.
	 * @param beneficiary The address of the beneficiary.
	 * @param index The index of the vesting schedule to be deleted.
	 */
	function _deleteSchedule(address beneficiary, uint256 index) internal {
		Schedule[] storage schedules = _beneficiarySchedules[beneficiary];
		if (index >= schedules.length) revert NoScheduleForBeneficiary(beneficiary, index);

		schedules[index] = schedules[schedules.length - 1];
		schedules.pop();

		emit ScheduleDeleted(beneficiary, index);
	}

	/**
	 * @notice Calculates the amount of token that can be claimed from a vesting schedule and if the vesting schedule is fully paid.
	 * @dev Calculates the amount of token that can be claimed from a vesting schedule and if the vesting schedule is fully paid.
	 * @param schedule The vesting schedule.
	 * @return claimable The amount of token that can be claimed from a vesting schedule.
	 * @return fullyPaid True if the vesting schedule is fully paid.
	 */
	function _scheduleClaimable(
		Schedule memory schedule
	) internal view returns (uint256 claimable, bool fullyPaid) {
		if (block.timestamp < schedule.start) return (0, false);

		uint256 elapsedTime = block.timestamp - schedule.start;
		uint256 vestedAmount = 0;
		if (elapsedTime >= schedule.duration) {
			vestedAmount = schedule.amount;
			fullyPaid = true;
		} else {
			vestedAmount = (schedule.amount * elapsedTime) / schedule.duration;
		}

		claimable = vestedAmount - schedule.releasedAmount;
	}
}
