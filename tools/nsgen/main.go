package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"text/template"

	"github.com/deiu/rdf2go"
	"github.com/pkg/errors"
	"go.uber.org/multierr"

	"github.com/meowpub/meow/ld"
)

// Base path to the github.com/meowpub/meow package.
var MeowBasePath = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "meowpub", "meow")

func LoadWithCache(uri, cachePath string) ([]byte, error) {
	if cacheData, err := ioutil.ReadFile(cachePath); err == nil {
		return cacheData, nil
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, ioutil.WriteFile(cachePath, data, 0644)
}

func DumpJSON(path string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

func Render(path string, tmpl *template.Template, rctx interface{}) error {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, rctx); err != nil {
		return errors.Wrap(err, "render")
	}
	data := buf.Bytes()
	src, err := format.Source(buf.Bytes())
	if err != nil {
		err = errors.Wrap(err, "gofmt")
		src = data
	}
	return multierr.Append(err, ioutil.WriteFile(path, src, 0644))
}

func TurtleToJSONLD(namespace string, data []byte) ([]byte, error) {
	as2 := rdf2go.NewGraph(namespace)
	if err := as2.Parse(bytes.NewReader(data), "text/turtle"); err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := as2.Serialize(&buf, "application/ld+json"); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GenNamespace(ns *Namespace) error {
	outdir := filepath.Join(MeowBasePath, "ld", ns.Short)

	// Make sure the output path exists.
	if err := os.MkdirAll(outdir, 0755); err != nil {
		return err
	}

	// Load and store the turtle data.
	turtlePath := filepath.Join(outdir, ns.Short+".ttl")
	turtleData, err := LoadWithCache(ns.Source, turtlePath)
	if err != nil {
		return err
	}

	// Turn it into JSON-LD.
	ldData, err := TurtleToJSONLD(ns.Long, turtleData)
	fragments, err := ld.NewObjects(ldData)
	if err != nil {
		return err
	}
	sort.SliceStable(fragments, func(i, j int) bool {
		idata, err := json.Marshal(fragments[i])
		if err != nil {
			panic(err)
		}
		jdata, err := json.Marshal(fragments[j])
		if err != nil {
			panic(err)
		}
		return string(idata) < string(jdata)
	})
	if err := DumpJSON(filepath.Join(outdir, ns.Short+".ld.json"), fragments); err != nil {
		return err
	}

	// Reassemble the fragments into usable Declarations.
	declMap := make(map[string]*Declaration)
	orderedKeys := []string{}
	for _, fragment := range fragments {
		key := fragment.ID()
		if obj, ok := declMap[key]; ok {
			if err := obj.Apply(fragment, true); err != nil {
				return err
			}
		} else {
			declMap[key] = &Declaration{fragment, ns}
			orderedKeys = append(orderedKeys, key)
		}
	}
	declarations := make([]*Declaration, 0, len(declMap))
	for _, key := range orderedKeys {
		isExcluded := false
		for _, exclusion := range ns.Exclude {
			if key == exclusion {
				isExcluded = true
				break
			}
		}
		if !isExcluded {
			declarations = append(declarations, declMap[key])
		}
	}

	// Generate files!
	rctx := &RenderContext{
		Namespace:    ns,
		Declarations: declarations,
	}
	return multierr.Combine(
		errors.Wrap(Render(filepath.Join(outdir, "ns.gen.go"), NSTemplate, rctx), "ns.gen.go"),
		errors.Wrap(Render(filepath.Join(outdir, "classes.gen.go"), ClassesTemplate, rctx), "classes.gen.go"),
		errors.Wrap(Render(filepath.Join(outdir, "datatypes.gen.go"), DataTypesTemplate, rctx), "datatypes.gen.go"),
	)
}

func Main() error {
	for _, ns := range Namespaces {
		if ns.Source == "" {
			continue
		}
		log.Printf("Generating: \"%s\" (%s)", ns.Short, ns.Long)
		if err := GenNamespace(ns); err != nil {
			return err
		}
	}

	log.Printf("Generating index...")
	indexPath := filepath.Join(MeowBasePath, "ld", "resolve", "index.gen.go")
	return Render(indexPath, IndexTemplate, Namespaces)
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	}
}
