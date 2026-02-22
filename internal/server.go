package internal

import (
	"log"
	"net/http"

	"github.com/ivange94/junkdb/internal/config"
	"github.com/ivange94/junkdb/internal/handlers"
	"github.com/ivange94/junkdb/pkg/router"
)

func Run(cfg *config.Config) error {
	r := router.New()

	r.Post("/api/v1/{key}", handlers.Write)
	r.Get("/api/v1/{key}", handlers.Read)

	log.Printf("Database server listening on %s", cfg.BindAddr)
	return http.ListenAndServe(cfg.BindAddr, r)
}
