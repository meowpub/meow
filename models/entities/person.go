package entities

import (
	"context"
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/lib/xrd"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/ns"
	"github.com/meowpub/meow/server/api"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

type Person struct {
	Object
	PreferredUsername jsonld.String `json:"https://www.w3.org/ns/activitystreams#preferredUsername,omitempty"`
	Outbox            jsonld.Ref    `json:"https://www.w3.org/ns/activitystreams#outbox,omitempty"`
	Following         jsonld.Ref    `json:"https://www.w3.org/ns/activitystreams#following,omitempty"`
	Followers         jsonld.Ref    `json:"https://www.w3.org/ns/activitystreams#followers,omitempty"`
	Liked             jsonld.Ref    `json:"https://www.w3.org/ns/activitystreams#liked,omitempty"`
	Streams           jsonld.Ref    `json:"https://www.w3.org/ns/activitystreams#streams,omitempty"`
	Endpoints         jsonld.Ref    `json:"https://www.w3.org/ns/activitystreams#endpoints,omitempty"`
}

var personKind = &EntityKind{
	Name: "person",
	New: func(flake snowflake.ID, id string) (Entity, error) {
		return &Person{
			Object: Object{
				Base: Base{
					Meta:      jsonld.Meta{},
					Snowflake: flake,
					ID:        id,
					Type:      []string{ns.AS("Person")},
				},
			},
		}, nil
	},
	Unmarshall: func(flake snowflake.ID, obj map[string]interface{}) (Entity, error) {
		v := &Person{Object: Object{Base: Base{
			Snowflake: flake,
		}}}
		return v, jsonld.Unmarshal(obj, v)
	},
	Marshall: func(e Entity) (map[string]interface{}, error) {
		v, err := jsonld.Marshal(e.(*Person))
		if err != nil {
			return nil, err
		} else {
			return v.(map[string]interface{}), nil
		}
	},
}

func (self *Person) Init(ctx context.Context) error {
	store := GetStore(ctx)

	inbox, err := self.GetInbox(ctx)
	if err != nil {
		return err
	}

	outbox, err := self.GetOutbox(ctx)
	if err != nil {
		return err
	}

	creation_, err := store.NewChildEntity(ctx, outbox, "activity", "")
	if err != nil {
		return errors.Wrap(err, "Creating Create activity")
	}

	creation := creation_.(*Activity)
	creation.Type = []string{ns.AS("Create")}
	creation.Actor = jsonld.ToRef(self.GetID())
	creation.Obj = jsonld.ToRef(self.GetID())

	return multierr.Combine(
		store.Save(creation),
		outbox.InsertItem(ctx, creation),
		inbox.InsertItem(ctx, creation),
		store.Save(self))
}

func (*Person) GetKind() *EntityKind {
	return personKind
}

func (self *Person) Hydrate(ctx context.Context, stack []snowflake.ID) (interface{}, error) {
	stack = append([]snowflake.ID{self.GetSnowflake()}, stack...)
	o, err := jsonld.Marshal(self)
	if err != nil {
		return nil, err
	}

	hydrateChildren(ctx, o, stack,
		ns.AS("icon"),
		ns.AS("image"),
		ns.AS("location"),
		ns.AS("url"),
		ns.LDP("inbox"),
		ns.AS("outbox"),
		ns.AS("following"),
		ns.AS("followers"),
		ns.AS("liked"),
		ns.AS("streams"),
		ns.AS("endpoints"))

	return o, nil
}

func (p *Person) GetUser(store models.UserStore) (*models.User, error) {
	return store.GetBySnowflake(p.GetSnowflake())
}

// Return ourselves
func (o *Person) HandleRequest(req api.Request) api.Response {
	return handleEntityGetRequest(req, o)
}

func (o *Person) Finger(req api.Request) (*xrd.XRD, error) {
	url := req.ResourceURL()

	acctUrl := fmt.Sprintf("acct:%s@%s", o.PreferredUsername.String(), url.Hostname())
	return &xrd.XRD{
		Subject: url.String(),
		Aliases: []string{
			acctUrl,
		},
		Links: []xrd.Link{
			xrd.Link{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: "text/html",
				HRef: url.String(),
			},
			xrd.Link{
				Rel:  "self",
				Type: "text/html",
				HRef: url.String(),
			},
			xrd.Link{
				Rel:  "self",
				Type: "application/activity+json",
				HRef: url.String(),
			},
			xrd.Link{
				Rel:  "self",
				Type: "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"",
				HRef: url.String(),
			},
			xrd.Link{
				Rel:  "self",
				HRef: acctUrl,
			},
			xrd.Link{
				Rel:  "canonical",
				HRef: url.String(),
			},
		},
	}, nil
}

func (self *Person) GetOutbox(ctx context.Context) (*Stream, error) {
	store := GetStore(ctx)
	if len(self.Outbox) > 0 {
		outbox, err := store.GetByID(self.Outbox[0].ID)
		return outbox.(*Stream), err
	}

	outbox_, err := store.NewChildEntity(ctx, self, "stream", "outbox")
	if err != nil {
		return nil, errors.Wrap(err, "Getting outbox")
	}
	outbox := outbox_.(*Stream)
	outbox.AttributedTo = jsonld.ToRef(self.GetID())
	self.Outbox = jsonld.ToRef(outbox.GetID())

	return outbox, multierr.Combine(
		store.Save(self),
		store.Save(outbox))

}

func init() {
	RegisterKind(personKind)
}
