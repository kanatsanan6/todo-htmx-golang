package server

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/kanatsanan6/todo-htmx-go/internal/repo"
	"github.com/kanatsanan6/todo-htmx-go/internal/views/task"
	"github.com/kanatsanan6/todo-htmx-go/types"
)

type TaskHandler struct {
	tr repo.TaskRepo
}

func NewTaskHandler(tr repo.TaskRepo) *TaskHandler {
	return &TaskHandler{tr: tr}
}

func (t *TaskHandler) Index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, _ := t.tr.GetAll()

		templ.Handler(task.Index(tasks)).ServeHTTP(w, r)
	}
}

func (t *TaskHandler) Create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tk := &types.Task{
			Title: r.FormValue("title"),
		}

		t.tr.Create(tk)

		templ.Handler(task.TaskCard(tk)).ServeHTTP(w, r)
	}
}

func (t *TaskHandler) Update() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)

		tk, _ := t.tr.GetByID(id)
		tk.Title = r.FormValue("title")

		t.tr.Update(tk)

		templ.Handler(task.TaskCard(tk)).ServeHTTP(w, r)
	}
}

func (t *TaskHandler) Destroy() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)

		t.tr.Destroy(id)
	}
}

func (t *TaskHandler) Toggle() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)

		tk, _ := t.tr.GetByID(id)
		tk.Completed.Bool = !tk.Completed.Bool

		t.tr.Update(tk)

		templ.Handler(task.TaskCard(tk)).ServeHTTP(w, r)
	}
}
