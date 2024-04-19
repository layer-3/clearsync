// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package secta_v3_router

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

// IApproveAndCallIncreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IApproveAndCallIncreaseLiquidityParams struct {
	Token0     common.Address
	Token1     common.Address
	TokenId    *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
}

// IApproveAndCallMintParams is an auto generated low-level Go binding around an user-defined struct.
type IApproveAndCallMintParams struct {
	Token0     common.Address
	Token1     common.Address
	Fee        *big.Int
	TickLower  *big.Int
	TickUpper  *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
	Recipient  common.Address
}

// IV3SwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// IV3SwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IV3SwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// IV3SwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// SectaV3RouterMetaData contains all meta data concerning the SectaV3Router contract.
var SectaV3RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factoryV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryV3\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_positionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stableFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stableInfo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"info\",\"type\":\"address\"}],\"name\":\"SetStableSwap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveMax\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveMaxMinusOne\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveZeroThenMax\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveZeroThenMaxMinusOne\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callPositionManager\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"paths\",\"type\":\"bytes[]\"},{\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"},{\"internalType\":\"uint24\",\"name\":\"maximumTickDivergence\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"secondsAgo\",\"type\":\"uint32\"}],\"name\":\"checkOracleSlippage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint24\",\"name\":\"maximumTickDivergence\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"secondsAgo\",\"type\":\"uint32\"}],\"name\":\"checkOracleSlippage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structIV3SwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIV3SwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"flag\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"exactInputStableSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structIV3SwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIV3SwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"flag\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"exactOutputStableSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getApprovalType\",\"outputs\":[{\"internalType\":\"enumIApproveAndCall.ApprovalType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"}],\"internalType\":\"structIApproveAndCall.IncreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"increaseLiquidity\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIApproveAndCall.MintParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"previousBlockhash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"pancakeV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_info\",\"type\":\"address\"}],\"name\":\"setStableSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableSwapFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableSwapInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"wrapETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101206040526000196002553480156200001857600080fd5b506040516200616d3803806200616d8339810160408190526200003b9162000125565b6001600160601b0319606088811b821660805285811b821660a05287811b821660c05286811b821660e05282901b1661010052828260006200007c62000104565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060018055600380546001600160a01b039384166001600160a01b0319918216179091556004805492909316911617905550620001b995505050505050565b3390565b80516001600160a01b03811681146200012057600080fd5b919050565b600080600080600080600060e0888a03121562000140578283fd5b6200014b8862000108565b96506200015b6020890162000108565b95506200016b6040890162000108565b94506200017b6060890162000108565b93506200018b6080890162000108565b92506200019b60a0890162000108565b9150620001ab60c0890162000108565b905092959891949750929550565b60805160601c60a05160601c60c05160601c60e05160601c6101005160601c615efc620002716000398061022c5280610a7452806114a85280611593528061162052806118db52806119c65280613002528061304852806130bc525080612441525080610b9f52806124715280612a845280612c9f52806144cb52508061188c5280611c8652806124a8528061365e525080610e825280610f48528061125852806117ce528061329f52806135075250615efc6000f3fe60806040526004361061021c5760003560e01c806304e45aaf1461029157806309b81346146102ba57806311ed56c9146102cd57806312210e8a146102ed5780631c58db4f146102f55780631f0464d11461030857806323a69e751461032857806324dec034146103485780633068c5541461036857806342712a671461037b5780634659a4941461038e578063472b43f3146103a157806349404b7c146103b45780634aa4a4fc146103c75780635023b4df146103e9578063571ac8b0146103fc57806357c799611461040f5780635ae401dc14610424578063639d71a91461043757806368e0d4e11461044a578063715018a61461045f578063791b98bc146104745780638da5cb5b146104895780639b2c0a371461049e578063a4a78f0c146104b1578063ab3fdd50146104c4578063ac9650d8146104d7578063b3a2af13146104ea578063b4554231146104fd578063b4c4e55514610510578063b858183f14610523578063b85aa7af14610536578063c2e3140a1461054b578063c45a01551461055e578063cab372ce14610573578063d4ef38de14610586578063d5f3948814610599578063dee00f35146105ae578063df2ab5bb146105db578063e0e189a0146105ee578063e90a182f14610601578063efdeed8e14610614578063f100b20514610634578063f25801a714610647578063f2d5d56b14610667578063f2fde38b1461067a578063f3995c671461069a5761028c565b3661028c57336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461028a576040805162461bcd60e51b81526020600482015260096024820152684e6f7420574554483960b81b604482015290519081900360640190fd5b005b600080fd5b6102a461029f36600461545e565b6106ad565b6040516102b19190615cd2565b60405180910390f35b6102a46102c83660046154f8565b610819565b6102e06102db366004615551565b6108fd565b6040516102b19190615b6d565b61028a610a60565b61028a6103033660046156d3565b610a72565b61031b6103163660046151c8565b610ae9565b6040516102b19190615b0d565b34801561033457600080fd5b5061028a6103433660046152d6565b610b43565b34801561035457600080fd5b5061028a610363366004614da4565b610ce0565b61028a610376366004614ed4565b610dcf565b6102a46103893660046157bf565b610de2565b61028a61039c366004614f1d565b61107a565b6102a46103af3660046157bf565b611108565b61028a6103c2366004615703565b6114a4565b3480156103d357600080fd5b506103dc61161e565b6040516102b19190615962565b6102a46103f736600461552f565b611642565b61028a61040a366004614d65565b61173c565b34801561041b57600080fd5b506103dc611754565b61031b6104323660046151c8565b611763565b61028a610445366004614d65565b6117b8565b34801561045657600080fd5b506103dc6117cc565b34801561046b57600080fd5b5061028a6117f0565b34801561048057600080fd5b506103dc61188a565b34801561049557600080fd5b506103dc6118ae565b61028a6104ac366004615727565b6118bd565b61028a6104bf366004614f1d565b611a87565b61028a6104d2366004614d65565b611b18565b61031b6104e5366004615010565b611b38565b6102e06104f8366004615211565b611c80565b6102a461050b366004614f7d565b611d37565b6102a461051e366004614f7d565b611fc1565b6102a46105313660046153b9565b6121f1565b34801561054257600080fd5b506103dc6123a1565b61028a610559366004614f1d565b6123b0565b34801561056a57600080fd5b506103dc61243f565b61028a610581366004614d65565b611b2c565b61028a610594366004615765565b612463565b3480156105a557600080fd5b506103dc61246f565b3480156105ba57600080fd5b506105ce6105c9366004614e0a565b612493565b6040516102b19190615b80565b61028a6105e9366004614e35565b6125bb565b61028a6105fc366004614e76565b612693565b61028a61060f366004614e0a565b6127ba565b34801561062057600080fd5b5061028a61062f36600461504f565b6127c9565b6102e0610642366004615540565b6127ee565b34801561065357600080fd5b5061028a610662366004615243565b612870565b61028a610675366004614e0a565b612894565b34801561068657600080fd5b5061028a610695366004614d65565b6128a0565b61028a6106a8366004614f1d565b612990565b6000600260015414156106f5576040805162461bcd60e51b815260206004820152601f6024820152600080516020615e41833981519152604482015290519081900360640190fd5b6002600155608082015160009061078d575081516040516370a0823160e01b81526001916001600160a01b0316906370a0823190610737903090600401615962565b60206040518083038186803b15801561074f57600080fd5b505afa158015610763573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078791906156eb565b60808401525b6107fc836080015184606001518560c001516040518060400160405280886000015189604001518a602001516040516020016107cb93929190615900565b6040516020818303038152906040528152602001866107ea57336107ec565b305b6001600160a01b03169052612a02565b91508260a0015182101561080f57600080fd5b5060018055919050565b600060026001541415610861576040805162461bcd60e51b815260206004820152601f6024820152600080516020615e41833981519152604482015290519081900360640190fd5b60026001556108d86040830180359061087d9060208601614d65565b60408051808201909152600090806108958880615d39565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525033602090910152612c1d565b505060025460608201358111156108ee57600080fd5b60001960025560018055919050565b604080516101608101909152606090610a5890634418b22b60e11b90806109276020870187614d65565b6001600160a01b031681526020018560200160208101906109489190614d65565b6001600160a01b0316815260200161096660608701604088016156b9565b62ffffff1681526020016109806080870160608801615297565b60020b815260200161099860a0870160808801615297565b60020b81526020908101906109b8906109b390880188614d65565b612e61565b81526020016109d38660200160208101906109b39190614d65565b815260a0860135602082015260c086013560408201526060016109fd610100870160e08801614d65565b6001600160a01b03168152602001600019815250604051602401610a219190615bd8565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152611c80565b90505b919050565b4715610a7057610a703347612ee0565b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b158015610acd57600080fd5b505af1158015610ae1573d6000803e3d6000fd5b505050505050565b60608380600143034014610b30576040805162461bcd60e51b8152602060048201526009602482015268084d8dec6d6d0c2e6d60bb1b604482015290519081900360640190fd5b610b3a8484611b38565b95945050505050565b6000841380610b525750600083135b610b5b57600080fd5b6000610b6982840184615563565b90506000806000610b7d8460000151612fcf565b92509250925073__$d75017e8592c813c72614beebf2a3b9fe6$__638bdb19257f00000000000000000000000000000000000000000000000000000000000000008585856040518563ffffffff1660e01b8152600401610be094939291906159b3565b60206040518083038186803b158015610bf857600080fd5b505af4158015610c0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c309190614d88565b5060008060008a13610c5757846001600160a01b0316846001600160a01b03161089610c6e565b836001600160a01b0316856001600160a01b0316108a5b915091508115610c8d57610c888587602001513384613000565b610cd4565b8551610c9890613190565b15610cbd578551610ca890613198565b8652610cb78133600089612c1d565b50610cd4565b80600281905550610cd48487602001513384613000565b50505050505050505050565b610ce86131af565b6001600160a01b0316610cf96118ae565b6001600160a01b031614610d42576040805162461bcd60e51b81526020600482018190526024820152600080516020615e87833981519152604482015290519081900360640190fd5b6001600160a01b03821615801590610d6257506001600160a01b03811615155b610d6b57600080fd5b600380546001600160a01b038085166001600160a01b0319928316179283905560048054858316931692909217918290556040519181169216907f26e41379222b54b0470031bc11852ad23058ffb8983f7cc0e18257d6f7afca9d90600090a35050565b610ddc8484338585612693565b50505050565b600060026001541415610e2a576040805162461bcd60e51b815260206004820152601f6024820152600080516020615e41833981519152604482015290519081900360640190fd5b6002600155600084848281610e3b57fe5b9050602002016020810190610e509190614d65565b604051630c90945960e11b815290915073__$d75017e8592c813c72614beebf2a3b9fe6$__9063192128b290610eb0907f0000000000000000000000000000000000000000000000000000000000000000908b908a908a90600401615adb565b60006040518083038186803b158015610ec857600080fd5b505af4158015610edc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f04919081019061511f565b600081518110610f1057fe5b6020026020010151915085821115610f2757600080fd5b610ffe813373__$d75017e8592c813c72614beebf2a3b9fe6$__636d91c0e27f0000000000000000000000000000000000000000000000000000000000000000868b8b6001818110610f7557fe5b9050602002016020810190610f8a9190614d65565b6040518463ffffffff1660e01b8152600401610fa893929190615990565b60206040518083038186803b158015610fc057600080fd5b505af4158015610fd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff89190614d88565b85613000565b6001600160a01b038316600114156110185733925061102e565b6001600160a01b0383166002141561102e573092505b61106c8585808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508792506131b3915050565b506001805595945050505050565b604080516323f2ebc360e21b815233600482015230602482015260448101879052606481018690526001608482015260ff851660a482015260c4810184905260e4810183905290516001600160a01b03881691638fcbaf0c9161010480830192600092919082900301818387803b1580156110f457600080fd5b505af1158015610cd4573d6000803e3d6000fd5b600060026001541415611150576040805162461bcd60e51b815260206004820152601f6024820152600080516020615e41833981519152604482015290519081900360640190fd5b600260015560008484828161116157fe5b90506020020160208101906111769190614d65565b905060008585600019810181811061118a57fe5b905060200201602081019061119f9190614d65565b905060008861122b57506040516370a0823160e01b81526001906001600160a01b038416906370a08231906111d8903090600401615962565b60206040518083038186803b1580156111f057600080fd5b505afa158015611204573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061122891906156eb565b98505b61130e838261123a573361123c565b305b73__$d75017e8592c813c72614beebf2a3b9fe6$__636d91c0e27f0000000000000000000000000000000000000000000000000000000000000000888d8d600181811061128557fe5b905060200201602081019061129a9190614d65565b6040518463ffffffff1660e01b81526004016112b893929190615990565b60206040518083038186803b1580156112d057600080fd5b505af41580156112e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113089190614d88565b8c613000565b6001600160a01b038516600114156113285733945061133e565b6001600160a01b0385166002141561133e573094505b6040516370a0823160e01b81526000906001600160a01b038416906370a082319061136d908990600401615962565b60206040518083038186803b15801561138557600080fd5b505afa158015611399573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113bd91906156eb565b90506113fd8888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a92506131b3915050565b61148381846001600160a01b03166370a08231896040518263ffffffff1660e01b815260040161142d9190615962565b60206040518083038186803b15801561144557600080fd5b505afa158015611459573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061147d91906156eb565b90613635565b94508885101561149257600080fd5b50506001805550909695505050505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561151357600080fd5b505afa158015611527573d6000803e3d6000fd5b505050506040513d602081101561153d57600080fd5b505190508281101561158b576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e7420574554483960701b604482015290519081900360640190fd5b8015611619577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b1580156115f757600080fd5b505af115801561160b573d6000803e3d6000fd5b505050506116198282612ee0565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006002600154141561168a576040805162461bcd60e51b815260206004820152601f6024820152600080516020615e41833981519152604482015290519081900360640190fd5b6002600155611729608083018035906116a69060608601614d65565b6116b660e0860160c08701614d65565b60405180604001604052808760200160208101906116d49190614d65565b6116e460608a0160408b016156b9565b6116f160208b018b614d65565b60405160200161170393929190615900565b6040516020818303038152906040528152602001336001600160a01b0316815250612c1d565b90508160a001358111156108ee57600080fd5b61174881600019613645565b61175157600080fd5b50565b6003546001600160a01b031681565b6060838061176f613739565b1115610b30576040805162461bcd60e51b8152602060048201526013602482015272151c985b9cd858dd1a5bdb881d1bdbc81bdb19606a1b604482015290519081900360640190fd5b6117c3816000613645565b61173c57600080fd5b7f000000000000000000000000000000000000000000000000000000000000000081565b6117f86131af565b6001600160a01b03166118096118ae565b6001600160a01b031614611852576040805162461bcd60e51b81526020600482018190526024820152600080516020615e87833981519152604482015290519081900360640190fd5b600080546040516001600160a01b0390911690600080516020615ea7833981519152908390a3600080546001600160a01b0319169055565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000546001600160a01b031690565b6000821180156118ce575060648211155b6118d757600080fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561194657600080fd5b505afa15801561195a573d6000803e3d6000fd5b505050506040513d602081101561197057600080fd5b50519050848110156119be576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e7420574554483960701b604482015290519081900360640190fd5b8015611a80577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015611a2a57600080fd5b505af1158015611a3e573d6000803e3d6000fd5b505050506000612710611a5a858461373d90919063ffffffff16565b81611a6157fe5b0490508015611a7457611a748382612ee0565b610ae185828403612ee0565b5050505050565b60408051636eb1769f60e11b81523360048201523060248201529051600019916001600160a01b0389169163dd62ed3e91604480820192602092909190829003018186803b158015611ad857600080fd5b505afa158015611aec573d6000803e3d6000fd5b505050506040513d6020811015611b0257600080fd5b50511015610ae157610ae186868686868661107a565b611b23816000613645565b611b2c57600080fd5b61174881600119613645565b6060816001600160401b0381118015611b5057600080fd5b50604051908082528060200260200182016040528015611b8457816020015b6060815260200190600190039081611b6f5790505b50905060005b82811015611c795760008030868685818110611ba257fe5b9050602002810190611bb49190615d39565b604051611bc2929190615936565b600060405180830381855af49150503d8060008114611bfd576040519150601f19603f3d011682016040523d82523d6000602084013e611c02565b606091505b509150915081611c5757604481511015611c1b57600080fd5b60048101905080806020019051810190611c359190615350565b60405162461bcd60e51b8152600401611c4e9190615b6d565b60405180910390fd5b80848481518110611c6457fe5b60209081029190910101525050600101611b8a565b5092915050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031683604051611cbc9190615946565b6000604051808303816000865af19150503d8060008114611cf9576040519150601f19603f3d011682016040523d82523d6000602084013e611cfe565b606091505b509250905080611d3157604482511015611d1757600080fd5b60048201915081806020019051810190611c359190615350565b50919050565b600060026001541415611d7f576040805162461bcd60e51b815260206004820152601f6024820152600080516020615e41833981519152604482015290519081900360640190fd5b6002600155600088888281611d9057fe5b9050602002016020810190611da59190614d65565b9050600089896000198101818110611db957fe5b9050602002016020810190611dce9190614d65565b9050600086611e5a57506040516370a0823160e01b81526001906001600160a01b038416906370a0823190611e07903090600401615962565b60206040518083038186803b158015611e1f57600080fd5b505afa158015611e33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e5791906156eb565b96505b80611e6b57611e6b8333308a613000565b611ed88b8b8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808f0282810182019093528e82529093508e92508d91829185019084908082843760009201919091525061376192505050565b6040516370a0823160e01b81526001600160a01b038316906370a0823190611f04903090600401615962565b60206040518083038186803b158015611f1c57600080fd5b505afa158015611f30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f5491906156eb565b935085841015611f6357600080fd5b6001600160a01b03851660011415611f7d57339450611f93565b6001600160a01b03851660021415611f93573094505b6001600160a01b0385163014611faf57611faf82308787613000565b50506001805550979650505050505050565b600060026001541415612009576040805162461bcd60e51b815260206004820152601f6024820152600080516020615e41833981519152604482015290519081900360640190fd5b600260015560035460048054604051635923cab360e01b815273__$d75017e8592c813c72614beebf2a3b9fe6$__93635923cab393612061936001600160a01b03928316939216918e918e918e918e918e9101615a0b565b60006040518083038186803b15801561207957600080fd5b505af415801561208d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526120b5919081019061511f565b6000815181106120c157fe5b60200260200101519050828111156120d857600080fd5b612105888860008181106120e857fe5b90506020020160208101906120fd9190614d65565b333084613000565b61217288888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808c0282810182019093528b82529093508b92508a91829185019084908082843760009201919091525061376192505050565b6001600160a01b0382166001141561218c573391506121a2565b6001600160a01b038216600214156121a2573091505b6001600160a01b03821630146121e2576121e2888860001981018181106121c557fe5b90506020020160208101906121da9190614d65565b308487613000565b60018055979650505050505050565b600060026001541415612239576040805162461bcd60e51b815260206004820152601f6024820152600080516020615e41833981519152604482015290519081900360640190fd5b600260015560408201516000906122e55760019050600061225d8460000151612fcf565b50506040516370a0823160e01b81529091506001600160a01b038216906370a082319061228e903090600401615962565b60206040518083038186803b1580156122a657600080fd5b505afa1580156122ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122de91906156eb565b6040850152505b6000816122f257336122f4565b305b90505b60006123068560000151613190565b905061235285604001518261231f578660200151612321565b305b6000604051806040016040528061233b8b60000151613972565b8152602001876001600160a01b0316815250612a02565b6040860152801561237257845130925061236b90613198565b855261237f565b8460400151935050612385565b506122f7565b836060015183101561239657600080fd5b505060018055919050565b6004546001600160a01b031681565b60408051636eb1769f60e11b8152336004820152306024820152905186916001600160a01b0389169163dd62ed3e91604480820192602092909190829003018186803b1580156123ff57600080fd5b505afa158015612413573d6000803e3d6000fd5b505050506040513d602081101561242957600080fd5b50511015610ae157610ae1868686868686612990565b7f000000000000000000000000000000000000000000000000000000000000000081565b611619833384846118bd565b7f000000000000000000000000000000000000000000000000000000000000000081565b600081836001600160a01b031663dd62ed3e307f00000000000000000000000000000000000000000000000000000000000000006040518363ffffffff1660e01b81526004016124e4929190615976565b60206040518083038186803b1580156124fc57600080fd5b505afa158015612510573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061253491906156eb565b10612541575060006125b5565b61254d83600019613645565b1561255a575060016125b5565b61256683600119613645565b15612573575060026125b5565b61257e836000613645565b61258757600080fd5b61259383600019613645565b156125a0575060036125b5565b6125ac83600119613645565b1561028c575060045b92915050565b6000836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561260a57600080fd5b505afa15801561261e573d6000803e3d6000fd5b505050506040513d602081101561263457600080fd5b5051905082811015612682576040805162461bcd60e51b815260206004820152601260248201527124b739bab33334b1b4b2b73a103a37b5b2b760711b604482015290519081900360640190fd5b8015610ddc57610ddc848383613981565b6000821180156126a4575060648211155b6126ad57600080fd5b6000856001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156126fc57600080fd5b505afa158015612710573d6000803e3d6000fd5b505050506040513d602081101561272657600080fd5b5051905084811015612774576040805162461bcd60e51b815260206004820152601260248201527124b739bab33334b1b4b2b73a103a37b5b2b760711b604482015290519081900360640190fd5b8015610ae1576000612710612789838661373d565b8161279057fe5b04905080156127a4576127a4878483613981565b6127b18786838503613981565b50505050505050565b6127c58282336125bb565b5050565b6000806127d7868685613ac8565b915091508362ffffff1681830312610ae157600080fd5b6060610a5863219f5d1760e01b6040518060c00160405280856040013581526020016128268660000160208101906109b39190614d65565b81526020016128418660200160208101906109b39190614d65565b81526020018560600135815260200185608001358152602001600019815250604051602401610a219190615b94565b60008061287d8584613cbe565b915091508362ffffff1681830312611a8057600080fd5b6127c582333084613e95565b6128a86131af565b6001600160a01b03166128b96118ae565b6001600160a01b031614612902576040805162461bcd60e51b81526020600482018190526024820152600080516020615e87833981519152604482015290519081900360640190fd5b6001600160a01b0381166129475760405162461bcd60e51b8152600401808060200182810382526026815260200180615e616026913960400191505060405180910390fd5b600080546040516001600160a01b0380851693921691600080516020615ea783398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6040805163d505accf60e01b8152336004820152306024820152604481018790526064810186905260ff8516608482015260a4810184905260c4810183905290516001600160a01b0388169163d505accf9160e480830192600092919082900301818387803b1580156110f457600080fd5b60006001600160a01b03841660011415612a1e57339350612a34565b6001600160a01b03841660021415612a34573093505b6000806000612a468560000151612fcf565b9250925092506000826001600160a01b0316846001600160a01b031610905060008073__$d75017e8592c813c72614beebf2a3b9fe6$__634e6c8ed87f00000000000000000000000000000000000000000000000000000000000000008888886040518563ffffffff1660e01b8152600401612ac594939291906159b3565b60206040518083038186803b158015612add57600080fd5b505af4158015612af1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b159190614d88565b6001600160a01b031663128acb088b85612b2e8f613fe5565b6001600160a01b038e1615612b43578d612b69565b87612b625773fffd8963efd1fc6a506488495d951d5263988d25612b69565b6401000276a45b8d604051602001612b7a9190615c86565b6040516020818303038152906040526040518663ffffffff1660e01b8152600401612ba9959493929190615a7c565b6040805180830381600087803b158015612bc257600080fd5b505af1158015612bd6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bfa91906152b3565b9150915082612c095781612c0b565b805b6000039b9a5050505050505050505050565b60006001600160a01b03841660011415612c3957339350612c4f565b6001600160a01b03841660021415612c4f573093505b6000806000612c618560000151612fcf565b9250925092506000836001600160a01b0316836001600160a01b031610905060008073__$d75017e8592c813c72614beebf2a3b9fe6$__634e6c8ed87f00000000000000000000000000000000000000000000000000000000000000008789886040518563ffffffff1660e01b8152600401612ce094939291906159b3565b60206040518083038186803b158015612cf857600080fd5b505af4158015612d0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d309190614d88565b6001600160a01b031663128acb088b85612d498f613fe5565b6000036001600160a01b038e1615612d61578d612d87565b87612d805773fffd8963efd1fc6a506488495d951d5263988d25612d87565b6401000276a45b8d604051602001612d989190615c86565b6040516020818303038152906040526040518663ffffffff1660e01b8152600401612dc7959493929190615a7c565b6040805180830381600087803b158015612de057600080fd5b505af1158015612df4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e1891906152b3565b91509150600083612e2d578183600003612e33565b82826000035b90985090506001600160a01b038a16612e52578b8114612e5257600080fd5b50505050505050949350505050565b6040516370a0823160e01b81526000906001600160a01b038316906370a0823190612e90903090600401615962565b60206040518083038186803b158015612ea857600080fd5b505afa158015612ebc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5891906156eb565b604080516000808252602082019092526001600160a01b0384169083906040518082805190602001908083835b60208310612f2c5780518252601f199092019160209182019101612f0d565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114612f8e576040519150601f19603f3d011682016040523d82523d6000602084013e612f93565b606091505b5050905080611619576040805162461bcd60e51b815260206004820152600360248201526253544560e81b604482015290519081900360640190fd5b60008080612fdd8482613ffb565b9250612fea8460146140ab565b9050612ff7846017613ffb565b91509193909250565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316846001600160a01b03161480156130415750804710155b15613163577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b1580156130a157600080fd5b505af11580156130b5573d6000803e3d6000fd5b50505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb83836040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561313157600080fd5b505af1158015613145573d6000803e3d6000fd5b505050506040513d602081101561315b57600080fd5b50610ddc9050565b6001600160a01b0383163014156131845761317f848383613981565b610ddc565b610ddc84848484613e95565b516042111590565b8051606090610a5890839060179060161901614152565b3390565b60005b6001835103811015611619576000808483815181106131d157fe5b60200260200101518584600101815181106131e857fe5b602002602001015191509150600073__$d75017e8592c813c72614beebf2a3b9fe6$__63544caa5684846040518363ffffffff1660e01b815260040161322f929190615976565b604080518083038186803b15801561324657600080fd5b505af415801561325a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061327e9190614ddc565b509050600073__$d75017e8592c813c72614beebf2a3b9fe6$__636d91c0e27f000000000000000000000000000000000000000000000000000000000000000086866040518463ffffffff1660e01b81526004016132de93929190615990565b60206040518083038186803b1580156132f657600080fd5b505af415801561330a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061332e9190614d88565b9050600080600080846001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b15801561336f57600080fd5b505afa158015613383573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133a791906155f2565b506001600160701b031691506001600160701b03169150600080876001600160a01b03168a6001600160a01b0316146133e15782846133e4565b83835b91509150613418828b6001600160a01b03166370a082318a6040518263ffffffff1660e01b815260040161142d9190615962565b604051630153543560e21b815290965073__$d75017e8592c813c72614beebf2a3b9fe6$__9063054d50d49061345690899086908690600401615d08565b60206040518083038186803b15801561346e57600080fd5b505af4158015613482573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134a691906156eb565b945050505050600080856001600160a01b0316886001600160a01b0316146134d0578260006134d4565b6000835b91509150600060028c51038a106134eb578a6135ac565b73__$d75017e8592c813c72614beebf2a3b9fe6$__636d91c0e27f00000000000000000000000000000000000000000000000000000000000000008a8f8e6002018151811061353657fe5b60200260200101516040518463ffffffff1660e01b815260040161355c93929190615990565b60206040518083038186803b15801561357457600080fd5b505af4158015613588573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135ac9190614d88565b6040805160008152602081019182905263022c0d9f60e01b9091529091506001600160a01b0387169063022c0d9f906135ee9086908690869060248101615cdb565b600060405180830381600087803b15801561360857600080fd5b505af115801561361c573d6000803e3d6000fd5b50506001909b019a506131b69950505050505050505050565b808203828111156125b557600080fd5b6000806000846001600160a01b031663095ea7b360e01b7f00000000000000000000000000000000000000000000000000000000000000008660405160240161368f929190615ac2565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516136cd9190615946565b6000604051808303816000865af19150503d806000811461370a576040519150601f19603f3d011682016040523d82523d6000602084013e61370f565b606091505b5091509150818015610b3a575080511580610b3a575080806020019051810190610b3a91906151ae565b4290565b60008215806137585750508181028183828161375557fe5b04145b6125b557600080fd5b805160018351031461377257600080fd5b60005b81518110156116195760008084838151811061378d57fe5b60200260200101518584600101815181106137a457fe5b602002602001015191509150600080600073__$d75017e8592c813c72614beebf2a3b9fe6$__63b735aecd600360009054906101000a90046001600160a01b031687878b8b815181106137f357fe5b60200260200101516040518563ffffffff1660e01b815260040161381a94939291906159e1565b60606040518083038186803b15801561383257600080fd5b505af4158015613846573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061386a9190615792565b9250925092506000856001600160a01b03166370a08231306040518263ffffffff1660e01b815260040161389e9190615962565b60206040518083038186803b1580156138b657600080fd5b505afa1580156138ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138ee91906156eb565b90506138fb8683836142a3565b604051630b68372160e31b81526001600160a01b03831690635b41b9089061392e90879087908690600090600401615d1e565b600060405180830381600087803b15801561394857600080fd5b505af115801561395c573d6000803e3d6000fd5b5050600190980197506137759650505050505050565b6060610a58826000602b614152565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b1781529251825160009485949389169392918291908083835b602083106139fd5780518252601f1990920191602091820191016139de565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114613a5f576040519150601f19603f3d011682016040523d82523d6000602084013e613a64565b606091505b5091509150818015613a92575080511580613a925750808060200190516020811015613a8f57600080fd5b50515b611a80576040805162461bcd60e51b815260206004820152600260248201526114d560f21b604482015290519081900360640190fd5b6000808351855114613ad957600080fd5b600085516001600160401b0381118015613af257600080fd5b50604051908082528060200260200182016040528015613b2c57816020015b613b19614bb8565b815260200190600190039081613b115790505b509050600086516001600160401b0381118015613b4857600080fd5b50604051908082528060200260200182016040528015613b8257816020015b613b6f614bb8565b815260200190600190039081613b675790505b50905060005b8751811015613c9757600080613bb18a8481518110613ba357fe5b602002602001015189613cbe565b91509150613bbe826143ea565b858481518110613bca57fe5b60200260200101516000019060020b908160020b81525050613beb816143ea565b848481518110613bf757fe5b60200260200101516000019060020b908160020b81525050888381518110613c1b57fe5b6020026020010151858481518110613c2f57fe5b6020026020010151602001906001600160801b031690816001600160801b031681525050888381518110613c5f57fe5b6020026020010151848481518110613c7357fe5b6020908102919091018101516001600160801b039092169101525050600101613b88565b50613ca1826143fb565b60020b9350613caf816143fb565b60020b92505050935093915050565b600080600080613ccd866144b7565b90506000805b82811015613e76576000806000613ce98b612fcf565b9250925092506000613cfc8484846144c4565b905060008063ffffffff8d16613d2557613d1583614502565b600291820b9350900b9050613dba565b613d2f838e614719565b8160020b91505080925050826001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b158015613d7357600080fd5b505afa158015613d87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613dab919061562d565b50505060029290920b93505050505b60018903871415613de157846001600160a01b0316866001600160a01b0316109950613df0565b613dea8e613198565b9d508597505b6000871580613e435750866001600160a01b0316896001600160a01b031610613e2d57866001600160a01b0316866001600160a01b031610613e43565b856001600160a01b0316876001600160a01b0316105b90508015613e58579b82019b9a81019a613e63565b828d039c50818c039b505b505060019095019450613cd39350505050565b5082613e8b5760001985029450600019840293505b5050509250929050565b604080516001600160a01b0385811660248301528481166044830152606480830185905283518084039091018152608490920183526020820180516001600160e01b03166323b872dd60e01b178152925182516000948594938a169392918291908083835b60208310613f195780518252601f199092019160209182019101613efa565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114613f7b576040519150601f19603f3d011682016040523d82523d6000602084013e613f80565b606091505b5091509150818015613fae575080511580613fae5750808060200190516020811015613fab57600080fd5b50515b610ae1576040805162461bcd60e51b815260206004820152600360248201526229aa2360e91b604482015290519081900360640190fd5b6000600160ff1b8210613ff757600080fd5b5090565b60008182601401101561404a576040805162461bcd60e51b8152602060048201526012602482015271746f416464726573735f6f766572666c6f7760701b604482015290519081900360640190fd5b816014018351101561409b576040805162461bcd60e51b8152602060048201526015602482015274746f416464726573735f6f75744f66426f756e647360581b604482015290519081900360640190fd5b500160200151600160601b900490565b6000818260030110156140f9576040805162461bcd60e51b8152602060048201526011602482015270746f55696e7432345f6f766572666c6f7760781b604482015290519081900360640190fd5b8160030183511015614149576040805162461bcd60e51b8152602060048201526014602482015273746f55696e7432345f6f75744f66426f756e647360601b604482015290519081900360640190fd5b50016003015190565b60608182601f01101561419d576040805162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b604482015290519081900360640190fd5b8282840110156141e5576040805162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b604482015290519081900360640190fd5b81830184511015614231576040805162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b604482015290519081900360640190fd5b606082158015614250576040519150600082526020820160405261429a565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015614289578051835260209283019201614271565b5050858452601f01601f1916604052505b50949350505050565b604080516001600160a01b038481166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663095ea7b360e01b1781529251825160009485949389169392918291908083835b6020831061431f5780518252601f199092019160209182019101614300565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114614381576040519150601f19603f3d011682016040523d82523d6000602084013e614386565b606091505b50915091508180156143b45750805115806143b457508080602001905160208110156143b157600080fd5b50515b611a80576040805162461bcd60e51b8152602060048201526002602482015261534160f01b604482015290519081900360640190fd5b80600281900b8114610a5b57600080fd5b6000806000805b845181101561447e5784818151811061441757fe5b6020026020010151602001516001600160801b031685828151811061443857fe5b60200260200101516000015160020b028301925084818151811061445857fe5b6020026020010151602001516001600160801b0316820191508080600101915050614402565b5080828161448857fe5b0592506000821280156144a3575080828161449f57fe5b0715155b156144b057600019909201915b5050919050565b5160176013199091010490565b60006144fa7f00000000000000000000000000000000000000000000000000000000000000006144f5868686614a83565b614ad9565b949350505050565b600080600080846001600160a01b0316633850c7bd6040518163ffffffff1660e01b815260040160e06040518083038186803b15801561454157600080fd5b505afa158015614555573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614579919061562d565b50939750919550935050600161ffff8416119150614598905057600080fd5b600080866001600160a01b031663252c09d7856040518263ffffffff1660e01b81526004016145c79190615cc3565b60806040518083038186803b1580156145df57600080fd5b505afa1580156145f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146179190615819565b505091509150614625613739565b63ffffffff168263ffffffff161461463f57849550614710565b60008361ffff1660018561ffff168761ffff1601038161465b57fe5b06905060008060008a6001600160a01b031663252c09d7856040518263ffffffff1660e01b815260040161468f9190615cd2565b60806040518083038186803b1580156146a757600080fd5b505afa1580156146bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146df9190615819565b93505092509250806146f057600080fd5b82860363ffffffff811683870360060b8161470757fe5b059a5050505050505b50505050915091565b60008063ffffffff8316614759576040805162461bcd60e51b8152602060048201526002602482015261042560f41b604482015290519081900360640190fd5b604080516002808252606082018352600092602083019080368337019050509050838160008151811061478857fe5b602002602001019063ffffffff16908163ffffffff16815250506000816001815181106147b157fe5b63ffffffff90921660209283029190910182015260405163883bdbfd60e01b81526004810182815283516024830152835160009384936001600160a01b038b169363883bdbfd9388939192839260449091019185820191028083838b5b8381101561482657818101518382015260200161480e565b505050509050019250505060006040518083038186803b15801561484957600080fd5b505afa15801561485d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604090815281101561488657600080fd5b8101908080516040519392919084600160201b8211156148a557600080fd5b9083019060208201858111156148ba57600080fd5b82518660208202830111600160201b821117156148d657600080fd5b82525081516020918201928201910280838360005b838110156149035781810151838201526020016148eb565b5050505090500160405260200180516040519392919084600160201b82111561492b57600080fd5b90830190602082018581111561494057600080fd5b82518660208202830111600160201b8211171561495c57600080fd5b82525081516020918201928201910280838360005b83811015614989578181015183820152602001614971565b50505050905001604052505050915091506000826000815181106149a957fe5b6020026020010151836001815181106149be57fe5b60200260200101510390506000826000815181106149d857fe5b6020026020010151836001815181106149ed57fe5b60200260200101510390508763ffffffff168260060b81614a0a57fe5b05965060008260060b128015614a3457508763ffffffff168260060b81614a2d57fe5b0760060b15155b15614a4157600019909601955b63ffffffff88166001600160a01b0302600160201b600160c01b03602083901b166001600160c01b03821681614a7357fe5b0496505050505050509250929050565b614a8b614bcf565b826001600160a01b0316846001600160a01b03161115614aa9579192915b50604080516060810182526001600160a01b03948516815292909316602083015262ffffff169181019190915290565b600081602001516001600160a01b031682600001516001600160a01b031610614b0157600080fd5b50805160208083015160409384015184516001600160a01b0394851681850152939091168385015262ffffff166060808401919091528351808403820181526080840185528051908301206001600160f81b031960a085015294901b6001600160601b03191660a183015260b58201939093527f6ce8eb472fa82df5469c6ab6d485f17c3ad13c8cd7af59b3d4a8026c5ce0f7e260d5808301919091528251808303909101815260f5909101909152805191012090565b604080518082019091526000808252602082015290565b604080516060810182526000808252602082018190529181019190915290565b8035610a5b81615e0a565b60008083601f840112614c0b578182fd5b5081356001600160401b03811115614c21578182fd5b6020830191508360208083028501011115614c3b57600080fd5b9250929050565b600082601f830112614c52578081fd5b81356020614c67614c6283615da0565b615d7d565b8281528181019085830183850287018401881015614c83578586fd5b855b85811015614cb55781356001600160801b0381168114614ca3578788fd5b84529284019290840190600101614c85565b5090979650505050505050565b80518015158114610a5b57600080fd5b600082601f830112614ce2578081fd5b8135614cf0614c6282615dbd565b818152846020838601011115614d04578283fd5b816020850160208301379081016020019190915292915050565b80516001600160701b0381168114610a5b57600080fd5b805161ffff81168114610a5b57600080fd5b803562ffffff81168114610a5b57600080fd5b8035610a5b81615e2e565b600060208284031215614d76578081fd5b8135614d8181615e0a565b9392505050565b600060208284031215614d99578081fd5b8151614d8181615e0a565b60008060408385031215614db6578081fd5b8235614dc181615e0a565b91506020830135614dd181615e0a565b809150509250929050565b60008060408385031215614dee578182fd5b8251614df981615e0a565b6020840151909250614dd181615e0a565b60008060408385031215614e1c578182fd5b8235614e2781615e0a565b946020939093013593505050565b600080600060608486031215614e49578081fd5b8335614e5481615e0a565b9250602084013591506040840135614e6b81615e0a565b809150509250925092565b600080600080600060a08688031215614e8d578283fd5b8535614e9881615e0a565b9450602086013593506040860135614eaf81615e0a565b9250606086013591506080860135614ec681615e0a565b809150509295509295909350565b60008060008060808587031215614ee9578182fd5b8435614ef481615e0a565b935060208501359250604085013591506060850135614f1281615e0a565b939692955090935050565b60008060008060008060c08789031215614f35578384fd5b8635614f4081615e0a565b95506020870135945060408701359350606087013560ff81168114614f63578182fd5b9598949750929560808101359460a0909101359350915050565b600080600080600080600060a0888a031215614f97578485fd5b87356001600160401b0380821115614fad578687fd5b614fb98b838c01614bfa565b909950975060208a0135915080821115614fd1578687fd5b50614fde8a828b01614bfa565b9096509450506040880135925060608801359150608088013561500081615e0a565b8091505092959891949750929550565b60008060208385031215615022578182fd5b82356001600160401b03811115615037578283fd5b61504385828601614bfa565b90969095509350505050565b60008060008060808587031215615064578182fd5b84356001600160401b038082111561507a578384fd5b818701915087601f83011261508d578384fd5b8135602061509d614c6283615da0565b82815281810190858301885b858110156150d2576150c08e8684358b0101614cd2565b845292840192908401906001016150a9565b509099505050880135925050808211156150ea578384fd5b506150f787828801614c42565b93505061510660408601614d47565b915061511460608601614d5a565b905092959194509250565b60006020808385031215615131578182fd5b82516001600160401b03811115615146578283fd5b8301601f81018513615156578283fd5b8051615164614c6282615da0565b8181528381019083850185840285018601891015615180578687fd5b8694505b838510156151a2578051835260019490940193918501918501615184565b50979650505050505050565b6000602082840312156151bf578081fd5b614d8182614cc2565b6000806000604084860312156151dc578081fd5b8335925060208401356001600160401b038111156151f8578182fd5b61520486828701614bfa565b9497909650939450505050565b600060208284031215615222578081fd5b81356001600160401b03811115615237578182fd5b6144fa84828501614cd2565b600080600060608486031215615257578081fd5b83356001600160401b0381111561526c578182fd5b61527886828701614cd2565b93505061528760208501614d47565b91506040840135614e6b81615e2e565b6000602082840312156152a8578081fd5b8135614d8181615e1f565b600080604083850312156152c5578182fd5b505080516020909101519092909150565b600080600080606085870312156152eb578182fd5b843593506020850135925060408501356001600160401b038082111561530f578384fd5b818701915087601f830112615322578384fd5b813581811115615330578485fd5b886020828501011115615341578485fd5b95989497505060200194505050565b600060208284031215615361578081fd5b81516001600160401b03811115615376578182fd5b8201601f81018413615386578182fd5b8051615394614c6282615dbd565b8181528560208385010111156153a8578384fd5b610b3a826020830160208601615dde565b6000602082840312156153ca578081fd5b81356001600160401b03808211156153e0578283fd5b90830190608082860312156153f3578283fd5b60405160808101818110838211171561540857fe5b604052823582811115615419578485fd5b61542587828601614cd2565b8252506020830135915061543882615e0a565b816020820152604083013560408201526060830135606082015280935050505092915050565b600060e0828403121561546f578081fd5b60405160e081016001600160401b038111828210171561548b57fe5b60405261549783614bef565b81526154a560208401614bef565b60208201526154b660408401614d47565b60408201526154c760608401614bef565b60608201526080830135608082015260a083013560a08201526154ec60c08401614bef565b60c08201529392505050565b600060208284031215615509578081fd5b81356001600160401b0381111561551e578182fd5b820160808185031215614d81578182fd5b600060e08284031215611d31578081fd5b600060a08284031215611d31578081fd5b60006101008284031215611d31578081fd5b600060208284031215615574578081fd5b81356001600160401b038082111561558a578283fd5b908301906040828603121561559d578283fd5b6040516040810181811083821117156155b257fe5b6040528235828111156155c3578485fd5b6155cf87828601614cd2565b825250602083013592506155e283615e0a565b6020810192909252509392505050565b600080600060608486031215615606578081fd5b61560f84614d1e565b925061561d60208501614d1e565b91506040840151614e6b81615e2e565b600080600080600080600060e0888a031215615647578081fd5b875161565281615e0a565b602089015190975061566381615e1f565b955061567160408901614d35565b945061567f60608901614d35565b935061568d60808901614d35565b925060a088015161569d81615e2e565b91506156ab60c08901614cc2565b905092959891949750929550565b6000602082840312156156ca578081fd5b614d8182614d47565b6000602082840312156156e4578081fd5b5035919050565b6000602082840312156156fc578081fd5b5051919050565b60008060408385031215615715578182fd5b823591506020830135614dd181615e0a565b6000806000806080858703121561573c578182fd5b84359350602085013561574e81615e0a565b9250604085013591506060850135614f1281615e0a565b600080600060608486031215615779578081fd5b83359250602084013591506040840135614e6b81615e0a565b6000806000606084860312156157a6578081fd5b83519250602084015191506040840151614e6b81615e0a565b6000806000806000608086880312156157d6578283fd5b853594506020860135935060408601356001600160401b038111156157f9578384fd5b61580588828901614bfa565b9094509250506060860135614ec681615e0a565b6000806000806080858703121561582e578182fd5b845161583981615e2e565b8094505060208501518060060b8114615850578283fd5b604086015190935061586181615e0a565b915061511460608601614cc2565b6001600160a01b03169052565b60008284526020808501945082825b858110156158b957813561589e81615e0a565b6001600160a01b03168752958201959082019060010161588b565b509495945050505050565b600081518084526158dc816020860160208601615dde565b601f01601f19169290920160200192915050565b60020b9052565b62ffffff169052565b606093841b6001600160601b0319908116825260e89390931b6001600160e81b0319166014820152921b166017820152602b0190565b6000828483379101908152919050565b60008251615958818460208701615dde565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b0393841681529183166020830152909116604082015260600190565b6001600160a01b03948516815292841660208401529216604082015262ffffff909116606082015260800190565b6001600160a01b039485168152928416602084015292166040820152606081019190915260800190565b6001600160a01b0388811682528716602082015260a060408201819052600090615a38908301878961587c565b82810360608401528481526001600160fb1b03851115615a56578182fd5b602085028087602084013701602001908152608091909101919091529695505050505050565b6001600160a01b0386811682528515156020830152604082018590528316606082015260a060808201819052600090615ab7908301846158c4565b979650505050505050565b6001600160a01b03929092168252602082015260400190565b600060018060a01b038616825284602083015260606040830152615b0360608301848661587c565b9695505050505050565b6000602080830181845280855180835260408601915060408482028701019250838701855b82811015615b6057603f19888603018452615b4e8583516158c4565b94509285019290850190600101615b32565b5092979650505050505050565b600060208252614d8160208301846158c4565b6020810160058310615b8e57fe5b91905290565b600060c082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015292915050565b600061016082019050615bec82845161586f565b6020830151615bfe602084018261586f565b506040830151615c1160408401826158f7565b506060830151615c2460608401826158f0565b506080830151615c3760808401826158f0565b5060a083015160a083015260c083015160c083015260e083015160e083015261010080840151818401525061012080840151615c758285018261586f565b505061014092830151919092015290565b600060208252825160406020840152615ca260608401826158c4565b602094909401516001600160a01b0316604093909301929092525090919050565b61ffff91909116815260200190565b90815260200190565b600085825284602083015260018060a01b038416604083015260806060830152615b0360808301846158c4565b9283526020830191909152604082015260600190565b93845260208401929092526040830152606082015260800190565b6000808335601e19843603018112615d4f578283fd5b8301803591506001600160401b03821115615d68578283fd5b602001915036819003821315614c3b57600080fd5b6040518181016001600160401b0381118282101715615d9857fe5b604052919050565b60006001600160401b03821115615db357fe5b5060209081020190565b60006001600160401b03821115615dd057fe5b50601f01601f191660200190565b60005b83811015615df9578181015183820152602001615de1565b83811115610ddc5750506000910152565b6001600160a01b038116811461175157600080fd5b8060020b811461175157600080fd5b63ffffffff8116811461175157600080fdfe5265656e7472616e637947756172643a207265656e7472616e742063616c6c004f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65728be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a26469706673582212205895e27cbf1d81d8130f314234ccbd2ca0c396cf4151f6bdcd109cfff720b9b464736f6c63430007060033",
}

