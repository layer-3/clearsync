// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adjudicator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ExitFormatAllocation is an auto generated low-level Go binding around an user-defined struct.
type ExitFormatAllocation struct {
	Destination    [32]byte
	Amount         *big.Int
	AllocationType uint8
	Metadata       []byte
}

// ExitFormatAssetMetadata is an auto generated low-level Go binding around an user-defined struct.
type ExitFormatAssetMetadata struct {
	AssetType uint8
	Metadata  []byte
}

// ExitFormatSingleAssetExit is an auto generated low-level Go binding around an user-defined struct.
type ExitFormatSingleAssetExit struct {
	Asset         common.Address
	AssetMetadata ExitFormatAssetMetadata
	Allocations   []ExitFormatAllocation
}

// IMultiAssetHolderReclaimArgs is an auto generated low-level Go binding around an user-defined struct.
type IMultiAssetHolderReclaimArgs struct {
	SourceChannelId       [32]byte
	SourceStateHash       [32]byte
	SourceOutcomeBytes    []byte
	SourceAssetIndex      *big.Int
	IndexOfTargetInSource *big.Int
	TargetStateHash       [32]byte
	TargetOutcomeBytes    []byte
	TargetAssetIndex      *big.Int
}

// INitroTypesFixedPart is an auto generated low-level Go binding around an user-defined struct.
type INitroTypesFixedPart struct {
	Participants      []common.Address
	ChannelNonce      uint64
	AppDefinition     common.Address
	ChallengeDuration *big.Int
}

// INitroTypesSignature is an auto generated low-level Go binding around an user-defined struct.
type INitroTypesSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// INitroTypesSignedVariablePart is an auto generated low-level Go binding around an user-defined struct.
type INitroTypesSignedVariablePart struct {
	VariablePart INitroTypesVariablePart
	Sigs         []INitroTypesSignature
}

// INitroTypesVariablePart is an auto generated low-level Go binding around an user-defined struct.
type INitroTypesVariablePart struct {
	Outcome []ExitFormatSingleAssetExit
	AppData []byte
	TurnNum *big.Int
	IsFinal bool
}

