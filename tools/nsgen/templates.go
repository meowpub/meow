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
	Declarations []*Declaration

	Namespaces []*Namespace
	Packages   map[string][]*Declaration
}

func (rctx RenderContext) ResolvePrefix(id, pfx string) string {
	for pkg, decls := range rctx.Packages {
		for _, decl := range decls {
			if decl.ID() == id {
				t := pfx + decl.TypeName()
				if pkg != rctx.Package {
					t = pkg + "." + t
				}
				return t
			}
		}
	}
	panic(id)
}

func (rctx RenderContext) Resolve(id string) string {
	return rctx.ResolvePrefix(id, "")
}

func (rctx RenderContext) OfType(ts ...string) (matches []*Declaration) {
nextDecl:
	for _, decl := range rctx.Declarations {
		rdfTypes := decl.RDFTypes()
		for _, t := range ts {
			for _, rdfT := range rdfTypes {
				if rdfT == t {
					matches = append(matches, decl)
					continue nextDecl
				}
			}
		}
	}
	return
}

func (rctx RenderContext) Ontology() *Declaration {
	for _, decl := range rctx.OfType("http://www.w3.org/2002/07/owl#Ontology") {
		return decl
	}
	return nil
}

func (rctx RenderContext) Classes() []*Declaration {
	return rctx.OfType(
		"http://www.w3.org/2000/01/rdf-schema#Class",
		"http://www.w3.org/2002/07/owl#Class",
	)
}

func (rctx RenderContext) Properties() []*Declaration {
	return rctx.OfType(
		"http://www.w3.org/1999/02/22-rdf-syntax-ns#Property",
		"http://www.w3.org/2002/07/owl#AnnotationProperty",
		"http://www.w3.org/2002/07/owl#DatatypeProperty",
		"http://www.w3.org/2002/07/owl#ObjectProperty",
		"http://www.w3.org/2002/07/owl#OntologyProperty",
	)
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
nextDecl:
	for _, decl := range rctx.Declarations {
		for _, rdfT := range decl.RDFTypes() {
			switch rdfT {
			case "http://www.w3.org/2000/01/rdf-schema#Class",
				"http://www.w3.org/1999/02/22-rdf-syntax-ns#Property",
				"http://www.w3.org/2000/01/rdf-schema#Datatype",
				"http://www.w3.org/2002/07/owl#Ontology",
				"http://www.w3.org/2002/07/owl#Class",
				"http://www.w3.org/2002/07/owl#AnnotationProperty",
				"http://www.w3.org/2002/07/owl#DatatypeProperty",
				"http://www.w3.org/2002/07/owl#ObjectProperty",
				"http://www.w3.org/2002/07/owl#OntologyProperty":
				continue nextDecl
			default:
				matches = append(matches, decl)
			}
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
	"github.com/meowpub/meow/ld" {{range .Namespaces}}
	"github.com/meowpub/meow/ld/ns/{{.Package}}" {{end}}
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
// {{.ID}} - {{.RDFTypes}}
{{end}}
`[1:]))

var ClassesTemplate = template.Must(template.New("classes.gen.go").Funcs(Funcs).Parse(`
// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package {{.Package}}

import (
	"github.com/meowpub/meow/ld" {{range .Namespaces}}
	"github.com/meowpub/meow/ld/ns/{{.Package}}" {{end}}
)

{{range $i, $cls := .Classes}}
{{comment ($cls.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
type {{$cls.TypeName}} struct { {{range .SubClassOf}}{{$.Resolve .}}; {{else}}o *ld.Object{{end}} }

func As{{$cls.TypeName}}(obj *ld.Object) {{$cls.TypeName}} {
	return {{$cls.TypeName}}{ {{range .SubClassOf}}{{$.ResolvePrefix . "As"}}(obj),{{else}}o: obj{{end}} }
}

func Is{{$cls.TypeName}}(obj *ld.Object) bool {
	return ld.Is(obj, "{{.ID}}")
}

{{if not .SubClassOf}}
// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj {{$cls.TypeName}}) Obj() *ld.Object {
	return obj.o
}

// Returns the object's @id. Implements ld.Entity.
func (obj {{$cls.TypeName}}) ID() string {
	return obj.o.ID()
}

// Returns the object's @value. Implements ld.Entity.
func (obj {{$cls.TypeName}}) Value() string {
	return obj.o.Value()
}

// Returns the object's @type. Implements ld.Entity.
func (obj {{$cls.TypeName}}) Type() []string {
	return obj.o.Type()
}

// Returns the named attribute. Implements ld.Entity.
func (obj {{$cls.TypeName}}) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj {{$cls.TypeName}}) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}
{{end}}

{{range $.PropertiesOf $cls}}
{{comment (.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
func (obj {{$cls.TypeName}}) {{.FuncName}}() interface{} {
	return obj.Get("{{.ID}}")
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
