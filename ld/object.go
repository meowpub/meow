package ld

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/meowpub/meow/lib"
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

func NewObject(id string, types ...string) *Object {
	obj := &Object{}
	if id != "" {
		obj.SetID(id)
	}
	if types != nil {
		obj.SetType(types...)
	}
	return obj
}

func FetchObject(ctx context.Context, id string) (*Object, error) {
	rdoc, err := NewDocumentLoader(ctx).LoadDocument(id)
	if err != nil {
		return nil, lib.Code(err, http.StatusUnprocessableEntity)
	}
	v, ok := rdoc.Document.(map[string]interface{})
	if !ok {
		return nil, lib.Errorf(http.StatusInternalServerError, "can't cast %T to map[string]interface{}", v)
	}
	return &Object{V: v}, nil
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

func (obj *Object) IsNull() bool {
	return obj == nil
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

// Sets the object's @id.
func (obj *Object) SetID(id string) {
	obj.Set("@id", id)
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

// Sets the object's @value.
func (obj *Object) SetValue(v string) {
	obj.Set("@value", v)
}

// Returns the object's @type, or nil.
func (obj *Object) Type() []string {
	if obj == nil {
		return nil
	}
	if obj.typ == nil {
		obj.typ = ToStringSlice(obj.Get("@type"))
	}
	return obj.typ
}

// Sets the object's @type.
func (obj *Object) SetType(ts ...string) {
	obj.Set("@type", ts)
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
	if obj == nil {
		return
	}
	if obj.V == nil {
		obj.V = map[string]interface{}{key: value}
		return
	}
	switch v := value.(type) {
	case Entity:
		value = v.Obj().V
	case *Object:
		value = v.V
	}
	obj.V[key] = value
	switch key {
	case "@id":
		obj.id = ""
	case "@value":
		obj.value = ""
	case "@type":
		obj.typ = nil
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
			obj.V[k] = append(ToSlice(obj.Get(k)), arr...)
		} else {
			obj.V[k] = append(ToSlice(obj.Get(k)), v)
		}
	}
	return nil
}
