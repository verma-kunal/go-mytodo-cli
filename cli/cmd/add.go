package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// global vars
var user string
var todoItem string

var addTodo = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo item",
	Run: func(cmd *cobra.Command, args []string) {

		if user == "" {
			log.Println("'user' flag is required")
			return
		}
		if todoItem == "" {
			log.Println("'item' flag is required")
			return
		}

		// prepare data to send in request
		data := map[string]string{
			"owner": user,
			"title": todoItem,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		response, err := http.Post("http://localhost:8080/api/todos", "application/json", bytes.NewBuffer(jsonData))
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

		fmt.Println("Todo item added")

	},
}

func init() {

	/*
		two flags:
		1. --user "kunal"
		2. --item "send postcard"
	*/

	addTodo.Flags().StringVarP(&user, "user", "u", "", "Name of the owner of the task (type: string)")
	addTodo.Flags().StringVarP(&todoItem, "item", "", "", "Todo item to be added (type: string)")

	rootCmd.AddCommand(addTodo)

}
