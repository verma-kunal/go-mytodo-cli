package cmd

import "github.com/spf13/cobra"

var viewAllTodos = &cobra.Command{
	Use: "view-all",
	Short: "View all todos",
	Run: func(cmd *cobra.Command, args []string){
		// get request to API
		
	},
}

func init() {
	rootCmd.AddCommand(viewAllTodos)
}
