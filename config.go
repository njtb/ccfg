package ccfg

import (
	"crypto"
)

type Algorithm int

type Config struct {
	Crypto crypto.Hash
	Hash   []byte
	Data   []byte
}

func (c Config) Verify() error {
	//skip check if not set
	if c.Crypto == 0 || len(c.Hash) == 0 {
		return nil
	}
	return VerifyCrypto(c.Crypto, c.Hash, c.Data)
}
