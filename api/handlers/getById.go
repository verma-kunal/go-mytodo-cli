package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	todoModel "github.com/verma-kunal/go-mytodo/api"
)

func GetTodoById(c *gin.Context) {

	// get id from request
	todoIdStr := c.Param("id") // id as a string

	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	// open the json file
	jsonFile, err := os.Open("api/data/todos.json")
	if err != nil {
		log.Fatalf("error in opening JSON file")
	}
	// defer the closing of the file
	defer jsonFile.Close()

	// decode/unmarshal JSON to the Go struct
	var todoList todoModel.Todos
	decoder := json.NewDecoder(jsonFile) // reads JSON from the file
	decoder.Decode(&todoList) // converts JSON to Go struct (without first loading the entire JSON content into memory)

	for _, todoItem := range todoList.Todos {
		if todoItem.Id == todoId {
			// if matching todo found, return it
			c.JSON(http.StatusFound, todoItem)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo item not found"})

}
