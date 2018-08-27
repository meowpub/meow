// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package owl

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/rdf"
)

// The class of collections of pairwise different individuals.
type AllDifferent struct{ rdf.Resource }

// Ducktypes the object into a AllDifferent.
func AsAllDifferent(obj *ld.Object) AllDifferent { return AllDifferent{rdf.AsResource(obj)} }

// Does the object quack like a AllDifferent?
func IsAllDifferent(obj *ld.Object) bool { return ld.Is(obj, TypeAllDifferent) }

// The property that determines the collection of pairwise different individuals in a owl:AllDifferent axiom.
func (obj AllDifferent) DistinctMembers() interface{} { return obj.Get(PropDistinctMembers) }

// The class of collections of pairwise disjoint classes.
type AllDisjointClasses struct{ rdf.Resource }

// Ducktypes the object into a AllDisjointClasses.
func AsAllDisjointClasses(obj *ld.Object) AllDisjointClasses {
	return AllDisjointClasses{rdf.AsResource(obj)}
}

// Does the object quack like a AllDisjointClasses?
func IsAllDisjointClasses(obj *ld.Object) bool { return ld.Is(obj, TypeAllDisjointClasses) }

// The class of collections of pairwise disjoint properties.
type AllDisjointProperties struct{ rdf.Resource }

// Ducktypes the object into a AllDisjointProperties.
func AsAllDisjointProperties(obj *ld.Object) AllDisjointProperties {
	return AllDisjointProperties{rdf.AsResource(obj)}
}

// Does the object quack like a AllDisjointProperties?
func IsAllDisjointProperties(obj *ld.Object) bool { return ld.Is(obj, TypeAllDisjointProperties) }

// The class of annotated annotations for which the RDF serialization consists of an annotated subject, predicate and object.
type Annotation struct{ rdf.Resource }

// Ducktypes the object into a Annotation.
func AsAnnotation(obj *ld.Object) Annotation { return Annotation{rdf.AsResource(obj)} }

// Does the object quack like a Annotation?
func IsAnnotation(obj *ld.Object) bool { return ld.Is(obj, TypeAnnotation) }

// The class of annotation properties.
type AnnotationProperty struct{ rdf.Property }

// Ducktypes the object into a AnnotationProperty.
func AsAnnotationProperty(obj *ld.Object) AnnotationProperty {
	return AnnotationProperty{rdf.AsProperty(obj)}
}

// Does the object quack like a AnnotationProperty?
func IsAnnotationProperty(obj *ld.Object) bool { return ld.Is(obj, TypeAnnotationProperty) }

// The class of asymmetric properties.
type AsymmetricProperty struct{ ObjectProperty }

// Ducktypes the object into a AsymmetricProperty.
func AsAsymmetricProperty(obj *ld.Object) AsymmetricProperty {
	return AsymmetricProperty{AsObjectProperty(obj)}
}

// Does the object quack like a AsymmetricProperty?
func IsAsymmetricProperty(obj *ld.Object) bool { return ld.Is(obj, TypeAsymmetricProperty) }

// The class of annotated axioms for which the RDF serialization consists of an annotated subject, predicate and object.
type Axiom struct{ rdf.Resource }

// Ducktypes the object into a Axiom.
func AsAxiom(obj *ld.Object) Axiom { return Axiom{rdf.AsResource(obj)} }

// Does the object quack like a Axiom?
func IsAxiom(obj *ld.Object) bool { return ld.Is(obj, TypeAxiom) }

// The class of OWL classes.
type Class struct{ rdf.Class }

// Ducktypes the object into a Class.
func AsClass(obj *ld.Object) Class { return Class{rdf.AsClass(obj)} }

// Does the object quack like a Class?
func IsClass(obj *ld.Object) bool { return ld.Is(obj, TypeClass) }

// The property that determines that a given class is the complement of another class.
func (obj Class) ComplementOf() interface{} { return obj.Get(PropComplementOf) }

// The property that determines that a given class is equivalent to the disjoint union of a collection of other classes.
func (obj Class) DisjointUnionOf() interface{} { return obj.Get(PropDisjointUnionOf) }

// The property that determines that two given classes are disjoint.
func (obj Class) DisjointWith() interface{} { return obj.Get(PropDisjointWith) }

