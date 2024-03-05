package session_key

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/layer-3/clearsync/pkg/userop"
)

const (
	ValidatorApprovedStruct = "ValidatorApproved(bytes4 sig,uint256 validatorData,address executor,bytes enableData)"
	DomainStruct            = "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
	KernelDomainName        = "Kernel"
	KernelDomainVersion     = "0.2.2"
	KernelEnableDataLength  = 20 + 32 + 6 + 6 + 20 + 32
)

var (
	KernelExecuteSig      = [4]byte(userop.KernelExecuteABI.Methods["execute"].ID)
	KernelExecuteBatchSig = [4]byte(userop.KernelExecuteABI.Methods["executeBatch"].ID)
)

type SessionData struct {
	SessionKey common.Address
	ValidAfter time.Time
	ValidUntil time.Time

	// should be generated from the list of permissions
	MerkleRoot []byte

	// address(0) means accept userOp without paymaster,
	// address(1) means reject userOp with paymaster,
	// other address means accept userOp with paymaster with the address
	Paymaster common.Address

	// `SessionKeyValidator.nonces.lastNonce++` -> used in permissionKey to track executions
	Nonce *big.Int
}

func (sd SessionData) PackEnableData() []byte {
	validAfterEncoded := big.NewInt(sd.ValidAfter.Unix()).FillBytes(make([]byte, 6))
	validUntilEncoded := big.NewInt(sd.ValidUntil.Unix()).FillBytes(make([]byte, 6))
	nonceEncoded := sd.Nonce.FillBytes(make([]byte, 32))

	enableData := make([]byte, 0, KernelEnableDataLength)
	enableData = append(enableData, sd.SessionKey.Bytes()...)
	enableData = append(enableData, sd.MerkleRoot[:common.HashLength]...)
	enableData = append(enableData, validAfterEncoded...)
	enableData = append(enableData, validUntilEncoded...)
	enableData = append(enableData, sd.Paymaster.Bytes()...)
	enableData = append(enableData, nonceEncoded...)

	return enableData
}

func UnpackEnableData(signature []byte) (SessionData, error) {
	offset := 4 + 6 + 6 + 20 + 20 + 32
	enableData := signature[offset : offset+KernelEnableDataLength]

	return SessionData{
		SessionKey: common.BytesToAddress(enableData[:common.AddressLength]),
		MerkleRoot: enableData[common.AddressLength : common.AddressLength+common.HashLength],
		ValidAfter: time.Unix(int64(big.NewInt(0).SetBytes(enableData[common.AddressLength+common.HashLength:common.AddressLength+common.HashLength+6]).Uint64()), 0),
		ValidUntil: time.Unix(int64(big.NewInt(0).SetBytes(enableData[common.AddressLength+common.HashLength+6:common.AddressLength+common.HashLength+12]).Uint64()), 0),
		Paymaster:  common.BytesToAddress(enableData[common.AddressLength+common.HashLength+12 : common.AddressLength+common.HashLength+32]),
		Nonce:      big.NewInt(0).SetBytes(enableData[common.AddressLength+common.HashLength+32:]),
	}, nil
}

func getKernelSessionDataHash(sessionData SessionData, sig [4]byte, chainId *big.Int, kernelAddress, validator, executor common.Address) []byte {
	enableData := sessionData.PackEnableData()
	enableDataHash := getEnableDataHash(enableData, sig, sessionData.SessionKey, validator, executor)
	domainSeparator := getKernelDomainSeparator(chainId, kernelAddress)

	typedData := make([]byte, 0, 2+32+32)
	typedData = append(typedData, []byte{0x19, 0x01}...)
	typedData = append(typedData, domainSeparator...)
	typedData = append(typedData, enableDataHash...)

	return crypto.Keccak256(typedData)
}

func getKernelDomainSeparator(chainId *big.Int, kernelAddress common.Address) []byte {
	domainSeparator := make([]byte, 0, 32+32+32+32+32)

	domainSeparator = append(domainSeparator, crypto.Keccak256([]byte(DomainStruct))...)
	domainSeparator = append(domainSeparator, crypto.Keccak256([]byte(KernelDomainName))...)
	domainSeparator = append(domainSeparator, crypto.Keccak256([]byte(KernelDomainVersion))...)
	domainSeparator = append(domainSeparator, chainId.FillBytes(make([]byte, 32))...)
	domainSeparator = append(domainSeparator, make([]byte, 12)...)
	domainSeparator = append(domainSeparator, kernelAddress.Bytes()...)

	return crypto.Keccak256(domainSeparator)

}

func getEnableDataHash(enableData []byte, sig [4]byte, sessionKey, validator, executor common.Address) []byte {
	digest := make([]byte, 0, 32+32+32+32+32)

	digest = append(digest, crypto.Keccak256([]byte(ValidatorApprovedStruct))...)

	digest = append(digest, sig[:]...)
	digest = append(digest, make([]byte, 28)...)

	digest = append(digest, enableData[52:58]...)
	digest = append(digest, enableData[58:64]...)
	digest = append(digest, validator.Bytes()...)

	digest = append(digest, make([]byte, 12)...)
	digest = append(digest, executor.Bytes()...)

	digest = append(digest, crypto.Keccak256(enableData)...)

	return crypto.Keccak256(digest)
}
