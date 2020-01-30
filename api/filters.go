package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/monzo/typhon"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func StoresFilter(db *gorm.DB, r *redis.Client) typhon.Filter {
	return func(req typhon.Request, next typhon.Service) typhon.Response {
		req.Context = lib.WithDB(req, db)
		req.Context = lib.WithRedis(req, r)
		req.Context = models.WithStores(req, models.NewStores(db, r, config.RedisKeyspace()))
		req.Request = *req.WithContext(req)
		return next(req)
	}
}

func ErrorFilter(req typhon.Request, next typhon.Service) typhon.Response {
	var rsp typhon.Response
	if err := req.Err(); err != nil {
		rsp = typhon.NewResponse(req)
		rsp.Error = err
	} else {
		rsp = next(req)
	}

	if rsp.Response == nil {
		rsp.Response = &http.Response{}
	}
	if rsp.Request == nil {
		rsp.Request = &req
	}

	if rsp.Error != nil && rsp.Error.Error() != "" {
		// TODO: Grab stack traces out of the error, if possible.
		errV := map[string]interface{}{
			"error": rsp.Error.Error(),
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetIndent("", "  ")
		if err := enc.Encode(errV); err != nil {
			lib.GetLogger(req).DPanic("Failed to marshal error response", zap.Error(err))
		}

		if rsp.Body != nil {
			rsp.Body.Close()
		}
		rsp.Body = ioutil.NopCloser(&buf)

		if rsp.StatusCode == 0 || rsp.StatusCode == http.StatusOK {
			rsp.StatusCode = ErrorCode(rsp.Error)
		}
	}
	return rsp
}

func PanicFilter(req typhon.Request, next typhon.Service) (rsp typhon.Response) {
	defer func() {
		if v := recover(); v != nil {
			lib.GetLogger(req).Error("panic: " + fmt.Sprint(v))
			err, ok := v.(error)
			if !ok {
				err = errors.Errorf("panic: %v", v)
			}
			rsp = typhon.Response{Error: err}
		}
	}()
	return next(req)
}
