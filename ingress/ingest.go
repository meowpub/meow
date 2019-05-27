package ingress

import (
	"context"

	"github.com/meowpub/meow/ingress/rules"
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/as"
	"github.com/meowpub/meow/models"
)

// Ingests an Entity posted to an Inbox (S2S) or Outbox (C2S), as directed by `src`.
// The caller is responsible for validating the actor's identity by whatever means.
func Ingest(ctx context.Context, stores models.Stores, src ld.Source, actor, stream *models.Entity, e ld.Entity) error {
	activity := as.AsActivity(e)

	// Rules may reject or modify submitted activities, and broadly have two roles:
	// - S2S: Filtering out unwanted content, eg. federation black/whitelists.
	// - C2S: Validation, generating IDs, wrapping plain objects in Create objects.
	if err := rules.Run(ctx, rules.Rules, src, actor, stream, &activity); err != nil {
		return err
	}

	// TODO: Rule-like system for "side effects", which is what the spec calls things like
	// a Create activity creating the contained object, or a Delete activity deleting it.

	entity := models.NewEntityFrom(models.ObjectEntity, activity)
	return stores.Entities().Save(entity)
}
