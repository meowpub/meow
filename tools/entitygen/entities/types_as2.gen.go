package entities

// Actor accepts the Object
type AS2Accept struct{}

func (AS2Accept) Type() string {
	return "http://www.w3.org/ns/activitystreams#Accept"
}

// An Object representing some form of Action that has been taken
type AS2Activity struct{}

func (AS2Activity) Type() string {
	return "http://www.w3.org/ns/activitystreams#Activity"
}

// To Add an Object or Link to Something
type AS2Add struct{}

func (AS2Add) Type() string {
	return "http://www.w3.org/ns/activitystreams#Add"
}

// Actor announces the object to the target
type AS2Announce struct{}

func (AS2Announce) Type() string {
	return "http://www.w3.org/ns/activitystreams#Announce"
}

// Represents a software application of any sort
type AS2Application struct{}

func (AS2Application) Type() string {
	return "http://www.w3.org/ns/activitystreams#Application"
}

// To Arrive Somewhere (can be used, for instance, to indicate that a particular entity is currently located somewhere, e.g. a "check-in")
type AS2Arrive struct{}

func (AS2Arrive) Type() string {
	return "http://www.w3.org/ns/activitystreams#Arrive"
}

// A written work. Typically several paragraphs long. For example, a blog post or a news article.
type AS2Article struct{}

func (AS2Article) Type() string {
	return "http://www.w3.org/ns/activitystreams#Article"
}

// An audio file
type AS2Audio struct{}

func (AS2Audio) Type() string {
	return "http://www.w3.org/ns/activitystreams#Audio"
}

//
type AS2Block struct{}

func (AS2Block) Type() string {
	return "http://www.w3.org/ns/activitystreams#Block"
}

// An ordered or unordered collection of Objects or Links
type AS2Collection struct{}

func (AS2Collection) Type() string {
	return "http://www.w3.org/ns/activitystreams#Collection"
}

// A subset of items from a Collection
type AS2CollectionPage struct{}

func (AS2CollectionPage) Type() string {
	return "http://www.w3.org/ns/activitystreams#CollectionPage"
}

// To Create Something
type AS2Create struct{}

func (AS2Create) Type() string {
	return "http://www.w3.org/ns/activitystreams#Create"
}

// To Delete Something
type AS2Delete struct{}

func (AS2Delete) Type() string {
	return "http://www.w3.org/ns/activitystreams#Delete"
}

// The actor dislikes the object
type AS2Dislike struct{}

func (AS2Dislike) Type() string {
	return "http://www.w3.org/ns/activitystreams#Dislike"
}

// Represents a digital document/file of any sort
type AS2Document struct{}

func (AS2Document) Type() string {
	return "http://www.w3.org/ns/activitystreams#Document"
}

// An Event of any kind
type AS2Event struct{}

func (AS2Event) Type() string {
	return "http://www.w3.org/ns/activitystreams#Event"
}

// To flag something (e.g. flag as inappropriate, flag as spam, etc)
type AS2Flag struct{}

func (AS2Flag) Type() string {
	return "http://www.w3.org/ns/activitystreams#Flag"
}

// To Express Interest in Something
type AS2Follow struct{}

func (AS2Follow) Type() string {
	return "http://www.w3.org/ns/activitystreams#Follow"
}

// A Group of any kind.
type AS2Group struct{}

func (AS2Group) Type() string {
	return "http://www.w3.org/ns/activitystreams#Group"
}

// Actor is ignoring the Object
type AS2Ignore struct{}

func (AS2Ignore) Type() string {
	return "http://www.w3.org/ns/activitystreams#Ignore"
}

// An Image file
type AS2Image struct{}

func (AS2Image) Type() string {
	return "http://www.w3.org/ns/activitystreams#Image"
}

// An Activity that has no direct object
type AS2IntransitiveActivity struct{}

func (AS2IntransitiveActivity) Type() string {
	return "http://www.w3.org/ns/activitystreams#IntransitiveActivity"
}

// To invite someone or something to something
type AS2Invite struct{}

