// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package owl

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/rdf"
)

// The class of collections of pairwise different individuals.
type AllDifferent struct{ rdf.Resource }

// The property that determines the collection of pairwise different individuals in a owl:AllDifferent axiom.
func (obj AllDifferent) DistinctMembers() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#distinctMembers"]
}

// The class of collections of pairwise disjoint classes.
type AllDisjointClasses struct{ rdf.Resource }

// The class of collections of pairwise disjoint properties.
type AllDisjointProperties struct{ rdf.Resource }

// The class of annotated annotations for which the RDF serialization consists of an annotated subject, predicate and object.
type Annotation struct{ rdf.Resource }

// The class of annotation properties.
type AnnotationProperty struct{ rdf.Property }

// The class of asymmetric properties.
type AsymmetricProperty struct{ ObjectProperty }

// The class of annotated axioms for which the RDF serialization consists of an annotated subject, predicate and object.
type Axiom struct{ rdf.Resource }

// The class of OWL classes.
type Class struct{ rdf.Class }

// The property that determines that a given class is the complement of another class.
func (obj Class) ComplementOf() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#complementOf"]
}

// The property that determines that a given class is equivalent to the disjoint union of a collection of other classes.
func (obj Class) DisjointUnionOf() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#disjointUnionOf"]
}

// The property that determines that two given classes are disjoint.
func (obj Class) DisjointWith() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#disjointWith"]
}

// The property that determines the collection of properties that jointly build a key.
func (obj Class) HasKey() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#hasKey"]
}

// The class of OWL data ranges, which are special kinds of datatypes. Note: The use of the IRI owl:DataRange has been deprecated as of OWL 2. The IRI rdfs:Datatype SHOULD be used instead.
type DataRange struct{ rdf.Datatype }

// The class of data properties.
type DatatypeProperty struct{ rdf.Property }

// The class of deprecated classes.
type DeprecatedClass struct{ rdf.Class }

// The class of deprecated properties.
type DeprecatedProperty struct{ rdf.Property }

// The class of functional properties.
type FunctionalProperty struct{ rdf.Property }

// The class of inverse-functional properties.
type InverseFunctionalProperty struct{ ObjectProperty }

// The class of irreflexive properties.
type IrreflexiveProperty struct{ ObjectProperty }

// The class of named individuals.
type NamedIndividual struct{ Thing }

// The class of negative property assertions.
type NegativePropertyAssertion struct{ rdf.Resource }

// The property that determines the predicate of a negative property assertion.
func (obj NegativePropertyAssertion) AssertionProperty() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#assertionProperty"]
}

// The property that determines the subject of a negative property assertion.
func (obj NegativePropertyAssertion) SourceIndividual() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#sourceIndividual"]
}

// The property that determines the object of a negative object property assertion.
func (obj NegativePropertyAssertion) TargetIndividual() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#targetIndividual"]
}

// The property that determines the value of a negative data property assertion.
func (obj NegativePropertyAssertion) TargetValue() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#targetValue"]
}

// The class of object properties.
type ObjectProperty struct{ rdf.Property }

// The property that determines that two given properties are inverse.
func (obj ObjectProperty) InverseOf() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#inverseOf"]
}

// The property that determines the n-tuple of properties that build a sub property chain of a given property.
func (obj ObjectProperty) PropertyChainAxiom() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#propertyChainAxiom"]
}

// The class of ontologies.
type Ontology struct{ rdf.Resource }

// The class of ontology properties.
type OntologyProperty struct{ rdf.Property }

// The class of reflexive properties.
type ReflexiveProperty struct{ ObjectProperty }

// The class of property restrictions.
type Restriction struct{ Class }

// The property that determines the class that a universal property restriction refers to.
func (obj Restriction) AllValuesFrom() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#allValuesFrom"]
}

// The property that determines the cardinality of an exact cardinality restriction.
func (obj Restriction) Cardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#cardinality"]
}

// The property that determines the property that a self restriction refers to.
func (obj Restriction) HasSelf() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#hasSelf"]
}

// The property that determines the individual that a has-value restriction refers to.
func (obj Restriction) HasValue() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#hasValue"]
}

// The property that determines the cardinality of a maximum cardinality restriction.
func (obj Restriction) MaxCardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#maxCardinality"]
}

// The property that determines the cardinality of a maximum qualified cardinality restriction.
func (obj Restriction) MaxQualifiedCardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#maxQualifiedCardinality"]
}

// The property that determines the cardinality of a minimum cardinality restriction.
func (obj Restriction) MinCardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#minCardinality"]
}

// The property that determines the cardinality of a minimum qualified cardinality restriction.
func (obj Restriction) MinQualifiedCardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#minQualifiedCardinality"]
}

// The property that determines the class that a qualified object cardinality restriction refers to.
func (obj Restriction) OnClass() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#onClass"]
}

// The property that determines the data range that a qualified data cardinality restriction refers to.
func (obj Restriction) OnDataRange() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#onDataRange"]
}

// The property that determines the n-tuple of properties that a property restriction on an n-ary data range refers to.
func (obj Restriction) OnProperties() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#onProperties"]
}

// The property that determines the property that a property restriction refers to.
func (obj Restriction) OnProperty() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#onProperty"]
}

// The property that determines the cardinality of an exact qualified cardinality restriction.
func (obj Restriction) QualifiedCardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#qualifiedCardinality"]
}

// The property that determines the class that an existential property restriction refers to.
func (obj Restriction) SomeValuesFrom() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#someValuesFrom"]
}

// The class of symmetric properties.
type SymmetricProperty struct{ ObjectProperty }

// The class of transitive properties.
type TransitiveProperty struct{ ObjectProperty }

var (
	_ ld.Entity = AllDifferent{}
	_ ld.Entity = AllDisjointClasses{}
	_ ld.Entity = AllDisjointProperties{}
	_ ld.Entity = Annotation{}
	_ ld.Entity = AnnotationProperty{}
	_ ld.Entity = AsymmetricProperty{}
	_ ld.Entity = Axiom{}
	_ ld.Entity = Class{}
	_ ld.Entity = DataRange{}
	_ ld.Entity = DatatypeProperty{}
	_ ld.Entity = DeprecatedClass{}
	_ ld.Entity = DeprecatedProperty{}
	_ ld.Entity = FunctionalProperty{}
	_ ld.Entity = InverseFunctionalProperty{}
	_ ld.Entity = IrreflexiveProperty{}
	_ ld.Entity = NamedIndividual{}
	_ ld.Entity = NegativePropertyAssertion{}
	_ ld.Entity = ObjectProperty{}
	_ ld.Entity = Ontology{}
	_ ld.Entity = OntologyProperty{}
	_ ld.Entity = ReflexiveProperty{}
	_ ld.Entity = Restriction{}
	_ ld.Entity = SymmetricProperty{}
	_ ld.Entity = TransitiveProperty{}
)
