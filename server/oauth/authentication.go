package oauth

import (
	"context"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/server/api"
	"github.com/meowpub/meow/server/middleware"
)

const BearerPrefix = "Bearer "
const accessTokenKey ctxKey = "AccessToken"

// Authenticated retrieves the user credentials for the request (if any)
func Authenticate(ctx context.Context, req *http.Request) (*models.AccessToken, error) {
	l := lib.GetLogger(ctx)

	// Pull an access token from either the Authorization header (with the "Bearer " prefix),
	// or failing that, from the "access_token" session variable (no prefix).
	accessToken := strings.TrimPrefix(req.Header.Get("Authorization"), BearerPrefix)
	if accessToken == "" {
		accessToken, _ = middleware.GetSession(ctx).Values["access_token"].(string)
	}
	if accessToken == "" {
		return nil, nil
	}

	// Look up the token data, reject as unauthorized if not found.
	token, err := models.GetStores(ctx).AccessTokens().Get(accessToken)
	if err != nil {
		if !models.IsNotFound(err) {
			l.Error("Access Token lookup failed", zap.Error(err), zap.String("token", accessToken))
			return nil, err
		}

		l.Debug("Access Token is invalid/expired", zap.Error(err), zap.String("token", accessToken))
		return nil, nil
	}

	return token, nil
}

func WithToken(ctx context.Context, tok *models.AccessToken) context.Context {
	return context.WithValue(ctx, accessTokenKey, tok)
}

func GetToken(ctx context.Context) *models.AccessToken {
	tok := ctx.Value(accessTokenKey)
	if tok != nil {
		return tok.(*models.AccessToken)
	} else {
		return nil
	}
}

func AuthenticationMiddleware(next api.Handler) api.Handler {
	return api.HandlerFunc(func(ctx context.Context, req *http.Request) api.Response {
		tok, err := Authenticate(ctx, req)
		if err != nil {
			return api.ErrorResponse(err)
		}

		ctx = WithToken(ctx, tok)

		return next.HandleRequest(ctx, req)
	})
}
