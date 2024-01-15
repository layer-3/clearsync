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
	Bin: "0x608060405234801561001057600080fd5b506144dd806100206000396000f3fe6080604052600436106100dd5760003560e01c80635685b7dc1161007f578063c7df14e211610059578063c7df14e21461029f578063d3c4e738146102cc578063ec346235146102ec578063ee049b501461030c57600080fd5b80635685b7dc146102315780636d2a9c921461025f5780638286a0601461027f57600080fd5b80633033730e116100bb5780633033730e1461017657806331afa0b414610196578063552cfa50146101b6578063566d54c61461020457600080fd5b806311e9f178146100e2578063166e56cd1461011b5780632fb1d27014610161575b600080fd5b3480156100ee57600080fd5b506101026100fd366004612eaa565b61032c565b6040516101129493929190612feb565b60405180910390f35b34801561012757600080fd5b5061015361013636600461304f565b600160209081526000928352604080842090915290825290205481565b604051908152602001610112565b61017461016f36600461307b565b61076c565b005b34801561018257600080fd5b506101746101913660046130b6565b610918565b3480156101a257600080fd5b506101746101b1366004613277565b6109a4565b3480156101c257600080fd5b506101d66101d13660046132c6565b610df5565b6040805165ffffffffffff94851681529390921660208401526001600160a01b031690820152606001610112565b34801561021057600080fd5b5061022461021f3660046132df565b610e10565b604051610112919061333a565b34801561023d57600080fd5b5061025161024c36600461335f565b611246565b604051610112929190613425565b34801561026b57600080fd5b5061017461027a366004613775565b61130d565b34801561028b57600080fd5b5061017461029a3660046137f2565b611469565b3480156102ab57600080fd5b506101536102ba3660046132c6565b60006020819052908152604090205481565b3480156102d857600080fd5b506101746102e736600461388b565b6115f5565b3480156102f857600080fd5b5061017461030736600461395b565b611677565b34801561031857600080fd5b5061017461032736600461395b565b61169b565b606060006060600080855111610343578551610346565b84515b6001600160401b0381111561035d5761035d612bdf565b6040519080825280602002602001820160405280156103ae57816020015b604080516080810182526000808252602080830182905292820152606080820152825260001990920191018161037b5790505b5091506000905085516001600160401b038111156103ce576103ce612bdf565b60405190808252806020026020018201604052801561041f57816020015b60408051608081018252600080825260208083018290529282015260608082015282526000199092019101816103ec5790505b50935060019250866000805b885181101561076057888181518110610446576104466139be565b602002602001015160000151878281518110610464576104646139be565b60200260200101516000018181525050888181518110610486576104866139be565b6020026020010151604001518782815181106104a4576104a46139be565b60200260200101516040019060ff16908160ff16815250508881815181106104ce576104ce6139be565b6020026020010151606001518782815181106104ec576104ec6139be565b60200260200101516060018190525060006105248a8381518110610512576105126139be565b602002602001015160200151856116a5565b9050885160001480610559575088518310801561055957508189848151811061054f5761054f6139be565b6020026020010151145b156106d557600260ff168a8481518110610575576105756139be565b60200260200101516040015160ff16036105d65760405162461bcd60e51b815260206004820152601b60248201527f63616e6e6f74207472616e7366657220612067756172616e746565000000000060448201526064015b60405180910390fd5b808a83815181106105e9576105e96139be565b6020026020010151602001516105ff9190613a00565b888381518110610611576106116139be565b6020026020010151602001818152505060405180608001604052808b848151811061063e5761063e6139be565b60200260200101516000015181526020018281526020018b8481518110610667576106676139be565b60200260200101516040015160ff1681526020018b848151811061068d5761068d6139be565b6020026020010151606001518152508684815181106106ae576106ae6139be565b60209081029190910101526106c38186613a13565b94506106ce83613a26565b9250610716565b8982815181106106e7576106e76139be565b602002602001015160200151888381518110610705576107056139be565b602002602001015160200181815250505b878281518110610728576107286139be565b60200260200101516020015160001461074057600096505b61074a8185613a00565b935050808061075890613a26565b91505061042b565b50505093509350935093565b6107778360a01c1590565b156107c45760405162461bcd60e51b815260206004820152601f60248201527f4465706f73697420746f2065787465726e616c2064657374696e6174696f6e0060448201526064016105cd565b6001600160a01b038416600090815260016020908152604080832086845290915290205482811461082e5760405162461bcd60e51b81526020600482015260146024820152731a195b1908084f48195e1c1958dd195912195b1960621b60448201526064016105cd565b6001600160a01b0385166108905781341461088b5760405162461bcd60e51b815260206004820152601f60248201527f496e636f7272656374206d73672e76616c756520666f72206465706f7369740060448201526064016105cd565b6108a5565b6108a56001600160a01b0386163330856116bf565b6108af8282613a13565b6001600160a01b03861660008181526001602090815260408083208984528252918290208490558151928352820183905291925085917f87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b91015b60405180910390a25050505050565b600080600061092a888589888a611719565b925092509250600080600061095d84878d8151811061094b5761094b6139be565b6020026020010151604001518961032c565b935093505092506109748b868c8b8a888a8861179d565b610997868c81518110610989576109896139be565b60200260200101518361188d565b5050505050505050505050565b6109ad836118c5565b6109c0816109ba8461192a565b85611943565b81516001906000906001600160401b038111156109df576109df612bdf565b604051908082528060200260200182016040528015610a1857816020015b610a05612b9d565b8152602001906001900390816109fd5790505b509050600084516001600160401b03811115610a3657610a36612bdf565b604051908082528060200260200182016040528015610a5f578160200160208202803683370190505b509050600085516001600160401b03811115610a7d57610a7d612bdf565b604051908082528060200260200182016040528015610aa6578160200160208202803683370190505b50905060005b8651811015610c71576000878281518110610ac957610ac96139be565b602002602001015190506000816040015190506000898481518110610af057610af06139be565b602002602001015160000151905060016000826001600160a01b03166001600160a01b0316815260200190815260200160002060008c815260200190815260200160002054868581518110610b4757610b476139be565b602002602001018181525050600080600080610bbf8a8981518110610b6e57610b6e6139be565b60200260200101518760006001600160401b03811115610b9057610b90612bdf565b604051908082528060200260200182016040528015610bb9578160200160208202803683370190505b5061032c565b935093509350935082610bd15760009b505b80898981518110610be457610be46139be565b602002602001018181525050838e8981518110610c0357610c036139be565b6020026020010151604001819052506040518060600160405280866001600160a01b0316815260200188602001518152602001838152508b8981518110610c4c57610c4c6139be565b6020026020010181905250505050505050508080610c6990613a26565b915050610aac565b5060005b8651811015610db5576000878281518110610c9257610c926139be565b6020026020010151600001519050828281518110610cb257610cb26139be565b602002602001015160016000836001600160a01b03166001600160a01b0316815260200190815260200160002060008b81526020019081526020016000206000828254610cff9190613a00565b92505081905550887fc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a83868581518110610d3b57610d3b6139be565b602002602001015160016000866001600160a01b03166001600160a01b0316815260200190815260200160002060008e815260200190815260200160002054604051610d9a939291909283526020830191909152604082015260600190565b60405180910390a25080610dad81613a26565b915050610c75565b508315610dd057600087815260208190526040812055610de3565b610de38786610dde8961192a565b6119da565b610dec83611a40565b50505050505050565b6000806000610e0384611a80565b9196909550909350915050565b6060600060018551610e229190613a00565b6001600160401b03811115610e3957610e39612bdf565b604051908082528060200260200182016040528015610e8a57816020015b6040805160808101825260008082526020808301829052928201526060808201528252600019909201910181610e575790505b5090506000858481518110610ea157610ea16139be565b602002602001015190506000610eba8260600151611acd565b9050600080808080805b8c51811015611114578a8103610edd5760019550611102565b60405180608001604052808e8381518110610efa57610efa6139be565b60200260200101516000015181526020018e8381518110610f1d57610f1d6139be565b60200260200101516020015181526020018e8381518110610f4057610f406139be565b60200260200101516040015160ff1681526020018e8381518110610f6657610f666139be565b602002602001015160600151815250898381518110610f8757610f876139be565b602002602001018190525084158015610fc0575086600001518d8281518110610fb257610fb26139be565b602002602001015160000151145b15611042578b600081518110610fd857610fd86139be565b602002602001015160200151898381518110610ff657610ff66139be565b602002602001015160200181815161100e9190613a13565b9052508b518c90600090611024576110246139be565b6020026020010151602001518361103b9190613a13565b9250600194505b83158015611070575086602001518d8281518110611062576110626139be565b602002602001015160000151145b156110f4578b600181518110611088576110886139be565b6020026020010151602001518983815181106110a6576110a66139be565b60200260200101516020018181516110be9190613a13565b9052508b518c9060019081106110d6576110d66139be565b602002602001015160200151836110ed9190613a13565b9250600193505b816110fe81613a26565b9250505b8061110c81613a26565b915050610ec4565b508461115a5760405162461bcd60e51b815260206004820152601560248201527418dbdd5b19081b9bdd08199a5b99081d185c99d95d605a1b60448201526064016105cd565b8361119d5760405162461bcd60e51b815260206004820152601360248201527218dbdd5b19081b9bdd08199a5b99081b19599d606a1b60448201526064016105cd565b826111e15760405162461bcd60e51b815260206004820152601460248201527318dbdd5b19081b9bdd08199a5b99081c9a59da1d60621b60448201526064016105cd565b866020015182146112345760405162461bcd60e51b815260206004820181905260248201527f746f74616c5265636c61696d6564213d67756172616e7465652e616d6f756e7460448201526064016105cd565b509596505050505050505b9392505050565b6000606061125986820160408801613a3f565b6001600160a01b0316639936d8128761128361127482613a5c565b61127e898b613a68565b611af5565b61129d61128f8b613a5c565b61129889613a75565b611be7565b6040518463ffffffff1660e01b81526004016112bb93929190613cd6565b600060405180830381865afa1580156112d8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526113009190810190613de3565b9150915094509492505050565b600061131884611d01565b82516040015190915061132a82611d47565b6113348282611da5565b60006060611343878787611e13565b909250905080826113675760405162461bcd60e51b81526004016105cd9190613e3d565b50600061137385611eae565b90506113b560405180608001604052808665ffffffffffff168152602001600065ffffffffffff1681526020016000801b81526020016000801b815250611ef8565b6000868152602081905260408120919091558160028111156113d9576113d96139d4565b036114205760405165ffffffffffff8516815285907ff3f2d5574c50e581f1a2371fac7dee87f7c6d599a496765fbfa2547ce7fd5f1a9060200160405180910390a261145f565b60405165ffffffffffff8516815285907f07da0a0674fb921e484018c8b81d80e292745e5d8ed134b580c8b9c631c5e9e0906020015b60405180910390a25b5050505050505050565b600061147485611d01565b835160400151909150600061148883611eae565b6002811115611499576114996139d4565b036114ad576114a88282611f8f565b6114e1565b60016114b883611eae565b60028111156114c9576114c96139d4565b036114d8576114a88282611da5565b6114e182611d47565b600060606114f0888888611e13565b909250905080826115145760405162461bcd60e51b81526004016105cd9190613e3d565b506000611525898860000151611ffe565b9050611536818a600001518861204a565b847f0aa12461ee6c137332989aa12cec79f4772ab2c1a8732a382aada7e9f3ec9d348a60600151426115689190613e50565b8a8a60405161157993929190613eee565b60405180910390a26115d860405180608001604052808665ffffffffffff1681526020018b60600151426115ad9190613e50565b65ffffffffffff1681526020018381526020016115d18a600001516000015161192a565b9052611ef8565b600095865260208690526040909520949094555050505050505050565b600080611601836120f7565b91509150606060008385606001518151811061161f5761161f6139be565b60200260200101516040015190506000838660e0015181518110611645576116456139be565b602002602001015160400151905061166282828860800151610e10565b925050506116718484836122ef565b50505050565b6000611683838361236c565b82515190915061169690829060006109a4565b505050565b611696828261236c565b60008183116116b457826116b6565b815b90505b92915050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526116719085906124c1565b606060008061172787612596565b611730866118c5565b61174285858051906020012088611943565b61174b84612642565b925082888151811061175f5761175f6139be565b602090810291909101810151516001600160a01b03811660009081526001835260408082209982529890925296902054929895975091955050505050565b6001600160a01b0387166000908152600160209081526040808320898452909152812080548392906117d0908490613a00565b92505081905550828489815181106117ea576117ea6139be565b60200260200101516040018190525061182a86868660405160200161180f9190613f69565b604051602081830303815290604052805190602001206119da565b6001600160a01b03871660009081526001602090815260408083208984528252918290205482518b81529182018590529181019190915286907fc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a90606001611456565b6118c1604051806060016040528084600001516001600160a01b031681526020018460200151815260200183815250612658565b5050565b60026118d082611eae565b60028111156118e1576118e16139d4565b146119275760405162461bcd60e51b815260206004820152601660248201527521b430b73732b6103737ba103334b730b634bd32b21760511b60448201526064016105cd565b50565b600061193582612724565b805190602001209050919050565b600061194e82611a80565b6040805160208082018a90528183018990528251808303840181526060909201909252805191012090935091506119829050565b6001600160a01b0316816001600160a01b0316146116715760405162461bcd60e51b81526020600482015260156024820152741a5b98dbdc9c9958dd08199a5b99d95c9c1c9a5b9d605a1b60448201526064016105cd565b6000806119e685611a80565b50915091506000611a2660405180608001604052808565ffffffffffff1681526020018465ffffffffffff16815260200187815260200186815250611ef8565b600096875260208790526040909620959095555050505050565b60005b81518110156118c157611a6e828281518110611a6157611a616139be565b6020026020010151612658565b80611a7881613a26565b915050611a43565b60008181526020819052604081205481908190610100611aa1603082613f7c565b61ffff811683901c95509050611ab8603082613f7c565b949661ffff90951682901c9550909392505050565b6040805180820190915260008082526020820152818060200190518101906116b99190613f97565b6060600082516001600160401b03811115611b1257611b12612bdf565b604051908082528060200260200182016040528015611b7d57816020015b611b6a6040805160c08101825260609181018281528282019290925260006080820181905260a08201529081908152602001600081525090565b815260200190600190039081611b305790505b50905060005b8351811015611bdf57611baf85858381518110611ba257611ba26139be565b6020026020010151611be7565b828281518110611bc157611bc16139be565b60200260200101819052508080611bd790613a26565b915050611b83565b509392505050565b611c216040805160c08101825260609181018281528282019290925260006080820181905260a08201529081908152602001600081525090565b60408051808201909152825181526000602082018190525b836020015151811015611bdf576000611c7c611c59878760000151611ffe565b86602001518481518110611c6f57611c6f6139be565b602002602001015161274d565b905060005b865151811015611cec578651805182908110611c9f57611c9f6139be565b60200260200101516001600160a01b0316826001600160a01b031603611cda57611cca8160026140ad565b6020850180519091179052611cec565b80611ce481613a26565b915050611c81565b50508080611cf990613a26565b915050611c39565b60008160000151826020015183604001518460600151604051602001611d2a94939291906140f2565b604051602081830303815290604052805190602001209050919050565b6002611d5282611eae565b6002811115611d6357611d636139d4565b036119275760405162461bcd60e51b815260206004820152601260248201527121b430b73732b6103334b730b634bd32b21760711b60448201526064016105cd565b6000611db083611a80565b505090508065ffffffffffff168265ffffffffffff16116116965760405162461bcd60e51b815260206004820152601c60248201527f7475726e4e756d5265636f7264206e6f7420696e637265617365642e0000000060448201526064016105cd565b6000606084604001516001600160a01b0316639936d81286611e358888611af5565b611e3f8988611be7565b6040518463ffffffff1660e01b8152600401611e5d9392919061413b565b600060405180830381865afa158015611e7a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ea29190810190613de3565b91509150935093915050565b600080611eba83611a80565b509150508065ffffffffffff16600003611ed75750600092915050565b428165ffffffffffff1611611eef5750600292915050565b50600192915050565b600080610100611f09603082613f7c565b845165ffffffffffff1661ffff82161b92509050611f28603082613f7c565b90508061ffff16846020015165ffffffffffff16901b82179150611f7b84604001518560600151604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b6001600160a01b0316919091179392505050565b6000611f9a83611a80565b505090508065ffffffffffff168265ffffffffffff1610156116965760405162461bcd60e51b815260206004820152601860248201527f7475726e4e756d5265636f7264206465637265617365642e000000000000000060448201526064016105cd565b600061200983611d01565b60208084015184516040808701516060880151915161202c9695919291016141a3565b60405160208183030381529060405280519060200120905092915050565b600061209f8460405160200161208391815260406020820181905260099082015268666f7263654d6f766560b81b606082015260800190565b604051602081830303815290604052805190602001208361274d565b90506120ab8184612868565b6116715760405162461bcd60e51b815260206004820152601f60248201527f4368616c6c656e676572206973206e6f742061207061727469636970616e740060448201526064016105cd565b8051604082015160608381015160c085015160e086015192948594909390929190612121856118c5565b6121378860200151858051906020012087611943565b61214084612642565b965061214b82612642565b95506000878481518110612161576121616139be565b6020908102919091010151519050600260ff16888581518110612186576121866139be565b6020026020010151604001518a60800151815181106121a7576121a76139be565b60200260200101516040015160ff16146122035760405162461bcd60e51b815260206004820152601a60248201527f6e6f7420612067756172616e74656520616c6c6f636174696f6e00000000000060448201526064016105cd565b6000888581518110612217576122176139be565b6020026020010151604001518a6080015181518110612238576122386139be565b6020026020010151600001519050816001600160a01b0316888481518110612262576122626139be565b6020026020010151600001516001600160a01b0316146122c45760405162461bcd60e51b815260206004820152601d60248201527f746172676574417373657420213d2067756172616e746565417373657400000060448201526064016105cd565b6122cd816118c5565b6122e38a60a00151858051906020012083611943565b50505050505050915091565b825160608401518351839085908390811061230c5761230c6139be565b6020026020010151604001819052506123358286602001518660405160200161180f9190613f69565b845160608601516040519081527f4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b90602001610909565b600061237783611d01565b905061238281611d47565b8151606001516123ca5760405162461bcd60e51b815260206004820152601360248201527214dd185d19481b5d5cdd08189948199a5b985b606a1b60448201526064016105cd565b60006123d68484611be7565b90508360000151516123eb82602001516128cd565b60ff16146124285760405162461bcd60e51b815260206004820152600a60248201526921756e616e696d6f757360b01b60448201526064016105cd565b61246d6040518060800160405280600065ffffffffffff1681526020014265ffffffffffff1681526020016000801b81526020016115d186600001516000015161192a565b60008381526020818152604091829020929092555165ffffffffffff4216815283917f4f465027a3d06ea73dd12be0f5c5fc0a34e21f19d6eaed4834a7a944edabc901910160405180910390a25092915050565b6000612516826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166128f89092919063ffffffff16565b905080516000148061253757508080602001905181019061253791906141f0565b6116965760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016105cd565b60005b81516125a6826001613a13565b10156118c157816125b8826001613a13565b815181106125c8576125c86139be565b60200260200101518282815181106125e2576125e26139be565b6020026020010151106126305760405162461bcd60e51b8152602060048201526016602482015275125b991a58d95cc81b5d5cdd081899481cdbdc9d195960521b60448201526064016105cd565b8061263a81613a26565b915050612599565b6060818060200190518101906116b9919061430e565b805160005b82604001515181101561169657600083604001518281518110612682576126826139be565b60200260200101516000015190506000846040015183815181106126a8576126a86139be565b60200260200101516020015190506126c18260a01c1590565b156126d6576126d1848383612907565b61270f565b6001600160a01b038416600090815260016020908152604080832085845290915281208054839290612709908490613a13565b90915550505b5050808061271c90613a26565b91505061265d565b6060816040516020016127379190613f69565b6040516020818303038152906040529050919050565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018390526000908190605c016040516020818303038152906040528051906020012090506000600182856000015186602001518760400151604051600081526020016040526040516127e7949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015612809573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166128605760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b60448201526064016105cd565b949350505050565b6000805b82518110156128c357828181518110612887576128876139be565b60200260200101516001600160a01b0316846001600160a01b0316036128b15760019150506116b9565b806128bb81613a26565b91505061286c565b5060009392505050565b6000805b82156116b9576128e2600184613a00565b90921691806128f08161446c565b9150506128d1565b60606128608484600085612a24565b6001600160a01b0383166129b1576000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114612962576040519150601f19603f3d011682016040523d82523d6000602084013e612967565b606091505b50509050806116715760405162461bcd60e51b8152602060048201526016602482015275086deead8c840dcdee840e8e4c2dce6cccae4408aa8960531b60448201526064016105cd565b60405163a9059cbb60e01b81526001600160a01b0383811660048301526024820183905284169063a9059cbb906044016020604051808303816000875af1158015612a00573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061167191906141f0565b606082471015612a855760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016105cd565b600080866001600160a01b03168587604051612aa1919061448b565b60006040518083038185875af1925050503d8060008114612ade576040519150601f19603f3d011682016040523d82523d6000602084013e612ae3565b606091505b5091509150612af487838387612aff565b979650505050505050565b60608315612b6e578251600003612b67576001600160a01b0385163b612b675760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016105cd565b5081612860565b6128608383815115612b835781518083602001fd5b8060405162461bcd60e51b81526004016105cd9190613e3d565b604051806060016040528060006001600160a01b03168152602001612bc0612bcd565b8152602001606081525090565b60408051808201909152806000612bc0565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715612c1757612c17612bdf565b60405290565b604051606081016001600160401b0381118282101715612c1757612c17612bdf565b604080519081016001600160401b0381118282101715612c1757612c17612bdf565b60405161010081016001600160401b0381118282101715612c1757612c17612bdf565b604051601f8201601f191681016001600160401b0381118282101715612cac57612cac612bdf565b604052919050565b60006001600160401b03821115612ccd57612ccd612bdf565b5060051b60200190565b60ff8116811461192757600080fd5b60006001600160401b03821115612cff57612cff612bdf565b50601f01601f191660200190565b600082601f830112612d1e57600080fd5b8135612d31612d2c82612ce6565b612c84565b818152846020838601011115612d4657600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112612d7457600080fd5b81356020612d84612d2c83612cb4565b82815260059290921b84018101918181019086841115612da357600080fd5b8286015b84811015612e445780356001600160401b0380821115612dc75760008081fd5b908801906080828b03601f1901811315612de15760008081fd5b612de9612bf5565b87840135815260408085013589830152606080860135612e0881612cd7565b83830152928501359284841115612e2157600091508182fd5b612e2f8e8b86890101612d0d565b90830152508652505050918301918301612da7565b509695505050505050565b600082601f830112612e6057600080fd5b81356020612e70612d2c83612cb4565b82815260059290921b84018101918181019086841115612e8f57600080fd5b8286015b84811015612e445780358352918301918301612e93565b600080600060608486031215612ebf57600080fd5b8335925060208401356001600160401b0380821115612edd57600080fd5b612ee987838801612d63565b93506040860135915080821115612eff57600080fd5b50612f0c86828701612e4f565b9150509250925092565b60005b83811015612f31578181015183820152602001612f19565b50506000910152565b60008151808452612f52816020860160208601612f16565b601f01601f19169290920160200192915050565b600082825180855260208086019550808260051b84010181860160005b84811015612fde57858303601f19018952815180518452848101518585015260408082015160ff1690850152606090810151608091850182905290612fca81860183612f3a565b9a86019a9450505090830190600101612f83565b5090979650505050505050565b608081526000612ffe6080830187612f66565b851515602084015282810360408401526130188186612f66565b91505082606083015295945050505050565b6001600160a01b038116811461192757600080fd5b803561304a8161302a565b919050565b6000806040838503121561306257600080fd5b823561306d8161302a565b946020939093013593505050565b6000806000806080858703121561309157600080fd5b843561309c8161302a565b966020860135965060408601359560600135945092505050565b600080600080600060a086880312156130ce57600080fd5b853594506020860135935060408601356001600160401b03808211156130f357600080fd5b6130ff89838a01612d0d565b945060608801359350608088013591508082111561311c57600080fd5b5061312988828901612e4f565b9150509295509295909350565b6004811061192757600080fd5b600082601f83011261315457600080fd5b81356020613164612d2c83612cb4565b82815260059290921b8401810191818101908684111561318357600080fd5b8286015b84811015612e445780356001600160401b03808211156131a657600080fd5b90880190601f196060838c03820112156131bf57600080fd5b6131c7612c1d565b878401356131d48161302a565b81526040840135838111156131e857600080fd5b84016040818e03840112156131fc57600080fd5b613204612c3f565b92508881013561321381613136565b835260408101358481111561322757600080fd5b6132358e8b83850101612d0d565b8a85015250508188820152606084013591508282111561325457600080fd5b6132628c8984870101612d63565b60408201528652505050918301918301613187565b60008060006060848603121561328c57600080fd5b8335925060208401356001600160401b038111156132a957600080fd5b6132b586828701613143565b925050604084013590509250925092565b6000602082840312156132d857600080fd5b5035919050565b6000806000606084860312156132f457600080fd5b83356001600160401b038082111561330b57600080fd5b61331787838801612d63565b9450602086013591508082111561332d57600080fd5b506132b586828701612d63565b6020815260006116b66020830184612f66565b6000604082840312156116b457600080fd5b6000806000806060858703121561337557600080fd5b84356001600160401b038082111561338c57600080fd5b90860190608082890312156133a057600080fd5b909450602086013590808211156133b657600080fd5b818701915087601f8301126133ca57600080fd5b8135818111156133d957600080fd5b8860208260051b85010111156133ee57600080fd5b60208301955080945050604087013591508082111561340c57600080fd5b506134198782880161334d565b91505092959194509250565b82151581526040602082015260006128606040830184612f3a565b80356001600160401b038116811461304a57600080fd5b803565ffffffffffff8116811461304a57600080fd5b60006080828403121561347f57600080fd5b613487612bf5565b905081356001600160401b0381111561349f57600080fd5b8201601f810184136134b057600080fd5b803560206134c0612d2c83612cb4565b82815260059290921b830181019181810190878411156134df57600080fd5b938201935b838510156135065784356134f78161302a565b825293820193908201906134e4565b855250613514858201613440565b818501525050506135276040830161303f565b604082015261353860608301613457565b606082015292915050565b801515811461192757600080fd5b60006060828403121561356357600080fd5b61356b612c1d565b9050813561357881612cd7565b80825250602082013560208201526040820135604082015292915050565b600082601f8301126135a757600080fd5b813560206135b7612d2c83612cb4565b828152606092830285018201928282019190878511156135d657600080fd5b8387015b85811015612fde576135ec8982613551565b84529284019281016135da565b60006040828403121561360b57600080fd5b613613612c3f565b905081356001600160401b038082111561362c57600080fd5b908301906080828603121561364057600080fd5b613648612bf5565b82358281111561365757600080fd5b61366387828601613143565b82525060208301358281111561367857600080fd5b61368487828601612d0d565b60208301525061369660408401613457565b6040820152606083013592506136ab83613543565b8260608201528084525060208401359150808211156136c957600080fd5b506136d684828501613596565b60208301525092915050565b60006136f0612d2c84612cb4565b8381529050602080820190600585901b84018681111561370f57600080fd5b845b8181101561374a5780356001600160401b038111156137305760008081fd5b61373c898289016135f9565b855250928201928201613711565b505050509392505050565b600082601f83011261376657600080fd5b6116b6838335602085016136e2565b60008060006060848603121561378a57600080fd5b83356001600160401b03808211156137a157600080fd5b6137ad8783880161346d565b945060208601359150808211156137c357600080fd5b6137cf87838801613755565b935060408601359150808211156137e557600080fd5b50612f0c868287016135f9565b60008060008060c0858703121561380857600080fd5b84356001600160401b038082111561381f57600080fd5b61382b8883890161346d565b9550602087013591508082111561384157600080fd5b61384d88838901613755565b9450604087013591508082111561386357600080fd5b50613870878288016135f9565b9250506138808660608701613551565b905092959194509250565b60006020828403121561389d57600080fd5b81356001600160401b03808211156138b457600080fd5b9083019061010082860312156138c957600080fd5b6138d1612c61565b82358152602083013560208201526040830135828111156138f157600080fd5b6138fd87828601612d0d565b604083015250606083013560608201526080830135608082015260a083013560a082015260c08301358281111561393357600080fd5b61393f87828601612d0d565b60c08301525060e083013560e082015280935050505092915050565b6000806040838503121561396e57600080fd5b82356001600160401b038082111561398557600080fd5b6139918683870161346d565b935060208501359150808211156139a757600080fd5b506139b4858286016135f9565b9150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b818103818111156116b9576116b96139ea565b808201808211156116b9576116b96139ea565b600060018201613a3857613a386139ea565b5060010190565b600060208284031215613a5157600080fd5b813561123f8161302a565b60006116b9368361346d565b60006116b63684846136e2565b60006116b936836135f9565b8183526000602080850194508260005b85811015613abf578135613aa48161302a565b6001600160a01b031687529582019590820190600101613a91565b509495945050505050565b600081518084526020808501808196508360051b8101915082860160005b85811015613b3f578284038952815180518552858101518686015260408082015160ff1690860152606090810151608091860182905290613b2b81870183612f3a565b9a87019a9550505090840190600101613ae8565b5091979650505050505050565b600081518084526020808501808196508360051b810191508286016000805b86811015613bfe578385038a528251606060018060a01b03825116875287820151818989015280516004808210613baf57634e487b7160e01b875260218152602487fd5b5091880191909152870151604060808801819052613bd060a0890183612f3a565b91508083015192508782038189015250613bea8183613aca565b9b88019b9650505091850191600101613b6b565b509298975050505050505050565b6000815160808452613c216080850182613b4c565b905060208301518482036020860152613c3a8282612f3a565b91505065ffffffffffff60408401511660408501526060830151151560608501528091505092915050565b6000815160408452613c7a6040850182613c0c565b602093840151949093019390935250919050565b600081518084526020808501808196508360051b8101915082860160005b85811015613b3f578284038952613cc4848351613c65565b98850198935090840190600101613cac565b6060815260008435601e19863603018112613cf057600080fd5b85016020810190356001600160401b03811115613d0c57600080fd5b8060051b3603821315613d1e57600080fd5b60806060850152613d3360e085018284613a81565b915050613d4260208701613440565b6001600160401b03166080840152613d5c6040870161303f565b6001600160a01b031660a0840152613d7660608701613457565b65ffffffffffff1660c08401528281036020840152613d958186613c8e565b90508281036040840152613da98185613c65565b9695505050505050565b6000613dc1612d2c84612ce6565b9050828152838383011115613dd557600080fd5b61123f836020830184612f16565b60008060408385031215613df657600080fd5b8251613e0181613543565b60208401519092506001600160401b03811115613e1d57600080fd5b8301601f81018513613e2e57600080fd5b6139b485825160208401613db3565b6020815260006116b66020830184612f3a565b65ffffffffffff818116838216019080821115613e6f57613e6f6139ea565b5092915050565b600060408251818552613e8b82860182613c0c565b60208581015187830388830152805180845290820193509091600091908301905b80831015613ee2578451805160ff1683528481015185840152860151868301529383019360019290920191606090910190613eac565b50979650505050505050565b60006060820165ffffffffffff86168352602060608185015281865180845260808601915060808160051b870101935082880160005b82811015613f5257607f19888703018452613f40868351613e76565b95509284019290840190600101613f24565b50505050508281036040840152613da98185613e76565b6020815260006116b66020830184613b4c565b61ffff828116828216039080821115613e6f57613e6f6139ea565b600060408284031215613fa957600080fd5b613fb1612c3f565b82518152602083015160208201528091505092915050565b600181815b80851115614004578160001904821115613fea57613fea6139ea565b80851615613ff757918102915b93841c9390800290613fce565b509250929050565b60008261401b575060016116b9565b81614028575060006116b9565b816001811461403e576002811461404857614064565b60019150506116b9565b60ff841115614059576140596139ea565b50506001821b6116b9565b5060208310610133831016604e8410600b8410161715614087575081810a6116b9565b6140918383613fc9565b80600019048211156140a5576140a56139ea565b029392505050565b60006116b6838361400c565b600081518084526020808501945080840160005b83811015613abf5781516001600160a01b0316875295820195908201906001016140cd565b60808152600061410560808301876140b9565b6001600160401b03959095166020830152506001600160a01b0392909216604083015265ffffffffffff16606090910152919050565b60608152600084516080606084015261415760e08401826140b9565b6020878101516001600160401b0316608086015260408801516001600160a01b031660a0860152606088015165ffffffffffff1660c0860152848203908501529050613d958186613c8e565b85815260a0602082015260006141bc60a0830187612f3a565b82810360408401526141ce8187613b4c565b65ffffffffffff95909516606084015250509015156080909101529392505050565b60006020828403121561420257600080fd5b815161123f81613543565b600082601f83011261421e57600080fd5b6116b683835160208501613db3565b600082601f83011261423e57600080fd5b8151602061424e612d2c83612cb4565b82815260059290921b8401810191818101908684111561426d57600080fd5b8286015b84811015612e445780516001600160401b03808211156142915760008081fd5b908801906080828b03601f19018113156142ab5760008081fd5b6142b3612bf5565b878401518152604080850151898301526060808601516142d281612cd7565b838301529285015192848411156142eb57600091508182fd5b6142f98e8b8689010161420d565b90830152508652505050918301918301614271565b60006020828403121561432057600080fd5b81516001600160401b038082111561433757600080fd5b818401915084601f83011261434b57600080fd5b8151614359612d2c82612cb4565b8082825260208201915060208360051b86010192508783111561437b57600080fd5b602085015b83811015613ee25780518581111561439757600080fd5b8601601f196060828c03820112156143ae57600080fd5b6143b6612c1d565b60208301516143c48161302a565b81526040830151888111156143d857600080fd5b83016040818e03840112156143ec57600080fd5b6143f4612c3f565b9250602081015161440481613136565b835260408101518981111561441857600080fd5b6144278e60208385010161420d565b60208501525050816020820152606083015191508782111561444857600080fd5b6144578c60208486010161422d565b60408201528552505060209283019201614380565b600060ff821660ff8103614482576144826139ea565b60010192915050565b6000825161449d818460208701612f16565b919091019291505056fea26469706673582212200c3f23d06e4d077392c9a51da9f112dfa4ec40ea2dd45d3fd1ee4591348c631a64736f6c63430008140033",
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
