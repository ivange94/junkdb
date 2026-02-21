package cmd

import (
	"log"

	"github.com/ivange94/junkdb/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the db server",
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}
