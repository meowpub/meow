// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package as

import (
	"github.com/meowpub/meow/ld"
)

// Actor accepts the Object
type Accept struct{ Activity }

func NewAccept(id string) Accept { return AsAccept(ld.NewObject(id, Class_Accept.ID)) }

// Ducktypes the object into a(n) Accept.
func AsAccept(e ld.Entity) Accept { return Accept{AsActivity(e)} }

// Does the object quack like a(n) Accept?
func IsAccept(e ld.Entity) bool { return ld.Is(e, Class_Accept.ID) }

// An Object representing some form of Action that has been taken
type Activity struct{ Object }

func NewActivity(id string) Activity { return AsActivity(ld.NewObject(id, Class_Activity.ID)) }

// Ducktypes the object into a(n) Activity.
func AsActivity(e ld.Entity) Activity { return Activity{AsObject(e)} }

// Does the object quack like a(n) Activity?
func IsActivity(e ld.Entity) bool { return ld.Is(e, Class_Activity.ID) }

// Subproperty of as:attributedTo that identifies the primary actor
func (obj Activity) Actor() interface{} { return GetActor(obj) }

func (obj Activity) SetActor(v interface{}) { SetActor(obj, v) }

// Indentifies an object used (or to be used) to complete an activity
func (obj Activity) Instrument() interface{} { return GetInstrument(obj) }

func (obj Activity) SetInstrument(v interface{}) { SetInstrument(obj, v) }

// For certain activities, specifies the entity from which the action is directed.
func (obj Activity) Origin() interface{} { return GetOrigin(obj) }

func (obj Activity) SetOrigin(v interface{}) { SetOrigin(obj, v) }

func (obj Activity) Result() interface{} { return GetResult(obj) }

func (obj Activity) SetResult(v interface{}) { SetResult(obj, v) }

func (obj Activity) Target() interface{} { return GetTarget(obj) }

func (obj Activity) SetTarget(v interface{}) { SetTarget(obj, v) }

func (obj Activity) Verb() interface{} { return GetVerb(obj) }

func (obj Activity) SetVerb(v interface{}) { SetVerb(obj, v) }

// To Add an Object or Link to Something
type Add struct{ Activity }

func NewAdd(id string) Add { return AsAdd(ld.NewObject(id, Class_Add.ID)) }

// Ducktypes the object into a(n) Add.
func AsAdd(e ld.Entity) Add { return Add{AsActivity(e)} }

// Does the object quack like a(n) Add?
func IsAdd(e ld.Entity) bool { return ld.Is(e, Class_Add.ID) }

// Actor announces the object to the target
type Announce struct{ Activity }

func NewAnnounce(id string) Announce { return AsAnnounce(ld.NewObject(id, Class_Announce.ID)) }

// Ducktypes the object into a(n) Announce.
func AsAnnounce(e ld.Entity) Announce { return Announce{AsActivity(e)} }

// Does the object quack like a(n) Announce?
func IsAnnounce(e ld.Entity) bool { return ld.Is(e, Class_Announce.ID) }

// Represents a software application of any sort
type Application struct{ Object }

func NewApplication(id string) Application {
	return AsApplication(ld.NewObject(id, Class_Application.ID))
}

// Ducktypes the object into a(n) Application.
func AsApplication(e ld.Entity) Application { return Application{AsObject(e)} }

// Does the object quack like a(n) Application?
func IsApplication(e ld.Entity) bool { return ld.Is(e, Class_Application.ID) }

// To Arrive Somewhere (can be used, for instance, to indicate that a particular entity is currently located somewhere, e.g. a "check-in")
type Arrive struct{ IntransitiveActivity }

func NewArrive(id string) Arrive { return AsArrive(ld.NewObject(id, Class_Arrive.ID)) }

// Ducktypes the object into a(n) Arrive.
func AsArrive(e ld.Entity) Arrive { return Arrive{AsIntransitiveActivity(e)} }

// Does the object quack like a(n) Arrive?
func IsArrive(e ld.Entity) bool { return ld.Is(e, Class_Arrive.ID) }

// A written work. Typically several paragraphs long. For example, a blog post or a news article.
type Article struct{ Object }

func NewArticle(id string) Article { return AsArticle(ld.NewObject(id, Class_Article.ID)) }

// Ducktypes the object into a(n) Article.
func AsArticle(e ld.Entity) Article { return Article{AsObject(e)} }

// Does the object quack like a(n) Article?
func IsArticle(e ld.Entity) bool { return ld.Is(e, Class_Article.ID) }

// An audio file
type Audio struct{ Document }

func NewAudio(id string) Audio { return AsAudio(ld.NewObject(id, Class_Audio.ID)) }

// Ducktypes the object into a(n) Audio.
func AsAudio(e ld.Entity) Audio { return Audio{AsDocument(e)} }

// Does the object quack like a(n) Audio?
func IsAudio(e ld.Entity) bool { return ld.Is(e, Class_Audio.ID) }

type Block struct{ Ignore }

func NewBlock(id string) Block { return AsBlock(ld.NewObject(id, Class_Block.ID)) }

// Ducktypes the object into a(n) Block.
func AsBlock(e ld.Entity) Block { return Block{AsIgnore(e)} }

// Does the object quack like a(n) Block?
func IsBlock(e ld.Entity) bool { return ld.Is(e, Class_Block.ID) }

// An ordered or unordered collection of Objects or Links
type Collection struct{ Object }

func NewCollection(id string) Collection { return AsCollection(ld.NewObject(id, Class_Collection.ID)) }

