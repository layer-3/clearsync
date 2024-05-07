package voucher

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/layer-3/clearsync/pkg/abi/ivoucher"
)

var (
	addressTy = must(abi.NewType("address", "", nil))
	uint8Ty   = must(abi.NewType("uint8", "", nil))
	uint32Ty  = must(abi.NewType("uint32", "", nil))
	uint64Ty  = must(abi.NewType("uint64", "", nil))
	bytes32Ty = must(abi.NewType("bytes32", "", nil))
	bytesTy   = must(abi.NewType("bytes", "", nil))

	voucherArgs = abi.Arguments{
		{Name: "target", Type: addressTy},
		{Name: "action", Type: uint8Ty},
		{Name: "beneficiary", Type: addressTy},
		{Name: "expire", Type: uint64Ty},
		{Name: "chainId", Type: uint32Ty},
		{Name: "voucherCodeHash", Type: bytes32Ty},
		{Name: "encodedParams", Type: bytesTy},
	}
)

func must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

// Encode encodes the Voucher into a byte slice according to Ethereum ABI.
func Encode(voucher ivoucher.IVoucherVoucher) ([]byte, error) {
	packed, err := voucherArgs.Pack(
		voucher.Target,
		voucher.Action,
		voucher.Beneficiary,
		voucher.Expire,
		voucher.ChainId,
		voucher.VoucherCodeHash,
		voucher.EncodedParams,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to encode: %w", err)
	}

	return packed, nil
}

// Decode decodes a byte slice into a Voucher struct according to Ethereum ABI.
func Decode(voucher []byte) (ivoucher.IVoucherVoucher, error) {
	data := make(map[string]any)
	if err := voucherArgs.UnpackIntoMap(data, voucher); err != nil {
		return ivoucher.IVoucherVoucher{}, fmt.Errorf("failed to decode: %w", err)
	}

	var result ivoucher.IVoucherVoucher
	fields := map[string]reflect.Type{
		"target":          reflect.TypeOf(result.Target),
		"action":          reflect.TypeOf(result.Action),
		"beneficiary":     reflect.TypeOf(result.Beneficiary),
		"expire":          reflect.TypeOf(result.Expire),
		"chainId":         reflect.TypeOf(result.ChainId),
		"voucherCodeHash": reflect.TypeOf(result.VoucherCodeHash),
		"encodedParams":   reflect.TypeOf(result.EncodedParams),
	}

	valResult := reflect.ValueOf(&result).Elem()
	for key, expectedType := range fields {
		rawVal, ok := data[key]
		if !ok {
			return ivoucher.IVoucherVoucher{}, fmt.Errorf("missing %s", key)
		}
		val := reflect.ValueOf(rawVal)
		if !val.Type().AssignableTo(expectedType) {
			return ivoucher.IVoucherVoucher{}, fmt.Errorf("%s field has wrong type", key)
		}

		fieldVal := valResult.FieldByName(strings.Title(key))
		if !fieldVal.IsValid() {
			return ivoucher.IVoucherVoucher{}, fmt.Errorf("no such field: %s in result", key)
		}
		if !fieldVal.CanSet() {
			return ivoucher.IVoucherVoucher{}, fmt.Errorf("cannot set field: %s", key)
		}
		fieldVal.Set(val)
	}

	return result, nil
}
