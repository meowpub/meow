package main

import (
	"strings"
	"text/template"

	"github.com/meowpub/meow/ld"
)

var Funcs = template.FuncMap{
	"toobject": ld.ToObject,
	"tostring": ld.ToString,
	"tovalue": func(v interface{}) string {
		if obj := ld.ToObject(v); obj != nil {
			return obj.Value()
		}
		return ""
	},
	"comment": func(v interface{}) string {
		if s := ld.ToString(v); s != "" {
			lines := strings.Split(s, "\n")
			for i, line := range lines {
				lines[i] = strings.TrimSpace(line)
			}
			return "// " + strings.Join(lines, "\n// ")
		}
		return ""
	},
}

type RenderContext struct {
	Package      string
	Namespaces   []*Namespace
	Declarations []*Declaration
}

func (rctx RenderContext) OfType(t string) (matches []*Declaration) {
	for _, decl := range rctx.Declarations {
		if decl.RDFType() == t {
			matches = append(matches, decl)
		}
	}
	return
}

func (rctx RenderContext) Ontology() *Declaration {
	for _, decl := range rctx.Declarations {
		if decl.RDFType() == "http://www.w3.org/2002/07/owl#Ontology" {
			return decl
		}
	}
	return nil
}

func (rctx RenderContext) Classes() []*Declaration {
	return rctx.OfType("http://www.w3.org/2000/01/rdf-schema#Class")
}

func (rctx RenderContext) Properties() []*Declaration {
	return rctx.OfType("http://www.w3.org/1999/02/22-rdf-syntax-ns#Property")
}

func (rctx RenderContext) PropertiesOf(of *Declaration) (matches []*Declaration) {
	for _, decl := range rctx.Properties() {
		if decl.Domain() == of.ID() {
			matches = append(matches, decl)
		}
	}
	return
}

func (rctx RenderContext) DataTypes() []*Declaration {
	return rctx.OfType("http://www.w3.org/2000/01/rdf-schema#Datatype")
}

func (rctx RenderContext) Misc() (matches []*Declaration) {
	for _, decl := range rctx.Declarations {
		switch decl.RDFType() {
		case "http://www.w3.org/2000/01/rdf-schema#Class":
		case "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property":
		case "http://www.w3.org/2000/01/rdf-schema#Datatype":
		case "http://www.w3.org/2002/07/owl#Ontology":
		default:
			matches = append(matches, decl)
		}
	}
	return
}

// func (rctx RenderContext) NS() map[string]*ld.Namespace {
// 	return ns.Namespaces
// }

var NSTemplate = template.Must(template.New("ns.gen.go").Funcs(Funcs).Parse(`
// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package {{.Package}}

import (
	"github.com/meowpub/meow/ld"
)

{{if .Properties}}
const ( {{range .Properties}}
	{{comment (.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
	Prop{{.TypeName}} = "{{.ID}}"
	{{end}}
)
{{end}}

{{range .Misc}}
{{comment (.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
// {{.ID}} - {{.RDFType}}
{{end}}
`[1:]))

var ClassesTemplate = template.Must(template.New("classes.gen.go").Funcs(Funcs).Parse(`
// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package {{.Package}}

import (
	"github.com/meowpub/meow/ld"
)

{{range $i, $cls := .Classes}}
{{comment ($cls.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
type {{$cls.TypeName}} struct { O *ld.Object }

func (obj {{$cls.TypeName}}) Obj() *ld.Object { return obj.O }
func (obj {{$cls.TypeName}}) ID() string { return obj.O.ID() }
func (obj {{$cls.TypeName}}) Value() string { return obj.O.Value() }
func (obj {{$cls.TypeName}}) Type() []string { return obj.O.Type() }
func (obj {{$cls.TypeName}}) Get(key string) interface{} { return obj.O.Get(key) }

func (obj {{$cls.TypeName}}) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

{{range $.PropertiesOf $cls}}
{{comment (.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
func (obj {{$cls.TypeName}}) {{.FuncName}}() interface{} {
	return obj.O.V["{{.ID}}"]
}
{{end}}
{{end}}

var ({{range .Classes}}
	_ ld.Entity = {{.TypeName}}{}
	{{- end}}
)
`[1:]))

var DataTypesTemplate = template.Must(template.New("datatypes.gen.go").Funcs(Funcs).Parse(`
// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package {{.Package}}

{{range $i, $dt := .DataTypes}}
{{comment ($dt.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
type {{$dt.TypeName}} interface{}
{{end}}
`[1:]))

// var IndexTemplate = template.Must(template.New("index.gen.go").Funcs(Funcs).Parse(`
// // GENERATED FILE, DO NOT EDIT.
// // Please refer to: tools/nsgen/templates.go
// package ns

// import (
// 	"github.com/meowpub/meow/ld"{{range .}}
// 	"github.com/meowpub/meow/ld/ns/{{.Short}}"{{end}}
// )

// // Namespace.
// {{range .}}
// var ns_{{.Short}} = &ld.Namespace{
// 	ID: "{{.Long}}",
// 	Short: "{{.Short}}",
// 	Props: map[string]string{ {{range .Properties}}
// 		"{{.Short}}": "{{.ID}}",
// 		{{- end}}
// 	},
// }
// {{end}}

// var Namespaces = map[string]*ld.Namespace{ {{range .}}
// 	"{{.Long}}": ns_{{.Short}},
// 	"{{.Short}}": ns_{{.Short}}, {{end}}
// }
// `[1:]))
