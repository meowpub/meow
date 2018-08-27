package entities

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/ns"
	"github.com/meowpub/meow/server/api"
)

type Stream struct {
	Object
}

var streamKind = &EntityKind{
	Name: "stream",
	New: func(flake snowflake.ID, id string) (Entity, error) {
		return &Stream{
			Object: Object{
				Base: Base{
					Meta:      jsonld.Meta{},
					Snowflake: flake,
					ID:        id,
					Type:      []string{ns.AS("Collection")},
				},
			},
		}, nil
	},
	Unmarshall: func(flake snowflake.ID, obj map[string]interface{}) (Entity, error) {
		v := &Stream{Object: Object{Base: Base{
			Snowflake: flake,
		}}}
		return v, jsonld.Unmarshal(obj, v)
	},
	Marshall: func(e Entity) (map[string]interface{}, error) {
		v, err := jsonld.Marshal(e.(*Stream))
		if err != nil {
			return nil, err
		} else {
			return v.(map[string]interface{}), nil
		}
	},
}

func (*Stream) GetKind() *EntityKind {
	return streamKind
}

func (self *Stream) Hydrate(ctx context.Context, stack []snowflake.ID) (interface{}, error) {
	stack = append([]snowflake.ID{self.GetSnowflake()}, stack...)
	o_, err := jsonld.Marshal(self)
	if err != nil {
		return nil, err
	}
	o := o_.(map[string]interface{})

	hydrateChildren(ctx, o, stack,
		ns.AS("attributedTo"),
		ns.AS("icon"),
		ns.AS("image"))

	if len(stack) == 1 {
		itemStore := models.GetStores(ctx).StreamItems()
		eStore := GetStore(ctx)

		items, err := itemStore.GetItems(self.GetSnowflake(), models.End, models.Before, 20)
		if err != nil {
			return nil, err
		}

		ents := make([]interface{}, len(items))
		for i := 0; i < len(items); i++ {
			e, err := eStore.GetBySnowflake(items[i].EntityID)
			if err != nil {
				return nil, err
			}
			o, err := e.Hydrate(ctx, stack)
			if err != nil {
				return nil, err
			}
			ents[i] = o
		}

		m := map[string]interface{}{
			"@list": ents,
		}

		o[ns.AS("items")] = []interface{}{m}
	}

	return o, nil
}

func (self *Stream) InsertItem(ctx context.Context, ent Entity) error {
	stores := models.GetStores(ctx)
	_, _, err := stores.StreamItems().TryInsertItem(self.GetSnowflake(), ent.GetSnowflake())
	return err
}

// Return ourselves
func (o *Stream) HandleRequest(req api.Request) api.Response {
	return handleEntityGetRequest(req, o)
}

func init() {
	RegisterKind(streamKind)
}
