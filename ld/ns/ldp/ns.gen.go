// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package ldp

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/meta"
)

var (
	LDP = &meta.Namespace{
		ID:    "http://www.w3.org/ns/ldp#",
		Short: "ldp",
		Props: []*meta.Prop{
			Prop_ConstrainedBy,
			Prop_Contains,
			Prop_HasMemberRelation,
			Prop_Inbox,
			Prop_InsertedContentRelation,
			Prop_IsMemberOfRelation,
			Prop_Member,
			Prop_MembershipResource,
			Prop_PageSequence,
			Prop_PageSortCollation,
			Prop_PageSortCriteria,
			Prop_PageSortOrder,
			Prop_PageSortPredicate,
		},
		Types: map[string]*meta.Type{
			"http://www.w3.org/ns/ldp#BasicContainer":    Class_BasicContainer,
			"http://www.w3.org/ns/ldp#Container":         Class_Container,
			"http://www.w3.org/ns/ldp#DirectContainer":   Class_DirectContainer,
			"http://www.w3.org/ns/ldp#IndirectContainer": Class_IndirectContainer,
			"http://www.w3.org/ns/ldp#NonRDFSource":      Class_NonRDFSource,
			"http://www.w3.org/ns/ldp#Page":              Class_Page,
			"http://www.w3.org/ns/ldp#PageSortCriterion": Class_PageSortCriterion,
			"http://www.w3.org/ns/ldp#RDFSource":         Class_RDFSource,
			"http://www.w3.org/ns/ldp#Resource":          Class_Resource,
		},
	}

	// Links a resource with constraints that the server requires requests like creation and update to conform to.
	Prop_ConstrainedBy = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#constrainedBy",
		Short:   "constrainedBy",
		Comment: "Links a resource with constraints that the server requires requests like creation and update to conform to.",
	}

	// Links a container with resources created through the container.
	Prop_Contains = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#contains",
		Short:   "contains",
		Comment: "Links a container with resources created through the container.",
	}

	// Indicates which predicate is used in membership triples, and that the membership triple pattern is < membership-constant-URI , object-of-hasMemberRelation, member-URI >.
	Prop_HasMemberRelation = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#hasMemberRelation",
		Short:   "hasMemberRelation",
		Comment: "Indicates which predicate is used in membership triples, and that the membership triple pattern is < membership-constant-URI , object-of-hasMemberRelation, member-URI >.",
	}

	// Links a resource to a container where notifications for the resource can be created and discovered.
	Prop_Inbox = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#inbox",
		Short:   "inbox",
		Comment: "Links a resource to a container where notifications for the resource can be created and discovered.",
	}

	// Indicates which triple in a creation request should be used as the member-URI value in the membership triple added when the creation request is successful.
	Prop_InsertedContentRelation = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#insertedContentRelation",
		Short:   "insertedContentRelation",
		Comment: "Indicates which triple in a creation request should be used as the member-URI value in the membership triple added when the creation request is successful.",
	}

	// Indicates which predicate is used in membership triples, and that the membership triple pattern is < member-URI , object-of-isMemberOfRelation, membership-constant-URI >.
	Prop_IsMemberOfRelation = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#isMemberOfRelation",
		Short:   "isMemberOfRelation",
		Comment: "Indicates which predicate is used in membership triples, and that the membership triple pattern is < member-URI , object-of-isMemberOfRelation, membership-constant-URI >.",
	}

	// LDP servers should use this predicate as the membership predicate if there is no obvious predicate from an application vocabulary to use.
	Prop_Member = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#member",
		Short:   "member",
		Comment: "LDP servers should use this predicate as the membership predicate if there is no obvious predicate from an application vocabulary to use.",
	}

	// Indicates the membership-constant-URI in a membership triple.  Depending upon the membership triple pattern a container uses, as indicated by the presence of ldp:hasMemberRelation or ldp:isMemberOfRelation, the membership-constant-URI might occupy either the subject or object position in membership triples.
	Prop_MembershipResource = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#membershipResource",
		Short:   "membershipResource",
		Comment: "Indicates the membership-constant-URI in a membership triple.  Depending upon the membership triple pattern a container uses, as indicated by the presence of ldp:hasMemberRelation or ldp:isMemberOfRelation, the membership-constant-URI might occupy either the subject or object position in membership triples.",
	}

	// Link to a page sequence resource, as defined by LDP Paging.  Typically used to communicate the sorting criteria used to allocate LDPC members to pages.
	Prop_PageSequence = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#pageSequence",
		Short:   "pageSequence",
		Comment: "Link to a page sequence resource, as defined by LDP Paging.  Typically used to communicate the sorting criteria used to allocate LDPC members to pages.",
	}

	// The collation used to order the members across pages in a page sequence when comparing strings.
	Prop_PageSortCollation = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#pageSortCollation",
		Short:   "pageSortCollation",
		Comment: "The collation used to order the members across pages in a page sequence when comparing strings.",
	}

	// Link to the list of sorting criteria used by the server in a representation.  Typically used on Link response headers as an extension link relation URI in the rel= parameter.
	Prop_PageSortCriteria = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#pageSortCriteria",
		Short:   "pageSortCriteria",
		Comment: "Link to the list of sorting criteria used by the server in a representation.  Typically used on Link response headers as an extension link relation URI in the rel= parameter.",
	}

	// The ascending/descending/etc order used to order the members across pages in a page sequence.
	Prop_PageSortOrder = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#pageSortOrder",
		Short:   "pageSortOrder",
		Comment: "The ascending/descending/etc order used to order the members across pages in a page sequence.",
	}

	// Predicate used to specify the order of the members across a page sequence's in-sequence page resources; it asserts nothing about the order of members in the representation of a single page.
	Prop_PageSortPredicate = &meta.Prop{
		ID:      "http://www.w3.org/ns/ldp#pageSortPredicate",
		Short:   "pageSortPredicate",
		Comment: "Predicate used to specify the order of the members across a page sequence's in-sequence page resources; it asserts nothing about the order of members in the representation of a single page.",
	}

	// An LDPC that uses a predefined predicate to simply link to its contained resources.
	Class_BasicContainer = &meta.Type{
		ID:    "http://www.w3.org/ns/ldp#BasicContainer",
		Short: "BasicContainer",
		Cast:  func(e ld.Entity) ld.Entity { return AsBasicContainer(e) },
		SubClassOf: []*meta.Type{
			Class_Container,
		},
		Props: []*meta.Prop{},
	}

	// A Linked Data Platform RDF Source (LDP-RS) that also conforms to additional patterns and conventions for managing membership. Readers should refer to the specification defining this ontology for the list of behaviors associated with it.
	Class_Container = &meta.Type{
		ID:    "http://www.w3.org/ns/ldp#Container",
		Short: "Container",
		Cast:  func(e ld.Entity) ld.Entity { return AsContainer(e) },
		SubClassOf: []*meta.Type{
			Class_RDFSource,
		},
		Props: []*meta.Prop{
			Prop_Contains,
			Prop_HasMemberRelation,
			Prop_InsertedContentRelation,
			Prop_IsMemberOfRelation,
			Prop_MembershipResource,
		},
	}

	// An LDPC that is similar to a LDP-DC but it allows an indirection with the ability to list as member a resource, such as a URI representing a real-world object, that is different from the resource that is created.
	Class_DirectContainer = &meta.Type{
		ID:    "http://www.w3.org/ns/ldp#DirectContainer",
		Short: "DirectContainer",
		Cast:  func(e ld.Entity) ld.Entity { return AsDirectContainer(e) },
		SubClassOf: []*meta.Type{
			Class_Container,
		},
		Props: []*meta.Prop{},
	}

	// An LDPC that has the flexibility of choosing what form the membership triples take.
	Class_IndirectContainer = &meta.Type{
		ID:    "http://www.w3.org/ns/ldp#IndirectContainer",
		Short: "IndirectContainer",
		Cast:  func(e ld.Entity) ld.Entity { return AsIndirectContainer(e) },
		SubClassOf: []*meta.Type{
			Class_Container,
		},
		Props: []*meta.Prop{},
	}

	// A Linked Data Platform Resource (LDPR) whose state is NOT represented as RDF.
	Class_NonRDFSource = &meta.Type{
		ID:    "http://www.w3.org/ns/ldp#NonRDFSource",
		Short: "NonRDFSource",
		Cast:  func(e ld.Entity) ld.Entity { return AsNonRDFSource(e) },
		SubClassOf: []*meta.Type{
			Class_Resource,
		},
		Props: []*meta.Prop{},
	}

	// URI signifying that the resource is an in-sequence page resource, as defined by LDP Paging.  Typically used on Link rel='type' response headers.
	Class_Page = &meta.Type{
		ID:         "http://www.w3.org/ns/ldp#Page",
		Short:      "Page",
		Cast:       func(e ld.Entity) ld.Entity { return AsPage(e) },
		SubClassOf: []*meta.Type{},
		Props: []*meta.Prop{
			Prop_Inbox,
			Prop_PageSequence,
			Prop_PageSortCriteria,
		},
	}

	// Element in the list of sorting criteria used by the server to assign container members to pages.
	Class_PageSortCriterion = &meta.Type{
		ID:         "http://www.w3.org/ns/ldp#PageSortCriterion",
		Short:      "PageSortCriterion",
		Cast:       func(e ld.Entity) ld.Entity { return AsPageSortCriterion(e) },
		SubClassOf: []*meta.Type{},
		Props: []*meta.Prop{
			Prop_Inbox,
			Prop_PageSequence,
			Prop_PageSortCollation,
			Prop_PageSortOrder,
			Prop_PageSortPredicate,
		},
	}

	// A Linked Data Platform Resource (LDPR) whose state is represented as RDF.
	Class_RDFSource = &meta.Type{
		ID:    "http://www.w3.org/ns/ldp#RDFSource",
		Short: "RDFSource",
		Cast:  func(e ld.Entity) ld.Entity { return AsRDFSource(e) },
		SubClassOf: []*meta.Type{
			Class_Resource,
		},
		Props: []*meta.Prop{},
	}

	// A HTTP-addressable resource whose lifecycle is managed by a LDP server.
	Class_Resource = &meta.Type{
		ID:         "http://www.w3.org/ns/ldp#Resource",
		Short:      "Resource",
		Cast:       func(e ld.Entity) ld.Entity { return AsResource(e) },
		SubClassOf: []*meta.Type{},
		Props: []*meta.Prop{
			Prop_ConstrainedBy,
			Prop_Inbox,
			Prop_Member,
			Prop_PageSequence,
		},
	}
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
