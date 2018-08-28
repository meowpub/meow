// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

import (
	"github.com/meowpub/meow/ld"
)

// The first item in the subject RDF list.
func GetFirst(e ld.Entity) interface{} { return e.Get(PropFirst) }

func SetFirst(e ld.Entity, v interface{}) { e.Set(PropFirst, v) }

// The object of the subject RDF statement.
func GetObject(e ld.Entity) interface{} { return e.Get(PropObject) }

func SetObject(e ld.Entity, v interface{}) { e.Set(PropObject, v) }

// The predicate of the subject RDF statement.
func GetPredicate(e ld.Entity) interface{} { return e.Get(PropPredicate) }

func SetPredicate(e ld.Entity, v interface{}) { e.Set(PropPredicate, v) }

// The rest of the subject RDF list after the first item.
func GetRest(e ld.Entity) interface{} { return e.Get(PropRest) }

func SetRest(e ld.Entity, v interface{}) { e.Set(PropRest, v) }

// The subject of the subject RDF statement.
func GetSubject(e ld.Entity) interface{} { return e.Get(PropSubject) }

func SetSubject(e ld.Entity, v interface{}) { e.Set(PropSubject, v) }

// The subject is an instance of a class.
func GetType(e ld.Entity) interface{} { return e.Get(PropType) }

func SetType(e ld.Entity, v interface{}) { e.Set(PropType, v) }

// Idiomatic property used for structured values.
func GetValue(e ld.Entity) interface{} { return e.Get(PropValue) }

func SetValue(e ld.Entity, v interface{}) { e.Set(PropValue, v) }

// A description of the subject resource.
func GetComment(e ld.Entity) interface{} { return e.Get(PropComment) }

func SetComment(e ld.Entity, v interface{}) { e.Set(PropComment, v) }

// A domain of the subject property.
func GetDomain(e ld.Entity) interface{} { return e.Get(PropDomain) }

func SetDomain(e ld.Entity, v interface{}) { e.Set(PropDomain, v) }

// The defininition of the subject resource.
func GetIsDefinedBy(e ld.Entity) interface{} { return e.Get(PropIsDefinedBy) }

func SetIsDefinedBy(e ld.Entity, v interface{}) { e.Set(PropIsDefinedBy, v) }

// A human-readable name for the subject.
func GetLabel(e ld.Entity) interface{} { return e.Get(PropLabel) }

func SetLabel(e ld.Entity, v interface{}) { e.Set(PropLabel, v) }

// A member of the subject resource.
func GetMember(e ld.Entity) interface{} { return e.Get(PropMember) }

func SetMember(e ld.Entity, v interface{}) { e.Set(PropMember, v) }

// A range of the subject property.
func GetRange(e ld.Entity) interface{} { return e.Get(PropRange) }

func SetRange(e ld.Entity, v interface{}) { e.Set(PropRange, v) }

// Further information about the subject resource.
func GetSeeAlso(e ld.Entity) interface{} { return e.Get(PropSeeAlso) }

func SetSeeAlso(e ld.Entity, v interface{}) { e.Set(PropSeeAlso, v) }

// The subject is a subclass of a class.
func GetSubClassOf(e ld.Entity) interface{} { return e.Get(PropSubClassOf) }

func SetSubClassOf(e ld.Entity, v interface{}) { e.Set(PropSubClassOf, v) }

// The subject is a subproperty of a property.
func GetSubPropertyOf(e ld.Entity) interface{} { return e.Get(PropSubPropertyOf) }

func SetSubPropertyOf(e ld.Entity, v interface{}) { e.Set(PropSubPropertyOf, v) }
