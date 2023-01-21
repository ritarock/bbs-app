.PHONY: buf.gen buf.lint run.dev

buf.lint:
	@go run github.com/bufbuild/buf/cmd/buf@latest lint .

buf.gen:
	@go run github.com/bufbuild/buf/cmd/buf@latest generate ./proto/v1

run.server:
	go run main.go
