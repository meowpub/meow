// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package owl

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/rdf"
)

// The class of collections of pairwise different individuals.
type AllDifferent struct{ rdf.Resource }

// Ducktypes the object into a(n) AllDifferent.
func AsAllDifferent(e ld.Entity) AllDifferent { return AllDifferent{rdf.AsResource(e)} }

// Does the object quack like a(n) AllDifferent?
func IsAllDifferent(e ld.Entity) bool { return ld.Is(e, TypeAllDifferent) }

// The property that determines the collection of pairwise different individuals in a owl:AllDifferent axiom.
func (obj AllDifferent) DistinctMembers() interface{} { return GetDistinctMembers(obj) }

func (obj AllDifferent) SetDistinctMembers(v interface{}) { SetDistinctMembers(obj, v) }

// The class of collections of pairwise disjoint classes.
type AllDisjointClasses struct{ rdf.Resource }

// Ducktypes the object into a(n) AllDisjointClasses.
func AsAllDisjointClasses(e ld.Entity) AllDisjointClasses {
	return AllDisjointClasses{rdf.AsResource(e)}
}

// Does the object quack like a(n) AllDisjointClasses?
func IsAllDisjointClasses(e ld.Entity) bool { return ld.Is(e, TypeAllDisjointClasses) }

// The class of collections of pairwise disjoint properties.
type AllDisjointProperties struct{ rdf.Resource }

// Ducktypes the object into a(n) AllDisjointProperties.
func AsAllDisjointProperties(e ld.Entity) AllDisjointProperties {
	return AllDisjointProperties{rdf.AsResource(e)}
}

// Does the object quack like a(n) AllDisjointProperties?
func IsAllDisjointProperties(e ld.Entity) bool { return ld.Is(e, TypeAllDisjointProperties) }

// The class of annotated annotations for which the RDF serialization consists of an annotated subject, predicate and object.
type Annotation struct{ rdf.Resource }

// Ducktypes the object into a(n) Annotation.
func AsAnnotation(e ld.Entity) Annotation { return Annotation{rdf.AsResource(e)} }

// Does the object quack like a(n) Annotation?
func IsAnnotation(e ld.Entity) bool { return ld.Is(e, TypeAnnotation) }

// The class of annotation properties.
type AnnotationProperty struct{ rdf.Property }

// Ducktypes the object into a(n) AnnotationProperty.
func AsAnnotationProperty(e ld.Entity) AnnotationProperty {
	return AnnotationProperty{rdf.AsProperty(e)}
}

// Does the object quack like a(n) AnnotationProperty?
func IsAnnotationProperty(e ld.Entity) bool { return ld.Is(e, TypeAnnotationProperty) }

// The class of asymmetric properties.
type AsymmetricProperty struct{ ObjectProperty }

// Ducktypes the object into a(n) AsymmetricProperty.
func AsAsymmetricProperty(e ld.Entity) AsymmetricProperty {
	return AsymmetricProperty{AsObjectProperty(e)}
}

// Does the object quack like a(n) AsymmetricProperty?
func IsAsymmetricProperty(e ld.Entity) bool { return ld.Is(e, TypeAsymmetricProperty) }

// The class of annotated axioms for which the RDF serialization consists of an annotated subject, predicate and object.
type Axiom struct{ rdf.Resource }

// Ducktypes the object into a(n) Axiom.
func AsAxiom(e ld.Entity) Axiom { return Axiom{rdf.AsResource(e)} }

// Does the object quack like a(n) Axiom?
func IsAxiom(e ld.Entity) bool { return ld.Is(e, TypeAxiom) }

// The class of OWL classes.
type Class struct{ rdf.Class }

// Ducktypes the object into a(n) Class.
func AsClass(e ld.Entity) Class { return Class{rdf.AsClass(e)} }

// Does the object quack like a(n) Class?
func IsClass(e ld.Entity) bool { return ld.Is(e, TypeClass) }

// The property that determines that a given class is the complement of another class.
func (obj Class) ComplementOf() interface{} { return GetComplementOf(obj) }

func (obj Class) SetComplementOf(v interface{}) { SetComplementOf(obj, v) }

// The property that determines that a given class is equivalent to the disjoint union of a collection of other classes.
func (obj Class) DisjointUnionOf() interface{} { return GetDisjointUnionOf(obj) }

