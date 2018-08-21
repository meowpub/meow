// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package owl

import (
	"github.com/meowpub/meow/ld"
)

type AllDifferent struct{ O *ld.Object }

func (obj AllDifferent) Obj() *ld.Object            { return obj.O }
func (obj AllDifferent) ID() string                 { return obj.O.ID() }
func (obj AllDifferent) Value() string              { return obj.O.Value() }
func (obj AllDifferent) Type() []string             { return obj.O.Type() }
func (obj AllDifferent) Get(key string) interface{} { return obj.O.Get(key) }

func (obj AllDifferent) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

func (obj AllDifferent) DistinctMembers() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#distinctMembers"]
}

type AllDisjointClasses struct{ O *ld.Object }

func (obj AllDisjointClasses) Obj() *ld.Object            { return obj.O }
func (obj AllDisjointClasses) ID() string                 { return obj.O.ID() }
func (obj AllDisjointClasses) Value() string              { return obj.O.Value() }
func (obj AllDisjointClasses) Type() []string             { return obj.O.Type() }
func (obj AllDisjointClasses) Get(key string) interface{} { return obj.O.Get(key) }

func (obj AllDisjointClasses) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type AllDisjointProperties struct{ O *ld.Object }

func (obj AllDisjointProperties) Obj() *ld.Object            { return obj.O }
func (obj AllDisjointProperties) ID() string                 { return obj.O.ID() }
func (obj AllDisjointProperties) Value() string              { return obj.O.Value() }
func (obj AllDisjointProperties) Type() []string             { return obj.O.Type() }
func (obj AllDisjointProperties) Get(key string) interface{} { return obj.O.Get(key) }

func (obj AllDisjointProperties) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type Annotation struct{ O *ld.Object }

func (obj Annotation) Obj() *ld.Object            { return obj.O }
func (obj Annotation) ID() string                 { return obj.O.ID() }
func (obj Annotation) Value() string              { return obj.O.Value() }
func (obj Annotation) Type() []string             { return obj.O.Type() }
func (obj Annotation) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Annotation) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type AnnotationProperty struct{ O *ld.Object }

func (obj AnnotationProperty) Obj() *ld.Object            { return obj.O }
func (obj AnnotationProperty) ID() string                 { return obj.O.ID() }
func (obj AnnotationProperty) Value() string              { return obj.O.Value() }
func (obj AnnotationProperty) Type() []string             { return obj.O.Type() }
func (obj AnnotationProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj AnnotationProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type AsymmetricProperty struct{ O *ld.Object }

func (obj AsymmetricProperty) Obj() *ld.Object            { return obj.O }
func (obj AsymmetricProperty) ID() string                 { return obj.O.ID() }
func (obj AsymmetricProperty) Value() string              { return obj.O.Value() }
func (obj AsymmetricProperty) Type() []string             { return obj.O.Type() }
func (obj AsymmetricProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj AsymmetricProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type Axiom struct{ O *ld.Object }

func (obj Axiom) Obj() *ld.Object            { return obj.O }
func (obj Axiom) ID() string                 { return obj.O.ID() }
func (obj Axiom) Value() string              { return obj.O.Value() }
func (obj Axiom) Type() []string             { return obj.O.Type() }
func (obj Axiom) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Axiom) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type Class struct{ O *ld.Object }

func (obj Class) Obj() *ld.Object            { return obj.O }
func (obj Class) ID() string                 { return obj.O.ID() }
func (obj Class) Value() string              { return obj.O.Value() }
func (obj Class) Type() []string             { return obj.O.Type() }
func (obj Class) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Class) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

func (obj Class) ComplementOf() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#complementOf"]
}

func (obj Class) DisjointUnionOf() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#disjointUnionOf"]
}

func (obj Class) DisjointWith() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#disjointWith"]
}

func (obj Class) HasKey() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#hasKey"]
}

type DataRange struct{ O *ld.Object }

