// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package as

import (
	"github.com/meowpub/meow/ld"
)

// Specifies the accuracy around the point established by the longitude and latitude
func GetAccuracy(e ld.Entity) interface{} { return e.Get(Prop_Accuracy.ID) }

func SetAccuracy(e ld.Entity, v interface{}) { e.Set(Prop_Accuracy.ID, v) }

// Subproperty of as:attributedTo that identifies the primary actor
func GetActor(e ld.Entity) interface{} { return e.Get(Prop_Actor.ID) }

func SetActor(e ld.Entity, v interface{}) { e.Set(Prop_Actor.ID, v) }

// The altitude of a place
func GetAltitude(e ld.Entity) interface{} { return e.Get(Prop_Altitude.ID) }

func SetAltitude(e ld.Entity, v interface{}) { e.Set(Prop_Altitude.ID, v) }

// Describes a possible inclusive answer or option for a question.
func GetAnyOf(e ld.Entity) interface{} { return e.Get(Prop_AnyOf.ID) }

func SetAnyOf(e ld.Entity, v interface{}) { e.Set(Prop_AnyOf.ID, v) }

func GetAttachment(e ld.Entity) interface{} { return e.Get(Prop_Attachment.ID) }

func SetAttachment(e ld.Entity, v interface{}) { e.Set(Prop_Attachment.ID, v) }

func GetAttachments(e ld.Entity) interface{} { return e.Get(Prop_Attachments.ID) }

func SetAttachments(e ld.Entity, v interface{}) { e.Set(Prop_Attachments.ID, v) }

// Identifies an entity to which an object is attributed
func GetAttributedTo(e ld.Entity) interface{} { return e.Get(Prop_AttributedTo.ID) }

func SetAttributedTo(e ld.Entity, v interface{}) { e.Set(Prop_AttributedTo.ID, v) }

func GetAudience(e ld.Entity) interface{} { return e.Get(Prop_Audience.ID) }

func SetAudience(e ld.Entity, v interface{}) { e.Set(Prop_Audience.ID, v) }

// Identifies the author of an object. Deprecated. Use as:attributedTo instead
func GetAuthor(e ld.Entity) interface{} { return e.Get(Prop_Author.ID) }

func SetAuthor(e ld.Entity, v interface{}) { e.Set(Prop_Author.ID, v) }

func GetBcc(e ld.Entity) interface{} { return e.Get(Prop_Bcc.ID) }

func SetBcc(e ld.Entity, v interface{}) { e.Set(Prop_Bcc.ID, v) }

func GetBto(e ld.Entity) interface{} { return e.Get(Prop_Bto.ID) }

func SetBto(e ld.Entity, v interface{}) { e.Set(Prop_Bto.ID, v) }

func GetCc(e ld.Entity) interface{} { return e.Get(Prop_Cc.ID) }

func SetCc(e ld.Entity, v interface{}) { e.Set(Prop_Cc.ID, v) }

// The content of the object.
func GetContent(e ld.Entity) interface{} { return e.Get(Prop_Content.ID) }

func SetContent(e ld.Entity, v interface{}) { e.Set(Prop_Content.ID, v) }

// Specifies the context within which an object exists or an activity was performed
func GetContext(e ld.Entity) interface{} { return e.Get(Prop_Context.ID) }

func SetContext(e ld.Entity, v interface{}) { e.Set(Prop_Context.ID, v) }

func GetCurrent(e ld.Entity) interface{} { return e.Get(Prop_Current.ID) }

func SetCurrent(e ld.Entity, v interface{}) { e.Set(Prop_Current.ID, v) }

// Specifies the date and time the object was deleted
func GetDeleted(e ld.Entity) interface{} { return e.Get(Prop_Deleted.ID) }

func SetDeleted(e ld.Entity, v interface{}) { e.Set(Prop_Deleted.ID, v) }

// On a Profile object, describes the object described by the profile
func GetDescribes(e ld.Entity) interface{} { return e.Get(Prop_Describes.ID) }