func (AS2Invite) Type() string {
	return "http://www.w3.org/ns/activitystreams#Invite"
}

// To Join Something
type AS2Join struct{}

func (AS2Join) Type() string {
	return "http://www.w3.org/ns/activitystreams#Join"
}

// To Leave Something
type AS2Leave struct{}

func (AS2Leave) Type() string {
	return "http://www.w3.org/ns/activitystreams#Leave"
}

// To Like Something
type AS2Like struct{}

func (AS2Like) Type() string {
	return "http://www.w3.org/ns/activitystreams#Like"
}

// Represents a qualified reference to another resource. Patterned after the RFC5988 Web Linking Model
type AS2Link struct{}

func (AS2Link) Type() string {
	return "http://www.w3.org/ns/activitystreams#Link"
}

// The actor listened to the object
type AS2Listen struct{}

func (AS2Listen) Type() string {
	return "http://www.w3.org/ns/activitystreams#Listen"
}

// A specialized Link that represents an @mention
type AS2Mention struct{}

func (AS2Mention) Type() string {
	return "http://www.w3.org/ns/activitystreams#Mention"
}

// The actor is moving the object. The target specifies where the object is moving to. The origin specifies where the object is moving from.
type AS2Move struct{}

func (AS2Move) Type() string {
	return "http://www.w3.org/ns/activitystreams#Move"
}

// A Short note, typically less than a single paragraph. A "tweet" is an example, or a "status update"
type AS2Note struct{}

func (AS2Note) Type() string {
	return "http://www.w3.org/ns/activitystreams#Note"
}

//
type AS2Object struct{}

func (AS2Object) Type() string {
	return "http://www.w3.org/ns/activitystreams#Object"
}

// To Offer something to someone or something
type AS2Offer struct{}

func (AS2Offer) Type() string {
	return "http://www.w3.org/ns/activitystreams#Offer"
}

// A variation of Collection in which items are strictly ordered
type AS2OrderedCollection struct{}

func (AS2OrderedCollection) Type() string {
	return "http://www.w3.org/ns/activitystreams#OrderedCollection"
}

// An ordered subset of items from an OrderedCollection
type AS2OrderedCollectionPage struct{}

func (AS2OrderedCollectionPage) Type() string {
	return "http://www.w3.org/ns/activitystreams#OrderedCollectionPage"
}

// A rdf:List variant for Objects and Links
type AS2OrderedItems struct{}

func (AS2OrderedItems) Type() string {
	return "http://www.w3.org/ns/activitystreams#OrderedItems"
}

// An Organization
type AS2Organization struct{}

func (AS2Organization) Type() string {
	return "http://www.w3.org/ns/activitystreams#Organization"
}

// A Web Page
type AS2Page struct{}

func (AS2Page) Type() string {
	return "http://www.w3.org/ns/activitystreams#Page"
}

// A Person
type AS2Person struct{}

func (AS2Person) Type() string {
	return "http://www.w3.org/ns/activitystreams#Person"
}

// A physical or logical location
type AS2Place struct{}

func (AS2Place) Type() string {
	return "http://www.w3.org/ns/activitystreams#Place"
}

// A Profile Document
type AS2Profile struct{}

func (AS2Profile) Type() string {
	return "http://www.w3.org/ns/activitystreams#Profile"
}

// A question of any sort.
type AS2Question struct{}

func (AS2Question) Type() string {
	return "http://www.w3.org/ns/activitystreams#Question"
}

// The actor read the object
type AS2Read struct{}

func (AS2Read) Type() string {
	return "http://www.w3.org/ns/activitystreams#Read"
}

// Actor rejects the Object
type AS2Reject struct{}

func (AS2Reject) Type() string {
	return "http://www.w3.org/ns/activitystreams#Reject"
}

// Represents a Social Graph relationship between two Individuals (indicated by the 'a' and 'b' properties)
type AS2Relationship struct{}

func (AS2Relationship) Type() string {
	return "http://www.w3.org/ns/activitystreams#Relationship"
}

