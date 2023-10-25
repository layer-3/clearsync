// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iadjudicator

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

// IYellowAdjudicatorMetaData contains all meta data concerning the IYellowAdjudicator contract.
var IYellowAdjudicatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialHoldings\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalHoldings\",\"type\":\"uint256\"}],\"name\":\"AllocationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newTurnNumRecord\",\"type\":\"uint48\"}],\"name\":\"ChallengeCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"finalizesAt\",\"type\":\"uint48\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"ChallengeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"finalizesAt\",\"type\":\"uint48\"}],\"name\":\"Concluded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationHoldings\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"}],\"name\":\"Reclaimed\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature\",\"name\":\"challengerSig\",\"type\":\"tuple\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"checkpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"conclude\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"concludeAndTransferAllAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedHeld\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sourceOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sourceAssetIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexOfTargetInSource\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"targetStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"targetOutcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"targetAssetIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIMultiAssetHolder.ReclaimArgs\",\"name\":\"reclaimArgs\",\"type\":\"tuple\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"participants\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"channelNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"appDefinition\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"challengeDuration\",\"type\":\"uint48\"}],\"internalType\":\"structINitroTypes.FixedPart\",\"name\":\"fixedPart\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart[]\",\"name\":\"proof\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"uint48\",\"name\":\"turnNum\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structINitroTypes.VariablePart\",\"name\":\"variablePart\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structINitroTypes.Signature[]\",\"name\":\"sigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structINitroTypes.SignedVariablePart\",\"name\":\"candidate\",\"type\":\"tuple\"}],\"name\":\"stateIsSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"fromChannelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"outcomeBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumExitFormat.AssetType\",\"name\":\"assetType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.AssetMetadata\",\"name\":\"assetMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"allocationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structExitFormat.Allocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"internalType\":\"structExitFormat.SingleAssetExit[]\",\"name\":\"outcome\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"}],\"name\":\"transferAllAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IYellowAdjudicatorABI is the input ABI used to generate the binding from.
// Deprecated: Use IYellowAdjudicatorMetaData.ABI instead.
var IYellowAdjudicatorABI = IYellowAdjudicatorMetaData.ABI

// IYellowAdjudicator is an auto generated Go binding around an Ethereum contract.
type IYellowAdjudicator struct {
	IYellowAdjudicatorCaller     // Read-only binding to the contract
	IYellowAdjudicatorTransactor // Write-only binding to the contract
	IYellowAdjudicatorFilterer   // Log filterer for contract events
}

// IYellowAdjudicatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IYellowAdjudicatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYellowAdjudicatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IYellowAdjudicatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYellowAdjudicatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IYellowAdjudicatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYellowAdjudicatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IYellowAdjudicatorSession struct {
	Contract     *IYellowAdjudicator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IYellowAdjudicatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IYellowAdjudicatorCallerSession struct {
	Contract *IYellowAdjudicatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IYellowAdjudicatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IYellowAdjudicatorTransactorSession struct {
	Contract     *IYellowAdjudicatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IYellowAdjudicatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IYellowAdjudicatorRaw struct {
	Contract *IYellowAdjudicator // Generic contract binding to access the raw methods on
}

// IYellowAdjudicatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IYellowAdjudicatorCallerRaw struct {
	Contract *IYellowAdjudicatorCaller // Generic read-only contract binding to access the raw methods on
}

// IYellowAdjudicatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IYellowAdjudicatorTransactorRaw struct {
	Contract *IYellowAdjudicatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIYellowAdjudicator creates a new instance of IYellowAdjudicator, bound to a specific deployed contract.
func NewIYellowAdjudicator(address common.Address, backend bind.ContractBackend) (*IYellowAdjudicator, error) {
	contract, err := bindIYellowAdjudicator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IYellowAdjudicator{IYellowAdjudicatorCaller: IYellowAdjudicatorCaller{contract: contract}, IYellowAdjudicatorTransactor: IYellowAdjudicatorTransactor{contract: contract}, IYellowAdjudicatorFilterer: IYellowAdjudicatorFilterer{contract: contract}}, nil
}

// NewIYellowAdjudicatorCaller creates a new read-only instance of IYellowAdjudicator, bound to a specific deployed contract.
func NewIYellowAdjudicatorCaller(address common.Address, caller bind.ContractCaller) (*IYellowAdjudicatorCaller, error) {
	contract, err := bindIYellowAdjudicator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IYellowAdjudicatorCaller{contract: contract}, nil
}

// NewIYellowAdjudicatorTransactor creates a new write-only instance of IYellowAdjudicator, bound to a specific deployed contract.
func NewIYellowAdjudicatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IYellowAdjudicatorTransactor, error) {
	contract, err := bindIYellowAdjudicator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IYellowAdjudicatorTransactor{contract: contract}, nil
}

