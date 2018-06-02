package config

import (
	"github.com/spf13/viper"
)

// IsProd returns whether we're running in production mode.
func IsProd() bool {
	return viper.GetBool("prod")
}

// Secret returns the master secret, as a hex string.
// Do not use this key directly; use config/secrets to generate per-task derivative keys.
func Secret() string {
	return viper.GetString("secret")
}

// DB returns the database connection string.
func DB() string {
	return viper.GetString("db")
}

// Redis returns the redis connection string.
func Redis() string {
	return viper.GetString("redis")
}

// RedisKeyspace returns the keyspace for redis, eg. "meow" will prefix all redis keys with "meow:".
func RedisKeyspace() string {
	return viper.GetString("redis-keyspace")
}

// NodeID returns the node ID used for distributed snowflake generation.
func NodeID() int64 {
	return viper.GetInt64("node-id")
}

// Syntax style for JSON output.
func HighlightStyle() string {
	return viper.GetString("highlight-style")
}

// Disable colour in the CLI.
func NoColour() bool {
	return viper.GetBool("no-colour")
}
