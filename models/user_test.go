package models

import (
	"testing"

	"github.com/liclac/meow/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserConflictClause(t *testing.T) {
	assert.Equal(t, "ON CONFLICT (id) DO UPDATE SET entity_id=EXCLUDED.entity_id, email=EXCLUDED.email, password_hash=EXCLUDED.password_hash", userOnConflict)
}

func TestUserPasswords(t *testing.T) {
	user := &User{}
	require.NoError(t, user.SetPassword("password"))

	ok, err := user.CheckPassword("abc123")
	assert.NoError(t, err)
	assert.False(t, ok, "incorrect password accepted!!")

	ok, err = user.CheckPassword("")
	assert.NoError(t, err)
	assert.False(t, ok, "blank password accepted!!")

	ok, err = user.CheckPassword("password")
	assert.NoError(t, err)
	assert.True(t, ok, "correct password rejected")
}

func TestUserStore(t *testing.T) {
	tx := TestDB.Begin()
	defer tx.Rollback()

	store := NewUserStore(tx)
	estore := NewEntityStore(tx)

	t.Run("Not Found", func(t *testing.T) {
		t.Run("Snowflake", func(t *testing.T) {
			_, err := store.GetBySnowflake(12345)
			require.EqualError(t, err, "record not found")
		})
		t.Run("EntityID", func(t *testing.T) {
			_, err := store.GetByEntityID(12345)
			require.EqualError(t, err, "record not found")
		})
		t.Run("Email", func(t *testing.T) {
			_, err := store.GetByEmail("test@example.com")
			require.EqualError(t, err, "record not found")
		})
	})

	eid, err := lib.GenSnowflake(0)
	require.NoError(t, err)
	entity := &Entity{ID: eid, Data: []byte(`{
		"@id": "https://example.com/~jsmith",
		"@type": ["http://schema.org/Person"],
		"http://schema.org/name": [{"@value": "John Smith"}]
	}`)}
	require.NoError(t, estore.Save(entity))

	id, err := lib.GenSnowflake(0)
	require.NoError(t, err)
	user := &User{
		ID:       id,
		EntityID: eid,
		Email:    "jsmith@example.com",
	}
	assert.NoError(t, user.SetPassword("password"))

	t.Run("Password", func(t *testing.T) {
		ok, err := user.CheckPassword("password")
		assert.NoError(t, err)
		assert.True(t, ok, "password not accepted")
	})

	t.Run("Create", func(t *testing.T) {
		require.NoError(t, store.Save(user))
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("Snowflake", func(t *testing.T) {
			u2, err := store.GetBySnowflake(id)
			require.NoError(t, err)
			assert.Equal(t, user, u2)
		})
		t.Run("EntityID", func(t *testing.T) {
			u2, err := store.GetByEntityID(entity.ID)
			require.NoError(t, err)
			assert.Equal(t, user, u2)
		})
		t.Run("Email", func(t *testing.T) {
			u2, err := store.GetByEmail("jsmith@example.com")
			require.NoError(t, err)
			assert.Equal(t, user, u2)
		})
	})

	t.Run("Change Email", func(t *testing.T) {
		user.Email = "jsmith2@example.com"
		require.NoError(t, store.Save(user))

		u2, err := store.GetBySnowflake(id)
		require.NoError(t, err)
		assert.Equal(t, user, u2)
	})

	t.Run("Change Password", func(t *testing.T) {
		require.NoError(t, user.SetPassword("new_password"))
		require.NoError(t, store.Save(user))

		u2, err := store.GetBySnowflake(id)
		require.NoError(t, err)
		assert.Equal(t, user, u2)

		t.Run("Password", func(t *testing.T) {
			ok, err := u2.CheckPassword("password")
			assert.NoError(t, err)
			assert.False(t, ok, "old password accepted!!")

			ok, err = u2.CheckPassword("new_password")
			assert.NoError(t, err)
			assert.True(t, ok, "new password not accepted")
		})
	})
}
