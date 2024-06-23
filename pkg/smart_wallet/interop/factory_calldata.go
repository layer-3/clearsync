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
} SmartWalletConfig_GetFactoryCallDataParams;

typedef struct {
  SmartWalletConfig_GetFactoryCallDataParams config;
  char* owner_address;
  int index;
} GetFactoryCallDataParams;

typedef struct {
  char* owner_address;
  int index;
  char* account_logic;
  char* ecdsa_validator;
} GetKernelFactoryCallDataParams;

typedef struct {
  char* owner_address;
  int index;
  char* ecdsa_validator;
} GetBiconomyFactoryCallDataParams;
*/
import "C"
import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/layer-3/clearsync/pkg/smart_wallet"
)

//export GetFactoryCallData
func GetFactoryCallData(params C.GetFactoryCallDataParams, err **C.char) *C.char {
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

	ownerAddress := common.HexToAddress(C.GoString(params.owner_address))
	index := decimal.NewFromInt(int64(params.index))

	callData, callErr := smart_wallet.GetFactoryCallData(config, ownerAddress, index)
	if callErr != nil {
		*err = C.CString(callErr.Error())
		return nil
	}

	return C.CString(string(callData))
}

//export GetKernelFactoryCallData
func GetKernelFactoryCallData(params C.GetKernelFactoryCallDataParams, err **C.char) *C.char {
	ownerAddress := common.HexToAddress(C.GoString(params.owner_address))
	index := decimal.NewFromInt(int64(params.index))
	accountLogic := common.HexToAddress(C.GoString(params.account_logic))
	ecdsaValidator := common.HexToAddress(C.GoString(params.ecdsa_validator))

	callData, callErr := smart_wallet.GetKernelFactoryCallData(ownerAddress, index, accountLogic, ecdsaValidator)
	if callErr != nil {
		*err = C.CString(callErr.Error())
		return nil
	}

	return C.CString(string(callData))
}

//export GetBiconomyFactoryCallData
func GetBiconomyFactoryCallData(params C.GetBiconomyFactoryCallDataParams, err **C.char) *C.char {
	ownerAddress := common.HexToAddress(C.GoString(params.owner_address))
	index := decimal.NewFromInt(int64(params.index))
	ecdsaValidator := common.HexToAddress(C.GoString(params.ecdsa_validator))

	callData, callErr := smart_wallet.GetBiconomyFactoryCallData(ownerAddress, index, ecdsaValidator)
	if callErr != nil {
		*err = C.CString(callErr.Error())
		return nil
	}

	return C.CString(string(callData))
}
