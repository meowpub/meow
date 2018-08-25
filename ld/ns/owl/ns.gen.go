// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package owl

import (
	"github.com/meowpub/meow/ld"
)

const Namespace = "http://www.w3.org/2002/07/owl#"

const (
	// The property that determines the class that a universal property restriction refers to.
	PropAllValuesFrom = "http://www.w3.org/2002/07/owl#allValuesFrom"

	// The property that determines the predicate of an annotated axiom or annotated annotation.
	PropAnnotatedProperty = "http://www.w3.org/2002/07/owl#annotatedProperty"

	// The property that determines the subject of an annotated axiom or annotated annotation.
	PropAnnotatedSource = "http://www.w3.org/2002/07/owl#annotatedSource"

	// The property that determines the object of an annotated axiom or annotated annotation.
	PropAnnotatedTarget = "http://www.w3.org/2002/07/owl#annotatedTarget"

	// The property that determines the predicate of a negative property assertion.
	PropAssertionProperty = "http://www.w3.org/2002/07/owl#assertionProperty"

	// The property that determines the cardinality of an exact cardinality restriction.
	PropCardinality = "http://www.w3.org/2002/07/owl#cardinality"

	// The property that determines that a given class is the complement of another class.
	PropComplementOf = "http://www.w3.org/2002/07/owl#complementOf"

	// The property that determines that a given data range is the complement of another data range with respect to the data domain.
	PropDatatypeComplementOf = "http://www.w3.org/2002/07/owl#datatypeComplementOf"

	// The property that determines that two given individuals are different.
	PropDifferentFrom = "http://www.w3.org/2002/07/owl#differentFrom"

	// The property that determines that a given class is equivalent to the disjoint union of a collection of other classes.
	PropDisjointUnionOf = "http://www.w3.org/2002/07/owl#disjointUnionOf"

	// The property that determines that two given classes are disjoint.
	PropDisjointWith = "http://www.w3.org/2002/07/owl#disjointWith"

	// The property that determines the collection of pairwise different individuals in a owl:AllDifferent axiom.
	PropDistinctMembers = "http://www.w3.org/2002/07/owl#distinctMembers"

	// The property that determines that two given classes are equivalent, and that is used to specify datatype definitions.
	PropEquivalentClass = "http://www.w3.org/2002/07/owl#equivalentClass"

	// The property that determines that two given properties are equivalent.
	PropEquivalentProperty = "http://www.w3.org/2002/07/owl#equivalentProperty"

	// The property that determines the collection of properties that jointly build a key.
	PropHasKey = "http://www.w3.org/2002/07/owl#hasKey"

	// The property that determines the property that a self restriction refers to.
	PropHasSelf = "http://www.w3.org/2002/07/owl#hasSelf"

	// The property that determines the individual that a has-value restriction refers to.
	PropHasValue = "http://www.w3.org/2002/07/owl#hasValue"

	// The property that determines the collection of classes or data ranges that build an intersection.
	PropIntersectionOf = "http://www.w3.org/2002/07/owl#intersectionOf"

	// The property that determines that two given properties are inverse.
	PropInverseOf = "http://www.w3.org/2002/07/owl#inverseOf"

	// The property that determines the cardinality of a maximum cardinality restriction.
	PropMaxCardinality = "http://www.w3.org/2002/07/owl#maxCardinality"

	// The property that determines the cardinality of a maximum qualified cardinality restriction.
	PropMaxQualifiedCardinality = "http://www.w3.org/2002/07/owl#maxQualifiedCardinality"

	// The property that determines the collection of members in either a owl:AllDifferent, owl:AllDisjointClasses or owl:AllDisjointProperties axiom.
	PropMembers = "http://www.w3.org/2002/07/owl#members"

	// The property that determines the cardinality of a minimum cardinality restriction.
	PropMinCardinality = "http://www.w3.org/2002/07/owl#minCardinality"

	// The property that determines the cardinality of a minimum qualified cardinality restriction.
	PropMinQualifiedCardinality = "http://www.w3.org/2002/07/owl#minQualifiedCardinality"

	// The property that determines the class that a qualified object cardinality restriction refers to.
	PropOnClass = "http://www.w3.org/2002/07/owl#onClass"

	// The property that determines the data range that a qualified data cardinality restriction refers to.
	PropOnDataRange = "http://www.w3.org/2002/07/owl#onDataRange"

	// The property that determines the datatype that a datatype restriction refers to.
	PropOnDatatype = "http://www.w3.org/2002/07/owl#onDatatype"

	// The property that determines the n-tuple of properties that a property restriction on an n-ary data range refers to.
	PropOnProperties = "http://www.w3.org/2002/07/owl#onProperties"

	// The property that determines the property that a property restriction refers to.
	PropOnProperty = "http://www.w3.org/2002/07/owl#onProperty"

	// The property that determines the collection of individuals or data values that build an enumeration.
	PropOneOf = "http://www.w3.org/2002/07/owl#oneOf"

	// The property that determines the n-tuple of properties that build a sub property chain of a given property.
	PropPropertyChainAxiom = "http://www.w3.org/2002/07/owl#propertyChainAxiom"

	// The property that determines that two given properties are disjoint.
	PropPropertyDisjointWith = "http://www.w3.org/2002/07/owl#propertyDisjointWith"

	// The property that determines the cardinality of an exact qualified cardinality restriction.
	PropQualifiedCardinality = "http://www.w3.org/2002/07/owl#qualifiedCardinality"

	// The property that determines that two given individuals are equal.
	PropSameAs = "http://www.w3.org/2002/07/owl#sameAs"

	// The property that determines the class that an existential property restriction refers to.
	PropSomeValuesFrom = "http://www.w3.org/2002/07/owl#someValuesFrom"

	// The property that determines the subject of a negative property assertion.
	PropSourceIndividual = "http://www.w3.org/2002/07/owl#sourceIndividual"

	// The property that determines the object of a negative object property assertion.
	PropTargetIndividual = "http://www.w3.org/2002/07/owl#targetIndividual"

	// The property that determines the value of a negative data property assertion.
	PropTargetValue = "http://www.w3.org/2002/07/owl#targetValue"

	// The property that determines the collection of classes or data ranges that build a union.
	PropUnionOf = "http://www.w3.org/2002/07/owl#unionOf"

	// The property that determines the collection of facet-value pairs that define a datatype restriction.
	PropWithRestrictions = "http://www.w3.org/2002/07/owl#withRestrictions"
)

