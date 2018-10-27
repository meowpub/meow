package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_resolveShortType(t *testing.T) {
	t.Run("as.Note", func(t *testing.T) {
		full, err := resolveShortType("as:Note")
		assert.NoError(t, err)
		assert.Equal(t, "http://www.w3.org/ns/activitystreams#Note", full)
	})
	t.Run("as2.Note", func(t *testing.T) {
		_, err := resolveShortType("as2:Note")
		assert.EqualError(t, err, "unknown namespace: as2")
	})
	t.Run("as.Notee", func(t *testing.T) {
		_, err := resolveShortType("as:Notee")
		assert.EqualError(t, err, "unknown type: http://www.w3.org/ns/activitystreams#Notee")
	})
}

func Test_parseTypedValue(t *testing.T) {
	testdata := map[string]map[string]interface{}{
		"a string": {"@value": "a string"},
	}
	for s, out := range testdata {
		t.Run(s, func(t *testing.T) {
			v, err := parseTypedValue(s)
			require.NoError(t, err)
			assert.Equal(t, out, v)
		})
	}
}
