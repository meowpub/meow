// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package as

const (
	// Specifies the accuracy around the point established by the longitude and latitude
	PropAccuracy = "http://www.w3.org/ns/activitystreams#accuracy"

	// Subproperty of as:attributedTo that identifies the primary actor
	PropActor = "http://www.w3.org/ns/activitystreams#actor"

	// The altitude of a place
	PropAltitude = "http://www.w3.org/ns/activitystreams#altitude"

	// Describes a possible inclusive answer or option for a question.
	PropAnyOf = "http://www.w3.org/ns/activitystreams#anyOf"

	PropAttachment = "http://www.w3.org/ns/activitystreams#attachment"

	// Identifies an entity to which an object is attributed
	PropAttributedTo = "http://www.w3.org/ns/activitystreams#attributedTo"

	PropAudience = "http://www.w3.org/ns/activitystreams#audience"

	PropBcc = "http://www.w3.org/ns/activitystreams#bcc"

	PropBto = "http://www.w3.org/ns/activitystreams#bto"

	PropCc = "http://www.w3.org/ns/activitystreams#cc"

	// The content of the object.
	PropContent = "http://www.w3.org/ns/activitystreams#content"

	// Specifies the context within which an object exists or an activity was performed
	PropContext = "http://www.w3.org/ns/activitystreams#context"

	// Specifies the date and time the object was deleted
	PropDeleted = "http://www.w3.org/ns/activitystreams#deleted"

	PropDownstreamDuplicates = "http://www.w3.org/ns/activitystreams#downstreamDuplicates"

	// The duration of the object
	PropDuration = "http://www.w3.org/ns/activitystreams#duration"

	// The ending time of the object
	PropEndTime = "http://www.w3.org/ns/activitystreams#endTime"

	PropGenerator = "http://www.w3.org/ns/activitystreams#generator"

	// The display height expressed as device independent pixels
	PropHeight = "http://www.w3.org/ns/activitystreams#height"

	// The target URI of the Link
	PropHref = "http://www.w3.org/ns/activitystreams#href"

	// A hint about the language of the referenced resource
	PropHreflang = "http://www.w3.org/ns/activitystreams#hreflang"

	PropIcon = "http://www.w3.org/ns/activitystreams#icon"

	PropId = "http://www.w3.org/ns/activitystreams#id"

	PropImage = "http://www.w3.org/ns/activitystreams#image"

	PropInReplyTo = "http://www.w3.org/ns/activitystreams#inReplyTo"

	// Indentifies an object used (or to be used) to complete an activity
	PropInstrument = "http://www.w3.org/ns/activitystreams#instrument"

	PropItems = "http://www.w3.org/ns/activitystreams#items"

	// The latitude
	PropLatitude = "http://www.w3.org/ns/activitystreams#latitude"

	PropLocation = "http://www.w3.org/ns/activitystreams#location"

	// The longitude
	PropLongitude = "http://www.w3.org/ns/activitystreams#longitude"

	// The MIME Media Type
	PropMediaType = "http://www.w3.org/ns/activitystreams#mediaType"

	PropName = "http://www.w3.org/ns/activitystreams#name"

	PropObject = "http://www.w3.org/ns/activitystreams#object"

	PropObjectType = "http://www.w3.org/ns/activitystreams#objectType"

	// Describes a possible exclusive answer or option for a question.
	PropOneOf = "http://www.w3.org/ns/activitystreams#oneOf"

	// For certain activities, specifies the entity from which the action is directed.
	PropOrigin = "http://www.w3.org/ns/activitystreams#origin"

	PropPreview = "http://www.w3.org/ns/activitystreams#preview"

	// Specifies the date and time the object was published
	PropPublished = "http://www.w3.org/ns/activitystreams#published"

	// Specifies a radius around the point established by the longitude and latitude
	PropRadius = "http://www.w3.org/ns/activitystreams#radius"

	// A numeric rating (>= 0.0, <= 5.0) for the object
	PropRating = "http://www.w3.org/ns/activitystreams#rating"

	// The RFC 5988 or HTML5 Link Relation associated with the Link
	PropRel = "http://www.w3.org/ns/activitystreams#rel"

	// On a Relationship object, describes the type of relationship
	PropRelationship = "http://www.w3.org/ns/activitystreams#relationship"

	PropReplies = "http://www.w3.org/ns/activitystreams#replies"

	PropResult = "http://www.w3.org/ns/activitystreams#result"

	// In a strictly ordered logical collection, specifies the index position of the first item in the items list
	PropStartIndex = "http://www.w3.org/ns/activitystreams#startIndex"

	// The starting time of the object
	PropStartTime = "http://www.w3.org/ns/activitystreams#startTime"

	// A short summary of the object
	PropSummary = "http://www.w3.org/ns/activitystreams#summary"

	PropTag = "http://www.w3.org/ns/activitystreams#tag"

	PropTarget = "http://www.w3.org/ns/activitystreams#target"

	PropTo = "http://www.w3.org/ns/activitystreams#to"

	// The total number of items in a logical collection
	PropTotalItems = "http://www.w3.org/ns/activitystreams#totalItems"

	// Identifies the unit of measurement used by the radius, altitude and accuracy properties. The value can be expressed either as one of a set of predefined units or as a well-known common URI that identifies units.
	PropUnits = "http://www.w3.org/ns/activitystreams#units"

	// Specifies when the object was last updated
	PropUpdated = "http://www.w3.org/ns/activitystreams#updated"

	PropUpstreamDuplicates = "http://www.w3.org/ns/activitystreams#upstreamDuplicates"

	// Specifies a link to a specific representation of the Object
	PropUrl = "http://www.w3.org/ns/activitystreams#url"

	PropVerb = "http://www.w3.org/ns/activitystreams#verb"

	// Specifies the preferred display width of the content, expressed in terms of device independent pixels.
	PropWidth = "http://www.w3.org/ns/activitystreams#width"
)

