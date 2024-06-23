package main

/*
#include <stdint.h>

typedef struct {
    char* typ;
	char* ecdsa_validator;
	char* logic;
	char* factory;
} SmartWalletConfig;
*/
import "C"
import (
	"context"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/smart_wallet"
	"github.com/layer-3/clearsync/pkg/universal_sigver"
)

//export NewUniversalSigVer
func NewUniversalSigVer(providerURL *C.char, smartWalletConfig C.SmartWalletConfig, entryPointAddress *C.char) C.uintptr_t {
	config := smart_wallet.Config{
		ECDSAValidator: common.HexToAddress(C.GoString(smartWalletConfig.ecdsa_validator)),
		Logic:          common.HexToAddress(C.GoString(smartWalletConfig.logic)),
		Factory:        common.HexToAddress(C.GoString(smartWalletConfig.factory)),
	}

	typ := C.GoString(smartWalletConfig.typ)
	if typ == "simple_account" {
		config.Type = &smart_wallet.SimpleAccountType
	} else if typ == "biconomy" {
		config.Type = &smart_wallet.BiconomyType
	} else if typ == "kernel" {
		config.Type = &smart_wallet.KernelType
	} else {
		config.Type = nil // setting it explicitly just in case
	}

	address := common.HexToAddress(C.GoString(entryPointAddress))
	client, err := universal_sigver.NewUniversalSigVer(C.GoString(providerURL), config, address)
	if err != nil {
		return 0
	}
	return C.uintptr_t(uintptr(unsafe.Pointer(&client)))
}

//export UniversalSigVer_Verify
func UniversalSigVer_Verify(client C.uintptr_t, signer *C.char, messageHash *C.char, signature *C.char) (C.int, *C.char) {
	c := *(*universal_sigver.Client)(unsafe.Pointer(uintptr(client)))
	ctx := context.Background()
	signerAddr := common.HexToAddress(C.GoString(signer))
	msgHash := common.HexToHash(C.GoString(messageHash))
	sig := []byte(C.GoString(signature))
	valid, err := c.Verify(ctx, signerAddr, msgHash, sig)
	if err != nil {
		return 0, C.CString(err.Error())
	}

	if valid {
		return 1, nil // 1 is true in C
	}
	return 0, nil // 0 is false in C
}

//export UniversalSigVer_PackERC6492Sig
func UniversalSigVer_PackERC6492Sig(client C.uintptr_t, ownerAddress *C.char, index *C.char, sig *C.char) (*C.char, *C.char) {
	c := *(*universal_sigver.Client)(unsafe.Pointer(uintptr(client)))
	ctx := context.Background()
	ownerAddr := common.HexToAddress(C.GoString(ownerAddress))
	idx, _ := decimal.NewFromString(C.GoString(index))
	signature := []byte(C.GoString(sig))
	packedSig, err := c.PackERC6492Sig(ctx, ownerAddr, idx, signature)
	if err != nil {
		return nil, C.CString(err.Error())
	}
	return C.CString(string(packedSig)), nil
}
