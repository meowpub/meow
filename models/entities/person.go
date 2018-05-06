package entities

import (
	"github.com/liclac/meow/jsonld"
	"github.com/liclac/meow/models"
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
	Unmarshall: func(data []byte) (Entity, error) {
		v := &Person{}
		if err := v.Meta.Unmarshal(data, v); err != nil {
			return nil, err
		} else {
			return v, nil
		}
	},

	// Marshall this object into an Entity
	Marshall: func(e Entity) ([]byte, error) {
		return e.(*Person).Meta.Marshal(e.(*Person))
	},
}

func (*Person) GetKind() *EntityKind {
	return personKind
}

func (p *Person) GetUser(store models.UserStore) (*models.User, error) {
	return store.GetBySnowflake(p.GetSnowflake())
}

func (o *Person) UnmarshalJSON(data []byte) error {
	return o.Meta.Unmarshal(data, o)
}

func (o *Person) MarshalJSON() ([]byte, error) {
	return o.Marshal(o)
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
