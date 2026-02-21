package handlers

import (
	"fmt"

	"github.com/ivange94/junkdb/pkg/router"
)

func Write(ctx *router.Ctx) error {
	key := ctx.Param("key")
	bodyBytes, err := ctx.Body()
	if err != nil {
		return fmt.Errorf("error reading request body: %w", err)
	}
	fmt.Printf("Writing %s,%s\n", key, string(bodyBytes))
	return nil
}

func Read(ctx *router.Ctx) error {
	key := ctx.Param("key")
	fmt.Printf("Asking for the value for: %s\n", key)
	return nil
}
