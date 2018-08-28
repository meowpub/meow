// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package as

import (
	"github.com/meowpub/meow/ld"
)

// Specifies the accuracy around the point established by the longitude and latitude
func GetAccuracy(e ld.Entity) interface{} { return e.Get(PropAccuracy) }

func SetAccuracy(e ld.Entity, v interface{}) { e.Set(PropAccuracy, v) }

// Subproperty of as:attributedTo that identifies the primary actor
func GetActor(e ld.Entity) interface{} { return e.Get(PropActor) }

func SetActor(e ld.Entity, v interface{}) { e.Set(PropActor, v) }

// The altitude of a place
func GetAltitude(e ld.Entity) interface{} { return e.Get(PropAltitude) }

func SetAltitude(e ld.Entity, v interface{}) { e.Set(PropAltitude, v) }

// Describes a possible inclusive answer or option for a question.
func GetAnyOf(e ld.Entity) interface{} { return e.Get(PropAnyOf) }

func SetAnyOf(e ld.Entity, v interface{}) { e.Set(PropAnyOf, v) }

func GetAttachment(e ld.Entity) interface{} { return e.Get(PropAttachment) }

func SetAttachment(e ld.Entity, v interface{}) { e.Set(PropAttachment, v) }

func GetAttachments(e ld.Entity) interface{} { return e.Get(PropAttachments) }

func SetAttachments(e ld.Entity, v interface{}) { e.Set(PropAttachments, v) }

// Identifies an entity to which an object is attributed
func GetAttributedTo(e ld.Entity) interface{} { return e.Get(PropAttributedTo) }

func SetAttributedTo(e ld.Entity, v interface{}) { e.Set(PropAttributedTo, v) }

func GetAudience(e ld.Entity) interface{} { return e.Get(PropAudience) }

func SetAudience(e ld.Entity, v interface{}) { e.Set(PropAudience, v) }

// Identifies the author of an object. Deprecated. Use as:attributedTo instead
func GetAuthor(e ld.Entity) interface{} { return e.Get(PropAuthor) }

func SetAuthor(e ld.Entity, v interface{}) { e.Set(PropAuthor, v) }

func GetBcc(e ld.Entity) interface{} { return e.Get(PropBcc) }

func SetBcc(e ld.Entity, v interface{}) { e.Set(PropBcc, v) }

func GetBto(e ld.Entity) interface{} { return e.Get(PropBto) }

func SetBto(e ld.Entity, v interface{}) { e.Set(PropBto, v) }

func GetCc(e ld.Entity) interface{} { return e.Get(PropCc) }

func SetCc(e ld.Entity, v interface{}) { e.Set(PropCc, v) }

// The content of the object.
func GetContent(e ld.Entity) interface{} { return e.Get(PropContent) }

func SetContent(e ld.Entity, v interface{}) { e.Set(PropContent, v) }

// Specifies the context within which an object exists or an activity was performed
func GetContext(e ld.Entity) interface{} { return e.Get(PropContext) }

func SetContext(e ld.Entity, v interface{}) { e.Set(PropContext, v) }

func GetCurrent(e ld.Entity) interface{} { return e.Get(PropCurrent) }

func SetCurrent(e ld.Entity, v interface{}) { e.Set(PropCurrent, v) }

// Specifies the date and time the object was deleted
func GetDeleted(e ld.Entity) interface{} { return e.Get(PropDeleted) }

func SetDeleted(e ld.Entity, v interface{}) { e.Set(PropDeleted, v) }

// On a Profile object, describes the object described by the profile
func GetDescribes(e ld.Entity) interface{} { return e.Get(PropDescribes) }

func SetDescribes(e ld.Entity, v interface{}) { e.Set(PropDescribes, v) }

func GetDownstreamDuplicates(e ld.Entity) interface{} { return e.Get(PropDownstreamDuplicates) }

func SetDownstreamDuplicates(e ld.Entity, v interface{}) { e.Set(PropDownstreamDuplicates, v) }

// The duration of the object
func GetDuration(e ld.Entity) interface{} { return e.Get(PropDuration) }

func SetDuration(e ld.Entity, v interface{}) { e.Set(PropDuration, v) }

// The ending time of the object
func GetEndTime(e ld.Entity) interface{} { return e.Get(PropEndTime) }

func SetEndTime(e ld.Entity, v interface{}) { e.Set(PropEndTime, v) }

func GetFirst(e ld.Entity) interface{} { return e.Get(PropFirst) }

func SetFirst(e ld.Entity, v interface{}) { e.Set(PropFirst, v) }

// On a Tombstone object, describes the former type of the deleted object
func GetFormerType(e ld.Entity) interface{} { return e.Get(PropFormerType) }

func SetFormerType(e ld.Entity, v interface{}) { e.Set(PropFormerType, v) }

