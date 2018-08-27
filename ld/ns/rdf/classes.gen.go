// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

import (
	"github.com/meowpub/meow/ld"
)

// The class of containers of alternatives.
type Alt struct{ Container }

func AsAlt(obj *ld.Object) Alt {
	return Alt{AsContainer(obj)}
}

// The class of unordered containers.
type Bag struct{ Container }

func AsBag(obj *ld.Object) Bag {
	return Bag{AsContainer(obj)}
}

// The class of RDF Lists.
type List struct{ Resource }

func AsList(obj *ld.Object) List {
	return List{AsResource(obj)}
}

// The first item in the subject RDF list.
func (obj List) First() interface{} {
	return obj.Get("http://www.w3.org/1999/02/22-rdf-syntax-ns#first")
}

// The rest of the subject RDF list after the first item.
func (obj List) Rest() interface{} {
	return obj.Get("http://www.w3.org/1999/02/22-rdf-syntax-ns#rest")
}

// The class of RDF properties.
type Property struct{ Resource }

func AsProperty(obj *ld.Object) Property {
	return Property{AsResource(obj)}
}

// A domain of the subject property.
func (obj Property) Domain() interface{} {
	return obj.Get("http://www.w3.org/2000/01/rdf-schema#domain")
}

// A range of the subject property.
func (obj Property) Range() interface{} {
	return obj.Get("http://www.w3.org/2000/01/rdf-schema#range")
}

// The subject is a subproperty of a property.
func (obj Property) SubPropertyOf() interface{} {
	return obj.Get("http://www.w3.org/2000/01/rdf-schema#subPropertyOf")
}

// The class of ordered containers.
type Seq struct{ Container }

func AsSeq(obj *ld.Object) Seq {
	return Seq{AsContainer(obj)}
}

// The class of RDF statements.
type Statement struct{ Resource }

func AsStatement(obj *ld.Object) Statement {
	return Statement{AsResource(obj)}
}

// The object of the subject RDF statement.
func (obj Statement) Object() interface{} {
	return obj.Get("http://www.w3.org/1999/02/22-rdf-syntax-ns#object")
}

// The predicate of the subject RDF statement.
func (obj Statement) Predicate() interface{} {
	return obj.Get("http://www.w3.org/1999/02/22-rdf-syntax-ns#predicate")
}

// The subject of the subject RDF statement.
func (obj Statement) Subject() interface{} {
	return obj.Get("http://www.w3.org/1999/02/22-rdf-syntax-ns#subject")
}

// The class of classes.
type Class struct{ Resource }

func AsClass(obj *ld.Object) Class {
	return Class{AsResource(obj)}
}

// The subject is a subclass of a class.
func (obj Class) SubClassOf() interface{} {
	return obj.Get("http://www.w3.org/2000/01/rdf-schema#subClassOf")
}

// The class of RDF containers.
type Container struct{ Resource }

func AsContainer(obj *ld.Object) Container {
	return Container{AsResource(obj)}
}

// The class of container membership properties, rdf:_1, rdf:_2, ...,
// all of which are sub-properties of 'member'.
type ContainerMembershipProperty struct{ Property }

func AsContainerMembershipProperty(obj *ld.Object) ContainerMembershipProperty {
	return ContainerMembershipProperty{AsProperty(obj)}
}

// The class of RDF datatypes.
type Datatype struct{ Class }

func AsDatatype(obj *ld.Object) Datatype {
	return Datatype{AsClass(obj)}
}

// The class of literal values, eg. textual strings and integers.
type Literal struct{ Resource }

func AsLiteral(obj *ld.Object) Literal {
	return Literal{AsResource(obj)}
}

// The class resource, everything.
type Resource struct{ o *ld.Object }

func AsResource(obj *ld.Object) Resource {
	return Resource{o: obj}
}

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Resource) Obj() *ld.Object {
	return obj.o
}

// Returns the object's @id. Implements ld.Entity.
func (obj Resource) ID() string {
	return obj.o.ID()
}

// Returns the object's @value. Implements ld.Entity.
func (obj Resource) Value() string {
	return obj.o.Value()
}

// Returns the object's @type. Implements ld.Entity.
func (obj Resource) Type() []string {
	return obj.o.Type()
}

// Returns the named attribute. Implements ld.Entity.
func (obj Resource) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Resource) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// The subject is an instance of a class.
func (obj Resource) Type_() interface{} {
	return obj.Get("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
}

// Idiomatic property used for structured values.
func (obj Resource) Value_() interface{} {
	return obj.Get("http://www.w3.org/1999/02/22-rdf-syntax-ns#value")
}

// A description of the subject resource.
func (obj Resource) Comment() interface{} {
	return obj.Get("http://www.w3.org/2000/01/rdf-schema#comment")
}

// The defininition of the subject resource.
func (obj Resource) IsDefinedBy() interface{} {
	return obj.Get("http://www.w3.org/2000/01/rdf-schema#isDefinedBy")
}

// A human-readable name for the subject.
func (obj Resource) Label() interface{} {
	return obj.Get("http://www.w3.org/2000/01/rdf-schema#label")
}

// A member of the subject resource.
func (obj Resource) Member() interface{} {
	return obj.Get("http://www.w3.org/2000/01/rdf-schema#member")
}

// Further information about the subject resource.
func (obj Resource) SeeAlso() interface{} {
	return obj.Get("http://www.w3.org/2000/01/rdf-schema#seeAlso")
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
