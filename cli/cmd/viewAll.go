package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var viewAllTodos = &cobra.Command{
	Use:   "list",
	Short: "List all the todo items",
	Run: func(cmd *cobra.Command, args []string) {
		// get request to API
		response, err := http.Get("http://localhost:8080/api/todos")
		if err != nil {
			log.Fatalf("Failed to make the request: %v", err)
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			log.Fatalf("Request failed with status: %v", response.Status)
			return
		}

		// read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
			return
		}

		// pretty print JSON
		var prettyJSON interface{}
		jsonErr := json.Unmarshal(body, &prettyJSON)
		if jsonErr != nil {
			log.Fatalf("failed to unmarshal JSON: %v", jsonErr)
		}

		formattedJSON, _ := json.MarshalIndent(prettyJSON, "", "   ")
		fmt.Println(string(formattedJSON))

	},
}

func init() {
	rootCmd.AddCommand(viewAllTodos)
}
