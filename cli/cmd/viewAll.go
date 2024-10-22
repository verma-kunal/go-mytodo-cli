package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/olekukonko/tablewriter"

	todoModel "github.com/verma-kunal/go-mytodo/api/model"
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

		// parse JSON data
		var resp todoModel.Todos
		jsonErr := json.Unmarshal(body, &resp)
		if jsonErr != nil {
			log.Fatalf("failed to unmarshal JSON: %v", jsonErr)
		}

		// convert to [][]string format
		var result [][]string
		for _, todo := range resp.Todos {
			result = append(result, []string{
				fmt.Sprint(todo.Id),
				todo.Owner,
				*todo.Title,
				StatusToEmoji(*todo.Status),
			})
		}

		// format CLI response to table
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Id",
			"Owner",
			"Todo Item",
			"Status",
		})
		for _, vals := range result {
			table.Append(vals)
		}
		table.Render() // Send output

	},
}

func StatusToEmoji(status string) string {
	switch status {
	case "not started":
		return "⬜"
	case "in progress":
		return "🟨"
	case "completed":
		return "✅"
	default:
		return "⬜"
	}
}

func init() {
	rootCmd.AddCommand(viewAllTodos)
}
