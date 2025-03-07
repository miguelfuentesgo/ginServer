package infrastructure

import (
	"log"
	"server/internal/domain"
	"server/internal/repository"

	"gorm.io/gorm"
)

type TodoRepoDB struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &TodoRepoDB{
		db: db,
	}
}

func (todoDB *TodoRepoDB) GetAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	err := todoDB.db.Find(&todos).Error

	if err != nil {
		log.Fatal("Error getting all todos")
		return nil, err
	}
	return todos, nil
}

func (todoDB *TodoRepoDB) GetByID(id string) (domain.Todo, error) {
	var foundTodo domain.Todo
	err := todoDB.db.First(&foundTodo, id).Error

	if err != nil {
		log.Fatal("Error getting todo by id")
		return domain.Todo{}, err
	}
	return foundTodo, nil
}

func (todoDB *TodoRepoDB) Update(todo domain.Todo) error {
	err := todoDB.db.Save(&todo).Error

	if err != nil {
		log.Fatal("Error updating todo")
		return err
	}
	return nil
}

func (todoDB *TodoRepoDB) Delete(id string) error {
	err := todoDB.db.Delete(&domain.Todo{}, id).Error

	if err != nil {
		log.Fatal("Error deleting todo")
		return err
	}
	return nil
}

func (todoDB *TodoRepoDB) Create(todo domain.Todo) error {

	err := todoDB.db.Create(&todo).Error

	if err != nil {
		log.Fatal("Error creating todo")
		return err
	}
	return nil
}
