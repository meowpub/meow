// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

import (
	"github.com/meowpub/meow/ld"
)

const Namespace = "http://www.w3.org/1999/02/22-rdf-syntax-ns#"

const (
	PropFirst = "http://www.w3.org/1999/02/22-rdf-syntax-ns#first"

	PropObject = "http://www.w3.org/1999/02/22-rdf-syntax-ns#object"

	PropPredicate = "http://www.w3.org/1999/02/22-rdf-syntax-ns#predicate"

	PropRest = "http://www.w3.org/1999/02/22-rdf-syntax-ns#rest"

	PropSubject = "http://www.w3.org/1999/02/22-rdf-syntax-ns#subject"

	PropType = "http://www.w3.org/1999/02/22-rdf-syntax-ns#type"

	PropValue = "http://www.w3.org/1999/02/22-rdf-syntax-ns#value"
)

// Namespace.
var NS = &ld.Namespace{
	ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
	Short: "rdf",
	Props: map[string]string{
		"first":     "http://www.w3.org/1999/02/22-rdf-syntax-ns#first",
		"object":    "http://www.w3.org/1999/02/22-rdf-syntax-ns#object",
		"predicate": "http://www.w3.org/1999/02/22-rdf-syntax-ns#predicate",
		"rest":      "http://www.w3.org/1999/02/22-rdf-syntax-ns#rest",
		"subject":   "http://www.w3.org/1999/02/22-rdf-syntax-ns#subject",
		"type":      "http://www.w3.org/1999/02/22-rdf-syntax-ns#type",
		"value":     "http://www.w3.org/1999/02/22-rdf-syntax-ns#value",
	},
	Classes: map[string]func(ld.Entity) ld.Entity{
		"Alt":       func(e ld.Entity) ld.Entity { return &Alt{O: e.Obj()} },
		"Bag":       func(e ld.Entity) ld.Entity { return &Bag{O: e.Obj()} },
		"List":      func(e ld.Entity) ld.Entity { return &List{O: e.Obj()} },
		"Property":  func(e ld.Entity) ld.Entity { return &Property{O: e.Obj()} },
		"Seq":       func(e ld.Entity) ld.Entity { return &Seq{O: e.Obj()} },
		"Statement": func(e ld.Entity) ld.Entity { return &Statement{O: e.Obj()} },
	},
}
