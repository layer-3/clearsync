// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

library Flags {
	uint8 public constant IS_TRANSFERABLE_FLAG = 0;

	function raiseFlag(uint8 self, uint8 flag) internal pure returns (uint8) {
		return uint8(self | (1 << flag));
	}

	function clearFlag(uint8 self, uint8 flag) internal pure returns (uint8) {
		return uint8(self & ~(1 << flag));
	}

	function setFlag(uint8 self, uint8 flag, bool value) internal pure returns (uint8) {
		if (value) {
			return raiseFlag(self, flag);
		} else {
			return clearFlag(self, flag);
		}
	}

	function getFlag(uint8 flags, uint8 flag) internal pure returns (bool) {
		return flags & (1 << flag) > 0;
	}
}
