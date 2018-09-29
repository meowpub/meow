package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
	"go.uber.org/zap"

	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/config/secrets"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/server/api"
	"github.com/meowpub/meow/server/middleware"
	"github.com/meowpub/meow/server/oauth"
	"github.com/meowpub/meow/server/well_known"
)

// New returns a new API router.
func New(db *gorm.DB, r *redis.Client, keyspace string) http.Handler {
	renderOpts := render.Options{
		Directory:     "templates",
		Extensions:    []string{".html"},
		IndentJSON:    true,
		IsDevelopment: !config.IsProd(),
	}

	mux := api.NewRouter(Lookup, renderOpts)
	mux.Use(middleware.AddDB(db))
	mux.Use(middleware.AddRedis(r))
	mux.Use(middleware.AddStores())
	mux.Use(middleware.AddSession(sessions.NewCookieStore(secrets.SessionKey())))
	mux.Use(oauth.AuthenticationMiddleware)

	mux.Mount("favicon.ico", nil)
	mux.Mount(".well-known", well_known.WellKnown)
	mux.Mount("-", &api.Node{
		Self: api.HandlerFunc(func(req api.Request) api.Response {
			return api.Response{Data: "hi"}
		}),
		Children: map[string]api.Handler{
			"login": api.HandlerFunc(HandleLogin),
		},
	})

	return mux
}

func HandleNotFound(req api.Request) api.Response {
	return api.Response{Error: api.Wrap(errors.New("not found"), http.StatusNotFound)}
}

func Lookup(ctx context.Context, url string) (api.Traversible, error) {
	e, err := models.GetStores(ctx).Entities().GetByID(url)
	if err != nil {
		lib.GetLogger(ctx).Debug("Lookup failed", zap.String("url", url), zap.Error(err))
		return nil, err
	}
	if e == nil {
		lib.GetLogger(ctx).Debug("Lookup found nothing", zap.String("url", url))
		return nil, nil
	}
	// TODO: Instantiate things properly instead.
	handlerFn := api.HandlerFunc(func(api.Request) api.Response {
		return api.Response{Data: e.Obj}
	})
	return api.Node{Self: handlerFn}, nil
}
