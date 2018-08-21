// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdfs

import (
	"github.com/meowpub/meow/ld"
)

type Class struct{ O *ld.Object }

func (obj Class) Obj() *ld.Object            { return obj.O }
func (obj Class) ID() string                 { return obj.O.ID() }
func (obj Class) Value() string              { return obj.O.Value() }
func (obj Class) Type() []string             { return obj.O.Type() }
func (obj Class) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Class) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

func (obj Class) SubClassOf() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#subClassOf"]
}

type Container struct{ O *ld.Object }

func (obj Container) Obj() *ld.Object            { return obj.O }
func (obj Container) ID() string                 { return obj.O.ID() }
func (obj Container) Value() string              { return obj.O.Value() }
func (obj Container) Type() []string             { return obj.O.Type() }
func (obj Container) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Container) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type ContainerMembershipProperty struct{ O *ld.Object }

func (obj ContainerMembershipProperty) Obj() *ld.Object            { return obj.O }
func (obj ContainerMembershipProperty) ID() string                 { return obj.O.ID() }
func (obj ContainerMembershipProperty) Value() string              { return obj.O.Value() }
func (obj ContainerMembershipProperty) Type() []string             { return obj.O.Type() }
func (obj ContainerMembershipProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj ContainerMembershipProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type Datatype struct{ O *ld.Object }

func (obj Datatype) Obj() *ld.Object            { return obj.O }
func (obj Datatype) ID() string                 { return obj.O.ID() }
func (obj Datatype) Value() string              { return obj.O.Value() }
func (obj Datatype) Type() []string             { return obj.O.Type() }
func (obj Datatype) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Datatype) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type Literal struct{ O *ld.Object }

func (obj Literal) Obj() *ld.Object            { return obj.O }
func (obj Literal) ID() string                 { return obj.O.ID() }
func (obj Literal) Value() string              { return obj.O.Value() }
func (obj Literal) Type() []string             { return obj.O.Type() }
func (obj Literal) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Literal) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type Resource struct{ O *ld.Object }

func (obj Resource) Obj() *ld.Object            { return obj.O }
func (obj Resource) ID() string                 { return obj.O.ID() }
func (obj Resource) Value() string              { return obj.O.Value() }
func (obj Resource) Type() []string             { return obj.O.Type() }
func (obj Resource) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Resource) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

func (obj Resource) Comment() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#comment"]
}

func (obj Resource) IsDefinedBy() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#isDefinedBy"]
}

func (obj Resource) Label() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#label"]
}

func (obj Resource) Member() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#member"]
}

func (obj Resource) SeeAlso() interface{} {
	return obj.O.V["http://www.w3.org/2000/01/rdf-schema#seeAlso"]
}

var (
	_ ld.Entity = Class{}
	_ ld.Entity = Container{}
	_ ld.Entity = ContainerMembershipProperty{}
	_ ld.Entity = Datatype{}
	_ ld.Entity = Literal{}
	_ ld.Entity = Resource{}
)
