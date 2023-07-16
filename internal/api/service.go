package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type MetroServiceRouter struct {
	Router chi.Router
	// TODO: add logging and mongo support
}

func (mr MetroServiceRouter) Test(res http.ResponseWriter, req *http.Request) {

	var payload = map[string]interface{}{
		"status": true,
		"error":  nil,
		"data":   "test_data",
	}
	writeJSON(res, req, http.StatusOK, payload)

}

func (mr MetroServiceRouter) RegisterRoutes(r chi.Router) {
	mr.Router = r
	reqGroup := r.Group(func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.Recoverer)
	})

	reqGroup.Get("/test", mr.Test)

}
