package models

import (
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/stretchr/testify/assert"
)

func Test_genOnConflict(t *testing.T) {
	var v struct {
		ID        snowflake.ID
		FirstName string
		LastName  string
		Email     string `gorm:"column:e_mail"`
		Priv      string `gorm:"-"`
		priv      string
	}
	assert.Equal(t, "ON CONFLICT (id) DO UPDATE SET first_name=EXCLUDED.first_name, last_name=EXCLUDED.last_name, e_mail=EXCLUDED.e_mail", genOnConflict(v, "id"))
}
