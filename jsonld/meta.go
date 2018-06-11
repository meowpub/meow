package jsonld

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

var metaType = reflect.TypeOf(Meta{})

type Meta struct {
	Extra map[string]json.RawMessage `json:"-"`
}

func Unmarshal(obj, target interface{}) error {
	_, err := UnmarshalValue(obj, reflect.ValueOf(target))
	return err
}

// UnmarshalValue unmarshalls obj into rV, returning rV
// (the returned value of rV may differ from the passed in value, for example
// if it is an array)
func UnmarshalValue(obj interface{}, rV reflect.Value) (reflect.Value, error) {
	rV = reflect.Indirect(rV)
	rT := rV.Type()

	switch rT.Kind() {
	case reflect.Bool:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.String:
		rV.Set(reflect.ValueOf(obj).Convert(rT))

	case reflect.Slice:
		arr, ok := obj.([]interface{})
		if !ok {
			return rV, errors.New("Expected to decode an array but got a non-array")
		}

		for i := 0; i < len(arr); i++ {
			childV := reflect.New(rT.Elem())
			childV, err := UnmarshalValue(arr[i], childV)
			if err != nil {
				return rV, errors.Wrapf(err, "Element %d", i)
			}
			rV = reflect.Append(rV, childV)
		}

	case reflect.Map:
		map_, ok := obj.(map[string]interface{})
		if !ok {
			return rV, errors.New("Expected to decode a map but got something else")
		}

		for k, v := range map_ {
			childV := reflect.New(rT.Elem())
			childV, err := UnmarshalValue(v, childV)
			if err != nil {
				return rV, errors.Wrap(err, k)
			}
			rV.SetMapIndex(reflect.ValueOf(k), childV)
		}

	case reflect.Struct:
		fields, ok := obj.(map[string]interface{})
		if !ok {
			return rV, errors.New("Expected to decode a struct but got a non-Object")
		}

		m := findMeta(rV)
		if m == nil {
			return rV, errors.New("Could not find Meta to unmarshall struct")
		}

		err := m.unmarshalFields(fields, rV)
		if err != nil {
			return rV, err
		}

	default:
		panic(fmt.Sprintf("Don't know how to marshal a %s", rT.Kind()))
	}

	return rV, nil
}

func (m *Meta) unmarshalFields(fields map[string]interface{}, rV reflect.Value) error {
	rT := rV.Type()

	for i := 0; i < rT.NumField(); i++ {
		field := rT.Field(i)
		fieldV := rV.Field(i)

		// Skip the Meta embed.
		if field.Type == metaType {
			continue
		}

		// If this is an embedded field, descend into it using the same Meta.
		if field.Anonymous {
			if err := m.unmarshalFields(fields, fieldV); err != nil {
				return err
			}
			continue
		}

		// Horribly mangle the `json` tag, just like what encoding/json does.
		jField := parseJsonField(field.Name, field.Tag.Get("json"))
		if jField.Name == "" || jField.Name == "-" {
			continue
		}

		// Unmarshal the field if we have any data for it
		fieldData, ok := fields[jField.Name]
		if !ok {
			continue
		}

		v, err := UnmarshalValue(fieldData, fieldV)
		if err != nil {
			return errors.Wrapf(err, "Field %s", jField.Name)
		}
		fieldV.Set(v)
	}

	return nil
}

func Marshal(v interface{}) (interface{}, error) {
	return marshal(reflect.ValueOf(v))
}

func marshal(rV reflect.Value) (interface{}, error) {
	rT := rV.Type()

	switch rT.Kind() {
	case reflect.Bool:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.String:
		return rV.Interface(), nil

	case reflect.Slice:
		slice := make([]interface{}, rV.Len())
		for i := 0; i < rV.Len(); i++ {
			v, err := marshal(rV.Index(i))
			if err != nil {
				return nil, errors.Wrapf(err, "Item %d", i)
			} else {
				slice[i] = v
			}
		}
		return slice, nil

	case reflect.Map:
		m := make(map[string]interface{})
		keys := rV.MapKeys()
		for i := 0; i < len(keys); i++ {
			k := keys[i].String()
			v, err := marshal(rV.MapIndex(keys[i]))
			if err != nil {
				return nil, errors.Wrap(err, k)
			} else {
				m[k] = v
			}
		}
		return m, nil

	case reflect.Struct:
		m := findMeta(rV)
		if m == nil {
			return nil, errors.New("Trying to marshal struct but couldn't find Meta?")
		}

		map_ := make(map[string]interface{})
		err := m.marshalFields(rV, map_)
		if err != nil {
			return nil, err
		}
		return map_, nil

	case reflect.Ptr:
		return marshal(reflect.Indirect(rV))

	default:
		panic(fmt.Sprintf("Don't know how to marshal a %s", rT.Kind()))
	}

	return nil, nil
}

func (m *Meta) marshalFields(rV reflect.Value, fields map[string]interface{}) error {
	rT := rV.Type()

	for i := 0; i < rT.NumField(); i++ {
		field := rT.Field(i)
		fieldV := rV.Field(i)
		fieldT := fieldV.Type()

		// Skip the Meta embed.
		if field.Type == metaType {
			continue
		}

		// If this is an embedded field, descend into it using the same Meta.
		// and encode it into the same body
		if field.Anonymous {
			if err := m.marshalFields(fieldV, fields); err != nil {
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

		v, err := marshal(fieldV)
		if err != nil {
			return errors.Wrap(err, jField.Name)
		}

		fields[jField.Name] = v
	}

	return nil
}

func findMeta(rV reflect.Value) *Meta {
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
		} else if rF.Anonymous && rF.Type.Kind() == reflect.Struct {
			if child := findMeta(rV.Field(i)); child != nil {
				return child
			}
		}
	}
	return nil
}
