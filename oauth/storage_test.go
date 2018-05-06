package oauth

import (
	"testing"
	"time"

	"github.com/RangelReale/osin"
	"github.com/bwmarrin/snowflake"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/liclac/meow/models"
)

func newMockStorage(ctrl *gomock.Controller) *Storage {
	return NewStorage(
		models.NewMockClientStore(ctrl),
		models.NewMockAuthorizationStore(ctrl),
	)
}

func mockClientStore(s *Storage) *models.MockClientStore {
	return s.Clients.(*models.MockClientStore)
}

func mockAuthStore(s *Storage) *models.MockAuthorizationStore {
	return s.Auths.(*models.MockAuthorizationStore)
}

func TestStorage(t *testing.T) {
	t.Run("Clone", func(t *testing.T) {
		store := &Storage{}
		require.Equal(t, store, store.Clone())
	})
	t.Run("Close", func(t *testing.T) {
		// This is a NOP lol
		Storage{}.Close()
	})
}

func TestStorageGetClient(t *testing.T) {
	cl := &models.Client{
		ID:          353894652568535040,
		RedirectURI: "https://example.com/",
		Secret:      "hello world",
		ClientUserData: models.ClientUserData{
			Name:        "my client",
			Description: "lorem ipsum dolor sit amet",
		},
	}

	t.Run("NotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := newMockStorage(ctrl)
		mockClientStore(store).EXPECT().Get(
			snowflake.ID(353894652568535040),
		).Return(nil, gorm.ErrRecordNotFound)

		_, err := store.GetClient("353894652568535040")
		assert.Equal(t, osin.ErrNotFound, err)
	})

	t.Run("Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := newMockStorage(ctrl)
		mockClientStore(store).EXPECT().Get(
			snowflake.ID(353894652568535040),
		).Return(cl, nil)

		cl2, err := store.GetClient("353894652568535040")
		require.NoError(t, err)
		assert.Equal(t, cl, cl2)
	})
}

func TestStorageSaveAuthorize(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := newMockStorage(ctrl)
		mockAuthStore(store).EXPECT().Set(&models.Authorization{
			Code:        "my_code",
			ClientID:    "353894652568535040",
			Scope:       "my_scope",
			RedirectURI: "https://example.com/",
			State:       "hi",
			AuthorizationUserData: models.AuthorizationUserData{
				UserID: 12345,
			},
		}, 15*time.Second).Return(nil)

		assert.NoError(t, store.SaveAuthorize(&osin.AuthorizeData{
			Client:      &models.Client{ID: 353894652568535040},
			Code:        "my_code",
			ExpiresIn:   15,
			Scope:       "my_scope",
			RedirectUri: "https://example.com/",
			State:       "hi",
			UserData: models.AuthorizationUserData{
				UserID: 12345,
			},
		}))
	})

	t.Run("No TTL", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := newMockStorage(ctrl)
		mockAuthStore(store).EXPECT().Set(gomock.Any(), 0*time.Second).Return(models.ErrNoTTL)

		assert.EqualError(t, store.SaveAuthorize(&osin.AuthorizeData{
			Client:      &models.Client{ID: 353894652568535040},
			Code:        "my_code",
			Scope:       "my_scope",
			RedirectUri: "https://example.com/",
			State:       "hi",
			UserData: models.AuthorizationUserData{
				UserID: 12345,
			},
		}), "a TTL is required, but not given")
	})
}

func TestStorageLoadAuthorize(t *testing.T) {
	t.Run("Not Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := newMockStorage(ctrl)
		mockAuthStore(store).EXPECT().Get("my_code").Return(nil, models.NotFoundError("not found"))

		_, err := store.LoadAuthorize("my_code")
		assert.Equal(t, err, osin.ErrNotFound)
	})

	t.Run("Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := newMockStorage(ctrl)
		gomock.InOrder(
			mockAuthStore(store).EXPECT().Get("my_code").Return(&models.Authorization{
				ClientID:    "353894652568535040",
				Scope:       "my_scope",
				RedirectURI: "https://example.com/",
				State:       "hi",
				AuthorizationUserData: models.AuthorizationUserData{
					UserID: 12345,
				},
			}, nil),
			mockClientStore(store).EXPECT().Get(
				snowflake.ID(353894652568535040),
			).Return(&models.Client{
				ID: 353894652568535040,
			}, nil),
		)

		auth, err := store.LoadAuthorize("my_code")
		require.NoError(t, err)
		assert.Equal(t, &osin.AuthorizeData{
			Client:      &models.Client{ID: 353894652568535040},
			Code:        "my_code",
			Scope:       "my_scope",
			RedirectUri: "https://example.com/",
			State:       "hi",
			UserData: models.AuthorizationUserData{
				UserID: 12345,
			},
		}, auth)
	})
}

func TestStorageRemoveAuthorize(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := newMockStorage(ctrl)
	mockAuthStore(store).EXPECT().Delete("my_code")
	assert.NoError(t, store.RemoveAuthorize("my_code"))
}
