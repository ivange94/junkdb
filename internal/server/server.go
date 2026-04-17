package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/ivange94/junkdb/internal/config"
	"github.com/ivange94/junkdb/internal/storage"
)

type Server struct {
	Engine *storage.Engine
}

func (s *Server) putHandler(ctx fiber.Ctx) error {
	key := ctx.Params("key")
	value := string(ctx.Body())

	if err := s.Engine.Put(key, value); err != nil {
		return err
	}
	ctx.Status(http.StatusAccepted)
	return nil
}

func (s *Server) getHandler(ctx fiber.Ctx) error {
	key := ctx.Params("key")
	value, err := s.Engine.Get(key)
	if err != nil {
		if errors.Is(err, storage.ErrKeyNotFound) {
			ctx.Status(http.StatusNotFound)
			return ctx.SendString(err.Error())
		}
		return fmt.Errorf("error getting value for key %s: %w", key, err) // TODO: add a custom error type for missing key
	}
	return ctx.SendString(value)
}

func Run(cfg *config.Config) error {
	e, err := storage.NewEngine()
	if err != nil {
		return fmt.Errorf("connect to storage engine: %w", err)
	}

	server := &Server{Engine: e}
	app := fiber.New()

	app.Post("api/v1/:key", server.putHandler)
	app.Get("api/v1/:key", server.getHandler)

	return app.Listen(cfg.BindAddr)
}
