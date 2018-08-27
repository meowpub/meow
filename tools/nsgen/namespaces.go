package main

type Namespace struct {
	Package string
	Short   string
	Long    string
	Source  string
}

var Namespaces = []*Namespace{
	{
		Package: "rdf",
		Short:   "rdf",
		Long:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
		Source:  "https://www.w3.org/1999/02/22-rdf-syntax-ns",
	},
	{
		Package: "rdf",
		Short:   "rdfs",
		Long:    "http://www.w3.org/2000/01/rdf-schema#",
		Source:  "https://www.w3.org/2000/01/rdf-schema",
	},
	{
		Package: "owl",
		Short:   "owl",
		Long:    "http://www.w3.org/2002/07/owl#",
		Source:  "http://www.w3.org/2002/07/owl",
	},
	{
		Package: "as",
		Short:   "as",
		Long:    "http://www.w3.org/ns/activitystreams#",
		Source:  "https://raw.githubusercontent.com/w3c/activitystreams/master/vocabulary/activitystreams2.owl",
	},
	{
		Package: "ldp",
		Short:   "ldp",
		Long:    "http://www.w3.org/ns/ldp#",
		Source:  "http://www.w3.org/ns/ldp.ttl",
	},

	// These aren't available in Turtle anywhere, so we have to implement them manually.
	// Well, XMLSchema is available in XML schema and... no. Just, no. Not parsing that.
	// {Package: "xsd", Short: "xsd", Long: "http://www.w3.org/2001/XMLSchema#"},
	// {Package: "ldp", Short: "ldp", Long: "http://www.w3.org/ns/ldp#"},
}
