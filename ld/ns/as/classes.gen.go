// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package as

import (
	"github.com/meowpub/meow/ld"
)

// Actor accepts the Object
type Accept struct{ Activity }

// Ducktypes the object into a(n) Accept.
func AsAccept(e ld.Entity) Accept { return Accept{AsActivity(e)} }

// Does the object quack like a(n) Accept?
func IsAccept(e ld.Entity) bool { return ld.Is(e, TypeAccept) }

// An Object representing some form of Action that has been taken
type Activity struct{ Object }

// Ducktypes the object into a(n) Activity.
func AsActivity(e ld.Entity) Activity { return Activity{AsObject(e)} }

// Does the object quack like a(n) Activity?
func IsActivity(e ld.Entity) bool { return ld.Is(e, TypeActivity) }

// Subproperty of as:attributedTo that identifies the primary actor
func (obj Activity) Actor() interface{} { return GetActor(obj) }

// Indentifies an object used (or to be used) to complete an activity
func (obj Activity) Instrument() interface{} { return GetInstrument(obj) }

// For certain activities, specifies the entity from which the action is directed.
func (obj Activity) Origin() interface{} { return GetOrigin(obj) }

func (obj Activity) Result() interface{} { return GetResult(obj) }

func (obj Activity) Target() interface{} { return GetTarget(obj) }

func (obj Activity) Verb() interface{} { return GetVerb(obj) }

// To Add an Object or Link to Something
type Add struct{ Activity }

// Ducktypes the object into a(n) Add.
func AsAdd(e ld.Entity) Add { return Add{AsActivity(e)} }

// Does the object quack like a(n) Add?
func IsAdd(e ld.Entity) bool { return ld.Is(e, TypeAdd) }

// Actor announces the object to the target
type Announce struct{ Activity }

// Ducktypes the object into a(n) Announce.
func AsAnnounce(e ld.Entity) Announce { return Announce{AsActivity(e)} }

// Does the object quack like a(n) Announce?
func IsAnnounce(e ld.Entity) bool { return ld.Is(e, TypeAnnounce) }

// Represents a software application of any sort
type Application struct{ Object }

// Ducktypes the object into a(n) Application.
func AsApplication(e ld.Entity) Application { return Application{AsObject(e)} }

// Does the object quack like a(n) Application?
func IsApplication(e ld.Entity) bool { return ld.Is(e, TypeApplication) }

// To Arrive Somewhere (can be used, for instance, to indicate that a particular entity is currently located somewhere, e.g. a "check-in")
type Arrive struct{ IntransitiveActivity }

// Ducktypes the object into a(n) Arrive.
func AsArrive(e ld.Entity) Arrive { return Arrive{AsIntransitiveActivity(e)} }

// Does the object quack like a(n) Arrive?
func IsArrive(e ld.Entity) bool { return ld.Is(e, TypeArrive) }

// A written work. Typically several paragraphs long. For example, a blog post or a news article.
type Article struct{ Object }

// Ducktypes the object into a(n) Article.
func AsArticle(e ld.Entity) Article { return Article{AsObject(e)} }

// Does the object quack like a(n) Article?
func IsArticle(e ld.Entity) bool { return ld.Is(e, TypeArticle) }

// An audio file
type Audio struct{ Document }

// Ducktypes the object into a(n) Audio.
func AsAudio(e ld.Entity) Audio { return Audio{AsDocument(e)} }

// Does the object quack like a(n) Audio?
func IsAudio(e ld.Entity) bool { return ld.Is(e, TypeAudio) }

type Block struct{ Ignore }

// Ducktypes the object into a(n) Block.
func AsBlock(e ld.Entity) Block { return Block{AsIgnore(e)} }

// Does the object quack like a(n) Block?
func IsBlock(e ld.Entity) bool { return ld.Is(e, TypeBlock) }

// An ordered or unordered collection of Objects or Links
type Collection struct{ Object }

// Ducktypes the object into a(n) Collection.
func AsCollection(e ld.Entity) Collection { return Collection{AsObject(e)} }

