package models

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestIsNotFound(t *testing.T) {
	nf := NotFoundError("Not found!")

	require.True(t, IsNotFound(nf), "NotFound error should be IsNotFound")

	wrapped := errors.Wrap(nf, "Foo")
	require.True(t, IsNotFound(wrapped), "Wrapped NotFound error should be IsNotFound")
}
