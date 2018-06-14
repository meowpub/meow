package models

import (
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/stretchr/testify/require"
)

func TestStreamItemStore(t *testing.T) {
	tx := TestDB.Begin()
	defer tx.Rollback()

	store := NewStreamItemStore(tx)

	// Empty DB - everything should be not found

	t.Run("Not Found", func(t *testing.T) {
		_, err := store.GetItem(snowflake.ID(353894652568535040))
		require.EqualError(t, err, "record not found")
	})

	t.Run("Empty", func(t *testing.T) {
		l, err := store.GetItems(snowflake.ID(353894652568535040), Beginning, After, 20)
		require.NoError(t, err, "GetItems")
		require.Equal(t, 0, len(l))
	})

	t.Run("Not Found by Entity ID", func(t *testing.T) {
		_, err := store.GetItemByEntityID(snowflake.ID(353894652568535040), snowflake.ID(384408932061417472))
		require.EqualError(t, err, "record not found")
	})

	// Now we should insert some items into a stream
	t.Run("Insert", func(t *testing.T) {
		t.Run("Candles", func(t *testing.T) {
			b, i, err := store.TryInsertItem(snowflake.ID(353894652568535040), snowflake.ID(384408932061417472))
			require.NoError(t, err, "TryInsertItem")
			require.Equal(t, true, b, "Should have inserted")
			require.Equal(t, snowflake.ID(353894652568535040), i.StreamID, "stream ID should match")
			require.Equal(t, snowflake.ID(384408932061417472), i.EntityID, "entity ID should match")
		})

		t.Run("No", func(t *testing.T) {
			b, i, err := store.TryInsertItem(snowflake.ID(353894652568535040), snowflake.ID(384411458794057728))
			require.NoError(t, err, "TryInsertItem")
			require.Equal(t, true, b, "Should have inserted")
			require.Equal(t, snowflake.ID(353894652568535040), i.StreamID, "stream ID should match")
			require.Equal(t, snowflake.ID(384411458794057728), i.EntityID, "entity ID should match")
		})
	})

	// Check we don't insert duplicates
	t.Run("Duplicates", func(t *testing.T) {
		t.Run("Candles", func(t *testing.T) {
			b, i, err := store.TryInsertItem(snowflake.ID(353894652568535040), snowflake.ID(384408932061417472))
			require.NoError(t, err, "TryInsertItem")
			require.Equal(t, false, b, "Should not have inserted")
			require.Equal(t, snowflake.ID(353894652568535040), i.StreamID, "stream ID should match")
			require.Equal(t, snowflake.ID(384408932061417472), i.EntityID, "entity ID should match")
		})

		t.Run("No", func(t *testing.T) {
			b, i, err := store.TryInsertItem(snowflake.ID(353894652568535040), snowflake.ID(384411458794057728))
			require.NoError(t, err, "TryInsertItem")
			require.Equal(t, false, b, "Should not have inserted")
			require.Equal(t, snowflake.ID(353894652568535040), i.StreamID, "stream ID should match")
			require.Equal(t, snowflake.ID(384411458794057728), i.EntityID, "entity ID should match")
		})
	})

	// Fetch items back out of stream
	t.Run("Fetch", func(t *testing.T) {
		t.Run("Forwards", func(t *testing.T) {
			l, err := store.GetItems(snowflake.ID(353894652568535040), Beginning, After, 20)
			require.NoError(t, err, "GetItems")
			require.Equal(t, 2, len(l), "2 items")
			require.Equal(t, snowflake.ID(384408932061417472), l[0].EntityID, "id 1")
			require.Equal(t, snowflake.ID(384411458794057728), l[1].EntityID, "id 2")
		})

		t.Run("Forwards from End", func(t *testing.T) {
			l, err := store.GetItems(snowflake.ID(353894652568535040), End, After, 20)
			require.NoError(t, err, "GetItems")
			require.Equal(t, 0, len(l), "0 items")
		})

		t.Run("Backwards", func(t *testing.T) {
			l, err := store.GetItems(snowflake.ID(353894652568535040), End, Before, 20)
			require.NoError(t, err, "GetItems")
			require.Equal(t, 2, len(l), "2 items")
			require.Equal(t, snowflake.ID(384411458794057728), l[0].EntityID, "id 1")
			require.Equal(t, snowflake.ID(384408932061417472), l[1].EntityID, "id 2")
		})

		t.Run("Backwards from Beginning", func(t *testing.T) {
			l, err := store.GetItems(snowflake.ID(353894652568535040), Beginning, Before, 20)
			require.NoError(t, err, "GetItems")
			require.Equal(t, 0, len(l), "0 items")
		})

		t.Run("After First", func(t *testing.T) {
			i, err := store.GetItemByEntityID(snowflake.ID(353894652568535040), snowflake.ID(384408932061417472))
			require.NoError(t, err, "first record should be found")

			l, err := store.GetItems(snowflake.ID(353894652568535040), i.ItemID, After, 20)
			require.NoError(t, err, "GetItems")
			require.Equal(t, 1, len(l), "1 items")
			require.Equal(t, snowflake.ID(384411458794057728), l[0].EntityID, "id 1")
		})

		t.Run("After Second", func(t *testing.T) {
			i, err := store.GetItemByEntityID(snowflake.ID(353894652568535040), snowflake.ID(384411458794057728))
			require.NoError(t, err, "second record should be found")

			l, err := store.GetItems(snowflake.ID(353894652568535040), i.ItemID, After, 20)
			require.NoError(t, err, "GetItems")
			require.Equal(t, 0, len(l), "0 items")
		})

		t.Run("Before Second", func(t *testing.T) {
			i, err := store.GetItemByEntityID(snowflake.ID(353894652568535040), snowflake.ID(384411458794057728))
			require.NoError(t, err, "second record should be found")

			l, err := store.GetItems(snowflake.ID(353894652568535040), i.ItemID, Before, 20)
			require.NoError(t, err, "GetItems")
			require.Equal(t, 1, len(l), "1 items")
			require.Equal(t, snowflake.ID(384408932061417472), l[0].EntityID, "id 1")
		})

		t.Run("Before First", func(t *testing.T) {
			i, err := store.GetItemByEntityID(snowflake.ID(353894652568535040), snowflake.ID(384408932061417472))
			require.NoError(t, err, "first record should be found")

			l, err := store.GetItems(snowflake.ID(353894652568535040), i.ItemID, Before, 20)
			require.NoError(t, err, "GetItems")
			require.Equal(t, 0, len(l), "0 items")
		})

		t.Run("Empty Stream", func(t *testing.T) {
			l, err := store.GetItems(snowflake.ID(384408932061417472), Beginning, After, 20)
			require.NoError(t, err, "GetItems")
			require.Equal(t, 0, len(l), "0 items")
		})
	})
}
