package cmd

import "github.com/spf13/cobra"

var setCmd = &cobra.Command{
	Use: "set",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
