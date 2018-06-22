package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/deiu/rdf2go"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

const AS2SpecURL = "https://raw.githubusercontent.com/w3c/activitystreams/master/vocabulary/activitystreams2.owl"

var flagOutDir = pflag.StringP("out-dir", "O", "./models/entities", "output directory for generated files")

var (
	topLevel = rdf2go.NewBlankNode(0)
)

func Main() error {
	pflag.Parse()

	// Resolve outdir path, make sure it exists.
	outDir, err := filepath.Abs(*flagOutDir)
	if err != nil {
		return errors.Wrapf(err, "couldn't canonicalize output directory (%s)", *flagOutDir)
	}
	if _, err := os.Stat(outDir); err != nil {
		return errors.Wrapf(err, "couldn't stat output directory")
	}

	// Load the ActivityStreams 2 RDF spec from github.
	log("Loading the ActivityStreams 2 spec...")
	as2Data, err := load(AS2SpecURL)
	if err != nil {
		return errors.Wrap(err, "couldn't load ActivityStreams 2 spec")
	}
	log("Parsing the ActivityStreams 2 spec...")

	// Parse it into an RDF graph.
	as2 := rdf2go.NewGraph(AS2SpecURL)
	if err := as2.Parse(bytes.NewReader(as2Data), "text/turtle"); err != nil {
		return errors.Wrap(err, "couldn't read ActivityStreams 2 spec")
	}
	log("-> ActivityStreams 2 spec loaded!")

	// Build up a graph of types.
	classes := map[string]*Class{}
	dtypes := map[string]*DataType{}
	props := map[string]*Property{}
	for _, triple := range as2.All(
		nil,
		rdf2go.NewResource(RDF_type),
		nil,
	) {
		if triple.Subject.Equal(topLevel) || triple.Subject.RawValue() == RDF_nil {
			continue
		}

		n := Node{
			ID:    triple.Subject.RawValue(),
			Attrs: make(map[string][]string),
		}
		for _, tp := range as2.All(triple.Subject, nil, nil) {
			attr := tp.Predicate.RawValue()
			n.Attrs[attr] = append(n.Attrs[attr], tp.Object.RawValue())
		}

		n.Types = n.PluckAll(RDF_type)
		n.Label = n.Pluck(RDFS_label)
		n.Comment = n.Pluck(RDFS_comment)

		thisType := triple.Object.RawValue()
		switch thisType {
		case RDFS_Datatype:
			dtypes[n.ID] = &DataType{
				Node: n,
			}
		case OWL_Class:
			classes[n.ID] = &Class{
				Node:       n,
				subClassOf: n.Pluck(RDFS_subClassOf),
			}
		case OWL_DatatypeProperty,
			OWL_ObjectProperty:
			props[n.ID] = &Property{
				Node:               n,
				Domain:             n.Pluck(RDFS_domain),
				Range:              n.PluckAll(RDFS_range),
				SubPropertyOf:      n.Pluck(RDFS_subPropertyOf),
				EquivalentProperty: n.PluckAll(OWL_equivalentProperty),
			}
		case OWL_FunctionalProperty,
			OWL_DeprecatedProperty,
			RDF_Statement,
			OWL_Ontology:
			continue
		default:
			return errors.Errorf("unknown type: %s; %v", n.Types, n)
		}

		for k, v := range n.Attrs {
			switch k {
			default:
				log("? unknown attr %s = %v, on %s (%s)", k, v, n.ID, n.Types)
			}
		}
	}

	// Resolve references.
	for _, cls := range classes {
		if cls.subClassOf != "" && cls.subClassOf != "0" {
			cls.SubClassOf = classes[cls.subClassOf]
			if cls.SubClassOf == nil {
				return errors.Errorf("%s has unknown parent type: %s", cls.ID, cls.SubClassOf)
			}
		}
	}

	for _, prop := range props {
		if prop.SubPropertyOf != "" && prop.SubPropertyOf != "0" {
			if parent, ok := props[prop.SubPropertyOf]; ok {
				parent.SubProperties = append(parent.SubProperties, prop)
			} else {
				log("! WARN: Prop %s has unknown parent type: %s", prop.ID, prop.SubPropertyOf)
			}
		}
		if prop.Domain != "0" {
			if dt, ok := dtypes[prop.Domain]; ok {
				dt.Properties = append(dt.Properties, prop)
			} else if cls, ok := classes[prop.Domain]; ok {
				cls.Properties = append(cls.Properties, prop)
			} else {
				return errors.Errorf("datatype property on nonexistent type: %s: %v", prop.Domain, prop)
			}
		}
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(map[string]interface{}{
		"Classes":   classes,
		"DataTypes": dtypes,
	}); err != nil {
		return err
	}

	return render(filepath.Join(outDir, "types_as2.gen.go"), AS2GenTmpl, classes)
}

func render(outfile string, tmpl *template.Template, data interface{}) error {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return ioutil.WriteFile(outfile, src, 0644)
}

func pluck(attrs map[string][]string, key string) []string {
	value := attrs[key]
	delete(attrs, key)
	return value
}

func pad(str string, width int) string {
	return str + strings.Repeat(" ", width-len(str))
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintf(os.Stderr, "=> ERROR: %s\n", err)
		os.Exit(1)
	}
}
