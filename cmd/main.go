package main

import (
	"log"

	"github.com/kanatsanan6/todo-htmx-go/config"
	"github.com/kanatsanan6/todo-htmx-go/db"
	"github.com/kanatsanan6/todo-htmx-go/internal/server"
)

func main() {
	env, err := config.NewEnv()
	if err != nil {
		log.Fatalf("initialize env: %s", err.Error())
	}

	database, err := db.NewDatabase(env)
	if err != nil {
		log.Fatalf("connect db: %s", err.Error())
	}

	err = database.MigrateUp()
	if err != nil {
		log.Fatalf("migrate db: %s", err.Error())
	}

	s := server.NewServer()
	if err = s.Start(env.App.Port, database); err != nil {
		log.Fatalf("initialize server: %s", err.Error())
	}
}
