package server

import (
	"context"
	"net/http"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/models/entities"
	"github.com/meowpub/meow/server/api"
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
	stores := models.GetStores(ctx)

	// If the user is already logged in, redirect them to their user page.
	if tok := oauth.GetToken(ctx); tok != nil {
		user, err := stores.Users().GetByEntityID(tok.UserID)
		if err != nil {
			if err := oauth.LogOut(ctx); err != nil {
				return api.Response{Error: err}
			}
			if !models.IsNotFound(err) {
				return api.Response{Error: err}
			}
		} else {
			entity, err := entities.GetStore(ctx).GetBySnowflake(user.EntityID)
			if err != nil {
				if err := oauth.LogOut(ctx); err != nil {
					return api.Response{Error: err}
				}
				if !models.IsNotFound(err) {
					return api.Response{Error: err}
				}
			} else {
				return api.RedirectResponse(entity.GetID())
			}
		}
	}

	if req.Method == "POST" {
		var body LoginRequest
		if err := api.DecodeForm(req, &body); err != nil {
			return api.Response{Error: err}
		}

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

		// Log the user in.
		if err := oauth.LogIn(ctx, user); err != nil {
			return api.Response{Error: err}
		}

		return api.Response{Data: "OK"}
	}

	return api.Response{Template: "login"}
}
