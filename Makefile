PHONY: wire.build create.table create.mock test

MOCK_GEN = go run -mod=mod go.uber.org/mock/mockgen@latest

wire.build:
	go run github.com/google/wire/cmd/wire@latest ./injector

create.table:
	sqlite3 data.sqlite < init.sql

create.mock:
	$(MOCK_GEN) \
	-source=./port/post.go \
	-destination=./testing/mock/post.go \
	-package=mock

	$(MOCK_GEN) \
	-source=./port/comment.go \
	-destination=./testing/mock/comment.go \
	-package=mock

	$(MOCK_GEN) \
	-source=./port/user.go \
	-destination=./testing/mock/user.go \
	-package=mock

test:
	go test ./...
