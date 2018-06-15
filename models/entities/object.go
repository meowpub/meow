package entities

import (
	"context"
	"net/http"

	"github.com/bwmarrin/snowflake"
	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/ns"
	"github.com/meowpub/meow/server/api"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
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
	New: func(flake snowflake.ID, id string) (Entity, error) {
		return &Object{
			Base: Base{
				Meta:      jsonld.Meta{},
				Snowflake: flake,
				ID:        id,
				Type:      []string{ns.AS("Object")},
			},
		}, nil
	},
	Unmarshall: func(flake snowflake.ID, obj map[string]interface{}) (Entity, error) {
		v := &Object{Base: Base{
			Snowflake: flake,
		}}
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
		ns.AS("attachment"),
		ns.AS("attributedTo"),
		ns.AS("context"),
		ns.AS("generator"),
		ns.AS("icon"),
		ns.AS("image"),
		ns.AS("inReplyTo"),
		ns.AS("location"),
		ns.AS("preview"),
		ns.AS("replies"),
		ns.AS("tag"),
		ns.AS("url"),
		ns.LDP("inbox"))

	return o, nil
}

func (self *Object) GetInbox(ctx context.Context) (*Stream, error) {
	store := GetStore(ctx)
	if len(self.Inbox) > 0 {
		inbox, err := store.GetByID(self.Inbox[0].ID)
		return inbox.(*Stream), err
	}

	inbox_, err := store.NewChildEntity(ctx, self, "stream", "inbox")
	if err != nil {
		return nil, errors.Wrap(err, "Getting inbox")
	}
	inbox := inbox_.(*Stream)
	inbox.AttributedTo = jsonld.ToRef(self.GetID())
	self.Inbox = jsonld.ToRef(inbox.GetID())

	return inbox, multierr.Combine(
		store.Save(self),
		store.Save(inbox))

}

func init() {
	RegisterKind(objectKind)
}