// Does the object quack like a(n) Collection?
func IsCollection(e ld.Entity) bool { return ld.Is(e, TypeCollection) }

func (obj Collection) Current() interface{} { return GetCurrent(obj) }

func (obj Collection) First() interface{} { return GetFirst(obj) }

func (obj Collection) Items() interface{} { return GetItems(obj) }

func (obj Collection) Last() interface{} { return GetLast(obj) }

// The total number of items in a logical collection
func (obj Collection) TotalItems() interface{} { return GetTotalItems(obj) }

// A subset of items from a Collection
type CollectionPage struct{ Collection }

// Ducktypes the object into a(n) CollectionPage.
func AsCollectionPage(e ld.Entity) CollectionPage { return CollectionPage{AsCollection(e)} }

// Does the object quack like a(n) CollectionPage?
func IsCollectionPage(e ld.Entity) bool { return ld.Is(e, TypeCollectionPage) }

func (obj CollectionPage) Next() interface{} { return GetNext(obj) }

func (obj CollectionPage) PartOf() interface{} { return GetPartOf(obj) }

func (obj CollectionPage) Prev() interface{} { return GetPrev(obj) }

// To Create Something
type Create struct{ Activity }

// Ducktypes the object into a(n) Create.
func AsCreate(e ld.Entity) Create { return Create{AsActivity(e)} }

// Does the object quack like a(n) Create?
func IsCreate(e ld.Entity) bool { return ld.Is(e, TypeCreate) }

// To Delete Something
type Delete struct{ Activity }

// Ducktypes the object into a(n) Delete.
func AsDelete(e ld.Entity) Delete { return Delete{AsActivity(e)} }

// Does the object quack like a(n) Delete?
func IsDelete(e ld.Entity) bool { return ld.Is(e, TypeDelete) }

// The actor dislikes the object
type Dislike struct{ Activity }

// Ducktypes the object into a(n) Dislike.
func AsDislike(e ld.Entity) Dislike { return Dislike{AsActivity(e)} }

// Does the object quack like a(n) Dislike?
func IsDislike(e ld.Entity) bool { return ld.Is(e, TypeDislike) }

// Represents a digital document/file of any sort
type Document struct{ Object }

// Ducktypes the object into a(n) Document.
func AsDocument(e ld.Entity) Document { return Document{AsObject(e)} }

// Does the object quack like a(n) Document?
func IsDocument(e ld.Entity) bool { return ld.Is(e, TypeDocument) }

// An Event of any kind
type Event struct{ Object }

// Ducktypes the object into a(n) Event.
func AsEvent(e ld.Entity) Event { return Event{AsObject(e)} }

// Does the object quack like a(n) Event?
func IsEvent(e ld.Entity) bool { return ld.Is(e, TypeEvent) }

// To flag something (e.g. flag as inappropriate, flag as spam, etc)
type Flag struct{ Activity }

// Ducktypes the object into a(n) Flag.
func AsFlag(e ld.Entity) Flag { return Flag{AsActivity(e)} }

// Does the object quack like a(n) Flag?
func IsFlag(e ld.Entity) bool { return ld.Is(e, TypeFlag) }

// To Express Interest in Something
type Follow struct{ Activity }

// Ducktypes the object into a(n) Follow.
func AsFollow(e ld.Entity) Follow { return Follow{AsActivity(e)} }

// Does the object quack like a(n) Follow?
func IsFollow(e ld.Entity) bool { return ld.Is(e, TypeFollow) }

// A Group of any kind.
type Group struct{ Object }

// Ducktypes the object into a(n) Group.
func AsGroup(e ld.Entity) Group { return Group{AsObject(e)} }

// Does the object quack like a(n) Group?
func IsGroup(e ld.Entity) bool { return ld.Is(e, TypeGroup) }

// Actor is ignoring the Object
type Ignore struct{ Activity }

// Ducktypes the object into a(n) Ignore.
func AsIgnore(e ld.Entity) Ignore { return Ignore{AsActivity(e)} }

// Does the object quack like a(n) Ignore?
func IsIgnore(e ld.Entity) bool { return ld.Is(e, TypeIgnore) }

