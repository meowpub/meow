package jsonld

import (
	"reflect"
	"strings"
)

type jsonField struct {
	Name      string
	OmitEmpty bool
}

func parseJsonField(fieldName, tag string) jsonField {
	field := jsonField{
		Name:      fieldName,
		OmitEmpty: false,
	}

	if tag != "" {
		parts := strings.Split(tag, ",")
		field.Name = parts[0]
		for i := 1; i < len(parts); i++ {
			if parts[i] == "omitempty" {
				field.OmitEmpty = true
			}
		}
	}

	return field
}

// Stolen from encoding/json
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
