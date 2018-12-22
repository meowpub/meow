package middleware

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/server/api"
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
		return api.HandlerFunc(func(req api.Request) api.Response {
			sess, err := store.Get(req.Request, "session")
			if err != nil {
				return api.Response{Error: lib.Code(err, http.StatusBadRequest)}
			}
			req = req.WithContext(WithSession(req.Context(), sess))
			resp := next.HandleRequest(req)
			if err := sess.Save(req.Request, &resp); err != nil {
				return api.Response{Error: lib.Code(err, http.StatusBadRequest)}
			}
			return resp
		})
	}
}