// YellowAdjudicatorMetaData contains all meta data concerning the YellowAdjudicator contract.
var YellowAdjudicatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialHoldings\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalHoldings\",\"type\":\"uint256\"}],\"name\":\"AllocationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newTurnNumRecord\",\"type\":\"uint48\"}],\"name\":\"ChallengeCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"finalizesAt\",\"type\":\"uint48\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"ChallengeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newTurnNumRecord\",\"type\":\"uint48\"}],\"name\":\"Checkpointed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"finalizesAt\",\"type\":\"uint48\"}],\"name\":\"Concluded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationHoldings\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"}],\"name\":\"Reclaimed\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature\",\"name\":\"challengerSig\",\"type\":\"tuple\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"checkpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"sourceAllocations\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"targetAllocations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"indexOfTargetInSource\",\"type\":\"uint256\"}],\"name\":\"compute_reclaim_effects\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialHoldings\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"compute_transfer_effects_and_interactions\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"newAllocations\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"allocatesOnlyZeros\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"exitAllocations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalPayouts\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"conclude\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"concludeAndTransferAllAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeld\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"holdings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sourceOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceAssetIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexOfTargetInSource\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"targetStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"targetOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"targetAssetIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIMultiAssetHolder.ReclaimArgs\",\"name\":\"reclaimArgs\",\"type\":\"tuple\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"stateIsSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"statusOf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"fromChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"outcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"}],\"name\":\"transferAllAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"}],\"name\":\"unpackStatus\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"turnNumRecord\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"finalizesAt\",\"type\":\"uint48\"},{\"internalType\":\"uint160\",\"name\":\"fingerprint\",\"type\":\"uint160\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506143b7806100206000396000f3fe6080604052600436106100dd5760003560e01c80635685b7dc1161007f578063c7df14e211610059578063c7df14e21461029f578063d3c4e738146102cc578063ec346235146102ec578063ee049b501461030c57600080fd5b80635685b7dc146102315780636d2a9c921461025f5780638286a0601461027f57600080fd5b80633033730e116100bb5780633033730e1461017657806331afa0b414610196578063552cfa50146101b6578063566d54c61461020457600080fd5b806311e9f178146100e2578063166e56cd1461011b5780632fb1d27014610161575b600080fd5b3480156100ee57600080fd5b506101026100fd366004612d80565b61032c565b6040516101129493929190612ec1565b60405180910390f35b34801561012757600080fd5b50610153610136366004612f25565b600160209081526000928352604080842090915290825290205481565b604051908152602001610112565b61017461016f366004612f51565b610761565b005b34801561018257600080fd5b50610174610191366004612f8c565b61090d565b3480156101a257600080fd5b506101746101b136600461314d565b610999565b3480156101c257600080fd5b506101d66101d136600461319c565b610dda565b6040805165ffffffffffff94851681529390921660208401526001600160a01b031690820152606001610112565b34801561021057600080fd5b5061022461021f3660046131b5565b610df5565b6040516101129190613210565b34801561023d57600080fd5b5061025161024c366004613235565b611221565b6040516101129291906132fb565b34801561026b57600080fd5b5061017461027a366004613651565b6112e8565b34801561028b57600080fd5b5061017461029a3660046136ce565b611444565b3480156102ab57600080fd5b506101536102ba36600461319c565b60006020819052908152604090205481565b3480156102d857600080fd5b506101746102e7366004613767565b6115d0565b3480156102f857600080fd5b50610174610307366004613837565b611652565b34801561031857600080fd5b50610174610327366004613837565b611676565b606060006060600080855111610343578551610346565b84515b6001600160401b0381111561035d5761035d612ab1565b6040519080825280602002602001820160405280156103ae57816020015b604080516080810182526000808252602080830182905292820152606080820152825260001990920191018161037b5790505b5091506000905085516001600160401b038111156103ce576103ce612ab1565b60405190808252806020026020018201604052801561041f57816020015b60408051608081018252600080825260208083018290529282015260608082015282526000199092019101816103ec5790505b50935060019250866000805b8851811015610755578881815181106104465761044661389a565b6020026020010151600001518782815181106104645761046461389a565b602002602001015160000181815250508881815181106104865761048661389a565b6020026020010151604001518782815181106104a4576104a461389a565b60200260200101516040019060ff16908160ff16815250508881815181106104ce576104ce61389a565b6020026020010151606001518782815181106104ec576104ec61389a565b60200260200101516060018190525060006105248a83815181106105125761051261389a565b60200260200101516020015185611680565b9050885160001480610559575088518310801561055957508189848151811061054f5761054f61389a565b6020026020010151145b156106d557600260ff168a84815181106105755761057561389a565b60200260200101516040015160ff16036105d65760405162461bcd60e51b815260206004820152601b60248201527f63616e6e6f74207472616e7366657220612067756172616e746565000000000060448201526064015b60405180910390fd5b808a83815181106105e9576105e961389a565b6020026020010151602001516105ff91906138dc565b8883815181106106115761061161389a565b6020026020010151602001818152505060405180608001604052808b848151811061063e5761063e61389a565b60200260200101516000015181526020018281526020018b84815181106106675761066761389a565b60200260200101516040015160ff1681526020018b848151811061068d5761068d61389a565b6020026020010151606001518152508684815181106106ae576106ae61389a565b60209081029190910101526106c381866138ef565b94506106ce83613902565b9250610716565b8982815181106106e7576106e761389a565b6020026020010151602001518883815181106107055761070561389a565b602002602001015160200181815250505b8782815181106107285761072861389a565b60200260200101516020015160001461074057600096505b61074a81856138dc565b93505060010161042b565b50505093509350935093565b61076c8360a01c1590565b156107b95760405162461bcd60e51b815260206004820152601f60248201527f4465706f73697420746f2065787465726e616c2064657374696e6174696f6e0060448201526064016105cd565b6001600160a01b03841660009081526001602090815260408083208684529091529020548281146108235760405162461bcd60e51b81526020600482015260146024820152731a195b1908084f48195e1c1958dd195912195b1960621b60448201526064016105cd565b6001600160a01b038516610885578134146108805760405162461bcd60e51b815260206004820152601f60248201527f496e636f7272656374206d73672e76616c756520666f72206465706f7369740060448201526064016105cd565b61089a565b61089a6001600160a01b03861633308561169a565b6108a482826138ef565b6001600160a01b03861660008181526001602090815260408083208984528252918290208490558151928352820183905291925085917f87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b91015b60405180910390a25050505050565b600080600061091f888589888a6116f4565b925092509250600080600061095284878d815181106109405761094061389a565b6020026020010151604001518961032c565b935093505092506109698b868c8b8a888a88611778565b61098c868c8151811061097e5761097e61389a565b602002602001015183611868565b5050505050505050505050565b6109a2836118a0565b6109b5816109af84611905565b8561191e565b81516001906000906001600160401b038111156109d4576109d4612ab1565b604051908082528060200260200182016040528015610a0d57816020015b6109fa612a6f565b8152602001906001900390816109f25790505b509050600084516001600160401b03811115610a2b57610a2b612ab1565b604051908082528060200260200182016040528015610a54578160200160208202803683370190505b509050600085516001600160401b03811115610a7257610a72612ab1565b604051908082528060200260200182016040528015610a9b578160200160208202803683370190505b50905060005b8651811015610c60576000878281518110610abe57610abe61389a565b602002602001015190506000816040015190506000898481518110610ae557610ae561389a565b602002602001015160000151905060016000826001600160a01b03166001600160a01b0316815260200190815260200160002060008c815260200190815260200160002054868581518110610b3c57610b3c61389a565b602002602001018181525050600080600080610bb48a8981518110610b6357610b6361389a565b60200260200101518760006001600160401b03811115610b8557610b85612ab1565b604051908082528060200260200182016040528015610bae578160200160208202803683370190505b5061032c565b935093509350935082610bc65760009b505b80898981518110610bd957610bd961389a565b602002602001018181525050838e8981518110610bf857610bf861389a565b6020026020010151604001819052506040518060600160405280866001600160a01b0316815260200188602001518152602001838152508b8981518110610c4157610c4161389a565b6020026020010181905250505050505050508080600101915050610aa1565b5060005b8651811015610d9a576000878281518110610c8157610c8161389a565b6020026020010151600001519050828281518110610ca157610ca161389a565b602002602001015160016000836001600160a01b03166001600160a01b0316815260200190815260200160002060008b81526020019081526020016000206000828254610cee91906138dc565b92505081905550887fc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a83868581518110610d2a57610d2a61389a565b602002602001015160016000866001600160a01b03166001600160a01b0316815260200190815260200160002060008e815260200190815260200160002054604051610d89939291909283526020830191909152604082015260600190565b60405180910390a250600101610c64565b508315610db557600087815260208190526040812055610dc8565b610dc88786610dc389611905565b6119b5565b610dd183611a1b565b50505050505050565b6000806000610de884611a51565b9196909550909350915050565b6060600060018551610e0791906138dc565b6001600160401b03811115610e1e57610e1e612ab1565b604051908082528060200260200182016040528015610e6f57816020015b6040805160808101825260008082526020808301829052928201526060808201528252600019909201910181610e3c5790505b5090506000858481518110610e8657610e8661389a565b602002602001015190506000610e9f8260600151611a9e565b9050600080808080805b8c518110156110ef578a8103610ec257600195506110e7565b60405180608001604052808e8381518110610edf57610edf61389a565b60200260200101516000015181526020018e8381518110610f0257610f0261389a565b60200260200101516020015181526020018e8381518110610f2557610f2561389a565b60200260200101516040015160ff1681526020018e8381518110610f4b57610f4b61389a565b602002602001015160600151815250898381518110610f6c57610f6c61389a565b602002602001018190525084158015610fa5575086600001518d8281518110610f9757610f9761389a565b602002602001015160000151145b15611027578b600081518110610fbd57610fbd61389a565b602002602001015160200151898381518110610fdb57610fdb61389a565b6020026020010151602001818151610ff391906138ef565b9052508b518c906000906110095761100961389a565b6020026020010151602001518361102091906138ef565b9250600194505b83158015611055575086602001518d82815181106110475761104761389a565b602002602001015160000151145b156110d9578b60018151811061106d5761106d61389a565b60200260200101516020015189838151811061108b5761108b61389a565b60200260200101516020018181516110a391906138ef565b9052508b518c9060019081106110bb576110bb61389a565b602002602001015160200151836110d291906138ef565b9250600193505b816110e381613902565b9250505b600101610ea9565b50846111355760405162461bcd60e51b815260206004820152601560248201527418dbdd5b19081b9bdd08199a5b99081d185c99d95d605a1b60448201526064016105cd565b836111785760405162461bcd60e51b815260206004820152601360248201527218dbdd5b19081b9bdd08199a5b99081b19599d606a1b60448201526064016105cd565b826111bc5760405162461bcd60e51b815260206004820152601460248201527318dbdd5b19081b9bdd08199a5b99081c9a59da1d60621b60448201526064016105cd565b8660200151821461120f5760405162461bcd60e51b815260206004820181905260248201527f746f74616c5265636c61696d6564213d67756172616e7465652e616d6f756e7460448201526064016105cd565b509596505050505050505b9392505050565b600060606112348682016040880161391b565b6001600160a01b0316639936d8128761125e61124f82613938565b611259898b613944565b611ac6565b61127861126a8b613938565b61127389613951565b611bad565b6040518463ffffffff1660e01b815260040161129693929190613bb5565b600060405180830381865afa1580156112b3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526112db9190810190613cb8565b9150915094509492505050565b60006112f384611cb2565b82516040015190915061130582611cf8565b61130f8282611d56565b6000606061131e878787611dc4565b909250905080826113425760405162461bcd60e51b81526004016105cd9190613d12565b50600061134e85611e5f565b905061139060405180608001604052808665ffffffffffff168152602001600065ffffffffffff1681526020016000801b81526020016000801b815250611ea9565b6000868152602081905260408120919091558160028111156113b4576113b46138b0565b036113fb5760405165ffffffffffff8516815285907ff3f2d5574c50e581f1a2371fac7dee87f7c6d599a496765fbfa2547ce7fd5f1a9060200160405180910390a261143a565b60405165ffffffffffff8516815285907f07da0a0674fb921e484018c8b81d80e292745e5d8ed134b580c8b9c631c5e9e0906020015b60405180910390a25b5050505050505050565b600061144f85611cb2565b835160400151909150600061146383611e5f565b6002811115611474576114746138b0565b03611488576114838282611f40565b6114bc565b600161149383611e5f565b60028111156114a4576114a46138b0565b036114b3576114838282611d56565b6114bc82611cf8565b600060606114cb888888611dc4565b909250905080826114ef5760405162461bcd60e51b81526004016105cd9190613d12565b506000611500898860000151611faf565b9050611511818a6000015188611ffb565b847f0aa12461ee6c137332989aa12cec79f4772ab2c1a8732a382aada7e9f3ec9d348a60600151426115439190613d25565b8a8a60405161155493929190613dc5565b60405180910390a26115b360405180608001604052808665ffffffffffff1681526020018b60600151426115889190613d25565b65ffffffffffff1681526020018381526020016115ac8a6000015160000151611905565b9052611ea9565b600095865260208690526040909520949094555050505050505050565b6000806115dc836120a8565b9150915060606000838560600151815181106115fa576115fa61389a565b60200260200101516040015190506000838660e00151815181106116205761162061389a565b602002602001015160400151905061163d82828860800151610df5565b9250505061164c8484836122a0565b50505050565b600061165e838361231d565b8251519091506116719082906000610999565b505050565b611671828261231d565b600081831161168f5782611691565b815b90505b92915050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b17905261164c908590612472565b6060600080611702876124d5565b61170b866118a0565b61171d8585805190602001208861191e565b61172684612581565b925082888151811061173a5761173a61389a565b602090810291909101810151516001600160a01b03811660009081526001835260408082209982529890925296902054929895975091955050505050565b6001600160a01b0387166000908152600160209081526040808320898452909152812080548392906117ab9084906138dc565b92505081905550828489815181106117c5576117c561389a565b6020026020010151604001819052506118058686866040516020016117ea9190613e42565b604051602081830303815290604052805190602001206119b5565b6001600160a01b03871660009081526001602090815260408083208984528252918290205482518b81529182018590529181019190915286907fc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a90606001611431565b61189c604051806060016040528084600001516001600160a01b031681526020018460200151815260200183815250612597565b5050565b60026118ab82611e5f565b60028111156118bc576118bc6138b0565b146119025760405162461bcd60e51b815260206004820152601660248201527521b430b73732b6103737ba103334b730b634bd32b21760511b60448201526064016105cd565b50565b600061191082612658565b805190602001209050919050565b600061192982611a51565b6040805160208082018a905281830189905282518083038401815260609092019092528051910120909350915061195d9050565b6001600160a01b0316816001600160a01b03161461164c5760405162461bcd60e51b81526020600482015260156024820152741a5b98dbdc9c9958dd08199a5b99d95c9c1c9a5b9d605a1b60448201526064016105cd565b6000806119c185611a51565b50915091506000611a0160405180608001604052808565ffffffffffff1681526020018465ffffffffffff16815260200187815260200186815250611ea9565b600096875260208790526040909620959095555050505050565b60005b815181101561189c57611a49828281518110611a3c57611a3c61389a565b6020026020010151612597565b600101611a1e565b60008181526020819052604081205481908190610100611a72603082613e55565b61ffff811683901c95509050611a89603082613e55565b949661ffff90951682901c9550909392505050565b6040805180820190915260008082526020820152818060200190518101906116949190613e70565b6060600082516001600160401b03811115611ae357611ae3612ab1565b604051908082528060200260200182016040528015611b4e57816020015b611b3b6040805160c08101825260609181018281528282019290925260006080820181905260a08201529081908152602001600081525090565b815260200190600190039081611b015790505b50905060005b8351811015611ba557611b8085858381518110611b7357611b7361389a565b6020026020010151611bad565b828281518110611b9257611b9261389a565b6020908102919091010152600101611b54565b509392505050565b611be76040805160c08101825260609181018281528282019290925260006080820181905260a08201529081908152602001600081525090565b60408051808201909152825181526000602082018190525b836020015151811015611ba5576000611c42611c1f878760000151611faf565b86602001518481518110611c3557611c3561389a565b6020026020010151612681565b905060005b865151811015611ca8578651805182908110611c6557611c6561389a565b60200260200101516001600160a01b0316826001600160a01b031603611ca057611c90816002613f86565b6020850180519091179052611ca8565b600101611c47565b5050600101611bff565b60008160000151826020015183604001518460600151604051602001611cdb9493929190613fcc565b604051602081830303815290604052805190602001209050919050565b6002611d0382611e5f565b6002811115611d1457611d146138b0565b036119025760405162461bcd60e51b815260206004820152601260248201527121b430b73732b6103334b730b634bd32b21760711b60448201526064016105cd565b6000611d6183611a51565b505090508065ffffffffffff168265ffffffffffff16116116715760405162461bcd60e51b815260206004820152601c60248201527f7475726e4e756d5265636f7264206e6f7420696e637265617365642e0000000060448201526064016105cd565b6000606084604001516001600160a01b0316639936d81286611de68888611ac6565b611df08988611bad565b6040518463ffffffff1660e01b8152600401611e0e93929190614015565b600060405180830381865afa158015611e2b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611e539190810190613cb8565b91509150935093915050565b600080611e6b83611a51565b509150508065ffffffffffff16600003611e885750600092915050565b428165ffffffffffff1611611ea05750600292915050565b50600192915050565b600080610100611eba603082613e55565b845165ffffffffffff1661ffff82161b92509050611ed9603082613e55565b90508061ffff16846020015165ffffffffffff16901b82179150611f2c84604001518560600151604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b6001600160a01b0316919091179392505050565b6000611f4b83611a51565b505090508065ffffffffffff168265ffffffffffff1610156116715760405162461bcd60e51b815260206004820152601860248201527f7475726e4e756d5265636f7264206465637265617365642e000000000000000060448201526064016105cd565b6000611fba83611cb2565b602080840151845160408087015160608801519151611fdd96959192910161407d565b60405160208183030381529060405280519060200120905092915050565b60006120508460405160200161203491815260406020820181905260099082015268666f7263654d6f766560b81b606082015260800190565b6040516020818303038152906040528051906020012083612681565b905061205c818461279c565b61164c5760405162461bcd60e51b815260206004820152601f60248201527f4368616c6c656e676572206973206e6f742061207061727469636970616e740060448201526064016105cd565b8051604082015160608381015160c085015160e0860151929485949093909291906120d2856118a0565b6120e8886020015185805190602001208761191e565b6120f184612581565b96506120fc82612581565b955060008784815181106121125761211261389a565b6020908102919091010151519050600260ff168885815181106121375761213761389a565b6020026020010151604001518a60800151815181106121585761215861389a565b60200260200101516040015160ff16146121b45760405162461bcd60e51b815260206004820152601a60248201527f6e6f7420612067756172616e74656520616c6c6f636174696f6e00000000000060448201526064016105cd565b60008885815181106121c8576121c861389a565b6020026020010151604001518a60800151815181106121e9576121e961389a565b6020026020010151600001519050816001600160a01b03168884815181106122135761221361389a565b6020026020010151600001516001600160a01b0316146122755760405162461bcd60e51b815260206004820152601d60248201527f746172676574417373657420213d2067756172616e746565417373657400000060448201526064016105cd565b61227e816118a0565b6122948a60a0015185805190602001208361191e565b50505050505050915091565b82516060840151835183908590839081106122bd576122bd61389a565b6020026020010151604001819052506122e6828660200151866040516020016117ea9190613e42565b845160608601516040519081527f4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b906020016108fe565b600061232883611cb2565b905061233381611cf8565b81516060015161237b5760405162461bcd60e51b815260206004820152601360248201527214dd185d19481b5d5cdd08189948199a5b985b606a1b60448201526064016105cd565b60006123878484611bad565b905083600001515161239c82602001516127f7565b60ff16146123d95760405162461bcd60e51b815260206004820152600a60248201526921756e616e696d6f757360b01b60448201526064016105cd565b61241e6040518060800160405280600065ffffffffffff1681526020014265ffffffffffff1681526020016000801b81526020016115ac866000015160000151611905565b60008381526020818152604091829020929092555165ffffffffffff4216815283917f4f465027a3d06ea73dd12be0f5c5fc0a34e21f19d6eaed4834a7a944edabc901910160405180910390a25092915050565b60006124876001600160a01b03841683612822565b905080516000141580156124ac5750808060200190518101906124aa91906140ca565b155b1561167157604051635274afe760e01b81526001600160a01b03841660048201526024016105cd565b60005b81516124e58260016138ef565b101561189c57816124f78260016138ef565b815181106125075761250761389a565b60200260200101518282815181106125215761252161389a565b60200260200101511061256f5760405162461bcd60e51b8152602060048201526016602482015275125b991a58d95cc81b5d5cdd081899481cdbdc9d195960521b60448201526064016105cd565b8061257981613902565b9150506124d8565b60608180602001905181019061169491906141e8565b805160005b826040015151811015611671576000836040015182815181106125c1576125c161389a565b60200260200101516000015190506000846040015183815181106125e7576125e761389a565b60200260200101516020015190506126008260a01c1590565b1561261557612610848383612830565b61264e565b6001600160a01b0384166000908152600160209081526040808320858452909152812080548392906126489084906138ef565b90915550505b505060010161259c565b60608160405160200161266b9190613e42565b6040516020818303038152906040529050919050565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018390526000908190605c0160405160208183030381529060405280519060200120905060006001828560000151866020015187604001516040516000815260200160405260405161271b949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa15801561273d573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166127945760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b60448201526064016105cd565b949350505050565b6000805b82518110156127ed578281815181106127bb576127bb61389a565b60200260200101516001600160a01b0316846001600160a01b0316036127e5576001915050611694565b6001016127a0565b5060009392505050565b6000805b82156116945761280c6001846138dc565b909216918061281a81614346565b9150506127fb565b60606116918383600061294d565b6001600160a01b0383166128da576000826001600160a01b03168260405160006040518083038185875af1925050503d806000811461288b576040519150601f19603f3d011682016040523d82523d6000602084013e612890565b606091505b505090508061164c5760405162461bcd60e51b8152602060048201526016602482015275086deead8c840dcdee840e8e4c2dce6cccae4408aa8960531b60448201526064016105cd565b60405163a9059cbb60e01b81526001600160a01b0383811660048301526024820183905284169063a9059cbb906044016020604051808303816000875af1158015612929573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061164c91906140ca565b6060814710156129725760405163cd78605960e01b81523060048201526024016105cd565b600080856001600160a01b0316848660405161298e9190614365565b60006040518083038185875af1925050503d80600081146129cb576040519150601f19603f3d011682016040523d82523d6000602084013e6129d0565b606091505b50915091506129e08683836129ea565b9695505050505050565b6060826129ff576129fa82612a46565b61121a565b8151158015612a1657506001600160a01b0384163b155b15612a3f57604051639996b31560e01b81526001600160a01b03851660048201526024016105cd565b508061121a565b805115612a565780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b604051806060016040528060006001600160a01b03168152602001612a92612a9f565b8152602001606081525090565b60408051808201909152806000612a92565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715612ae957612ae9612ab1565b60405290565b604051606081016001600160401b0381118282101715612ae957612ae9612ab1565b604080519081016001600160401b0381118282101715612ae957612ae9612ab1565b60405161010081016001600160401b0381118282101715612ae957612ae9612ab1565b604051601f8201601f191681016001600160401b0381118282101715612b7e57612b7e612ab1565b604052919050565b60006001600160401b03821115612b9f57612b9f612ab1565b5060051b60200190565b60ff8116811461190257600080fd5b60006001600160401b03821115612bd157612bd1612ab1565b50601f01601f191660200190565b600082601f830112612bf057600080fd5b8135612c03612bfe82612bb8565b612b56565b818152846020838601011115612c1857600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112612c4657600080fd5b81356020612c56612bfe83612b86565b82815260059290921b84018101918181019086841115612c7557600080fd5b8286015b84811015612d165780356001600160401b0380821115612c995760008081fd5b908801906080828b03601f1901811315612cb35760008081fd5b612cbb612ac7565b87840135815260408085013589830152606080860135612cda81612ba9565b83830152928501359284841115612cf357600091508182fd5b612d018e8b86890101612bdf565b90830152508652505050918301918301612c79565b509695505050505050565b600082601f830112612d3257600080fd5b81356020612d42612bfe83612b86565b8083825260208201915060208460051b870101935086841115612d6457600080fd5b602086015b84811015612d165780358352918301918301612d69565b600080600060608486031215612d9557600080fd5b8335925060208401356001600160401b0380821115612db357600080fd5b612dbf87838801612c35565b93506040860135915080821115612dd557600080fd5b50612de286828701612d21565b9150509250925092565b60005b83811015612e07578181015183820152602001612def565b50506000910152565b60008151808452612e28816020860160208601612dec565b601f01601f19169290920160200192915050565b600082825180855260208086019550808260051b84010181860160005b84811015612eb457858303601f19018952815180518452848101518585015260408082015160ff1690850152606090810151608091850182905290612ea081860183612e10565b9a86019a9450505090830190600101612e59565b5090979650505050505050565b608081526000612ed46080830187612e3c565b85151560208401528281036040840152612eee8186612e3c565b91505082606083015295945050505050565b6001600160a01b038116811461190257600080fd5b8035612f2081612f00565b919050565b60008060408385031215612f3857600080fd5b8235612f4381612f00565b946020939093013593505050565b60008060008060808587031215612f6757600080fd5b8435612f7281612f00565b966020860135965060408601359560600135945092505050565b600080600080600060a08688031215612fa457600080fd5b853594506020860135935060408601356001600160401b0380821115612fc957600080fd5b612fd589838a01612bdf565b9450606088013593506080880135915080821115612ff257600080fd5b50612fff88828901612d21565b9150509295509295909350565b6004811061190257600080fd5b600082601f83011261302a57600080fd5b8135602061303a612bfe83612b86565b82815260059290921b8401810191818101908684111561305957600080fd5b8286015b84811015612d165780356001600160401b038082111561307c57600080fd5b90880190601f196060838c038201121561309557600080fd5b61309d612aef565b878401356130aa81612f00565b81526040840135838111156130be57600080fd5b84016040818e03840112156130d257600080fd5b6130da612b11565b9250888101356130e98161300c565b83526040810135848111156130fd57600080fd5b61310b8e8b83850101612bdf565b8a85015250508188820152606084013591508282111561312a57600080fd5b6131388c8984870101612c35565b6040820152865250505091830191830161305d565b60008060006060848603121561316257600080fd5b8335925060208401356001600160401b0381111561317f57600080fd5b61318b86828701613019565b925050604084013590509250925092565b6000602082840312156131ae57600080fd5b5035919050565b6000806000606084860312156131ca57600080fd5b83356001600160401b03808211156131e157600080fd5b6131ed87838801612c35565b9450602086013591508082111561320357600080fd5b5061318b86828701612c35565b6020815260006116916020830184612e3c565b60006040828403121561168f57600080fd5b6000806000806060858703121561324b57600080fd5b84356001600160401b038082111561326257600080fd5b908601906080828903121561327657600080fd5b9094506020860135908082111561328c57600080fd5b818701915087601f8301126132a057600080fd5b8135818111156132af57600080fd5b8860208260051b85010111156132c457600080fd5b6020830195508094505060408701359150808211156132e257600080fd5b506132ef87828801613223565b91505092959194509250565b82151581526040602082015260006127946040830184612e10565b80356001600160401b0381168114612f2057600080fd5b803565ffffffffffff81168114612f2057600080fd5b60006080828403121561335557600080fd5b61335d612ac7565b905081356001600160401b0381111561337557600080fd5b8201601f8101841361338657600080fd5b80356020613396612bfe83612b86565b82815260059290921b830181019181810190878411156133b557600080fd5b938201935b838510156133dc5784356133cd81612f00565b825293820193908201906133ba565b8552506133ea858201613316565b818501525050506133fd60408301612f15565b604082015261340e6060830161332d565b606082015292915050565b801515811461190257600080fd5b60006060828403121561343957600080fd5b613441612aef565b9050813561344e81612ba9565b80825250602082013560208201526040820135604082015292915050565b600082601f83011261347d57600080fd5b8135602061348d612bfe83612b86565b8083825260208201915060606020606086028801019450878511156134b157600080fd5b602087015b85811015612eb4576134c88982613427565b84529284019281016134b6565b6000604082840312156134e757600080fd5b6134ef612b11565b905081356001600160401b038082111561350857600080fd5b908301906080828603121561351c57600080fd5b613524612ac7565b82358281111561353357600080fd5b61353f87828601613019565b82525060208301358281111561355457600080fd5b61356087828601612bdf565b6020830152506135726040840161332d565b60408201526060830135925061358783613419565b8260608201528084525060208401359150808211156135a557600080fd5b506135b28482850161346c565b60208301525092915050565b60006135cc612bfe84612b86565b8381529050602080820190600585901b8401868111156135eb57600080fd5b845b818110156136265780356001600160401b0381111561360c5760008081fd5b613618898289016134d5565b8552509282019282016135ed565b505050509392505050565b600082601f83011261364257600080fd5b611691838335602085016135be565b60008060006060848603121561366657600080fd5b83356001600160401b038082111561367d57600080fd5b61368987838801613343565b9450602086013591508082111561369f57600080fd5b6136ab87838801613631565b935060408601359150808211156136c157600080fd5b50612de2868287016134d5565b60008060008060c085870312156136e457600080fd5b84356001600160401b03808211156136fb57600080fd5b61370788838901613343565b9550602087013591508082111561371d57600080fd5b61372988838901613631565b9450604087013591508082111561373f57600080fd5b5061374c878288016134d5565b92505061375c8660608701613427565b905092959194509250565b60006020828403121561377957600080fd5b81356001600160401b038082111561379057600080fd5b9083019061010082860312156137a557600080fd5b6137ad612b33565b82358152602083013560208201526040830135828111156137cd57600080fd5b6137d987828601612bdf565b604083015250606083013560608201526080830135608082015260a083013560a082015260c08301358281111561380f57600080fd5b61381b87828601612bdf565b60c08301525060e083013560e082015280935050505092915050565b6000806040838503121561384a57600080fd5b82356001600160401b038082111561386157600080fd5b61386d86838701613343565b9350602085013591508082111561388357600080fd5b50613890858286016134d5565b9150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b81810381811115611694576116946138c6565b80820180821115611694576116946138c6565b600060018201613914576139146138c6565b5060010190565b60006020828403121561392d57600080fd5b813561121a81612f00565b60006116943683613343565b60006116913684846135be565b600061169436836134d5565b8183526000602080850194508260005b8581101561399b57813561398081612f00565b6001600160a01b03168752958201959082019060010161396d565b509495945050505050565b600082825180855260208086019550808260051b84010181860160005b84811015612eb457858303601f19018952815180518452848101518585015260408082015160ff1690850152606090810151608091850182905290613a0a81860183612e10565b9a86019a94505050908301906001016139c3565b600082825180855260208086019550808260051b8401018186016000805b85811015613ad857868403601f19018a52825180516001600160a01b0316855285810151606087870181905281519091906004808210613a8957634e487b7160e01b875260218152602487fd5b5091870191909152860151604060808701819052613aaa60a0880183612e10565b91508083015192508682038188015250613ac481836139a6565b9b87019b9550505091840191600101613a3c565b509198975050505050505050565b6000815160808452613afb6080850182613a1e565b905060208301518482036020860152613b148282612e10565b91505065ffffffffffff60408401511660408501526060830151151560608501528091505092915050565b6000815160408452613b546040850182613ae6565b602093840151949093019390935250919050565b60008282518085526020808601955060208260051b8401016020860160005b84811015612eb457601f19868403018952613ba3838351613b3f565b98840198925090830190600101613b87565b6060815260008435601e19863603018112613bcf57600080fd5b85016020810190356001600160401b03811115613beb57600080fd5b8060051b3603821315613bfd57600080fd5b60806060850152613c1260e08501828461395d565b915050613c2160208701613316565b6001600160401b03166080840152613c3b60408701612f15565b6001600160a01b031660a0840152613c556060870161332d565b65ffffffffffff1660c08401528281036020840152613c748186613b68565b905082810360408401526129e08185613b3f565b6000613c96612bfe84612bb8565b9050828152838383011115613caa57600080fd5b61121a836020830184612dec565b60008060408385031215613ccb57600080fd5b8251613cd681613419565b60208401519092506001600160401b03811115613cf257600080fd5b8301601f81018513613d0357600080fd5b61389085825160208401613c88565b6020815260006116916020830184612e10565b65ffffffffffff818116838216019080821115613d4457613d446138c6565b5092915050565b60006040825160408552613d626040860182613ae6565b60208581015187830388830152805180845290820193509091600091908301905b80831015613db9578451805160ff1683528481015185840152860151868301529383019360019290920191606090910190613d83565b50979650505050505050565b60006060820165ffffffffffff8616835260206060602085015281865180845260808601915060808160051b87010193506020880160005b82811015613e2b57607f19888703018452613e19868351613d4b565b95509284019290840190600101613dfd565b505050505082810360408401526129e08185613d4b565b6020815260006116916020830184613a1e565b61ffff828116828216039080821115613d4457613d446138c6565b600060408284031215613e8257600080fd5b613e8a612b11565b82518152602083015160208201528091505092915050565b600181815b80851115613edd578160001904821115613ec357613ec36138c6565b80851615613ed057918102915b93841c9390800290613ea7565b509250929050565b600082613ef457506001611694565b81613f0157506000611694565b8160018114613f175760028114613f2157613f3d565b6001915050611694565b60ff841115613f3257613f326138c6565b50506001821b611694565b5060208310610133831016604e8410600b8410161715613f60575081810a611694565b613f6a8383613ea2565b8060001904821115613f7e57613f7e6138c6565b029392505050565b60006116918383613ee5565b60008151808452602080850194506020840160005b8381101561399b5781516001600160a01b031687529582019590820190600101613fa7565b608081526000613fdf6080830187613f92565b6001600160401b03959095166020830152506001600160a01b0392909216604083015265ffffffffffff16606090910152919050565b60608152600084516080606084015261403160e0840182613f92565b6020878101516001600160401b0316608086015260408801516001600160a01b031660a0860152606088015165ffffffffffff1660c0860152848203908501529050613c748186613b68565b85815260a06020820152600061409660a0830187612e10565b82810360408401526140a88187613a1e565b65ffffffffffff95909516606084015250509015156080909101529392505050565b6000602082840312156140dc57600080fd5b815161121a81613419565b600082601f8301126140f857600080fd5b61169183835160208501613c88565b600082601f83011261411857600080fd5b81516020614128612bfe83612b86565b82815260059290921b8401810191818101908684111561414757600080fd5b8286015b84811015612d165780516001600160401b038082111561416b5760008081fd5b908801906080828b03601f19018113156141855760008081fd5b61418d612ac7565b878401518152604080850151898301526060808601516141ac81612ba9565b838301529285015192848411156141c557600091508182fd5b6141d38e8b868901016140e7565b9083015250865250505091830191830161414b565b6000602082840312156141fa57600080fd5b81516001600160401b038082111561421157600080fd5b818401915084601f83011261422557600080fd5b8151614233612bfe82612b86565b8082825260208201915060208360051b86010192508783111561425557600080fd5b602085015b83811015613db95780518581111561427157600080fd5b8601601f196060828c038201121561428857600080fd5b614290612aef565b602083015161429e81612f00565b81526040830151888111156142b257600080fd5b83016040818e03840112156142c657600080fd5b6142ce612b11565b925060208101516142de8161300c565b83526040810151898111156142f257600080fd5b6143018e6020838501016140e7565b60208501525050816020820152606083015191508782111561432257600080fd5b6143318c602084860101614107565b6040820152855250506020928301920161425a565b600060ff821660ff810361435c5761435c6138c6565b60010192915050565b60008251614377818460208701612dec565b919091019291505056fea26469706673582212206abf76fb44ea97969a084a62e0752c9848746986d48e47a741e74a76c628222664736f6c63430008160033",
}