// An Image file
type Image struct{ Document }

// Ducktypes the object into a(n) Image.
func AsImage(e ld.Entity) Image { return Image{AsDocument(e)} }

// Does the object quack like a(n) Image?
func IsImage(e ld.Entity) bool { return ld.Is(e, TypeImage) }

// An Activity that has no direct object
type IntransitiveActivity struct{ Activity }

// Ducktypes the object into a(n) IntransitiveActivity.
func AsIntransitiveActivity(e ld.Entity) IntransitiveActivity {
	return IntransitiveActivity{AsActivity(e)}
}

// Does the object quack like a(n) IntransitiveActivity?
func IsIntransitiveActivity(e ld.Entity) bool { return ld.Is(e, TypeIntransitiveActivity) }

// To invite someone or something to something
type Invite struct{ Offer }

// Ducktypes the object into a(n) Invite.
func AsInvite(e ld.Entity) Invite { return Invite{AsOffer(e)} }

// Does the object quack like a(n) Invite?
func IsInvite(e ld.Entity) bool { return ld.Is(e, TypeInvite) }

// To Join Something
type Join struct{ Activity }

// Ducktypes the object into a(n) Join.
func AsJoin(e ld.Entity) Join { return Join{AsActivity(e)} }

// Does the object quack like a(n) Join?
func IsJoin(e ld.Entity) bool { return ld.Is(e, TypeJoin) }

// To Leave Something
type Leave struct{ Activity }

// Ducktypes the object into a(n) Leave.
func AsLeave(e ld.Entity) Leave { return Leave{AsActivity(e)} }

// Does the object quack like a(n) Leave?
func IsLeave(e ld.Entity) bool { return ld.Is(e, TypeLeave) }

// To Like Something
type Like struct{ Activity }

// Ducktypes the object into a(n) Like.
func AsLike(e ld.Entity) Like { return Like{AsActivity(e)} }

// Does the object quack like a(n) Like?
func IsLike(e ld.Entity) bool { return ld.Is(e, TypeLike) }

// Represents a qualified reference to another resource. Patterned after the RFC5988 Web Linking Model
type Link struct{ o *ld.Object }

// Ducktypes the object into a(n) Link.
func AsLink(e ld.Entity) Link { return Link{o: e.Obj()} }

// Does the object quack like a(n) Link?
func IsLink(e ld.Entity) bool { return ld.Is(e, TypeLink) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Link) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj Link) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Link) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Link) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Link) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Link) Apply(other ld.Entity, mergeArrays bool) error { return obj.o.Apply(other, mergeArrays) }

// The display height expressed as device independent pixels
func (obj Link) Height() interface{} { return GetHeight(obj) }

// The target URI of the Link
func (obj Link) Href() interface{} { return GetHref(obj) }

// A hint about the language of the referenced resource
func (obj Link) Hreflang() interface{} { return GetHreflang(obj) }

// The RFC 5988 or HTML5 Link Relation associated with the Link
func (obj Link) Rel() interface{} { return GetRel(obj) }

// Specifies the preferred display width of the content, expressed in terms of device independent pixels.
func (obj Link) Width() interface{} { return GetWidth(obj) }

// The actor listened to the object
type Listen struct{ Activity }

// Ducktypes the object into a(n) Listen.
func AsListen(e ld.Entity) Listen { return Listen{AsActivity(e)} }

// Does the object quack like a(n) Listen?
func IsListen(e ld.Entity) bool { return ld.Is(e, TypeListen) }

// A specialized Link that represents an @mention
type Mention struct{ Link }

// Ducktypes the object into a(n) Mention.
func AsMention(e ld.Entity) Mention { return Mention{AsLink(e)} }

// Does the object quack like a(n) Mention?
func IsMention(e ld.Entity) bool { return ld.Is(e, TypeMention) }

// The actor is moving the object. The target specifies where the object is moving to. The origin specifies where the object is moving from.
type Move struct{ Activity }

// Ducktypes the object into a(n) Move.
func AsMove(e ld.Entity) Move { return Move{AsActivity(e)} }

