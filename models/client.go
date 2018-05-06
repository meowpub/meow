package models

import (
	"crypto/rand"
	"encoding/base64"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/liclac/meow/config"
	"github.com/liclac/meow/lib"
)

//go:generate mockgen -package=models -source=client.go -destination=client.mock.go

// Client is an OAuth2 client application.
// Implements the osin.Client interface for the oauth package.
type Client struct {
	ID          snowflake.ID `json:"id"`
	RedirectURI string       `json:"redirect_uri"`
	Secret      string       `json:"-"`
	ClientUserData
}

// ClientUserData contains additional fields for a client to be included as UserData.
type ClientUserData struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	OwnerID     snowflake.ID `json:"owner_id"`
}

func NewClient(ud ClientUserData, redirectURI string) (*Client, error) {
	// Generate a snowflake.
	id, err := lib.GenSnowflake(config.NodeID())
	if err != nil {
		return nil, err
	}

	// Generate a Secret.
	secretData := make([]byte, 64)
	if _, err := rand.Read(secretData); err != nil {
		return nil, err
	}
	secret := base64.RawURLEncoding.EncodeToString(secretData)

	return &Client{
		ID:             id,
		RedirectURI:    redirectURI,
		Secret:         secret,
		ClientUserData: ud,
	}, nil
}

func (c Client) GetId() string            { return strconv.FormatInt(int64(c.ID), 10) }
func (c Client) GetSecret() string        { return c.Secret }
func (c Client) GetRedirectUri() string   { return c.RedirectURI }
func (c Client) GetUserData() interface{} { return c.ClientUserData }

// ClientStore abstracts data access to Clients.
type ClientStore interface {
	// Create creates a new client.
	Create(cl *Client) error

	// Get returns a client by ID, or an error if it doesn't exist.
	Get(id snowflake.ID) (*Client, error)
}

type clientStore struct {
	DB *gorm.DB
}

// NewClientStore returns a ClientStore that accesses the given DB.
func NewClientStore(db *gorm.DB) ClientStore {
	return &clientStore{db}
}

func (s clientStore) Create(cl *Client) error {
	return s.DB.Create(cl).Error
}

func (s clientStore) Get(id snowflake.ID) (*Client, error) {
	var cl Client
	return &cl, s.DB.First(&cl, Client{ID: id}).Error
}
