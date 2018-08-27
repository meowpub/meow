// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

import (
	"github.com/meowpub/meow/ld"
)

// The class of containers of alternatives.
type Alt struct{ Container }

func (obj Alt) Obj() *ld.Object            { return obj.O }
func (obj Alt) ID() string                 { return obj.O.ID() }
func (obj Alt) Value() string              { return obj.O.Value() }
func (obj Alt) Type() []string             { return obj.O.Type() }
func (obj Alt) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Alt) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of unordered containers.
type Bag struct{ Container }

func (obj Bag) Obj() *ld.Object            { return obj.O }
func (obj Bag) ID() string                 { return obj.O.ID() }
func (obj Bag) Value() string              { return obj.O.Value() }
func (obj Bag) Type() []string             { return obj.O.Type() }
func (obj Bag) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Bag) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of RDF Lists.
type List struct{ Resource }

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
type Property struct{ Resource }

func (obj Property) Obj() *ld.Object            { return obj.O }
func (obj Property) ID() string                 { return obj.O.ID() }
func (obj Property) Value() string              { return obj.O.Value() }
func (obj Property) Type() []string             { return obj.O.Type() }
func (obj Property) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Property) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// A domain of the subject property.
func (obj Property) Domain() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#domain"]
}

// A range of the subject property.
func (obj Property) Range() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#range"]
}

// The subject is a subproperty of a property.
func (obj Property) SubPropertyOf() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#subPropertyOf"]
}

// The class of ordered containers.
type Seq struct{ Container }

func (obj Seq) Obj() *ld.Object            { return obj.O }
func (obj Seq) ID() string                 { return obj.O.ID() }
func (obj Seq) Value() string              { return obj.O.Value() }
func (obj Seq) Type() []string             { return obj.O.Type() }
func (obj Seq) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Seq) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of RDF statements.
type Statement struct{ Resource }

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

// The class of classes.
type Class struct{ Resource }

func (obj Class) Obj() *ld.Object            { return obj.O }
func (obj Class) ID() string                 { return obj.O.ID() }
func (obj Class) Value() string              { return obj.O.Value() }
func (obj Class) Type() []string             { return obj.O.Type() }
func (obj Class) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Class) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The subject is a subclass of a class.
func (obj Class) SubClassOf() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#subClassOf"]
}

// The class of RDF containers.
type Container struct{ Resource }

func (obj Container) Obj() *ld.Object            { return obj.O }
func (obj Container) ID() string                 { return obj.O.ID() }
func (obj Container) Value() string              { return obj.O.Value() }
func (obj Container) Type() []string             { return obj.O.Type() }
func (obj Container) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Container) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of container membership properties, rdf:_1, rdf:_2, ...,
// all of which are sub-properties of 'member'.
type ContainerMembershipProperty struct{ Property }

func (obj ContainerMembershipProperty) Obj() *ld.Object            { return obj.O }
func (obj ContainerMembershipProperty) ID() string                 { return obj.O.ID() }
func (obj ContainerMembershipProperty) Value() string              { return obj.O.Value() }
func (obj ContainerMembershipProperty) Type() []string             { return obj.O.Type() }
func (obj ContainerMembershipProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj ContainerMembershipProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of RDF datatypes.
type Datatype struct{ Class }

func (obj Datatype) Obj() *ld.Object            { return obj.O }
func (obj Datatype) ID() string                 { return obj.O.ID() }
func (obj Datatype) Value() string              { return obj.O.Value() }
func (obj Datatype) Type() []string             { return obj.O.Type() }
func (obj Datatype) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Datatype) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of literal values, eg. textual strings and integers.
type Literal struct{ Resource }

func (obj Literal) Obj() *ld.Object            { return obj.O }
func (obj Literal) ID() string                 { return obj.O.ID() }
func (obj Literal) Value() string              { return obj.O.Value() }
func (obj Literal) Type() []string             { return obj.O.Type() }
func (obj Literal) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Literal) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class resource, everything.
type Resource struct{ O *ld.Object }

func (obj Resource) Obj() *ld.Object            { return obj.O }
func (obj Resource) ID() string                 { return obj.O.ID() }
func (obj Resource) Value() string              { return obj.O.Value() }
func (obj Resource) Type() []string             { return obj.O.Type() }
func (obj Resource) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Resource) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The subject is an instance of a class.
func (obj Resource) Type_() interface{} {
	return obj.O.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#type"]
}

// Idiomatic property used for structured values.
func (obj Resource) Value_() interface{} {
	return obj.O.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#value"]
}

// A description of the subject resource.
func (obj Resource) Comment() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#comment"]
}

// The defininition of the subject resource.
func (obj Resource) IsDefinedBy() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#isDefinedBy"]
}

// A human-readable name for the subject.
func (obj Resource) Label() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#label"]
}

// A member of the subject resource.
func (obj Resource) Member() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#member"]
}

// Further information about the subject resource.
func (obj Resource) SeeAlso() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#seeAlso"]
}

var (
	_ ld.Entity = Alt{}
	_ ld.Entity = Bag{}
	_ ld.Entity = List{}
	_ ld.Entity = Property{}
	_ ld.Entity = Seq{}
	_ ld.Entity = Statement{}
	_ ld.Entity = Class{}
	_ ld.Entity = Container{}
	_ ld.Entity = ContainerMembershipProperty{}
	_ ld.Entity = Datatype{}
	_ ld.Entity = Literal{}
	_ ld.Entity = Resource{}
)
