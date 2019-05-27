package lib

import (
	"github.com/bwmarrin/snowflake"
)

//go:generate mockgen -package=lib -source=snowflakes.go -destination=snowflakes.mock.go

var DefaultSnowflakeGenerator SnowflakeGenerator

func init() {
	node, err := snowflake.NewNode(0)
	if err != nil {
		panic(err)
	}
	DefaultSnowflakeGenerator = node
}

type SnowflakeGenerator interface {
	Generate() snowflake.ID
}

// GenSnowflake generates a snowflake ID.
func GenSnowflake() snowflake.ID {
	return DefaultSnowflakeGenerator.Generate()
}
