package internal

import (
	"net/http"

	"github.com/ivange94/junkdb/internal/handlers"
	"github.com/ivange94/junkdb/pkg/router"
)

func Run() error {
	r := router.New()

	r.Post("/api/v1/{key}", handlers.Write)
	r.Get("/api/v1/{key}", handlers.Read)

	return http.ListenAndServe(":9000", r)
}
