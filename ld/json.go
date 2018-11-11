package ld

import (
	"encoding/json"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/piprate/json-gold/ld"
	"github.com/pkg/errors"
)

type documentLoader struct {
	def *ld.DefaultDocumentLoader
}

// Note:
//  * Do not mark changing documents as known (e.g. the as2 context)
//  * If a document has good caching, it is not important to include here
var knownDocuments = map[string]string{
	"https://w3id.org/security/v1": "security-v1.json",
	"https://w3id.org/security/v2": "security-v2.json",
}

func NewDocumentLoader(cli *http.Client) ld.DocumentLoader {
	return documentLoader{
		def: ld.NewDefaultDocumentLoader(cli),
	}
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
	} else {
		return l.def.LoadDocument(u)
	}

}

func Options(cli *http.Client, uri string) *ld.JsonLdOptions {
	opts := ld.NewJsonLdOptions(uri)
	opts.DocumentLoader = NewDocumentLoader(cli)
	return opts
}

func Expand(cli *http.Client, doc map[string]interface{}, uri string) (map[string]interface{}, error) {
	proc := ld.NewJsonLdProcessor()
	opts := Options(cli, uri)

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

func Compact(cli *http.Client, doc map[string]interface{}, uri string, context interface{}) (map[string]interface{}, error) {
	proc := ld.NewJsonLdProcessor()
	opts := Options(cli, uri)

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
