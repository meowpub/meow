package models

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthorizationStore(t *testing.T) {
	minir, err := miniredis.Run()
	require.NoError(t, err)

	r := redis.NewClient(&redis.Options{Addr: minir.Addr()})
	store := NewAuthorizationStore("meow", r)
	auth := &Authorization{
		Code:        "c049a780-8e07-44aa-9b2c-7bdae5af728d",
		ClientID:    "f41d0dea-7a3b-4ca6-8b16-2eefe07b28f0",
		Scope:       "myscope",
		RedirectURI: "https://google.com/",
		State:       "weh",
		AuthorizationUserData: AuthorizationUserData{
			UserID: 12345,
		},
	}

	require.Len(t, r.Keys("*").Val(), 0)

	t.Run("Not Found", func(t *testing.T) {
		_, err := store.Get(auth.Code)
		assert.EqualError(t, err, "invalid authorization code")
		assert.True(t, IsNotFound(err))
	})

	t.Run("Set", func(t *testing.T) {
		t.Run("NoTTL", func(t *testing.T) {
			assert.EqualError(t, store.Set(auth, 0), "a TTL is required, but not given")
		})

		assert.NoError(t, store.Set(auth, 15*time.Minute))
		assert.Equal(t, r.Keys("*").Val(), []string{"meow:oauth2:authorizations:" + auth.Code})
	})

	t.Run("Get", func(t *testing.T) {
		auth2, err := store.Get(auth.Code)
		require.NoError(t, err)
		assert.Equal(t, auth, auth2)
	})

	t.Run("Delete", func(t *testing.T) {
		assert.NoError(t, store.Delete(auth.Code))
		require.Len(t, r.Keys("*").Val(), 0)
	})
}
