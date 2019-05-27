// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

import (
	"github.com/meowpub/meow/ld"
)

// The class of containers of alternatives.
type Alt struct{ Container }

func NewAlt(id string) Alt { return AsAlt(ld.NewObject(id, Class_Alt.ID)) }

// Ducktypes the object into a(n) Alt.
func AsAlt(e ld.Entity) Alt { return Alt{AsContainer(e)} }

// Does the object quack like a(n) Alt?
func IsAlt(e ld.Entity) bool { return ld.Is(e, Class_Alt.ID) }

// The class of unordered containers.
type Bag struct{ Container }

func NewBag(id string) Bag { return AsBag(ld.NewObject(id, Class_Bag.ID)) }

// Ducktypes the object into a(n) Bag.
func AsBag(e ld.Entity) Bag { return Bag{AsContainer(e)} }

// Does the object quack like a(n) Bag?
func IsBag(e ld.Entity) bool { return ld.Is(e, Class_Bag.ID) }

// The class of RDF Lists.
type List struct{ Resource }

func NewList(id string) List { return AsList(ld.NewObject(id, Class_List.ID)) }

// Ducktypes the object into a(n) List.
func AsList(e ld.Entity) List { return List{AsResource(e)} }

// Does the object quack like a(n) List?
func IsList(e ld.Entity) bool { return ld.Is(e, Class_List.ID) }

// The first item in the subject RDF list.
func (obj List) First() interface{} { return GetFirst(obj) }

func (obj List) SetFirst(v interface{}) { SetFirst(obj, v) }

// The rest of the subject RDF list after the first item.
func (obj List) Rest() interface{} { return GetRest(obj) }

func (obj List) SetRest(v interface{}) { SetRest(obj, v) }

// The class of RDF properties.
type Property struct{ Resource }

func NewProperty(id string) Property { return AsProperty(ld.NewObject(id, Class_Property.ID)) }

// Ducktypes the object into a(n) Property.
func AsProperty(e ld.Entity) Property { return Property{AsResource(e)} }

// Does the object quack like a(n) Property?
func IsProperty(e ld.Entity) bool { return ld.Is(e, Class_Property.ID) }

// A domain of the subject property.
func (obj Property) Domain() interface{} { return GetDomain(obj) }

func (obj Property) SetDomain(v interface{}) { SetDomain(obj, v) }

// A range of the subject property.
func (obj Property) Range() interface{} { return GetRange(obj) }

func (obj Property) SetRange(v interface{}) { SetRange(obj, v) }

// The subject is a subproperty of a property.
func (obj Property) SubPropertyOf() interface{} { return GetSubPropertyOf(obj) }

func (obj Property) SetSubPropertyOf(v interface{}) { SetSubPropertyOf(obj, v) }

// The class of ordered containers.
type Seq struct{ Container }

func NewSeq(id string) Seq { return AsSeq(ld.NewObject(id, Class_Seq.ID)) }

// Ducktypes the object into a(n) Seq.
func AsSeq(e ld.Entity) Seq { return Seq{AsContainer(e)} }

// Does the object quack like a(n) Seq?
func IsSeq(e ld.Entity) bool { return ld.Is(e, Class_Seq.ID) }

// The class of RDF statements.
type Statement struct{ Resource }

func NewStatement(id string) Statement { return AsStatement(ld.NewObject(id, Class_Statement.ID)) }

// Ducktypes the object into a(n) Statement.
func AsStatement(e ld.Entity) Statement { return Statement{AsResource(e)} }

// Does the object quack like a(n) Statement?
func IsStatement(e ld.Entity) bool { return ld.Is(e, Class_Statement.ID) }

// The object of the subject RDF statement.
func (obj Statement) Object() interface{} { return GetObject(obj) }

func (obj Statement) SetObject(v interface{}) { SetObject(obj, v) }

// The predicate of the subject RDF statement.
func (obj Statement) Predicate() interface{} { return GetPredicate(obj) }

func (obj Statement) SetPredicate(v interface{}) { SetPredicate(obj, v) }

// The subject of the subject RDF statement.
func (obj Statement) Subject() interface{} { return GetSubject(obj) }

func (obj Statement) SetSubject(v interface{}) { SetSubject(obj, v) }

// The class of classes.
type Class struct{ Resource }

func NewClass(id string) Class { return AsClass(ld.NewObject(id, Class_Class.ID)) }

// Ducktypes the object into a(n) Class.
func AsClass(e ld.Entity) Class { return Class{AsResource(e)} }

// Does the object quack like a(n) Class?
func IsClass(e ld.Entity) bool { return ld.Is(e, Class_Class.ID) }

