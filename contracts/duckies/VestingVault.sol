// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title VestingVault
 * @dev A token vesting contract that supports multiple vesting schedules for each beneficiary.
 */
contract VestingVault is Ownable {
    using SafeMath for uint256;

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
    event ScheduleAdded(address indexed beneficiary, uint256 amount, uint256 start, uint256 duration);
    event ScheduleDeleted(address indexed beneficiary, uint256 index);
    event TokensReleased(address indexed beneficiary, uint256 amount);

    /**
     * @dev Initializes the contract with the given ERC20 token.
     * @param _token The address of the ERC20 token.
     */
    constructor(IERC20 _token) {
        require(address(_token) != address(0), "TokenVesting: invalid token address");
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
    function addSchedule(address _beneficiary, uint256 _amount, uint256 _start, uint256 _duration) public onlyOwner {
        require(_beneficiary != address(0), "TokenVesting: invalid beneficiary address");
        require(_amount > 0, "TokenVesting: amount must be greater than 0");
        require(_duration > 0, "TokenVesting: duration must be greater than 0");

        Schedule memory newSchedule = Schedule({
            amount: _amount,
            start: _start,
            duration: _duration,
            released: 0
        });

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
        require(_index < schedules.length, "TokenVesting: index out of range");

        schedules[_index] = schedules[schedules.length - 1];
        schedules.pop();

        emit ScheduleDeleted(_beneficiary, _index);
    }

    /**
     * @dev Releases vested tokens for the calling beneficiary.
     * Can only be called by a beneficiary.
     */
    function release() public {
        uint256 totalUnreleasedAmount = 0;
        Schedule[] storage schedules = beneficiarySchedules[msg.sender];

        require(schedules.length > 0, "TokenVesting: no schedules for the beneficiary");

        for (uint256 i = 0; i < schedules.length; i++) {
            Schedule storage schedule = schedules[i];

            uint256 elapsedTime = block.timestamp.sub(schedule.start);
            uint256 vestedAmount = schedule.amount.mul(elapsedTime).div(schedule.duration);
            uint256 unreleasedAmount = vestedAmount.sub(schedule.released);

            if (unreleasedAmount > 0) {
                schedule.released = schedule.released.add(unreleasedAmount);
                totalUnreleasedAmount = totalUnreleasedAmount.add(unreleasedAmount);
            }
        }

        require(totalUnreleasedAmount > 0, "TokenVesting: no tokens to release");

        token.transfer(msg.sender, totalUnreleasedAmount);

        emit TokensReleased(msg.sender, totalUnreleasedAmount);
    }
}
