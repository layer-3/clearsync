package main

// #include <stdint.h>
// #include <string.h>
import "C"
import (
	"unsafe"

	"github.com/ethereum/go-ethereum/common"

	"github.com/layer-3/clearsync/pkg/universal_sigver"
)

//export IsERC6492Sig
func IsERC6492Sig(sig *C.uchar, length C.int) C.int {
	goSig := C.GoBytes(unsafe.Pointer(sig), length)

	if universal_sigver.IsERC6492Sig(goSig) {
		return 1
	}
	return 0
}

//export PackERC6492Sig
func PackERC6492Sig(factoryAddress *C.char, factoryCalldata *C.char, sig *C.char) *C.char {
	goFactoryAddress := common.HexToAddress(C.GoString(factoryAddress))
	goFactoryCalldata := C.GoBytes(unsafe.Pointer(factoryCalldata), C.int(C.strlen(factoryCalldata)))
	goSig := C.GoBytes(unsafe.Pointer(sig), C.int(C.strlen(sig)))
	packedSig := universal_sigver.PackERC6492Sig(goFactoryAddress, goFactoryCalldata, goSig)
	return C.CString(string(packedSig))
}

//export UnpackERC6492Sig
func UnpackERC6492Sig(sig *C.char, factoryAddress **C.char, factoryCalldata **C.char, signature **C.char, err **C.char) C.int {
	goSig := C.GoBytes(unsafe.Pointer(sig), C.int(C.strlen(sig)))
	factoryAddr_, factoryCalldata_, sig_, unpackErr := universal_sigver.UnpackERC6492Sig(goSig)
	if unpackErr != nil {
		*err = C.CString(unpackErr.Error())
		return 0
	}
	*factoryAddress = C.CString(factoryAddr_.Hex())
	*factoryCalldata = C.CString(string(factoryCalldata_))
	*signature = C.CString(string(sig_))
	return 1
}

func main() {}
