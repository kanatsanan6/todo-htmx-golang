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
	row, err := t.db.Query(`SELECT * FROM tasks ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var task types.Task
		if err := row.Scan(
			&task.ID,
			&task.Title,
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

func (t *taskRepo) GetByID(id int64) (*types.Task, error) {
	var task types.Task
	row := t.db.QueryRow(`SELECT * FROM tasks WHERE (id) = ($1) LIMIT 1`, id)
	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	return &task, err
}

func (t *taskRepo) Create(task *types.Task) error {
	return t.db.QueryRow(
		`INSERT INTO tasks (title) VALUES ($1) RETURNING *`,
		task.Title,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
}

func (t *taskRepo) Update(task *types.Task) error {
	query := `
		UPDATE tasks SET
			title = COALESCE($2, title),
			completed = COALESCE($3, completed)
		 WHERE id = $1 RETURNING *
	`
	return t.db.QueryRow(
		query,
		task.ID,
		task.Title,
		task.Completed,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
}

func (t *taskRepo) Destroy(id int64) error {
	row := t.db.QueryRow(`DELETE FROM tasks WHERE (id) = ($1)`, id)
	return row.Err()
}
