package jsonld

type Ref []RefItem

type RefItem struct {
	Meta
	ID   ID   `json:"@id,omitempty"`
	Type Type `json:"@type,omitempty"`
}

func (ref RefItem) MarshalJSON() ([]byte, error) {
	return ref.Meta.Marshal(ref)
}

func (ref *RefItem) UnmarshalJSON(data []byte) error {
	return ref.Meta.Unmarshal(data, ref)
}

func ToRef(id ID, typ Type) Ref {
	return Ref{RefItem{ID: id, Type: typ}}
}
