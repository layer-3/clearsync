package universal_sigver

import "errors"

var (
	ErrNotERC6492Sig       = errors.New("not an ERC-6492 signature")
	ErrCorruptedERC6492Sig = errors.New("corrupted ERC-6492 signature")
)
