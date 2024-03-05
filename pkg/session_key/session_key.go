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
	KernelEnableDataLength  = 20 + 32 + 6 + 6 + 20 + 32
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

func (sd SessionData) Encode() []byte {
	validAfterEncoded := make([]byte, 6)
	big.NewInt(sd.ValidAfter.Unix()).FillBytes(validAfterEncoded)

	validUntilEncoded := make([]byte, 6)
	big.NewInt(sd.ValidUntil.Unix()).FillBytes(validUntilEncoded)

	// TODO:
	nonceEncoded := make([]byte, 32)
	sd.Nonce.FillBytes(nonceEncoded)

	enableData := make([]byte, 0, KernelEnableDataLength)
	enableData = append(enableData, sd.SessionKey.Bytes()...)
	enableData = append(enableData, sd.MerkleRoot[:common.HashLength]...)
	enableData = append(enableData, validAfterEncoded...)
	enableData = append(enableData, validUntilEncoded...)
	enableData = append(enableData, sd.Paymaster.Bytes()...)
	enableData = append(enableData, nonceEncoded...)

	return enableData
}

func GetKernelSessionDataHash(sessionData SessionData, sig [4]byte, chainId *big.Int, kernelAddress, validator, executor common.Address) []byte {
	enableData := sessionData.Encode()
	enableDataHash := getEnableDataHash(enableData, sig, sessionData.SessionKey, validator, executor)
	domainSeparator := getKernelDomainSeparator(chainId, kernelAddress)

	typedData := make([]byte, 0, 2+32+32)
	// "use given validator" (0x00000001) mode
	// see https://github.com/zerodevapp/kernel/blob/807b75a4da6fea6311a3573bc8b8964a34074d94/src/Kernel.sol#L127
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