// SectaV3RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use SectaV3RouterMetaData.ABI instead.
var SectaV3RouterABI = SectaV3RouterMetaData.ABI

// SectaV3RouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SectaV3RouterMetaData.Bin instead.
var SectaV3RouterBin = SectaV3RouterMetaData.Bin

// DeploySectaV3Router deploys a new Ethereum contract, binding an instance of SectaV3Router to it.
func DeploySectaV3Router(auth *bind.TransactOpts, backend bind.ContractBackend, _factoryV2 common.Address, _deployer common.Address, _factoryV3 common.Address, _positionManager common.Address, _stableFactory common.Address, _stableInfo common.Address, _WETH9 common.Address) (common.Address, *types.Transaction, *SectaV3Router, error) {
	parsed, err := SectaV3RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SectaV3RouterBin), backend, _factoryV2, _deployer, _factoryV3, _positionManager, _stableFactory, _stableInfo, _WETH9)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SectaV3Router{SectaV3RouterCaller: SectaV3RouterCaller{contract: contract}, SectaV3RouterTransactor: SectaV3RouterTransactor{contract: contract}, SectaV3RouterFilterer: SectaV3RouterFilterer{contract: contract}}, nil
}

// SectaV3Router is an auto generated Go binding around an Ethereum contract.
type SectaV3Router struct {
	SectaV3RouterCaller     // Read-only binding to the contract
	SectaV3RouterTransactor // Write-only binding to the contract
	SectaV3RouterFilterer   // Log filterer for contract events
}

