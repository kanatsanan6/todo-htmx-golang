package repo

import "github.com/kanatsanan6/todo-htmx-go/types"

type TaskRepo interface {
	GetAll() ([]*types.Task, error)
	GetByID(id int64) (*types.Task, error)
	Create(task *types.Task) error
	Update(task *types.Task) error
	Destroy(id int64) error
}
