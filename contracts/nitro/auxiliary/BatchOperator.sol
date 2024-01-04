// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import {MultiAssetHolder} from '../MultiAssetHolder.sol';

import {SafeERC20} from '@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol';
import {IERC20} from '@openzeppelin/contracts/token/ERC20/IERC20.sol';

string constant lengthsErr = 'Array lengths must match';

/**
@dev This contract is a proxy for the NitroAdjudicator, providing endpoints for batched operations.
 */
contract BatchOperator {
    using SafeERC20 for IERC20;
    MultiAssetHolder public adjudicator;

    constructor(address adjudicatorAddr) {
        adjudicator = MultiAssetHolder(adjudicatorAddr);
    }

    /**
     * @dev Deposits ETH (or native token for other runtime) into the adjudicator for multiple channels.
     */
    function deposit_batch_eth(
        bytes32[] calldata channelIds,
        uint256[] calldata expectedHelds,
        uint256[] calldata amounts
    ) external payable virtual {
        require(
            channelIds.length == expectedHelds.length && expectedHelds.length == amounts.length,
            lengthsErr
        );
        for (uint256 i = 0; i < channelIds.length; i++) {
            adjudicator.deposit{value: amounts[i]}(
                address(0),
                channelIds[i],
                expectedHelds[i],
                amounts[i]
            );
        }
    }

    /**
     * @dev Deposits ERC20 tokens into the adjudicator for multiple channels.
     */
    function deposit_batch_erc20(
        address asset,
        bytes32[] calldata channelIds,
        uint256[] calldata expectedHelds,
        uint256[] calldata amounts,
        uint256 totalAmount
    ) external payable virtual {
        require(
            channelIds.length == expectedHelds.length && expectedHelds.length == amounts.length,
            lengthsErr
        );

        IERC20(asset).safeTransferFrom(msg.sender, address(this), totalAmount);
        IERC20(asset).safeApprove(address(adjudicator), totalAmount);

        for (uint256 i = 0; i < channelIds.length; i++) {
            adjudicator.deposit(asset, channelIds[i], expectedHelds[i], amounts[i]);
        }
    }
}
