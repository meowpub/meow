package ld

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/piprate/json-gold/ld"
	"github.com/pkg/errors"
)

const (
	// Recommended Accept from the AS2 spec; json-gold defaults to the more inclusive:
	// "application/ld+json, application/json;q=0.9, application/javascript;q=0.5,
	// text/javascript;q=0.5, text/plain;q=0.2, */*;q=0.1"
	acceptHeader = `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`
)

type documentLoader struct {
	ctx context.Context
}

// Note:
//  * Do not mark changing documents as known (e.g. the as2 context)
//  * If a document has good caching, it is not important to include here
var knownDocuments = map[string]string{
	"https://w3id.org/security/v1": "security-v1.json",
	"https://w3id.org/security/v2": "security-v2.json",
}

// Returns a DocumentLoader for the json-gold library. It's basically a simplified version of
// the default DocumentLoader, which can't be tricked into touching the local filesystem, and
// with the addition of proper request contexts.
func NewDocumentLoader(ctx context.Context) ld.DocumentLoader {
	return documentLoader{ctx: ctx}
}

func (l documentLoader) LoadDocument(u string) (*ld.RemoteDocument, error) {
	box := rice.MustFindBox("contexts")
	if name, ok := knownDocuments[u]; ok {
		var obj interface{}

		buf := box.MustBytes(name)
		if err := json.Unmarshal(buf, &obj); err != nil {
			return nil, err
		}
		doc := &ld.RemoteDocument{
			DocumentURL: u,
			Document:    obj,
			ContextURL:  "",
		}
		return doc, nil
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(l.ctx)
	req.Header.Set("Accept", acceptHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		// Remember to drain the response body, or H/2 and connection pooling break horribly!
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, ld.NewJsonLdError(ld.LoadingDocumentFailed,
			errors.Errorf("fetching remote document returned HTTP %d %s: %s",
				resp.StatusCode, resp.Status, req.URL.String()))
	}

	doc, err := ld.DocumentFromReader(resp.Body)
	if err != nil {
		return nil, ld.NewJsonLdError(ld.LoadingDocumentFailed, err)
	}
	return &ld.RemoteDocument{
		DocumentURL: resp.Request.URL.String(), // Final URL after redirections
		Document:    doc,
	}, nil
}

func Options(ctx context.Context, uri string) *ld.JsonLdOptions {
	opts := ld.NewJsonLdOptions(uri)
	opts.DocumentLoader = NewDocumentLoader(ctx)
	return opts
}

func Expand(ctx context.Context, doc map[string]interface{}, uri string) (map[string]interface{}, error) {
	proc := ld.NewJsonLdProcessor()
	opts := Options(ctx, uri)

	res, err := proc.Expand(doc, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "expanding '%s'", uri)
	}

	if len(res) == 1 {
		return res[0].(map[string]interface{}), nil
	} else {
		return nil, errors.Errorf("Expansion of '%s' returned an array of length %d (need 1)", uri, len(res))
	}
}

func Compact(ctx context.Context, doc map[string]interface{}, uri string, context interface{}) (map[string]interface{}, error) {
	proc := ld.NewJsonLdProcessor()
	opts := Options(ctx, uri)

	arr := []interface{}{doc}
	res, err := proc.Compact(arr, context, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "compacting '%s'", uri)
	}

	res["@context"] = context

	if err != nil {
		return nil, errors.Wrapf(err, "marshalling '%s' after compaction", uri)
	} else {
		return res, nil
	}
}
