package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/megajon/gqlgen-todos2/graph/model"
)

type TodosRepo struct {
	DB *pg.DB
}

func (t *TodosRepo) GetTodos() ([]*model.Todo, error) {
	var todos []*model.Todo
	err := t.DB.Model(&todos).Select()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *TodosRepo) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	_, err := t.DB.Model(todo).Returning("*").Insert()
	return todo, err
}
