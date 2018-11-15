package ns

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/meta"
)

// Returns whether the entity quacks like a certain type.
func QuacksLike(t *meta.Type, e ld.Entity) bool {
	for _, tID := range e.Type() {
		if t.ID == tID {
			return true
		}
		typ, ok := Classes[tID]
		if !ok {
			continue
		}
		if typ.IsSubClassOf(t) {
			return true
		}
	}
	return false
}

// Manifests concrete, typed Entities from an Object based on its Type() entries.
// It's worth noting that the Object is not copied; all returned Entities refer to the same
// underlying Object, and motifications to one will reflect on all of them.
func Manifest(obj *ld.Object) (entities []ld.Entity) {
	for _, typ := range obj.Type() {
		if typ, ok := Classes[typ]; ok {
			entities = append(entities, typ.Cast(obj))
		}
	}
	return
}