func (obj Class) SetDisjointUnionOf(v interface{}) { SetDisjointUnionOf(obj, v) }

// The property that determines that two given classes are disjoint.
func (obj Class) DisjointWith() interface{} { return GetDisjointWith(obj) }

func (obj Class) SetDisjointWith(v interface{}) { SetDisjointWith(obj, v) }

// The property that determines the collection of properties that jointly build a key.
func (obj Class) HasKey() interface{} { return GetHasKey(obj) }

func (obj Class) SetHasKey(v interface{}) { SetHasKey(obj, v) }

// The class of OWL data ranges, which are special kinds of datatypes. Note: The use of the IRI owl:DataRange has been deprecated as of OWL 2. The IRI rdfs:Datatype SHOULD be used instead.
type DataRange struct{ rdf.Datatype }

// Ducktypes the object into a(n) DataRange.
func AsDataRange(e ld.Entity) DataRange { return DataRange{rdf.AsDatatype(e)} }

// Does the object quack like a(n) DataRange?
func IsDataRange(e ld.Entity) bool { return ld.Is(e, TypeDataRange) }

// The class of data properties.
type DatatypeProperty struct{ rdf.Property }

// Ducktypes the object into a(n) DatatypeProperty.
func AsDatatypeProperty(e ld.Entity) DatatypeProperty { return DatatypeProperty{rdf.AsProperty(e)} }

// Does the object quack like a(n) DatatypeProperty?
func IsDatatypeProperty(e ld.Entity) bool { return ld.Is(e, TypeDatatypeProperty) }

// The class of deprecated classes.
type DeprecatedClass struct{ rdf.Class }

// Ducktypes the object into a(n) DeprecatedClass.
func AsDeprecatedClass(e ld.Entity) DeprecatedClass { return DeprecatedClass{rdf.AsClass(e)} }

// Does the object quack like a(n) DeprecatedClass?
func IsDeprecatedClass(e ld.Entity) bool { return ld.Is(e, TypeDeprecatedClass) }

// The class of deprecated properties.
type DeprecatedProperty struct{ rdf.Property }

// Ducktypes the object into a(n) DeprecatedProperty.
func AsDeprecatedProperty(e ld.Entity) DeprecatedProperty {
	return DeprecatedProperty{rdf.AsProperty(e)}
}

// Does the object quack like a(n) DeprecatedProperty?
func IsDeprecatedProperty(e ld.Entity) bool { return ld.Is(e, TypeDeprecatedProperty) }

// The class of functional properties.
type FunctionalProperty struct{ rdf.Property }

// Ducktypes the object into a(n) FunctionalProperty.
func AsFunctionalProperty(e ld.Entity) FunctionalProperty {
	return FunctionalProperty{rdf.AsProperty(e)}
}

// Does the object quack like a(n) FunctionalProperty?
func IsFunctionalProperty(e ld.Entity) bool { return ld.Is(e, TypeFunctionalProperty) }

// The class of inverse-functional properties.
type InverseFunctionalProperty struct{ ObjectProperty }

// Ducktypes the object into a(n) InverseFunctionalProperty.
func AsInverseFunctionalProperty(e ld.Entity) InverseFunctionalProperty {
	return InverseFunctionalProperty{AsObjectProperty(e)}
}

// Does the object quack like a(n) InverseFunctionalProperty?
func IsInverseFunctionalProperty(e ld.Entity) bool { return ld.Is(e, TypeInverseFunctionalProperty) }

// The class of irreflexive properties.
type IrreflexiveProperty struct{ ObjectProperty }

// Ducktypes the object into a(n) IrreflexiveProperty.
func AsIrreflexiveProperty(e ld.Entity) IrreflexiveProperty {
	return IrreflexiveProperty{AsObjectProperty(e)}
}

// Does the object quack like a(n) IrreflexiveProperty?
func IsIrreflexiveProperty(e ld.Entity) bool { return ld.Is(e, TypeIrreflexiveProperty) }

// The class of named individuals.
type NamedIndividual struct{ Thing }

// Ducktypes the object into a(n) NamedIndividual.
func AsNamedIndividual(e ld.Entity) NamedIndividual { return NamedIndividual{AsThing(e)} }

// Does the object quack like a(n) NamedIndividual?
func IsNamedIndividual(e ld.Entity) bool { return ld.Is(e, TypeNamedIndividual) }

