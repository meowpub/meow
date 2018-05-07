package jsonld

import (
	"encoding/json"
	"github.com/GeertJohan/go.rice"
	"github.com/piprate/json-gold/ld"
	"github.com/pkg/errors"
	"net/http"
)

type documentLoader struct {
	def *ld.DefaultDocumentLoader
}

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

func Expand(cli *http.Client, doc []byte, uri string) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	opts := Options(cli, uri)

	var obj map[string]interface{}
	if err := json.Unmarshal(doc, &obj); err != nil {
		return nil, errors.Wrapf(err, "unmarshalling '%s' for expansion", uri)
	}

	res, err := proc.Expand(obj, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "expanding '%s'", uri)
	}

	if len(res) == 1 {
		return json.Marshal(res[0])
	} else {
		return nil, errors.Errorf("Expansion of '%s' returned an array of length %d (need 1)", uri, len(res))
	}
}

func Compact(cli *http.Client, doc []byte, uri string, context interface{}) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	opts := Options(cli, uri)

	var obj map[string]interface{}
	if err := json.Unmarshal(doc, &obj); err != nil {
		return nil, errors.Wrapf(err, "unmarshalling '%s' for compaction", uri)
	}

	arr := []interface{}{obj}
	res, err := proc.Compact(arr, context, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "compacting '%s'", uri)
	}

	res["@context"] = context

	buf, err := json.Marshal(res)
	if err != nil {
		return nil, errors.Wrapf(err, "marshalling '%s' after compaction", uri)
	} else {
		return buf, nil
	}
}
