// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package owl

import (
	"github.com/meowpub/meow/ld"
)

// The property that determines the class that a universal property restriction refers to.
func GetAllValuesFrom(e ld.Entity) interface{} { return e.Get(PropAllValuesFrom) }

func SetAllValuesFrom(e ld.Entity, v interface{}) { e.Set(PropAllValuesFrom, v) }

// The property that determines the predicate of an annotated axiom or annotated annotation.
func GetAnnotatedProperty(e ld.Entity) interface{} { return e.Get(PropAnnotatedProperty) }

func SetAnnotatedProperty(e ld.Entity, v interface{}) { e.Set(PropAnnotatedProperty, v) }

// The property that determines the subject of an annotated axiom or annotated annotation.
func GetAnnotatedSource(e ld.Entity) interface{} { return e.Get(PropAnnotatedSource) }

func SetAnnotatedSource(e ld.Entity, v interface{}) { e.Set(PropAnnotatedSource, v) }

// The property that determines the object of an annotated axiom or annotated annotation.
func GetAnnotatedTarget(e ld.Entity) interface{} { return e.Get(PropAnnotatedTarget) }

func SetAnnotatedTarget(e ld.Entity, v interface{}) { e.Set(PropAnnotatedTarget, v) }

// The property that determines the predicate of a negative property assertion.
func GetAssertionProperty(e ld.Entity) interface{} { return e.Get(PropAssertionProperty) }

func SetAssertionProperty(e ld.Entity, v interface{}) { e.Set(PropAssertionProperty, v) }

// The annotation property that indicates that a given ontology is backward compatible with another ontology.
func GetBackwardCompatibleWith(e ld.Entity) interface{} { return e.Get(PropBackwardCompatibleWith) }

func SetBackwardCompatibleWith(e ld.Entity, v interface{}) { e.Set(PropBackwardCompatibleWith, v) }

// The data property that does not relate any individual to any data value.
func GetBottomDataProperty(e ld.Entity) interface{} { return e.Get(PropBottomDataProperty) }

func SetBottomDataProperty(e ld.Entity, v interface{}) { e.Set(PropBottomDataProperty, v) }

// The object property that does not relate any two individuals.
func GetBottomObjectProperty(e ld.Entity) interface{} { return e.Get(PropBottomObjectProperty) }

func SetBottomObjectProperty(e ld.Entity, v interface{}) { e.Set(PropBottomObjectProperty, v) }

// The property that determines the cardinality of an exact cardinality restriction.
func GetCardinality(e ld.Entity) interface{} { return e.Get(PropCardinality) }

func SetCardinality(e ld.Entity, v interface{}) { e.Set(PropCardinality, v) }

// The property that determines that a given class is the complement of another class.
func GetComplementOf(e ld.Entity) interface{} { return e.Get(PropComplementOf) }

func SetComplementOf(e ld.Entity, v interface{}) { e.Set(PropComplementOf, v) }

// The property that determines that a given data range is the complement of another data range with respect to the data domain.
func GetDatatypeComplementOf(e ld.Entity) interface{} { return e.Get(PropDatatypeComplementOf) }

func SetDatatypeComplementOf(e ld.Entity, v interface{}) { e.Set(PropDatatypeComplementOf, v) }

// The annotation property that indicates that a given entity has been deprecated.
func GetDeprecated(e ld.Entity) interface{} { return e.Get(PropDeprecated) }

func SetDeprecated(e ld.Entity, v interface{}) { e.Set(PropDeprecated, v) }

// The property that determines that two given individuals are different.
func GetDifferentFrom(e ld.Entity) interface{} { return e.Get(PropDifferentFrom) }

func SetDifferentFrom(e ld.Entity, v interface{}) { e.Set(PropDifferentFrom, v) }

// The property that determines that a given class is equivalent to the disjoint union of a collection of other classes.
func GetDisjointUnionOf(e ld.Entity) interface{} { return e.Get(PropDisjointUnionOf) }

func SetDisjointUnionOf(e ld.Entity, v interface{}) { e.Set(PropDisjointUnionOf, v) }

// The property that determines that two given classes are disjoint.
func GetDisjointWith(e ld.Entity) interface{} { return e.Get(PropDisjointWith) }

func SetDisjointWith(e ld.Entity, v interface{}) { e.Set(PropDisjointWith, v) }

// The property that determines the collection of pairwise different individuals in a owl:AllDifferent axiom.
func GetDistinctMembers(e ld.Entity) interface{} { return e.Get(PropDistinctMembers) }

func SetDistinctMembers(e ld.Entity, v interface{}) { e.Set(PropDistinctMembers, v) }

// The property that determines that two given classes are equivalent, and that is used to specify datatype definitions.
func GetEquivalentClass(e ld.Entity) interface{} { return e.Get(PropEquivalentClass) }