// NewIYellowAdjudicatorFilterer creates a new log filterer instance of IYellowAdjudicator, bound to a specific deployed contract.
func NewIYellowAdjudicatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IYellowAdjudicatorFilterer, error) {
	contract, err := bindIYellowAdjudicator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IYellowAdjudicatorFilterer{contract: contract}, nil
}

// bindIYellowAdjudicator binds a generic wrapper to an already deployed contract.
func bindIYellowAdjudicator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IYellowAdjudicatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYellowAdjudicator *IYellowAdjudicatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYellowAdjudicator.Contract.IYellowAdjudicatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYellowAdjudicator *IYellowAdjudicatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.IYellowAdjudicatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYellowAdjudicator *IYellowAdjudicatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.IYellowAdjudicatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYellowAdjudicator *IYellowAdjudicatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYellowAdjudicator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYellowAdjudicator *IYellowAdjudicatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYellowAdjudicator *IYellowAdjudicatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.contract.Transact(opts, method, params...)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x5685b7dc.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) view returns(bool, string)
func (_IYellowAdjudicator *IYellowAdjudicatorCaller) StateIsSupported(opts *bind.CallOpts, fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (bool, string, error) {
	var out []interface{}
	err := _IYellowAdjudicator.contract.Call(opts, &out, "stateIsSupported", fixedPart, proof, candidate)

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
func (_IYellowAdjudicator *IYellowAdjudicatorSession) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (bool, string, error) {
	return _IYellowAdjudicator.Contract.StateIsSupported(&_IYellowAdjudicator.CallOpts, fixedPart, proof, candidate)
}

// StateIsSupported is a free data retrieval call binding the contract method 0x5685b7dc.
//
// Solidity: function stateIsSupported((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) view returns(bool, string)
func (_IYellowAdjudicator *IYellowAdjudicatorCallerSession) StateIsSupported(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (bool, string, error) {
	return _IYellowAdjudicator.Contract.StateIsSupported(&_IYellowAdjudicator.CallOpts, fixedPart, proof, candidate)
}

// Challenge is a paid mutator transaction binding the contract method 0x8286a060.
//
// Solidity: function challenge((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate, (uint8,bytes32,bytes32) challengerSig) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactor) Challenge(opts *bind.TransactOpts, fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart, challengerSig INitroTypesSignature) (*types.Transaction, error) {
	return _IYellowAdjudicator.contract.Transact(opts, "challenge", fixedPart, proof, candidate, challengerSig)
}

// Challenge is a paid mutator transaction binding the contract method 0x8286a060.
//
// Solidity: function challenge((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate, (uint8,bytes32,bytes32) challengerSig) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorSession) Challenge(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart, challengerSig INitroTypesSignature) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Challenge(&_IYellowAdjudicator.TransactOpts, fixedPart, proof, candidate, challengerSig)
}

// Challenge is a paid mutator transaction binding the contract method 0x8286a060.
//
// Solidity: function challenge((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate, (uint8,bytes32,bytes32) challengerSig) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactorSession) Challenge(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart, challengerSig INitroTypesSignature) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Challenge(&_IYellowAdjudicator.TransactOpts, fixedPart, proof, candidate, challengerSig)
}

