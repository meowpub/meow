package models

import (
	"encoding/json"
	"testing"

	"github.com/liclac/meow/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntityStore(t *testing.T) {
	tx := TestDB.Begin()
	defer tx.Rollback()

	store := NewEntityStore(TestDB)

	t.Run("Not Found", func(t *testing.T) {
		t.Run("Snowflake", func(t *testing.T) {
			_, err := store.GetBySnowflake(12345)
			require.EqualError(t, err, "record not found")
		})
		t.Run("ID", func(t *testing.T) {
			_, err := store.GetByID("https://example.com/@jsmith")
			require.EqualError(t, err, "record not found")
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("No Snowflake", func(t *testing.T) {
			require.EqualError(t, store.Save(&Entity{Data: json.RawMessage(`{
				"@id": "https://example.com/@jsmith",
				"@type": ["http://schema.org/Person"],
				"http://schema.org/name": [{"@value": "John Smith"}]
			}`)}), "pq: null value in column \"id\" violates not-null constraint")
		})

		id, err := lib.GenSnowflake(0)
		require.NoError(t, err)

		require.NoError(t, store.Save(&Entity{ID: id, Data: json.RawMessage(`{
			"@id": "https://example.com/@jsmith",
			"@type": ["http://schema.org/Person"],
			"http://schema.org/name": [{"@value": "John Smith"}]
		}`)}))

		t.Run("Get", func(t *testing.T) {
			se, err := store.GetBySnowflake(id)
			require.NoError(t, err)
			ie, err := store.GetByID("https://example.com/@jsmith")
			require.NoError(t, err)

			assert.Equal(t, ie, se)
			assert.Equal(t, id, se.ID)

			var data map[string]interface{}
			require.NoError(t, json.Unmarshal(se.Data, &data))
			assert.Equal(t, map[string]interface{}{
				"@id":   "https://example.com/@jsmith",
				"@type": []interface{}{"http://schema.org/Person"},
				"http://schema.org/name": []interface{}{
					map[string]interface{}{"@value": "John Smith"},
				},
			}, data)
		})

		t.Run("Snowflake Reuse", func(t *testing.T) {
			require.EqualError(t, store.Save(&Entity{ID: id, Data: json.RawMessage(`{
				"@id": "https://example.com/@jdoe",
				"@type": ["http://schema.org/Person"],
				"http://schema.org/name": [{"@value": "John Smith"}]
			}`)}), "pq: duplicate key value violates unique constraint \"entities_pkey\"")
		})

		t.Run("Update", func(t *testing.T) {
			t.Run("No Snowflake", func(t *testing.T) {
				require.EqualError(t, store.Save(&Entity{Data: json.RawMessage(`{
					"@id": "https://example.com/@jsmith",
					"@type": ["http://schema.org/Person"],
					"http://schema.org/name": [{"@value": "Jane Smith"}]
				}`)}), "pq: null value in column \"id\" violates not-null constraint")
			})

			newID, err := lib.GenSnowflake(0)
			require.NoError(t, err)

			require.NoError(t, store.Save(&Entity{ID: newID, Data: json.RawMessage(`{
				"@id": "https://example.com/@jsmith",
				"@type": ["http://schema.org/Person"],
				"http://schema.org/name": [{"@value": "Jane Smith"}]
			}`)}))

			t.Run("Get", func(t *testing.T) {
				_, err := store.GetBySnowflake(newID)
				require.EqualError(t, err, "record not found")
				se, err := store.GetBySnowflake(id)
				require.NoError(t, err)
				ie, err := store.GetByID("https://example.com/@jsmith")
				require.NoError(t, err)

				assert.Equal(t, ie, se)
				assert.Equal(t, id, se.ID)

				var data map[string]interface{}
				require.NoError(t, json.Unmarshal(se.Data, &data))
				assert.Equal(t, map[string]interface{}{
					"@id":   "https://example.com/@jsmith",
					"@type": []interface{}{"http://schema.org/Person"},
					"http://schema.org/name": []interface{}{
						map[string]interface{}{"@value": "Jane Smith"},
					},
				}, data)
			})
		})
	})
}
