package entities

import (
	"context"
	"net/http"

	"github.com/liclac/meow/jsonld"
	"github.com/liclac/meow/server/api"
)

type Object struct {
	Base

	Attachment   jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#attachment,omitempty"`
	AttributedTo jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#attributedTo,omitempty"`
	Audience     jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#audience,omitempty"`
	Content      jsonld.LangString `json:"https://www.w3.org/ns/activitystreams#content,omitempty"`
	Context      jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#context,omitempty"`
	Name         jsonld.LangString `json:"https://www.w3.org/ns/activitystreams#name,omitempty"`
	EndTime      jsonld.DateTime   `json:"https://www.w3.org/ns/activitystreams#endTime,omitempty"`
	Generator    jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#generator,omitempty"`
	Icon         jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#icon,omitempty"`
	Image        jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#image,omitempty"`
	InReplyTo    jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#inReplyTo,omitempty"`
	Location     jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#location,omitempty"`
	Preview      jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#preview,omitempty"`
	Published    jsonld.DateTime   `json:"https://www.w3.org/ns/activitystreams#published,omitempty"`
	Replies      jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#replies,omitempty"`
	StartTime    jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#startTime,omitempty"`
	Summary      jsonld.LangString `json:"https://www.w3.org/ns/activitystreams#summary,omitempty"`
	Tag          jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#tag,omitempty"`
	Updated      jsonld.DateTime   `json:"https://www.w3.org/ns/activitystreams#updated,omitempty"`
	Url          jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#url,omitempty"`
	To           jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#to,omitempty"`
	BTo          jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#bto,omitempty"`
	CC           jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#cc,omitempty"`
	BCC          jsonld.Ref        `json:"https://www.w3.org/ns/activitystreams#bcc,omitempty"`
	MediaType    jsonld.String     `json:"https://www.w3.org/ns/activitystreams#mediaType,omitempty"`
	Inbox        jsonld.Ref        `json:"https://www.w3.org/ns/ldp#inbox,omitempty"`
}

var objectKind = &EntityKind{
	Name: "object",
	Unmarshall: func(data []byte) (Entity, error) {
		v := &Object{}
		if err := v.Meta.Unmarshal(data, v); err != nil {
			return nil, err
		} else {
			return v, nil
		}
	},

	// Marshall this object into an Entity
	Marshall: func(e Entity) ([]byte, error) {
		return e.(*Object).Meta.Marshal(e.(*Object))
	},
}

func (*Object) GetKind() *EntityKind {
	return objectKind
}

func (o *Object) UnmarshalJSON(data []byte) error {
	return o.Meta.Unmarshal(data, o)
}

func (o *Object) MarshalJSON() ([]byte, error) {
	return o.Marshal(o)
}

// api.Handler

// Return ourselves serialized as a json blob
// TOOD: Compact!
func (o *Object) HandleRequest(ctx context.Context, req *http.Request) api.Response {
	return api.Response{
		Data: o,
	}
}

func NewObject(store *Store, id string, types []string) (*Object, error) {
	obj := &Object{
		Base: Base{
			Meta: jsonld.Meta{},
			ID:   id,
			Type: types,
		},
	}

	if err := store.Insert(obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func init() {
	RegisterKind(objectKind)
}
