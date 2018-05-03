package jsonld

// LangString represents a property which contains rdf:langString values
// https://www.w3.org/TR/rdf11-concepts/#dfn-language-tagged-string
type LangString []LangStringItem

func (s LangString) String() string {
	if len(s) == 0 {
		return ""
	}
	return s[0].Value
}

// LangStringItm represents a single language tagged rdf:langString
type LangStringItem struct {
	Meta
	Value    string `json:"@value"`
	Language string `json:"@language,omitempty"`
}

// ToLangString converts a string to an untagged LangString instance
func ToLangString(s string) LangString {
	return LangString{LangStringItem{Value: s}}
}

// String returns the raw string inside a LangStringItem
func (s LangStringItem) String() string {
	return s.Value
}

func (s LangStringItem) MarshalJSON() ([]byte, error) {
	return s.Meta.Marshal(&s)
}

func (s *LangStringItem) UnmarshalJSON(data []byte) error {
	return s.Meta.Unmarshal(data, s)
}