// The property that determines the collection of properties that jointly build a key.
func (obj Class) HasKey() interface{} { return obj.Get(PropHasKey) }

// The class of OWL data ranges, which are special kinds of datatypes. Note: The use of the IRI owl:DataRange has been deprecated as of OWL 2. The IRI rdfs:Datatype SHOULD be used instead.
type DataRange struct{ rdf.Datatype }

// Ducktypes the object into a DataRange.
func AsDataRange(obj *ld.Object) DataRange { return DataRange{rdf.AsDatatype(obj)} }

// Does the object quack like a DataRange?
func IsDataRange(obj *ld.Object) bool { return ld.Is(obj, TypeDataRange) }

// The class of data properties.
type DatatypeProperty struct{ rdf.Property }

// Ducktypes the object into a DatatypeProperty.
func AsDatatypeProperty(obj *ld.Object) DatatypeProperty { return DatatypeProperty{rdf.AsProperty(obj)} }

// Does the object quack like a DatatypeProperty?
func IsDatatypeProperty(obj *ld.Object) bool { return ld.Is(obj, TypeDatatypeProperty) }

// The class of deprecated classes.
type DeprecatedClass struct{ rdf.Class }

// Ducktypes the object into a DeprecatedClass.
func AsDeprecatedClass(obj *ld.Object) DeprecatedClass { return DeprecatedClass{rdf.AsClass(obj)} }

// Does the object quack like a DeprecatedClass?
func IsDeprecatedClass(obj *ld.Object) bool { return ld.Is(obj, TypeDeprecatedClass) }

// The class of deprecated properties.
type DeprecatedProperty struct{ rdf.Property }

// Ducktypes the object into a DeprecatedProperty.
func AsDeprecatedProperty(obj *ld.Object) DeprecatedProperty {
	return DeprecatedProperty{rdf.AsProperty(obj)}
}

// Does the object quack like a DeprecatedProperty?
func IsDeprecatedProperty(obj *ld.Object) bool { return ld.Is(obj, TypeDeprecatedProperty) }

// The class of functional properties.
type FunctionalProperty struct{ rdf.Property }

// Ducktypes the object into a FunctionalProperty.
func AsFunctionalProperty(obj *ld.Object) FunctionalProperty {
	return FunctionalProperty{rdf.AsProperty(obj)}
}

// Does the object quack like a FunctionalProperty?
func IsFunctionalProperty(obj *ld.Object) bool { return ld.Is(obj, TypeFunctionalProperty) }

// The class of inverse-functional properties.
type InverseFunctionalProperty struct{ ObjectProperty }

// Ducktypes the object into a InverseFunctionalProperty.
func AsInverseFunctionalProperty(obj *ld.Object) InverseFunctionalProperty {
	return InverseFunctionalProperty{AsObjectProperty(obj)}
}

// Does the object quack like a InverseFunctionalProperty?
func IsInverseFunctionalProperty(obj *ld.Object) bool {
	return ld.Is(obj, TypeInverseFunctionalProperty)
}

// The class of irreflexive properties.
type IrreflexiveProperty struct{ ObjectProperty }

// Ducktypes the object into a IrreflexiveProperty.
func AsIrreflexiveProperty(obj *ld.Object) IrreflexiveProperty {
	return IrreflexiveProperty{AsObjectProperty(obj)}
}

// Does the object quack like a IrreflexiveProperty?
func IsIrreflexiveProperty(obj *ld.Object) bool { return ld.Is(obj, TypeIrreflexiveProperty) }

// The class of named individuals.
type NamedIndividual struct{ Thing }

// Ducktypes the object into a NamedIndividual.
func AsNamedIndividual(obj *ld.Object) NamedIndividual { return NamedIndividual{AsThing(obj)} }

// Does the object quack like a NamedIndividual?
func IsNamedIndividual(obj *ld.Object) bool { return ld.Is(obj, TypeNamedIndividual) }

// The class of negative property assertions.
type NegativePropertyAssertion struct{ rdf.Resource }

// Ducktypes the object into a NegativePropertyAssertion.
func AsNegativePropertyAssertion(obj *ld.Object) NegativePropertyAssertion {
	return NegativePropertyAssertion{rdf.AsResource(obj)}
}

// Does the object quack like a NegativePropertyAssertion?
func IsNegativePropertyAssertion(obj *ld.Object) bool {
	return ld.Is(obj, TypeNegativePropertyAssertion)
}

