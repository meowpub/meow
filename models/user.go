package models

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/liclac/meow/config"
	"github.com/liclac/meow/lib"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -package=models -source=user.go -destination=user.mock.go

var userOnConflict = genOnConflict(User{}, "id")

// A User stores authentication data for a local user. This is separate from profile information,
// which is stored as an Entity. This should never be sent to clients, and only used internally.
// To get the public ID, etc., please look up the related Entity.
type User struct {
	ID           snowflake.ID `json:"id"`
	EntityID     snowflake.ID `json:"entity_id"`
	Email        string       `json:"email"`
	PasswordHash string       `json:"-"`
}

func NewUser(profileID snowflake.ID, email, password string) (*User, error) {
	id, err := lib.GenSnowflake(config.NodeID())
	if err != nil {
		return nil, err
	}

	user := User{
		ID:       id,
		EntityID: profileID,
		Email:    email,
	}
	return &user, user.SetPassword(password)
}

// SetPassword updates the user's PasswordHash by bcrypt'ing the password.
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	u.PasswordHash = string(hash)
	return err
}

// CheckPassword compares the given password to the stored hash. Returns whether the comparison
// succeeded, and any errors produced. Please note that if an error is produced here, you should
// NOT RELAY IT VERBATIM TO THE USER; log it, check the bool and return a generic "unauthorized"
// response if it's false.
func (u *User) CheckPassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	return err == nil, err
}

type UserStore interface {
	GetBySnowflake(id snowflake.ID) (*User, error)
	GetByEntityID(id snowflake.ID) (*User, error)
	GetByEmail(email string) (*User, error)
	Save(u *User) error
}

type userStore struct {
	DB *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{db}
}

func (s *userStore) GetBySnowflake(id snowflake.ID) (*User, error) {
	var u User
	return &u, s.DB.First(&u, User{ID: id}).Error
}

func (s *userStore) GetByEntityID(eid snowflake.ID) (*User, error) {
	var u User
	return &u, s.DB.First(&u, User{EntityID: eid}).Error
}

func (s *userStore) GetByEmail(email string) (*User, error) {
	var u User
	return &u, s.DB.First(&u, User{Email: email}).Error
}

func (s *userStore) Save(u *User) error {
	return s.DB.Set(gormInsertOption, userOnConflict).Create(u).Error
}
