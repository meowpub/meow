// GENERATED FILE, DO NOT EDIT.
// Please refer to: tools/nsgen/templates.go
package as

import (
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/meta"
)

var (
	AS = &meta.Namespace{
		ID:    "http://www.w3.org/ns/activitystreams#",
		Short: "as",
		Props: []*meta.Prop{
			Prop_Accuracy,
			Prop_Actor,
			Prop_Altitude,
			Prop_AnyOf,
			Prop_Attachment,
			Prop_Attachments,
			Prop_AttributedTo,
			Prop_Audience,
			Prop_Author,
			Prop_Bcc,
			Prop_Bto,
			Prop_Cc,
			Prop_Content,
			Prop_Context,
			Prop_Current,
			Prop_Deleted,
			Prop_Describes,
			Prop_DownstreamDuplicates,
			Prop_Duration,
			Prop_EndTime,
			Prop_First,
			Prop_FormerType,
			Prop_Generator,
			Prop_Height,
			Prop_Href,
			Prop_Hreflang,
			Prop_Icon,
			Prop_Id,
			Prop_Image,
			Prop_InReplyTo,
			Prop_Instrument,
			Prop_Items,
			Prop_Last,
			Prop_Latitude,
			Prop_Location,
			Prop_Longitude,
			Prop_MediaType,
			Prop_Name,
			Prop_Next,
			Prop_Object,
			Prop_ObjectType,
			Prop_OneOf,
			Prop_Origin,
			Prop_PartOf,
			Prop_Prev,
			Prop_Preview,
			Prop_Provider,
			Prop_Published,
			Prop_Radius,
			Prop_Rating,
			Prop_Rel,
			Prop_Relationship,
			Prop_Replies,
			Prop_Result,
			Prop_StartIndex,
			Prop_StartTime,
			Prop_Subject,
			Prop_Summary,
			Prop_Tag,
			Prop_Tags,
			Prop_Target,
			Prop_To,
			Prop_TotalItems,
			Prop_Units,
			Prop_Updated,
			Prop_UpstreamDuplicates,
			Prop_Url,
			Prop_Verb,
			Prop_Width,
		},
		Types: map[string]*meta.Type{
			"http://www.w3.org/ns/activitystreams#Accept":                Class_Accept,
			"http://www.w3.org/ns/activitystreams#Activity":              Class_Activity,
			"http://www.w3.org/ns/activitystreams#Add":                   Class_Add,
			"http://www.w3.org/ns/activitystreams#Announce":              Class_Announce,
			"http://www.w3.org/ns/activitystreams#Application":           Class_Application,
			"http://www.w3.org/ns/activitystreams#Arrive":                Class_Arrive,
			"http://www.w3.org/ns/activitystreams#Article":               Class_Article,
			"http://www.w3.org/ns/activitystreams#Audio":                 Class_Audio,
			"http://www.w3.org/ns/activitystreams#Block":                 Class_Block,
			"http://www.w3.org/ns/activitystreams#Collection":            Class_Collection,
			"http://www.w3.org/ns/activitystreams#CollectionPage":        Class_CollectionPage,
			"http://www.w3.org/ns/activitystreams#Create":                Class_Create,
			"http://www.w3.org/ns/activitystreams#Delete":                Class_Delete,
			"http://www.w3.org/ns/activitystreams#Dislike":               Class_Dislike,
			"http://www.w3.org/ns/activitystreams#Document":              Class_Document,
			"http://www.w3.org/ns/activitystreams#Event":                 Class_Event,
			"http://www.w3.org/ns/activitystreams#Flag":                  Class_Flag,
			"http://www.w3.org/ns/activitystreams#Follow":                Class_Follow,
			"http://www.w3.org/ns/activitystreams#Group":                 Class_Group,
			"http://www.w3.org/ns/activitystreams#Ignore":                Class_Ignore,
			"http://www.w3.org/ns/activitystreams#Image":                 Class_Image,
			"http://www.w3.org/ns/activitystreams#IntransitiveActivity":  Class_IntransitiveActivity,
			"http://www.w3.org/ns/activitystreams#Invite":                Class_Invite,
			"http://www.w3.org/ns/activitystreams#Join":                  Class_Join,
			"http://www.w3.org/ns/activitystreams#Leave":                 Class_Leave,
			"http://www.w3.org/ns/activitystreams#Like":                  Class_Like,
			"http://www.w3.org/ns/activitystreams#Link":                  Class_Link,
			"http://www.w3.org/ns/activitystreams#Listen":                Class_Listen,
			"http://www.w3.org/ns/activitystreams#Mention":               Class_Mention,
			"http://www.w3.org/ns/activitystreams#Move":                  Class_Move,
			"http://www.w3.org/ns/activitystreams#Note":                  Class_Note,
			"http://www.w3.org/ns/activitystreams#Object":                Class_Object,
			"http://www.w3.org/ns/activitystreams#Offer":                 Class_Offer,
			"http://www.w3.org/ns/activitystreams#OrderedCollection":     Class_OrderedCollection,
			"http://www.w3.org/ns/activitystreams#OrderedCollectionPage": Class_OrderedCollectionPage,
			"http://www.w3.org/ns/activitystreams#OrderedItems":          Class_OrderedItems,
			"http://www.w3.org/ns/activitystreams#Organization":          Class_Organization,
			"http://www.w3.org/ns/activitystreams#Page":                  Class_Page,
			"http://www.w3.org/ns/activitystreams#Person":                Class_Person,
			"http://www.w3.org/ns/activitystreams#Place":                 Class_Place,
			"http://www.w3.org/ns/activitystreams#Profile":               Class_Profile,
			"http://www.w3.org/ns/activitystreams#Question":              Class_Question,
			"http://www.w3.org/ns/activitystreams#Read":                  Class_Read,
			"http://www.w3.org/ns/activitystreams#Reject":                Class_Reject,
			"http://www.w3.org/ns/activitystreams#Relationship":          Class_Relationship,
			"http://www.w3.org/ns/activitystreams#Remove":                Class_Remove,
			"http://www.w3.org/ns/activitystreams#Service":               Class_Service,
			"http://www.w3.org/ns/activitystreams#TentativeAccept":       Class_TentativeAccept,
			"http://www.w3.org/ns/activitystreams#TentativeReject":       Class_TentativeReject,
			"http://www.w3.org/ns/activitystreams#Tombstone":             Class_Tombstone,
			"http://www.w3.org/ns/activitystreams#Travel":                Class_Travel,
			"http://www.w3.org/ns/activitystreams#Undo":                  Class_Undo,
			"http://www.w3.org/ns/activitystreams#Update":                Class_Update,
			"http://www.w3.org/ns/activitystreams#Video":                 Class_Video,
			"http://www.w3.org/ns/activitystreams#View":                  Class_View,
		},
	}

	// Specifies the accuracy around the point established by the longitude and latitude
	Prop_Accuracy = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#accuracy",
		Short:   "accuracy",
		Comment: "Specifies the accuracy around the point established by the longitude and latitude",
	}

	// Subproperty of as:attributedTo that identifies the primary actor
	Prop_Actor = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#actor",
		Short:   "actor",
		Comment: "Subproperty of as:attributedTo that identifies the primary actor",
	}

	// The altitude of a place
	Prop_Altitude = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#altitude",
		Short:   "altitude",
		Comment: "The altitude of a place",
	}

	// Describes a possible inclusive answer or option for a question.
	Prop_AnyOf = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#anyOf",
		Short:   "anyOf",
		Comment: "Describes a possible inclusive answer or option for a question.",
	}

	Prop_Attachment = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#attachment",
		Short:   "attachment",
		Comment: "",
	}

	Prop_Attachments = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#attachments",
		Short:   "attachments",
		Comment: "",
	}

	// Identifies an entity to which an object is attributed
	Prop_AttributedTo = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#attributedTo",
		Short:   "attributedTo",
		Comment: "Identifies an entity to which an object is attributed",
	}

	Prop_Audience = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#audience",
		Short:   "audience",
		Comment: "",
	}

	// Identifies the author of an object. Deprecated. Use as:attributedTo instead
	Prop_Author = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#author",
		Short:   "author",
		Comment: "Identifies the author of an object. Deprecated. Use as:attributedTo instead",
	}

	Prop_Bcc = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#bcc",
		Short:   "bcc",
		Comment: "",
	}

	Prop_Bto = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#bto",
		Short:   "bto",
		Comment: "",
	}

	Prop_Cc = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#cc",
		Short:   "cc",
		Comment: "",
	}

	// The content of the object.
	Prop_Content = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#content",
		Short:   "content",
		Comment: "The content of the object.",
	}

	// Specifies the context within which an object exists or an activity was performed
	Prop_Context = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#context",
		Short:   "context",
		Comment: "Specifies the context within which an object exists or an activity was performed",
	}

	Prop_Current = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#current",
		Short:   "current",
		Comment: "",
	}

	// Specifies the date and time the object was deleted
	Prop_Deleted = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#deleted",
		Short:   "deleted",
		Comment: "Specifies the date and time the object was deleted",
	}

	// On a Profile object, describes the object described by the profile
	Prop_Describes = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#describes",
		Short:   "describes",
		Comment: "On a Profile object, describes the object described by the profile",
	}

	Prop_DownstreamDuplicates = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#downstreamDuplicates",
		Short:   "downstreamDuplicates",
		Comment: "",
	}

	// The duration of the object
	Prop_Duration = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#duration",
		Short:   "duration",
		Comment: "The duration of the object",
	}

	// The ending time of the object
	Prop_EndTime = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#endTime",
		Short:   "endTime",
		Comment: "The ending time of the object",
	}

	Prop_First = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#first",
		Short:   "first",
		Comment: "",
	}

	// On a Tombstone object, describes the former type of the deleted object
	Prop_FormerType = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#formerType",
		Short:   "formerType",
		Comment: "On a Tombstone object, describes the former type of the deleted object",
	}

	Prop_Generator = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#generator",
		Short:   "generator",
		Comment: "",
	}

	// The display height expressed as device independent pixels
	Prop_Height = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#height",
		Short:   "height",
		Comment: "The display height expressed as device independent pixels",
	}

	// The target URI of the Link
	Prop_Href = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#href",
		Short:   "href",
		Comment: "The target URI of the Link",
	}

	// A hint about the language of the referenced resource
	Prop_Hreflang = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#hreflang",
		Short:   "hreflang",
		Comment: "A hint about the language of the referenced resource",
	}

	Prop_Icon = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#icon",
		Short:   "icon",
		Comment: "",
	}

	Prop_Id = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#id",
		Short:   "id",
		Comment: "",
	}

	Prop_Image = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#image",
		Short:   "image",
		Comment: "",
	}

	Prop_InReplyTo = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#inReplyTo",
		Short:   "inReplyTo",
		Comment: "",
	}

	// Indentifies an object used (or to be used) to complete an activity
	Prop_Instrument = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#instrument",
		Short:   "instrument",
		Comment: "Indentifies an object used (or to be used) to complete an activity",
	}

	Prop_Items = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#items",
		Short:   "items",
		Comment: "",
	}

	Prop_Last = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#last",
		Short:   "last",
		Comment: "",
	}

	// The latitude
	Prop_Latitude = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#latitude",
		Short:   "latitude",
		Comment: "The latitude",
	}

	Prop_Location = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#location",
		Short:   "location",
		Comment: "",
	}

	// The longitude
	Prop_Longitude = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#longitude",
		Short:   "longitude",
		Comment: "The longitude",
	}

	// The MIME Media Type
	Prop_MediaType = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#mediaType",
		Short:   "mediaType",
		Comment: "The MIME Media Type",
	}

	Prop_Name = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#name",
		Short:   "name",
		Comment: "",
	}

	Prop_Next = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#next",
		Short:   "next",
		Comment: "",
	}

	Prop_Object = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#object",
		Short:   "object",
		Comment: "",
	}

	Prop_ObjectType = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#objectType",
		Short:   "objectType",
		Comment: "",
	}

	// Describes a possible exclusive answer or option for a question.
	Prop_OneOf = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#oneOf",
		Short:   "oneOf",
		Comment: "Describes a possible exclusive answer or option for a question.",
	}

	// For certain activities, specifies the entity from which the action is directed.
	Prop_Origin = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#origin",
		Short:   "origin",
		Comment: "For certain activities, specifies the entity from which the action is directed.",
	}

	Prop_PartOf = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#partOf",
		Short:   "partOf",
		Comment: "",
	}

	Prop_Prev = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#prev",
		Short:   "prev",
		Comment: "",
	}

	Prop_Preview = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#preview",
		Short:   "preview",
		Comment: "",
	}

	Prop_Provider = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#provider",
		Short:   "provider",
		Comment: "",
	}

	// Specifies the date and time the object was published
	Prop_Published = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#published",
		Short:   "published",
		Comment: "Specifies the date and time the object was published",
	}

	// Specifies a radius around the point established by the longitude and latitude
	Prop_Radius = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#radius",
		Short:   "radius",
		Comment: "Specifies a radius around the point established by the longitude and latitude",
	}

	// A numeric rating (>= 0.0, <= 5.0) for the object
	Prop_Rating = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#rating",
		Short:   "rating",
		Comment: "A numeric rating (>= 0.0, <= 5.0) for the object",
	}

	// The RFC 5988 or HTML5 Link Relation associated with the Link
	Prop_Rel = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#rel",
		Short:   "rel",
		Comment: "The RFC 5988 or HTML5 Link Relation associated with the Link",
	}

	// On a Relationship object, describes the type of relationship
	Prop_Relationship = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#relationship",
		Short:   "relationship",
		Comment: "On a Relationship object, describes the type of relationship",
	}

	Prop_Replies = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#replies",
		Short:   "replies",
		Comment: "",
	}

	Prop_Result = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#result",
		Short:   "result",
		Comment: "",
	}

	// In a strictly ordered logical collection, specifies the index position of the first item in the items list
	Prop_StartIndex = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#startIndex",
		Short:   "startIndex",
		Comment: "In a strictly ordered logical collection, specifies the index position of the first item in the items list",
	}

	// The starting time of the object
	Prop_StartTime = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#startTime",
		Short:   "startTime",
		Comment: "The starting time of the object",
	}

	// On a Relationship object, identifies the subject. e.g. when saying "John is connected to Sally", 'subject' refers to 'John'
	Prop_Subject = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#subject",
		Short:   "subject",
		Comment: "On a Relationship object, identifies the subject. e.g. when saying \"John is connected to Sally\", 'subject' refers to 'John'",
	}

	// A short summary of the object
	Prop_Summary = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#summary",
		Short:   "summary",
		Comment: "A short summary of the object",
	}

	Prop_Tag = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#tag",
		Short:   "tag",
		Comment: "",
	}

	Prop_Tags = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#tags",
		Short:   "tags",
		Comment: "",
	}

	Prop_Target = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#target",
		Short:   "target",
		Comment: "",
	}

	Prop_To = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#to",
		Short:   "to",
		Comment: "",
	}

	// The total number of items in a logical collection
	Prop_TotalItems = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#totalItems",
		Short:   "totalItems",
		Comment: "The total number of items in a logical collection",
	}

	// Identifies the unit of measurement used by the radius, altitude and accuracy properties. The value can be expressed either as one of a set of predefined units or as a well-known common URI that identifies units.
	Prop_Units = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#units",
		Short:   "units",
		Comment: "Identifies the unit of measurement used by the radius, altitude and accuracy properties. The value can be expressed either as one of a set of predefined units or as a well-known common URI that identifies units.",
	}

	// Specifies when the object was last updated
	Prop_Updated = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#updated",
		Short:   "updated",
		Comment: "Specifies when the object was last updated",
	}

	Prop_UpstreamDuplicates = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#upstreamDuplicates",
		Short:   "upstreamDuplicates",
		Comment: "",
	}

	// Specifies a link to a specific representation of the Object
	Prop_Url = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#url",
		Short:   "url",
		Comment: "Specifies a link to a specific representation of the Object",
	}

	Prop_Verb = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#verb",
		Short:   "verb",
		Comment: "",
	}

	// Specifies the preferred display width of the content, expressed in terms of device independent pixels.
	Prop_Width = &meta.Prop{
		ID:      "http://www.w3.org/ns/activitystreams#width",
		Short:   "width",
		Comment: "Specifies the preferred display width of the content, expressed in terms of device independent pixels.",
	}

	// Actor accepts the Object
	Class_Accept = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Accept",
		Short: "Accept",
		Cast:  func(e ld.Entity) ld.Entity { return AsAccept(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// An Object representing some form of Action that has been taken
	Class_Activity = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Activity",
		Short: "Activity",
		Cast:  func(e ld.Entity) ld.Entity { return AsActivity(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{
			Prop_Actor,
			Prop_Instrument,
			Prop_Origin,
			Prop_Result,
			Prop_Target,
			Prop_Verb,
		},
	}

	// To Add an Object or Link to Something
	Class_Add = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Add",
		Short: "Add",
		Cast:  func(e ld.Entity) ld.Entity { return AsAdd(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// Actor announces the object to the target
	Class_Announce = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Announce",
		Short: "Announce",
		Cast:  func(e ld.Entity) ld.Entity { return AsAnnounce(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// Represents a software application of any sort
	Class_Application = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Application",
		Short: "Application",
		Cast:  func(e ld.Entity) ld.Entity { return AsApplication(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{},
	}

	// To Arrive Somewhere (can be used, for instance, to indicate that a particular entity is currently located somewhere, e.g. a "check-in")
	Class_Arrive = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Arrive",
		Short: "Arrive",
		Cast:  func(e ld.Entity) ld.Entity { return AsArrive(e) },
		SubClassOf: []*meta.Type{
			Class_IntransitiveActivity,
		},
		Props: []*meta.Prop{},
	}

	// A written work. Typically several paragraphs long. For example, a blog post or a news article.
	Class_Article = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Article",
		Short: "Article",
		Cast:  func(e ld.Entity) ld.Entity { return AsArticle(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{},
	}

	// An audio file
	Class_Audio = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Audio",
		Short: "Audio",
		Cast:  func(e ld.Entity) ld.Entity { return AsAudio(e) },
		SubClassOf: []*meta.Type{
			Class_Document,
		},
		Props: []*meta.Prop{},
	}

	Class_Block = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Block",
		Short: "Block",
		Cast:  func(e ld.Entity) ld.Entity { return AsBlock(e) },
		SubClassOf: []*meta.Type{
			Class_Ignore,
		},
		Props: []*meta.Prop{},
	}

	// An ordered or unordered collection of Objects or Links
	Class_Collection = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Collection",
		Short: "Collection",
		Cast:  func(e ld.Entity) ld.Entity { return AsCollection(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{
			Prop_Current,
			Prop_First,
			Prop_Items,
			Prop_Last,
			Prop_TotalItems,
		},
	}

	// A subset of items from a Collection
	Class_CollectionPage = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#CollectionPage",
		Short: "CollectionPage",
		Cast:  func(e ld.Entity) ld.Entity { return AsCollectionPage(e) },
		SubClassOf: []*meta.Type{
			Class_Collection,
		},
		Props: []*meta.Prop{
			Prop_Next,
			Prop_PartOf,
			Prop_Prev,
		},
	}

	// To Create Something
	Class_Create = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Create",
		Short: "Create",
		Cast:  func(e ld.Entity) ld.Entity { return AsCreate(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// To Delete Something
	Class_Delete = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Delete",
		Short: "Delete",
		Cast:  func(e ld.Entity) ld.Entity { return AsDelete(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// The actor dislikes the object
	Class_Dislike = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Dislike",
		Short: "Dislike",
		Cast:  func(e ld.Entity) ld.Entity { return AsDislike(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// Represents a digital document/file of any sort
	Class_Document = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Document",
		Short: "Document",
		Cast:  func(e ld.Entity) ld.Entity { return AsDocument(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{},
	}

	// An Event of any kind
	Class_Event = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Event",
		Short: "Event",
		Cast:  func(e ld.Entity) ld.Entity { return AsEvent(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{},
	}

	// To flag something (e.g. flag as inappropriate, flag as spam, etc)
	Class_Flag = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Flag",
		Short: "Flag",
		Cast:  func(e ld.Entity) ld.Entity { return AsFlag(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// To Express Interest in Something
	Class_Follow = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Follow",
		Short: "Follow",
		Cast:  func(e ld.Entity) ld.Entity { return AsFollow(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// A Group of any kind.
	Class_Group = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Group",
		Short: "Group",
		Cast:  func(e ld.Entity) ld.Entity { return AsGroup(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{},
	}

	// Actor is ignoring the Object
	Class_Ignore = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Ignore",
		Short: "Ignore",
		Cast:  func(e ld.Entity) ld.Entity { return AsIgnore(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// An Image file
	Class_Image = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Image",
		Short: "Image",
		Cast:  func(e ld.Entity) ld.Entity { return AsImage(e) },
		SubClassOf: []*meta.Type{
			Class_Document,
		},
		Props: []*meta.Prop{},
	}

	// An Activity that has no direct object
	Class_IntransitiveActivity = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#IntransitiveActivity",
		Short: "IntransitiveActivity",
		Cast:  func(e ld.Entity) ld.Entity { return AsIntransitiveActivity(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// To invite someone or something to something
	Class_Invite = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Invite",
		Short: "Invite",
		Cast:  func(e ld.Entity) ld.Entity { return AsInvite(e) },
		SubClassOf: []*meta.Type{
			Class_Offer,
		},
		Props: []*meta.Prop{},
	}

	// To Join Something
	Class_Join = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Join",
		Short: "Join",
		Cast:  func(e ld.Entity) ld.Entity { return AsJoin(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// To Leave Something
	Class_Leave = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Leave",
		Short: "Leave",
		Cast:  func(e ld.Entity) ld.Entity { return AsLeave(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// To Like Something
	Class_Like = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Like",
		Short: "Like",
		Cast:  func(e ld.Entity) ld.Entity { return AsLike(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// Represents a qualified reference to another resource. Patterned after the RFC5988 Web Linking Model
	Class_Link = &meta.Type{
		ID:         "http://www.w3.org/ns/activitystreams#Link",
		Short:      "Link",
		Cast:       func(e ld.Entity) ld.Entity { return AsLink(e) },
		SubClassOf: []*meta.Type{},
		Props: []*meta.Prop{
			Prop_AttributedTo,
			Prop_Height,
			Prop_Href,
			Prop_Hreflang,
			Prop_Id,
			Prop_MediaType,
			Prop_Name,
			Prop_Object,
			Prop_Preview,
			Prop_Rel,
			Prop_Width,
		},
	}

	// The actor listened to the object
	Class_Listen = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Listen",
		Short: "Listen",
		Cast:  func(e ld.Entity) ld.Entity { return AsListen(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// A specialized Link that represents an @mention
	Class_Mention = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Mention",
		Short: "Mention",
		Cast:  func(e ld.Entity) ld.Entity { return AsMention(e) },
		SubClassOf: []*meta.Type{
			Class_Link,
		},
		Props: []*meta.Prop{},
	}

	// The actor is moving the object. The target specifies where the object is moving to. The origin specifies where the object is moving from.
	Class_Move = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Move",
		Short: "Move",
		Cast:  func(e ld.Entity) ld.Entity { return AsMove(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// A Short note, typically less than a single paragraph. A "tweet" is an example, or a "status update"
	Class_Note = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Note",
		Short: "Note",
		Cast:  func(e ld.Entity) ld.Entity { return AsNote(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{},
	}

	Class_Object = &meta.Type{
		ID:         "http://www.w3.org/ns/activitystreams#Object",
		Short:      "Object",
		Cast:       func(e ld.Entity) ld.Entity { return AsObject(e) },
		SubClassOf: []*meta.Type{},
		Props: []*meta.Prop{
			Prop_Attachment,
			Prop_Attachments,
			Prop_AttributedTo,
			Prop_Audience,
			Prop_Author,
			Prop_Bcc,
			Prop_Bto,
			Prop_Cc,
			Prop_Content,
			Prop_Context,
			Prop_DownstreamDuplicates,
			Prop_Duration,
			Prop_EndTime,
			Prop_Generator,
			Prop_Icon,
			Prop_Id,
			Prop_Image,
			Prop_InReplyTo,
			Prop_Location,
			Prop_MediaType,
			Prop_Name,
			Prop_Object,
			Prop_ObjectType,
			Prop_Preview,
			Prop_Provider,
			Prop_Published,
			Prop_Rating,
			Prop_Replies,
			Prop_StartTime,
			Prop_Summary,
			Prop_Tag,
			Prop_Tags,
			Prop_To,
			Prop_Updated,
			Prop_UpstreamDuplicates,
			Prop_Url,
		},
	}

	// To Offer something to someone or something
	Class_Offer = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Offer",
		Short: "Offer",
		Cast:  func(e ld.Entity) ld.Entity { return AsOffer(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// A variation of Collection in which items are strictly ordered
	Class_OrderedCollection = &meta.Type{
		ID:         "http://www.w3.org/ns/activitystreams#OrderedCollection",
		Short:      "OrderedCollection",
		Cast:       func(e ld.Entity) ld.Entity { return AsOrderedCollection(e) },
		SubClassOf: []*meta.Type{},
		Props: []*meta.Prop{
			Prop_AttributedTo,
			Prop_Id,
			Prop_MediaType,
			Prop_Name,
			Prop_Object,
			Prop_Preview,
		},
	}

	// An ordered subset of items from an OrderedCollection
	Class_OrderedCollectionPage = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#OrderedCollectionPage",
		Short: "OrderedCollectionPage",
		Cast:  func(e ld.Entity) ld.Entity { return AsOrderedCollectionPage(e) },
		SubClassOf: []*meta.Type{
			Class_CollectionPage,
			Class_OrderedCollection,
		},
		Props: []*meta.Prop{
			Prop_StartIndex,
		},
	}

	// A rdf:List variant for Objects and Links
	Class_OrderedItems = &meta.Type{
		ID:         "http://www.w3.org/ns/activitystreams#OrderedItems",
		Short:      "OrderedItems",
		Cast:       func(e ld.Entity) ld.Entity { return AsOrderedItems(e) },
		SubClassOf: []*meta.Type{},
		Props: []*meta.Prop{
			Prop_AttributedTo,
			Prop_Id,
			Prop_MediaType,
			Prop_Name,
			Prop_Object,
			Prop_Preview,
		},
	}

	// An Organization
	Class_Organization = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Organization",
		Short: "Organization",
		Cast:  func(e ld.Entity) ld.Entity { return AsOrganization(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{},
	}

	// A Web Page
	Class_Page = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Page",
		Short: "Page",
		Cast:  func(e ld.Entity) ld.Entity { return AsPage(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{},
	}

	// A Person
	Class_Person = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Person",
		Short: "Person",
		Cast:  func(e ld.Entity) ld.Entity { return AsPerson(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{},
	}

	// A physical or logical location
	Class_Place = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Place",
		Short: "Place",
		Cast:  func(e ld.Entity) ld.Entity { return AsPlace(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{
			Prop_Accuracy,
			Prop_Altitude,
			Prop_Latitude,
			Prop_Longitude,
			Prop_Radius,
			Prop_Units,
		},
	}

	// A Profile Document
	Class_Profile = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Profile",
		Short: "Profile",
		Cast:  func(e ld.Entity) ld.Entity { return AsProfile(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{
			Prop_Describes,
		},
	}

	// A question of any sort.
	Class_Question = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Question",
		Short: "Question",
		Cast:  func(e ld.Entity) ld.Entity { return AsQuestion(e) },
		SubClassOf: []*meta.Type{
			Class_IntransitiveActivity,
		},
		Props: []*meta.Prop{
			Prop_AnyOf,
			Prop_OneOf,
		},
	}

	// The actor read the object
	Class_Read = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Read",
		Short: "Read",
		Cast:  func(e ld.Entity) ld.Entity { return AsRead(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// Actor rejects the Object
	Class_Reject = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Reject",
		Short: "Reject",
		Cast:  func(e ld.Entity) ld.Entity { return AsReject(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// Represents a Social Graph relationship between two Individuals (indicated by the 'a' and 'b' properties)
	Class_Relationship = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Relationship",
		Short: "Relationship",
		Cast:  func(e ld.Entity) ld.Entity { return AsRelationship(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{
			Prop_Relationship,
			Prop_Subject,
		},
	}

	// To Remove Something
	Class_Remove = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Remove",
		Short: "Remove",
		Cast:  func(e ld.Entity) ld.Entity { return AsRemove(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// A service provided by some entity
	Class_Service = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Service",
		Short: "Service",
		Cast:  func(e ld.Entity) ld.Entity { return AsService(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{},
	}

	// Actor tentatively accepts the Object
	Class_TentativeAccept = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#TentativeAccept",
		Short: "TentativeAccept",
		Cast:  func(e ld.Entity) ld.Entity { return AsTentativeAccept(e) },
		SubClassOf: []*meta.Type{
			Class_Accept,
		},
		Props: []*meta.Prop{},
	}

	// Actor tentatively rejects the object
	Class_TentativeReject = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#TentativeReject",
		Short: "TentativeReject",
		Cast:  func(e ld.Entity) ld.Entity { return AsTentativeReject(e) },
		SubClassOf: []*meta.Type{
			Class_Reject,
		},
		Props: []*meta.Prop{},
	}

	// A placeholder for a deleted object
	Class_Tombstone = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Tombstone",
		Short: "Tombstone",
		Cast:  func(e ld.Entity) ld.Entity { return AsTombstone(e) },
		SubClassOf: []*meta.Type{
			Class_Object,
		},
		Props: []*meta.Prop{
			Prop_Deleted,
			Prop_FormerType,
		},
	}

	// The actor is traveling to the target. The origin specifies where the actor is traveling from.
	Class_Travel = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Travel",
		Short: "Travel",
		Cast:  func(e ld.Entity) ld.Entity { return AsTravel(e) },
		SubClassOf: []*meta.Type{
			Class_IntransitiveActivity,
		},
		Props: []*meta.Prop{},
	}

	// To Undo Something. This would typically be used to indicate that a previous Activity has been undone.
	Class_Undo = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Undo",
		Short: "Undo",
		Cast:  func(e ld.Entity) ld.Entity { return AsUndo(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// To Update/Modify Something
	Class_Update = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Update",
		Short: "Update",
		Cast:  func(e ld.Entity) ld.Entity { return AsUpdate(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}

	// A Video document of any kind.
	Class_Video = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#Video",
		Short: "Video",
		Cast:  func(e ld.Entity) ld.Entity { return AsVideo(e) },
		SubClassOf: []*meta.Type{
			Class_Document,
		},
		Props: []*meta.Prop{},
	}

	// The actor viewed the object
	Class_View = &meta.Type{
		ID:    "http://www.w3.org/ns/activitystreams#View",
		Short: "View",
		Cast:  func(e ld.Entity) ld.Entity { return AsView(e) },
		SubClassOf: []*meta.Type{
			Class_Activity,
		},
		Props: []*meta.Prop{},
	}
)
