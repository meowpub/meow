package middleware

import (
	"net/http"

	"github.com/meowpub/meow/server/api"
)

func LocalDomains(domains []string) func(next api.Handler) api.Handler {
	return func(next api.Handler) api.Handler {
		return api.HandlerFunc(func(req api.Request) api.Response {
			hostname := req.URL.Hostname()
			isLocal := false
			for _, domain := range domains {
				if hostname == domain {
					isLocal = true
					break
				}
			}
			if !isLocal {
				return api.Response{Status: http.StatusMisdirectedRequest}
			}
			return next.HandleRequest(req)
		})
	}
}
