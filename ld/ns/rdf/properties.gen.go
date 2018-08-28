// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

import (
	"github.com/meowpub/meow/ld"
)

// The first item in the subject RDF list.
func GetFirst(e ld.Entity) interface{} { return e.Get(PropFirst) }

// The object of the subject RDF statement.
func GetObject(e ld.Entity) interface{} { return e.Get(PropObject) }

// The predicate of the subject RDF statement.
func GetPredicate(e ld.Entity) interface{} { return e.Get(PropPredicate) }

// The rest of the subject RDF list after the first item.
func GetRest(e ld.Entity) interface{} { return e.Get(PropRest) }

// The subject of the subject RDF statement.
func GetSubject(e ld.Entity) interface{} { return e.Get(PropSubject) }

// The subject is an instance of a class.
func GetType(e ld.Entity) interface{} { return e.Get(PropType) }

// Idiomatic property used for structured values.
func GetValue(e ld.Entity) interface{} { return e.Get(PropValue) }

// A description of the subject resource.
func GetComment(e ld.Entity) interface{} { return e.Get(PropComment) }

// A domain of the subject property.
func GetDomain(e ld.Entity) interface{} { return e.Get(PropDomain) }

// The defininition of the subject resource.
func GetIsDefinedBy(e ld.Entity) interface{} { return e.Get(PropIsDefinedBy) }

// A human-readable name for the subject.
func GetLabel(e ld.Entity) interface{} { return e.Get(PropLabel) }

// A member of the subject resource.
func GetMember(e ld.Entity) interface{} { return e.Get(PropMember) }

// A range of the subject property.
func GetRange(e ld.Entity) interface{} { return e.Get(PropRange) }

// Further information about the subject resource.
func GetSeeAlso(e ld.Entity) interface{} { return e.Get(PropSeeAlso) }

// The subject is a subclass of a class.
func GetSubClassOf(e ld.Entity) interface{} { return e.Get(PropSubClassOf) }

// The subject is a subproperty of a property.
func GetSubPropertyOf(e ld.Entity) interface{} { return e.Get(PropSubPropertyOf) }
