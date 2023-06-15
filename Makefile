.PHONY: generate create.table test run.backend run.frontend run

generate:
	go generate ./...

create.table:
	sqlite3 data.sqlite < init.sql

test:
	go test ./...

run.backend:
	go run cmd/main.go

run.frontend:
	cd web; yarn dev
