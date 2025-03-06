package infrastructure

import (
	"server/internal/domain"
	"server/internal/repository"
)

type TodoRepoMemory struct {
	todos []domain.Todo
}

func NewTodoRepository() repository.TodoRepository {
	return &TodoRepoMemory{
		todos: []domain.Todo{},
	}
}

func (memory *TodoRepoMemory) GetAll() ([]domain.Todo, error) {
	return memory.todos, nil
}

func (memory *TodoRepoMemory) GetByID(id string) (domain.Todo, error) {
	for _, todo := range memory.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return domain.Todo{}, nil
}

func (memory *TodoRepoMemory) Update(todo domain.Todo) error {
	for i, t := range memory.todos {
		if t.ID == todo.ID {
			memory.todos[i] = todo
			return nil
		}
	}
	return nil
}

func (memory *TodoRepoMemory) Delete(id string) error {
	for i, todo := range memory.todos {
		if todo.ID == id {
			memory.todos = append(memory.todos[:i], memory.todos[i+1:]...)
			return nil
		}
	}
	return nil
}

func (memory *TodoRepoMemory) Create(todo domain.Todo) error {
	memory.todos = append(memory.todos, todo)
	return nil
}
