package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/kanatsanan6/todo-htmx-go/internal/repo"
	"github.com/kanatsanan6/todo-htmx-go/internal/views/task"
)

type TaskHandler struct {
	tr repo.TaskRepo
}

func NewTaskHandler(tr repo.TaskRepo) *TaskHandler {
	return &TaskHandler{tr: tr}
}

func (t *TaskHandler) Index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks := t.tr.GetAll()

		templ.Handler(task.Index(tasks)).ServeHTTP(w, r)
	}
}

func (t *TaskHandler) Create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
