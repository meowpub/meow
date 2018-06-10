package entities

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/models"
)

func newStore(t *testing.T) (*gomock.Controller, *models.MockEntityStore, *Store) {
	ctrl := gomock.NewController(t)
	rawStore := models.NewMockEntityStore(ctrl)
	store := NewStore(rawStore)

	return ctrl, rawStore, store
}

func TestStore(t *testing.T) {
	t.Run("Create Entity", func(t *testing.T) {
		ctrl, raw, store := newStore(t)
		defer ctrl.Finish()

		obj, err := NewObject(store, "https://example.com/", []string{"https://www.w3.org/ns/activitystreams#Note"})
		require.NoError(t, err)
		assert.Equal(t, obj.GetID(), "https://example.com/", "ID should match")
		if assert.Equal(t, len(obj.GetType()), 1, "Should have one type") {
			assert.Equal(t, obj.GetType()[0], "https://www.w3.org/ns/activitystreams#Note", "Should be a note")
		}

		obj2, err := store.GetByID("https://example.com/")
		if assert.NoError(t, err, "GetByID") {
			assert.Equal(t, obj2, obj, "Fetching object should return same value")
		}

		raw.EXPECT().Save(gomock.Any())
		require.NoError(t, store.Save(obj))
	})
	t.Run("Fetch Entity", func(t *testing.T) {
		ctrl, raw, store := newStore(t)
		defer ctrl.Finish()

		raw.EXPECT().GetByID(gomock.Eq("https://example.net/")).Return(&models.Entity{
			ID:   353894652568535040,
			Kind: "object",
			Data: models.JSONB(`{
				"@id": "https://example.net/",
				"@type": ["https://w3c.org/ns/activitystreams#Article"],
				"https://www.w3.org/ns/activitystreams#name": [{"@value": "Catgirls: how?"}],
				"https://www.w3.org/ns/activitystreams#content": [{"@value": "my disrespectful teen son somehow got  hold of a gluten product and now he wants to become a cat girl"}]
			}`),
		}, nil)
		gluten, err := store.GetByID("https://example.net/")
		glutenObj := gluten.(*Object)
		require.NotNil(t, gluten)
		require.NoError(t, err, "Should return a value")
		require.Equal(t, "https://example.net/", gluten.GetID())
		require.Equal(t, []string{"https://w3c.org/ns/activitystreams#Article"}, gluten.GetType())
		require.Equal(t, "Catgirls: how?", glutenObj.Name.String())
		require.Equal(t, "my disrespectful teen son somehow got  hold of a gluten product and now he wants to become a cat girl", glutenObj.Content.String())

		glutenObj.Url = jsonld.ToRef("http://catgirl.how/", jsonld.Type{})

		raw.EXPECT().Save(gomock.Eq(models.Entity{
			ID:   353894652568535040,
			Kind: "object",
			Data: models.JSONB(`
{"@id"
:"https://example.net/"
,"@type"
:["https://w3c.org/ns/activitystreams#Article"]
,"https://www.w3.org/ns/activitystreams#content"
:[{"@value":"my disrespectful teen son somehow got  hold of a gluten product and now he wants to become a cat girl"}]
,"https://www.w3.org/ns/activitystreams#name"
:[{"@value":"Catgirls: how?"}]
,"https://www.w3.org/ns/activitystreams#url"
:[{"@id":"http://catgirl.how/"}]
}`[1:]),
		})).Return(nil)
		err = store.Save(glutenObj)
		require.NoError(t, err, "Saving modified entity")
	})
	t.Run("Not found", func(t *testing.T) {
		ctrl, raw, store := newStore(t)
		defer ctrl.Finish()

		raw.EXPECT().GetByID(gomock.Eq("https://example.com/")).Return(nil, errors.New("Not there"))
		obj, err := store.GetByID("https://example.com/")
		assert.Equal(t, nil, obj, "Not found should return nil object")
		assert.Error(t, err, "Should return an error")
	})
}
