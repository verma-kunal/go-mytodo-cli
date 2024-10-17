package handlers

import (
	"encoding/json"
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

	// slice of todo items
	var todoList todoModel.Todos

	// decode/unmarshal the JSON file into Go struct
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&todoList)
	if err != nil {
		log.Fatalf("error in decoding JSON file")
	}

	// return todos in response
	c.JSON(200, todoList)
}
