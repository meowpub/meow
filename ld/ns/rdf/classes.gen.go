// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

import (
	"github.com/meowpub/meow/ld"
)

// The class of containers of alternatives.
type Alt struct{ Container }

// Ducktypes the object into a(n) Alt.
func AsAlt(e ld.Entity) Alt { return Alt{AsContainer(e)} }

// Does the object quack like a(n) Alt?
func IsAlt(e ld.Entity) bool { return ld.Is(e, TypeAlt) }

// The class of unordered containers.
type Bag struct{ Container }

// Ducktypes the object into a(n) Bag.
func AsBag(e ld.Entity) Bag { return Bag{AsContainer(e)} }

// Does the object quack like a(n) Bag?
func IsBag(e ld.Entity) bool { return ld.Is(e, TypeBag) }

// The class of RDF Lists.
type List struct{ Resource }

// Ducktypes the object into a(n) List.
func AsList(e ld.Entity) List { return List{AsResource(e)} }

// Does the object quack like a(n) List?
func IsList(e ld.Entity) bool { return ld.Is(e, TypeList) }

// The first item in the subject RDF list.
func (obj List) First() interface{} { return GetFirst(obj) }

// The rest of the subject RDF list after the first item.
func (obj List) Rest() interface{} { return GetRest(obj) }

// The class of RDF properties.
type Property struct{ Resource }

// Ducktypes the object into a(n) Property.
func AsProperty(e ld.Entity) Property { return Property{AsResource(e)} }

// Does the object quack like a(n) Property?
func IsProperty(e ld.Entity) bool { return ld.Is(e, TypeProperty) }

// A domain of the subject property.
func (obj Property) Domain() interface{} { return GetDomain(obj) }

// A range of the subject property.
func (obj Property) Range() interface{} { return GetRange(obj) }

// The subject is a subproperty of a property.
func (obj Property) SubPropertyOf() interface{} { return GetSubPropertyOf(obj) }

// The class of ordered containers.
type Seq struct{ Container }

// Ducktypes the object into a(n) Seq.
func AsSeq(e ld.Entity) Seq { return Seq{AsContainer(e)} }

// Does the object quack like a(n) Seq?
func IsSeq(e ld.Entity) bool { return ld.Is(e, TypeSeq) }

// The class of RDF statements.
type Statement struct{ Resource }

// Ducktypes the object into a(n) Statement.
func AsStatement(e ld.Entity) Statement { return Statement{AsResource(e)} }

// Does the object quack like a(n) Statement?
func IsStatement(e ld.Entity) bool { return ld.Is(e, TypeStatement) }

// The object of the subject RDF statement.
func (obj Statement) Object() interface{} { return GetObject(obj) }

// The predicate of the subject RDF statement.
func (obj Statement) Predicate() interface{} { return GetPredicate(obj) }

// The subject of the subject RDF statement.
func (obj Statement) Subject() interface{} { return GetSubject(obj) }

// The class of classes.
type Class struct{ Resource }

// Ducktypes the object into a(n) Class.
func AsClass(e ld.Entity) Class { return Class{AsResource(e)} }

// Does the object quack like a(n) Class?
func IsClass(e ld.Entity) bool { return ld.Is(e, TypeClass) }

// The subject is a subclass of a class.
func (obj Class) SubClassOf() interface{} { return GetSubClassOf(obj) }

// The class of RDF containers.
type Container struct{ Resource }

// Ducktypes the object into a(n) Container.
func AsContainer(e ld.Entity) Container { return Container{AsResource(e)} }

// Does the object quack like a(n) Container?
func IsContainer(e ld.Entity) bool { return ld.Is(e, TypeContainer) }

// The class of container membership properties, rdf:_1, rdf:_2, ...,
// all of which are sub-properties of 'member'.
type ContainerMembershipProperty struct{ Property }

// Ducktypes the object into a(n) ContainerMembershipProperty.
func AsContainerMembershipProperty(e ld.Entity) ContainerMembershipProperty {
	return ContainerMembershipProperty{AsProperty(e)}
}

// Does the object quack like a(n) ContainerMembershipProperty?
func IsContainerMembershipProperty(e ld.Entity) bool { return ld.Is(e, TypeContainerMembershipProperty) }

// The class of RDF datatypes.
type Datatype struct{ Class }

// Ducktypes the object into a(n) Datatype.
func AsDatatype(e ld.Entity) Datatype { return Datatype{AsClass(e)} }

// Does the object quack like a(n) Datatype?
func IsDatatype(e ld.Entity) bool { return ld.Is(e, TypeDatatype) }

// The class of literal values, eg. textual strings and integers.
type Literal struct{ Resource }

// Ducktypes the object into a(n) Literal.
func AsLiteral(e ld.Entity) Literal { return Literal{AsResource(e)} }

// Does the object quack like a(n) Literal?
func IsLiteral(e ld.Entity) bool { return ld.Is(e, TypeLiteral) }

// The class resource, everything.
type Resource struct{ o *ld.Object }

// Ducktypes the object into a(n) Resource.
func AsResource(e ld.Entity) Resource { return Resource{o: e.Obj()} }

// Does the object quack like a(n) Resource?
func IsResource(e ld.Entity) bool { return ld.Is(e, TypeResource) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Resource) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj Resource) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Resource) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Resource) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Resource) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Resource) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// The subject is an instance of a class.
func (obj Resource) Type_() interface{} { return GetType(obj) }

// Idiomatic property used for structured values.
func (obj Resource) Value_() interface{} { return GetValue(obj) }

// A description of the subject resource.
func (obj Resource) Comment() interface{} { return GetComment(obj) }

// The defininition of the subject resource.
func (obj Resource) IsDefinedBy() interface{} { return GetIsDefinedBy(obj) }

// A human-readable name for the subject.
func (obj Resource) Label() interface{} { return GetLabel(obj) }

// A member of the subject resource.
func (obj Resource) Member() interface{} { return GetMember(obj) }

// Further information about the subject resource.
func (obj Resource) SeeAlso() interface{} { return GetSeeAlso(obj) }

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
