package config

import (
	"github.com/spf13/viper"
)

// IsProd returns whether we're running in production mode.
func IsProd() bool {
	return viper.GetBool("prod")
}

// DB returns the database connection string.
func DB() string {
	return viper.GetString("db")
}

// Redis returns the redis connection string.
func Redis() string {
	return viper.GetString("redis")
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
