package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameOrigin(t *testing.T) {
	assert.True(t, IsSameOrigin("https://example.com/~a", "https://example.com/~b"))
	assert.True(t, IsSameOrigin("https://example.com/~a", "http://example.com/~b"))
	assert.False(t, IsSameOrigin("https://example.com/~a", "http://example.co.uk/~b"))
}

func TestSlugify(t *testing.T) {
	testdata := map[string]string{
		"i'm gay":      "im-gay",
		"lovely cafés": "lovely-cafes",
		"According to all known Laws of Aviation,": "according-to-all-known-laws-of-aviation",
	}
	for input, expected := range testdata {
		t.Run(input, func(t *testing.T) {
			output, err := Slugify(input)
			assert.NoError(t, err)
			assert.Equal(t, expected, output)
		})
	}
}
