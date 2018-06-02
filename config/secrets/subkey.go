package secrets

import (
	"hash"
	"io"

	"github.com/pkg/errors"
	"golang.org/x/crypto/hkdf"
)

type subkey struct {
	Name   string
	Length int

	value []byte
}

// Derives a value for this subkey from the master key.
func (k *subkey) init(hsh func() hash.Hash, masterKey []byte) error {
	// Use HKDF to derive a subkey from the master key.
	r := hkdf.New(hsh, masterKey, nil, []byte(k.Name))
	key := make([]byte, k.Length)
	if _, err := io.ReadFull(r, key); err != nil {
		return err
	}

	// Sanity check the heck out of it to be safe.
	if len(key) != k.Length {
		return errors.Errorf("derived key has the wrong length: %d != %d", len(key), k.Length)
	}
	if isZero(key) {
		return errors.New("derived key is all zeroes?????")
	}

	// Only now export the key.
	k.value = key
	return nil
}

// Returns the value, or panics if it's uninitialized or invalid in some way.
func (k *subkey) access() []byte {
	if isZero(k.value) {
		panic(k.Name + ": key uninitialized; did you call secrets.Init() properly?")
	}
	if len(k.value) != k.Length {
		panic(k.Name + ": key has an invalid length, what the heck did you do???")
	}
	return k.value
}
