package userop

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	initCode, err := hexutil.Decode("0xbeefdead")
	require.NoError(t, err)

	callData, err := hexutil.Decode("0x34fcd5be0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000002a8b51821884cf9a7ea1a24c72e46ff52dcb4f16000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000224142cfda800000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000890000000000000000000000002a8b51821884cf9a7ea1a24c72e46ff52dcb4f160000000000000000000000000fb43b1ce0016df92e945155a7eadd3c9f2b2830000000000000000000000000dbb20123ccc4bc5cc283948969a196cbc573b5f5000000000000000000000000000000000000000000000000000000006748786a000000000000000000000000000000000000000000000000180c26068ef516d1000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000005d00000000000000000000000000000000000000000000000000000000000000418e8c6ccd5097e1d6f2dde1f940a9c23ac903c4471a57167b95221230cfacadd10754594046c0fc6feb98ab9dc1d9e1758e7aecc434a9397cffd3bb7f83a043201c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)

	paymasterAndData, err := hexutil.Decode("0xdeadbeef")
	require.NoError(t, err)

	signature, err := hexutil.Decode("0x00000000d00449d477e0ba7bce872b7bb85e64a3c97d7bfe0fb9b2d93ecddd28c295c0d93b9e18a567d688bccb9757e8c9e4dcb2a638889e3f76887a366ee56f476a1cda1c")
	require.NoError(t, err)

	userOp := UserOperation{
		Sender:               common.HexToAddress("0xDBB20123Ccc4Bc5cC283948969a196cBc573b5f5"),
		Nonce:                decimal.NewFromInt(130),
		InitCode:             initCode,
		CallData:             callData,
		CallGasLimit:         decimal.NewFromInt(132850),
		VerificationGasLimit: decimal.NewFromInt(540000),
		PreVerificationGas:   decimal.NewFromInt(370305),
		MaxFeePerGas:         decimal.NewFromInt(123456789),
		MaxPriorityFeePerGas: decimal.NewFromInt(10000000000),
		PaymasterAndData:     paymasterAndData,
		Signature:            signature,
	}

	marshalled, err := userOp.MarshalJSON()
	require.NoError(t, err)

	var userOpDTO UserOperationDTO
	err = json.Unmarshal(marshalled, &userOpDTO)
	require.NoError(t, err)

	assert.Equal(t, userOp.Sender.String(), userOpDTO.Sender)
	assert.Equal(t, "0x"+userOp.Nonce.BigInt().Text(16), userOpDTO.Nonce)
	assert.Equal(t, hexutil.Encode(userOp.InitCode), userOpDTO.InitCode)
	assert.Equal(t, hexutil.Encode(userOp.CallData), userOpDTO.CallData)
	assert.Equal(t, "0x"+userOp.CallGasLimit.BigInt().Text(16), userOpDTO.CallGasLimit)
	assert.Equal(t, "0x"+userOp.VerificationGasLimit.BigInt().Text(16), userOpDTO.VerificationGasLimit)
	assert.Equal(t, "0x"+userOp.PreVerificationGas.BigInt().Text(16), userOpDTO.PreVerificationGas)
	assert.Equal(t, "0x"+userOp.MaxFeePerGas.BigInt().Text(16), userOpDTO.MaxFeePerGas)
	assert.Equal(t, "0x"+userOp.MaxPriorityFeePerGas.BigInt().Text(16), userOpDTO.MaxPriorityFeePerGas)
	assert.Equal(t, hexutil.Encode(userOp.PaymasterAndData), userOpDTO.PaymasterAndData)
	assert.Equal(t, hexutil.Encode(userOp.Signature), userOpDTO.Signature)
}

