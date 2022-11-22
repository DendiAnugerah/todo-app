package repository

import (
	"github.com/DendiAnugerah/Todo/model"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return TodoRepository{db}
}

func (t *TodoRepository) AddTodo(todo model.Todo) error {
	return t.db.Create(&todo).Error
}

func (t *TodoRepository) ReadTodo() ([]model.Todo, error) {
	todos := []model.Todo{}
	err := t.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *TodoRepository) UpdateTodoStatus(id int, status bool) error {
	todos := model.Todo{}
	err := t.db.Model(&todos).Where("id = ", id).Update("done", status).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *TodoRepository) DeleteTodo(id int) error {
	todos := model.Todo{}
	err := t.db.Where("id = ", id).Delete(&todos).Error
	if err != nil {
		return err
	}

	return nil
}
