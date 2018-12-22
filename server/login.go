package server

import (
	"net/http"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/server/api"
	"github.com/meowpub/meow/server/oauth"
	"go.uber.org/zap"
)

var ErrInvalidUsernameOrPassword = lib.Error(
	http.StatusUnauthorized,
	"invalid username or password",
)

type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`

	RedirectURI string `form:"redirect_uri"`
}

type LoginResponse struct {
	RedirectURI string `form:"redirect_uri"`
}

func HandleLogin(req api.Request) api.Response {
	L := lib.GetLogger(req)
	stores := models.GetStores(req)

	// If the user is already logged in, redirect them to their user page.
	if tok := oauth.GetToken(req); tok != nil {
		user, err := stores.Users().GetByEntityID(tok.UserID)
		if err != nil {
			if err := oauth.LogOut(req); err != nil {
				return api.Response{Error: err}
			}
			if !models.IsNotFound(err) {
				return api.Response{Error: err}
			}
		} else {
			entity, err := stores.Entities().GetBySnowflake(user.EntityID)
			if err != nil {
				if err := oauth.LogOut(req); err != nil {
					return api.Response{Error: err}
				}
				if !models.IsNotFound(err) {
					return api.Response{Error: err}
				}
			} else {
				return api.RedirectResponse(entity.Obj.ID())
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
			L.Info("Login attempt with invalid email",
				zap.String("email", body.Email),
				zap.Error(err),
			)
			return api.Response{Error: ErrInvalidUsernameOrPassword}
		}

		// Look up the user's profile.
		profile, err := stores.Entities().GetBySnowflake(user.EntityID)
		if err != nil {
			L.Error("Failed to look up user's profile",
				zap.Int64("user_id", user.ID.Int64()),
				zap.Int64("entity_id", user.EntityID.Int64()),
				zap.Error(err),
			)
			return api.Response{Error: err}
		}

		// Validate the password.
		ok, err := user.CheckPassword(body.Password)
		if err != nil {
			L.Error("Password check failed, is the password corrupt in the DB?",
				zap.String("id", profile.Obj.ID()),
				zap.Int64("user_id", user.ID.Int64()),
				zap.Error(err),
			)
			return api.Response{Error: err}
		}
		if !ok {
			L.Warn("Login attempt with invalid password",
				zap.String("id", profile.Obj.ID()),
				zap.Int64("user_id", user.ID.Int64()),
				zap.Error(err),
			)
			return api.Response{Error: ErrInvalidUsernameOrPassword}
		}

		// Log the user in.
		if err := oauth.LogIn(req, user); err != nil {
			L.Error("User authenticated successfully, but session login failed",
				zap.String("id", profile.Obj.ID()),
				zap.Int64("user_id", user.ID.Int64()),
				zap.Int64("profile_id", user.EntityID.Int64()),
				zap.Error(err),
			)
			return api.Response{Error: err}
		}

		L.Info("User logged in!",
			zap.String("id", profile.Obj.ID()),
			zap.Int64("user_id", user.ID.Int64()),
			zap.Int64("entity_id", user.EntityID.Int64()),
		)

		// Redirect to the user's profile, or to a given redirect URI *if it's safe*.
		redirectURI := profile.Obj.ID()
		if body.RedirectURI != "" {
			uri, err := SanitizeRedirectURI(body.RedirectURI)
			if err != nil {
				L.Warn("Ignoring unsafe redirect URI",
					zap.String("uri", body.RedirectURI),
					zap.Error(err))
			} else {
				L.Debug("Using redirect URI from request",
					zap.String("raw_uri", body.RedirectURI),
					zap.String("uri", uri))
				redirectURI = uri
			}
		}

		L.Debug("Redirecting after login...",
			zap.String("redirect_uri", redirectURI))
		return api.RedirectResponse(redirectURI)
	}

	// Decode the query.
	var body LoginRequest
	if err := api.DecodeQuery(req, &body); err != nil {
		return api.Response{Error: err}
	}
	return api.Response{Template: "login", Data: LoginResponse{
		RedirectURI: body.RedirectURI,
	}}
}
