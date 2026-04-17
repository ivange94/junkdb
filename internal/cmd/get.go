package cmd

import (
	"fmt"

	"github.com/ivange94/junkdb/internal/client"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:  "get <key>",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := mustConfig(cmd.Context())
		value, err := client.New(cfg).Get(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(cmd.OutOrStdout(), value)
		return err
	},
}
