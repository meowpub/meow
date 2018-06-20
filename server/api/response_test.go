package api

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedirectResponse(t *testing.T) {
	testdata := map[string]struct {
		To string
		Q  []url.Values

		Location string
	}{
		"Path": {"/-/login", nil, "/-/login"},
		"Query": {"/-/login", []url.Values{
			{"a": []string{"1"}},
		}, "/-/login?a=1"},
		"Queries": {"/-/login", []url.Values{
			{"a": []string{"1"}},
			{"b": []string{"2"}},
		}, "/-/login?a=1&b=2"},
		"Merge": {"/-/login", []url.Values{
			{"a": []string{"1", "2"}},
			{"a": []string{"3"}, "b": []string{"b"}},
			{"b": []string{"bb"}, "a": []string{"4"}, "c": []string{"hi"}},
		}, "/-/login?a=1&a=2&a=3&a=4&b=b&b=bb&c=hi"},
	}
	for name, tdata := range testdata {
		t.Run(name, func(t *testing.T) {
			resp := RedirectResponse(tdata.To, tdata.Q...)
			assert.Nil(t, resp.Data)
			assert.NoError(t, resp.Error)
			assert.Empty(t, resp.Template)
			assert.Equal(t, http.StatusTemporaryRedirect, resp.Status)
			assert.Equal(t, http.Header{
				"Location": []string{tdata.Location},
			}, resp.Header())
		})
	}
}
