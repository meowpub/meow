// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package owl

import (
	"github.com/meowpub/meow/ld"
)

// The property that determines the class that a universal property restriction refers to.
func GetAllValuesFrom(e ld.Entity) interface{} { return e.Get(PropAllValuesFrom) }

// The property that determines the predicate of an annotated axiom or annotated annotation.
func GetAnnotatedProperty(e ld.Entity) interface{} { return e.Get(PropAnnotatedProperty) }

// The property that determines the subject of an annotated axiom or annotated annotation.
func GetAnnotatedSource(e ld.Entity) interface{} { return e.Get(PropAnnotatedSource) }

// The property that determines the object of an annotated axiom or annotated annotation.
func GetAnnotatedTarget(e ld.Entity) interface{} { return e.Get(PropAnnotatedTarget) }

// The property that determines the predicate of a negative property assertion.
func GetAssertionProperty(e ld.Entity) interface{} { return e.Get(PropAssertionProperty) }

// The annotation property that indicates that a given ontology is backward compatible with another ontology.
func GetBackwardCompatibleWith(e ld.Entity) interface{} { return e.Get(PropBackwardCompatibleWith) }

// The data property that does not relate any individual to any data value.
func GetBottomDataProperty(e ld.Entity) interface{} { return e.Get(PropBottomDataProperty) }

// The object property that does not relate any two individuals.
func GetBottomObjectProperty(e ld.Entity) interface{} { return e.Get(PropBottomObjectProperty) }

// The property that determines the cardinality of an exact cardinality restriction.
func GetCardinality(e ld.Entity) interface{} { return e.Get(PropCardinality) }

// The property that determines that a given class is the complement of another class.
func GetComplementOf(e ld.Entity) interface{} { return e.Get(PropComplementOf) }

// The property that determines that a given data range is the complement of another data range with respect to the data domain.
func GetDatatypeComplementOf(e ld.Entity) interface{} { return e.Get(PropDatatypeComplementOf) }

// The annotation property that indicates that a given entity has been deprecated.
func GetDeprecated(e ld.Entity) interface{} { return e.Get(PropDeprecated) }

// The property that determines that two given individuals are different.
func GetDifferentFrom(e ld.Entity) interface{} { return e.Get(PropDifferentFrom) }

// The property that determines that a given class is equivalent to the disjoint union of a collection of other classes.
func GetDisjointUnionOf(e ld.Entity) interface{} { return e.Get(PropDisjointUnionOf) }

// The property that determines that two given classes are disjoint.
func GetDisjointWith(e ld.Entity) interface{} { return e.Get(PropDisjointWith) }

// The property that determines the collection of pairwise different individuals in a owl:AllDifferent axiom.
func GetDistinctMembers(e ld.Entity) interface{} { return e.Get(PropDistinctMembers) }

// The property that determines that two given classes are equivalent, and that is used to specify datatype definitions.
func GetEquivalentClass(e ld.Entity) interface{} { return e.Get(PropEquivalentClass) }

// The property that determines that two given properties are equivalent.
func GetEquivalentProperty(e ld.Entity) interface{} { return e.Get(PropEquivalentProperty) }

// The property that determines the collection of properties that jointly build a key.
func GetHasKey(e ld.Entity) interface{} { return e.Get(PropHasKey) }

// The property that determines the property that a self restriction refers to.
func GetHasSelf(e ld.Entity) interface{} { return e.Get(PropHasSelf) }

// The property that determines the individual that a has-value restriction refers to.
func GetHasValue(e ld.Entity) interface{} { return e.Get(PropHasValue) }

// The property that is used for importing other ontologies into a given ontology.
func GetImports(e ld.Entity) interface{} { return e.Get(PropImports) }

// The annotation property that indicates that a given ontology is incompatible with another ontology.
func GetIncompatibleWith(e ld.Entity) interface{} { return e.Get(PropIncompatibleWith) }

// The property that determines the collection of classes or data ranges that build an intersection.
func GetIntersectionOf(e ld.Entity) interface{} { return e.Get(PropIntersectionOf) }

// The property that determines that two given properties are inverse.
func GetInverseOf(e ld.Entity) interface{} { return e.Get(PropInverseOf) }