// Ducktypes the object into a(n) Collection.
func AsCollection(e ld.Entity) Collection { return Collection{AsObject(e)} }

// Does the object quack like a(n) Collection?
func IsCollection(e ld.Entity) bool { return ld.Is(e, Class_Collection.ID) }

func (obj Collection) Current() interface{} { return GetCurrent(obj) }

func (obj Collection) SetCurrent(v interface{}) { SetCurrent(obj, v) }

func (obj Collection) First() interface{} { return GetFirst(obj) }

func (obj Collection) SetFirst(v interface{}) { SetFirst(obj, v) }

func (obj Collection) Items() interface{} { return GetItems(obj) }

func (obj Collection) SetItems(v interface{}) { SetItems(obj, v) }

func (obj Collection) Last() interface{} { return GetLast(obj) }

func (obj Collection) SetLast(v interface{}) { SetLast(obj, v) }

// The total number of items in a logical collection
func (obj Collection) TotalItems() interface{} { return GetTotalItems(obj) }

func (obj Collection) SetTotalItems(v interface{}) { SetTotalItems(obj, v) }

// A subset of items from a Collection
type CollectionPage struct{ Collection }

func NewCollectionPage(id string) CollectionPage {
	return AsCollectionPage(ld.NewObject(id, Class_CollectionPage.ID))
}

// Ducktypes the object into a(n) CollectionPage.
func AsCollectionPage(e ld.Entity) CollectionPage { return CollectionPage{AsCollection(e)} }

// Does the object quack like a(n) CollectionPage?
func IsCollectionPage(e ld.Entity) bool { return ld.Is(e, Class_CollectionPage.ID) }

func (obj CollectionPage) Next() interface{} { return GetNext(obj) }

func (obj CollectionPage) SetNext(v interface{}) { SetNext(obj, v) }

func (obj CollectionPage) PartOf() interface{} { return GetPartOf(obj) }

func (obj CollectionPage) SetPartOf(v interface{}) { SetPartOf(obj, v) }

func (obj CollectionPage) Prev() interface{} { return GetPrev(obj) }

func (obj CollectionPage) SetPrev(v interface{}) { SetPrev(obj, v) }

// To Create Something
type Create struct{ Activity }

func NewCreate(id string) Create { return AsCreate(ld.NewObject(id, Class_Create.ID)) }

// Ducktypes the object into a(n) Create.
func AsCreate(e ld.Entity) Create { return Create{AsActivity(e)} }

// Does the object quack like a(n) Create?
func IsCreate(e ld.Entity) bool { return ld.Is(e, Class_Create.ID) }

// To Delete Something
type Delete struct{ Activity }

func NewDelete(id string) Delete { return AsDelete(ld.NewObject(id, Class_Delete.ID)) }

// Ducktypes the object into a(n) Delete.
func AsDelete(e ld.Entity) Delete { return Delete{AsActivity(e)} }

// Does the object quack like a(n) Delete?
func IsDelete(e ld.Entity) bool { return ld.Is(e, Class_Delete.ID) }

// The actor dislikes the object
type Dislike struct{ Activity }

func NewDislike(id string) Dislike { return AsDislike(ld.NewObject(id, Class_Dislike.ID)) }

// Ducktypes the object into a(n) Dislike.
func AsDislike(e ld.Entity) Dislike { return Dislike{AsActivity(e)} }

// Does the object quack like a(n) Dislike?
func IsDislike(e ld.Entity) bool { return ld.Is(e, Class_Dislike.ID) }

// Represents a digital document/file of any sort
type Document struct{ Object }

func NewDocument(id string) Document { return AsDocument(ld.NewObject(id, Class_Document.ID)) }

// Ducktypes the object into a(n) Document.
func AsDocument(e ld.Entity) Document { return Document{AsObject(e)} }

// Does the object quack like a(n) Document?
func IsDocument(e ld.Entity) bool { return ld.Is(e, Class_Document.ID) }

// An Event of any kind
type Event struct{ Object }

func NewEvent(id string) Event { return AsEvent(ld.NewObject(id, Class_Event.ID)) }

// Ducktypes the object into a(n) Event.
func AsEvent(e ld.Entity) Event { return Event{AsObject(e)} }

// Does the object quack like a(n) Event?
func IsEvent(e ld.Entity) bool { return ld.Is(e, Class_Event.ID) }

// To flag something (e.g. flag as inappropriate, flag as spam, etc)
type Flag struct{ Activity }

func NewFlag(id string) Flag { return AsFlag(ld.NewObject(id, Class_Flag.ID)) }

// Ducktypes the object into a(n) Flag.
func AsFlag(e ld.Entity) Flag { return Flag{AsActivity(e)} }

// Does the object quack like a(n) Flag?
func IsFlag(e ld.Entity) bool { return ld.Is(e, Class_Flag.ID) }

// To Express Interest in Something
type Follow struct{ Activity }

func NewFollow(id string) Follow { return AsFollow(ld.NewObject(id, Class_Follow.ID)) }

// Ducktypes the object into a(n) Follow.
func AsFollow(e ld.Entity) Follow { return Follow{AsActivity(e)} }

// Does the object quack like a(n) Follow?
func IsFollow(e ld.Entity) bool { return ld.Is(e, Class_Follow.ID) }

// A Group of any kind.
type Group struct{ Object }

func NewGroup(id string) Group { return AsGroup(ld.NewObject(id, Class_Group.ID)) }

// Ducktypes the object into a(n) Group.
func AsGroup(e ld.Entity) Group { return Group{AsObject(e)} }

