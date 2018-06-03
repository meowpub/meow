package lib

import (
	"crypto/rand"
	"io"

	"github.com/keybase/saltpack/encoding/basex"
)

// Generates a random, 64-bit token, encoded as base62.
func GenToken() (string, error) {
	var data [64]byte
	if _, err := io.ReadFull(rand.Reader, data[:]); err != nil {
		return "", err
	}
	return basex.Base62StdEncoding.EncodeToString(data[:]), nil
}