// Does the object quack like a(n) Move?
func IsMove(e ld.Entity) bool { return ld.Is(e, TypeMove) }

// A Short note, typically less than a single paragraph. A "tweet" is an example, or a "status update"
type Note struct{ Object }

// Ducktypes the object into a(n) Note.
func AsNote(e ld.Entity) Note { return Note{AsObject(e)} }

// Does the object quack like a(n) Note?
func IsNote(e ld.Entity) bool { return ld.Is(e, TypeNote) }

type Object struct{ o *ld.Object }

// Ducktypes the object into a(n) Object.
func AsObject(e ld.Entity) Object { return Object{o: e.Obj()} }

// Does the object quack like a(n) Object?
func IsObject(e ld.Entity) bool { return ld.Is(e, TypeObject) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Object) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj Object) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Object) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Object) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Object) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Object) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

func (obj Object) Attachment() interface{} { return GetAttachment(obj) }

func (obj Object) Attachments() interface{} { return GetAttachments(obj) }

func (obj Object) Audience() interface{} { return GetAudience(obj) }

// Identifies the author of an object. Deprecated. Use as:attributedTo instead
func (obj Object) Author() interface{} { return GetAuthor(obj) }

func (obj Object) Bcc() interface{} { return GetBcc(obj) }

func (obj Object) Bto() interface{} { return GetBto(obj) }

func (obj Object) Cc() interface{} { return GetCc(obj) }

// The content of the object.
func (obj Object) Content() interface{} { return GetContent(obj) }

// Specifies the context within which an object exists or an activity was performed
func (obj Object) Context() interface{} { return GetContext(obj) }

func (obj Object) DownstreamDuplicates() interface{} { return GetDownstreamDuplicates(obj) }

// The duration of the object
func (obj Object) Duration() interface{} { return GetDuration(obj) }

// The ending time of the object
func (obj Object) EndTime() interface{} { return GetEndTime(obj) }

func (obj Object) Generator() interface{} { return GetGenerator(obj) }

func (obj Object) Icon() interface{} { return GetIcon(obj) }

func (obj Object) Image() interface{} { return GetImage(obj) }

func (obj Object) InReplyTo() interface{} { return GetInReplyTo(obj) }

func (obj Object) Location() interface{} { return GetLocation(obj) }

func (obj Object) ObjectType() interface{} { return GetObjectType(obj) }

func (obj Object) Provider() interface{} { return GetProvider(obj) }

// Specifies the date and time the object was published
func (obj Object) Published() interface{} { return GetPublished(obj) }

// A numeric rating (>= 0.0, <= 5.0) for the object
func (obj Object) Rating() interface{} { return GetRating(obj) }

func (obj Object) Replies() interface{} { return GetReplies(obj) }

// The starting time of the object
func (obj Object) StartTime() interface{} { return GetStartTime(obj) }

// A short summary of the object
func (obj Object) Summary() interface{} { return GetSummary(obj) }

func (obj Object) Tag() interface{} { return GetTag(obj) }

func (obj Object) Tags() interface{} { return GetTags(obj) }

func (obj Object) To() interface{} { return GetTo(obj) }

// Specifies when the object was last updated
func (obj Object) Updated() interface{} { return GetUpdated(obj) }

func (obj Object) UpstreamDuplicates() interface{} { return GetUpstreamDuplicates(obj) }

// Specifies a link to a specific representation of the Object
func (obj Object) Url() interface{} { return GetUrl(obj) }

// To Offer something to someone or something
type Offer struct{ Activity }

// Ducktypes the object into a(n) Offer.
func AsOffer(e ld.Entity) Offer { return Offer{AsActivity(e)} }

// Does the object quack like a(n) Offer?
func IsOffer(e ld.Entity) bool { return ld.Is(e, TypeOffer) }

// A variation of Collection in which items are strictly ordered
type OrderedCollection struct{ o *ld.Object }

// Ducktypes the object into a(n) OrderedCollection.
func AsOrderedCollection(e ld.Entity) OrderedCollection { return OrderedCollection{o: e.Obj()} }

