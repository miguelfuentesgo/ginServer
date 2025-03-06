package repository

import "server/internal/domain"

type TodoRepository interface {
	GetAll() ([]domain.Todo, error)
	GetByID(id string) (domain.Todo, error)
	Update(todo domain.Todo) error
	Delete(id string) error
	Create(todo domain.Todo) error
}
