package entities

import (
	"context"
	"net/http"

	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/server/api"
)

type Person struct {
	Object
	PreferredUsername jsonld.String `json:"https://www.w3c.org/ns/activitystreams#preferredUsername,omitempty"`
	Outbox            jsonld.Ref    `json:"https://www.w3c.org/ns/activitystreams#outbox,omitempty"`
	Following         jsonld.Ref    `json:"https://www.w3c.org/ns/activitystreams#following,omitempty"`
	Followers         jsonld.Ref    `json:"https://www.w3c.org/ns/activitystreams#followers,omitempty"`
	Liked             jsonld.Ref    `json:"https://www.w3c.org/ns/activitystreams#liked,omitempty"`
	Streams           jsonld.Ref    `json:"https://www.w3c.org/ns/activitystreams#streams,omitempty"`
	Endpoints         jsonld.Ref    `json:"https://www.w3c.org/ns/activitystreams#endpoints,omitempty"`
}

var personKind = &EntityKind{
	Name: "person",
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

func (*Person) GetKind() *EntityKind {
	return personKind
}

func (p *Person) GetUser(store models.UserStore) (*models.User, error) {
	return store.GetBySnowflake(p.GetSnowflake())
}

// Return ourselves serialized as a json blob
// TOOD: Compact!
func (o *Person) HandleRequest(ctx context.Context, req *http.Request) api.Response {
	return api.Response{
		Data: o,
	}
}

func NewPerson(store *Store, id string) (*Person, error) {
	obj := &Person{
		Object: Object{
			Base: Base{
				Meta: jsonld.Meta{},
				ID:   id,
				Type: []string{"https://www.w3c.org/ns/activitystreams#Person"},
			},
		},
	}

	if err := store.Insert(obj); err != nil {
		return nil, err
	}

	// Revisit: Create Inbox/Outbox/etc

	return obj, nil
}

func init() {
	RegisterKind(personKind)
}
