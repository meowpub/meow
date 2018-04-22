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
