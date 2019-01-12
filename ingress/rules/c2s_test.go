package rules

import (
	"context"
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/as"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
)

// may woz ere 2k19
func TestC2SAutoCreateActivity(t *testing.T) {
	actor := models.NewEntityFrom(models.ObjectEntity,
		as.NewPerson("https://example.com/~jsmith"))
	outbox := models.NewEntityFrom(models.ObjectEntity,
		as.NewOrderedCollection("https://example.com/~jsmith/outbox"))

	note := as.NewNote("https://example.com/~jsmith/123456789")
	note.SetContent(ld.Value("Lorem ipsum dolor sit amet"))
	note.SetTo(ld.ID("https://example.com/~to"))
	note.SetBto(ld.ID("https://example.com/~bto"))
	note.SetCc(ld.ID("https://example.com/~cc"))
	note.SetBcc(ld.ID("https://example.com/~bcc"))
	note.SetAudience(ld.ID("https://example.com/~audience"))

	create := as.NewCreate("")
	create.SetObject(note)
	create.SetTo(ld.ID("https://example.com/~to"))
	create.SetBto(ld.ID("https://example.com/~bto"))
	create.SetCc(ld.ID("https://example.com/~cc"))
	create.SetBcc(ld.ID("https://example.com/~bcc"))
	create.SetAudience(ld.ID("https://example.com/~audience"))

	t.Run("Activity", func(t *testing.T) {
		activity := as.AsActivity(create)
		assert.NoError(t, C2SAutoCreateActivity(context.Background(), ld.ClientToServer, actor, outbox, &activity))
		assert.Equal(t, create.Obj().V, activity.Obj().V)
	})

	t.Run("Object", func(t *testing.T) {
		activity := as.AsActivity(note)
		assert.NoError(t, C2SAutoCreateActivity(context.Background(), ld.ClientToServer, actor, outbox, &activity))
		assert.Equal(t, create.Obj().V, activity.Obj().V)
	})
}

func TestC2SRequiredActivityProperties(t *testing.T) {
	ctx := context.Background()

	actor := models.NewEntityFrom(models.ObjectEntity,
		as.NewPerson("https://example.com/~jsmith"))
	outbox := models.NewEntityFrom(models.ObjectEntity,
		as.NewOrderedCollection("https://example.com/~jsmith/outbox"))

	t.Run("S2S", func(t *testing.T) {
		activity := as.AsActivity(as.NewCreate(""))
		assert.NoError(t, C2SRequiredActivityProperties(ctx, ld.ServerToServer, actor, outbox, &activity))
	})

	t.Run("Create.object", func(t *testing.T) {
		activity := as.AsActivity(as.NewCreate(""))
		assert.EqualError(t, C2SRequiredActivityProperties(ctx, ld.ClientToServer, actor, outbox, &activity), "activity of type [http://www.w3.org/ns/activitystreams#Create] is missing required property: http://www.w3.org/ns/activitystreams#object")

		activity.SetObject(as.NewNote(""))
		assert.NoError(t, C2SRequiredActivityProperties(ctx, ld.ClientToServer, actor, outbox, &activity))
	})

	t.Run("Add.object/Add.target", func(t *testing.T) {
		activity := as.AsActivity(as.NewAdd(""))
		assert.EqualError(t, C2SRequiredActivityProperties(ctx, ld.ClientToServer, actor, outbox, &activity), "activity of type [http://www.w3.org/ns/activitystreams#Add] is missing required property: http://www.w3.org/ns/activitystreams#object; activity of type [http://www.w3.org/ns/activitystreams#Add] is missing required property: http://www.w3.org/ns/activitystreams#target")

		activity.SetObject(as.NewNote(""))
		assert.EqualError(t, C2SRequiredActivityProperties(ctx, ld.ClientToServer, actor, outbox, &activity), "activity of type [http://www.w3.org/ns/activitystreams#Add] is missing required property: http://www.w3.org/ns/activitystreams#target")

		activity.SetTarget(as.NewCollection(""))
		assert.NoError(t, C2SRequiredActivityProperties(ctx, ld.ClientToServer, actor, outbox, &activity))
	})
}

func TestC2SGenerateIDs(t *testing.T) {
	oldFlakeGen := lib.DefaultSnowflakeGenerator
	defer func() { lib.DefaultSnowflakeGenerator = oldFlakeGen }()

	ctx := context.Background()

	actor := models.NewEntityFrom(models.ObjectEntity,
		as.NewPerson("https://example.com/~jsmith"))
	outbox := models.NewEntityFrom(models.ObjectEntity,
		as.NewOrderedCollection("https://example.com/~jsmith/outbox"))

	t.Run("Create", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		flakeGen := lib.NewMockSnowflakeGenerator(ctrl)
		lib.DefaultSnowflakeGenerator = flakeGen

		gomock.InOrder(
			flakeGen.EXPECT().Generate().Return(snowflake.ID(1)).Times(1),
			flakeGen.EXPECT().Generate().Return(snowflake.ID(2)).Times(1),
		)

		note := as.NewNote("")
		note.SetContent(ld.Value("Lorem ipsum dolor sit amet"))
		create := as.NewCreate("")
		create.SetObject(note)

		activity := as.AsActivity(create)
		assert.NoError(t, C2SGenerateIDs(ctx, ld.ClientToServer, actor, outbox, &activity))
		assert.Equal(t, "https://example.com/~jsmith/outbox/1", create.ID())
		assert.Equal(t, "https://example.com/~jsmith/2", note.ID())
	})

	t.Run("Name", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		flakeGen := lib.NewMockSnowflakeGenerator(ctrl)
		lib.DefaultSnowflakeGenerator = flakeGen

		gomock.InOrder(
			flakeGen.EXPECT().Generate().Return(snowflake.ID(1)).Times(1),
		)

		note := as.NewNote("")
		note.SetName(ld.Value("i'm gay"))
		note.SetContent(ld.Value("Lorem ipsum dolor sit amet"))
		create := as.NewCreate("")
		create.SetObject(note)

		activity := as.AsActivity(create)
		assert.NoError(t, C2SGenerateIDs(ctx, ld.ClientToServer, actor, outbox, &activity))
		assert.Equal(t, "https://example.com/~jsmith/outbox/1", create.ID())
		assert.Equal(t, "https://example.com/~jsmith/im-gay", note.ID())
	})

	t.Run("Long Name", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		flakeGen := lib.NewMockSnowflakeGenerator(ctrl)
		lib.DefaultSnowflakeGenerator = flakeGen

		gomock.InOrder(
			flakeGen.EXPECT().Generate().Return(snowflake.ID(1)).Times(1),
		)

		note := as.NewNote("")
		note.SetName(ld.Value("According to all known laws of aviation, there's no way a bee should be able to fly"))
		note.SetContent(ld.Value("Lorem ipsum dolor sit amet"))
		create := as.NewCreate("")
		create.SetObject(note)

		activity := as.AsActivity(create)
		assert.NoError(t, C2SGenerateIDs(ctx, ld.ClientToServer, actor, outbox, &activity))
		assert.Equal(t, "https://example.com/~jsmith/outbox/1", create.ID())
		assert.Equal(t, "https://example.com/~jsmith/according-to-all-known-laws-of-aviation-theres-no", note.ID())
	})
}
