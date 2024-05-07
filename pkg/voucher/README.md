# Voucher Package

The `voucher` package provides functionality to encode and decode voucher data structures using Ethereum's ABI (
Application Binary Interface) encoding. It specifically handles the encoding and decoding of `IVoucherVoucher`
structures as defined by the `clearsync` Ethereum contract interface (see `github.com/layer-3/clearsync/pkg/abi/ivoucher`).

## Usage

The package provides two main functions:

1. **`Encode(voucher ivoucher.IVoucherVoucher) ([]byte, error)`**: This function takes a `IVoucherVoucher` structure and
   returns its ABI-encoded byte slice. It returns an error if the encoding fails.

2. **`Decode(voucher []byte) (ivoucher.IVoucherVoucher, error)`**: This function decodes an ABI-encoded byte slice into
   a `IVoucherVoucher` structure. It returns an error if the decoding fails.

### Example

Here's a basic example on how to use this package to encode and decode a `IVoucherVoucher`:

```go
package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/layer-3/clearsync/pkg/abi/ivoucher"
	"github.com/layer-3/clearsync/pkg/voucher"
)

func main() {
	v := ivoucher.IVoucherVoucher{
		Target:          common.HexToAddress("0xabc123abc123abc123abc123abc123abc123abc1"),
		Action:          3,
		Beneficiary:     common.HexToAddress("0xdef456def456def456def456def456def456def4"),
		Expire:          1715100785,
		ChainId:         59144,
		VoucherCodeHash: [32]byte{/*--snip--*/},
		EncodedParams:   []byte("paramData"),
	}

	// Encode
	encoded, err := voucher.Encode(v)
	if err != nil {
		log.Fatal("Failed to encode voucher: ", err)
	}
	fmt.Println("Encoded voucher:", common.Bytes2Hex(encoded))

	// Decode
	decoded, err := voucher.Decode(encoded)
	if err != nil {
		log.Fatal("Failed to decode voucher:", err)
		return
	}
	fmt.Println("Result:", decoded)
}
```
