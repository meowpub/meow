package server

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Safety checks a redirect URI to prevent open redirection vulnerabilities.
func SanitizeRedirectURI(rawuri string, req *http.Request) (string, error) {
	if rawuri == "" {
		return "", nil
	}

	u, err := url.ParseRequestURI(rawuri)
	if err != nil {
		return "", err
	}

	if u.Host != "" || u.Scheme != "" {
		return "", errors.Errorf("absolute urls not allowed in redirect uris: %s", u.String())
	}

	return u.String(), nil
}