// SectaV3RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SectaV3RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SectaV3RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SectaV3RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SectaV3RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SectaV3RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SectaV3RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SectaV3RouterSession struct {
	Contract     *SectaV3Router    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SectaV3RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SectaV3RouterCallerSession struct {
	Contract *SectaV3RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SectaV3RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SectaV3RouterTransactorSession struct {
	Contract     *SectaV3RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SectaV3RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SectaV3RouterRaw struct {
	Contract *SectaV3Router // Generic contract binding to access the raw methods on
}

// SectaV3RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SectaV3RouterCallerRaw struct {
	Contract *SectaV3RouterCaller // Generic read-only contract binding to access the raw methods on
}

// SectaV3RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SectaV3RouterTransactorRaw struct {
	Contract *SectaV3RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSectaV3Router creates a new instance of SectaV3Router, bound to a specific deployed contract.
func NewSectaV3Router(address common.Address, backend bind.ContractBackend) (*SectaV3Router, error) {
	contract, err := bindSectaV3Router(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SectaV3Router{SectaV3RouterCaller: SectaV3RouterCaller{contract: contract}, SectaV3RouterTransactor: SectaV3RouterTransactor{contract: contract}, SectaV3RouterFilterer: SectaV3RouterFilterer{contract: contract}}, nil
}

// NewSectaV3RouterCaller creates a new read-only instance of SectaV3Router, bound to a specific deployed contract.
func NewSectaV3RouterCaller(address common.Address, caller bind.ContractCaller) (*SectaV3RouterCaller, error) {
	contract, err := bindSectaV3Router(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SectaV3RouterCaller{contract: contract}, nil
}

// NewSectaV3RouterTransactor creates a new write-only instance of SectaV3Router, bound to a specific deployed contract.
func NewSectaV3RouterTransactor(address common.Address, transactor bind.ContractTransactor) (*SectaV3RouterTransactor, error) {
	contract, err := bindSectaV3Router(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SectaV3RouterTransactor{contract: contract}, nil
}

// NewSectaV3RouterFilterer creates a new log filterer instance of SectaV3Router, bound to a specific deployed contract.
func NewSectaV3RouterFilterer(address common.Address, filterer bind.ContractFilterer) (*SectaV3RouterFilterer, error) {
	contract, err := bindSectaV3Router(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SectaV3RouterFilterer{contract: contract}, nil
}

// bindSectaV3Router binds a generic wrapper to an already deployed contract.
func bindSectaV3Router(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SectaV3RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SectaV3Router *SectaV3RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SectaV3Router.Contract.SectaV3RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SectaV3Router *SectaV3RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SectaV3RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SectaV3Router *SectaV3RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SectaV3RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SectaV3Router *SectaV3RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SectaV3Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SectaV3Router *SectaV3RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SectaV3Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SectaV3Router *SectaV3RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SectaV3Router.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_SectaV3Router *SectaV3RouterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SectaV3Router.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_SectaV3Router *SectaV3RouterSession) WETH9() (common.Address, error) {
	return _SectaV3Router.Contract.WETH9(&_SectaV3Router.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_SectaV3Router *SectaV3RouterCallerSession) WETH9() (common.Address, error) {
	return _SectaV3Router.Contract.WETH9(&_SectaV3Router.CallOpts)
}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SectaV3Router *SectaV3RouterCaller) CheckOracleSlippage(opts *bind.CallOpts, paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	var out []interface{}
	err := _SectaV3Router.contract.Call(opts, &out, "checkOracleSlippage", paths, amounts, maximumTickDivergence, secondsAgo)

	if err != nil {
		return err
	}

	return err

}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SectaV3Router *SectaV3RouterSession) CheckOracleSlippage(paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _SectaV3Router.Contract.CheckOracleSlippage(&_SectaV3Router.CallOpts, paths, amounts, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage is a free data retrieval call binding the contract method 0xefdeed8e.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SectaV3Router *SectaV3RouterCallerSession) CheckOracleSlippage(paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _SectaV3Router.Contract.CheckOracleSlippage(&_SectaV3Router.CallOpts, paths, amounts, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SectaV3Router *SectaV3RouterCaller) CheckOracleSlippage0(opts *bind.CallOpts, path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	var out []interface{}
	err := _SectaV3Router.contract.Call(opts, &out, "checkOracleSlippage0", path, maximumTickDivergence, secondsAgo)

	if err != nil {
		return err
	}

	return err

}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SectaV3Router *SectaV3RouterSession) CheckOracleSlippage0(path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _SectaV3Router.Contract.CheckOracleSlippage0(&_SectaV3Router.CallOpts, path, maximumTickDivergence, secondsAgo)
}

// CheckOracleSlippage0 is a free data retrieval call binding the contract method 0xf25801a7.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (_SectaV3Router *SectaV3RouterCallerSession) CheckOracleSlippage0(path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) error {
	return _SectaV3Router.Contract.CheckOracleSlippage0(&_SectaV3Router.CallOpts, path, maximumTickDivergence, secondsAgo)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_SectaV3Router *SectaV3RouterCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SectaV3Router.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_SectaV3Router *SectaV3RouterSession) Deployer() (common.Address, error) {
	return _SectaV3Router.Contract.Deployer(&_SectaV3Router.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_SectaV3Router *SectaV3RouterCallerSession) Deployer() (common.Address, error) {
	return _SectaV3Router.Contract.Deployer(&_SectaV3Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SectaV3Router *SectaV3RouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SectaV3Router.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SectaV3Router *SectaV3RouterSession) Factory() (common.Address, error) {
	return _SectaV3Router.Contract.Factory(&_SectaV3Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_SectaV3Router *SectaV3RouterCallerSession) Factory() (common.Address, error) {
	return _SectaV3Router.Contract.Factory(&_SectaV3Router.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_SectaV3Router *SectaV3RouterCaller) FactoryV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SectaV3Router.contract.Call(opts, &out, "factoryV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_SectaV3Router *SectaV3RouterSession) FactoryV2() (common.Address, error) {
	return _SectaV3Router.Contract.FactoryV2(&_SectaV3Router.CallOpts)
}

// FactoryV2 is a free data retrieval call binding the contract method 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (_SectaV3Router *SectaV3RouterCallerSession) FactoryV2() (common.Address, error) {
	return _SectaV3Router.Contract.FactoryV2(&_SectaV3Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SectaV3Router *SectaV3RouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SectaV3Router.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SectaV3Router *SectaV3RouterSession) Owner() (common.Address, error) {
	return _SectaV3Router.Contract.Owner(&_SectaV3Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SectaV3Router *SectaV3RouterCallerSession) Owner() (common.Address, error) {
	return _SectaV3Router.Contract.Owner(&_SectaV3Router.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_SectaV3Router *SectaV3RouterCaller) PositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SectaV3Router.contract.Call(opts, &out, "positionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_SectaV3Router *SectaV3RouterSession) PositionManager() (common.Address, error) {
	return _SectaV3Router.Contract.PositionManager(&_SectaV3Router.CallOpts)
}

// PositionManager is a free data retrieval call binding the contract method 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (_SectaV3Router *SectaV3RouterCallerSession) PositionManager() (common.Address, error) {
	return _SectaV3Router.Contract.PositionManager(&_SectaV3Router.CallOpts)
}

// StableSwapFactory is a free data retrieval call binding the contract method 0x57c79961.
//
// Solidity: function stableSwapFactory() view returns(address)
func (_SectaV3Router *SectaV3RouterCaller) StableSwapFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SectaV3Router.contract.Call(opts, &out, "stableSwapFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StableSwapFactory is a free data retrieval call binding the contract method 0x57c79961.
//
// Solidity: function stableSwapFactory() view returns(address)
func (_SectaV3Router *SectaV3RouterSession) StableSwapFactory() (common.Address, error) {
	return _SectaV3Router.Contract.StableSwapFactory(&_SectaV3Router.CallOpts)
}

// StableSwapFactory is a free data retrieval call binding the contract method 0x57c79961.
//
// Solidity: function stableSwapFactory() view returns(address)
func (_SectaV3Router *SectaV3RouterCallerSession) StableSwapFactory() (common.Address, error) {
	return _SectaV3Router.Contract.StableSwapFactory(&_SectaV3Router.CallOpts)
}

// StableSwapInfo is a free data retrieval call binding the contract method 0xb85aa7af.
//
// Solidity: function stableSwapInfo() view returns(address)
func (_SectaV3Router *SectaV3RouterCaller) StableSwapInfo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SectaV3Router.contract.Call(opts, &out, "stableSwapInfo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StableSwapInfo is a free data retrieval call binding the contract method 0xb85aa7af.
//
// Solidity: function stableSwapInfo() view returns(address)
func (_SectaV3Router *SectaV3RouterSession) StableSwapInfo() (common.Address, error) {
	return _SectaV3Router.Contract.StableSwapInfo(&_SectaV3Router.CallOpts)
}

// StableSwapInfo is a free data retrieval call binding the contract method 0xb85aa7af.
//
// Solidity: function stableSwapInfo() view returns(address)
func (_SectaV3Router *SectaV3RouterCallerSession) StableSwapInfo() (common.Address, error) {
	return _SectaV3Router.Contract.StableSwapInfo(&_SectaV3Router.CallOpts)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) ApproveMax(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "approveMax", token)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_SectaV3Router *SectaV3RouterSession) ApproveMax(token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ApproveMax(&_SectaV3Router.TransactOpts, token)
}

// ApproveMax is a paid mutator transaction binding the contract method 0x571ac8b0.
//
// Solidity: function approveMax(address token) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) ApproveMax(token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ApproveMax(&_SectaV3Router.TransactOpts, token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) ApproveMaxMinusOne(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "approveMaxMinusOne", token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_SectaV3Router *SectaV3RouterSession) ApproveMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ApproveMaxMinusOne(&_SectaV3Router.TransactOpts, token)
}

// ApproveMaxMinusOne is a paid mutator transaction binding the contract method 0xcab372ce.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) ApproveMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ApproveMaxMinusOne(&_SectaV3Router.TransactOpts, token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) ApproveZeroThenMax(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "approveZeroThenMax", token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_SectaV3Router *SectaV3RouterSession) ApproveZeroThenMax(token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ApproveZeroThenMax(&_SectaV3Router.TransactOpts, token)
}

// ApproveZeroThenMax is a paid mutator transaction binding the contract method 0x639d71a9.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) ApproveZeroThenMax(token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ApproveZeroThenMax(&_SectaV3Router.TransactOpts, token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) ApproveZeroThenMaxMinusOne(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "approveZeroThenMaxMinusOne", token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_SectaV3Router *SectaV3RouterSession) ApproveZeroThenMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ApproveZeroThenMaxMinusOne(&_SectaV3Router.TransactOpts, token)
}

// ApproveZeroThenMaxMinusOne is a paid mutator transaction binding the contract method 0xab3fdd50.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) ApproveZeroThenMaxMinusOne(token common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ApproveZeroThenMaxMinusOne(&_SectaV3Router.TransactOpts, token)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_SectaV3Router *SectaV3RouterTransactor) CallPositionManager(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "callPositionManager", data)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_SectaV3Router *SectaV3RouterSession) CallPositionManager(data []byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.CallPositionManager(&_SectaV3Router.TransactOpts, data)
}

// CallPositionManager is a paid mutator transaction binding the contract method 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (_SectaV3Router *SectaV3RouterTransactorSession) CallPositionManager(data []byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.CallPositionManager(&_SectaV3Router.TransactOpts, data)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterTransactor) ExactInput(opts *bind.TransactOpts, params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterSession) ExactInput(params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactInput(&_SectaV3Router.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterTransactorSession) ExactInput(params IV3SwapRouterExactInputParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactInput(&_SectaV3Router.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterSession) ExactInputSingle(params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactInputSingle(&_SectaV3Router.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterTransactorSession) ExactInputSingle(params IV3SwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactInputSingle(&_SectaV3Router.TransactOpts, params)
}

// ExactInputStableSwap is a paid mutator transaction binding the contract method 0xb4554231.
//
// Solidity: function exactInputStableSwap(address[] path, uint256[] flag, uint256 amountIn, uint256 amountOutMin, address to) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterTransactor) ExactInputStableSwap(opts *bind.TransactOpts, path []common.Address, flag []*big.Int, amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "exactInputStableSwap", path, flag, amountIn, amountOutMin, to)
}

// ExactInputStableSwap is a paid mutator transaction binding the contract method 0xb4554231.
//
// Solidity: function exactInputStableSwap(address[] path, uint256[] flag, uint256 amountIn, uint256 amountOutMin, address to) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterSession) ExactInputStableSwap(path []common.Address, flag []*big.Int, amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactInputStableSwap(&_SectaV3Router.TransactOpts, path, flag, amountIn, amountOutMin, to)
}

// ExactInputStableSwap is a paid mutator transaction binding the contract method 0xb4554231.
//
// Solidity: function exactInputStableSwap(address[] path, uint256[] flag, uint256 amountIn, uint256 amountOutMin, address to) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterTransactorSession) ExactInputStableSwap(path []common.Address, flag []*big.Int, amountIn *big.Int, amountOutMin *big.Int, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactInputStableSwap(&_SectaV3Router.TransactOpts, path, flag, amountIn, amountOutMin, to)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterTransactor) ExactOutput(opts *bind.TransactOpts, params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterSession) ExactOutput(params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactOutput(&_SectaV3Router.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterTransactorSession) ExactOutput(params IV3SwapRouterExactOutputParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactOutput(&_SectaV3Router.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterSession) ExactOutputSingle(params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactOutputSingle(&_SectaV3Router.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterTransactorSession) ExactOutputSingle(params IV3SwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactOutputSingle(&_SectaV3Router.TransactOpts, params)
}

// ExactOutputStableSwap is a paid mutator transaction binding the contract method 0xb4c4e555.
//
// Solidity: function exactOutputStableSwap(address[] path, uint256[] flag, uint256 amountOut, uint256 amountInMax, address to) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterTransactor) ExactOutputStableSwap(opts *bind.TransactOpts, path []common.Address, flag []*big.Int, amountOut *big.Int, amountInMax *big.Int, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "exactOutputStableSwap", path, flag, amountOut, amountInMax, to)
}

// ExactOutputStableSwap is a paid mutator transaction binding the contract method 0xb4c4e555.
//
// Solidity: function exactOutputStableSwap(address[] path, uint256[] flag, uint256 amountOut, uint256 amountInMax, address to) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterSession) ExactOutputStableSwap(path []common.Address, flag []*big.Int, amountOut *big.Int, amountInMax *big.Int, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactOutputStableSwap(&_SectaV3Router.TransactOpts, path, flag, amountOut, amountInMax, to)
}

