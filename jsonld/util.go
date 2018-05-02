package jsonld

import (
	"strings"
)

// jsonKey returns the JSON key for a struct field, given a field name and a `json:"..."` tag.
func jsonKey(fieldName, tag string) string {
	if tag != "" {
		if idx := strings.Index(tag, ","); idx != -1 {
			return tag[:idx]
		}
		return tag
	}
	return fieldName
}
