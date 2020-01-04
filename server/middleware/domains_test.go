package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/meowpub/meow/server/api"
	"github.com/stretchr/testify/assert"
)

func TestLocalDomains(t *testing.T) {
	handler := api.HandlerFunc(func(api.Request) api.Response {
		return api.Response{Data: map[string]string{"hello": "world"}}
	})
	t.Run("Local", func(t *testing.T) {
		req := httptest.NewRequest("GET", "https://example.com/abc123", nil)
		assert.Equal(t,
			api.Response{Data: map[string]string{"hello": "world"}},
			LocalDomains([]string{"example.com"})(handler).HandleRequest(api.Request{Request: req}),
		)
	})
	t.Run("Remote", func(t *testing.T) {
		req := httptest.NewRequest("GET", "https://example2.com/abc123", nil)
		assert.EqualError(t,
			LocalDomains([]string{"example.com"})(handler).HandleRequest(api.Request{Request: req}).Error,
			"The server was not configured to handle the domain: 'example2.com'",
		)
	})
}
