// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package as

import (
	"github.com/meowpub/meow/ld"
)

// Actor accepts the Object
type Accept struct{ Activity }

func AsAccept(obj *ld.Object) Accept {
	return Accept{AsActivity(obj)}
}

func IsAccept(obj *ld.Object) bool {
	return ld.Is(obj, TypeAccept)
}

// An Object representing some form of Action that has been taken
type Activity struct{ Object }

func AsActivity(obj *ld.Object) Activity {
	return Activity{AsObject(obj)}
}

func IsActivity(obj *ld.Object) bool {
	return ld.Is(obj, TypeActivity)
}

// Subproperty of as:attributedTo that identifies the primary actor
func (obj Activity) Actor() interface{} {
	return obj.Get(PropActor)
}

// Indentifies an object used (or to be used) to complete an activity
func (obj Activity) Instrument() interface{} {
	return obj.Get(PropInstrument)
}

// For certain activities, specifies the entity from which the action is directed.
func (obj Activity) Origin() interface{} {
	return obj.Get(PropOrigin)
}

func (obj Activity) Result() interface{} {
	return obj.Get(PropResult)
}

func (obj Activity) Target() interface{} {
	return obj.Get(PropTarget)
}

func (obj Activity) Verb() interface{} {
	return obj.Get(PropVerb)
}

// To Add an Object or Link to Something
type Add struct{ Activity }

func AsAdd(obj *ld.Object) Add {
	return Add{AsActivity(obj)}
}

func IsAdd(obj *ld.Object) bool {
	return ld.Is(obj, TypeAdd)
}

// Actor announces the object to the target
type Announce struct{ Activity }

func AsAnnounce(obj *ld.Object) Announce {
	return Announce{AsActivity(obj)}
}

func IsAnnounce(obj *ld.Object) bool {
	return ld.Is(obj, TypeAnnounce)
}

// Represents a software application of any sort
type Application struct{ Object }

func AsApplication(obj *ld.Object) Application {
	return Application{AsObject(obj)}
}

func IsApplication(obj *ld.Object) bool {
	return ld.Is(obj, TypeApplication)
}

// To Arrive Somewhere (can be used, for instance, to indicate that a particular entity is currently located somewhere, e.g. a "check-in")
type Arrive struct{ IntransitiveActivity }

func AsArrive(obj *ld.Object) Arrive {
	return Arrive{AsIntransitiveActivity(obj)}
}

func IsArrive(obj *ld.Object) bool {
	return ld.Is(obj, TypeArrive)
}

// A written work. Typically several paragraphs long. For example, a blog post or a news article.
type Article struct{ Object }

func AsArticle(obj *ld.Object) Article {
	return Article{AsObject(obj)}
}

func IsArticle(obj *ld.Object) bool {
	return ld.Is(obj, TypeArticle)
}

// An audio file
type Audio struct{ Document }

func AsAudio(obj *ld.Object) Audio {
	return Audio{AsDocument(obj)}
}

func IsAudio(obj *ld.Object) bool {
	return ld.Is(obj, TypeAudio)
}

type Block struct{ Ignore }

func AsBlock(obj *ld.Object) Block {
	return Block{AsIgnore(obj)}
}

func IsBlock(obj *ld.Object) bool {
	return ld.Is(obj, TypeBlock)
}

// An ordered or unordered collection of Objects or Links
type Collection struct{ Object }

func AsCollection(obj *ld.Object) Collection {
	return Collection{AsObject(obj)}
}

func IsCollection(obj *ld.Object) bool {
	return ld.Is(obj, TypeCollection)
}

func (obj Collection) Current() interface{} {
	return obj.Get(PropCurrent)
}

func (obj Collection) First() interface{} {
	return obj.Get(PropFirst)
}

func (obj Collection) Items() interface{} {
	return obj.Get(PropItems)
}

func (obj Collection) Last() interface{} {
	return obj.Get(PropLast)
}

// The total number of items in a logical collection
func (obj Collection) TotalItems() interface{} {
	return obj.Get(PropTotalItems)
}

// A subset of items from a Collection
type CollectionPage struct{ Collection }