// Does the object quack like a(n) Group?
func IsGroup(e ld.Entity) bool { return ld.Is(e, Class_Group.ID) }

// Actor is ignoring the Object
type Ignore struct{ Activity }

func NewIgnore(id string) Ignore { return AsIgnore(ld.NewObject(id, Class_Ignore.ID)) }

// Ducktypes the object into a(n) Ignore.
func AsIgnore(e ld.Entity) Ignore { return Ignore{AsActivity(e)} }

// Does the object quack like a(n) Ignore?
func IsIgnore(e ld.Entity) bool { return ld.Is(e, Class_Ignore.ID) }

// An Image file
type Image struct{ Document }

func NewImage(id string) Image { return AsImage(ld.NewObject(id, Class_Image.ID)) }

// Ducktypes the object into a(n) Image.
func AsImage(e ld.Entity) Image { return Image{AsDocument(e)} }

// Does the object quack like a(n) Image?
func IsImage(e ld.Entity) bool { return ld.Is(e, Class_Image.ID) }

// An Activity that has no direct object
type IntransitiveActivity struct{ Activity }

func NewIntransitiveActivity(id string) IntransitiveActivity {
	return AsIntransitiveActivity(ld.NewObject(id, Class_IntransitiveActivity.ID))
}

// Ducktypes the object into a(n) IntransitiveActivity.
func AsIntransitiveActivity(e ld.Entity) IntransitiveActivity {
	return IntransitiveActivity{AsActivity(e)}
}

// Does the object quack like a(n) IntransitiveActivity?
func IsIntransitiveActivity(e ld.Entity) bool { return ld.Is(e, Class_IntransitiveActivity.ID) }

// To invite someone or something to something
type Invite struct{ Offer }

func NewInvite(id string) Invite { return AsInvite(ld.NewObject(id, Class_Invite.ID)) }

// Ducktypes the object into a(n) Invite.
func AsInvite(e ld.Entity) Invite { return Invite{AsOffer(e)} }

// Does the object quack like a(n) Invite?
func IsInvite(e ld.Entity) bool { return ld.Is(e, Class_Invite.ID) }

// To Join Something
type Join struct{ Activity }

func NewJoin(id string) Join { return AsJoin(ld.NewObject(id, Class_Join.ID)) }

// Ducktypes the object into a(n) Join.
func AsJoin(e ld.Entity) Join { return Join{AsActivity(e)} }

// Does the object quack like a(n) Join?
func IsJoin(e ld.Entity) bool { return ld.Is(e, Class_Join.ID) }

// To Leave Something
type Leave struct{ Activity }

func NewLeave(id string) Leave { return AsLeave(ld.NewObject(id, Class_Leave.ID)) }

// Ducktypes the object into a(n) Leave.
func AsLeave(e ld.Entity) Leave { return Leave{AsActivity(e)} }

// Does the object quack like a(n) Leave?
func IsLeave(e ld.Entity) bool { return ld.Is(e, Class_Leave.ID) }

// To Like Something
type Like struct{ Activity }

func NewLike(id string) Like { return AsLike(ld.NewObject(id, Class_Like.ID)) }

// Ducktypes the object into a(n) Like.
func AsLike(e ld.Entity) Like { return Like{AsActivity(e)} }

// Does the object quack like a(n) Like?
func IsLike(e ld.Entity) bool { return ld.Is(e, Class_Like.ID) }

// Represents a qualified reference to another resource. Patterned after the RFC5988 Web Linking Model
type Link struct{ o *ld.Object }

func NewLink(id string) Link { return AsLink(ld.NewObject(id, Class_Link.ID)) }

// Ducktypes the object into a(n) Link.
func AsLink(e ld.Entity) Link { return Link{o: e.Obj()} }

// Does the object quack like a(n) Link?
func IsLink(e ld.Entity) bool { return ld.Is(e, Class_Link.ID) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Link) Obj() *ld.Object { return obj.o }

// Returns whether the wrapped object is null. Implements ld.Entity.
func (obj Link) IsNull() bool { return obj.o == nil }

// Returns the object's @id. Implements ld.Entity.
func (obj Link) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Link) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Link) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Link) Get(key string) interface{} { return obj.o.Get(key) }

// Sets the named attribute. Implements ld.Entity.
func (obj Link) Set(key string, v interface{}) { obj.o.Set(key, v) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Link) Apply(other ld.Entity, mergeArrays bool) error { return obj.o.Apply(other, mergeArrays) }

// Identifies an entity to which an object is attributed
func (obj Link) AttributedTo() interface{} { return GetAttributedTo(obj) }

func (obj Link) SetAttributedTo(v interface{}) { SetAttributedTo(obj, v) }

// The display height expressed as device independent pixels
func (obj Link) Height() interface{} { return GetHeight(obj) }

func (obj Link) SetHeight(v interface{}) { SetHeight(obj, v) }

// The target URI of the Link
func (obj Link) Href() interface{} { return GetHref(obj) }

func (obj Link) SetHref(v interface{}) { SetHref(obj, v) }

// A hint about the language of the referenced resource
func (obj Link) Hreflang() interface{} { return GetHreflang(obj) }

func (obj Link) SetHreflang(v interface{}) { SetHreflang(obj, v) }

func (obj Link) Id() interface{} { return GetId(obj) }

func (obj Link) SetId(v interface{}) { SetId(obj, v) }

// The MIME Media Type
func (obj Link) MediaType() interface{} { return GetMediaType(obj) }

func (obj Link) SetMediaType(v interface{}) { SetMediaType(obj, v) }

func (obj Link) Name() interface{} { return GetName(obj) }

