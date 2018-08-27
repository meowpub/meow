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
