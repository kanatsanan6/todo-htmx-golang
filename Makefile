setup:
	which air | go install github.com/air-verse/air@latest
	which templ | go install github.com/a-h/templ/cmd/templ@latest
run.dev:
	air
db.create:
	docker exec -it todo-htmx-go-postgres-1 psql -U postgres -c "CREATE DATABASE todo_htmx_go"
db.drop:
	docker exec -it todo-htmx-go-postgres-1 psql -U postgres -c "DROP DATABASE todo_htmx_go"
