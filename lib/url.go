package lib

import (
	"net/url"
	"strings"
)

func NormalizeURL(u url.URL) url.URL {
	// Remove query/fragment portions as we cannot route on them
	u.RawQuery = ""
	u.ForceQuery = false
	u.Fragment = ""

	// Normalize path to remove any trailing slashes
	u.Path = strings.TrimSuffix(u.Path, "/")
	u.RawPath = u.EscapedPath()

	return u
}
