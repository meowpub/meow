package lib

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gregjones/httpcache"

	"github.com/liclac/meow/config"
	"github.com/liclac/meow/lib/rediscache"
)

// CreateHttpClient creates a http.Client instance which
// will cache responses inside our Redis database
func CreateHttpClient(conn redis.Conn) *http.Client {
	// Build a cache
	cache := rediscache.NewWithClient(conn, config.RedisKeyspace()+"httpcache:")
	cacheTransport := httpcache.NewTransport(cache)

	// Mark cached responses for debugging purposes
	cacheTransport.MarkCachedResponses = true

	// Return the resulting client
	return cacheTransport.Client()
}