// ExactOutputStableSwap is a paid mutator transaction binding the contract method 0xb4c4e555.
//
// Solidity: function exactOutputStableSwap(address[] path, uint256[] flag, uint256 amountOut, uint256 amountInMax, address to) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterTransactorSession) ExactOutputStableSwap(path []common.Address, flag []*big.Int, amountOut *big.Int, amountInMax *big.Int, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.ExactOutputStableSwap(&_SectaV3Router.TransactOpts, path, flag, amountOut, amountInMax, to)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_SectaV3Router *SectaV3RouterTransactor) GetApprovalType(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "getApprovalType", token, amount)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_SectaV3Router *SectaV3RouterSession) GetApprovalType(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.Contract.GetApprovalType(&_SectaV3Router.TransactOpts, token, amount)
}

// GetApprovalType is a paid mutator transaction binding the contract method 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (_SectaV3Router *SectaV3RouterTransactorSession) GetApprovalType(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.Contract.GetApprovalType(&_SectaV3Router.TransactOpts, token, amount)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_SectaV3Router *SectaV3RouterTransactor) IncreaseLiquidity(opts *bind.TransactOpts, params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "increaseLiquidity", params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_SectaV3Router *SectaV3RouterSession) IncreaseLiquidity(params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.IncreaseLiquidity(&_SectaV3Router.TransactOpts, params)
}

