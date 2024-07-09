package types

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          int64
	Title       string
	Description sql.NullString
	Completed   sql.NullBool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
