package session_key

import (
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

const SessionKeyValidatorAddress = "0x5C06CE2b673fD5E6e56076e40DD46aB67f5a72A5"
const ECDSAValidatorAddress = "0xd9AB5096a832b9ce79914329DAEE236f8Eea0390"

type Config struct {
	ProviderUrl                url.URL
	SessionKeyValidAfter       uint64
	SessionKeyValidUntil       uint64
	SessionKeyValidatorAddress common.Address
	ExecutorAddress            common.Address
	PaymasterAddress           common.Address
	Permissions                []Permission
}