// The property that determines the predicate of a negative property assertion.
func (obj NegativePropertyAssertion) AssertionProperty() interface{} {
	return obj.Get(PropAssertionProperty)
}

// The property that determines the subject of a negative property assertion.
func (obj NegativePropertyAssertion) SourceIndividual() interface{} {
	return obj.Get(PropSourceIndividual)
}

// The property that determines the object of a negative object property assertion.
func (obj NegativePropertyAssertion) TargetIndividual() interface{} {
	return obj.Get(PropTargetIndividual)
}

// The property that determines the value of a negative data property assertion.
func (obj NegativePropertyAssertion) TargetValue() interface{} { return obj.Get(PropTargetValue) }

// This is the empty class.
type Nothing struct{ Thing }

// Ducktypes the object into a Nothing.
func AsNothing(obj *ld.Object) Nothing { return Nothing{AsThing(obj)} }

// Does the object quack like a Nothing?
func IsNothing(obj *ld.Object) bool { return ld.Is(obj, TypeNothing) }

// The class of object properties.
type ObjectProperty struct{ rdf.Property }

// Ducktypes the object into a ObjectProperty.
func AsObjectProperty(obj *ld.Object) ObjectProperty { return ObjectProperty{rdf.AsProperty(obj)} }

// Does the object quack like a ObjectProperty?
func IsObjectProperty(obj *ld.Object) bool { return ld.Is(obj, TypeObjectProperty) }

// The property that determines that two given properties are inverse.
func (obj ObjectProperty) InverseOf() interface{} { return obj.Get(PropInverseOf) }

// The property that determines the n-tuple of properties that build a sub property chain of a given property.
func (obj ObjectProperty) PropertyChainAxiom() interface{} { return obj.Get(PropPropertyChainAxiom) }

// The class of ontologies.
type Ontology struct{ rdf.Resource }

// Ducktypes the object into a Ontology.
func AsOntology(obj *ld.Object) Ontology { return Ontology{rdf.AsResource(obj)} }

// Does the object quack like a Ontology?
func IsOntology(obj *ld.Object) bool { return ld.Is(obj, TypeOntology) }

// The annotation property that indicates that a given ontology is backward compatible with another ontology.
func (obj Ontology) BackwardCompatibleWith() interface{} { return obj.Get(PropBackwardCompatibleWith) }

// The property that is used for importing other ontologies into a given ontology.
func (obj Ontology) Imports() interface{} { return obj.Get(PropImports) }

// The annotation property that indicates that a given ontology is incompatible with another ontology.
func (obj Ontology) IncompatibleWith() interface{} { return obj.Get(PropIncompatibleWith) }

// The annotation property that indicates the predecessor ontology of a given ontology.
func (obj Ontology) PriorVersion() interface{} { return obj.Get(PropPriorVersion) }

// The property that identifies the version IRI of an ontology.
func (obj Ontology) VersionIRI() interface{} { return obj.Get(PropVersionIRI) }

// The class of ontology properties.
type OntologyProperty struct{ rdf.Property }

// Ducktypes the object into a OntologyProperty.
func AsOntologyProperty(obj *ld.Object) OntologyProperty { return OntologyProperty{rdf.AsProperty(obj)} }

// Does the object quack like a OntologyProperty?
func IsOntologyProperty(obj *ld.Object) bool { return ld.Is(obj, TypeOntologyProperty) }

// The class of reflexive properties.
type ReflexiveProperty struct{ ObjectProperty }

// Ducktypes the object into a ReflexiveProperty.
func AsReflexiveProperty(obj *ld.Object) ReflexiveProperty {
	return ReflexiveProperty{AsObjectProperty(obj)}
}

// Does the object quack like a ReflexiveProperty?
func IsReflexiveProperty(obj *ld.Object) bool { return ld.Is(obj, TypeReflexiveProperty) }

// The class of property restrictions.
type Restriction struct{ Class }

// Ducktypes the object into a Restriction.
func AsRestriction(obj *ld.Object) Restriction { return Restriction{AsClass(obj)} }

// Does the object quack like a Restriction?
func IsRestriction(obj *ld.Object) bool { return ld.Is(obj, TypeRestriction) }

// The property that determines the class that a universal property restriction refers to.
func (obj Restriction) AllValuesFrom() interface{} { return obj.Get(PropAllValuesFrom) }

