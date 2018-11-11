package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_varsafe(t *testing.T) {
	assert.Equal(t, varsafe("http://example.com#Type"), "http_example_com_Type")
}
