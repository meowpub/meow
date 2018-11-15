// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/meta"
)

var (
	RDF = &meta.Namespace{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
		Short: "rdf",
		Props: []*meta.Prop{
			Prop_First,
			Prop_Object,
			Prop_Predicate,
			Prop_Rest,
			Prop_Subject,
			Prop_Type,
			Prop_Value,
		},
		Types: map[string]*meta.Type{
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#Alt":       Class_Alt,
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag":       Class_Bag,
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#List":      Class_List,
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#Property":  Class_Property,
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq":       Class_Seq,
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement": Class_Statement,
		},
	}
	RDFS = &meta.Namespace{
		ID:    "http://www.w3.org/2000/01/rdf-schema#",
		Short: "rdfs",
		Props: []*meta.Prop{
			Prop_Comment,
			Prop_Domain,
			Prop_IsDefinedBy,
			Prop_Label,
			Prop_Member,
			Prop_Range,
			Prop_SeeAlso,
			Prop_SubClassOf,
			Prop_SubPropertyOf,
		},
		Types: map[string]*meta.Type{
			"http://www.w3.org/2000/01/rdf-schema#Class":                       Class_Class,
			"http://www.w3.org/2000/01/rdf-schema#Container":                   Class_Container,
			"http://www.w3.org/2000/01/rdf-schema#ContainerMembershipProperty": Class_ContainerMembershipProperty,
			"http://www.w3.org/2000/01/rdf-schema#Datatype":                    Class_Datatype,
			"http://www.w3.org/2000/01/rdf-schema#Literal":                     Class_Literal,
			"http://www.w3.org/2000/01/rdf-schema#Resource":                    Class_Resource,
		},
	}

	// The first item in the subject RDF list.
	Prop_First = &meta.Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#first",
		Short:   "first",
		Comment: "The first item in the subject RDF list.",
	}

	// The object of the subject RDF statement.
	Prop_Object = &meta.Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#object",
		Short:   "object",
		Comment: "The object of the subject RDF statement.",
	}

	// The predicate of the subject RDF statement.
	Prop_Predicate = &meta.Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#predicate",
		Short:   "predicate",
		Comment: "The predicate of the subject RDF statement.",
	}

	// The rest of the subject RDF list after the first item.
	Prop_Rest = &meta.Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#rest",
		Short:   "rest",
		Comment: "The rest of the subject RDF list after the first item.",
	}

	// The subject of the subject RDF statement.
	Prop_Subject = &meta.Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#subject",
		Short:   "subject",
		Comment: "The subject of the subject RDF statement.",
	}

	// The subject is an instance of a class.
	Prop_Type = &meta.Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#type",
		Short:   "type",
		Comment: "The subject is an instance of a class.",
	}

	// Idiomatic property used for structured values.
	Prop_Value = &meta.Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#value",
		Short:   "value",
		Comment: "Idiomatic property used for structured values.",
	}

	// A description of the subject resource.
	Prop_Comment = &meta.Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#comment",
		Short:   "comment",
		Comment: "A description of the subject resource.",
	}

	// A domain of the subject property.
	Prop_Domain = &meta.Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#domain",
		Short:   "domain",
		Comment: "A domain of the subject property.",
	}

	// The defininition of the subject resource.
	Prop_IsDefinedBy = &meta.Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#isDefinedBy",
		Short:   "isDefinedBy",
		Comment: "The defininition of the subject resource.",
	}

	// A human-readable name for the subject.
	Prop_Label = &meta.Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#label",
		Short:   "label",
		Comment: "A human-readable name for the subject.",
	}

	// A member of the subject resource.
	Prop_Member = &meta.Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#member",
		Short:   "member",
		Comment: "A member of the subject resource.",
	}

	// A range of the subject property.
	Prop_Range = &meta.Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#range",
		Short:   "range",
		Comment: "A range of the subject property.",
	}

	// Further information about the subject resource.
	Prop_SeeAlso = &meta.Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#seeAlso",
		Short:   "seeAlso",
		Comment: "Further information about the subject resource.",
	}

	// The subject is a subclass of a class.
	Prop_SubClassOf = &meta.Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#subClassOf",
		Short:   "subClassOf",
		Comment: "The subject is a subclass of a class.",
	}

	// The subject is a subproperty of a property.
	Prop_SubPropertyOf = &meta.Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#subPropertyOf",
		Short:   "subPropertyOf",
		Comment: "The subject is a subproperty of a property.",
	}

	// The class of containers of alternatives.
	Class_Alt = &meta.Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#Alt",
		Short: "Alt",
		Cast:  func(e ld.Entity) ld.Entity { return AsAlt(e) },
		SubClassOf: []*meta.Type{
			Class_Container,
		},
		Props: []*meta.Prop{},
	}

	// The class of unordered containers.
	Class_Bag = &meta.Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag",
		Short: "Bag",
		Cast:  func(e ld.Entity) ld.Entity { return AsBag(e) },
		SubClassOf: []*meta.Type{
			Class_Container,
		},
		Props: []*meta.Prop{},
	}

	// The class of RDF Lists.
	Class_List = &meta.Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#List",
		Short: "List",
		Cast:  func(e ld.Entity) ld.Entity { return AsList(e) },
		SubClassOf: []*meta.Type{
			Class_Resource,
		},
		Props: []*meta.Prop{
			Prop_First,
			Prop_Rest,
		},
	}

	// The class of RDF properties.
	Class_Property = &meta.Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property",
		Short: "Property",
		Cast:  func(e ld.Entity) ld.Entity { return AsProperty(e) },
		SubClassOf: []*meta.Type{
			Class_Resource,
		},
		Props: []*meta.Prop{
			Prop_Domain,
			Prop_Range,
			Prop_SubPropertyOf,
		},
	}

	// The class of ordered containers.
	Class_Seq = &meta.Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq",
		Short: "Seq",
		Cast:  func(e ld.Entity) ld.Entity { return AsSeq(e) },
		SubClassOf: []*meta.Type{
			Class_Container,
		},
		Props: []*meta.Prop{},
	}

	// The class of RDF statements.
	Class_Statement = &meta.Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement",
		Short: "Statement",
		Cast:  func(e ld.Entity) ld.Entity { return AsStatement(e) },
		SubClassOf: []*meta.Type{
			Class_Resource,
		},
		Props: []*meta.Prop{
			Prop_Object,
			Prop_Predicate,
			Prop_Subject,
		},
	}

	// The class of classes.
	Class_Class = &meta.Type{
		ID:    "http://www.w3.org/2000/01/rdf-schema#Class",
		Short: "Class",
		Cast:  func(e ld.Entity) ld.Entity { return AsClass(e) },
		SubClassOf: []*meta.Type{
			Class_Resource,
		},
		Props: []*meta.Prop{
			Prop_SubClassOf,
		},
	}

	// The class of RDF containers.
	Class_Container = &meta.Type{
		ID:    "http://www.w3.org/2000/01/rdf-schema#Container",
		Short: "Container",
		Cast:  func(e ld.Entity) ld.Entity { return AsContainer(e) },
		SubClassOf: []*meta.Type{
			Class_Resource,
		},
		Props: []*meta.Prop{},
	}

	// The class of container membership properties, rdf:_1, rdf:_2, ...,
	// all of which are sub-properties of 'member'.
	Class_ContainerMembershipProperty = &meta.Type{
		ID:    "http://www.w3.org/2000/01/rdf-schema#ContainerMembershipProperty",
		Short: "ContainerMembershipProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsContainerMembershipProperty(e) },
		SubClassOf: []*meta.Type{
			Class_Property,
		},
		Props: []*meta.Prop{},
	}

	// The class of RDF datatypes.
	Class_Datatype = &meta.Type{
		ID:    "http://www.w3.org/2000/01/rdf-schema#Datatype",
		Short: "Datatype",
		Cast:  func(e ld.Entity) ld.Entity { return AsDatatype(e) },
		SubClassOf: []*meta.Type{
			Class_Class,
		},
		Props: []*meta.Prop{},
	}

	// The class of literal values, eg. textual strings and integers.
	Class_Literal = &meta.Type{
		ID:    "http://www.w3.org/2000/01/rdf-schema#Literal",
		Short: "Literal",
		Cast:  func(e ld.Entity) ld.Entity { return AsLiteral(e) },
		SubClassOf: []*meta.Type{
			Class_Resource,
		},
		Props: []*meta.Prop{},
	}

	// The class resource, everything.
	Class_Resource = &meta.Type{
		ID:         "http://www.w3.org/2000/01/rdf-schema#Resource",
		Short:      "Resource",
		Cast:       func(e ld.Entity) ld.Entity { return AsResource(e) },
		SubClassOf: []*meta.Type{},
		Props: []*meta.Prop{
			Prop_Type,
			Prop_Value,
			Prop_Comment,
			Prop_IsDefinedBy,
			Prop_Label,
			Prop_Member,
			Prop_SeeAlso,
		},
	}
)

// The empty list, with no items in it. If the rest of a list is nil then the list has no more items in it.
// http://www.w3.org/1999/02/22-rdf-syntax-ns#nil - [http://www.w3.org/1999/02/22-rdf-syntax-ns#List http://www.w3.org/ns/activitystreams#OrderedItems]
