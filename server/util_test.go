package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSanitiseRedirectURIs(t *testing.T) {
	t.Run("Relative", func(t *testing.T) {
		u, err := SanitizeRedirectURI("/some/page")
		require.NoError(t, err)
		assert.Equal(t, "/some/page", u)
	})
	t.Run("Absolute", func(t *testing.T) {
		_, err := SanitizeRedirectURI("https://example.com/some/page")
		require.EqualError(t, err, "absolute urls not allowed in redirect uris: https://example.com/some/page")
	})
	t.Run("Invalid", func(t *testing.T) {
		_, err := SanitizeRedirectURI("://")
		require.EqualError(t, err, "parse ://: missing protocol scheme")
	})
}