// IncreaseLiquidity is a paid mutator transaction binding the contract method 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (_SectaV3Router *SectaV3RouterTransactorSession) IncreaseLiquidity(params IApproveAndCallIncreaseLiquidityParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.IncreaseLiquidity(&_SectaV3Router.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_SectaV3Router *SectaV3RouterTransactor) Mint(opts *bind.TransactOpts, params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "mint", params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_SectaV3Router *SectaV3RouterSession) Mint(params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.Mint(&_SectaV3Router.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (_SectaV3Router *SectaV3RouterTransactorSession) Mint(params IApproveAndCallMintParams) (*types.Transaction, error) {
	return _SectaV3Router.Contract.Mint(&_SectaV3Router.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_SectaV3Router *SectaV3RouterTransactor) Multicall(opts *bind.TransactOpts, previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "multicall", previousBlockhash, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_SectaV3Router *SectaV3RouterSession) Multicall(previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.Multicall(&_SectaV3Router.TransactOpts, previousBlockhash, data)
}

// Multicall is a paid mutator transaction binding the contract method 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (_SectaV3Router *SectaV3RouterTransactorSession) Multicall(previousBlockhash [32]byte, data [][]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.Multicall(&_SectaV3Router.TransactOpts, previousBlockhash, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_SectaV3Router *SectaV3RouterTransactor) Multicall0(opts *bind.TransactOpts, deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "multicall0", deadline, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_SectaV3Router *SectaV3RouterSession) Multicall0(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.Multicall0(&_SectaV3Router.TransactOpts, deadline, data)
}

