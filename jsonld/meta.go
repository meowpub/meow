package jsonld

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

var metaType = reflect.TypeOf(Meta{})

type Meta struct {
	Extra map[string]json.RawMessage `json:"-"`
}

func (m Meta) UnmarshalJSON(data []byte) error {
	return errors.New("if you embed Meta, you must override UnmarshalJSON and have it call Meta.Unmarshal")
}

func (m Meta) MarshalJSON() ([]byte, error) {
	return nil, errors.New("if you embed Meta, you must override MarshalJSON and have it call Meta.Marshal")
}

func (m *Meta) Unmarshal(data []byte, v interface{}) error {
	return m.unmarshal(data, reflect.ValueOf(v))
}

func (m *Meta) unmarshal(data json.RawMessage, rV reflect.Value) error {
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

		// Grab the Meta and skip if we have one.
		if field.Type == metaType {
			meta = rV.Field(i).Addr().Interface().(*Meta)
			continue
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
		if err := m.unmarshal(fieldData, rV.Field(i)); err != nil {
			return errors.Wrap(err, key)
		}
		delete(fields, key)
	}

	if meta != nil {
		// Set it to nil if there are no leftover fields, to save some memory and make test cases
		// simpler to write.
		if len(fields) == 0 {
			fields = nil
		}
		m.Extra = fields
	}

	return nil
}

func (m *Meta) Marshal(v interface{}) ([]byte, error) {
	return nil, nil
}
