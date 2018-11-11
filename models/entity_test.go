package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/meowpub/meow/lib"
)

func TestEntityConflictClause(t *testing.T) {
	assert.Equal(t, `ON CONFLICT ((data->>'@id')) DO UPDATE SET data=EXCLUDED.data, kind=EXCLUDED.kind`, entityOnConflict)
}

func TestEntityStore(t *testing.T) {
	tx := TestDB.Begin()
	defer tx.Rollback()

	store := NewEntityStore(TestDB)

	t.Run("Not Found", func(t *testing.T) {
		t.Run("Snowflake", func(t *testing.T) {
			_, err := store.GetBySnowflake(12345)
			require.EqualError(t, err, "record not found")
			assert.True(t, IsNotFound(err))
		})
		t.Run("ID", func(t *testing.T) {
			_, err := store.GetByID("https://example.com/@jsmith")
			require.EqualError(t, err, "record not found")
			assert.True(t, IsNotFound(err))
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("No Snowflake", func(t *testing.T) {
			require.EqualError(t, store.Save(&Entity{Kind: ObjectEntity, Data: JSONB(`{
				"@id": "https://example.com/@jsmith",
				"@type": ["http://schema.org/Person"],
				"http://schema.org/name": [{"@value": "John Smith"}]
			}`)}), "pq: null value in column \"id\" violates not-null constraint")
		})

		id, err := lib.GenSnowflake(0)
		require.NoError(t, err)

		t.Run("Empty", func(t *testing.T) {
			require.EqualError(t, store.Save(&Entity{ID: id, Kind: ObjectEntity, Data: JSONB(`{}`)}), "pq: new row for relation \"entities\" violates check constraint \"entities_data_id_not_null_check\"")
		})

		t.Run("Empty ID", func(t *testing.T) {
			require.EqualError(t, store.Save(&Entity{ID: id, Kind: ObjectEntity, Data: JSONB(`{"@id":""}`)}), "pq: new row for relation \"entities\" violates check constraint \"entities_data_id_not_empty_check\"")
		})

		t.Run("Invalid ID", func(t *testing.T) {
			require.EqualError(t, store.Save(&Entity{ID: id, Kind: ObjectEntity, Data: JSONB(`{"@id":"test"}`)}), "pq: new row for relation \"entities\" violates check constraint \"entities_data_id_protocol_check\"")
		})

		t.Run("Invalid ID Protocol", func(t *testing.T) {
			require.EqualError(t, store.Save(&Entity{ID: id, Kind: ObjectEntity, Data: JSONB(`{"@id":"ftp://example.com/~jsmith"}`)}), "pq: new row for relation \"entities\" violates check constraint \"entities_data_id_protocol_check\"")
		})

		require.NoError(t, store.Save(&Entity{ID: id, Data: JSONB(`{
			"@id": "https://example.com/@jsmith",
			"@type": ["http://schema.org/Person"],
			"http://schema.org/name": [{"@value": "John Smith"}]
		}`), Kind: "weird"}))

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

			t.Run("Object", func(t *testing.T) {
				assert.Equal(t, "https://example.com/@jsmith", se.Obj.ID())
				assert.Equal(t, se.Obj.ID(), ie.Obj.ID())

				assert.Equal(t, []string{"http://schema.org/Person"}, se.Obj.Type())
				assert.Equal(t, se.Obj.Type(), ie.Obj.Type())

				assert.Equal(t, map[string]interface{}{
					"@id":   "https://example.com/@jsmith",
					"@type": []interface{}{"http://schema.org/Person"},
					"http://schema.org/name": []interface{}{
						map[string]interface{}{"@value": "John Smith"},
					},
				}, se.Obj.V)
				assert.Equal(t, se.Obj.V, ie.Obj.V)

				t.Run("Update", func(t *testing.T) {
					se.Obj.V["http://schema.org/name"] = []interface{}{
						map[string]interface{}{"@value": "Jane Smith"},
					}
					require.NoError(t, store.Save(se))

					t.Run("Get", func(t *testing.T) {
						se2, err := store.GetBySnowflake(id)
						require.NoError(t, err)
						ie2, err := store.GetByID("https://example.com/@jsmith")
						require.NoError(t, err)

						assert.Equal(t, []interface{}{
							map[string]interface{}{"@value": "Jane Smith"},
						}, se2.Obj.Get("http://schema.org/name"))
						assert.Equal(t, se2.Obj.V, ie2.Obj.V)
					})
				})
			})
		})

		t.Run("Update", func(t *testing.T) {
			newID, err := lib.GenSnowflake(0)
			require.NoError(t, err)

			require.NoError(t, store.Save(&Entity{ID: newID, Data: JSONB(`{
				"@id": "https://example.com/@jsmith",
				"@type": ["http://schema.org/Person"],
				"http://schema.org/name": [{"@value": "Jane Smith"}]
			}`), Kind: "Success!"}))

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
						map[string]interface{}{"@value": "Jane Smith"},
					},
				}, data)
			})
		})
	})
}
