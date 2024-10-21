package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteTodoId int

var deleteTodo = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing todo item data",
	Run: func(cmd *cobra.Command, args []string) {

		if deleteTodoId == 0 {
			log.Println("'id' flag is required")
			return
		}

		// convert int ID to string
		idToString := strconv.Itoa(deleteTodoId)

		url := "http://localhost:8080/api/todos/" + idToString

		// create new delete request
		request, _ := http.NewRequest(http.MethodDelete, url, nil)

		// create client to send request
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			log.Fatalf("Failed to make the POST request: %v", err)
			return
		}
		defer response.Body.Close()

		// Check if the request was successful
		if response.StatusCode != http.StatusOK {
			fmt.Printf("Request failed with status: %v\n", response.Status)
			return
		}

		fmt.Printf("Todo item with id: %d has been deleted", deleteTodoId)
	},
}

func init() {

	deleteTodo.Flags().IntVarP(&deleteTodoId, "id", "i", 0, "Id of the todo item (type: number)")
	rootCmd.AddCommand(deleteTodo)
}
