package api

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/meowpub/meow/models"
)

func TestErrorStatusAndWrap(t *testing.T) {
	assert.Equal(t, http.StatusOK, ErrorStatus(nil))
	assert.Equal(t, http.StatusInternalServerError, ErrorStatus(errors.New("error")))
	assert.Equal(t, http.StatusNotFound, ErrorStatus(models.NotFoundError("error")))
	assert.Equal(t, http.StatusBadRequest, ErrorStatus(errors.Wrap(Wrap(errors.New("error"), http.StatusBadRequest), "wrapped error")))
}

func TestStackTrace(t *testing.T) {
	assert.Nil(t, StackTrace(nil))
	assert.Nil(t, StackTrace(models.NotFoundError("error")))
	assert.NotEmpty(t, StackTrace(errors.New("error")))
	assert.NotEmpty(t, StackTrace(errors.Wrap(errors.New("error"), "wrapped error")))
	assert.NotEmpty(t, StackTrace(Wrap(errors.Wrap(errors.New("error"), "wrapped error"), http.StatusBadRequest)))
}