func (obj Link) SetName(v interface{}) { SetName(obj, v) }

func (obj Link) Object() interface{} { return GetObject(obj) }

func (obj Link) SetObject(v interface{}) { SetObject(obj, v) }

func (obj Link) Preview() interface{} { return GetPreview(obj) }

func (obj Link) SetPreview(v interface{}) { SetPreview(obj, v) }

// The RFC 5988 or HTML5 Link Relation associated with the Link
func (obj Link) Rel() interface{} { return GetRel(obj) }

func (obj Link) SetRel(v interface{}) { SetRel(obj, v) }

// Specifies the preferred display width of the content, expressed in terms of device independent pixels.
func (obj Link) Width() interface{} { return GetWidth(obj) }

func (obj Link) SetWidth(v interface{}) { SetWidth(obj, v) }

// The actor listened to the object
type Listen struct{ Activity }

func NewListen(id string) Listen { return AsListen(ld.NewObject(id, Class_Listen.ID)) }

// Ducktypes the object into a(n) Listen.
func AsListen(e ld.Entity) Listen { return Listen{AsActivity(e)} }

// Does the object quack like a(n) Listen?
func IsListen(e ld.Entity) bool { return ld.Is(e, Class_Listen.ID) }

// A specialized Link that represents an @mention
type Mention struct{ Link }

func NewMention(id string) Mention { return AsMention(ld.NewObject(id, Class_Mention.ID)) }

// Ducktypes the object into a(n) Mention.
func AsMention(e ld.Entity) Mention { return Mention{AsLink(e)} }

// Does the object quack like a(n) Mention?
func IsMention(e ld.Entity) bool { return ld.Is(e, Class_Mention.ID) }

// The actor is moving the object. The target specifies where the object is moving to. The origin specifies where the object is moving from.
type Move struct{ Activity }

func NewMove(id string) Move { return AsMove(ld.NewObject(id, Class_Move.ID)) }

// Ducktypes the object into a(n) Move.
func AsMove(e ld.Entity) Move { return Move{AsActivity(e)} }

// Does the object quack like a(n) Move?
func IsMove(e ld.Entity) bool { return ld.Is(e, Class_Move.ID) }

// A Short note, typically less than a single paragraph. A "tweet" is an example, or a "status update"
type Note struct{ Object }

func NewNote(id string) Note { return AsNote(ld.NewObject(id, Class_Note.ID)) }

// Ducktypes the object into a(n) Note.
func AsNote(e ld.Entity) Note { return Note{AsObject(e)} }

// Does the object quack like a(n) Note?
func IsNote(e ld.Entity) bool { return ld.Is(e, Class_Note.ID) }

type Object struct{ o *ld.Object }

func NewObject(id string) Object { return AsObject(ld.NewObject(id, Class_Object.ID)) }

// Ducktypes the object into a(n) Object.
func AsObject(e ld.Entity) Object { return Object{o: e.Obj()} }

// Does the object quack like a(n) Object?
func IsObject(e ld.Entity) bool { return ld.Is(e, Class_Object.ID) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj Object) Obj() *ld.Object { return obj.o }

// Returns whether the wrapped object is null. Implements ld.Entity.
func (obj Object) IsNull() bool { return obj.o == nil }

// Returns the object's @id. Implements ld.Entity.
func (obj Object) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj Object) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj Object) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj Object) Get(key string) interface{} { return obj.o.Get(key) }

// Sets the named attribute. Implements ld.Entity.
func (obj Object) Set(key string, v interface{}) { obj.o.Set(key, v) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj Object) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

func (obj Object) Attachment() interface{} { return GetAttachment(obj) }

func (obj Object) SetAttachment(v interface{}) { SetAttachment(obj, v) }

func (obj Object) Attachments() interface{} { return GetAttachments(obj) }

func (obj Object) SetAttachments(v interface{}) { SetAttachments(obj, v) }

// Identifies an entity to which an object is attributed
func (obj Object) AttributedTo() interface{} { return GetAttributedTo(obj) }

func (obj Object) SetAttributedTo(v interface{}) { SetAttributedTo(obj, v) }

func (obj Object) Audience() interface{} { return GetAudience(obj) }

func (obj Object) SetAudience(v interface{}) { SetAudience(obj, v) }

// Identifies the author of an object. Deprecated. Use as:attributedTo instead
func (obj Object) Author() interface{} { return GetAuthor(obj) }

func (obj Object) SetAuthor(v interface{}) { SetAuthor(obj, v) }

func (obj Object) Bcc() interface{} { return GetBcc(obj) }

func (obj Object) SetBcc(v interface{}) { SetBcc(obj, v) }

func (obj Object) Bto() interface{} { return GetBto(obj) }

func (obj Object) SetBto(v interface{}) { SetBto(obj, v) }

func (obj Object) Cc() interface{} { return GetCc(obj) }

func (obj Object) SetCc(v interface{}) { SetCc(obj, v) }

// The content of the object.
func (obj Object) Content() interface{} { return GetContent(obj) }

func (obj Object) SetContent(v interface{}) { SetContent(obj, v) }

// Specifies the context within which an object exists or an activity was performed
func (obj Object) Context() interface{} { return GetContext(obj) }

func (obj Object) SetContext(v interface{}) { SetContext(obj, v) }

func (obj Object) DownstreamDuplicates() interface{} { return GetDownstreamDuplicates(obj) }

func (obj Object) SetDownstreamDuplicates(v interface{}) { SetDownstreamDuplicates(obj, v) }

// The duration of the object
func (obj Object) Duration() interface{} { return GetDuration(obj) }

