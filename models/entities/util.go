package entities

import (
	"context"
	"net/http"

	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/server/api"
)

func compact(ctx context.Context, data interface{}) (interface{}, error) {
	return jsonld.Compact(lib.GetHttpClient(ctx), data.(map[string]interface{}), "", "https://www.w3.org/ns/activitystreams")
}

func handleEntityGetRequest(req api.Request, e Entity) api.Response {
	if req.Method != http.MethodGet && req.Method != http.MethodOptions && req.Method != http.MethodHead {
		return api.Response{
			Status: http.StatusMethodNotAllowed,
			Error:  errors.New("Unsupported request method " + req.Method),
		}
	}

	d, err := e.Hydrate(req, []snowflake.ID{})
	if err != nil {
		return api.ErrorResponse(err)
	}

	o, err := compact(req, d)
	if err != nil {
		return api.ErrorResponse(err)
	}

	return api.Response{
		Data: o,
	}
}

func hydrateChildren(ctx context.Context, ob interface{}, stack []snowflake.ID, keys ...string) {
	log := lib.GetLogger(ctx)

	if len(stack) > 3 {
		return
	}

	o, ok := ob.(map[string]interface{})
	if !ok {
		return
	}

	estore := GetStore(ctx)

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		s, ok := o[key]
		if !ok {
			continue
		}

		slice, ok := s.([]interface{})
		if !ok {
			log.Debug("Not hydrating child because not a slice", zap.String("key", key))
			continue
		}

	ent:
		for j := 0; j < len(slice); j++ {
			m, ok := slice[j].(map[string]interface{})
			if !ok || len(m) != 1 {
				log.Debug("Not hydrating child of key because not a map", zap.String("key", key))
				continue
			}
			if id, ok := m["@id"]; ok {
				id_str, ok := id.(string)
				if !ok {
					log.Debug("Not hydrating child of key because id isn't a string", zap.String("key", key))
					continue
				}

				ent, err := estore.GetByID(id_str)
				if err != nil {
					log.Warn("Couldn't get object", zap.String("key", key), zap.String("id", id_str), zap.Error(err))
					continue
				}

				for k := 0; k < len(stack); k++ {
					if stack[k] == ent.GetSnowflake() {
						continue ent
					}
				}

				if h, err := ent.Hydrate(ctx, stack); err == nil {
					slice[j] = h
				} else {
					log.Warn("Error hydrating entity", zap.String("key", key), zap.String("id", id_str), zap.Error(err))
				}
			}
		}
	}
}
