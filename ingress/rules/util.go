package rules

import (
	"net/http"
	"net/url"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns"
	"github.com/meowpub/meow/ld/ns/meta"
	"github.com/meowpub/meow/lib"
)

// Returns whether two IDs are from the same origin.
func IsSameOrigin(a, b string) bool {
	ua, _ := url.Parse(a)
	ub, _ := url.Parse(b)
	return ua != nil && ub != nil && ua.Hostname() == ub.Hostname()
}

func Slugify(input string) (string, error) {
	s, _, err := transform.String(
		transform.Chain(
			// Normalise to NFD form, splitting things out into base forms and modifiers.
			norm.NFD,

			// Make all of it lowercase.
			runes.Map(unicode.ToLower),

			// Filter out characters to clean up the output.
			runes.Remove(runes.In(unicode.M)), // Marks, eg. the ´ in é.
			runes.Remove(runes.In(unicode.P)), // Punctuation.

			// Replace non-URL-safe characters (including whitespace) with '-'.
			runes.Map(func(r rune) rune {
				if (r >= 'a' && r <= 'z') ||
					(r >= 'A' && r <= 'Z') ||
					(r >= '0' && r <= '9') {
					return r
				}
				return '-'
			}),
		),
		input)
	return s, err
}

func Copy(from, to ld.Entity, prop *meta.Prop) {
	if v := from.Get(prop.ID); v != nil {
		to.Set(prop.ID, v)
	}
}

func ValidateQuacksLike(e ld.Entity, t *meta.Type) error {
	if !ns.QuacksLike(e, t) {
		return lib.Errorf(http.StatusBadRequest, "@type=[%s] does not include required type (or a known subtype thereof): %s", e.Type(), t.ID)
	}
	return nil
}
