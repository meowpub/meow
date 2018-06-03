package oauth

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/liclac/meow/models"
	"github.com/liclac/meow/server/api"
)

// NewRouter creates a new handler for OAuth endpoints. Mount it on a subpath.
func New(stores models.Stores) http.Handler {
	// config := osin.NewServerConfig()
	// storage := NewStorage(stores)
	// server := osin.NewServer(config, storage)

	// The library we're using makes it very easy to use regular http handlers.
	// TODO: Look into rewriting this with API scaffolding instead.
	mux := chi.NewMux()
	// mux.Get("/authorize", api.WrapHandler(api.HandlerFunc(HandleGETAuthorize)))
	// mux.Post("/authorize", func(rw http.ResponseWriter, r *http.Request) {
	// 	resp := server.NewResponse()
	// 	defer resp.Close()

	// 	server.HandleAuthorizeRequest(resp, r)
	// })
	return mux
}

func HandleGETAuthorize(ctx context.Context, req *http.Request) api.Response {
	return api.Response{Template: "oauth/authorize"}
}