// Does the object quack like a(n) OrderedCollection?
func IsOrderedCollection(e ld.Entity) bool { return ld.Is(e, TypeOrderedCollection) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj OrderedCollection) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj OrderedCollection) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj OrderedCollection) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj OrderedCollection) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj OrderedCollection) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj OrderedCollection) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// An ordered subset of items from an OrderedCollection
type OrderedCollectionPage struct {
	CollectionPage
	OrderedCollection
}

// Ducktypes the object into a(n) OrderedCollectionPage.
func AsOrderedCollectionPage(e ld.Entity) OrderedCollectionPage {
	return OrderedCollectionPage{AsCollectionPage(e), AsOrderedCollection(e)}
}

// Does the object quack like a(n) OrderedCollectionPage?
func IsOrderedCollectionPage(e ld.Entity) bool { return ld.Is(e, TypeOrderedCollectionPage) }

// In a strictly ordered logical collection, specifies the index position of the first item in the items list
func (obj OrderedCollectionPage) StartIndex() interface{} { return GetStartIndex(obj) }

// A rdf:List variant for Objects and Links
type OrderedItems struct{ o *ld.Object }

// Ducktypes the object into a(n) OrderedItems.
func AsOrderedItems(e ld.Entity) OrderedItems { return OrderedItems{o: e.Obj()} }

// Does the object quack like a(n) OrderedItems?
func IsOrderedItems(e ld.Entity) bool { return ld.Is(e, TypeOrderedItems) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj OrderedItems) Obj() *ld.Object { return obj.o }

// Returns the object's @id. Implements ld.Entity.
func (obj OrderedItems) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj OrderedItems) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj OrderedItems) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj OrderedItems) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj OrderedItems) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// An Organization
type Organization struct{ Object }

// Ducktypes the object into a(n) Organization.
func AsOrganization(e ld.Entity) Organization { return Organization{AsObject(e)} }

// Does the object quack like a(n) Organization?
func IsOrganization(e ld.Entity) bool { return ld.Is(e, TypeOrganization) }

// A Web Page
type Page struct{ Object }

// Ducktypes the object into a(n) Page.
func AsPage(e ld.Entity) Page { return Page{AsObject(e)} }

// Does the object quack like a(n) Page?
func IsPage(e ld.Entity) bool { return ld.Is(e, TypePage) }

// A Person
type Person struct{ Object }

// Ducktypes the object into a(n) Person.
func AsPerson(e ld.Entity) Person { return Person{AsObject(e)} }

// Does the object quack like a(n) Person?
func IsPerson(e ld.Entity) bool { return ld.Is(e, TypePerson) }

// A physical or logical location
type Place struct{ Object }

// Ducktypes the object into a(n) Place.
func AsPlace(e ld.Entity) Place { return Place{AsObject(e)} }

// Does the object quack like a(n) Place?
func IsPlace(e ld.Entity) bool { return ld.Is(e, TypePlace) }

// Specifies the accuracy around the point established by the longitude and latitude
func (obj Place) Accuracy() interface{} { return GetAccuracy(obj) }

// The altitude of a place
func (obj Place) Altitude() interface{} { return GetAltitude(obj) }

// The latitude
func (obj Place) Latitude() interface{} { return GetLatitude(obj) }

// The longitude
func (obj Place) Longitude() interface{} { return GetLongitude(obj) }

// Specifies a radius around the point established by the longitude and latitude
func (obj Place) Radius() interface{} { return GetRadius(obj) }

// Identifies the unit of measurement used by the radius, altitude and accuracy properties. The value can be expressed either as one of a set of predefined units or as a well-known common URI that identifies units.
func (obj Place) Units() interface{} { return GetUnits(obj) }

// A Profile Document
type Profile struct{ Object }

// Ducktypes the object into a(n) Profile.
func AsProfile(e ld.Entity) Profile { return Profile{AsObject(e)} }

// Does the object quack like a(n) Profile?
func IsProfile(e ld.Entity) bool { return ld.Is(e, TypeProfile) }

