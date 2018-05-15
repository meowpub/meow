package lib

import (
	"context"

	"go.uber.org/zap"
)

// Report logs an error to the global logger. Useful for `defer conn.Close()`-type constructs,
// where there's not really anything useful to do with the error, but you still want to log it.
func Report(ctx context.Context, err error, msg string) {
	if err != nil {
		GetLogger(ctx).Error(msg, zap.Error(err))
	}
}
