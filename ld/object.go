package ld

import (
	"encoding/json"
)

var _ Entity = &Object{}

// The base type of all JSON-LD objects.
type Object struct {
	V map[string]interface{}

	// Cached attributes.
	id    string
	value string
	typ   []string
}

// Creates a new object from a JSON object.
func ParseObject(data []byte) (*Object, error) {
	var obj Object
	return &obj, json.Unmarshal(data, &obj.V)
}

// Creates a new list of objects from a JSON array.
func ParseObjects(data []byte) ([]*Object, error) {
	var vs []map[string]interface{}
	if err := json.Unmarshal(data, &vs); err != nil {
		return nil, err
	}
	objs := make([]*Object, len(vs))
	for i, v := range vs {
		objs[i] = &Object{V: v}
	}
	return objs, nil
}

func (obj Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(obj.V)
}

func (obj *Object) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	obj.V = v
	return nil
}

// Revolver Ocelot.
func (obj *Object) Obj() *Object {
	return obj
}

// Returns the object's @id, or "".
func (obj *Object) ID() string {
	if obj == nil {
		return ""
	}
	if obj.id == "" {
		obj.id = ToString(obj.Get("@id"))
	}
	return obj.id
}

// Returns the object's @value, or "".
func (obj *Object) Value() string {
	if obj == nil {
		return ""
	}
	if obj.value == "" {
		obj.value = ToString(obj.Get("@value"))
	}
	return obj.value
}

// Returns the object's @type, or nil.
func (obj *Object) Type() []string {
	if obj == nil {
		return nil
	}
	if obj.typ == nil {
		obj.typ = ToStringArray(obj.Get("@type"))
	}
	return obj.typ
}

// Nil-safe getter for attributes.
func (obj *Object) Get(key string) interface{} {
	if obj != nil && obj.V != nil {
		return obj.V[key]
	}
	return nil
}

// Nil-safe setter for attributes.
func (obj *Object) Set(key string, value interface{}) {
	if obj != nil {
		obj.V[key] = value
	}
}

// Applies another object as a patch to this object. The mergeArrays argument specifies whether
// values for existing keys should replace the existing value, or be added as an additional value.
func (obj *Object) Apply(patch Entity, mergeArrays bool) error {
	patchV := patch.Obj().V
	for k, v := range patchV {
		if k == "@id" || !mergeArrays {
			obj.V[k] = v
		} else if arr, ok := v.([]interface{}); ok {
			obj.V[k] = append(ToArray(obj.Get(k)), arr...)
		} else {
			obj.V[k] = append(ToArray(obj.Get(k)), v)
		}
	}
	return nil
}
