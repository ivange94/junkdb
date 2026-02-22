package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ivange94/junkdb/internal"
	"github.com/ivange94/junkdb/internal/client"
	"github.com/ivange94/junkdb/internal/config"
	"github.com/urfave/cli/v3"
)

type cfgKey struct{}

func withConfig(ctx context.Context, cfg *config.Config) context.Context {
	return context.WithValue(ctx, cfgKey{}, cfg)
}

func mustConfig(ctx context.Context) *config.Config {
	cfg, ok := ctx.Value(cfgKey{}).(*config.Config)
	if !ok || cfg == nil {
		panic("config missing from context")
	}
	return cfg
}

func Execute() {
	cmd := &cli.Command{
		Name: "junkdb",
		Before: func(ctx context.Context, c *cli.Command) (context.Context, error) {
			return withConfig(ctx, config.MustLoad()), nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "config.json",
				Sources: cli.EnvVars("CONFIG_FILE"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "Start database server",
				Action: func(ctx context.Context, c *cli.Command) error {
					return internal.Run(mustConfig(ctx))
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
