// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package as

import (
	"github.com/meowpub/meow/ld"
)

// Actor accepts the Object
type Accept struct{ Activity }

// An Object representing some form of Action that has been taken
type Activity struct{ Object }

// Subproperty of as:attributedTo that identifies the primary actor
func (obj Activity) Actor() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#actor")
}

// Indentifies an object used (or to be used) to complete an activity
func (obj Activity) Instrument() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#instrument")
}

// For certain activities, specifies the entity from which the action is directed.
func (obj Activity) Origin() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#origin")
}

func (obj Activity) Result() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#result")
}

func (obj Activity) Target() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#target")
}

func (obj Activity) Verb() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#verb")
}

// To Add an Object or Link to Something
type Add struct{ Activity }

// Actor announces the object to the target
type Announce struct{ Activity }

// Represents a software application of any sort
type Application struct{ Object }

// To Arrive Somewhere (can be used, for instance, to indicate that a particular entity is currently located somewhere, e.g. a "check-in")
type Arrive struct{ IntransitiveActivity }

// A written work. Typically several paragraphs long. For example, a blog post or a news article.
type Article struct{ Object }

// An audio file
type Audio struct{ Document }

type Block struct{ Ignore }

// An ordered or unordered collection of Objects or Links
type Collection struct{ Object }

func (obj Collection) Current() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#current")
}

func (obj Collection) First() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#first")
}

func (obj Collection) Items() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#items")
}

func (obj Collection) Last() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#last")
}

// The total number of items in a logical collection
func (obj Collection) TotalItems() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#totalItems")
}

// A subset of items from a Collection
type CollectionPage struct{ Collection }

func (obj CollectionPage) Next() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#next")
}

func (obj CollectionPage) PartOf() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#partOf")
}

func (obj CollectionPage) Prev() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#prev")
}

// To Create Something
type Create struct{ Activity }

// To Delete Something
type Delete struct{ Activity }

// The actor dislikes the object
type Dislike struct{ Activity }

// Represents a digital document/file of any sort
type Document struct{ Object }

// An Event of any kind
type Event struct{ Object }

// To flag something (e.g. flag as inappropriate, flag as spam, etc)
type Flag struct{ Activity }

// To Express Interest in Something
type Follow struct{ Activity }

// A Group of any kind.
type Group struct{ Object }

// Actor is ignoring the Object
type Ignore struct{ Activity }

// An Image file
type Image struct{ Document }

// An Activity that has no direct object
type IntransitiveActivity struct{ Activity }

// To invite someone or something to something
type Invite struct{ Offer }

// To Join Something
type Join struct{ Activity }

// To Leave Something
type Leave struct{ Activity }

// To Like Something
type Like struct{ Activity }

// Represents a qualified reference to another resource. Patterned after the RFC5988 Web Linking Model
type Link struct{ O *ld.Object }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Link) Obj() *ld.Object {
	return obj.O
}

// Returns the object's @id. Implements ld.Entity.
func (obj Link) ID() string {
	return obj.O.ID()
}

// Returns the object's @value. Implements ld.Entity.
func (obj Link) Value() string {
	return obj.O.Value()
}

// Returns the object's @type. Implements ld.Entity.
func (obj Link) Type() []string {
	return obj.O.Type()
}

// Returns the named attribute. Implements ld.Entity.
func (obj Link) Get(key string) interface{} { return obj.O.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Link) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// The display height expressed as device independent pixels
func (obj Link) Height() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#height")
}

// The target URI of the Link
func (obj Link) Href() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#href")
}

// A hint about the language of the referenced resource
func (obj Link) Hreflang() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#hreflang")
}

// The RFC 5988 or HTML5 Link Relation associated with the Link
func (obj Link) Rel() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#rel")
}

// Specifies the preferred display width of the content, expressed in terms of device independent pixels.
func (obj Link) Width() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#width")
}

// The actor listened to the object
type Listen struct{ Activity }

// A specialized Link that represents an @mention
type Mention struct{ Link }

// The actor is moving the object. The target specifies where the object is moving to. The origin specifies where the object is moving from.
type Move struct{ Activity }

// A Short note, typically less than a single paragraph. A "tweet" is an example, or a "status update"
type Note struct{ Object }

type Object struct{ O *ld.Object }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Object) Obj() *ld.Object {
	return obj.O
}

// Returns the object's @id. Implements ld.Entity.
func (obj Object) ID() string {
	return obj.O.ID()
}

// Returns the object's @value. Implements ld.Entity.
func (obj Object) Value() string {
	return obj.O.Value()
}

// Returns the object's @type. Implements ld.Entity.
func (obj Object) Type() []string {
	return obj.O.Type()
}

// Returns the named attribute. Implements ld.Entity.
func (obj Object) Get(key string) interface{} { return obj.O.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Object) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

func (obj Object) Attachment() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#attachment")
}

func (obj Object) Attachments() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#attachments")
}

func (obj Object) Audience() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#audience")
}

// Identifies the author of an object. Deprecated. Use as:attributedTo instead
func (obj Object) Author() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#author")
}

func (obj Object) Bcc() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#bcc")
}

func (obj Object) Bto() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#bto")
}

func (obj Object) Cc() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#cc")
}

// The content of the object.
func (obj Object) Content() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#content")
}

// Specifies the context within which an object exists or an activity was performed
func (obj Object) Context() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#context")
}

func (obj Object) DownstreamDuplicates() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#downstreamDuplicates")
}

// The duration of the object
func (obj Object) Duration() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#duration")
}

// The ending time of the object
func (obj Object) EndTime() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#endTime")
}

func (obj Object) Generator() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#generator")
}

