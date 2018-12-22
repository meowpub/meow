package well_known

import (
	"net/http"
	"net/url"
	"strings"

	"go.uber.org/zap"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/lib/clog"
	"github.com/meowpub/meow/lib/xrd"
	"github.com/meowpub/meow/server/api"
)

func WebFingerHandler(req api.Request) api.Response {
	q := req.URL.Query()
	resource := q.Get("resource")
	if resource == "" {
		return api.ErrorResponse(lib.Error(http.StatusBadRequest, "No resource specified"))
	}

	rootUrl := req.RootURL()

	resourceUrl, err := url.Parse(resource)
	if err != nil {
		return api.ErrorResponse(lib.Code(err, http.StatusBadRequest))
	}
	*resourceUrl = lib.NormalizeURL(*rootUrl.ResolveReference(resourceUrl))

	if resourceUrl.Host != rootUrl.Host {
		return api.ErrorResponse(lib.Error(http.StatusNotFound, "Resource on wrong host"))
	}

	subReq := req.WithContext(lib.WithLoggerFields(req.Context(),
		zap.Any("webfinger_resource", resourceUrl)))
	subReq.Method = "FINGER"
	subReq.URL = resourceUrl

	clog.Debug(req, "Performing WebFinger dicovery")

	var xrd *xrd.XRD

	switch resourceUrl.Scheme {
	case "https":
		origTC := api.GetTraversalContext(req)
		root := origTC.TraversedPath[0]

		tc, err := api.TraverseFrom(subReq, root, strings.Split(resourceUrl.Path, "/"))
		if err != nil {
			clog.Error(req, "Traversal error", zap.Error(err))
			return api.ErrorResponse(err)
		}

		subReq = subReq.WithContext(
			lib.WithLoggerFields(req.Context(),
				zap.Any("webfinger_traversal_context", tc)))

		if f, ok := api.UnwrapHandler(tc.FoundHandler).(api.Fingerable); ok {
			xrd, err = f.Finger(subReq)
			if err != nil {
				return api.ErrorResponse(err)
			}
		} else {
			clog.Error(subReq, "Not found: not traversible")
			return api.ErrorResponse(lib.Errorf(http.StatusNotFound, "%s: Not found", resourceUrl.String()))
		}

	default:
		clog.Error(req, "Unsupported scheme")
		return api.ErrorResponse(lib.Error(http.StatusNotFound, "Unsupported scheme"))
	}

	return RenderXRD(req, xrd, "jrd")
}
