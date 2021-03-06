package ld

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestObjectSet(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		obj := NewObject("https://example.com")
		obj.Set("key", "value")
		assert.Equal(t, "value", obj.Get("key"))
	})
	t.Run("Object", func(t *testing.T) {
		obj := NewObject("https://example.com")
		obj.Set("key", NewObject("https://example.com/other"))
		assert.Equal(t, map[string]interface{}{
			"@id": "https://example.com/other",
		}, obj.Get("key"))
	})
	t.Run("Entity", func(t *testing.T) {
		obj := NewObject("https://example.com")
		obj.Set("key", Entity(NewObject("https://example.com/other")))
		assert.Equal(t, map[string]interface{}{
			"@id": "https://example.com/other",
		}, obj.Get("key"))
	})
}

func TestObjectID(t *testing.T) {
	t.Run("None", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{}`))
		require.NoError(t, err)
		assert.Equal(t, "", obj.ID())
	})

	t.Run("Null", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{
			"@id": null
		}`))
		require.NoError(t, err)
		assert.Equal(t, "", obj.ID())
	})

	t.Run("Number", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{
			"@id": 12345
		}`))
		require.NoError(t, err)
		assert.Equal(t, "12345", obj.ID())
	})

	t.Run("String", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{
			"@id": "https://example.com"
		}`))
		require.NoError(t, err)
		assert.Equal(t, "https://example.com", obj.ID())
	})

	t.Run("Array", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{
			"@id": ["https://example.com/1", "https://example.com/2"]
		}`))
		require.NoError(t, err)
		assert.Equal(t, `https://example.com/1,https://example.com/2`, obj.ID())
	})
}

func TestObjectSetID(t *testing.T) {
	obj := NewObject("https://example.com")
	assert.Equal(t, "https://example.com", obj.ID())
	obj.SetID("https://example.com/boop")
	assert.Equal(t, "https://example.com/boop", obj.ID())
}

func TestObjectType(t *testing.T) {
	t.Run("None", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{}`))
		require.NoError(t, err)
		assert.Empty(t, obj.Type())
	})

	t.Run("Null", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{"@type": null}`))
		require.NoError(t, err)
		assert.Empty(t, obj.Type())
	})

	t.Run("Number", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{
			"@type": 12345
		}`))
		require.NoError(t, err)
		assert.Equal(t, []string{
			"12345",
		}, obj.Type())
	})

	t.Run("String", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{
			"@type": "http://www.w3.org/ns/activitystreams#Note"
		}`))
		require.NoError(t, err)
		assert.Equal(t, []string{
			"http://www.w3.org/ns/activitystreams#Note",
		}, obj.Type())
	})

	t.Run("Array", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{
			"@type": [
				"http://www.w3.org/ns/activitystreams#Note",
				"http://www.w3.org/ns/activitystreams#Image"
			]
		}`))
		require.NoError(t, err)
		assert.Equal(t, []string{
			"http://www.w3.org/ns/activitystreams#Note",
			"http://www.w3.org/ns/activitystreams#Image",
		}, obj.Type())
	})
}

func TestObjectSetType(t *testing.T) {
	obj := NewObject("https://example.com", "http://www.w3.org/ns/activitystreams#Note")
	assert.Equal(t, []string{"http://www.w3.org/ns/activitystreams#Note"}, obj.Type())
	obj.SetType("http://www.w3.org/ns/activitystreams#Image")
	assert.Equal(t, []string{"http://www.w3.org/ns/activitystreams#Image"}, obj.Type())
}

func TestObjectMerge(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{
			"@id": "https://example.com"
		}`))
		require.NoError(t, err)
		patch, err := ParseObject([]byte(`{
			"@id": "https://example.com",
			"http://www.w3.org/ns/activitystreams#content": "hi"
		}`))
		require.NoError(t, err)
		assert.NoError(t, obj.Apply(patch, false))
		assert.Equal(t, obj.V, map[string]interface{}{
			"@id": "https://example.com",
			"http://www.w3.org/ns/activitystreams#content": "hi",
		})
	})

	t.Run("Replace", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{
			"@id": "https://example.com",
			"http://www.w3.org/ns/activitystreams#content": "hi"
		}`))
		require.NoError(t, err)
		patch, err := ParseObject([]byte(`{
			"@id": "https://example.com",
			"http://www.w3.org/ns/activitystreams#content": "bye"
		}`))
		require.NoError(t, err)
		assert.NoError(t, obj.Apply(patch, false))
		assert.Equal(t, obj.V, map[string]interface{}{
			"@id": "https://example.com",
			"http://www.w3.org/ns/activitystreams#content": "bye",
		})
	})

	t.Run("Merge", func(t *testing.T) {
		obj, err := ParseObject([]byte(`{
			"@id": "https://example.com",
			"http://www.w3.org/ns/activitystreams#content": "hi"
		}`))
		require.NoError(t, err)
		patch, err := ParseObject([]byte(`{
			"@id": "https://example.com",
			"http://www.w3.org/ns/activitystreams#content": "bye"
		}`))
		require.NoError(t, err)
		assert.NoError(t, obj.Apply(patch, true))
		assert.Equal(t, obj.V, map[string]interface{}{
			"@id": "https://example.com",
			"http://www.w3.org/ns/activitystreams#content": []interface{}{
				"hi",
				"bye",
			},
		})
	})
}