func (obj Object) SetDuration(v interface{}) { SetDuration(obj, v) }

// The ending time of the object
func (obj Object) EndTime() interface{} { return GetEndTime(obj) }

func (obj Object) SetEndTime(v interface{}) { SetEndTime(obj, v) }

func (obj Object) Generator() interface{} { return GetGenerator(obj) }

func (obj Object) SetGenerator(v interface{}) { SetGenerator(obj, v) }

func (obj Object) Icon() interface{} { return GetIcon(obj) }

func (obj Object) SetIcon(v interface{}) { SetIcon(obj, v) }

func (obj Object) Id() interface{} { return GetId(obj) }

func (obj Object) SetId(v interface{}) { SetId(obj, v) }

func (obj Object) Image() interface{} { return GetImage(obj) }

func (obj Object) SetImage(v interface{}) { SetImage(obj, v) }

func (obj Object) InReplyTo() interface{} { return GetInReplyTo(obj) }

func (obj Object) SetInReplyTo(v interface{}) { SetInReplyTo(obj, v) }

func (obj Object) Location() interface{} { return GetLocation(obj) }

func (obj Object) SetLocation(v interface{}) { SetLocation(obj, v) }

// The MIME Media Type
func (obj Object) MediaType() interface{} { return GetMediaType(obj) }

func (obj Object) SetMediaType(v interface{}) { SetMediaType(obj, v) }

func (obj Object) Name() interface{} { return GetName(obj) }

func (obj Object) SetName(v interface{}) { SetName(obj, v) }

func (obj Object) Object() interface{} { return GetObject(obj) }

func (obj Object) SetObject(v interface{}) { SetObject(obj, v) }

func (obj Object) ObjectType() interface{} { return GetObjectType(obj) }

func (obj Object) SetObjectType(v interface{}) { SetObjectType(obj, v) }

func (obj Object) Preview() interface{} { return GetPreview(obj) }

func (obj Object) SetPreview(v interface{}) { SetPreview(obj, v) }

func (obj Object) Provider() interface{} { return GetProvider(obj) }

func (obj Object) SetProvider(v interface{}) { SetProvider(obj, v) }

// Specifies the date and time the object was published
func (obj Object) Published() interface{} { return GetPublished(obj) }

func (obj Object) SetPublished(v interface{}) { SetPublished(obj, v) }

// A numeric rating (>= 0.0, <= 5.0) for the object
func (obj Object) Rating() interface{} { return GetRating(obj) }

func (obj Object) SetRating(v interface{}) { SetRating(obj, v) }

func (obj Object) Replies() interface{} { return GetReplies(obj) }

func (obj Object) SetReplies(v interface{}) { SetReplies(obj, v) }

// The starting time of the object
func (obj Object) StartTime() interface{} { return GetStartTime(obj) }

func (obj Object) SetStartTime(v interface{}) { SetStartTime(obj, v) }

// A short summary of the object
func (obj Object) Summary() interface{} { return GetSummary(obj) }

func (obj Object) SetSummary(v interface{}) { SetSummary(obj, v) }

func (obj Object) Tag() interface{} { return GetTag(obj) }

func (obj Object) SetTag(v interface{}) { SetTag(obj, v) }

func (obj Object) Tags() interface{} { return GetTags(obj) }

func (obj Object) SetTags(v interface{}) { SetTags(obj, v) }

func (obj Object) To() interface{} { return GetTo(obj) }

func (obj Object) SetTo(v interface{}) { SetTo(obj, v) }

// Specifies when the object was last updated
func (obj Object) Updated() interface{} { return GetUpdated(obj) }

func (obj Object) SetUpdated(v interface{}) { SetUpdated(obj, v) }

func (obj Object) UpstreamDuplicates() interface{} { return GetUpstreamDuplicates(obj) }

func (obj Object) SetUpstreamDuplicates(v interface{}) { SetUpstreamDuplicates(obj, v) }

// Specifies a link to a specific representation of the Object
func (obj Object) Url() interface{} { return GetUrl(obj) }

func (obj Object) SetUrl(v interface{}) { SetUrl(obj, v) }

// To Offer something to someone or something
type Offer struct{ Activity }

func NewOffer(id string) Offer { return AsOffer(ld.NewObject(id, Class_Offer.ID)) }

// Ducktypes the object into a(n) Offer.
func AsOffer(e ld.Entity) Offer { return Offer{AsActivity(e)} }

// Does the object quack like a(n) Offer?
func IsOffer(e ld.Entity) bool { return ld.Is(e, Class_Offer.ID) }

// A variation of Collection in which items are strictly ordered
type OrderedCollection struct{ o *ld.Object }

func NewOrderedCollection(id string) OrderedCollection {
	return AsOrderedCollection(ld.NewObject(id, Class_OrderedCollection.ID))
}

// Ducktypes the object into a(n) OrderedCollection.
func AsOrderedCollection(e ld.Entity) OrderedCollection { return OrderedCollection{o: e.Obj()} }

// Does the object quack like a(n) OrderedCollection?
func IsOrderedCollection(e ld.Entity) bool { return ld.Is(e, Class_OrderedCollection.ID) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj OrderedCollection) Obj() *ld.Object { return obj.o }

// Returns whether the wrapped object is null. Implements ld.Entity.
func (obj OrderedCollection) IsNull() bool { return obj.o == nil }

// Returns the object's @id. Implements ld.Entity.
func (obj OrderedCollection) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj OrderedCollection) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj OrderedCollection) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj OrderedCollection) Get(key string) interface{} { return obj.o.Get(key) }