// Multicall0 is a paid mutator transaction binding the contract method 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (_SectaV3Router *SectaV3RouterTransactorSession) Multicall0(deadline *big.Int, data [][]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.Multicall0(&_SectaV3Router.TransactOpts, deadline, data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SectaV3Router *SectaV3RouterTransactor) Multicall1(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "multicall1", data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SectaV3Router *SectaV3RouterSession) Multicall1(data [][]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.Multicall1(&_SectaV3Router.TransactOpts, data)
}

// Multicall1 is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_SectaV3Router *SectaV3RouterTransactorSession) Multicall1(data [][]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.Multicall1(&_SectaV3Router.TransactOpts, data)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SectaV3Router *SectaV3RouterTransactor) PancakeV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "pancakeV3SwapCallback", amount0Delta, amount1Delta, _data)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SectaV3Router *SectaV3RouterSession) PancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.PancakeV3SwapCallback(&_SectaV3Router.TransactOpts, amount0Delta, amount1Delta, _data)
}

// PancakeV3SwapCallback is a paid mutator transaction binding the contract method 0x23a69e75.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) PancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.PancakeV3SwapCallback(&_SectaV3Router.TransactOpts, amount0Delta, amount1Delta, _data)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) Pull(opts *bind.TransactOpts, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "pull", token, value)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_SectaV3Router *SectaV3RouterSession) Pull(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.Contract.Pull(&_SectaV3Router.TransactOpts, token, value)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) Pull(token common.Address, value *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.Contract.Pull(&_SectaV3Router.TransactOpts, token, value)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SectaV3Router *SectaV3RouterSession) RefundETH() (*types.Transaction, error) {
	return _SectaV3Router.Contract.RefundETH(&_SectaV3Router.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) RefundETH() (*types.Transaction, error) {
	return _SectaV3Router.Contract.RefundETH(&_SectaV3Router.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SectaV3Router *SectaV3RouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SectaV3Router *SectaV3RouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _SectaV3Router.Contract.RenounceOwnership(&_SectaV3Router.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SectaV3Router.Contract.RenounceOwnership(&_SectaV3Router.TransactOpts)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SelfPermit(&_SectaV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SelfPermit(&_SectaV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SelfPermitAllowed(&_SectaV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SelfPermitAllowed(&_SectaV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SelfPermitAllowedIfNecessary(&_SectaV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SelfPermitAllowedIfNecessary(&_SectaV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SelfPermitIfNecessary(&_SectaV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SelfPermitIfNecessary(&_SectaV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x24dec034.
//
// Solidity: function setStableSwap(address _factory, address _info) returns()
func (_SectaV3Router *SectaV3RouterTransactor) SetStableSwap(opts *bind.TransactOpts, _factory common.Address, _info common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "setStableSwap", _factory, _info)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x24dec034.
//
// Solidity: function setStableSwap(address _factory, address _info) returns()
func (_SectaV3Router *SectaV3RouterSession) SetStableSwap(_factory common.Address, _info common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SetStableSwap(&_SectaV3Router.TransactOpts, _factory, _info)
}

// SetStableSwap is a paid mutator transaction binding the contract method 0x24dec034.
//
// Solidity: function setStableSwap(address _factory, address _info) returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) SetStableSwap(_factory common.Address, _info common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SetStableSwap(&_SectaV3Router.TransactOpts, _factory, _info)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SwapExactTokensForTokens(&_SectaV3Router.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (_SectaV3Router *SectaV3RouterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SwapExactTokensForTokens(&_SectaV3Router.TransactOpts, amountIn, amountOutMin, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SwapTokensForExactTokens(&_SectaV3Router.TransactOpts, amountOut, amountInMax, path, to)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (_SectaV3Router *SectaV3RouterTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SwapTokensForExactTokens(&_SectaV3Router.TransactOpts, amountOut, amountInMax, path, to)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SectaV3Router *SectaV3RouterSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SweepToken(&_SectaV3Router.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SweepToken(&_SectaV3Router.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) SweepToken0(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "sweepToken0", token, amountMinimum)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_SectaV3Router *SectaV3RouterSession) SweepToken0(token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SweepToken0(&_SectaV3Router.TransactOpts, token, amountMinimum)
}

// SweepToken0 is a paid mutator transaction binding the contract method 0xe90a182f.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) SweepToken0(token common.Address, amountMinimum *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SweepToken0(&_SectaV3Router.TransactOpts, token, amountMinimum)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SweepTokenWithFee(&_SectaV3Router.TransactOpts, token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0x3068c554.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SweepTokenWithFee(&_SectaV3Router.TransactOpts, token, amountMinimum, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) SweepTokenWithFee0(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "sweepTokenWithFee0", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterSession) SweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SweepTokenWithFee0(&_SectaV3Router.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee0 is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) SweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.SweepTokenWithFee0(&_SectaV3Router.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SectaV3Router *SectaV3RouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SectaV3Router *SectaV3RouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.TransferOwnership(&_SectaV3Router.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.TransferOwnership(&_SectaV3Router.TransactOpts, newOwner)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_SectaV3Router *SectaV3RouterSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.UnwrapWETH9(&_SectaV3Router.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.UnwrapWETH9(&_SectaV3Router.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) UnwrapWETH9WithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.UnwrapWETH9WithFee(&_SectaV3Router.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.UnwrapWETH9WithFee(&_SectaV3Router.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) UnwrapWETH9WithFee0(opts *bind.TransactOpts, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "unwrapWETH9WithFee0", amountMinimum, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterSession) UnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.UnwrapWETH9WithFee0(&_SectaV3Router.TransactOpts, amountMinimum, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee0 is a paid mutator transaction binding the contract method 0xd4ef38de.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) UnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _SectaV3Router.Contract.UnwrapWETH9WithFee0(&_SectaV3Router.TransactOpts, amountMinimum, feeBips, feeRecipient)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) WrapETH(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.contract.Transact(opts, "wrapETH", value)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_SectaV3Router *SectaV3RouterSession) WrapETH(value *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.Contract.WrapETH(&_SectaV3Router.TransactOpts, value)
}

// WrapETH is a paid mutator transaction binding the contract method 0x1c58db4f.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) WrapETH(value *big.Int) (*types.Transaction, error) {
	return _SectaV3Router.Contract.WrapETH(&_SectaV3Router.TransactOpts, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SectaV3Router *SectaV3RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SectaV3Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SectaV3Router *SectaV3RouterSession) Receive() (*types.Transaction, error) {
	return _SectaV3Router.Contract.Receive(&_SectaV3Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SectaV3Router *SectaV3RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _SectaV3Router.Contract.Receive(&_SectaV3Router.TransactOpts)
}

// SectaV3RouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SectaV3Router contract.
type SectaV3RouterOwnershipTransferredIterator struct {
	Event *SectaV3RouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SectaV3RouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SectaV3RouterOwnershipTransferred)
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
		it.Event = new(SectaV3RouterOwnershipTransferred)
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
func (it *SectaV3RouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SectaV3RouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SectaV3RouterOwnershipTransferred represents a OwnershipTransferred event raised by the SectaV3Router contract.
type SectaV3RouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SectaV3Router *SectaV3RouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SectaV3RouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SectaV3Router.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SectaV3RouterOwnershipTransferredIterator{contract: _SectaV3Router.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SectaV3Router *SectaV3RouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SectaV3RouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SectaV3Router.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SectaV3RouterOwnershipTransferred)
				if err := _SectaV3Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SectaV3Router *SectaV3RouterFilterer) ParseOwnershipTransferred(log types.Log) (*SectaV3RouterOwnershipTransferred, error) {
	event := new(SectaV3RouterOwnershipTransferred)
	if err := _SectaV3Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SectaV3RouterSetStableSwapIterator is returned from FilterSetStableSwap and is used to iterate over the raw logs and unpacked data for SetStableSwap events raised by the SectaV3Router contract.
type SectaV3RouterSetStableSwapIterator struct {
	Event *SectaV3RouterSetStableSwap // Event containing the contract specifics and raw log

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
func (it *SectaV3RouterSetStableSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SectaV3RouterSetStableSwap)
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
		it.Event = new(SectaV3RouterSetStableSwap)
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
func (it *SectaV3RouterSetStableSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SectaV3RouterSetStableSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SectaV3RouterSetStableSwap represents a SetStableSwap event raised by the SectaV3Router contract.
type SectaV3RouterSetStableSwap struct {
	Factory common.Address
	Info    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetStableSwap is a free log retrieval operation binding the contract event 0x26e41379222b54b0470031bc11852ad23058ffb8983f7cc0e18257d6f7afca9d.
//
// Solidity: event SetStableSwap(address indexed factory, address indexed info)
func (_SectaV3Router *SectaV3RouterFilterer) FilterSetStableSwap(opts *bind.FilterOpts, factory []common.Address, info []common.Address) (*SectaV3RouterSetStableSwapIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}
	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _SectaV3Router.contract.FilterLogs(opts, "SetStableSwap", factoryRule, infoRule)
	if err != nil {
		return nil, err
	}
	return &SectaV3RouterSetStableSwapIterator{contract: _SectaV3Router.contract, event: "SetStableSwap", logs: logs, sub: sub}, nil
}

// WatchSetStableSwap is a free log subscription operation binding the contract event 0x26e41379222b54b0470031bc11852ad23058ffb8983f7cc0e18257d6f7afca9d.
//
// Solidity: event SetStableSwap(address indexed factory, address indexed info)
func (_SectaV3Router *SectaV3RouterFilterer) WatchSetStableSwap(opts *bind.WatchOpts, sink chan<- *SectaV3RouterSetStableSwap, factory []common.Address, info []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}
	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _SectaV3Router.contract.WatchLogs(opts, "SetStableSwap", factoryRule, infoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SectaV3RouterSetStableSwap)
				if err := _SectaV3Router.contract.UnpackLog(event, "SetStableSwap", log); err != nil {
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

// ParseSetStableSwap is a log parse operation binding the contract event 0x26e41379222b54b0470031bc11852ad23058ffb8983f7cc0e18257d6f7afca9d.
//
// Solidity: event SetStableSwap(address indexed factory, address indexed info)
func (_SectaV3Router *SectaV3RouterFilterer) ParseSetStableSwap(log types.Log) (*SectaV3RouterSetStableSwap, error) {
	event := new(SectaV3RouterSetStableSwap)
	if err := _SectaV3Router.contract.UnpackLog(event, "SetStableSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
