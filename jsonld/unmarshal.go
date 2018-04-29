package jsonld

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

var metaType = reflect.TypeOf(Meta{})

func Unmarshal(data []byte, v interface{}) error {
	return unmarshal(data, reflect.ValueOf(v))
}

func unmarshal(data json.RawMessage, rV reflect.Value) error {
	rV = reflect.Indirect(rV)
	rT := rV.Type()

	// Use regular ol' encoding/json unmarshalling for non-structs.
	if rT.Kind() != reflect.Struct {
		return json.Unmarshal(data, rV.Addr().Interface())
	}

	// Unmarshal all of the fields into RawMessage objects.
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var meta *Meta
	for i := 0; i < rT.NumField(); i++ {
		field := rT.Field(i)

		// Find the Meta embed if there is one, we put unhandled fields in there for round trips.
		if field.Type == metaType {
			meta = rV.Field(i).Addr().Interface().(*Meta)
		}

		// Horribly mangle the `json` tag, just like what encoding/json does.
		key := field.Name
		if tag, ok := field.Tag.Lookup("json"); ok {
			if idx := strings.Index(tag, ","); idx != -1 {
				tag = tag[:idx]
			}
			key = tag
		}

		// Unmarshal the field if we have any data for it.
		fieldData, ok := fields[key]
		if !ok {
			continue
		}
		if err := unmarshal(fieldData, rV.Field(i)); err != nil {
			return errors.Wrap(err, key)
		}
	}

	if meta != nil {
		meta.Extra = fields
	}

	return nil
}
