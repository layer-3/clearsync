package session_key

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	ValidatorApprovedStruct = "ValidatorApproved(bytes4 sig,uint256 validatorData,address executor,bytes enableData)"
	DomainStruct            = "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
	KernelDomainName        = "Kernel"
	KernelDomainVersion     = "0.2.2"
)

type SessionData struct {
	ValidAfter time.Time
	ValidUntil time.Time

	// should be generated from the list of permissions
	MerkleRoot []byte

	// address(0) means accept userOp without paymaster,
	// address(1) means reject userOp with paymaster,
	// other address means accept userOp with paymaster with the address
	Paymaster common.Address

	// ? -> used in permissionKey to track executions
	Nonce *big.Int
}

func (sd SessionData) Encode(sessionKey common.Address) []byte {
	validAfterEncoded := make([]byte, 6)
	big.NewInt(sd.ValidAfter.Unix()).FillBytes(validAfterEncoded)

	validUntilEncoded := make([]byte, 6)
	big.NewInt(sd.ValidUntil.Unix()).FillBytes(validUntilEncoded)

	nonceEncoded := make([]byte, 32)
	sd.Nonce.FillBytes(nonceEncoded)

	enableData := make([]byte, 0, 20+32+6+6+20+32)
	enableData = append(enableData, sessionKey.Bytes()...)
	enableData = append(enableData, sd.MerkleRoot[:common.HashLength]...)
	enableData = append(enableData, validAfterEncoded...)
	enableData = append(enableData, validUntilEncoded...)
	enableData = append(enableData, sd.Paymaster.Bytes()...)
	enableData = append(enableData, nonceEncoded...)

	return enableData
}

func ComputeKernelSessionDataHash(sessionData SessionData, sig [4]byte, chainId *big.Int, kernelAddress, sessionKey, validator, executor common.Address) []byte {
	enableData := sessionData.Encode(sessionKey)
	enableDataHash := buildEnableDataHash(enableData, sig, sessionKey, validator, executor)
	domainSeparator := buildKernelDomainSeparator(chainId, kernelAddress)

	typedData := make([]byte, 0, 2+32+32)
	typedData = append(typedData, []byte{0x19, 0x01}...)
	typedData = append(typedData, domainSeparator...)
	typedData = append(typedData, enableDataHash...)

	return crypto.Keccak256(typedData)
}

func buildKernelDomainSeparator(chainId *big.Int, kernelAddress common.Address) []byte {
	domainSeparator := make([]byte, 0, 32+32+32+32+32)

	domainSeparator = append(domainSeparator, crypto.Keccak256([]byte(DomainStruct))...)
	domainSeparator = append(domainSeparator, crypto.Keccak256([]byte(KernelDomainName))...)
	domainSeparator = append(domainSeparator, crypto.Keccak256([]byte(KernelDomainVersion))...)
	domainSeparator = append(domainSeparator, chainId.FillBytes(make([]byte, 32))...)
	domainSeparator = append(domainSeparator, make([]byte, 12)...)
	domainSeparator = append(domainSeparator, kernelAddress.Bytes()...)

	return crypto.Keccak256(domainSeparator)

}

func buildEnableDataHash(enableData []byte, sig [4]byte, sessionKey, validator, executor common.Address) []byte {
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
