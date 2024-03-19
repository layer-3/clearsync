package universal_sigver

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ERC-6492 defined suffix to identify the ERC-6492 signature.
const erc6492MagicValue = "0x6492649264926492649264926492649264926492649264926492649264926492"

// IsERC6492Sig checks if the signature contains the ERC-6492 suffix, meaning it is an ERC-6492 signature.
//
// Parameters:
// - sig: the signature
//
// Returns:
// - true if the signature is an ERC-6492 signature, false otherwise
func IsERC6492Sig(sig []byte) bool {
	return len(sig) >= 32 && hexutil.Encode(sig[len(sig)-32:]) == erc6492MagicValue
}

// PackERC6492Sig packs an ERC-6492 signature by adding smart wallet factory address, calldata and ERC-6492 suffix.
//
// Parameters:
// - factoryAddress: the address of the smart wallet factory
// - factoryCalldata: the calldata for the smart wallet factory
// - sig: the signature
//
// Returns:
// - the packed signature
func PackERC6492Sig(factoryAddress common.Address, factoryCalldata, sig []byte) []byte {
	args := abi.Arguments{
		{Name: "factory", Type: address},
		{Name: "calldata", Type: bytes},
		{Name: "signature", Type: bytes},
	}

	packed, err := args.Pack(factoryAddress, factoryCalldata, sig)
	if err != nil {
		panic(fmt.Errorf("failed to pack ERC-6492 sig: %w", err))
	}

	return append(packed, hexutil.MustDecode(erc6492MagicValue)...)
}

// UnpackERC6492Sig unpacks an ERC-6492 signature into a smart wallet factory address, calldata and signature.
//
// Parameters:
// - sig: the signature
//
// Returns:
// - the smart wallet factory address
// - the calldata for the smart wallet factory
// - the signature
// - an error if the signature is not an ERC-6492 signature or it is corrupted
func UnpackERC6492Sig(sig []byte) (factoryAddress common.Address, factoryCalldata, signature []byte, err error) {
	if !IsERC6492Sig(sig) {
		return common.Address{}, nil, nil, ErrNotERC6492Sig
	}

	args := abi.Arguments{
		{Name: "factory", Type: address},
		{Name: "calldata", Type: bytes},
		{Name: "signature", Type: bytes},
	}

	unpacked, err := args.Unpack(sig[:len(sig)-32])
	if err != nil {
		return common.Address{}, nil, nil, ErrCorruptedERC6492Sig
	}

	var ok bool
	factoryAddress, ok = unpacked[0].(common.Address)
	if !ok {
		return common.Address{}, nil, nil, ErrCorruptedERC6492Sig
	}
	factoryCalldata, ok = unpacked[1].([]byte)
	if !ok {
		return common.Address{}, nil, nil, ErrCorruptedERC6492Sig
	}
	signature, ok = unpacked[2].([]byte)
	if !ok {
		return common.Address{}, nil, nil, ErrCorruptedERC6492Sig
	}

	return
}
