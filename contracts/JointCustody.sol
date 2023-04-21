// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import '@openzeppelin/contracts/utils/Counters.sol';
import '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';
import './interfaces/ICustody.sol';
import './interfaces/IClearing.sol';
import './interfaces/ILegacyERC20.sol';
import './Quorum.sol';

/**
 * @notice Custody for brokers allowing users to deposit their funds to or withdraw from, and brokers to swap currencies according do SwapCall agreed.
 */
contract JointCustody is ICustody, IClearing, Quorum {
    using Counters for Counters.Counter;
    using ECDSA for bytes32;

    // how many deposits, withdrawals and swaps were performed
    Counters.Counter private _lastInteractionId;

    // broker => asset => amount
    mapping(address => mapping(address => uint256)) private _balances;

    mapping(bytes32 => bool) private _sigsUsed;

    // constructor
    constructor(address[] memory validators, uint8 quorum) Quorum(validators, quorum) {}

    function lastInteractionId() external view returns (uint256) {
        return _lastInteractionId.current();
    }

    // deposit
    function deposit(
        Payload calldata payload,
        bytes calldata brokerSignature,
        bytes[] calldata signatures
    ) external payable {
        address issuer = msg.sender;

        // check signatures
        _requireSigsNotUsed(signatures);

        bytes memory encodedPayload = abi.encode(payload);
        _requireValidSignature(payload.broker, encodedPayload, brokerSignature);
        requireQuorum(encodedPayload, signatures);

        // check payload
        require(payload.action == Actions.deposit, 'Invalid action');
        require(payload.account == issuer, 'Invalid destination');
        _requireValidPayload(payload);

        // use signature
        _useSignatures(signatures);

        // deposit tokens
        address broker = payload.broker;
        address asset = payload.asset;
        uint256 amount = payload.amount;

        if (asset == address(0)) {
            require(msg.value == amount, 'Incorrect msg.value');
        } else {
            ILegacyERC20(asset).transferFrom(issuer, address(this), amount);

            // protected from reentrancy by marking signatures as used
            require(_retrieveTransferResult(), 'Could not deposit ERC20');
        }

        _balances[broker][asset] += amount;

        emit Deposited(payload.nonce, broker, issuer, asset, amount);
    }

    function withdraw(
        Payload calldata payload,
        bytes calldata brokerSignature,
        bytes[] calldata signatures
    ) external payable {
        address issuer = msg.sender;

        // check signatures
        _requireSigsNotUsed(signatures);

        bytes memory encodedPayload = abi.encode(payload);
        _requireValidSignature(payload.broker, encodedPayload, brokerSignature);
        requireQuorum(encodedPayload, signatures);

        // check payload
        require(payload.action == Actions.withdraw, 'Invalid action');
        _requireValidPayload(payload);

        // use signature
        _useSignatures(signatures);

        // deposit tokens
        address broker = payload.broker;
        address asset = payload.asset;
        uint256 amount = payload.amount;

        if (asset == address(0)) {
            (bool success, ) = payload.account.call{value: amount}(''); //solhint-disable-line avoid-low-level-calls

            require(success, 'Could not transfer ETH');
        } else {
            ILegacyERC20(asset).transfer(payload.account, amount);

            // protected from reentrancy by marking signatures as used
            require(_retrieveTransferResult(), 'Could not withdraw ERC20');
        }

        _balances[broker][asset] -= amount;

        emit Withdrawn(payload.nonce, broker, issuer, asset, amount);
    }

    function swap(SignedSwapCall calldata sSC, bytes32 channelID) external payable {
        // check this swap has not been performed
        require(_sigsUsed[keccak256(abi.encode(sSC.swapCall))] == false, 'swap already performed');

        // check expire is < now and != 0
        require(sSC.swapCall.expire != 0, 'expire = 0');
        require(sSC.swapCall.expire > block.timestamp, 'swap call has already expired');

        // correct signatures
        uint256 leaderIdx = uint256(IClearing.MarginIndices.Leader);
        address leader = sSC.swapCall.brokers[leaderIdx];
        _requireValidSignature(
            leader,
            abi.encode(channelID, sSC.swapCall),
            _vrsToBytes(sSC.sigs[leaderIdx])
        );

        uint256 followerIdx = uint256(IClearing.MarginIndices.Follower);
        address follower = sSC.swapCall.brokers[followerIdx];
        _requireValidSignature(
            follower,
            abi.encode(channelID, sSC.swapCall),
            _vrsToBytes(sSC.sigs[followerIdx])
        );

        // mark swap as performed
        _sigsUsed[keccak256(abi.encode(sSC.swapCall))] = true;

        // perform swaps from leader (take from the leader, give the follower)
        for (uint256 j = 0; j < sSC.swapCall.swaps[leaderIdx].length; j++) {
            address token = sSC.swapCall.swaps[leaderIdx][j].token;
            uint256 amount = sSC.swapCall.swaps[leaderIdx][j].amount;

            _balances[leader][token] -= amount;
            _balances[follower][token] += amount;
        }

        // perform swaps from follower (take from the follower, give the leader)
        for (uint256 j = 0; j < sSC.swapCall.swaps[followerIdx].length; j++) {
            address token = sSC.swapCall.swaps[followerIdx][j].token;
            uint256 amount = sSC.swapCall.swaps[followerIdx][j].amount;

            _balances[follower][token] -= amount;
            _balances[leader][token] += amount;
        }
    }

    // internal
    function _requireValidSignature(
        address signer,
        bytes memory encodedData,
        bytes memory signature
    ) internal pure {
        require(_recoverSigner(encodedData, signature) == signer, 'Invalid signature');
    }

    function _requireSigsNotUsed(bytes[] memory signatures) internal view {
        for (uint256 sIdx = 0; sIdx < signatures.length; sIdx++) {
            require(!_sigsUsed[keccak256(abi.encode(signatures[sIdx]))], 'Signature already used');
        }
    }

    function _getChainId() internal view returns (uint256) {
        uint256 id;
        /* solhint-disable no-inline-assembly */
        assembly {
            id := chainid()
        }

        /* solhint-disable no-inline-assembly */
        return id;
    }

    function _requireValidPayload(Payload memory payload) internal view {
        _requireValidAddress(payload.broker);
        _requireValidAddress(payload.account);
        _requireValidAddress(payload.asset);

        if (payload.action == Actions.withdraw) {
            require(
                _balances[payload.broker][payload.asset] >= payload.amount,
                'Insufficient holdings'
            );
        }

        require(payload.chainId == _getChainId(), 'Incorrect chainId');
        require(block.timestamp < payload.expire, 'Request expired');
    }

    function _useSignature(bytes memory signature) internal {
        _sigsUsed[keccak256(signature)] = true;
    }

    function _useSignatures(bytes[] memory signatures) internal {
        for (uint256 sIdx = 0; sIdx < signatures.length; sIdx++) {
            _useSignature(signatures[sIdx]);
        }
    }

    function _vrsToBytes(INitroTypes.Signature memory sig) internal pure returns (bytes memory) {
        return abi.encodePacked(sig.v, int160(0x20 + 0x80), sig.v, int160(0x20 + 0x80), sig.s);
    }

    /**
     * @notice Retrieve the result of `transfer` or `transferFrom` function, supposing it is the latest called function.
     * @dev Tackle the inconsistency in ERC20 implementations regarding the return value of `transfer` and `transferFrom`. More: https://github.com/ethereum/solidity/issues/4116.
     * @return result Result of `transfer` or `transferFrom` function.
     */
    function _retrieveTransferResult() internal pure returns (bool result) {
        assembly {
            switch returndatasize()
            case 0 {
                // This is LegacyToken
                result := not(0) // result is true
            }
            case 32 {
                // This is ERC20 compliant token
                returndatacopy(0, 0, 32)
                result := mload(0) // result == return data of external call
            }
            default {
                // This is not an ERC20 token
                revert(0, 0)
            }
        }
    }
}