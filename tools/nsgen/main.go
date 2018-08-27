package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/deiu/rdf2go"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"golang.org/x/tools/imports"

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
	data, err := imports.Process(path, data, nil)
	if err != nil {
		return errors.Wrap(err, "goimports")
	}
	return ioutil.WriteFile(path, data, 0644)
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

func GetAndCreatePackagePath(pkg string) (string, error) {
	outdir := filepath.Join(MeowBasePath, "ld", "ns", pkg)
	return outdir, os.MkdirAll(outdir, 0755)
}

func NamespaceFragments(ns *Namespace) ([]*ld.Object, error) {
	pkgdir, err := GetAndCreatePackagePath(ns.Package)
	if err != nil {
		return nil, err
	}

	// Load and store the turtle data.
	turtlePath := filepath.Join(pkgdir, ns.Short+".ttl")
	turtleData, err := LoadWithCache(ns.Source, turtlePath)
	if err != nil {
		return nil, err
	}

	// Turn it into JSON-LD fragments.
	ldData, err := TurtleToJSONLD(ns.Long, turtleData)
	fragments, err := ld.NewObjects(ldData)
	if err != nil {
		return nil, err
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

	return fragments, DumpJSON(filepath.Join(pkgdir, ns.Short+".ld.json"), fragments)
}

func Main() error {
	var fragments []*ld.Object
	for _, ns := range Namespaces {
		if ns.Source == "" {
			continue
		}
		log.Printf("Loading: %s: %s", ns.Short, ns.Long)
		nsFragments, err := NamespaceFragments(ns)
		if err != nil {
			return errors.Wrap(err, ns.Short)
		}
		fragments = append(fragments, nsFragments...)
	}

	// Reassemble the fragments into usable Declarations.
	log.Printf("Reassembling fragments...")
	declMap := make(map[string]*Declaration)
	orderedKeys := []string{}
	for _, fragment := range fragments {
		key := fragment.ID()
		if obj, ok := declMap[key]; ok {
			if err := obj.Apply(fragment, true); err != nil {
				return err
			}
		} else {
			declMap[key] = &Declaration{fragment}
			orderedKeys = append(orderedKeys, key)
		}
	}
	declarations := make([]*Declaration, 0, len(declMap))
	for _, key := range orderedKeys {
		declarations = append(declarations, declMap[key])
	}

	// Generate packages!
	pkgs := make(map[string][]*Declaration)
	var namespaces []*Namespace
	for _, ns := range Namespaces {
		namespaces = append(namespaces, ns)
		for _, decl := range declarations {
			if strings.HasPrefix(decl.ID(), ns.Long) {
				pkgs[ns.Package] = append(pkgs[ns.Package], decl)
			}
		}
	}
	var errs []error
	for pkg, pkgdecls := range pkgs {
		outdir, err := GetAndCreatePackagePath(pkg)
		if err != nil {
			return errors.Wrap(err, pkg)
		}
		log.Printf("Generating package: %s: %s", pkg, outdir)
		rctx := &RenderContext{
			Package:      pkg,
			Declarations: pkgdecls,
			Packages:     pkgs,
			Namespaces:   namespaces,
		}
		errs = append(errs,
			errors.Wrap(Render(filepath.Join(outdir, "ns.gen.go"), NSTemplate, rctx), "ns.gen.go"),
			errors.Wrap(Render(filepath.Join(outdir, "classes.gen.go"), ClassesTemplate, rctx), "classes.gen.go"),
			errors.Wrap(Render(filepath.Join(outdir, "datatypes.gen.go"), DataTypesTemplate, rctx), "datatypes.gen.go"),
		)
	}
	// if err := multierr.Combine(errs...); err != nil {
	// 	return err
	// }

	// log.Printf("Generating index...")
	// indexPath := filepath.Join(MeowBasePath, "ld", "ns", "index.gen.go")
	// return Render(indexPath, IndexTemplate, Namespaces)
	return multierr.Combine(errs...)
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	}
}