// On a Profile object, describes the object described by the profile
func (obj Profile) Describes() interface{} { return GetDescribes(obj) }

// A question of any sort.
type Question struct{ IntransitiveActivity }

// Ducktypes the object into a(n) Question.
func AsQuestion(e ld.Entity) Question { return Question{AsIntransitiveActivity(e)} }

// Does the object quack like a(n) Question?
func IsQuestion(e ld.Entity) bool { return ld.Is(e, TypeQuestion) }

// Describes a possible inclusive answer or option for a question.
func (obj Question) AnyOf() interface{} { return GetAnyOf(obj) }

// Describes a possible exclusive answer or option for a question.
func (obj Question) OneOf() interface{} { return GetOneOf(obj) }

// The actor read the object
type Read struct{ Activity }

// Ducktypes the object into a(n) Read.
func AsRead(e ld.Entity) Read { return Read{AsActivity(e)} }

// Does the object quack like a(n) Read?
func IsRead(e ld.Entity) bool { return ld.Is(e, TypeRead) }

// Actor rejects the Object
type Reject struct{ Activity }

// Ducktypes the object into a(n) Reject.
func AsReject(e ld.Entity) Reject { return Reject{AsActivity(e)} }

// Does the object quack like a(n) Reject?
func IsReject(e ld.Entity) bool { return ld.Is(e, TypeReject) }

// Represents a Social Graph relationship between two Individuals (indicated by the 'a' and 'b' properties)
type Relationship struct{ Object }

// Ducktypes the object into a(n) Relationship.
func AsRelationship(e ld.Entity) Relationship { return Relationship{AsObject(e)} }

// Does the object quack like a(n) Relationship?
func IsRelationship(e ld.Entity) bool { return ld.Is(e, TypeRelationship) }

// On a Relationship object, describes the type of relationship
func (obj Relationship) Relationship() interface{} { return GetRelationship(obj) }

// On a Relationship object, identifies the subject. e.g. when saying "John is connected to Sally", 'subject' refers to 'John'
func (obj Relationship) Subject() interface{} { return GetSubject(obj) }

// To Remove Something
type Remove struct{ Activity }

// Ducktypes the object into a(n) Remove.
func AsRemove(e ld.Entity) Remove { return Remove{AsActivity(e)} }

// Does the object quack like a(n) Remove?
func IsRemove(e ld.Entity) bool { return ld.Is(e, TypeRemove) }

// A service provided by some entity
type Service struct{ Object }

// Ducktypes the object into a(n) Service.
func AsService(e ld.Entity) Service { return Service{AsObject(e)} }

// Does the object quack like a(n) Service?
func IsService(e ld.Entity) bool { return ld.Is(e, TypeService) }

// Actor tentatively accepts the Object
type TentativeAccept struct{ Accept }

// Ducktypes the object into a(n) TentativeAccept.
func AsTentativeAccept(e ld.Entity) TentativeAccept { return TentativeAccept{AsAccept(e)} }

// Does the object quack like a(n) TentativeAccept?
func IsTentativeAccept(e ld.Entity) bool { return ld.Is(e, TypeTentativeAccept) }

// Actor tentatively rejects the object
type TentativeReject struct{ Reject }

// Ducktypes the object into a(n) TentativeReject.
func AsTentativeReject(e ld.Entity) TentativeReject { return TentativeReject{AsReject(e)} }

// Does the object quack like a(n) TentativeReject?
func IsTentativeReject(e ld.Entity) bool { return ld.Is(e, TypeTentativeReject) }

// A placeholder for a deleted object
type Tombstone struct{ Object }

// Ducktypes the object into a(n) Tombstone.
func AsTombstone(e ld.Entity) Tombstone { return Tombstone{AsObject(e)} }

// Does the object quack like a(n) Tombstone?
func IsTombstone(e ld.Entity) bool { return ld.Is(e, TypeTombstone) }

// Specifies the date and time the object was deleted
func (obj Tombstone) Deleted() interface{} { return GetDeleted(obj) }

