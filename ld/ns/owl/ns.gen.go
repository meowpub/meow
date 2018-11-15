// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package owl

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/meta"
	"github.com/meowpub/meow/ld/ns/rdf"
)

var (
	OWL = &meta.Namespace{
		ID:    "http://www.w3.org/2002/07/owl#",
		Short: "owl",
		Props: []*meta.Prop{
			Prop_AllValuesFrom,
			Prop_AnnotatedProperty,
			Prop_AnnotatedSource,
			Prop_AnnotatedTarget,
			Prop_AssertionProperty,
			Prop_BackwardCompatibleWith,
			Prop_BottomDataProperty,
			Prop_BottomObjectProperty,
			Prop_Cardinality,
			Prop_ComplementOf,
			Prop_DatatypeComplementOf,
			Prop_Deprecated,
			Prop_DifferentFrom,
			Prop_DisjointUnionOf,
			Prop_DisjointWith,
			Prop_DistinctMembers,
			Prop_EquivalentClass,
			Prop_EquivalentProperty,
			Prop_HasKey,
			Prop_HasSelf,
			Prop_HasValue,
			Prop_Imports,
			Prop_IncompatibleWith,
			Prop_IntersectionOf,
			Prop_InverseOf,
			Prop_MaxCardinality,
			Prop_MaxQualifiedCardinality,
			Prop_Members,
			Prop_MinCardinality,
			Prop_MinQualifiedCardinality,
			Prop_OnClass,
			Prop_OnDataRange,
			Prop_OnDatatype,
			Prop_OnProperties,
			Prop_OnProperty,
			Prop_OneOf,
			Prop_PriorVersion,
			Prop_PropertyChainAxiom,
			Prop_PropertyDisjointWith,
			Prop_QualifiedCardinality,
			Prop_SameAs,
			Prop_SomeValuesFrom,
			Prop_SourceIndividual,
			Prop_TargetIndividual,
			Prop_TargetValue,
			Prop_TopDataProperty,
			Prop_TopObjectProperty,
			Prop_UnionOf,
			Prop_VersionIRI,
			Prop_VersionInfo,
			Prop_WithRestrictions,
		},
		Types: map[string]*meta.Type{
			"http://www.w3.org/2002/07/owl#AllDifferent":              Class_AllDifferent,
			"http://www.w3.org/2002/07/owl#AllDisjointClasses":        Class_AllDisjointClasses,
			"http://www.w3.org/2002/07/owl#AllDisjointProperties":     Class_AllDisjointProperties,
			"http://www.w3.org/2002/07/owl#Annotation":                Class_Annotation,
			"http://www.w3.org/2002/07/owl#AnnotationProperty":        Class_AnnotationProperty,
			"http://www.w3.org/2002/07/owl#AsymmetricProperty":        Class_AsymmetricProperty,
			"http://www.w3.org/2002/07/owl#Axiom":                     Class_Axiom,
			"http://www.w3.org/2002/07/owl#Class":                     Class_Class,
			"http://www.w3.org/2002/07/owl#DataRange":                 Class_DataRange,
			"http://www.w3.org/2002/07/owl#DatatypeProperty":          Class_DatatypeProperty,
			"http://www.w3.org/2002/07/owl#DeprecatedClass":           Class_DeprecatedClass,
			"http://www.w3.org/2002/07/owl#DeprecatedProperty":        Class_DeprecatedProperty,
			"http://www.w3.org/2002/07/owl#FunctionalProperty":        Class_FunctionalProperty,
			"http://www.w3.org/2002/07/owl#InverseFunctionalProperty": Class_InverseFunctionalProperty,
			"http://www.w3.org/2002/07/owl#IrreflexiveProperty":       Class_IrreflexiveProperty,
			"http://www.w3.org/2002/07/owl#NamedIndividual":           Class_NamedIndividual,
			"http://www.w3.org/2002/07/owl#NegativePropertyAssertion": Class_NegativePropertyAssertion,
			"http://www.w3.org/2002/07/owl#Nothing":                   Class_Nothing,
			"http://www.w3.org/2002/07/owl#ObjectProperty":            Class_ObjectProperty,
			"http://www.w3.org/2002/07/owl#Ontology":                  Class_Ontology,
			"http://www.w3.org/2002/07/owl#OntologyProperty":          Class_OntologyProperty,
			"http://www.w3.org/2002/07/owl#ReflexiveProperty":         Class_ReflexiveProperty,
			"http://www.w3.org/2002/07/owl#Restriction":               Class_Restriction,
			"http://www.w3.org/2002/07/owl#SymmetricProperty":         Class_SymmetricProperty,
			"http://www.w3.org/2002/07/owl#Thing":                     Class_Thing,
			"http://www.w3.org/2002/07/owl#TransitiveProperty":        Class_TransitiveProperty,
		},
	}

	// The property that determines the class that a universal property restriction refers to.
	Prop_AllValuesFrom = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#allValuesFrom",
		Short:   "allValuesFrom",
		Comment: "The property that determines the class that a universal property restriction refers to.",
	}

	// The property that determines the predicate of an annotated axiom or annotated annotation.
	Prop_AnnotatedProperty = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#annotatedProperty",
		Short:   "annotatedProperty",
		Comment: "The property that determines the predicate of an annotated axiom or annotated annotation.",
	}

	// The property that determines the subject of an annotated axiom or annotated annotation.
	Prop_AnnotatedSource = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#annotatedSource",
		Short:   "annotatedSource",
		Comment: "The property that determines the subject of an annotated axiom or annotated annotation.",
	}

	// The property that determines the object of an annotated axiom or annotated annotation.
	Prop_AnnotatedTarget = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#annotatedTarget",
		Short:   "annotatedTarget",
		Comment: "The property that determines the object of an annotated axiom or annotated annotation.",
	}

	// The property that determines the predicate of a negative property assertion.
	Prop_AssertionProperty = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#assertionProperty",
		Short:   "assertionProperty",
		Comment: "The property that determines the predicate of a negative property assertion.",
	}

	// The annotation property that indicates that a given ontology is backward compatible with another ontology.
	Prop_BackwardCompatibleWith = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#backwardCompatibleWith",
		Short:   "backwardCompatibleWith",
		Comment: "The annotation property that indicates that a given ontology is backward compatible with another ontology.",
	}

	// The data property that does not relate any individual to any data value.
	Prop_BottomDataProperty = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#bottomDataProperty",
		Short:   "bottomDataProperty",
		Comment: "The data property that does not relate any individual to any data value.",
	}

	// The object property that does not relate any two individuals.
	Prop_BottomObjectProperty = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#bottomObjectProperty",
		Short:   "bottomObjectProperty",
		Comment: "The object property that does not relate any two individuals.",
	}

	// The property that determines the cardinality of an exact cardinality restriction.
	Prop_Cardinality = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#cardinality",
		Short:   "cardinality",
		Comment: "The property that determines the cardinality of an exact cardinality restriction.",
	}

	// The property that determines that a given class is the complement of another class.
	Prop_ComplementOf = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#complementOf",
		Short:   "complementOf",
		Comment: "The property that determines that a given class is the complement of another class.",
	}

	// The property that determines that a given data range is the complement of another data range with respect to the data domain.
	Prop_DatatypeComplementOf = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#datatypeComplementOf",
		Short:   "datatypeComplementOf",
		Comment: "The property that determines that a given data range is the complement of another data range with respect to the data domain.",
	}

	// The annotation property that indicates that a given entity has been deprecated.
	Prop_Deprecated = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#deprecated",
		Short:   "deprecated",
		Comment: "The annotation property that indicates that a given entity has been deprecated.",
	}

	// The property that determines that two given individuals are different.
	Prop_DifferentFrom = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#differentFrom",
		Short:   "differentFrom",
		Comment: "The property that determines that two given individuals are different.",
	}

	// The property that determines that a given class is equivalent to the disjoint union of a collection of other classes.
	Prop_DisjointUnionOf = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#disjointUnionOf",
		Short:   "disjointUnionOf",
		Comment: "The property that determines that a given class is equivalent to the disjoint union of a collection of other classes.",
	}

	// The property that determines that two given classes are disjoint.
	Prop_DisjointWith = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#disjointWith",
		Short:   "disjointWith",
		Comment: "The property that determines that two given classes are disjoint.",
	}

	// The property that determines the collection of pairwise different individuals in a owl:AllDifferent axiom.
	Prop_DistinctMembers = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#distinctMembers",
		Short:   "distinctMembers",
		Comment: "The property that determines the collection of pairwise different individuals in a owl:AllDifferent axiom.",
	}

	// The property that determines that two given classes are equivalent, and that is used to specify datatype definitions.
	Prop_EquivalentClass = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#equivalentClass",
		Short:   "equivalentClass",
		Comment: "The property that determines that two given classes are equivalent, and that is used to specify datatype definitions.",
	}

	// The property that determines that two given properties are equivalent.
	Prop_EquivalentProperty = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#equivalentProperty",
		Short:   "equivalentProperty",
		Comment: "The property that determines that two given properties are equivalent.",
	}

	// The property that determines the collection of properties that jointly build a key.
	Prop_HasKey = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#hasKey",
		Short:   "hasKey",
		Comment: "The property that determines the collection of properties that jointly build a key.",
	}

	// The property that determines the property that a self restriction refers to.
	Prop_HasSelf = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#hasSelf",
		Short:   "hasSelf",
		Comment: "The property that determines the property that a self restriction refers to.",
	}

	// The property that determines the individual that a has-value restriction refers to.
	Prop_HasValue = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#hasValue",
		Short:   "hasValue",
		Comment: "The property that determines the individual that a has-value restriction refers to.",
	}

	// The property that is used for importing other ontologies into a given ontology.
	Prop_Imports = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#imports",
		Short:   "imports",
		Comment: "The property that is used for importing other ontologies into a given ontology.",
	}

	// The annotation property that indicates that a given ontology is incompatible with another ontology.
	Prop_IncompatibleWith = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#incompatibleWith",
		Short:   "incompatibleWith",
		Comment: "The annotation property that indicates that a given ontology is incompatible with another ontology.",
	}

	// The property that determines the collection of classes or data ranges that build an intersection.
	Prop_IntersectionOf = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#intersectionOf",
		Short:   "intersectionOf",
		Comment: "The property that determines the collection of classes or data ranges that build an intersection.",
	}

	// The property that determines that two given properties are inverse.
	Prop_InverseOf = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#inverseOf",
		Short:   "inverseOf",
		Comment: "The property that determines that two given properties are inverse.",
	}

	// The property that determines the cardinality of a maximum cardinality restriction.
	Prop_MaxCardinality = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#maxCardinality",
		Short:   "maxCardinality",
		Comment: "The property that determines the cardinality of a maximum cardinality restriction.",
	}

	// The property that determines the cardinality of a maximum qualified cardinality restriction.
	Prop_MaxQualifiedCardinality = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#maxQualifiedCardinality",
		Short:   "maxQualifiedCardinality",
		Comment: "The property that determines the cardinality of a maximum qualified cardinality restriction.",
	}

	// The property that determines the collection of members in either a owl:AllDifferent, owl:AllDisjointClasses or owl:AllDisjointProperties axiom.
	Prop_Members = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#members",
		Short:   "members",
		Comment: "The property that determines the collection of members in either a owl:AllDifferent, owl:AllDisjointClasses or owl:AllDisjointProperties axiom.",
	}

	// The property that determines the cardinality of a minimum cardinality restriction.
	Prop_MinCardinality = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#minCardinality",
		Short:   "minCardinality",
		Comment: "The property that determines the cardinality of a minimum cardinality restriction.",
	}

	// The property that determines the cardinality of a minimum qualified cardinality restriction.
	Prop_MinQualifiedCardinality = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#minQualifiedCardinality",
		Short:   "minQualifiedCardinality",
		Comment: "The property that determines the cardinality of a minimum qualified cardinality restriction.",
	}

	// The property that determines the class that a qualified object cardinality restriction refers to.
	Prop_OnClass = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#onClass",
		Short:   "onClass",
		Comment: "The property that determines the class that a qualified object cardinality restriction refers to.",
	}

	// The property that determines the data range that a qualified data cardinality restriction refers to.
	Prop_OnDataRange = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#onDataRange",
		Short:   "onDataRange",
		Comment: "The property that determines the data range that a qualified data cardinality restriction refers to.",
	}

	// The property that determines the datatype that a datatype restriction refers to.
	Prop_OnDatatype = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#onDatatype",
		Short:   "onDatatype",
		Comment: "The property that determines the datatype that a datatype restriction refers to.",
	}

	// The property that determines the n-tuple of properties that a property restriction on an n-ary data range refers to.
	Prop_OnProperties = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#onProperties",
		Short:   "onProperties",
		Comment: "The property that determines the n-tuple of properties that a property restriction on an n-ary data range refers to.",
	}

	// The property that determines the property that a property restriction refers to.
	Prop_OnProperty = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#onProperty",
		Short:   "onProperty",
		Comment: "The property that determines the property that a property restriction refers to.",
	}

	// The property that determines the collection of individuals or data values that build an enumeration.
	Prop_OneOf = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#oneOf",
		Short:   "oneOf",
		Comment: "The property that determines the collection of individuals or data values that build an enumeration.",
	}

	// The annotation property that indicates the predecessor ontology of a given ontology.
	Prop_PriorVersion = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#priorVersion",
		Short:   "priorVersion",
		Comment: "The annotation property that indicates the predecessor ontology of a given ontology.",
	}

	// The property that determines the n-tuple of properties that build a sub property chain of a given property.
	Prop_PropertyChainAxiom = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#propertyChainAxiom",
		Short:   "propertyChainAxiom",
		Comment: "The property that determines the n-tuple of properties that build a sub property chain of a given property.",
	}

	// The property that determines that two given properties are disjoint.
	Prop_PropertyDisjointWith = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#propertyDisjointWith",
		Short:   "propertyDisjointWith",
		Comment: "The property that determines that two given properties are disjoint.",
	}

	// The property that determines the cardinality of an exact qualified cardinality restriction.
	Prop_QualifiedCardinality = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#qualifiedCardinality",
		Short:   "qualifiedCardinality",
		Comment: "The property that determines the cardinality of an exact qualified cardinality restriction.",
	}

	// The property that determines that two given individuals are equal.
	Prop_SameAs = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#sameAs",
		Short:   "sameAs",
		Comment: "The property that determines that two given individuals are equal.",
	}

	// The property that determines the class that an existential property restriction refers to.
	Prop_SomeValuesFrom = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#someValuesFrom",
		Short:   "someValuesFrom",
		Comment: "The property that determines the class that an existential property restriction refers to.",
	}

	// The property that determines the subject of a negative property assertion.
	Prop_SourceIndividual = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#sourceIndividual",
		Short:   "sourceIndividual",
		Comment: "The property that determines the subject of a negative property assertion.",
	}

	// The property that determines the object of a negative object property assertion.
	Prop_TargetIndividual = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#targetIndividual",
		Short:   "targetIndividual",
		Comment: "The property that determines the object of a negative object property assertion.",
	}

	// The property that determines the value of a negative data property assertion.
	Prop_TargetValue = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#targetValue",
		Short:   "targetValue",
		Comment: "The property that determines the value of a negative data property assertion.",
	}

	// The data property that relates every individual to every data value.
	Prop_TopDataProperty = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#topDataProperty",
		Short:   "topDataProperty",
		Comment: "The data property that relates every individual to every data value.",
	}

	// The object property that relates every two individuals.
	Prop_TopObjectProperty = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#topObjectProperty",
		Short:   "topObjectProperty",
		Comment: "The object property that relates every two individuals.",
	}

	// The property that determines the collection of classes or data ranges that build a union.
	Prop_UnionOf = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#unionOf",
		Short:   "unionOf",
		Comment: "The property that determines the collection of classes or data ranges that build a union.",
	}

	// The property that identifies the version IRI of an ontology.
	Prop_VersionIRI = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#versionIRI",
		Short:   "versionIRI",
		Comment: "The property that identifies the version IRI of an ontology.",
	}

	// The annotation property that provides version information for an ontology or another OWL construct.
	Prop_VersionInfo = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#versionInfo",
		Short:   "versionInfo",
		Comment: "The annotation property that provides version information for an ontology or another OWL construct.",
	}

	// The property that determines the collection of facet-value pairs that define a datatype restriction.
	Prop_WithRestrictions = &meta.Prop{
		ID:      "http://www.w3.org/2002/07/owl#withRestrictions",
		Short:   "withRestrictions",
		Comment: "The property that determines the collection of facet-value pairs that define a datatype restriction.",
	}

	// The class of collections of pairwise different individuals.
	Class_AllDifferent = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#AllDifferent",
		Short: "AllDifferent",
		Cast:  func(e ld.Entity) ld.Entity { return AsAllDifferent(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Resource,
		},
		Props: []*meta.Prop{
			Prop_DistinctMembers,
		},
	}

	// The class of collections of pairwise disjoint classes.
	Class_AllDisjointClasses = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#AllDisjointClasses",
		Short: "AllDisjointClasses",
		Cast:  func(e ld.Entity) ld.Entity { return AsAllDisjointClasses(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Resource,
		},
		Props: []*meta.Prop{},
	}

	// The class of collections of pairwise disjoint properties.
	Class_AllDisjointProperties = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#AllDisjointProperties",
		Short: "AllDisjointProperties",
		Cast:  func(e ld.Entity) ld.Entity { return AsAllDisjointProperties(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Resource,
		},
		Props: []*meta.Prop{},
	}

	// The class of annotated annotations for which the RDF serialization consists of an annotated subject, predicate and object.
	Class_Annotation = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#Annotation",
		Short: "Annotation",
		Cast:  func(e ld.Entity) ld.Entity { return AsAnnotation(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Resource,
		},
		Props: []*meta.Prop{},
	}

	// The class of annotation properties.
	Class_AnnotationProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#AnnotationProperty",
		Short: "AnnotationProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsAnnotationProperty(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Property,
		},
		Props: []*meta.Prop{},
	}

	// The class of asymmetric properties.
	Class_AsymmetricProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#AsymmetricProperty",
		Short: "AsymmetricProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsAsymmetricProperty(e) },
		SubClassOf: []*meta.Type{
			Class_ObjectProperty,
		},
		Props: []*meta.Prop{},
	}

	// The class of annotated axioms for which the RDF serialization consists of an annotated subject, predicate and object.
	Class_Axiom = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#Axiom",
		Short: "Axiom",
		Cast:  func(e ld.Entity) ld.Entity { return AsAxiom(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Resource,
		},
		Props: []*meta.Prop{},
	}

	// The class of OWL classes.
	Class_Class = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#Class",
		Short: "Class",
		Cast:  func(e ld.Entity) ld.Entity { return AsClass(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Class,
		},
		Props: []*meta.Prop{
			Prop_ComplementOf,
			Prop_DisjointUnionOf,
			Prop_DisjointWith,
			Prop_HasKey,
		},
	}

	// The class of OWL data ranges, which are special kinds of datatypes. Note: The use of the IRI owl:DataRange has been deprecated as of OWL 2. The IRI rdfs:Datatype SHOULD be used instead.
	Class_DataRange = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#DataRange",
		Short: "DataRange",
		Cast:  func(e ld.Entity) ld.Entity { return AsDataRange(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Datatype,
		},
		Props: []*meta.Prop{},
	}

	// The class of data properties.
	Class_DatatypeProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#DatatypeProperty",
		Short: "DatatypeProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsDatatypeProperty(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Property,
		},
		Props: []*meta.Prop{},
	}

	// The class of deprecated classes.
	Class_DeprecatedClass = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#DeprecatedClass",
		Short: "DeprecatedClass",
		Cast:  func(e ld.Entity) ld.Entity { return AsDeprecatedClass(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Class,
		},
		Props: []*meta.Prop{},
	}

	// The class of deprecated properties.
	Class_DeprecatedProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#DeprecatedProperty",
		Short: "DeprecatedProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsDeprecatedProperty(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Property,
		},
		Props: []*meta.Prop{},
	}

	// The class of functional properties.
	Class_FunctionalProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#FunctionalProperty",
		Short: "FunctionalProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsFunctionalProperty(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Property,
		},
		Props: []*meta.Prop{},
	}

	// The class of inverse-functional properties.
	Class_InverseFunctionalProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#InverseFunctionalProperty",
		Short: "InverseFunctionalProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsInverseFunctionalProperty(e) },
		SubClassOf: []*meta.Type{
			Class_ObjectProperty,
		},
		Props: []*meta.Prop{},
	}

	// The class of irreflexive properties.
	Class_IrreflexiveProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#IrreflexiveProperty",
		Short: "IrreflexiveProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsIrreflexiveProperty(e) },
		SubClassOf: []*meta.Type{
			Class_ObjectProperty,
		},
		Props: []*meta.Prop{},
	}

	// The class of named individuals.
	Class_NamedIndividual = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#NamedIndividual",
		Short: "NamedIndividual",
		Cast:  func(e ld.Entity) ld.Entity { return AsNamedIndividual(e) },
		SubClassOf: []*meta.Type{
			Class_Thing,
		},
		Props: []*meta.Prop{},
	}

	// The class of negative property assertions.
	Class_NegativePropertyAssertion = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#NegativePropertyAssertion",
		Short: "NegativePropertyAssertion",
		Cast:  func(e ld.Entity) ld.Entity { return AsNegativePropertyAssertion(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Resource,
		},
		Props: []*meta.Prop{
			Prop_AssertionProperty,
			Prop_SourceIndividual,
			Prop_TargetIndividual,
			Prop_TargetValue,
		},
	}

	// This is the empty class.
	Class_Nothing = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#Nothing",
		Short: "Nothing",
		Cast:  func(e ld.Entity) ld.Entity { return AsNothing(e) },
		SubClassOf: []*meta.Type{
			Class_Thing,
		},
		Props: []*meta.Prop{},
	}

	// The class of object properties.
	Class_ObjectProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#ObjectProperty",
		Short: "ObjectProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsObjectProperty(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Property,
		},
		Props: []*meta.Prop{
			Prop_InverseOf,
			Prop_PropertyChainAxiom,
		},
	}

	// The class of ontologies.
	Class_Ontology = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#Ontology",
		Short: "Ontology",
		Cast:  func(e ld.Entity) ld.Entity { return AsOntology(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Resource,
		},
		Props: []*meta.Prop{
			Prop_BackwardCompatibleWith,
			Prop_Imports,
			Prop_IncompatibleWith,
			Prop_PriorVersion,
			Prop_VersionIRI,
		},
	}

	// The class of ontology properties.
	Class_OntologyProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#OntologyProperty",
		Short: "OntologyProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsOntologyProperty(e) },
		SubClassOf: []*meta.Type{
			rdf.Class_Property,
		},
		Props: []*meta.Prop{},
	}

	// The class of reflexive properties.
	Class_ReflexiveProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#ReflexiveProperty",
		Short: "ReflexiveProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsReflexiveProperty(e) },
		SubClassOf: []*meta.Type{
			Class_ObjectProperty,
		},
		Props: []*meta.Prop{},
	}

	// The class of property restrictions.
	Class_Restriction = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#Restriction",
		Short: "Restriction",
		Cast:  func(e ld.Entity) ld.Entity { return AsRestriction(e) },
		SubClassOf: []*meta.Type{
			Class_Class,
		},
		Props: []*meta.Prop{
			Prop_AllValuesFrom,
			Prop_Cardinality,
			Prop_HasSelf,
			Prop_HasValue,
			Prop_MaxCardinality,
			Prop_MaxQualifiedCardinality,
			Prop_MinCardinality,
			Prop_MinQualifiedCardinality,
			Prop_OnClass,
			Prop_OnDataRange,
			Prop_OnProperties,
			Prop_OnProperty,
			Prop_QualifiedCardinality,
			Prop_SomeValuesFrom,
		},
	}

	// The class of symmetric properties.
	Class_SymmetricProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#SymmetricProperty",
		Short: "SymmetricProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsSymmetricProperty(e) },
		SubClassOf: []*meta.Type{
			Class_ObjectProperty,
		},
		Props: []*meta.Prop{},
	}

	// The class of OWL individuals.
	Class_Thing = &meta.Type{
		ID:         "http://www.w3.org/2002/07/owl#Thing",
		Short:      "Thing",
		Cast:       func(e ld.Entity) ld.Entity { return AsThing(e) },
		SubClassOf: []*meta.Type{},
		Props: []*meta.Prop{
			Prop_BottomDataProperty,
			Prop_BottomObjectProperty,
			Prop_DifferentFrom,
			Prop_SameAs,
			Prop_TopDataProperty,
			Prop_TopObjectProperty,
		},
	}

	// The class of transitive properties.
	Class_TransitiveProperty = &meta.Type{
		ID:    "http://www.w3.org/2002/07/owl#TransitiveProperty",
		Short: "TransitiveProperty",
		Cast:  func(e ld.Entity) ld.Entity { return AsTransitiveProperty(e) },
		SubClassOf: []*meta.Type{
			Class_ObjectProperty,
		},
		Props: []*meta.Prop{},
	}
)