func GetGenerator(e ld.Entity) interface{} { return e.Get(PropGenerator) }

func SetGenerator(e ld.Entity, v interface{}) { e.Set(PropGenerator, v) }

// The display height expressed as device independent pixels
func GetHeight(e ld.Entity) interface{} { return e.Get(PropHeight) }

func SetHeight(e ld.Entity, v interface{}) { e.Set(PropHeight, v) }

// The target URI of the Link
func GetHref(e ld.Entity) interface{} { return e.Get(PropHref) }

func SetHref(e ld.Entity, v interface{}) { e.Set(PropHref, v) }

// A hint about the language of the referenced resource
func GetHreflang(e ld.Entity) interface{} { return e.Get(PropHreflang) }

func SetHreflang(e ld.Entity, v interface{}) { e.Set(PropHreflang, v) }

func GetIcon(e ld.Entity) interface{} { return e.Get(PropIcon) }

func SetIcon(e ld.Entity, v interface{}) { e.Set(PropIcon, v) }

func GetId(e ld.Entity) interface{} { return e.Get(PropId) }

func SetId(e ld.Entity, v interface{}) { e.Set(PropId, v) }

func GetImage(e ld.Entity) interface{} { return e.Get(PropImage) }

func SetImage(e ld.Entity, v interface{}) { e.Set(PropImage, v) }

func GetInReplyTo(e ld.Entity) interface{} { return e.Get(PropInReplyTo) }

func SetInReplyTo(e ld.Entity, v interface{}) { e.Set(PropInReplyTo, v) }

// Indentifies an object used (or to be used) to complete an activity
func GetInstrument(e ld.Entity) interface{} { return e.Get(PropInstrument) }

func SetInstrument(e ld.Entity, v interface{}) { e.Set(PropInstrument, v) }

func GetItems(e ld.Entity) interface{} { return e.Get(PropItems) }

func SetItems(e ld.Entity, v interface{}) { e.Set(PropItems, v) }

func GetLast(e ld.Entity) interface{} { return e.Get(PropLast) }

func SetLast(e ld.Entity, v interface{}) { e.Set(PropLast, v) }

// The latitude
func GetLatitude(e ld.Entity) interface{} { return e.Get(PropLatitude) }

func SetLatitude(e ld.Entity, v interface{}) { e.Set(PropLatitude, v) }

func GetLocation(e ld.Entity) interface{} { return e.Get(PropLocation) }

func SetLocation(e ld.Entity, v interface{}) { e.Set(PropLocation, v) }

// The longitude
func GetLongitude(e ld.Entity) interface{} { return e.Get(PropLongitude) }

func SetLongitude(e ld.Entity, v interface{}) { e.Set(PropLongitude, v) }

// The MIME Media Type
func GetMediaType(e ld.Entity) interface{} { return e.Get(PropMediaType) }

func SetMediaType(e ld.Entity, v interface{}) { e.Set(PropMediaType, v) }

func GetName(e ld.Entity) interface{} { return e.Get(PropName) }

func SetName(e ld.Entity, v interface{}) { e.Set(PropName, v) }

func GetNext(e ld.Entity) interface{} { return e.Get(PropNext) }

func SetNext(e ld.Entity, v interface{}) { e.Set(PropNext, v) }

func GetObject(e ld.Entity) interface{} { return e.Get(PropObject) }

func SetObject(e ld.Entity, v interface{}) { e.Set(PropObject, v) }

func GetObjectType(e ld.Entity) interface{} { return e.Get(PropObjectType) }

func SetObjectType(e ld.Entity, v interface{}) { e.Set(PropObjectType, v) }

// Describes a possible exclusive answer or option for a question.
func GetOneOf(e ld.Entity) interface{} { return e.Get(PropOneOf) }

func SetOneOf(e ld.Entity, v interface{}) { e.Set(PropOneOf, v) }

// For certain activities, specifies the entity from which the action is directed.
func GetOrigin(e ld.Entity) interface{} { return e.Get(PropOrigin) }

func SetOrigin(e ld.Entity, v interface{}) { e.Set(PropOrigin, v) }

func GetPartOf(e ld.Entity) interface{} { return e.Get(PropPartOf) }

func SetPartOf(e ld.Entity, v interface{}) { e.Set(PropPartOf, v) }

func GetPrev(e ld.Entity) interface{} { return e.Get(PropPrev) }

func SetPrev(e ld.Entity, v interface{}) { e.Set(PropPrev, v) }

func GetPreview(e ld.Entity) interface{} { return e.Get(PropPreview) }

func SetPreview(e ld.Entity, v interface{}) { e.Set(PropPreview, v) }

func GetProvider(e ld.Entity) interface{} { return e.Get(PropProvider) }

func SetProvider(e ld.Entity, v interface{}) { e.Set(PropProvider, v) }

