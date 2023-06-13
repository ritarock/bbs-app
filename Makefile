.PHONY: generate test

go-generate:
	go generate ./...

test:
	go test ./...
