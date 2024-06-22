package main

// #include <stdint.h>
import "C"
import (
	"unsafe"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const Erc6492MagicValue = "0x6492649264926492649264926492649264926492649264926492649264926492"

//export IsERC6492SigWrapper
func IsERC6492SigWrapper(sig *C.uchar, length C.int) C.int {
	goSig := C.GoBytes(unsafe.Pointer(sig), length)

	if len(goSig) >= 32 && hexutil.Encode(goSig[len(goSig)-32:]) == Erc6492MagicValue {
		return 1
	}
	return 0
}

func main() {}
