package jsonld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jsonKey(t *testing.T) {
	assert.Equal(t, "FieldName", jsonKey("FieldName", ""))
	assert.Equal(t, "field_name", jsonKey("FieldName", "field_name"))
	assert.Equal(t, "field_name", jsonKey("FieldName", "field_name,omitempty"))
}
