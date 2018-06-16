package oauth

import (
	"context"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/server/middleware"
)

// Sets an access token in the session for the specified user.
func LogIn(ctx context.Context, user *models.User) error {
	stores := models.GetStores(ctx)

	// Generate and store an access token.
	token, err := lib.GenToken()
	if err != nil {
		return err
	}
	tok := &models.AccessToken{
		Token:    token,
		ClientID: WebClient.GetId(),
		Scope:    "web",
		AuthorizationUserData: models.AuthorizationUserData{
			UserID: user.ID,
		},
	}
	if err := stores.AccessTokens().Set(tok, WebAccessTokenTTL); err != nil {
		return err
	}

	// Stick it in the session.
	middleware.GetSession(ctx).Values[AccessTokenSessionKey] = tok.Token

	return nil
}

// Clears and invalidates the access token in the session, if any.
func LogOut(ctx context.Context) error {
	stores := models.GetStores(ctx)
	sess := middleware.GetSession(ctx)

	// Retrieve the access token from the session, then drop it.
	tok, _ := sess.Values[AccessTokenSessionKey].(string)
	delete(sess.Values, AccessTokenSessionKey)

	// Invalidate it.
	return stores.AccessTokens().Delete(tok)
}
