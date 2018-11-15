// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package owl

import (
	"github.com/meowpub/meow/ld"
)

// The property that determines the class that a universal property restriction refers to.
func GetAllValuesFrom(e ld.Entity) interface{} { return e.Get(Prop_AllValuesFrom.ID) }

func SetAllValuesFrom(e ld.Entity, v interface{}) { e.Set(Prop_AllValuesFrom.ID, v) }

// The property that determines the predicate of an annotated axiom or annotated annotation.
func GetAnnotatedProperty(e ld.Entity) interface{} { return e.Get(Prop_AnnotatedProperty.ID) }

func SetAnnotatedProperty(e ld.Entity, v interface{}) { e.Set(Prop_AnnotatedProperty.ID, v) }

// The property that determines the subject of an annotated axiom or annotated annotation.
func GetAnnotatedSource(e ld.Entity) interface{} { return e.Get(Prop_AnnotatedSource.ID) }

func SetAnnotatedSource(e ld.Entity, v interface{}) { e.Set(Prop_AnnotatedSource.ID, v) }

// The property that determines the object of an annotated axiom or annotated annotation.
func GetAnnotatedTarget(e ld.Entity) interface{} { return e.Get(Prop_AnnotatedTarget.ID) }

func SetAnnotatedTarget(e ld.Entity, v interface{}) { e.Set(Prop_AnnotatedTarget.ID, v) }

// The property that determines the predicate of a negative property assertion.
func GetAssertionProperty(e ld.Entity) interface{} { return e.Get(Prop_AssertionProperty.ID) }

func SetAssertionProperty(e ld.Entity, v interface{}) { e.Set(Prop_AssertionProperty.ID, v) }

// The annotation property that indicates that a given ontology is backward compatible with another ontology.
func GetBackwardCompatibleWith(e ld.Entity) interface{} { return e.Get(Prop_BackwardCompatibleWith.ID) }

func SetBackwardCompatibleWith(e ld.Entity, v interface{}) { e.Set(Prop_BackwardCompatibleWith.ID, v) }

// The data property that does not relate any individual to any data value.
func GetBottomDataProperty(e ld.Entity) interface{} { return e.Get(Prop_BottomDataProperty.ID) }

func SetBottomDataProperty(e ld.Entity, v interface{}) { e.Set(Prop_BottomDataProperty.ID, v) }

// The object property that does not relate any two individuals.
func GetBottomObjectProperty(e ld.Entity) interface{} { return e.Get(Prop_BottomObjectProperty.ID) }

func SetBottomObjectProperty(e ld.Entity, v interface{}) { e.Set(Prop_BottomObjectProperty.ID, v) }

// The property that determines the cardinality of an exact cardinality restriction.
func GetCardinality(e ld.Entity) interface{} { return e.Get(Prop_Cardinality.ID) }

func SetCardinality(e ld.Entity, v interface{}) { e.Set(Prop_Cardinality.ID, v) }

// The property that determines that a given class is the complement of another class.
func GetComplementOf(e ld.Entity) interface{} { return e.Get(Prop_ComplementOf.ID) }

func SetComplementOf(e ld.Entity, v interface{}) { e.Set(Prop_ComplementOf.ID, v) }

// The property that determines that a given data range is the complement of another data range with respect to the data domain.
func GetDatatypeComplementOf(e ld.Entity) interface{} { return e.Get(Prop_DatatypeComplementOf.ID) }

func SetDatatypeComplementOf(e ld.Entity, v interface{}) { e.Set(Prop_DatatypeComplementOf.ID, v) }

// The annotation property that indicates that a given entity has been deprecated.
func GetDeprecated(e ld.Entity) interface{} { return e.Get(Prop_Deprecated.ID) }

func SetDeprecated(e ld.Entity, v interface{}) { e.Set(Prop_Deprecated.ID, v) }

