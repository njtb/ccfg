package ccfg

import (
	"bytes"
	"crypto"
	"errors"
)

// VerifyCrypto ...
// @param crypto.Hash
// @param []byte
// @param []byte
// @return error
func VerifyCrypto(c crypto.Hash, hash, data []byte) error {
	if bytes.Compare(c.New().Sum(data), data) == 0 {
		return nil
	}
	return errors.New("verify failed")
}
