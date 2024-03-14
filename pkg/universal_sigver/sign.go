package universal_sigver

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const erc6492MagicValue = "6492649264926492649264926492649264926492649264926492649264926492"

func packERC6492Sig(factoryAddress common.Address, factoryCalldata, sig []byte) []byte {
	args := abi.Arguments{
		{Name: "factory", Type: address},
		{Name: "calldata", Type: bytes},
		{Name: "signature", Type: bytes},
	}

	packed, err := args.Pack(factoryAddress, factoryCalldata, sig)
	if err != nil {
		panic(fmt.Errorf("failed to pack ERC-6492 sig: %w", err))
	}

	return append(packed, common.Hex2Bytes(erc6492MagicValue)...)
}
