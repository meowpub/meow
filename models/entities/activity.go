package entities

import (
	"context"
	"net/http"

	"github.com/bwmarrin/snowflake"
	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/server/api"
	"github.com/pkg/errors"
)

type Activity struct {
	Object

	Actor      jsonld.Ref `json:"https://www.w3.org/ns/activitystreams#actor,omitempty"`
	Obj        jsonld.Ref `json:"https://www.w3.org/ns/activitystreams#object,omitempty"`
	Target     jsonld.Ref `json:"https://www.w3.org/ns/activitystreams#target,omitempty"`
	Origin     jsonld.Ref `json:"https://www.w3.org/ns/activitystreams#origin,omitempty"`
	Result     jsonld.Ref `json:"https://www.w3.org/ns/activitystreams#result,omitemtpy"`
	Instrument jsonld.Ref `json:"https://www.w3.org/ns/activitystreams#instrument,omitempty"`
}

var _ Entity = &Activity{}

var activityKind = &EntityKind{
	Name: "activity",
	Unmarshall: func(obj map[string]interface{}) (Entity, error) {
		v := &Activity{}
		return v, jsonld.Unmarshal(obj, v)
	},
	Marshall: func(e Entity) (map[string]interface{}, error) {
		v, err := jsonld.Marshal(e.(*Activity))
		if err != nil {
			return nil, err
		} else {
			return v.(map[string]interface{}), nil
		}
	},
}

func (*Activity) GetKind() *EntityKind {
	return activityKind
}

// api.Handler

// Return ourselves
func (o *Activity) HandleRequest(ctx context.Context, req *http.Request) api.Response {
	return handleEntityGetRequest(ctx, o, req)
}

func (self *Activity) Hydrate(ctx context.Context, stack []snowflake.ID) (interface{}, error) {
	stack = append([]snowflake.ID{self.GetSnowflake()}, stack...)
	o, err := jsonld.Marshal(self)
	if err != nil {
		return nil, err
	}

	hydrateChildren(ctx, o, stack,
		as2("actor"),
		as2("object"),
		as2("target"),
		as2("origin"),
		as2("result"),
		as2("instrument"))

	return o, nil
}

func NewActivity(store *Store, parent Entity, types []string) (*Activity, error) {
	flake, err := lib.GenSnowflake(config.NodeID())
	if err != nil {
		return nil, errors.Wrap(err, "Generating snowflake for Activity")
	}

	id := parent.GetID() + "/" + flake.String()

	obj := &Activity{
		Object: Object{
			Base: Base{
				Meta: jsonld.Meta{},
				ID:   id,
				Type: types,
			},
		},
	}
	obj.SetSnowflake(flake)

	if err := store.Insert(obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func init() {
	RegisterKind(activityKind)
}
