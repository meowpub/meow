package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientStore(t *testing.T) {
	tx := TestDB.Begin()
	defer tx.Rollback()

	store := NewClientStore(TestDB)
	cl, err := NewClient(ClientUserData{
		Name:        "test client",
		Description: "lorem ipsum dolor sit amet",
	}, "https://google.com/")

	t.Run("NotFound", func(t *testing.T) {
		_, err := store.Get(cl.ID)
		require.EqualError(t, err, "record not found")
	})

	t.Run("Create", func(t *testing.T) {
		require.NoError(t, err)
		require.NoError(t, store.Create(cl))
	})

	t.Run("Get", func(t *testing.T) {
		cl2, err := store.Get(cl.ID)
		require.NoError(t, err)
		assert.Equal(t, cl, cl2)
	})
}
