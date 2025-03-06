package main

import (
	"server/internal/handler"
	"server/internal/infrastructure"
	"server/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := infrastructure.NewTodoRepository()
	todoUseCase := usecase.NewTodoUseCase(repo)
	todoHandler := handler.NewTodoHandler(*todoUseCase)

	r := gin.Default()

	r.GET("/todos", todoHandler.GetAllTodos)
	r.POST("/todos", todoHandler.CreateTodo)
	r.GET("/todos/:id", todoHandler.GetTodoByID)
	r.PUT("/todos", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodoByID)

	r.Run(":8080")
}
