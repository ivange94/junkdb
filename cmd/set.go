package cmd

import (
	"fmt"
	"log"

	"github.com/ivange94/junkdb/internal/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Write key,value pair to the db",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("need exactly two arguments")
			return
		}
		err := client.Put(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
	},
}
