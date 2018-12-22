package oauth

import (
	"context"
	"net/http"
	"strings"

	"github.com/meowpub/meow/lib"
)

// Authorize retrieves the access token from the context and ensures that it has the requested
// scope (if any). If the user is not authenticated, the returned error will have a status code
// of 401 Unauthorized, or if the user does not have the required scope, 403 Forbidden.
func Authorize(ctx context.Context, scope *Scope) error {
	l := lib.GetLogger(ctx)
	token := GetToken(ctx)

	if token == nil {
		l.Debug("No token in context")
		return lib.Error(
			http.StatusUnauthorized,
			"authentication required",
		)
	}

	// If we're authorizing with a scope, reject as unauthorized if we don't have it.
	scopes := strings.Split(token.Scope, " ")
	if scope != nil && !CheckScope(scopes, scope) {
		return lib.Errorf(
			http.StatusForbidden,
			"this endpoint requires the %s scope, but you only have grants for %s",
			scope.ID, strings.Join(scopes, ", "),
		)
	}

	return nil
}
