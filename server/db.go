package server

import (
	"github.com/jinzhu/gorm"
	"net/http"

	"github.com/liclac/meow/config"
	"github.com/liclac/meow/models"
	"github.com/liclac/meow/models/entities"
	"github.com/liclac/meow/server/api"
)

func DBMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			db, err := gorm.Open("postgres", config.DB())
			if err != nil {
				RenderResponse(rw, req, api.ErrorResponse(err))
			}

			rawEntityStore := models.NewEntityStore(db)
			entityStore := entities.NewStore(rawEntityStore)
			req = req.WithContext(entities.WithStore(req.Context(), entityStore))
			next.ServeHTTP(rw, req)
		})
	}
}
