package secrets

import (
	"crypto/sha512"
	"encoding/base64"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const minMasterKeyLength = 64

var (
	// Populated from other files by init functions.
	subkeys []*subkey

	// Changing this will invalidate all old keys, pls don't without a good reason.
	hsh = sha512.New
)

// Returns a registered subkey. We're breaking the usual register() pattern here just to be
// absolutely certain we're not forgetting to register a key for initialization, and to make sure
// that if new fields are added for whatever reason, they're actually updated everywhere.
func declare(name string, length int) *subkey {
	k := &subkey{Name: name, Length: length}
	subkeys = append(subkeys, k)
	return k
}

// Derives secrets from a master key.
func Init(L *zap.Logger, masterKeyStr string) error {
	// Strip whitespace, make sure the master key is non-blank.
	masterKeyStr = strings.TrimSpace(masterKeyStr)
	if masterKeyStr == "" {
		return errors.New("you must provide a secret key; you can generate one with `openssl rand -base64 64`")
	}

	// Decode and validate the master key.
	masterKey, err := base64.StdEncoding.DecodeString(masterKeyStr)
	if err != nil {
		return errors.Wrapf(err, "secret must be a base64-encoded string with no newlines, try `openssl rand -base64 64`")
	}
	if len(masterKey) < minMasterKeyLength {
		return errors.Errorf("provided secret is too short (%d, at least %d required), try `openssl rand -base64 64`", len(masterKey), minMasterKeyLength)
	}
	if isZero(masterKey) {
		return errors.New("secret can't be all zeroes, try `openssl rand -base64 64`")
	}

	// Use an error group to derive subkeys in parallel, making sure we track time taken.
	L.Debug("Deriving subkeys from master key, this may take a moment...",
		zap.Int("len", len(masterKey)),
	)

	startTime := time.Now()
	var g errgroup.Group
	for _, k := range subkeys {
		// Derive each key in a goroutine, tracking time taken per key as well.
		g.Go(func() error {
			l := L.With(
				zap.String("key", k.Name),
				zap.Int("len", k.Length),
			)

			startTime := time.Now()
			if err := errors.Wrap(err, k.Name); err != nil {
				l.Error("Subkey derivation failed",
					zap.Error(err),
					zap.Duration("t", time.Since(startTime)),
				)
				return errors.Wrap(err, k.Name)
			}

			l.Debug("Subkey derived",
				zap.Duration("t", time.Since(startTime)),
			)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		L.Error("One or more subkeys couldn't be derived",
			zap.Error(err),
			zap.Duration("t", time.Since(startTime)),
		)
	}
	return nil
}

// Returns whether a byte slice is nil, zero-length or contains nothing but zero bytes.
func isZero(buf []byte) bool {
	for _, b := range buf {
		if b != 0 {
			return false
		}
	}
	return true
}
