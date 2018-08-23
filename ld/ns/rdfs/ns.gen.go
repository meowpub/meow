// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdfs

import (
	"github.com/meowpub/meow/ld"
)

const Namespace = "http://www.w3.org/2000/01/rdf-schema#"

const (
	PropComment = "http://www.w3.org/2000/01/rdf-schema#comment"

	PropDomain = "http://www.w3.org/2000/01/rdf-schema#domain"

	PropIsDefinedBy = "http://www.w3.org/2000/01/rdf-schema#isDefinedBy"

	PropLabel = "http://www.w3.org/2000/01/rdf-schema#label"

	PropMember = "http://www.w3.org/2000/01/rdf-schema#member"

	PropRange = "http://www.w3.org/2000/01/rdf-schema#range"

	PropSeeAlso = "http://www.w3.org/2000/01/rdf-schema#seeAlso"

	PropSubClassOf = "http://www.w3.org/2000/01/rdf-schema#subClassOf"

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
