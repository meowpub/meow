package main

import (
	"text/template"
)

var AS2GenTmpl = template.Must(template.New("").Parse(`
package entities

{{range .}}
// {{.Comment}}
type AS2{{.Label}} struct {}

func (AS2{{.Label}}) Type() string {
	return "{{.ID}}"
}
{{end}}

// Compile-time assertions to make sure all generated types are valid.
var (
{{range .}}_ Type = AS2{{.Label}}{}
{{end}}
)
`[1:]))
