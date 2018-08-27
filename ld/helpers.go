package ld

func Is(obj *Object, typ string) bool {
	for _, typ2 := range obj.Type() {
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
