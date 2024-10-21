package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mytodo", // invoke command using this
	Short: "A simple todo list CLI to manage your tasks.",
	Run: func(cmd *cobra.Command, args []string){},

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("error executing mytodo command")
	}
	os.Exit(0)
}
