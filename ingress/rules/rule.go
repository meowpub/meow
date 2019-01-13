package rules

import (
	"context"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/as"
	"github.com/meowpub/meow/models"
)

var Rules = []Rule{
	// Client to Server rules are ran first, and only for C2S sources.
	C2SAutoCreateActivity,         // Handle as:Objects posted w/o an as:Create activity.
	C2SRequiredActivityProperties, // Check required properties.
	C2SGenerateIDs,                // Generate activity/object IDs, discarding old ones.

	RequireActivity, // Require input to be a proper Activity.
}

// A rule can process or reject an activity before it's ingested.
type Rule func(ctx context.Context, src ld.Source, actor, stream *models.Entity, activity *as.Activity) error

// Runs a series of rules on the input.
func Run(ctx context.Context, rules []Rule, src ld.Source, actor, stream *models.Entity, activity *as.Activity) error {
	for _, rule := range rules {
		if err := rule(ctx, src, actor, stream, activity); err != nil {
			return err
		}
	}
	return nil
}