func SetDescribes(e ld.Entity, v interface{}) { e.Set(Prop_Describes.ID, v) }

func GetDownstreamDuplicates(e ld.Entity) interface{} { return e.Get(Prop_DownstreamDuplicates.ID) }

func SetDownstreamDuplicates(e ld.Entity, v interface{}) { e.Set(Prop_DownstreamDuplicates.ID, v) }

// The duration of the object
func GetDuration(e ld.Entity) interface{} { return e.Get(Prop_Duration.ID) }

func SetDuration(e ld.Entity, v interface{}) { e.Set(Prop_Duration.ID, v) }

// The ending time of the object
func GetEndTime(e ld.Entity) interface{} { return e.Get(Prop_EndTime.ID) }

func SetEndTime(e ld.Entity, v interface{}) { e.Set(Prop_EndTime.ID, v) }

func GetFirst(e ld.Entity) interface{} { return e.Get(Prop_First.ID) }

func SetFirst(e ld.Entity, v interface{}) { e.Set(Prop_First.ID, v) }

// On a Tombstone object, describes the former type of the deleted object
func GetFormerType(e ld.Entity) interface{} { return e.Get(Prop_FormerType.ID) }

func SetFormerType(e ld.Entity, v interface{}) { e.Set(Prop_FormerType.ID, v) }

func GetGenerator(e ld.Entity) interface{} { return e.Get(Prop_Generator.ID) }

func SetGenerator(e ld.Entity, v interface{}) { e.Set(Prop_Generator.ID, v) }

// The display height expressed as device independent pixels
func GetHeight(e ld.Entity) interface{} { return e.Get(Prop_Height.ID) }

func SetHeight(e ld.Entity, v interface{}) { e.Set(Prop_Height.ID, v) }

// The target URI of the Link
func GetHref(e ld.Entity) interface{} { return e.Get(Prop_Href.ID) }

func SetHref(e ld.Entity, v interface{}) { e.Set(Prop_Href.ID, v) }

// A hint about the language of the referenced resource
func GetHreflang(e ld.Entity) interface{} { return e.Get(Prop_Hreflang.ID) }

func SetHreflang(e ld.Entity, v interface{}) { e.Set(Prop_Hreflang.ID, v) }

func GetIcon(e ld.Entity) interface{} { return e.Get(Prop_Icon.ID) }

func SetIcon(e ld.Entity, v interface{}) { e.Set(Prop_Icon.ID, v) }

func GetId(e ld.Entity) interface{} { return e.Get(Prop_Id.ID) }

func SetId(e ld.Entity, v interface{}) { e.Set(Prop_Id.ID, v) }

func GetImage(e ld.Entity) interface{} { return e.Get(Prop_Image.ID) }

func SetImage(e ld.Entity, v interface{}) { e.Set(Prop_Image.ID, v) }

func GetInReplyTo(e ld.Entity) interface{} { return e.Get(Prop_InReplyTo.ID) }

func SetInReplyTo(e ld.Entity, v interface{}) { e.Set(Prop_InReplyTo.ID, v) }

// Indentifies an object used (or to be used) to complete an activity
func GetInstrument(e ld.Entity) interface{} { return e.Get(Prop_Instrument.ID) }

func SetInstrument(e ld.Entity, v interface{}) { e.Set(Prop_Instrument.ID, v) }

func GetItems(e ld.Entity) interface{} { return e.Get(Prop_Items.ID) }

func SetItems(e ld.Entity, v interface{}) { e.Set(Prop_Items.ID, v) }

func GetLast(e ld.Entity) interface{} { return e.Get(Prop_Last.ID) }

func SetLast(e ld.Entity, v interface{}) { e.Set(Prop_Last.ID, v) }

// The latitude
func GetLatitude(e ld.Entity) interface{} { return e.Get(Prop_Latitude.ID) }

func SetLatitude(e ld.Entity, v interface{}) { e.Set(Prop_Latitude.ID, v) }

func GetLocation(e ld.Entity) interface{} { return e.Get(Prop_Location.ID) }

func SetLocation(e ld.Entity, v interface{}) { e.Set(Prop_Location.ID, v) }

