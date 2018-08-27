// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package rdf

const (
	// The first item in the subject RDF list.
	PropFirst = "http://www.w3.org/1999/02/22-rdf-syntax-ns#first"

	// The object of the subject RDF statement.
	PropObject = "http://www.w3.org/1999/02/22-rdf-syntax-ns#object"

	// The predicate of the subject RDF statement.
	PropPredicate = "http://www.w3.org/1999/02/22-rdf-syntax-ns#predicate"

	// The rest of the subject RDF list after the first item.
	PropRest = "http://www.w3.org/1999/02/22-rdf-syntax-ns#rest"

	// The subject of the subject RDF statement.
	PropSubject = "http://www.w3.org/1999/02/22-rdf-syntax-ns#subject"

	// The subject is an instance of a class.
	PropType = "http://www.w3.org/1999/02/22-rdf-syntax-ns#type"

	// Idiomatic property used for structured values.
	PropValue = "http://www.w3.org/1999/02/22-rdf-syntax-ns#value"

	// A description of the subject resource.
	PropComment = "http://www.w3.org/2000/01/rdf-schema#comment"

	// A domain of the subject property.
	PropDomain = "http://www.w3.org/2000/01/rdf-schema#domain"

	// The defininition of the subject resource.
	PropIsDefinedBy = "http://www.w3.org/2000/01/rdf-schema#isDefinedBy"

	// A human-readable name for the subject.
	PropLabel = "http://www.w3.org/2000/01/rdf-schema#label"

	// A member of the subject resource.
	PropMember = "http://www.w3.org/2000/01/rdf-schema#member"

	// A range of the subject property.
	PropRange = "http://www.w3.org/2000/01/rdf-schema#range"

	// Further information about the subject resource.
	PropSeeAlso = "http://www.w3.org/2000/01/rdf-schema#seeAlso"

	// The subject is a subclass of a class.
	PropSubClassOf = "http://www.w3.org/2000/01/rdf-schema#subClassOf"

	// The subject is a subproperty of a property.
	PropSubPropertyOf = "http://www.w3.org/2000/01/rdf-schema#subPropertyOf"
)

const (
	// The class of containers of alternatives.
	TypeAlt = "http://www.w3.org/1999/02/22-rdf-syntax-ns#Alt"

	// The class of unordered containers.
	TypeBag = "http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag"

	// The class of RDF Lists.
	TypeList = "http://www.w3.org/1999/02/22-rdf-syntax-ns#List"

	// The class of RDF properties.
	TypeProperty = "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property"

	// The class of ordered containers.
	TypeSeq = "http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq"

	// The class of RDF statements.
	TypeStatement = "http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement"

	// The class of classes.
	TypeClass = "http://www.w3.org/2000/01/rdf-schema#Class"

	// The class of RDF containers.
	TypeContainer = "http://www.w3.org/2000/01/rdf-schema#Container"

	// The class of container membership properties, rdf:_1, rdf:_2, ...,
	// all of which are sub-properties of 'member'.
	TypeContainerMembershipProperty = "http://www.w3.org/2000/01/rdf-schema#ContainerMembershipProperty"

	// The class of RDF datatypes.
	TypeDatatype = "http://www.w3.org/2000/01/rdf-schema#Datatype"

	// The class of literal values, eg. textual strings and integers.
	TypeLiteral = "http://www.w3.org/2000/01/rdf-schema#Literal"

	// The class resource, everything.
	TypeResource = "http://www.w3.org/2000/01/rdf-schema#Resource"
)

// The empty list, with no items in it. If the rest of a list is nil then the list has no more items in it.
// http://www.w3.org/1999/02/22-rdf-syntax-ns#nil - [http://www.w3.org/1999/02/22-rdf-syntax-ns#List http://www.w3.org/ns/activitystreams#OrderedItems]