func AsCollectionPage(obj *ld.Object) CollectionPage {
	return CollectionPage{AsCollection(obj)}
}

func IsCollectionPage(obj *ld.Object) bool {
	return ld.Is(obj, TypeCollectionPage)
}

func (obj CollectionPage) Next() interface{} {
	return obj.Get(PropNext)
}

func (obj CollectionPage) PartOf() interface{} {
	return obj.Get(PropPartOf)
}

func (obj CollectionPage) Prev() interface{} {
	return obj.Get(PropPrev)
}

// To Create Something
type Create struct{ Activity }

func AsCreate(obj *ld.Object) Create {
	return Create{AsActivity(obj)}
}

func IsCreate(obj *ld.Object) bool {
	return ld.Is(obj, TypeCreate)
}

// To Delete Something
type Delete struct{ Activity }

func AsDelete(obj *ld.Object) Delete {
	return Delete{AsActivity(obj)}
}

func IsDelete(obj *ld.Object) bool {
	return ld.Is(obj, TypeDelete)
}

// The actor dislikes the object
type Dislike struct{ Activity }

func AsDislike(obj *ld.Object) Dislike {
	return Dislike{AsActivity(obj)}
}

func IsDislike(obj *ld.Object) bool {
	return ld.Is(obj, TypeDislike)
}

// Represents a digital document/file of any sort
type Document struct{ Object }

func AsDocument(obj *ld.Object) Document {
	return Document{AsObject(obj)}
}

func IsDocument(obj *ld.Object) bool {
	return ld.Is(obj, TypeDocument)
}

// An Event of any kind
type Event struct{ Object }

func AsEvent(obj *ld.Object) Event {
	return Event{AsObject(obj)}
}

func IsEvent(obj *ld.Object) bool {
	return ld.Is(obj, TypeEvent)
}

// To flag something (e.g. flag as inappropriate, flag as spam, etc)
type Flag struct{ Activity }

func AsFlag(obj *ld.Object) Flag {
	return Flag{AsActivity(obj)}
}

func IsFlag(obj *ld.Object) bool {
	return ld.Is(obj, TypeFlag)
}

// To Express Interest in Something
type Follow struct{ Activity }

func AsFollow(obj *ld.Object) Follow {
	return Follow{AsActivity(obj)}
}

func IsFollow(obj *ld.Object) bool {
	return ld.Is(obj, TypeFollow)
}

// A Group of any kind.
type Group struct{ Object }

func AsGroup(obj *ld.Object) Group {
	return Group{AsObject(obj)}
}

func IsGroup(obj *ld.Object) bool {
	return ld.Is(obj, TypeGroup)
}

// Actor is ignoring the Object
type Ignore struct{ Activity }

func AsIgnore(obj *ld.Object) Ignore {
	return Ignore{AsActivity(obj)}
}

func IsIgnore(obj *ld.Object) bool {
	return ld.Is(obj, TypeIgnore)
}

// An Image file
type Image struct{ Document }

func AsImage(obj *ld.Object) Image {
	return Image{AsDocument(obj)}
}

func IsImage(obj *ld.Object) bool {
	return ld.Is(obj, TypeImage)
}

// An Activity that has no direct object
type IntransitiveActivity struct{ Activity }

func AsIntransitiveActivity(obj *ld.Object) IntransitiveActivity {
	return IntransitiveActivity{AsActivity(obj)}
}

func IsIntransitiveActivity(obj *ld.Object) bool {
	return ld.Is(obj, TypeIntransitiveActivity)
}

// To invite someone or something to something
type Invite struct{ Offer }

func AsInvite(obj *ld.Object) Invite {
	return Invite{AsOffer(obj)}
}

func IsInvite(obj *ld.Object) bool {
	return ld.Is(obj, TypeInvite)
}

// To Join Something
type Join struct{ Activity }

func AsJoin(obj *ld.Object) Join {
	return Join{AsActivity(obj)}
}

func IsJoin(obj *ld.Object) bool {
	return ld.Is(obj, TypeJoin)
}

// To Leave Something
type Leave struct{ Activity }