// The subject is a subclass of a class.
func (obj Class) SubClassOf() interface{} { return GetSubClassOf(obj) }

func (obj Class) SetSubClassOf(v interface{}) { SetSubClassOf(obj, v) }

// The class of RDF containers.
type Container struct{ Resource }

func NewContainer(id string) Container { return AsContainer(ld.NewObject(id, Class_Container.ID)) }

// Ducktypes the object into a(n) Container.
func AsContainer(e ld.Entity) Container { return Container{AsResource(e)} }

// Does the object quack like a(n) Container?
func IsContainer(e ld.Entity) bool { return ld.Is(e, Class_Container.ID) }

// The class of container membership properties, rdf:_1, rdf:_2, ...,
// all of which are sub-properties of 'member'.
type ContainerMembershipProperty struct{ Property }

func NewContainerMembershipProperty(id string) ContainerMembershipProperty {
	return AsContainerMembershipProperty(ld.NewObject(id, Class_ContainerMembershipProperty.ID))
}

// Ducktypes the object into a(n) ContainerMembershipProperty.
func AsContainerMembershipProperty(e ld.Entity) ContainerMembershipProperty {
	return ContainerMembershipProperty{AsProperty(e)}
}

// Does the object quack like a(n) ContainerMembershipProperty?
func IsContainerMembershipProperty(e ld.Entity) bool {
	return ld.Is(e, Class_ContainerMembershipProperty.ID)
}

// The class of RDF datatypes.
type Datatype struct{ Class }

func NewDatatype(id string) Datatype { return AsDatatype(ld.NewObject(id, Class_Datatype.ID)) }

// Ducktypes the object into a(n) Datatype.
func AsDatatype(e ld.Entity) Datatype { return Datatype{AsClass(e)} }

// Does the object quack like a(n) Datatype?
func IsDatatype(e ld.Entity) bool { return ld.Is(e, Class_Datatype.ID) }

// The class of literal values, eg. textual strings and integers.
type Literal struct{ Resource }

func NewLiteral(id string) Literal { return AsLiteral(ld.NewObject(id, Class_Literal.ID)) }

// Ducktypes the object into a(n) Literal.
func AsLiteral(e ld.Entity) Literal { return Literal{AsResource(e)} }

// Does the object quack like a(n) Literal?
func IsLiteral(e ld.Entity) bool { return ld.Is(e, Class_Literal.ID) }

// The class resource, everything.
type Resource struct{ o *ld.Object }

func NewResource(id string) Resource { return AsResource(ld.NewObject(id, Class_Resource.ID)) }

// Ducktypes the object into a(n) Resource.
func AsResource(e ld.Entity) Resource { return Resource{o: e.Obj()} }

// Does the object quack like a(n) Resource?
func IsResource(e ld.Entity) bool { return ld.Is(e, Class_Resource.ID) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Resource) Obj() *ld.Object { return obj.o }

// Returns whether the wrapped object is null. Implements ld.Entity.
func (obj Resource) IsNull() bool { return obj.o == nil }

// Returns the object's @id. Implements ld.Entity.
func (obj Resource) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Resource) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Resource) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Resource) Get(key string) interface{} { return obj.o.Get(key) }

// Sets the named attribute. Implements ld.Entity.
func (obj Resource) Set(key string, v interface{}) { obj.o.Set(key, v) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Resource) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// The subject is an instance of a class.
func (obj Resource) Type_() interface{} { return GetType(obj) }

func (obj Resource) SetType_(v interface{}) { SetType(obj, v) }

// Idiomatic property used for structured values.
func (obj Resource) Value_() interface{} { return GetValue(obj) }

func (obj Resource) SetValue_(v interface{}) { SetValue(obj, v) }

// A description of the subject resource.
func (obj Resource) Comment() interface{} { return GetComment(obj) }

func (obj Resource) SetComment(v interface{}) { SetComment(obj, v) }

// The defininition of the subject resource.
func (obj Resource) IsDefinedBy() interface{} { return GetIsDefinedBy(obj) }

func (obj Resource) SetIsDefinedBy(v interface{}) { SetIsDefinedBy(obj, v) }

// A human-readable name for the subject.
func (obj Resource) Label() interface{} { return GetLabel(obj) }

func (obj Resource) SetLabel(v interface{}) { SetLabel(obj, v) }

// A member of the subject resource.
func (obj Resource) Member() interface{} { return GetMember(obj) }

func (obj Resource) SetMember(v interface{}) { SetMember(obj, v) }

// Further information about the subject resource.
func (obj Resource) SeeAlso() interface{} { return GetSeeAlso(obj) }

func (obj Resource) SetSeeAlso(v interface{}) { SetSeeAlso(obj, v) }

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
