// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package ldp

const (
	// Links a resource with constraints that the server requires requests like creation and update to conform to.
	PropConstrainedBy = "http://www.w3.org/ns/ldp#constrainedBy"

	// Links a container with resources created through the container.
	PropContains = "http://www.w3.org/ns/ldp#contains"

	// Indicates which predicate is used in membership triples, and that the membership triple pattern is < membership-constant-URI , object-of-hasMemberRelation, member-URI >.
	PropHasMemberRelation = "http://www.w3.org/ns/ldp#hasMemberRelation"

	// Links a resource to a container where notifications for the resource can be created and discovered.
	PropInbox = "http://www.w3.org/ns/ldp#inbox"

	// Indicates which triple in a creation request should be used as the member-URI value in the membership triple added when the creation request is successful.
	PropInsertedContentRelation = "http://www.w3.org/ns/ldp#insertedContentRelation"

	// Indicates which predicate is used in membership triples, and that the membership triple pattern is < member-URI , object-of-isMemberOfRelation, membership-constant-URI >.
	PropIsMemberOfRelation = "http://www.w3.org/ns/ldp#isMemberOfRelation"

	// LDP servers should use this predicate as the membership predicate if there is no obvious predicate from an application vocabulary to use.
	PropMember = "http://www.w3.org/ns/ldp#member"

	// Indicates the membership-constant-URI in a membership triple.  Depending upon the membership triple pattern a container uses, as indicated by the presence of ldp:hasMemberRelation or ldp:isMemberOfRelation, the membership-constant-URI might occupy either the subject or object position in membership triples.
	PropMembershipResource = "http://www.w3.org/ns/ldp#membershipResource"

	// Link to a page sequence resource, as defined by LDP Paging.  Typically used to communicate the sorting criteria used to allocate LDPC members to pages.
	PropPageSequence = "http://www.w3.org/ns/ldp#pageSequence"

	// The collation used to order the members across pages in a page sequence when comparing strings.
	PropPageSortCollation = "http://www.w3.org/ns/ldp#pageSortCollation"

	// Link to the list of sorting criteria used by the server in a representation.  Typically used on Link response headers as an extension link relation URI in the rel= parameter.
	PropPageSortCriteria = "http://www.w3.org/ns/ldp#pageSortCriteria"

	// The ascending/descending/etc order used to order the members across pages in a page sequence.
	PropPageSortOrder = "http://www.w3.org/ns/ldp#pageSortOrder"

	// Predicate used to specify the order of the members across a page sequence's in-sequence page resources; it asserts nothing about the order of members in the representation of a single page.
	PropPageSortPredicate = "http://www.w3.org/ns/ldp#pageSortPredicate"
)

const (
	// An LDPC that uses a predefined predicate to simply link to its contained resources.
	TypeBasicContainer = "http://www.w3.org/ns/ldp#BasicContainer"

	// A Linked Data Platform RDF Source (LDP-RS) that also conforms to additional patterns and conventions for managing membership. Readers should refer to the specification defining this ontology for the list of behaviors associated with it.
	TypeContainer = "http://www.w3.org/ns/ldp#Container"

	// An LDPC that is similar to a LDP-DC but it allows an indirection with the ability to list as member a resource, such as a URI representing a real-world object, that is different from the resource that is created.
	TypeDirectContainer = "http://www.w3.org/ns/ldp#DirectContainer"

	// An LDPC that has the flexibility of choosing what form the membership triples take.
	TypeIndirectContainer = "http://www.w3.org/ns/ldp#IndirectContainer"

	// A Linked Data Platform Resource (LDPR) whose state is NOT represented as RDF.
	TypeNonRDFSource = "http://www.w3.org/ns/ldp#NonRDFSource"

	// URI signifying that the resource is an in-sequence page resource, as defined by LDP Paging.  Typically used on Link rel='type' response headers.
	TypePage = "http://www.w3.org/ns/ldp#Page"

	// Element in the list of sorting criteria used by the server to assign container members to pages.
	TypePageSortCriterion = "http://www.w3.org/ns/ldp#PageSortCriterion"

	// A Linked Data Platform Resource (LDPR) whose state is represented as RDF.
	TypeRDFSource = "http://www.w3.org/ns/ldp#RDFSource"

	// A HTTP-addressable resource whose lifecycle is managed by a LDP server.
	TypeResource = "http://www.w3.org/ns/ldp#Resource"
)

const (
	// Ascending order.
	Ascending = "http://www.w3.org/ns/ldp#Ascending"

	// Descending order.
	Descending = "http://www.w3.org/ns/ldp#Descending"

	// Used to indicate default and typical behavior for ldp:insertedContentRelation, where the member-URI value in the membership triple added when a creation request is successful is the URI assigned to the newly created resource.
	MemberSubject = "http://www.w3.org/ns/ldp#MemberSubject"

	// URI identifying a LDPC's containment triples, for example to allow clients to express interest in receiving them.
	PreferContainment = "http://www.w3.org/ns/ldp#PreferContainment"

	// Archaic alias for ldp:PreferMinimalContainer
	PreferEmptyContainer = "http://www.w3.org/ns/ldp#PreferEmptyContainer"

	// URI identifying a LDPC's membership triples, for example to allow clients to express interest in receiving them.
	PreferMembership = "http://www.w3.org/ns/ldp#PreferMembership"

	// URI identifying the subset of a LDPC's triples present in an empty LDPC, for example to allow clients to express interest in receiving them.  Currently this excludes containment and membership triples, but in the future other exclusions might be added.  This definition is written to automatically exclude those new classes of triples.
	PreferMinimalContainer = "http://www.w3.org/ns/ldp#PreferMinimalContainer"
)
