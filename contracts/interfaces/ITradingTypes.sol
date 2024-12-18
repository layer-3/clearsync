// SPDX-License-Identifier: MIT
pragma solidity 0.8.22;


interface ITradingTypes {
	struct Order {
		bytes32 orderID;
	}

	enum OrderResponseType {
		ACCEPT,
		REJECT
	}

	struct OrderResponse {
		OrderResponseType responseType;
		bytes32 orderID; // orderID making the trade
	}

	struct AssetAndAmount {
		address asset;
		uint256 amount;
	}

	struct Settlement {
		AssetAndAmount[] toTrader;
		AssetAndAmount[] toBroker;
	}
}
