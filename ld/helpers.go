package ld

func ObjectIDs(objs []*Object) []string {
	ids := make([]string, len(objs))
	for i, obj := range objs {
		ids[i] = obj.ID()
	}
	return ids
}
