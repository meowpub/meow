package jsonld

// String represents a property which contains xsd:string values
type String []StringItem

// StringItem returns a single string item
type StringItem struct {
	Meta
	Value string `json:"@value"`
}

// String converts a String into a string
func (s String) String() string {
	if len(s) == 0 {
		return ""
	}
	return s[0].Value
}

// ToString converts a string into a String
func ToString(s string) String {
	return String{StringItem{Value: s}}
}

// String converts a StringItem into a string
func (s StringItem) String() string {
	return s.Value
}

func (s StringItem) MarshalJSON() ([]byte, error) {
	return s.Meta.Marshal(&s)
}

func (s *StringItem) UnmarshalJSON(data []byte) error {
	return s.Meta.Unmarshal(data, s)
}