// The class of negative property assertions.
type NegativePropertyAssertion struct{ rdf.Resource }

// Ducktypes the object into a(n) NegativePropertyAssertion.
func AsNegativePropertyAssertion(e ld.Entity) NegativePropertyAssertion {
	return NegativePropertyAssertion{rdf.AsResource(e)}
}

// Does the object quack like a(n) NegativePropertyAssertion?
func IsNegativePropertyAssertion(e ld.Entity) bool { return ld.Is(e, TypeNegativePropertyAssertion) }

// The property that determines the predicate of a negative property assertion.
func (obj NegativePropertyAssertion) AssertionProperty() interface{} { return GetAssertionProperty(obj) }

func (obj NegativePropertyAssertion) SetAssertionProperty(v interface{}) { SetAssertionProperty(obj, v) }

// The property that determines the subject of a negative property assertion.
func (obj NegativePropertyAssertion) SourceIndividual() interface{} { return GetSourceIndividual(obj) }

func (obj NegativePropertyAssertion) SetSourceIndividual(v interface{}) { SetSourceIndividual(obj, v) }

// The property that determines the object of a negative object property assertion.
func (obj NegativePropertyAssertion) TargetIndividual() interface{} { return GetTargetIndividual(obj) }

func (obj NegativePropertyAssertion) SetTargetIndividual(v interface{}) { SetTargetIndividual(obj, v) }

// The property that determines the value of a negative data property assertion.
func (obj NegativePropertyAssertion) TargetValue() interface{} { return GetTargetValue(obj) }

func (obj NegativePropertyAssertion) SetTargetValue(v interface{}) { SetTargetValue(obj, v) }

// This is the empty class.
type Nothing struct{ Thing }

// Ducktypes the object into a(n) Nothing.
func AsNothing(e ld.Entity) Nothing { return Nothing{AsThing(e)} }

// Does the object quack like a(n) Nothing?
func IsNothing(e ld.Entity) bool { return ld.Is(e, TypeNothing) }

// The class of object properties.
type ObjectProperty struct{ rdf.Property }

// Ducktypes the object into a(n) ObjectProperty.
func AsObjectProperty(e ld.Entity) ObjectProperty { return ObjectProperty{rdf.AsProperty(e)} }

// Does the object quack like a(n) ObjectProperty?
func IsObjectProperty(e ld.Entity) bool { return ld.Is(e, TypeObjectProperty) }

// The property that determines that two given properties are inverse.
func (obj ObjectProperty) InverseOf() interface{} { return GetInverseOf(obj) }

func (obj ObjectProperty) SetInverseOf(v interface{}) { SetInverseOf(obj, v) }

// The property that determines the n-tuple of properties that build a sub property chain of a given property.
func (obj ObjectProperty) PropertyChainAxiom() interface{} { return GetPropertyChainAxiom(obj) }

func (obj ObjectProperty) SetPropertyChainAxiom(v interface{}) { SetPropertyChainAxiom(obj, v) }

// The class of ontologies.
type Ontology struct{ rdf.Resource }

// Ducktypes the object into a(n) Ontology.
func AsOntology(e ld.Entity) Ontology { return Ontology{rdf.AsResource(e)} }

// Does the object quack like a(n) Ontology?
func IsOntology(e ld.Entity) bool { return ld.Is(e, TypeOntology) }

// The annotation property that indicates that a given ontology is backward compatible with another ontology.
func (obj Ontology) BackwardCompatibleWith() interface{} { return GetBackwardCompatibleWith(obj) }

func (obj Ontology) SetBackwardCompatibleWith(v interface{}) { SetBackwardCompatibleWith(obj, v) }

// The property that is used for importing other ontologies into a given ontology.
func (obj Ontology) Imports() interface{} { return GetImports(obj) }

func (obj Ontology) SetImports(v interface{}) { SetImports(obj, v) }

// The annotation property that indicates that a given ontology is incompatible with another ontology.
func (obj Ontology) IncompatibleWith() interface{} { return GetIncompatibleWith(obj) }

func (obj Ontology) SetIncompatibleWith(v interface{}) { SetIncompatibleWith(obj, v) }

// The annotation property that indicates the predecessor ontology of a given ontology.
func (obj Ontology) PriorVersion() interface{} { return GetPriorVersion(obj) }