// The longitude
func GetLongitude(e ld.Entity) interface{} { return e.Get(Prop_Longitude.ID) }

func SetLongitude(e ld.Entity, v interface{}) { e.Set(Prop_Longitude.ID, v) }

// The MIME Media Type
func GetMediaType(e ld.Entity) interface{} { return e.Get(Prop_MediaType.ID) }

func SetMediaType(e ld.Entity, v interface{}) { e.Set(Prop_MediaType.ID, v) }

func GetName(e ld.Entity) interface{} { return e.Get(Prop_Name.ID) }

func SetName(e ld.Entity, v interface{}) { e.Set(Prop_Name.ID, v) }

func GetNext(e ld.Entity) interface{} { return e.Get(Prop_Next.ID) }

func SetNext(e ld.Entity, v interface{}) { e.Set(Prop_Next.ID, v) }

func GetObject(e ld.Entity) interface{} { return e.Get(Prop_Object.ID) }

func SetObject(e ld.Entity, v interface{}) { e.Set(Prop_Object.ID, v) }

func GetObjectType(e ld.Entity) interface{} { return e.Get(Prop_ObjectType.ID) }

func SetObjectType(e ld.Entity, v interface{}) { e.Set(Prop_ObjectType.ID, v) }

// Describes a possible exclusive answer or option for a question.
func GetOneOf(e ld.Entity) interface{} { return e.Get(Prop_OneOf.ID) }

func SetOneOf(e ld.Entity, v interface{}) { e.Set(Prop_OneOf.ID, v) }

// For certain activities, specifies the entity from which the action is directed.
func GetOrigin(e ld.Entity) interface{} { return e.Get(Prop_Origin.ID) }

func SetOrigin(e ld.Entity, v interface{}) { e.Set(Prop_Origin.ID, v) }

func GetPartOf(e ld.Entity) interface{} { return e.Get(Prop_PartOf.ID) }

func SetPartOf(e ld.Entity, v interface{}) { e.Set(Prop_PartOf.ID, v) }

func GetPrev(e ld.Entity) interface{} { return e.Get(Prop_Prev.ID) }

func SetPrev(e ld.Entity, v interface{}) { e.Set(Prop_Prev.ID, v) }

func GetPreview(e ld.Entity) interface{} { return e.Get(Prop_Preview.ID) }

func SetPreview(e ld.Entity, v interface{}) { e.Set(Prop_Preview.ID, v) }

func GetProvider(e ld.Entity) interface{} { return e.Get(Prop_Provider.ID) }

func SetProvider(e ld.Entity, v interface{}) { e.Set(Prop_Provider.ID, v) }

// Specifies the date and time the object was published
func GetPublished(e ld.Entity) interface{} { return e.Get(Prop_Published.ID) }

func SetPublished(e ld.Entity, v interface{}) { e.Set(Prop_Published.ID, v) }

// Specifies a radius around the point established by the longitude and latitude
func GetRadius(e ld.Entity) interface{} { return e.Get(Prop_Radius.ID) }

func SetRadius(e ld.Entity, v interface{}) { e.Set(Prop_Radius.ID, v) }

// A numeric rating (>= 0.0, <= 5.0) for the object
func GetRating(e ld.Entity) interface{} { return e.Get(Prop_Rating.ID) }

func SetRating(e ld.Entity, v interface{}) { e.Set(Prop_Rating.ID, v) }

// The RFC 5988 or HTML5 Link Relation associated with the Link
func GetRel(e ld.Entity) interface{} { return e.Get(Prop_Rel.ID) }

func SetRel(e ld.Entity, v interface{}) { e.Set(Prop_Rel.ID, v) }

// On a Relationship object, describes the type of relationship
func GetRelationship(e ld.Entity) interface{} { return e.Get(Prop_Relationship.ID) }

func SetRelationship(e ld.Entity, v interface{}) { e.Set(Prop_Relationship.ID, v) }

func GetReplies(e ld.Entity) interface{} { return e.Get(Prop_Replies.ID) }

