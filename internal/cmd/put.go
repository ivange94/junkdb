package cmd

import (
	"github.com/ivange94/junkdb/internal/client"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:  "put <key> <value>",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := mustConfig(cmd.Context())
		return client.New(cfg).Put(cmd.Context(), args[0], args[1])
	},
}
