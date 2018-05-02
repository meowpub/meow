package models

import (
	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/bcrypt"
)

// A User stores authentication data for a local user. This is separate from profile information,
// which is stored as an Entity. This should never be sent to clients, and only used internally.
type User struct {
	ID           snowflake.ID `json:"id"`
	EntityID     snowflake.ID `json:"entity_id"`
	Email        string       `json:"email"`
	PasswordHash string       `json:"-"`
	PrivateKey   string       `json:"-"`
}

// SetPassword updates the user's PasswordHash by bcrypt'ing the password.
func (u *User) SetPassword(plainPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 0)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}
