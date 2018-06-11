package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
	"go.uber.org/zap"

	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/config/secrets"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models/entities"
	"github.com/meowpub/meow/server/api"
	"github.com/meowpub/meow/server/middleware"
	"github.com/meowpub/meow/server/oauth"
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
	mux.Mount("-", &api.Node{
		Self: api.HandlerFunc(func(ctx context.Context, req *http.Request) api.Response {
			return api.Response{Data: "hi"}
		}),
		Children: map[string]api.Handler{
			"login": api.HandlerFunc(HandleLogin),
		},
	})

	return mux
}

func HandleNotFound(ctx context.Context, req *http.Request) api.Response {
	return api.Response{Error: api.Wrap(errors.New("not found"), http.StatusNotFound)}
}

func Lookup(ctx context.Context, url string) (api.Traversible, error) {
	e, err := entities.GetStore(ctx).GetByID(url)
	if err != nil {
		lib.GetLogger(ctx).Debug("Lookup failed", zap.String("url", url), zap.Error(err))
		return nil, err
	}
	if e == nil {
		lib.GetLogger(ctx).Debug("Lookup found nothing", zap.String("url", url))
		return nil, nil
	}
	return e, nil
}

func RouteRequest(ctx context.Context, req *http.Request) api.Response {
	host, _, err := net.SplitHostPort(req.Host)
	if err != nil {
		return api.ErrorResponse(err)
	}

	// Normalize the url
	url := *req.URL
	url.Scheme = "https"
	url.Host = host
	url = lib.NormalizeURL(url)

	// Build the Url of the root object for traversal
	rootUrl := url
	rootUrl.Path = ""
	rootUrl.RawPath = rootUrl.EscapedPath()

	// Grab the root object to start traversing from
	store := entities.GetStore(ctx)
	root, err := store.GetByID(rootUrl.String())

	if err != nil {
		return api.ErrorResponse(err)
	}

	path := strings.Split(url.Path, "/")
	tc, err := api.TraverseFrom(ctx, root, path)
	if err != nil {
		return api.ErrorResponse(err)
	}

	ctx = api.WithTraversalContext(ctx, tc)
	return tc.FoundHandler.HandleRequest(ctx, req)
}