func SetReplies(e ld.Entity, v interface{}) { e.Set(Prop_Replies.ID, v) }

func GetResult(e ld.Entity) interface{} { return e.Get(Prop_Result.ID) }

func SetResult(e ld.Entity, v interface{}) { e.Set(Prop_Result.ID, v) }

// In a strictly ordered logical collection, specifies the index position of the first item in the items list
func GetStartIndex(e ld.Entity) interface{} { return e.Get(Prop_StartIndex.ID) }

func SetStartIndex(e ld.Entity, v interface{}) { e.Set(Prop_StartIndex.ID, v) }

// The starting time of the object
func GetStartTime(e ld.Entity) interface{} { return e.Get(Prop_StartTime.ID) }

func SetStartTime(e ld.Entity, v interface{}) { e.Set(Prop_StartTime.ID, v) }

// On a Relationship object, identifies the subject. e.g. when saying "John is connected to Sally", 'subject' refers to 'John'
func GetSubject(e ld.Entity) interface{} { return e.Get(Prop_Subject.ID) }

func SetSubject(e ld.Entity, v interface{}) { e.Set(Prop_Subject.ID, v) }

// A short summary of the object
func GetSummary(e ld.Entity) interface{} { return e.Get(Prop_Summary.ID) }

func SetSummary(e ld.Entity, v interface{}) { e.Set(Prop_Summary.ID, v) }

func GetTag(e ld.Entity) interface{} { return e.Get(Prop_Tag.ID) }

func SetTag(e ld.Entity, v interface{}) { e.Set(Prop_Tag.ID, v) }

func GetTags(e ld.Entity) interface{} { return e.Get(Prop_Tags.ID) }

func SetTags(e ld.Entity, v interface{}) { e.Set(Prop_Tags.ID, v) }

func GetTarget(e ld.Entity) interface{} { return e.Get(Prop_Target.ID) }

func SetTarget(e ld.Entity, v interface{}) { e.Set(Prop_Target.ID, v) }

func GetTo(e ld.Entity) interface{} { return e.Get(Prop_To.ID) }

func SetTo(e ld.Entity, v interface{}) { e.Set(Prop_To.ID, v) }

// The total number of items in a logical collection
func GetTotalItems(e ld.Entity) interface{} { return e.Get(Prop_TotalItems.ID) }

func SetTotalItems(e ld.Entity, v interface{}) { e.Set(Prop_TotalItems.ID, v) }

// Identifies the unit of measurement used by the radius, altitude and accuracy properties. The value can be expressed either as one of a set of predefined units or as a well-known common URI that identifies units.
func GetUnits(e ld.Entity) interface{} { return e.Get(Prop_Units.ID) }

func SetUnits(e ld.Entity, v interface{}) { e.Set(Prop_Units.ID, v) }

// Specifies when the object was last updated
func GetUpdated(e ld.Entity) interface{} { return e.Get(Prop_Updated.ID) }

func SetUpdated(e ld.Entity, v interface{}) { e.Set(Prop_Updated.ID, v) }

func GetUpstreamDuplicates(e ld.Entity) interface{} { return e.Get(Prop_UpstreamDuplicates.ID) }

func SetUpstreamDuplicates(e ld.Entity, v interface{}) { e.Set(Prop_UpstreamDuplicates.ID, v) }

// Specifies a link to a specific representation of the Object
func GetUrl(e ld.Entity) interface{} { return e.Get(Prop_Url.ID) }

func SetUrl(e ld.Entity, v interface{}) { e.Set(Prop_Url.ID, v) }

func GetVerb(e ld.Entity) interface{} { return e.Get(Prop_Verb.ID) }

func SetVerb(e ld.Entity, v interface{}) { e.Set(Prop_Verb.ID, v) }

// Specifies the preferred display width of the content, expressed in terms of device independent pixels.
func GetWidth(e ld.Entity) interface{} { return e.Get(Prop_Width.ID) }

func SetWidth(e ld.Entity, v interface{}) { e.Set(Prop_Width.ID, v) }
