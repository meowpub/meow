package ld

// An interface for any kind of entity which embeds Object.
type Entity interface {
	Obj() *Object
	IsNull() bool
	ID() string
	Value() string
	Type() []string
	Get(key string) interface{}
	Set(key string, v interface{})
	Apply(other Entity, mergeArrays bool) error
}
