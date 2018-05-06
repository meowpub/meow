package models

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientAccessors(t *testing.T) {
	cl, err := NewClient(ClientUserData{
		Name:        "test client",
		Description: "lorem ipsum dolor sit amet",
	}, "https://google.com/")
	require.NoError(t, err)

	t.Run("GetId", func(t *testing.T) {
		idStr := cl.GetId()
		assert.Equal(t, fmt.Sprintf("%d", cl.ID), idStr)
		id, err := strconv.ParseInt(idStr, 10, 64)
		assert.NoError(t, err)
		assert.Equal(t, int64(cl.ID), id)
	})

	t.Run("GetSecret", func(t *testing.T) {
		assert.Equal(t, cl.Secret, cl.GetSecret())
	})

	t.Run("GetRedirectUri", func(t *testing.T) {
		assert.Equal(t, cl.RedirectURI, cl.GetRedirectUri())
	})

	t.Run("GetUserData", func(t *testing.T) {
		assert.Equal(t, cl.ClientUserData, cl.GetUserData())
	})
}

func TestClientStore(t *testing.T) {
	estore := NewEntityStore(TestDB)
	ustore := NewUserStore(TestDB)

	profile, err := NewEntity([]byte(`{"@id": "https://example.com/@client-dev"}`))
	require.NoError(t, err)
	require.NoError(t, estore.Save(profile))

	user, err := NewUser(profile.ID, "client-dev@example.com", "password")
	require.NoError(t, err)
	require.NoError(t, ustore.Save(user))

	store := NewClientStore(TestDB)
	cl, err := NewClient(ClientUserData{
		Name:        "test client",
		Description: "lorem ipsum dolor sit amet",
		OwnerID:     user.ID,
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