func SetEquivalentClass(e ld.Entity, v interface{}) { e.Set(PropEquivalentClass, v) }

// The property that determines that two given properties are equivalent.
func GetEquivalentProperty(e ld.Entity) interface{} { return e.Get(PropEquivalentProperty) }

func SetEquivalentProperty(e ld.Entity, v interface{}) { e.Set(PropEquivalentProperty, v) }

// The property that determines the collection of properties that jointly build a key.
func GetHasKey(e ld.Entity) interface{} { return e.Get(PropHasKey) }

func SetHasKey(e ld.Entity, v interface{}) { e.Set(PropHasKey, v) }

// The property that determines the property that a self restriction refers to.
func GetHasSelf(e ld.Entity) interface{} { return e.Get(PropHasSelf) }

func SetHasSelf(e ld.Entity, v interface{}) { e.Set(PropHasSelf, v) }

// The property that determines the individual that a has-value restriction refers to.
func GetHasValue(e ld.Entity) interface{} { return e.Get(PropHasValue) }

func SetHasValue(e ld.Entity, v interface{}) { e.Set(PropHasValue, v) }

// The property that is used for importing other ontologies into a given ontology.
func GetImports(e ld.Entity) interface{} { return e.Get(PropImports) }

func SetImports(e ld.Entity, v interface{}) { e.Set(PropImports, v) }

// The annotation property that indicates that a given ontology is incompatible with another ontology.
func GetIncompatibleWith(e ld.Entity) interface{} { return e.Get(PropIncompatibleWith) }

func SetIncompatibleWith(e ld.Entity, v interface{}) { e.Set(PropIncompatibleWith, v) }

// The property that determines the collection of classes or data ranges that build an intersection.
func GetIntersectionOf(e ld.Entity) interface{} { return e.Get(PropIntersectionOf) }

func SetIntersectionOf(e ld.Entity, v interface{}) { e.Set(PropIntersectionOf, v) }

// The property that determines that two given properties are inverse.
func GetInverseOf(e ld.Entity) interface{} { return e.Get(PropInverseOf) }

func SetInverseOf(e ld.Entity, v interface{}) { e.Set(PropInverseOf, v) }

// The property that determines the cardinality of a maximum cardinality restriction.
func GetMaxCardinality(e ld.Entity) interface{} { return e.Get(PropMaxCardinality) }

func SetMaxCardinality(e ld.Entity, v interface{}) { e.Set(PropMaxCardinality, v) }

// The property that determines the cardinality of a maximum qualified cardinality restriction.
func GetMaxQualifiedCardinality(e ld.Entity) interface{} { return e.Get(PropMaxQualifiedCardinality) }

func SetMaxQualifiedCardinality(e ld.Entity, v interface{}) { e.Set(PropMaxQualifiedCardinality, v) }

// The property that determines the collection of members in either a owl:AllDifferent, owl:AllDisjointClasses or owl:AllDisjointProperties axiom.
func GetMembers(e ld.Entity) interface{} { return e.Get(PropMembers) }

func SetMembers(e ld.Entity, v interface{}) { e.Set(PropMembers, v) }

// The property that determines the cardinality of a minimum cardinality restriction.
func GetMinCardinality(e ld.Entity) interface{} { return e.Get(PropMinCardinality) }

func SetMinCardinality(e ld.Entity, v interface{}) { e.Set(PropMinCardinality, v) }

// The property that determines the cardinality of a minimum qualified cardinality restriction.
func GetMinQualifiedCardinality(e ld.Entity) interface{} { return e.Get(PropMinQualifiedCardinality) }

func SetMinQualifiedCardinality(e ld.Entity, v interface{}) { e.Set(PropMinQualifiedCardinality, v) }

// The property that determines the class that a qualified object cardinality restriction refers to.
func GetOnClass(e ld.Entity) interface{} { return e.Get(PropOnClass) }

func SetOnClass(e ld.Entity, v interface{}) { e.Set(PropOnClass, v) }

// The property that determines the data range that a qualified data cardinality restriction refers to.
func GetOnDataRange(e ld.Entity) interface{} { return e.Get(PropOnDataRange) }

func SetOnDataRange(e ld.Entity, v interface{}) { e.Set(PropOnDataRange, v) }

// The property that determines the datatype that a datatype restriction refers to.
func GetOnDatatype(e ld.Entity) interface{} { return e.Get(PropOnDatatype) }

func SetOnDatatype(e ld.Entity, v interface{}) { e.Set(PropOnDatatype, v) }

// The property that determines the n-tuple of properties that a property restriction on an n-ary data range refers to.
func GetOnProperties(e ld.Entity) interface{} { return e.Get(PropOnProperties) }