func AsLeave(obj *ld.Object) Leave {
	return Leave{AsActivity(obj)}
}

func IsLeave(obj *ld.Object) bool {
	return ld.Is(obj, TypeLeave)
}

// To Like Something
type Like struct{ Activity }

func AsLike(obj *ld.Object) Like {
	return Like{AsActivity(obj)}
}

func IsLike(obj *ld.Object) bool {
	return ld.Is(obj, TypeLike)
}

// Represents a qualified reference to another resource. Patterned after the RFC5988 Web Linking Model
type Link struct{ o *ld.Object }

func AsLink(obj *ld.Object) Link {
	return Link{o: obj}
}

func IsLink(obj *ld.Object) bool {
	return ld.Is(obj, TypeLink)
}

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Link) Obj() *ld.Object {
	return obj.o
}

// Returns the object's @id. Implements ld.Entity.
func (obj Link) ID() string {
	return obj.o.ID()
}

// Returns the object's @value. Implements ld.Entity.
func (obj Link) Value() string {
	return obj.o.Value()
}

// Returns the object's @type. Implements ld.Entity.
func (obj Link) Type() []string {
	return obj.o.Type()
}

// Returns the named attribute. Implements ld.Entity.
func (obj Link) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Link) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// The display height expressed as device independent pixels
func (obj Link) Height() interface{} {
	return obj.Get(PropHeight)
}

// The target URI of the Link
func (obj Link) Href() interface{} {
	return obj.Get(PropHref)
}

// A hint about the language of the referenced resource
func (obj Link) Hreflang() interface{} {
	return obj.Get(PropHreflang)
}

// The RFC 5988 or HTML5 Link Relation associated with the Link
func (obj Link) Rel() interface{} {
	return obj.Get(PropRel)
}

// Specifies the preferred display width of the content, expressed in terms of device independent pixels.
func (obj Link) Width() interface{} {
	return obj.Get(PropWidth)
}

// The actor listened to the object
type Listen struct{ Activity }

func AsListen(obj *ld.Object) Listen {
	return Listen{AsActivity(obj)}
}

func IsListen(obj *ld.Object) bool {
	return ld.Is(obj, TypeListen)
}

// A specialized Link that represents an @mention
type Mention struct{ Link }

func AsMention(obj *ld.Object) Mention {
	return Mention{AsLink(obj)}
}

func IsMention(obj *ld.Object) bool {
	return ld.Is(obj, TypeMention)
}

// The actor is moving the object. The target specifies where the object is moving to. The origin specifies where the object is moving from.
type Move struct{ Activity }

func AsMove(obj *ld.Object) Move {
	return Move{AsActivity(obj)}
}

func IsMove(obj *ld.Object) bool {
	return ld.Is(obj, TypeMove)
}

// A Short note, typically less than a single paragraph. A "tweet" is an example, or a "status update"
type Note struct{ Object }

func AsNote(obj *ld.Object) Note {
	return Note{AsObject(obj)}
}

func IsNote(obj *ld.Object) bool {
	return ld.Is(obj, TypeNote)
}

type Object struct{ o *ld.Object }

func AsObject(obj *ld.Object) Object {
	return Object{o: obj}
}

func IsObject(obj *ld.Object) bool {
	return ld.Is(obj, TypeObject)
}

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Object) Obj() *ld.Object {
	return obj.o
}

// Returns the object's @id. Implements ld.Entity.
func (obj Object) ID() string {
	return obj.o.ID()
}

// Returns the object's @value. Implements ld.Entity.
func (obj Object) Value() string {
	return obj.o.Value()
}

// Returns the object's @type. Implements ld.Entity.
func (obj Object) Type() []string {
	return obj.o.Type()
}

// Returns the named attribute. Implements ld.Entity.
func (obj Object) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Object) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

func (obj Object) Attachment() interface{} {
	return obj.Get(PropAttachment)
}

func (obj Object) Attachments() interface{} {
	return obj.Get(PropAttachments)
}

func (obj Object) Audience() interface{} {
	return obj.Get(PropAudience)
}

// Identifies the author of an object. Deprecated. Use as:attributedTo instead
func (obj Object) Author() interface{} {
	return obj.Get(PropAuthor)
}

