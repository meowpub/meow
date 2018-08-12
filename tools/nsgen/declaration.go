package main

import (
	"strings"

	"github.com/meowpub/meow/ld"
)

const (
	RDF      = "http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	RDF_type = RDF + "type"

	RDFS             = "http://www.w3.org/2000/01/rdf-schema#"
	RDFS_comment     = RDFS + "comment"
	RDFS_isDefinedBy = RDFS + "isDefinedBy"
	RDFS_label       = RDFS + "label"
	RDFS_subClassOf  = RDFS + "subClassOf"
	RDFS_seeAlso     = RDFS + "seeAlso"
	RDFS_domain      = RDFS + "domain"
	RDFS_range       = RDFS + "range"
)

var _ ld.Entity = Declaration{}

type Declaration struct {
	*ld.Object
	Namespace *Namespace
}

func (t Declaration) Obj() *ld.Object {
	return t.Object
}

func (t Declaration) TypeName() string {
	s := strings.TrimPrefix(t.ID(), t.Namespace.Long)
	if len(s) > 0 {
		s = strings.ToUpper(s[0:1]) + s[1:]
	}
	return s
}

func (t Declaration) RDFType() string {
	return ld.ToObject(t.V[RDF_type]).ID()
}

func (t Declaration) Domain() string {
	return ld.ToObject(t.V[RDFS_domain]).ID()
}
