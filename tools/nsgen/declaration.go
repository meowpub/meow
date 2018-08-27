package main

import (
	"strings"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/rdf"
)

var _ ld.Entity = Declaration{}

type Declaration struct {
	*ld.Object
}

func (t Declaration) Obj() *ld.Object {
	return t.Object
}

func (t Declaration) Short() string {
	return strings.SplitN(t.ID(), "#", 2)[1]
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

func (t Declaration) RDFType() string {
	return ld.ToObject(t.V[rdf.PropType]).ID()
}

func (t Declaration) Domain() string {
	return ld.ToObject(t.V[rdf.PropDomain]).ID()
}
