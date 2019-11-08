package ld

import (
	"fmt"
)

func ToString(raw interface{}) string {
	switch v := raw.(type) {
	case nil:
		return ""
	case string:
		return v
	case []interface{}:
		out := ""
		for i, vv := range v {
			s := ToString(vv)
			if i == 0 {
				out = s
			} else {
				out = out + "," + s
			}
		}
		return out
	case map[string]interface{}:
		if val, ok := v["@value"]; ok {
			return ToString(val)
		}
		return fmt.Sprint(v)
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprint(v)
	}
}

// Casts to a slice.
// If the value is a slice, it's returned verbatim.
// If not, it's returned wrapped in a slice with one item.
func ToSlice(raw interface{}) []interface{} {
	switch v := raw.(type) {
	case nil:
		return nil
	case []interface{}:
		return v
	case []string:
		vs := make([]interface{}, len(v))
		for i, vv := range v {
			vs[i] = vv
		}
		return vs
	default:
		return []interface{}{raw}
	}
}

func ToStringSlice(raw interface{}) []string {
	switch v := raw.(type) {
	case []string:
		return v
	default:
		arr := ToSlice(raw)
		if arr == nil {
			return nil
		}
		strs := make([]string, len(arr))
		for i, v := range arr {
			strs[i] = ToString(v)
		}
		return strs
	}
}

func ToObject(raw interface{}) *Object {
	switch v := raw.(type) {
	case nil:
		return nil
	case map[string]interface{}:
		return &Object{V: v}
	case []interface{}:
		for _, item := range v {
			if obj := ToObject(item); obj != nil {
				return obj
			}
		}
		return nil
	default:
		return nil
	}
}

func ToObjects(raw interface{}) []*Object {
	switch v := raw.(type) {
	case nil:
		return nil
	case map[string]interface{}:
		return []*Object{&Object{V: v}}
	case []interface{}:
		var objs []*Object
		for _, item := range v {
			if obj := ToObject(item); obj != nil {
				objs = append(objs, obj)
			}
		}
		return objs
	default:
		return nil
	}
}
