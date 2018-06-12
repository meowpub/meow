package jsonld

import (
	"bytes"
	"encoding/json"

	// "github.com/stretchr/testify/assert"
	"testing"

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
