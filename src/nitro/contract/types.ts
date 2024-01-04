export type Address = string;

export type Byte = string; // 0x + val(length 4)
export type Bytes32 = string; // 0x + val(length 64)
export type Bytes = string;

// Ethersjs lets you pass, and returns a number, for solidity variables of
// The types uint8, uint16, and uint32
export type Uint8 = number;
export type Uint16 = number;
export type Uint24 = number;
export type Uint32 = number;
export type Uint40 = number;
export type Uint48 = number;
// These can only be safely stored as a hex string, which is the type that ethers returns
export type Uint56 = number;
export type Uint64 = string;
export type Uint128 = string;
export type Uint256 = string;
