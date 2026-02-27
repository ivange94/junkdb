package cmd

import (
	"github.com/ivange94/junkdb/internal/server"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := mustConfig(cmd.Context())
		return server.Run(cfg)
	},
}
