package smart_wallet

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func GetStubSignature(swt Type) ([]byte, error) {
	switch swt {
	case SimpleAccountType, BiconomyType:
		return hexutil.Decode("0xfffffffffffffffffffffffffffffff0000000000000000000000000000000007aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1c")
	case KernelType:
		return hexutil.Decode("0x00000000fffffffffffffffffffffffffffffff0000000000000000000000000000000007aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1c")
	default:
		return nil, fmt.Errorf("unknown smart wallet type: %s", swt)
	}
}
