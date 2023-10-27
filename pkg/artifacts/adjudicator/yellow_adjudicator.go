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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialHoldings\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalHoldings\",\"type\":\"uint256\"}],\"name\":\"AllocationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newTurnNumRecord\",\"type\":\"uint48\"}],\"name\":\"ChallengeCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"finalizesAt\",\"type\":\"uint48\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"ChallengeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newTurnNumRecord\",\"type\":\"uint48\"}],\"name\":\"Checkpointed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"finalizesAt\",\"type\":\"uint48\"}],\"name\":\"Concluded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationHoldings\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"}],\"name\":\"Reclaimed\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature\",\"name\":\"challengerSig\",\"type\":\"tuple\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"checkpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"sourceAllocations\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"targetAllocations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"indexOfTargetInSource\",\"type\":\"uint256\"}],\"name\":\"compute_reclaim_effects\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialHoldings\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"compute_transfer_effects_and_interactions\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"newAllocations\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"allocatesOnlyZeros\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"exitAllocations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalPayouts\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"conclude\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"concludeAndTransferAllAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeld\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"holdings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sourceOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceAssetIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexOfTargetInSource\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"targetStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"targetOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"targetAssetIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIMultiAssetHolder.ReclaimArgs\",\"name\":\"reclaimArgs\",\"type\":\"tuple\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"stateIsSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"statusOf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"fromChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"outcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"}],\"name\":\"transferAllAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"}],\"name\":\"unpackStatus\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"turnNumRecord\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"finalizesAt\",\"type\":\"uint48\"},{\"internalType\":\"uint160\",\"name\":\"fingerprint\",\"type\":\"uint160\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052346200001d575b604051614a4e6200002d8239614a4e90f35b62000026600080fd5b6200000b56fe60806040526004361015610018575b610016600080fd5b005b60003560e01c806311e9f17814610130578063166e56cd146101275780632fb1d2701461011e5780633033730e1461011557806331afa0b41461010c578063552cfa5014610103578063566d54c6146100fa5780635685b7dc146100f15780636d2a9c92146100e85780638286a060146100df578063c7df14e2146100d6578063d3c4e738146100cd578063ec346235146100c45763ee049b500361000e576100bf6114b0565b61000e565b506100bf611488565b506100bf6113ff565b506100bf6112c0565b506100bf611286565b506100bf6111c5565b506100bf610d01565b506100bf610b9c565b506100bf610aed565b506100bf610a7f565b506100bf610848565b506100bf6107a0565b506100bf61071a565b506100bf610603565b600080fd5b805b0361013957565b905035906101548261013e565b565b601f01601f191690565b50634e487b7160e01b600052604160045260246000fd5b90601f01601f191681019081106001600160401b0382111761019857604052565b6101a0610160565b604052565b906101546101b260405190565b9283610177565b602080916001600160401b0381116101d057020190565b6101d8610160565b020190565b60ff8116610140565b90503590610154826101dd565b6102116020916001600160401b03811161021557601f01601f191690565b0190565b610156610160565b90826000939282370152565b9291906101549161024161023c836101f3565b6101a5565b948286526020860191838201111561021d5761025b600080fd5b61021d565b9061027c9181601f8201121561027f575b602081359101610229565b90565b610287600080fd5b610271565b91906102f79060808482031261030b575b6102a760806101a5565b9360006102b48383610147565b9086015260206102c683828401610147565b9086015260406102d8838284016101e6565b908601526060810135906001600160401b0382116102fe575b01610260565b6060830152565b610306600080fd5b6102f1565b610313600080fd5b61029d565b9092919261032861023c826101b9565b9381855260208086019202830192818411610387575b80925b84841061034f575050505050565b6020809161036f8587356001600160401b03811161037a575b860161028c565b815201930192610341565b610382600080fd5b610368565b61038f600080fd5b61033e565b9061027c9181601f820112156103b0575b602081359101610318565b6103b8600080fd5b6103a5565b909291926103cd61023c826101b9565b938185526020808601920283019281841161040b575b915b8383106103f25750505050565b602080916104008486610147565b8152019201916103e5565b610413600080fd5b6103e3565b9061027c9181601f82011215610434575b6020813591016103bd565b61043c600080fd5b610429565b90916060828403126104b1575b61027c61045b8484610147565b9361047b8160208601356001600160401b0381116104a4575b8601610394565b936040810135906001600160401b038211610497575b01610418565b61049f600080fd5b610491565b6104ac600080fd5b610474565b6104b9600080fd5b61044e565b9052565b60005b8381106104d55750506000910152565b81810151838201526020016104c5565b610506610156602093610211936104fa815190565b80835293849260200190565b958691016104c2565b8051825261027c91608081019160609061052e60208201516020850152565b60408181015160ff169084015201519060608184039101526104e5565b9061027c9161050f565b9061056b610561835190565b8083529160200190565b908161057d6020830284019460200190565b926000915b83831061059157505050505090565b909192939460206105b46105ad8385600195038752895161054b565b9760200190565b9301930191939290610582565b9493916105fe90610154946105f16105e760609560808b01908b820360008d0152610555565b92151560208a0152565b8782036040890152610555565b940152565b503461063b575b61063761062161061b366004610441565b91612ef0565b9061062e94929460405190565b948594856105c1565b0390f35b610643600080fd5b61060a565b6001600160a01b031690565b6001600160a01b038116610140565b9050359061015482610654565b919061027c90604084820312610694575b61068b8185610663565b93602001610147565b61069c600080fd5b610681565b61027c90610648906001600160a01b031682565b61027c906106a1565b61027c906106b5565b906106d1906106be565b600052602052604060002090565b906106d1565b61027c916008021c81565b9061027c91546106e5565b61071561027c926107106001936000946106c7565b6106df565b6106f0565b5034610749575b610637610738610732366004610670565b906106fb565b6040515b9182918290815260200190565b610751600080fd5b610721565b608081830312610793575b61076b8282610663565b9261027c61077c8460208501610147565b9361078a8160408601610147565b93606001610147565b61079b600080fd5b610761565b506107b86107af366004610756565b92919091612656565b604051005b919060a08382031261083b575b6107d48184610147565b926107e28260208301610147565b9261027c6108058460408501356001600160401b03811161082e575b8501610260565b936108138160608601610147565b936080810135906001600160401b0382116104975701610418565b610836600080fd5b6107fe565b610843600080fd5b6107ca565b5034610866575b6107b861085d3660046107bd565b93929092612a7f565b61086e600080fd5b61084f565b6004111561013957565b9050359061015482610873565b91906108d0906040848203126108d7575b6108a560406101a5565b9360006108b2838361087d565b908601526020810135906001600160401b0382116102fe5701610260565b6020830152565b6108df600080fd5b61089b565b919061094f90606084820312610970575b6108ff60606101a5565b93600061090c8383610663565b9086015261092f8260208301356001600160401b038111610963575b830161088a565b60208601526040810135906001600160401b038211610956575b01610394565b6040830152565b61095e600080fd5b610949565b61096b600080fd5b610928565b610978600080fd5b6108f5565b9092919261098d61023c826101b9565b93818552602080860192028301928184116109ec575b80925b8484106109b4575050505050565b602080916109d48587356001600160401b0381116109df575b86016108e4565b8152019301926109a6565b6109e7600080fd5b6109cd565b6109f4600080fd5b6109a3565b9061027c9181601f82011215610a15575b60208135910161097d565b610a1d600080fd5b610a0a565b9091606082840312610a72575b61027c610a3c8484610147565b93610a5c8160208601356001600160401b038111610a65575b86016109f9565b93604001610147565b610a6d600080fd5b610a55565b610a7a600080fd5b610a2f565b5034610a9a575b6107b8610a94366004610a22565b91613f8d565b610aa2600080fd5b610a86565b9061027c916020818303121561014757610abf600080fd5b610147565b65ffffffffffff9182168152911660208201526001600160a01b03909116604082015260600190565b5034610b19575b610637610b0a610b05366004610aa7565b6114d8565b60405191939193849384610ac4565b610b21600080fd5b610af4565b9091606082840312610b7e575b61027c610b528484356001600160401b038111610b71575b8501610394565b93610a5c8160208601356001600160401b0381116104a4578601610394565b610b79600080fd5b610b4b565b610b86600080fd5b610b33565b602080825261027c92910190610555565b5034610bc6575b610637610bba610bb4366004610b26565b916136bf565b60405191829182610b8b565b610bce600080fd5b610ba3565b90816080910312610be15790565b61027c600080fd5b909182601f83011215610c2f575b60208235926001600160401b038411610c22575b019260208302840111610c1a57565b610154600080fd5b610c2a600080fd5b610c0b565b610c37600080fd5b610bf7565b90816040910312610be15790565b90606082820312610cda575b610c728183356001600160401b038111610ccd575b8401610bd3565b9261027c610c958360208601356001600160401b038111610cc0575b8601610be9565b9390946040810135906001600160401b038211610cb3575b01610c3c565b610cbb600080fd5b610cad565b610cc8600080fd5b610c8e565b610cd5600080fd5b610c6b565b610ce2600080fd5b610c56565b901515815260406020820181905261027c929101906104e5565b5034610d35575b610d1f610d16366004610c4a565b929190916143d5565b90610637610d2c60405190565b92839283610ce7565b610d3d600080fd5b610d08565b90929192610d5261023c826101b9565b9381855260208086019202830192818411610d90575b915b838310610d775750505050565b60208091610d858486610663565b815201920191610d6a565b610d98600080fd5b610d68565b9061027c9181601f82011215610db9575b602081359101610d42565b610dc1600080fd5b610dae565b6001600160401b038116610140565b9050359061015482610dc6565b65ffffffffffff8116610140565b9050359061015482610de2565b919091608081840312610e78575b610e65610e1860806101a5565b93610e358184356001600160401b038111610e6b575b8501610d9d565b85526020610e4582858301610dd5565b908601526040610e5782828601610663565b908601526060809301610df0565b90830152565b610e73600080fd5b610e2e565b610e80600080fd5b610e0b565b801515610140565b9050359061015482610e85565b919091608081840312610f20575b610e65610eb560806101a5565b93610ed28184356001600160401b038111610f13575b85016109f9565b8552610ef28160208501356001600160401b03811161082e578501610260565b60208601526040610f0582828601610df0565b908601526060809301610e8d565b610f1b600080fd5b610ecb565b610f28600080fd5b610ea8565b919091606081840312610f75575b610e65610f4860606101a5565b936000610f5582856101e6565b908601526020610f6782828601610147565b908601526040809301610147565b610f7d600080fd5b610f3b565b90929192610f9261023c826101b9565b938185526060602086019202830192818411610fd2575b915b838310610fb85750505050565b6020606091610fc78486610f2d565b815201920191610fab565b610fda600080fd5b610fa9565b9061027c9181601f82011215610ffb575b602081359101610f82565b611003600080fd5b610ff0565b91906108d090604084820312611077575b61102360406101a5565b936110408282356001600160401b03811161106a575b8301610e9a565b85526020810135906001600160401b03821161105d575b01610fdf565b611065600080fd5b611057565b611072600080fd5b611039565b61107f600080fd5b611019565b9092919261109461023c826101b9565b93818552602080860192028301928184116110f3575b80925b8484106110bb575050505050565b602080916110db8587356001600160401b0381116110e6575b8601611008565b8152019301926110ad565b6110ee600080fd5b6110d4565b6110fb600080fd5b6110aa565b9061027c9181601f8201121561111c575b602081359101611084565b611124600080fd5b611111565b90916060828403126111b8575b61027c6111558484356001600160401b0381116111ab575b8501610dfd565b936111758160208601356001600160401b03811161119e575b8601611100565b936040810135906001600160401b038211611191575b01611008565b611199600080fd5b61118b565b6111a6600080fd5b61116e565b6111b3600080fd5b61114e565b6111c0600080fd5b611136565b50346111e0575b6107b86111da366004611129565b91611a8e565b6111e8600080fd5b6111cc565b60c081830312611279575b6112148282356001600160401b03811161126c575b8301610dfd565b9261027c6112378460208501356001600160401b03811161125f575b8501611100565b936112568160408601356001600160401b0381116110e6578601611008565b93606001610f2d565b611267600080fd5b611230565b611274600080fd5b61120d565b611281600080fd5b6111f8565b50346112a4575b6107b861129b3660046111ed565b929190916118e1565b6112ac600080fd5b61128d565b600061071561027c92826106df565b50346112dd575b6106376107386112d8366004610aa7565b6112b1565b6112e5600080fd5b6112c7565b919091610100818403126113b0575b610e656113076101006101a5565b9360006113148285610147565b90860152602061132682828601610147565b908601526113488160408501356001600160401b03811161082e578501610260565b6040860152606061135b82828601610147565b90860152608061136d82828601610147565b9086015260a061137f82828601610147565b908601526113a18160c08501356001600160401b03811161082e578501610260565b60c086015260e0809301610147565b6113b8600080fd5b6112f9565b9061027c916020818303126113f2575b8035906001600160401b0382116113e5575b016112ea565b6113ed600080fd5b6113df565b6113fa600080fd5b6113cd565b5034611419575b6107b86114143660046113bd565b61338e565b611421600080fd5b611406565b919061027c9060408482031261147b575b6114538185356001600160401b03811161146e575b8601610dfd565b936020810135906001600160401b0382116111915701611008565b611476600080fd5b61144c565b611483600080fd5b611437565b50346114a3575b6107b861149d366004611426565b90613e92565b6114ab600080fd5b61148f565b50346114cb575b6107b86114c5366004611426565b90611bb7565b6114d3600080fd5b6114b7565b6114e59060005b50614638565b909192565b50634e487b7160e01b600052602160045260246000fd5b6003111561150b57565b6101546114ea565b9061015482611501565b602080825261027c929101906104e5565b156115365750565b6115589061154360405190565b62461bcd60e51b81529182916004830161151d565b0390fd5b61156961027c61027c9290565b65ffffffffffff1690565b50634e487b7160e01b600052601160045260246000fd5b6115a49065ffffffffffff165b9165ffffffffffff1690565b019065ffffffffffff82116115b557565b610154611574565b6004111561150b57565b90610154826115bd565b61027c906115c7565b6104be906115d1565b61027c9160206040820192611600600082015160008501906115da565b01519060208184039101526104e5565b9061161c610561835190565b908161162e6020830284019460200190565b926000915b83831061164257505050505090565b9091929394602061165e6105ad8385600195038752895161054b565b9301930191939290611633565b80516001600160a01b0316825261027c91604061169760608301602085015184820360208601526115e3565b920151906040818403910152611610565b9061027c9161166b565b906116be610561835190565b90816116d06020830284019460200190565b926000915b8383106116e457505050505090565b909192939460206117006105ad838560019503875289516116a8565b93019301919392906116d5565b9061027c9060608061174361173160808501600088015186820360008801526116b2565b602087015185820360208701526104e5565b60408087015165ffffffffffff16908501529401511515910152565b805160ff1682526101549190604090819061177f60208201516020860152565b0151910152565b906102118160609361175f565b906117b36117ac6117a2845190565b8084529260200190565b9260200190565b9060005b8181106117c45750505090565b9091926117de6117d76001928651611786565b9460200190565b9291016117b7565b8051604080845261027c9391602091611802919084019061170d565b920151906020818403910152611793565b9061027c916117e6565b90611829610561835190565b908161183b6020830284019460200190565b926000915b83831061184f57505050505090565b9091929394602061186b6105ad83856001950387528951611813565b9301930191939290611840565b65ffffffffffff909116815261027c9290916118a0906060840190848203602086015261181d565b9160408184039101526117e6565b61027c60806101a5565b90600019905b9181191691161790565b906118d661027c6118dd9290565b82546118b8565b9055565b6118ea816148ef565b83516040015190919065ffffffffffff1661190483614499565b9383600096879661191488611513565b9061191e90611513565b1460001498611a31611a06611a43996119fd611a3d996119f76101549f996102f798611a389b611a48576119528d8c612492565b6119666119608284876120f3565b9061152e565b8581019a6119836119788d51876149ac565b9a888701518c611d94565b606061198e4261155c565b9501926119aa6119a4855165ffffffffffff1690565b8761158b565b6119e86119d57f0aa12461ee6c137332989aa12cec79f4772ab2c1a8732a382aada7e9f3ec9d349490565b946119df60405190565b93849384611878565b0390a25165ffffffffffff1690565b9061158b565b95510151614a08565b93611a22611a126118ae565b65ffffffffffff9098168c890152565b65ffffffffffff166020870152565b6040850152565b614535565b926106df565b6118c8565b611a518b614499565b611a64611a5e6001611513565b91611513565b03611a7857611a738d8c612420565b611952565b611a738b6124ed565b61027c61027c61027c9290565b91611960611ac991611a9f856148ef565b81516040015190949065ffffffffffff1695611aba866124ed565b611ac48787612420565b6120f3565b611ad281614499565b611b37611a5e6000611b32611b28611ae983611a81565b611a38611af46118ae565b65ffffffffffff8b168682015291611b1d611b0e8761155c565b65ffffffffffff166020850152565b6102f7816040850152565b611a4387846106df565b611513565b03611b8c57611b87611b677ff3f2d5574c50e581f1a2371fac7dee87f7c6d599a496765fbfa2547ce7fd5f1a9290565b92611b7160405190565b9182918265ffffffffffff909116815260200190565b0390a2565b611b87611b677f07da0a0674fb921e484018c8b81d80e292745e5d8ed134b580c8b9c631c5e9e09290565b90611bc191611c4f565b50565b15611bcb57565b60405162461bcd60e51b815260206004820152601360248201527214dd185d19481b5d5cdd08189948199a5b985b606a1b6044820152606490fd5b61027c61027c61027c9260ff1690565b15611c1d57565b60405162461bcd60e51b815260206004820152600a60248201526921756e616e696d6f757360b01b6044820152606490fd5b9190611c5a836148ef565b611cb38194611c68836124ed565b835160600151611c79901515611bc4565b611cad611ca761027c6000611ca0611c9b6020611c968b896122d0565b015190565b614805565b9401515190565b91611c06565b14611c16565b611d1d6000611a4383611a3d611cc84261155c565b96611a38611ce48680611cda81611a81565b9401510151614a08565b6102f7611cef6118ae565b93611d0a611cfc8a61155c565b65ffffffffffff16868b0152565b65ffffffffffff8c166020860152611a31565b611b87611b677f4f465027a3d06ea73dd12be0f5c5fc0a34e21f19d6eaed4834a7a944edabc9019290565b15611d4f57565b60405162461bcd60e51b815260206004820152601f60248201527f4368616c6c656e676572206973206e6f742061207061727469636970616e74006044820152606490fd5b90611dfe61015493611e0393611de6611dac60405190565b602081019283526040808201526009606082015268666f7263654d6f766560b81b608082015291829060a082015b90810382520382610177565b611df8611df1825190565b9160200190565b20614746565b611e5f565b611d48565b6001906000198114611e18570190565b610211611574565b50634e487b7160e01b600052603260045260246000fd5b9060208091611e44845190565b811015611e52575b02010190565b611e5a611e20565b611e4c565b600091611e6b83611a81565b611e7661027c835190565b811015611ec057611e9a610648611e8d8385611e37565b516001600160a01b031690565b6001600160a01b03841614611eb757611eb290611e08565b611e6b565b50505050600190565b50505090565b9050519061015482610e85565b92919061015491611ee661023c836101f3565b94828652602086019183820111156104c257611f00600080fd5b6104c2565b9061027c9181601f82011215611f21575b602081519101611ed3565b611f29600080fd5b611f16565b919061027c90604084820312611f72575b611f498185611ec6565b936020810151906001600160401b038211611f65575b01611f05565b611f6d600080fd5b611f5f565b611f7a600080fd5b611f3f565b90611f8e6117ac6117a2845190565b9060005b818110611f9f5750505090565b909192611fbe6117d760019286516001600160a01b0316815260200190565b929101611f92565b9061027c90606080611fe76080840160008701518582036000870152611f7f565b6020808701516001600160401b031690850152946040818101516001600160a01b031690850152015165ffffffffffff16910152565b9061027c9060208061203e604084016000870151858203600087015261170d565b940151910152565b9061027c9161201d565b9061205c610561835190565b908161206e6020830284019460200190565b926000915b83831061208257505050505090565b9091929394602061209e6105ad83856001950387528951612046565b9301930191939290612073565b916120d8906120ca61027c959360608601908682036000880152611fc6565b908482036020860152612050565b91604081840391015261201d565b506040513d6000823e3d90fd5b91929160009161214861211b61211661211660408601516001600160a01b031690565b6106be565b91612153612138612131639936d812938761224f565b98866122d0565b6040519889968795869560e01b90565b8552600485016120ab565b03915afa801561218e575b600092839161216c57509190565b905061218a91923d8091833e6121828183610177565b810190611f2e565b9091565b6121966120e6565b61215e565b906121a861023c836101b9565b918252565b61027c60406101a5565b6121bf6118ae565b906060825260208080808501606081520160008152016000905250565b61027c6121b7565b6121ec6121ad565b906121f56121dc565b825260006020830152565b61027c6121e4565b60005b82811061221757505050565b602090612222612200565b818401520161220b565b9061015461224261223c8461219b565b936101b9565b601f190160208401612208565b61225f61225a835190565b61222c565b9161226a6000611a81565b61227561027c835190565b811015611ec0578061229461228d6122af9385611e37565b51856122d0565b61229e8287611e37565b526122a98186611e37565b50611e08565b61226a565b60ff81116122c3575b60020a90565b6122cb611574565b6122bd565b916122d9612200565b5081519060006122f16122ea6121ad565b9382850152565b6122fa81611a81565b91612306836020860152565b82945b6020810161231961027c82515190565b8710156123c9576123439061233c88612335878601518c6149ac565b9251611e37565b5190614746565b93805b84890161235561027c82515190565b8210156123bb57610648611e8d8361236d9351611e37565b6001600160a01b0387161461238a5761238590611e08565b612346565b6123b5929791955061239e6123b0916122b4565b60208801906123ab825190565b179052565b611e08565b94612309565b50509350946123b590611e08565b505094505050905090565b156123db57565b60405162461bcd60e51b815260206004820152601c60248201527f7475726e4e756d5265636f7264206e6f7420696e637265617365642e000000006044820152606490fd5b9061244061159861243361015494614638565b505065ffffffffffff1690565b116123d4565b1561244d57565b60405162461bcd60e51b815260206004820152601860248201527f7475726e4e756d5265636f7264206465637265617365642e00000000000000006044820152606490fd5b906124a561159861243361015494614638565b1015612446565b156124b357565b60405162461bcd60e51b815260206004820152601260248201527121b430b73732b6103334b730b634bd32b21760711b6044820152606490fd5b6124f961015491614499565b612506611a5e6002611513565b14156124ac565b1561251457565b60405162461bcd60e51b815260206004820152601f60248201527f4465706f73697420746f2065787465726e616c2064657374696e6174696f6e006044820152606490fd5b61027c9081565b61027c9054612559565b1561257157565b60405162461bcd60e51b81526020600482015260146024820152731a195b1908084f48195e1c1958dd195912195b1960621b6044820152606490fd5b61064861027c61027c9290565b61027c906125ad565b156125ca57565b60405162461bcd60e51b815260206004820152601f60248201527f496e636f7272656374206d73672e76616c756520666f72206465706f736974006044820152606490fd5b919061261a565b9290565b82018092116115b557565b906118d661027c6118dd92611a81565b6001600160a01b0390911681526040810192916101549160200152565b0152565b92836126e3836107106126db6126e89585979861268161267c61267887613b9b565b1590565b61250d565b6126a661269f61261661269a8861071060019c8d6106c7565b612560565b821461256a565b6126b361064860006125ba565b6001600160a01b038a1603612720576126d682346126d0565b9190565b146125c3565b61260f565b9586946106c7565b612625565b7f87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b9192611b8761271760405190565b92839283612635565b6126d68261272d8b6106be565b33612737306106be565b91612787565b61275661275061027c9263ffffffff1690565b60e01b90565b6001600160e01b03191690565b6001600160a01b039182168152911660208201526060810192916101549160400152565b906127cc906127bd610154956004956127a36323b872dd61273d565b936127ad60405190565b9788956020870190815201612763565b60208201810382520383610177565b612895565b906121a861023c836101f3565b6127e860206127d1565b7f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602082015290565b61027c6127de565b9061027c9160208183031215611ec657612831600080fd5b611ec6565b1561283d57565b60405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b906128a26128b1926106be565b906128ab612811565b916128e9565b80516128c06126cc6000611a81565b116128c85750565b6128e48160206128d9610154945190565b818301019101612819565b612836565b61027c92916128f86000611a81565b91612978565b1561290557565b60405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608490fd5b3d15612973576129683d6127d1565b903d6000602084013e565b606090565b90600061027c9493819261298a606090565b506129a1612997306106be565b83903110156128fe565b60208101905191855af16129b3612959565b91612a05565b156129c057565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b919260609115612a3a5750508151612a206126cc6000611a81565b14612a29575090565b612a3561027c91612a47565b6129b9565b9093926101549250612a59565b3b612a556126cc6000611a81565b1190565b9150612a63825190565b612a706126cc6000611a81565b11156115365750805190602001fd5b90612ac26101549594612ab6612a9b612ac79683868a89612acd565b9282969291996040612aad8a85611e37565b51015185612ef0565b96929491509a89613280565b611e37565b51613347565b94926000612b1b612b148597612b0f61027c97612b2a96612af961269a99612af3606090565b50613d6c565b612b0282613cc5565b8351602085012090613c4d565b612e52565b9788611e37565b5101516001600160a01b031690565b936107108560016106c7565b9050519061015482610654565b9050519061015482610873565b91906108d090604084820312612b96575b612b6b60406101a5565b936000612b788383612b43565b908601526020810151906001600160401b038211611f655701611f05565b612b9e600080fd5b612b61565b905051906101548261013e565b90505190610154826101dd565b91906102f790608084820312612c27575b612bd860806101a5565b936000612be58383612ba3565b908601526020612bf783828401612ba3565b908601526040612c0983828401612bb0565b908601526060810151906001600160401b038211611f655701611f05565b612c2f600080fd5b612bce565b90929192612c4461023c826101b9565b9381855260208086019202830192818411612ca3575b80925b848410612c6b575050505050565b60208091612c8b8587516001600160401b038111612c96575b8601612bbd565b815201930192612c5d565b612c9e600080fd5b612c84565b612cab600080fd5b612c5a565b9061027c9181601f82011215612ccc575b602081519101612c34565b612cd4600080fd5b612cc1565b919061094f90606084820312612d5e575b612cf460606101a5565b936000612d018383612b36565b90860152612d248260208301516001600160401b038111612d51575b8301612b50565b60208601526040810151906001600160401b038211612d44575b01612cb0565b612d4c600080fd5b612d3e565b612d59600080fd5b612d1d565b612d66600080fd5b612cea565b90929192612d7b61023c826101b9565b9381855260208086019202830192818411612dda575b80925b848410612da2575050505050565b60208091612dc28587516001600160401b038111612dcd575b8601612cd9565b815201930192612d94565b612dd5600080fd5b612dbb565b612de2600080fd5b612d91565b9061027c9181601f82011215612e03575b602081519101612d6b565b612e0b600080fd5b612df8565b9061027c91602081830312612e45575b8051906001600160401b038211612e38575b01612de7565b612e40600080fd5b612e32565b612e4d600080fd5b612e20565b61027c906020612e60825190565b818301019101612e10565b612e736118ae565b906000825260208080808501600081520160005b8152016060905250565b61027c612e6b565b60005b828110612ea857505050565b602090612eb3612e91565b8184015201612e9c565b90610154612ecd61223c8461219b565b601f190160208401612e99565b61027c90611513565b919082039182116115b557565b91805193600094612f036126cc87611a81565b11156131ee57612f1c612f14835190565b939293612ebd565b9283612f2787611a81565b968794612f3a612f35825190565b612ebd565b9384996001938491829b97819882945b612f5b575b50505050505050505050565b9091929394959697989a612f7061027c895190565b8610156131e857612f9e612f8e86612f88898c611e37565b51015190565b86612f998987611e37565b510152565b612fcf612fb96040612fb0898c611e37565b51015160ff1690565b6040612fc58987611e37565b51019060ff169052565b6060612fdb878a611e37565b5101516060612fea8886611e37565b510152613005826130006020612f888a8d611e37565b613de7565b908461301261027c8a5190565b1480156131af575b1561318a5761302e6040612fb08f8c611e37565b61304b61304461303e6002612eda565b60ff1690565b9160ff1690565b14613145576130f8826130fe928f6130f2908f8b8f8e6102f78261308f8f6020612f99866130898f61308485612f88856130e29e611e37565b612ee3565b93611e37565b6130cd60606130bb6130a588612f888887611e37565b956130b56040612fb08388611e37565b94611e37565b510151936130c76118ae565b96870152565b6130d88a6020870152565b60ff166040850152565b6130ec8383611e37565b52611e37565b5061260f565b9c611e08565b915b8461311361027c6020612f888b89611e37565b0361313b575b61312c9161312691612ee3565b95611e08565b93929190989796959487612f4a565b949d508d94613119565b60405162461bcd60e51b815260206004820152601b60248201527f63616e6e6f74207472616e7366657220612067756172616e74656500000000006044820152606490fd5b9b916131aa61319e6020612f888a8d611e37565b6020612f998a88611e37565b613100565b508c6131bf6126cc61027c8b5190565b108d816131cd575b5061301a565b6131e091506131dc908a611e37565b5190565b87148d6131c7565b9a612f4f565b612f1c612f14845190565b90613205610561835190565b90816132176020830284019460200190565b926000915b83831061322b57505050505090565b909192939460206132476105ad838560019503875289516116a8565b930193019193929061321c565b602080825261027c929101906131f9565b908152606081019392610154929091604091612652906108d0565b9690806132f961269a956132fe97869a9960406132c78e6132c16107109a6132bb6132b18e61071060019d8e6106c7565b9161308483612560565b90612625565b84611e37565b5101526132e66132d660405190565b8092611dda602083019182613254565b6132f1611df1825190565b209086613ce4565b6106c7565b92611b8761332a7fc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a9490565b9461333460405190565b93849384613265565b61027c60606101a5565b610154916133899061094f602061336585516001600160a01b031690565b94015161338261337361333d565b6001600160a01b039096168652565b6020850152565b6139af565b610154906133dd61339e8261347b565b90919060406133d0816133bb6133b5606089015190565b87611e37565b510151926133ca60e088015190565b90611e37565b5101516080850151610bb4565b9161391d565b156133ea57565b60405162461bcd60e51b815260206004820152601a60248201527f6e6f7420612067756172616e74656520616c6c6f636174696f6e0000000000006044820152606490fd5b1561343657565b60405162461bcd60e51b815260206004820152601d60248201527f746172676574417373657420213d2067756172616e74656541737365740000006044820152606490fd5b90613484825190565b90604083015160608401613496905190565b9260c08501519160e086016134a9905190565b916134b381613cc5565b6020870151906134c1835190565b60208401206134cf92613c4d565b6134d890612e52565b94856134e384612e52565b80966134ef8184611e37565b51516001600160a01b0316926135058282611e37565b5160400151916080860192613518845190565b61352191611e37565b516040015160ff166135336002612eda565b60ff169060ff1614613544906133e3565b61354d91611e37565b5160400151905161355d91611e37565b51519361356991611e37565b5151613582916001600160a01b0390811691161461342f565b61358b82613cc5565b60a0015191613598815190565b906020012061015492613c4d565b156135ad57565b60405162461bcd60e51b815260206004820152601560248201527418dbdd5b19081b9bdd08199a5b99081d185c99d95d605a1b6044820152606490fd5b156135f157565b60405162461bcd60e51b815260206004820152601360248201527218dbdd5b19081b9bdd08199a5b99081b19599d606a1b6044820152606490fd5b1561363357565b60405162461bcd60e51b815260206004820152601460248201527318dbdd5b19081b9bdd08199a5b99081c9a59da1d60621b6044820152606490fd5b1561367657565b60405162461bcd60e51b815280611558600482016020808252818101527f746f74616c5265636c61696d6564213d67756172616e7465652e616d6f756e74604082015260600190565b9192918051906136dd612f356001936136d785611a81565b90612ee3565b926136e88683611e37565b51916136f76060840151613e70565b966000948592818795889b61370a600090565b9961371481611a81565b80945b613756575b5050505050505050602061375093611c9661027c97989961374b6126cc9661374661027c976135a6565b6135ea565b61362c565b1461366f565b61376161027c875190565b851015613918578e8d89871461390557826137d791856130e28a6102f78d6130d861379086612f888685611e37565b936137d060606137be6137a86020612f888689611e37565b936137b86040612fb0838a611e37565b96611e37565b510151956137ca6118ae565b98890152565b6020870152565b508d8b15806138e0575b613895575b50155b8061386f575b61380e575b9061380461312688959493611e08565b9490919293613717565b9d50999061385e8693928f8e6104be602061383c613847946130b583612f886138368e611a81565b8d611e37565b5101916126d6835190565b6138586020612f886133b588611a81565b9061260f565b929d929a8e965091929091906137f4565b5061387e82612f888789611e37565b61388f6126cc61027c602087015190565b146137ef565b829d919b506138bb906104be602061383c6138d2966130b583612f888d6133ca8d611a81565b6138586020612f886138cc86611a81565b88611e37565b9a6137e9879a90508d6137e6565b506138ef84612f88898b611e37565b6138ff6126cc61027c8789015190565b146137e1565b5050939750908592916138048499611e08565b61371c565b613979916131dc613974926000810192613955602061393a865190565b936060810199604061394d6133b58d5190565b510152015190565b906139626132d660405190565b61396d611df1825190565b2091613ce4565b915190565b90611b876139a57f4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b9290565b9261073c60405190565b80516001600160a01b0316916000926139c784611a81565b604084016139d761027c82515190565b821015613a4d578181613a046020612f88846139fd8c612f8860409a613a279a51611e37565b9451611e37565b613a0d82613b9b565b15613a2e57613a216121166123b093613be4565b86613a9a565b90506139c7565b6132bb613a436123b0936107108960016106c7565b916126d683612560565b505050915050565b15613a5c57565b60405162461bcd60e51b8152602060048201526016602482015275086deead8c840dcdee840e8e4c2dce6cccae4408aa8960531b6044820152606490fd5b90600091613aaa610648846125ba565b6001600160a01b03821603613ade5750819061015493613ac960405190565b90818003925af1613ad8612959565b50613a55565b91613b0f613af06121166020956106be565b9163a9059cbb613b1a613b0260405190565b9788968795869460e01b90565b845260048401612635565b03925af18015613b53575b613b2c5750565b611bc19060203d8111613b4c575b613b448183610177565b810190612819565b503d613b3a565b613b5b6120e6565b613b25565b613b7661027c61027c926001600160601b031690565b6001600160601b031690565b61027c9060a01c613b60565b613b7661027c61027c9290565b613bb8613bbd91613baa600090565b506001600160a01b03191690565b613b82565b613bd7613bca6000613b8e565b916001600160601b031690565b1490565b61027c90611a81565b612116613c04613bff61027c93613bf9600090565b50613bdb565b6125ad565b6106b5565b15613c1057565b60405162461bcd60e51b81526020600482015260156024820152741a5b98dbdc9c9958dd08199a5b99d95c9c1c9a5b9d605a1b6044820152606490fd5b90613c6d610648613c7a92613c6461015496614638565b969150506145e5565b916001600160a01b031690565b14613c09565b15613c8757565b60405162461bcd60e51b815260206004820152601660248201527521b430b73732b6103737ba103334b730b634bd32b21760511b6044820152606490fd5b613cd161015491614499565b613cde611a5e6002611513565b14613c80565b90613d1f611a31611a38610154956102f7611a4395613d0288614638565b50611a22613d119792976118ae565b65ffffffffffff9098168852565b9160006106df565b15613d2e57565b60405162461bcd60e51b8152602060048201526016602482015275125b991a58d95cc81b5d5cdd081899481cdbdc9d195960521b6044820152606490fd5b613d766000611a81565b600190613d8b613d8583611a81565b8261260f565b613d996126cc61027c865190565b1015613de2576123b0613ddd92613dd76126cc61027c6131dc613dd1613dcb613dc56131dc8a8d611e37565b96611a81565b8861260f565b89611e37565b10613d27565b613d76565b505050565b81811115613df3575090565b905090565b613e006121ad565b9060006121f5565b61027c613df8565b919091604081840312613e46575b610e65613e2b60406101a5565b936000613e388285612ba3565b908601526020809301612ba3565b613e4e600080fd5b613e1e565b9061027c9160408183031215613e1057613e6b600080fd5b613e10565b61027c90613e7c613e08565b506020613e87825190565b818301019101613e53565b90600080613ea38361015495611c4f565b9201510151610a946000611a81565b613eba6121ad565b906000825260606020830152565b61027c613eb2565b613ed861333d565b9060008252602080808401612e87613ec8565b61027c613ed0565b60005b828110613f0257505050565b602090613f0d613eeb565b8184015201613ef6565b90610154613f2761223c8461219b565b601f190160208401613ef3565b369037565b90610154613f4961223c8461219b565b601f190160208401613f34565b9160001960089290920291821b911b6118be565b9190613f7961027c6118dd9390565b908354613f56565b61015491600091613f6a565b929192613f9981613cc5565b613fac81613fa684614a08565b86613c4d565b6001908194613fc1613fbc855190565b613f17565b93613fd2613fcd825190565b613f39565b96613fde613fcd835190565b9460009581613fec88611a81565b905b61412b575b5081613ffe88611a81565b905b614043575b50505061015496975060001461402f5750509061402561402a92826106df565b613f81565b614446565b61402a935061403d90614a08565b91613ce4565b61404e61027c865190565b81101561412657908188614063859488611e37565b5101516001600160a01b0316888d8361407c8187611e37565b5183614088868a6106c7565b90614092916106df565b9061409c82612560565b906140a691612ee3565b6140af91612625565b6140b891611e37565b51916140c490866106c7565b906140ce916106df565b6140d790612560565b907fc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a908a9261410560405190565b918291614113918784613265565b0390a261411f90611e08565b9091614000565b614005565b8261413761027c875190565b821015614236575087858c61414c8483611e37565b5160408101518461415d8786611e37565b5101516001600160a01b031692868d614176868c6106c7565b90614180916106df565b61418990612560565b6141938284611e37565b5261419d91611e37565b51906141a886611a81565b6141b190613f39565b906141bb92612ef0565b9095911561422c575b936133826020614211958a989560406141f28f9d8f98612ac26142269f9d6141ef8461094f9d611e37565b52565b51015201519161420061333d565b968701906001600160a01b03169052565b61421b828d611e37565b526122a9818c611e37565b90613fee565b95995089956141c4565b50613ff3565b3561027c81610654565b61027c903690610dfd565b61027c913691611084565b61027c903690611008565b9035601e1936839003018112156142ab575b01602081359101916001600160401b03821161429e575b6020820236038313610c1a57565b6142a6600080fd5b614290565b6142b3600080fd5b614279565b5061027c906020810190610663565b818352602090920191906000825b8282106142e3575050505090565b9091929361431261430b6001926142fa88866142b8565b6001600160a01b0316815260200190565b9560200190565b939201906142d5565b5061027c906020810190610dd5565b5061027c906020810190610df0565b9061027c9060606143a961435f608084016143548780614267565b8683038752906142c7565b94614380614370602083018361431b565b6001600160401b03166020860152565b6143a061439060408301836142b8565b6001600160a01b03166040860152565b8281019061432a565b65ffffffffffff16910152565b916120d8906120ca61027c959360608601908682036000880152614339565b9061443b6000939594956143e7600090565b5061215361213861442261440361211661211660408a0161423c565b9561441c639936d812956144168a614246565b92614251565b9061224f565b9861443561442f88614246565b9161425c565b906122d0565b8552600485016143b6565b906144516000611a81565b61445c61027c845190565b81101561447f57806123b061447461447a9386611e37565b516139af565b614451565b509050565b61027c61027c61027c9265ffffffffffff1690565b6144a49060006114df565b50600091506144b28261155c565b65ffffffffffff8216036144c4575090565b90506144d04291614484565b116144da57600290565b600190565b6144ec61027c61027c9290565b61ffff1690565b61ffff908116911690039061ffff82116115b557565b61027c9061451e6126cc61027c9461ffff1690565b901b90565b61027c9081906001600160a01b031681565b61027c906145cc6145c761454a6101006144df565b6145a9614565614560865165ffffffffffff1690565b614484565b6145a361458661457f61457860306144df565b80966144f3565b8093614509565b9361459d61456060208a015165ffffffffffff1690565b926144f3565b90614509565b17926145c160606145bb604084015190565b92015190565b906145e5565b614523565b17611a81565b9081526040810192916101549160200152565b611dda61460d61027c93613bff936145fb600090565b506040519384926020840192836145d2565b614618611df1825190565b20613bdb565b61027c906146336126cc61027c9461ffff1690565b901c90565b61269a61464f91614647600090565b5060006106df565b9061466461465e6101006144df565b92613bdb565b9061027c6146a461469361469e61468561467e60306144df565b80986144f3565b96614698614693898961461e565b61155c565b976144f3565b8561461e565b926125ad565b906102116121a86020937f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c0190565b6126526101549461094f6060949897956146fc608086019a6000870152565b60ff166020850152565b1561470d57565b60405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b6044820152606490fd5b60209160009161476761475860405190565b8092611dda87830191826146aa565b614772611df1825190565b206147a66147838484015160ff1690565b9261479360406145bb8884015190565b9061479d60405190565b948594856146dd565b838052039060015afa156147d9575b60005161027c6147c861064860006125ba565b6001600160a01b0383161415614706565b6147e16120e6565b6147b5565b61303e61027c61027c9290565b60019060ff1660ff8114611e18570190565b600090614811826147e6565b905b61481c83611a81565b81111561484c5761484461481c9161483d6148376001611a81565b82612ee3565b16926147f3565b919050614813565b50905090565b906148616117ac6117a2845190565b9060005b8181106148725750505090565b9091926148916117d760019286516001600160a01b0316815260200190565b929101614865565b6148e1610154946148d16148bf6060959998969960808601908682036000880152614852565b6001600160401b039099166020850152565b6001600160a01b03166040830152565b019065ffffffffffff169052565b80516149459061490960208401516001600160401b031690565b90611dda614934606061492660408801516001600160a01b031690565b96015165ffffffffffff1690565b604051958694602086019485614899565b614950611df1825190565b2090565b90614993610154959796946149856149a49360809661497860a08801926000890152565b86820360208801526104e5565b9084820360408601526131f9565b65ffffffffffff9097166060830152565b019015159052565b6149c1614945916149bb600090565b506148ef565b602083015190611dda6000850151946149f060606149e8604084015165ffffffffffff1690565b920151151590565b906149fa60405190565b968795602087019586614954565b6149459061027c6132d66040519056fea26469706673582212201b685687204d0d24b2447748cfb8573256c34fbe13ef0ad16636e4890bc34fb064736f6c63430008110033",
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
