package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unrolled/render"
)

func TestRouterHandleRequest(t *testing.T) {
	roots := map[string]*Node{
		"https://example.com": &Node{
			Self: HandlerFunc(func(ctx context.Context, req *http.Request) Response {
				return Response{Data: "OK"}
			}),
		},
	}
	r := NewRouter(func(ctx context.Context, url string) (Traversible, error) {
		if n, ok := roots[url]; ok {
			return n, nil
		}
		return nil, nil
	})

	r.Mount("-", &Node{
		Children: map[string]Handler{
			"route": HandlerFunc(func(ctx context.Context, req *http.Request) Response {
				return Response{Status: http.StatusTeapot, Data: "i'm a teapot"}
			}),
		},
	})

	t.Run("Root", func(t *testing.T) {
		req := httptest.NewRequest("GET", "https://example.com/", nil)
		assert.Equal(t, Response{Data: "OK"}, r.HandleRequest(req.Context(), req))

		t.Run("-", func(t *testing.T) {
			req := httptest.NewRequest("GET", "https://example.com/-/", nil)
			assert.Equal(t, Response{Status: http.StatusNotFound}, r.HandleRequest(req.Context(), req))

			t.Run("route", func(t *testing.T) {
				req := httptest.NewRequest("GET", "https://example.com/-/route", nil)
				assert.Equal(t, Response{Status: http.StatusTeapot, Data: "i'm a teapot"}, r.HandleRequest(req.Context(), req))
			})
		})
	})

	t.Run("Host with Port", func(t *testing.T) {
		req := httptest.NewRequest("GET", "https://example.com:443/", nil)
		assert.Equal(t, Response{Data: "OK"}, r.HandleRequest(req.Context(), req))
	})

	t.Run("Nonexistent Root", func(t *testing.T) {
		req := httptest.NewRequest("GET", "https://example.co.uk/", nil)
		assert.Equal(t, Response{Status: http.StatusNotFound}, r.HandleRequest(req.Context(), req))

		t.Run("-", func(t *testing.T) {
			req := httptest.NewRequest("GET", "https://example.co.uk/-/", nil)
			assert.Equal(t, Response{Status: http.StatusNotFound}, r.HandleRequest(req.Context(), req))

			t.Run("route", func(t *testing.T) {
				req := httptest.NewRequest("GET", "https://example.co.uk/-/route", nil)
				assert.Equal(t, Response{Status: http.StatusNotFound}, r.HandleRequest(req.Context(), req))
			})
		})
	})
}

func TestRouterServe(t *testing.T) {
	templates := map[string]string{
		"templates/template.tmpl": `Hello {{.}}`,
	}
	root := &Node{}
	r := NewRouter(func(ctx context.Context, url string) (Traversible, error) {
		return root, nil
	}, render.Options{
		Asset: func(name string) ([]byte, error) {
			if tpl, ok := templates[name]; ok {
				return []byte(tpl), nil
			}
			return nil, errors.Errorf("couldn't find template: %s", name)
		},
		AssetNames: func() (names []string) {
			for name := range templates {
				names = append(names, name)
			}
			return
		},
	})

	simpletestdata := map[string]struct {
		Data interface{}
		Body string
	}{
		"String": {"OK", "OK"},
		"Binary": {[]byte("OK"), "OK"},
		"Map":    {map[string]interface{}{"a": 1}, `{"a":1}`},
		"Empty":  {nil, ""},
	}
	for name, tdata := range simpletestdata {
		t.Run(name, func(t *testing.T) {
			statuses := map[int]int{
				0:                   http.StatusOK,
				http.StatusOK:       http.StatusOK,
				http.StatusNotFound: http.StatusNotFound,
			}
			for instatus, outstatus := range statuses {
				t.Run(strconv.Itoa(instatus), func(t *testing.T) {
					root.Self = HandlerFunc(func(ctx context.Context, req *http.Request) Response {
						return Response{Status: instatus, Data: tdata.Data}
					})

					rec := httptest.NewRecorder()
					r.ServeHTTP(rec, httptest.NewRequest("GET", "https://example.com/", nil))
					assert.Equal(t, outstatus, rec.Code)
					assert.Equal(t, tdata.Body, rec.Body.String())
				})
			}
		})
	}

	t.Run("Template", func(t *testing.T) {
		root.Self = HandlerFunc(func(ctx context.Context, req *http.Request) Response {
			return Response{Data: "World", Template: "template"}
		})

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, httptest.NewRequest("GET", "https://example.com/", nil))
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello World", rec.Body.String())
	})

	t.Run("Headers", func(t *testing.T) {
		root.Self = HandlerFunc(func(ctx context.Context, req *http.Request) Response {
			resp := Response{}
			resp.Header().Set("X-My-Header", "value")
			resp.WriteHeader(http.StatusTeapot)
			resp.Write([]byte("i'm a teapot"))
			return resp
		})

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, httptest.NewRequest("GET", "https://example.com/", nil))
		assert.Equal(t, http.StatusTeapot, rec.Code)
		assert.Equal(t, "value", rec.Header().Get("X-My-Header"))
		assert.Equal(t, "i'm a teapot", rec.Body.String())
	})

	t.Run("Error", func(t *testing.T) {
		root.Self = HandlerFunc(func(ctx context.Context, req *http.Request) Response {
			return Response{Error: Wrap(errors.New("i'm a teapot"), http.StatusTeapot)}
		})

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, httptest.NewRequest("GET", "https://example.com/", nil))
		assert.Equal(t, http.StatusTeapot, rec.Code)

		var resp errorResponse
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
		assert.Equal(t, "i'm a teapot", resp.Error)
		assert.Equal(t, http.StatusTeapot, resp.StatusCode)
	})

	t.Run("Template Error", func(t *testing.T) {
		root.Self = HandlerFunc(func(ctx context.Context, req *http.Request) Response {
			return Response{Template: "nonexistent"}
		})

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, httptest.NewRequest("GET", "https://example.com/", nil))
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var resp errorResponse
		fmt.Println(rec.Body.String())
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
		assert.Equal(t, `html/template: "nonexistent" is undefined`, resp.Error)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}

func TestRouterUse(t *testing.T) {
	r := NewRouter(func(ctx context.Context, url string) (Traversible, error) {
		return Node{Self: HandlerFunc(func(ctx context.Context, req *http.Request) Response {
			return Response{Data: "hi"}
		})}, nil
	})
	r.Use(func(next Handler) Handler {
		return HandlerFunc(func(ctx context.Context, req *http.Request) Response {
			resp := next.HandleRequest(ctx, req)
			resp.Status = http.StatusTeapot
			return resp
		})
	})
	r.Use(func(next Handler) Handler {
		return HandlerFunc(func(ctx context.Context, req *http.Request) Response {
			resp := next.HandleRequest(ctx, req)
			resp.Header().Set("X-My-Header", "value")
			return resp
		})
	})

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest("GET", "https://example.com/", nil))
	assert.Equal(t, http.StatusTeapot, rec.Code)
	assert.Equal(t, "value", rec.Header().Get("X-My-Header"))
	assert.Equal(t, "hi", rec.Body.String())
}
