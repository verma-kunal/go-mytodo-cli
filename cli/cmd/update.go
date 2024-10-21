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

var updatedItem string
var updateTodoId int

var updateTodo = &cobra.Command{
	Use:   "update",
	Short: "Update an existing todo item data",
	Run: func(cmd *cobra.Command, args []string) {

		if updateTodoId == 0 {
			log.Println("'id' flag is required")
			return
		}
		if updatedItem == "" {
			log.Println("'update' flag is required")
			return
		}

		// prepare data to send in request
		data := map[string]string{
			"title": updatedItem,
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
		if response.StatusCode != http.StatusCreated {
			fmt.Printf("Request failed with status: %v\n", response.Status)
			return
		}

		fmt.Printf("Todo item with id: %d has been updated", updateTodoId)

	},
}

func init() {

	/*
		two flags:
		1. -i 2
		2. --item "send postcard"
	*/

	updateTodo.Flags().IntVarP(&updateTodoId, "id", "i", 0, "Id of the todo item (type: number)")
	updateTodo.Flags().StringVarP(&updatedItem, "update", "", "", "Updated todo item title (type: string)")

	rootCmd.AddCommand(updateTodo)

}
