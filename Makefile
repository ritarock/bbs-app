PHONY: create.table create.mock test run.backend

MOCK_GEN = go run -mod=mod go.uber.org/mock/mockgen@latest

create.table:
	sqlite3 data.sqlite < init.sql

create.mock:
	$(MOCK_GEN) \
	-source=./domain/post.go \
	-destination=./testing/mock/post.go \
	-package=mock

	$(MOCK_GEN) \
	-source=./domain/comment.go \
	-destination=./testing/mock/comment.go \
	-package=mock

	$(MOCK_GEN) \
	-source=./domain/user.go \
	-destination=./testing/mock/user.go \
	-package=mock

test:
	go test ./...

run.backend:
	go run cmd/main.go

run.frontend:
	cd web; yarn dev
