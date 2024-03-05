package session_key

import (
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

// FIXME: v2.4
const SessionKeyValidatorAddress = "0x0000000000000000000000000000000000000000"
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
