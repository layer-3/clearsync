package main

/*
#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>

typedef struct {
  char* typ;
  char* ecdsa_validator;
  char* logic;
  char* factory;
} SmartWalletConfig_GetAccountAddressParams;

typedef struct {
  char* provider_url;
  SmartWalletConfig_GetAccountAddressParams config;
  char* entry_point_address;
  char* owner_address;
  int index;
} GetAccountAddressParams;
*/
import "C"
import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/smart_wallet"
)

//export GetAccountAddress
func GetAccountAddress(params C.GetAccountAddressParams, err **C.char) *C.char {
	provider, providerErr := ethclient.Dial(C.GoString(params.provider_url))
	if providerErr != nil {
		*err = C.CString(providerErr.Error())
		return nil
	}
	defer provider.Close()

	config := smart_wallet.Config{
		ECDSAValidator: common.HexToAddress(C.GoString(params.config.ecdsa_validator)),
		Logic:          common.HexToAddress(C.GoString(params.config.logic)),
		Factory:        common.HexToAddress(C.GoString(params.config.factory)),
	}

	typ := C.GoString(params.config.typ)
	if typ == "simple_account" {
		config.Type = &smart_wallet.SimpleAccountType
	} else if typ == "biconomy" {
		config.Type = &smart_wallet.BiconomyType
	} else if typ == "kernel" {
		config.Type = &smart_wallet.KernelType
	} else {
		config.Type = nil // setting it explicitly just in case
	}

	entryPointAddress := common.HexToAddress(C.GoString(params.entry_point_address))
	ownerAddress := common.HexToAddress(C.GoString(params.owner_address))
	index := decimal.NewFromInt(int64(params.index))

	swAddress, swErr := smart_wallet.GetAccountAddress(
		context.Background(),
		provider,
		config,
		entryPointAddress,
		ownerAddress,
		index,
	)
	if swErr != nil {
		*err = C.CString(swErr.Error())
		return nil
	}

	return C.CString(swAddress.Hex())
}

func main() {}
