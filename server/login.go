package server

import (
	"context"
	"net/http"

	"github.com/liclac/meow/lib"
	"github.com/liclac/meow/models"
	"github.com/liclac/meow/models/entities"
	"github.com/liclac/meow/server/api"
	"github.com/liclac/meow/server/middleware"
	"github.com/liclac/meow/server/oauth"
	"github.com/pkg/errors"
)

type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func RenderLogin(ctx context.Context, req *http.Request) api.Response {
	tok := oauth.GetToken(ctx)

	if tok != nil {
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

	return api.Response{Template: "login"}
}

func HandleLogin(ctx context.Context, req *http.Request) api.Response {
	var body LoginRequest
	if err := api.DecodeForm(req, &body); err != nil {
		return api.Response{Error: err}
	}

	stores := models.GetStores(ctx)
	sess := middleware.GetSession(ctx)

	// Look up the user.
	user, err := stores.Users().GetByEmail(body.Email)
	if err != nil {
		return api.Response{Error: err}
	}

	// Validate the password.
	ok, err := user.CheckPassword(body.Password)
	if err != nil {
		return api.Response{Error: err}
	}
	if !ok {
		err := errors.New("invalid password")
		return api.Response{Error: api.Wrap(err, http.StatusUnauthorized)}
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
