// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

interface IStatusManager {
	enum ChannelMode {
		Open,
		Challenge,
		Finalized
	}

	struct ChannelData {
		uint48 turnNumRecord;
		uint48 finalizesAt;
		bytes32 stateHash; // keccak256(abi.encode(State))
		bytes32 outcomeHash;
	}

	function unpackStatus(
		bytes32 channelId
	) external view returns (uint48 turnNumRecord, uint48 finalizesAt, uint160 fingerprint);
}
