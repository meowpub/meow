package jsonld

type Ref []RefItem

type RefItem struct {
	Meta
	ID string `json:"@id,omitempty"`
}

func ToRef(id string) Ref {
	return Ref{RefItem{ID: id}}
}
