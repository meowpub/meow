package entities

import (
	"context"
	"net/http"

	"github.com/bwmarrin/snowflake"
	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/server/api"
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
	Inbox        jsonld.Ref        `json:"http://www.w3.org/ns/ldp#inbox,omitempty"`
}

var _ Entity = &Object{}

var objectKind = &EntityKind{
	Name: "object",
	Unmarshall: func(obj map[string]interface{}) (Entity, error) {
		v := &Object{}
		return v, jsonld.Unmarshal(obj, v)
	},
	Marshall: func(e Entity) (map[string]interface{}, error) {
		v, err := jsonld.Marshal(e.(*Object))
		if err != nil {
			return nil, err
		} else {
			return v.(map[string]interface{}), nil
		}
	},
}

func (*Object) GetKind() *EntityKind {
	return objectKind
}

// api.Handler

// Return ourselves
func (o *Object) HandleRequest(ctx context.Context, req *http.Request) api.Response {
	return handleEntityGetRequest(ctx, o, req)
}

func (self *Object) Hydrate(ctx context.Context, stack []snowflake.ID) (interface{}, error) {
	stack = append([]snowflake.ID{self.GetSnowflake()}, stack...)

	o, err := jsonld.Marshal(self)
	if err != nil {
		return nil, err
	}

	hydrateChildren(ctx, o, stack,
		as2("attachment"),
		as2("attributedTo"),
		as2("context"),
		as2("generator"),
		as2("icon"),
		as2("image"),
		as2("inReplyTo"),
		as2("location"),
		as2("preview"),
		as2("replies"),
		as2("tag"),
		as2("url"),
		ldp("inbox"))

	return o, nil
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
