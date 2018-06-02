package oauth

import (
	"context"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/liclac/meow/lib"
	"github.com/liclac/meow/models"
	"github.com/liclac/meow/server/api"
	"github.com/liclac/meow/server/middleware"
)

const BearerPrefix = "Bearer "

// Authorize retrieves the access token from the request and ensures that it has the requested
// scope (if any). If the user is not authenticated, the returned error will have a status code
// of 401 Unauthorized, or if the user does not have the required scope, 403 Forbidden.
func Authorize(ctx context.Context, req *http.Request, scope *Scope) (*models.AccessToken, error) {
	l := lib.GetLogger(ctx)

	// Pull an access token from either the Authorization header (with the "Bearer " prefix),
	// or failing that, from the "access_token" session variable (no prefix).
	accessToken := strings.TrimPrefix(req.Header.Get("Authorization"), BearerPrefix)
	if accessToken == "" {
		accessToken, _ = middleware.GetSession(ctx).Values["access_token"].(string)
	}
	if accessToken == "" {
		return nil, api.Wrap(errors.New("authentication required"), http.StatusUnauthorized)
	}

	// Look up the token data, reject as unauthorized if not found.
	token, err := models.GetStores(ctx).AccessTokens().Get(accessToken)
	if err != nil {
		if !models.IsNotFound(err) {
			l.Error("Access Token lookup failed", zap.Error(err), zap.String("token", accessToken))
			return nil, err
		}
		l.Debug("Access Token is invalid/expired", zap.Error(err), zap.String("token", accessToken))
		return nil, api.Wrap(
			errors.New("authentication required"),
			http.StatusUnauthorized,
		)
	}

	// If we're authorizing with a scope, reject as unauthorized if we don't have it.
	scopes := strings.Split(token.Scope, " ")
	if scope != nil && !CheckScope(scopes, scope) {
		return nil, api.Wrap(
			errors.Errorf("this endpoint requires the %s scope, but you only have grants for %s", strings.Join(scopes, ", ")),
			http.StatusForbidden,
		)
	}

	return token, nil
}