func TestUnmarshal(t *testing.T) {
	userOpDTO := UserOperationDTO{
		Sender:               "0xDBB20123Ccc4Bc5cC283948969a196cBc573b5f5",
		Nonce:                "0x130",
		InitCode:             "0xbeefdead",
		CallData:             "0x34fcd5be0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000002a8b51821884cf9a7ea1a24c72e46ff52dcb4f16000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000224142cfda800000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000890000000000000000000000002a8b51821884cf9a7ea1a24c72e46ff52dcb4f160000000000000000000000000fb43b1ce0016df92e945155a7eadd3c9f2b2830000000000000000000000000dbb20123ccc4bc5cc283948969a196cbc573b5f5000000000000000000000000000000000000000000000000000000006748786a000000000000000000000000000000000000000000000000180c26068ef516d1000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000005d00000000000000000000000000000000000000000000000000000000000000418e8c6ccd5097e1d6f2dde1f940a9c23ac903c4471a57167b95221230cfacadd10754594046c0fc6feb98ab9dc1d9e1758e7aecc434a9397cffd3bb7f83a043201c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		CallGasLimit:         "0x144932",
		VerificationGasLimit: "0x83640",
		PreVerificationGas:   "0x59721",
		MaxFeePerGas:         "0x140842128458",
		MaxPriorityFeePerGas: "0x38750000000",
		PaymasterAndData:     "0xdeadbeef",
		Signature:            "0x00000000d00449d477e0ba7bce872b7bb85e64a3c97d7bfe0fb9b2d93ecddd28c295c0d93b9e18a567d688bccb9757e8c9e4dcb2a638889e3f76887a366ee56f476a1cda1c",
	}

	userOpBytes, err := json.Marshal(userOpDTO)
	require.NoError(t, err)

	var userOp UserOperation
	err = json.Unmarshal(userOpBytes, &userOp)
	require.NoError(t, err)
	assert.Equal(t, userOpDTO.Sender, userOp.Sender.String())
	assert.Equal(t, userOpDTO.Nonce, "0x"+userOp.Nonce.BigInt().Text(16))
	assert.Equal(t, userOpDTO.InitCode, hexutil.Encode(userOp.InitCode))
	assert.Equal(t, userOpDTO.CallData, hexutil.Encode(userOp.CallData))
	assert.Equal(t, userOpDTO.CallGasLimit, "0x"+userOp.CallGasLimit.BigInt().Text(16))
	assert.Equal(t, userOpDTO.VerificationGasLimit, "0x"+userOp.VerificationGasLimit.BigInt().Text(16))
	assert.Equal(t, userOpDTO.PreVerificationGas, "0x"+userOp.PreVerificationGas.BigInt().Text(16))
	assert.Equal(t, userOpDTO.MaxFeePerGas, "0x"+userOp.MaxFeePerGas.BigInt().Text(16))
	assert.Equal(t, userOpDTO.MaxPriorityFeePerGas, "0x"+userOp.MaxPriorityFeePerGas.BigInt().Text(16))
	assert.Equal(t, userOpDTO.PaymasterAndData, hexutil.Encode(userOp.PaymasterAndData))
	assert.Equal(t, userOpDTO.Signature, hexutil.Encode(userOp.Signature))
}

func TestUnmarshalEmpty(t *testing.T) {
	userOpDTO := UserOperationDTO{}

	userOpBytes, err := json.Marshal(userOpDTO)
	require.NoError(t, err)

	var userOp UserOperation
	err = json.Unmarshal(userOpBytes, &userOp)
	require.NoError(t, err)
	assert.Equal(t, common.Address{}.String(), userOp.Sender.String())
	assert.Equal(t, "0x0", "0x"+userOp.Nonce.BigInt().Text(16))
	assert.Equal(t, "0x", hexutil.Encode(userOp.InitCode))
	assert.Equal(t, "0x", hexutil.Encode(userOp.CallData))
	assert.Equal(t, "0x0", "0x"+userOp.CallGasLimit.BigInt().Text(16))
	assert.Equal(t, "0x0", "0x"+userOp.VerificationGasLimit.BigInt().Text(16))
	assert.Equal(t, "0x0", "0x"+userOp.PreVerificationGas.BigInt().Text(16))
	assert.Equal(t, "0x0", "0x"+userOp.MaxFeePerGas.BigInt().Text(16))
	assert.Equal(t, "0x0", "0x"+userOp.MaxPriorityFeePerGas.BigInt().Text(16))
	assert.Equal(t, "0x", hexutil.Encode(userOp.PaymasterAndData))
	assert.Equal(t, "0x", hexutil.Encode(userOp.Signature))
}

