package patches

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/rdf"
)

func patch(id, key, subkey, value string) *ld.Object {
	return &ld.Object{V: map[string]interface{}{
		"@id": id,
		key: []interface{}{
			map[string]interface{}{subkey: value},
		},
	}}
}

func comment(typ, comment string) *ld.Object {
	return patch(typ, rdf.Prop_Comment.ID, "@value", comment)
}
