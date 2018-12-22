package server

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"

	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/config/secrets"
	"github.com/meowpub/meow/lib"
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
	return api.Response{Error: lib.Error(http.StatusNotFound, "not found")}
}
