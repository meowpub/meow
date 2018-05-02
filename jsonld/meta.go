package jsonld

import (
	"bytes"
	"encoding/json"
	"reflect"

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
	rV := reflect.ValueOf(v)
	if m.findMeta(rV) != m {
		return errors.Errorf("Meta.Unmarshal called on a struct other than the one containing this Meta (of type '%s')",
			rV.Type().String())
	}
	return m.unmarshal(data, rV)
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

	for i := 0; i < rT.NumField(); i++ {
		field := rT.Field(i)
		fieldV := rV.Field(i)

		// Skip the Meta embed.
		if field.Type == metaType {
			continue
		}

		// If this is an embedded field, descend into it using the same Meta.
		if field.Anonymous {
			if err := m.unmarshal(data, fieldV); err != nil {
				return err
			}
			continue
		}

		// Horribly mangle the `json` tag, just like what encoding/json does.
		key := jsonKey(field.Name, field.Tag.Get("json"))
		if key == "" || key == "-" {
			continue
		}

		// Unmarshal the field if we have any data for it; using Meta.Unmarshal if it has a Meta,
		// otherwise plain encoding/json-style.
		fieldData, ok := fields[key]
		if !ok {
			continue
		}
		if meta := m.findMeta(fieldV); meta != nil {
			if err := meta.unmarshal(fieldData, fieldV); err != nil {
				return errors.Wrap(err, key)
			}
		} else {
			if err := json.Unmarshal(fieldData, fieldV.Addr().Interface()); err != nil {
				return errors.Wrap(err, key)
			}
		}
	}

	return nil
}

func (m *Meta) Marshal(v interface{}) ([]byte, error) {
	rV := reflect.ValueOf(v)
	if m.findMeta(rV) != m {
		return nil, errors.Errorf(
			"Meta.Marshal called on a struct other than the one containing this Meta (of type '%s')",
			rV.Type().Name())
	}

	var buf bytes.Buffer

	if err := m.marshal(reflect.Indirect(rV), &buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (m *Meta) marshal(rV reflect.Value, buf *bytes.Buffer) error {
	// rV = reflect.Indirect(rV)
	rT := rV.Type()

	enc := json.NewEncoder(buf)

	// Use regular ol' encoding/json marshalling for non-structs.
	if rT.Kind() != reflect.Struct {
		return enc.Encode(rV.Addr().Interface())
	}

	buf.WriteString("{")
	if err := m.marshalBody(rV, buf); err != nil {
		return err
	}
	buf.WriteString("}")

	return nil
}

func (m *Meta) marshalBody(rV reflect.Value, buf *bytes.Buffer) error {
	rT := rV.Type()
	first := true

	for i := 0; i < rT.NumField(); i++ {
		field := rT.Field(i)
		fieldV := rV.Field(i)
		fieldT := fieldV.Type()
		enc := json.NewEncoder(buf)

		// Skip the Meta embed.
		if field.Type == metaType {
			continue
		}

		// If this is an embedded field, descend into it using the same Meta.
		// and encode it into the same body
		if field.Anonymous {
			if err := m.marshalBody(fieldV, buf); err != nil {
				return err
			}
			continue
		}

		// Horribly mangle the `json` tag, just like what encoding/json does.
		jField := parseJsonField(field.Name, field.Tag.Get("json"))
		if jField.Name == "" || jField.Name == "-" {
			continue
		}

		// Marshall the field if we have any data for it

		// Indirect through pointers skipping nulls,
		if fieldT.Kind() == reflect.Ptr {
			if fieldV.IsNil() {
				continue
			} else {
				fieldT = fieldT.Elem()
				fieldV = reflect.Indirect(fieldV)
			}
		}

		if jField.OmitEmpty && isEmptyValue(fieldV) {
			continue
		}

		if !first {
			buf.WriteString(",")
		} else {
			first = false
		}

		if err := enc.Encode(jField.Name); err != nil {
			return errors.Wrap(err, jField.Name)
		}
		buf.WriteString(":")

		// If not a struct, marshall normally
		if fieldT.Kind() != reflect.Struct {
			if err := enc.Encode(fieldV.Addr().Interface()); err != nil {
				return errors.Wrap(err, jField.Name)
			}
			continue
		}

		// Else, marshall it as a Meta struct if it posesses one
		if meta := m.findMeta(fieldV); meta != nil {
			if err := meta.marshal(fieldV, buf); err != nil {
				return errors.Wrap(err, jField.Name)
			}
		} else {
			// Else, just encode it
			if err := enc.Encode(fieldV.Addr().Interface()); err != nil {
				return errors.Wrap(err, jField.Name)
			}
		}
	}

	return nil
}

func (m *Meta) findMeta(rV reflect.Value) *Meta {
	rT := rV.Type()

	if rT.Kind() == reflect.Ptr {
		rV = reflect.Indirect(rV)
		rT = rV.Type()
	}

	if rT.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < rT.NumField(); i++ {
		rF := rT.Field(i)
		if rF.Type == metaType {
			return rV.Field(i).Addr().Interface().(*Meta)
		}
	}
	return nil
}
