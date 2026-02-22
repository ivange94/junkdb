package handlers

import (
	"fmt"
	"os"
	"path"

	"github.com/ivange94/junkdb/pkg/router"
)

func Write(ctx *router.Ctx) error {
	key := ctx.Param("key")
	bodyBytes, err := ctx.Body()
	if err != nil {
		return fmt.Errorf("error reading request body: %w", err)
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error reading users home directory: %w", err)
	}
	dataFilePath := path.Join(home, ".junkdb", "data")
	file, err := os.OpenFile(dataFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // TODO: handle concurrent access
	if err != nil {
		return fmt.Errorf("error reading data file path: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	entry := fmt.Sprintf("%s,%s\n", key, string(bodyBytes))
	_, err = file.WriteString(entry)
	return err
}

func Read(ctx *router.Ctx) error {
	key := ctx.Param("key")
	fmt.Printf("Asking for the value for: %s\n", key)
	return nil
}
