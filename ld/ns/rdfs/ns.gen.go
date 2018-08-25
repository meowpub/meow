// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdfs

import (
	"github.com/meowpub/meow/ld"
)

const Namespace = "http://www.w3.org/2000/01/rdf-schema#"

const (
	// A description of the subject resource.
	PropComment = "http://www.w3.org/2000/01/rdf-schema#comment"

	// A domain of the subject property.
	PropDomain = "http://www.w3.org/2000/01/rdf-schema#domain"

	// The defininition of the subject resource.
	PropIsDefinedBy = "http://www.w3.org/2000/01/rdf-schema#isDefinedBy"

	// A human-readable name for the subject.
	PropLabel = "http://www.w3.org/2000/01/rdf-schema#label"

	// A member of the subject resource.
	PropMember = "http://www.w3.org/2000/01/rdf-schema#member"

	// A range of the subject property.
	PropRange = "http://www.w3.org/2000/01/rdf-schema#range"

	// Further information about the subject resource.
	PropSeeAlso = "http://www.w3.org/2000/01/rdf-schema#seeAlso"

	// The subject is a subclass of a class.
	PropSubClassOf = "http://www.w3.org/2000/01/rdf-schema#subClassOf"

	// The subject is a subproperty of a property.
	PropSubPropertyOf = "http://www.w3.org/2000/01/rdf-schema#subPropertyOf"
)

// Namespace.
var NS = &ld.Namespace{
	ID:    "http://www.w3.org/2000/01/rdf-schema#",
	Short: "rdfs",
	Props: map[string]string{
		"comment":       "http://www.w3.org/2000/01/rdf-schema#comment",
		"domain":        "http://www.w3.org/2000/01/rdf-schema#domain",
		"isDefinedBy":   "http://www.w3.org/2000/01/rdf-schema#isDefinedBy",
		"label":         "http://www.w3.org/2000/01/rdf-schema#label",
		"member":        "http://www.w3.org/2000/01/rdf-schema#member",
		"range":         "http://www.w3.org/2000/01/rdf-schema#range",
		"seeAlso":       "http://www.w3.org/2000/01/rdf-schema#seeAlso",
		"subClassOf":    "http://www.w3.org/2000/01/rdf-schema#subClassOf",
		"subPropertyOf": "http://www.w3.org/2000/01/rdf-schema#subPropertyOf",
	},
	Classes: map[string]func(ld.Entity) ld.Entity{
		"Class":                       func(e ld.Entity) ld.Entity { return &Class{O: e.Obj()} },
		"Container":                   func(e ld.Entity) ld.Entity { return &Container{O: e.Obj()} },
		"ContainerMembershipProperty": func(e ld.Entity) ld.Entity { return &ContainerMembershipProperty{O: e.Obj()} },
		"Datatype":                    func(e ld.Entity) ld.Entity { return &Datatype{O: e.Obj()} },
		"Literal":                     func(e ld.Entity) ld.Entity { return &Literal{O: e.Obj()} },
		"Resource":                    func(e ld.Entity) ld.Entity { return &Resource{O: e.Obj()} },
	},
}