// Sets the named attribute. Implements ld.Entity.
func (obj OrderedCollection) Set(key string, v interface{}) { obj.o.Set(key, v) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj OrderedCollection) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// Identifies an entity to which an object is attributed
func (obj OrderedCollection) AttributedTo() interface{} { return GetAttributedTo(obj) }

func (obj OrderedCollection) SetAttributedTo(v interface{}) { SetAttributedTo(obj, v) }

func (obj OrderedCollection) Id() interface{} { return GetId(obj) }

func (obj OrderedCollection) SetId(v interface{}) { SetId(obj, v) }

// The MIME Media Type
func (obj OrderedCollection) MediaType() interface{} { return GetMediaType(obj) }

func (obj OrderedCollection) SetMediaType(v interface{}) { SetMediaType(obj, v) }

func (obj OrderedCollection) Name() interface{} { return GetName(obj) }

func (obj OrderedCollection) SetName(v interface{}) { SetName(obj, v) }

func (obj OrderedCollection) Object() interface{} { return GetObject(obj) }

func (obj OrderedCollection) SetObject(v interface{}) { SetObject(obj, v) }

func (obj OrderedCollection) Preview() interface{} { return GetPreview(obj) }

func (obj OrderedCollection) SetPreview(v interface{}) { SetPreview(obj, v) }

// An ordered subset of items from an OrderedCollection
type OrderedCollectionPage struct {
	CollectionPage
	OrderedCollection
}

func NewOrderedCollectionPage(id string) OrderedCollectionPage {
	return AsOrderedCollectionPage(ld.NewObject(id, Class_OrderedCollectionPage.ID))
}

// Ducktypes the object into a(n) OrderedCollectionPage.
func AsOrderedCollectionPage(e ld.Entity) OrderedCollectionPage {
	return OrderedCollectionPage{AsCollectionPage(e), AsOrderedCollection(e)}
}

// Does the object quack like a(n) OrderedCollectionPage?
func IsOrderedCollectionPage(e ld.Entity) bool { return ld.Is(e, Class_OrderedCollectionPage.ID) }

// In a strictly ordered logical collection, specifies the index position of the first item in the items list
func (obj OrderedCollectionPage) StartIndex() interface{} { return GetStartIndex(obj) }

func (obj OrderedCollectionPage) SetStartIndex(v interface{}) { SetStartIndex(obj, v) }

// A rdf:List variant for Objects and Links
type OrderedItems struct{ o *ld.Object }

func NewOrderedItems(id string) OrderedItems {
	return AsOrderedItems(ld.NewObject(id, Class_OrderedItems.ID))
}

// Ducktypes the object into a(n) OrderedItems.
func AsOrderedItems(e ld.Entity) OrderedItems { return OrderedItems{o: e.Obj()} }

// Does the object quack like a(n) OrderedItems?
func IsOrderedItems(e ld.Entity) bool { return ld.Is(e, Class_OrderedItems.ID) }

// Returns the wrapped plain ld.Object. Implements ld.Entity.
func (obj OrderedItems) Obj() *ld.Object { return obj.o }

// Returns whether the wrapped object is null. Implements ld.Entity.
func (obj OrderedItems) IsNull() bool { return obj.o == nil }

// Returns the object's @id. Implements ld.Entity.
func (obj OrderedItems) ID() string { return obj.o.ID() }

// Returns the object's @value. Implements ld.Entity.
func (obj OrderedItems) Value() string { return obj.o.Value() }

// Returns the object's @type. Implements ld.Entity.
func (obj OrderedItems) Type() []string { return obj.o.Type() }

// Returns the named attribute. Implements ld.Entity.
func (obj OrderedItems) Get(key string) interface{} { return obj.o.Get(key) }

// Sets the named attribute. Implements ld.Entity.
func (obj OrderedItems) Set(key string, v interface{}) { obj.o.Set(key, v) }

// Applies another object as a patch to this one. Implements ld.Entity.
func (obj OrderedItems) Apply(other ld.Entity, mergeArrays bool) error {
	return obj.o.Apply(other, mergeArrays)
}

// Identifies an entity to which an object is attributed
func (obj OrderedItems) AttributedTo() interface{} { return GetAttributedTo(obj) }

func (obj OrderedItems) SetAttributedTo(v interface{}) { SetAttributedTo(obj, v) }

func (obj OrderedItems) Id() interface{} { return GetId(obj) }

func (obj OrderedItems) SetId(v interface{}) { SetId(obj, v) }

// The MIME Media Type
func (obj OrderedItems) MediaType() interface{} { return GetMediaType(obj) }

func (obj OrderedItems) SetMediaType(v interface{}) { SetMediaType(obj, v) }

func (obj OrderedItems) Name() interface{} { return GetName(obj) }

func (obj OrderedItems) SetName(v interface{}) { SetName(obj, v) }

func (obj OrderedItems) Object() interface{} { return GetObject(obj) }

func (obj OrderedItems) SetObject(v interface{}) { SetObject(obj, v) }

func (obj OrderedItems) Preview() interface{} { return GetPreview(obj) }

func (obj OrderedItems) SetPreview(v interface{}) { SetPreview(obj, v) }

// An Organization
type Organization struct{ Object }

func NewOrganization(id string) Organization {
	return AsOrganization(ld.NewObject(id, Class_Organization.ID))
}

// Ducktypes the object into a(n) Organization.
func AsOrganization(e ld.Entity) Organization { return Organization{AsObject(e)} }

// Does the object quack like a(n) Organization?
func IsOrganization(e ld.Entity) bool { return ld.Is(e, Class_Organization.ID) }

