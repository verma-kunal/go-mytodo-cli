package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	todoModel "github.com/verma-kunal/go-mytodo/api/model"
)

func DeleteTodo(c *gin.Context) {
	// get id from request
	todoIdStr := c.Param("id") // id as a string

	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	// read existing todos from the file
	jsonFile, err := os.Open("api/data/todos.json")
	if err != nil {
		log.Fatalf("error in opening JSON file")
	}
	// defer the closing of the file
	defer jsonFile.Close()

	// decode/unmarshal the JSON file into Go struct
	var todoList todoModel.Todos
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&todoList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read from file"})
		return
	}

	todoFound := false

	for index, todo := range todoList.Todos {
		if todoId == todo.Id {
			todoFound = true
			// Remove the todo item from the list
			todoList.Todos = append(todoList.Todos[:index], todoList.Todos[index+1:]...)
			break
		}
	}

	// if the todo with the specified ID was not found, return an error
	if !todoFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo item does not exist"})
		return
	}

	// write the list back to the JSON file
	// re-open the file in write mode with truncation to clear the old content
	jsonFile, err = os.OpenFile("api/data/todos.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("error in opening JSON file")
	}
	// defer the closing of the file
	defer jsonFile.Close()

	// write the updated todo list back to the JSON file
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", " ")
	err = encoder.Encode(todoList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not write to the file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo item deleted"})

}
