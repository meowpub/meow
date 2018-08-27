package entities

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/ns"
	"github.com/meowpub/meow/server/api"
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
	New: func(flake snowflake.ID, id string) (Entity, error) {
		return &Activity{
			Object: Object{
				Base: Base{
					Meta:      jsonld.Meta{},
					Snowflake: flake,
					ID:        id,
					Type:      []string{ns.AS("Activity")},
				},
			},
		}, nil
	},
	Unmarshall: func(flake snowflake.ID, obj map[string]interface{}) (Entity, error) {
		v := &Activity{Object: Object{Base: Base{
			Snowflake: flake,
		}}}
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
func (o *Activity) HandleRequest(req api.Request) api.Response {
	return handleEntityGetRequest(req, o)
}

func (self *Activity) Hydrate(ctx context.Context, stack []snowflake.ID) (interface{}, error) {
	stack = append([]snowflake.ID{self.GetSnowflake()}, stack...)
	o, err := jsonld.Marshal(self)
	if err != nil {
		return nil, err
	}

	hydrateChildren(ctx, o, stack,
		ns.AS("actor"),
		ns.AS("object"),
		ns.AS("target"),
		ns.AS("origin"),
		ns.AS("result"),
		ns.AS("instrument"))

	return o, nil
}

func init() {
	RegisterKind(activityKind)
}