// YellowAdjudicatorABI is the input ABI used to generate the binding from.
// Deprecated: Use YellowAdjudicatorMetaData.ABI instead.
var YellowAdjudicatorABI = YellowAdjudicatorMetaData.ABI

// YellowAdjudicatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use YellowAdjudicatorMetaData.Bin instead.
var YellowAdjudicatorBin = YellowAdjudicatorMetaData.Bin

// DeployYellowAdjudicator deploys a new Ethereum contract, binding an instance of YellowAdjudicator to it.
func DeployYellowAdjudicator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *YellowAdjudicator, error) {
	parsed, err := YellowAdjudicatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(YellowAdjudicatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &YellowAdjudicator{YellowAdjudicatorCaller: YellowAdjudicatorCaller{contract: contract}, YellowAdjudicatorTransactor: YellowAdjudicatorTransactor{contract: contract}, YellowAdjudicatorFilterer: YellowAdjudicatorFilterer{contract: contract}}, nil
}

// YellowAdjudicator is an auto generated Go binding around an Ethereum contract.
type YellowAdjudicator struct {
	YellowAdjudicatorCaller     // Read-only binding to the contract
	YellowAdjudicatorTransactor // Write-only binding to the contract
	YellowAdjudicatorFilterer   // Log filterer for contract events
}

// YellowAdjudicatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type YellowAdjudicatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YellowAdjudicatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YellowAdjudicatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YellowAdjudicatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YellowAdjudicatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YellowAdjudicatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YellowAdjudicatorSession struct {
	Contract     *YellowAdjudicator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// YellowAdjudicatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YellowAdjudicatorCallerSession struct {
	Contract *YellowAdjudicatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// YellowAdjudicatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YellowAdjudicatorTransactorSession struct {
	Contract     *YellowAdjudicatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// YellowAdjudicatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type YellowAdjudicatorRaw struct {
	Contract *YellowAdjudicator // Generic contract binding to access the raw methods on
}

// YellowAdjudicatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YellowAdjudicatorCallerRaw struct {
	Contract *YellowAdjudicatorCaller // Generic read-only contract binding to access the raw methods on
}

// YellowAdjudicatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YellowAdjudicatorTransactorRaw struct {
	Contract *YellowAdjudicatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYellowAdjudicator creates a new instance of YellowAdjudicator, bound to a specific deployed contract.
func NewYellowAdjudicator(address common.Address, backend bind.ContractBackend) (*YellowAdjudicator, error) {
	contract, err := bindYellowAdjudicator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YellowAdjudicator{YellowAdjudicatorCaller: YellowAdjudicatorCaller{contract: contract}, YellowAdjudicatorTransactor: YellowAdjudicatorTransactor{contract: contract}, YellowAdjudicatorFilterer: YellowAdjudicatorFilterer{contract: contract}}, nil
}

// NewYellowAdjudicatorCaller creates a new read-only instance of YellowAdjudicator, bound to a specific deployed contract.
func NewYellowAdjudicatorCaller(address common.Address, caller bind.ContractCaller) (*YellowAdjudicatorCaller, error) {
	contract, err := bindYellowAdjudicator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YellowAdjudicatorCaller{contract: contract}, nil
}

// NewYellowAdjudicatorTransactor creates a new write-only instance of YellowAdjudicator, bound to a specific deployed contract.
func NewYellowAdjudicatorTransactor(address common.Address, transactor bind.ContractTransactor) (*YellowAdjudicatorTransactor, error) {
	contract, err := bindYellowAdjudicator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YellowAdjudicatorTransactor{contract: contract}, nil
}

