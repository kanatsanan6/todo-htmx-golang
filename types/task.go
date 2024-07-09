package types

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        int64
	Title     string
	Completed sql.NullBool
	CreatedAt time.Time
	UpdatedAt time.Time
}
