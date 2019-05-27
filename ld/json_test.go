package ld

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const expanded = `
	{
		"@id": "https://example.com/something",
		"http://purl.org/dc/terms/creator": [
			{"@id": "https://example.com/someBODY"}
		],
		"https://w3id.org/security#Ed25519Signature2018": [
			{"@value": "foo"}
		]
	}
	`

const compacted = `
	{
		"@context":"https://w3id.org/security/v1",
		"Ed25519Signature2018":"foo",
		"creator":"someBODY",
		"id":"something"
	}
	`

func TestDocumentLoader(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "GET", req.Method, "Wrong Method")
		assert.Equal(t, "/something", req.RequestURI, "Wrong RequestURI")
		assert.Equal(t, acceptHeader, req.Header.Get("Accept"), "Wrong Accept header")
		fmt.Fprintln(rw, expanded)
	}))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	loader := NewDocumentLoader(ctx)
	rdoc, err := loader.LoadDocument(srv.URL + "/something")
	require.NoError(t, err)
	assert.Equal(t, srv.URL+"/something", rdoc.DocumentURL, "Wrong DocumentURL")
	assert.Equal(t, map[string]interface{}{
		"@id": "https://example.com/something",
		"http://purl.org/dc/terms/creator": []interface{}{
			map[string]interface{}{"@id": "https://example.com/someBODY"},
		},
		"https://w3id.org/security#Ed25519Signature2018": []interface{}{
			map[string]interface{}{"@value": "foo"},
		},
	}, rdoc.Document, "Wrong Document")

	cancel()

	_, err = loader.LoadDocument(srv.URL + "/something")
	assert.EqualError(t, err, "Get "+srv.URL+"/something: context canceled", "Wrong error from cancelled context")
}

func TestCompact(t *testing.T) {
	var compBuf bytes.Buffer
	require.NoError(t, json.Compact(&compBuf, []byte(compacted)), "reference must json compact")

	var v map[string]interface{}
	require.NoError(t, json.Unmarshal([]byte(expanded), &v), "json.Unmarshal")

	// We  can use a nil http.Client as no fetches should actually be done
	res, err := Compact(nil, v, "https://example.com/something", "https://w3id.org/security/v1")
	require.NoError(t, err, "Compact should succeed")

	buf, err := json.Marshal(res)
	require.NoError(t, err, "json.Marshal")

	require.Equal(t, compBuf.String(), string(buf), "should produce expected result")
}

func TestExpand(t *testing.T) {
	var expBuf bytes.Buffer
	require.NoError(t, json.Compact(&expBuf, []byte(expanded)), "reference must compact")

	var v map[string]interface{}
	require.NoError(t, json.Unmarshal([]byte(compacted), &v), "json.Unmarshal")

	res, err := Expand(nil, v, "https://example.com/something")
	require.NoError(t, err, "Expand should succeed")

	buf, err := json.Marshal(res)
	require.NoError(t, err, "json.Marshal")

	require.Equal(t, expBuf.String(), string(buf), "should produce expected result")
}
