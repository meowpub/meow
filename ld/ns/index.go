package ns

import (
	"github.com/meowpub/meow/ld"
)

// Describes a namespace.
type Namespace struct {
	ID    string           // "http://www.w3.org/ns/activitystreams#"
	Short string           // "as"
	Props []*Prop          // All properties in this namespace.
	Types map[string]*Type // All Types in this namespace.
}

// Describes a type.
type Type struct {
	ID    string  // "http://www.w3.org/ns/activitystreams#Note"
	Short string  // "Note"
	Props []*Prop // Property names.

	// Casts an arbitrary Entity to this type.
	// Underlying Objects are not copied - see the note on Manifest!
	Cast func(ld.Entity) ld.Entity
}

// Describes a property.
type Prop struct {
	ID      string
	Short   string
	Comment string
}

// Manifests concrete, typed Entities from an Object based on its Type() entries.
// It's worth noting that the Object is not copied; all returned Entities refer to the same
// underlying Object, and motifications to one will reflect on all of them.
func Manifest(obj *ld.Object) (entities []ld.Entity) {
	for _, typ := range obj.Type() {
		if typ, ok := Types[typ]; ok {
			entities = append(entities, typ.Cast(obj))
		}
	}
	return
}