// Checkpoint is a paid mutator transaction binding the contract method 0x6d2a9c92.
//
// Solidity: function checkpoint((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactor) Checkpoint(opts *bind.TransactOpts, fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _IYellowAdjudicator.contract.Transact(opts, "checkpoint", fixedPart, proof, candidate)
}

// Checkpoint is a paid mutator transaction binding the contract method 0x6d2a9c92.
//
// Solidity: function checkpoint((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorSession) Checkpoint(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Checkpoint(&_IYellowAdjudicator.TransactOpts, fixedPart, proof, candidate)
}

// Checkpoint is a paid mutator transaction binding the contract method 0x6d2a9c92.
//
// Solidity: function checkpoint((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactorSession) Checkpoint(fixedPart INitroTypesFixedPart, proof []INitroTypesSignedVariablePart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Checkpoint(&_IYellowAdjudicator.TransactOpts, fixedPart, proof, candidate)
}

// Conclude is a paid mutator transaction binding the contract method 0xee049b50.
//
// Solidity: function conclude((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactor) Conclude(opts *bind.TransactOpts, fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _IYellowAdjudicator.contract.Transact(opts, "conclude", fixedPart, candidate)
}

// Conclude is a paid mutator transaction binding the contract method 0xee049b50.
//
// Solidity: function conclude((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorSession) Conclude(fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Conclude(&_IYellowAdjudicator.TransactOpts, fixedPart, candidate)
}

// Conclude is a paid mutator transaction binding the contract method 0xee049b50.
//
// Solidity: function conclude((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactorSession) Conclude(fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Conclude(&_IYellowAdjudicator.TransactOpts, fixedPart, candidate)
}

// ConcludeAndTransferAllAssets is a paid mutator transaction binding the contract method 0xec346235.
//
// Solidity: function concludeAndTransferAllAssets((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactor) ConcludeAndTransferAllAssets(opts *bind.TransactOpts, fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _IYellowAdjudicator.contract.Transact(opts, "concludeAndTransferAllAssets", fixedPart, candidate)
}

// ConcludeAndTransferAllAssets is a paid mutator transaction binding the contract method 0xec346235.
//
// Solidity: function concludeAndTransferAllAssets((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorSession) ConcludeAndTransferAllAssets(fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.ConcludeAndTransferAllAssets(&_IYellowAdjudicator.TransactOpts, fixedPart, candidate)
}

// ConcludeAndTransferAllAssets is a paid mutator transaction binding the contract method 0xec346235.
//
// Solidity: function concludeAndTransferAllAssets((address[],uint64,address,uint48) fixedPart, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactorSession) ConcludeAndTransferAllAssets(fixedPart INitroTypesFixedPart, candidate INitroTypesSignedVariablePart) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.ConcludeAndTransferAllAssets(&_IYellowAdjudicator.TransactOpts, fixedPart, candidate)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 destination, uint256 expectedHeld, uint256 amount) payable returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, destination [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IYellowAdjudicator.contract.Transact(opts, "deposit", asset, destination, expectedHeld, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 destination, uint256 expectedHeld, uint256 amount) payable returns()
func (_IYellowAdjudicator *IYellowAdjudicatorSession) Deposit(asset common.Address, destination [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Deposit(&_IYellowAdjudicator.TransactOpts, asset, destination, expectedHeld, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x2fb1d270.
//
// Solidity: function deposit(address asset, bytes32 destination, uint256 expectedHeld, uint256 amount) payable returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactorSession) Deposit(asset common.Address, destination [32]byte, expectedHeld *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Deposit(&_IYellowAdjudicator.TransactOpts, asset, destination, expectedHeld, amount)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactor) Reclaim(opts *bind.TransactOpts, reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _IYellowAdjudicator.contract.Transact(opts, "reclaim", reclaimArgs)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorSession) Reclaim(reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Reclaim(&_IYellowAdjudicator.TransactOpts, reclaimArgs)
}

// Reclaim is a paid mutator transaction binding the contract method 0xd3c4e738.
//
// Solidity: function reclaim((bytes32,bytes32,bytes,uint256,uint256,bytes32,bytes,uint256) reclaimArgs) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactorSession) Reclaim(reclaimArgs IMultiAssetHolderReclaimArgs) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Reclaim(&_IYellowAdjudicator.TransactOpts, reclaimArgs)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactor) Transfer(opts *bind.TransactOpts, assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _IYellowAdjudicator.contract.Transact(opts, "transfer", assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorSession) Transfer(assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Transfer(&_IYellowAdjudicator.TransactOpts, assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// Transfer is a paid mutator transaction binding the contract method 0x3033730e.
//
// Solidity: function transfer(uint256 assetIndex, bytes32 fromChannelId, bytes outcomeBytes, bytes32 stateHash, uint256[] indices) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactorSession) Transfer(assetIndex *big.Int, fromChannelId [32]byte, outcomeBytes []byte, stateHash [32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.Transfer(&_IYellowAdjudicator.TransactOpts, assetIndex, fromChannelId, outcomeBytes, stateHash, indices)
}

// TransferAllAssets is a paid mutator transaction binding the contract method 0x31afa0b4.
//
// Solidity: function transferAllAssets(bytes32 channelId, (address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[] outcome, bytes32 stateHash) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactor) TransferAllAssets(opts *bind.TransactOpts, channelId [32]byte, outcome []ExitFormatSingleAssetExit, stateHash [32]byte) (*types.Transaction, error) {
	return _IYellowAdjudicator.contract.Transact(opts, "transferAllAssets", channelId, outcome, stateHash)
}

