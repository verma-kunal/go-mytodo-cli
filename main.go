package main

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/verma-kunal/go-mytodo/api/handlers"
)

func main() {
	router := gin.Default()

	// routes
	router.GET("/api/todos", handlers.GetTodos)
	router.GET("/api/todos/:id", handlers.GetTodoById)
	router.POST("/api/todos", handlers.AddTodo)
	router.PATCH("/api/todos/:id", handlers.UpdateTodo)
	router.DELETE("/api/todos/:id", handlers.DeleteTodo)

	router.Run("localhost:8080")
}
