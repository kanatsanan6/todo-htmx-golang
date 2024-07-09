package repo

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kanatsanan6/todo-htmx-go/types"
)

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) TaskRepo {
	return &taskRepo{db: db}
}

func (t *taskRepo) GetAll() []*types.Task {
	return []*types.Task{
		{
			ID:          1,
			Title:       "TODO 1",
			Description: "Description 1",
			Completed:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Title:       "TODO 2",
			Description: "Description 2",
			Completed:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
}

func (t *taskRepo) Create() {
}
