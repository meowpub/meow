// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package owl

import (
	"github.com/meowpub/meow/ld"
)

// The class of collections of pairwise different individuals.
type AllDifferent struct{ O *ld.Object }

func (obj AllDifferent) Obj() *ld.Object            { return obj.O }
func (obj AllDifferent) ID() string                 { return obj.O.ID() }
func (obj AllDifferent) Value() string              { return obj.O.Value() }
func (obj AllDifferent) Type() []string             { return obj.O.Type() }
func (obj AllDifferent) Get(key string) interface{} { return obj.O.Get(key) }

func (obj AllDifferent) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The property that determines the collection of pairwise different individuals in a owl:AllDifferent axiom.
func (obj AllDifferent) DistinctMembers() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#distinctMembers"]
}

// The class of collections of pairwise disjoint classes.
type AllDisjointClasses struct{ O *ld.Object }

func (obj AllDisjointClasses) Obj() *ld.Object            { return obj.O }
func (obj AllDisjointClasses) ID() string                 { return obj.O.ID() }
func (obj AllDisjointClasses) Value() string              { return obj.O.Value() }
func (obj AllDisjointClasses) Type() []string             { return obj.O.Type() }
func (obj AllDisjointClasses) Get(key string) interface{} { return obj.O.Get(key) }

func (obj AllDisjointClasses) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of collections of pairwise disjoint properties.
type AllDisjointProperties struct{ O *ld.Object }

func (obj AllDisjointProperties) Obj() *ld.Object            { return obj.O }
func (obj AllDisjointProperties) ID() string                 { return obj.O.ID() }
func (obj AllDisjointProperties) Value() string              { return obj.O.Value() }
func (obj AllDisjointProperties) Type() []string             { return obj.O.Type() }
func (obj AllDisjointProperties) Get(key string) interface{} { return obj.O.Get(key) }

func (obj AllDisjointProperties) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of annotated annotations for which the RDF serialization consists of an annotated subject, predicate and object.
type Annotation struct{ O *ld.Object }

func (obj Annotation) Obj() *ld.Object            { return obj.O }
func (obj Annotation) ID() string                 { return obj.O.ID() }
func (obj Annotation) Value() string              { return obj.O.Value() }
func (obj Annotation) Type() []string             { return obj.O.Type() }
func (obj Annotation) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Annotation) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of annotation properties.
type AnnotationProperty struct{ O *ld.Object }

func (obj AnnotationProperty) Obj() *ld.Object            { return obj.O }
func (obj AnnotationProperty) ID() string                 { return obj.O.ID() }
func (obj AnnotationProperty) Value() string              { return obj.O.Value() }
func (obj AnnotationProperty) Type() []string             { return obj.O.Type() }
func (obj AnnotationProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj AnnotationProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of asymmetric properties.
type AsymmetricProperty struct{ O *ld.Object }

func (obj AsymmetricProperty) Obj() *ld.Object            { return obj.O }
func (obj AsymmetricProperty) ID() string                 { return obj.O.ID() }
func (obj AsymmetricProperty) Value() string              { return obj.O.Value() }
func (obj AsymmetricProperty) Type() []string             { return obj.O.Type() }
func (obj AsymmetricProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj AsymmetricProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of annotated axioms for which the RDF serialization consists of an annotated subject, predicate and object.
type Axiom struct{ O *ld.Object }

func (obj Axiom) Obj() *ld.Object            { return obj.O }
func (obj Axiom) ID() string                 { return obj.O.ID() }
func (obj Axiom) Value() string              { return obj.O.Value() }
func (obj Axiom) Type() []string             { return obj.O.Type() }
func (obj Axiom) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Axiom) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of OWL classes.
type Class struct{ O *ld.Object }

func (obj Class) Obj() *ld.Object            { return obj.O }
func (obj Class) ID() string                 { return obj.O.ID() }
func (obj Class) Value() string              { return obj.O.Value() }
func (obj Class) Type() []string             { return obj.O.Type() }
func (obj Class) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Class) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

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
type DataRange struct{ O *ld.Object }

func (obj DataRange) Obj() *ld.Object            { return obj.O }
func (obj DataRange) ID() string                 { return obj.O.ID() }
func (obj DataRange) Value() string              { return obj.O.Value() }
func (obj DataRange) Type() []string             { return obj.O.Type() }
func (obj DataRange) Get(key string) interface{} { return obj.O.Get(key) }

func (obj DataRange) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of data properties.
type DatatypeProperty struct{ O *ld.Object }

