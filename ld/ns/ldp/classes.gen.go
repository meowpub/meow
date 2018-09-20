// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package ldp

import (
	"github.com/meowpub/meow/ld"
)

// An LDPC that uses a predefined predicate to simply link to its contained resources.
type BasicContainer struct{ Container }

// Ducktypes the object into a(n) BasicContainer.
func AsBasicContainer(e ld.Entity) BasicContainer { return BasicContainer{AsContainer(e)} }

// Does the object quack like a(n) BasicContainer?
func IsBasicContainer(e ld.Entity) bool { return ld.Is(e, TypeBasicContainer) }

// A Linked Data Platform RDF Source (LDP-RS) that also conforms to additional patterns and conventions for managing membership. Readers should refer to the specification defining this ontology for the list of behaviors associated with it.
type Container struct{ RDFSource }

// Ducktypes the object into a(n) Container.
func AsContainer(e ld.Entity) Container { return Container{AsRDFSource(e)} }

// Does the object quack like a(n) Container?
func IsContainer(e ld.Entity) bool { return ld.Is(e, TypeContainer) }

// Links a container with resources created through the container.
func (obj Container) Contains() interface{} { return GetContains(obj) }

func (obj Container) SetContains(v interface{}) { SetContains(obj, v) }

// Indicates which predicate is used in membership triples, and that the membership triple pattern is < membership-constant-URI , object-of-hasMemberRelation, member-URI >.
func (obj Container) HasMemberRelation() interface{} { return GetHasMemberRelation(obj) }

func (obj Container) SetHasMemberRelation(v interface{}) { SetHasMemberRelation(obj, v) }

// Indicates which triple in a creation request should be used as the member-URI value in the membership triple added when the creation request is successful.
func (obj Container) InsertedContentRelation() interface{} { return GetInsertedContentRelation(obj) }

func (obj Container) SetInsertedContentRelation(v interface{}) { SetInsertedContentRelation(obj, v) }

// Indicates which predicate is used in membership triples, and that the membership triple pattern is < member-URI , object-of-isMemberOfRelation, membership-constant-URI >.
func (obj Container) IsMemberOfRelation() interface{} { return GetIsMemberOfRelation(obj) }

func (obj Container) SetIsMemberOfRelation(v interface{}) { SetIsMemberOfRelation(obj, v) }

// Indicates the membership-constant-URI in a membership triple.  Depending upon the membership triple pattern a container uses, as indicated by the presence of ldp:hasMemberRelation or ldp:isMemberOfRelation, the membership-constant-URI might occupy either the subject or object position in membership triples.
func (obj Container) MembershipResource() interface{} { return GetMembershipResource(obj) }

func (obj Container) SetMembershipResource(v interface{}) { SetMembershipResource(obj, v) }

// An LDPC that is similar to a LDP-DC but it allows an indirection with the ability to list as member a resource, such as a URI representing a real-world object, that is different from the resource that is created.
type DirectContainer struct{ Container }

// Ducktypes the object into a(n) DirectContainer.
func AsDirectContainer(e ld.Entity) DirectContainer { return DirectContainer{AsContainer(e)} }

// Does the object quack like a(n) DirectContainer?
func IsDirectContainer(e ld.Entity) bool { return ld.Is(e, TypeDirectContainer) }

// An LDPC that has the flexibility of choosing what form the membership triples take.
type IndirectContainer struct{ Container }

// Ducktypes the object into a(n) IndirectContainer.
func AsIndirectContainer(e ld.Entity) IndirectContainer { return IndirectContainer{AsContainer(e)} }

// Does the object quack like a(n) IndirectContainer?
func IsIndirectContainer(e ld.Entity) bool { return ld.Is(e, TypeIndirectContainer) }

// A Linked Data Platform Resource (LDPR) whose state is NOT represented as RDF.
type NonRDFSource struct{ Resource }

// Ducktypes the object into a(n) NonRDFSource.
func AsNonRDFSource(e ld.Entity) NonRDFSource { return NonRDFSource{AsResource(e)} }

// Does the object quack like a(n) NonRDFSource?
func IsNonRDFSource(e ld.Entity) bool { return ld.Is(e, TypeNonRDFSource) }

// URI signifying that the resource is an in-sequence page resource, as defined by LDP Paging.  Typically used on Link rel='type' response headers.
type Page struct{ o *ld.Object }

// Ducktypes the object into a(n) Page.
func AsPage(e ld.Entity) Page { return Page{o: e.Obj()} }

// Does the object quack like a(n) Page?
func IsPage(e ld.Entity) bool { return ld.Is(e, TypePage) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Page) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj Page) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Page) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Page) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Page) Get(key string) interface{} { return obj.o.Get(key) }

// Sets the named attribute. Implements ld.Entity.
func (obj Page) Set(key string, v interface{}) { obj.o.Set(key, v) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Page) Apply(other ld.Entity, mergeArrays bool) error { return obj.o.Apply(other, mergeArrays) }

// Links a resource to a container where notifications for the resource can be created and discovered.
func (obj Page) Inbox() interface{} { return GetInbox(obj) }

func (obj Page) SetInbox(v interface{}) { SetInbox(obj, v) }

// Link to a page sequence resource, as defined by LDP Paging.  Typically used to communicate the sorting criteria used to allocate LDPC members to pages.
func (obj Page) PageSequence() interface{} { return GetPageSequence(obj) }

