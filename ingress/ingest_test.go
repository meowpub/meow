package ingress

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/as"
	"github.com/meowpub/meow/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIngest(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		actor := models.NewEntityFrom(models.ObjectEntity,
			as.NewPerson("https://remote.com/~jdoe"))
		stream := models.NewEntityFrom(models.ObjectEntity,
			as.NewPerson("https://example.com/~jsmith/inbox"))

		note := as.NewNote("https://example.com/lorem-ipsum")
		note.SetContent(ld.Value("lorem ipsum dolor sit amet"))
		create := as.NewCreate("https://example.com/lorem-ipsum/outbox/create")
		create.SetObject(note)

		for _, src := range []ld.Source{ld.ClientToServer, ld.ServerToServer} {
			t.Run(src.String(), func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				stores := models.NewMockStores(ctrl)
				stores.EntityStore = models.NewMockEntityStore(ctrl)
				stores.EntityStore.EXPECT().Save(gomock.Any()).Do(func(e *models.Entity) {
					expected := models.NewEntityFrom(models.ObjectEntity, create)
					expected.ID = e.ID
					assert.Equal(t, expected, e)
				})

				require.NoError(t, Ingest(context.Background(), stores, src, actor, stream, create))
			})
		}
	})

	t.Run("Object", func(t *testing.T) {
		actor := models.NewEntityFrom(models.ObjectEntity,
			as.NewPerson("https://example.com/~jsmith"))
		stream := models.NewEntityFrom(models.ObjectEntity,
			as.NewPerson("https://example.com/~jsmith/inbox"))

		note := as.NewNote("https://example.com/lorem-ipsum")
		note.SetContent(ld.Value("lorem ipsum dolor sit amet"))

		t.Run("ServerToServer", func(t *testing.T) {
			require.EqualError(t, Ingest(context.Background(), nil, ld.ServerToServer, actor, stream, note), "@type=[[http://www.w3.org/ns/activitystreams#Note]] does not include required type (or a known subtype thereof): http://www.w3.org/ns/activitystreams#Activity")
		})

		t.Run("ClientToServer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stores := models.NewMockStores(ctrl)
			stores.EntityStore = models.NewMockEntityStore(ctrl)
			stores.EntityStore.EXPECT().Save(gomock.Any()).Do(func(e *models.Entity) {
				ecreate := as.AsCreate(e.Obj)
				assert.Equal(t, []string{as.Class_Create.ID}, ecreate.Type())
				assert.Equal(t, note.Obj().V, ecreate.Object.Object())
			})

			require.NoError(t, Ingest(context.Background(), stores, ld.ClientToServer, actor, stream, note))
		})
	})
}