func (obj DatatypeProperty) Obj() *ld.Object            { return obj.O }
func (obj DatatypeProperty) ID() string                 { return obj.O.ID() }
func (obj DatatypeProperty) Value() string              { return obj.O.Value() }
func (obj DatatypeProperty) Type() []string             { return obj.O.Type() }
func (obj DatatypeProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj DatatypeProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of deprecated classes.
type DeprecatedClass struct{ O *ld.Object }

func (obj DeprecatedClass) Obj() *ld.Object            { return obj.O }
func (obj DeprecatedClass) ID() string                 { return obj.O.ID() }
func (obj DeprecatedClass) Value() string              { return obj.O.Value() }
func (obj DeprecatedClass) Type() []string             { return obj.O.Type() }
func (obj DeprecatedClass) Get(key string) interface{} { return obj.O.Get(key) }

func (obj DeprecatedClass) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of deprecated properties.
type DeprecatedProperty struct{ O *ld.Object }

func (obj DeprecatedProperty) Obj() *ld.Object            { return obj.O }
func (obj DeprecatedProperty) ID() string                 { return obj.O.ID() }
func (obj DeprecatedProperty) Value() string              { return obj.O.Value() }
func (obj DeprecatedProperty) Type() []string             { return obj.O.Type() }
func (obj DeprecatedProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj DeprecatedProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of functional properties.
type FunctionalProperty struct{ O *ld.Object }

func (obj FunctionalProperty) Obj() *ld.Object            { return obj.O }
func (obj FunctionalProperty) ID() string                 { return obj.O.ID() }
func (obj FunctionalProperty) Value() string              { return obj.O.Value() }
func (obj FunctionalProperty) Type() []string             { return obj.O.Type() }
func (obj FunctionalProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj FunctionalProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of inverse-functional properties.
type InverseFunctionalProperty struct{ O *ld.Object }

func (obj InverseFunctionalProperty) Obj() *ld.Object            { return obj.O }
func (obj InverseFunctionalProperty) ID() string                 { return obj.O.ID() }
func (obj InverseFunctionalProperty) Value() string              { return obj.O.Value() }
func (obj InverseFunctionalProperty) Type() []string             { return obj.O.Type() }
func (obj InverseFunctionalProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj InverseFunctionalProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of irreflexive properties.
type IrreflexiveProperty struct{ O *ld.Object }

func (obj IrreflexiveProperty) Obj() *ld.Object            { return obj.O }
func (obj IrreflexiveProperty) ID() string                 { return obj.O.ID() }
func (obj IrreflexiveProperty) Value() string              { return obj.O.Value() }
func (obj IrreflexiveProperty) Type() []string             { return obj.O.Type() }
func (obj IrreflexiveProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj IrreflexiveProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of named individuals.
type NamedIndividual struct{ O *ld.Object }

func (obj NamedIndividual) Obj() *ld.Object            { return obj.O }
func (obj NamedIndividual) ID() string                 { return obj.O.ID() }
func (obj NamedIndividual) Value() string              { return obj.O.Value() }
func (obj NamedIndividual) Type() []string             { return obj.O.Type() }
func (obj NamedIndividual) Get(key string) interface{} { return obj.O.Get(key) }

func (obj NamedIndividual) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of negative property assertions.
type NegativePropertyAssertion struct{ O *ld.Object }

func (obj NegativePropertyAssertion) Obj() *ld.Object            { return obj.O }
func (obj NegativePropertyAssertion) ID() string                 { return obj.O.ID() }
func (obj NegativePropertyAssertion) Value() string              { return obj.O.Value() }
func (obj NegativePropertyAssertion) Type() []string             { return obj.O.Type() }
func (obj NegativePropertyAssertion) Get(key string) interface{} { return obj.O.Get(key) }

func (obj NegativePropertyAssertion) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

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
type ObjectProperty struct{ O *ld.Object }

func (obj ObjectProperty) Obj() *ld.Object            { return obj.O }
func (obj ObjectProperty) ID() string                 { return obj.O.ID() }
func (obj ObjectProperty) Value() string              { return obj.O.Value() }
func (obj ObjectProperty) Type() []string             { return obj.O.Type() }
func (obj ObjectProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj ObjectProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The property that determines that two given properties are inverse.
func (obj ObjectProperty) InverseOf() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#inverseOf"]
}

// The property that determines the n-tuple of properties that build a sub property chain of a given property.
func (obj ObjectProperty) PropertyChainAxiom() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#propertyChainAxiom"]
}

// The class of ontologies.
type Ontology struct{ O *ld.Object }

func (obj Ontology) Obj() *ld.Object            { return obj.O }
func (obj Ontology) ID() string                 { return obj.O.ID() }
func (obj Ontology) Value() string              { return obj.O.Value() }
func (obj Ontology) Type() []string             { return obj.O.Type() }
func (obj Ontology) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Ontology) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of ontology properties.
type OntologyProperty struct{ O *ld.Object }

func (obj OntologyProperty) Obj() *ld.Object            { return obj.O }
func (obj OntologyProperty) ID() string                 { return obj.O.ID() }
func (obj OntologyProperty) Value() string              { return obj.O.Value() }
func (obj OntologyProperty) Type() []string             { return obj.O.Type() }
func (obj OntologyProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj OntologyProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of reflexive properties.
type ReflexiveProperty struct{ O *ld.Object }

func (obj ReflexiveProperty) Obj() *ld.Object            { return obj.O }
func (obj ReflexiveProperty) ID() string                 { return obj.O.ID() }
func (obj ReflexiveProperty) Value() string              { return obj.O.Value() }
func (obj ReflexiveProperty) Type() []string             { return obj.O.Type() }
func (obj ReflexiveProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj ReflexiveProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of property restrictions.
type Restriction struct{ O *ld.Object }

func (obj Restriction) Obj() *ld.Object            { return obj.O }
func (obj Restriction) ID() string                 { return obj.O.ID() }
func (obj Restriction) Value() string              { return obj.O.Value() }
func (obj Restriction) Type() []string             { return obj.O.Type() }
func (obj Restriction) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Restriction) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

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
type SymmetricProperty struct{ O *ld.Object }

func (obj SymmetricProperty) Obj() *ld.Object            { return obj.O }
func (obj SymmetricProperty) ID() string                 { return obj.O.ID() }
func (obj SymmetricProperty) Value() string              { return obj.O.Value() }
func (obj SymmetricProperty) Type() []string             { return obj.O.Type() }
func (obj SymmetricProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj SymmetricProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The class of transitive properties.
type TransitiveProperty struct{ O *ld.Object }

func (obj TransitiveProperty) Obj() *ld.Object            { return obj.O }
func (obj TransitiveProperty) ID() string                 { return obj.O.ID() }
func (obj TransitiveProperty) Value() string              { return obj.O.Value() }
func (obj TransitiveProperty) Type() []string             { return obj.O.Type() }
func (obj TransitiveProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj TransitiveProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

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
