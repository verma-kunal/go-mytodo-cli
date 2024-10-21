package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var singleTodoId int

var viewSingle = &cobra.Command{
	Use:   "view",
	Short: "View a single todo item",
	Run: func(cmd *cobra.Command, args []string) {

		// convert int ID to string
		idToString := strconv.Itoa(singleTodoId)

		url := "http://localhost:8080/api/todos/" + idToString

		response, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to make the request: %v", err)
			return
		}
		defer response.Body.Close()
		// fmt.Println("pass 1") // for debugging

		if response.StatusCode != http.StatusFound {
			log.Fatalf("Request failed with status: %v", response.Status)
			return
		}
		// fmt.Println("pass 2") // for debugging

		// read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
			return
		}
		// fmt.Println("pass 3") // for debugging

		// pretty print JSON
		var prettyJSON interface{}
		jsonErr := json.Unmarshal(body, &prettyJSON)
		if jsonErr != nil {
			log.Fatalf("failed to unmarshal JSON: %v", jsonErr)
		}

		// fmt.Println("pass 4") // for debugging

		formattedJSON, _ := json.MarshalIndent(prettyJSON, "", "   ")
		fmt.Println(string(formattedJSON))
	},
}

func init() {

	viewSingle.Flags().IntVarP(&singleTodoId, "id", "i", 0, "Id of the todo item (type: number)")
	rootCmd.AddCommand(viewSingle)
}
