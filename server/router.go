package server

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"

	"github.com/liclac/meow/config"
	"github.com/liclac/meow/lib"
	"github.com/liclac/meow/models"
	"github.com/liclac/meow/models/entities"
	"github.com/liclac/meow/server/api"
	"github.com/liclac/meow/server/middleware"
	"github.com/liclac/meow/server/oauth"
)

// New returns a new API router.
func New(db *gorm.DB, r *redis.Client, keyspace string) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.AddDB(db))
	mux.Use(middleware.AddRedis(r))
	mux.Use(middleware.AddStores())
	mux.Use(middleware.AddEntityStore())
	mux.Use(middleware.AddRender(render.New(render.Options{
		Directory:     "templates",
		Extensions:    []string{".html"},
		IndentJSON:    true,
		IsDevelopment: !config.IsProd(),
	})))

	mux.Get("/", api.WrapHandler(api.HandlerFunc(RouteRequest)))
	mux.NotFound(api.WrapHandler(api.HandlerFunc(RouteRequest)))
	mux.Mount("/oauth", oauth.New(models.NewStores(db, r, keyspace)))

	return mux
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
