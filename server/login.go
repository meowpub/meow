package server

import (
	"context"
	"net/http"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/models/entities"
	"github.com/meowpub/meow/server/api"
	"github.com/meowpub/meow/server/middleware"
	"github.com/meowpub/meow/server/oauth"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var ErrInvalidUsernameOrPassword = api.Wrap(
	errors.New("invalid username or password"),
	http.StatusUnauthorized,
)

type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func HandleLogin(ctx context.Context, req *http.Request) api.Response {
	// If the user is already logged in, redirect to their profile.
	if tok := oauth.GetToken(ctx); tok != nil {
		uid := tok.UserID
		stores := models.GetStores(ctx)
		usr, err := stores.Users().GetBySnowflake(uid)
		if err != nil {
			return api.ErrorResponse(err)
		}

		entities := entities.GetStore(ctx)
		usrEntity, err := entities.GetBySnowflake(usr.EntityID)
		if err != nil {
			return api.ErrorResponse(err)
		}

		return api.RedirectResponse(usrEntity.GetID())
	}

	if req.Method == "POST" {
		var body LoginRequest
		if err := api.DecodeForm(req, &body); err != nil {
			return api.Response{Error: err}
		}

		stores := models.GetStores(ctx)
		sess := middleware.GetSession(ctx)

		// Look up the user.
		user, err := stores.Users().GetByEmail(body.Email)
		if err != nil {
			lib.GetLogger(ctx).Info("Login attempt with invalid email",
				zap.String("email", body.Email))
			return api.Response{Error: ErrInvalidUsernameOrPassword}
		}

		// Validate the password.
		ok, err := user.CheckPassword(body.Password)
		if err != nil {
			lib.GetLogger(ctx).Error("Password check failed, is the password corrupt in the DB?",
				zap.String("email", body.Email))
			return api.Response{Error: err}
		}
		if !ok {
			lib.GetLogger(ctx).Info("Login attempt with invalid password",
				zap.String("email", body.Email))
			return api.Response{Error: ErrInvalidUsernameOrPassword}
		}

		// Generate and store an access token.
		token, err := lib.GenToken()
		if err != nil {
			return api.Response{Error: err}
		}
		tok := &models.AccessToken{
			Token:    token,
			ClientID: oauth.WebClient.GetId(),
			Scope:    "web",
			AuthorizationUserData: models.AuthorizationUserData{
				UserID: user.ID,
			},
		}
		if err := stores.AccessTokens().Set(tok, oauth.WebAccessTokenTTL); err != nil {
			return api.Response{Error: err}
		}

		// Store the access token in the session.
		sess.Values["access_token"] = token

		return api.Response{Data: "OK"}
	}

	return api.Response{Template: "login"}
}
