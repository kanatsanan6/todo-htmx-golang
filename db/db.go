package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kanatsanan6/todo-htmx-go/config"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sqlx.DB
}

const POSTGRES = "postgres"

func NewDatabase(env *config.Env) (*Database, error) {
	db, err := sqlx.Connect(POSTGRES, fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		env.Database.User, env.Database.Pass, env.Database.Host,
		env.Database.Port, env.Database.Name,
	))

	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (database *Database) GetDB() *sqlx.DB {
	return database.db
}
