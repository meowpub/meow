package middleware

import (
	"net"
	"net/http"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/server/api"
	"go.uber.org/zap"
)

func LocalDomains(domains []string) func(next api.Handler) api.Handler {
	return func(next api.Handler) api.Handler {
		return api.HandlerFunc(func(req api.Request) api.Response {
			hostname, _, _ := net.SplitHostPort(req.Host)
			if hostname == "" {
				hostname = req.Host
			}
			isLocal := false
			for _, domain := range domains {
				if hostname == domain {
					isLocal = true
					break
				}
			}
			if !isLocal {
				zap.L().Warn("rejecting request to unknown domain",
					zap.String("domain", hostname), zap.Strings("local", domains))
				return api.ErrorResponse(lib.Errorf(http.StatusMisdirectedRequest,
					"The server was not configured to handle the domain: '%s'", hostname))
			}
			return next.HandleRequest(req)
		})
	}
}