func SetOnProperties(e ld.Entity, v interface{}) { e.Set(PropOnProperties, v) }

// The property that determines the property that a property restriction refers to.
func GetOnProperty(e ld.Entity) interface{} { return e.Get(PropOnProperty) }

func SetOnProperty(e ld.Entity, v interface{}) { e.Set(PropOnProperty, v) }

// The property that determines the collection of individuals or data values that build an enumeration.
func GetOneOf(e ld.Entity) interface{} { return e.Get(PropOneOf) }

func SetOneOf(e ld.Entity, v interface{}) { e.Set(PropOneOf, v) }

// The annotation property that indicates the predecessor ontology of a given ontology.
func GetPriorVersion(e ld.Entity) interface{} { return e.Get(PropPriorVersion) }

func SetPriorVersion(e ld.Entity, v interface{}) { e.Set(PropPriorVersion, v) }

// The property that determines the n-tuple of properties that build a sub property chain of a given property.
func GetPropertyChainAxiom(e ld.Entity) interface{} { return e.Get(PropPropertyChainAxiom) }

func SetPropertyChainAxiom(e ld.Entity, v interface{}) { e.Set(PropPropertyChainAxiom, v) }

// The property that determines that two given properties are disjoint.
func GetPropertyDisjointWith(e ld.Entity) interface{} { return e.Get(PropPropertyDisjointWith) }

func SetPropertyDisjointWith(e ld.Entity, v interface{}) { e.Set(PropPropertyDisjointWith, v) }

// The property that determines the cardinality of an exact qualified cardinality restriction.
func GetQualifiedCardinality(e ld.Entity) interface{} { return e.Get(PropQualifiedCardinality) }

func SetQualifiedCardinality(e ld.Entity, v interface{}) { e.Set(PropQualifiedCardinality, v) }

// The property that determines that two given individuals are equal.
func GetSameAs(e ld.Entity) interface{} { return e.Get(PropSameAs) }

func SetSameAs(e ld.Entity, v interface{}) { e.Set(PropSameAs, v) }

// The property that determines the class that an existential property restriction refers to.
func GetSomeValuesFrom(e ld.Entity) interface{} { return e.Get(PropSomeValuesFrom) }

func SetSomeValuesFrom(e ld.Entity, v interface{}) { e.Set(PropSomeValuesFrom, v) }

// The property that determines the subject of a negative property assertion.
func GetSourceIndividual(e ld.Entity) interface{} { return e.Get(PropSourceIndividual) }

func SetSourceIndividual(e ld.Entity, v interface{}) { e.Set(PropSourceIndividual, v) }

// The property that determines the object of a negative object property assertion.
func GetTargetIndividual(e ld.Entity) interface{} { return e.Get(PropTargetIndividual) }

func SetTargetIndividual(e ld.Entity, v interface{}) { e.Set(PropTargetIndividual, v) }

// The property that determines the value of a negative data property assertion.
func GetTargetValue(e ld.Entity) interface{} { return e.Get(PropTargetValue) }

func SetTargetValue(e ld.Entity, v interface{}) { e.Set(PropTargetValue, v) }

// The data property that relates every individual to every data value.
func GetTopDataProperty(e ld.Entity) interface{} { return e.Get(PropTopDataProperty) }

func SetTopDataProperty(e ld.Entity, v interface{}) { e.Set(PropTopDataProperty, v) }

// The object property that relates every two individuals.
func GetTopObjectProperty(e ld.Entity) interface{} { return e.Get(PropTopObjectProperty) }

func SetTopObjectProperty(e ld.Entity, v interface{}) { e.Set(PropTopObjectProperty, v) }

// The property that determines the collection of classes or data ranges that build a union.
func GetUnionOf(e ld.Entity) interface{} { return e.Get(PropUnionOf) }

func SetUnionOf(e ld.Entity, v interface{}) { e.Set(PropUnionOf, v) }

// The property that identifies the version IRI of an ontology.
func GetVersionIRI(e ld.Entity) interface{} { return e.Get(PropVersionIRI) }

func SetVersionIRI(e ld.Entity, v interface{}) { e.Set(PropVersionIRI, v) }

// The annotation property that provides version information for an ontology or another OWL construct.
func GetVersionInfo(e ld.Entity) interface{} { return e.Get(PropVersionInfo) }

func SetVersionInfo(e ld.Entity, v interface{}) { e.Set(PropVersionInfo, v) }

// The property that determines the collection of facet-value pairs that define a datatype restriction.
func GetWithRestrictions(e ld.Entity) interface{} { return e.Get(PropWithRestrictions) }

func SetWithRestrictions(e ld.Entity, v interface{}) { e.Set(PropWithRestrictions, v) }
