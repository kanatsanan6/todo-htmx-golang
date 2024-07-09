package repo

import "github.com/kanatsanan6/todo-htmx-go/types"

type TaskRepo interface {
	GetAll() []*types.Task
}
