.PHONY: setup run.dev db.create db.drop tailwind.compile templ.gen

setup:
	which air | go install github.com/air-verse/air@latest
	which migrate | brew install golang-migrate
	which templ | go install github.com/a-h/templ/cmd/templ@latest
	which tailwindcss | npm install -g tailwindcss
run.dev:
	air
tailwind.compile:
	tailwindcss --config tailwind.config.js -i static/tailwind.css -o static/css/styles.css
templ.gen:
	templ generate
db.create:
	docker exec -it todo-htmx-go-postgres-1 psql -U postgres -c "CREATE DATABASE todo_htmx_go"
db.drop:
	docker exec -it todo-htmx-go-postgres-1 psql -U postgres -c "DROP DATABASE todo_htmx_go"
db.migrate.create:
	migrate create -ext sql -dir db/migrations -seq $(name)