// NewYellowAdjudicatorFilterer creates a new log filterer instance of YellowAdjudicator, bound to a specific deployed contract.
func NewYellowAdjudicatorFilterer(address common.Address, filterer bind.ContractFilterer) (*YellowAdjudicatorFilterer, error) {
	contract, err := bindYellowAdjudicator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YellowAdjudicatorFilterer{contract: contract}, nil
}

// bindYellowAdjudicator binds a generic wrapper to an already deployed contract.
func bindYellowAdjudicator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YellowAdjudicatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YellowAdjudicator *YellowAdjudicatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YellowAdjudicator.Contract.YellowAdjudicatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YellowAdjudicator *YellowAdjudicatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.YellowAdjudicatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YellowAdjudicator *YellowAdjudicatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.YellowAdjudicatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YellowAdjudicator *YellowAdjudicatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YellowAdjudicator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YellowAdjudicator *YellowAdjudicatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YellowAdjudicator *YellowAdjudicatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.contract.Transact(opts, method, params...)
}

// ComputeReclaimEffects is a free data retrieval call binding the contract method 0x566d54c6.
//
// Solidity: function compute_reclaim_effects((bytes32,uint256,uint8,bytes)[] sourceAllocations, (bytes32,uint256,uint8,bytes)[] targetAllocations, uint256 indexOfTargetInSource) pure returns((bytes32,uint256,uint8,bytes)[])
func (_YellowAdjudicator *YellowAdjudicatorCaller) ComputeReclaimEffects(opts *bind.CallOpts, sourceAllocations []ExitFormatAllocation, targetAllocations []ExitFormatAllocation, indexOfTargetInSource *big.Int) ([]ExitFormatAllocation, error) {
	var out []interface{}
	err := _YellowAdjudicator.contract.Call(opts, &out, "compute_reclaim_effects", sourceAllocations, targetAllocations, indexOfTargetInSource)

	if err != nil {
		return *new([]ExitFormatAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new([]ExitFormatAllocation)).(*[]ExitFormatAllocation)

	return out0, err

}

// ComputeReclaimEffects is a free data retrieval call binding the contract method 0x566d54c6.
//
// Solidity: function compute_reclaim_effects((bytes32,uint256,uint8,bytes)[] sourceAllocations, (bytes32,uint256,uint8,bytes)[] targetAllocations, uint256 indexOfTargetInSource) pure returns((bytes32,uint256,uint8,bytes)[])
func (_YellowAdjudicator *YellowAdjudicatorSession) ComputeReclaimEffects(sourceAllocations []ExitFormatAllocation, targetAllocations []ExitFormatAllocation, indexOfTargetInSource *big.Int) ([]ExitFormatAllocation, error) {
	return _YellowAdjudicator.Contract.ComputeReclaimEffects(&_YellowAdjudicator.CallOpts, sourceAllocations, targetAllocations, indexOfTargetInSource)
}

// ComputeReclaimEffects is a free data retrieval call binding the contract method 0x566d54c6.
//
// Solidity: function compute_reclaim_effects((bytes32,uint256,uint8,bytes)[] sourceAllocations, (bytes32,uint256,uint8,bytes)[] targetAllocations, uint256 indexOfTargetInSource) pure returns((bytes32,uint256,uint8,bytes)[])
func (_YellowAdjudicator *YellowAdjudicatorCallerSession) ComputeReclaimEffects(sourceAllocations []ExitFormatAllocation, targetAllocations []ExitFormatAllocation, indexOfTargetInSource *big.Int) ([]ExitFormatAllocation, error) {
	return _YellowAdjudicator.Contract.ComputeReclaimEffects(&_YellowAdjudicator.CallOpts, sourceAllocations, targetAllocations, indexOfTargetInSource)
}

// ComputeTransferEffectsAndInteractions is a free data retrieval call binding the contract method 0x11e9f178.
//
// Solidity: function compute_transfer_effects_and_interactions(uint256 initialHoldings, (bytes32,uint256,uint8,bytes)[] allocations, uint256[] indices) pure returns((bytes32,uint256,uint8,bytes)[] newAllocations, bool allocatesOnlyZeros, (bytes32,uint256,uint8,bytes)[] exitAllocations, uint256 totalPayouts)
func (_YellowAdjudicator *YellowAdjudicatorCaller) ComputeTransferEffectsAndInteractions(opts *bind.CallOpts, initialHoldings *big.Int, allocations []ExitFormatAllocation, indices []*big.Int) (struct {
	NewAllocations     []ExitFormatAllocation
	AllocatesOnlyZeros bool
	ExitAllocations    []ExitFormatAllocation
	TotalPayouts       *big.Int
}, error) {
	var out []interface{}
	err := _YellowAdjudicator.contract.Call(opts, &out, "compute_transfer_effects_and_interactions", initialHoldings, allocations, indices)

	outstruct := new(struct {
		NewAllocations     []ExitFormatAllocation
		AllocatesOnlyZeros bool
		ExitAllocations    []ExitFormatAllocation
		TotalPayouts       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAllocations = *abi.ConvertType(out[0], new([]ExitFormatAllocation)).(*[]ExitFormatAllocation)
	outstruct.AllocatesOnlyZeros = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.ExitAllocations = *abi.ConvertType(out[2], new([]ExitFormatAllocation)).(*[]ExitFormatAllocation)
	outstruct.TotalPayouts = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ComputeTransferEffectsAndInteractions is a free data retrieval call binding the contract method 0x11e9f178.
//
// Solidity: function compute_transfer_effects_and_interactions(uint256 initialHoldings, (bytes32,uint256,uint8,bytes)[] allocations, uint256[] indices) pure returns((bytes32,uint256,uint8,bytes)[] newAllocations, bool allocatesOnlyZeros, (bytes32,uint256,uint8,bytes)[] exitAllocations, uint256 totalPayouts)
func (_YellowAdjudicator *YellowAdjudicatorSession) ComputeTransferEffectsAndInteractions(initialHoldings *big.Int, allocations []ExitFormatAllocation, indices []*big.Int) (struct {
	NewAllocations     []ExitFormatAllocation
	AllocatesOnlyZeros bool
	ExitAllocations    []ExitFormatAllocation
	TotalPayouts       *big.Int
}, error) {
	return _YellowAdjudicator.Contract.ComputeTransferEffectsAndInteractions(&_YellowAdjudicator.CallOpts, initialHoldings, allocations, indices)
}

// ComputeTransferEffectsAndInteractions is a free data retrieval call binding the contract method 0x11e9f178.
//
// Solidity: function compute_transfer_effects_and_interactions(uint256 initialHoldings, (bytes32,uint256,uint8,bytes)[] allocations, uint256[] indices) pure returns((bytes32,uint256,uint8,bytes)[] newAllocations, bool allocatesOnlyZeros, (bytes32,uint256,uint8,bytes)[] exitAllocations, uint256 totalPayouts)
func (_YellowAdjudicator *YellowAdjudicatorCallerSession) ComputeTransferEffectsAndInteractions(initialHoldings *big.Int, allocations []ExitFormatAllocation, indices []*big.Int) (struct {
	NewAllocations     []ExitFormatAllocation
	AllocatesOnlyZeros bool
	ExitAllocations    []ExitFormatAllocation
	TotalPayouts       *big.Int
}, error) {
	return _YellowAdjudicator.Contract.ComputeTransferEffectsAndInteractions(&_YellowAdjudicator.CallOpts, initialHoldings, allocations, indices)
}

// Holdings is a free data retrieval call binding the contract method 0x166e56cd.
//
// Solidity: function holdings(address , bytes32 ) view returns(uint256)
func (_YellowAdjudicator *YellowAdjudicatorCaller) Holdings(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _YellowAdjudicator.contract.Call(opts, &out, "holdings", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Holdings is a free data retrieval call binding the contract method 0x166e56cd.
//
// Solidity: function holdings(address , bytes32 ) view returns(uint256)
func (_YellowAdjudicator *YellowAdjudicatorSession) Holdings(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _YellowAdjudicator.Contract.Holdings(&_YellowAdjudicator.CallOpts, arg0, arg1)
}

// Holdings is a free data retrieval call binding the contract method 0x166e56cd.
//
// Solidity: function holdings(address , bytes32 ) view returns(uint256)
func (_YellowAdjudicator *YellowAdjudicatorCallerSession) Holdings(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _YellowAdjudicator.Contract.Holdings(&_YellowAdjudicator.CallOpts, arg0, arg1)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x5685b7dc.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) view returns(bool, string)
func (_YellowAdjudicator *YellowAdjudicatorCaller) StateIsSupported(opts *bind.CallOpts, fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (bool, string, error) {
	var out []interface{}
	err := _YellowAdjudicator.contract.Call(opts, &out, "stateIsSupported", fixedPart, proof, candidate)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// StateIsSupported is a free data retrieval call binding the contract method 0x5685b7dc.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) view returns(bool, string)
func (_YellowAdjudicator *YellowAdjudicatorSession) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (bool, string, error) {
	return _YellowAdjudicator.Contract.StateIsSupported(&_YellowAdjudicator.CallOpts, fixedPart, proof, candidate)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x5685b7dc.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) view returns(bool, string)
func (_YellowAdjudicator *YellowAdjudicatorCallerSession) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (bool, string, error) {
	return _YellowAdjudicator.Contract.StateIsSupported(&_YellowAdjudicator.CallOpts, fixedPart, proof, candidate)
}

// StatusOf is a free data retrieval call binding the contract method 0xc7df14e2.
//
// Solidity: function statusOf(bytes32 ) view returns(bytes32)
func (_YellowAdjudicator *YellowAdjudicatorCaller) StatusOf(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _YellowAdjudicator.contract.Call(opts, &out, "statusOf", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StatusOf is a free data retrieval call binding the contract method 0xc7df14e2.
//
// Solidity: function statusOf(bytes32 ) view returns(bytes32)
func (_YellowAdjudicator *YellowAdjudicatorSession) StatusOf(arg0 [32]byte) ([32]byte, error) {
	return _YellowAdjudicator.Contract.StatusOf(&_YellowAdjudicator.CallOpts, arg0)
}

// StatusOf is a free data retrieval call binding the contract method 0xc7df14e2.
//
// Solidity: function statusOf(bytes32 ) view returns(bytes32)
func (_YellowAdjudicator *YellowAdjudicatorCallerSession) StatusOf(arg0 [32]byte) ([32]byte, error) {
	return _YellowAdjudicator.Contract.StatusOf(&_YellowAdjudicator.CallOpts, arg0)
}

// UnpackStatus is a free data retrieval call binding the contract method 0x552cfa50.
//
// Solidity: function unpackStatus(bytes32 channelId) view returns(uint48 turnNumRecord, uint48 finalizesAt, uint160 fingerprint)
func (_YellowAdjudicator *YellowAdjudicatorCaller) UnpackStatus(opts *bind.CallOpts, channelId [32]byte) (struct {
	TurnNumRecord *big.Int
	FinalizesAt   *big.Int
	Fingerprint   *big.Int
}, error) {
	var out []interface{}
	err := _YellowAdjudicator.contract.Call(opts, &out, "unpackStatus", channelId)

	outstruct := new(struct {
		TurnNumRecord *big.Int
		FinalizesAt   *big.Int
		Fingerprint   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TurnNumRecord = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FinalizesAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fingerprint = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UnpackStatus is a free data retrieval call binding the contract method 0x552cfa50.
//
// Solidity: function unpackStatus(bytes32 channelId) view returns(uint48 turnNumRecord, uint48 finalizesAt, uint160 fingerprint)
func (_YellowAdjudicator *YellowAdjudicatorSession) UnpackStatus(channelId [32]byte) (struct {
	TurnNumRecord *big.Int
	FinalizesAt   *big.Int
	Fingerprint   *big.Int
}, error) {
	return _YellowAdjudicator.Contract.UnpackStatus(&_YellowAdjudicator.CallOpts, channelId)
}

// UnpackStatus is a free data retrieval call binding the contract method 0x552cfa50.
//
// Solidity: function unpackStatus(bytes32 channelId) view returns(uint48 turnNumRecord, uint48 finalizesAt, uint160 fingerprint)
func (_YellowAdjudicator *YellowAdjudicatorCallerSession) UnpackStatus(channelId [32]byte) (struct {
	TurnNumRecord *big.Int
	FinalizesAt   *big.Int
	Fingerprint   *big.Int
}, error) {
	return _YellowAdjudicator.Contract.UnpackStatus(&_YellowAdjudicator.CallOpts, channelId)
}

// Challenge is a paid mutator transaction binding the contract method 0x8286a060.
//
// Solidity: function challenge((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate, (uint8,bytes32,bytes32) challengerSig) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactor) Challenge(opts *bind.TransactOpts, fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart, challengerSig INitroTypesSignature) (*types.Transaction, error) {
	return _YellowAdjudicator.contract.Transact(opts, "challenge", fixedPart, proof, candidate, challengerSig)
}

// Challenge is a paid mutator transaction binding the contract method 0x8286a060.
//
// Solidity: function challenge((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate, (uint8,bytes32,bytes32) challengerSig) returns()
func (_YellowAdjudicator *YellowAdjudicatorSession) Challenge(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart, challengerSig INitroTypesSignature) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Challenge(&_YellowAdjudicator.TransactOpts, fixedPart, proof, candidate, challengerSig)
}

// Challenge is a paid mutator transaction binding the contract method 0x8286a060.
//
// Solidity: function challenge((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate, (uint8,bytes32,bytes32) challengerSig) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactorSession) Challenge(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart, challengerSig INitroTypesSignature) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Challenge(&_YellowAdjudicator.TransactOpts, fixedPart, proof, candidate, challengerSig)
}

// Checkpoint is a paid mutator transaction binding the contract method 0x6d2a9c92.
//
// Solidity: function checkpoint((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactor) Checkpoint(opts *bind.TransactOpts, fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _YellowAdjudicator.contract.Transact(opts, "checkpoint", fixedPart, proof, candidate)
}

// Checkpoint is a paid mutator transaction binding the contract method 0x6d2a9c92.
//
// Solidity: function checkpoint((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_YellowAdjudicator *YellowAdjudicatorSession) Checkpoint(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Checkpoint(&_YellowAdjudicator.TransactOpts, fixedPart, proof, candidate)
}

// Checkpoint is a paid mutator transaction binding the contract method 0x6d2a9c92.
//
// Solidity: function checkpoint((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactorSession) Checkpoint(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Checkpoint(&_YellowAdjudicator.TransactOpts, fixedPart, proof, candidate)
}

// Conclude is a paid mutator transaction binding the contract method 0xee049b50.
//
// Solidity: function conclude((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactor) Conclude(opts *bind.TransactOpts, fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _YellowAdjudicator.contract.Transact(opts, "conclude", fixedPart, candidate)
}

// Conclude is a paid mutator transaction binding the contract method 0xee049b50.
//
// Solidity: function conclude((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_YellowAdjudicator *YellowAdjudicatorSession) Conclude(fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Conclude(&_YellowAdjudicator.TransactOpts, fixedPart, candidate)
}

// Conclude is a paid mutator transaction binding the contract method 0xee049b50.
//
// Solidity: function conclude((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactorSession) Conclude(fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Conclude(&_YellowAdjudicator.TransactOpts, fixedPart, candidate)
}

// ConcludeAndTransferAllAssets is a paid mutator transaction binding the contract method 0xec346235.
//
// Solidity: function concludeAndTransferAllAssets((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactor) ConcludeAndTransferAllAssets(opts *bind.TransactOpts, fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _YellowAdjudicator.contract.Transact(opts, "concludeAndTransferAllAssets", fixedPart, candidate)
}

// ConcludeAndTransferAllAssets is a paid mutator transaction binding the contract method 0xec346235.
//
// Solidity: function concludeAndTransferAllAssets((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_YellowAdjudicator *YellowAdjudicatorSession) ConcludeAndTransferAllAssets(fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.ConcludeAndTransferAllAssets(&_YellowAdjudicator.TransactOpts, fixedPart, candidate)
}

// ConcludeAndTransferAllAssets is a paid mutator transaction binding the contract method 0xec346235.
//
// Solidity: function concludeAndTransferAllAssets((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactorSession) ConcludeAndTransferAllAssets(fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.ConcludeAndTransferAllAssets(&_YellowAdjudicator.TransactOpts, fixedPart, candidate)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 channelId, uint256 expectedHeld, uint256 amount) payable returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, channelId [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _YellowAdjudicator.contract.Transact(opts, "deposit", asset, channelId, expectedHeld, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 channelId, uint256 expectedHeld, uint256 amount) payable returns()
func (_YellowAdjudicator *YellowAdjudicatorSession) Deposit(asset common.Address, channelId [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Deposit(&_YellowAdjudicator.TransactOpts, asset, channelId, expectedHeld, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 channelId, uint256 expectedHeld, uint256 amount) payable returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactorSession) Deposit(asset common.Address, channelId [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Deposit(&_YellowAdjudicator.TransactOpts, asset, channelId, expectedHeld, amount)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactor) Reclaim(opts *bind.TransactOpts, reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _YellowAdjudicator.contract.Transact(opts, "reclaim", reclaimArgs)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_YellowAdjudicator *YellowAdjudicatorSession) Reclaim(reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Reclaim(&_YellowAdjudicator.TransactOpts, reclaimArgs)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactorSession) Reclaim(reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Reclaim(&_YellowAdjudicator.TransactOpts, reclaimArgs)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactor) Transfer(opts *bind.TransactOpts, assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _YellowAdjudicator.contract.Transact(opts, "transfer", assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_YellowAdjudicator *YellowAdjudicatorSession) Transfer(assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Transfer(&_YellowAdjudicator.TransactOpts, assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactorSession) Transfer(assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.Transfer(&_YellowAdjudicator.TransactOpts, assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// TransferAllAssets is a paid mutator transaction binding the contract method 0x31afa0b4.
//
// Solidity: function transferAllAssets(bytes32 channelId, (address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[] outcome, bytes32 stateHash) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactor) TransferAllAssets(opts *bind.TransactOpts, channelId [32]byte, outcome []ExitFormatSingleAssetExit, stateHash [32]byte) (*types.Transaction, error) {
	return _YellowAdjudicator.contract.Transact(opts, "transferAllAssets", channelId, outcome, stateHash)
}

// TransferAllAssets is a paid mutator transaction binding the contract method 0x31afa0b4.
//
// Solidity: function transferAllAssets(bytes32 channelId, (address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[] outcome, bytes32 stateHash) returns()
func (_YellowAdjudicator *YellowAdjudicatorSession) TransferAllAssets(channelId [32]byte, outcome []ExitFormatSingleAssetExit, stateHash [32]byte) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.TransferAllAssets(&_YellowAdjudicator.TransactOpts, channelId, outcome, stateHash)
}

// TransferAllAssets is a paid mutator transaction binding the contract method 0x31afa0b4.
//
// Solidity: function transferAllAssets(bytes32 channelId, (address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[] outcome, bytes32 stateHash) returns()
func (_YellowAdjudicator *YellowAdjudicatorTransactorSession) TransferAllAssets(channelId [32]byte, outcome []ExitFormatSingleAssetExit, stateHash [32]byte) (*types.Transaction, error) {
	return _YellowAdjudicator.Contract.TransferAllAssets(&_YellowAdjudicator.TransactOpts, channelId, outcome, stateHash)
}

// YellowAdjudicatorAllocationUpdatedIterator is returned from FilterAllocationUpdated and is used to iterate over the raw logs and unpacked data for AllocationUpdated events raised by the YellowAdjudicator contract.
type YellowAdjudicatorAllocationUpdatedIterator struct {
	Event *YellowAdjudicatorAllocationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YellowAdjudicatorAllocationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowAdjudicatorAllocationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YellowAdjudicatorAllocationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YellowAdjudicatorAllocationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowAdjudicatorAllocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowAdjudicatorAllocationUpdated represents a AllocationUpdated event raised by the YellowAdjudicator contract.
type YellowAdjudicatorAllocationUpdated struct {
	ChannelId       [32]byte
	AssetIndex      *big.Int
	InitialHoldings *big.Int
	FinalHoldings   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAllocationUpdated is a free log retrieval operation binding the contract event 0xc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a.
//
// Solidity: event AllocationUpdated(bytes32 indexed channelId, uint256 assetIndex, uint256 initialHoldings, uint256 finalHoldings)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) FilterAllocationUpdated(opts *bind.FilterOpts, channelId [][32]byte) (*YellowAdjudicatorAllocationUpdatedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.FilterLogs(opts, "AllocationUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &YellowAdjudicatorAllocationUpdatedIterator{contract: _YellowAdjudicator.contract, event: "AllocationUpdated", logs: logs, sub: sub}, nil
}

// WatchAllocationUpdated is a free log subscription operation binding the contract event 0xc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a.
//
// Solidity: event AllocationUpdated(bytes32 indexed channelId, uint256 assetIndex, uint256 initialHoldings, uint256 finalHoldings)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) WatchAllocationUpdated(opts *bind.WatchOpts, sink chan<- *YellowAdjudicatorAllocationUpdated, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.WatchLogs(opts, "AllocationUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowAdjudicatorAllocationUpdated)
				if err := _YellowAdjudicator.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllocationUpdated is a log parse operation binding the contract event 0xc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a.
//
// Solidity: event AllocationUpdated(bytes32 indexed channelId, uint256 assetIndex, uint256 initialHoldings, uint256 finalHoldings)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) ParseAllocationUpdated(log types.Log) (*YellowAdjudicatorAllocationUpdated, error) {
	event := new(YellowAdjudicatorAllocationUpdated)
	if err := _YellowAdjudicator.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowAdjudicatorChallengeClearedIterator is returned from FilterChallengeCleared and is used to iterate over the raw logs and unpacked data for ChallengeCleared events raised by the YellowAdjudicator contract.
type YellowAdjudicatorChallengeClearedIterator struct {
	Event *YellowAdjudicatorChallengeCleared // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YellowAdjudicatorChallengeClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowAdjudicatorChallengeCleared)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YellowAdjudicatorChallengeCleared)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YellowAdjudicatorChallengeClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowAdjudicatorChallengeClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowAdjudicatorChallengeCleared represents a ChallengeCleared event raised by the YellowAdjudicator contract.
type YellowAdjudicatorChallengeCleared struct {
	ChannelId        [32]byte
	NewTurnNumRecord *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterChallengeCleared is a free log retrieval operation binding the contract event 0x07da0a0674fb921e484018c8b81d80e292745e5d8ed134b580c8b9c631c5e9e0.
//
// Solidity: event ChallengeCleared(bytes32 indexed channelId, uint48 newTurnNumRecord)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) FilterChallengeCleared(opts *bind.FilterOpts, channelId [][32]byte) (*YellowAdjudicatorChallengeClearedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.FilterLogs(opts, "ChallengeCleared", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &YellowAdjudicatorChallengeClearedIterator{contract: _YellowAdjudicator.contract, event: "ChallengeCleared", logs: logs, sub: sub}, nil
}

// WatchChallengeCleared is a free log subscription operation binding the contract event 0x07da0a0674fb921e484018c8b81d80e292745e5d8ed134b580c8b9c631c5e9e0.
//
// Solidity: event ChallengeCleared(bytes32 indexed channelId, uint48 newTurnNumRecord)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) WatchChallengeCleared(opts *bind.WatchOpts, sink chan<- *YellowAdjudicatorChallengeCleared, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.WatchLogs(opts, "ChallengeCleared", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowAdjudicatorChallengeCleared)
				if err := _YellowAdjudicator.contract.UnpackLog(event, "ChallengeCleared", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeCleared is a log parse operation binding the contract event 0x07da0a0674fb921e484018c8b81d80e292745e5d8ed134b580c8b9c631c5e9e0.
//
// Solidity: event ChallengeCleared(bytes32 indexed channelId, uint48 newTurnNumRecord)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) ParseChallengeCleared(log types.Log) (*YellowAdjudicatorChallengeCleared, error) {
	event := new(YellowAdjudicatorChallengeCleared)
	if err := _YellowAdjudicator.contract.UnpackLog(event, "ChallengeCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowAdjudicatorChallengeRegisteredIterator is returned from FilterChallengeRegistered and is used to iterate over the raw logs and unpacked data for ChallengeRegistered events raised by the YellowAdjudicator contract.
type YellowAdjudicatorChallengeRegisteredIterator struct {
	Event *YellowAdjudicatorChallengeRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YellowAdjudicatorChallengeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowAdjudicatorChallengeRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YellowAdjudicatorChallengeRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YellowAdjudicatorChallengeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowAdjudicatorChallengeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowAdjudicatorChallengeRegistered represents a ChallengeRegistered event raised by the YellowAdjudicator contract.
type YellowAdjudicatorChallengeRegistered struct {
	ChannelId   [32]byte
	FinalizesAt *big.Int
	Proof       []INitroTypesSignedVariablePart
	Candidate   INitroTypesSignedVariablePart
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChallengeRegistered is a free log retrieval operation binding the contract event 0x0aa12461ee6c137332989aa12cec79f4772ab2c1a8732a382aada7e9f3ec9d34.
//
// Solidity: event ChallengeRegistered(bytes32 indexed channelId, uint48 finalizesAt, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) FilterChallengeRegistered(opts *bind.FilterOpts, channelId [][32]byte) (*YellowAdjudicatorChallengeRegisteredIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.FilterLogs(opts, "ChallengeRegistered", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &YellowAdjudicatorChallengeRegisteredIterator{contract: _YellowAdjudicator.contract, event: "ChallengeRegistered", logs: logs, sub: sub}, nil
}

// WatchChallengeRegistered is a free log subscription operation binding the contract event 0x0aa12461ee6c137332989aa12cec79f4772ab2c1a8732a382aada7e9f3ec9d34.
//
// Solidity: event ChallengeRegistered(bytes32 indexed channelId, uint48 finalizesAt, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) WatchChallengeRegistered(opts *bind.WatchOpts, sink chan<- *YellowAdjudicatorChallengeRegistered, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.WatchLogs(opts, "ChallengeRegistered", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowAdjudicatorChallengeRegistered)
				if err := _YellowAdjudicator.contract.UnpackLog(event, "ChallengeRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeRegistered is a log parse operation binding the contract event 0x0aa12461ee6c137332989aa12cec79f4772ab2c1a8732a382aada7e9f3ec9d34.
//
// Solidity: event ChallengeRegistered(bytes32 indexed channelId, uint48 finalizesAt, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) ParseChallengeRegistered(log types.Log) (*YellowAdjudicatorChallengeRegistered, error) {
	event := new(YellowAdjudicatorChallengeRegistered)
	if err := _YellowAdjudicator.contract.UnpackLog(event, "ChallengeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowAdjudicatorCheckpointedIterator is returned from FilterCheckpointed and is used to iterate over the raw logs and unpacked data for Checkpointed events raised by the YellowAdjudicator contract.
type YellowAdjudicatorCheckpointedIterator struct {
	Event *YellowAdjudicatorCheckpointed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YellowAdjudicatorCheckpointedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowAdjudicatorCheckpointed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YellowAdjudicatorCheckpointed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YellowAdjudicatorCheckpointedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowAdjudicatorCheckpointedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowAdjudicatorCheckpointed represents a Checkpointed event raised by the YellowAdjudicator contract.
type YellowAdjudicatorCheckpointed struct {
	ChannelId        [32]byte
	NewTurnNumRecord *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCheckpointed is a free log retrieval operation binding the contract event 0xf3f2d5574c50e581f1a2371fac7dee87f7c6d599a496765fbfa2547ce7fd5f1a.
//
// Solidity: event Checkpointed(bytes32 indexed channelId, uint48 newTurnNumRecord)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) FilterCheckpointed(opts *bind.FilterOpts, channelId [][32]byte) (*YellowAdjudicatorCheckpointedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.FilterLogs(opts, "Checkpointed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &YellowAdjudicatorCheckpointedIterator{contract: _YellowAdjudicator.contract, event: "Checkpointed", logs: logs, sub: sub}, nil
}

// WatchCheckpointed is a free log subscription operation binding the contract event 0xf3f2d5574c50e581f1a2371fac7dee87f7c6d599a496765fbfa2547ce7fd5f1a.
//
// Solidity: event Checkpointed(bytes32 indexed channelId, uint48 newTurnNumRecord)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) WatchCheckpointed(opts *bind.WatchOpts, sink chan<- *YellowAdjudicatorCheckpointed, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.WatchLogs(opts, "Checkpointed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowAdjudicatorCheckpointed)
				if err := _YellowAdjudicator.contract.UnpackLog(event, "Checkpointed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCheckpointed is a log parse operation binding the contract event 0xf3f2d5574c50e581f1a2371fac7dee87f7c6d599a496765fbfa2547ce7fd5f1a.
//
// Solidity: event Checkpointed(bytes32 indexed channelId, uint48 newTurnNumRecord)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) ParseCheckpointed(log types.Log) (*YellowAdjudicatorCheckpointed, error) {
	event := new(YellowAdjudicatorCheckpointed)
	if err := _YellowAdjudicator.contract.UnpackLog(event, "Checkpointed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowAdjudicatorConcludedIterator is returned from FilterConcluded and is used to iterate over the raw logs and unpacked data for Concluded events raised by the YellowAdjudicator contract.
type YellowAdjudicatorConcludedIterator struct {
	Event *YellowAdjudicatorConcluded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YellowAdjudicatorConcludedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowAdjudicatorConcluded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YellowAdjudicatorConcluded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YellowAdjudicatorConcludedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowAdjudicatorConcludedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowAdjudicatorConcluded represents a Concluded event raised by the YellowAdjudicator contract.
type YellowAdjudicatorConcluded struct {
	ChannelId   [32]byte
	FinalizesAt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConcluded is a free log retrieval operation binding the contract event 0x4f465027a3d06ea73dd12be0f5c5fc0a34e21f19d6eaed4834a7a944edabc901.
//
// Solidity: event Concluded(bytes32 indexed channelId, uint48 finalizesAt)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) FilterConcluded(opts *bind.FilterOpts, channelId [][32]byte) (*YellowAdjudicatorConcludedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.FilterLogs(opts, "Concluded", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &YellowAdjudicatorConcludedIterator{contract: _YellowAdjudicator.contract, event: "Concluded", logs: logs, sub: sub}, nil
}

// WatchConcluded is a free log subscription operation binding the contract event 0x4f465027a3d06ea73dd12be0f5c5fc0a34e21f19d6eaed4834a7a944edabc901.
//
// Solidity: event Concluded(bytes32 indexed channelId, uint48 finalizesAt)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) WatchConcluded(opts *bind.WatchOpts, sink chan<- *YellowAdjudicatorConcluded, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.WatchLogs(opts, "Concluded", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowAdjudicatorConcluded)
				if err := _YellowAdjudicator.contract.UnpackLog(event, "Concluded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConcluded is a log parse operation binding the contract event 0x4f465027a3d06ea73dd12be0f5c5fc0a34e21f19d6eaed4834a7a944edabc901.
//
// Solidity: event Concluded(bytes32 indexed channelId, uint48 finalizesAt)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) ParseConcluded(log types.Log) (*YellowAdjudicatorConcluded, error) {
	event := new(YellowAdjudicatorConcluded)
	if err := _YellowAdjudicator.contract.UnpackLog(event, "Concluded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowAdjudicatorDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the YellowAdjudicator contract.
type YellowAdjudicatorDepositedIterator struct {
	Event *YellowAdjudicatorDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YellowAdjudicatorDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowAdjudicatorDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YellowAdjudicatorDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YellowAdjudicatorDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowAdjudicatorDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowAdjudicatorDeposited represents a Deposited event raised by the YellowAdjudicator contract.
type YellowAdjudicatorDeposited struct {
	Destination         [32]byte
	Asset               common.Address
	DestinationHoldings *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b.
//
// Solidity: event Deposited(bytes32 indexed destination, address asset, uint256 destinationHoldings)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) FilterDeposited(opts *bind.FilterOpts, destination [][32]byte) (*YellowAdjudicatorDepositedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.FilterLogs(opts, "Deposited", destinationRule)
	if err != nil {
		return nil, err
	}
	return &YellowAdjudicatorDepositedIterator{contract: _YellowAdjudicator.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b.
//
// Solidity: event Deposited(bytes32 indexed destination, address asset, uint256 destinationHoldings)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *YellowAdjudicatorDeposited, destination [][32]byte) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.WatchLogs(opts, "Deposited", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowAdjudicatorDeposited)
				if err := _YellowAdjudicator.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b.
//
// Solidity: event Deposited(bytes32 indexed destination, address asset, uint256 destinationHoldings)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) ParseDeposited(log types.Log) (*YellowAdjudicatorDeposited, error) {
	event := new(YellowAdjudicatorDeposited)
	if err := _YellowAdjudicator.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YellowAdjudicatorReclaimedIterator is returned from FilterReclaimed and is used to iterate over the raw logs and unpacked data for Reclaimed events raised by the YellowAdjudicator contract.
type YellowAdjudicatorReclaimedIterator struct {
	Event *YellowAdjudicatorReclaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *YellowAdjudicatorReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YellowAdjudicatorReclaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(YellowAdjudicatorReclaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *YellowAdjudicatorReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YellowAdjudicatorReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YellowAdjudicatorReclaimed represents a Reclaimed event raised by the YellowAdjudicator contract.
type YellowAdjudicatorReclaimed struct {
	ChannelId  [32]byte
	AssetIndex *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReclaimed is a free log retrieval operation binding the contract event 0x4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b.
//
// Solidity: event Reclaimed(bytes32 indexed channelId, uint256 assetIndex)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) FilterReclaimed(opts *bind.FilterOpts, channelId [][32]byte) (*YellowAdjudicatorReclaimedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.FilterLogs(opts, "Reclaimed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &YellowAdjudicatorReclaimedIterator{contract: _YellowAdjudicator.contract, event: "Reclaimed", logs: logs, sub: sub}, nil
}

// WatchReclaimed is a free log subscription operation binding the contract event 0x4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b.
//
// Solidity: event Reclaimed(bytes32 indexed channelId, uint256 assetIndex)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) WatchReclaimed(opts *bind.WatchOpts, sink chan<- *YellowAdjudicatorReclaimed, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _YellowAdjudicator.contract.WatchLogs(opts, "Reclaimed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YellowAdjudicatorReclaimed)
				if err := _YellowAdjudicator.contract.UnpackLog(event, "Reclaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReclaimed is a log parse operation binding the contract event 0x4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b.
//
// Solidity: event Reclaimed(bytes32 indexed channelId, uint256 assetIndex)
func (_YellowAdjudicator *YellowAdjudicatorFilterer) ParseReclaimed(log types.Log) (*YellowAdjudicatorReclaimed, error) {
	event := new(YellowAdjudicatorReclaimed)
	if err := _YellowAdjudicator.contract.UnpackLog(event, "Reclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
