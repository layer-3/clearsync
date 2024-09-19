package voucher_v2

import (
	"fmt"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/layer-3/clearsync/pkg/abi/ivoucher_v2"
	signer_pkg "github.com/layer-3/clearsync/pkg/signer"
)

type VoucherSigner struct {
	signer signer_pkg.Signer
	vs     []ivoucher_v2.IVoucherVoucher
}

func NewVoucherSigner() *VoucherSigner {
	return &VoucherSigner{}
}

func (vs *VoucherSigner) AddVoucher(v ivoucher_v2.IVoucherVoucher) *VoucherSigner {
	vs.vs = append(vs.vs, v)

	return vs
}

func (vs *VoucherSigner) WithSigner(signer signer_pkg.Signer) *VoucherSigner {
	vs.signer = signer

	return vs
}

func (vs *VoucherSigner) Sign() error {
	for i, v := range vs.vs {
		if v.Signature != nil {
			continue
		}

		unsignedData, err := Encode(v)
		if err != nil {
			return err
		}

		hashedBytes := ecrypto.Keccak256Hash(unsignedData).Bytes()
		sig, err := signer_pkg.SignEthMessage(vs.signer, hashedBytes)
		if err != nil {
			return err
		}

		vs.vs[i].Signature = sig.Raw()
	}

	return nil
}

func (vs *VoucherSigner) Encode() ([]byte, error) {
	voucherRouterABI, err := ivoucher_v2.VoucherRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	packed, err := voucherRouterABI.Methods["use"].Inputs[0:1].Pack(vs.vs)
	if err != nil {
		return nil, fmt.Errorf("failed to encode: %w", err)
	}

	return packed, nil
}

func (vs *VoucherSigner) Vouchers() []ivoucher_v2.IVoucherVoucher {
	return vs.vs
}

// Most of the time we are working with a single voucher
func (vs *VoucherSigner) First() (ivoucher_v2.IVoucherVoucher, error) {
	if len(vs.vs) == 0 {
		return ivoucher_v2.IVoucherVoucher{}, fmt.Errorf("no vouchers")
	}

	return vs.vs[0], nil
}
