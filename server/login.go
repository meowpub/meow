package server

import (
	"context"
	"net/http"

	"github.com/liclac/meow/server/api"
)

type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func RenderLogin(ctx context.Context, req *http.Request) api.Response {
	return api.Response{Template: "login"}
}

func HandleLogin(ctx context.Context, req *http.Request) api.Response {
	return RenderLogin(ctx, req)
}
