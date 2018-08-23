// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package owl

import (
	"github.com/meowpub/meow/ld"
)

const Namespace = "http://www.w3.org/2002/07/owl#"

const (
	PropAllValuesFrom = "http://www.w3.org/2002/07/owl#allValuesFrom"

	PropAnnotatedProperty = "http://www.w3.org/2002/07/owl#annotatedProperty"

	PropAnnotatedSource = "http://www.w3.org/2002/07/owl#annotatedSource"

	PropAnnotatedTarget = "http://www.w3.org/2002/07/owl#annotatedTarget"

	PropAssertionProperty = "http://www.w3.org/2002/07/owl#assertionProperty"

	PropCardinality = "http://www.w3.org/2002/07/owl#cardinality"

	PropComplementOf = "http://www.w3.org/2002/07/owl#complementOf"

	PropDatatypeComplementOf = "http://www.w3.org/2002/07/owl#datatypeComplementOf"

	PropDifferentFrom = "http://www.w3.org/2002/07/owl#differentFrom"

	PropDisjointUnionOf = "http://www.w3.org/2002/07/owl#disjointUnionOf"

	PropDisjointWith = "http://www.w3.org/2002/07/owl#disjointWith"

	PropDistinctMembers = "http://www.w3.org/2002/07/owl#distinctMembers"

	PropEquivalentClass = "http://www.w3.org/2002/07/owl#equivalentClass"

	PropEquivalentProperty = "http://www.w3.org/2002/07/owl#equivalentProperty"

	PropHasKey = "http://www.w3.org/2002/07/owl#hasKey"

	PropHasSelf = "http://www.w3.org/2002/07/owl#hasSelf"

	PropHasValue = "http://www.w3.org/2002/07/owl#hasValue"

	PropIntersectionOf = "http://www.w3.org/2002/07/owl#intersectionOf"

	PropInverseOf = "http://www.w3.org/2002/07/owl#inverseOf"

	PropMaxCardinality = "http://www.w3.org/2002/07/owl#maxCardinality"

	PropMaxQualifiedCardinality = "http://www.w3.org/2002/07/owl#maxQualifiedCardinality"

	PropMembers = "http://www.w3.org/2002/07/owl#members"

	PropMinCardinality = "http://www.w3.org/2002/07/owl#minCardinality"

	PropMinQualifiedCardinality = "http://www.w3.org/2002/07/owl#minQualifiedCardinality"

	PropOnClass = "http://www.w3.org/2002/07/owl#onClass"

	PropOnDataRange = "http://www.w3.org/2002/07/owl#onDataRange"

	PropOnDatatype = "http://www.w3.org/2002/07/owl#onDatatype"

	PropOnProperties = "http://www.w3.org/2002/07/owl#onProperties"

	PropOnProperty = "http://www.w3.org/2002/07/owl#onProperty"

	PropOneOf = "http://www.w3.org/2002/07/owl#oneOf"

	PropPropertyChainAxiom = "http://www.w3.org/2002/07/owl#propertyChainAxiom"

	PropPropertyDisjointWith = "http://www.w3.org/2002/07/owl#propertyDisjointWith"

	PropQualifiedCardinality = "http://www.w3.org/2002/07/owl#qualifiedCardinality"

	PropSameAs = "http://www.w3.org/2002/07/owl#sameAs"

	PropSomeValuesFrom = "http://www.w3.org/2002/07/owl#someValuesFrom"

	PropSourceIndividual = "http://www.w3.org/2002/07/owl#sourceIndividual"

	PropTargetIndividual = "http://www.w3.org/2002/07/owl#targetIndividual"

	PropTargetValue = "http://www.w3.org/2002/07/owl#targetValue"

	PropUnionOf = "http://www.w3.org/2002/07/owl#unionOf"

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

// http://www.w3.org/2002/07/owl#Nothing - http://www.w3.org/2002/07/owl#Class

// http://www.w3.org/2002/07/owl#Thing - http://www.w3.org/2002/07/owl#Class

// http://www.w3.org/2002/07/owl#backwardCompatibleWith - http://www.w3.org/2002/07/owl#AnnotationProperty

// http://www.w3.org/2002/07/owl#bottomDataProperty - http://www.w3.org/2002/07/owl#DatatypeProperty

// http://www.w3.org/2002/07/owl#bottomObjectProperty - http://www.w3.org/2002/07/owl#ObjectProperty

// http://www.w3.org/2002/07/owl#deprecated - http://www.w3.org/2002/07/owl#AnnotationProperty

// http://www.w3.org/2002/07/owl#imports - http://www.w3.org/2002/07/owl#OntologyProperty

// http://www.w3.org/2002/07/owl#incompatibleWith - http://www.w3.org/2002/07/owl#AnnotationProperty

// http://www.w3.org/2002/07/owl#priorVersion - http://www.w3.org/2002/07/owl#AnnotationProperty

// http://www.w3.org/2002/07/owl#topDataProperty - http://www.w3.org/2002/07/owl#DatatypeProperty

// http://www.w3.org/2002/07/owl#topObjectProperty - http://www.w3.org/2002/07/owl#ObjectProperty

// http://www.w3.org/2002/07/owl#versionIRI - http://www.w3.org/2002/07/owl#OntologyProperty

// http://www.w3.org/2002/07/owl#versionInfo - http://www.w3.org/2002/07/owl#AnnotationProperty
