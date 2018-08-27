package well_known

import (
	"net/url"

	"github.com/meowpub/meow/lib/xrd"
	"github.com/meowpub/meow/server/api"
)

func HostMetaHandler(req api.Request) api.Response {
	baseUrl := url.URL{
		Scheme: "https",
		Host:   req.Host,
		Path:   "/",
	}

	xrd := &xrd.XRD{
		ID:      baseUrl.String(),
		Subject: baseUrl.String(),
		Properties: map[string]string{
			"https://github.com/meowpub/meow": "data:text/plain;base64,bWxlbQ",
		},
		Links: []xrd.Link{
			xrd.Link{
				Rel:  "lrdd",
				Type: "application/xrd+xml",
				Template: baseUrl.ResolveReference(&url.URL{
					Path:     "/.well-known/webfinger.xrd",
					RawQuery: "resource={uri}",
				}).String(),
			},
			xrd.Link{
				Rel:  "lrdd",
				Type: "application/jrd+json",
				Template: baseUrl.ResolveReference(&url.URL{
					Path:     "/.well-known/webfinger.jrd",
					RawQuery: "resource={uri}",
				}).String(),
			},
		},
	}

	return RenderXRD(req, xrd, "xrd")
}
