// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

import (
	"github.com/meowpub/meow/ld"
)

// The class of containers of alternatives.
type Alt struct{ O *ld.Object }

func (obj Alt) Obj() *ld.Object            { return obj.O }
func (obj Alt) ID() string                 { return obj.O.ID() }
func (obj Alt) Value() string              { return obj.O.Value() }
func (obj Alt) Type() []string             { return obj.O.Type() }
func (obj Alt) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Alt) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of unordered containers.
type Bag struct{ O *ld.Object }

func (obj Bag) Obj() *ld.Object            { return obj.O }
func (obj Bag) ID() string                 { return obj.O.ID() }
func (obj Bag) Value() string              { return obj.O.Value() }
func (obj Bag) Type() []string             { return obj.O.Type() }
func (obj Bag) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Bag) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of RDF Lists.
type List struct{ O *ld.Object }

func (obj List) Obj() *ld.Object            { return obj.O }
func (obj List) ID() string                 { return obj.O.ID() }
func (obj List) Value() string              { return obj.O.Value() }
func (obj List) Type() []string             { return obj.O.Type() }
func (obj List) Get(key string) interface{} { return obj.O.Get(key) }

func (obj List) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The first item in the subject RDF list.
func (obj List) First() interface{} {
	return obj.O.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#first"]
}

// The rest of the subject RDF list after the first item.
func (obj List) Rest() interface{} {
	return obj.O.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#rest"]
}

// The class of RDF properties.
type Property struct{ O *ld.Object }

func (obj Property) Obj() *ld.Object            { return obj.O }
func (obj Property) ID() string                 { return obj.O.ID() }
func (obj Property) Value() string              { return obj.O.Value() }
func (obj Property) Type() []string             { return obj.O.Type() }
func (obj Property) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Property) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of ordered containers.
type Seq struct{ O *ld.Object }

func (obj Seq) Obj() *ld.Object            { return obj.O }
func (obj Seq) ID() string                 { return obj.O.ID() }
func (obj Seq) Value() string              { return obj.O.Value() }
func (obj Seq) Type() []string             { return obj.O.Type() }
func (obj Seq) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Seq) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of RDF statements.
type Statement struct{ O *ld.Object }

func (obj Statement) Obj() *ld.Object            { return obj.O }
func (obj Statement) ID() string                 { return obj.O.ID() }
func (obj Statement) Value() string              { return obj.O.Value() }
func (obj Statement) Type() []string             { return obj.O.Type() }
func (obj Statement) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Statement) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The object of the subject RDF statement.
func (obj Statement) Object() interface{} {
	return obj.O.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#object"]
}

// The predicate of the subject RDF statement.
func (obj Statement) Predicate() interface{} {
	return obj.O.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#predicate"]
}

// The subject of the subject RDF statement.
func (obj Statement) Subject() interface{} {
	return obj.O.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#subject"]
}

var (
	_ ld.Entity = Alt{}
	_ ld.Entity = Bag{}
	_ ld.Entity = List{}
	_ ld.Entity = Property{}
	_ ld.Entity = Seq{}
	_ ld.Entity = Statement{}
)
