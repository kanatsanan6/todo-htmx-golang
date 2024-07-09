package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/kanatsanan6/todo-htmx-go/types"
)

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) TaskRepo {
	return &taskRepo{db: db}
}

func (t *taskRepo) GetAll() ([]*types.Task, error) {
	var tasks []*types.Task
	row, err := t.db.Query(`SELECT * FROM tasks`)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var task types.Task
		if err := row.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *taskRepo) Create(task *types.Task) error {
	return t.db.QueryRow(
		`INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING *`,
		task.Title,
		task.Description,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
}