func (obj Object) Icon() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#icon")
}

func (obj Object) Image() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#image")
}

func (obj Object) InReplyTo() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#inReplyTo")
}

func (obj Object) Location() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#location")
}

func (obj Object) ObjectType() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#objectType")
}

func (obj Object) Provider() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#provider")
}

// Specifies the date and time the object was published
func (obj Object) Published() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#published")
}

// A numeric rating (>= 0.0, <= 5.0) for the object
func (obj Object) Rating() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#rating")
}

func (obj Object) Replies() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#replies")
}

// The starting time of the object
func (obj Object) StartTime() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#startTime")
}

// A short summary of the object
func (obj Object) Summary() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#summary")
}

func (obj Object) Tag() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#tag")
}

func (obj Object) Tags() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#tags")
}

func (obj Object) To() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#to")
}

// Specifies when the object was last updated
func (obj Object) Updated() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#updated")
}

func (obj Object) UpstreamDuplicates() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#upstreamDuplicates")
}

// Specifies a link to a specific representation of the Object
func (obj Object) Url() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#url")
}

// To Offer something to someone or something
type Offer struct{ Activity }

// A variation of Collection in which items are strictly ordered
type OrderedCollection struct{ O *ld.Object }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj OrderedCollection) Obj() *ld.Object {
	return obj.O
}

// Returns the object's @id. Implements ld.Entity.
func (obj OrderedCollection) ID() string {
	return obj.O.ID()
}

// Returns the object's @value. Implements ld.Entity.
func (obj OrderedCollection) Value() string {
	return obj.O.Value()
}

// Returns the object's @type. Implements ld.Entity.
func (obj OrderedCollection) Type() []string {
	return obj.O.Type()
}

// Returns the named attribute. Implements ld.Entity.
func (obj OrderedCollection) Get(key string) interface{} { return obj.O.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj OrderedCollection) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// An ordered subset of items from an OrderedCollection
type OrderedCollectionPage struct{ CollectionPage }

// In a strictly ordered logical collection, specifies the index position of the first item in the items list
func (obj OrderedCollectionPage) StartIndex() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#startIndex")
}

// A rdf:List variant for Objects and Links
type OrderedItems struct{ O *ld.Object }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj OrderedItems) Obj() *ld.Object {
	return obj.O
}

// Returns the object's @id. Implements ld.Entity.
func (obj OrderedItems) ID() string {
	return obj.O.ID()
}

// Returns the object's @value. Implements ld.Entity.
func (obj OrderedItems) Value() string {
	return obj.O.Value()
}

// Returns the object's @type. Implements ld.Entity.
func (obj OrderedItems) Type() []string {
	return obj.O.Type()
}

// Returns the named attribute. Implements ld.Entity.
func (obj OrderedItems) Get(key string) interface{} { return obj.O.Get(key) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj OrderedItems) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.O.Apply(other, mergeArrays)
}

// An Organization
type Organization struct{ Object }

// A Web Page
type Page struct{ Object }

// A Person
type Person struct{ Object }

// A physical or logical location
type Place struct{ Object }

// Specifies the accuracy around the point established by the longitude and latitude
func (obj Place) Accuracy() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#accuracy")
}

// The altitude of a place
func (obj Place) Altitude() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#altitude")
}

// The latitude
func (obj Place) Latitude() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#latitude")
}

// The longitude
func (obj Place) Longitude() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#longitude")
}

// Specifies a radius around the point established by the longitude and latitude
func (obj Place) Radius() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#radius")
}

// Identifies the unit of measurement used by the radius, altitude and accuracy properties. The value can be expressed either as one of a set of predefined units or as a well-known common URI that identifies units.
func (obj Place) Units() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#units")
}

// A Profile Document
type Profile struct{ Object }

// On a Profile object, describes the object described by the profile
func (obj Profile) Describes() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#describes")
}

// A question of any sort.
type Question struct{ IntransitiveActivity }

// Describes a possible inclusive answer or option for a question.
func (obj Question) AnyOf() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#anyOf")
}

// Describes a possible exclusive answer or option for a question.
func (obj Question) OneOf() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#oneOf")
}

// The actor read the object
type Read struct{ Activity }

// Actor rejects the Object
type Reject struct{ Activity }

// Represents a Social Graph relationship between two Individuals (indicated by the 'a' and 'b' properties)
type Relationship struct{ Object }

// On a Relationship object, describes the type of relationship
func (obj Relationship) Relationship() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#relationship")
}

// On a Relationship object, identifies the subject. e.g. when saying "John is connected to Sally", 'subject' refers to 'John'
func (obj Relationship) Subject() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#subject")
}

// To Remove Something
type Remove struct{ Activity }

// A service provided by some entity
type Service struct{ Object }

// Actor tentatively accepts the Object
type TentativeAccept struct{ Accept }

// Actor tentatively rejects the object
type TentativeReject struct{ Reject }

// A placeholder for a deleted object
type Tombstone struct{ Object }

// Specifies the date and time the object was deleted
func (obj Tombstone) Deleted() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#deleted")
}

// On a Tombstone object, describes the former type of the deleted object
func (obj Tombstone) FormerType() interface{} {
	return obj.Get("http://www.w3.org/ns/activitystreams#formerType")
}

// The actor is traveling to the target. The origin specifies where the actor is traveling from.
type Travel struct{ IntransitiveActivity }

// To Undo Something. This would typically be used to indicate that a previous Activity has been undone.
type Undo struct{ Activity }

// To Update/Modify Something
type Update struct{ Activity }

// A Video document of any kind.
type Video struct{ Document }

// The actor viewed the object
type View struct{ Activity }

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
