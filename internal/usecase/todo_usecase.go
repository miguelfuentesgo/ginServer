package usecase

import (
	"server/internal/domain"
	"server/internal/repository"

	"github.com/google/uuid"
)

type TodoUseCase struct {
	repo repository.TodoRepository
}

// Cuando usar el * y cuando no?
func NewTodoUseCase(repo repository.TodoRepository) *TodoUseCase {
	return &TodoUseCase{
		repo: repo,
	}
}

func (usecase *TodoUseCase) GetAll() ([]domain.Todo, error) {
	todos, err := usecase.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (usecase *TodoUseCase) GetByID(id string) (domain.Todo, error) {
	todo, err := usecase.repo.GetByID(id)
	if err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (usecase *TodoUseCase) Update(todo domain.Todo) error {
	err := usecase.repo.Update(todo)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *TodoUseCase) Delete(id string) error {
	err := usecase.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *TodoUseCase) Create(description string) error {
	newTodo := domain.Todo{
		ID:          uuid.New().String(),
		Description: description,
		Completed:   false,
	}
	err := usecase.repo.Create(newTodo)
	if err != nil {
		return err
	}
	return nil
}
