package lib

import (
	"context"
	"testing"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestContextLogger(t *testing.T) {
	l, err := zap.NewDevelopment()
	require.NoError(t, err)

	t.Run("Global", func(t *testing.T) {
		defer zap.ReplaceGlobals(l)()
		assert.Equal(t, l, GetLogger(context.Background()))

		t.Run("nil", func(t *testing.T) {
			assert.Equal(t, l, GetLogger(nil))
		})
	})

	t.Run("Named", func(t *testing.T) {
		l2 := GetLogger(WithNamedLogger(context.Background(), "sub"))
		assert.NotNil(t, l2)
		assert.NotEqual(t, l, l2)
	})

	t.Run("Local", func(t *testing.T) {
		assert.Equal(t, l, GetLogger(WithLogger(context.Background(), l)))
	})
}

func TestContextDB(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	assert.Equal(t, db, GetDB(WithDB(context.Background(), db)))
}

func TestContextRedis(t *testing.T) {
	r := redis.NewClient(&redis.Options{})
	assert.Equal(t, r, GetRedis(WithRedis(context.Background(), r)))
}
