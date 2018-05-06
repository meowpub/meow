// Package redis provides a redis interface for http caching.
package rediscache

import (
	"github.com/go-redis/redis"
	"github.com/gregjones/httpcache"
	"time"
)

// cache is an implementation of httpcache.Cache that caches responses in a
// redis server.
type cache struct {
	conn   redis.Conn
	prefix string
}

func (c cache) key(k string) string {
	return c.prefix + k
}

// Get returns the response corresponding to key if present.
func (c cache) Get(key string) (resp []byte, ok bool) {
	return c.Get(c.key(key))
}

// Set saves a response to the cache as key.
func (c cache) Set(key string, resp []byte) {
	c.conn.Set(c.key(key), resp, time.Duration(0))
}

// Delete removes the response with key from the cache.
func (c cache) Delete(key string) {
	c.Delete(c.key(key))
}

// NewWithClient returns a new Cache with the given redis connection.
func NewWithClient(client redis.Conn, prefix string) httpcache.Cache {
	return cache{
		conn:   client,
		prefix: prefix,
	}
}