// Specifies the date and time the object was published
func GetPublished(e ld.Entity) interface{} { return e.Get(PropPublished) }

func SetPublished(e ld.Entity, v interface{}) { e.Set(PropPublished, v) }

// Specifies a radius around the point established by the longitude and latitude
func GetRadius(e ld.Entity) interface{} { return e.Get(PropRadius) }

func SetRadius(e ld.Entity, v interface{}) { e.Set(PropRadius, v) }

// A numeric rating (>= 0.0, <= 5.0) for the object
func GetRating(e ld.Entity) interface{} { return e.Get(PropRating) }

func SetRating(e ld.Entity, v interface{}) { e.Set(PropRating, v) }

// The RFC 5988 or HTML5 Link Relation associated with the Link
func GetRel(e ld.Entity) interface{} { return e.Get(PropRel) }

func SetRel(e ld.Entity, v interface{}) { e.Set(PropRel, v) }

// On a Relationship object, describes the type of relationship
func GetRelationship(e ld.Entity) interface{} { return e.Get(PropRelationship) }

func SetRelationship(e ld.Entity, v interface{}) { e.Set(PropRelationship, v) }

func GetReplies(e ld.Entity) interface{} { return e.Get(PropReplies) }

func SetReplies(e ld.Entity, v interface{}) { e.Set(PropReplies, v) }

func GetResult(e ld.Entity) interface{} { return e.Get(PropResult) }

func SetResult(e ld.Entity, v interface{}) { e.Set(PropResult, v) }

// In a strictly ordered logical collection, specifies the index position of the first item in the items list
func GetStartIndex(e ld.Entity) interface{} { return e.Get(PropStartIndex) }

func SetStartIndex(e ld.Entity, v interface{}) { e.Set(PropStartIndex, v) }

// The starting time of the object
func GetStartTime(e ld.Entity) interface{} { return e.Get(PropStartTime) }

func SetStartTime(e ld.Entity, v interface{}) { e.Set(PropStartTime, v) }

// On a Relationship object, identifies the subject. e.g. when saying "John is connected to Sally", 'subject' refers to 'John'
func GetSubject(e ld.Entity) interface{} { return e.Get(PropSubject) }

func SetSubject(e ld.Entity, v interface{}) { e.Set(PropSubject, v) }

// A short summary of the object
func GetSummary(e ld.Entity) interface{} { return e.Get(PropSummary) }

func SetSummary(e ld.Entity, v interface{}) { e.Set(PropSummary, v) }

func GetTag(e ld.Entity) interface{} { return e.Get(PropTag) }

func SetTag(e ld.Entity, v interface{}) { e.Set(PropTag, v) }

func GetTags(e ld.Entity) interface{} { return e.Get(PropTags) }

func SetTags(e ld.Entity, v interface{}) { e.Set(PropTags, v) }

func GetTarget(e ld.Entity) interface{} { return e.Get(PropTarget) }

func SetTarget(e ld.Entity, v interface{}) { e.Set(PropTarget, v) }

func GetTo(e ld.Entity) interface{} { return e.Get(PropTo) }

func SetTo(e ld.Entity, v interface{}) { e.Set(PropTo, v) }

// The total number of items in a logical collection
func GetTotalItems(e ld.Entity) interface{} { return e.Get(PropTotalItems) }

func SetTotalItems(e ld.Entity, v interface{}) { e.Set(PropTotalItems, v) }

// Identifies the unit of measurement used by the radius, altitude and accuracy properties. The value can be expressed either as one of a set of predefined units or as a well-known common URI that identifies units.
func GetUnits(e ld.Entity) interface{} { return e.Get(PropUnits) }

func SetUnits(e ld.Entity, v interface{}) { e.Set(PropUnits, v) }

// Specifies when the object was last updated
func GetUpdated(e ld.Entity) interface{} { return e.Get(PropUpdated) }

func SetUpdated(e ld.Entity, v interface{}) { e.Set(PropUpdated, v) }

func GetUpstreamDuplicates(e ld.Entity) interface{} { return e.Get(PropUpstreamDuplicates) }

func SetUpstreamDuplicates(e ld.Entity, v interface{}) { e.Set(PropUpstreamDuplicates, v) }

// Specifies a link to a specific representation of the Object
func GetUrl(e ld.Entity) interface{} { return e.Get(PropUrl) }

func SetUrl(e ld.Entity, v interface{}) { e.Set(PropUrl, v) }

func GetVerb(e ld.Entity) interface{} { return e.Get(PropVerb) }

func SetVerb(e ld.Entity, v interface{}) { e.Set(PropVerb, v) }

// Specifies the preferred display width of the content, expressed in terms of device independent pixels.
func GetWidth(e ld.Entity) interface{} { return e.Get(PropWidth) }

func SetWidth(e ld.Entity, v interface{}) { e.Set(PropWidth, v) }
