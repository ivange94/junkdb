package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ivange94/junkdb/internal"
	"github.com/ivange94/junkdb/internal/client"
	"github.com/urfave/cli/v3"
)

func Execute() {
	cmd := &cli.Command{
		Name: "junkdb",
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "Start database server",
				Action: func(ctx context.Context, c *cli.Command) error {
					return internal.Run()
				},
			},
			{
				Name:  "set",
				Usage: "set [key] [value]",
				Action: func(ctx context.Context, c *cli.Command) error {
					if c.Args().Len() != 2 {
						return fmt.Errorf("set requires two arguments")
					}
					return client.Put(c.Args().Get(0), c.Args().Get(1))
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
