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

	"github.com/liclac/meow/config"
	"github.com/liclac/meow/config/secrets"
	"github.com/liclac/meow/lib"
	"github.com/liclac/meow/models/entities"
	"github.com/liclac/meow/server/api"
	"github.com/liclac/meow/server/middleware"
	"github.com/liclac/meow/server/oauth"
)

// New returns a new API router.
func New(db *gorm.DB, r *redis.Client, keyspace string) http.Handler {
	mux := api.NewRouter(render.Options{
		Directory:     "templates",
		Extensions:    []string{".html"},
		IndentJSON:    true,
		IsDevelopment: !config.IsProd(),
	})
	mux.Use(middleware.AddDB(db))
	mux.Use(middleware.AddRedis(r))
	mux.Use(middleware.AddStores())
	mux.Use(middleware.AddSession(sessions.NewCookieStore(secrets.SessionKey())))
	mux.Use(oauth.AuthenticationMiddleware)

	mux.GET("/", RouteRequest)
	mux.ANY("/favicon.ico", HandleNotFound)
	mux.GET("/-/login", RenderLogin)
	mux.POST("/-/login", HandleLogin)
	mux.NotFound(RouteRequest)

	return mux
}

func HandleNotFound(ctx context.Context, req *http.Request) api.Response {
	return api.Response{Error: api.Wrap(errors.New("not found"), http.StatusNotFound)}
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