func (obj Object) Bcc() interface{} {
	return obj.Get(PropBcc)
}

func (obj Object) Bto() interface{} {
	return obj.Get(PropBto)
}

func (obj Object) Cc() interface{} {
	return obj.Get(PropCc)
}

// The content of the object.
func (obj Object) Content() interface{} {
	return obj.Get(PropContent)
}

// Specifies the context within which an object exists or an activity was performed
func (obj Object) Context() interface{} {
	return obj.Get(PropContext)
}

func (obj Object) DownstreamDuplicates() interface{} {
	return obj.Get(PropDownstreamDuplicates)
}

// The duration of the object
func (obj Object) Duration() interface{} {
	return obj.Get(PropDuration)
}

// The ending time of the object
func (obj Object) EndTime() interface{} {
	return obj.Get(PropEndTime)
}

func (obj Object) Generator() interface{} {
	return obj.Get(PropGenerator)
}

func (obj Object) Icon() interface{} {
	return obj.Get(PropIcon)
}

func (obj Object) Image() interface{} {
	return obj.Get(PropImage)
}

func (obj Object) InReplyTo() interface{} {
	return obj.Get(PropInReplyTo)
}

func (obj Object) Location() interface{} {
	return obj.Get(PropLocation)
}

func (obj Object) ObjectType() interface{} {
	return obj.Get(PropObjectType)
}

func (obj Object) Provider() interface{} {
	return obj.Get(PropProvider)
}

// Specifies the date and time the object was published
func (obj Object) Published() interface{} {
	return obj.Get(PropPublished)
}

// A numeric rating (>= 0.0, <= 5.0) for the object
func (obj Object) Rating() interface{} {
	return obj.Get(PropRating)
}

func (obj Object) Replies() interface{} {
	return obj.Get(PropReplies)
}

// The starting time of the object
func (obj Object) StartTime() interface{} {
	return obj.Get(PropStartTime)
}

// A short summary of the object
func (obj Object) Summary() interface{} {
	return obj.Get(PropSummary)
}

func (obj Object) Tag() interface{} {
	return obj.Get(PropTag)
}

func (obj Object) Tags() interface{} {
	return obj.Get(PropTags)
}

func (obj Object) To() interface{} {
	return obj.Get(PropTo)
}

// Specifies when the object was last updated
func (obj Object) Updated() interface{} {
	return obj.Get(PropUpdated)
}

func (obj Object) UpstreamDuplicates() interface{} {
	return obj.Get(PropUpstreamDuplicates)
}

// Specifies a link to a specific representation of the Object
func (obj Object) Url() interface{} {
	return obj.Get(PropUrl)
}

// To Offer something to someone or something
type Offer struct{ Activity }

func AsOffer(obj *ld.Object) Offer {
	return Offer{AsActivity(obj)}
}

func IsOffer(obj *ld.Object) bool {
	return ld.Is(obj, TypeOffer)
}

// A variation of Collection in which items are strictly ordered
type OrderedCollection struct{ o *ld.Object }

func AsOrderedCollection(obj *ld.Object) OrderedCollection {
	return OrderedCollection{o: obj}
}

func IsOrderedCollection(obj *ld.Object) bool {
	return ld.Is(obj, TypeOrderedCollection)
}

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj OrderedCollection) Obj() *ld.Object {
	return obj.o
}

// Returns the object's @id. Implements ld.Entity.
func (obj OrderedCollection) ID() string {
	return obj.o.ID()
}

// Returns the object's @value. Implements ld.Entity.
func (obj OrderedCollection) Value() string {
	return obj.o.Value()
}

// Returns the object's @type. Implements ld.Entity.
func (obj OrderedCollection) Type() []string {
	return obj.o.Type()
}

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

func AsOrderedCollectionPage(obj *ld.Object) OrderedCollectionPage {
	return OrderedCollectionPage{AsCollectionPage(obj), AsOrderedCollection(obj)}
}

func IsOrderedCollectionPage(obj *ld.Object) bool {
	return ld.Is(obj, TypeOrderedCollectionPage)
}

