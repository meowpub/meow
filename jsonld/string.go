package jsonld

type String []StringItem

func (s String) String() string {
	if len(s) == 0 {
		return ""
	}
	return s[0].Value
}

type StringItem struct {
	Meta
	Value    string `json:"@value"`
	Language string `json:"@language"`
}

func ToString(s string) String {
	return String{StringItem{Value: s}}
}

func (s StringItem) String() string {
	return s.Value
}

func (s StringItem) MarshalJSON() ([]byte, error) {
	return s.Meta.Marshal(s)
}

func (s *StringItem) UnmarshalJSON(data []byte) error {
	return s.Meta.Unmarshal(data, s)
}