func (obj DataRange) Obj() *ld.Object            { return obj.O }
func (obj DataRange) ID() string                 { return obj.O.ID() }
func (obj DataRange) Value() string              { return obj.O.Value() }
func (obj DataRange) Type() []string             { return obj.O.Type() }
func (obj DataRange) Get(key string) interface{} { return obj.O.Get(key) }

func (obj DataRange) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type DatatypeProperty struct{ O *ld.Object }

func (obj DatatypeProperty) Obj() *ld.Object            { return obj.O }
func (obj DatatypeProperty) ID() string                 { return obj.O.ID() }
func (obj DatatypeProperty) Value() string              { return obj.O.Value() }
func (obj DatatypeProperty) Type() []string             { return obj.O.Type() }
func (obj DatatypeProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj DatatypeProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type DeprecatedClass struct{ O *ld.Object }

func (obj DeprecatedClass) Obj() *ld.Object            { return obj.O }
func (obj DeprecatedClass) ID() string                 { return obj.O.ID() }
func (obj DeprecatedClass) Value() string              { return obj.O.Value() }
func (obj DeprecatedClass) Type() []string             { return obj.O.Type() }
func (obj DeprecatedClass) Get(key string) interface{} { return obj.O.Get(key) }

func (obj DeprecatedClass) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type DeprecatedProperty struct{ O *ld.Object }

func (obj DeprecatedProperty) Obj() *ld.Object            { return obj.O }
func (obj DeprecatedProperty) ID() string                 { return obj.O.ID() }
func (obj DeprecatedProperty) Value() string              { return obj.O.Value() }
func (obj DeprecatedProperty) Type() []string             { return obj.O.Type() }
func (obj DeprecatedProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj DeprecatedProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type FunctionalProperty struct{ O *ld.Object }

func (obj FunctionalProperty) Obj() *ld.Object            { return obj.O }
func (obj FunctionalProperty) ID() string                 { return obj.O.ID() }
func (obj FunctionalProperty) Value() string              { return obj.O.Value() }
func (obj FunctionalProperty) Type() []string             { return obj.O.Type() }
func (obj FunctionalProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj FunctionalProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type InverseFunctionalProperty struct{ O *ld.Object }

func (obj InverseFunctionalProperty) Obj() *ld.Object            { return obj.O }
func (obj InverseFunctionalProperty) ID() string                 { return obj.O.ID() }
func (obj InverseFunctionalProperty) Value() string              { return obj.O.Value() }
func (obj InverseFunctionalProperty) Type() []string             { return obj.O.Type() }
func (obj InverseFunctionalProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj InverseFunctionalProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type IrreflexiveProperty struct{ O *ld.Object }

func (obj IrreflexiveProperty) Obj() *ld.Object            { return obj.O }
func (obj IrreflexiveProperty) ID() string                 { return obj.O.ID() }
func (obj IrreflexiveProperty) Value() string              { return obj.O.Value() }
func (obj IrreflexiveProperty) Type() []string             { return obj.O.Type() }
func (obj IrreflexiveProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj IrreflexiveProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type NamedIndividual struct{ O *ld.Object }

func (obj NamedIndividual) Obj() *ld.Object            { return obj.O }
func (obj NamedIndividual) ID() string                 { return obj.O.ID() }
func (obj NamedIndividual) Value() string              { return obj.O.Value() }
func (obj NamedIndividual) Type() []string             { return obj.O.Type() }
func (obj NamedIndividual) Get(key string) interface{} { return obj.O.Get(key) }

func (obj NamedIndividual) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type NegativePropertyAssertion struct{ O *ld.Object }

func (obj NegativePropertyAssertion) Obj() *ld.Object            { return obj.O }
func (obj NegativePropertyAssertion) ID() string                 { return obj.O.ID() }
func (obj NegativePropertyAssertion) Value() string              { return obj.O.Value() }
func (obj NegativePropertyAssertion) Type() []string             { return obj.O.Type() }
func (obj NegativePropertyAssertion) Get(key string) interface{} { return obj.O.Get(key) }

func (obj NegativePropertyAssertion) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

func (obj NegativePropertyAssertion) AssertionProperty() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#assertionProperty"]
}

func (obj NegativePropertyAssertion) SourceIndividual() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#sourceIndividual"]
}

func (obj NegativePropertyAssertion) TargetIndividual() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#targetIndividual"]
}

func (obj NegativePropertyAssertion) TargetValue() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#targetValue"]
}

type ObjectProperty struct{ O *ld.Object }

func (obj ObjectProperty) Obj() *ld.Object            { return obj.O }
func (obj ObjectProperty) ID() string                 { return obj.O.ID() }
func (obj ObjectProperty) Value() string              { return obj.O.Value() }
func (obj ObjectProperty) Type() []string             { return obj.O.Type() }
func (obj ObjectProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj ObjectProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

func (obj ObjectProperty) InverseOf() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#inverseOf"]
}

func (obj ObjectProperty) PropertyChainAxiom() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#propertyChainAxiom"]
}

