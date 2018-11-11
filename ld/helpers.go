package ld

func Is(e Entity, typ string) bool {
	for _, typ2 := range e.Type() {
		if typ == typ2 {
			return true
		}
	}
	return false
}

func ObjectIDs(objs []*Object) []string {
	ids := make([]string, len(objs))
	for i, obj := range objs {
		ids[i] = obj.ID()
	}
	return ids
}

// Shorthand to make {"@value": v} structures.
func Value(v interface{}) map[string]interface{} {
	return map[string]interface{}{"@value": v}
}

// Shorthand to make {"@id": v} structures.
func ID(id string) map[string]interface{} {
	return map[string]interface{}{"@id": id}
}