// Namespace.
var NS = &ld.Namespace{
	ID:    "http://www.w3.org/2002/07/owl#",
	Short: "owl",
	Props: map[string]string{
		"allValuesFrom":           "http://www.w3.org/2002/07/owl#allValuesFrom",
		"annotatedProperty":       "http://www.w3.org/2002/07/owl#annotatedProperty",
		"annotatedSource":         "http://www.w3.org/2002/07/owl#annotatedSource",
		"annotatedTarget":         "http://www.w3.org/2002/07/owl#annotatedTarget",
		"assertionProperty":       "http://www.w3.org/2002/07/owl#assertionProperty",
		"cardinality":             "http://www.w3.org/2002/07/owl#cardinality",
		"complementOf":            "http://www.w3.org/2002/07/owl#complementOf",
		"datatypeComplementOf":    "http://www.w3.org/2002/07/owl#datatypeComplementOf",
		"differentFrom":           "http://www.w3.org/2002/07/owl#differentFrom",
		"disjointUnionOf":         "http://www.w3.org/2002/07/owl#disjointUnionOf",
		"disjointWith":            "http://www.w3.org/2002/07/owl#disjointWith",
		"distinctMembers":         "http://www.w3.org/2002/07/owl#distinctMembers",
		"equivalentClass":         "http://www.w3.org/2002/07/owl#equivalentClass",
		"equivalentProperty":      "http://www.w3.org/2002/07/owl#equivalentProperty",
		"hasKey":                  "http://www.w3.org/2002/07/owl#hasKey",
		"hasSelf":                 "http://www.w3.org/2002/07/owl#hasSelf",
		"hasValue":                "http://www.w3.org/2002/07/owl#hasValue",
		"intersectionOf":          "http://www.w3.org/2002/07/owl#intersectionOf",
		"inverseOf":               "http://www.w3.org/2002/07/owl#inverseOf",
		"maxCardinality":          "http://www.w3.org/2002/07/owl#maxCardinality",
		"maxQualifiedCardinality": "http://www.w3.org/2002/07/owl#maxQualifiedCardinality",
		"members":                 "http://www.w3.org/2002/07/owl#members",
		"minCardinality":          "http://www.w3.org/2002/07/owl#minCardinality",
		"minQualifiedCardinality": "http://www.w3.org/2002/07/owl#minQualifiedCardinality",
		"onClass":                 "http://www.w3.org/2002/07/owl#onClass",
		"onDataRange":             "http://www.w3.org/2002/07/owl#onDataRange",
		"onDatatype":              "http://www.w3.org/2002/07/owl#onDatatype",
		"onProperties":            "http://www.w3.org/2002/07/owl#onProperties",
		"onProperty":              "http://www.w3.org/2002/07/owl#onProperty",
		"oneOf":                   "http://www.w3.org/2002/07/owl#oneOf",
		"propertyChainAxiom":      "http://www.w3.org/2002/07/owl#propertyChainAxiom",
		"propertyDisjointWith":    "http://www.w3.org/2002/07/owl#propertyDisjointWith",
		"qualifiedCardinality":    "http://www.w3.org/2002/07/owl#qualifiedCardinality",
		"sameAs":                  "http://www.w3.org/2002/07/owl#sameAs",
		"someValuesFrom":          "http://www.w3.org/2002/07/owl#someValuesFrom",
		"sourceIndividual":        "http://www.w3.org/2002/07/owl#sourceIndividual",
		"targetIndividual":        "http://www.w3.org/2002/07/owl#targetIndividual",
		"targetValue":             "http://www.w3.org/2002/07/owl#targetValue",
		"unionOf":                 "http://www.w3.org/2002/07/owl#unionOf",
		"withRestrictions":        "http://www.w3.org/2002/07/owl#withRestrictions",
	},
	Classes: map[string]func(ld.Entity) ld.Entity{
		"AllDifferent":              func(e ld.Entity) ld.Entity { return &AllDifferent{O: e.Obj()} },
		"AllDisjointClasses":        func(e ld.Entity) ld.Entity { return &AllDisjointClasses{O: e.Obj()} },
		"AllDisjointProperties":     func(e ld.Entity) ld.Entity { return &AllDisjointProperties{O: e.Obj()} },
		"Annotation":                func(e ld.Entity) ld.Entity { return &Annotation{O: e.Obj()} },
		"AnnotationProperty":        func(e ld.Entity) ld.Entity { return &AnnotationProperty{O: e.Obj()} },
		"AsymmetricProperty":        func(e ld.Entity) ld.Entity { return &AsymmetricProperty{O: e.Obj()} },
		"Axiom":                     func(e ld.Entity) ld.Entity { return &Axiom{O: e.Obj()} },
		"Class":                     func(e ld.Entity) ld.Entity { return &Class{O: e.Obj()} },
		"DataRange":                 func(e ld.Entity) ld.Entity { return &DataRange{O: e.Obj()} },
		"DatatypeProperty":          func(e ld.Entity) ld.Entity { return &DatatypeProperty{O: e.Obj()} },
		"DeprecatedClass":           func(e ld.Entity) ld.Entity { return &DeprecatedClass{O: e.Obj()} },
		"DeprecatedProperty":        func(e ld.Entity) ld.Entity { return &DeprecatedProperty{O: e.Obj()} },
		"FunctionalProperty":        func(e ld.Entity) ld.Entity { return &FunctionalProperty{O: e.Obj()} },
		"InverseFunctionalProperty": func(e ld.Entity) ld.Entity { return &InverseFunctionalProperty{O: e.Obj()} },
		"IrreflexiveProperty":       func(e ld.Entity) ld.Entity { return &IrreflexiveProperty{O: e.Obj()} },
		"NamedIndividual":           func(e ld.Entity) ld.Entity { return &NamedIndividual{O: e.Obj()} },
		"NegativePropertyAssertion": func(e ld.Entity) ld.Entity { return &NegativePropertyAssertion{O: e.Obj()} },
		"ObjectProperty":            func(e ld.Entity) ld.Entity { return &ObjectProperty{O: e.Obj()} },
		"Ontology":                  func(e ld.Entity) ld.Entity { return &Ontology{O: e.Obj()} },
		"OntologyProperty":          func(e ld.Entity) ld.Entity { return &OntologyProperty{O: e.Obj()} },
		"ReflexiveProperty":         func(e ld.Entity) ld.Entity { return &ReflexiveProperty{O: e.Obj()} },
		"Restriction":               func(e ld.Entity) ld.Entity { return &Restriction{O: e.Obj()} },
		"SymmetricProperty":         func(e ld.Entity) ld.Entity { return &SymmetricProperty{O: e.Obj()} },
		"TransitiveProperty":        func(e ld.Entity) ld.Entity { return &TransitiveProperty{O: e.Obj()} },
	},
}

