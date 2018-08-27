package xrd

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

type XRD struct {
	ID         string            `json:"-"`
	Expires    *time.Time        `json:"expires,omitempty"`
	Subject    string            `json:"subject,omitempty"`
	Aliases    []string          `json:"aliases,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	Links      []Link            `json:"links,omitempty"`
}

func (x *XRD) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	xm := xmlXRD{
		ID:         x.ID,
		Expires:    x.Expires,
		Subject:    x.Subject,
		Aliases:    x.Aliases,
		Properties: make([]xmlProperty, 0, len(x.Properties)),
		Links:      x.Links,
	}

	for k, v := range x.Properties {
		xm.Properties = append(xm.Properties, xmlProperty{
			Type:  k,
			Value: v,
		})
	}

	return e.EncodeElement(xm, start)
}

type xmlXRD struct {
	XMLName    xml.Name      `xml:"XRD"`
	ID         string        `xml:"http://www.w3.org/XML/1998/namespace id,attr,omitempty"`
	Expires    *time.Time    `xml:"Expires,omitempty"`
	Subject    string        `xml:"Subject,omitempty"`
	Aliases    []string      `xml:"Alias"`
	Properties []xmlProperty `xml:"Property"`
	Links      []Link        `xml:"Link"`
}

type xmlProperty struct {
	XMLName xml.Name `xml:"Property" json:"-"`

	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type Link struct {
	Rel        string            `json:"rel,omitempty"`
	Type       string            `json:"type,omitempty"`
	HRef       string            `json:"href,omitempty"`
	Template   string            `json:"template,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	Titles     map[string]string `json:"title,omitempty"`
}

func (l *Link) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	xl := xmlLink{
		Rel:        l.Rel,
		Type:       l.Type,
		HRef:       l.HRef,
		Template:   l.Template,
		Properties: make([]xmlProperty, 0, len(l.Properties)),
		Titles:     make([]xmlTitle, 0, len(l.Titles)),
	}

	for k, v := range l.Properties {
		xl.Properties = append(xl.Properties, xmlProperty{
			Type:  k,
			Value: v,
		})
	}

	for k, v := range l.Titles {
		if k == "und" {
			k = ""
		}

		xl.Titles = append(xl.Titles, xmlTitle{
			Lang:  k,
			Value: v,
		})
	}

	return e.EncodeElement(xl, start)
}

type xmlLink struct {
	XMLName    xml.Name      `xml:"Link" json:"-"`
	Rel        string        `xml:"rel,attr,omitempty"`
	Type       string        `xml:"type,attr,omitempty"`
	HRef       string        `xml:"href,attr,omitempty"`
	Template   string        `xml:"template,attr,omitempty"`
	Properties []xmlProperty `xml:"Property,omitempty"`
	Titles     []xmlTitle    `xml:"Title,omitempty"`
}

type xmlTitle struct {
	XMLName xml.Name `xml:"Title" json:"-"`
	Lang    string   `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

func UnmarshalXRD(buf []byte) (*XRD, error) {
	x := &XRD{}
	if err := xml.Unmarshal(buf, x); err != nil {
		return nil, err
	}
	return x, nil
}

func (x *XRD) MarshalXRD() ([]byte, error) {
	return xml.Marshal(x)
}

func (x *XRD) MarshalIndentXRD() ([]byte, error) {
	return xml.MarshalIndent(x, "", "\t")
}

func (x *XRD) MarshalJRD() ([]byte, error) {
	return json.Marshal(x)
}

func (x *XRD) MarshalIndentJRD() ([]byte, error) {
	return json.MarshalIndent(x, "", "\t")
}
