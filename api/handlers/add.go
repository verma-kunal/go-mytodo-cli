package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	todoModel "github.com/verma-kunal/go-mytodo/api"
)

func AddTodo(c *gin.Context) {

	// instance of todo struct
	var newTodo todoModel.Todo

	// bind JSON body from request to the todo struct
	err := c.ShouldBindJSON(&newTodo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	// generate new id for the new todo item (based on the existing todo items)
	newTodo.Id = len(todoList.Todos) + 1

	// append the new todo item to the list of todos
	todoList.Todos = append(todoList.Todos, newTodo)

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

	c.JSON(http.StatusCreated, gin.H{"message": "New todo item added", "todoId": newTodo.Id})

}
