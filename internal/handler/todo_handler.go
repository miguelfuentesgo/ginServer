package handler

import (
	"net/http"
	"server/internal/domain"
	"server/internal/dto"
	"server/internal/usecase"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	usecase usecase.TodoUseCase
}

func NewTodoHandler(usecase usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		usecase: usecase,
	}
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.usecase.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoByID(c *gin.Context) {

	todo, err := h.usecase.GetByID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var createTodoRequest dto.CreateTodoRequest

	if err := c.ShouldBindJSON(&createTodoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.usecase.Create(createTodoRequest.Description)

	c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully"})
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var todo domain.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Update(todo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

func (h *TodoHandler) DeleteTodoByID(c *gin.Context) {
	err := h.usecase.Delete(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