// The property that determines that two given individuals are different.
func GetDifferentFrom(e ld.Entity) interface{} { return e.Get(Prop_DifferentFrom.ID) }

func SetDifferentFrom(e ld.Entity, v interface{}) { e.Set(Prop_DifferentFrom.ID, v) }

// The property that determines that a given class is equivalent to the disjoint union of a collection of other classes.
func GetDisjointUnionOf(e ld.Entity) interface{} { return e.Get(Prop_DisjointUnionOf.ID) }

func SetDisjointUnionOf(e ld.Entity, v interface{}) { e.Set(Prop_DisjointUnionOf.ID, v) }

// The property that determines that two given classes are disjoint.
func GetDisjointWith(e ld.Entity) interface{} { return e.Get(Prop_DisjointWith.ID) }

func SetDisjointWith(e ld.Entity, v interface{}) { e.Set(Prop_DisjointWith.ID, v) }

// The property that determines the collection of pairwise different individuals in a owl:AllDifferent axiom.
func GetDistinctMembers(e ld.Entity) interface{} { return e.Get(Prop_DistinctMembers.ID) }

func SetDistinctMembers(e ld.Entity, v interface{}) { e.Set(Prop_DistinctMembers.ID, v) }

// The property that determines that two given classes are equivalent, and that is used to specify datatype definitions.
func GetEquivalentClass(e ld.Entity) interface{} { return e.Get(Prop_EquivalentClass.ID) }

func SetEquivalentClass(e ld.Entity, v interface{}) { e.Set(Prop_EquivalentClass.ID, v) }

// The property that determines that two given properties are equivalent.
func GetEquivalentProperty(e ld.Entity) interface{} { return e.Get(Prop_EquivalentProperty.ID) }

func SetEquivalentProperty(e ld.Entity, v interface{}) { e.Set(Prop_EquivalentProperty.ID, v) }

// The property that determines the collection of properties that jointly build a key.
func GetHasKey(e ld.Entity) interface{} { return e.Get(Prop_HasKey.ID) }

func SetHasKey(e ld.Entity, v interface{}) { e.Set(Prop_HasKey.ID, v) }

// The property that determines the property that a self restriction refers to.
func GetHasSelf(e ld.Entity) interface{} { return e.Get(Prop_HasSelf.ID) }

func SetHasSelf(e ld.Entity, v interface{}) { e.Set(Prop_HasSelf.ID, v) }

// The property that determines the individual that a has-value restriction refers to.
func GetHasValue(e ld.Entity) interface{} { return e.Get(Prop_HasValue.ID) }

func SetHasValue(e ld.Entity, v interface{}) { e.Set(Prop_HasValue.ID, v) }

// The property that is used for importing other ontologies into a given ontology.
func GetImports(e ld.Entity) interface{} { return e.Get(Prop_Imports.ID) }

func SetImports(e ld.Entity, v interface{}) { e.Set(Prop_Imports.ID, v) }

// The annotation property that indicates that a given ontology is incompatible with another ontology.
func GetIncompatibleWith(e ld.Entity) interface{} { return e.Get(Prop_IncompatibleWith.ID) }

func SetIncompatibleWith(e ld.Entity, v interface{}) { e.Set(Prop_IncompatibleWith.ID, v) }

// The property that determines the collection of classes or data ranges that build an intersection.
func GetIntersectionOf(e ld.Entity) interface{} { return e.Get(Prop_IntersectionOf.ID) }

func SetIntersectionOf(e ld.Entity, v interface{}) { e.Set(Prop_IntersectionOf.ID, v) }

// The property that determines that two given properties are inverse.
func GetInverseOf(e ld.Entity) interface{} { return e.Get(Prop_InverseOf.ID) }

func SetInverseOf(e ld.Entity, v interface{}) { e.Set(Prop_InverseOf.ID, v) }

// The property that determines the cardinality of a maximum cardinality restriction.
func GetMaxCardinality(e ld.Entity) interface{} { return e.Get(Prop_MaxCardinality.ID) }

