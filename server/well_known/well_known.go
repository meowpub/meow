package well_known

import (
	"errors"
	"strings"

	"github.com/meowpub/meow/lib/xrd"
	"github.com/meowpub/meow/server/api"
)

var WellKnown = api.Node{
	Self: nil,
	Children: map[string]api.Handler{
		"host-meta":      api.HandlerFunc(HostMetaHandler),
		"host-meta.xml":  api.HandlerFunc(HostMetaHandler),
		"host-meta.xrd":  api.HandlerFunc(HostMetaHandler),
		"host-meta.json": api.HandlerFunc(HostMetaHandler),
		"host-meta.jrd":  api.HandlerFunc(HostMetaHandler),
		"webfinger":      api.HandlerFunc(WebFingerHandler),
		"webfinger.xml":  api.HandlerFunc(WebFingerHandler),
		"webfinger.xrd":  api.HandlerFunc(WebFingerHandler),
		"webfinger.json": api.HandlerFunc(WebFingerHandler),
		"webfinger.jrd":  api.HandlerFunc(WebFingerHandler),
	},
}

func RenderXRD(req api.Request, xrd *xrd.XRD, defaultFormat string) api.Response {
	format := defaultFormat

	switch {
	case strings.HasSuffix(req.URL.Path, ".json"),
		strings.HasSuffix(req.URL.Path, ".jrd"):
		format = "jrd"

	case strings.HasSuffix(req.URL.Path, ".xml"),
		strings.HasSuffix(req.URL.Path, ".xrd"):
		format = "xrd"
	}

	switch format {
	case "xrd":
		data, err := xrd.MarshalIndentXRD()
		if err != nil {
			return api.ErrorResponse(err)
		}

		r := api.Response{Data: data}
		r.Header().Add("Content-Type", "application/xml")
		return r

	case "jrd":
		data, err := xrd.MarshalIndentJRD()
		if err != nil {
			return api.ErrorResponse(err)
		}

		r := api.Response{Data: data}
		r.Header().Add("Content-Type", "application/json")
		return r

	default:
		return api.ErrorResponse(errors.New("Unknown format " + format))
	}
}