// This is the empty class.
// http://www.w3.org/2002/07/owl#Nothing - http://www.w3.org/2002/07/owl#Class

// The class of OWL individuals.
// http://www.w3.org/2002/07/owl#Thing - http://www.w3.org/2002/07/owl#Class

// The annotation property that indicates that a given ontology is backward compatible with another ontology.
// http://www.w3.org/2002/07/owl#backwardCompatibleWith - http://www.w3.org/2002/07/owl#AnnotationProperty

// The data property that does not relate any individual to any data value.
// http://www.w3.org/2002/07/owl#bottomDataProperty - http://www.w3.org/2002/07/owl#DatatypeProperty

// The object property that does not relate any two individuals.
// http://www.w3.org/2002/07/owl#bottomObjectProperty - http://www.w3.org/2002/07/owl#ObjectProperty

// The annotation property that indicates that a given entity has been deprecated.
// http://www.w3.org/2002/07/owl#deprecated - http://www.w3.org/2002/07/owl#AnnotationProperty

// The property that is used for importing other ontologies into a given ontology.
// http://www.w3.org/2002/07/owl#imports - http://www.w3.org/2002/07/owl#OntologyProperty

// The annotation property that indicates that a given ontology is incompatible with another ontology.
// http://www.w3.org/2002/07/owl#incompatibleWith - http://www.w3.org/2002/07/owl#AnnotationProperty

// The annotation property that indicates the predecessor ontology of a given ontology.
// http://www.w3.org/2002/07/owl#priorVersion - http://www.w3.org/2002/07/owl#AnnotationProperty

// The data property that relates every individual to every data value.
// http://www.w3.org/2002/07/owl#topDataProperty - http://www.w3.org/2002/07/owl#DatatypeProperty

// The object property that relates every two individuals.
// http://www.w3.org/2002/07/owl#topObjectProperty - http://www.w3.org/2002/07/owl#ObjectProperty

// The property that identifies the version IRI of an ontology.
// http://www.w3.org/2002/07/owl#versionIRI - http://www.w3.org/2002/07/owl#OntologyProperty

// The annotation property that provides version information for an ontology or another OWL construct.
// http://www.w3.org/2002/07/owl#versionInfo - http://www.w3.org/2002/07/owl#AnnotationProperty