func (obj Page) SetPageSequence(v interface{}) { SetPageSequence(obj, v) }

// Link to the list of sorting criteria used by the server in a representation.  Typically used on Link response headers as an extension link relation URI in the rel= parameter.
func (obj Page) PageSortCriteria() interface{} { return GetPageSortCriteria(obj) }

func (obj Page) SetPageSortCriteria(v interface{}) { SetPageSortCriteria(obj, v) }

// Element in the list of sorting criteria used by the server to assign container members to pages.
type PageSortCriterion struct{ o *ld.Object }

// Ducktypes the object into a(n) PageSortCriterion.
func AsPageSortCriterion(e ld.Entity) PageSortCriterion { return PageSortCriterion{o: e.Obj()} }

// Does the object quack like a(n) PageSortCriterion?
func IsPageSortCriterion(e ld.Entity) bool { return ld.Is(e, TypePageSortCriterion) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj PageSortCriterion) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj PageSortCriterion) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj PageSortCriterion) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj PageSortCriterion) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj PageSortCriterion) Get(key string) interface{} { return obj.o.Get(key) }

// Sets the named attribute. Implements ld.Entity.
func (obj PageSortCriterion) Set(key string, v interface{}) { obj.o.Set(key, v) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj PageSortCriterion) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// Links a resource to a container where notifications for the resource can be created and discovered.
func (obj PageSortCriterion) Inbox() interface{} { return GetInbox(obj) }

func (obj PageSortCriterion) SetInbox(v interface{}) { SetInbox(obj, v) }

// Link to a page sequence resource, as defined by LDP Paging.  Typically used to communicate the sorting criteria used to allocate LDPC members to pages.
func (obj PageSortCriterion) PageSequence() interface{} { return GetPageSequence(obj) }

func (obj PageSortCriterion) SetPageSequence(v interface{}) { SetPageSequence(obj, v) }

// The collation used to order the members across pages in a page sequence when comparing strings.
func (obj PageSortCriterion) PageSortCollation() interface{} { return GetPageSortCollation(obj) }

func (obj PageSortCriterion) SetPageSortCollation(v interface{}) { SetPageSortCollation(obj, v) }

// The ascending/descending/etc order used to order the members across pages in a page sequence.
func (obj PageSortCriterion) PageSortOrder() interface{} { return GetPageSortOrder(obj) }

func (obj PageSortCriterion) SetPageSortOrder(v interface{}) { SetPageSortOrder(obj, v) }

// Predicate used to specify the order of the members across a page sequence's in-sequence page resources; it asserts nothing about the order of members in the representation of a single page.
func (obj PageSortCriterion) PageSortPredicate() interface{} { return GetPageSortPredicate(obj) }

func (obj PageSortCriterion) SetPageSortPredicate(v interface{}) { SetPageSortPredicate(obj, v) }

// A Linked Data Platform Resource (LDPR) whose state is represented as RDF.
type RDFSource struct{ Resource }

// Ducktypes the object into a(n) RDFSource.
func AsRDFSource(e ld.Entity) RDFSource { return RDFSource{AsResource(e)} }

// Does the object quack like a(n) RDFSource?
func IsRDFSource(e ld.Entity) bool { return ld.Is(e, TypeRDFSource) }

// A HTTP-addressable resource whose lifecycle is managed by a LDP server.
type Resource struct{ o *ld.Object }

// Ducktypes the object into a(n) Resource.
func AsResource(e ld.Entity) Resource { return Resource{o: e.Obj()} }

// Does the object quack like a(n) Resource?
func IsResource(e ld.Entity) bool { return ld.Is(e, TypeResource) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Resource) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj Resource) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Resource) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Resource) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Resource) Get(key string) interface{} { return obj.o.Get(key) }

// Sets the named attribute. Implements ld.Entity.
func (obj Resource) Set(key string, v interface{}) { obj.o.Set(key, v) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Resource) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// Links a resource with constraints that the server requires requests like creation and update to conform to.
func (obj Resource) ConstrainedBy() interface{} { return GetConstrainedBy(obj) }

func (obj Resource) SetConstrainedBy(v interface{}) { SetConstrainedBy(obj, v) }

// Links a resource to a container where notifications for the resource can be created and discovered.
func (obj Resource) Inbox() interface{} { return GetInbox(obj) }

func (obj Resource) SetInbox(v interface{}) { SetInbox(obj, v) }

// LDP servers should use this predicate as the membership predicate if there is no obvious predicate from an application vocabulary to use.
func (obj Resource) Member() interface{} { return GetMember(obj) }

func (obj Resource) SetMember(v interface{}) { SetMember(obj, v) }

// Link to a page sequence resource, as defined by LDP Paging.  Typically used to communicate the sorting criteria used to allocate LDPC members to pages.
func (obj Resource) PageSequence() interface{} { return GetPageSequence(obj) }

func (obj Resource) SetPageSequence(v interface{}) { SetPageSequence(obj, v) }

var (
	_ ld.Entity = BasicContainer{}
	_ ld.Entity = Container{}
	_ ld.Entity = DirectContainer{}
	_ ld.Entity = IndirectContainer{}
	_ ld.Entity = NonRDFSource{}
	_ ld.Entity = Page{}
	_ ld.Entity = PageSortCriterion{}
	_ ld.Entity = RDFSource{}
	_ ld.Entity = Resource{}
)
