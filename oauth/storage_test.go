package oauth

import (
	"testing"
	"time"

	"github.com/RangelReale/osin"
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
		ID:          "my_client",
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
		mockClientStore(store).EXPECT().Get("my_client").Return(nil, gorm.ErrRecordNotFound)

		_, err := store.GetClient("my_client")
		assert.Equal(t, err, osin.ErrNotFound)
	})

	t.Run("Found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := newMockStorage(ctrl)
		mockClientStore(store).EXPECT().Get("my_client").Return(cl, nil)

		cl2, err := store.GetClient("my_client")
		require.NoError(t, err)
		assert.Equal(t, cl, cl2)
	})
}

func TestStorageSaveAuthorize(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := newMockStorage(ctrl)
		mockAuthStore(store).EXPECT().Set("my_code", &models.Authorization{
			ClientID:    "my_client",
			Scope:       "my_scope",
			RedirectURI: "https://example.com/",
			State:       "hi",
		}, 15*time.Second).Return(nil)

		assert.NoError(t, store.SaveAuthorize(&osin.AuthorizeData{
			Client:      &models.Client{ID: "my_client"},
			Code:        "my_code",
			ExpiresIn:   15,
			Scope:       "my_scope",
			RedirectUri: "https://example.com/",
			State:       "hi",
		}))
	})

	t.Run("No TTL", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := newMockStorage(ctrl)
		mockAuthStore(store).EXPECT().Set("my_code", gomock.Any(), 0*time.Second).Return(models.ErrNoTTL)

		assert.EqualError(t, store.SaveAuthorize(&osin.AuthorizeData{
			Client:      &models.Client{ID: "my_client"},
			Code:        "my_code",
			Scope:       "my_scope",
			RedirectUri: "https://example.com/",
			State:       "hi",
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
				ClientID:    "my_client",
				Scope:       "my_scope",
				RedirectURI: "https://example.com/",
				State:       "hi",
			}, nil),
			mockClientStore(store).EXPECT().Get("my_client").Return(&models.Client{
				ID: "my_client",
			}, nil),
		)

		auth, err := store.LoadAuthorize("my_code")
		require.NoError(t, err)
		assert.Equal(t, &osin.AuthorizeData{
			Client:      &models.Client{ID: "my_client"},
			Code:        "my_code",
			Scope:       "my_scope",
			RedirectUri: "https://example.com/",
			State:       "hi",
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
