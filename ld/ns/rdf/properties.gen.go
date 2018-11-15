// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

import (
	"github.com/meowpub/meow/ld"
)

// The first item in the subject RDF list.
func GetFirst(e ld.Entity) interface{} { return e.Get(Prop_First.ID) }

func SetFirst(e ld.Entity, v interface{}) { e.Set(Prop_First.ID, v) }

// The object of the subject RDF statement.
func GetObject(e ld.Entity) interface{} { return e.Get(Prop_Object.ID) }

func SetObject(e ld.Entity, v interface{}) { e.Set(Prop_Object.ID, v) }

// The predicate of the subject RDF statement.
func GetPredicate(e ld.Entity) interface{} { return e.Get(Prop_Predicate.ID) }

func SetPredicate(e ld.Entity, v interface{}) { e.Set(Prop_Predicate.ID, v) }

// The rest of the subject RDF list after the first item.
func GetRest(e ld.Entity) interface{} { return e.Get(Prop_Rest.ID) }

func SetRest(e ld.Entity, v interface{}) { e.Set(Prop_Rest.ID, v) }

// The subject of the subject RDF statement.
func GetSubject(e ld.Entity) interface{} { return e.Get(Prop_Subject.ID) }

func SetSubject(e ld.Entity, v interface{}) { e.Set(Prop_Subject.ID, v) }

// The subject is an instance of a class.
func GetType(e ld.Entity) interface{} { return e.Get(Prop_Type.ID) }

func SetType(e ld.Entity, v interface{}) { e.Set(Prop_Type.ID, v) }

// Idiomatic property used for structured values.
func GetValue(e ld.Entity) interface{} { return e.Get(Prop_Value.ID) }

func SetValue(e ld.Entity, v interface{}) { e.Set(Prop_Value.ID, v) }

// A description of the subject resource.
func GetComment(e ld.Entity) interface{} { return e.Get(Prop_Comment.ID) }

func SetComment(e ld.Entity, v interface{}) { e.Set(Prop_Comment.ID, v) }

// A domain of the subject property.
func GetDomain(e ld.Entity) interface{} { return e.Get(Prop_Domain.ID) }

func SetDomain(e ld.Entity, v interface{}) { e.Set(Prop_Domain.ID, v) }

// The defininition of the subject resource.
func GetIsDefinedBy(e ld.Entity) interface{} { return e.Get(Prop_IsDefinedBy.ID) }

func SetIsDefinedBy(e ld.Entity, v interface{}) { e.Set(Prop_IsDefinedBy.ID, v) }

// A human-readable name for the subject.
func GetLabel(e ld.Entity) interface{} { return e.Get(Prop_Label.ID) }

func SetLabel(e ld.Entity, v interface{}) { e.Set(Prop_Label.ID, v) }

// A member of the subject resource.
func GetMember(e ld.Entity) interface{} { return e.Get(Prop_Member.ID) }

func SetMember(e ld.Entity, v interface{}) { e.Set(Prop_Member.ID, v) }

// A range of the subject property.
func GetRange(e ld.Entity) interface{} { return e.Get(Prop_Range.ID) }

func SetRange(e ld.Entity, v interface{}) { e.Set(Prop_Range.ID, v) }

// Further information about the subject resource.
func GetSeeAlso(e ld.Entity) interface{} { return e.Get(Prop_SeeAlso.ID) }

func SetSeeAlso(e ld.Entity, v interface{}) { e.Set(Prop_SeeAlso.ID, v) }

// The subject is a subclass of a class.
func GetSubClassOf(e ld.Entity) interface{} { return e.Get(Prop_SubClassOf.ID) }

func SetSubClassOf(e ld.Entity, v interface{}) { e.Set(Prop_SubClassOf.ID, v) }

// The subject is a subproperty of a property.
func GetSubPropertyOf(e ld.Entity) interface{} { return e.Get(Prop_SubPropertyOf.ID) }

func SetSubPropertyOf(e ld.Entity, v interface{}) { e.Set(Prop_SubPropertyOf.ID, v) }
