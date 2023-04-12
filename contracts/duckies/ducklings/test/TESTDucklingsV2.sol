// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '../DucklingsV1.sol';

contract TESTDucklingsV2 is DucklingsV1 {
	function isV2() external pure returns (bool) {
		return true;
	}
}