func SetMaxCardinality(e ld.Entity, v interface{}) { e.Set(Prop_MaxCardinality.ID, v) }

// The property that determines the cardinality of a maximum qualified cardinality restriction.
func GetMaxQualifiedCardinality(e ld.Entity) interface{} {
	return e.Get(Prop_MaxQualifiedCardinality.ID)
}

func SetMaxQualifiedCardinality(e ld.Entity, v interface{}) { e.Set(Prop_MaxQualifiedCardinality.ID, v) }

// The property that determines the collection of members in either a owl:AllDifferent, owl:AllDisjointClasses or owl:AllDisjointProperties axiom.
func GetMembers(e ld.Entity) interface{} { return e.Get(Prop_Members.ID) }

func SetMembers(e ld.Entity, v interface{}) { e.Set(Prop_Members.ID, v) }

// The property that determines the cardinality of a minimum cardinality restriction.
func GetMinCardinality(e ld.Entity) interface{} { return e.Get(Prop_MinCardinality.ID) }

func SetMinCardinality(e ld.Entity, v interface{}) { e.Set(Prop_MinCardinality.ID, v) }

// The property that determines the cardinality of a minimum qualified cardinality restriction.
func GetMinQualifiedCardinality(e ld.Entity) interface{} {
	return e.Get(Prop_MinQualifiedCardinality.ID)
}

func SetMinQualifiedCardinality(e ld.Entity, v interface{}) { e.Set(Prop_MinQualifiedCardinality.ID, v) }

// The property that determines the class that a qualified object cardinality restriction refers to.
func GetOnClass(e ld.Entity) interface{} { return e.Get(Prop_OnClass.ID) }

func SetOnClass(e ld.Entity, v interface{}) { e.Set(Prop_OnClass.ID, v) }

// The property that determines the data range that a qualified data cardinality restriction refers to.
func GetOnDataRange(e ld.Entity) interface{} { return e.Get(Prop_OnDataRange.ID) }

func SetOnDataRange(e ld.Entity, v interface{}) { e.Set(Prop_OnDataRange.ID, v) }

// The property that determines the datatype that a datatype restriction refers to.
func GetOnDatatype(e ld.Entity) interface{} { return e.Get(Prop_OnDatatype.ID) }

func SetOnDatatype(e ld.Entity, v interface{}) { e.Set(Prop_OnDatatype.ID, v) }

// The property that determines the n-tuple of properties that a property restriction on an n-ary data range refers to.
func GetOnProperties(e ld.Entity) interface{} { return e.Get(Prop_OnProperties.ID) }

func SetOnProperties(e ld.Entity, v interface{}) { e.Set(Prop_OnProperties.ID, v) }

// The property that determines the property that a property restriction refers to.
func GetOnProperty(e ld.Entity) interface{} { return e.Get(Prop_OnProperty.ID) }

func SetOnProperty(e ld.Entity, v interface{}) { e.Set(Prop_OnProperty.ID, v) }

// The property that determines the collection of individuals or data values that build an enumeration.
func GetOneOf(e ld.Entity) interface{} { return e.Get(Prop_OneOf.ID) }

func SetOneOf(e ld.Entity, v interface{}) { e.Set(Prop_OneOf.ID, v) }

// The annotation property that indicates the predecessor ontology of a given ontology.
func GetPriorVersion(e ld.Entity) interface{} { return e.Get(Prop_PriorVersion.ID) }

func SetPriorVersion(e ld.Entity, v interface{}) { e.Set(Prop_PriorVersion.ID, v) }

// The property that determines the n-tuple of properties that build a sub property chain of a given property.
func GetPropertyChainAxiom(e ld.Entity) interface{} { return e.Get(Prop_PropertyChainAxiom.ID) }

func SetPropertyChainAxiom(e ld.Entity, v interface{}) { e.Set(Prop_PropertyChainAxiom.ID, v) }

