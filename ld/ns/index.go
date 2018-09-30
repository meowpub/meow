package ns

import (
	"github.com/meowpub/meow/ld"
)

// Manifests concrete, typed Entities from an Object based on its Type() entries.
// It's worth noting that the Object is not copied; all returned Entities refer to the same
// underlying Object, and motifications to one will reflect on all of them.
func Manifest(obj *ld.Object) (entities []ld.Entity) {
	for _, typ := range obj.Type() {
		if cast, ok := Index[typ]; ok {
			entities = append(entities, cast(obj))
		}
	}
	return
}
