package entities

import (
	"context"
	"net/http"

	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/server/api"
	"github.com/pkg/errors"
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
	Unmarshall: func(obj map[string]interface{}) (Entity, error) {
		v := &Person{}
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

func (*Person) GetKind() *EntityKind {
	return personKind
}

func (self *Person) Hydrate(ctx context.Context) (map[string]interface{}, error) {
	if o, err := jsonld.Marshal(self); err == nil {
		return jsonld.Compact(lib.GetHttpClient(ctx), o.(map[string]interface{}), "", "https://www.w3.org/ns/activitystreams")
	} else {
		return nil, err
	}
}

func (p *Person) GetUser(store models.UserStore) (*models.User, error) {
	return store.GetBySnowflake(p.GetSnowflake())
}

// Return ourselves
func (o *Person) HandleRequest(ctx context.Context, req *http.Request) api.Response {
	d, err := o.Hydrate(ctx)
	if err != nil {
		return api.ErrorResponse(err)
	}

	return api.Response{
		Data: d,
	}
}

func NewPerson(ctx context.Context, id string) (*Person, error) {
	store := GetStore(ctx)
	stores := models.GetStores(ctx)

	obj := &Person{
		Object: Object{
			Base: Base{
				Meta: jsonld.Meta{},
				ID:   id,
				Type: []string{"https://www.w3.org/ns/activitystreams#Person"},
			},
		},
	}

	if err := store.Insert(obj); err != nil {
		return nil, err
	}

	// TOOD: These hould be Inbox/Outbox kinds but this will do For Now (TM)
	inbox, err := NewStream(store, id+"/inbox")
	if err != nil {
		return nil, errors.Wrap(err, "Creating inbox")
	}

	outbox, err := NewStream(store, id+"/outbox")
	if err != nil {
		return nil, errors.Wrap(err, "Creating outbox")
	}

	inbox.AttributedTo = jsonld.ToRef(id)
	outbox.AttributedTo = jsonld.ToRef(id)
	obj.Inbox = jsonld.ToRef(inbox.GetID())
	obj.Outbox = jsonld.ToRef(outbox.GetID())

	creation, err := NewActivity(store, outbox, []string{as2("Create")})
	if err != nil {
		return nil, errors.Wrap(err, "Creating Create activity")
	}

	creation.Actor = jsonld.ToRef(id)
	creation.Obj = jsonld.ToRef(id)

	if err := store.Save(inbox); err != nil {
		return nil, errors.Wrap(err, "Saving inbox")
	}
	if err := store.Save(outbox); err != nil {
		return nil, errors.Wrap(err, "Saving outbox")
	}
	if err := store.Save(creation); err != nil {
		return nil, errors.Wrap(err, "Saving Create")
	}
	if _, _, err := stores.StreamItems().TryInsertItem(outbox.GetSnowflake(), creation.GetSnowflake()); err != nil {
		return nil, errors.Wrap(err, "Inserting Create into Outbox")
	}
	if _, _, err := stores.StreamItems().TryInsertItem(inbox.GetSnowflake(), creation.GetSnowflake()); err != nil {
		return nil, errors.Wrap(err, "Inserting Create into Inbox")
	}

	return obj, nil
}

func init() {
	RegisterKind(personKind)
}