// The property that determines that two given properties are disjoint.
func GetPropertyDisjointWith(e ld.Entity) interface{} { return e.Get(Prop_PropertyDisjointWith.ID) }

func SetPropertyDisjointWith(e ld.Entity, v interface{}) { e.Set(Prop_PropertyDisjointWith.ID, v) }

// The property that determines the cardinality of an exact qualified cardinality restriction.
func GetQualifiedCardinality(e ld.Entity) interface{} { return e.Get(Prop_QualifiedCardinality.ID) }

func SetQualifiedCardinality(e ld.Entity, v interface{}) { e.Set(Prop_QualifiedCardinality.ID, v) }

// The property that determines that two given individuals are equal.
func GetSameAs(e ld.Entity) interface{} { return e.Get(Prop_SameAs.ID) }

func SetSameAs(e ld.Entity, v interface{}) { e.Set(Prop_SameAs.ID, v) }

// The property that determines the class that an existential property restriction refers to.
func GetSomeValuesFrom(e ld.Entity) interface{} { return e.Get(Prop_SomeValuesFrom.ID) }

func SetSomeValuesFrom(e ld.Entity, v interface{}) { e.Set(Prop_SomeValuesFrom.ID, v) }

// The property that determines the subject of a negative property assertion.
func GetSourceIndividual(e ld.Entity) interface{} { return e.Get(Prop_SourceIndividual.ID) }

func SetSourceIndividual(e ld.Entity, v interface{}) { e.Set(Prop_SourceIndividual.ID, v) }

// The property that determines the object of a negative object property assertion.
func GetTargetIndividual(e ld.Entity) interface{} { return e.Get(Prop_TargetIndividual.ID) }

func SetTargetIndividual(e ld.Entity, v interface{}) { e.Set(Prop_TargetIndividual.ID, v) }

// The property that determines the value of a negative data property assertion.
func GetTargetValue(e ld.Entity) interface{} { return e.Get(Prop_TargetValue.ID) }

func SetTargetValue(e ld.Entity, v interface{}) { e.Set(Prop_TargetValue.ID, v) }

// The data property that relates every individual to every data value.
func GetTopDataProperty(e ld.Entity) interface{} { return e.Get(Prop_TopDataProperty.ID) }

func SetTopDataProperty(e ld.Entity, v interface{}) { e.Set(Prop_TopDataProperty.ID, v) }

// The object property that relates every two individuals.
func GetTopObjectProperty(e ld.Entity) interface{} { return e.Get(Prop_TopObjectProperty.ID) }

func SetTopObjectProperty(e ld.Entity, v interface{}) { e.Set(Prop_TopObjectProperty.ID, v) }

// The property that determines the collection of classes or data ranges that build a union.
func GetUnionOf(e ld.Entity) interface{} { return e.Get(Prop_UnionOf.ID) }

func SetUnionOf(e ld.Entity, v interface{}) { e.Set(Prop_UnionOf.ID, v) }

// The property that identifies the version IRI of an ontology.
func GetVersionIRI(e ld.Entity) interface{} { return e.Get(Prop_VersionIRI.ID) }

func SetVersionIRI(e ld.Entity, v interface{}) { e.Set(Prop_VersionIRI.ID, v) }

// The annotation property that provides version information for an ontology or another OWL construct.
func GetVersionInfo(e ld.Entity) interface{} { return e.Get(Prop_VersionInfo.ID) }

func SetVersionInfo(e ld.Entity, v interface{}) { e.Set(Prop_VersionInfo.ID, v) }

// The property that determines the collection of facet-value pairs that define a datatype restriction.
func GetWithRestrictions(e ld.Entity) interface{} { return e.Get(Prop_WithRestrictions.ID) }

func SetWithRestrictions(e ld.Entity, v interface{}) { e.Set(Prop_WithRestrictions.ID, v) }
