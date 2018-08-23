// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package ns

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/ldp"
	"github.com/meowpub/meow/ld/ns/owl"
	"github.com/meowpub/meow/ld/ns/rdf"
	"github.com/meowpub/meow/ld/ns/rdfs"
	"github.com/meowpub/meow/ld/ns/xsd"
)

var Namespaces = map[string]*ld.Namespace{
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#": rdf.NS,
	"rdf": rdf.NS,
	"http://www.w3.org/2000/01/rdf-schema#": rdfs.NS,
	"rdfs": rdfs.NS,
	"http://www.w3.org/2002/07/owl#": owl.NS,
	"owl": owl.NS,
	"http://www.w3.org/2001/XMLSchema#": xsd.NS,
	"xsd": xsd.NS,
	"http://www.w3.org/ns/ldp#": ldp.NS,
	"ldp": ldp.NS,
}