// Represents a Social Graph relationship between two Individuals (indicated by the 'a' and 'b' properties)
// http://www.w3.org/ns/activitystreams#Relationship - http://www.w3.org/1999/02/22-rdf-syntax-ns#Statement

// http://www.w3.org/ns/activitystreams#attachments - http://www.w3.org/2002/07/owl#DeprecatedProperty

// Identifies the author of an object. Deprecated. Use as:attributedTo instead
// http://www.w3.org/ns/activitystreams#author - http://www.w3.org/2002/07/owl#DeprecatedProperty

// http://www.w3.org/ns/activitystreams#current - http://www.w3.org/2002/07/owl#FunctionalProperty

// On a Profile object, describes the object described by the profile
// http://www.w3.org/ns/activitystreams#describes - http://www.w3.org/2002/07/owl#FunctionalProperty

// http://www.w3.org/ns/activitystreams#first - http://www.w3.org/2002/07/owl#FunctionalProperty

// On a Tombstone object, describes the former type of the deleted object
// http://www.w3.org/ns/activitystreams#formerType - http://www.w3.org/2002/07/owl#FunctionalProperty

// http://www.w3.org/ns/activitystreams#last - http://www.w3.org/2002/07/owl#FunctionalProperty

// http://www.w3.org/ns/activitystreams#next - http://www.w3.org/2002/07/owl#FunctionalProperty

// http://www.w3.org/ns/activitystreams#partOf - http://www.w3.org/2002/07/owl#FunctionalProperty

// http://www.w3.org/ns/activitystreams#prev - http://www.w3.org/2002/07/owl#FunctionalProperty

// http://www.w3.org/ns/activitystreams#provider - http://www.w3.org/2002/07/owl#DeprecatedProperty

// On a Relationship object, identifies the subject. e.g. when saying "John is connected to Sally", 'subject' refers to 'John'
// http://www.w3.org/ns/activitystreams#subject - http://www.w3.org/2002/07/owl#FunctionalProperty

// http://www.w3.org/ns/activitystreams#tags - http://www.w3.org/2002/07/owl#DeprecatedProperty
