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

func (s StringItem) String() string {
	return s.Value
}

func ToString(s string) String {
	return String{StringItem{Value: s}}
}