func (obj Ontology) SetPriorVersion(v interface{}) { SetPriorVersion(obj, v) }

// The property that identifies the version IRI of an ontology.
func (obj Ontology) VersionIRI() interface{} { return GetVersionIRI(obj) }

func (obj Ontology) SetVersionIRI(v interface{}) { SetVersionIRI(obj, v) }

// The class of ontology properties.
type OntologyProperty struct{ rdf.Property }

// Ducktypes the object into a(n) OntologyProperty.
func AsOntologyProperty(e ld.Entity) OntologyProperty { return OntologyProperty{rdf.AsProperty(e)} }

// Does the object quack like a(n) OntologyProperty?
func IsOntologyProperty(e ld.Entity) bool { return ld.Is(e, TypeOntologyProperty) }

// The class of reflexive properties.
type ReflexiveProperty struct{ ObjectProperty }

// Ducktypes the object into a(n) ReflexiveProperty.
func AsReflexiveProperty(e ld.Entity) ReflexiveProperty { return ReflexiveProperty{AsObjectProperty(e)} }

// Does the object quack like a(n) ReflexiveProperty?
func IsReflexiveProperty(e ld.Entity) bool { return ld.Is(e, TypeReflexiveProperty) }

// The class of property restrictions.
type Restriction struct{ Class }

// Ducktypes the object into a(n) Restriction.
func AsRestriction(e ld.Entity) Restriction { return Restriction{AsClass(e)} }

// Does the object quack like a(n) Restriction?
func IsRestriction(e ld.Entity) bool { return ld.Is(e, TypeRestriction) }

// The property that determines the class that a universal property restriction refers to.
func (obj Restriction) AllValuesFrom() interface{} { return GetAllValuesFrom(obj) }

func (obj Restriction) SetAllValuesFrom(v interface{}) { SetAllValuesFrom(obj, v) }

// The property that determines the cardinality of an exact cardinality restriction.
func (obj Restriction) Cardinality() interface{} { return GetCardinality(obj) }

func (obj Restriction) SetCardinality(v interface{}) { SetCardinality(obj, v) }

// The property that determines the property that a self restriction refers to.
func (obj Restriction) HasSelf() interface{} { return GetHasSelf(obj) }

func (obj Restriction) SetHasSelf(v interface{}) { SetHasSelf(obj, v) }

// The property that determines the individual that a has-value restriction refers to.
func (obj Restriction) HasValue() interface{} { return GetHasValue(obj) }

func (obj Restriction) SetHasValue(v interface{}) { SetHasValue(obj, v) }

// The property that determines the cardinality of a maximum cardinality restriction.
func (obj Restriction) MaxCardinality() interface{} { return GetMaxCardinality(obj) }

func (obj Restriction) SetMaxCardinality(v interface{}) { SetMaxCardinality(obj, v) }

// The property that determines the cardinality of a maximum qualified cardinality restriction.
func (obj Restriction) MaxQualifiedCardinality() interface{} { return GetMaxQualifiedCardinality(obj) }

func (obj Restriction) SetMaxQualifiedCardinality(v interface{}) { SetMaxQualifiedCardinality(obj, v) }

// The property that determines the cardinality of a minimum cardinality restriction.
func (obj Restriction) MinCardinality() interface{} { return GetMinCardinality(obj) }

func (obj Restriction) SetMinCardinality(v interface{}) { SetMinCardinality(obj, v) }

// The property that determines the cardinality of a minimum qualified cardinality restriction.
func (obj Restriction) MinQualifiedCardinality() interface{} { return GetMinQualifiedCardinality(obj) }

func (obj Restriction) SetMinQualifiedCardinality(v interface{}) { SetMinQualifiedCardinality(obj, v) }

// The property that determines the class that a qualified object cardinality restriction refers to.
func (obj Restriction) OnClass() interface{} { return GetOnClass(obj) }

func (obj Restriction) SetOnClass(v interface{}) { SetOnClass(obj, v) }

// The property that determines the data range that a qualified data cardinality restriction refers to.
func (obj Restriction) OnDataRange() interface{} { return GetOnDataRange(obj) }

func (obj Restriction) SetOnDataRange(v interface{}) { SetOnDataRange(obj, v) }

// The property that determines the n-tuple of properties that a property restriction on an n-ary data range refers to.
func (obj Restriction) OnProperties() interface{} { return GetOnProperties(obj) }

