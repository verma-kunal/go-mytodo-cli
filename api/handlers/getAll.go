package handlers

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	todoModel "github.com/verma-kunal/go-mytodo/api"
)

func GetTodos(c *gin.Context) {

	// open the json file
	jsonFile, err := os.Open("api/data/todos.json")
	if err != nil {
		log.Fatalf("error in opening JSON file")
	}
	// defer the closing of the file
	defer jsonFile.Close()

	// read the JSON file and convert to bytes
	jsonToBytes, _ := io.ReadAll(jsonFile)

	// variable (of type 'slice of Todo') to hold the unmarshalled json data
	var todos todoModel.Todos
	json.Unmarshal(jsonToBytes, &todos)

	// return todos in response
	c.JSON(200, todos)
}
