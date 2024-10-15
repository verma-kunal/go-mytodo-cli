package handlers

import (
	"encoding/json"
	"io"
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

	// read the JSON file and convert to bytes
	jsonToBytes, _ := io.ReadAll(jsonFile)

	// unmarshall
	var todoList todoModel.Todos
	json.Unmarshal(jsonToBytes, &todoList)

	for _, todo := range todoList.Todos {
		if todo.Id == todoId {
			// if matching todo found, return it
			c.JSON(http.StatusFound, todo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo item not found"})

}
