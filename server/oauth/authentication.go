package oauth

import (
	"context"
	"strings"

	"go.uber.org/zap"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/server/api"
	"github.com/meowpub/meow/server/middleware"
)

const BearerPrefix = "Bearer "

const AccessTokenSessionKey = "access_token"

const accessTokenKey ctxKey = "AccessToken"

// Authenticated retrieves the user credentials for the request (if any)
func Authenticate(req api.Request) (*models.AccessToken, error) {
	l := lib.GetLogger(req)

	// Pull an access token from either the Authorization header (with the "Bearer " prefix),
	// or failing that, from the "access_token" session variable (no prefix).
	accessToken := strings.TrimPrefix(req.Header.Get("Authorization"), BearerPrefix)
	if accessToken == "" {
		accessToken, _ = middleware.GetSession(req).Values["access_token"].(string)
	}
	if accessToken == "" {
		return nil, nil
	}

	// Look up the token data, reject as unauthorized if not found.
	token, err := models.GetStores(req).AccessTokens().Get(accessToken)
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
	tok, _ := ctx.Value(accessTokenKey).(*models.AccessToken)
	return tok
}

func AuthenticationMiddleware(next api.Handler) api.Handler {
	return api.HandlerFunc(func(req api.Request) api.Response {
		tok, err := Authenticate(req)
		if err != nil {
			return api.ErrorResponse(err)
		}

		req = req.WithContext(WithToken(req.Context(), tok))
		return next.HandleRequest(req)
	})
}
