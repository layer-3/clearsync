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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialHoldings\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalHoldings\",\"type\":\"uint256\"}],\"name\":\"AllocationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newTurnNumRecord\",\"type\":\"uint48\"}],\"name\":\"ChallengeCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"finalizesAt\",\"type\":\"uint48\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"ChallengeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"finalizesAt\",\"type\":\"uint48\"}],\"name\":\"Concluded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationHoldings\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"}],\"name\":\"Reclaimed\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature\",\"name\":\"challengerSig\",\"type\":\"tuple\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"checkpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"sourceAllocations\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"targetAllocations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"indexOfTargetInSource\",\"type\":\"uint256\"}],\"name\":\"compute_reclaim_effects\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialHoldings\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"compute_transfer_effects_and_interactions\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"newAllocations\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"allocatesOnlyZeros\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"exitAllocations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalPayouts\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"conclude\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"concludeAndTransferAllAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeld\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"holdings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sourceOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceAssetIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexOfTargetInSource\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"targetStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"targetOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"targetAssetIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIMultiAssetHolder.ReclaimArgs\",\"name\":\"reclaimArgs\",\"type\":\"tuple\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"stateIsSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"statusOf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"fromChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"outcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"}],\"name\":\"transferAllAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"}],\"name\":\"unpackStatus\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"turnNumRecord\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"finalizesAt\",\"type\":\"uint48\"},{\"internalType\":\"uint160\",\"name\":\"fingerprint\",\"type\":\"uint160\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052346200001d575b604051614a0c6200002d8239614a0c90f35b62000026600080fd5b6200000b56fe60806040526004361015610018575b610016600080fd5b005b60003560e01c806311e9f17814610130578063166e56cd146101275780632fb1d2701461011e5780633033730e1461011557806331afa0b41461010c578063552cfa5014610103578063566d54c6146100fa5780635685b7dc146100f15780636d2a9c92146100e85780638286a060146100df578063c7df14e2146100d6578063d3c4e738146100cd578063ec346235146100c45763ee049b500361000e576100bf6114b0565b61000e565b506100bf611488565b506100bf6113ff565b506100bf6112c0565b506100bf611286565b506100bf6111c5565b506100bf610d01565b506100bf610b9c565b506100bf610aed565b506100bf610a7f565b506100bf610848565b506100bf6107a0565b506100bf61071a565b506100bf610603565b600080fd5b805b0361013957565b905035906101548261013e565b565b601f01601f191690565b50634e487b7160e01b600052604160045260246000fd5b90601f01601f191681019081106001600160401b0382111761019857604052565b6101a0610160565b604052565b906101546101b260405190565b9283610177565b602080916001600160401b0381116101d057020190565b6101d8610160565b020190565b60ff8116610140565b90503590610154826101dd565b6102116020916001600160401b03811161021557601f01601f191690565b0190565b610156610160565b90826000939282370152565b9291906101549161024161023c836101f3565b6101a5565b948286526020860191838201111561021d5761025b600080fd5b61021d565b9061027c9181601f8201121561027f575b602081359101610229565b90565b610287600080fd5b610271565b91906102f79060808482031261030b575b6102a760806101a5565b9360006102b48383610147565b9086015260206102c683828401610147565b9086015260406102d8838284016101e6565b908601526060810135906001600160401b0382116102fe575b01610260565b6060830152565b610306600080fd5b6102f1565b610313600080fd5b61029d565b9092919261032861023c826101b9565b9381855260208086019202830192818411610387575b80925b84841061034f575050505050565b6020809161036f8587356001600160401b03811161037a575b860161028c565b815201930192610341565b610382600080fd5b610368565b61038f600080fd5b61033e565b9061027c9181601f820112156103b0575b602081359101610318565b6103b8600080fd5b6103a5565b909291926103cd61023c826101b9565b938185526020808601920283019281841161040b575b915b8383106103f25750505050565b602080916104008486610147565b8152019201916103e5565b610413600080fd5b6103e3565b9061027c9181601f82011215610434575b6020813591016103bd565b61043c600080fd5b610429565b90916060828403126104b1575b61027c61045b8484610147565b9361047b8160208601356001600160401b0381116104a4575b8601610394565b936040810135906001600160401b038211610497575b01610418565b61049f600080fd5b610491565b6104ac600080fd5b610474565b6104b9600080fd5b61044e565b9052565b60005b8381106104d55750506000910152565b81810151838201526020016104c5565b610506610156602093610211936104fa815190565b80835293849260200190565b958691016104c2565b8051825261027c91608081019160609061052e60208201516020850152565b60408181015160ff169084015201519060608184039101526104e5565b9061027c9161050f565b9061056b610561835190565b8083529160200190565b908161057d6020830284019460200190565b926000915b83831061059157505050505090565b909192939460206105b46105ad8385600195038752895161054b565b9760200190565b9301930191939290610582565b9493916105fe90610154946105f16105e760609560808b01908b820360008d0152610555565b92151560208a0152565b8782036040890152610555565b940152565b503461063b575b61063761062161061b366004610441565b91612eae565b9061062e94929460405190565b948594856105c1565b0390f35b610643600080fd5b61060a565b6001600160a01b031690565b6001600160a01b038116610140565b9050359061015482610654565b919061027c90604084820312610694575b61068b8185610663565b93602001610147565b61069c600080fd5b610681565b61027c90610648906001600160a01b031682565b61027c906106a1565b61027c906106b5565b906106d1906106be565b600052602052604060002090565b906106d1565b61027c916008021c81565b9061027c91546106e5565b61071561027c926107106001936000946106c7565b6106df565b6106f0565b5034610749575b610637610738610732366004610670565b906106fb565b6040515b9182918290815260200190565b610751600080fd5b610721565b608081830312610793575b61076b8282610663565b9261027c61077c8460208501610147565b9361078a8160408601610147565b93606001610147565b61079b600080fd5b610761565b506107b86107af366004610756565b92919091612614565b604051005b919060a08382031261083b575b6107d48184610147565b926107e28260208301610147565b9261027c6108058460408501356001600160401b03811161082e575b8501610260565b936108138160608601610147565b936080810135906001600160401b0382116104975701610418565b610836600080fd5b6107fe565b610843600080fd5b6107ca565b5034610866575b6107b861085d3660046107bd565b93929092612a3d565b61086e600080fd5b61084f565b6004111561013957565b9050359061015482610873565b91906108d0906040848203126108d7575b6108a560406101a5565b9360006108b2838361087d565b908601526020810135906001600160401b0382116102fe5701610260565b6020830152565b6108df600080fd5b61089b565b919061094f90606084820312610970575b6108ff60606101a5565b93600061090c8383610663565b9086015261092f8260208301356001600160401b038111610963575b830161088a565b60208601526040810135906001600160401b038211610956575b01610394565b6040830152565b61095e600080fd5b610949565b61096b600080fd5b610928565b610978600080fd5b6108f5565b9092919261098d61023c826101b9565b93818552602080860192028301928184116109ec575b80925b8484106109b4575050505050565b602080916109d48587356001600160401b0381116109df575b86016108e4565b8152019301926109a6565b6109e7600080fd5b6109cd565b6109f4600080fd5b6109a3565b9061027c9181601f82011215610a15575b60208135910161097d565b610a1d600080fd5b610a0a565b9091606082840312610a72575b61027c610a3c8484610147565b93610a5c8160208601356001600160401b038111610a65575b86016109f9565b93604001610147565b610a6d600080fd5b610a55565b610a7a600080fd5b610a2f565b5034610a9a575b6107b8610a94366004610a22565b91613f4b565b610aa2600080fd5b610a86565b9061027c916020818303121561014757610abf600080fd5b610147565b65ffffffffffff9182168152911660208201526001600160a01b03909116604082015260600190565b5034610b19575b610637610b0a610b05366004610aa7565b6114d8565b60405191939193849384610ac4565b610b21600080fd5b610af4565b9091606082840312610b7e575b61027c610b528484356001600160401b038111610b71575b8501610394565b93610a5c8160208601356001600160401b0381116104a4578601610394565b610b79600080fd5b610b4b565b610b86600080fd5b610b33565b602080825261027c92910190610555565b5034610bc6575b610637610bba610bb4366004610b26565b9161367d565b60405191829182610b8b565b610bce600080fd5b610ba3565b90816080910312610be15790565b61027c600080fd5b909182601f83011215610c2f575b60208235926001600160401b038411610c22575b019260208302840111610c1a57565b610154600080fd5b610c2a600080fd5b610c0b565b610c37600080fd5b610bf7565b90816040910312610be15790565b90606082820312610cda575b610c728183356001600160401b038111610ccd575b8401610bd3565b9261027c610c958360208601356001600160401b038111610cc0575b8601610be9565b9390946040810135906001600160401b038211610cb3575b01610c3c565b610cbb600080fd5b610cad565b610cc8600080fd5b610c8e565b610cd5600080fd5b610c6b565b610ce2600080fd5b610c56565b901515815260406020820181905261027c929101906104e5565b5034610d35575b610d1f610d16366004610c4a565b92919091614393565b90610637610d2c60405190565b92839283610ce7565b610d3d600080fd5b610d08565b90929192610d5261023c826101b9565b9381855260208086019202830192818411610d90575b915b838310610d775750505050565b60208091610d858486610663565b815201920191610d6a565b610d98600080fd5b610d68565b9061027c9181601f82011215610db9575b602081359101610d42565b610dc1600080fd5b610dae565b6001600160401b038116610140565b9050359061015482610dc6565b65ffffffffffff8116610140565b9050359061015482610de2565b919091608081840312610e78575b610e65610e1860806101a5565b93610e358184356001600160401b038111610e6b575b8501610d9d565b85526020610e4582858301610dd5565b908601526040610e5782828601610663565b908601526060809301610df0565b90830152565b610e73600080fd5b610e2e565b610e80600080fd5b610e0b565b801515610140565b9050359061015482610e85565b919091608081840312610f20575b610e65610eb560806101a5565b93610ed28184356001600160401b038111610f13575b85016109f9565b8552610ef28160208501356001600160401b03811161082e578501610260565b60208601526040610f0582828601610df0565b908601526060809301610e8d565b610f1b600080fd5b610ecb565b610f28600080fd5b610ea8565b919091606081840312610f75575b610e65610f4860606101a5565b936000610f5582856101e6565b908601526020610f6782828601610147565b908601526040809301610147565b610f7d600080fd5b610f3b565b90929192610f9261023c826101b9565b938185526060602086019202830192818411610fd2575b915b838310610fb85750505050565b6020606091610fc78486610f2d565b815201920191610fab565b610fda600080fd5b610fa9565b9061027c9181601f82011215610ffb575b602081359101610f82565b611003600080fd5b610ff0565b91906108d090604084820312611077575b61102360406101a5565b936110408282356001600160401b03811161106a575b8301610e9a565b85526020810135906001600160401b03821161105d575b01610fdf565b611065600080fd5b611057565b611072600080fd5b611039565b61107f600080fd5b611019565b9092919261109461023c826101b9565b93818552602080860192028301928184116110f3575b80925b8484106110bb575050505050565b602080916110db8587356001600160401b0381116110e6575b8601611008565b8152019301926110ad565b6110ee600080fd5b6110d4565b6110fb600080fd5b6110aa565b9061027c9181601f8201121561111c575b602081359101611084565b611124600080fd5b611111565b90916060828403126111b8575b61027c6111558484356001600160401b0381116111ab575b8501610dfd565b936111758160208601356001600160401b03811161119e575b8601611100565b936040810135906001600160401b038211611191575b01611008565b611199600080fd5b61118b565b6111a6600080fd5b61116e565b6111b3600080fd5b61114e565b6111c0600080fd5b611136565b50346111e0575b6107b86111da366004611129565b91611a81565b6111e8600080fd5b6111cc565b60c081830312611279575b6112148282356001600160401b03811161126c575b8301610dfd565b9261027c6112378460208501356001600160401b03811161125f575b8501611100565b936112568160408601356001600160401b0381116110e6578601611008565b93606001610f2d565b611267600080fd5b611230565b611274600080fd5b61120d565b611281600080fd5b6111f8565b50346112a4575b6107b861129b3660046111ed565b929190916118e1565b6112ac600080fd5b61128d565b600061071561027c92826106df565b50346112dd575b6106376107386112d8366004610aa7565b6112b1565b6112e5600080fd5b6112c7565b919091610100818403126113b0575b610e656113076101006101a5565b9360006113148285610147565b90860152602061132682828601610147565b908601526113488160408501356001600160401b03811161082e578501610260565b6040860152606061135b82828601610147565b90860152608061136d82828601610147565b9086015260a061137f82828601610147565b908601526113a18160c08501356001600160401b03811161082e578501610260565b60c086015260e0809301610147565b6113b8600080fd5b6112f9565b9061027c916020818303126113f2575b8035906001600160401b0382116113e5575b016112ea565b6113ed600080fd5b6113df565b6113fa600080fd5b6113cd565b5034611419575b6107b86114143660046113bd565b61334c565b611421600080fd5b611406565b919061027c9060408482031261147b575b6114538185356001600160401b03811161146e575b8601610dfd565b936020810135906001600160401b0382116111915701611008565b611476600080fd5b61144c565b611483600080fd5b611437565b50346114a3575b6107b861149d366004611426565b90613e50565b6114ab600080fd5b61148f565b50346114cb575b6107b86114c5366004611426565b90611ac4565b6114d3600080fd5b6114b7565b6114e59060005b506145f6565b909192565b50634e487b7160e01b600052602160045260246000fd5b6003111561150b57565b6101546114ea565b9061015482611501565b602080825261027c929101906104e5565b156115365750565b6115589061154360405190565b62461bcd60e51b81529182916004830161151d565b0390fd5b61156961027c61027c9290565b65ffffffffffff1690565b50634e487b7160e01b600052601160045260246000fd5b6115a49065ffffffffffff165b9165ffffffffffff1690565b019065ffffffffffff82116115b557565b610154611574565b6004111561150b57565b90610154826115bd565b61027c906115c7565b6104be906115d1565b61027c9160206040820192611600600082015160008501906115da565b01519060208184039101526104e5565b9061161c610561835190565b908161162e6020830284019460200190565b926000915b83831061164257505050505090565b9091929394602061165e6105ad8385600195038752895161054b565b9301930191939290611633565b80516001600160a01b0316825261027c91604061169760608301602085015184820360208601526115e3565b920151906040818403910152611610565b9061027c9161166b565b906116be610561835190565b90816116d06020830284019460200190565b926000915b8383106116e457505050505090565b909192939460206117006105ad838560019503875289516116a8565b93019301919392906116d5565b9061027c9060608061174361173160808501600088015186820360008801526116b2565b602087015185820360208701526104e5565b60408087015165ffffffffffff16908501529401511515910152565b805160ff1682526101549190604090819061177f60208201516020860152565b0151910152565b906102118160609361175f565b906117b36117ac6117a2845190565b8084529260200190565b9260200190565b9060005b8181106117c45750505090565b9091926117de6117d76001928651611786565b9460200190565b9291016117b7565b8051604080845261027c9391602091611802919084019061170d565b920151906020818403910152611793565b9061027c916117e6565b90611829610561835190565b908161183b6020830284019460200190565b926000915b83831061184f57505050505090565b9091929394602061186b6105ad83856001950387528951611813565b9301930191939290611840565b65ffffffffffff909116815261027c9290916118a0906060840190848203602086015261181d565b9160408184039101526117e6565b61027c60806101a5565b90600019905b9181191691161790565b906118d661027c6118dd9290565b82546118b8565b9055565b6118ea816148ad565b83516040015190919065ffffffffffff1661190483614457565b9383600096879661191488611513565b9061191e90611513565b1460001498611a31611a06611a43996119fd611a3d996119f76101549f996102f798611a389b611a48576119528d8c612450565b611966611960828487612032565b9061152e565b8581019a6119836119788d518761496a565b9a888701518c611cd3565b606061198e4261155c565b9501926119aa6119a4855165ffffffffffff1690565b8761158b565b6119e86119d57f0aa12461ee6c137332989aa12cec79f4772ab2c1a8732a382aada7e9f3ec9d349490565b946119df60405190565b93849384611878565b0390a25165ffffffffffff1690565b9061158b565b955101516149c6565b93611a22611a126118ae565b65ffffffffffff9098168c890152565b65ffffffffffff166020870152565b6040850152565b6144f3565b926106df565b6118c8565b611a518b614457565b611a64611a5e6001611513565b91611513565b03611a7857611a738d8c6123de565b611952565b611a738b6124ab565b90611abf61196061015494611a95856148ad565b81516040015190949065ffffffffffff1695611ab0866124ab565b611aba87876123de565b612032565b612313565b90611ace91611b69565b50565b15611ad857565b60405162461bcd60e51b815260206004820152601360248201527214dd185d19481b5d5cdd08189948199a5b985b606a1b6044820152606490fd5b61027c61027c61027c9260ff1690565b15611b2a57565b60405162461bcd60e51b815260206004820152600a60248201526921756e616e696d6f757360b01b6044820152606490fd5b61027c61027c61027c9290565b9190611b74836148ad565b611bcd8194611b82836124ab565b835160600151611b93901515611ad1565b611bc7611bc161027c6000611bba611bb56020611bb08b8961220f565b015190565b6147c3565b9401515190565b91611b13565b14611b23565b611c376000611a4383611a3d611be24261155c565b96611a38611bfe8680611bf481611b5c565b94015101516149c6565b6102f7611c096118ae565b93611c24611c168a61155c565b65ffffffffffff16868b0152565b65ffffffffffff8c166020860152611a31565b611c82611c627f4f465027a3d06ea73dd12be0f5c5fc0a34e21f19d6eaed4834a7a944edabc9019290565b92611c6c60405190565b9182918265ffffffffffff909116815260200190565b0390a2565b15611c8e57565b60405162461bcd60e51b815260206004820152601f60248201527f4368616c6c656e676572206973206e6f742061207061727469636970616e74006044820152606490fd5b90611d3d61015493611d4293611d25611ceb60405190565b602081019283526040808201526009606082015268666f7263654d6f766560b81b608082015291829060a082015b90810382520382610177565b611d37611d30825190565b9160200190565b20614704565b611d9e565b611c87565b6001906000198114611d57570190565b610211611574565b50634e487b7160e01b600052603260045260246000fd5b9060208091611d83845190565b811015611d91575b02010190565b611d99611d5f565b611d8b565b600091611daa83611b5c565b611db561027c835190565b811015611dff57611dd9610648611dcc8385611d76565b516001600160a01b031690565b6001600160a01b03841614611df657611df190611d47565b611daa565b50505050600190565b50505090565b9050519061015482610e85565b92919061015491611e2561023c836101f3565b94828652602086019183820111156104c257611e3f600080fd5b6104c2565b9061027c9181601f82011215611e60575b602081519101611e12565b611e68600080fd5b611e55565b919061027c90604084820312611eb1575b611e888185611e05565b936020810151906001600160401b038211611ea4575b01611e44565b611eac600080fd5b611e9e565b611eb9600080fd5b611e7e565b90611ecd6117ac6117a2845190565b9060005b818110611ede5750505090565b909192611efd6117d760019286516001600160a01b0316815260200190565b929101611ed1565b9061027c90606080611f266080840160008701518582036000870152611ebe565b6020808701516001600160401b031690850152946040818101516001600160a01b031690850152015165ffffffffffff16910152565b9061027c90602080611f7d604084016000870151858203600087015261170d565b940151910152565b9061027c91611f5c565b90611f9b610561835190565b9081611fad6020830284019460200190565b926000915b838310611fc157505050505090565b90919293946020611fdd6105ad83856001950387528951611f85565b9301930191939290611fb2565b916120179061200961027c959360608601908682036000880152611f05565b908482036020860152611f8f565b916040818403910152611f5c565b506040513d6000823e3d90fd5b91929160009161208761205a61205561205560408601516001600160a01b031690565b6106be565b91612092612077612070639936d812938761218e565b988661220f565b6040519889968795869560e01b90565b855260048501611fea565b03915afa80156120cd575b60009283916120ab57509190565b90506120c991923d8091833e6120c18183610177565b810190611e6d565b9091565b6120d5612025565b61209d565b906120e761023c836101b9565b918252565b61027c60406101a5565b6120fe6118ae565b906060825260208080808501606081520160008152016000905250565b61027c6120f6565b61212b6120ec565b9061213461211b565b825260006020830152565b61027c612123565b60005b82811061215657505050565b60209061216161213f565b818401520161214a565b9061015461218161217b846120da565b936101b9565b601f190160208401612147565b61219e612199835190565b61216b565b916121a96000611b5c565b6121b461027c835190565b811015611dff57806121d36121cc6121ee9385611d76565b518561220f565b6121dd8287611d76565b526121e88186611d76565b50611d47565b6121a9565b60ff8111612202575b60020a90565b61220a611574565b6121fc565b9161221861213f565b5081519060006122306122296120ec565b9382850152565b61223981611b5c565b91612245836020860152565b82945b6020810161225861027c82515190565b871015612308576122829061227b88612274878601518c61496a565b9251611d76565b5190614704565b93805b84890161229461027c82515190565b8210156122fa57610648611dcc836122ac9351611d76565b6001600160a01b038716146122c9576122c490611d47565b612285565b6122f492979195506122dd6122ef916121f3565b60208801906122ea825190565b179052565b611d47565b94612248565b50509350946122f490611d47565b505094505050905090565b6123676000611a4383611a3d61232884611b5c565b611a386123336118ae565b65ffffffffffff8a16878201529161235c61234d8861155c565b65ffffffffffff166020850152565b6102f7816040850152565b611c82611c627f07da0a0674fb921e484018c8b81d80e292745e5d8ed134b580c8b9c631c5e9e09290565b1561239957565b60405162461bcd60e51b815260206004820152601c60248201527f7475726e4e756d5265636f7264206e6f7420696e637265617365642e000000006044820152606490fd5b906123fe6115986123f1610154946145f6565b505065ffffffffffff1690565b11612392565b1561240b57565b60405162461bcd60e51b815260206004820152601860248201527f7475726e4e756d5265636f7264206465637265617365642e00000000000000006044820152606490fd5b906124636115986123f1610154946145f6565b1015612404565b1561247157565b60405162461bcd60e51b815260206004820152601260248201527121b430b73732b6103334b730b634bd32b21760711b6044820152606490fd5b6124b761015491614457565b6124c4611a5e6002611513565b141561246a565b156124d257565b60405162461bcd60e51b815260206004820152601f60248201527f4465706f73697420746f2065787465726e616c2064657374696e6174696f6e006044820152606490fd5b61027c9081565b61027c9054612517565b1561252f57565b60405162461bcd60e51b81526020600482015260146024820152731a195b1908084f48195e1c1958dd195912195b1960621b6044820152606490fd5b61064861027c61027c9290565b61027c9061256b565b1561258857565b60405162461bcd60e51b815260206004820152601f60248201527f496e636f7272656374206d73672e76616c756520666f72206465706f736974006044820152606490fd5b91906125d8565b9290565b82018092116115b557565b906118d661027c6118dd92611b5c565b6001600160a01b0390911681526040810192916101549160200152565b0152565b92836126a1836107106126996126a69585979861263f61263a61263687613b59565b1590565b6124cb565b61266461265d6125d46126588861071060019c8d6106c7565b61251e565b8214612528565b6126716106486000612578565b6001600160a01b038a16036126de57612694823461268e565b9190565b14612581565b6125cd565b9586946106c7565b6125e3565b7f87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b9192611c826126d560405190565b928392836125f3565b612694826126eb8b6106be565b336126f5306106be565b91612745565b61271461270e61027c9263ffffffff1690565b60e01b90565b6001600160e01b03191690565b6001600160a01b039182168152911660208201526060810192916101549160400152565b9061278a9061277b610154956004956127616323b872dd6126fb565b9361276b60405190565b9788956020870190815201612721565b60208201810382520383610177565b612853565b906120e761023c836101f3565b6127a6602061278f565b7f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602082015290565b61027c61279c565b9061027c9160208183031215611e05576127ef600080fd5b611e05565b156127fb57565b60405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b9061286061286f926106be565b906128696127cf565b916128a7565b805161287e61268a6000611b5c565b116128865750565b6128a2816020612897610154945190565b8183010191016127d7565b6127f4565b61027c92916128b66000611b5c565b91612936565b156128c357565b60405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608490fd5b3d15612931576129263d61278f565b903d6000602084013e565b606090565b90600061027c94938192612948606090565b5061295f612955306106be565b83903110156128bc565b60208101905191855af1612971612917565b916129c3565b1561297e57565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b9192606091156129f857505081516129de61268a6000611b5c565b146129e7575090565b6129f361027c91612a05565b612977565b9093926101549250612a17565b3b612a1361268a6000611b5c565b1190565b9150612a21825190565b612a2e61268a6000611b5c565b11156115365750805190602001fd5b90612a806101549594612a74612a59612a859683868a89612a8b565b9282969291996040612a6b8a85611d76565b51015185612eae565b96929491509a8961323e565b611d76565b51613305565b94926000612ad9612ad28597612acd61027c97612ae896612ab761265899612ab1606090565b50613d2a565b612ac082613c83565b8351602085012090613c0b565b612e10565b9788611d76565b5101516001600160a01b031690565b936107108560016106c7565b9050519061015482610654565b9050519061015482610873565b91906108d090604084820312612b54575b612b2960406101a5565b936000612b368383612b01565b908601526020810151906001600160401b038211611ea45701611e44565b612b5c600080fd5b612b1f565b905051906101548261013e565b90505190610154826101dd565b91906102f790608084820312612be5575b612b9660806101a5565b936000612ba38383612b61565b908601526020612bb583828401612b61565b908601526040612bc783828401612b6e565b908601526060810151906001600160401b038211611ea45701611e44565b612bed600080fd5b612b8c565b90929192612c0261023c826101b9565b9381855260208086019202830192818411612c61575b80925b848410612c29575050505050565b60208091612c498587516001600160401b038111612c54575b8601612b7b565b815201930192612c1b565b612c5c600080fd5b612c42565b612c69600080fd5b612c18565b9061027c9181601f82011215612c8a575b602081519101612bf2565b612c92600080fd5b612c7f565b919061094f90606084820312612d1c575b612cb260606101a5565b936000612cbf8383612af4565b90860152612ce28260208301516001600160401b038111612d0f575b8301612b0e565b60208601526040810151906001600160401b038211612d02575b01612c6e565b612d0a600080fd5b612cfc565b612d17600080fd5b612cdb565b612d24600080fd5b612ca8565b90929192612d3961023c826101b9565b9381855260208086019202830192818411612d98575b80925b848410612d60575050505050565b60208091612d808587516001600160401b038111612d8b575b8601612c97565b815201930192612d52565b612d93600080fd5b612d79565b612da0600080fd5b612d4f565b9061027c9181601f82011215612dc1575b602081519101612d29565b612dc9600080fd5b612db6565b9061027c91602081830312612e03575b8051906001600160401b038211612df6575b01612da5565b612dfe600080fd5b612df0565b612e0b600080fd5b612dde565b61027c906020612e1e825190565b818301019101612dce565b612e316118ae565b906000825260208080808501600081520160005b8152016060905250565b61027c612e29565b60005b828110612e6657505050565b602090612e71612e4f565b8184015201612e5a565b90610154612e8b61217b846120da565b601f190160208401612e57565b61027c90611513565b919082039182116115b557565b91805193600094612ec161268a87611b5c565b11156131ac57612eda612ed2835190565b939293612e7b565b9283612ee587611b5c565b968794612ef8612ef3825190565b612e7b565b9384996001938491829b97819882945b612f19575b50505050505050505050565b9091929394959697989a612f2e61027c895190565b8610156131a657612f5c612f4c86612f46898c611d76565b51015190565b86612f578987611d76565b510152565b612f8d612f776040612f6e898c611d76565b51015160ff1690565b6040612f838987611d76565b51019060ff169052565b6060612f99878a611d76565b5101516060612fa88886611d76565b510152612fc382612fbe6020612f468a8d611d76565b613da5565b9084612fd061027c8a5190565b14801561316d575b1561314857612fec6040612f6e8f8c611d76565b613009613002612ffc6002612e98565b60ff1690565b9160ff1690565b14613103576130b6826130bc928f6130b0908f8b8f8e6102f78261304d8f6020612f57866130478f61304285612f46856130a09e611d76565b612ea1565b93611d76565b61308b606061307961306388612f468887611d76565b956130736040612f6e8388611d76565b94611d76565b510151936130856118ae565b96870152565b6130968a6020870152565b60ff166040850152565b6130aa8383611d76565b52611d76565b506125cd565b9c611d47565b915b846130d161027c6020612f468b89611d76565b036130f9575b6130ea916130e491612ea1565b95611d47565b93929190989796959487612f08565b949d508d946130d7565b60405162461bcd60e51b815260206004820152601b60248201527f63616e6e6f74207472616e7366657220612067756172616e74656500000000006044820152606490fd5b9b9161316861315c6020612f468a8d611d76565b6020612f578a88611d76565b6130be565b508c61317d61268a61027c8b5190565b108d8161318b575b50612fd8565b61319e915061319a908a611d76565b5190565b87148d613185565b9a612f0d565b612eda612ed2845190565b906131c3610561835190565b90816131d56020830284019460200190565b926000915b8383106131e957505050505090565b909192939460206132056105ad838560019503875289516116a8565b93019301919392906131da565b602080825261027c929101906131b7565b908152606081019392610154929091604091612610906108d0565b9690806132b7612658956132bc97869a9960406132858e61327f6107109a61327961326f8e61071060019d8e6106c7565b916130428361251e565b906125e3565b84611d76565b5101526132a461329460405190565b8092611d19602083019182613212565b6132af611d30825190565b209086613ca2565b6106c7565b92611c826132e87fc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a9490565b946132f260405190565b93849384613223565b61027c60606101a5565b610154916133479061094f602061332385516001600160a01b031690565b9401516133406133316132fb565b6001600160a01b039096168652565b6020850152565b61396d565b6101549061339b61335c82613439565b909190604061338e81613379613373606089015190565b87611d76565b5101519261338860e088015190565b90611d76565b5101516080850151610bb4565b916138db565b156133a857565b60405162461bcd60e51b815260206004820152601a60248201527f6e6f7420612067756172616e74656520616c6c6f636174696f6e0000000000006044820152606490fd5b156133f457565b60405162461bcd60e51b815260206004820152601d60248201527f746172676574417373657420213d2067756172616e74656541737365740000006044820152606490fd5b90613442825190565b90604083015160608401613454905190565b9260c08501519160e08601613467905190565b9161347181613c83565b60208701519061347f835190565b602084012061348d92613c0b565b61349690612e10565b94856134a184612e10565b80966134ad8184611d76565b51516001600160a01b0316926134c38282611d76565b51604001519160808601926134d6845190565b6134df91611d76565b516040015160ff166134f16002612e98565b60ff169060ff1614613502906133a1565b61350b91611d76565b5160400151905161351b91611d76565b51519361352791611d76565b5151613540916001600160a01b039081169116146133ed565b61354982613c83565b60a0015191613556815190565b906020012061015492613c0b565b1561356b57565b60405162461bcd60e51b815260206004820152601560248201527418dbdd5b19081b9bdd08199a5b99081d185c99d95d605a1b6044820152606490fd5b156135af57565b60405162461bcd60e51b815260206004820152601360248201527218dbdd5b19081b9bdd08199a5b99081b19599d606a1b6044820152606490fd5b156135f157565b60405162461bcd60e51b815260206004820152601460248201527318dbdd5b19081b9bdd08199a5b99081c9a59da1d60621b6044820152606490fd5b1561363457565b60405162461bcd60e51b815280611558600482016020808252818101527f746f74616c5265636c61696d6564213d67756172616e7465652e616d6f756e74604082015260600190565b91929180519061369b612ef360019361369585611b5c565b90612ea1565b926136a68683611d76565b51916136b56060840151613e2e565b966000948592818795889b6136c8600090565b996136d281611b5c565b80945b613714575b5050505050505050602061370e93611bb061027c97989961370961268a9661370461027c97613564565b6135a8565b6135ea565b1461362d565b61371f61027c875190565b8510156138d6578e8d8987146138c3578261379591856130a08a6102f78d61309661374e86612f468685611d76565b9361378e606061377c6137666020612f468689611d76565b936137766040612f6e838a611d76565b96611d76565b510151956137886118ae565b98890152565b6020870152565b508d8b158061389e575b613853575b50155b8061382d575b6137cc575b906137c26130e488959493611d47565b94909192936136d5565b9d50999061381c8693928f8e6104be60206137fa6138059461307383612f466137f48e611b5c565b8d611d76565b510191612694835190565b6138166020612f4661337388611b5c565b906125cd565b929d929a8e965091929091906137b2565b5061383c82612f468789611d76565b61384d61268a61027c602087015190565b146137ad565b829d919b50613879906104be60206137fa6138909661307383612f468d6133888d611b5c565b6138166020612f4661388a86611b5c565b88611d76565b9a6137a7879a90508d6137a4565b506138ad84612f46898b611d76565b6138bd61268a61027c8789015190565b1461379f565b5050939750908592916137c28499611d47565b6136da565b6139379161319a61393292600081019261391360206138f8865190565b936060810199604061390b6133738d5190565b510152015190565b9061392061329460405190565b61392b611d30825190565b2091613ca2565b915190565b90611c826139637f4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b9290565b9261073c60405190565b80516001600160a01b03169160009261398584611b5c565b6040840161399561027c82515190565b821015613a0b5781816139c26020612f46846139bb8c612f4660409a6139e59a51611d76565b9451611d76565b6139cb82613b59565b156139ec576139df6120556122ef93613ba2565b86613a58565b9050613985565b613279613a016122ef936107108960016106c7565b916126948361251e565b505050915050565b15613a1a57565b60405162461bcd60e51b8152602060048201526016602482015275086deead8c840dcdee840e8e4c2dce6cccae4408aa8960531b6044820152606490fd5b90600091613a6861064884612578565b6001600160a01b03821603613a9c5750819061015493613a8760405190565b90818003925af1613a96612917565b50613a13565b91613acd613aae6120556020956106be565b9163a9059cbb613ad8613ac060405190565b9788968795869460e01b90565b8452600484016125f3565b03925af18015613b11575b613aea5750565b611ace9060203d8111613b0a575b613b028183610177565b8101906127d7565b503d613af8565b613b19612025565b613ae3565b613b3461027c61027c926001600160601b031690565b6001600160601b031690565b61027c9060a01c613b1e565b613b3461027c61027c9290565b613b76613b7b91613b68600090565b506001600160a01b03191690565b613b40565b613b95613b886000613b4c565b916001600160601b031690565b1490565b61027c90611b5c565b612055613bc2613bbd61027c93613bb7600090565b50613b99565b61256b565b6106b5565b15613bce57565b60405162461bcd60e51b81526020600482015260156024820152741a5b98dbdc9c9958dd08199a5b99d95c9c1c9a5b9d605a1b6044820152606490fd5b90613c2b610648613c3892613c22610154966145f6565b969150506145a3565b916001600160a01b031690565b14613bc7565b15613c4557565b60405162461bcd60e51b815260206004820152601660248201527521b430b73732b6103737ba103334b730b634bd32b21760511b6044820152606490fd5b613c8f61015491614457565b613c9c611a5e6002611513565b14613c3e565b90613cdd611a31611a38610154956102f7611a4395613cc0886145f6565b50611a22613ccf9792976118ae565b65ffffffffffff9098168852565b9160006106df565b15613cec57565b60405162461bcd60e51b8152602060048201526016602482015275125b991a58d95cc81b5d5cdd081899481cdbdc9d195960521b6044820152606490fd5b613d346000611b5c565b600190613d49613d4383611b5c565b826125cd565b613d5761268a61027c865190565b1015613da0576122ef613d9b92613d9561268a61027c61319a613d8f613d89613d8361319a8a8d611d76565b96611b5c565b886125cd565b89611d76565b10613ce5565b613d34565b505050565b81811115613db1575090565b905090565b613dbe6120ec565b906000612134565b61027c613db6565b919091604081840312613e04575b610e65613de960406101a5565b936000613df68285612b61565b908601526020809301612b61565b613e0c600080fd5b613ddc565b9061027c9160408183031215613dce57613e29600080fd5b613dce565b61027c90613e3a613dc6565b506020613e45825190565b818301019101613e11565b90600080613e618361015495611b69565b9201510151610a946000611b5c565b613e786120ec565b906000825260606020830152565b61027c613e70565b613e966132fb565b9060008252602080808401612e45613e86565b61027c613e8e565b60005b828110613ec057505050565b602090613ecb613ea9565b8184015201613eb4565b90610154613ee561217b846120da565b601f190160208401613eb1565b369037565b90610154613f0761217b846120da565b601f190160208401613ef2565b9160001960089290920291821b911b6118be565b9190613f3761027c6118dd9390565b908354613f14565b61015491600091613f28565b929192613f5781613c83565b613f6a81613f64846149c6565b86613c0b565b6001908194613f7f613f7a855190565b613ed5565b93613f90613f8b825190565b613ef7565b96613f9c613f8b835190565b9460009581613faa88611b5c565b905b6140e9575b5081613fbc88611b5c565b905b614001575b505050610154969750600014613fed57505090613fe3613fe892826106df565b613f3f565b614404565b613fe89350613ffb906149c6565b91613ca2565b61400c61027c865190565b8110156140e457908188614021859488611d76565b5101516001600160a01b0316888d8361403a8187611d76565b5183614046868a6106c7565b90614050916106df565b9061405a8261251e565b9061406491612ea1565b61406d916125e3565b61407691611d76565b519161408290866106c7565b9061408c916106df565b6140959061251e565b907fc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a908a926140c360405190565b9182916140d1918784613223565b0390a26140dd90611d47565b9091613fbe565b613fc3565b826140f561027c875190565b8210156141f4575087858c61410a8483611d76565b5160408101518461411b8786611d76565b5101516001600160a01b031692868d614134868c6106c7565b9061413e916106df565b6141479061251e565b6141518284611d76565b5261415b91611d76565b519061416686611b5c565b61416f90613ef7565b9061417992612eae565b909591156141ea575b9361334060206141cf958a989560406141b08f9d8f98612a806141e49f9d6141ad8461094f9d611d76565b52565b5101520151916141be6132fb565b968701906001600160a01b03169052565b6141d9828d611d76565b526121e8818c611d76565b90613fac565b9599508995614182565b50613fb1565b3561027c81610654565b61027c903690610dfd565b61027c913691611084565b61027c903690611008565b9035601e193683900301811215614269575b01602081359101916001600160401b03821161425c575b6020820236038313610c1a57565b614264600080fd5b61424e565b614271600080fd5b614237565b5061027c906020810190610663565b818352602090920191906000825b8282106142a1575050505090565b909192936142d06142c96001926142b88886614276565b6001600160a01b0316815260200190565b9560200190565b93920190614293565b5061027c906020810190610dd5565b5061027c906020810190610df0565b9061027c90606061436761431d608084016143128780614225565b868303875290614285565b9461433e61432e60208301836142d9565b6001600160401b03166020860152565b61435e61434e6040830183614276565b6001600160a01b03166040860152565b828101906142e8565b65ffffffffffff16910152565b916120179061200961027c9593606086019086820360008801526142f7565b906143f96000939594956143a5600090565b506120926120776143e06143c161205561205560408a016141fa565b956143da639936d812956143d48a614204565b9261420f565b9061218e565b986143f36143ed88614204565b9161421a565b9061220f565b855260048501614374565b9061440f6000611b5c565b61441a61027c845190565b81101561443d57806122ef6144326144389386611d76565b5161396d565b61440f565b509050565b61027c61027c61027c9265ffffffffffff1690565b6144629060006114df565b50600091506144708261155c565b65ffffffffffff821603614482575090565b905061448e4291614442565b1161449857600290565b600190565b6144aa61027c61027c9290565b61ffff1690565b61ffff908116911690039061ffff82116115b557565b61027c906144dc61268a61027c9461ffff1690565b901b90565b61027c9081906001600160a01b031681565b61027c9061458a61458561450861010061449d565b61456761452361451e865165ffffffffffff1690565b614442565b61456161454461453d614536603061449d565b80966144b1565b80936144c7565b9361455b61451e60208a015165ffffffffffff1690565b926144b1565b906144c7565b179261457f6060614579604084015190565b92015190565b906145a3565b6144e1565b17611b5c565b9081526040810192916101549160200152565b611d196145cb61027c93613bbd936145b9600090565b50604051938492602084019283614590565b6145d6611d30825190565b20613b99565b61027c906145f161268a61027c9461ffff1690565b901c90565b61265861460d91614605600090565b5060006106df565b9061462261461c61010061449d565b92613b99565b9061027c61466261465161465c61464361463c603061449d565b80986144b1565b9661465661465189896145dc565b61155c565b976144b1565b856145dc565b9261256b565b906102116120e76020937f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c0190565b6126106101549461094f6060949897956146ba608086019a6000870152565b60ff166020850152565b156146cb57565b60405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b6044820152606490fd5b60209160009161472561471660405190565b8092611d198783019182614668565b614730611d30825190565b206147646147418484015160ff1690565b9261475160406145798884015190565b9061475b60405190565b9485948561469b565b838052039060015afa15614797575b60005161027c6147866106486000612578565b6001600160a01b03831614156146c4565b61479f612025565b614773565b612ffc61027c61027c9290565b60019060ff1660ff8114611d57570190565b6000906147cf826147a4565b905b6147da83611b5c565b81111561480a576148026147da916147fb6147f56001611b5c565b82612ea1565b16926147b1565b9190506147d1565b50905090565b9061481f6117ac6117a2845190565b9060005b8181106148305750505090565b90919261484f6117d760019286516001600160a01b0316815260200190565b929101614823565b61489f6101549461488f61487d6060959998969960808601908682036000880152614810565b6001600160401b039099166020850152565b6001600160a01b03166040830152565b019065ffffffffffff169052565b8051614903906148c760208401516001600160401b031690565b90611d196148f260606148e460408801516001600160a01b031690565b96015165ffffffffffff1690565b604051958694602086019485614857565b61490e611d30825190565b2090565b90614951610154959796946149436149629360809661493660a08801926000890152565b86820360208801526104e5565b9084820360408601526131b7565b65ffffffffffff9097166060830152565b019015159052565b61497f61490391614979600090565b506148ad565b602083015190611d196000850151946149ae60606149a6604084015165ffffffffffff1690565b920151151590565b906149b860405190565b968795602087019586614912565b6149039061027c6132946040519056fea264697066735822122062051ac53465318ca83b8d0ee639edf5e0d59416dfd972d00ca0b013878e809e64736f6c63430008110033",
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
