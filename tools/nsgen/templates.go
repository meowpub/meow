package main

import (
	"regexp"
	"strings"
	"text/template"

	"github.com/meowpub/meow/ld"
)

var varSafeRE = regexp.MustCompile(`([^a-zA-Z0-9])+`)

var Funcs = template.FuncMap{
	"toobject": ld.ToObject,
	"tostring": ld.ToString,
	"tovalue": func(v interface{}) string {
		if obj := ld.ToObject(v); obj != nil {
			return obj.Value()
		}
		return ""
	},
	"toupper": strings.ToUpper,
	"varsafe": varsafe,
	"quotesafe": func(v interface{}) string {
		s := ld.ToString(v)
		s = strings.Replace(s, "\"", "\\\"", -1)
		s = strings.Replace(s, "\n", " ", -1)
		return s
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

func varsafe(s string) string {
	return varSafeRE.ReplaceAllString(s, "_")
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
		domain := decl.Domain()
		if domain == of.ID() {
			matches = append(matches, decl)
		}
		if domain == "" && len(of.SubClassOf()) == 0 {
			matches = append(matches, decl)
		}
	}
	return
}

func (rctx RenderContext) DataTypes() []*Declaration {
	return rctx.OfType("http://www.w3.org/2000/01/rdf-schema#Datatype")
}

func (rctx RenderContext) Constants() []*Declaration {
	// This isn't even a real type. NamedIndividual is a thing, Individual isn't. Pls.
	return rctx.OfType("http://www.w3.org/2002/07/owl#Individual")
}

func (rctx RenderContext) Misc() (matches []*Declaration) {
	for _, decl := range rctx.Declarations {
		isMisc := len(decl.RDFTypes()) > 0
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
				"http://www.w3.org/2002/07/owl#OntologyProperty",
				"http://www.w3.org/2002/07/owl#Individual":
				isMisc = false
			}
		}
		if isMisc {
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
	"github.com/meowpub/meow/ld/ns/meta" {{range .Namespaces}}
	"github.com/meowpub/meow/ld/ns/{{.Package}}" {{end}}
)

var ( {{range $i, $ns := .Namespaces}} {{if eq $ns.Package $.Package}}
	{{$ns.Short|toupper}} = &meta.Namespace{
	ID: "{{$ns.Long}}",
	Short: "{{$ns.Short}}",
	Props: []*meta.Prop{ {{range $.Properties}}{{if eq .NS.Long $ns.Long}}
			Prop_{{.TypeName}},
		{{- end}}{{end}}
		},
	Types: map[string]*meta.Type{ {{range $.Classes}}{{if eq .NS.Long $ns.Long}}
	"{{.ID}}": Class_{{.TypeName}},
	{{- end}}{{end}}
	},
	}
{{- end}}{{end}}

{{range .Properties}}
	{{comment (.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
	Prop_{{.TypeName}} = &meta.Prop{
		ID: "{{.ID}}",
		Short: "{{.Short}}",
		Comment: "{{.Get "http://www.w3.org/2000/01/rdf-schema#comment"|quotesafe}}",
	}
{{end}}

{{range .Classes}}
	{{comment (.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
	Class_{{.TypeName}} = &meta.Type{
		ID: "{{.ID}}",
		Short: "{{.Short}}",
		Cast: func(e ld.Entity) ld.Entity { return As{{.TypeName}}(e) },
		SubClassOf: []*meta.Type{ {{range .SubClassOf}}
			{{$.ResolvePrefix . "Class_"}},
		{{- end}}
		},
		Props: []*meta.Prop{ {{range $.PropertiesOf .}}
			Prop_{{.TypeName}},
		{{- end}}
		},
	}
{{end}}
)

{{if .Constants}}
const ( {{range .Constants}}
	{{comment (.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
	{{.TypeName}} = "{{.ID}}"
	{{end}}
)
{{end}}

{{range .Misc}}
{{comment (.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
// {{.ID}} - {{.RDFTypes}}
{{end}}
`[1:]))

var PropertiesTemplate = template.Must(template.New("properties.gen.go").Funcs(Funcs).Parse(`
// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package {{.Package}}

import (
	"github.com/meowpub/meow/ld"
)

{{range .Properties}}
{{comment (.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
func Get{{.TypeName}}(e ld.Entity) interface{} { return e.Get(Prop_{{.TypeName}}.ID) }

func Set{{.TypeName}}(e ld.Entity, v interface{}) { e.Set(Prop_{{.TypeName}}.ID, v) }
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

func New{{$cls.TypeName}}(id string) {{$cls.TypeName}} { return As{{$cls.TypeName}}(ld.NewObject(id, Class_{{.TypeName}}.ID)) }

// Ducktypes the object into a(n) {{.TypeName}}.
func As{{$cls.TypeName}}(e ld.Entity) {{$cls.TypeName}} { return {{$cls.TypeName}}{ {{range .SubClassOf}}{{$.ResolvePrefix . "As"}}(e),{{else}}o: e.Obj(){{end}} } }

// Does the object quack like a(n) {{.TypeName}}?
func Is{{$cls.TypeName}}(e ld.Entity) bool { return ld.Is(e, Class_{{.TypeName}}.ID) }

{{if not .SubClassOf}}
// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj {{$cls.TypeName}}) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj {{$cls.TypeName}}) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj {{$cls.TypeName}}) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj {{$cls.TypeName}}) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj {{$cls.TypeName}}) Get(key string) interface{} { return obj.o.Get(key) }

// Sets the named attribute. Implements ld.Entity.
func (obj {{$cls.TypeName}}) Set(key string, v interface{}) { obj.o.Set(key, v) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj {{$cls.TypeName}}) Apply(other ld.Entity, mergeArrays bool) error { return obj.o.Apply(other, mergeArrays) }
{{end}}

{{range $.PropertiesOf $cls}}
{{comment (.Get "http://www.w3.org/2000/01/rdf-schema#comment")}}
func (obj {{$cls.TypeName}}) {{.FuncName}}() interface{} { return Get{{.TypeName}}(obj) }

func (obj {{$cls.TypeName}}) Set{{.FuncName}}(v interface{}) { Set{{.TypeName}}(obj, v) }
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

var IndexTemplate = template.Must(template.New("index.gen.go").Funcs(Funcs).Parse(`
// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package ns

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/meta" {{range .Namespaces}}
	"github.com/meowpub/meow/ld/ns/{{.Package}}" {{end}}
)

var Namespaces = map[string]*meta.Namespace{ {{range .Namespaces}}
	"{{.Long}}": {{.Package}}.{{.Short|toupper}},
{{- end}}
}

var Classes = map[string]*meta.Type{ {{range .Classes}}
	"{{.ID}}": {{.Package}}.Class_{{.TypeName}},
{{- end}}
}
`[1:]))
