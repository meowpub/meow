package oauth

import (
	"github.com/RangelReale/osin"
	"github.com/keybase/saltpack/encoding/basex"

	"github.com/meowpub/meow/config/secrets"
	"github.com/meowpub/meow/models"
)

var WebClient osin.Client = &webClient{}

type webClient struct {
	secret string
}

func (cl *webClient) GetId() string {
	return "__web__"
}

func (cl *webClient) GetSecret() string {
	if cl.secret == "" {
		secret := secrets.WebClientSecret()
		cl.secret = basex.Base62StdEncoding.EncodeToString(secret)
	}
	return cl.secret
}

func (cl *webClient) GetRedirectUri() string {
	return ""
}

func (cl *webClient) GetUserData() interface{} {
	return models.ClientUserData{}
}
