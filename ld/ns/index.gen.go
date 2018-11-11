// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package ns

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/as"
	"github.com/meowpub/meow/ld/ns/ldp"
	"github.com/meowpub/meow/ld/ns/owl"
	"github.com/meowpub/meow/ld/ns/rdf"
	"github.com/meowpub/meow/ld/ns/sec"
)

var Namespaces = map[string]*Namespace{
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#": RDF,
	"http://www.w3.org/2000/01/rdf-schema#":       RDFS,
	"http://www.w3.org/2002/07/owl#":              OWL,
	"http://www.w3.org/ns/activitystreams#":       AS,
	"http://www.w3.org/ns/ldp#":                   LDP,
	"https://w3id.org/security#":                  SEC,
}

var Types = map[string]*Type{
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#Alt":                   http_www_w3_org_1999_02_22_rdf_syntax_ns_Alt,
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag":                   http_www_w3_org_1999_02_22_rdf_syntax_ns_Bag,
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#List":                  http_www_w3_org_1999_02_22_rdf_syntax_ns_List,
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#Property":              http_www_w3_org_1999_02_22_rdf_syntax_ns_Property,
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq":                   http_www_w3_org_1999_02_22_rdf_syntax_ns_Seq,
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement":             http_www_w3_org_1999_02_22_rdf_syntax_ns_Statement,
	"http://www.w3.org/2000/01/rdf-schema#Class":                       http_www_w3_org_2000_01_rdf_schema_Class,
	"http://www.w3.org/2000/01/rdf-schema#Container":                   http_www_w3_org_2000_01_rdf_schema_Container,
	"http://www.w3.org/2000/01/rdf-schema#ContainerMembershipProperty": http_www_w3_org_2000_01_rdf_schema_ContainerMembershipProperty,
	"http://www.w3.org/2000/01/rdf-schema#Datatype":                    http_www_w3_org_2000_01_rdf_schema_Datatype,
	"http://www.w3.org/2000/01/rdf-schema#Literal":                     http_www_w3_org_2000_01_rdf_schema_Literal,
	"http://www.w3.org/2000/01/rdf-schema#Resource":                    http_www_w3_org_2000_01_rdf_schema_Resource,
	"http://www.w3.org/2002/07/owl#AllDifferent":                       http_www_w3_org_2002_07_owl_AllDifferent,
	"http://www.w3.org/2002/07/owl#AllDisjointClasses":                 http_www_w3_org_2002_07_owl_AllDisjointClasses,
	"http://www.w3.org/2002/07/owl#AllDisjointProperties":              http_www_w3_org_2002_07_owl_AllDisjointProperties,
	"http://www.w3.org/2002/07/owl#Annotation":                         http_www_w3_org_2002_07_owl_Annotation,
	"http://www.w3.org/2002/07/owl#AnnotationProperty":                 http_www_w3_org_2002_07_owl_AnnotationProperty,
	"http://www.w3.org/2002/07/owl#AsymmetricProperty":                 http_www_w3_org_2002_07_owl_AsymmetricProperty,
	"http://www.w3.org/2002/07/owl#Axiom":                              http_www_w3_org_2002_07_owl_Axiom,
	"http://www.w3.org/2002/07/owl#Class":                              http_www_w3_org_2002_07_owl_Class,
	"http://www.w3.org/2002/07/owl#DataRange":                          http_www_w3_org_2002_07_owl_DataRange,
	"http://www.w3.org/2002/07/owl#DatatypeProperty":                   http_www_w3_org_2002_07_owl_DatatypeProperty,
	"http://www.w3.org/2002/07/owl#DeprecatedClass":                    http_www_w3_org_2002_07_owl_DeprecatedClass,
	"http://www.w3.org/2002/07/owl#DeprecatedProperty":                 http_www_w3_org_2002_07_owl_DeprecatedProperty,
	"http://www.w3.org/2002/07/owl#FunctionalProperty":                 http_www_w3_org_2002_07_owl_FunctionalProperty,
	"http://www.w3.org/2002/07/owl#InverseFunctionalProperty":          http_www_w3_org_2002_07_owl_InverseFunctionalProperty,
	"http://www.w3.org/2002/07/owl#IrreflexiveProperty":                http_www_w3_org_2002_07_owl_IrreflexiveProperty,
	"http://www.w3.org/2002/07/owl#NamedIndividual":                    http_www_w3_org_2002_07_owl_NamedIndividual,
	"http://www.w3.org/2002/07/owl#NegativePropertyAssertion":          http_www_w3_org_2002_07_owl_NegativePropertyAssertion,
	"http://www.w3.org/2002/07/owl#Nothing":                            http_www_w3_org_2002_07_owl_Nothing,
	"http://www.w3.org/2002/07/owl#ObjectProperty":                     http_www_w3_org_2002_07_owl_ObjectProperty,
	"http://www.w3.org/2002/07/owl#Ontology":                           http_www_w3_org_2002_07_owl_Ontology,
	"http://www.w3.org/2002/07/owl#OntologyProperty":                   http_www_w3_org_2002_07_owl_OntologyProperty,
	"http://www.w3.org/2002/07/owl#ReflexiveProperty":                  http_www_w3_org_2002_07_owl_ReflexiveProperty,
	"http://www.w3.org/2002/07/owl#Restriction":                        http_www_w3_org_2002_07_owl_Restriction,
	"http://www.w3.org/2002/07/owl#SymmetricProperty":                  http_www_w3_org_2002_07_owl_SymmetricProperty,
	"http://www.w3.org/2002/07/owl#Thing":                              http_www_w3_org_2002_07_owl_Thing,
	"http://www.w3.org/2002/07/owl#TransitiveProperty":                 http_www_w3_org_2002_07_owl_TransitiveProperty,
	"http://www.w3.org/ns/activitystreams#Accept":                      http_www_w3_org_ns_activitystreams_Accept,
	"http://www.w3.org/ns/activitystreams#Activity":                    http_www_w3_org_ns_activitystreams_Activity,
	"http://www.w3.org/ns/activitystreams#Add":                         http_www_w3_org_ns_activitystreams_Add,
	"http://www.w3.org/ns/activitystreams#Announce":                    http_www_w3_org_ns_activitystreams_Announce,
	"http://www.w3.org/ns/activitystreams#Application":                 http_www_w3_org_ns_activitystreams_Application,
	"http://www.w3.org/ns/activitystreams#Arrive":                      http_www_w3_org_ns_activitystreams_Arrive,
	"http://www.w3.org/ns/activitystreams#Article":                     http_www_w3_org_ns_activitystreams_Article,
	"http://www.w3.org/ns/activitystreams#Audio":                       http_www_w3_org_ns_activitystreams_Audio,
	"http://www.w3.org/ns/activitystreams#Block":                       http_www_w3_org_ns_activitystreams_Block,
	"http://www.w3.org/ns/activitystreams#Collection":                  http_www_w3_org_ns_activitystreams_Collection,
	"http://www.w3.org/ns/activitystreams#CollectionPage":              http_www_w3_org_ns_activitystreams_CollectionPage,
	"http://www.w3.org/ns/activitystreams#Create":                      http_www_w3_org_ns_activitystreams_Create,
	"http://www.w3.org/ns/activitystreams#Delete":                      http_www_w3_org_ns_activitystreams_Delete,
	"http://www.w3.org/ns/activitystreams#Dislike":                     http_www_w3_org_ns_activitystreams_Dislike,
	"http://www.w3.org/ns/activitystreams#Document":                    http_www_w3_org_ns_activitystreams_Document,
	"http://www.w3.org/ns/activitystreams#Event":                       http_www_w3_org_ns_activitystreams_Event,
	"http://www.w3.org/ns/activitystreams#Flag":                        http_www_w3_org_ns_activitystreams_Flag,
	"http://www.w3.org/ns/activitystreams#Follow":                      http_www_w3_org_ns_activitystreams_Follow,
	"http://www.w3.org/ns/activitystreams#Group":                       http_www_w3_org_ns_activitystreams_Group,
	"http://www.w3.org/ns/activitystreams#Ignore":                      http_www_w3_org_ns_activitystreams_Ignore,
	"http://www.w3.org/ns/activitystreams#Image":                       http_www_w3_org_ns_activitystreams_Image,
	"http://www.w3.org/ns/activitystreams#IntransitiveActivity":        http_www_w3_org_ns_activitystreams_IntransitiveActivity,
	"http://www.w3.org/ns/activitystreams#Invite":                      http_www_w3_org_ns_activitystreams_Invite,
	"http://www.w3.org/ns/activitystreams#Join":                        http_www_w3_org_ns_activitystreams_Join,
	"http://www.w3.org/ns/activitystreams#Leave":                       http_www_w3_org_ns_activitystreams_Leave,
	"http://www.w3.org/ns/activitystreams#Like":                        http_www_w3_org_ns_activitystreams_Like,
	"http://www.w3.org/ns/activitystreams#Link":                        http_www_w3_org_ns_activitystreams_Link,
	"http://www.w3.org/ns/activitystreams#Listen":                      http_www_w3_org_ns_activitystreams_Listen,
	"http://www.w3.org/ns/activitystreams#Mention":                     http_www_w3_org_ns_activitystreams_Mention,
	"http://www.w3.org/ns/activitystreams#Move":                        http_www_w3_org_ns_activitystreams_Move,
	"http://www.w3.org/ns/activitystreams#Note":                        http_www_w3_org_ns_activitystreams_Note,
	"http://www.w3.org/ns/activitystreams#Object":                      http_www_w3_org_ns_activitystreams_Object,
	"http://www.w3.org/ns/activitystreams#Offer":                       http_www_w3_org_ns_activitystreams_Offer,
	"http://www.w3.org/ns/activitystreams#OrderedCollection":           http_www_w3_org_ns_activitystreams_OrderedCollection,
	"http://www.w3.org/ns/activitystreams#OrderedCollectionPage":       http_www_w3_org_ns_activitystreams_OrderedCollectionPage,
	"http://www.w3.org/ns/activitystreams#OrderedItems":                http_www_w3_org_ns_activitystreams_OrderedItems,
	"http://www.w3.org/ns/activitystreams#Organization":                http_www_w3_org_ns_activitystreams_Organization,
	"http://www.w3.org/ns/activitystreams#Page":                        http_www_w3_org_ns_activitystreams_Page,
	"http://www.w3.org/ns/activitystreams#Person":                      http_www_w3_org_ns_activitystreams_Person,
	"http://www.w3.org/ns/activitystreams#Place":                       http_www_w3_org_ns_activitystreams_Place,
	"http://www.w3.org/ns/activitystreams#Profile":                     http_www_w3_org_ns_activitystreams_Profile,
	"http://www.w3.org/ns/activitystreams#Question":                    http_www_w3_org_ns_activitystreams_Question,
	"http://www.w3.org/ns/activitystreams#Read":                        http_www_w3_org_ns_activitystreams_Read,
	"http://www.w3.org/ns/activitystreams#Reject":                      http_www_w3_org_ns_activitystreams_Reject,
	"http://www.w3.org/ns/activitystreams#Relationship":                http_www_w3_org_ns_activitystreams_Relationship,
	"http://www.w3.org/ns/activitystreams#Remove":                      http_www_w3_org_ns_activitystreams_Remove,
	"http://www.w3.org/ns/activitystreams#Service":                     http_www_w3_org_ns_activitystreams_Service,
	"http://www.w3.org/ns/activitystreams#TentativeAccept":             http_www_w3_org_ns_activitystreams_TentativeAccept,
	"http://www.w3.org/ns/activitystreams#TentativeReject":             http_www_w3_org_ns_activitystreams_TentativeReject,
	"http://www.w3.org/ns/activitystreams#Tombstone":                   http_www_w3_org_ns_activitystreams_Tombstone,
	"http://www.w3.org/ns/activitystreams#Travel":                      http_www_w3_org_ns_activitystreams_Travel,
	"http://www.w3.org/ns/activitystreams#Undo":                        http_www_w3_org_ns_activitystreams_Undo,
	"http://www.w3.org/ns/activitystreams#Update":                      http_www_w3_org_ns_activitystreams_Update,
	"http://www.w3.org/ns/activitystreams#Video":                       http_www_w3_org_ns_activitystreams_Video,
	"http://www.w3.org/ns/activitystreams#View":                        http_www_w3_org_ns_activitystreams_View,
	"http://www.w3.org/ns/ldp#BasicContainer":                          http_www_w3_org_ns_ldp_BasicContainer,
	"http://www.w3.org/ns/ldp#Container":                               http_www_w3_org_ns_ldp_Container,
	"http://www.w3.org/ns/ldp#DirectContainer":                         http_www_w3_org_ns_ldp_DirectContainer,
	"http://www.w3.org/ns/ldp#IndirectContainer":                       http_www_w3_org_ns_ldp_IndirectContainer,
	"http://www.w3.org/ns/ldp#NonRDFSource":                            http_www_w3_org_ns_ldp_NonRDFSource,
	"http://www.w3.org/ns/ldp#Page":                                    http_www_w3_org_ns_ldp_Page,
	"http://www.w3.org/ns/ldp#PageSortCriterion":                       http_www_w3_org_ns_ldp_PageSortCriterion,
	"http://www.w3.org/ns/ldp#RDFSource":                               http_www_w3_org_ns_ldp_RDFSource,
	"http://www.w3.org/ns/ldp#Resource":                                http_www_w3_org_ns_ldp_Resource,
	"https://w3id.org/security#Digest":                                 https_w3id_org_security_Digest,
	"https://w3id.org/security#EncryptedMessage":                       https_w3id_org_security_EncryptedMessage,
	"https://w3id.org/security#GraphSignature2012":                     https_w3id_org_security_GraphSignature2012,
	"https://w3id.org/security#Key":                                    https_w3id_org_security_Key,
	"https://w3id.org/security#LinkedDataSignature2015":                https_w3id_org_security_LinkedDataSignature2015,
	"https://w3id.org/security#LinkedDataSignature2016":                https_w3id_org_security_LinkedDataSignature2016,
	"https://w3id.org/security#Signature":                              https_w3id_org_security_Signature,
}

