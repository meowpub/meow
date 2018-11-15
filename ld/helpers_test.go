package ld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	obj := NewObject("https://example.com", "http://www.w3.org/ns/activitystreams#Note")
	assert.True(t, Is(obj, "http://www.w3.org/ns/activitystreams#Note"))
	assert.False(t, Is(obj, "http://www.w3.org/ns/activitystreams#Image"))
}
