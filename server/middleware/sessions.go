package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/liclac/meow/server/api"
)

const ctxKeySession ctxKey = "session"

func WithSession(ctx context.Context, sess *sessions.Session) context.Context {
	return context.WithValue(ctx, ctxKeySession, sess)
}

func GetSession(ctx context.Context) *sessions.Session {
	sess, _ := ctx.Value(ctxKeySession).(*sessions.Session)
	return sess
}

func AddSession(store sessions.Store) func(next api.Handler) api.Handler {
	return func(next api.Handler) api.Handler {
		return api.HandlerFunc(func(ctx context.Context, req *http.Request) api.Response {
			sess, err := store.Get(req, "session")
			if err != nil {
				return api.Response{Error: api.Wrap(err, http.StatusBadRequest)}
			}
			ctx = WithSession(ctx, sess)
			req = req.WithContext(ctx)
			resp := next.HandleRequest(ctx, req)
			if err := sess.Save(req, &resp); err != nil {
				return api.Response{Error: api.Wrap(err, http.StatusBadRequest)}
			}
			return resp
		})
	}
}