var (
	RDF = &Namespace{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
		Short: "rdf",
		Props: []*Prop{
			http_www_w3_org_1999_02_22_rdf_syntax_ns_first,
			http_www_w3_org_1999_02_22_rdf_syntax_ns_object,
			http_www_w3_org_1999_02_22_rdf_syntax_ns_predicate,
			http_www_w3_org_1999_02_22_rdf_syntax_ns_rest,
			http_www_w3_org_1999_02_22_rdf_syntax_ns_subject,
			http_www_w3_org_1999_02_22_rdf_syntax_ns_type,
			http_www_w3_org_1999_02_22_rdf_syntax_ns_value,
		},
		Types: map[string]*Type{
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#Alt":       http_www_w3_org_1999_02_22_rdf_syntax_ns_Alt,
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag":       http_www_w3_org_1999_02_22_rdf_syntax_ns_Bag,
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#List":      http_www_w3_org_1999_02_22_rdf_syntax_ns_List,
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#Property":  http_www_w3_org_1999_02_22_rdf_syntax_ns_Property,
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq":       http_www_w3_org_1999_02_22_rdf_syntax_ns_Seq,
			"http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement": http_www_w3_org_1999_02_22_rdf_syntax_ns_Statement,
		},
	}
	RDFS = &Namespace{
		ID:    "http://www.w3.org/2000/01/rdf-schema#",
		Short: "rdfs",
		Props: []*Prop{
			http_www_w3_org_2000_01_rdf_schema_comment,
			http_www_w3_org_2000_01_rdf_schema_domain,
			http_www_w3_org_2000_01_rdf_schema_isDefinedBy,
			http_www_w3_org_2000_01_rdf_schema_label,
			http_www_w3_org_2000_01_rdf_schema_member,
			http_www_w3_org_2000_01_rdf_schema_range,
			http_www_w3_org_2000_01_rdf_schema_seeAlso,
			http_www_w3_org_2000_01_rdf_schema_subClassOf,
			http_www_w3_org_2000_01_rdf_schema_subPropertyOf,
		},
		Types: map[string]*Type{
			"http://www.w3.org/2000/01/rdf-schema#Class":                       http_www_w3_org_2000_01_rdf_schema_Class,
			"http://www.w3.org/2000/01/rdf-schema#Container":                   http_www_w3_org_2000_01_rdf_schema_Container,
			"http://www.w3.org/2000/01/rdf-schema#ContainerMembershipProperty": http_www_w3_org_2000_01_rdf_schema_ContainerMembershipProperty,
			"http://www.w3.org/2000/01/rdf-schema#Datatype":                    http_www_w3_org_2000_01_rdf_schema_Datatype,
			"http://www.w3.org/2000/01/rdf-schema#Literal":                     http_www_w3_org_2000_01_rdf_schema_Literal,
			"http://www.w3.org/2000/01/rdf-schema#Resource":                    http_www_w3_org_2000_01_rdf_schema_Resource,
		},
	}
	OWL = &Namespace{
		ID:    "http://www.w3.org/2002/07/owl#",
		Short: "owl",
		Props: []*Prop{
			http_www_w3_org_2002_07_owl_allValuesFrom,
			http_www_w3_org_2002_07_owl_annotatedProperty,
			http_www_w3_org_2002_07_owl_annotatedSource,
			http_www_w3_org_2002_07_owl_annotatedTarget,
			http_www_w3_org_2002_07_owl_assertionProperty,
			http_www_w3_org_2002_07_owl_backwardCompatibleWith,
			http_www_w3_org_2002_07_owl_bottomDataProperty,
			http_www_w3_org_2002_07_owl_bottomObjectProperty,
			http_www_w3_org_2002_07_owl_cardinality,
			http_www_w3_org_2002_07_owl_complementOf,
			http_www_w3_org_2002_07_owl_datatypeComplementOf,
			http_www_w3_org_2002_07_owl_deprecated,
			http_www_w3_org_2002_07_owl_differentFrom,
			http_www_w3_org_2002_07_owl_disjointUnionOf,
			http_www_w3_org_2002_07_owl_disjointWith,
			http_www_w3_org_2002_07_owl_distinctMembers,
			http_www_w3_org_2002_07_owl_equivalentClass,
			http_www_w3_org_2002_07_owl_equivalentProperty,
			http_www_w3_org_2002_07_owl_hasKey,
			http_www_w3_org_2002_07_owl_hasSelf,
			http_www_w3_org_2002_07_owl_hasValue,
			http_www_w3_org_2002_07_owl_imports,
			http_www_w3_org_2002_07_owl_incompatibleWith,
			http_www_w3_org_2002_07_owl_intersectionOf,
			http_www_w3_org_2002_07_owl_inverseOf,
			http_www_w3_org_2002_07_owl_maxCardinality,
			http_www_w3_org_2002_07_owl_maxQualifiedCardinality,
			http_www_w3_org_2002_07_owl_members,
			http_www_w3_org_2002_07_owl_minCardinality,
			http_www_w3_org_2002_07_owl_minQualifiedCardinality,
			http_www_w3_org_2002_07_owl_onClass,
			http_www_w3_org_2002_07_owl_onDataRange,
			http_www_w3_org_2002_07_owl_onDatatype,
			http_www_w3_org_2002_07_owl_onProperties,
			http_www_w3_org_2002_07_owl_onProperty,
			http_www_w3_org_2002_07_owl_oneOf,
			http_www_w3_org_2002_07_owl_priorVersion,
			http_www_w3_org_2002_07_owl_propertyChainAxiom,
			http_www_w3_org_2002_07_owl_propertyDisjointWith,
			http_www_w3_org_2002_07_owl_qualifiedCardinality,
			http_www_w3_org_2002_07_owl_sameAs,
			http_www_w3_org_2002_07_owl_someValuesFrom,
			http_www_w3_org_2002_07_owl_sourceIndividual,
			http_www_w3_org_2002_07_owl_targetIndividual,
			http_www_w3_org_2002_07_owl_targetValue,
			http_www_w3_org_2002_07_owl_topDataProperty,
			http_www_w3_org_2002_07_owl_topObjectProperty,
			http_www_w3_org_2002_07_owl_unionOf,
			http_www_w3_org_2002_07_owl_versionIRI,
			http_www_w3_org_2002_07_owl_versionInfo,
			http_www_w3_org_2002_07_owl_withRestrictions,
		},
		Types: map[string]*Type{
			"http://www.w3.org/2002/07/owl#AllDifferent":              http_www_w3_org_2002_07_owl_AllDifferent,
			"http://www.w3.org/2002/07/owl#AllDisjointClasses":        http_www_w3_org_2002_07_owl_AllDisjointClasses,
			"http://www.w3.org/2002/07/owl#AllDisjointProperties":     http_www_w3_org_2002_07_owl_AllDisjointProperties,
			"http://www.w3.org/2002/07/owl#Annotation":                http_www_w3_org_2002_07_owl_Annotation,
			"http://www.w3.org/2002/07/owl#AnnotationProperty":        http_www_w3_org_2002_07_owl_AnnotationProperty,
			"http://www.w3.org/2002/07/owl#AsymmetricProperty":        http_www_w3_org_2002_07_owl_AsymmetricProperty,
			"http://www.w3.org/2002/07/owl#Axiom":                     http_www_w3_org_2002_07_owl_Axiom,
			"http://www.w3.org/2002/07/owl#Class":                     http_www_w3_org_2002_07_owl_Class,
			"http://www.w3.org/2002/07/owl#DataRange":                 http_www_w3_org_2002_07_owl_DataRange,
			"http://www.w3.org/2002/07/owl#DatatypeProperty":          http_www_w3_org_2002_07_owl_DatatypeProperty,
			"http://www.w3.org/2002/07/owl#DeprecatedClass":           http_www_w3_org_2002_07_owl_DeprecatedClass,
			"http://www.w3.org/2002/07/owl#DeprecatedProperty":        http_www_w3_org_2002_07_owl_DeprecatedProperty,
			"http://www.w3.org/2002/07/owl#FunctionalProperty":        http_www_w3_org_2002_07_owl_FunctionalProperty,
			"http://www.w3.org/2002/07/owl#InverseFunctionalProperty": http_www_w3_org_2002_07_owl_InverseFunctionalProperty,
			"http://www.w3.org/2002/07/owl#IrreflexiveProperty":       http_www_w3_org_2002_07_owl_IrreflexiveProperty,
			"http://www.w3.org/2002/07/owl#NamedIndividual":           http_www_w3_org_2002_07_owl_NamedIndividual,
			"http://www.w3.org/2002/07/owl#NegativePropertyAssertion": http_www_w3_org_2002_07_owl_NegativePropertyAssertion,
			"http://www.w3.org/2002/07/owl#Nothing":                   http_www_w3_org_2002_07_owl_Nothing,
			"http://www.w3.org/2002/07/owl#ObjectProperty":            http_www_w3_org_2002_07_owl_ObjectProperty,
			"http://www.w3.org/2002/07/owl#Ontology":                  http_www_w3_org_2002_07_owl_Ontology,
			"http://www.w3.org/2002/07/owl#OntologyProperty":          http_www_w3_org_2002_07_owl_OntologyProperty,
			"http://www.w3.org/2002/07/owl#ReflexiveProperty":         http_www_w3_org_2002_07_owl_ReflexiveProperty,
			"http://www.w3.org/2002/07/owl#Restriction":               http_www_w3_org_2002_07_owl_Restriction,
			"http://www.w3.org/2002/07/owl#SymmetricProperty":         http_www_w3_org_2002_07_owl_SymmetricProperty,
			"http://www.w3.org/2002/07/owl#Thing":                     http_www_w3_org_2002_07_owl_Thing,
			"http://www.w3.org/2002/07/owl#TransitiveProperty":        http_www_w3_org_2002_07_owl_TransitiveProperty,
		},
	}
	AS = &Namespace{
		ID:    "http://www.w3.org/ns/activitystreams#",
		Short: "as",
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_accuracy,
			http_www_w3_org_ns_activitystreams_actor,
			http_www_w3_org_ns_activitystreams_altitude,
			http_www_w3_org_ns_activitystreams_anyOf,
			http_www_w3_org_ns_activitystreams_attachment,
			http_www_w3_org_ns_activitystreams_attachments,
			http_www_w3_org_ns_activitystreams_attributedTo,
			http_www_w3_org_ns_activitystreams_audience,
			http_www_w3_org_ns_activitystreams_author,
			http_www_w3_org_ns_activitystreams_bcc,
			http_www_w3_org_ns_activitystreams_bto,
			http_www_w3_org_ns_activitystreams_cc,
			http_www_w3_org_ns_activitystreams_content,
			http_www_w3_org_ns_activitystreams_context,
			http_www_w3_org_ns_activitystreams_current,
			http_www_w3_org_ns_activitystreams_deleted,
			http_www_w3_org_ns_activitystreams_describes,
			http_www_w3_org_ns_activitystreams_downstreamDuplicates,
			http_www_w3_org_ns_activitystreams_duration,
			http_www_w3_org_ns_activitystreams_endTime,
			http_www_w3_org_ns_activitystreams_first,
			http_www_w3_org_ns_activitystreams_formerType,
			http_www_w3_org_ns_activitystreams_generator,
			http_www_w3_org_ns_activitystreams_height,
			http_www_w3_org_ns_activitystreams_href,
			http_www_w3_org_ns_activitystreams_hreflang,
			http_www_w3_org_ns_activitystreams_icon,
			http_www_w3_org_ns_activitystreams_id,
			http_www_w3_org_ns_activitystreams_image,
			http_www_w3_org_ns_activitystreams_inReplyTo,
			http_www_w3_org_ns_activitystreams_instrument,
			http_www_w3_org_ns_activitystreams_items,
			http_www_w3_org_ns_activitystreams_last,
			http_www_w3_org_ns_activitystreams_latitude,
			http_www_w3_org_ns_activitystreams_location,
			http_www_w3_org_ns_activitystreams_longitude,
			http_www_w3_org_ns_activitystreams_mediaType,
			http_www_w3_org_ns_activitystreams_name,
			http_www_w3_org_ns_activitystreams_next,
			http_www_w3_org_ns_activitystreams_object,
			http_www_w3_org_ns_activitystreams_objectType,
			http_www_w3_org_ns_activitystreams_oneOf,
			http_www_w3_org_ns_activitystreams_origin,
			http_www_w3_org_ns_activitystreams_partOf,
			http_www_w3_org_ns_activitystreams_prev,
			http_www_w3_org_ns_activitystreams_preview,
			http_www_w3_org_ns_activitystreams_provider,
			http_www_w3_org_ns_activitystreams_published,
			http_www_w3_org_ns_activitystreams_radius,
			http_www_w3_org_ns_activitystreams_rating,
			http_www_w3_org_ns_activitystreams_rel,
			http_www_w3_org_ns_activitystreams_relationship,
			http_www_w3_org_ns_activitystreams_replies,
			http_www_w3_org_ns_activitystreams_result,
			http_www_w3_org_ns_activitystreams_startIndex,
			http_www_w3_org_ns_activitystreams_startTime,
			http_www_w3_org_ns_activitystreams_subject,
			http_www_w3_org_ns_activitystreams_summary,
			http_www_w3_org_ns_activitystreams_tag,
			http_www_w3_org_ns_activitystreams_tags,
			http_www_w3_org_ns_activitystreams_target,
			http_www_w3_org_ns_activitystreams_to,
			http_www_w3_org_ns_activitystreams_totalItems,
			http_www_w3_org_ns_activitystreams_units,
			http_www_w3_org_ns_activitystreams_updated,
			http_www_w3_org_ns_activitystreams_upstreamDuplicates,
			http_www_w3_org_ns_activitystreams_url,
			http_www_w3_org_ns_activitystreams_verb,
			http_www_w3_org_ns_activitystreams_width,
		},
		Types: map[string]*Type{
			"http://www.w3.org/ns/activitystreams#Accept":                http_www_w3_org_ns_activitystreams_Accept,
			"http://www.w3.org/ns/activitystreams#Activity":              http_www_w3_org_ns_activitystreams_Activity,
			"http://www.w3.org/ns/activitystreams#Add":                   http_www_w3_org_ns_activitystreams_Add,
			"http://www.w3.org/ns/activitystreams#Announce":              http_www_w3_org_ns_activitystreams_Announce,
			"http://www.w3.org/ns/activitystreams#Application":           http_www_w3_org_ns_activitystreams_Application,
			"http://www.w3.org/ns/activitystreams#Arrive":                http_www_w3_org_ns_activitystreams_Arrive,
			"http://www.w3.org/ns/activitystreams#Article":               http_www_w3_org_ns_activitystreams_Article,
			"http://www.w3.org/ns/activitystreams#Audio":                 http_www_w3_org_ns_activitystreams_Audio,
			"http://www.w3.org/ns/activitystreams#Block":                 http_www_w3_org_ns_activitystreams_Block,
			"http://www.w3.org/ns/activitystreams#Collection":            http_www_w3_org_ns_activitystreams_Collection,
			"http://www.w3.org/ns/activitystreams#CollectionPage":        http_www_w3_org_ns_activitystreams_CollectionPage,
			"http://www.w3.org/ns/activitystreams#Create":                http_www_w3_org_ns_activitystreams_Create,
			"http://www.w3.org/ns/activitystreams#Delete":                http_www_w3_org_ns_activitystreams_Delete,
			"http://www.w3.org/ns/activitystreams#Dislike":               http_www_w3_org_ns_activitystreams_Dislike,
			"http://www.w3.org/ns/activitystreams#Document":              http_www_w3_org_ns_activitystreams_Document,
			"http://www.w3.org/ns/activitystreams#Event":                 http_www_w3_org_ns_activitystreams_Event,
			"http://www.w3.org/ns/activitystreams#Flag":                  http_www_w3_org_ns_activitystreams_Flag,
			"http://www.w3.org/ns/activitystreams#Follow":                http_www_w3_org_ns_activitystreams_Follow,
			"http://www.w3.org/ns/activitystreams#Group":                 http_www_w3_org_ns_activitystreams_Group,
			"http://www.w3.org/ns/activitystreams#Ignore":                http_www_w3_org_ns_activitystreams_Ignore,
			"http://www.w3.org/ns/activitystreams#Image":                 http_www_w3_org_ns_activitystreams_Image,
			"http://www.w3.org/ns/activitystreams#IntransitiveActivity":  http_www_w3_org_ns_activitystreams_IntransitiveActivity,
			"http://www.w3.org/ns/activitystreams#Invite":                http_www_w3_org_ns_activitystreams_Invite,
			"http://www.w3.org/ns/activitystreams#Join":                  http_www_w3_org_ns_activitystreams_Join,
			"http://www.w3.org/ns/activitystreams#Leave":                 http_www_w3_org_ns_activitystreams_Leave,
			"http://www.w3.org/ns/activitystreams#Like":                  http_www_w3_org_ns_activitystreams_Like,
			"http://www.w3.org/ns/activitystreams#Link":                  http_www_w3_org_ns_activitystreams_Link,
			"http://www.w3.org/ns/activitystreams#Listen":                http_www_w3_org_ns_activitystreams_Listen,
			"http://www.w3.org/ns/activitystreams#Mention":               http_www_w3_org_ns_activitystreams_Mention,
			"http://www.w3.org/ns/activitystreams#Move":                  http_www_w3_org_ns_activitystreams_Move,
			"http://www.w3.org/ns/activitystreams#Note":                  http_www_w3_org_ns_activitystreams_Note,
			"http://www.w3.org/ns/activitystreams#Object":                http_www_w3_org_ns_activitystreams_Object,
			"http://www.w3.org/ns/activitystreams#Offer":                 http_www_w3_org_ns_activitystreams_Offer,
			"http://www.w3.org/ns/activitystreams#OrderedCollection":     http_www_w3_org_ns_activitystreams_OrderedCollection,
			"http://www.w3.org/ns/activitystreams#OrderedCollectionPage": http_www_w3_org_ns_activitystreams_OrderedCollectionPage,
			"http://www.w3.org/ns/activitystreams#OrderedItems":          http_www_w3_org_ns_activitystreams_OrderedItems,
			"http://www.w3.org/ns/activitystreams#Organization":          http_www_w3_org_ns_activitystreams_Organization,
			"http://www.w3.org/ns/activitystreams#Page":                  http_www_w3_org_ns_activitystreams_Page,
			"http://www.w3.org/ns/activitystreams#Person":                http_www_w3_org_ns_activitystreams_Person,
			"http://www.w3.org/ns/activitystreams#Place":                 http_www_w3_org_ns_activitystreams_Place,
			"http://www.w3.org/ns/activitystreams#Profile":               http_www_w3_org_ns_activitystreams_Profile,
			"http://www.w3.org/ns/activitystreams#Question":              http_www_w3_org_ns_activitystreams_Question,
			"http://www.w3.org/ns/activitystreams#Read":                  http_www_w3_org_ns_activitystreams_Read,
			"http://www.w3.org/ns/activitystreams#Reject":                http_www_w3_org_ns_activitystreams_Reject,
			"http://www.w3.org/ns/activitystreams#Relationship":          http_www_w3_org_ns_activitystreams_Relationship,
			"http://www.w3.org/ns/activitystreams#Remove":                http_www_w3_org_ns_activitystreams_Remove,
			"http://www.w3.org/ns/activitystreams#Service":               http_www_w3_org_ns_activitystreams_Service,
			"http://www.w3.org/ns/activitystreams#TentativeAccept":       http_www_w3_org_ns_activitystreams_TentativeAccept,
			"http://www.w3.org/ns/activitystreams#TentativeReject":       http_www_w3_org_ns_activitystreams_TentativeReject,
			"http://www.w3.org/ns/activitystreams#Tombstone":             http_www_w3_org_ns_activitystreams_Tombstone,
			"http://www.w3.org/ns/activitystreams#Travel":                http_www_w3_org_ns_activitystreams_Travel,
			"http://www.w3.org/ns/activitystreams#Undo":                  http_www_w3_org_ns_activitystreams_Undo,
			"http://www.w3.org/ns/activitystreams#Update":                http_www_w3_org_ns_activitystreams_Update,
			"http://www.w3.org/ns/activitystreams#Video":                 http_www_w3_org_ns_activitystreams_Video,
			"http://www.w3.org/ns/activitystreams#View":                  http_www_w3_org_ns_activitystreams_View,
		},
	}
	LDP = &Namespace{
		ID:    "http://www.w3.org/ns/ldp#",
		Short: "ldp",
		Props: []*Prop{
			http_www_w3_org_ns_ldp_constrainedBy,
			http_www_w3_org_ns_ldp_contains,
			http_www_w3_org_ns_ldp_hasMemberRelation,
			http_www_w3_org_ns_ldp_inbox,
			http_www_w3_org_ns_ldp_insertedContentRelation,
			http_www_w3_org_ns_ldp_isMemberOfRelation,
			http_www_w3_org_ns_ldp_member,
			http_www_w3_org_ns_ldp_membershipResource,
			http_www_w3_org_ns_ldp_pageSequence,
			http_www_w3_org_ns_ldp_pageSortCollation,
			http_www_w3_org_ns_ldp_pageSortCriteria,
			http_www_w3_org_ns_ldp_pageSortOrder,
			http_www_w3_org_ns_ldp_pageSortPredicate,
		},
		Types: map[string]*Type{
			"http://www.w3.org/ns/ldp#BasicContainer":    http_www_w3_org_ns_ldp_BasicContainer,
			"http://www.w3.org/ns/ldp#Container":         http_www_w3_org_ns_ldp_Container,
			"http://www.w3.org/ns/ldp#DirectContainer":   http_www_w3_org_ns_ldp_DirectContainer,
			"http://www.w3.org/ns/ldp#IndirectContainer": http_www_w3_org_ns_ldp_IndirectContainer,
			"http://www.w3.org/ns/ldp#NonRDFSource":      http_www_w3_org_ns_ldp_NonRDFSource,
			"http://www.w3.org/ns/ldp#Page":              http_www_w3_org_ns_ldp_Page,
			"http://www.w3.org/ns/ldp#PageSortCriterion": http_www_w3_org_ns_ldp_PageSortCriterion,
			"http://www.w3.org/ns/ldp#RDFSource":         http_www_w3_org_ns_ldp_RDFSource,
			"http://www.w3.org/ns/ldp#Resource":          http_www_w3_org_ns_ldp_Resource,
		},
	}
	SEC = &Namespace{
		ID:    "https://w3id.org/security#",
		Short: "sec",
		Props: []*Prop{
			https_w3id_org_security_authenticationTag,
			https_w3id_org_security_canonicalizationAlgorithm,
			https_w3id_org_security_cipherAlgorithm,
			https_w3id_org_security_cipherData,
			https_w3id_org_security_cipherKey,
			https_w3id_org_security_creator,
			https_w3id_org_security_digestAlgorithm,
			https_w3id_org_security_digestValue,
			https_w3id_org_security_expires,
			https_w3id_org_security_initializationVector,
			https_w3id_org_security_nonce,
			https_w3id_org_security_owner,
			https_w3id_org_security_password,
			https_w3id_org_security_privateKeyPem,
			https_w3id_org_security_publicKey,
			https_w3id_org_security_publicKeyPem,
			https_w3id_org_security_publicKeyService,
			https_w3id_org_security_revoked,
			https_w3id_org_security_signature,
			https_w3id_org_security_signatureAlgorithm,
			https_w3id_org_security_signatureValue,
		},
		Types: map[string]*Type{
			"https://w3id.org/security#Digest":                  https_w3id_org_security_Digest,
			"https://w3id.org/security#EncryptedMessage":        https_w3id_org_security_EncryptedMessage,
			"https://w3id.org/security#GraphSignature2012":      https_w3id_org_security_GraphSignature2012,
			"https://w3id.org/security#Key":                     https_w3id_org_security_Key,
			"https://w3id.org/security#LinkedDataSignature2015": https_w3id_org_security_LinkedDataSignature2015,
			"https://w3id.org/security#LinkedDataSignature2016": https_w3id_org_security_LinkedDataSignature2016,
			"https://w3id.org/security#Signature":               https_w3id_org_security_Signature,
		},
	}

	http_www_w3_org_1999_02_22_rdf_syntax_ns_Alt = &Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#Alt",
		Short: "Alt",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsAlt(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_1999_02_22_rdf_syntax_ns_Bag = &Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag",
		Short: "Bag",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsBag(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_1999_02_22_rdf_syntax_ns_List = &Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#List",
		Short: "List",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsList(e)
		},
		Props: []*Prop{
			http_www_w3_org_1999_02_22_rdf_syntax_ns_first,
			http_www_w3_org_1999_02_22_rdf_syntax_ns_rest,
		},
	}
	http_www_w3_org_1999_02_22_rdf_syntax_ns_Property = &Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#Property",
		Short: "Property",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsProperty(e)
		},
		Props: []*Prop{
			http_www_w3_org_2000_01_rdf_schema_domain,
			http_www_w3_org_2000_01_rdf_schema_range,
			http_www_w3_org_2000_01_rdf_schema_subPropertyOf,
			http_www_w3_org_2002_07_owl_equivalentProperty,
			http_www_w3_org_2002_07_owl_propertyDisjointWith,
		},
	}
	http_www_w3_org_1999_02_22_rdf_syntax_ns_Seq = &Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq",
		Short: "Seq",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsSeq(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_1999_02_22_rdf_syntax_ns_Statement = &Type{
		ID:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement",
		Short: "Statement",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsStatement(e)
		},
		Props: []*Prop{
			http_www_w3_org_1999_02_22_rdf_syntax_ns_object,
			http_www_w3_org_1999_02_22_rdf_syntax_ns_predicate,
			http_www_w3_org_1999_02_22_rdf_syntax_ns_subject,
		},
	}
	http_www_w3_org_2000_01_rdf_schema_Class = &Type{
		ID:    "http://www.w3.org/2000/01/rdf-schema#Class",
		Short: "Class",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsClass(e)
		},
		Props: []*Prop{
			http_www_w3_org_2000_01_rdf_schema_subClassOf,
			http_www_w3_org_2002_07_owl_equivalentClass,
			http_www_w3_org_2002_07_owl_intersectionOf,
			http_www_w3_org_2002_07_owl_oneOf,
			http_www_w3_org_2002_07_owl_unionOf,
		},
	}
	http_www_w3_org_2000_01_rdf_schema_Container = &Type{
		ID:    "http://www.w3.org/2000/01/rdf-schema#Container",
		Short: "Container",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsContainer(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2000_01_rdf_schema_ContainerMembershipProperty = &Type{
		ID:    "http://www.w3.org/2000/01/rdf-schema#ContainerMembershipProperty",
		Short: "ContainerMembershipProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsContainerMembershipProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2000_01_rdf_schema_Datatype = &Type{
		ID:    "http://www.w3.org/2000/01/rdf-schema#Datatype",
		Short: "Datatype",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsDatatype(e)
		},
		Props: []*Prop{
			http_www_w3_org_2002_07_owl_datatypeComplementOf,
			http_www_w3_org_2002_07_owl_onDatatype,
			http_www_w3_org_2002_07_owl_withRestrictions,
		},
	}
	http_www_w3_org_2000_01_rdf_schema_Literal = &Type{
		ID:    "http://www.w3.org/2000/01/rdf-schema#Literal",
		Short: "Literal",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsLiteral(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2000_01_rdf_schema_Resource = &Type{
		ID:    "http://www.w3.org/2000/01/rdf-schema#Resource",
		Short: "Resource",
		Cast: func(e ld.Entity) ld.Entity {
			return rdf.AsResource(e)
		},
		Props: []*Prop{
			http_www_w3_org_1999_02_22_rdf_syntax_ns_type,
			http_www_w3_org_1999_02_22_rdf_syntax_ns_value,
			http_www_w3_org_2000_01_rdf_schema_comment,
			http_www_w3_org_2000_01_rdf_schema_isDefinedBy,
			http_www_w3_org_2000_01_rdf_schema_label,
			http_www_w3_org_2000_01_rdf_schema_member,
			http_www_w3_org_2000_01_rdf_schema_seeAlso,
			http_www_w3_org_2002_07_owl_annotatedProperty,
			http_www_w3_org_2002_07_owl_annotatedSource,
			http_www_w3_org_2002_07_owl_annotatedTarget,
			http_www_w3_org_2002_07_owl_deprecated,
			http_www_w3_org_2002_07_owl_members,
			http_www_w3_org_2002_07_owl_versionInfo,
		},
	}
	http_www_w3_org_2002_07_owl_AllDifferent = &Type{
		ID:    "http://www.w3.org/2002/07/owl#AllDifferent",
		Short: "AllDifferent",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsAllDifferent(e)
		},
		Props: []*Prop{
			http_www_w3_org_2002_07_owl_distinctMembers,
		},
	}
	http_www_w3_org_2002_07_owl_AllDisjointClasses = &Type{
		ID:    "http://www.w3.org/2002/07/owl#AllDisjointClasses",
		Short: "AllDisjointClasses",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsAllDisjointClasses(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_AllDisjointProperties = &Type{
		ID:    "http://www.w3.org/2002/07/owl#AllDisjointProperties",
		Short: "AllDisjointProperties",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsAllDisjointProperties(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_Annotation = &Type{
		ID:    "http://www.w3.org/2002/07/owl#Annotation",
		Short: "Annotation",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsAnnotation(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_AnnotationProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#AnnotationProperty",
		Short: "AnnotationProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsAnnotationProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_AsymmetricProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#AsymmetricProperty",
		Short: "AsymmetricProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsAsymmetricProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_Axiom = &Type{
		ID:    "http://www.w3.org/2002/07/owl#Axiom",
		Short: "Axiom",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsAxiom(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_Class = &Type{
		ID:    "http://www.w3.org/2002/07/owl#Class",
		Short: "Class",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsClass(e)
		},
		Props: []*Prop{
			http_www_w3_org_2002_07_owl_complementOf,
			http_www_w3_org_2002_07_owl_disjointUnionOf,
			http_www_w3_org_2002_07_owl_disjointWith,
			http_www_w3_org_2002_07_owl_hasKey,
		},
	}
	http_www_w3_org_2002_07_owl_DataRange = &Type{
		ID:    "http://www.w3.org/2002/07/owl#DataRange",
		Short: "DataRange",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsDataRange(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_DatatypeProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#DatatypeProperty",
		Short: "DatatypeProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsDatatypeProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_DeprecatedClass = &Type{
		ID:    "http://www.w3.org/2002/07/owl#DeprecatedClass",
		Short: "DeprecatedClass",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsDeprecatedClass(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_DeprecatedProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#DeprecatedProperty",
		Short: "DeprecatedProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsDeprecatedProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_FunctionalProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#FunctionalProperty",
		Short: "FunctionalProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsFunctionalProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_InverseFunctionalProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#InverseFunctionalProperty",
		Short: "InverseFunctionalProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsInverseFunctionalProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_IrreflexiveProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#IrreflexiveProperty",
		Short: "IrreflexiveProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsIrreflexiveProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_NamedIndividual = &Type{
		ID:    "http://www.w3.org/2002/07/owl#NamedIndividual",
		Short: "NamedIndividual",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsNamedIndividual(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_NegativePropertyAssertion = &Type{
		ID:    "http://www.w3.org/2002/07/owl#NegativePropertyAssertion",
		Short: "NegativePropertyAssertion",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsNegativePropertyAssertion(e)
		},
		Props: []*Prop{
			http_www_w3_org_2002_07_owl_assertionProperty,
			http_www_w3_org_2002_07_owl_sourceIndividual,
			http_www_w3_org_2002_07_owl_targetIndividual,
			http_www_w3_org_2002_07_owl_targetValue,
		},
	}
	http_www_w3_org_2002_07_owl_Nothing = &Type{
		ID:    "http://www.w3.org/2002/07/owl#Nothing",
		Short: "Nothing",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsNothing(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_ObjectProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#ObjectProperty",
		Short: "ObjectProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsObjectProperty(e)
		},
		Props: []*Prop{
			http_www_w3_org_2002_07_owl_inverseOf,
			http_www_w3_org_2002_07_owl_propertyChainAxiom,
		},
	}
	http_www_w3_org_2002_07_owl_Ontology = &Type{
		ID:    "http://www.w3.org/2002/07/owl#Ontology",
		Short: "Ontology",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsOntology(e)
		},
		Props: []*Prop{
			http_www_w3_org_2002_07_owl_backwardCompatibleWith,
			http_www_w3_org_2002_07_owl_imports,
			http_www_w3_org_2002_07_owl_incompatibleWith,
			http_www_w3_org_2002_07_owl_priorVersion,
			http_www_w3_org_2002_07_owl_versionIRI,
		},
	}
	http_www_w3_org_2002_07_owl_OntologyProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#OntologyProperty",
		Short: "OntologyProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsOntologyProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_ReflexiveProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#ReflexiveProperty",
		Short: "ReflexiveProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsReflexiveProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_Restriction = &Type{
		ID:    "http://www.w3.org/2002/07/owl#Restriction",
		Short: "Restriction",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsRestriction(e)
		},
		Props: []*Prop{
			http_www_w3_org_2002_07_owl_allValuesFrom,
			http_www_w3_org_2002_07_owl_cardinality,
			http_www_w3_org_2002_07_owl_hasSelf,
			http_www_w3_org_2002_07_owl_hasValue,
			http_www_w3_org_2002_07_owl_maxCardinality,
			http_www_w3_org_2002_07_owl_maxQualifiedCardinality,
			http_www_w3_org_2002_07_owl_minCardinality,
			http_www_w3_org_2002_07_owl_minQualifiedCardinality,
			http_www_w3_org_2002_07_owl_onClass,
			http_www_w3_org_2002_07_owl_onDataRange,
			http_www_w3_org_2002_07_owl_onProperties,
			http_www_w3_org_2002_07_owl_onProperty,
			http_www_w3_org_2002_07_owl_qualifiedCardinality,
			http_www_w3_org_2002_07_owl_someValuesFrom,
		},
	}
	http_www_w3_org_2002_07_owl_SymmetricProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#SymmetricProperty",
		Short: "SymmetricProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsSymmetricProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_2002_07_owl_Thing = &Type{
		ID:    "http://www.w3.org/2002/07/owl#Thing",
		Short: "Thing",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsThing(e)
		},
		Props: []*Prop{
			http_www_w3_org_2002_07_owl_bottomDataProperty,
			http_www_w3_org_2002_07_owl_bottomObjectProperty,
			http_www_w3_org_2002_07_owl_differentFrom,
			http_www_w3_org_2002_07_owl_sameAs,
			http_www_w3_org_2002_07_owl_topDataProperty,
			http_www_w3_org_2002_07_owl_topObjectProperty,
		},
	}
	http_www_w3_org_2002_07_owl_TransitiveProperty = &Type{
		ID:    "http://www.w3.org/2002/07/owl#TransitiveProperty",
		Short: "TransitiveProperty",
		Cast: func(e ld.Entity) ld.Entity {
			return owl.AsTransitiveProperty(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Accept = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Accept",
		Short: "Accept",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsAccept(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Activity = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Activity",
		Short: "Activity",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsActivity(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_actor,
			http_www_w3_org_ns_activitystreams_instrument,
			http_www_w3_org_ns_activitystreams_origin,
			http_www_w3_org_ns_activitystreams_result,
			http_www_w3_org_ns_activitystreams_target,
			http_www_w3_org_ns_activitystreams_verb,
		},
	}
	http_www_w3_org_ns_activitystreams_Add = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Add",
		Short: "Add",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsAdd(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Announce = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Announce",
		Short: "Announce",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsAnnounce(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Application = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Application",
		Short: "Application",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsApplication(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Arrive = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Arrive",
		Short: "Arrive",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsArrive(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Article = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Article",
		Short: "Article",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsArticle(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Audio = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Audio",
		Short: "Audio",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsAudio(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Block = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Block",
		Short: "Block",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsBlock(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Collection = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Collection",
		Short: "Collection",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsCollection(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_current,
			http_www_w3_org_ns_activitystreams_first,
			http_www_w3_org_ns_activitystreams_items,
			http_www_w3_org_ns_activitystreams_last,
			http_www_w3_org_ns_activitystreams_totalItems,
		},
	}
	http_www_w3_org_ns_activitystreams_CollectionPage = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#CollectionPage",
		Short: "CollectionPage",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsCollectionPage(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_next,
			http_www_w3_org_ns_activitystreams_partOf,
			http_www_w3_org_ns_activitystreams_prev,
		},
	}
	http_www_w3_org_ns_activitystreams_Create = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Create",
		Short: "Create",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsCreate(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Delete = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Delete",
		Short: "Delete",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsDelete(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Dislike = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Dislike",
		Short: "Dislike",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsDislike(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Document = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Document",
		Short: "Document",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsDocument(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Event = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Event",
		Short: "Event",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsEvent(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Flag = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Flag",
		Short: "Flag",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsFlag(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Follow = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Follow",
		Short: "Follow",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsFollow(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Group = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Group",
		Short: "Group",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsGroup(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Ignore = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Ignore",
		Short: "Ignore",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsIgnore(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Image = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Image",
		Short: "Image",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsImage(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_IntransitiveActivity = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#IntransitiveActivity",
		Short: "IntransitiveActivity",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsIntransitiveActivity(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Invite = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Invite",
		Short: "Invite",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsInvite(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Join = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Join",
		Short: "Join",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsJoin(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Leave = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Leave",
		Short: "Leave",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsLeave(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Like = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Like",
		Short: "Like",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsLike(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Link = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Link",
		Short: "Link",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsLink(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_height,
			http_www_w3_org_ns_activitystreams_href,
			http_www_w3_org_ns_activitystreams_hreflang,
			http_www_w3_org_ns_activitystreams_rel,
			http_www_w3_org_ns_activitystreams_width,
		},
	}
	http_www_w3_org_ns_activitystreams_Listen = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Listen",
		Short: "Listen",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsListen(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Mention = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Mention",
		Short: "Mention",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsMention(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Move = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Move",
		Short: "Move",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsMove(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Note = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Note",
		Short: "Note",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsNote(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Object = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Object",
		Short: "Object",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsObject(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_attachment,
			http_www_w3_org_ns_activitystreams_attachments,
			http_www_w3_org_ns_activitystreams_audience,
			http_www_w3_org_ns_activitystreams_author,
			http_www_w3_org_ns_activitystreams_bcc,
			http_www_w3_org_ns_activitystreams_bto,
			http_www_w3_org_ns_activitystreams_cc,
			http_www_w3_org_ns_activitystreams_content,
			http_www_w3_org_ns_activitystreams_context,
			http_www_w3_org_ns_activitystreams_downstreamDuplicates,
			http_www_w3_org_ns_activitystreams_duration,
			http_www_w3_org_ns_activitystreams_endTime,
			http_www_w3_org_ns_activitystreams_generator,
			http_www_w3_org_ns_activitystreams_icon,
			http_www_w3_org_ns_activitystreams_image,
			http_www_w3_org_ns_activitystreams_inReplyTo,
			http_www_w3_org_ns_activitystreams_location,
			http_www_w3_org_ns_activitystreams_objectType,
			http_www_w3_org_ns_activitystreams_provider,
			http_www_w3_org_ns_activitystreams_published,
			http_www_w3_org_ns_activitystreams_rating,
			http_www_w3_org_ns_activitystreams_replies,
			http_www_w3_org_ns_activitystreams_startTime,
			http_www_w3_org_ns_activitystreams_summary,
			http_www_w3_org_ns_activitystreams_tag,
			http_www_w3_org_ns_activitystreams_tags,
			http_www_w3_org_ns_activitystreams_to,
			http_www_w3_org_ns_activitystreams_updated,
			http_www_w3_org_ns_activitystreams_upstreamDuplicates,
			http_www_w3_org_ns_activitystreams_url,
		},
	}
	http_www_w3_org_ns_activitystreams_Offer = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Offer",
		Short: "Offer",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsOffer(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_OrderedCollection = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#OrderedCollection",
		Short: "OrderedCollection",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsOrderedCollection(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_OrderedCollectionPage = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#OrderedCollectionPage",
		Short: "OrderedCollectionPage",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsOrderedCollectionPage(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_startIndex,
		},
	}
	http_www_w3_org_ns_activitystreams_OrderedItems = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#OrderedItems",
		Short: "OrderedItems",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsOrderedItems(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Organization = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Organization",
		Short: "Organization",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsOrganization(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Page = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Page",
		Short: "Page",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsPage(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Person = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Person",
		Short: "Person",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsPerson(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Place = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Place",
		Short: "Place",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsPlace(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_accuracy,
			http_www_w3_org_ns_activitystreams_altitude,
			http_www_w3_org_ns_activitystreams_latitude,
			http_www_w3_org_ns_activitystreams_longitude,
			http_www_w3_org_ns_activitystreams_radius,
			http_www_w3_org_ns_activitystreams_units,
		},
	}
	http_www_w3_org_ns_activitystreams_Profile = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Profile",
		Short: "Profile",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsProfile(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_describes,
		},
	}
	http_www_w3_org_ns_activitystreams_Question = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Question",
		Short: "Question",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsQuestion(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_anyOf,
			http_www_w3_org_ns_activitystreams_oneOf,
		},
	}
	http_www_w3_org_ns_activitystreams_Read = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Read",
		Short: "Read",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsRead(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Reject = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Reject",
		Short: "Reject",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsReject(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Relationship = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Relationship",
		Short: "Relationship",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsRelationship(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_relationship,
			http_www_w3_org_ns_activitystreams_subject,
		},
	}
	http_www_w3_org_ns_activitystreams_Remove = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Remove",
		Short: "Remove",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsRemove(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Service = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Service",
		Short: "Service",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsService(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_TentativeAccept = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#TentativeAccept",
		Short: "TentativeAccept",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsTentativeAccept(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_TentativeReject = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#TentativeReject",
		Short: "TentativeReject",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsTentativeReject(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Tombstone = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Tombstone",
		Short: "Tombstone",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsTombstone(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_activitystreams_deleted,
			http_www_w3_org_ns_activitystreams_formerType,
		},
	}
	http_www_w3_org_ns_activitystreams_Travel = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Travel",
		Short: "Travel",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsTravel(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Undo = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Undo",
		Short: "Undo",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsUndo(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Update = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Update",
		Short: "Update",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsUpdate(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_Video = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#Video",
		Short: "Video",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsVideo(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_activitystreams_View = &Type{
		ID:    "http://www.w3.org/ns/activitystreams#View",
		Short: "View",
		Cast: func(e ld.Entity) ld.Entity {
			return as.AsView(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_ldp_BasicContainer = &Type{
		ID:    "http://www.w3.org/ns/ldp#BasicContainer",
		Short: "BasicContainer",
		Cast: func(e ld.Entity) ld.Entity {
			return ldp.AsBasicContainer(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_ldp_Container = &Type{
		ID:    "http://www.w3.org/ns/ldp#Container",
		Short: "Container",
		Cast: func(e ld.Entity) ld.Entity {
			return ldp.AsContainer(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_ldp_contains,
			http_www_w3_org_ns_ldp_hasMemberRelation,
			http_www_w3_org_ns_ldp_insertedContentRelation,
			http_www_w3_org_ns_ldp_isMemberOfRelation,
			http_www_w3_org_ns_ldp_membershipResource,
		},
	}
	http_www_w3_org_ns_ldp_DirectContainer = &Type{
		ID:    "http://www.w3.org/ns/ldp#DirectContainer",
		Short: "DirectContainer",
		Cast: func(e ld.Entity) ld.Entity {
			return ldp.AsDirectContainer(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_ldp_IndirectContainer = &Type{
		ID:    "http://www.w3.org/ns/ldp#IndirectContainer",
		Short: "IndirectContainer",
		Cast: func(e ld.Entity) ld.Entity {
			return ldp.AsIndirectContainer(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_ldp_NonRDFSource = &Type{
		ID:    "http://www.w3.org/ns/ldp#NonRDFSource",
		Short: "NonRDFSource",
		Cast: func(e ld.Entity) ld.Entity {
			return ldp.AsNonRDFSource(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_ldp_Page = &Type{
		ID:    "http://www.w3.org/ns/ldp#Page",
		Short: "Page",
		Cast: func(e ld.Entity) ld.Entity {
			return ldp.AsPage(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_ldp_pageSortCriteria,
		},
	}
	http_www_w3_org_ns_ldp_PageSortCriterion = &Type{
		ID:    "http://www.w3.org/ns/ldp#PageSortCriterion",
		Short: "PageSortCriterion",
		Cast: func(e ld.Entity) ld.Entity {
			return ldp.AsPageSortCriterion(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_ldp_pageSortCollation,
			http_www_w3_org_ns_ldp_pageSortOrder,
			http_www_w3_org_ns_ldp_pageSortPredicate,
		},
	}
	http_www_w3_org_ns_ldp_RDFSource = &Type{
		ID:    "http://www.w3.org/ns/ldp#RDFSource",
		Short: "RDFSource",
		Cast: func(e ld.Entity) ld.Entity {
			return ldp.AsRDFSource(e)
		},
		Props: []*Prop{},
	}
	http_www_w3_org_ns_ldp_Resource = &Type{
		ID:    "http://www.w3.org/ns/ldp#Resource",
		Short: "Resource",
		Cast: func(e ld.Entity) ld.Entity {
			return ldp.AsResource(e)
		},
		Props: []*Prop{
			http_www_w3_org_ns_ldp_constrainedBy,
			http_www_w3_org_ns_ldp_member,
		},
	}
	https_w3id_org_security_Digest = &Type{
		ID:    "https://w3id.org/security#Digest",
		Short: "Digest",
		Cast: func(e ld.Entity) ld.Entity {
			return sec.AsDigest(e)
		},
		Props: []*Prop{
			https_w3id_org_security_digestAlgorithm,
			https_w3id_org_security_digestValue,
		},
	}
	https_w3id_org_security_EncryptedMessage = &Type{
		ID:    "https://w3id.org/security#EncryptedMessage",
		Short: "EncryptedMessage",
		Cast: func(e ld.Entity) ld.Entity {
			return sec.AsEncryptedMessage(e)
		},
		Props: []*Prop{
			https_w3id_org_security_authenticationTag,
			https_w3id_org_security_cipherAlgorithm,
			https_w3id_org_security_cipherData,
			https_w3id_org_security_cipherKey,
			https_w3id_org_security_initializationVector,
			https_w3id_org_security_publicKey,
		},
	}
	https_w3id_org_security_GraphSignature2012 = &Type{
		ID:    "https://w3id.org/security#GraphSignature2012",
		Short: "GraphSignature2012",
		Cast: func(e ld.Entity) ld.Entity {
			return sec.AsGraphSignature2012(e)
		},
		Props: []*Prop{},
	}
	https_w3id_org_security_Key = &Type{
		ID:    "https://w3id.org/security#Key",
		Short: "Key",
		Cast: func(e ld.Entity) ld.Entity {
			return sec.AsKey(e)
		},
		Props: []*Prop{
			https_w3id_org_security_owner,
			https_w3id_org_security_privateKeyPem,
			https_w3id_org_security_publicKeyPem,
		},
	}
	https_w3id_org_security_LinkedDataSignature2015 = &Type{
		ID:    "https://w3id.org/security#LinkedDataSignature2015",
		Short: "LinkedDataSignature2015",
		Cast: func(e ld.Entity) ld.Entity {
			return sec.AsLinkedDataSignature2015(e)
		},
		Props: []*Prop{},
	}
	https_w3id_org_security_LinkedDataSignature2016 = &Type{
		ID:    "https://w3id.org/security#LinkedDataSignature2016",
		Short: "LinkedDataSignature2016",
		Cast: func(e ld.Entity) ld.Entity {
			return sec.AsLinkedDataSignature2016(e)
		},
		Props: []*Prop{},
	}
	https_w3id_org_security_Signature = &Type{
		ID:    "https://w3id.org/security#Signature",
		Short: "Signature",
		Cast: func(e ld.Entity) ld.Entity {
			return sec.AsSignature(e)
		},
		Props: []*Prop{
			https_w3id_org_security_creator,
			https_w3id_org_security_signatureValue,
		},
	}

	http_www_w3_org_1999_02_22_rdf_syntax_ns_first = &Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#first",
		Short:   "first",
		Comment: "The first item in the subject RDF list.",
	}
	http_www_w3_org_1999_02_22_rdf_syntax_ns_object = &Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#object",
		Short:   "object",
		Comment: "The object of the subject RDF statement.",
	}
	http_www_w3_org_1999_02_22_rdf_syntax_ns_predicate = &Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#predicate",
		Short:   "predicate",
		Comment: "The predicate of the subject RDF statement.",
	}
	http_www_w3_org_1999_02_22_rdf_syntax_ns_rest = &Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#rest",
		Short:   "rest",
		Comment: "The rest of the subject RDF list after the first item.",
	}
	http_www_w3_org_1999_02_22_rdf_syntax_ns_subject = &Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#subject",
		Short:   "subject",
		Comment: "The subject of the subject RDF statement.",
	}
	http_www_w3_org_1999_02_22_rdf_syntax_ns_type = &Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#type",
		Short:   "type",
		Comment: "The subject is an instance of a class.",
	}
	http_www_w3_org_1999_02_22_rdf_syntax_ns_value = &Prop{
		ID:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#value",
		Short:   "value",
		Comment: "Idiomatic property used for structured values.",
	}
	http_www_w3_org_2000_01_rdf_schema_comment = &Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#comment",
		Short:   "comment",
		Comment: "A description of the subject resource.",
	}
	http_www_w3_org_2000_01_rdf_schema_domain = &Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#domain",
		Short:   "domain",
		Comment: "A domain of the subject property.",
	}
	http_www_w3_org_2000_01_rdf_schema_isDefinedBy = &Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#isDefinedBy",
		Short:   "isDefinedBy",
		Comment: "The defininition of the subject resource.",
	}
	http_www_w3_org_2000_01_rdf_schema_label = &Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#label",
		Short:   "label",
		Comment: "A human-readable name for the subject.",
	}
	http_www_w3_org_2000_01_rdf_schema_member = &Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#member",
		Short:   "member",
		Comment: "A member of the subject resource.",
	}
	http_www_w3_org_2000_01_rdf_schema_range = &Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#range",
		Short:   "range",
		Comment: "A range of the subject property.",
	}
	http_www_w3_org_2000_01_rdf_schema_seeAlso = &Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#seeAlso",
		Short:   "seeAlso",
		Comment: "Further information about the subject resource.",
	}
	http_www_w3_org_2000_01_rdf_schema_subClassOf = &Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#subClassOf",
		Short:   "subClassOf",
		Comment: "The subject is a subclass of a class.",
	}
	http_www_w3_org_2000_01_rdf_schema_subPropertyOf = &Prop{
		ID:      "http://www.w3.org/2000/01/rdf-schema#subPropertyOf",
		Short:   "subPropertyOf",
		Comment: "The subject is a subproperty of a property.",
	}
	http_www_w3_org_2002_07_owl_allValuesFrom = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#allValuesFrom",
		Short:   "allValuesFrom",
		Comment: "The property that determines the class that a universal property restriction refers to.",
	}
	http_www_w3_org_2002_07_owl_annotatedProperty = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#annotatedProperty",
		Short:   "annotatedProperty",
		Comment: "The property that determines the predicate of an annotated axiom or annotated annotation.",
	}
	http_www_w3_org_2002_07_owl_annotatedSource = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#annotatedSource",
		Short:   "annotatedSource",
		Comment: "The property that determines the subject of an annotated axiom or annotated annotation.",
	}
	http_www_w3_org_2002_07_owl_annotatedTarget = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#annotatedTarget",
		Short:   "annotatedTarget",
		Comment: "The property that determines the object of an annotated axiom or annotated annotation.",
	}
	http_www_w3_org_2002_07_owl_assertionProperty = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#assertionProperty",
		Short:   "assertionProperty",
		Comment: "The property that determines the predicate of a negative property assertion.",
	}
	http_www_w3_org_2002_07_owl_backwardCompatibleWith = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#backwardCompatibleWith",
		Short:   "backwardCompatibleWith",
		Comment: "The annotation property that indicates that a given ontology is backward compatible with another ontology.",
	}
	http_www_w3_org_2002_07_owl_bottomDataProperty = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#bottomDataProperty",
		Short:   "bottomDataProperty",
		Comment: "The data property that does not relate any individual to any data value.",
	}
	http_www_w3_org_2002_07_owl_bottomObjectProperty = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#bottomObjectProperty",
		Short:   "bottomObjectProperty",
		Comment: "The object property that does not relate any two individuals.",
	}
	http_www_w3_org_2002_07_owl_cardinality = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#cardinality",
		Short:   "cardinality",
		Comment: "The property that determines the cardinality of an exact cardinality restriction.",
	}
	http_www_w3_org_2002_07_owl_complementOf = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#complementOf",
		Short:   "complementOf",
		Comment: "The property that determines that a given class is the complement of another class.",
	}
	http_www_w3_org_2002_07_owl_datatypeComplementOf = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#datatypeComplementOf",
		Short:   "datatypeComplementOf",
		Comment: "The property that determines that a given data range is the complement of another data range with respect to the data domain.",
	}
	http_www_w3_org_2002_07_owl_deprecated = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#deprecated",
		Short:   "deprecated",
		Comment: "The annotation property that indicates that a given entity has been deprecated.",
	}
	http_www_w3_org_2002_07_owl_differentFrom = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#differentFrom",
		Short:   "differentFrom",
		Comment: "The property that determines that two given individuals are different.",
	}
	http_www_w3_org_2002_07_owl_disjointUnionOf = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#disjointUnionOf",
		Short:   "disjointUnionOf",
		Comment: "The property that determines that a given class is equivalent to the disjoint union of a collection of other classes.",
	}
	http_www_w3_org_2002_07_owl_disjointWith = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#disjointWith",
		Short:   "disjointWith",
		Comment: "The property that determines that two given classes are disjoint.",
	}
	http_www_w3_org_2002_07_owl_distinctMembers = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#distinctMembers",
		Short:   "distinctMembers",
		Comment: "The property that determines the collection of pairwise different individuals in a owl:AllDifferent axiom.",
	}
	http_www_w3_org_2002_07_owl_equivalentClass = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#equivalentClass",
		Short:   "equivalentClass",
		Comment: "The property that determines that two given classes are equivalent, and that is used to specify datatype definitions.",
	}
	http_www_w3_org_2002_07_owl_equivalentProperty = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#equivalentProperty",
		Short:   "equivalentProperty",
		Comment: "The property that determines that two given properties are equivalent.",
	}
	http_www_w3_org_2002_07_owl_hasKey = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#hasKey",
		Short:   "hasKey",
		Comment: "The property that determines the collection of properties that jointly build a key.",
	}
	http_www_w3_org_2002_07_owl_hasSelf = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#hasSelf",
		Short:   "hasSelf",
		Comment: "The property that determines the property that a self restriction refers to.",
	}
	http_www_w3_org_2002_07_owl_hasValue = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#hasValue",
		Short:   "hasValue",
		Comment: "The property that determines the individual that a has-value restriction refers to.",
	}
	http_www_w3_org_2002_07_owl_imports = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#imports",
		Short:   "imports",
		Comment: "The property that is used for importing other ontologies into a given ontology.",
	}
	http_www_w3_org_2002_07_owl_incompatibleWith = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#incompatibleWith",
		Short:   "incompatibleWith",
		Comment: "The annotation property that indicates that a given ontology is incompatible with another ontology.",
	}
	http_www_w3_org_2002_07_owl_intersectionOf = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#intersectionOf",
		Short:   "intersectionOf",
		Comment: "The property that determines the collection of classes or data ranges that build an intersection.",
	}
	http_www_w3_org_2002_07_owl_inverseOf = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#inverseOf",
		Short:   "inverseOf",
		Comment: "The property that determines that two given properties are inverse.",
	}
	http_www_w3_org_2002_07_owl_maxCardinality = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#maxCardinality",
		Short:   "maxCardinality",
		Comment: "The property that determines the cardinality of a maximum cardinality restriction.",
	}
	http_www_w3_org_2002_07_owl_maxQualifiedCardinality = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#maxQualifiedCardinality",
		Short:   "maxQualifiedCardinality",
		Comment: "The property that determines the cardinality of a maximum qualified cardinality restriction.",
	}
	http_www_w3_org_2002_07_owl_members = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#members",
		Short:   "members",
		Comment: "The property that determines the collection of members in either a owl:AllDifferent, owl:AllDisjointClasses or owl:AllDisjointProperties axiom.",
	}
	http_www_w3_org_2002_07_owl_minCardinality = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#minCardinality",
		Short:   "minCardinality",
		Comment: "The property that determines the cardinality of a minimum cardinality restriction.",
	}
	http_www_w3_org_2002_07_owl_minQualifiedCardinality = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#minQualifiedCardinality",
		Short:   "minQualifiedCardinality",
		Comment: "The property that determines the cardinality of a minimum qualified cardinality restriction.",
	}
	http_www_w3_org_2002_07_owl_onClass = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#onClass",
		Short:   "onClass",
		Comment: "The property that determines the class that a qualified object cardinality restriction refers to.",
	}
	http_www_w3_org_2002_07_owl_onDataRange = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#onDataRange",
		Short:   "onDataRange",
		Comment: "The property that determines the data range that a qualified data cardinality restriction refers to.",
	}
	http_www_w3_org_2002_07_owl_onDatatype = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#onDatatype",
		Short:   "onDatatype",
		Comment: "The property that determines the datatype that a datatype restriction refers to.",
	}
	http_www_w3_org_2002_07_owl_onProperties = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#onProperties",
		Short:   "onProperties",
		Comment: "The property that determines the n-tuple of properties that a property restriction on an n-ary data range refers to.",
	}
	http_www_w3_org_2002_07_owl_onProperty = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#onProperty",
		Short:   "onProperty",
		Comment: "The property that determines the property that a property restriction refers to.",
	}
	http_www_w3_org_2002_07_owl_oneOf = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#oneOf",
		Short:   "oneOf",
		Comment: "The property that determines the collection of individuals or data values that build an enumeration.",
	}
	http_www_w3_org_2002_07_owl_priorVersion = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#priorVersion",
		Short:   "priorVersion",
		Comment: "The annotation property that indicates the predecessor ontology of a given ontology.",
	}
	http_www_w3_org_2002_07_owl_propertyChainAxiom = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#propertyChainAxiom",
		Short:   "propertyChainAxiom",
		Comment: "The property that determines the n-tuple of properties that build a sub property chain of a given property.",
	}
	http_www_w3_org_2002_07_owl_propertyDisjointWith = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#propertyDisjointWith",
		Short:   "propertyDisjointWith",
		Comment: "The property that determines that two given properties are disjoint.",
	}
	http_www_w3_org_2002_07_owl_qualifiedCardinality = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#qualifiedCardinality",
		Short:   "qualifiedCardinality",
		Comment: "The property that determines the cardinality of an exact qualified cardinality restriction.",
	}
	http_www_w3_org_2002_07_owl_sameAs = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#sameAs",
		Short:   "sameAs",
		Comment: "The property that determines that two given individuals are equal.",
	}
	http_www_w3_org_2002_07_owl_someValuesFrom = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#someValuesFrom",
		Short:   "someValuesFrom",
		Comment: "The property that determines the class that an existential property restriction refers to.",
	}
	http_www_w3_org_2002_07_owl_sourceIndividual = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#sourceIndividual",
		Short:   "sourceIndividual",
		Comment: "The property that determines the subject of a negative property assertion.",
	}
	http_www_w3_org_2002_07_owl_targetIndividual = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#targetIndividual",
		Short:   "targetIndividual",
		Comment: "The property that determines the object of a negative object property assertion.",
	}
	http_www_w3_org_2002_07_owl_targetValue = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#targetValue",
		Short:   "targetValue",
		Comment: "The property that determines the value of a negative data property assertion.",
	}
	http_www_w3_org_2002_07_owl_topDataProperty = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#topDataProperty",
		Short:   "topDataProperty",
		Comment: "The data property that relates every individual to every data value.",
	}
	http_www_w3_org_2002_07_owl_topObjectProperty = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#topObjectProperty",
		Short:   "topObjectProperty",
		Comment: "The object property that relates every two individuals.",
	}
	http_www_w3_org_2002_07_owl_unionOf = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#unionOf",
		Short:   "unionOf",
		Comment: "The property that determines the collection of classes or data ranges that build a union.",
	}
	http_www_w3_org_2002_07_owl_versionIRI = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#versionIRI",
		Short:   "versionIRI",
		Comment: "The property that identifies the version IRI of an ontology.",
	}
	http_www_w3_org_2002_07_owl_versionInfo = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#versionInfo",
		Short:   "versionInfo",
		Comment: "The annotation property that provides version information for an ontology or another OWL construct.",
	}
	http_www_w3_org_2002_07_owl_withRestrictions = &Prop{
		ID:      "http://www.w3.org/2002/07/owl#withRestrictions",
		Short:   "withRestrictions",
		Comment: "The property that determines the collection of facet-value pairs that define a datatype restriction.",
	}
	http_www_w3_org_ns_activitystreams_accuracy = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#accuracy",
		Short:   "accuracy",
		Comment: "Specifies the accuracy around the point established by the longitude and latitude",
	}
	http_www_w3_org_ns_activitystreams_actor = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#actor",
		Short:   "actor",
		Comment: "Subproperty of as:attributedTo that identifies the primary actor",
	}
	http_www_w3_org_ns_activitystreams_altitude = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#altitude",
		Short:   "altitude",
		Comment: "The altitude of a place",
	}
	http_www_w3_org_ns_activitystreams_anyOf = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#anyOf",
		Short:   "anyOf",
		Comment: "Describes a possible inclusive answer or option for a question.",
	}
	http_www_w3_org_ns_activitystreams_attachment = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#attachment",
		Short:   "attachment",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_attachments = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#attachments",
		Short:   "attachments",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_attributedTo = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#attributedTo",
		Short:   "attributedTo",
		Comment: "Identifies an entity to which an object is attributed",
	}
	http_www_w3_org_ns_activitystreams_audience = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#audience",
		Short:   "audience",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_author = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#author",
		Short:   "author",
		Comment: "Identifies the author of an object. Deprecated. Use as:attributedTo instead",
	}
	http_www_w3_org_ns_activitystreams_bcc = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#bcc",
		Short:   "bcc",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_bto = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#bto",
		Short:   "bto",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_cc = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#cc",
		Short:   "cc",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_content = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#content",
		Short:   "content",
		Comment: "The content of the object.",
	}
	http_www_w3_org_ns_activitystreams_context = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#context",
		Short:   "context",
		Comment: "Specifies the context within which an object exists or an activity was performed",
	}
	http_www_w3_org_ns_activitystreams_current = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#current",
		Short:   "current",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_deleted = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#deleted",
		Short:   "deleted",
		Comment: "Specifies the date and time the object was deleted",
	}
	http_www_w3_org_ns_activitystreams_describes = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#describes",
		Short:   "describes",
		Comment: "On a Profile object, describes the object described by the profile",
	}
	http_www_w3_org_ns_activitystreams_downstreamDuplicates = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#downstreamDuplicates",
		Short:   "downstreamDuplicates",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_duration = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#duration",
		Short:   "duration",
		Comment: "The duration of the object",
	}
	http_www_w3_org_ns_activitystreams_endTime = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#endTime",
		Short:   "endTime",
		Comment: "The ending time of the object",
	}
	http_www_w3_org_ns_activitystreams_first = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#first",
		Short:   "first",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_formerType = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#formerType",
		Short:   "formerType",
		Comment: "On a Tombstone object, describes the former type of the deleted object",
	}
	http_www_w3_org_ns_activitystreams_generator = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#generator",
		Short:   "generator",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_height = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#height",
		Short:   "height",
		Comment: "The display height expressed as device independent pixels",
	}
	http_www_w3_org_ns_activitystreams_href = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#href",
		Short:   "href",
		Comment: "The target URI of the Link",
	}
	http_www_w3_org_ns_activitystreams_hreflang = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#hreflang",
		Short:   "hreflang",
		Comment: "A hint about the language of the referenced resource",
	}
	http_www_w3_org_ns_activitystreams_icon = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#icon",
		Short:   "icon",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_id = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#id",
		Short:   "id",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_image = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#image",
		Short:   "image",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_inReplyTo = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#inReplyTo",
		Short:   "inReplyTo",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_instrument = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#instrument",
		Short:   "instrument",
		Comment: "Indentifies an object used (or to be used) to complete an activity",
	}
	http_www_w3_org_ns_activitystreams_items = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#items",
		Short:   "items",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_last = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#last",
		Short:   "last",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_latitude = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#latitude",
		Short:   "latitude",
		Comment: "The latitude",
	}
	http_www_w3_org_ns_activitystreams_location = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#location",
		Short:   "location",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_longitude = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#longitude",
		Short:   "longitude",
		Comment: "The longitude",
	}
	http_www_w3_org_ns_activitystreams_mediaType = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#mediaType",
		Short:   "mediaType",
		Comment: "The MIME Media Type",
	}
	http_www_w3_org_ns_activitystreams_name = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#name",
		Short:   "name",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_next = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#next",
		Short:   "next",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_object = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#object",
		Short:   "object",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_objectType = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#objectType",
		Short:   "objectType",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_oneOf = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#oneOf",
		Short:   "oneOf",
		Comment: "Describes a possible exclusive answer or option for a question.",
	}
	http_www_w3_org_ns_activitystreams_origin = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#origin",
		Short:   "origin",
		Comment: "For certain activities, specifies the entity from which the action is directed.",
	}
	http_www_w3_org_ns_activitystreams_partOf = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#partOf",
		Short:   "partOf",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_prev = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#prev",
		Short:   "prev",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_preview = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#preview",
		Short:   "preview",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_provider = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#provider",
		Short:   "provider",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_published = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#published",
		Short:   "published",
		Comment: "Specifies the date and time the object was published",
	}
	http_www_w3_org_ns_activitystreams_radius = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#radius",
		Short:   "radius",
		Comment: "Specifies a radius around the point established by the longitude and latitude",
	}
	http_www_w3_org_ns_activitystreams_rating = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#rating",
		Short:   "rating",
		Comment: "A numeric rating (>= 0.0, <= 5.0) for the object",
	}
	http_www_w3_org_ns_activitystreams_rel = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#rel",
		Short:   "rel",
		Comment: "The RFC 5988 or HTML5 Link Relation associated with the Link",
	}
	http_www_w3_org_ns_activitystreams_relationship = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#relationship",
		Short:   "relationship",
		Comment: "On a Relationship object, describes the type of relationship",
	}
	http_www_w3_org_ns_activitystreams_replies = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#replies",
		Short:   "replies",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_result = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#result",
		Short:   "result",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_startIndex = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#startIndex",
		Short:   "startIndex",
		Comment: "In a strictly ordered logical collection, specifies the index position of the first item in the items list",
	}
	http_www_w3_org_ns_activitystreams_startTime = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#startTime",
		Short:   "startTime",
		Comment: "The starting time of the object",
	}
	http_www_w3_org_ns_activitystreams_subject = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#subject",
		Short:   "subject",
		Comment: "On a Relationship object, identifies the subject. e.g. when saying \"John is connected to Sally\", 'subject' refers to 'John'",
	}
	http_www_w3_org_ns_activitystreams_summary = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#summary",
		Short:   "summary",
		Comment: "A short summary of the object",
	}
	http_www_w3_org_ns_activitystreams_tag = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#tag",
		Short:   "tag",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_tags = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#tags",
		Short:   "tags",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_target = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#target",
		Short:   "target",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_to = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#to",
		Short:   "to",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_totalItems = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#totalItems",
		Short:   "totalItems",
		Comment: "The total number of items in a logical collection",
	}
	http_www_w3_org_ns_activitystreams_units = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#units",
		Short:   "units",
		Comment: "Identifies the unit of measurement used by the radius, altitude and accuracy properties. The value can be expressed either as one of a set of predefined units or as a well-known common URI that identifies units.",
	}
	http_www_w3_org_ns_activitystreams_updated = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#updated",
		Short:   "updated",
		Comment: "Specifies when the object was last updated",
	}
	http_www_w3_org_ns_activitystreams_upstreamDuplicates = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#upstreamDuplicates",
		Short:   "upstreamDuplicates",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_url = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#url",
		Short:   "url",
		Comment: "Specifies a link to a specific representation of the Object",
	}
	http_www_w3_org_ns_activitystreams_verb = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#verb",
		Short:   "verb",
		Comment: "",
	}
	http_www_w3_org_ns_activitystreams_width = &Prop{
		ID:      "http://www.w3.org/ns/activitystreams#width",
		Short:   "width",
		Comment: "Specifies the preferred display width of the content, expressed in terms of device independent pixels.",
	}
	http_www_w3_org_ns_ldp_constrainedBy = &Prop{
		ID:      "http://www.w3.org/ns/ldp#constrainedBy",
		Short:   "constrainedBy",
		Comment: "Links a resource with constraints that the server requires requests like creation and update to conform to.",
	}
	http_www_w3_org_ns_ldp_contains = &Prop{
		ID:      "http://www.w3.org/ns/ldp#contains",
		Short:   "contains",
		Comment: "Links a container with resources created through the container.",
	}
	http_www_w3_org_ns_ldp_hasMemberRelation = &Prop{
		ID:      "http://www.w3.org/ns/ldp#hasMemberRelation",
		Short:   "hasMemberRelation",
		Comment: "Indicates which predicate is used in membership triples, and that the membership triple pattern is < membership-constant-URI , object-of-hasMemberRelation, member-URI >.",
	}
	http_www_w3_org_ns_ldp_inbox = &Prop{
		ID:      "http://www.w3.org/ns/ldp#inbox",
		Short:   "inbox",
		Comment: "Links a resource to a container where notifications for the resource can be created and discovered.",
	}
	http_www_w3_org_ns_ldp_insertedContentRelation = &Prop{
		ID:      "http://www.w3.org/ns/ldp#insertedContentRelation",
		Short:   "insertedContentRelation",
		Comment: "Indicates which triple in a creation request should be used as the member-URI value in the membership triple added when the creation request is successful.",
	}
	http_www_w3_org_ns_ldp_isMemberOfRelation = &Prop{
		ID:      "http://www.w3.org/ns/ldp#isMemberOfRelation",
		Short:   "isMemberOfRelation",
		Comment: "Indicates which predicate is used in membership triples, and that the membership triple pattern is < member-URI , object-of-isMemberOfRelation, membership-constant-URI >.",
	}
	http_www_w3_org_ns_ldp_member = &Prop{
		ID:      "http://www.w3.org/ns/ldp#member",
		Short:   "member",
		Comment: "LDP servers should use this predicate as the membership predicate if there is no obvious predicate from an application vocabulary to use.",
	}
	http_www_w3_org_ns_ldp_membershipResource = &Prop{
		ID:      "http://www.w3.org/ns/ldp#membershipResource",
		Short:   "membershipResource",
		Comment: "Indicates the membership-constant-URI in a membership triple.  Depending upon the membership triple pattern a container uses, as indicated by the presence of ldp:hasMemberRelation or ldp:isMemberOfRelation, the membership-constant-URI might occupy either the subject or object position in membership triples.",
	}
	http_www_w3_org_ns_ldp_pageSequence = &Prop{
		ID:      "http://www.w3.org/ns/ldp#pageSequence",
		Short:   "pageSequence",
		Comment: "Link to a page sequence resource, as defined by LDP Paging.  Typically used to communicate the sorting criteria used to allocate LDPC members to pages.",
	}
	http_www_w3_org_ns_ldp_pageSortCollation = &Prop{
		ID:      "http://www.w3.org/ns/ldp#pageSortCollation",
		Short:   "pageSortCollation",
		Comment: "The collation used to order the members across pages in a page sequence when comparing strings.",
	}
	http_www_w3_org_ns_ldp_pageSortCriteria = &Prop{
		ID:      "http://www.w3.org/ns/ldp#pageSortCriteria",
		Short:   "pageSortCriteria",
		Comment: "Link to the list of sorting criteria used by the server in a representation.  Typically used on Link response headers as an extension link relation URI in the rel= parameter.",
	}
	http_www_w3_org_ns_ldp_pageSortOrder = &Prop{
		ID:      "http://www.w3.org/ns/ldp#pageSortOrder",
		Short:   "pageSortOrder",
		Comment: "The ascending/descending/etc order used to order the members across pages in a page sequence.",
	}
	http_www_w3_org_ns_ldp_pageSortPredicate = &Prop{
		ID:      "http://www.w3.org/ns/ldp#pageSortPredicate",
		Short:   "pageSortPredicate",
		Comment: "Predicate used to specify the order of the members across a page sequence's in-sequence page resources; it asserts nothing about the order of members in the representation of a single page.",
	}
	https_w3id_org_security_authenticationTag = &Prop{
		ID:      "https://w3id.org/security#authenticationTag",
		Short:   "authenticationTag",
		Comment: "",
	}
	https_w3id_org_security_canonicalizationAlgorithm = &Prop{
		ID:      "https://w3id.org/security#canonicalizationAlgorithm",
		Short:   "canonicalizationAlgorithm",
		Comment: "The canonicalization algorithm is used to transform the input data into a form that can be passed to a cryptographic digest method. The digest is then digitally signed using a digital signature algorithm. Canonicalization ensures that a piece of software that is generating a digital signature is able to do so on the same set of information in a deterministic manner.",
	}
	https_w3id_org_security_cipherAlgorithm = &Prop{
		ID:      "https://w3id.org/security#cipherAlgorithm",
		Short:   "cipherAlgorithm",
		Comment: "The cipher algorithm describes the mechanism used to encrypt a message. It is typically a string expressing the cipher suite, the strength of the cipher, and a block cipher mode.",
	}
	https_w3id_org_security_cipherData = &Prop{
		ID:      "https://w3id.org/security#cipherData",
		Short:   "cipherData",
		Comment: "Cipher data an opaque blob of information that is used to specify an encrypted message.",
	}
	https_w3id_org_security_cipherKey = &Prop{
		ID:      "https://w3id.org/security#cipherKey",
		Short:   "cipherKey",
		Comment: "A cipher key is a symmetric key that is used to encrypt or decrypt a piece of information. The key itself may be expressed in clear text or encrypted.",
	}
	https_w3id_org_security_creator = &Prop{
		ID:      "https://w3id.org/security#creator",
		Short:   "creator",
		Comment: "",
	}
	https_w3id_org_security_digestAlgorithm = &Prop{
		ID:      "https://w3id.org/security#digestAlgorithm",
		Short:   "digestAlgorithm",
		Comment: "The digest algorithm is used to specify the cryptographic function to use when generating the data to be digitally signed. Typically, data that is to be signed goes through three steps: 1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm that should be used for step #2. A signature class typically specifies a default digest method, so this property is typically used to specify information for a signature algorithm.",
	}
	https_w3id_org_security_digestValue = &Prop{
		ID:      "https://w3id.org/security#digestValue",
		Short:   "digestValue",
		Comment: "The digest value is used to express the output of the digest algorithm expressed in Base-1 (hexadecimal) format.",
	}
	https_w3id_org_security_expires = &Prop{
		ID:      "https://w3id.org/security#expires",
		Short:   "expires",
		Comment: "The expiration time is typically associated with a Key and specifies when the validity of the key will expire. It is considered a best practice to only create keys that have very definite expiration periods. This period is typically set to between six months and two years. An digital signature created using an expired key MUST be marked as invalid by any software attempting to verify the signature.",
	}
	https_w3id_org_security_initializationVector = &Prop{
		ID:      "https://w3id.org/security#initializationVector",
		Short:   "initializationVector",
		Comment: "The initialization vector (IV) is a byte stream that is typically used to initialize certain block cipher encryption schemes. For a receiving application to be able to decrypt a message, it must know the decryption key and the initialization vector. The value is typically base-64 encoded.",
	}
	https_w3id_org_security_nonce = &Prop{
		ID:      "https://w3id.org/security#nonce",
		Short:   "nonce",
		Comment: "This property is used in conjunction with the input to the signature hashing function in order to protect against replay attacks. Typically, receivers need to track all nonce values used within a certain time period in order to ensure that an attacker cannot merely re-send a compromised packet in order to execute a privileged request.",
	}
	https_w3id_org_security_owner = &Prop{
		ID:      "https://w3id.org/security#owner",
		Short:   "owner",
		Comment: "An owner is an entity that claims control over a particular resource. Note that ownership is best validated as a two-way relationship where the owner claims ownership over a particular resource, and the resource clearly identifies its owner.",
	}
	https_w3id_org_security_password = &Prop{
		ID:      "https://w3id.org/security#password",
		Short:   "password",
		Comment: "A secret that is used to generate a key that can be used to encrypt or decrypt message. It is typically a string value.",
	}
	https_w3id_org_security_privateKeyPem = &Prop{
		ID:      "https://w3id.org/security#privateKeyPem",
		Short:   "privateKeyPem",
		Comment: "A private key PEM property is used to specify the PEM-encoded version of the private key. This encoding is compatible with almost every Secure Sockets Layer library implementation and typically plugs directly into functions intializing private keys.",
	}
	https_w3id_org_security_publicKey = &Prop{
		ID:      "https://w3id.org/security#publicKey",
		Short:   "publicKey",
		Comment: "A public key property is used to specify a URL that contains information about a public key.",
	}
	https_w3id_org_security_publicKeyPem = &Prop{
		ID:      "https://w3id.org/security#publicKeyPem",
		Short:   "publicKeyPem",
		Comment: "A public key PEM property is used to specify the PEM-encoded version of the public key. This encoding is compatible with almost every Secure Sockets Layer library implementation and typically plugs directly into functions intializing public keys.",
	}
	https_w3id_org_security_publicKeyService = &Prop{
		ID:      "https://w3id.org/security#publicKeyService",
		Short:   "publicKeyService",
		Comment: "The publicKeyService property is used to express the REST URL that provides public key management services as defined by the Web Key [SECURE-MESSAGING] specification.",
	}
	https_w3id_org_security_revoked = &Prop{
		ID:      "https://w3id.org/security#revoked",
		Short:   "revoked",
		Comment: "The revocation time is typically associated with a Key that has been marked as invalid as of the date and time associated with the property. Key revocations are often used when a key is compromised, such as the theft of the private key, or during the course of best-practice key rotation schedules.",
	}
	https_w3id_org_security_signature = &Prop{
		ID:      "https://w3id.org/security#signature",
		Short:   "signature",
		Comment: "The signature property is used to associate a signature with a graph of information. The signature property is typically not included in the canonicalized graph that is then digested, and digitally signed.",
	}
	https_w3id_org_security_signatureAlgorithm = &Prop{
		ID:      "https://w3id.org/security#signatureAlgorithm",
		Short:   "signatureAlgorithm",
		Comment: "The signature algorithm is used to specify the cryptographic signature function to use when digitally signing the digest data. Typically, text to be signed goes through three steps: 1) canonicalization, 2) digest, and 3) signature. This property is used to specify the algorithm that should be used for step #3. A signature class typically specifies a default signature algorithm, so this property rarely needs to be used in practice when specifying digital signatures.",
	}
	https_w3id_org_security_signatureValue = &Prop{
		ID:      "https://w3id.org/security#signatureValue",
		Short:   "signatureValue",
		Comment: "The signature value is used to express the output of the signature algorithm expressed in base-64 format.",
	}
)