type Ontology struct{ O *ld.Object }

func (obj Ontology) Obj() *ld.Object            { return obj.O }
func (obj Ontology) ID() string                 { return obj.O.ID() }
func (obj Ontology) Value() string              { return obj.O.Value() }
func (obj Ontology) Type() []string             { return obj.O.Type() }
func (obj Ontology) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Ontology) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type OntologyProperty struct{ O *ld.Object }

func (obj OntologyProperty) Obj() *ld.Object            { return obj.O }
func (obj OntologyProperty) ID() string                 { return obj.O.ID() }
func (obj OntologyProperty) Value() string              { return obj.O.Value() }
func (obj OntologyProperty) Type() []string             { return obj.O.Type() }
func (obj OntologyProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj OntologyProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type ReflexiveProperty struct{ O *ld.Object }

func (obj ReflexiveProperty) Obj() *ld.Object            { return obj.O }
func (obj ReflexiveProperty) ID() string                 { return obj.O.ID() }
func (obj ReflexiveProperty) Value() string              { return obj.O.Value() }
func (obj ReflexiveProperty) Type() []string             { return obj.O.Type() }
func (obj ReflexiveProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj ReflexiveProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

type Restriction struct{ O *ld.Object }

func (obj Restriction) Obj() *ld.Object            { return obj.O }
func (obj Restriction) ID() string                 { return obj.O.ID() }
func (obj Restriction) Value() string              { return obj.O.Value() }
func (obj Restriction) Type() []string             { return obj.O.Type() }
func (obj Restriction) Get(key string) interface{} { return obj.O.Get(key) }

func (obj Restriction) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

func (obj Restriction) AllValuesFrom() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#allValuesFrom"]
}

func (obj Restriction) Cardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#cardinality"]
}

func (obj Restriction) HasSelf() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#hasSelf"]
}

func (obj Restriction) HasValue() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#hasValue"]
}

func (obj Restriction) MaxCardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#maxCardinality"]
}

func (obj Restriction) MaxQualifiedCardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#maxQualifiedCardinality"]
}

func (obj Restriction) MinCardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#minCardinality"]
}

func (obj Restriction) MinQualifiedCardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#minQualifiedCardinality"]
}

func (obj Restriction) OnClass() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#onClass"]
}

func (obj Restriction) OnDataRange() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#onDataRange"]
}

func (obj Restriction) OnProperties() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#onProperties"]
}

func (obj Restriction) OnProperty() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#onProperty"]
}

func (obj Restriction) QualifiedCardinality() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#qualifiedCardinality"]
}

func (obj Restriction) SomeValuesFrom() interface{} {
	return obj.O.V["http://www.w3.org/2002/07/owl#someValuesFrom"]
}

type SymmetricProperty struct{ O *ld.Object }

func (obj SymmetricProperty) Obj() *ld.Object            { return obj.O }
func (obj SymmetricProperty) ID() string                 { return obj.O.ID() }
func (obj SymmetricProperty) Value() string              { return obj.O.Value() }
func (obj SymmetricProperty) Type() []string             { return obj.O.Type() }
func (obj SymmetricProperty) Get(key string) interface{} { return obj.O.Get(key) }

func (obj SymmetricProperty) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

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
