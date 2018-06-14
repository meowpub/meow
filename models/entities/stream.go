package entities

import (
	"context"
	"net/http"

	"github.com/bwmarrin/snowflake"
	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/server/api"
)

type Stream struct {
	Object
}

var streamKind = &EntityKind{
	Name: "stream",
	Unmarshall: func(obj map[string]interface{}) (Entity, error) {
		v := &Stream{}
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
		as2("attributedTo"),
		as2("icon"),
		as2("image"))

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

		o[as2("items")] = []interface{}{m}
	}

	return o, nil
}

// Return ourselves
func (o *Stream) HandleRequest(ctx context.Context, req *http.Request) api.Response {
	return handleEntityGetRequest(ctx, o, req)
}

func NewStream(store *Store, id string) (*Stream, error) {
	obj := &Stream{
		Object: Object{
			Base: Base{
				Meta: jsonld.Meta{},
				ID:   id,
				Type: []string{as2("Collection")},
			},
		},
	}

	return obj, store.Insert(obj)
}

func init() {
	RegisterKind(streamKind)
}
