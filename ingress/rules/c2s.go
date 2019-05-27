package rules

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strings"

	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns"
	"github.com/meowpub/meow/ld/ns/as"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
)

// This rule ensures that the processed activity is indeed an activity, and wraps it as
// the as:Object in an as:Create event if it's not.
//
// NOTE: This rule is ran before C2SGenerateIDs, do not depend on things having IDs.
//
// ActivityPub §6: Client To Server Interactions
//   The body of the POST request MUST contain a single Activity (which MAY contain embedded
//   objects), or a single non-Activity object which will be wrapped in a Create activity
//   by the server.
//
// ActivityPub §6.2.1: Object creation without a Create Activity
//   For client to server posting, it is possible to submit an object for creation without a
//   surrounding activity. The server MUST accept a valid [ActivityStreams] object that isn't
//   a subtype of Activity in the POST request to the outbox. The server then MUST attach this
//   object as the object of a Create Activity.
//   ...
//    Any to, bto, cc, bcc, and audience properties specified on the object MUST be copied over
//    to the new Create activity by the server.
func C2SAutoCreateActivity(ctx context.Context, src ld.Source, actor, stream *models.Entity, activity *as.Activity) error {
	if src != ld.ClientToServer {
		return nil
	}

	// Just ignore if this is already an Activity.
	if ns.QuacksLike(activity, as.Class_Activity) {
		return nil
	}
	obj := as.AsObject(activity)

	// Create a new as:Create activity.
	create := as.NewCreate("")
	create.SetObject(obj)

	// Copy addressing properties off the object.
	Copy(obj, create, as.Prop_To)
	Copy(obj, create, as.Prop_Bto)
	Copy(obj, create, as.Prop_Cc)
	Copy(obj, create, as.Prop_Bcc)
	Copy(obj, create, as.Prop_Audience)

	*activity = as.AsActivity(create)
	return nil
}

// Ensure that activities have required attributes for their type as dictated by the spec.
//
// TODO: Should this be enforced for federated activities as well as client-submitted ones?
//
// NOTE: This rule is ran before C2SGenerateIDs, do not depend on things having IDs.
//
// ActivityPub §6.1 Client Addressing:
//   Clients submitting the following activities to an outbox MUST provide the object property
//   in the activity: Create, Update, Delete, Follow, Add, Remove, Like, Block, Undo.
//   Additionally, clients submitting the following activities to an outbox MUST also provide
//   the target property: Add, Remove.
func C2SRequiredActivityProperties(ctx context.Context, src ld.Source, actor, stream *models.Entity, activity *as.Activity) (err error) {
	// Ignore if this isn't a client-to-server request.
	if src != ld.ClientToServer {
		return nil
	}

	// Activity types that must have an object property
	if ns.QuacksLikeAny(activity,
		as.Class_Create, as.Class_Update, as.Class_Delete,
		as.Class_Follow, as.Class_Add, as.Class_Remove,
		as.Class_Like, as.Class_Block, as.Class_Undo,
	) && activity.Object.Object() == nil {
		err = multierr.Append(err, lib.Errorf(http.StatusBadRequest,
			"activity of type %v is missing required property: %s",
			activity.Type(), as.Prop_Object.ID))
	}

	// Activity types that must have a target property
	if ns.QuacksLikeAny(activity, as.Class_Add, as.Class_Remove) && activity.Target() == nil {
		err = multierr.Append(err, lib.Errorf(http.StatusBadRequest,
			"activity of type %v is missing required property: %s",
			activity.Type(), as.Prop_Target.ID))
	}

	return err
}

// This rule generates local IDs for client-to-server submissions.
//
// ActivityPub §6: Client To Server Interactions
//   If an Activity is submitted with a value in the id property, servers MUST ignore this
//   and generate a new id for the Activity.
//
// ActivityPub §6.2.1: Object creation without a Create activity
//   For non-transient objects, the server MUST attach an id to both the wrapping Create and
//   its wrapped Object.
//
// ^ I'm choosing to interpret this as having been placed in the wrong section of the spec
//   (§6.2.1 rather than §6.2: Create Activity), because it makes no sense for it only to
//   apply to objects created without Create activities.
func C2SGenerateIDs(ctx context.Context, src ld.Source, actor, stream *models.Entity, activity *as.Activity) error {
	if src != ld.ClientToServer {
		return nil
	}

	// Activities go under the outbox, eg. https://example.com/@jsmith/outbox/123456789.
	streamURL, err := url.Parse(stream.Obj.ID())
	if err != nil {
		return lib.Wrap(err, http.StatusInternalServerError, "stream ID is invalid")
	}
	activityURL := streamURL.ResolveReference(&url.URL{
		Path: path.Join(streamURL.Path, lib.GenSnowflake().String()),
	})
	activity.Obj().SetID(activityURL.String())

	// If this is a Create activity, the object gets an ID as well.
	if ns.QuacksLike(activity, as.Class_Create) {
		create := as.AsCreate(activity)
		object := as.AsObject(ld.ToObject(create.Object.Object()))

		// Allow an as:name property to override the last part of the URL.
		// This is gonna need some handling for collissions at some point...
		objectID := ""
		if name := ld.ToString(object.Name()); name != "" {
			slug, err := Slugify(name)
			if err != nil {
				lib.GetLogger(ctx).Warn("Couldn't slugify object name",
					zap.String("objectID", objectID),
					zap.String("name", name),
					zap.String("slug", slug),
					zap.Error(err),
				)
			} else {
				// Take the first (at most) 50 characters, make sure there isn't a trailing -.
				if len(slug) > 50 {
					slug = strings.Trim(slug[:50], "-")
				}
				objectID += slug
			}
		} else {
			objectID = lib.GenSnowflake().String()
		}

		// Objects are placed under the actor, eg. https://example.com/@jsmith/123456789.
		actorURL, err := url.Parse(actor.Obj.ID())
		if err != nil {
			return lib.Wrap(err, http.StatusInternalServerError, "actor ID is invalid")
		}
		objectURL := actorURL.ResolveReference(&url.URL{
			Path: path.Join(actorURL.Path, objectID),
		})
		object.Obj().SetID(objectURL.String())
	}

	return nil
}