// A Web Page
type Page struct{ Object }

func NewPage(id string) Page { return AsPage(ld.NewObject(id, Class_Page.ID)) }

// Ducktypes the object into a(n) Page.
func AsPage(e ld.Entity) Page { return Page{AsObject(e)} }

// Does the object quack like a(n) Page?
func IsPage(e ld.Entity) bool { return ld.Is(e, Class_Page.ID) }

// A Person
type Person struct{ Object }

func NewPerson(id string) Person { return AsPerson(ld.NewObject(id, Class_Person.ID)) }

// Ducktypes the object into a(n) Person.
func AsPerson(e ld.Entity) Person { return Person{AsObject(e)} }

// Does the object quack like a(n) Person?
func IsPerson(e ld.Entity) bool { return ld.Is(e, Class_Person.ID) }

// A physical or logical location
type Place struct{ Object }

func NewPlace(id string) Place { return AsPlace(ld.NewObject(id, Class_Place.ID)) }

// Ducktypes the object into a(n) Place.
func AsPlace(e ld.Entity) Place { return Place{AsObject(e)} }

// Does the object quack like a(n) Place?
func IsPlace(e ld.Entity) bool { return ld.Is(e, Class_Place.ID) }

// Specifies the accuracy around the point established by the longitude and latitude
func (obj Place) Accuracy() interface{} { return GetAccuracy(obj) }

func (obj Place) SetAccuracy(v interface{}) { SetAccuracy(obj, v) }

// The altitude of a place
func (obj Place) Altitude() interface{} { return GetAltitude(obj) }

func (obj Place) SetAltitude(v interface{}) { SetAltitude(obj, v) }

// The latitude
func (obj Place) Latitude() interface{} { return GetLatitude(obj) }

func (obj Place) SetLatitude(v interface{}) { SetLatitude(obj, v) }

// The longitude
func (obj Place) Longitude() interface{} { return GetLongitude(obj) }

func (obj Place) SetLongitude(v interface{}) { SetLongitude(obj, v) }

// Specifies a radius around the point established by the longitude and latitude
func (obj Place) Radius() interface{} { return GetRadius(obj) }

func (obj Place) SetRadius(v interface{}) { SetRadius(obj, v) }

// Identifies the unit of measurement used by the radius, altitude and accuracy properties. The value can be expressed either as one of a set of predefined units or as a well-known common URI that identifies units.
func (obj Place) Units() interface{} { return GetUnits(obj) }

func (obj Place) SetUnits(v interface{}) { SetUnits(obj, v) }

// A Profile Document
type Profile struct{ Object }

func NewProfile(id string) Profile { return AsProfile(ld.NewObject(id, Class_Profile.ID)) }

// Ducktypes the object into a(n) Profile.
func AsProfile(e ld.Entity) Profile { return Profile{AsObject(e)} }

// Does the object quack like a(n) Profile?
func IsProfile(e ld.Entity) bool { return ld.Is(e, Class_Profile.ID) }

// On a Profile object, describes the object described by the profile
func (obj Profile) Describes() interface{} { return GetDescribes(obj) }

func (obj Profile) SetDescribes(v interface{}) { SetDescribes(obj, v) }

// A question of any sort.
type Question struct{ IntransitiveActivity }

func NewQuestion(id string) Question { return AsQuestion(ld.NewObject(id, Class_Question.ID)) }

// Ducktypes the object into a(n) Question.
func AsQuestion(e ld.Entity) Question { return Question{AsIntransitiveActivity(e)} }

// Does the object quack like a(n) Question?
func IsQuestion(e ld.Entity) bool { return ld.Is(e, Class_Question.ID) }

// Describes a possible inclusive answer or option for a question.
func (obj Question) AnyOf() interface{} { return GetAnyOf(obj) }

func (obj Question) SetAnyOf(v interface{}) { SetAnyOf(obj, v) }

// Describes a possible exclusive answer or option for a question.
func (obj Question) OneOf() interface{} { return GetOneOf(obj) }

func (obj Question) SetOneOf(v interface{}) { SetOneOf(obj, v) }

// The actor read the object
type Read struct{ Activity }

func NewRead(id string) Read { return AsRead(ld.NewObject(id, Class_Read.ID)) }

// Ducktypes the object into a(n) Read.
func AsRead(e ld.Entity) Read { return Read{AsActivity(e)} }

// Does the object quack like a(n) Read?
func IsRead(e ld.Entity) bool { return ld.Is(e, Class_Read.ID) }

// Actor rejects the Object
type Reject struct{ Activity }

func NewReject(id string) Reject { return AsReject(ld.NewObject(id, Class_Reject.ID)) }

// Ducktypes the object into a(n) Reject.
func AsReject(e ld.Entity) Reject { return Reject{AsActivity(e)} }

// Does the object quack like a(n) Reject?
func IsReject(e ld.Entity) bool { return ld.Is(e, Class_Reject.ID) }

// Represents a Social Graph relationship between two Individuals (indicated by the 'a' and 'b' properties)
type Relationship struct{ Object }

func NewRelationship(id string) Relationship {
	return AsRelationship(ld.NewObject(id, Class_Relationship.ID))
}

// Ducktypes the object into a(n) Relationship.
func AsRelationship(e ld.Entity) Relationship { return Relationship{AsObject(e)} }

// Does the object quack like a(n) Relationship?
func IsRelationship(e ld.Entity) bool { return ld.Is(e, Class_Relationship.ID) }

// On a Relationship object, describes the type of relationship
func (obj Relationship) Relationship() interface{} { return GetRelationship(obj) }