// In a strictly ordered logical collection, specifies the index position of the first item in the items list
func (obj OrderedCollectionPage) StartIndex() interface{} {
	return obj.Get(PropStartIndex)
}

// A rdf:List variant for Objects and Links
type OrderedItems struct{ o *ld.Object }

func AsOrderedItems(obj *ld.Object) OrderedItems {
	return OrderedItems{o: obj}
}

func IsOrderedItems(obj *ld.Object) bool {
	return ld.Is(obj, TypeOrderedItems)
}

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj OrderedItems) Obj() *ld.Object {
	return obj.o
}

// Returns the object's @id. Implements ld.Entity.
func (obj OrderedItems) ID() string {
	return obj.o.ID()
}

// Returns the object's @value. Implements ld.Entity.
func (obj OrderedItems) Value() string {
	return obj.o.Value()
}

// Returns the object's @type. Implements ld.Entity.
func (obj OrderedItems) Type() []string {
	return obj.o.Type()
}

// Returns the named attribute. Implements ld.Entity.
func (obj OrderedItems) Get(key string) interface{} { return obj.o.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj OrderedItems) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// An Organization
type Organization struct{ Object }

func AsOrganization(obj *ld.Object) Organization {
	return Organization{AsObject(obj)}
}

func IsOrganization(obj *ld.Object) bool {
	return ld.Is(obj, TypeOrganization)
}

// A Web Page
type Page struct{ Object }

func AsPage(obj *ld.Object) Page {
	return Page{AsObject(obj)}
}

func IsPage(obj *ld.Object) bool {
	return ld.Is(obj, TypePage)
}

// A Person
type Person struct{ Object }

func AsPerson(obj *ld.Object) Person {
	return Person{AsObject(obj)}
}

func IsPerson(obj *ld.Object) bool {
	return ld.Is(obj, TypePerson)
}

// A physical or logical location
type Place struct{ Object }

func AsPlace(obj *ld.Object) Place {
	return Place{AsObject(obj)}
}

func IsPlace(obj *ld.Object) bool {
	return ld.Is(obj, TypePlace)
}

// Specifies the accuracy around the point established by the longitude and latitude
func (obj Place) Accuracy() interface{} {
	return obj.Get(PropAccuracy)
}

// The altitude of a place
func (obj Place) Altitude() interface{} {
	return obj.Get(PropAltitude)
}

// The latitude
func (obj Place) Latitude() interface{} {
	return obj.Get(PropLatitude)
}

// The longitude
func (obj Place) Longitude() interface{} {
	return obj.Get(PropLongitude)
}

// Specifies a radius around the point established by the longitude and latitude
func (obj Place) Radius() interface{} {
	return obj.Get(PropRadius)
}

// Identifies the unit of measurement used by the radius, altitude and accuracy properties. The value can be expressed either as one of a set of predefined units or as a well-known common URI that identifies units.
func (obj Place) Units() interface{} {
	return obj.Get(PropUnits)
}

// A Profile Document
type Profile struct{ Object }

func AsProfile(obj *ld.Object) Profile {
	return Profile{AsObject(obj)}
}

func IsProfile(obj *ld.Object) bool {
	return ld.Is(obj, TypeProfile)
}

// On a Profile object, describes the object described by the profile
func (obj Profile) Describes() interface{} {
	return obj.Get(PropDescribes)
}

// A question of any sort.
type Question struct{ IntransitiveActivity }

func AsQuestion(obj *ld.Object) Question {
	return Question{AsIntransitiveActivity(obj)}
}

func IsQuestion(obj *ld.Object) bool {
	return ld.Is(obj, TypeQuestion)
}

// Describes a possible inclusive answer or option for a question.
func (obj Question) AnyOf() interface{} {
	return obj.Get(PropAnyOf)
}

// Describes a possible exclusive answer or option for a question.
func (obj Question) OneOf() interface{} {
	return obj.Get(PropOneOf)
}

// The actor read the object
type Read struct{ Activity }

func AsRead(obj *ld.Object) Read {
	return Read{AsActivity(obj)}
}

func IsRead(obj *ld.Object) bool {
	return ld.Is(obj, TypeRead)
}

