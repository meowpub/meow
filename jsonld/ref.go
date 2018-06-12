package jsonld

type Ref []RefItem

type RefItem struct {
	Meta
	ID   ID   `json:"@id,omitempty"`
	Type Type `json:"@type,omitempty"`
}

func ToRef(id ID, typ Type) Ref {
	return Ref{RefItem{ID: id, Type: typ}}
}
