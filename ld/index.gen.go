// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package ld

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/as"
	"github.com/meowpub/meow/ld/ns/ldp"
	"github.com/meowpub/meow/ld/ns/owl"
	"github.com/meowpub/meow/ld/ns/rdf"
	"github.com/meowpub/meow/ld/ns/sec"
)

var Index = map[string]func(e ld.Entity) ld.Entity{
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#Alt": func(e ld.Entity) ld.Entity {
		return rdf.AsAlt(e)
	},
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#Bag": func(e ld.Entity) ld.Entity {
		return rdf.AsBag(e)
	},
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#List": func(e ld.Entity) ld.Entity {
		return rdf.AsList(e)
	},
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#Property": func(e ld.Entity) ld.Entity {
		return rdf.AsProperty(e)
	},
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#Seq": func(e ld.Entity) ld.Entity {
		return rdf.AsSeq(e)
	},
	"http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement": func(e ld.Entity) ld.Entity {
		return rdf.AsStatement(e)
	},
	"http://www.w3.org/2000/01/rdf-schema#Class": func(e ld.Entity) ld.Entity {
		return rdf.AsClass(e)
	},
	"http://www.w3.org/2000/01/rdf-schema#Container": func(e ld.Entity) ld.Entity {
		return rdf.AsContainer(e)
	},
	"http://www.w3.org/2000/01/rdf-schema#ContainerMembershipProperty": func(e ld.Entity) ld.Entity {
		return rdf.AsContainerMembershipProperty(e)
	},
	"http://www.w3.org/2000/01/rdf-schema#Datatype": func(e ld.Entity) ld.Entity {
		return rdf.AsDatatype(e)
	},
	"http://www.w3.org/2000/01/rdf-schema#Literal": func(e ld.Entity) ld.Entity {
		return rdf.AsLiteral(e)
	},
	"http://www.w3.org/2000/01/rdf-schema#Resource": func(e ld.Entity) ld.Entity {
		return rdf.AsResource(e)
	},
	"http://www.w3.org/2002/07/owl#AllDifferent": func(e ld.Entity) ld.Entity {
		return owl.AsAllDifferent(e)
	},
	"http://www.w3.org/2002/07/owl#AllDisjointClasses": func(e ld.Entity) ld.Entity {
		return owl.AsAllDisjointClasses(e)
	},
	"http://www.w3.org/2002/07/owl#AllDisjointProperties": func(e ld.Entity) ld.Entity {
		return owl.AsAllDisjointProperties(e)
	},
	"http://www.w3.org/2002/07/owl#Annotation": func(e ld.Entity) ld.Entity {
		return owl.AsAnnotation(e)
	},
	"http://www.w3.org/2002/07/owl#AnnotationProperty": func(e ld.Entity) ld.Entity {
		return owl.AsAnnotationProperty(e)
	},
	"http://www.w3.org/2002/07/owl#AsymmetricProperty": func(e ld.Entity) ld.Entity {
		return owl.AsAsymmetricProperty(e)
	},
	"http://www.w3.org/2002/07/owl#Axiom": func(e ld.Entity) ld.Entity {
		return owl.AsAxiom(e)
	},
	"http://www.w3.org/2002/07/owl#Class": func(e ld.Entity) ld.Entity {
		return owl.AsClass(e)
	},
	"http://www.w3.org/2002/07/owl#DataRange": func(e ld.Entity) ld.Entity {
		return owl.AsDataRange(e)
	},
	"http://www.w3.org/2002/07/owl#DatatypeProperty": func(e ld.Entity) ld.Entity {
		return owl.AsDatatypeProperty(e)
	},
	"http://www.w3.org/2002/07/owl#DeprecatedClass": func(e ld.Entity) ld.Entity {
		return owl.AsDeprecatedClass(e)
	},
	"http://www.w3.org/2002/07/owl#DeprecatedProperty": func(e ld.Entity) ld.Entity {
		return owl.AsDeprecatedProperty(e)
	},
	"http://www.w3.org/2002/07/owl#FunctionalProperty": func(e ld.Entity) ld.Entity {
		return owl.AsFunctionalProperty(e)
	},
	"http://www.w3.org/2002/07/owl#InverseFunctionalProperty": func(e ld.Entity) ld.Entity {
		return owl.AsInverseFunctionalProperty(e)
	},
	"http://www.w3.org/2002/07/owl#IrreflexiveProperty": func(e ld.Entity) ld.Entity {
		return owl.AsIrreflexiveProperty(e)
	},
	"http://www.w3.org/2002/07/owl#NamedIndividual": func(e ld.Entity) ld.Entity {
		return owl.AsNamedIndividual(e)
	},
	"http://www.w3.org/2002/07/owl#NegativePropertyAssertion": func(e ld.Entity) ld.Entity {
		return owl.AsNegativePropertyAssertion(e)
	},
	"http://www.w3.org/2002/07/owl#Nothing": func(e ld.Entity) ld.Entity {
		return owl.AsNothing(e)
	},
	"http://www.w3.org/2002/07/owl#ObjectProperty": func(e ld.Entity) ld.Entity {
		return owl.AsObjectProperty(e)
	},
	"http://www.w3.org/2002/07/owl#Ontology": func(e ld.Entity) ld.Entity {
		return owl.AsOntology(e)
	},
	"http://www.w3.org/2002/07/owl#OntologyProperty": func(e ld.Entity) ld.Entity {
		return owl.AsOntologyProperty(e)
	},
	"http://www.w3.org/2002/07/owl#ReflexiveProperty": func(e ld.Entity) ld.Entity {
		return owl.AsReflexiveProperty(e)
	},
	"http://www.w3.org/2002/07/owl#Restriction": func(e ld.Entity) ld.Entity {
		return owl.AsRestriction(e)
	},
	"http://www.w3.org/2002/07/owl#SymmetricProperty": func(e ld.Entity) ld.Entity {
		return owl.AsSymmetricProperty(e)
	},
	"http://www.w3.org/2002/07/owl#Thing": func(e ld.Entity) ld.Entity {
		return owl.AsThing(e)
	},
	"http://www.w3.org/2002/07/owl#TransitiveProperty": func(e ld.Entity) ld.Entity {
		return owl.AsTransitiveProperty(e)
	},
	"http://www.w3.org/ns/activitystreams#Accept": func(e ld.Entity) ld.Entity {
		return as.AsAccept(e)
	},
	"http://www.w3.org/ns/activitystreams#Activity": func(e ld.Entity) ld.Entity {
		return as.AsActivity(e)
	},
	"http://www.w3.org/ns/activitystreams#Add": func(e ld.Entity) ld.Entity {
		return as.AsAdd(e)
	},
	"http://www.w3.org/ns/activitystreams#Announce": func(e ld.Entity) ld.Entity {
		return as.AsAnnounce(e)
	},
	"http://www.w3.org/ns/activitystreams#Application": func(e ld.Entity) ld.Entity {
		return as.AsApplication(e)
	},
	"http://www.w3.org/ns/activitystreams#Arrive": func(e ld.Entity) ld.Entity {
		return as.AsArrive(e)
	},
	"http://www.w3.org/ns/activitystreams#Article": func(e ld.Entity) ld.Entity {
		return as.AsArticle(e)
	},
	"http://www.w3.org/ns/activitystreams#Audio": func(e ld.Entity) ld.Entity {
		return as.AsAudio(e)
	},
	"http://www.w3.org/ns/activitystreams#Block": func(e ld.Entity) ld.Entity {
		return as.AsBlock(e)
	},
	"http://www.w3.org/ns/activitystreams#Collection": func(e ld.Entity) ld.Entity {
		return as.AsCollection(e)
	},
	"http://www.w3.org/ns/activitystreams#CollectionPage": func(e ld.Entity) ld.Entity {
		return as.AsCollectionPage(e)
	},
	"http://www.w3.org/ns/activitystreams#Create": func(e ld.Entity) ld.Entity {
		return as.AsCreate(e)
	},
	"http://www.w3.org/ns/activitystreams#Delete": func(e ld.Entity) ld.Entity {
		return as.AsDelete(e)
	},
	"http://www.w3.org/ns/activitystreams#Dislike": func(e ld.Entity) ld.Entity {
		return as.AsDislike(e)
	},
	"http://www.w3.org/ns/activitystreams#Document": func(e ld.Entity) ld.Entity {
		return as.AsDocument(e)
	},
	"http://www.w3.org/ns/activitystreams#Event": func(e ld.Entity) ld.Entity {
		return as.AsEvent(e)
	},
	"http://www.w3.org/ns/activitystreams#Flag": func(e ld.Entity) ld.Entity {
		return as.AsFlag(e)
	},
	"http://www.w3.org/ns/activitystreams#Follow": func(e ld.Entity) ld.Entity {
		return as.AsFollow(e)
	},
	"http://www.w3.org/ns/activitystreams#Group": func(e ld.Entity) ld.Entity {
		return as.AsGroup(e)
	},
	"http://www.w3.org/ns/activitystreams#Ignore": func(e ld.Entity) ld.Entity {
		return as.AsIgnore(e)
	},
	"http://www.w3.org/ns/activitystreams#Image": func(e ld.Entity) ld.Entity {
		return as.AsImage(e)
	},
	"http://www.w3.org/ns/activitystreams#IntransitiveActivity": func(e ld.Entity) ld.Entity {
		return as.AsIntransitiveActivity(e)
	},
	"http://www.w3.org/ns/activitystreams#Invite": func(e ld.Entity) ld.Entity {
		return as.AsInvite(e)
	},
	"http://www.w3.org/ns/activitystreams#Join": func(e ld.Entity) ld.Entity {
		return as.AsJoin(e)
	},
	"http://www.w3.org/ns/activitystreams#Leave": func(e ld.Entity) ld.Entity {
		return as.AsLeave(e)
	},
	"http://www.w3.org/ns/activitystreams#Like": func(e ld.Entity) ld.Entity {
		return as.AsLike(e)
	},
	"http://www.w3.org/ns/activitystreams#Link": func(e ld.Entity) ld.Entity {
		return as.AsLink(e)
	},
	"http://www.w3.org/ns/activitystreams#Listen": func(e ld.Entity) ld.Entity {
		return as.AsListen(e)
	},
	"http://www.w3.org/ns/activitystreams#Mention": func(e ld.Entity) ld.Entity {
		return as.AsMention(e)
	},
	"http://www.w3.org/ns/activitystreams#Move": func(e ld.Entity) ld.Entity {
		return as.AsMove(e)
	},
	"http://www.w3.org/ns/activitystreams#Note": func(e ld.Entity) ld.Entity {
		return as.AsNote(e)
	},
	"http://www.w3.org/ns/activitystreams#Object": func(e ld.Entity) ld.Entity {
		return as.AsObject(e)
	},
	"http://www.w3.org/ns/activitystreams#Offer": func(e ld.Entity) ld.Entity {
		return as.AsOffer(e)
	},
	"http://www.w3.org/ns/activitystreams#OrderedCollection": func(e ld.Entity) ld.Entity {
		return as.AsOrderedCollection(e)
	},
	"http://www.w3.org/ns/activitystreams#OrderedCollectionPage": func(e ld.Entity) ld.Entity {
		return as.AsOrderedCollectionPage(e)
	},
	"http://www.w3.org/ns/activitystreams#OrderedItems": func(e ld.Entity) ld.Entity {
		return as.AsOrderedItems(e)
	},
	"http://www.w3.org/ns/activitystreams#Organization": func(e ld.Entity) ld.Entity {
		return as.AsOrganization(e)
	},
	"http://www.w3.org/ns/activitystreams#Page": func(e ld.Entity) ld.Entity {
		return as.AsPage(e)
	},
	"http://www.w3.org/ns/activitystreams#Person": func(e ld.Entity) ld.Entity {
		return as.AsPerson(e)
	},
	"http://www.w3.org/ns/activitystreams#Place": func(e ld.Entity) ld.Entity {
		return as.AsPlace(e)
	},
	"http://www.w3.org/ns/activitystreams#Profile": func(e ld.Entity) ld.Entity {
		return as.AsProfile(e)
	},
	"http://www.w3.org/ns/activitystreams#Question": func(e ld.Entity) ld.Entity {
		return as.AsQuestion(e)
	},
	"http://www.w3.org/ns/activitystreams#Read": func(e ld.Entity) ld.Entity {
		return as.AsRead(e)
	},
	"http://www.w3.org/ns/activitystreams#Reject": func(e ld.Entity) ld.Entity {
		return as.AsReject(e)
	},
	"http://www.w3.org/ns/activitystreams#Relationship": func(e ld.Entity) ld.Entity {
		return as.AsRelationship(e)
	},
	"http://www.w3.org/ns/activitystreams#Remove": func(e ld.Entity) ld.Entity {
		return as.AsRemove(e)
	},
	"http://www.w3.org/ns/activitystreams#Service": func(e ld.Entity) ld.Entity {
		return as.AsService(e)
	},
	"http://www.w3.org/ns/activitystreams#TentativeAccept": func(e ld.Entity) ld.Entity {
		return as.AsTentativeAccept(e)
	},
	"http://www.w3.org/ns/activitystreams#TentativeReject": func(e ld.Entity) ld.Entity {
		return as.AsTentativeReject(e)
	},
	"http://www.w3.org/ns/activitystreams#Tombstone": func(e ld.Entity) ld.Entity {
		return as.AsTombstone(e)
	},
	"http://www.w3.org/ns/activitystreams#Travel": func(e ld.Entity) ld.Entity {
		return as.AsTravel(e)
	},
	"http://www.w3.org/ns/activitystreams#Undo": func(e ld.Entity) ld.Entity {
		return as.AsUndo(e)
	},
	"http://www.w3.org/ns/activitystreams#Update": func(e ld.Entity) ld.Entity {
		return as.AsUpdate(e)
	},
	"http://www.w3.org/ns/activitystreams#Video": func(e ld.Entity) ld.Entity {
		return as.AsVideo(e)
	},
	"http://www.w3.org/ns/activitystreams#View": func(e ld.Entity) ld.Entity {
		return as.AsView(e)
	},
	"http://www.w3.org/ns/ldp#BasicContainer": func(e ld.Entity) ld.Entity {
		return ldp.AsBasicContainer(e)
	},
	"http://www.w3.org/ns/ldp#Container": func(e ld.Entity) ld.Entity {
		return ldp.AsContainer(e)
	},
	"http://www.w3.org/ns/ldp#DirectContainer": func(e ld.Entity) ld.Entity {
		return ldp.AsDirectContainer(e)
	},
	"http://www.w3.org/ns/ldp#IndirectContainer": func(e ld.Entity) ld.Entity {
		return ldp.AsIndirectContainer(e)
	},
	"http://www.w3.org/ns/ldp#NonRDFSource": func(e ld.Entity) ld.Entity {
		return ldp.AsNonRDFSource(e)
	},
	"http://www.w3.org/ns/ldp#Page": func(e ld.Entity) ld.Entity {
		return ldp.AsPage(e)
	},
	"http://www.w3.org/ns/ldp#PageSortCriterion": func(e ld.Entity) ld.Entity {
		return ldp.AsPageSortCriterion(e)
	},
	"http://www.w3.org/ns/ldp#RDFSource": func(e ld.Entity) ld.Entity {
		return ldp.AsRDFSource(e)
	},
	"http://www.w3.org/ns/ldp#Resource": func(e ld.Entity) ld.Entity {
		return ldp.AsResource(e)
	},
	"https://w3id.org/security#Digest": func(e ld.Entity) ld.Entity {
		return sec.AsDigest(e)
	},
	"https://w3id.org/security#EncryptedMessage": func(e ld.Entity) ld.Entity {
		return sec.AsEncryptedMessage(e)
	},
	"https://w3id.org/security#GraphSignature2012": func(e ld.Entity) ld.Entity {
		return sec.AsGraphSignature2012(e)
	},
	"https://w3id.org/security#Key": func(e ld.Entity) ld.Entity {
		return sec.AsKey(e)
	},
	"https://w3id.org/security#LinkedDataSignature2015": func(e ld.Entity) ld.Entity {
		return sec.AsLinkedDataSignature2015(e)
	},
	"https://w3id.org/security#LinkedDataSignature2016": func(e ld.Entity) ld.Entity {
		return sec.AsLinkedDataSignature2016(e)
	},
	"https://w3id.org/security#Signature": func(e ld.Entity) ld.Entity {
		return sec.AsSignature(e)
	},
}
