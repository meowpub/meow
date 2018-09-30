package ns

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns/as"
)

func TestManifest(t *testing.T) {
	t.Run("Unknown Type", func(t *testing.T) {
		o, err := ld.NewObject([]byte(`{
			"@id": "https://example.com/@jsmith",
			"@type": ["http://schema.org/Person"],
			"http://schema.org/name": [{"@value": "John Smith"}]
		}`))
		require.NoError(t, err)
		assert.Empty(t, Manifest(o))
	})
	t.Run("One Type", func(t *testing.T) {
		o, err := ld.NewObject([]byte(`{
			"@id": "https://example.com/@jsmith",
			"@type": ["http://www.w3.org/ns/activitystreams#Person"],
			"http://www.w3.org/ns/activitystreams#name": [{"@value": "John Smith"}]
		}`))
		require.NoError(t, err)
		assert.True(t, as.IsPerson(o))
		if entities := Manifest(o); assert.Len(t, entities, 1) {
			assert.IsType(t, as.Person{}, entities[0])
		}
	})
	t.Run("Two Types", func(t *testing.T) {
		// A person who is also a place... how mysterious.
		o, err := ld.NewObject([]byte(`{
			"@id": "https://example.com/@jsmith",
			"@type": [
				"http://www.w3.org/ns/activitystreams#Person",
				"http://www.w3.org/ns/activitystreams#Place"
			],
			"http://www.w3.org/ns/activitystreams#name": [{"@value": "John Smith"}]
		}`))
		require.NoError(t, err)
		assert.True(t, as.IsPerson(o))
		assert.True(t, as.IsPlace(o))
		if entities := Manifest(o); assert.Len(t, entities, 2) {
			assert.IsType(t, as.Person{}, entities[0])
			assert.IsType(t, as.Place{}, entities[1])
		}
	})
}
