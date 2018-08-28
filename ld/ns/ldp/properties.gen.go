// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package ldp

import (
	"github.com/meowpub/meow/ld"
)

// Links a resource with constraints that the server requires requests like creation and update to conform to.
func GetConstrainedBy(e ld.Entity) interface{} { return e.Get(PropConstrainedBy) }

func SetConstrainedBy(e ld.Entity, v interface{}) { e.Set(PropConstrainedBy, v) }

// Links a container with resources created through the container.
func GetContains(e ld.Entity) interface{} { return e.Get(PropContains) }

func SetContains(e ld.Entity, v interface{}) { e.Set(PropContains, v) }

// Indicates which predicate is used in membership triples, and that the membership triple pattern is < membership-constant-URI , object-of-hasMemberRelation, member-URI >.
func GetHasMemberRelation(e ld.Entity) interface{} { return e.Get(PropHasMemberRelation) }

func SetHasMemberRelation(e ld.Entity, v interface{}) { e.Set(PropHasMemberRelation, v) }

// Links a resource to a container where notifications for the resource can be created and discovered.
func GetInbox(e ld.Entity) interface{} { return e.Get(PropInbox) }

func SetInbox(e ld.Entity, v interface{}) { e.Set(PropInbox, v) }

// Indicates which triple in a creation request should be used as the member-URI value in the membership triple added when the creation request is successful.
func GetInsertedContentRelation(e ld.Entity) interface{} { return e.Get(PropInsertedContentRelation) }

func SetInsertedContentRelation(e ld.Entity, v interface{}) { e.Set(PropInsertedContentRelation, v) }

// Indicates which predicate is used in membership triples, and that the membership triple pattern is < member-URI , object-of-isMemberOfRelation, membership-constant-URI >.
func GetIsMemberOfRelation(e ld.Entity) interface{} { return e.Get(PropIsMemberOfRelation) }

func SetIsMemberOfRelation(e ld.Entity, v interface{}) { e.Set(PropIsMemberOfRelation, v) }

// LDP servers should use this predicate as the membership predicate if there is no obvious predicate from an application vocabulary to use.
func GetMember(e ld.Entity) interface{} { return e.Get(PropMember) }

func SetMember(e ld.Entity, v interface{}) { e.Set(PropMember, v) }

// Indicates the membership-constant-URI in a membership triple.  Depending upon the membership triple pattern a container uses, as indicated by the presence of ldp:hasMemberRelation or ldp:isMemberOfRelation, the membership-constant-URI might occupy either the subject or object position in membership triples.
func GetMembershipResource(e ld.Entity) interface{} { return e.Get(PropMembershipResource) }

func SetMembershipResource(e ld.Entity, v interface{}) { e.Set(PropMembershipResource, v) }

// Link to a page sequence resource, as defined by LDP Paging.  Typically used to communicate the sorting criteria used to allocate LDPC members to pages.
func GetPageSequence(e ld.Entity) interface{} { return e.Get(PropPageSequence) }

func SetPageSequence(e ld.Entity, v interface{}) { e.Set(PropPageSequence, v) }

// The collation used to order the members across pages in a page sequence when comparing strings.
func GetPageSortCollation(e ld.Entity) interface{} { return e.Get(PropPageSortCollation) }

func SetPageSortCollation(e ld.Entity, v interface{}) { e.Set(PropPageSortCollation, v) }

// Link to the list of sorting criteria used by the server in a representation.  Typically used on Link response headers as an extension link relation URI in the rel= parameter.
func GetPageSortCriteria(e ld.Entity) interface{} { return e.Get(PropPageSortCriteria) }

func SetPageSortCriteria(e ld.Entity, v interface{}) { e.Set(PropPageSortCriteria, v) }

// The ascending/descending/etc order used to order the members across pages in a page sequence.
func GetPageSortOrder(e ld.Entity) interface{} { return e.Get(PropPageSortOrder) }

func SetPageSortOrder(e ld.Entity, v interface{}) { e.Set(PropPageSortOrder, v) }

// Predicate used to specify the order of the members across a page sequence's in-sequence page resources; it asserts nothing about the order of members in the representation of a single page.
func GetPageSortPredicate(e ld.Entity) interface{} { return e.Get(PropPageSortPredicate) }

func SetPageSortPredicate(e ld.Entity, v interface{}) { e.Set(PropPageSortPredicate, v) }
