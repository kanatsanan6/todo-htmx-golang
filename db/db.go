package db

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func (database *Database) MigrateUp() error {
	driver, err := postgres.WithInstance(database.db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	_, b, _, _ := runtime.Caller(0)
	migrationDir := "file://" + filepath.Dir(b) + "/migrations"

	m, err := migrate.NewWithDatabaseInstance(migrationDir, POSTGRES, driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migration changed")
			return nil
		} else {
			return err
		}
	}
	log.Println("Migration has been executed")
	return nil
}