// The property that determines the cardinality of a maximum cardinality restriction.
func GetMaxCardinality(e ld.Entity) interface{} { return e.Get(PropMaxCardinality) }

// The property that determines the cardinality of a maximum qualified cardinality restriction.
func GetMaxQualifiedCardinality(e ld.Entity) interface{} { return e.Get(PropMaxQualifiedCardinality) }

// The property that determines the collection of members in either a owl:AllDifferent, owl:AllDisjointClasses or owl:AllDisjointProperties axiom.
func GetMembers(e ld.Entity) interface{} { return e.Get(PropMembers) }

// The property that determines the cardinality of a minimum cardinality restriction.
func GetMinCardinality(e ld.Entity) interface{} { return e.Get(PropMinCardinality) }

// The property that determines the cardinality of a minimum qualified cardinality restriction.
func GetMinQualifiedCardinality(e ld.Entity) interface{} { return e.Get(PropMinQualifiedCardinality) }

// The property that determines the class that a qualified object cardinality restriction refers to.
func GetOnClass(e ld.Entity) interface{} { return e.Get(PropOnClass) }

// The property that determines the data range that a qualified data cardinality restriction refers to.
func GetOnDataRange(e ld.Entity) interface{} { return e.Get(PropOnDataRange) }

// The property that determines the datatype that a datatype restriction refers to.
func GetOnDatatype(e ld.Entity) interface{} { return e.Get(PropOnDatatype) }

// The property that determines the n-tuple of properties that a property restriction on an n-ary data range refers to.
func GetOnProperties(e ld.Entity) interface{} { return e.Get(PropOnProperties) }

// The property that determines the property that a property restriction refers to.
func GetOnProperty(e ld.Entity) interface{} { return e.Get(PropOnProperty) }

// The property that determines the collection of individuals or data values that build an enumeration.
func GetOneOf(e ld.Entity) interface{} { return e.Get(PropOneOf) }

// The annotation property that indicates the predecessor ontology of a given ontology.
func GetPriorVersion(e ld.Entity) interface{} { return e.Get(PropPriorVersion) }

// The property that determines the n-tuple of properties that build a sub property chain of a given property.
func GetPropertyChainAxiom(e ld.Entity) interface{} { return e.Get(PropPropertyChainAxiom) }

// The property that determines that two given properties are disjoint.
func GetPropertyDisjointWith(e ld.Entity) interface{} { return e.Get(PropPropertyDisjointWith) }

// The property that determines the cardinality of an exact qualified cardinality restriction.
func GetQualifiedCardinality(e ld.Entity) interface{} { return e.Get(PropQualifiedCardinality) }

// The property that determines that two given individuals are equal.
func GetSameAs(e ld.Entity) interface{} { return e.Get(PropSameAs) }

// The property that determines the class that an existential property restriction refers to.
func GetSomeValuesFrom(e ld.Entity) interface{} { return e.Get(PropSomeValuesFrom) }

// The property that determines the subject of a negative property assertion.
func GetSourceIndividual(e ld.Entity) interface{} { return e.Get(PropSourceIndividual) }

// The property that determines the object of a negative object property assertion.
func GetTargetIndividual(e ld.Entity) interface{} { return e.Get(PropTargetIndividual) }

// The property that determines the value of a negative data property assertion.
func GetTargetValue(e ld.Entity) interface{} { return e.Get(PropTargetValue) }

// The data property that relates every individual to every data value.
func GetTopDataProperty(e ld.Entity) interface{} { return e.Get(PropTopDataProperty) }

// The object property that relates every two individuals.
func GetTopObjectProperty(e ld.Entity) interface{} { return e.Get(PropTopObjectProperty) }

// The property that determines the collection of classes or data ranges that build a union.
func GetUnionOf(e ld.Entity) interface{} { return e.Get(PropUnionOf) }

// The property that identifies the version IRI of an ontology.
func GetVersionIRI(e ld.Entity) interface{} { return e.Get(PropVersionIRI) }

// The annotation property that provides version information for an ontology or another OWL construct.
func GetVersionInfo(e ld.Entity) interface{} { return e.Get(PropVersionInfo) }

// The property that determines the collection of facet-value pairs that define a datatype restriction.
func GetWithRestrictions(e ld.Entity) interface{} { return e.Get(PropWithRestrictions) }
