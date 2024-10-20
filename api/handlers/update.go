package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	todoModel "github.com/verma-kunal/go-mytodo/api/model"
)

func UpdateTodo(c *gin.Context) {

	// get id from request
	todoIdStr := c.Param("id") // id as a string

	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	// instance of todo struct
	var newTodo todoModel.Todo

	// bind JSON body from request to the todo struct
	bindErr := c.ShouldBindJSON(&newTodo)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
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

	// variable to track whether todo item was found
	todoFound := false

	// find the target todo item
	for index, todo := range todoList.Todos {
		fmt.Println(todoId, todo.Id)
		// id from request matches any existing todo item
		if todoId == todo.Id {
			// update title from the original todo list struct
			todoList.Todos[index].Title = newTodo.Title
			todoFound = true
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

	c.JSON(http.StatusCreated, gin.H{"message": "todo item updated"})

}
