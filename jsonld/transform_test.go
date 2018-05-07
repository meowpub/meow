package jsonld

import (
	"bytes"
	"encoding/json"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
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
	require.NoError(t, json.Compact(&compBuf, []byte(compacted)), "reference must compact")

	// We  can use a nil http.Client as no fetches should actually be done
	res, err := Compact(nil, []byte(expanded), "https://example.com/something", "https://w3id.org/security/v1")

	require.NoError(t, err, "Compact should succeed")
	require.Equal(t, compBuf.String(), string(res), "should produce expected result")
}

func TestExpand(t *testing.T) {
	var expBuf bytes.Buffer
	require.NoError(t, json.Compact(&expBuf, []byte(expanded)), "reference must compact")

	res, err := Expand(nil, []byte(compacted), "https://example.com/something")

	require.NoError(t, err, "Expand should succeed")
	require.Equal(t, expBuf.String(), string(res), "should produce expected result")
}
