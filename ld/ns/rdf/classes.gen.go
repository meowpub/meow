// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

import (
	"github.com/meowpub/meow/ld"
)

// The class of containers of alternatives.
type Alt struct{ Container }

// Ducktypes the object into a Alt.
func AsAlt(obj *ld.Object) Alt { return Alt{AsContainer(obj)} }

// Does the object quack like a Alt?
func IsAlt(obj *ld.Object) bool { return ld.Is(obj, TypeAlt) }

// The class of unordered containers.
type Bag struct{ Container }

// Ducktypes the object into a Bag.
func AsBag(obj *ld.Object) Bag { return Bag{AsContainer(obj)} }

// Does the object quack like a Bag?
func IsBag(obj *ld.Object) bool { return ld.Is(obj, TypeBag) }

// The class of RDF Lists.
type List struct{ Resource }

// Ducktypes the object into a List.
func AsList(obj *ld.Object) List { return List{AsResource(obj)} }

// Does the object quack like a List?
func IsList(obj *ld.Object) bool { return ld.Is(obj, TypeList) }

// The first item in the subject RDF list.
func (obj List) First() interface{} { return obj.Get(PropFirst) }

// The rest of the subject RDF list after the first item.
func (obj List) Rest() interface{} { return obj.Get(PropRest) }

// The class of RDF properties.
type Property struct{ Resource }

// Ducktypes the object into a Property.
func AsProperty(obj *ld.Object) Property { return Property{AsResource(obj)} }

// Does the object quack like a Property?
func IsProperty(obj *ld.Object) bool { return ld.Is(obj, TypeProperty) }

// A domain of the subject property.
func (obj Property) Domain() interface{} { return obj.Get(PropDomain) }

// A range of the subject property.
func (obj Property) Range() interface{} { return obj.Get(PropRange) }

// The subject is a subproperty of a property.
func (obj Property) SubPropertyOf() interface{} { return obj.Get(PropSubPropertyOf) }

// The class of ordered containers.
type Seq struct{ Container }

// Ducktypes the object into a Seq.
func AsSeq(obj *ld.Object) Seq { return Seq{AsContainer(obj)} }

// Does the object quack like a Seq?
func IsSeq(obj *ld.Object) bool { return ld.Is(obj, TypeSeq) }

// The class of RDF statements.
type Statement struct{ Resource }

// Ducktypes the object into a Statement.
func AsStatement(obj *ld.Object) Statement { return Statement{AsResource(obj)} }

// Does the object quack like a Statement?
func IsStatement(obj *ld.Object) bool { return ld.Is(obj, TypeStatement) }

// The object of the subject RDF statement.
func (obj Statement) Object() interface{} { return obj.Get(PropObject) }

// The predicate of the subject RDF statement.
func (obj Statement) Predicate() interface{} { return obj.Get(PropPredicate) }

// The subject of the subject RDF statement.
func (obj Statement) Subject() interface{} { return obj.Get(PropSubject) }

// The class of classes.
type Class struct{ Resource }

// Ducktypes the object into a Class.
func AsClass(obj *ld.Object) Class { return Class{AsResource(obj)} }

// Does the object quack like a Class?
func IsClass(obj *ld.Object) bool { return ld.Is(obj, TypeClass) }

// The subject is a subclass of a class.
func (obj Class) SubClassOf() interface{} { return obj.Get(PropSubClassOf) }

// The class of RDF containers.
type Container struct{ Resource }

// Ducktypes the object into a Container.
func AsContainer(obj *ld.Object) Container { return Container{AsResource(obj)} }

// Does the object quack like a Container?
func IsContainer(obj *ld.Object) bool { return ld.Is(obj, TypeContainer) }

// The class of container membership properties, rdf:_1, rdf:_2, ...,
// all of which are sub-properties of 'member'.
type ContainerMembershipProperty struct{ Property }

// Ducktypes the object into a ContainerMembershipProperty.
func AsContainerMembershipProperty(obj *ld.Object) ContainerMembershipProperty {
	return ContainerMembershipProperty{AsProperty(obj)}
}

// Does the object quack like a ContainerMembershipProperty?
func IsContainerMembershipProperty(obj *ld.Object) bool {
	return ld.Is(obj, TypeContainerMembershipProperty)
}

// The class of RDF datatypes.
type Datatype struct{ Class }

// Ducktypes the object into a Datatype.
func AsDatatype(obj *ld.Object) Datatype { return Datatype{AsClass(obj)} }

// Does the object quack like a Datatype?
func IsDatatype(obj *ld.Object) bool { return ld.Is(obj, TypeDatatype) }

// The class of literal values, eg. textual strings and integers.
type Literal struct{ Resource }

// Ducktypes the object into a Literal.
func AsLiteral(obj *ld.Object) Literal { return Literal{AsResource(obj)} }

// Does the object quack like a Literal?
func IsLiteral(obj *ld.Object) bool { return ld.Is(obj, TypeLiteral) }

// The class resource, everything.
type Resource struct{ o *ld.Object }

// Ducktypes the object into a Resource.
func AsResource(obj *ld.Object) Resource { return Resource{o: obj} }

// Does the object quack like a Resource?
func IsResource(obj *ld.Object) bool { return ld.Is(obj, TypeResource) }

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
func (obj Resource) Type_() interface{} { return obj.Get(PropType) }

// Idiomatic property used for structured values.
func (obj Resource) Value_() interface{} { return obj.Get(PropValue) }

// A description of the subject resource.
func (obj Resource) Comment() interface{} { return obj.Get(PropComment) }

// The defininition of the subject resource.
func (obj Resource) IsDefinedBy() interface{} { return obj.Get(PropIsDefinedBy) }

// A human-readable name for the subject.
func (obj Resource) Label() interface{} { return obj.Get(PropLabel) }

// A member of the subject resource.
func (obj Resource) Member() interface{} { return obj.Get(PropMember) }

// Further information about the subject resource.
func (obj Resource) SeeAlso() interface{} { return obj.Get(PropSeeAlso) }

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
