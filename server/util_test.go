package server

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSanitiseRedirectURIs(t *testing.T) {
	t.Run("Relative", func(t *testing.T) {
		req := httptest.NewRequest("GET", "https://example.com/page", nil)
		u, err := SanitizeRedirectURI("/some/page", req)
		require.NoError(t, err)
		assert.Equal(t, "/some/page", u)
	})
	t.Run("Absolute", func(t *testing.T) {
		req := httptest.NewRequest("GET", "https://example.com/page", nil)
		_, err := SanitizeRedirectURI("https://example.com/some/page", req)
		require.EqualError(t, err, "absolute urls not allowed in redirect uris: https://example.com/some/page")
	})
	t.Run("Invalid", func(t *testing.T) {
		req := httptest.NewRequest("GET", "https://example.com/page", nil)
		_, err := SanitizeRedirectURI("://", req)
		require.EqualError(t, err, "parse ://: missing protocol scheme")
	})
}
