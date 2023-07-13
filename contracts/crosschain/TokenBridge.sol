// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import '@openzeppelin/contracts/token/ERC20/IERC20.sol';
import '@openzeppelin/contracts/access/AccessControl.sol';
import '@layerzerolabs/solidity-examples/contracts/lzApp/NonblockingLzApp.sol';

import '../interfaces/ITokenBridge.sol';
import '../interfaces/IERC20MintableBurnable.sol';

contract TokenBridge is ITokenBridge, NonblockingLzApp, AccessControl {
	bytes32 public constant BRIDGER_ROLE = keccak256('BRIDGER_ROLE');

	struct TokenConfig {
		bool isSupported;
		bool isRoot;
		mapping(uint16 => address) dstTokenLookup;
	}

	mapping(address => TokenConfig) public tokensLookup;

	constructor(address endpoint) NonblockingLzApp(endpoint) {
		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(BRIDGER_ROLE, msg.sender);
	}

	// -------- Public / external --------

	function addToken(address token, bool isRoot) external onlyRole(DEFAULT_ADMIN_ROLE) {
		TokenConfig storage tokenConfig = tokensLookup[token];

		if (tokenConfig.isSupported) revert TokenAlreadySupported(token);

		tokenConfig.isSupported = true;
		tokenConfig.isRoot = isRoot;
	}

	function removeToken(address token) external onlyRole(DEFAULT_ADMIN_ROLE) {
		if (!tokensLookup[token].isSupported) revert TokenNotSupported(token);

		delete tokensLookup[token];
	}

	function setDstToken(
		address token,
		uint16 dstChainId,
		address dstToken
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		TokenConfig storage tokenConfig = tokensLookup[token];

		if (!tokenConfig.isSupported) revert TokenNotSupported(token);

		tokenConfig.dstTokenLookup[dstChainId] = dstToken;
	}

	// does not check whether bridging is possible for supplied parameters
	function estimateFees(
		uint16 dstChainId,
		address token,
		address receiver,
		uint256 amount,
		bool payInZRO,
		bytes calldata adapterParams
	) public view returns (uint nativeFee, uint zroFee) {
		return
			lzEndpoint.estimateFees(
				dstChainId,
				address(this),
				abi.encode(token, receiver, amount),
				payInZRO,
				adapterParams
			);
	}

	// NOTE: chainIds are proprietary to LayerZero protocol and can be found on their docs
	function bridge(
		uint16 dstChainId,
		address token,
		address receiver,
		uint256 amount,
		address zroPaymentAddress,
		bytes calldata adapterParams
	) external payable {
		TokenConfig storage tokenConfig = tokensLookup[token];

		if (!tokenConfig.isSupported) revert TokenNotSupported(token);

		address dstToken = tokenConfig.dstTokenLookup[dstChainId];
		if (dstToken == address(0)) revert NoDstToken(token, dstChainId);

		if (!_isAuthorizedForBridging(msg.sender, token))
			revert BridgingUnauthorized(msg.sender, token);

		IERC20MintableBurnable tokenContract = IERC20MintableBurnable(token);

		if (tokenConfig.isRoot) {
			tokenContract.transferFrom(msg.sender, address(this), amount);
		} else {
			tokenContract.burnFrom(msg.sender, amount);
		}

		_lzSend(
			dstChainId, // chainId
			abi.encode(dstToken, receiver, amount), // payload
			payable(msg.sender), // refundAddress
			zroPaymentAddress, // zroPaymentAddress
			adapterParams, // adapterParams
			msg.value // nativeFee
		);

		emit BridgeOut(
			dstChainId,
			lzEndpoint.getOutboundNonce(dstChainId, address(this)),
			token,
			msg.sender,
			amount
		);
	}

	// -------- Internal --------

	// non-blocking logic override
	function _nonblockingLzReceive(
		uint16 srcChainId,
		bytes memory, // srcAddress
		uint64 nonce,
		bytes memory payload
	) internal override {
		(address token, address receiver, uint256 amount) = abi.decode(
			payload,
			(address, address, uint256)
		);

		TokenConfig storage tokenConfig = tokensLookup[token];

		if (!tokenConfig.isSupported) revert TokenNotSupported(token);

		if (!_isAuthorizedForBridging(receiver, token))
			revert BridgingUnauthorized(receiver, token);

		IERC20MintableBurnable tokenContract = IERC20MintableBurnable(token);

		if (tokenConfig.isRoot) {
			// NOTE: Bridge should have enough tokens as the only ability for token to appear on other chains is to be transferred to the bridge
			tokenContract.transfer(receiver, amount);
		} else {
			tokenContract.mint(receiver, amount);
		}

		emit BridgeIn(srcChainId, nonce, token, receiver, amount);
	}

	function _isAuthorizedForBridging(
		address _address,
		address token
	) internal view returns (bool) {
		return hasRole(BRIDGER_ROLE, _address) || hasRole(keccak256(abi.encode(token)), _address);
	}
}
