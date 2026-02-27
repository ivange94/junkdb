package main

import (
	"context"
	"log"

	"github.com/ivange94/junkdb/internal/cmd"
)

func main() {
	if err := cmd.Execute(context.Background()); err != nil {
		log.Fatal(err)
	}
}
