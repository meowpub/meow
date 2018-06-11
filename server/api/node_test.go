package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNode(t *testing.T) {
	leaf := &Node{
		Self: &Node{
			Self: HandlerFunc(func(ctx context.Context, req *http.Request) Response {
				return Response{Data: "leaf"}
			}),
			Children: map[string]Handler{
				"self-leaf": HandlerFunc(func(ctx context.Context, req *http.Request) Response {
					return Response{Data: "self-leaf"}
				}),
			},
		},
	}
	root := &Node{
		Self: HandlerFunc(func(ctx context.Context, req *http.Request) Response {
			return Response{Data: "OK"}
		}),
		Children: map[string]Handler{
			"leaf": leaf,
		},
	}

	t.Run("HandleRequest", func(t *testing.T) {
		req := httptest.NewRequest("GET", "https://example.com/", nil)
		assert.Equal(t, Response{Data: "OK"}, root.HandleRequest(req.Context(), req))
	})

	t.Run("Traverse", func(t *testing.T) {
		t.Run("Not Found", func(t *testing.T) {
			h, err := root.Traverse(context.Background(), "not-found")
			require.NoError(t, err)
			assert.Nil(t, h)
		})
		t.Run("Found", func(t *testing.T) {
			h, err := root.Traverse(context.Background(), "leaf")
			require.NoError(t, err)
			assert.Equal(t, leaf, h)

			t.Run("HandleRequest", func(t *testing.T) {
				req := httptest.NewRequest("GET", "https://example.com/leaf", nil)
				assert.Equal(t, Response{Data: "leaf"}, h.HandleRequest(req.Context(), req))
			})

			t.Run("Self Traverse", func(t *testing.T) {
				h2, err := h.(*Node).Traverse(context.Background(), "self-leaf")
				require.NoError(t, err)
				req := httptest.NewRequest("GET", "https://example.com/leaf/self-leaf", nil)
				assert.Equal(t, Response{Data: "self-leaf"}, h2.HandleRequest(req.Context(), req))
			})
		})
	})
}

func TestNodeNoSelf(t *testing.T) {
	req := httptest.NewRequest("GET", "https://example.com/", nil)
	assert.Equal(t, Response{Status: http.StatusNotFound}, (&Node{}).HandleRequest(req.Context(), req))
}
