package jsonld

type Ref []RefItem

type RefItem struct {
	Meta
	ID   ID   `json:"@id"`
	Type Type `json:"@type"`
}

func ToRef(id ID, typ Type) Ref {
	return Ref{RefItem{ID: id, Type: typ}}
}