func TestDeepCopy(t *testing.T) {
	initCode, err := hexutil.Decode("0xbeefdead")
	require.NoError(t, err)

	callData, err := hexutil.Decode("0x34fcd5be0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000002a8b51821884cf9a7ea1a24c72e46ff52dcb4f16000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000224142cfda800000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000890000000000000000000000002a8b51821884cf9a7ea1a24c72e46ff52dcb4f160000000000000000000000000fb43b1ce0016df92e945155a7eadd3c9f2b2830000000000000000000000000dbb20123ccc4bc5cc283948969a196cbc573b5f5000000000000000000000000000000000000000000000000000000006748786a000000000000000000000000000000000000000000000000180c26068ef516d1000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000005d00000000000000000000000000000000000000000000000000000000000000418e8c6ccd5097e1d6f2dde1f940a9c23ac903c4471a57167b95221230cfacadd10754594046c0fc6feb98ab9dc1d9e1758e7aecc434a9397cffd3bb7f83a043201c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)

	paymasterAndData, err := hexutil.Decode("0xdeadbeef")
	require.NoError(t, err)

	signature, err := hexutil.Decode("0x00000000d00449d477e0ba7bce872b7bb85e64a3c97d7bfe0fb9b2d93ecddd28c295c0d93b9e18a567d688bccb9757e8c9e4dcb2a638889e3f76887a366ee56f476a1cda1c")
	require.NoError(t, err)

	userOp := UserOperation{
		Sender:               common.HexToAddress("0xDBB20123Ccc4Bc5cC283948969a196cBc573b5f5"),
		Nonce:                decimal.NewFromInt(130),
		InitCode:             initCode,
		CallData:             callData,
		CallGasLimit:         decimal.NewFromInt(132850),
		VerificationGasLimit: decimal.NewFromInt(540000),
		PreVerificationGas:   decimal.NewFromInt(370305),
		MaxFeePerGas:         decimal.NewFromInt(123456789),
		MaxPriorityFeePerGas: decimal.NewFromInt(10000000000),
		PaymasterAndData:     paymasterAndData,
		Signature:            signature,
	}

	userOpCopy := userOp.DeepCopy()

	// Verify pointers are different for slices
	require.False(t, &userOp.InitCode[0] == &userOpCopy.InitCode[0], "InitCode slice points to the same memory")
	require.False(t, &userOp.CallData[0] == &userOpCopy.CallData[0], "CallData slice points to the same memory")
	require.False(t, &userOp.PaymasterAndData[0] == &userOpCopy.PaymasterAndData[0], "PaymasterAndData slice points to the same memory")
	require.False(t, &userOp.Signature[0] == &userOpCopy.Signature[0], "Signature slice points to the same memory")

	// Modify the cloned object and ensure the original object remains unchanged
	userOpCopy.InitCode[0] = 0xAA
	assert.NotEqual(t, userOp.InitCode[0], userOpCopy.InitCode[0])

	userOpCopy.CallData[0] = 0xBB
	assert.NotEqual(t, userOp.CallData[0], userOpCopy.CallData[0])

	userOpCopy.PaymasterAndData[0] = 0xCC
	assert.NotEqual(t, userOp.PaymasterAndData[0], userOpCopy.PaymasterAndData[0])

	userOpCopy.Signature[0] = 0xDD
	assert.NotEqual(t, userOp.Signature[0], userOpCopy.Signature[0])
}
