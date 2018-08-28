// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package as

import (
	"github.com/meowpub/meow/ld"
)

// Specifies the accuracy around the point established by the longitude and latitude
func GetAccuracy(e ld.Entity) interface{} { return e.Get(PropAccuracy) }

// Subproperty of as:attributedTo that identifies the primary actor
func GetActor(e ld.Entity) interface{} { return e.Get(PropActor) }

// The altitude of a place
func GetAltitude(e ld.Entity) interface{} { return e.Get(PropAltitude) }

// Describes a possible inclusive answer or option for a question.
func GetAnyOf(e ld.Entity) interface{} { return e.Get(PropAnyOf) }

func GetAttachment(e ld.Entity) interface{} { return e.Get(PropAttachment) }

func GetAttachments(e ld.Entity) interface{} { return e.Get(PropAttachments) }

// Identifies an entity to which an object is attributed
func GetAttributedTo(e ld.Entity) interface{} { return e.Get(PropAttributedTo) }

func GetAudience(e ld.Entity) interface{} { return e.Get(PropAudience) }

// Identifies the author of an object. Deprecated. Use as:attributedTo instead
func GetAuthor(e ld.Entity) interface{} { return e.Get(PropAuthor) }

func GetBcc(e ld.Entity) interface{} { return e.Get(PropBcc) }

func GetBto(e ld.Entity) interface{} { return e.Get(PropBto) }

func GetCc(e ld.Entity) interface{} { return e.Get(PropCc) }

// The content of the object.
func GetContent(e ld.Entity) interface{} { return e.Get(PropContent) }

// Specifies the context within which an object exists or an activity was performed
func GetContext(e ld.Entity) interface{} { return e.Get(PropContext) }

func GetCurrent(e ld.Entity) interface{} { return e.Get(PropCurrent) }

// Specifies the date and time the object was deleted
func GetDeleted(e ld.Entity) interface{} { return e.Get(PropDeleted) }

// On a Profile object, describes the object described by the profile
func GetDescribes(e ld.Entity) interface{} { return e.Get(PropDescribes) }

func GetDownstreamDuplicates(e ld.Entity) interface{} { return e.Get(PropDownstreamDuplicates) }

// The duration of the object
func GetDuration(e ld.Entity) interface{} { return e.Get(PropDuration) }

// The ending time of the object
func GetEndTime(e ld.Entity) interface{} { return e.Get(PropEndTime) }

func GetFirst(e ld.Entity) interface{} { return e.Get(PropFirst) }

// On a Tombstone object, describes the former type of the deleted object
func GetFormerType(e ld.Entity) interface{} { return e.Get(PropFormerType) }

func GetGenerator(e ld.Entity) interface{} { return e.Get(PropGenerator) }

// The display height expressed as device independent pixels
func GetHeight(e ld.Entity) interface{} { return e.Get(PropHeight) }

// The target URI of the Link
func GetHref(e ld.Entity) interface{} { return e.Get(PropHref) }

// A hint about the language of the referenced resource
func GetHreflang(e ld.Entity) interface{} { return e.Get(PropHreflang) }

func GetIcon(e ld.Entity) interface{} { return e.Get(PropIcon) }

func GetId(e ld.Entity) interface{} { return e.Get(PropId) }

func GetImage(e ld.Entity) interface{} { return e.Get(PropImage) }

func GetInReplyTo(e ld.Entity) interface{} { return e.Get(PropInReplyTo) }

// Indentifies an object used (or to be used) to complete an activity
func GetInstrument(e ld.Entity) interface{} { return e.Get(PropInstrument) }

func GetItems(e ld.Entity) interface{} { return e.Get(PropItems) }

func GetLast(e ld.Entity) interface{} { return e.Get(PropLast) }

// The latitude
func GetLatitude(e ld.Entity) interface{} { return e.Get(PropLatitude) }

func GetLocation(e ld.Entity) interface{} { return e.Get(PropLocation) }

