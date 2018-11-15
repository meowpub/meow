package main

import (
	"strings"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/rdf"
)

var _ ld.Entity = Declaration{}

type Declaration struct {
	*ld.Object
	NS *Namespace
}

func (t Declaration) Obj() *ld.Object {
	return t.Object
}

func (t Declaration) Short() string {
	if parts := strings.SplitN(t.ID(), "#", 2); len(parts) > 1 {
		return parts[1]
	}
	return t.ID()
}

func (t Declaration) Package() string {
	if t.NS != nil {
		return t.NS.Package
	}
	return ""
}

func (t Declaration) TypeName() string {
	if s := t.Short(); s != "" {
		return strings.ToUpper(s[0:1]) + s[1:]
	}
	return ""
}

func (t Declaration) FuncName() string {
	tname := t.TypeName()
	switch tname {
	case "Value":
		return "Value_"
	case "Type":
		return "Type_"
	default:
		return tname
	}
}

func (t Declaration) RDFTypes() []string {
	return ld.ObjectIDs(ld.ToObjects(t.V[rdf.Prop_Type.ID]))
}

func (t Declaration) Domain() string {
	return ld.ToObject(t.V[rdf.Prop_Domain.ID]).ID()
}

func (t Declaration) SubClassOf() []string {
	return ld.ObjectIDs(ld.ToObjects(t.V[rdf.Prop_SubClassOf.ID]))
}