// On a Tombstone object, describes the former type of the deleted object
func (obj Tombstone) FormerType() interface{} { return GetFormerType(obj) }

// The actor is traveling to the target. The origin specifies where the actor is traveling from.
type Travel struct{ IntransitiveActivity }

// Ducktypes the object into a(n) Travel.
func AsTravel(e ld.Entity) Travel { return Travel{AsIntransitiveActivity(e)} }

// Does the object quack like a(n) Travel?
func IsTravel(e ld.Entity) bool { return ld.Is(e, TypeTravel) }

// To Undo Something. This would typically be used to indicate that a previous Activity has been undone.
type Undo struct{ Activity }

// Ducktypes the object into a(n) Undo.
func AsUndo(e ld.Entity) Undo { return Undo{AsActivity(e)} }

// Does the object quack like a(n) Undo?
func IsUndo(e ld.Entity) bool { return ld.Is(e, TypeUndo) }

// To Update/Modify Something
type Update struct{ Activity }

// Ducktypes the object into a(n) Update.
func AsUpdate(e ld.Entity) Update { return Update{AsActivity(e)} }

// Does the object quack like a(n) Update?
func IsUpdate(e ld.Entity) bool { return ld.Is(e, TypeUpdate) }

// A Video document of any kind.
type Video struct{ Document }

// Ducktypes the object into a(n) Video.
func AsVideo(e ld.Entity) Video { return Video{AsDocument(e)} }

// Does the object quack like a(n) Video?
func IsVideo(e ld.Entity) bool { return ld.Is(e, TypeVideo) }

// The actor viewed the object
type View struct{ Activity }

// Ducktypes the object into a(n) View.
func AsView(e ld.Entity) View { return View{AsActivity(e)} }

// Does the object quack like a(n) View?
func IsView(e ld.Entity) bool { return ld.Is(e, TypeView) }

var (
	_ ld.Entity = Accept{}
	_ ld.Entity = Activity{}
	_ ld.Entity = Add{}
	_ ld.Entity = Announce{}
	_ ld.Entity = Application{}
	_ ld.Entity = Arrive{}
	_ ld.Entity = Article{}
	_ ld.Entity = Audio{}
	_ ld.Entity = Block{}
	_ ld.Entity = Collection{}
	_ ld.Entity = CollectionPage{}
	_ ld.Entity = Create{}
	_ ld.Entity = Delete{}
	_ ld.Entity = Dislike{}
	_ ld.Entity = Document{}
	_ ld.Entity = Event{}
	_ ld.Entity = Flag{}
	_ ld.Entity = Follow{}
	_ ld.Entity = Group{}
	_ ld.Entity = Ignore{}
	_ ld.Entity = Image{}
	_ ld.Entity = IntransitiveActivity{}
	_ ld.Entity = Invite{}
	_ ld.Entity = Join{}
	_ ld.Entity = Leave{}
	_ ld.Entity = Like{}
	_ ld.Entity = Link{}
	_ ld.Entity = Listen{}
	_ ld.Entity = Mention{}
	_ ld.Entity = Move{}
	_ ld.Entity = Note{}
	_ ld.Entity = Object{}
	_ ld.Entity = Offer{}
	_ ld.Entity = OrderedCollection{}
	_ ld.Entity = OrderedCollectionPage{}
	_ ld.Entity = OrderedItems{}
	_ ld.Entity = Organization{}
	_ ld.Entity = Page{}
	_ ld.Entity = Person{}
	_ ld.Entity = Place{}
	_ ld.Entity = Profile{}
	_ ld.Entity = Question{}
	_ ld.Entity = Read{}
	_ ld.Entity = Reject{}
	_ ld.Entity = Relationship{}
	_ ld.Entity = Remove{}
	_ ld.Entity = Service{}
	_ ld.Entity = TentativeAccept{}
	_ ld.Entity = TentativeReject{}
	_ ld.Entity = Tombstone{}
	_ ld.Entity = Travel{}
	_ ld.Entity = Undo{}
	_ ld.Entity = Update{}
	_ ld.Entity = Video{}
	_ ld.Entity = View{}
)