// TransferAllAssets is a paid mutator transaction binding the contract method 0x31afa0b4.
//
// Solidity: function transferAllAssets(bytes32 channelId, (address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[] outcome, bytes32 stateHash) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorSession) TransferAllAssets(channelId [32]byte, outcome []ExitFormatSingleAssetExit, stateHash [32]byte) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.TransferAllAssets(&_IYellowAdjudicator.TransactOpts, channelId, outcome, stateHash)
}

// TransferAllAssets is a paid mutator transaction binding the contract method 0x31afa0b4.
//
// Solidity: function transferAllAssets(bytes32 channelId, (address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[] outcome, bytes32 stateHash) returns()
func (_IYellowAdjudicator *IYellowAdjudicatorTransactorSession) TransferAllAssets(channelId [32]byte, outcome []ExitFormatSingleAssetExit, stateHash [32]byte) (*types.Transaction, error) {
	return _IYellowAdjudicator.Contract.TransferAllAssets(&_IYellowAdjudicator.TransactOpts, channelId, outcome, stateHash)
}

// IYellowAdjudicatorAllocationUpdatedIterator is returned from FilterAllocationUpdated and is used to iterate over the raw logs and unpacked data for AllocationUpdated events raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorAllocationUpdatedIterator struct {
	Event *IYellowAdjudicatorAllocationUpdated // Event containing the contract specifics and raw log

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
func (it *IYellowAdjudicatorAllocationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IYellowAdjudicatorAllocationUpdated)
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
		it.Event = new(IYellowAdjudicatorAllocationUpdated)
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
func (it *IYellowAdjudicatorAllocationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IYellowAdjudicatorAllocationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IYellowAdjudicatorAllocationUpdated represents a AllocationUpdated event raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorAllocationUpdated struct {
	ChannelId       [32]byte
	AssetIndex      *big.Int
	InitialHoldings *big.Int
	FinalHoldings   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAllocationUpdated is a free log retrieval operation binding the contract event 0xc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a.
//
// Solidity: event AllocationUpdated(bytes32 indexed channelId, uint256 assetIndex, uint256 initialHoldings, uint256 finalHoldings)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) FilterAllocationUpdated(opts *bind.FilterOpts, channelId [][32]byte) (*IYellowAdjudicatorAllocationUpdatedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.FilterLogs(opts, "AllocationUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &IYellowAdjudicatorAllocationUpdatedIterator{contract: _IYellowAdjudicator.contract, event: "AllocationUpdated", logs: logs, sub: sub}, nil
}

// WatchAllocationUpdated is a free log subscription operation binding the contract event 0xc36da2054c5669d6dac211b7366d59f2d369151c21edf4940468614b449e0b9a.
//
// Solidity: event AllocationUpdated(bytes32 indexed channelId, uint256 assetIndex, uint256 initialHoldings, uint256 finalHoldings)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) WatchAllocationUpdated(opts *bind.WatchOpts, sink chan<- *IYellowAdjudicatorAllocationUpdated, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.WatchLogs(opts, "AllocationUpdated", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IYellowAdjudicatorAllocationUpdated)
				if err := _IYellowAdjudicator.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
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
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) ParseAllocationUpdated(log types.Log) (*IYellowAdjudicatorAllocationUpdated, error) {
	event := new(IYellowAdjudicatorAllocationUpdated)
	if err := _IYellowAdjudicator.contract.UnpackLog(event, "AllocationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IYellowAdjudicatorChallengeClearedIterator is returned from FilterChallengeCleared and is used to iterate over the raw logs and unpacked data for ChallengeCleared events raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorChallengeClearedIterator struct {
	Event *IYellowAdjudicatorChallengeCleared // Event containing the contract specifics and raw log

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
func (it *IYellowAdjudicatorChallengeClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IYellowAdjudicatorChallengeCleared)
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
		it.Event = new(IYellowAdjudicatorChallengeCleared)
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
func (it *IYellowAdjudicatorChallengeClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IYellowAdjudicatorChallengeClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IYellowAdjudicatorChallengeCleared represents a ChallengeCleared event raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorChallengeCleared struct {
	ChannelId        [32]byte
	NewTurnNumRecord *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterChallengeCleared is a free log retrieval operation binding the contract event 0x07da0a0674fb921e484018c8b81d80e292745e5d8ed134b580c8b9c631c5e9e0.
//
// Solidity: event ChallengeCleared(bytes32 indexed channelId, uint48 newTurnNumRecord)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) FilterChallengeCleared(opts *bind.FilterOpts, channelId [][32]byte) (*IYellowAdjudicatorChallengeClearedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.FilterLogs(opts, "ChallengeCleared", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &IYellowAdjudicatorChallengeClearedIterator{contract: _IYellowAdjudicator.contract, event: "ChallengeCleared", logs: logs, sub: sub}, nil
}

// WatchChallengeCleared is a free log subscription operation binding the contract event 0x07da0a0674fb921e484018c8b81d80e292745e5d8ed134b580c8b9c631c5e9e0.
//
// Solidity: event ChallengeCleared(bytes32 indexed channelId, uint48 newTurnNumRecord)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) WatchChallengeCleared(opts *bind.WatchOpts, sink chan<- *IYellowAdjudicatorChallengeCleared, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.WatchLogs(opts, "ChallengeCleared", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IYellowAdjudicatorChallengeCleared)
				if err := _IYellowAdjudicator.contract.UnpackLog(event, "ChallengeCleared", log); err != nil {
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
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) ParseChallengeCleared(log types.Log) (*IYellowAdjudicatorChallengeCleared, error) {
	event := new(IYellowAdjudicatorChallengeCleared)
	if err := _IYellowAdjudicator.contract.UnpackLog(event, "ChallengeCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IYellowAdjudicatorChallengeRegisteredIterator is returned from FilterChallengeRegistered and is used to iterate over the raw logs and unpacked data for ChallengeRegistered events raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorChallengeRegisteredIterator struct {
	Event *IYellowAdjudicatorChallengeRegistered // Event containing the contract specifics and raw log

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
func (it *IYellowAdjudicatorChallengeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IYellowAdjudicatorChallengeRegistered)
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
		it.Event = new(IYellowAdjudicatorChallengeRegistered)
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
func (it *IYellowAdjudicatorChallengeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IYellowAdjudicatorChallengeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IYellowAdjudicatorChallengeRegistered represents a ChallengeRegistered event raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorChallengeRegistered struct {
	ChannelId   [32]byte
	FinalizesAt *big.Int
	Proof       []INitroTypesSignedVariablePart
	Candidate   INitroTypesSignedVariablePart
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChallengeRegistered is a free log retrieval operation binding the contract event 0x0aa12461ee6c137332989aa12cec79f4772ab2c1a8732a382aada7e9f3ec9d34.
//
// Solidity: event ChallengeRegistered(bytes32 indexed channelId, uint48 finalizesAt, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) FilterChallengeRegistered(opts *bind.FilterOpts, channelId [][32]byte) (*IYellowAdjudicatorChallengeRegisteredIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.FilterLogs(opts, "ChallengeRegistered", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &IYellowAdjudicatorChallengeRegisteredIterator{contract: _IYellowAdjudicator.contract, event: "ChallengeRegistered", logs: logs, sub: sub}, nil
}

// WatchChallengeRegistered is a free log subscription operation binding the contract event 0x0aa12461ee6c137332989aa12cec79f4772ab2c1a8732a382aada7e9f3ec9d34.
//
// Solidity: event ChallengeRegistered(bytes32 indexed channelId, uint48 finalizesAt, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[])[] proof, (((address,(uint8,bytes),(bytes32,uint256,uint8,bytes)[])[],bytes,uint48,bool),(uint8,bytes32,bytes32)[]) candidate)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) WatchChallengeRegistered(opts *bind.WatchOpts, sink chan<- *IYellowAdjudicatorChallengeRegistered, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.WatchLogs(opts, "ChallengeRegistered", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IYellowAdjudicatorChallengeRegistered)
				if err := _IYellowAdjudicator.contract.UnpackLog(event, "ChallengeRegistered", log); err != nil {
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
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) ParseChallengeRegistered(log types.Log) (*IYellowAdjudicatorChallengeRegistered, error) {
	event := new(IYellowAdjudicatorChallengeRegistered)
	if err := _IYellowAdjudicator.contract.UnpackLog(event, "ChallengeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IYellowAdjudicatorConcludedIterator is returned from FilterConcluded and is used to iterate over the raw logs and unpacked data for Concluded events raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorConcludedIterator struct {
	Event *IYellowAdjudicatorConcluded // Event containing the contract specifics and raw log

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
func (it *IYellowAdjudicatorConcludedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IYellowAdjudicatorConcluded)
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
		it.Event = new(IYellowAdjudicatorConcluded)
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
func (it *IYellowAdjudicatorConcludedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IYellowAdjudicatorConcludedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IYellowAdjudicatorConcluded represents a Concluded event raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorConcluded struct {
	ChannelId   [32]byte
	FinalizesAt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConcluded is a free log retrieval operation binding the contract event 0x4f465027a3d06ea73dd12be0f5c5fc0a34e21f19d6eaed4834a7a944edabc901.
//
// Solidity: event Concluded(bytes32 indexed channelId, uint48 finalizesAt)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) FilterConcluded(opts *bind.FilterOpts, channelId [][32]byte) (*IYellowAdjudicatorConcludedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.FilterLogs(opts, "Concluded", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &IYellowAdjudicatorConcludedIterator{contract: _IYellowAdjudicator.contract, event: "Concluded", logs: logs, sub: sub}, nil
}

// WatchConcluded is a free log subscription operation binding the contract event 0x4f465027a3d06ea73dd12be0f5c5fc0a34e21f19d6eaed4834a7a944edabc901.
//
// Solidity: event Concluded(bytes32 indexed channelId, uint48 finalizesAt)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) WatchConcluded(opts *bind.WatchOpts, sink chan<- *IYellowAdjudicatorConcluded, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.WatchLogs(opts, "Concluded", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IYellowAdjudicatorConcluded)
				if err := _IYellowAdjudicator.contract.UnpackLog(event, "Concluded", log); err != nil {
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
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) ParseConcluded(log types.Log) (*IYellowAdjudicatorConcluded, error) {
	event := new(IYellowAdjudicatorConcluded)
	if err := _IYellowAdjudicator.contract.UnpackLog(event, "Concluded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IYellowAdjudicatorDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorDepositedIterator struct {
	Event *IYellowAdjudicatorDeposited // Event containing the contract specifics and raw log

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
func (it *IYellowAdjudicatorDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IYellowAdjudicatorDeposited)
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
		it.Event = new(IYellowAdjudicatorDeposited)
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
func (it *IYellowAdjudicatorDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IYellowAdjudicatorDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IYellowAdjudicatorDeposited represents a Deposited event raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorDeposited struct {
	Destination         [32]byte
	Asset               common.Address
	DestinationHoldings *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b.
//
// Solidity: event Deposited(bytes32 indexed destination, address asset, uint256 destinationHoldings)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) FilterDeposited(opts *bind.FilterOpts, destination [][32]byte) (*IYellowAdjudicatorDepositedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.FilterLogs(opts, "Deposited", destinationRule)
	if err != nil {
		return nil, err
	}
	return &IYellowAdjudicatorDepositedIterator{contract: _IYellowAdjudicator.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x87d4c0b5e30d6808bc8a94ba1c4d839b29d664151551a31753387ee9ef48429b.
//
// Solidity: event Deposited(bytes32 indexed destination, address asset, uint256 destinationHoldings)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *IYellowAdjudicatorDeposited, destination [][32]byte) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.WatchLogs(opts, "Deposited", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IYellowAdjudicatorDeposited)
				if err := _IYellowAdjudicator.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) ParseDeposited(log types.Log) (*IYellowAdjudicatorDeposited, error) {
	event := new(IYellowAdjudicatorDeposited)
	if err := _IYellowAdjudicator.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IYellowAdjudicatorReclaimedIterator is returned from FilterReclaimed and is used to iterate over the raw logs and unpacked data for Reclaimed events raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorReclaimedIterator struct {
	Event *IYellowAdjudicatorReclaimed // Event containing the contract specifics and raw log

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
func (it *IYellowAdjudicatorReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IYellowAdjudicatorReclaimed)
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
		it.Event = new(IYellowAdjudicatorReclaimed)
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
func (it *IYellowAdjudicatorReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IYellowAdjudicatorReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IYellowAdjudicatorReclaimed represents a Reclaimed event raised by the IYellowAdjudicator contract.
type IYellowAdjudicatorReclaimed struct {
	ChannelId  [32]byte
	AssetIndex *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReclaimed is a free log retrieval operation binding the contract event 0x4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b.
//
// Solidity: event Reclaimed(bytes32 indexed channelId, uint256 assetIndex)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) FilterReclaimed(opts *bind.FilterOpts, channelId [][32]byte) (*IYellowAdjudicatorReclaimedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.FilterLogs(opts, "Reclaimed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &IYellowAdjudicatorReclaimedIterator{contract: _IYellowAdjudicator.contract, event: "Reclaimed", logs: logs, sub: sub}, nil
}

// WatchReclaimed is a free log subscription operation binding the contract event 0x4d3754632451ebba9812a9305e7bca17b67a17186a5cff93d2e9ae1b01e3d27b.
//
// Solidity: event Reclaimed(bytes32 indexed channelId, uint256 assetIndex)
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) WatchReclaimed(opts *bind.WatchOpts, sink chan<- *IYellowAdjudicatorReclaimed, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IYellowAdjudicator.contract.WatchLogs(opts, "Reclaimed", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IYellowAdjudicatorReclaimed)
				if err := _IYellowAdjudicator.contract.UnpackLog(event, "Reclaimed", log); err != nil {
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
func (_IYellowAdjudicator *IYellowAdjudicatorFilterer) ParseReclaimed(log types.Log) (*IYellowAdjudicatorReclaimed, error) {
	event := new(IYellowAdjudicatorReclaimed)
	if err := _IYellowAdjudicator.contract.UnpackLog(event, "Reclaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
