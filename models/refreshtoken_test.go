package models

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRefreshTokenStore(t *testing.T) {
	minir, err := miniredis.Run()
	require.NoError(t, err)

	r := redis.NewClient(&redis.Options{Addr: minir.Addr()})
	store := NewRefreshTokenStore("meow", r)
	token := &RefreshToken{
		Token:    "8f9f21e6-d7d9-44f4-a018-295916935375",
		ClientID: "f41d0dea-7a3b-4ca6-8b16-2eefe07b28f0",
		Scope:    "myscope",
		AuthorizationUserData: AuthorizationUserData{
			UserID: 12345,
		},
	}

	require.Len(t, r.Keys("*").Val(), 0)

	t.Run("Not Found", func(t *testing.T) {
		_, err := store.Get(token.Token)
		assert.EqualError(t, err, "token is invalid or expired")
		assert.True(t, IsNotFound(err))
	})

	t.Run("Set", func(t *testing.T) {
		t.Run("NoTTL", func(t *testing.T) {
			assert.EqualError(t, store.Set(token, 0), "a TTL is required, but not given")
		})

		assert.NoError(t, store.Set(token, 15*time.Minute))
		assert.Equal(t, r.Keys("*").Val(), []string{"meow:oauth2:refresh_tokens:" + token.Token})
	})

	t.Run("Get", func(t *testing.T) {
		token2, err := store.Get(token.Token)
		require.NoError(t, err)
		assert.Equal(t, token, token2)
	})

	t.Run("Delete", func(t *testing.T) {
		assert.NoError(t, store.Delete(token.Token))
		require.Len(t, r.Keys("*").Val(), 0)
	})
}
