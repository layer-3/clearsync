package session_key

import (
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

// FIXME: v2.4
const SessionKeyValidatorAddress = "0x0000000000000000000000000000000000000000"
const ECDSAValidatorAddress = "0xd9AB5096a832b9ce79914329DAEE236f8Eea0390"

// TODO: add sessionKeyValidAfter, sessionKeyValidUntil (offset or smth); executor, paymaster addresses
type Config struct {
	ProviderUrl                url.URL
	SessionKeyValidatorAddress common.Address
	Permissions                []Permission
}