// Actor rejects the Object
type Reject struct{ Activity }

func AsReject(obj *ld.Object) Reject {
	return Reject{AsActivity(obj)}
}

func IsReject(obj *ld.Object) bool {
	return ld.Is(obj, TypeReject)
}

// Represents a Social Graph relationship between two Individuals (indicated by the 'a' and 'b' properties)
type Relationship struct{ Object }

func AsRelationship(obj *ld.Object) Relationship {
	return Relationship{AsObject(obj)}
}

func IsRelationship(obj *ld.Object) bool {
	return ld.Is(obj, TypeRelationship)
}

// On a Relationship object, describes the type of relationship
func (obj Relationship) Relationship() interface{} {
	return obj.Get(PropRelationship)
}

// On a Relationship object, identifies the subject. e.g. when saying "John is connected to Sally", 'subject' refers to 'John'
func (obj Relationship) Subject() interface{} {
	return obj.Get(PropSubject)
}

// To Remove Something
type Remove struct{ Activity }

func AsRemove(obj *ld.Object) Remove {
	return Remove{AsActivity(obj)}
}

func IsRemove(obj *ld.Object) bool {
	return ld.Is(obj, TypeRemove)
}

// A service provided by some entity
type Service struct{ Object }

func AsService(obj *ld.Object) Service {
	return Service{AsObject(obj)}
}

func IsService(obj *ld.Object) bool {
	return ld.Is(obj, TypeService)
}

// Actor tentatively accepts the Object
type TentativeAccept struct{ Accept }

func AsTentativeAccept(obj *ld.Object) TentativeAccept {
	return TentativeAccept{AsAccept(obj)}
}

func IsTentativeAccept(obj *ld.Object) bool {
	return ld.Is(obj, TypeTentativeAccept)
}

// Actor tentatively rejects the object
type TentativeReject struct{ Reject }

func AsTentativeReject(obj *ld.Object) TentativeReject {
	return TentativeReject{AsReject(obj)}
}

func IsTentativeReject(obj *ld.Object) bool {
	return ld.Is(obj, TypeTentativeReject)
}

// A placeholder for a deleted object
type Tombstone struct{ Object }

func AsTombstone(obj *ld.Object) Tombstone {
	return Tombstone{AsObject(obj)}
}

func IsTombstone(obj *ld.Object) bool {
	return ld.Is(obj, TypeTombstone)
}

// Specifies the date and time the object was deleted
func (obj Tombstone) Deleted() interface{} {
	return obj.Get(PropDeleted)
}

// On a Tombstone object, describes the former type of the deleted object
func (obj Tombstone) FormerType() interface{} {
	return obj.Get(PropFormerType)
}

// The actor is traveling to the target. The origin specifies where the actor is traveling from.
type Travel struct{ IntransitiveActivity }

func AsTravel(obj *ld.Object) Travel {
	return Travel{AsIntransitiveActivity(obj)}
}

func IsTravel(obj *ld.Object) bool {
	return ld.Is(obj, TypeTravel)
}

// To Undo Something. This would typically be used to indicate that a previous Activity has been undone.
type Undo struct{ Activity }

func AsUndo(obj *ld.Object) Undo {
	return Undo{AsActivity(obj)}
}

func IsUndo(obj *ld.Object) bool {
	return ld.Is(obj, TypeUndo)
}

// To Update/Modify Something
type Update struct{ Activity }

func AsUpdate(obj *ld.Object) Update {
	return Update{AsActivity(obj)}
}

func IsUpdate(obj *ld.Object) bool {
	return ld.Is(obj, TypeUpdate)
}

// A Video document of any kind.
type Video struct{ Document }

func AsVideo(obj *ld.Object) Video {
	return Video{AsDocument(obj)}
}

func IsVideo(obj *ld.Object) bool {
	return ld.Is(obj, TypeVideo)
}

// The actor viewed the object
type View struct{ Activity }

func AsView(obj *ld.Object) View {
	return View{AsActivity(obj)}
}

func IsView(obj *ld.Object) bool {
	return ld.Is(obj, TypeView)
}

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
