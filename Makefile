.PHONY: go-generate create-table test run-backend

go-generate:
	go generate ./...

create-table:
	sqlite3 data.sqlite < init.sql

test:
	go test ./...

run-backend:
	go run cmd/main.go