func (obj Restriction) SetOnProperties(v interface{}) { SetOnProperties(obj, v) }

// The property that determines the property that a property restriction refers to.
func (obj Restriction) OnProperty() interface{} { return GetOnProperty(obj) }

func (obj Restriction) SetOnProperty(v interface{}) { SetOnProperty(obj, v) }

// The property that determines the cardinality of an exact qualified cardinality restriction.
func (obj Restriction) QualifiedCardinality() interface{} { return GetQualifiedCardinality(obj) }

func (obj Restriction) SetQualifiedCardinality(v interface{}) { SetQualifiedCardinality(obj, v) }

// The property that determines the class that an existential property restriction refers to.
func (obj Restriction) SomeValuesFrom() interface{} { return GetSomeValuesFrom(obj) }

func (obj Restriction) SetSomeValuesFrom(v interface{}) { SetSomeValuesFrom(obj, v) }

// The class of symmetric properties.
type SymmetricProperty struct{ ObjectProperty }

// Ducktypes the object into a(n) SymmetricProperty.
func AsSymmetricProperty(e ld.Entity) SymmetricProperty { return SymmetricProperty{AsObjectProperty(e)} }

// Does the object quack like a(n) SymmetricProperty?
func IsSymmetricProperty(e ld.Entity) bool { return ld.Is(e, TypeSymmetricProperty) }

// The class of OWL individuals.
type Thing struct{ o *ld.Object }

// Ducktypes the object into a(n) Thing.
func AsThing(e ld.Entity) Thing { return Thing{o: e.Obj()} }

// Does the object quack like a(n) Thing?
func IsThing(e ld.Entity) bool { return ld.Is(e, TypeThing) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Thing) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj Thing) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Thing) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Thing) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Thing) Get(key string) interface{} { return obj.o.Get(key) }

// Sets the named attribute. Implements ld.Entity.
func (obj Thing) Set(key string, v interface{}) { obj.o.Set(key, v) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Thing) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// The data property that does not relate any individual to any data value.
func (obj Thing) BottomDataProperty() interface{} { return GetBottomDataProperty(obj) }

func (obj Thing) SetBottomDataProperty(v interface{}) { SetBottomDataProperty(obj, v) }

// The object property that does not relate any two individuals.
func (obj Thing) BottomObjectProperty() interface{} { return GetBottomObjectProperty(obj) }

func (obj Thing) SetBottomObjectProperty(v interface{}) { SetBottomObjectProperty(obj, v) }

// The property that determines that two given individuals are different.
func (obj Thing) DifferentFrom() interface{} { return GetDifferentFrom(obj) }

func (obj Thing) SetDifferentFrom(v interface{}) { SetDifferentFrom(obj, v) }

// The property that determines that two given individuals are equal.
func (obj Thing) SameAs() interface{} { return GetSameAs(obj) }

func (obj Thing) SetSameAs(v interface{}) { SetSameAs(obj, v) }

// The data property that relates every individual to every data value.
func (obj Thing) TopDataProperty() interface{} { return GetTopDataProperty(obj) }

func (obj Thing) SetTopDataProperty(v interface{}) { SetTopDataProperty(obj, v) }

// The object property that relates every two individuals.
func (obj Thing) TopObjectProperty() interface{} { return GetTopObjectProperty(obj) }

func (obj Thing) SetTopObjectProperty(v interface{}) { SetTopObjectProperty(obj, v) }

// The class of transitive properties.
type TransitiveProperty struct{ ObjectProperty }

// Ducktypes the object into a(n) TransitiveProperty.
func AsTransitiveProperty(e ld.Entity) TransitiveProperty {
	return TransitiveProperty{AsObjectProperty(e)}
}

// Does the object quack like a(n) TransitiveProperty?
func IsTransitiveProperty(e ld.Entity) bool { return ld.Is(e, TypeTransitiveProperty) }

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
	_ ld.Entity = Nothing{}
	_ ld.Entity = ObjectProperty{}
	_ ld.Entity = Ontology{}
	_ ld.Entity = OntologyProperty{}
	_ ld.Entity = ReflexiveProperty{}
	_ ld.Entity = Restriction{}
	_ ld.Entity = SymmetricProperty{}
	_ ld.Entity = Thing{}
	_ ld.Entity = TransitiveProperty{}
)
