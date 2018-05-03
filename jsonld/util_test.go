package jsonld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseJsonField(t *testing.T) {
	assert.Equal(t, jsonField{"FieldName", false}, parseJsonField("FieldName", ""))
	assert.Equal(t, jsonField{"field_name", false}, parseJsonField("FieldName", "field_name"))
	assert.Equal(t, jsonField{"field_name", true}, parseJsonField("FieldName", "field_name,omitempty"))
}