// The longitude
func GetLongitude(e ld.Entity) interface{} { return e.Get(PropLongitude) }

// The MIME Media Type
func GetMediaType(e ld.Entity) interface{} { return e.Get(PropMediaType) }

func GetName(e ld.Entity) interface{} { return e.Get(PropName) }

func GetNext(e ld.Entity) interface{} { return e.Get(PropNext) }

func GetObject(e ld.Entity) interface{} { return e.Get(PropObject) }

func GetObjectType(e ld.Entity) interface{} { return e.Get(PropObjectType) }

// Describes a possible exclusive answer or option for a question.
func GetOneOf(e ld.Entity) interface{} { return e.Get(PropOneOf) }

// For certain activities, specifies the entity from which the action is directed.
func GetOrigin(e ld.Entity) interface{} { return e.Get(PropOrigin) }

func GetPartOf(e ld.Entity) interface{} { return e.Get(PropPartOf) }

func GetPrev(e ld.Entity) interface{} { return e.Get(PropPrev) }

func GetPreview(e ld.Entity) interface{} { return e.Get(PropPreview) }

func GetProvider(e ld.Entity) interface{} { return e.Get(PropProvider) }

// Specifies the date and time the object was published
func GetPublished(e ld.Entity) interface{} { return e.Get(PropPublished) }

// Specifies a radius around the point established by the longitude and latitude
func GetRadius(e ld.Entity) interface{} { return e.Get(PropRadius) }

// A numeric rating (>= 0.0, <= 5.0) for the object
func GetRating(e ld.Entity) interface{} { return e.Get(PropRating) }

// The RFC 5988 or HTML5 Link Relation associated with the Link
func GetRel(e ld.Entity) interface{} { return e.Get(PropRel) }

// On a Relationship object, describes the type of relationship
func GetRelationship(e ld.Entity) interface{} { return e.Get(PropRelationship) }

func GetReplies(e ld.Entity) interface{} { return e.Get(PropReplies) }

func GetResult(e ld.Entity) interface{} { return e.Get(PropResult) }

// In a strictly ordered logical collection, specifies the index position of the first item in the items list
func GetStartIndex(e ld.Entity) interface{} { return e.Get(PropStartIndex) }

// The starting time of the object
func GetStartTime(e ld.Entity) interface{} { return e.Get(PropStartTime) }

// On a Relationship object, identifies the subject. e.g. when saying "John is connected to Sally", 'subject' refers to 'John'
func GetSubject(e ld.Entity) interface{} { return e.Get(PropSubject) }

// A short summary of the object
func GetSummary(e ld.Entity) interface{} { return e.Get(PropSummary) }

func GetTag(e ld.Entity) interface{} { return e.Get(PropTag) }

func GetTags(e ld.Entity) interface{} { return e.Get(PropTags) }

func GetTarget(e ld.Entity) interface{} { return e.Get(PropTarget) }

func GetTo(e ld.Entity) interface{} { return e.Get(PropTo) }

// The total number of items in a logical collection
func GetTotalItems(e ld.Entity) interface{} { return e.Get(PropTotalItems) }

// Identifies the unit of measurement used by the radius, altitude and accuracy properties. The value can be expressed either as one of a set of predefined units or as a well-known common URI that identifies units.
func GetUnits(e ld.Entity) interface{} { return e.Get(PropUnits) }

// Specifies when the object was last updated
func GetUpdated(e ld.Entity) interface{} { return e.Get(PropUpdated) }

func GetUpstreamDuplicates(e ld.Entity) interface{} { return e.Get(PropUpstreamDuplicates) }

// Specifies a link to a specific representation of the Object
func GetUrl(e ld.Entity) interface{} { return e.Get(PropUrl) }

func GetVerb(e ld.Entity) interface{} { return e.Get(PropVerb) }

// Specifies the preferred display width of the content, expressed in terms of device independent pixels.
func GetWidth(e ld.Entity) interface{} { return e.Get(PropWidth) }
