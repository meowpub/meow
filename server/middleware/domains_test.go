package middleware

import (
	"net/http"
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
		assert.Equal(t,
			api.Response{Status: http.StatusMisdirectedRequest},
			LocalDomains([]string{"example.com"})(handler).HandleRequest(api.Request{Request: req}),
		)
	})
}
