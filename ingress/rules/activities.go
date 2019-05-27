package rules

import (
	"context"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/as"
	"github.com/meowpub/meow/models"
)

func RequireActivity(ctx context.Context, src ld.Source, actor, stream *models.Entity, activity *as.Activity) (err error) {
	return ValidateQuacksLike(activity, as.Class_Activity)
}