func (obj Relationship) SetRelationship(v interface{}) { SetRelationship(obj, v) }

// On a Relationship object, identifies the subject. e.g. when saying "John is connected to Sally", 'subject' refers to 'John'
func (obj Relationship) Subject() interface{} { return GetSubject(obj) }

func (obj Relationship) SetSubject(v interface{}) { SetSubject(obj, v) }

// To Remove Something
type Remove struct{ Activity }

func NewRemove(id string) Remove { return AsRemove(ld.NewObject(id, Class_Remove.ID)) }

// Ducktypes the object into a(n) Remove.
func AsRemove(e ld.Entity) Remove { return Remove{AsActivity(e)} }

// Does the object quack like a(n) Remove?
func IsRemove(e ld.Entity) bool { return ld.Is(e, Class_Remove.ID) }

// A service provided by some entity
type Service struct{ Object }

func NewService(id string) Service { return AsService(ld.NewObject(id, Class_Service.ID)) }

// Ducktypes the object into a(n) Service.
func AsService(e ld.Entity) Service { return Service{AsObject(e)} }

// Does the object quack like a(n) Service?
func IsService(e ld.Entity) bool { return ld.Is(e, Class_Service.ID) }

// Actor tentatively accepts the Object
type TentativeAccept struct{ Accept }

func NewTentativeAccept(id string) TentativeAccept {
	return AsTentativeAccept(ld.NewObject(id, Class_TentativeAccept.ID))
}

// Ducktypes the object into a(n) TentativeAccept.
func AsTentativeAccept(e ld.Entity) TentativeAccept { return TentativeAccept{AsAccept(e)} }

// Does the object quack like a(n) TentativeAccept?
func IsTentativeAccept(e ld.Entity) bool { return ld.Is(e, Class_TentativeAccept.ID) }

// Actor tentatively rejects the object
type TentativeReject struct{ Reject }

func NewTentativeReject(id string) TentativeReject {
	return AsTentativeReject(ld.NewObject(id, Class_TentativeReject.ID))
}

// Ducktypes the object into a(n) TentativeReject.
func AsTentativeReject(e ld.Entity) TentativeReject { return TentativeReject{AsReject(e)} }

// Does the object quack like a(n) TentativeReject?
func IsTentativeReject(e ld.Entity) bool { return ld.Is(e, Class_TentativeReject.ID) }

// A placeholder for a deleted object
type Tombstone struct{ Object }

func NewTombstone(id string) Tombstone { return AsTombstone(ld.NewObject(id, Class_Tombstone.ID)) }

// Ducktypes the object into a(n) Tombstone.
func AsTombstone(e ld.Entity) Tombstone { return Tombstone{AsObject(e)} }

// Does the object quack like a(n) Tombstone?
func IsTombstone(e ld.Entity) bool { return ld.Is(e, Class_Tombstone.ID) }

// Specifies the date and time the object was deleted
func (obj Tombstone) Deleted() interface{} { return GetDeleted(obj) }

func (obj Tombstone) SetDeleted(v interface{}) { SetDeleted(obj, v) }

// On a Tombstone object, describes the former type of the deleted object
func (obj Tombstone) FormerType() interface{} { return GetFormerType(obj) }

func (obj Tombstone) SetFormerType(v interface{}) { SetFormerType(obj, v) }

// The actor is traveling to the target. The origin specifies where the actor is traveling from.
type Travel struct{ IntransitiveActivity }

func NewTravel(id string) Travel { return AsTravel(ld.NewObject(id, Class_Travel.ID)) }

// Ducktypes the object into a(n) Travel.
func AsTravel(e ld.Entity) Travel { return Travel{AsIntransitiveActivity(e)} }

// Does the object quack like a(n) Travel?
func IsTravel(e ld.Entity) bool { return ld.Is(e, Class_Travel.ID) }

// To Undo Something. This would typically be used to indicate that a previous Activity has been undone.
type Undo struct{ Activity }

func NewUndo(id string) Undo { return AsUndo(ld.NewObject(id, Class_Undo.ID)) }

// Ducktypes the object into a(n) Undo.
func AsUndo(e ld.Entity) Undo { return Undo{AsActivity(e)} }

// Does the object quack like a(n) Undo?
func IsUndo(e ld.Entity) bool { return ld.Is(e, Class_Undo.ID) }

// To Update/Modify Something
type Update struct{ Activity }

func NewUpdate(id string) Update { return AsUpdate(ld.NewObject(id, Class_Update.ID)) }

// Ducktypes the object into a(n) Update.
func AsUpdate(e ld.Entity) Update { return Update{AsActivity(e)} }

// Does the object quack like a(n) Update?
func IsUpdate(e ld.Entity) bool { return ld.Is(e, Class_Update.ID) }

// A Video document of any kind.
type Video struct{ Document }

func NewVideo(id string) Video { return AsVideo(ld.NewObject(id, Class_Video.ID)) }

// Ducktypes the object into a(n) Video.
func AsVideo(e ld.Entity) Video { return Video{AsDocument(e)} }

// Does the object quack like a(n) Video?
func IsVideo(e ld.Entity) bool { return ld.Is(e, Class_Video.ID) }

// The actor viewed the object
type View struct{ Activity }

func NewView(id string) View { return AsView(ld.NewObject(id, Class_View.ID)) }

// Ducktypes the object into a(n) View.
func AsView(e ld.Entity) View { return View{AsActivity(e)} }

// Does the object quack like a(n) View?
func IsView(e ld.Entity) bool { return ld.Is(e, Class_View.ID) }

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
