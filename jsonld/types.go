package jsonld

import (
	"strings"
)

type Type []string

func (t Type) String() string {
	if len(t) == 0 {
		return ""
	}
	return strings.Join(t, ";")
}
