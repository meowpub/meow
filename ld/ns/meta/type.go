package meta

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

// Describes a property.
type Prop struct {
	ID      string
	Short   string
	Comment string
}

// Describes a type.
type Type struct {
	ID         string  // "http://www.w3.org/ns/activitystreams#Note"
	Short      string  // "Note"
	SubClassOf []*Type // Supertypes, if any.
	Props      []*Prop // Property names.

	// Casts an arbitrary Entity to this type.
	// Underlying Objects are not copied - see the note on Manifest!
	Cast func(ld.Entity) ld.Entity
}

// Returns whether the type inherits from another type, somewhere up the inheritance chain.
func (t *Type) IsSubClassOf(other *Type) bool {
	for _, sup := range t.SubClassOf {
		if sup.ID == other.ID {
			return true
		}
	}
	return false
}