// To Remove Something
type AS2Remove struct{}

func (AS2Remove) Type() string {
	return "http://www.w3.org/ns/activitystreams#Remove"
}

// A service provided by some entity
type AS2Service struct{}

func (AS2Service) Type() string {
	return "http://www.w3.org/ns/activitystreams#Service"
}

// Actor tentatively accepts the Object
type AS2TentativeAccept struct{}

func (AS2TentativeAccept) Type() string {
	return "http://www.w3.org/ns/activitystreams#TentativeAccept"
}

// Actor tentatively rejects the object
type AS2TentativeReject struct{}

func (AS2TentativeReject) Type() string {
	return "http://www.w3.org/ns/activitystreams#TentativeReject"
}

// A placeholder for a deleted object
type AS2Tombstone struct{}

func (AS2Tombstone) Type() string {
	return "http://www.w3.org/ns/activitystreams#Tombstone"
}

// The actor is traveling to the target. The origin specifies where the actor is traveling from.
type AS2Travel struct{}

func (AS2Travel) Type() string {
	return "http://www.w3.org/ns/activitystreams#Travel"
}

// To Undo Something. This would typically be used to indicate that a previous Activity has been undone.
type AS2Undo struct{}

func (AS2Undo) Type() string {
	return "http://www.w3.org/ns/activitystreams#Undo"
}

// To Update/Modify Something
type AS2Update struct{}

func (AS2Update) Type() string {
	return "http://www.w3.org/ns/activitystreams#Update"
}

// A Video document of any kind.
type AS2Video struct{}

func (AS2Video) Type() string {
	return "http://www.w3.org/ns/activitystreams#Video"
}

// The actor viewed the object
type AS2View struct{}

func (AS2View) Type() string {
	return "http://www.w3.org/ns/activitystreams#View"
}

// Compile-time assertions to make sure all generated types are valid.
var (
	_ Type = AS2Accept{}
	_ Type = AS2Activity{}
	_ Type = AS2Add{}
	_ Type = AS2Announce{}
	_ Type = AS2Application{}
	_ Type = AS2Arrive{}
	_ Type = AS2Article{}
	_ Type = AS2Audio{}
	_ Type = AS2Block{}
	_ Type = AS2Collection{}
	_ Type = AS2CollectionPage{}
	_ Type = AS2Create{}
	_ Type = AS2Delete{}
	_ Type = AS2Dislike{}
	_ Type = AS2Document{}
	_ Type = AS2Event{}
	_ Type = AS2Flag{}
	_ Type = AS2Follow{}
	_ Type = AS2Group{}
	_ Type = AS2Ignore{}
	_ Type = AS2Image{}
	_ Type = AS2IntransitiveActivity{}
	_ Type = AS2Invite{}
	_ Type = AS2Join{}
	_ Type = AS2Leave{}
	_ Type = AS2Like{}
	_ Type = AS2Link{}
	_ Type = AS2Listen{}
	_ Type = AS2Mention{}
	_ Type = AS2Move{}
	_ Type = AS2Note{}
	_ Type = AS2Object{}
	_ Type = AS2Offer{}
	_ Type = AS2OrderedCollection{}
	_ Type = AS2OrderedCollectionPage{}
	_ Type = AS2OrderedItems{}
	_ Type = AS2Organization{}
	_ Type = AS2Page{}
	_ Type = AS2Person{}
	_ Type = AS2Place{}
	_ Type = AS2Profile{}
	_ Type = AS2Question{}
	_ Type = AS2Read{}
	_ Type = AS2Reject{}
	_ Type = AS2Relationship{}
	_ Type = AS2Remove{}
	_ Type = AS2Service{}
	_ Type = AS2TentativeAccept{}
	_ Type = AS2TentativeReject{}
	_ Type = AS2Tombstone{}
	_ Type = AS2Travel{}
	_ Type = AS2Undo{}
	_ Type = AS2Update{}
	_ Type = AS2Video{}
	_ Type = AS2View{}
)