// The property that determines the cardinality of an exact cardinality restriction.
func (obj Restriction) Cardinality() interface{} { return obj.Get(PropCardinality) }

// The property that determines the property that a self restriction refers to.
func (obj Restriction) HasSelf() interface{} { return obj.Get(PropHasSelf) }

// The property that determines the individual that a has-value restriction refers to.
func (obj Restriction) HasValue() interface{} { return obj.Get(PropHasValue) }

// The property that determines the cardinality of a maximum cardinality restriction.
func (obj Restriction) MaxCardinality() interface{} { return obj.Get(PropMaxCardinality) }

// The property that determines the cardinality of a maximum qualified cardinality restriction.
func (obj Restriction) MaxQualifiedCardinality() interface{} {
	return obj.Get(PropMaxQualifiedCardinality)
}

// The property that determines the cardinality of a minimum cardinality restriction.
func (obj Restriction) MinCardinality() interface{} { return obj.Get(PropMinCardinality) }

// The property that determines the cardinality of a minimum qualified cardinality restriction.
func (obj Restriction) MinQualifiedCardinality() interface{} {
	return obj.Get(PropMinQualifiedCardinality)
}

// The property that determines the class that a qualified object cardinality restriction refers to.
func (obj Restriction) OnClass() interface{} { return obj.Get(PropOnClass) }

// The property that determines the data range that a qualified data cardinality restriction refers to.
func (obj Restriction) OnDataRange() interface{} { return obj.Get(PropOnDataRange) }

// The property that determines the n-tuple of properties that a property restriction on an n-ary data range refers to.
func (obj Restriction) OnProperties() interface{} { return obj.Get(PropOnProperties) }

// The property that determines the property that a property restriction refers to.
func (obj Restriction) OnProperty() interface{} { return obj.Get(PropOnProperty) }

// The property that determines the cardinality of an exact qualified cardinality restriction.
func (obj Restriction) QualifiedCardinality() interface{} { return obj.Get(PropQualifiedCardinality) }

// The property that determines the class that an existential property restriction refers to.
func (obj Restriction) SomeValuesFrom() interface{} { return obj.Get(PropSomeValuesFrom) }

// The class of symmetric properties.
type SymmetricProperty struct{ ObjectProperty }

// Ducktypes the object into a SymmetricProperty.
func AsSymmetricProperty(obj *ld.Object) SymmetricProperty {
	return SymmetricProperty{AsObjectProperty(obj)}
}

// Does the object quack like a SymmetricProperty?
func IsSymmetricProperty(obj *ld.Object) bool { return ld.Is(obj, TypeSymmetricProperty) }

// The class of OWL individuals.
type Thing struct{ o *ld.Object }

// Ducktypes the object into a Thing.
func AsThing(obj *ld.Object) Thing { return Thing{o: obj} }

// Does the object quack like a Thing?
func IsThing(obj *ld.Object) bool { return ld.Is(obj, TypeThing) }

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

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Thing) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// The data property that does not relate any individual to any data value.
func (obj Thing) BottomDataProperty() interface{} { return obj.Get(PropBottomDataProperty) }

// The object property that does not relate any two individuals.
func (obj Thing) BottomObjectProperty() interface{} { return obj.Get(PropBottomObjectProperty) }

// The property that determines that two given individuals are different.
func (obj Thing) DifferentFrom() interface{} { return obj.Get(PropDifferentFrom) }

// The property that determines that two given individuals are equal.
func (obj Thing) SameAs() interface{} { return obj.Get(PropSameAs) }

// The data property that relates every individual to every data value.
func (obj Thing) TopDataProperty() interface{} { return obj.Get(PropTopDataProperty) }

// The object property that relates every two individuals.
func (obj Thing) TopObjectProperty() interface{} { return obj.Get(PropTopObjectProperty) }

// The class of transitive properties.
type TransitiveProperty struct{ ObjectProperty }

// Ducktypes the object into a TransitiveProperty.
func AsTransitiveProperty(obj *ld.Object) TransitiveProperty {
	return TransitiveProperty{AsObjectProperty(obj)}
}

// Does the object quack like a TransitiveProperty?
func IsTransitiveProperty(obj *ld.Object) bool { return ld.Is(obj, TypeTransitiveProperty) }

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
