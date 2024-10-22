package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var updatedData string
var updateTodoId int
var todoStatus string

var updateTodo = &cobra.Command{
	Use:   "update",
	Short: "Update an existing todo item",
	Run: func(cmd *cobra.Command, args []string) {

		if updateTodoId == 0 {
			log.Println("'id' flag is required")
			return
		}

		// prepare data to send based on flag
		data := make(map[string]string)
		if updatedData != "" { // If not empty, add to map
			data["title"] = updatedData
		}
		if todoStatus != "" { // If not empty, add to map
			data["status"] = todoStatus
		}

		// If no data to update, notify the user and exit
		if len(data) == 0 {
			log.Println("Either 'data' or 'status' flag must be provided")
			return
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		// convert int ID to string
		idToString := strconv.Itoa(updateTodoId)

		url := "http://localhost:8080/api/todos/" + idToString

		// create PATCH request
		request, _ := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json")

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

		fmt.Printf("Todo item with id: %d has been updated", updateTodoId)

	},
}

func init() {

	/*
		three flags:
		1. -i 2
		2. --item "send postcard"
		3. --status
	*/

	updateTodo.Flags().IntVarP(&updateTodoId, "id", "i", 0, "Id of the todo item (type: number)")
	updateTodo.Flags().StringVarP(&updatedData, "data", "d", "", "Updated todo item data (type: string)")
	updateTodo.Flags().StringVarP(&todoStatus, "status", "s", "",
		"Update status of todo item (type: string).\n"+
			"Supported values:\n"+
			"  - 'not started' (default)\n"+
			"  - 'in progress'\n"+
			"  - 'completed'")

	rootCmd.AddCommand(updateTodo)

}
