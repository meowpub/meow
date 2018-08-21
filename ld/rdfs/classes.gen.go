package rdfs

import (
	"github.com/meowpub/meow/ld"
)

type Class struct {
	*ld.Object
}

func (obj Class) Obj() *ld.Object {
	return obj.Object
}

func (obj Class) SubClassOf() interface{} {
	return obj.V["http://www.w3.org/2000/01/rdf-schema#subClassOf"]
}

type Datatype struct {
	*ld.Object
}

func (obj Datatype) Obj() *ld.Object {
	return obj.Object
}

type Resource struct {
	*ld.Object
}

func (obj Resource) Obj() *ld.Object {
	return obj.Object
}

func (obj Resource) Comment() interface{} {
	return obj.V["http://www.w3.org/2000/01/rdf-schema#comment"]
}

func (obj Resource) SeeAlso() interface{} {
	return obj.V["http://www.w3.org/2000/01/rdf-schema#seeAlso"]
}

func (obj Resource) Member() interface{} {
	return obj.V["http://www.w3.org/2000/01/rdf-schema#member"]
}

func (obj Resource) Label() interface{} {
	return obj.V["http://www.w3.org/2000/01/rdf-schema#label"]
}

func (obj Resource) IsDefinedBy() interface{} {
	return obj.V["http://www.w3.org/2000/01/rdf-schema#isDefinedBy"]
}

type Container struct {
	*ld.Object
}

func (obj Container) Obj() *ld.Object {
	return obj.Object
}

type ContainerMembershipProperty struct {
	*ld.Object
}

func (obj ContainerMembershipProperty) Obj() *ld.Object {
	return obj.Object
}

type Literal struct {
	*ld.Object
}

func (obj Literal) Obj() *ld.Object {
	return obj.Object
}

var (
	_ ld.Entity = Class{}
	_ ld.Entity = Datatype{}
	_ ld.Entity = Resource{}
	_ ld.Entity = Container{}
	_ ld.Entity = ContainerMembershipProperty{}
	_ ld.Entity = Literal{}
)
