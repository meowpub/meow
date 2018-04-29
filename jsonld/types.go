package jsonld

import (
	"strings"
)

type ID string

type Type []string

func (t Type) String() string {
	if len(t) == 0 {
		return ""
	}
	return strings.Join(t, ";")
}
