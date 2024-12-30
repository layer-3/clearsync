// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Test, console} from 'forge-std/Test.sol';

import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';

import {TradingApp} from '../../contracts/clearing/TradingApp.sol';
import {ITradingTypes} from '../../contracts/interfaces/ITradingTypes.sol';
import {INitroTypes} from '../../contracts/nitro/interfaces/INitroTypes.sol';

contract TradingAppTest_stateIsSupported is Test {
	TradingApp public tradingApp;
	INitroTypes.FixedPart public fixedPart;
	address traderAddress = vm.createWallet('trader').addr;
	address brokerAddress = vm.createWallet('broker').addr;

	function setUp() public {
		tradingApp = new TradingApp();
		address[] memory participants = new address[](2);
		participants[0] = traderAddress;
		participants[1] = brokerAddress;

		fixedPart = INitroTypes.FixedPart({
			participants: participants,
			channelNonce: 42,
			appDefinition: address(tradingApp),
			challengeDuration: 42
		});
	}

	function createRVP(
		bytes memory appData,
		uint48 turnNum,
		bool isFinal,
		uint8[] memory signedByIndices
	) public pure returns (INitroTypes.RecoveredVariablePart memory) {
		Outcome.SingleAssetExit[] memory outcome = new Outcome.SingleAssetExit[](0);
		INitroTypes.VariablePart memory variablePart = INitroTypes.VariablePart({
			outcome: outcome,
			appData: appData,
			turnNum: turnNum,
			isFinal: isFinal
		});
		uint256 signedBy;
		for (uint8 i = 0; i < signedByIndices.length; i++) {
			signedBy += 2 ** signedByIndices[i];
		}
		return INitroTypes.RecoveredVariablePart({variablePart: variablePart, signedBy: signedBy});
	}

	function newUint8_1(uint8 num) public pure returns (uint8[] memory) {
		uint8[] memory arr = new uint8[](1);
		arr[0] = num;
		return arr;
	}

	function newUint8_2(uint8 num1, uint8 num2) public pure returns (uint8[] memory) {
		uint8[] memory arr = new uint8[](2);
		arr[0] = num1;
		arr[1] = num2;
		return arr;
	}

	function test_supported_firstOrder() public view {
		INitroTypes.RecoveredVariablePart[] memory proof = new INitroTypes.RecoveredVariablePart[](
			1
		);
		// postfund
		proof[0] = createRVP(new bytes(0), 1, false, newUint8_2(0, 1));

		INitroTypes.RecoveredVariablePart memory candidate = createRVP(
			abi.encode(ITradingTypes.Order({orderID: bytes32('order1')})),
			2,
			false,
			newUint8_1(0)
		);

		(bool supported, string memory reason) = tradingApp.stateIsSupported(
			fixedPart,
			proof,
			candidate
		);
		assertTrue(supported);
		assertEq(reason, '');
	}

	function test_supported_orderResponsePair() public view {
		INitroTypes.RecoveredVariablePart[] memory proof = new INitroTypes.RecoveredVariablePart[](
			1
		);
		// order
		proof[0] = createRVP(
			abi.encode(ITradingTypes.Order({orderID: bytes32('order1')})),
			2,
			false,
			newUint8_1(0)
		);

		INitroTypes.RecoveredVariablePart memory candidate = createRVP(
			abi.encode(
				ITradingTypes.OrderResponse({
					orderID: bytes32('order1'),
					responseType: ITradingTypes.OrderResponseType.ACCEPT
				})
			),
			3,
			false,
			newUint8_1(1)
		);

		(bool supported, string memory reason) = tradingApp.stateIsSupported(
			fixedPart,
			proof,
			candidate
		);
		assertTrue(supported);
		assertEq(reason, '');
	}

	function test_supported_secondOrder() public view {
		INitroTypes.RecoveredVariablePart[] memory proof = new INitroTypes.RecoveredVariablePart[](
			1
		);
		// order
		proof[0] = createRVP(
			abi.encode(
				ITradingTypes.OrderResponse({
					orderID: bytes32('order1'),
					responseType: ITradingTypes.OrderResponseType.ACCEPT
				})
			),
			3,
			false,
			newUint8_1(1)
		);

		INitroTypes.RecoveredVariablePart memory candidate = createRVP(
			abi.encode(ITradingTypes.Order({orderID: bytes32('order2')})),
			4,
			false,
			newUint8_1(0)
		);

		(bool supported, string memory reason) = tradingApp.stateIsSupported(
			fixedPart,
			proof,
			candidate
		);
		assertTrue(supported);
		assertEq(reason, '');
	}

	function test_supported_liquidation() public view {
		ITradingTypes.Order memory order1 = ITradingTypes.Order({orderID: bytes32('order1')});
		ITradingTypes.OrderResponse memory response1 = ITradingTypes.OrderResponse({
			orderID: bytes32('order1'),
			responseType: ITradingTypes.OrderResponseType.ACCEPT
		});

		INitroTypes.RecoveredVariablePart[] memory proof = new INitroTypes.RecoveredVariablePart[](
			2
		);
		proof[0] = createRVP(abi.encode(order1), 2, false, newUint8_1(0));
		proof[1] = createRVP(abi.encode(response1), 3, false, newUint8_1(1));

		INitroTypes.RecoveredVariablePart memory candidate = createRVP(
			new bytes(0),
			4,
			false,
			newUint8_1(1)
		);

		// console.log(NitroUtils.getClaimedSignersNum(candidate.signedBy));

		(bool supported, string memory reason) = tradingApp.stateIsSupported(
			fixedPart,
			proof,
			candidate
		);
		assertTrue(supported);
		assertEq(reason, '');
	}

	function test_supported_settlement() public view {
		ITradingTypes.Order memory order1 = ITradingTypes.Order({orderID: bytes32('order1')});
		ITradingTypes.OrderResponse memory response1 = ITradingTypes.OrderResponse({
			orderID: bytes32('order1'),
			responseType: ITradingTypes.OrderResponseType.ACCEPT
		});
		ITradingTypes.Order memory order2 = ITradingTypes.Order({orderID: bytes32('order2')});
		ITradingTypes.OrderResponse memory response2 = ITradingTypes.OrderResponse({
			orderID: bytes32('order2'),
			responseType: ITradingTypes.OrderResponseType.ACCEPT
		});

		ITradingTypes.AssetAndAmount[] memory toTrader = new ITradingTypes.AssetAndAmount[](2);
		toTrader[0] = ITradingTypes.AssetAndAmount({asset: address(42), amount: 1});
		toTrader[1] = ITradingTypes.AssetAndAmount({asset: address(43), amount: 2});

		ITradingTypes.AssetAndAmount[] memory toBroker = new ITradingTypes.AssetAndAmount[](2);
		toBroker[0] = ITradingTypes.AssetAndAmount({asset: address(44), amount: 3});
		toBroker[1] = ITradingTypes.AssetAndAmount({asset: address(45), amount: 4});

		// NOTE: dynamic array as it is used in TradingApp
		bytes32[] memory proofDataHashes = new bytes32[](4);
		proofDataHashes[0] = keccak256(abi.encode(order1));
		proofDataHashes[1] = keccak256(abi.encode(response1));
		proofDataHashes[2] = keccak256(abi.encode(order2));
		proofDataHashes[3] = keccak256(abi.encode(response2));

		bytes32 checksum = keccak256(abi.encode(proofDataHashes));
		ITradingTypes.Settlement memory settlement = ITradingTypes.Settlement({
			toTrader: toTrader,
			toBroker: toBroker,
			ordersChecksum: checksum
		});

		INitroTypes.RecoveredVariablePart[] memory proof = new INitroTypes.RecoveredVariablePart[](
			4
		);
		proof[0] = createRVP(abi.encode(order1), 2, false, newUint8_1(0));
		proof[1] = createRVP(abi.encode(response1), 3, false, newUint8_1(1));
		proof[2] = createRVP(abi.encode(order2), 4, false, newUint8_1(0));
		proof[3] = createRVP(abi.encode(response2), 5, false, newUint8_1(1));

		INitroTypes.RecoveredVariablePart memory candidate = createRVP(
			abi.encode(settlement),
			6,
			false,
			newUint8_2(0, 1)
		);

		(bool supported, string memory reason) = tradingApp.stateIsSupported(
			fixedPart,
			proof,
			candidate
		);
		assertTrue(supported);
		assertEq(reason, '');
	}
}
